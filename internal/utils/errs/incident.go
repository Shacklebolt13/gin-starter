package errs

import (
	"fmt"
	"net/http"
)

type IncidentType int

const (
	InternalError          IncidentType = http.StatusInternalServerError
	DatabaseError          IncidentType = http.StatusInternalServerError
	NotFoundError          IncidentType = http.StatusNotFound
	RequestValidationError IncidentType = http.StatusBadRequest
	AuthenticationError    IncidentType = http.StatusUnauthorized
	Redirect               IncidentType = http.StatusMovedPermanently
	Success                IncidentType = http.StatusOK
)

type Incident struct {
	Type    IncidentType `json:"type,omitempty"`
	Message string       `json:"message,omitempty"`
	Err     error        `json:"-"`
	Detail  any          `json:"detail,omitempty"`
}

func (e *Incident) Error() string {
	return fmt.Sprintf("Incident: %d, %s", e.Type, e.Message)
}

func (e *Incident) ToMap() map[string]any {
	return map[string]any{
		"type":    fmt.Sprintf("%d", e.Type),
		"message": e.Message,
		"detail":  e.Detail,
	}
}
