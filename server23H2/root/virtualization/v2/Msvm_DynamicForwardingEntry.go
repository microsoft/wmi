// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_DynamicForwardingEntry struct
type Msvm_DynamicForwardingEntry struct {
	*CIM_DynamicForwardingEntry

	//
	VlanId uint16
}

func NewMsvm_DynamicForwardingEntryEx1(instance *cim.WmiInstance) (newInstance *Msvm_DynamicForwardingEntry, err error) {
	tmp, err := NewCIM_DynamicForwardingEntryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_DynamicForwardingEntry{
		CIM_DynamicForwardingEntry: tmp,
	}
	return
}

func NewMsvm_DynamicForwardingEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_DynamicForwardingEntry, err error) {
	tmp, err := NewCIM_DynamicForwardingEntryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_DynamicForwardingEntry{
		CIM_DynamicForwardingEntry: tmp,
	}
	return
}

// SetVlanId sets the value of VlanId for the instance
func (instance *Msvm_DynamicForwardingEntry) SetPropertyVlanId(value uint16) (err error) {
	return instance.SetProperty("VlanId", (value))
}

// GetVlanId gets the value of VlanId for the instance
func (instance *Msvm_DynamicForwardingEntry) GetPropertyVlanId() (value uint16, err error) {
	retValue, err := instance.GetProperty("VlanId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
