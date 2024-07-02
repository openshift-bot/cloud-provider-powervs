// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

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

// PcloudVolumesClonePostReader is a Reader for the PcloudVolumesClonePost structure.
type PcloudVolumesClonePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVolumesClonePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVolumesClonePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVolumesClonePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVolumesClonePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVolumesClonePostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVolumesClonePostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudVolumesClonePostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVolumesClonePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone] pcloud.volumes.clone.post", response, response.Code())
	}
}

// NewPcloudVolumesClonePostOK creates a PcloudVolumesClonePostOK with default headers values
func NewPcloudVolumesClonePostOK() *PcloudVolumesClonePostOK {
	return &PcloudVolumesClonePostOK{}
}

/*
PcloudVolumesClonePostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVolumesClonePostOK struct {
	Payload *models.VolumesCloneResponse
}

// IsSuccess returns true when this pcloud volumes clone post o k response has a 2xx status code
func (o *PcloudVolumesClonePostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud volumes clone post o k response has a 3xx status code
func (o *PcloudVolumesClonePostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumes clone post o k response has a 4xx status code
func (o *PcloudVolumesClonePostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volumes clone post o k response has a 5xx status code
func (o *PcloudVolumesClonePostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumes clone post o k response a status code equal to that given
func (o *PcloudVolumesClonePostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud volumes clone post o k response
func (o *PcloudVolumesClonePostOK) Code() int {
	return 200
}

func (o *PcloudVolumesClonePostOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostOK %s", 200, payload)
}

func (o *PcloudVolumesClonePostOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostOK %s", 200, payload)
}

func (o *PcloudVolumesClonePostOK) GetPayload() *models.VolumesCloneResponse {
	return o.Payload
}

func (o *PcloudVolumesClonePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumesCloneResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumesClonePostBadRequest creates a PcloudVolumesClonePostBadRequest with default headers values
func NewPcloudVolumesClonePostBadRequest() *PcloudVolumesClonePostBadRequest {
	return &PcloudVolumesClonePostBadRequest{}
}

/*
PcloudVolumesClonePostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVolumesClonePostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumes clone post bad request response has a 2xx status code
func (o *PcloudVolumesClonePostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumes clone post bad request response has a 3xx status code
func (o *PcloudVolumesClonePostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumes clone post bad request response has a 4xx status code
func (o *PcloudVolumesClonePostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumes clone post bad request response has a 5xx status code
func (o *PcloudVolumesClonePostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumes clone post bad request response a status code equal to that given
func (o *PcloudVolumesClonePostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud volumes clone post bad request response
func (o *PcloudVolumesClonePostBadRequest) Code() int {
	return 400
}

func (o *PcloudVolumesClonePostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostBadRequest %s", 400, payload)
}

func (o *PcloudVolumesClonePostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostBadRequest %s", 400, payload)
}

func (o *PcloudVolumesClonePostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumesClonePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumesClonePostUnauthorized creates a PcloudVolumesClonePostUnauthorized with default headers values
func NewPcloudVolumesClonePostUnauthorized() *PcloudVolumesClonePostUnauthorized {
	return &PcloudVolumesClonePostUnauthorized{}
}

/*
PcloudVolumesClonePostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVolumesClonePostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumes clone post unauthorized response has a 2xx status code
func (o *PcloudVolumesClonePostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumes clone post unauthorized response has a 3xx status code
func (o *PcloudVolumesClonePostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumes clone post unauthorized response has a 4xx status code
func (o *PcloudVolumesClonePostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumes clone post unauthorized response has a 5xx status code
func (o *PcloudVolumesClonePostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumes clone post unauthorized response a status code equal to that given
func (o *PcloudVolumesClonePostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud volumes clone post unauthorized response
func (o *PcloudVolumesClonePostUnauthorized) Code() int {
	return 401
}

func (o *PcloudVolumesClonePostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostUnauthorized %s", 401, payload)
}

func (o *PcloudVolumesClonePostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostUnauthorized %s", 401, payload)
}

func (o *PcloudVolumesClonePostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumesClonePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumesClonePostForbidden creates a PcloudVolumesClonePostForbidden with default headers values
func NewPcloudVolumesClonePostForbidden() *PcloudVolumesClonePostForbidden {
	return &PcloudVolumesClonePostForbidden{}
}

/*
PcloudVolumesClonePostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVolumesClonePostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumes clone post forbidden response has a 2xx status code
func (o *PcloudVolumesClonePostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumes clone post forbidden response has a 3xx status code
func (o *PcloudVolumesClonePostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumes clone post forbidden response has a 4xx status code
func (o *PcloudVolumesClonePostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumes clone post forbidden response has a 5xx status code
func (o *PcloudVolumesClonePostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumes clone post forbidden response a status code equal to that given
func (o *PcloudVolumesClonePostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud volumes clone post forbidden response
func (o *PcloudVolumesClonePostForbidden) Code() int {
	return 403
}

func (o *PcloudVolumesClonePostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostForbidden %s", 403, payload)
}

func (o *PcloudVolumesClonePostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostForbidden %s", 403, payload)
}

func (o *PcloudVolumesClonePostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumesClonePostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumesClonePostNotFound creates a PcloudVolumesClonePostNotFound with default headers values
func NewPcloudVolumesClonePostNotFound() *PcloudVolumesClonePostNotFound {
	return &PcloudVolumesClonePostNotFound{}
}

/*
PcloudVolumesClonePostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVolumesClonePostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumes clone post not found response has a 2xx status code
func (o *PcloudVolumesClonePostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumes clone post not found response has a 3xx status code
func (o *PcloudVolumesClonePostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumes clone post not found response has a 4xx status code
func (o *PcloudVolumesClonePostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumes clone post not found response has a 5xx status code
func (o *PcloudVolumesClonePostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumes clone post not found response a status code equal to that given
func (o *PcloudVolumesClonePostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud volumes clone post not found response
func (o *PcloudVolumesClonePostNotFound) Code() int {
	return 404
}

func (o *PcloudVolumesClonePostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostNotFound %s", 404, payload)
}

func (o *PcloudVolumesClonePostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostNotFound %s", 404, payload)
}

func (o *PcloudVolumesClonePostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumesClonePostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumesClonePostConflict creates a PcloudVolumesClonePostConflict with default headers values
func NewPcloudVolumesClonePostConflict() *PcloudVolumesClonePostConflict {
	return &PcloudVolumesClonePostConflict{}
}

/*
PcloudVolumesClonePostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudVolumesClonePostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumes clone post conflict response has a 2xx status code
func (o *PcloudVolumesClonePostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumes clone post conflict response has a 3xx status code
func (o *PcloudVolumesClonePostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumes clone post conflict response has a 4xx status code
func (o *PcloudVolumesClonePostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumes clone post conflict response has a 5xx status code
func (o *PcloudVolumesClonePostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumes clone post conflict response a status code equal to that given
func (o *PcloudVolumesClonePostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud volumes clone post conflict response
func (o *PcloudVolumesClonePostConflict) Code() int {
	return 409
}

func (o *PcloudVolumesClonePostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostConflict %s", 409, payload)
}

func (o *PcloudVolumesClonePostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostConflict %s", 409, payload)
}

func (o *PcloudVolumesClonePostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumesClonePostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumesClonePostInternalServerError creates a PcloudVolumesClonePostInternalServerError with default headers values
func NewPcloudVolumesClonePostInternalServerError() *PcloudVolumesClonePostInternalServerError {
	return &PcloudVolumesClonePostInternalServerError{}
}

/*
PcloudVolumesClonePostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVolumesClonePostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumes clone post internal server error response has a 2xx status code
func (o *PcloudVolumesClonePostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumes clone post internal server error response has a 3xx status code
func (o *PcloudVolumesClonePostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumes clone post internal server error response has a 4xx status code
func (o *PcloudVolumesClonePostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volumes clone post internal server error response has a 5xx status code
func (o *PcloudVolumesClonePostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud volumes clone post internal server error response a status code equal to that given
func (o *PcloudVolumesClonePostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud volumes clone post internal server error response
func (o *PcloudVolumesClonePostInternalServerError) Code() int {
	return 500
}

func (o *PcloudVolumesClonePostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostInternalServerError %s", 500, payload)
}

func (o *PcloudVolumesClonePostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostInternalServerError %s", 500, payload)
}

func (o *PcloudVolumesClonePostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumesClonePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
