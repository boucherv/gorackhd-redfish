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

// GetManagerReader is a Reader for the GetManager structure.
type GetManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetManagerReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetManagerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetManagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetManagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetManagerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetManagerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetManagerOK creates a GetManagerOK with default headers values
func NewGetManagerOK() *GetManagerOK {
	return &GetManagerOK{}
}

/*GetManagerOK handles this case with default header values.

Success
*/
type GetManagerOK struct {
	Payload *models.Manager100Manager
}

func (o *GetManagerOK) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}][%d] getManagerOK  %+v", 200, o.Payload)
}

func (o *GetManagerOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manager100Manager)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetManagerBadRequest creates a GetManagerBadRequest with default headers values
func NewGetManagerBadRequest() *GetManagerBadRequest {
	return &GetManagerBadRequest{}
}

/*GetManagerBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information  (such as validation error on an input field, a missing required value, and so on).  An extended error shall be returned in the response body, as defined in section Extended  Error Handling.

*/
type GetManagerBadRequest struct {
}

func (o *GetManagerBadRequest) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}][%d] getManagerBadRequest ", 400)
}

func (o *GetManagerBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetManagerUnauthorized creates a GetManagerUnauthorized with default headers values
func NewGetManagerUnauthorized() *GetManagerUnauthorized {
	return &GetManagerUnauthorized{}
}

/*GetManagerUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetManagerUnauthorized struct {
}

func (o *GetManagerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}][%d] getManagerUnauthorized ", 401)
}

func (o *GetManagerUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetManagerForbidden creates a GetManagerForbidden with default headers values
func NewGetManagerForbidden() *GetManagerForbidden {
	return &GetManagerForbidden{}
}

/*GetManagerForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type GetManagerForbidden struct {
}

func (o *GetManagerForbidden) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}][%d] getManagerForbidden ", 403)
}

func (o *GetManagerForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetManagerNotFound creates a GetManagerNotFound with default headers values
func NewGetManagerNotFound() *GetManagerNotFound {
	return &GetManagerNotFound{}
}

/*GetManagerNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type GetManagerNotFound struct {
}

func (o *GetManagerNotFound) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}][%d] getManagerNotFound ", 404)
}

func (o *GetManagerNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetManagerInternalServerError creates a GetManagerInternalServerError with default headers values
func NewGetManagerInternalServerError() *GetManagerInternalServerError {
	return &GetManagerInternalServerError{}
}

/*GetManagerInternalServerError handles this case with default header values.

Error
*/
type GetManagerInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetManagerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}][%d] getManagerInternalServerError  %+v", 500, o.Payload)
}

func (o *GetManagerInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
