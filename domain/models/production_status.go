package models

import "errors"

type Status int

const (
	Pending Status = iota
	InProgress
	Completed
	Failed
)

func (s Status) String() string {
	return [...]string{"Pending", "InProgress", "Completed", "Failed"}[s]
}

func StatusFromString(str string) (Status, error) {
	switch str {
	case "Pending":
		return Pending, nil
	case "InProgress":
		return InProgress, nil
	case "Completed":
		return Completed, nil
	case "Failed":
		return Failed, nil
	default:
		return 0, errors.New("invalid status")
	}
}

func StatusFromInt(status int) string {
	return [...]string{"Pending", "InProgress", "Completed", "Failed"}[status]
}
