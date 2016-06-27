package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
)

/*VirtualMedia100VirtualMedia This is the schema definition for the Virtual Media Service.

swagger:model VirtualMedia.1.0.0_VirtualMedia
*/
type VirtualMedia100VirtualMedia struct {

	/* at odata context
	 */
	AtOdataContext Odata400Context `json:"@odata.context,omitempty"`

	/* at odata id
	 */
	AtOdataID Odata400ID `json:"@odata.id,omitempty"`

	/* at odata type
	 */
	AtOdataType Odata400Type `json:"@odata.type,omitempty"`

	/* Current virtual media connection methods

	Read Only: true
	*/
	ConnectedVia VirtualMedia100ConnectedVia `json:"ConnectedVia,omitempty"`

	/* description
	 */
	Description ResourceDescription `json:"Description,omitempty"`

	/* Id
	 */
	ID ResourceID `json:"Id,omitempty"`

	/* A URI providing the location of the selected image

	Read Only: true
	*/
	Image string `json:"Image,omitempty"`

	/* The current image name

	Read Only: true
	*/
	ImageName string `json:"ImageName,omitempty"`

	/* Indicates if virtual media is inserted in the virtual device.

	Read Only: true
	*/
	Inserted *bool `json:"Inserted,omitempty"`

	/* This is the media types supported as virtual media.

	Read Only: true
	*/
	MediaTypes []VirtualMedia100MediaType `json:"MediaTypes,omitempty"`

	/* name
	 */
	Name ResourceName `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* Indicates the media is write protected.

	Read Only: true
	*/
	WriteProtected *bool `json:"WriteProtected,omitempty"`
}

// Validate validates this virtual media 1 0 0 virtual media
func (m *VirtualMedia100VirtualMedia) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectedVia(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMediaTypes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMedia100VirtualMedia) validateConnectedVia(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectedVia) { // not required
		return nil
	}

	if err := m.ConnectedVia.Validate(formats); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMedia100VirtualMedia) validateMediaTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaTypes) { // not required
		return nil
	}

	return nil
}
