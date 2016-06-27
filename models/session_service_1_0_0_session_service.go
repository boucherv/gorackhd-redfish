package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*SessionService100SessionService This is the schema definition for the Session Service.  It represents the properties for the service itself and has links to the actual list of sessions.

swagger:model SessionService.1.0.0_SessionService
*/
type SessionService100SessionService struct {

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

	/* Id
	 */
	ID ResourceID `json:"Id,omitempty"`

	/* name
	 */
	Name ResourceName `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* This indicates whether this service is enabled.
	 */
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	/* This is the number of seconds of inactivity that a session may have before the session service closes the session due to inactivity.

	Maximum: 86400
	Minimum: 30
	*/
	SessionTimeout float64 `json:"SessionTimeout,omitempty"`

	/* Link to a collection of Sessions

	Read Only: true
	*/
	Sessions *SessionCollectionSessionCollection `json:"Sessions,omitempty"`

	/* status
	 */
	Status *ResourceStatus `json:"Status,omitempty"`
}

// Validate validates this session service 1 0 0 session service
func (m *SessionService100SessionService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSessionTimeout(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SessionService100SessionService) validateSessionTimeout(formats strfmt.Registry) error {

	if swag.IsZero(m.SessionTimeout) { // not required
		return nil
	}

	if err := validate.Minimum("SessionTimeout", "body", float64(m.SessionTimeout), 30, false); err != nil {
		return err
	}

	if err := validate.Maximum("SessionTimeout", "body", float64(m.SessionTimeout), 86400, false); err != nil {
		return err
	}

	return nil
}
