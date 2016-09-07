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

// NewGetAccountParams creates a new GetAccountParams object
// with the default values initialized.
func NewGetAccountParams() *GetAccountParams {
	var ()
	return &GetAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountParamsWithTimeout creates a new GetAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountParamsWithTimeout(timeout time.Duration) *GetAccountParams {
	var ()
	return &GetAccountParams{

		timeout: timeout,
	}
}

/*GetAccountParams contains all the parameters to send to the API endpoint
for the get account operation typically these are written to a http.Request
*/
type GetAccountParams struct {

	/*Name*/
	Name string

	timeout time.Duration
}

// WithName adds the name to the get account params
func (o *GetAccountParams) WithName(name string) *GetAccountParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
