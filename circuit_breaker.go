package circuit_breaker

import "errors"

var errNotAllowed = errors.New("circuitbreaker: not allowed for circuit open")

func IsErrNotAllowed(err error) bool {
	return errors.Is(err, errNotAllowed)
}

func ErrNotAllowed() error {
	return errNotAllowed
}

// CircuitBreaker is a circuit breaker.
type CircuitBreaker interface {
	Allow() error
	MarkSuccess()
	MarkFailed()
}
