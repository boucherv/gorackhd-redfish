package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd-redfish/models"
)

// NewCreateAccountParams creates a new CreateAccountParams object
// with the default values initialized.
func NewCreateAccountParams() *CreateAccountParams {
	var ()
	return &CreateAccountParams{}
}

/*CreateAccountParams contains all the parameters to send to the API endpoint
for the create account operation typically these are written to a http.Request
*/
type CreateAccountParams struct {

	/*Payload*/
	Payload *models.ManagerAccount100ManagerAccount
}

// WithPayload adds the payload to the create account params
func (o *CreateAccountParams) WithPayload(payload *models.ManagerAccount100ManagerAccount) *CreateAccountParams {
	o.Payload = payload
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccountParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Payload == nil {
		o.Payload = new(models.ManagerAccount100ManagerAccount)
	}

	if err := r.SetBodyParam(o.Payload); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
