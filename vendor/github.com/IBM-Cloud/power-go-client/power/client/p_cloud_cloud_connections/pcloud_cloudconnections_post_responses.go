// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudconnectionsPostReader is a Reader for the PcloudCloudconnectionsPost structure.
type PcloudCloudconnectionsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudconnectionsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPcloudCloudconnectionsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPcloudCloudconnectionsPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudconnectionsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudconnectionsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudconnectionsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudconnectionsPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPcloudCloudconnectionsPostRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudCloudconnectionsPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudCloudconnectionsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudconnectionsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPcloudCloudconnectionsPostServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPcloudCloudconnectionsPostGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections] pcloud.cloudconnections.post", response, response.Code())
	}
}

// NewPcloudCloudconnectionsPostOK creates a PcloudCloudconnectionsPostOK with default headers values
func NewPcloudCloudconnectionsPostOK() *PcloudCloudconnectionsPostOK {
	return &PcloudCloudconnectionsPostOK{}
}

/*
PcloudCloudconnectionsPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudconnectionsPostOK struct {
	Payload *models.CloudConnection
}

// IsSuccess returns true when this pcloud cloudconnections post o k response has a 2xx status code
func (o *PcloudCloudconnectionsPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudconnections post o k response has a 3xx status code
func (o *PcloudCloudconnectionsPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections post o k response has a 4xx status code
func (o *PcloudCloudconnectionsPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections post o k response has a 5xx status code
func (o *PcloudCloudconnectionsPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections post o k response a status code equal to that given
func (o *PcloudCloudconnectionsPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud cloudconnections post o k response
func (o *PcloudCloudconnectionsPostOK) Code() int {
	return 200
}

func (o *PcloudCloudconnectionsPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudconnectionsPostOK) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudconnectionsPostOK) GetPayload() *models.CloudConnection {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostCreated creates a PcloudCloudconnectionsPostCreated with default headers values
func NewPcloudCloudconnectionsPostCreated() *PcloudCloudconnectionsPostCreated {
	return &PcloudCloudconnectionsPostCreated{}
}

/*
PcloudCloudconnectionsPostCreated describes a response with status code 201, with default header values.

Created
*/
type PcloudCloudconnectionsPostCreated struct {
	Payload *models.CloudConnection
}

