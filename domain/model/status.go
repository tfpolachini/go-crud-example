package model

import "github.com/pkg/errors"

var (
	ENABLED  = Status{"ENABLED"}
	DISABLED = Status{"DISABLED"}
)

var statusValues = []Status{
	ENABLED,
	DISABLED,
}

type Status struct {
	name string
}

func (s Status) String() string {
	return s.name
}

func NewStatusFromString(statusStr string) (Status, error) {
	for _, status := range statusValues {
		if status.String() == statusStr {
			return status, nil
		}
	}
	return Status{}, errors.Errorf("unknown '%s' status", statusStr)
}
