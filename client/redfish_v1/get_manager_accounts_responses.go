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

// GetManagerAccountsReader is a Reader for the GetManagerAccounts structure.
type GetManagerAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetManagerAccountsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetManagerAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetManagerAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetManagerAccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetManagerAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetManagerAccountsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetManagerAccountsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetManagerAccountsOK creates a GetManagerAccountsOK with default headers values
func NewGetManagerAccountsOK() *GetManagerAccountsOK {
	return &GetManagerAccountsOK{}
}

/*GetManagerAccountsOK handles this case with default header values.

Success
*/
type GetManagerAccountsOK struct {
	Payload *models.ManagerAccountCollectionManagerAccountCollection
}

func (o *GetManagerAccountsOK) Error() string {
	return fmt.Sprintf("[GET /ManagerAccounts][%d] getManagerAccountsOK  %+v", 200, o.Payload)
}

func (o *GetManagerAccountsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagerAccountCollectionManagerAccountCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetManagerAccountsBadRequest creates a GetManagerAccountsBadRequest with default headers values
func NewGetManagerAccountsBadRequest() *GetManagerAccountsBadRequest {
	return &GetManagerAccountsBadRequest{}
}

/*GetManagerAccountsBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information  (such as validation error on an input field, a missing required value, and so on).  An extended error shall be returned in the response body, as defined in section Extended  Error Handling.

*/
type GetManagerAccountsBadRequest struct {
}

func (o *GetManagerAccountsBadRequest) Error() string {
	return fmt.Sprintf("[GET /ManagerAccounts][%d] getManagerAccountsBadRequest ", 400)
}

func (o *GetManagerAccountsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetManagerAccountsUnauthorized creates a GetManagerAccountsUnauthorized with default headers values
func NewGetManagerAccountsUnauthorized() *GetManagerAccountsUnauthorized {
	return &GetManagerAccountsUnauthorized{}
}

/*GetManagerAccountsUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetManagerAccountsUnauthorized struct {
}

func (o *GetManagerAccountsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ManagerAccounts][%d] getManagerAccountsUnauthorized ", 401)
}

func (o *GetManagerAccountsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetManagerAccountsForbidden creates a GetManagerAccountsForbidden with default headers values
func NewGetManagerAccountsForbidden() *GetManagerAccountsForbidden {
	return &GetManagerAccountsForbidden{}
}

/*GetManagerAccountsForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type GetManagerAccountsForbidden struct {
}

func (o *GetManagerAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /ManagerAccounts][%d] getManagerAccountsForbidden ", 403)
}

func (o *GetManagerAccountsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetManagerAccountsNotFound creates a GetManagerAccountsNotFound with default headers values
func NewGetManagerAccountsNotFound() *GetManagerAccountsNotFound {
	return &GetManagerAccountsNotFound{}
}

/*GetManagerAccountsNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type GetManagerAccountsNotFound struct {
}

func (o *GetManagerAccountsNotFound) Error() string {
	return fmt.Sprintf("[GET /ManagerAccounts][%d] getManagerAccountsNotFound ", 404)
}

func (o *GetManagerAccountsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetManagerAccountsInternalServerError creates a GetManagerAccountsInternalServerError with default headers values
func NewGetManagerAccountsInternalServerError() *GetManagerAccountsInternalServerError {
	return &GetManagerAccountsInternalServerError{}
}

/*GetManagerAccountsInternalServerError handles this case with default header values.

Error
*/
type GetManagerAccountsInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetManagerAccountsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ManagerAccounts][%d] getManagerAccountsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetManagerAccountsInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
