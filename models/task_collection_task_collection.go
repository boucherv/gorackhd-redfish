package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*TaskCollectionTaskCollection task collection task collection

swagger:model TaskCollection_TaskCollection
*/
type TaskCollectionTaskCollection struct {

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

	/* Provides a description of this resource and is used for commonality  in the schema definitions.

	Read Only: true
	*/
	Description string `json:"Description,omitempty"`

	/* Contains the members of this collection.

	Read Only: true
	*/
	Members []*Odata400IDRef `json:"Members,omitempty"`

	/* members at odata count

	Read Only: true
	*/
	MembersAtOdataCount float64 `json:"Members@odata.count,omitempty"`

	/* members at odata navigation link
	 */
	MembersAtOdataNavigationLink *Odata400IDRef `json:"Members@odata.navigationLink,omitempty"`

	/* The name of the resource or array element.

	Read Only: true
	*/
	Name string `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`
}

// Validate validates this task collection task collection
func (m *TaskCollectionTaskCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMembersAtOdataNavigationLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskCollectionTaskCollection) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {

		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {

			if err := m.Members[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *TaskCollectionTaskCollection) validateMembersAtOdataNavigationLink(formats strfmt.Registry) error {

	if swag.IsZero(m.MembersAtOdataNavigationLink) { // not required
		return nil
	}

	if m.MembersAtOdataNavigationLink != nil {

		if err := m.MembersAtOdataNavigationLink.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
