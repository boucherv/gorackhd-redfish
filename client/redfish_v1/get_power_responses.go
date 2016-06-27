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

// GetPowerReader is a Reader for the GetPower structure.
type GetPowerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPowerReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPowerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPowerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetPowerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetPowerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetPowerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPowerOK creates a GetPowerOK with default headers values
func NewGetPowerOK() *GetPowerOK {
	return &GetPowerOK{}
}

/*GetPowerOK handles this case with default header values.

Success
*/
type GetPowerOK struct {
	Payload *models.Power100Power
}

func (o *GetPowerOK) Error() string {
	return fmt.Sprintf("[GET /Chassis/{identifier}/Power][%d] getPowerOK  %+v", 200, o.Payload)
}

func (o *GetPowerOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Power100Power)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPowerBadRequest creates a GetPowerBadRequest with default headers values
func NewGetPowerBadRequest() *GetPowerBadRequest {
	return &GetPowerBadRequest{}
}

/*GetPowerBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information  (such as validation error on an input field, a missing required value, and so on).  An extended error shall be returned in the response body, as defined in section Extended  Error Handling.

*/
type GetPowerBadRequest struct {
}

func (o *GetPowerBadRequest) Error() string {
	return fmt.Sprintf("[GET /Chassis/{identifier}/Power][%d] getPowerBadRequest ", 400)
}

func (o *GetPowerBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPowerUnauthorized creates a GetPowerUnauthorized with default headers values
func NewGetPowerUnauthorized() *GetPowerUnauthorized {
	return &GetPowerUnauthorized{}
}

/*GetPowerUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetPowerUnauthorized struct {
}

func (o *GetPowerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Chassis/{identifier}/Power][%d] getPowerUnauthorized ", 401)
}

func (o *GetPowerUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPowerForbidden creates a GetPowerForbidden with default headers values
func NewGetPowerForbidden() *GetPowerForbidden {
	return &GetPowerForbidden{}
}

/*GetPowerForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type GetPowerForbidden struct {
}

func (o *GetPowerForbidden) Error() string {
	return fmt.Sprintf("[GET /Chassis/{identifier}/Power][%d] getPowerForbidden ", 403)
}

func (o *GetPowerForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPowerInternalServerError creates a GetPowerInternalServerError with default headers values
func NewGetPowerInternalServerError() *GetPowerInternalServerError {
	return &GetPowerInternalServerError{}
}

/*GetPowerInternalServerError handles this case with default header values.

Error
*/
type GetPowerInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetPowerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Chassis/{identifier}/Power][%d] getPowerInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPowerInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
