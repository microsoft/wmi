// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

// MSFT_TaskPrincipal2 struct
type MSFT_TaskPrincipal2 struct {
	MSFT_TaskPrincipal

	//
	ProcessTokenSidType TaskPrincipal2_ProcessTokenSidType

	//
	RequiredPrivilege []string
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
