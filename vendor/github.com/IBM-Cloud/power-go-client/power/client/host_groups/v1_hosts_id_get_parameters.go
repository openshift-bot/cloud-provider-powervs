// Code generated by go-swagger; DO NOT EDIT.

package host_groups

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

// NewV1HostsIDGetParams creates a new V1HostsIDGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1HostsIDGetParams() *V1HostsIDGetParams {
	return &V1HostsIDGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1HostsIDGetParamsWithTimeout creates a new V1HostsIDGetParams object
// with the ability to set a timeout on a request.
func NewV1HostsIDGetParamsWithTimeout(timeout time.Duration) *V1HostsIDGetParams {
	return &V1HostsIDGetParams{
		timeout: timeout,
	}
}

// NewV1HostsIDGetParamsWithContext creates a new V1HostsIDGetParams object
// with the ability to set a context for a request.
func NewV1HostsIDGetParamsWithContext(ctx context.Context) *V1HostsIDGetParams {
	return &V1HostsIDGetParams{
		Context: ctx,
	}
}

// NewV1HostsIDGetParamsWithHTTPClient creates a new V1HostsIDGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1HostsIDGetParamsWithHTTPClient(client *http.Client) *V1HostsIDGetParams {
	return &V1HostsIDGetParams{
		HTTPClient: client,
	}
}

/*
V1HostsIDGetParams contains all the parameters to send to the API endpoint

	for the v1 hosts id get operation.

	Typically these are written to a http.Request.
*/
type V1HostsIDGetParams struct {

	/* HostID.

	   Host ID
	*/
	HostID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 hosts id get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1HostsIDGetParams) WithDefaults() *V1HostsIDGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 hosts id get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1HostsIDGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 hosts id get params
func (o *V1HostsIDGetParams) WithTimeout(timeout time.Duration) *V1HostsIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 hosts id get params
func (o *V1HostsIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 hosts id get params
func (o *V1HostsIDGetParams) WithContext(ctx context.Context) *V1HostsIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 hosts id get params
func (o *V1HostsIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 hosts id get params
func (o *V1HostsIDGetParams) WithHTTPClient(client *http.Client) *V1HostsIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 hosts id get params
func (o *V1HostsIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHostID adds the hostID to the v1 hosts id get params
func (o *V1HostsIDGetParams) WithHostID(hostID string) *V1HostsIDGetParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the v1 hosts id get params
func (o *V1HostsIDGetParams) SetHostID(hostID string) {
	o.HostID = hostID
}

// WriteToRequest writes these params to a swagger request
func (o *V1HostsIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}