package services

import "backend/domain"

func GetActivityById(id string) domain.Activity {
	activity := domain.Activity{
		ID:          id,
		Title:       "Spinning Avanzado",
		Category:    "Spinning",
		Day:         "Martes",
		Time:        "18:00",
		Duration:    60,
		Instructor:  "Juan Pérez",
		Capacity:    20,
		Description: "Clase de spinning de alta intensidad",
		ImageURL:    "https://example.com/spinning.jpg",
	}
	return activity
}

func GetUserById(id string) domain.Person {
	user := domain.Person{
		ID:          id,
		Name:        "Lucía Gómez",
		Age:         30,
		DateOfBirth: domain.Date{Day: 15, Month: 6, Year: 1993},
		Role:        "socio",
		DNI:         "40123456",
		Mail:        "lucia@example.com",
		Phone:       "1123456789",
	}
	return user
}

func GetGym() domain.Gym {
	location := domain.Location{
		Country: "Argentina",
		City:    "Buenos Aires",
		Street:  "Calle Falsa",
		Number:  123,
		ZipCode: "1000",
	}

	gym := domain.Gym{
		ID:       "1",
		Name:     "Gimnasio Energía",
		Location: location,
		Trainers: []domain.Person{GetUserById("trainer1")},
		Clients:  []domain.Person{GetUserById("client1")},
		Activities: []domain.Activity{
			GetActivityById("a1"),
		},
	}
	return gym
}

func GetSchedule() domain.Schedule {
	return domain.Schedule{
		DayOfWeek: "Lunes",
		StartTime: "09:00",
		EndTime:   "10:00",
	}
}

func GetMembershipById(id string) domain.Membership {
	return domain.Membership{
		ID:             id,
		Name:           "Plan Mensual",
		Price:          15000.00,
		DurationInDays: 30,
		AccessLevel:    "Total",
	}
}

func GetReservationById(id string) domain.Reservation {
	return domain.Reservation{
		ID:         id,
		ClientID:   "client1",
		ActivityID: "a1",
		Date:       domain.Date{Day: 25, Month: 4, Year: 2025},
		Status:     "confirmed",
	}
}

func GetAttendanceById(id string) domain.Attendance {
	return domain.Attendance{
		ID:       id,
		PersonID: "client1",
		Date:     domain.Date{Day: 24, Month: 4, Year: 2025},
		CheckIn:  "10:00",
		CheckOut: "11:15",
	}
}

func GetPaymentById(id string) domain.Payment {
	return domain.Payment{
		ID:           id,
		ClientID:     "client1",
		MembershipID: "m1",
		Amount:       15000.00,
		Date:         domain.Date{Day: 1, Month: 4, Year: 2025},
		Method:       "credit card",
	}
}

func GetWorkoutRoutineById(id string) domain.WorkoutRoutine {
	exercises := []domain.Exercise{
		{Name: "Sentadillas", Reps: 12, Sets: 4, RestTime: 60, MuscleGroup: "Piernas"},
		{Name: "Press de banca", Reps: 10, Sets: 4, RestTime: 90, MuscleGroup: "Pecho"},
	}

	return domain.WorkoutRoutine{
		ID:        id,
		ClientID:  "client1",
		TrainerID: "trainer1",
		Name:      "Rutina Fuerza Básica",
		Exercises: exercises,
	}
}

func GetEquipmentById(id string) domain.Equipment {
	return domain.Equipment{
		ID:       id,
		Name:     "Bicicleta Estática",
		Quantity: 10,
		Status:   "available",
	}
}
