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

// WHEAPolicyManagementMethods struct
type WHEAPolicyManagementMethods struct {
	*WHEA

	//
	Active bool

	//
	InstanceName string
}

func NewWHEAPolicyManagementMethodsEx1(instance *cim.WmiInstance) (newInstance *WHEAPolicyManagementMethods, err error) {
	tmp, err := NewWHEAEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WHEAPolicyManagementMethods{
		WHEA: tmp,
	}
	return
}

func NewWHEAPolicyManagementMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WHEAPolicyManagementMethods, err error) {
	tmp, err := NewWHEAEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WHEAPolicyManagementMethods{
		WHEA: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *WHEAPolicyManagementMethods) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WHEAPolicyManagementMethods) GetPropertyActive() (value bool, err error) {
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
func (instance *WHEAPolicyManagementMethods) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WHEAPolicyManagementMethods) GetPropertyInstanceName() (value string, err error) {
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

// <param name="Count" type="uint32 "></param>
// <param name="Length" type="uint32 "></param>
// <param name="Status" type="uint32 "></param>
// <param name="Values" type="uint8 []"></param>
func (instance *WHEAPolicyManagementMethods) WheaGetAllPolicyRtn( /* OUT */ Status uint32,
	/* OUT */ Count uint32,
	/* OUT */ Length uint32,
	/* OUT */ Values []uint8) (err error) {
	_, err = instance.InvokeMethod("WheaGetAllPolicyRtn")
	if err != nil {
		return
	}
	return

}

//

// <param name="Type" type="uint32 "></param>

// <param name="Status" type="uint32 "></param>
// <param name="Value" type="uint32 "></param>
func (instance *WHEAPolicyManagementMethods) WheaGetPolicyRtn( /* OUT */ Status uint32,
	/* OPTIONAL IN */ Type uint32,
	/* OUT */ Value uint32) (err error) {
	_, err = instance.InvokeMethod("WheaGetPolicyRtn", Type)
	if err != nil {
		return
	}
	return

}

//

// <param name="Type" type="uint32 "></param>
// <param name="Value" type="uint32 "></param>

// <param name="Status" type="uint32 "></param>
func (instance *WHEAPolicyManagementMethods) WheaSetPolicyRtn( /* OUT */ Status uint32,
	/* OPTIONAL IN */ Type uint32,
	/* OPTIONAL IN */ Value uint32) (err error) {
	_, err = instance.InvokeMethod("WheaSetPolicyRtn", Type, Value)
	if err != nil {
		return
	}
	return

}

//

// <param name="Status" type="uint32 "></param>
func (instance *WHEAPolicyManagementMethods) WheaCommitPolicyRtn( /* OUT */ Status uint32) (err error) {
	_, err = instance.InvokeMethod("WheaCommitPolicyRtn")
	if err != nil {
		return
	}
	return

}

//

// <param name="Status" type="uint32 "></param>
func (instance *WHEAPolicyManagementMethods) WheaResetPolicyRtn( /* OUT */ Status uint32) (err error) {
	_, err = instance.InvokeMethod("WheaResetPolicyRtn")
	if err != nil {
		return
	}
	return

}
