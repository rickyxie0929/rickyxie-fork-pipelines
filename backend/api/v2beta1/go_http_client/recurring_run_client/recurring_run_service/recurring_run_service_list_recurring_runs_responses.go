// Code generated by go-swagger; DO NOT EDIT.

package recurring_run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	recurring_run_model "github.com/kubeflow/pipelines/backend/api/v2beta1/go_http_client/recurring_run_model"
)

// RecurringRunServiceListRecurringRunsReader is a Reader for the RecurringRunServiceListRecurringRuns structure.
type RecurringRunServiceListRecurringRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecurringRunServiceListRecurringRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRecurringRunServiceListRecurringRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRecurringRunServiceListRecurringRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRecurringRunServiceListRecurringRunsOK creates a RecurringRunServiceListRecurringRunsOK with default headers values
func NewRecurringRunServiceListRecurringRunsOK() *RecurringRunServiceListRecurringRunsOK {
	return &RecurringRunServiceListRecurringRunsOK{}
}

/*RecurringRunServiceListRecurringRunsOK handles this case with default header values.

A successful response.
*/
type RecurringRunServiceListRecurringRunsOK struct {
	Payload *recurring_run_model.V2beta1ListRecurringRunsResponse
}

func (o *RecurringRunServiceListRecurringRunsOK) Error() string {
	return fmt.Sprintf("[GET /apis/v2beta1/recurringruns][%d] recurringRunServiceListRecurringRunsOK  %+v", 200, o.Payload)
}

func (o *RecurringRunServiceListRecurringRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(recurring_run_model.V2beta1ListRecurringRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecurringRunServiceListRecurringRunsDefault creates a RecurringRunServiceListRecurringRunsDefault with default headers values
func NewRecurringRunServiceListRecurringRunsDefault(code int) *RecurringRunServiceListRecurringRunsDefault {
	return &RecurringRunServiceListRecurringRunsDefault{
		_statusCode: code,
	}
}

/*RecurringRunServiceListRecurringRunsDefault handles this case with default header values.

An unexpected error response.
*/
type RecurringRunServiceListRecurringRunsDefault struct {
	_statusCode int

	Payload *recurring_run_model.RuntimeError
}

// Code gets the status code for the recurring run service list recurring runs default response
func (o *RecurringRunServiceListRecurringRunsDefault) Code() int {
	return o._statusCode
}

func (o *RecurringRunServiceListRecurringRunsDefault) Error() string {
	return fmt.Sprintf("[GET /apis/v2beta1/recurringruns][%d] RecurringRunService_ListRecurringRuns default  %+v", o._statusCode, o.Payload)
}

func (o *RecurringRunServiceListRecurringRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(recurring_run_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}