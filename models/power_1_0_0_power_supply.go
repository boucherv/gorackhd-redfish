package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*Power100PowerSupply Details of a power supplies associated with this system or device

swagger:model Power.1.0.0_PowerSupply
*/
type Power100PowerSupply struct {

	/* The firmware version for this Power Supply

	Read Only: true
	*/
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	/* The average power output of this Power Supply

	Read Only: true
	Minimum: 0
	*/
	LastPowerOutputWatts float64 `json:"LastPowerOutputWatts,omitempty"`

	/* The line input voltage at which the Power Supply is operating

	Read Only: true
	*/
	LineInputVoltage float64 `json:"LineInputVoltage,omitempty"`

	/* The line voltage type supported as an input to this Power Supply

	Read Only: true
	*/
	LineInputVoltageType Power100LineInputVoltageType `json:"LineInputVoltageType,omitempty"`

	/* This is the identifier for the member within the collection.
	 */
	MemberID string `json:"MemberId,omitempty"`

	/* The model number for this Power Supply

	Read Only: true
	*/
	Model string `json:"Model,omitempty"`

	/* The name of the Power Supply

	Read Only: true
	*/
	Name string `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* The part number for this Power Supply

	Read Only: true
	*/
	PartNumber string `json:"PartNumber,omitempty"`

	/* The maximum capacity of this Power Supply

	Read Only: true
	Minimum: 0
	*/
	PowerCapacityWatts float64 `json:"PowerCapacityWatts,omitempty"`

	/* The Power Supply type (AC or DC)

	Read Only: true
	*/
	PowerSupplyType Power100PowerSupplyType `json:"PowerSupplyType,omitempty"`

	/* This structure is used to show redundancy for fans.  The Component ids will reference the members of the redundancy groups.

	Read Only: true
	*/
	Redundancy []RedundancyRedundancy `json:"Redundancy,omitempty"`

	/* redundancy at odata count
	 */
	RedundancyAtOdataCount Odata400Count `json:"Redundancy@odata.count,omitempty"`

	/* redundancy at odata navigation link
	 */
	RedundancyAtOdataNavigationLink *Odata400IDRef `json:"Redundancy@odata.navigationLink,omitempty"`

	/* The ID(s) of the resources associated with this Power Limit
	 */
	RelatedItem []*Odata400IDRef `json:"RelatedItem,omitempty"`

	/* related item at odata count
	 */
	RelatedItemAtOdataCount Odata400Count `json:"RelatedItem@odata.count,omitempty"`

	/* related item at odata navigation link
	 */
	RelatedItemAtOdataNavigationLink *Odata400IDRef `json:"RelatedItem@odata.navigationLink,omitempty"`

	/* The serial number for this Power Supply

	Read Only: true
	*/
	SerialNumber string `json:"SerialNumber,omitempty"`

	/* The spare part number for this Power Supply

	Read Only: true
	*/
	SparePartNumber string `json:"SparePartNumber,omitempty"`

	/* status
	 */
	Status *ResourceStatus `json:"Status,omitempty"`
}

// Validate validates this power 1 0 0 power supply
func (m *Power100PowerSupply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastPowerOutputWatts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLineInputVoltageType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePowerCapacityWatts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePowerSupplyType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRedundancy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRelatedItem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Power100PowerSupply) validateLastPowerOutputWatts(formats strfmt.Registry) error {

	if swag.IsZero(m.LastPowerOutputWatts) { // not required
		return nil
	}

	if err := validate.Minimum("LastPowerOutputWatts", "body", float64(m.LastPowerOutputWatts), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Power100PowerSupply) validateLineInputVoltageType(formats strfmt.Registry) error {

	if swag.IsZero(m.LineInputVoltageType) { // not required
		return nil
	}

	if err := m.LineInputVoltageType.Validate(formats); err != nil {
		return err
	}

	return nil
}

func (m *Power100PowerSupply) validatePowerCapacityWatts(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerCapacityWatts) { // not required
		return nil
	}

	if err := validate.Minimum("PowerCapacityWatts", "body", float64(m.PowerCapacityWatts), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Power100PowerSupply) validatePowerSupplyType(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerSupplyType) { // not required
		return nil
	}

	if err := m.PowerSupplyType.Validate(formats); err != nil {
		return err
	}

	return nil
}

func (m *Power100PowerSupply) validateRedundancy(formats strfmt.Registry) error {

	if swag.IsZero(m.Redundancy) { // not required
		return nil
	}

	return nil
}

func (m *Power100PowerSupply) validateRelatedItem(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedItem) { // not required
		return nil
	}

	return nil
}
