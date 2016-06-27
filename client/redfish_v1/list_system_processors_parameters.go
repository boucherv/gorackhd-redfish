package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewListSystemProcessorsParams creates a new ListSystemProcessorsParams object
// with the default values initialized.
func NewListSystemProcessorsParams() *ListSystemProcessorsParams {
	var ()
	return &ListSystemProcessorsParams{}
}

/*ListSystemProcessorsParams contains all the parameters to send to the API endpoint
for the list system processors operation typically these are written to a http.Request
*/
type ListSystemProcessorsParams struct {

	/*Identifier*/
	Identifier string
}

// WithIdentifier adds the identifier to the list system processors params
func (o *ListSystemProcessorsParams) WithIdentifier(identifier string) *ListSystemProcessorsParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ListSystemProcessorsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
