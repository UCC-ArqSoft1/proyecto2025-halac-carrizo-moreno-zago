package domain

type Gym struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Location   Location   `json:"location"`
	Trainers   []Person   `json:"trainers"`
	Clients    []Person   `json:"clients"`
	Activities []Activity `json:"activities"`
}

type Location struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Number  int    `json:"number"`
	ZipCode string `json:"zip_code"`
}

type Activity struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Duration  int        `json:"duration"`
	Intensity string     `json:"intensity"`
	TrainerID string     `json:"trainer_id"`
	Schedule  []Schedule `json:"schedule"`
  }
  

  type Schedule struct {
	DayOfWeek string `json:"day_of_week"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
  }

type Person struct {
	ID          string
	Name        string
	Age         int
	DateOfBirth Date
	Role        string // "trainer", "client", etc.
	DNI         string
	Mail        string
	Phone       string
}

type Date struct {
	Day   int
	Month int
	Year  int
}


type Inscription struct {
	UserID     string `json:"user_id"`
	ActivityID string `json:"activity_id"`
}