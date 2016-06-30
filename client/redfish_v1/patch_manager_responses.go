package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd-redfish/models"
)

// PatchManagerReader is a Reader for the PatchManager structure.
type PatchManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchManagerReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchManagerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchManagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPatchManagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchManagerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPatchManagerMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPatchManagerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchManagerOK creates a PatchManagerOK with default headers values
func NewPatchManagerOK() *PatchManagerOK {
	return &PatchManagerOK{}
}

/*PatchManagerOK handles this case with default header values.

Success
*/
type PatchManagerOK struct {
	Payload *models.Manager100Manager
}

func (o *PatchManagerOK) Error() string {
	return fmt.Sprintf("[PATCH /Managers/{identifier}][%d] patchManagerOK  %+v", 200, o.Payload)
}

func (o *PatchManagerOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manager100Manager)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchManagerBadRequest creates a PatchManagerBadRequest with default headers values
func NewPatchManagerBadRequest() *PatchManagerBadRequest {
	return &PatchManagerBadRequest{}
}

/*PatchManagerBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information  (such as validation error on an input field, a missing required value, and so on).  An extended error shall be returned in the response body, as defined in section Extended  Error Handling.

*/
type PatchManagerBadRequest struct {
}

func (o *PatchManagerBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /Managers/{identifier}][%d] patchManagerBadRequest ", 400)
}

func (o *PatchManagerBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchManagerUnauthorized creates a PatchManagerUnauthorized with default headers values
func NewPatchManagerUnauthorized() *PatchManagerUnauthorized {
	return &PatchManagerUnauthorized{}
}

/*PatchManagerUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type PatchManagerUnauthorized struct {
}

func (o *PatchManagerUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /Managers/{identifier}][%d] patchManagerUnauthorized ", 401)
}

func (o *PatchManagerUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchManagerForbidden creates a PatchManagerForbidden with default headers values
func NewPatchManagerForbidden() *PatchManagerForbidden {
	return &PatchManagerForbidden{}
}

/*PatchManagerForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type PatchManagerForbidden struct {
}

func (o *PatchManagerForbidden) Error() string {
	return fmt.Sprintf("[PATCH /Managers/{identifier}][%d] patchManagerForbidden ", 403)
}

func (o *PatchManagerForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchManagerNotFound creates a PatchManagerNotFound with default headers values
func NewPatchManagerNotFound() *PatchManagerNotFound {
	return &PatchManagerNotFound{}
}

/*PatchManagerNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type PatchManagerNotFound struct {
}

func (o *PatchManagerNotFound) Error() string {
	return fmt.Sprintf("[PATCH /Managers/{identifier}][%d] patchManagerNotFound ", 404)
}

func (o *PatchManagerNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchManagerMethodNotAllowed creates a PatchManagerMethodNotAllowed with default headers values
func NewPatchManagerMethodNotAllowed() *PatchManagerMethodNotAllowed {
	return &PatchManagerMethodNotAllowed{}
}

/*PatchManagerMethodNotAllowed handles this case with default header values.

The method specified in the Request-Line is not allowed for the resource identified by the Request-URI.

*/
type PatchManagerMethodNotAllowed struct {
}

func (o *PatchManagerMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /Managers/{identifier}][%d] patchManagerMethodNotAllowed ", 405)
}

func (o *PatchManagerMethodNotAllowed) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchManagerInternalServerError creates a PatchManagerInternalServerError with default headers values
func NewPatchManagerInternalServerError() *PatchManagerInternalServerError {
	return &PatchManagerInternalServerError{}
}

/*PatchManagerInternalServerError handles this case with default header values.

Error
*/
type PatchManagerInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PatchManagerInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /Managers/{identifier}][%d] patchManagerInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchManagerInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
