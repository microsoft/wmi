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

// MSNdis_SetInterruptModeration struct
type MSNdis_SetInterruptModeration struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string
}

func NewMSNdis_SetInterruptModerationEx1(instance *cim.WmiInstance) (newInstance *MSNdis_SetInterruptModeration, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_SetInterruptModeration{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_SetInterruptModerationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_SetInterruptModeration, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_SetInterruptModeration{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_SetInterruptModeration) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_SetInterruptModeration) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSNdis_SetInterruptModeration) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_SetInterruptModeration) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

//

// <param name="InterruptModeration" type="MSNdis_InterruptModerationParameters "></param>
// <param name="MethodHeader" type="MSNdis_WmiMethodHeader "></param>

// <param name="OutputInfo" type="MSNdis_WmiOutputInfo "></param>
func (instance *MSNdis_SetInterruptModeration) WmiSetInterruptModeration( /* IN */ MethodHeader MSNdis_WmiMethodHeader,
	/* IN */ InterruptModeration MSNdis_InterruptModerationParameters,
	/* OUT */ OutputInfo MSNdis_WmiOutputInfo) (err error) {
	_, err = instance.InvokeMethod("WmiSetInterruptModeration", MethodHeader, InterruptModeration)
	if err != nil {
		return
	}
	return

}
