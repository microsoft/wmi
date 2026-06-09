// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package errors

import (
	"errors"
	"fmt"
	"strings"
)

const (
	wmiError = "WMI Error 0x"
)

var (
	NotFound         error = errors.New("Not Found")
	Timedout         error = errors.New("Timedout")
	InvalidInput     error = errors.New("Invalid Input")
	InvalidType      error = errors.New("Invalid Type")
	NotSupported     error = errors.New("Not Supported")
	AlreadyExists    error = errors.New("Already Exists")
	InvalidFilter    error = errors.New("Invalid Filter")
	Failed           error = errors.New("Failed")
	NotImplemented   error = errors.New("Not Implemented")
	Unknown          error = errors.New("Unknown Reason")
	DvdDriveNotFound error = errors.New("DVDDriveNotFound")
)

func Wrap(cause error, message string) error {
	if cause == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, cause)
}

func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
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
func IsDvdDriveNotFound(err error) bool {
	return checkError(err, DvdDriveNotFound)
}

func IsWMIError(err error) bool {
	for e := err; e != nil; e = errors.Unwrap(e) {
		if strings.HasPrefix(e.Error(), wmiError) {
			return true
		}
	}
	return false
}

func checkError(wrappedError, err error) bool {
	if wrappedError == nil {
		return false
	}
	return errors.Is(wrappedError, err)
}

func New(errString string) error {
	return errors.New(errString)
}

func NewWMIError(errorCode uint16) error {
	return fmt.Errorf(wmiError+"%08x", errorCode)
}
