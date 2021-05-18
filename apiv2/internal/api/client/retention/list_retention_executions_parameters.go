// Code generated by go-swagger; DO NOT EDIT.

package retention

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
	"github.com/go-openapi/swag"
)

// NewListRetentionExecutionsParams creates a new ListRetentionExecutionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListRetentionExecutionsParams() *ListRetentionExecutionsParams {
	return &ListRetentionExecutionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListRetentionExecutionsParamsWithTimeout creates a new ListRetentionExecutionsParams object
// with the ability to set a timeout on a request.
func NewListRetentionExecutionsParamsWithTimeout(timeout time.Duration) *ListRetentionExecutionsParams {
	return &ListRetentionExecutionsParams{
		timeout: timeout,
	}
}

// NewListRetentionExecutionsParamsWithContext creates a new ListRetentionExecutionsParams object
// with the ability to set a context for a request.
func NewListRetentionExecutionsParamsWithContext(ctx context.Context) *ListRetentionExecutionsParams {
	return &ListRetentionExecutionsParams{
		Context: ctx,
	}
}

// NewListRetentionExecutionsParamsWithHTTPClient creates a new ListRetentionExecutionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListRetentionExecutionsParamsWithHTTPClient(client *http.Client) *ListRetentionExecutionsParams {
	return &ListRetentionExecutionsParams{
		HTTPClient: client,
	}
}

/* ListRetentionExecutionsParams contains all the parameters to send to the API endpoint
   for the list retention executions operation.

   Typically these are written to a http.Request.
*/
type ListRetentionExecutionsParams struct {

	/* ID.

	   Retention ID.

	   Format: int64
	*/
	ID int64

	/* Page.

	   The page number.

	   Format: int64
	*/
	Page *int64

	/* PageSize.

	   The size of per page.

	   Format: int64
	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list retention executions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRetentionExecutionsParams) WithDefaults() *ListRetentionExecutionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list retention executions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRetentionExecutionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list retention executions params
func (o *ListRetentionExecutionsParams) WithTimeout(timeout time.Duration) *ListRetentionExecutionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list retention executions params
func (o *ListRetentionExecutionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list retention executions params
func (o *ListRetentionExecutionsParams) WithContext(ctx context.Context) *ListRetentionExecutionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list retention executions params
func (o *ListRetentionExecutionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list retention executions params
func (o *ListRetentionExecutionsParams) WithHTTPClient(client *http.Client) *ListRetentionExecutionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list retention executions params
func (o *ListRetentionExecutionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list retention executions params
func (o *ListRetentionExecutionsParams) WithID(id int64) *ListRetentionExecutionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list retention executions params
func (o *ListRetentionExecutionsParams) SetID(id int64) {
	o.ID = id
}

// WithPage adds the page to the list retention executions params
func (o *ListRetentionExecutionsParams) WithPage(page *int64) *ListRetentionExecutionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list retention executions params
func (o *ListRetentionExecutionsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list retention executions params
func (o *ListRetentionExecutionsParams) WithPageSize(pageSize *int64) *ListRetentionExecutionsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list retention executions params
func (o *ListRetentionExecutionsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *ListRetentionExecutionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
