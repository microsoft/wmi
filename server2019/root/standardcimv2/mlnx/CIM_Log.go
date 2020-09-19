// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_Log struct
type CIM_Log struct {
	*CIM_EnabledLogicalElement

	//
	CurrentNumberOfRecords uint64

	//
	LogState Log_LogState

	//
	MaxNumberOfRecords uint64

	//
	OverwritePolicy Log_OverwritePolicy
}

func NewCIM_LogEx1(instance *cim.WmiInstance) (newInstance *CIM_Log, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Log{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

func NewCIM_LogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Log, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Log{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

// SetCurrentNumberOfRecords sets the value of CurrentNumberOfRecords for the instance
func (instance *CIM_Log) SetPropertyCurrentNumberOfRecords(value uint64) (err error) {
	return instance.SetProperty("CurrentNumberOfRecords", (value))
}

// GetCurrentNumberOfRecords gets the value of CurrentNumberOfRecords for the instance
func (instance *CIM_Log) GetPropertyCurrentNumberOfRecords() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentNumberOfRecords")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetLogState sets the value of LogState for the instance
func (instance *CIM_Log) SetPropertyLogState(value Log_LogState) (err error) {
	return instance.SetProperty("LogState", (value))
}

// GetLogState gets the value of LogState for the instance
func (instance *CIM_Log) GetPropertyLogState() (value Log_LogState, err error) {
	retValue, err := instance.GetProperty("LogState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Log_LogState(valuetmp)

	return
}

// SetMaxNumberOfRecords sets the value of MaxNumberOfRecords for the instance
func (instance *CIM_Log) SetPropertyMaxNumberOfRecords(value uint64) (err error) {
	return instance.SetProperty("MaxNumberOfRecords", (value))
}

// GetMaxNumberOfRecords gets the value of MaxNumberOfRecords for the instance
func (instance *CIM_Log) GetPropertyMaxNumberOfRecords() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxNumberOfRecords")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetOverwritePolicy sets the value of OverwritePolicy for the instance
func (instance *CIM_Log) SetPropertyOverwritePolicy(value Log_OverwritePolicy) (err error) {
	return instance.SetProperty("OverwritePolicy", (value))
}

// GetOverwritePolicy gets the value of OverwritePolicy for the instance
func (instance *CIM_Log) GetPropertyOverwritePolicy() (value Log_OverwritePolicy, err error) {
	retValue, err := instance.GetProperty("OverwritePolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Log_OverwritePolicy(valuetmp)

	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Log) ClearLog() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ClearLog")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
