// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VSIDConfiguration struct
type VSIDConfiguration struct {
	*cim.WmiInstance

	//
	InterfaceName string

	//
	Ipv4TransportInfo []uint8

	//
	Ipv6TransportInfo []uint8

	//
	LUID uint64

	//
	MTNICName string

	//
	RDID string

	//
	VSID uint32
}

func NewVSIDConfigurationEx1(instance *cim.WmiInstance) (newInstance *VSIDConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VSIDConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewVSIDConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VSIDConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VSIDConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetInterfaceName sets the value of InterfaceName for the instance
func (instance *VSIDConfiguration) SetPropertyInterfaceName(value string) (err error) {
	return instance.SetProperty("InterfaceName", (value))
}

// GetInterfaceName gets the value of InterfaceName for the instance
func (instance *VSIDConfiguration) GetPropertyInterfaceName() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceName")
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

// SetIpv4TransportInfo sets the value of Ipv4TransportInfo for the instance
func (instance *VSIDConfiguration) SetPropertyIpv4TransportInfo(value []uint8) (err error) {
	return instance.SetProperty("Ipv4TransportInfo", (value))
}

// GetIpv4TransportInfo gets the value of Ipv4TransportInfo for the instance
func (instance *VSIDConfiguration) GetPropertyIpv4TransportInfo() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Ipv4TransportInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetIpv6TransportInfo sets the value of Ipv6TransportInfo for the instance
func (instance *VSIDConfiguration) SetPropertyIpv6TransportInfo(value []uint8) (err error) {
	return instance.SetProperty("Ipv6TransportInfo", (value))
}

// GetIpv6TransportInfo gets the value of Ipv6TransportInfo for the instance
func (instance *VSIDConfiguration) GetPropertyIpv6TransportInfo() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Ipv6TransportInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetLUID sets the value of LUID for the instance
func (instance *VSIDConfiguration) SetPropertyLUID(value uint64) (err error) {
	return instance.SetProperty("LUID", (value))
}

// GetLUID gets the value of LUID for the instance
func (instance *VSIDConfiguration) GetPropertyLUID() (value uint64, err error) {
	retValue, err := instance.GetProperty("LUID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMTNICName sets the value of MTNICName for the instance
func (instance *VSIDConfiguration) SetPropertyMTNICName(value string) (err error) {
	return instance.SetProperty("MTNICName", (value))
}

// GetMTNICName gets the value of MTNICName for the instance
func (instance *VSIDConfiguration) GetPropertyMTNICName() (value string, err error) {
	retValue, err := instance.GetProperty("MTNICName")
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

// SetRDID sets the value of RDID for the instance
func (instance *VSIDConfiguration) SetPropertyRDID(value string) (err error) {
	return instance.SetProperty("RDID", (value))
}

// GetRDID gets the value of RDID for the instance
func (instance *VSIDConfiguration) GetPropertyRDID() (value string, err error) {
	retValue, err := instance.GetProperty("RDID")
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

// SetVSID sets the value of VSID for the instance
func (instance *VSIDConfiguration) SetPropertyVSID(value uint32) (err error) {
	return instance.SetProperty("VSID", (value))
}

// GetVSID gets the value of VSID for the instance
func (instance *VSIDConfiguration) GetPropertyVSID() (value uint32, err error) {
	retValue, err := instance.GetProperty("VSID")
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
