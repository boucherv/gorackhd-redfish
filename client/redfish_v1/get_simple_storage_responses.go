package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd-redfish/models"
)

// GetSimpleStorageReader is a Reader for the GetSimpleStorage structure.
type GetSimpleStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSimpleStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSimpleStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSimpleStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetSimpleStorageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSimpleStorageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSimpleStorageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSimpleStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSimpleStorageOK creates a GetSimpleStorageOK with default headers values
func NewGetSimpleStorageOK() *GetSimpleStorageOK {
	return &GetSimpleStorageOK{}
}

/*GetSimpleStorageOK handles this case with default header values.

Success
*/
type GetSimpleStorageOK struct {
	Payload *models.SimpleStorage100SimpleStorage
}

func (o *GetSimpleStorageOK) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/SimpleStorage/{index}][%d] getSimpleStorageOK  %+v", 200, o.Payload)
}

func (o *GetSimpleStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleStorage100SimpleStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSimpleStorageBadRequest creates a GetSimpleStorageBadRequest with default headers values
func NewGetSimpleStorageBadRequest() *GetSimpleStorageBadRequest {
	return &GetSimpleStorageBadRequest{}
}

/*GetSimpleStorageBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information  (such as validation error on an input field, a missing required value, and so on).  An extended error shall be returned in the response body, as defined in section Extended  Error Handling.

*/
type GetSimpleStorageBadRequest struct {
}

func (o *GetSimpleStorageBadRequest) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/SimpleStorage/{index}][%d] getSimpleStorageBadRequest ", 400)
}

func (o *GetSimpleStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSimpleStorageUnauthorized creates a GetSimpleStorageUnauthorized with default headers values
func NewGetSimpleStorageUnauthorized() *GetSimpleStorageUnauthorized {
	return &GetSimpleStorageUnauthorized{}
}

/*GetSimpleStorageUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetSimpleStorageUnauthorized struct {
}

func (o *GetSimpleStorageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/SimpleStorage/{index}][%d] getSimpleStorageUnauthorized ", 401)
}

func (o *GetSimpleStorageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSimpleStorageForbidden creates a GetSimpleStorageForbidden with default headers values
func NewGetSimpleStorageForbidden() *GetSimpleStorageForbidden {
	return &GetSimpleStorageForbidden{}
}

/*GetSimpleStorageForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type GetSimpleStorageForbidden struct {
}

func (o *GetSimpleStorageForbidden) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/SimpleStorage/{index}][%d] getSimpleStorageForbidden ", 403)
}

func (o *GetSimpleStorageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSimpleStorageNotFound creates a GetSimpleStorageNotFound with default headers values
func NewGetSimpleStorageNotFound() *GetSimpleStorageNotFound {
	return &GetSimpleStorageNotFound{}
}

/*GetSimpleStorageNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type GetSimpleStorageNotFound struct {
}

func (o *GetSimpleStorageNotFound) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/SimpleStorage/{index}][%d] getSimpleStorageNotFound ", 404)
}

func (o *GetSimpleStorageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSimpleStorageInternalServerError creates a GetSimpleStorageInternalServerError with default headers values
func NewGetSimpleStorageInternalServerError() *GetSimpleStorageInternalServerError {
	return &GetSimpleStorageInternalServerError{}
}

/*GetSimpleStorageInternalServerError handles this case with default header values.

Error
*/
type GetSimpleStorageInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetSimpleStorageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/SimpleStorage/{index}][%d] getSimpleStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSimpleStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
