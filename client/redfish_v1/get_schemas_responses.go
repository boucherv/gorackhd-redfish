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

// GetSchemasReader is a Reader for the GetSchemas structure.
type GetSchemasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSchemasReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSchemasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetSchemasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSchemasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewGetSchemasNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSchemasOK creates a GetSchemasOK with default headers values
func NewGetSchemasOK() *GetSchemasOK {
	return &GetSchemasOK{}
}

/*GetSchemasOK handles this case with default header values.

Success
*/
type GetSchemasOK struct {
	Payload *models.JSONSchemaFileCollectionJSONSchemaFileCollection
}

func (o *GetSchemasOK) Error() string {
	return fmt.Sprintf("[GET /Schemas][%d] getSchemasOK  %+v", 200, o.Payload)
}

func (o *GetSchemasOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONSchemaFileCollectionJSONSchemaFileCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemasUnauthorized creates a GetSchemasUnauthorized with default headers values
func NewGetSchemasUnauthorized() *GetSchemasUnauthorized {
	return &GetSchemasUnauthorized{}
}

/*GetSchemasUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetSchemasUnauthorized struct {
}

func (o *GetSchemasUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Schemas][%d] getSchemasUnauthorized ", 401)
}

func (o *GetSchemasUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSchemasForbidden creates a GetSchemasForbidden with default headers values
func NewGetSchemasForbidden() *GetSchemasForbidden {
	return &GetSchemasForbidden{}
}

/*GetSchemasForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type GetSchemasForbidden struct {
}

func (o *GetSchemasForbidden) Error() string {
	return fmt.Sprintf("[GET /Schemas][%d] getSchemasForbidden ", 403)
}

func (o *GetSchemasForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSchemasNotImplemented creates a GetSchemasNotImplemented with default headers values
func NewGetSchemasNotImplemented() *GetSchemasNotImplemented {
	return &GetSchemasNotImplemented{}
}

/*GetSchemasNotImplemented handles this case with default header values.

The server does not (currently) support the functionality required to fulfill the request.  This is the appropriate response when the server does not recognize the request method  and is not capable of supporting the method for any resource.

*/
type GetSchemasNotImplemented struct {
}

func (o *GetSchemasNotImplemented) Error() string {
	return fmt.Sprintf("[GET /Schemas][%d] getSchemasNotImplemented ", 501)
}

func (o *GetSchemasNotImplemented) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
