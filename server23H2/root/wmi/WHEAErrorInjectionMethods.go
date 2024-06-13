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

// WHEAErrorInjectionMethods struct
type WHEAErrorInjectionMethods struct {
	*WHEA

	//
	Active bool

	//
	InstanceName string
}

func NewWHEAErrorInjectionMethodsEx1(instance *cim.WmiInstance) (newInstance *WHEAErrorInjectionMethods, err error) {
	tmp, err := NewWHEAEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WHEAErrorInjectionMethods{
		WHEA: tmp,
	}
	return
}

func NewWHEAErrorInjectionMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WHEAErrorInjectionMethods, err error) {
	tmp, err := NewWHEAEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WHEAErrorInjectionMethods{
		WHEA: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *WHEAErrorInjectionMethods) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WHEAErrorInjectionMethods) GetPropertyActive() (value bool, err error) {
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
func (instance *WHEAErrorInjectionMethods) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WHEAErrorInjectionMethods) GetPropertyInstanceName() (value string, err error) {
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

// <param name="Capabilities" type="uint32 "></param>
// <param name="Status" type="uint32 "></param>
func (instance *WHEAErrorInjectionMethods) GetErrorInjectionCapabilitiesRtn( /* OUT */ Status uint32,
	/* OUT */ Capabilities uint32) (err error) {
	_, err = instance.InvokeMethod("GetErrorInjectionCapabilitiesRtn")
	if err != nil {
		return
	}
	return

}

//

// <param name="ErrorType" type="uint32 "></param>
// <param name="Parameter1" type="uint64 "></param>
// <param name="Parameter2" type="uint64 "></param>
// <param name="Parameter3" type="uint64 "></param>
// <param name="Parameter4" type="uint64 "></param>

// <param name="Status" type="uint32 "></param>
func (instance *WHEAErrorInjectionMethods) InjectErrorRtn( /* IN */ ErrorType uint32,
	/* IN */ Parameter1 uint64,
	/* IN */ Parameter2 uint64,
	/* IN */ Parameter3 uint64,
	/* IN */ Parameter4 uint64,
	/* OUT */ Status uint32) (err error) {
	_, err = instance.InvokeMethod("InjectErrorRtn", ErrorType, Parameter1, Parameter2, Parameter3, Parameter4)
	if err != nil {
		return
	}
	return

}
