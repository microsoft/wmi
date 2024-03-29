// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_LogManagesRecord struct
type CIM_LogManagesRecord struct {
	*cim.WmiInstance

	//
	Log CIM_Log

	//
	Record CIM_RecordForLog
}

func NewCIM_LogManagesRecordEx1(instance *cim.WmiInstance) (newInstance *CIM_LogManagesRecord, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_LogManagesRecord{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_LogManagesRecordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_LogManagesRecord, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_LogManagesRecord{
		WmiInstance: tmp,
	}
	return
}

// SetLog sets the value of Log for the instance
func (instance *CIM_LogManagesRecord) SetPropertyLog(value CIM_Log) (err error) {
	return instance.SetProperty("Log", (value))
}

// GetLog gets the value of Log for the instance
func (instance *CIM_LogManagesRecord) GetPropertyLog() (value CIM_Log, err error) {
	retValue, err := instance.GetProperty("Log")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_Log)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_Log is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_Log(valuetmp)

	return
}

// SetRecord sets the value of Record for the instance
func (instance *CIM_LogManagesRecord) SetPropertyRecord(value CIM_RecordForLog) (err error) {
	return instance.SetProperty("Record", (value))
}

// GetRecord gets the value of Record for the instance
func (instance *CIM_LogManagesRecord) GetPropertyRecord() (value CIM_RecordForLog, err error) {
	retValue, err := instance.GetProperty("Record")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_RecordForLog)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_RecordForLog is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_RecordForLog(valuetmp)

	return
}
