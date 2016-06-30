package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*EthernetInterface100EthernetInterface This schema defines a simple ethernet NIC resource.

swagger:model EthernetInterface.1.0.0_EthernetInterface
*/
type EthernetInterface100EthernetInterface struct {

	/* at odata context

	Read Only: true
	*/
	AtOdataContext strfmt.URI `json:"@odata.context,omitempty"`

	/* at odata id

	Read Only: true
	*/
	AtOdataID strfmt.URI `json:"@odata.id,omitempty"`

	/* at odata type

	Read Only: true
	*/
	AtOdataType string `json:"@odata.type,omitempty"`

	/* This indicates if the speed and duplex are automatically negotiated and configured on this interface.
	 */
	AutoNeg bool `json:"AutoNeg,omitempty"`

	/* Provides a description of this resource and is used for commonality  in the schema definitions.

	Read Only: true
	*/
	Description string `json:"Description,omitempty"`

	/* This is the complete, fully qualified domain name obtained by DNS for this interface.
	 */
	FQDN string `json:"FQDN,omitempty"`

	/* This indicates if the interface is in Full Duplex mode or not.
	 */
	FullDuplex bool `json:"FullDuplex,omitempty"`

	/* The DNS Host Name, without any domain information
	 */
	HostName string `json:"HostName,omitempty"`

	/* The IPv4 addresses assigned to this interface

	Read Only: true
	*/
	IPV4Addresses []*IPAddresses100IPV4Address `json:"IPv4Addresses,omitempty"`

	/* An array representing the RFC3484 Address Selection Policy Table.
	 */
	IPV6AddressPolicyTable []*EthernetInterface100IPV6AddressPolicyEntry `json:"IPv6AddressPolicyTable,omitempty"`

	/* This array of objects enumerates all of the currently assigned IPv6 addresses on this interface.

	Read Only: true
	*/
	IPV6Addresses []*IPAddresses100IPV6Address `json:"IPv6Addresses,omitempty"`

	/* This is the IPv6 default gateway address that is currently in use on this interface.

	Read Only: true
	*/
	IPV6DefaultGateway string `json:"IPv6DefaultGateway,omitempty"`

	/* This array of objects represents all of the IPv6 static addresses to be assigned on this interface.

	Read Only: true
	*/
	IPV6StaticAddresses []*IPAddresses100IPV6StaticAddress `json:"IPv6StaticAddresses,omitempty"`

	/* Uniquely identifies the resource within the collection of like resources.

	Read Only: true
	*/
	ID string `json:"Id,omitempty"`

	/* This indicates whether this interface is enabled.
	 */
	InterfaceEnabled bool `json:"InterfaceEnabled,omitempty"`

	/* This is the currently configured MAC address of the (logical port) interface.

	Pattern: ^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$
	*/
	MACAddress string `json:"MACAddress,omitempty"`

	/* This is the currently configured Maximum Transmission Unit (MTU) in bytes on this interface.
	 */
	MTUSize float64 `json:"MTUSize,omitempty"`

	/* This indicates the maximum number of Static IPv6 addresses that can be configured on this interface.

	Read Only: true
	*/
	MaxIPV6StaticAddresses float64 `json:"MaxIPv6StaticAddresses,omitempty"`

	/* The name of the resource or array element.

	Read Only: true
	*/
	Name string `json:"Name,omitempty"`

	/* This represents DNS name servers that are currently in use on this interface.

	Read Only: true
	*/
	NameServers []string `json:"NameServers,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* This is the permanent MAC address assigned to this interface (port)

	Read Only: true
	Pattern: ^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$
	*/
	PermanentMACAddress string `json:"PermanentMACAddress,omitempty"`

	/* This is the current speed in Mbps of this interface.
	 */
	SpeedMbps float64 `json:"SpeedMbps,omitempty"`

	/* status
	 */
	Status *ResourceStatus `json:"Status,omitempty"`

	/* The UEFI device path for this interface

	Read Only: true
	*/
	UefiDevicePath string `json:"UefiDevicePath,omitempty"`

	/* If this Network Interface supports more than one VLAN, this property will not be present and the client should look for VLANs collection in the link section of this resource.
	 */
	VLAN *VLanNetworkInterface100VLAN `json:"VLAN,omitempty"`

	/* This is a reference to a collection of VLANs and is only used if the interface supports more than one VLANs.

	Read Only: true
	*/
	VLANs *VLanNetworkInterfaceCollectionVLanNetworkInterfaceCollection `json:"VLANs,omitempty"`
}

// Validate validates this ethernet interface 1 0 0 ethernet interface
func (m *EthernetInterface100EthernetInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPV4Addresses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPV6AddressPolicyTable(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPV6Addresses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPV6StaticAddresses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMACAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNameServers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePermanentMACAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EthernetInterface100EthernetInterface) validateIPV4Addresses(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV4Addresses) { // not required
		return nil
	}

	return nil
}

func (m *EthernetInterface100EthernetInterface) validateIPV6AddressPolicyTable(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6AddressPolicyTable) { // not required
		return nil
	}

	return nil
}

func (m *EthernetInterface100EthernetInterface) validateIPV6Addresses(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6Addresses) { // not required
		return nil
	}

	return nil
}

func (m *EthernetInterface100EthernetInterface) validateIPV6StaticAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6StaticAddresses) { // not required
		return nil
	}

	return nil
}

func (m *EthernetInterface100EthernetInterface) validateMACAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.MACAddress) { // not required
		return nil
	}

	if err := validate.Pattern("MACAddress", "body", string(m.MACAddress), `^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`); err != nil {
		return err
	}

	return nil
}

func (m *EthernetInterface100EthernetInterface) validateNameServers(formats strfmt.Registry) error {

	if swag.IsZero(m.NameServers) { // not required
		return nil
	}

	return nil
}

func (m *EthernetInterface100EthernetInterface) validatePermanentMACAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.PermanentMACAddress) { // not required
		return nil
	}

	if err := validate.Pattern("PermanentMACAddress", "body", string(m.PermanentMACAddress), `^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`); err != nil {
		return err
	}

	return nil
}
