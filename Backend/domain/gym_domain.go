package domain

type Gym struct {
	ID         string
	Name       string
	Location   Location
	Trainers   []Person
	Clients    []Person
	Activities []Activity
}

type Location struct {
	Country string
	City    string
	Street  string
	Number  int
	ZipCode string
}

type Activity struct {
	ID        string
	Name      string
	Duration  int    // in minutes
	Intensity string // low, medium, high
	TrainerID string
	Schedule  []Schedule
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
