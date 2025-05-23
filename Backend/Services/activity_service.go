package services

import "backend/domain"

var activities []domain.Activity = []domain.Activity{
	{
		ID:        "a1",
		Name:      "Spinning Avanzado",
		Duration:  60,
		Intensity: "high",
		TrainerID: "trainer1",
		Schedule: []domain.Schedule{
			{DayOfWeek: "Monday", StartTime: "18:00", EndTime: "19:00"},
			{DayOfWeek: "Wednesday", StartTime: "18:00", EndTime: "19:00"},
		},
	},
}

// GetActivities retorna todas las actividades creadas
func GetActivities() []domain.Activity {
	return activities
}

// GetActivityById busca una actividad por su ID
func GetActivityById(id string) domain.Activity {
	for _, a := range activities {
		if a.ID == id {
			return a
		}
	}
	return domain.Activity{}
}

// CreateActivity agrega una nueva actividad
func CreateActivity(a domain.Activity) {
	activities = append(activities, a)
}

// UpdateActivity actualiza una actividad existente por ID
func UpdateActivity(id string, updated domain.Activity) {
	for i, a := range activities {
		if a.ID == id {
			activities[i] = updated
			return
		}
	}
}

// DeleteActivity elimina una actividad por ID
func DeleteActivity(id string) {
	for i, a := range activities {
		if a.ID == id {
			activities = append(activities[:i], activities[i+1:]...)
			return
		}
	}
}
