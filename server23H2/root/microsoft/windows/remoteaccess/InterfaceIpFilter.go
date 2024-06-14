// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// InterfaceIpFilter struct
type InterfaceIpFilter struct {
	*cim.WmiInstance

	//
	Action uint32

	//
	AddressFamily uint32

	//
	Direction uint32

	//
	InterfaceAlias string

	//
	List []string
}

func NewInterfaceIpFilterEx1(instance *cim.WmiInstance) (newInstance *InterfaceIpFilter, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &InterfaceIpFilter{
		WmiInstance: tmp,
	}
	return
}

func NewInterfaceIpFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *InterfaceIpFilter, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &InterfaceIpFilter{
		WmiInstance: tmp,
	}
	return
}

// SetAction sets the value of Action for the instance
func (instance *InterfaceIpFilter) SetPropertyAction(value uint32) (err error) {
	return instance.SetProperty("Action", (value))
}

// GetAction gets the value of Action for the instance
func (instance *InterfaceIpFilter) GetPropertyAction() (value uint32, err error) {
	retValue, err := instance.GetProperty("Action")
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

// SetAddressFamily sets the value of AddressFamily for the instance
func (instance *InterfaceIpFilter) SetPropertyAddressFamily(value uint32) (err error) {
	return instance.SetProperty("AddressFamily", (value))
}

// GetAddressFamily gets the value of AddressFamily for the instance
func (instance *InterfaceIpFilter) GetPropertyAddressFamily() (value uint32, err error) {
	retValue, err := instance.GetProperty("AddressFamily")
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

// SetDirection sets the value of Direction for the instance
func (instance *InterfaceIpFilter) SetPropertyDirection(value uint32) (err error) {
	return instance.SetProperty("Direction", (value))
}

// GetDirection gets the value of Direction for the instance
func (instance *InterfaceIpFilter) GetPropertyDirection() (value uint32, err error) {
	retValue, err := instance.GetProperty("Direction")
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

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *InterfaceIpFilter) SetPropertyInterfaceAlias(value string) (err error) {
	return instance.SetProperty("InterfaceAlias", (value))
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *InterfaceIpFilter) GetPropertyInterfaceAlias() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceAlias")
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

// SetList sets the value of List for the instance
func (instance *InterfaceIpFilter) SetPropertyList(value []string) (err error) {
	return instance.SetProperty("List", (value))
}

// GetList gets the value of List for the instance
func (instance *InterfaceIpFilter) GetPropertyList() (value []string, err error) {
	retValue, err := instance.GetProperty("List")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
