// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// GetAWSRegionsReader is a Reader for the GetAWSRegions structure.
type GetAWSRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAWSRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAWSRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAWSRegionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAWSRegionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAWSRegionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAWSRegionsOK creates a GetAWSRegionsOK with default headers values
func NewGetAWSRegionsOK() *GetAWSRegionsOK {
	return &GetAWSRegionsOK{}
}

/*GetAWSRegionsOK handles this case with default header values.

Successful retrieval of AWS regions
*/
type GetAWSRegionsOK struct {
	Payload []string
}

func (o *GetAWSRegionsOK) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/regions][%d] getAWSRegionsOK  %+v", 200, o.Payload)
}

func (o *GetAWSRegionsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAWSRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSRegionsBadRequest creates a GetAWSRegionsBadRequest with default headers values
func NewGetAWSRegionsBadRequest() *GetAWSRegionsBadRequest {
	return &GetAWSRegionsBadRequest{}
}

/*GetAWSRegionsBadRequest handles this case with default header values.

Bad request
*/
type GetAWSRegionsBadRequest struct {
	Payload *models.Error
}

func (o *GetAWSRegionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/regions][%d] getAWSRegionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAWSRegionsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAWSRegionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSRegionsUnauthorized creates a GetAWSRegionsUnauthorized with default headers values
func NewGetAWSRegionsUnauthorized() *GetAWSRegionsUnauthorized {
	return &GetAWSRegionsUnauthorized{}
}

/*GetAWSRegionsUnauthorized handles this case with default header values.

Incorrect credentials
*/
type GetAWSRegionsUnauthorized struct {
	Payload *models.Error
}

func (o *GetAWSRegionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/regions][%d] getAWSRegionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAWSRegionsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAWSRegionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSRegionsInternalServerError creates a GetAWSRegionsInternalServerError with default headers values
func NewGetAWSRegionsInternalServerError() *GetAWSRegionsInternalServerError {
	return &GetAWSRegionsInternalServerError{}
}

/*GetAWSRegionsInternalServerError handles this case with default header values.

Internal server error
*/
type GetAWSRegionsInternalServerError struct {
	Payload *models.Error
}

func (o *GetAWSRegionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/regions][%d] getAWSRegionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAWSRegionsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAWSRegionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}