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

// GetEventsCollectionReader is a Reader for the GetEventsCollection structure.
type GetEventsCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetEventsCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventsCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetEventsCollectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetEventsCollectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewGetEventsCollectionNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventsCollectionOK creates a GetEventsCollectionOK with default headers values
func NewGetEventsCollectionOK() *GetEventsCollectionOK {
	return &GetEventsCollectionOK{}
}

/*GetEventsCollectionOK handles this case with default header values.

Success
*/
type GetEventsCollectionOK struct {
	Payload *models.EventDestinationCollectionEventDestinationCollection
}

func (o *GetEventsCollectionOK) Error() string {
	return fmt.Sprintf("[GET /EventService/Subscriptions][%d] getEventsCollectionOK  %+v", 200, o.Payload)
}

func (o *GetEventsCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventDestinationCollectionEventDestinationCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsCollectionUnauthorized creates a GetEventsCollectionUnauthorized with default headers values
func NewGetEventsCollectionUnauthorized() *GetEventsCollectionUnauthorized {
	return &GetEventsCollectionUnauthorized{}
}

/*GetEventsCollectionUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetEventsCollectionUnauthorized struct {
}

func (o *GetEventsCollectionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /EventService/Subscriptions][%d] getEventsCollectionUnauthorized ", 401)
}

func (o *GetEventsCollectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEventsCollectionForbidden creates a GetEventsCollectionForbidden with default headers values
func NewGetEventsCollectionForbidden() *GetEventsCollectionForbidden {
	return &GetEventsCollectionForbidden{}
}

/*GetEventsCollectionForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type GetEventsCollectionForbidden struct {
}

func (o *GetEventsCollectionForbidden) Error() string {
	return fmt.Sprintf("[GET /EventService/Subscriptions][%d] getEventsCollectionForbidden ", 403)
}

func (o *GetEventsCollectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEventsCollectionNotImplemented creates a GetEventsCollectionNotImplemented with default headers values
func NewGetEventsCollectionNotImplemented() *GetEventsCollectionNotImplemented {
	return &GetEventsCollectionNotImplemented{}
}

/*GetEventsCollectionNotImplemented handles this case with default header values.

The server does not (currently) support the functionality required to fulfill the request.  This is the appropriate response when the server does not recognize the request method  and is not capable of supporting the method for any resource.

*/
type GetEventsCollectionNotImplemented struct {
}

func (o *GetEventsCollectionNotImplemented) Error() string {
	return fmt.Sprintf("[GET /EventService/Subscriptions][%d] getEventsCollectionNotImplemented ", 501)
}

func (o *GetEventsCollectionNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
