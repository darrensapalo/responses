package responses

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// WriteErrorResponse is a convenience function for writing error responses
func WriteErrorResponse(res http.ResponseWriter, status int, jsonErrs ...JSONError) {
	res.WriteHeader(status)
	if err := json.NewEncoder(res).Encode(jsonErrs); err != nil {
		log.WithError(err).Error("unable to send error response, the JSON error cound not be encoded")
	}
}

// WriteResponse is a convenience function for writing responses
func WriteResponse(res http.ResponseWriter, status int, entity interface{}) {
	res.WriteHeader(status)
	if err := json.NewEncoder(res).Encode(entity); err != nil {
		log.WithError(err).Error("unable to send response, the entity could not be encoded")
	}
}

// WriteOK writes an ok response without a body
func WriteOK(res http.ResponseWriter) {
	res.WriteHeader(http.StatusOK)
}

// WriteOKWithEntity writes an ok response with the provided entity
func WriteOKWithEntity(res http.ResponseWriter, entity interface{}) {
	WriteResponse(res, http.StatusOK, entity)
}

// WriteCreated writes a CREATED response with no entity
func WriteCreated(res http.ResponseWriter) {
	res.WriteHeader(http.StatusCreated)
}

// WriteGatewayTimeout is when an external gateway cannot be accessed
func WriteGatewayTimeout(res http.ResponseWriter) {
	status := http.StatusGatewayTimeout
	err := ErrorBuilder().
		Status(status).
		Title("Gateway Timeout").
		Detail("An external resource is either inaccessible or down").
		Build()
	WriteErrorResponse(res, status, err)
}

// WriteUnknownError is when you don't know what the error is, this shouldn't be the case and is a bug in the code
func WriteUnknownError(res http.ResponseWriter) {
	status := http.StatusInternalServerError
	err := ErrorBuilder().
		Status(status).
		Title("Unknown Error").
		Detail("An unknown error has occurred, contact the system administrator").
		Build()
	WriteErrorResponse(res, status, err)
}

// WriteUnreadableRequestError is when the request body cannot be pased into JSON
func WriteUnreadableRequestError(res http.ResponseWriter) {
	status := http.StatusBadRequest
	jsonErr := ErrorBuilder().
		Status(status).
		Title("Request cannot be read").
		Detail("Request is not a valid JSON format").
		SourcePointer("/").
		Build()
	WriteErrorResponse(res, status, jsonErr)
}

// WriteUnauthorizedError is when the access to the endpoint is unauthorized
func WriteUnauthorizedError(res http.ResponseWriter) {
	status := http.StatusUnauthorized
	jsonErr := ErrorBuilder().
		Status(status).
		Title("Unauthorized").
		Detail("Access to the endpoint requires authorization").
		Build()
	WriteErrorResponse(res, status, jsonErr)
}
