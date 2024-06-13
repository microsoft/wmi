// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_DeviceStatus_NetworkIdentifiers01_01 struct
type MDM_DeviceStatus_NetworkIdentifiers01_01 struct {
	*cim.WmiInstance

	//
	InstanceID string

	//
	IPAddressV4 string

	//
	IPAddressV6 string

	//
	IsConnected bool

	//
	ParentID string

	//
	Type int32
}

func NewMDM_DeviceStatus_NetworkIdentifiers01_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_DeviceStatus_NetworkIdentifiers01_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceStatus_NetworkIdentifiers01_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_DeviceStatus_NetworkIdentifiers01_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_DeviceStatus_NetworkIdentifiers01_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceStatus_NetworkIdentifiers01_01{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetIPAddressV4 sets the value of IPAddressV4 for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) SetPropertyIPAddressV4(value string) (err error) {
	return instance.SetProperty("IPAddressV4", (value))
}

// GetIPAddressV4 gets the value of IPAddressV4 for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) GetPropertyIPAddressV4() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddressV4")
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

// SetIPAddressV6 sets the value of IPAddressV6 for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) SetPropertyIPAddressV6(value string) (err error) {
	return instance.SetProperty("IPAddressV6", (value))
}

// GetIPAddressV6 gets the value of IPAddressV6 for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) GetPropertyIPAddressV6() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddressV6")
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

// SetIsConnected sets the value of IsConnected for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) SetPropertyIsConnected(value bool) (err error) {
	return instance.SetProperty("IsConnected", (value))
}

// GetIsConnected gets the value of IsConnected for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) GetPropertyIsConnected() (value bool, err error) {
	retValue, err := instance.GetProperty("IsConnected")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetType sets the value of Type for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) SetPropertyType(value int32) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MDM_DeviceStatus_NetworkIdentifiers01_01) GetPropertyType() (value int32, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
