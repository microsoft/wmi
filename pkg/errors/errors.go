// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package errors

import (
	"errors"
	"fmt"

	mocerror "github.com/microsoft/moc/pkg/errors"
	perrors "github.com/pkg/errors"
)

const (
	wmiError                = "WMI Error 0x"
	OutOfMemoryErrorSummary = "Could not initialize memory: Not enough memory resources are available to complete this operation. (0x8007000E)"
)

var (
	ERROR_OUTOFMEMORY uint16 = 0xe
)

var (
	NotFound       error = errors.New("Not Found")
	Timedout       error = errors.New("Timedout")
	InvalidInput   error = errors.New("Invalid Input")
	InvalidType    error = errors.New("Invalid Type")
	NotSupported   error = errors.New("Not Supported")
	AlreadyExists  error = errors.New("Already Exists")
	InvalidFilter  error = errors.New("Invalid Filter")
	Failed         error = errors.New("Failed")
	NotImplemented error = errors.New("Not Implemented")
	Unknown        error = errors.New("Unknown Reason")
)

func Wrap(cause error, message string) error {
	return perrors.Wrap(cause, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return perrors.Wrapf(err, format, args...)
}

func IsNotFound(err error) bool {
	return checkError(err, NotFound)
}
func IsAlreadyExists(err error) bool {
	return checkError(err, AlreadyExists)
}
func IsTimedout(err error) bool {
	return checkError(err, Timedout)
}
func IsInvalidInput(err error) bool {
	return checkError(err, InvalidInput)
}
func IsInvalidType(err error) bool {
	return checkError(err, InvalidType)
}
func IsNotSupported(err error) bool {
	return checkError(err, NotSupported)
}
func IsInvalidFilter(err error) bool {
	return checkError(err, InvalidFilter)
}
func IsFailed(err error) bool {
	return checkError(err, Failed)
}
func IsNotImplemented(err error) bool {
	return checkError(err, NotImplemented)
}
func IsUnknown(err error) bool {
	return checkError(err, Unknown)
}

func checkError(wrappedError, err error) bool {
	if wrappedError == nil {
		return false
	}
	if wrappedError == err {
		return true
	}
	cerr := perrors.Cause(wrappedError)
	if cerr != nil && cerr == err {
		return true
	}

	return false

}

func New(errString string) error {
	return errors.New(errString)
}

func NewWMIError(errorCode uint16) error {
	switch errorCode {
	case ERROR_OUTOFMEMORY:
		return mocerror.OutOfMemory
	default:
		return fmt.Errorf(wmiError+"%08x", errorCode)
	}
}
