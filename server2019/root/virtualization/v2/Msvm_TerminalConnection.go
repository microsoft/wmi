// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_TerminalConnection struct
type Msvm_TerminalConnection struct {
	*CIM_EnabledLogicalElement

	//
	ConnectionID string
}

func NewMsvm_TerminalConnectionEx1(instance *cim.WmiInstance) (newInstance *Msvm_TerminalConnection, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_TerminalConnection{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

func NewMsvm_TerminalConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_TerminalConnection, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_TerminalConnection{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

// SetConnectionID sets the value of ConnectionID for the instance
func (instance *Msvm_TerminalConnection) SetPropertyConnectionID(value string) (err error) {
	return instance.SetProperty("ConnectionID", value)
}

// GetConnectionID gets the value of ConnectionID for the instance
func (instance *Msvm_TerminalConnection) GetPropertyConnectionID() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
