// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ALPC_Unwait struct
type ALPC_Unwait struct {
	*ALPC

	//
	Status uint32
}

func NewALPC_UnwaitEx1(instance *cim.WmiInstance) (newInstance *ALPC_Unwait, err error) {
	tmp, err := NewALPCEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ALPC_Unwait{
		ALPC: tmp,
	}
	return
}

func NewALPC_UnwaitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ALPC_Unwait, err error) {
	tmp, err := NewALPCEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ALPC_Unwait{
		ALPC: tmp,
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *ALPC_Unwait) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *ALPC_Unwait) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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
