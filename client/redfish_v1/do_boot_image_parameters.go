package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd-redfish/models"
)

// NewDoBootImageParams creates a new DoBootImageParams object
// with the default values initialized.
func NewDoBootImageParams() *DoBootImageParams {
	var ()
	return &DoBootImageParams{}
}

/*DoBootImageParams contains all the parameters to send to the API endpoint
for the do boot image operation typically these are written to a http.Request
*/
type DoBootImageParams struct {

	/*Identifier*/
	Identifier string
	/*Payload*/
	Payload *models.RackHDBootImageBootImage
}

// WithIdentifier adds the identifier to the do boot image params
func (o *DoBootImageParams) WithIdentifier(identifier string) *DoBootImageParams {
	o.Identifier = identifier
	return o
}

// WithPayload adds the payload to the do boot image params
func (o *DoBootImageParams) WithPayload(payload *models.RackHDBootImageBootImage) *DoBootImageParams {
	o.Payload = payload
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DoBootImageParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if o.Payload == nil {
		o.Payload = new(models.RackHDBootImageBootImage)
	}

	if err := r.SetBodyParam(o.Payload); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
