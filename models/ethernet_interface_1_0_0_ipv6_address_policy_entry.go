package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*EthernetInterface100IPV6AddressPolicyEntry ethernet interface 1 0 0 ipv6 address policy entry

swagger:model EthernetInterface.1.0.0_IPv6AddressPolicyEntry
*/
type EthernetInterface100IPV6AddressPolicyEntry struct {

	/* The IPv6 Label (as defined in RFC 6724 section 2.1)

	Maximum: 100
	Minimum: 0
	*/
	Label *float64 `json:"Label,omitempty"`

	/* The IPv6 Precedence (as defined in RFC 6724 section 2.1

	Maximum: 100
	Minimum: 1
	*/
	Precedence float64 `json:"Precedence,omitempty"`

	/* The IPv6 Address Prefix (as defined in RFC 3484 section 2.1)
	 */
	Prefix string `json:"Prefix,omitempty"`
}

// Validate validates this ethernet interface 1 0 0 ipv6 address policy entry
func (m *EthernetInterface100IPV6AddressPolicyEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrecedence(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EthernetInterface100IPV6AddressPolicyEntry) validateLabel(formats strfmt.Registry) error {

	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.Minimum("Label", "body", float64(*m.Label), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("Label", "body", float64(*m.Label), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *EthernetInterface100IPV6AddressPolicyEntry) validatePrecedence(formats strfmt.Registry) error {

	if swag.IsZero(m.Precedence) { // not required
		return nil
	}

	if err := validate.Minimum("Precedence", "body", float64(m.Precedence), 1, false); err != nil {
		return err
	}

	if err := validate.Maximum("Precedence", "body", float64(m.Precedence), 100, false); err != nil {
		return err
	}

	return nil
}
