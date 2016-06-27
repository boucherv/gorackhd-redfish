package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Manager100ModifyRedundancySet manager 1 0 0 modify redundancy set

swagger:model Manager.1.0.0_ModifyRedundancySet
*/
type Manager100ModifyRedundancySet struct {

	/* Link to invoke action
	 */
	Target strfmt.URI `json:"target,omitempty"`

	/* Friendly action name
	 */
	Title string `json:"title,omitempty"`
}

// Validate validates this manager 1 0 0 modify redundancy set
func (m *Manager100ModifyRedundancySet) Validate(formats strfmt.Registry) error {
	return nil
}
