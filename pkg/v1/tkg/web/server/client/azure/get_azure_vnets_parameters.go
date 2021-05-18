// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

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

// NewGetAzureVnetsParams creates a new GetAzureVnetsParams object
// with the default values initialized.
func NewGetAzureVnetsParams() *GetAzureVnetsParams {
	var ()
	return &GetAzureVnetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAzureVnetsParamsWithTimeout creates a new GetAzureVnetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAzureVnetsParamsWithTimeout(timeout time.Duration) *GetAzureVnetsParams {
	var ()
	return &GetAzureVnetsParams{

		timeout: timeout,
	}
}

// NewGetAzureVnetsParamsWithContext creates a new GetAzureVnetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAzureVnetsParamsWithContext(ctx context.Context) *GetAzureVnetsParams {
	var ()
	return &GetAzureVnetsParams{

		Context: ctx,
	}
}

// NewGetAzureVnetsParamsWithHTTPClient creates a new GetAzureVnetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAzureVnetsParamsWithHTTPClient(client *http.Client) *GetAzureVnetsParams {
	var ()
	return &GetAzureVnetsParams{
		HTTPClient: client,
	}
}

/*GetAzureVnetsParams contains all the parameters to send to the API endpoint
for the get azure vnets operation typically these are written to a http.Request
*/
type GetAzureVnetsParams struct {

	/*Location
	  Azure region

	*/
	Location string
	/*ResourceGroupName
	  Name of the Azure resource group

	*/
	ResourceGroupName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get azure vnets params
func (o *GetAzureVnetsParams) WithTimeout(timeout time.Duration) *GetAzureVnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get azure vnets params
func (o *GetAzureVnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get azure vnets params
func (o *GetAzureVnetsParams) WithContext(ctx context.Context) *GetAzureVnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get azure vnets params
func (o *GetAzureVnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get azure vnets params
func (o *GetAzureVnetsParams) WithHTTPClient(client *http.Client) *GetAzureVnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get azure vnets params
func (o *GetAzureVnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocation adds the location to the get azure vnets params
func (o *GetAzureVnetsParams) WithLocation(location string) *GetAzureVnetsParams {
	o.SetLocation(location)
	return o
}

// SetLocation adds the location to the get azure vnets params
func (o *GetAzureVnetsParams) SetLocation(location string) {
	o.Location = location
}

// WithResourceGroupName adds the resourceGroupName to the get azure vnets params
func (o *GetAzureVnetsParams) WithResourceGroupName(resourceGroupName string) *GetAzureVnetsParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the get azure vnets params
func (o *GetAzureVnetsParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WriteToRequest writes these params to a swagger request
func (o *GetAzureVnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param location
	qrLocation := o.Location
	qLocation := qrLocation
	if qLocation != "" {
		if err := r.SetQueryParam("location", qLocation); err != nil {
			return err
		}
	}

	// path param resourceGroupName
	if err := r.SetPathParam("resourceGroupName", o.ResourceGroupName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}