// Code generated by go-swagger; DO NOT EDIT.

package power_edge_router

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

// V1PoweredgerouterActionPostReader is a Reader for the V1PoweredgerouterActionPost structure.
type V1PoweredgerouterActionPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1PoweredgerouterActionPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1PoweredgerouterActionPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1PoweredgerouterActionPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1PoweredgerouterActionPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1PoweredgerouterActionPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1PoweredgerouterActionPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV1PoweredgerouterActionPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1PoweredgerouterActionPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/workspaces/{workspace_id}/power-edge-router/action] v1.poweredgerouter.action.post", response, response.Code())
	}
}

// NewV1PoweredgerouterActionPostOK creates a V1PoweredgerouterActionPostOK with default headers values
func NewV1PoweredgerouterActionPostOK() *V1PoweredgerouterActionPostOK {
	return &V1PoweredgerouterActionPostOK{}
}

/*
V1PoweredgerouterActionPostOK describes a response with status code 200, with default header values.

OK
*/
type V1PoweredgerouterActionPostOK struct {
	Payload models.Object
}

// IsSuccess returns true when this v1 poweredgerouter action post o k response has a 2xx status code
func (o *V1PoweredgerouterActionPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 poweredgerouter action post o k response has a 3xx status code
func (o *V1PoweredgerouterActionPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 poweredgerouter action post o k response has a 4xx status code
func (o *V1PoweredgerouterActionPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 poweredgerouter action post o k response has a 5xx status code
func (o *V1PoweredgerouterActionPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 poweredgerouter action post o k response a status code equal to that given
func (o *V1PoweredgerouterActionPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 poweredgerouter action post o k response
func (o *V1PoweredgerouterActionPostOK) Code() int {
	return 200
}

func (o *V1PoweredgerouterActionPostOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostOK %s", 200, payload)
}

func (o *V1PoweredgerouterActionPostOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostOK %s", 200, payload)
}

func (o *V1PoweredgerouterActionPostOK) GetPayload() models.Object {
	return o.Payload
}

func (o *V1PoweredgerouterActionPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1PoweredgerouterActionPostBadRequest creates a V1PoweredgerouterActionPostBadRequest with default headers values
func NewV1PoweredgerouterActionPostBadRequest() *V1PoweredgerouterActionPostBadRequest {
	return &V1PoweredgerouterActionPostBadRequest{}
}

/*
V1PoweredgerouterActionPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1PoweredgerouterActionPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 poweredgerouter action post bad request response has a 2xx status code
func (o *V1PoweredgerouterActionPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 poweredgerouter action post bad request response has a 3xx status code
func (o *V1PoweredgerouterActionPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 poweredgerouter action post bad request response has a 4xx status code
func (o *V1PoweredgerouterActionPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 poweredgerouter action post bad request response has a 5xx status code
func (o *V1PoweredgerouterActionPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 poweredgerouter action post bad request response a status code equal to that given
func (o *V1PoweredgerouterActionPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 poweredgerouter action post bad request response
func (o *V1PoweredgerouterActionPostBadRequest) Code() int {
	return 400
}

func (o *V1PoweredgerouterActionPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostBadRequest %s", 400, payload)
}

func (o *V1PoweredgerouterActionPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostBadRequest %s", 400, payload)
}

func (o *V1PoweredgerouterActionPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1PoweredgerouterActionPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1PoweredgerouterActionPostUnauthorized creates a V1PoweredgerouterActionPostUnauthorized with default headers values
func NewV1PoweredgerouterActionPostUnauthorized() *V1PoweredgerouterActionPostUnauthorized {
	return &V1PoweredgerouterActionPostUnauthorized{}
}

/*
V1PoweredgerouterActionPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1PoweredgerouterActionPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 poweredgerouter action post unauthorized response has a 2xx status code
func (o *V1PoweredgerouterActionPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 poweredgerouter action post unauthorized response has a 3xx status code
func (o *V1PoweredgerouterActionPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 poweredgerouter action post unauthorized response has a 4xx status code
func (o *V1PoweredgerouterActionPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 poweredgerouter action post unauthorized response has a 5xx status code
func (o *V1PoweredgerouterActionPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 poweredgerouter action post unauthorized response a status code equal to that given
func (o *V1PoweredgerouterActionPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 poweredgerouter action post unauthorized response
func (o *V1PoweredgerouterActionPostUnauthorized) Code() int {
	return 401
}

func (o *V1PoweredgerouterActionPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostUnauthorized %s", 401, payload)
}

func (o *V1PoweredgerouterActionPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostUnauthorized %s", 401, payload)
}

func (o *V1PoweredgerouterActionPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1PoweredgerouterActionPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1PoweredgerouterActionPostForbidden creates a V1PoweredgerouterActionPostForbidden with default headers values
func NewV1PoweredgerouterActionPostForbidden() *V1PoweredgerouterActionPostForbidden {
	return &V1PoweredgerouterActionPostForbidden{}
}

/*
V1PoweredgerouterActionPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1PoweredgerouterActionPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 poweredgerouter action post forbidden response has a 2xx status code
func (o *V1PoweredgerouterActionPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 poweredgerouter action post forbidden response has a 3xx status code
func (o *V1PoweredgerouterActionPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 poweredgerouter action post forbidden response has a 4xx status code
func (o *V1PoweredgerouterActionPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 poweredgerouter action post forbidden response has a 5xx status code
func (o *V1PoweredgerouterActionPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 poweredgerouter action post forbidden response a status code equal to that given
func (o *V1PoweredgerouterActionPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 poweredgerouter action post forbidden response
func (o *V1PoweredgerouterActionPostForbidden) Code() int {
	return 403
}

func (o *V1PoweredgerouterActionPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostForbidden %s", 403, payload)
}

func (o *V1PoweredgerouterActionPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostForbidden %s", 403, payload)
}

func (o *V1PoweredgerouterActionPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1PoweredgerouterActionPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1PoweredgerouterActionPostNotFound creates a V1PoweredgerouterActionPostNotFound with default headers values
func NewV1PoweredgerouterActionPostNotFound() *V1PoweredgerouterActionPostNotFound {
	return &V1PoweredgerouterActionPostNotFound{}
}

/*
V1PoweredgerouterActionPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1PoweredgerouterActionPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 poweredgerouter action post not found response has a 2xx status code
func (o *V1PoweredgerouterActionPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 poweredgerouter action post not found response has a 3xx status code
func (o *V1PoweredgerouterActionPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 poweredgerouter action post not found response has a 4xx status code
func (o *V1PoweredgerouterActionPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 poweredgerouter action post not found response has a 5xx status code
func (o *V1PoweredgerouterActionPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 poweredgerouter action post not found response a status code equal to that given
func (o *V1PoweredgerouterActionPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 poweredgerouter action post not found response
func (o *V1PoweredgerouterActionPostNotFound) Code() int {
	return 404
}

func (o *V1PoweredgerouterActionPostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostNotFound %s", 404, payload)
}

func (o *V1PoweredgerouterActionPostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostNotFound %s", 404, payload)
}

func (o *V1PoweredgerouterActionPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1PoweredgerouterActionPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1PoweredgerouterActionPostConflict creates a V1PoweredgerouterActionPostConflict with default headers values
func NewV1PoweredgerouterActionPostConflict() *V1PoweredgerouterActionPostConflict {
	return &V1PoweredgerouterActionPostConflict{}
}

/*
V1PoweredgerouterActionPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type V1PoweredgerouterActionPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 poweredgerouter action post conflict response has a 2xx status code
func (o *V1PoweredgerouterActionPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 poweredgerouter action post conflict response has a 3xx status code
func (o *V1PoweredgerouterActionPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 poweredgerouter action post conflict response has a 4xx status code
func (o *V1PoweredgerouterActionPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 poweredgerouter action post conflict response has a 5xx status code
func (o *V1PoweredgerouterActionPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 poweredgerouter action post conflict response a status code equal to that given
func (o *V1PoweredgerouterActionPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the v1 poweredgerouter action post conflict response
func (o *V1PoweredgerouterActionPostConflict) Code() int {
	return 409
}

func (o *V1PoweredgerouterActionPostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostConflict %s", 409, payload)
}

func (o *V1PoweredgerouterActionPostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostConflict %s", 409, payload)
}

func (o *V1PoweredgerouterActionPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1PoweredgerouterActionPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1PoweredgerouterActionPostInternalServerError creates a V1PoweredgerouterActionPostInternalServerError with default headers values
func NewV1PoweredgerouterActionPostInternalServerError() *V1PoweredgerouterActionPostInternalServerError {
	return &V1PoweredgerouterActionPostInternalServerError{}
}

/*
V1PoweredgerouterActionPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1PoweredgerouterActionPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 poweredgerouter action post internal server error response has a 2xx status code
func (o *V1PoweredgerouterActionPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 poweredgerouter action post internal server error response has a 3xx status code
func (o *V1PoweredgerouterActionPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 poweredgerouter action post internal server error response has a 4xx status code
func (o *V1PoweredgerouterActionPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 poweredgerouter action post internal server error response has a 5xx status code
func (o *V1PoweredgerouterActionPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 poweredgerouter action post internal server error response a status code equal to that given
func (o *V1PoweredgerouterActionPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 poweredgerouter action post internal server error response
func (o *V1PoweredgerouterActionPostInternalServerError) Code() int {
	return 500
}

func (o *V1PoweredgerouterActionPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostInternalServerError %s", 500, payload)
}

func (o *V1PoweredgerouterActionPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/workspaces/{workspace_id}/power-edge-router/action][%d] v1PoweredgerouterActionPostInternalServerError %s", 500, payload)
}

func (o *V1PoweredgerouterActionPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1PoweredgerouterActionPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
