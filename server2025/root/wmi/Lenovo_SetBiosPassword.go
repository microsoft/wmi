// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Lenovo_SetBiosPassword struct
type Lenovo_SetBiosPassword struct {
	*Lenovo_BIOSElement

	//
	Active bool

	//
	InstanceName string
}

func NewLenovo_SetBiosPasswordEx1(instance *cim.WmiInstance) (newInstance *Lenovo_SetBiosPassword, err error) {
	tmp, err := NewLenovo_BIOSElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Lenovo_SetBiosPassword{
		Lenovo_BIOSElement: tmp,
	}
	return
}

func NewLenovo_SetBiosPasswordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Lenovo_SetBiosPassword, err error) {
	tmp, err := NewLenovo_BIOSElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Lenovo_SetBiosPassword{
		Lenovo_BIOSElement: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *Lenovo_SetBiosPassword) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *Lenovo_SetBiosPassword) GetPropertyActive() (value bool, err error) {
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
func (instance *Lenovo_SetBiosPassword) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *Lenovo_SetBiosPassword) GetPropertyInstanceName() (value string, err error) {
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

// <param name="Parameter" type="string "></param>

// <param name="Return" type="string "></param>
func (instance *Lenovo_SetBiosPassword) SetBiosPassword( /* IN */ Parameter string,
	/* OUT */ Return string) (err error) {
	_, err = instance.InvokeMethod("SetBiosPassword", Parameter)
	if err != nil {
		return
	}
	return

}
