package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*VirtualMedia100MediaType virtual media 1 0 0 media type

swagger:model VirtualMedia.1.0.0_MediaType
*/
type VirtualMedia100MediaType string

// for schema
var virtualMedia100MediaTypeEnum []interface{}

func (m VirtualMedia100MediaType) validateVirtualMedia100MediaTypeEnum(path, location string, value VirtualMedia100MediaType) error {
	if virtualMedia100MediaTypeEnum == nil {
		var res []VirtualMedia100MediaType
		if err := json.Unmarshal([]byte(`["CD","Floppy","USBStick","DVD"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			virtualMedia100MediaTypeEnum = append(virtualMedia100MediaTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, virtualMedia100MediaTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this virtual media 1 0 0 media type
func (m VirtualMedia100MediaType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVirtualMedia100MediaTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
