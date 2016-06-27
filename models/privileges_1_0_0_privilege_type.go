package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*Privileges100PrivilegeType privileges 1 0 0 privilege type

swagger:model Privileges.1.0.0_PrivilegeType
*/
type Privileges100PrivilegeType string

// for schema
var privileges100PrivilegeTypeEnum []interface{}

func (m Privileges100PrivilegeType) validatePrivileges100PrivilegeTypeEnum(path, location string, value Privileges100PrivilegeType) error {
	if privileges100PrivilegeTypeEnum == nil {
		var res []Privileges100PrivilegeType
		if err := json.Unmarshal([]byte(`["Login","ConfigureManager","ConfigureUsers","ConfigureSelf","ConfigureComponents"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			privileges100PrivilegeTypeEnum = append(privileges100PrivilegeTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, privileges100PrivilegeTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this privileges 1 0 0 privilege type
func (m Privileges100PrivilegeType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePrivileges100PrivilegeTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
