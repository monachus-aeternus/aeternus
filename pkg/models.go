package pkg

import "time"

type HealthMacros struct {
	CarbsInGms   int
	ProteinInGms int
	FatsInGms    int
}
type HealthCoordinates struct {
	HealthMacros
	TotalCalories float32
	Density       float32
}
type HealthJournalEntryLine struct {
	CreationTime time.Time
	FoodItem     string
	FoodUnit     int
	FoodQty      float32
	Coordinates  HealthCoordinates
}
type HealthJournalEntry struct {
	Lines []HealthJournalEntryLine
	HealthCoordinates
}

type HealthJournal struct {
	entries []HealthJournalEntry
	HealthCoordinates
}
