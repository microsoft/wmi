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

// WHEAPageOfflineMethods struct
type WHEAPageOfflineMethods struct {
	*WHEA

	//
	Active bool

	//
	InstanceName string
}

func NewWHEAPageOfflineMethodsEx1(instance *cim.WmiInstance) (newInstance *WHEAPageOfflineMethods, err error) {
	tmp, err := NewWHEAEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WHEAPageOfflineMethods{
		WHEA: tmp,
	}
	return
}

func NewWHEAPageOfflineMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WHEAPageOfflineMethods, err error) {
	tmp, err := NewWHEAEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WHEAPageOfflineMethods{
		WHEA: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *WHEAPageOfflineMethods) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WHEAPageOfflineMethods) GetPropertyActive() (value bool, err error) {
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
func (instance *WHEAPageOfflineMethods) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WHEAPageOfflineMethods) GetPropertyInstanceName() (value string, err error) {
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

// <param name="InputBuffer" type="uint8 []"></param>
// <param name="InputLength" type="uint32 "></param>

// <param name="Status" type="uint32 "></param>
func (instance *WHEAPageOfflineMethods) WheaPageOfflineInterfaceRtn( /* IN */ InputLength uint32,
	/* IN */ InputBuffer []uint8,
	/* OUT */ Status uint32) (err error) {
	_, err = instance.InvokeMethod("WheaPageOfflineInterfaceRtn", InputLength, InputBuffer)
	if err != nil {
		return
	}
	return

}
