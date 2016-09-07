package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ManagerNetworkProtocol100SSDProtocol manager network protocol 1 0 0 s s d protocol

swagger:model ManagerNetworkProtocol.1.0.0_SSDProtocol
*/
type ManagerNetworkProtocol100SSDProtocol struct {

	/* Indicates the scope for the IPv6 Notify messages for SSDP.
	 */
	NotifyIPV6Scope string `json:"NotifyIPv6Scope,omitempty"`

	/* Indicates how often the Multicast is done from this service for SSDP.

	Minimum: 0
	*/
	NotifyMulticastIntervalSeconds *float64 `json:"NotifyMulticastIntervalSeconds,omitempty"`

	/* Indicates the time to live hop count for SSDPs Notify messages.

	Minimum: 1
	*/
	NotifyTTL float64 `json:"NotifyTTL,omitempty"`

	/* Indicates the protocol port.

	Minimum: 0
	*/
	Port *float64 `json:"Port,omitempty"`

	/* Indicates if the protocol is enabled or disabled
	 */
	ProtocolEnabled bool `json:"ProtocolEnabled,omitempty"`
}

// Validate validates this manager network protocol 1 0 0 s s d protocol
func (m *ManagerNetworkProtocol100SSDProtocol) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotifyIPV6Scope(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNotifyMulticastIntervalSeconds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNotifyTTL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var managerNetworkProtocol100SSDProtocolTypeNotifyIPV6ScopePropEnum []interface{}

// prop value enum
func (m *ManagerNetworkProtocol100SSDProtocol) validateNotifyIPV6ScopeEnum(path, location string, value string) error {
	if managerNetworkProtocol100SSDProtocolTypeNotifyIPV6ScopePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Link","Site","Organization"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			managerNetworkProtocol100SSDProtocolTypeNotifyIPV6ScopePropEnum = append(managerNetworkProtocol100SSDProtocolTypeNotifyIPV6ScopePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, managerNetworkProtocol100SSDProtocolTypeNotifyIPV6ScopePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ManagerNetworkProtocol100SSDProtocol) validateNotifyIPV6Scope(formats strfmt.Registry) error {

	if swag.IsZero(m.NotifyIPV6Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateNotifyIPV6ScopeEnum("NotifyIPv6Scope", "body", m.NotifyIPV6Scope); err != nil {
		return err
	}

	return nil
}

func (m *ManagerNetworkProtocol100SSDProtocol) validateNotifyMulticastIntervalSeconds(formats strfmt.Registry) error {

	if swag.IsZero(m.NotifyMulticastIntervalSeconds) { // not required
		return nil
	}

	if err := validate.Minimum("NotifyMulticastIntervalSeconds", "body", float64(*m.NotifyMulticastIntervalSeconds), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ManagerNetworkProtocol100SSDProtocol) validateNotifyTTL(formats strfmt.Registry) error {

	if swag.IsZero(m.NotifyTTL) { // not required
		return nil
	}

	if err := validate.Minimum("NotifyTTL", "body", float64(m.NotifyTTL), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ManagerNetworkProtocol100SSDProtocol) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.Minimum("Port", "body", float64(*m.Port), 0, false); err != nil {
		return err
	}

	return nil
}