// IsSuccess returns true when this pcloud cloudconnections post created response has a 2xx status code
func (o *PcloudCloudconnectionsPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudconnections post created response has a 3xx status code
func (o *PcloudCloudconnectionsPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections post created response has a 4xx status code
func (o *PcloudCloudconnectionsPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections post created response has a 5xx status code
func (o *PcloudCloudconnectionsPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections post created response a status code equal to that given
func (o *PcloudCloudconnectionsPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the pcloud cloudconnections post created response
func (o *PcloudCloudconnectionsPostCreated) Code() int {
	return 201
}

func (o *PcloudCloudconnectionsPostCreated) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostCreated  %+v", 201, o.Payload)
}

func (o *PcloudCloudconnectionsPostCreated) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostCreated  %+v", 201, o.Payload)
}

func (o *PcloudCloudconnectionsPostCreated) GetPayload() *models.CloudConnection {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostAccepted creates a PcloudCloudconnectionsPostAccepted with default headers values
func NewPcloudCloudconnectionsPostAccepted() *PcloudCloudconnectionsPostAccepted {
	return &PcloudCloudconnectionsPostAccepted{}
}

/*
PcloudCloudconnectionsPostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudCloudconnectionsPostAccepted struct {
	Payload *models.CloudConnectionCreateResponse
}

// IsSuccess returns true when this pcloud cloudconnections post accepted response has a 2xx status code
func (o *PcloudCloudconnectionsPostAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudconnections post accepted response has a 3xx status code
func (o *PcloudCloudconnectionsPostAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections post accepted response has a 4xx status code
func (o *PcloudCloudconnectionsPostAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections post accepted response has a 5xx status code
func (o *PcloudCloudconnectionsPostAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections post accepted response a status code equal to that given
func (o *PcloudCloudconnectionsPostAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud cloudconnections post accepted response
func (o *PcloudCloudconnectionsPostAccepted) Code() int {
	return 202
}

func (o *PcloudCloudconnectionsPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudCloudconnectionsPostAccepted) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudCloudconnectionsPostAccepted) GetPayload() *models.CloudConnectionCreateResponse {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnectionCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostBadRequest creates a PcloudCloudconnectionsPostBadRequest with default headers values
func NewPcloudCloudconnectionsPostBadRequest() *PcloudCloudconnectionsPostBadRequest {
	return &PcloudCloudconnectionsPostBadRequest{}
}

/*
PcloudCloudconnectionsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudconnectionsPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections post bad request response has a 2xx status code
func (o *PcloudCloudconnectionsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections post bad request response has a 3xx status code
func (o *PcloudCloudconnectionsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections post bad request response has a 4xx status code
func (o *PcloudCloudconnectionsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections post bad request response has a 5xx status code
func (o *PcloudCloudconnectionsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections post bad request response a status code equal to that given
func (o *PcloudCloudconnectionsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudconnections post bad request response
func (o *PcloudCloudconnectionsPostBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudconnectionsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudconnectionsPostBadRequest) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudconnectionsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostUnauthorized creates a PcloudCloudconnectionsPostUnauthorized with default headers values
func NewPcloudCloudconnectionsPostUnauthorized() *PcloudCloudconnectionsPostUnauthorized {
	return &PcloudCloudconnectionsPostUnauthorized{}
}

/*
PcloudCloudconnectionsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudconnectionsPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections post unauthorized response has a 2xx status code
func (o *PcloudCloudconnectionsPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections post unauthorized response has a 3xx status code
func (o *PcloudCloudconnectionsPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections post unauthorized response has a 4xx status code
func (o *PcloudCloudconnectionsPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections post unauthorized response has a 5xx status code
func (o *PcloudCloudconnectionsPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections post unauthorized response a status code equal to that given
func (o *PcloudCloudconnectionsPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudconnections post unauthorized response
func (o *PcloudCloudconnectionsPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudconnectionsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudconnectionsPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudconnectionsPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostForbidden creates a PcloudCloudconnectionsPostForbidden with default headers values
func NewPcloudCloudconnectionsPostForbidden() *PcloudCloudconnectionsPostForbidden {
	return &PcloudCloudconnectionsPostForbidden{}
}

/*
PcloudCloudconnectionsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudconnectionsPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections post forbidden response has a 2xx status code
func (o *PcloudCloudconnectionsPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections post forbidden response has a 3xx status code
func (o *PcloudCloudconnectionsPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections post forbidden response has a 4xx status code
func (o *PcloudCloudconnectionsPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections post forbidden response has a 5xx status code
func (o *PcloudCloudconnectionsPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections post forbidden response a status code equal to that given
func (o *PcloudCloudconnectionsPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud cloudconnections post forbidden response
func (o *PcloudCloudconnectionsPostForbidden) Code() int {
	return 403
}

func (o *PcloudCloudconnectionsPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudconnectionsPostForbidden) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudconnectionsPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostNotFound creates a PcloudCloudconnectionsPostNotFound with default headers values
func NewPcloudCloudconnectionsPostNotFound() *PcloudCloudconnectionsPostNotFound {
	return &PcloudCloudconnectionsPostNotFound{}
}

/*
PcloudCloudconnectionsPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudconnectionsPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections post not found response has a 2xx status code
func (o *PcloudCloudconnectionsPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections post not found response has a 3xx status code
func (o *PcloudCloudconnectionsPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections post not found response has a 4xx status code
func (o *PcloudCloudconnectionsPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections post not found response has a 5xx status code
func (o *PcloudCloudconnectionsPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections post not found response a status code equal to that given
func (o *PcloudCloudconnectionsPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud cloudconnections post not found response
func (o *PcloudCloudconnectionsPostNotFound) Code() int {
	return 404
}

func (o *PcloudCloudconnectionsPostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudconnectionsPostNotFound) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudconnectionsPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostRequestTimeout creates a PcloudCloudconnectionsPostRequestTimeout with default headers values
func NewPcloudCloudconnectionsPostRequestTimeout() *PcloudCloudconnectionsPostRequestTimeout {
	return &PcloudCloudconnectionsPostRequestTimeout{}
}

/*
PcloudCloudconnectionsPostRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PcloudCloudconnectionsPostRequestTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections post request timeout response has a 2xx status code
func (o *PcloudCloudconnectionsPostRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections post request timeout response has a 3xx status code
func (o *PcloudCloudconnectionsPostRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections post request timeout response has a 4xx status code
func (o *PcloudCloudconnectionsPostRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections post request timeout response has a 5xx status code
func (o *PcloudCloudconnectionsPostRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections post request timeout response a status code equal to that given
func (o *PcloudCloudconnectionsPostRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the pcloud cloudconnections post request timeout response
func (o *PcloudCloudconnectionsPostRequestTimeout) Code() int {
	return 408
}

func (o *PcloudCloudconnectionsPostRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostRequestTimeout  %+v", 408, o.Payload)
}

func (o *PcloudCloudconnectionsPostRequestTimeout) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostRequestTimeout  %+v", 408, o.Payload)
}

func (o *PcloudCloudconnectionsPostRequestTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostConflict creates a PcloudCloudconnectionsPostConflict with default headers values
func NewPcloudCloudconnectionsPostConflict() *PcloudCloudconnectionsPostConflict {
	return &PcloudCloudconnectionsPostConflict{}
}

/*
PcloudCloudconnectionsPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudCloudconnectionsPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections post conflict response has a 2xx status code
func (o *PcloudCloudconnectionsPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections post conflict response has a 3xx status code
func (o *PcloudCloudconnectionsPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections post conflict response has a 4xx status code
func (o *PcloudCloudconnectionsPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections post conflict response has a 5xx status code
func (o *PcloudCloudconnectionsPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections post conflict response a status code equal to that given
func (o *PcloudCloudconnectionsPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud cloudconnections post conflict response
func (o *PcloudCloudconnectionsPostConflict) Code() int {
	return 409
}

func (o *PcloudCloudconnectionsPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudCloudconnectionsPostConflict) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudCloudconnectionsPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostUnprocessableEntity creates a PcloudCloudconnectionsPostUnprocessableEntity with default headers values
func NewPcloudCloudconnectionsPostUnprocessableEntity() *PcloudCloudconnectionsPostUnprocessableEntity {
	return &PcloudCloudconnectionsPostUnprocessableEntity{}
}

/*
PcloudCloudconnectionsPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudCloudconnectionsPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections post unprocessable entity response has a 2xx status code
func (o *PcloudCloudconnectionsPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections post unprocessable entity response has a 3xx status code
func (o *PcloudCloudconnectionsPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections post unprocessable entity response has a 4xx status code
func (o *PcloudCloudconnectionsPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections post unprocessable entity response has a 5xx status code
func (o *PcloudCloudconnectionsPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections post unprocessable entity response a status code equal to that given
func (o *PcloudCloudconnectionsPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud cloudconnections post unprocessable entity response
func (o *PcloudCloudconnectionsPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudCloudconnectionsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudCloudconnectionsPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudCloudconnectionsPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostInternalServerError creates a PcloudCloudconnectionsPostInternalServerError with default headers values
func NewPcloudCloudconnectionsPostInternalServerError() *PcloudCloudconnectionsPostInternalServerError {
	return &PcloudCloudconnectionsPostInternalServerError{}
}

/*
PcloudCloudconnectionsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections post internal server error response has a 2xx status code
func (o *PcloudCloudconnectionsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections post internal server error response has a 3xx status code
func (o *PcloudCloudconnectionsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections post internal server error response has a 4xx status code
func (o *PcloudCloudconnectionsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections post internal server error response has a 5xx status code
func (o *PcloudCloudconnectionsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudconnections post internal server error response a status code equal to that given
func (o *PcloudCloudconnectionsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudconnections post internal server error response
func (o *PcloudCloudconnectionsPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudconnectionsPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudconnectionsPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudconnectionsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostServiceUnavailable creates a PcloudCloudconnectionsPostServiceUnavailable with default headers values
func NewPcloudCloudconnectionsPostServiceUnavailable() *PcloudCloudconnectionsPostServiceUnavailable {
	return &PcloudCloudconnectionsPostServiceUnavailable{}
}

/*
PcloudCloudconnectionsPostServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable
*/
type PcloudCloudconnectionsPostServiceUnavailable struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections post service unavailable response has a 2xx status code
func (o *PcloudCloudconnectionsPostServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections post service unavailable response has a 3xx status code
func (o *PcloudCloudconnectionsPostServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections post service unavailable response has a 4xx status code
func (o *PcloudCloudconnectionsPostServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections post service unavailable response has a 5xx status code
func (o *PcloudCloudconnectionsPostServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudconnections post service unavailable response a status code equal to that given
func (o *PcloudCloudconnectionsPostServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the pcloud cloudconnections post service unavailable response
func (o *PcloudCloudconnectionsPostServiceUnavailable) Code() int {
	return 503
}

func (o *PcloudCloudconnectionsPostServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PcloudCloudconnectionsPostServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PcloudCloudconnectionsPostServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsPostGatewayTimeout creates a PcloudCloudconnectionsPostGatewayTimeout with default headers values
func NewPcloudCloudconnectionsPostGatewayTimeout() *PcloudCloudconnectionsPostGatewayTimeout {
	return &PcloudCloudconnectionsPostGatewayTimeout{}
}

/*
PcloudCloudconnectionsPostGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout
*/
type PcloudCloudconnectionsPostGatewayTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections post gateway timeout response has a 2xx status code
func (o *PcloudCloudconnectionsPostGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections post gateway timeout response has a 3xx status code
func (o *PcloudCloudconnectionsPostGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections post gateway timeout response has a 4xx status code
func (o *PcloudCloudconnectionsPostGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections post gateway timeout response has a 5xx status code
func (o *PcloudCloudconnectionsPostGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudconnections post gateway timeout response a status code equal to that given
func (o *PcloudCloudconnectionsPostGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the pcloud cloudconnections post gateway timeout response
func (o *PcloudCloudconnectionsPostGatewayTimeout) Code() int {
	return 504
}

func (o *PcloudCloudconnectionsPostGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PcloudCloudconnectionsPostGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsPostGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PcloudCloudconnectionsPostGatewayTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsPostGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
