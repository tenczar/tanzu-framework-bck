// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAzureEndpointParams creates a new GetAzureEndpointParams object
// with the default values initialized.
func NewGetAzureEndpointParams() *GetAzureEndpointParams {

	return &GetAzureEndpointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAzureEndpointParamsWithTimeout creates a new GetAzureEndpointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAzureEndpointParamsWithTimeout(timeout time.Duration) *GetAzureEndpointParams {

	return &GetAzureEndpointParams{

		timeout: timeout,
	}
}

// NewGetAzureEndpointParamsWithContext creates a new GetAzureEndpointParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAzureEndpointParamsWithContext(ctx context.Context) *GetAzureEndpointParams {

	return &GetAzureEndpointParams{

		Context: ctx,
	}
}

// NewGetAzureEndpointParamsWithHTTPClient creates a new GetAzureEndpointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAzureEndpointParamsWithHTTPClient(client *http.Client) *GetAzureEndpointParams {

	return &GetAzureEndpointParams{
		HTTPClient: client,
	}
}

/*GetAzureEndpointParams contains all the parameters to send to the API endpoint
for the get azure endpoint operation typically these are written to a http.Request
*/
type GetAzureEndpointParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get azure endpoint params
func (o *GetAzureEndpointParams) WithTimeout(timeout time.Duration) *GetAzureEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get azure endpoint params
func (o *GetAzureEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get azure endpoint params
func (o *GetAzureEndpointParams) WithContext(ctx context.Context) *GetAzureEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get azure endpoint params
func (o *GetAzureEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get azure endpoint params
func (o *GetAzureEndpointParams) WithHTTPClient(client *http.Client) *GetAzureEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get azure endpoint params
func (o *GetAzureEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAzureEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
