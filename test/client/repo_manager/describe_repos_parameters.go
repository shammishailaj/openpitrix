// Code generated by go-swagger; DO NOT EDIT.

package repo_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeReposParams creates a new DescribeReposParams object
// with the default values initialized.
func NewDescribeReposParams() *DescribeReposParams {
	var ()
	return &DescribeReposParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeReposParamsWithTimeout creates a new DescribeReposParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeReposParamsWithTimeout(timeout time.Duration) *DescribeReposParams {
	var ()
	return &DescribeReposParams{

		timeout: timeout,
	}
}

// NewDescribeReposParamsWithContext creates a new DescribeReposParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeReposParamsWithContext(ctx context.Context) *DescribeReposParams {
	var ()
	return &DescribeReposParams{

		Context: ctx,
	}
}

// NewDescribeReposParamsWithHTTPClient creates a new DescribeReposParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeReposParamsWithHTTPClient(client *http.Client) *DescribeReposParams {
	var ()
	return &DescribeReposParams{
		HTTPClient: client,
	}
}

/*DescribeReposParams contains all the parameters to send to the API endpoint
for the describe repos operation typically these are written to a http.Request
*/
type DescribeReposParams struct {

	/*Label*/
	Label *string
	/*Limit*/
	Limit *int64
	/*Name*/
	Name []string
	/*Offset*/
	Offset *int64
	/*RepoID*/
	RepoID []string
	/*Selector*/
	Selector *string
	/*Status*/
	Status []string
	/*Type*/
	Type []string
	/*Visibility*/
	Visibility []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe repos params
func (o *DescribeReposParams) WithTimeout(timeout time.Duration) *DescribeReposParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe repos params
func (o *DescribeReposParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe repos params
func (o *DescribeReposParams) WithContext(ctx context.Context) *DescribeReposParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe repos params
func (o *DescribeReposParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe repos params
func (o *DescribeReposParams) WithHTTPClient(client *http.Client) *DescribeReposParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe repos params
func (o *DescribeReposParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabel adds the label to the describe repos params
func (o *DescribeReposParams) WithLabel(label *string) *DescribeReposParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the describe repos params
func (o *DescribeReposParams) SetLabel(label *string) {
	o.Label = label
}

// WithLimit adds the limit to the describe repos params
func (o *DescribeReposParams) WithLimit(limit *int64) *DescribeReposParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe repos params
func (o *DescribeReposParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the describe repos params
func (o *DescribeReposParams) WithName(name []string) *DescribeReposParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the describe repos params
func (o *DescribeReposParams) SetName(name []string) {
	o.Name = name
}

// WithOffset adds the offset to the describe repos params
func (o *DescribeReposParams) WithOffset(offset *int64) *DescribeReposParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe repos params
func (o *DescribeReposParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithRepoID adds the repoID to the describe repos params
func (o *DescribeReposParams) WithRepoID(repoID []string) *DescribeReposParams {
	o.SetRepoID(repoID)
	return o
}

// SetRepoID adds the repoId to the describe repos params
func (o *DescribeReposParams) SetRepoID(repoID []string) {
	o.RepoID = repoID
}

// WithSelector adds the selector to the describe repos params
func (o *DescribeReposParams) WithSelector(selector *string) *DescribeReposParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the describe repos params
func (o *DescribeReposParams) SetSelector(selector *string) {
	o.Selector = selector
}

// WithStatus adds the status to the describe repos params
func (o *DescribeReposParams) WithStatus(status []string) *DescribeReposParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe repos params
func (o *DescribeReposParams) SetStatus(status []string) {
	o.Status = status
}

// WithType adds the typeVar to the describe repos params
func (o *DescribeReposParams) WithType(typeVar []string) *DescribeReposParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the describe repos params
func (o *DescribeReposParams) SetType(typeVar []string) {
	o.Type = typeVar
}

// WithVisibility adds the visibility to the describe repos params
func (o *DescribeReposParams) WithVisibility(visibility []string) *DescribeReposParams {
	o.SetVisibility(visibility)
	return o
}

// SetVisibility adds the visibility to the describe repos params
func (o *DescribeReposParams) SetVisibility(visibility []string) {
	o.Visibility = visibility
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeReposParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Label != nil {

		// query param label
		var qrLabel string
		if o.Label != nil {
			qrLabel = *o.Label
		}
		qLabel := qrLabel
		if qLabel != "" {
			if err := r.SetQueryParam("label", qLabel); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	valuesName := o.Name

	joinedName := swag.JoinByFormat(valuesName, "")
	// query array param name
	if err := r.SetQueryParam("name", joinedName...); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesRepoID := o.RepoID

	joinedRepoID := swag.JoinByFormat(valuesRepoID, "")
	// query array param repo_id
	if err := r.SetQueryParam("repo_id", joinedRepoID...); err != nil {
		return err
	}

	if o.Selector != nil {

		// query param selector
		var qrSelector string
		if o.Selector != nil {
			qrSelector = *o.Selector
		}
		qSelector := qrSelector
		if qSelector != "" {
			if err := r.SetQueryParam("selector", qSelector); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	valuesType := o.Type

	joinedType := swag.JoinByFormat(valuesType, "")
	// query array param type
	if err := r.SetQueryParam("type", joinedType...); err != nil {
		return err
	}

	valuesVisibility := o.Visibility

	joinedVisibility := swag.JoinByFormat(valuesVisibility, "")
	// query array param visibility
	if err := r.SetQueryParam("visibility", joinedVisibility...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
