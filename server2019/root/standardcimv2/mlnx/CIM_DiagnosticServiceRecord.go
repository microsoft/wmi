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

// CIM_DiagnosticServiceRecord struct
type CIM_DiagnosticServiceRecord struct {
	*CIM_DiagnosticRecord

	//
	ErrorCode []string

	//
	ErrorCount []uint32

	//
	LoopsFailed uint32

	//
	LoopsPassed uint32
}

func NewCIM_DiagnosticServiceRecordEx1(instance *cim.WmiInstance) (newInstance *CIM_DiagnosticServiceRecord, err error) {
	tmp, err := NewCIM_DiagnosticRecordEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticServiceRecord{
		CIM_DiagnosticRecord: tmp,
	}
	return
}

func NewCIM_DiagnosticServiceRecordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DiagnosticServiceRecord, err error) {
	tmp, err := NewCIM_DiagnosticRecordEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticServiceRecord{
		CIM_DiagnosticRecord: tmp,
	}
	return
}

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *CIM_DiagnosticServiceRecord) SetPropertyErrorCode(value []string) (err error) {
	return instance.SetProperty("ErrorCode", (value))
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *CIM_DiagnosticServiceRecord) GetPropertyErrorCode() (value []string, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetErrorCount sets the value of ErrorCount for the instance
func (instance *CIM_DiagnosticServiceRecord) SetPropertyErrorCount(value []uint32) (err error) {
	return instance.SetProperty("ErrorCount", (value))
}

// GetErrorCount gets the value of ErrorCount for the instance
func (instance *CIM_DiagnosticServiceRecord) GetPropertyErrorCount() (value []uint32, err error) {
	retValue, err := instance.GetProperty("ErrorCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetLoopsFailed sets the value of LoopsFailed for the instance
func (instance *CIM_DiagnosticServiceRecord) SetPropertyLoopsFailed(value uint32) (err error) {
	return instance.SetProperty("LoopsFailed", (value))
}

// GetLoopsFailed gets the value of LoopsFailed for the instance
func (instance *CIM_DiagnosticServiceRecord) GetPropertyLoopsFailed() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoopsFailed")
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

// SetLoopsPassed sets the value of LoopsPassed for the instance
func (instance *CIM_DiagnosticServiceRecord) SetPropertyLoopsPassed(value uint32) (err error) {
	return instance.SetProperty("LoopsPassed", (value))
}

// GetLoopsPassed gets the value of LoopsPassed for the instance
func (instance *CIM_DiagnosticServiceRecord) GetPropertyLoopsPassed() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoopsPassed")
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
