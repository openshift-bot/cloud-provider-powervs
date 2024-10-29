// Code generated by go-swagger; DO NOT EDIT.

package snapshots

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

// NewV1SnapshotsGetParams creates a new V1SnapshotsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1SnapshotsGetParams() *V1SnapshotsGetParams {
	return &V1SnapshotsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1SnapshotsGetParamsWithTimeout creates a new V1SnapshotsGetParams object
// with the ability to set a timeout on a request.
func NewV1SnapshotsGetParamsWithTimeout(timeout time.Duration) *V1SnapshotsGetParams {
	return &V1SnapshotsGetParams{
		timeout: timeout,
	}
}

// NewV1SnapshotsGetParamsWithContext creates a new V1SnapshotsGetParams object
// with the ability to set a context for a request.
func NewV1SnapshotsGetParamsWithContext(ctx context.Context) *V1SnapshotsGetParams {
	return &V1SnapshotsGetParams{
		Context: ctx,
	}
}

// NewV1SnapshotsGetParamsWithHTTPClient creates a new V1SnapshotsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1SnapshotsGetParamsWithHTTPClient(client *http.Client) *V1SnapshotsGetParams {
	return &V1SnapshotsGetParams{
		HTTPClient: client,
	}
}

/*
V1SnapshotsGetParams contains all the parameters to send to the API endpoint

	for the v1 snapshots get operation.

	Typically these are written to a http.Request.
*/
type V1SnapshotsGetParams struct {

	/* SnapshotID.

	   PVM Instance snapshot id
	*/
	SnapshotID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 snapshots get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1SnapshotsGetParams) WithDefaults() *V1SnapshotsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 snapshots get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1SnapshotsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 snapshots get params
func (o *V1SnapshotsGetParams) WithTimeout(timeout time.Duration) *V1SnapshotsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 snapshots get params
func (o *V1SnapshotsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 snapshots get params
func (o *V1SnapshotsGetParams) WithContext(ctx context.Context) *V1SnapshotsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 snapshots get params
func (o *V1SnapshotsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 snapshots get params
func (o *V1SnapshotsGetParams) WithHTTPClient(client *http.Client) *V1SnapshotsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 snapshots get params
func (o *V1SnapshotsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotID adds the snapshotID to the v1 snapshots get params
func (o *V1SnapshotsGetParams) WithSnapshotID(snapshotID string) *V1SnapshotsGetParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the v1 snapshots get params
func (o *V1SnapshotsGetParams) SetSnapshotID(snapshotID string) {
	o.SnapshotID = snapshotID
}

// WriteToRequest writes these params to a swagger request
func (o *V1SnapshotsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param snapshot_id
	if err := r.SetPathParam("snapshot_id", o.SnapshotID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}