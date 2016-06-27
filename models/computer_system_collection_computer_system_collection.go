package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
)

/*ComputerSystemCollectionComputerSystemCollection computer system collection computer system collection

swagger:model ComputerSystemCollection_ComputerSystemCollection
*/
type ComputerSystemCollectionComputerSystemCollection struct {

	/* at odata context
	 */
	AtOdataContext Odata400Context `json:"@odata.context,omitempty"`

	/* at odata id
	 */
	AtOdataID Odata400ID `json:"@odata.id,omitempty"`

	/* at odata type
	 */
	AtOdataType Odata400Type `json:"@odata.type,omitempty"`

	/* description
	 */
	Description ResourceDescription `json:"Description,omitempty"`

	/* Contains the members of this collection.

	Read Only: true
	*/
	Members []ComputerSystemComputerSystem `json:"Members,omitempty"`

	/* members at odata count
	 */
	MembersAtOdataCount Odata400Count `json:"Members@odata.count,omitempty"`

	/* members at odata navigation link
	 */
	MembersAtOdataNavigationLink *Odata400IDRef `json:"Members@odata.navigationLink,omitempty"`

	/* name
	 */
	Name ResourceName `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`
}

// Validate validates this computer system collection computer system collection
func (m *ComputerSystemCollectionComputerSystemCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerSystemCollectionComputerSystemCollection) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	return nil
}
