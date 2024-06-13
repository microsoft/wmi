// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_TaskPrincipal2 struct
type MSFT_TaskPrincipal2 struct {
	*MSFT_TaskPrincipal

	//
	ProcessTokenSidType TaskPrincipal2_ProcessTokenSidType

	//
	RequiredPrivilege []string
}

func NewMSFT_TaskPrincipal2Ex1(instance *cim.WmiInstance) (newInstance *MSFT_TaskPrincipal2, err error) {
	tmp, err := NewMSFT_TaskPrincipalEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskPrincipal2{
		MSFT_TaskPrincipal: tmp,
	}
	return
}

func NewMSFT_TaskPrincipal2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskPrincipal2, err error) {
	tmp, err := NewMSFT_TaskPrincipalEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskPrincipal2{
		MSFT_TaskPrincipal: tmp,
	}
	return
}

// SetProcessTokenSidType sets the value of ProcessTokenSidType for the instance
func (instance *MSFT_TaskPrincipal2) SetPropertyProcessTokenSidType(value TaskPrincipal2_ProcessTokenSidType) (err error) {
	return instance.SetProperty("ProcessTokenSidType", (value))
}

// GetProcessTokenSidType gets the value of ProcessTokenSidType for the instance
func (instance *MSFT_TaskPrincipal2) GetPropertyProcessTokenSidType() (value TaskPrincipal2_ProcessTokenSidType, err error) {
	retValue, err := instance.GetProperty("ProcessTokenSidType")
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

	value = TaskPrincipal2_ProcessTokenSidType(valuetmp)

	return
}

// SetRequiredPrivilege sets the value of RequiredPrivilege for the instance
func (instance *MSFT_TaskPrincipal2) SetPropertyRequiredPrivilege(value []string) (err error) {
	return instance.SetProperty("RequiredPrivilege", (value))
}

// GetRequiredPrivilege gets the value of RequiredPrivilege for the instance
func (instance *MSFT_TaskPrincipal2) GetPropertyRequiredPrivilege() (value []string, err error) {
	retValue, err := instance.GetProperty("RequiredPrivilege")
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
