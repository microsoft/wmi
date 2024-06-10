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

// ALPC_Receive_Message struct
type ALPC_Receive_Message struct {
	*ALPC

	//
	MessageID uint32
}

func NewALPC_Receive_MessageEx1(instance *cim.WmiInstance) (newInstance *ALPC_Receive_Message, err error) {
	tmp, err := NewALPCEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ALPC_Receive_Message{
		ALPC: tmp,
	}
	return
}

func NewALPC_Receive_MessageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ALPC_Receive_Message, err error) {
	tmp, err := NewALPCEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ALPC_Receive_Message{
		ALPC: tmp,
	}
	return
}

// SetMessageID sets the value of MessageID for the instance
func (instance *ALPC_Receive_Message) SetPropertyMessageID(value uint32) (err error) {
	return instance.SetProperty("MessageID", (value))
}

// GetMessageID gets the value of MessageID for the instance
func (instance *ALPC_Receive_Message) GetPropertyMessageID() (value uint32, err error) {
	retValue, err := instance.GetProperty("MessageID")
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
