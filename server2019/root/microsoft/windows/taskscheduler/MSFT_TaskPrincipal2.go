// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("ProcessTokenSidType", value)
}

// GetProcessTokenSidType gets the value of ProcessTokenSidType for the instance
func (instance *MSFT_TaskPrincipal2) GetPropertyProcessTokenSidType() (value TaskPrincipal2_ProcessTokenSidType, err error) {
	retValue, err := instance.GetProperty("ProcessTokenSidType")
	if err != nil {
		return
	}
	value, ok := retValue.(TaskPrincipal2_ProcessTokenSidType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequiredPrivilege sets the value of RequiredPrivilege for the instance
func (instance *MSFT_TaskPrincipal2) SetPropertyRequiredPrivilege(value []string) (err error) {
	return instance.SetProperty("RequiredPrivilege", value)
}

// GetRequiredPrivilege gets the value of RequiredPrivilege for the instance
func (instance *MSFT_TaskPrincipal2) GetPropertyRequiredPrivilege() (value []string, err error) {
	retValue, err := instance.GetProperty("RequiredPrivilege")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
