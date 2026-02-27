package services

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"
	"errors"

	"gorm.io/gorm"
)

var ErrNoCapacity = errors.New("no hay cupos disponibles")

// RegisterUserToActivity registra una inscripción persistida en la base de datos
func RegisterUserToActivity(userID int, activityID string, dayOfWeek string) error {
	// Evitar duplicados: mismo usuario, actividad y día
	var existing dao.Inscription
	err := clients.DB.Where("user_id = ? AND activity_id = ? AND day_of_week = ?", userID, activityID, dayOfWeek).
		First(&existing).Error
	if err == nil {
		// Ya existe la inscripción → no hacemos nada
		return nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Buscar el horario correspondiente para conocer los cupos
	var schedule dao.Schedule
	if err := clients.DB.Where("activity_id = ? AND day_of_week = ?", activityID, dayOfWeek).
		First(&schedule).Error; err != nil {
		return err
	}

	// Validar cupos si hay capacidad configurada
	if schedule.Capacity > 0 {
		var count int64
		clients.DB.Model(&dao.Inscription{}).
			Where("activity_id = ? AND day_of_week = ?", activityID, dayOfWeek).
			Count(&count)

		if int(count) >= schedule.Capacity {
			return ErrNoCapacity
		}
	}

	insc := dao.Inscription{
		UserID:     userID,
		ActivityID: activityID,
		DayOfWeek:  dayOfWeek,
	}
	return clients.DB.Create(&insc).Error
}

// GetUserInscriptionsDetailed devuelve las actividades inscriptas de un usuario con su horario puntual
func GetUserInscriptionsDetailed(userID int) []domain.UserActivityResponse {
	var inscriptions []dao.Inscription
	clients.DB.Where("user_id = ?", userID).Find(&inscriptions)

	var result []domain.UserActivityResponse
	for _, insc := range inscriptions {
		activity, err := GetActivityById(insc.ActivityID)
		if err != nil || activity == nil || activity.ID == "" {
			continue
		}
		for _, s := range activity.Schedule {
			if s.DayOfWeek == insc.DayOfWeek {
				result = append(result, domain.UserActivityResponse{
					ActivityID: activity.ID,
					Name:       activity.Name,
					Duration:   activity.Duration,
					Intensity:  activity.Intensity,
					TrainerID:  activity.TrainerID,
					DayOfWeek:  s.DayOfWeek,
					StartTime:  s.StartTime,
					EndTime:    s.EndTime,
				})
			}
		}
	}
	return result
}
