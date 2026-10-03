// Package user contains the user domain: model, persistence, service and
// HTTP handlers. This file defines the domain errors.
package user

import "errors"

// NotFoundMessage is the exact detail message the source service raised
// (UserServiceImp.fetchUserById) and which the error handler echoes in the
// 404 response body. Kept verbatim (including grammar) for wire compatibility.
const NotFoundMessage = "User are not available"

// ErrNotFound is the sentinel returned when a requested user does not exist.
// Callers should wrap it with fmt.Errorf("...: %w", ErrNotFound) or return it
// directly, and detect it with errors.Is(err, user.ErrNotFound). Its message is
// exactly what the source put in the 404 body.
var ErrNotFound = errors.New(NotFoundMessage)

// NotFoundError is the Go equivalent of com.smartContact.error.UserNotFoundException
// for callers that need a custom detail message and/or an underlying cause.
// Any *NotFoundError satisfies errors.Is(err, ErrNotFound), so HTTP layers
// only need to check the sentinel to map it to 404.
type NotFoundError struct {
	// Message is the detail message (Java getMessage()). May be empty.
	Message string
	// Cause is the underlying error (Java getCause()). May be nil.
	Cause error
}

// NewNotFoundError constructs a NotFoundError with no detail message and no
// cause (Java: new UserNotFoundException()).
func NewNotFoundError() *NotFoundError {
	return &NotFoundError{}
}

// NewNotFoundErrorMessage constructs a NotFoundError with the given detail
// message and no cause (Java: new UserNotFoundException(message)).
func NewNotFoundErrorMessage(message string) *NotFoundError {
	return &NotFoundError{Message: message}
}

// WrapNotFoundError constructs a NotFoundError with the given detail message
// and cause (Java: new UserNotFoundException(message, cause)).
func WrapNotFoundError(message string, cause error) *NotFoundError {
	return &NotFoundError{Message: message, Cause: cause}
}

// NotFoundErrorFromCause constructs a NotFoundError whose detail message is
// derived from the cause (Java: new UserNotFoundException(cause), where the
// message becomes cause.toString(), or null when cause is null).
//
// MIGRATION_NOTE: Java's Throwable.toString() is "<ClassName>: <message>";
// Go errors have no class name, so the cause's Error() text is used instead.
func NotFoundErrorFromCause(cause error) *NotFoundError {
	e := &NotFoundError{Cause: cause}
	if cause != nil {
		e.Message = cause.Error()
	}
	return e
}

// MIGRATION_NOTE: the protected constructor
// UserNotFoundException(String, Throwable, boolean enableSuppression,
// boolean writableStackTrace) has no Go equivalent: Go errors carry neither
// suppressed exceptions nor stack traces. WrapNotFoundError covers the
// message+cause part; the two boolean flags are intentionally dropped.

// Error implements the error interface and returns the detail message.
// When no message was supplied it returns an empty string, mirroring Java's
// null getMessage() as closely as Go allows.
func (e *NotFoundError) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}

// Unwrap returns the underlying cause so errors.Is/errors.As can traverse it.
func (e *NotFoundError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Cause
}

// Is reports whether target is ErrNotFound, so every NotFoundError is
// classified as "user not found" regardless of its custom message.
func (e *NotFoundError) Is(target error) bool {
	return target == ErrNotFound
}
