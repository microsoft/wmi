// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WHEAErrorSourceMethods struct
type WHEAErrorSourceMethods struct {
	*WHEA

	//
	Active bool

	//
	InstanceName string
}

func NewWHEAErrorSourceMethodsEx1(instance *cim.WmiInstance) (newInstance *WHEAErrorSourceMethods, err error) {
	tmp, err := NewWHEAEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WHEAErrorSourceMethods{
		WHEA: tmp,
	}
	return
}

func NewWHEAErrorSourceMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WHEAErrorSourceMethods, err error) {
	tmp, err := NewWHEAEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WHEAErrorSourceMethods{
		WHEA: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *WHEAErrorSourceMethods) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WHEAErrorSourceMethods) GetPropertyActive() (value bool, err error) {
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
func (instance *WHEAErrorSourceMethods) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WHEAErrorSourceMethods) GetPropertyInstanceName() (value string, err error) {
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
// <param name="ErrorSourceArray" type="uint8 []"></param>
// <param name="Length" type="uint32 "></param>
// <param name="Status" type="uint32 "></param>
func (instance *WHEAErrorSourceMethods) GetAllErrorSourcesRtn( /* OUT */ Status uint32,
	/* OUT */ Count uint32,
	/* OUT */ Length uint32,
	/* OUT */ ErrorSourceArray []uint8) (err error) {
	_, err = instance.InvokeMethod("GetAllErrorSourcesRtn")
	if err != nil {
		return
	}
	return

}

//

// <param name="ErrorSourceId" type="uint32 "></param>

// <param name="ErrorSourceInfo" type="uint8 []"></param>
// <param name="Length" type="uint32 "></param>
// <param name="Status" type="uint32 "></param>
func (instance *WHEAErrorSourceMethods) GetErrorSourceInfoRtn( /* OUT */ Status uint32,
	/* OPTIONAL IN */ ErrorSourceId uint32,
	/* OUT */ Length uint32,
	/* OUT */ ErrorSourceInfo []uint8) (err error) {
	_, err = instance.InvokeMethod("GetErrorSourceInfoRtn", ErrorSourceId)
	if err != nil {
		return
	}
	return

}

//

// <param name="ErrorSourceInfo" type="uint8 []"></param>
// <param name="Length" type="uint32 "></param>

// <param name="Status" type="uint32 "></param>
func (instance *WHEAErrorSourceMethods) SetErrorSourceInfoRtn( /* OUT */ Status uint32,
	/* OPTIONAL IN */ Length uint32,
	/* OPTIONAL IN */ ErrorSourceInfo []uint8) (err error) {
	_, err = instance.InvokeMethod("SetErrorSourceInfoRtn", Length, ErrorSourceInfo)
	if err != nil {
		return
	}
	return

}

//

// <param name="ErrorSourceId" type="uint32 "></param>

// <param name="Status" type="uint32 "></param>
func (instance *WHEAErrorSourceMethods) EnableErrorSourceRtn( /* OUT */ Status uint32,
	/* OPTIONAL IN */ ErrorSourceId uint32) (err error) {
	_, err = instance.InvokeMethod("EnableErrorSourceRtn", ErrorSourceId)
	if err != nil {
		return
	}
	return

}

//

// <param name="ErrorSourceId" type="uint32 "></param>

// <param name="Status" type="uint32 "></param>
func (instance *WHEAErrorSourceMethods) DisableErrorSourceRtn( /* OUT */ Status uint32,
	/* OPTIONAL IN */ ErrorSourceId uint32) (err error) {
	_, err = instance.InvokeMethod("DisableErrorSourceRtn", ErrorSourceId)
	if err != nil {
		return
	}
	return

}
