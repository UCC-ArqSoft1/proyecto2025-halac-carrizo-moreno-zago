package services

import "backend/domain"

// GetGym - Retorna la información del gimnasio con sus actividades, entrenadores y clientes
func GetGym() domain.Gym {
	location := domain.Location{
		Country: "Argentina",
		City:    "Buenos Aires",
		Street:  "Calle Falsa",
		Number:  123,
		ZipCode: "1000",
	}

	trainers := []domain.Person{
		{
			ID:          "trainer1",
			Name:        "Juan Pérez",
			Age:         35,
			DateOfBirth: domain.Date{Day: 15, Month: 6, Year: 1988},
			Role:        "trainer",
			DNI:         "12345678",
			Mail:        "juan@example.com",
			Phone:       "1122334455",
		},
	}

	clients := []domain.Person{
		{
			ID:          "client1",
			Name:        "Lucía Gómez",
			Age:         30,
			DateOfBirth: domain.Date{Day: 15, Month: 6, Year: 1993},
			Role:        "client",
			DNI:         "40123456",
			Mail:        "lucia@example.com",
			Phone:       "1123456789",
		},
	}

	activities := []domain.Activity{
		{
			ID:        "a1",
			Name:      "Spinning Avanzado",
			Duration:  60,
			Intensity: "high",
			TrainerID: "trainer1",
			Schedule: []domain.Schedule{
				{
					DayOfWeek: "Monday",
					StartTime: "18:00",
					EndTime:   "19:00",
				},
				{
					DayOfWeek: "Wednesday",
					StartTime: "18:00",
					EndTime:   "19:00",
				},
			},
		},
	}

	gym := domain.Gym{
		ID:         "1",
		Name:       "Gimnasio Energía",
		Location:   location,
		Trainers:   trainers,
		Clients:    clients,
		Activities: activities,
	}

	return gym
}

// GetActivityById - Obtiene una actividad por su ID
func GetActivityById(id string) domain.Activity {
	activity := domain.Activity{
		ID:        id,
		Name:      "Spinning Avanzado",
		Duration:  60,
		Intensity: "high",
		TrainerID: "trainer1",
		Schedule: []domain.Schedule{
			{
				DayOfWeek: "Monday",
				StartTime: "18:00",
				EndTime:   "19:00",
			},
			{
				DayOfWeek: "Wednesday",
				StartTime: "18:00",
				EndTime:   "19:00",
			},
		},
	}
	return activity
}

// GetPersonById - Obtiene la información de una persona (cliente o entrenador) por su ID
func GetPersonById(id string) domain.Person {
	// Este sería un ejemplo simple, pero en una implementación real buscaríamos en una base de datos
	person := domain.Person{
		ID:          id,
		Name:        "Lucía Gómez",
		Age:         30,
		DateOfBirth: domain.Date{Day: 15, Month: 6, Year: 1993},
		Role:        "client",
		DNI:         "40123456",
		Mail:        "lucia@example.com",
		Phone:       "1123456789",
	}
	return person
}

// GetTrainerById - Obtiene un entrenador por su ID
func GetTrainerById(id string) domain.Person {
	trainer := domain.Person{
		ID:          id,
		Name:        "Juan Pérez",
		Age:         35,
		DateOfBirth: domain.Date{Day: 15, Month: 6, Year: 1988},
		Role:        "trainer",
		DNI:         "12345678",
		Mail:        "juan@example.com",
		Phone:       "1122334455",
	}
	return trainer
}

// GetClientById - Obtiene un cliente por su ID
func GetClientById(id string) domain.Person {
	client := domain.Person{
		ID:          id,
		Name:        "Lucía Gómez",
		Age:         30,
		DateOfBirth: domain.Date{Day: 15, Month: 6, Year: 1993},
		Role:        "client",
		DNI:         "40123456",
		Mail:        "lucia@example.com",
		Phone:       "1123456789",
	}
	return client
}

// GetScheduleForActivity - Obtiene el horario para una actividad por su ID
func GetScheduleForActivity(id string) []domain.Schedule {
	// Simulación de datos; en una base de datos real, se buscarían los horarios de la actividad
	schedule := []domain.Schedule{
		{
			DayOfWeek: "Monday",
			StartTime: "18:00",
			EndTime:   "19:00",
		},
		{
			DayOfWeek: "Wednesday",
			StartTime: "18:00",
			EndTime:   "19:00",
		},
	}
	return schedule
}

// GetActivitiesForTrainer - Obtiene las actividades de un entrenador por su ID
func GetActivitiesForTrainer(trainerId string) []domain.Activity {
	// Simulación de datos; se deben recuperar las actividades asociadas al entrenador
	activities := []domain.Activity{
		{
			ID:        "a1",
			Name:      "Spinning Avanzado",
			Duration:  60,
			Intensity: "high",
			TrainerID: trainerId,
			Schedule: []domain.Schedule{
				{
					DayOfWeek: "Monday",
					StartTime: "18:00",
					EndTime:   "19:00",
				},
				{
					DayOfWeek: "Wednesday",
					StartTime: "18:00",
					EndTime:   "19:00",
				},
			},
		},
	}
	return activities
}
