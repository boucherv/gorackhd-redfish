package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetSchemaContentParams creates a new GetSchemaContentParams object
// with the default values initialized.
func NewGetSchemaContentParams() *GetSchemaContentParams {
	var ()
	return &GetSchemaContentParams{}
}

/*GetSchemaContentParams contains all the parameters to send to the API endpoint
for the get schema content operation typically these are written to a http.Request
*/
type GetSchemaContentParams struct {

	/*Identifier*/
	Identifier string
}

// WithIdentifier adds the identifier to the get schema content params
func (o *GetSchemaContentParams) WithIdentifier(identifier string) *GetSchemaContentParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetSchemaContentParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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
