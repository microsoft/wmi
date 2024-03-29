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

// CIM_DiagnosticCompletionRecord struct
type CIM_DiagnosticCompletionRecord struct {
	*CIM_DiagnosticServiceRecord

	//
	CompletionState DiagnosticCompletionRecord_CompletionState

	//
	OtherCompletionStateDescription string
}

func NewCIM_DiagnosticCompletionRecordEx1(instance *cim.WmiInstance) (newInstance *CIM_DiagnosticCompletionRecord, err error) {
	tmp, err := NewCIM_DiagnosticServiceRecordEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticCompletionRecord{
		CIM_DiagnosticServiceRecord: tmp,
	}
	return
}

func NewCIM_DiagnosticCompletionRecordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DiagnosticCompletionRecord, err error) {
	tmp, err := NewCIM_DiagnosticServiceRecordEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticCompletionRecord{
		CIM_DiagnosticServiceRecord: tmp,
	}
	return
}

// SetCompletionState sets the value of CompletionState for the instance
func (instance *CIM_DiagnosticCompletionRecord) SetPropertyCompletionState(value DiagnosticCompletionRecord_CompletionState) (err error) {
	return instance.SetProperty("CompletionState", (value))
}

// GetCompletionState gets the value of CompletionState for the instance
func (instance *CIM_DiagnosticCompletionRecord) GetPropertyCompletionState() (value DiagnosticCompletionRecord_CompletionState, err error) {
	retValue, err := instance.GetProperty("CompletionState")
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

	value = DiagnosticCompletionRecord_CompletionState(valuetmp)

	return
}

// SetOtherCompletionStateDescription sets the value of OtherCompletionStateDescription for the instance
func (instance *CIM_DiagnosticCompletionRecord) SetPropertyOtherCompletionStateDescription(value string) (err error) {
	return instance.SetProperty("OtherCompletionStateDescription", (value))
}

// GetOtherCompletionStateDescription gets the value of OtherCompletionStateDescription for the instance
func (instance *CIM_DiagnosticCompletionRecord) GetPropertyOtherCompletionStateDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherCompletionStateDescription")
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
