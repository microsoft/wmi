// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ALPC_Wait_For_New_Message struct
type ALPC_Wait_For_New_Message struct {
	*ALPC

	//
	IsServerPort uint32

	//
	PortName string
}

func NewALPC_Wait_For_New_MessageEx1(instance *cim.WmiInstance) (newInstance *ALPC_Wait_For_New_Message, err error) {
	tmp, err := NewALPCEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ALPC_Wait_For_New_Message{
		ALPC: tmp,
	}
	return
}

func NewALPC_Wait_For_New_MessageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ALPC_Wait_For_New_Message, err error) {
	tmp, err := NewALPCEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ALPC_Wait_For_New_Message{
		ALPC: tmp,
	}
	return
}

// SetIsServerPort sets the value of IsServerPort for the instance
func (instance *ALPC_Wait_For_New_Message) SetPropertyIsServerPort(value uint32) (err error) {
	return instance.SetProperty("IsServerPort", (value))
}

// GetIsServerPort gets the value of IsServerPort for the instance
func (instance *ALPC_Wait_For_New_Message) GetPropertyIsServerPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsServerPort")
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

// SetPortName sets the value of PortName for the instance
func (instance *ALPC_Wait_For_New_Message) SetPropertyPortName(value string) (err error) {
	return instance.SetProperty("PortName", (value))
}

// GetPortName gets the value of PortName for the instance
func (instance *ALPC_Wait_For_New_Message) GetPropertyPortName() (value string, err error) {
	retValue, err := instance.GetProperty("PortName")
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
