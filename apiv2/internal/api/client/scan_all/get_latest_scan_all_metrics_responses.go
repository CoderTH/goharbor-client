// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v3/apiv2/model"
)

// GetLatestScanAllMetricsReader is a Reader for the GetLatestScanAllMetrics structure.
type GetLatestScanAllMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestScanAllMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatestScanAllMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLatestScanAllMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLatestScanAllMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewGetLatestScanAllMetricsPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLatestScanAllMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLatestScanAllMetricsOK creates a GetLatestScanAllMetricsOK with default headers values
func NewGetLatestScanAllMetricsOK() *GetLatestScanAllMetricsOK {
	return &GetLatestScanAllMetricsOK{}
}

/* GetLatestScanAllMetricsOK describes a response with status code 200, with default header values.

OK
*/
type GetLatestScanAllMetricsOK struct {
	Payload *model.Stats
}

func (o *GetLatestScanAllMetricsOK) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsOK  %+v", 200, o.Payload)
}
func (o *GetLatestScanAllMetricsOK) GetPayload() *model.Stats {
	return o.Payload
}

func (o *GetLatestScanAllMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Stats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestScanAllMetricsUnauthorized creates a GetLatestScanAllMetricsUnauthorized with default headers values
func NewGetLatestScanAllMetricsUnauthorized() *GetLatestScanAllMetricsUnauthorized {
	return &GetLatestScanAllMetricsUnauthorized{}
}

/* GetLatestScanAllMetricsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetLatestScanAllMetricsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLatestScanAllMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetLatestScanAllMetricsUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLatestScanAllMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestScanAllMetricsForbidden creates a GetLatestScanAllMetricsForbidden with default headers values
func NewGetLatestScanAllMetricsForbidden() *GetLatestScanAllMetricsForbidden {
	return &GetLatestScanAllMetricsForbidden{}
}

/* GetLatestScanAllMetricsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLatestScanAllMetricsForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLatestScanAllMetricsForbidden) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsForbidden  %+v", 403, o.Payload)
}
func (o *GetLatestScanAllMetricsForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLatestScanAllMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestScanAllMetricsPreconditionFailed creates a GetLatestScanAllMetricsPreconditionFailed with default headers values
func NewGetLatestScanAllMetricsPreconditionFailed() *GetLatestScanAllMetricsPreconditionFailed {
	return &GetLatestScanAllMetricsPreconditionFailed{}
}

/* GetLatestScanAllMetricsPreconditionFailed describes a response with status code 412, with default header values.

Precondition failed
*/
type GetLatestScanAllMetricsPreconditionFailed struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLatestScanAllMetricsPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsPreconditionFailed  %+v", 412, o.Payload)
}
func (o *GetLatestScanAllMetricsPreconditionFailed) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLatestScanAllMetricsPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestScanAllMetricsInternalServerError creates a GetLatestScanAllMetricsInternalServerError with default headers values
func NewGetLatestScanAllMetricsInternalServerError() *GetLatestScanAllMetricsInternalServerError {
	return &GetLatestScanAllMetricsInternalServerError{}
}

/* GetLatestScanAllMetricsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetLatestScanAllMetricsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLatestScanAllMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetLatestScanAllMetricsInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLatestScanAllMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
