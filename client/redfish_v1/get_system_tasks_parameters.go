package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSystemTasksParams creates a new GetSystemTasksParams object
// with the default values initialized.
func NewGetSystemTasksParams() *GetSystemTasksParams {
	var ()
	return &GetSystemTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemTasksParamsWithTimeout creates a new GetSystemTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSystemTasksParamsWithTimeout(timeout time.Duration) *GetSystemTasksParams {
	var ()
	return &GetSystemTasksParams{

		timeout: timeout,
	}
}

/*GetSystemTasksParams contains all the parameters to send to the API endpoint
for the get system tasks operation typically these are written to a http.Request
*/
type GetSystemTasksParams struct {

	/*Identifier*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the get system tasks params
func (o *GetSystemTasksParams) WithIdentifier(identifier string) *GetSystemTasksParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
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
