package handlers

import (
	"encoding/json"
	"net/http"
)

// Define a simple JSON error response
type ErrorResponse struct {
	Message string `json:"message"`
}

type ErrorHandler struct {
}

func NewErrorHandler() ErrorHandler {
	return ErrorHandler{}
}

// The server cannot find the requested resource.
func (errorHandler *ErrorHandler) NotFound(w http.ResponseWriter, message string) {
	errorHandler.write(w, message, http.StatusNotFound)
}

// This error response means that the server, while working as a gateway to
// get a response needed to handle the request, got an invalid response.
func (errorHandler *ErrorHandler) BadGateway(w http.ResponseWriter, message string) {
	errorHandler.write(w, message, http.StatusBadGateway)
}

// The server cannot or will not process the request due to something
// that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).
func (errorHandler *ErrorHandler) BadRequest(w http.ResponseWriter, message string) {
	errorHandler.write(w, message, http.StatusBadRequest)
}

// The HTTP 405 Method Not Allowed.
func (errorHandler *ErrorHandler) MethodNotAllowed(w http.ResponseWriter) {
	errorHandler.write(w, "The HTTP 405 Method Not Allowed", http.StatusMethodNotAllowed)
}

// The server has encountered a situation it does not know how to handle.
func (errorHandler *ErrorHandler) InternalServerError(w http.ResponseWriter) {
	errorHandler.write(w, "The server has encountered a situation it does not know how to handle", http.StatusInternalServerError)
}

func (errorHandler *ErrorHandler) write(w http.ResponseWriter, message string, errCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}
