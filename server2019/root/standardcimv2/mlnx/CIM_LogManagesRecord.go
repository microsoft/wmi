// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_LogManagesRecord struct
type CIM_LogManagesRecord struct {
	cim.WmiInstance

	//
	Log CIM_Log

	//
	Record CIM_RecordForLog
}

// SetLog sets the value of Log for the instance
func (instance *CIM_LogManagesRecord) SetPropertyLog(value CIM_Log) (err error) {
	return instance.SetProperty("Log", value)
}

// GetLog gets the value of Log for the instance
func (instance *CIM_LogManagesRecord) GetPropertyLog() (value CIM_Log, err error) {
	retValue, err := instance.GetProperty("Log")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Log)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecord sets the value of Record for the instance
func (instance *CIM_LogManagesRecord) SetPropertyRecord(value CIM_RecordForLog) (err error) {
	return instance.SetProperty("Record", value)
}

// GetRecord gets the value of Record for the instance
func (instance *CIM_LogManagesRecord) GetPropertyRecord() (value CIM_RecordForLog, err error) {
	retValue, err := instance.GetProperty("Record")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_RecordForLog)
	if !ok {
		// TODO: Set an error
	}
	return
}
