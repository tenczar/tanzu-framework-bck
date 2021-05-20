// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// GenerateTKGConfigForVsphereReader is a Reader for the GenerateTKGConfigForVsphere structure.
type GenerateTKGConfigForVsphereReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateTKGConfigForVsphereReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateTKGConfigForVsphereOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGenerateTKGConfigForVsphereBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGenerateTKGConfigForVsphereUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGenerateTKGConfigForVsphereInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGenerateTKGConfigForVsphereOK creates a GenerateTKGConfigForVsphereOK with default headers values
func NewGenerateTKGConfigForVsphereOK() *GenerateTKGConfigForVsphereOK {
	return &GenerateTKGConfigForVsphereOK{}
}

/*GenerateTKGConfigForVsphereOK handles this case with default header values.

Generated TKG configuration successfully
*/
type GenerateTKGConfigForVsphereOK struct {
	Payload string
}

func (o *GenerateTKGConfigForVsphereOK) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere/generate][%d] generateTKGConfigForVsphereOK  %+v", 200, o.Payload)
}

func (o *GenerateTKGConfigForVsphereOK) GetPayload() string {
	return o.Payload
}

func (o *GenerateTKGConfigForVsphereOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateTKGConfigForVsphereBadRequest creates a GenerateTKGConfigForVsphereBadRequest with default headers values
func NewGenerateTKGConfigForVsphereBadRequest() *GenerateTKGConfigForVsphereBadRequest {
	return &GenerateTKGConfigForVsphereBadRequest{}
}

/*GenerateTKGConfigForVsphereBadRequest handles this case with default header values.

Bad request
*/
type GenerateTKGConfigForVsphereBadRequest struct {
	Payload *models.Error
}

func (o *GenerateTKGConfigForVsphereBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere/generate][%d] generateTKGConfigForVsphereBadRequest  %+v", 400, o.Payload)
}

func (o *GenerateTKGConfigForVsphereBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GenerateTKGConfigForVsphereBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateTKGConfigForVsphereUnauthorized creates a GenerateTKGConfigForVsphereUnauthorized with default headers values
func NewGenerateTKGConfigForVsphereUnauthorized() *GenerateTKGConfigForVsphereUnauthorized {
	return &GenerateTKGConfigForVsphereUnauthorized{}
}

/*GenerateTKGConfigForVsphereUnauthorized handles this case with default header values.

Incorrect credentials
*/
type GenerateTKGConfigForVsphereUnauthorized struct {
	Payload *models.Error
}

func (o *GenerateTKGConfigForVsphereUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere/generate][%d] generateTKGConfigForVsphereUnauthorized  %+v", 401, o.Payload)
}

func (o *GenerateTKGConfigForVsphereUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GenerateTKGConfigForVsphereUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateTKGConfigForVsphereInternalServerError creates a GenerateTKGConfigForVsphereInternalServerError with default headers values
func NewGenerateTKGConfigForVsphereInternalServerError() *GenerateTKGConfigForVsphereInternalServerError {
	return &GenerateTKGConfigForVsphereInternalServerError{}
}

/*GenerateTKGConfigForVsphereInternalServerError handles this case with default header values.

Internal server error
*/
type GenerateTKGConfigForVsphereInternalServerError struct {
	Payload *models.Error
}

func (o *GenerateTKGConfigForVsphereInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere/generate][%d] generateTKGConfigForVsphereInternalServerError  %+v", 500, o.Payload)
}

func (o *GenerateTKGConfigForVsphereInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GenerateTKGConfigForVsphereInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
