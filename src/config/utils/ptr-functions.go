package utils

import "time"

func PtrTime(t time.Time) *time.Time {
	return &t
}

func PtrString(s string) *string {
	return &s
}

func PtrInt(i int) *int {
	return &i
}

func PtrBool(b bool) *bool {
	return &b
}

func PtrInt64(i int64) *int64 {
	return &i
}

func PtrFloat64(f float64) *float64 {
	return &f
}

func PtrInt32(i int32) *int32 {
	return &i
}

func PtrInt8(i int8) *int8 {
	return &i
}

func PtrToSlice[T any](input []*T) []T {
	output := make([]T, len(input))
	for i, ptr := range input {
		if ptr != nil {
			output[i] = *ptr
		}
	}
	return output
}

func PtrToObject[T any](input *T) T {
	if input == nil {
		return *new(T)
	}
	return *input
}

func PtrToString(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}
