// Code generated by go-swagger; DO NOT EDIT.

package host_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// V1HostsIDGetReader is a Reader for the V1HostsIDGet structure.
type V1HostsIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1HostsIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1HostsIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1HostsIDGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1HostsIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1HostsIDGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1HostsIDGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1HostsIDGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewV1HostsIDGetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/hosts/{host_id}] v1.hosts.id.get", response, response.Code())
	}
}

// NewV1HostsIDGetOK creates a V1HostsIDGetOK with default headers values
func NewV1HostsIDGetOK() *V1HostsIDGetOK {
	return &V1HostsIDGetOK{}
}

/*
V1HostsIDGetOK describes a response with status code 200, with default header values.

OK
*/
type V1HostsIDGetOK struct {
	Payload *models.Host
}

// IsSuccess returns true when this v1 hosts Id get o k response has a 2xx status code
func (o *V1HostsIDGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 hosts Id get o k response has a 3xx status code
func (o *V1HostsIDGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hosts Id get o k response has a 4xx status code
func (o *V1HostsIDGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 hosts Id get o k response has a 5xx status code
func (o *V1HostsIDGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hosts Id get o k response a status code equal to that given
func (o *V1HostsIDGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 hosts Id get o k response
func (o *V1HostsIDGetOK) Code() int {
	return 200
}

func (o *V1HostsIDGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetOK %s", 200, payload)
}

func (o *V1HostsIDGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetOK %s", 200, payload)
}

func (o *V1HostsIDGetOK) GetPayload() *models.Host {
	return o.Payload
}

func (o *V1HostsIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Host)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostsIDGetBadRequest creates a V1HostsIDGetBadRequest with default headers values
func NewV1HostsIDGetBadRequest() *V1HostsIDGetBadRequest {
	return &V1HostsIDGetBadRequest{}
}

/*
V1HostsIDGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1HostsIDGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hosts Id get bad request response has a 2xx status code
func (o *V1HostsIDGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hosts Id get bad request response has a 3xx status code
func (o *V1HostsIDGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hosts Id get bad request response has a 4xx status code
func (o *V1HostsIDGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 hosts Id get bad request response has a 5xx status code
func (o *V1HostsIDGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hosts Id get bad request response a status code equal to that given
func (o *V1HostsIDGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 hosts Id get bad request response
func (o *V1HostsIDGetBadRequest) Code() int {
	return 400
}

func (o *V1HostsIDGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetBadRequest %s", 400, payload)
}

func (o *V1HostsIDGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetBadRequest %s", 400, payload)
}

func (o *V1HostsIDGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostsIDGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostsIDGetUnauthorized creates a V1HostsIDGetUnauthorized with default headers values
func NewV1HostsIDGetUnauthorized() *V1HostsIDGetUnauthorized {
	return &V1HostsIDGetUnauthorized{}
}

/*
V1HostsIDGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1HostsIDGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hosts Id get unauthorized response has a 2xx status code
func (o *V1HostsIDGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hosts Id get unauthorized response has a 3xx status code
func (o *V1HostsIDGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hosts Id get unauthorized response has a 4xx status code
func (o *V1HostsIDGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 hosts Id get unauthorized response has a 5xx status code
func (o *V1HostsIDGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hosts Id get unauthorized response a status code equal to that given
func (o *V1HostsIDGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 hosts Id get unauthorized response
func (o *V1HostsIDGetUnauthorized) Code() int {
	return 401
}

func (o *V1HostsIDGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetUnauthorized %s", 401, payload)
}

func (o *V1HostsIDGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetUnauthorized %s", 401, payload)
}

func (o *V1HostsIDGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostsIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostsIDGetForbidden creates a V1HostsIDGetForbidden with default headers values
func NewV1HostsIDGetForbidden() *V1HostsIDGetForbidden {
	return &V1HostsIDGetForbidden{}
}

/*
V1HostsIDGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1HostsIDGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hosts Id get forbidden response has a 2xx status code
func (o *V1HostsIDGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hosts Id get forbidden response has a 3xx status code
func (o *V1HostsIDGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hosts Id get forbidden response has a 4xx status code
func (o *V1HostsIDGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 hosts Id get forbidden response has a 5xx status code
func (o *V1HostsIDGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hosts Id get forbidden response a status code equal to that given
func (o *V1HostsIDGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 hosts Id get forbidden response
func (o *V1HostsIDGetForbidden) Code() int {
	return 403
}

func (o *V1HostsIDGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetForbidden %s", 403, payload)
}

func (o *V1HostsIDGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetForbidden %s", 403, payload)
}

func (o *V1HostsIDGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostsIDGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostsIDGetNotFound creates a V1HostsIDGetNotFound with default headers values
func NewV1HostsIDGetNotFound() *V1HostsIDGetNotFound {
	return &V1HostsIDGetNotFound{}
}

/*
V1HostsIDGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1HostsIDGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hosts Id get not found response has a 2xx status code
func (o *V1HostsIDGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hosts Id get not found response has a 3xx status code
func (o *V1HostsIDGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hosts Id get not found response has a 4xx status code
func (o *V1HostsIDGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 hosts Id get not found response has a 5xx status code
func (o *V1HostsIDGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hosts Id get not found response a status code equal to that given
func (o *V1HostsIDGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 hosts Id get not found response
func (o *V1HostsIDGetNotFound) Code() int {
	return 404
}

func (o *V1HostsIDGetNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetNotFound %s", 404, payload)
}

func (o *V1HostsIDGetNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetNotFound %s", 404, payload)
}

func (o *V1HostsIDGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostsIDGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostsIDGetInternalServerError creates a V1HostsIDGetInternalServerError with default headers values
func NewV1HostsIDGetInternalServerError() *V1HostsIDGetInternalServerError {
	return &V1HostsIDGetInternalServerError{}
}

/*
V1HostsIDGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1HostsIDGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hosts Id get internal server error response has a 2xx status code
func (o *V1HostsIDGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hosts Id get internal server error response has a 3xx status code
func (o *V1HostsIDGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hosts Id get internal server error response has a 4xx status code
func (o *V1HostsIDGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 hosts Id get internal server error response has a 5xx status code
func (o *V1HostsIDGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 hosts Id get internal server error response a status code equal to that given
func (o *V1HostsIDGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 hosts Id get internal server error response
func (o *V1HostsIDGetInternalServerError) Code() int {
	return 500
}

func (o *V1HostsIDGetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetInternalServerError %s", 500, payload)
}

func (o *V1HostsIDGetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetInternalServerError %s", 500, payload)
}

func (o *V1HostsIDGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostsIDGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostsIDGetGatewayTimeout creates a V1HostsIDGetGatewayTimeout with default headers values
func NewV1HostsIDGetGatewayTimeout() *V1HostsIDGetGatewayTimeout {
	return &V1HostsIDGetGatewayTimeout{}
}

/*
V1HostsIDGetGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. Request is still processing and taking longer than expected.
*/
type V1HostsIDGetGatewayTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hosts Id get gateway timeout response has a 2xx status code
func (o *V1HostsIDGetGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hosts Id get gateway timeout response has a 3xx status code
func (o *V1HostsIDGetGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hosts Id get gateway timeout response has a 4xx status code
func (o *V1HostsIDGetGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 hosts Id get gateway timeout response has a 5xx status code
func (o *V1HostsIDGetGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 hosts Id get gateway timeout response a status code equal to that given
func (o *V1HostsIDGetGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the v1 hosts Id get gateway timeout response
func (o *V1HostsIDGetGatewayTimeout) Code() int {
	return 504
}

func (o *V1HostsIDGetGatewayTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetGatewayTimeout %s", 504, payload)
}

func (o *V1HostsIDGetGatewayTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{host_id}][%d] v1HostsIdGetGatewayTimeout %s", 504, payload)
}

func (o *V1HostsIDGetGatewayTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostsIDGetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
