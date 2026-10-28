// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.Intel_ME
//////////////////////////////////////////////
package intel_me

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_LogEntry struct
type CIM_LogEntry struct {
	*CIM_RecordForLog

	// A LogEntry may include a timestamp for the entry.
	CreationTimeStamp string

	// The containing Log's InstanceID.
	LogInstanceID string

	// The containing Log's Name. This property is available for backwards continuity with CIM_LogRecord.
	LogName string

	// RecordID may be used to provide a representation of log entry ordering or pointers/handles for log entries.
	RecordID string
}

func NewCIM_LogEntryEx1(instance *cim.WmiInstance) (newInstance *CIM_LogEntry, err error) {
	tmp, err := NewCIM_RecordForLogEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_LogEntry{
		CIM_RecordForLog: tmp,
	}
	return
}

func NewCIM_LogEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_LogEntry, err error) {
	tmp, err := NewCIM_RecordForLogEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_LogEntry{
		CIM_RecordForLog: tmp,
	}
	return
}

// SetCreationTimeStamp sets the value of CreationTimeStamp for the instance
func (instance *CIM_LogEntry) SetPropertyCreationTimeStamp(value string) (err error) {
	return instance.SetProperty("CreationTimeStamp", (value))
}

// GetCreationTimeStamp gets the value of CreationTimeStamp for the instance
func (instance *CIM_LogEntry) GetPropertyCreationTimeStamp() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTimeStamp")
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

// SetLogInstanceID sets the value of LogInstanceID for the instance
func (instance *CIM_LogEntry) SetPropertyLogInstanceID(value string) (err error) {
	return instance.SetProperty("LogInstanceID", (value))
}

// GetLogInstanceID gets the value of LogInstanceID for the instance
func (instance *CIM_LogEntry) GetPropertyLogInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("LogInstanceID")
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

// SetLogName sets the value of LogName for the instance
func (instance *CIM_LogEntry) SetPropertyLogName(value string) (err error) {
	return instance.SetProperty("LogName", (value))
}

// GetLogName gets the value of LogName for the instance
func (instance *CIM_LogEntry) GetPropertyLogName() (value string, err error) {
	retValue, err := instance.GetProperty("LogName")
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

// SetRecordID sets the value of RecordID for the instance
func (instance *CIM_LogEntry) SetPropertyRecordID(value string) (err error) {
	return instance.SetProperty("RecordID", (value))
}

// GetRecordID gets the value of RecordID for the instance
func (instance *CIM_LogEntry) GetPropertyRecordID() (value string, err error) {
	retValue, err := instance.GetProperty("RecordID")
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
