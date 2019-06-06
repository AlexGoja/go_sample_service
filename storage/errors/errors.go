package errors

import "errors"

var (
	ErrCreateTransaction = errors.New("Error creating a transaction")
	ErrQueryExecution    = errors.New("Error executing a query")
	ErrScanRows          = errors.New("Error scanning rows")
	ErrMarshallData      = errors.New("Error marshalling entity data")
)
