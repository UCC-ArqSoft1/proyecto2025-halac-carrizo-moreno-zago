package services

import "backend/domain"

var inscriptions []domain.Inscription

func RegisterUserToActivity(userID string, activityID string) {
	// Evitar duplicados
	for _, insc := range inscriptions {
		if insc.UserID == userID && insc.ActivityID == activityID {
			return
		}
	}
	inscriptions = append(inscriptions, domain.Inscription{
		UserID:     userID,
		ActivityID: activityID,
	})
}

func GetUserInscriptions(userID string) []domain.Activity {
	var result []domain.Activity
	for _, insc := range inscriptions {
		if insc.UserID == userID {
			activity := GetActivityById(insc.ActivityID)
			if activity.ID != "" {
				result = append(result, activity)
			}
		}
	}
	return result
}
