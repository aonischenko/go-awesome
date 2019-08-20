package model

import "time"

type ApiError struct {
	Status    int       `json:"status"`
	Message   string    `json:"message"`
	Details   string    `json:"details"`
	Timestamp time.Time `json:"ts"`
}

func NewApiError(status int, message string, details string) ApiError {
	return ApiError{
		Status:    status,
		Message:   message,
		Details:   details,
		Timestamp: time.Now().UTC(),
	}
}
