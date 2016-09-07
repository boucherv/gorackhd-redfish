package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*RackHDBootImageNetworkDevice rack h d boot image network device

swagger:model RackHD.BootImage_NetworkDevice
*/
type RackHDBootImageNetworkDevice struct {

	/* device
	 */
	Device string `json:"device,omitempty"`

	/* ipv4
	 */
	IPV4 *RackHDBootImageNetworkAddress `json:"ipv4,omitempty"`

	/* ipv6
	 */
	IPV6 *RackHDBootImageNetworkAddress `json:"ipv6,omitempty"`
}

// Validate validates this rack h d boot image network device
func (m *RackHDBootImageNetworkDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPV4(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPV6(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackHDBootImageNetworkDevice) validateIPV4(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV4) { // not required
		return nil
	}

	if m.IPV4 != nil {

		if err := m.IPV4.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *RackHDBootImageNetworkDevice) validateIPV6(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6) { // not required
		return nil
	}

	if m.IPV6 != nil {

		if err := m.IPV6.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
