package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*VLanNetworkInterface100VLAN v lan network interface 1 0 0 v l a n

swagger:model VLanNetworkInterface.1.0.0_VLAN
*/
type VLanNetworkInterface100VLAN struct {

	/* This indicates if this VLAN is enabled.
	 */
	VLANEnable bool `json:"VLANEnable,omitempty"`

	/* This indicates the VLAN identifier for this VLAN.

	Maximum: 4095
	Minimum: 0
	*/
	VLANID *float64 `json:"VLANId,omitempty"`
}

// Validate validates this v lan network interface 1 0 0 v l a n
func (m *VLanNetworkInterface100VLAN) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVLANID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VLanNetworkInterface100VLAN) validateVLANID(formats strfmt.Registry) error {

	if swag.IsZero(m.VLANID) { // not required
		return nil
	}

	if err := validate.Minimum("VLANId", "body", float64(*m.VLANID), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("VLANId", "body", float64(*m.VLANID), 4095, false); err != nil {
		return err
	}

	return nil
}
