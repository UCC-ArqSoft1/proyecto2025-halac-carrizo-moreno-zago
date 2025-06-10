package services

import (
    "backend/clients"
    "backend/dao"
    "fmt"
)

// SeedActivities recorre un slice de actividades en memoria y las inserta en la BD.
func SeedActivities() {
    activities := []dao.Activity{
        {
            ID:        "a1",
            Name:      "Spinning Avanzado",
            Duration:  60,
            Intensity: "high",
            TrainerID: "trainer1",
            Schedules: []dao.Schedule{
                {DayOfWeek: "Monday", StartTime: "18:00", EndTime: "19:00"},
                {DayOfWeek: "Wednesday", StartTime: "18:00", EndTime: "19:00"},
            },
        },
        {
            ID:        "a2",
            Name:      "Yoga Básico",
            Duration:  45,
            Intensity: "low",
            TrainerID: "trainer2",
            Schedules: []dao.Schedule{
                {DayOfWeek: "Tuesday", StartTime: "10:00", EndTime: "10:45"},
                {DayOfWeek: "Thursday", StartTime: "10:00", EndTime: "10:45"},
            },
        },
        // Agregar mas actividades aca
    }

    for _, a := range activities {
        var existing dao.Activity
        
        err := clients.DB.Preload("Schedules").First(&existing, "id = ?", a.ID).Error
        if err == nil {
            fmt.Printf("⚠️ Actividad %s ya existe. Saltando...\n", a.ID)
            continue
        }

        
        if err := clients.DB.Create(&a).Error; err != nil {
            fmt.Printf("⚠️ Error creando actividad %s: %v\n", a.ID, err)
        } else {
            fmt.Printf("✅ Actividad creada: %s (%s)\n", a.Name, a.ID)
        }
    }
}
