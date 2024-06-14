// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// TraceFailedRequestsLogging struct
type TraceFailedRequestsLogging struct {
	*EmbeddedObject

	//
	CustomActionsEnabled bool

	//
	Directory string

	//
	Enabled bool

	//
	MaxLogFiles uint32

	//
	MaxLogFileSizeKB uint32
}

func NewTraceFailedRequestsLoggingEx1(instance *cim.WmiInstance) (newInstance *TraceFailedRequestsLogging, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TraceFailedRequestsLogging{
		EmbeddedObject: tmp,
	}
	return
}

func NewTraceFailedRequestsLoggingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TraceFailedRequestsLogging, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TraceFailedRequestsLogging{
		EmbeddedObject: tmp,
	}
	return
}

// SetCustomActionsEnabled sets the value of CustomActionsEnabled for the instance
func (instance *TraceFailedRequestsLogging) SetPropertyCustomActionsEnabled(value bool) (err error) {
	return instance.SetProperty("CustomActionsEnabled", (value))
}

// GetCustomActionsEnabled gets the value of CustomActionsEnabled for the instance
func (instance *TraceFailedRequestsLogging) GetPropertyCustomActionsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("CustomActionsEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDirectory sets the value of Directory for the instance
func (instance *TraceFailedRequestsLogging) SetPropertyDirectory(value string) (err error) {
	return instance.SetProperty("Directory", (value))
}

// GetDirectory gets the value of Directory for the instance
func (instance *TraceFailedRequestsLogging) GetPropertyDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("Directory")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *TraceFailedRequestsLogging) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *TraceFailedRequestsLogging) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetMaxLogFiles sets the value of MaxLogFiles for the instance
func (instance *TraceFailedRequestsLogging) SetPropertyMaxLogFiles(value uint32) (err error) {
	return instance.SetProperty("MaxLogFiles", (value))
}

// GetMaxLogFiles gets the value of MaxLogFiles for the instance
func (instance *TraceFailedRequestsLogging) GetPropertyMaxLogFiles() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxLogFiles")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetMaxLogFileSizeKB sets the value of MaxLogFileSizeKB for the instance
func (instance *TraceFailedRequestsLogging) SetPropertyMaxLogFileSizeKB(value uint32) (err error) {
	return instance.SetProperty("MaxLogFileSizeKB", (value))
}

// GetMaxLogFileSizeKB gets the value of MaxLogFileSizeKB for the instance
func (instance *TraceFailedRequestsLogging) GetPropertyMaxLogFileSizeKB() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxLogFileSizeKB")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
