package models

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*LogEntryLogEntry log entry log entry

swagger:model LogEntry_LogEntry
*/
type LogEntryLogEntry struct {
	LogEntry100LogEntry
}

// Validate validates this log entry log entry
func (m *LogEntryLogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.LogEntry100LogEntry.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
