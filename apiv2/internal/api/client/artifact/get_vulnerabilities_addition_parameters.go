// Code generated by go-swagger; DO NOT EDIT.

package artifact

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

// NewGetVulnerabilitiesAdditionParams creates a new GetVulnerabilitiesAdditionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVulnerabilitiesAdditionParams() *GetVulnerabilitiesAdditionParams {
	return &GetVulnerabilitiesAdditionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVulnerabilitiesAdditionParamsWithTimeout creates a new GetVulnerabilitiesAdditionParams object
// with the ability to set a timeout on a request.
func NewGetVulnerabilitiesAdditionParamsWithTimeout(timeout time.Duration) *GetVulnerabilitiesAdditionParams {
	return &GetVulnerabilitiesAdditionParams{
		timeout: timeout,
	}
}

// NewGetVulnerabilitiesAdditionParamsWithContext creates a new GetVulnerabilitiesAdditionParams object
// with the ability to set a context for a request.
func NewGetVulnerabilitiesAdditionParamsWithContext(ctx context.Context) *GetVulnerabilitiesAdditionParams {
	return &GetVulnerabilitiesAdditionParams{
		Context: ctx,
	}
}

// NewGetVulnerabilitiesAdditionParamsWithHTTPClient creates a new GetVulnerabilitiesAdditionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVulnerabilitiesAdditionParamsWithHTTPClient(client *http.Client) *GetVulnerabilitiesAdditionParams {
	return &GetVulnerabilitiesAdditionParams{
		HTTPClient: client,
	}
}

/* GetVulnerabilitiesAdditionParams contains all the parameters to send to the API endpoint
   for the get vulnerabilities addition operation.

   Typically these are written to a http.Request.
*/
type GetVulnerabilitiesAdditionParams struct {

	/* XAcceptVulnerabilities.

	     A comma-separated lists of MIME types for the scan report or scan summary. The first mime type will be used when the report found for it.
	Currently the mime type supports 'application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0' and 'application/vnd.security.vulnerability.report; version=1.1'

	     Default: "application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0"
	*/
	XAcceptVulnerabilities *string

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* ProjectName.

	   The name of the project
	*/
	ProjectName string

	/* Reference.

	   The reference of the artifact, can be digest or tag
	*/
	Reference string

	/* RepositoryName.

	   The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -> a%252Fb
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vulnerabilities addition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVulnerabilitiesAdditionParams) WithDefaults() *GetVulnerabilitiesAdditionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vulnerabilities addition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVulnerabilitiesAdditionParams) SetDefaults() {
	var (
		xAcceptVulnerabilitiesDefault = string("application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0")
	)

	val := GetVulnerabilitiesAdditionParams{
		XAcceptVulnerabilities: &xAcceptVulnerabilitiesDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) WithTimeout(timeout time.Duration) *GetVulnerabilitiesAdditionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) WithContext(ctx context.Context) *GetVulnerabilitiesAdditionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) WithHTTPClient(client *http.Client) *GetVulnerabilitiesAdditionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAcceptVulnerabilities adds the xAcceptVulnerabilities to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) WithXAcceptVulnerabilities(xAcceptVulnerabilities *string) *GetVulnerabilitiesAdditionParams {
	o.SetXAcceptVulnerabilities(xAcceptVulnerabilities)
	return o
}

// SetXAcceptVulnerabilities adds the xAcceptVulnerabilities to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) SetXAcceptVulnerabilities(xAcceptVulnerabilities *string) {
	o.XAcceptVulnerabilities = xAcceptVulnerabilities
}

// WithXRequestID adds the xRequestID to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) WithXRequestID(xRequestID *string) *GetVulnerabilitiesAdditionParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectName adds the projectName to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) WithProjectName(projectName string) *GetVulnerabilitiesAdditionParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithReference adds the reference to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) WithReference(reference string) *GetVulnerabilitiesAdditionParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) SetReference(reference string) {
	o.Reference = reference
}

// WithRepositoryName adds the repositoryName to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) WithRepositoryName(repositoryName string) *GetVulnerabilitiesAdditionParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get vulnerabilities addition params
func (o *GetVulnerabilitiesAdditionParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetVulnerabilitiesAdditionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XAcceptVulnerabilities != nil {

		// header param X-Accept-Vulnerabilities
		if err := r.SetHeaderParam("X-Accept-Vulnerabilities", *o.XAcceptVulnerabilities); err != nil {
			return err
		}
	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	// path param reference
	if err := r.SetPathParam("reference", o.Reference); err != nil {
		return err
	}

	// path param repository_name
	if err := r.SetPathParam("repository_name", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
