// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_EASPolicy struct
type MDM_EASPolicy struct {
	*cim.WmiInstance

	//
	key uint32
}

func NewMDM_EASPolicyEx1(instance *cim.WmiInstance) (newInstance *MDM_EASPolicy, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_EASPolicy{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_EASPolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_EASPolicy, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_EASPolicy{
		WmiInstance: tmp,
	}
	return
}

// Setkey sets the value of key for the instance
func (instance *MDM_EASPolicy) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *MDM_EASPolicy) GetPropertykey() (value uint32, err error) {
	retValue, err := instance.GetProperty("key")
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

//

// <param name="NamedValuesList" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_EASPolicy) SetValues( /* IN */ NamedValuesList string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetValues", NamedValuesList)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
