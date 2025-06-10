package services

import "backend/domain"

var inscriptions []domain.Inscription

func RegisterUserToActivity(userID string, activityID string, dayOfWeek string) {
    for _, insc := range inscriptions {
        if insc.UserID == userID && insc.ActivityID == activityID && insc.DayOfWeek == dayOfWeek {
            return // Ya inscripto a ese día
        }
    }
    inscriptions = append(inscriptions, domain.Inscription{
        UserID:     userID,
        ActivityID: activityID,
        DayOfWeek:  dayOfWeek,
    })
}


func GetUserInscriptions(userID string) []domain.Activity {
	var result []domain.Activity
	for _, insc := range inscriptions {
		if insc.UserID == userID {
			activity, err := GetActivityById(insc.ActivityID) 
			if err == nil && activity != nil && activity.ID != "" {
				result = append(result, *activity)
			}
		}
	}
	return result
}
