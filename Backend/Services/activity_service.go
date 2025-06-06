package services

import (
    "backend/clients"
    "backend/dao"
    "backend/domain"
)

// GetActivities consulta la base y retorna TODAS las actividades
func GetActivities() []domain.Activity {
    var dbActivities []dao.Activity
    clients.DB.Preload("Schedules").Find(&dbActivities)

    var activities []domain.Activity
    for _, a := range dbActivities {
        var schedules []domain.Schedule
        for _, s := range a.Schedules {
            schedules = append(schedules, domain.Schedule{
                DayOfWeek: s.DayOfWeek,
                StartTime: s.StartTime,
                EndTime:   s.EndTime,
            })
        }
        activities = append(activities, domain.Activity{
            ID:        a.ID,
            Name:      a.Name,
            Duration:  a.Duration,
            Intensity: a.Intensity,
            TrainerID: a.TrainerID,
            Schedule:  schedules,
        })
    }
    return activities
}

// GetActivityById busca una actividad por su ID en la base
func GetActivityById(id string) (*domain.Activity, error) {
    var dbActivity dao.Activity
    result := clients.DB.Preload("Schedules").First(&dbActivity, "id = ?", id)
    if result.Error != nil {
        return nil, result.Error
    }
    var schedules []domain.Schedule
    for _, s := range dbActivity.Schedules {
        schedules = append(schedules, domain.Schedule{
            DayOfWeek: s.DayOfWeek,
            StartTime: s.StartTime,
            EndTime:   s.EndTime,
        })
    }
    act := &domain.Activity{
        ID:        dbActivity.ID,
        Name:      dbActivity.Name,
        Duration:  dbActivity.Duration,
        Intensity: dbActivity.Intensity,
        TrainerID: dbActivity.TrainerID,
        Schedule:  schedules,
    }
    return act, nil
}

// CreateActivity agrega una nueva actividad a la base
func CreateActivity(a domain.Activity) error {
    dbActivity := dao.Activity{
        ID:        a.ID,
        Name:      a.Name,
        Duration:  a.Duration,
        Intensity: a.Intensity,
        TrainerID: a.TrainerID,
    }

    // Primero la actividad, después los schedules
    err := clients.DB.Create(&dbActivity).Error
    if err != nil {
        return err
    }

    for _, sch := range a.Schedule {
        dbSchedule := dao.Schedule{
            ActivityID: a.ID,
            DayOfWeek:  sch.DayOfWeek,
            StartTime:  sch.StartTime,
            EndTime:    sch.EndTime,
        }
        clients.DB.Create(&dbSchedule)
    }
    return nil
}

// UpdateActivity actualiza una actividad existente por ID
func UpdateActivity(id string, updated domain.Activity) error {
    var dbActivity dao.Activity
    result := clients.DB.First(&dbActivity, "id = ?", id)
    if result.Error != nil {
        return result.Error
    }

    dbActivity.Name = updated.Name
    dbActivity.Duration = updated.Duration
    dbActivity.Intensity = updated.Intensity
    dbActivity.TrainerID = updated.TrainerID

    if err := clients.DB.Save(&dbActivity).Error; err != nil {
        return err
    }

    // Si querés también actualizar los schedules, primero los eliminás y después los volvés a crear
    clients.DB.Where("activity_id = ?", id).Delete(&dao.Schedule{})
    for _, sch := range updated.Schedule {
        dbSchedule := dao.Schedule{
            ActivityID: id,
            DayOfWeek:  sch.DayOfWeek,
            StartTime:  sch.StartTime,
            EndTime:    sch.EndTime,
        }
        clients.DB.Create(&dbSchedule)
    }

    return nil
}

// DeleteActivity elimina una actividad por ID de la base
func DeleteActivity(id string) error {
    // Eliminá los schedules primero (por FK)
    clients.DB.Where("activity_id = ?", id).Delete(&dao.Schedule{})
    // Después la actividad
    return clients.DB.Delete(&dao.Activity{}, "id = ?", id).Error
}
