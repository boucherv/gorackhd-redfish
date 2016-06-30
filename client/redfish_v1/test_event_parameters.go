package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd-redfish/models"
)

// NewTestEventParams creates a new TestEventParams object
// with the default values initialized.
func NewTestEventParams() *TestEventParams {
	var ()
	return &TestEventParams{}
}

/*TestEventParams contains all the parameters to send to the API endpoint
for the test event operation typically these are written to a http.Request
*/
type TestEventParams struct {

	/*Body*/
	Body models.TestEvent
}

// WithBody adds the body to the test event params
func (o *TestEventParams) WithBody(body models.TestEvent) *TestEventParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *TestEventParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
