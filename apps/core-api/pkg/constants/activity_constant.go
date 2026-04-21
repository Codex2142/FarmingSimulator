package constants

// ActivityType merepresentasikan tipe aktivitas yang dicatat dalam siklus budidaya
type ActivityType string

const (
	Feeding     ActivityType = "feeding"
	Fertilizing ActivityType = "fertilizing"
	Cleaning    ActivityType = "cleaning"
	Note        ActivityType = "note"
)
