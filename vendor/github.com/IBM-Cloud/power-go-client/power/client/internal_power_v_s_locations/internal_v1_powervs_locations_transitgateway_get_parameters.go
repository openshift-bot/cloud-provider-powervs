// Code generated by go-swagger; DO NOT EDIT.

package internal_power_v_s_locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewInternalV1PowervsLocationsTransitgatewayGetParams creates a new InternalV1PowervsLocationsTransitgatewayGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInternalV1PowervsLocationsTransitgatewayGetParams() *InternalV1PowervsLocationsTransitgatewayGetParams {
	return &InternalV1PowervsLocationsTransitgatewayGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInternalV1PowervsLocationsTransitgatewayGetParamsWithTimeout creates a new InternalV1PowervsLocationsTransitgatewayGetParams object
// with the ability to set a timeout on a request.
func NewInternalV1PowervsLocationsTransitgatewayGetParamsWithTimeout(timeout time.Duration) *InternalV1PowervsLocationsTransitgatewayGetParams {
	return &InternalV1PowervsLocationsTransitgatewayGetParams{
		timeout: timeout,
	}
}

// NewInternalV1PowervsLocationsTransitgatewayGetParamsWithContext creates a new InternalV1PowervsLocationsTransitgatewayGetParams object
// with the ability to set a context for a request.
func NewInternalV1PowervsLocationsTransitgatewayGetParamsWithContext(ctx context.Context) *InternalV1PowervsLocationsTransitgatewayGetParams {
	return &InternalV1PowervsLocationsTransitgatewayGetParams{
		Context: ctx,
	}
}

// NewInternalV1PowervsLocationsTransitgatewayGetParamsWithHTTPClient creates a new InternalV1PowervsLocationsTransitgatewayGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewInternalV1PowervsLocationsTransitgatewayGetParamsWithHTTPClient(client *http.Client) *InternalV1PowervsLocationsTransitgatewayGetParams {
	return &InternalV1PowervsLocationsTransitgatewayGetParams{
		HTTPClient: client,
	}
}

/*
InternalV1PowervsLocationsTransitgatewayGetParams contains all the parameters to send to the API endpoint

	for the internal v1 powervs locations transitgateway get operation.

	Typically these are written to a http.Request.
*/
type InternalV1PowervsLocationsTransitgatewayGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the internal v1 powervs locations transitgateway get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1PowervsLocationsTransitgatewayGetParams) WithDefaults() *InternalV1PowervsLocationsTransitgatewayGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the internal v1 powervs locations transitgateway get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1PowervsLocationsTransitgatewayGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the internal v1 powervs locations transitgateway get params
func (o *InternalV1PowervsLocationsTransitgatewayGetParams) WithTimeout(timeout time.Duration) *InternalV1PowervsLocationsTransitgatewayGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the internal v1 powervs locations transitgateway get params
func (o *InternalV1PowervsLocationsTransitgatewayGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the internal v1 powervs locations transitgateway get params
func (o *InternalV1PowervsLocationsTransitgatewayGetParams) WithContext(ctx context.Context) *InternalV1PowervsLocationsTransitgatewayGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the internal v1 powervs locations transitgateway get params
func (o *InternalV1PowervsLocationsTransitgatewayGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the internal v1 powervs locations transitgateway get params
func (o *InternalV1PowervsLocationsTransitgatewayGetParams) WithHTTPClient(client *http.Client) *InternalV1PowervsLocationsTransitgatewayGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the internal v1 powervs locations transitgateway get params
func (o *InternalV1PowervsLocationsTransitgatewayGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *InternalV1PowervsLocationsTransitgatewayGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
