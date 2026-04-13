package utils

import "time"

// Helper function untuk pointer
func Ptr[T any](v T) *T {
	return &v
}

// Helper function untuk pointer time
func TimePtr(t time.Time) *time.Time {
	return &t
}

// Helper functions untuk pointer
func PtrFloat(f float64) *float64 {
	return &f
}

func PtrStr(s string) *string {
	return &s
}
