package errors

import "errors"

var (
	ErrNilMetrics      = errors.New("metrics is nil")
	ErrUnknownType     = errors.New("unknown metric type")
	ErrCounterNoDelta  = errors.New("counter metric requires delta")
	ErrCounterHasValue = errors.New("counter metric must not have value")
	ErrGaugeNoValue    = errors.New("gauge metric requires value")
	ErrGaugeHasDelta   = errors.New("gauge metric must not have delta")
)
