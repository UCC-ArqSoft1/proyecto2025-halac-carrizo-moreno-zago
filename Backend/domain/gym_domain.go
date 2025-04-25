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
	Duration  int        // in minutes
	Intensity string     // low, medium, high
	TrainerID string     `json:"trainer_id"`
	Schedule  []Schedule `json:"schedule"`
}

type Schedule struct {
	DayOfWeek string // e.g. "Monday"
	StartTime string // "15:00"
	EndTime   string // "16:00"
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
