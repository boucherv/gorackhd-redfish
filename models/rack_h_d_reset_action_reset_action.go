package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*RackHDResetActionResetAction This is the base type for the reset action.

swagger:model RackHD.ResetAction_ResetAction
*/
type RackHDResetActionResetAction struct {

	/* reset type

	Required: true
	*/
	ResetType *string `json:"reset_type"`
}

// Validate validates this rack h d reset action reset action
func (m *RackHDResetActionResetAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResetType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rackHDResetActionResetActionTypeResetTypePropEnum []interface{}

// prop value enum
func (m *RackHDResetActionResetAction) validateResetTypeEnum(path, location string, value string) error {
	if rackHDResetActionResetActionTypeResetTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["On","ForceOff","GracefulShutdown","GracefulRestart","ForceRestart","Nmi","ForceOn","PushPowerButton"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			rackHDResetActionResetActionTypeResetTypePropEnum = append(rackHDResetActionResetActionTypeResetTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, rackHDResetActionResetActionTypeResetTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RackHDResetActionResetAction) validateResetType(formats strfmt.Registry) error {

	if err := validate.Required("reset_type", "body", m.ResetType); err != nil {
		return err
	}

	// value enum
	if err := m.validateResetTypeEnum("reset_type", "body", *m.ResetType); err != nil {
		return err
	}

	return nil
}
