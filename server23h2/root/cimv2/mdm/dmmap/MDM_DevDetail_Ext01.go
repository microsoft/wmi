// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// MDM_DevDetail_Ext01 struct
type MDM_DevDetail_Ext01 struct {
	*cim.WmiInstance

	//
	DeviceHardwareData string

	//
	InstanceID string

	//
	ParentID string

	//
	WLANMACAddress string
}

func NewMDM_DevDetail_Ext01Ex1(instance *cim.WmiInstance) (newInstance *MDM_DevDetail_Ext01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_DevDetail_Ext01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_DevDetail_Ext01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_DevDetail_Ext01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_DevDetail_Ext01{
		WmiInstance: tmp,
	}
	return
}

// SetDeviceHardwareData sets the value of DeviceHardwareData for the instance
func (instance *MDM_DevDetail_Ext01) SetPropertyDeviceHardwareData(value string) (err error) {
	return instance.SetProperty("DeviceHardwareData", (value))
}

// GetDeviceHardwareData gets the value of DeviceHardwareData for the instance
func (instance *MDM_DevDetail_Ext01) GetPropertyDeviceHardwareData() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceHardwareData")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_DevDetail_Ext01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DevDetail_Ext01) GetPropertyInstanceID() (value string, err error) {
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_DevDetail_Ext01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DevDetail_Ext01) GetPropertyParentID() (value string, err error) {
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

// SetWLANMACAddress sets the value of WLANMACAddress for the instance
func (instance *MDM_DevDetail_Ext01) SetPropertyWLANMACAddress(value string) (err error) {
	return instance.SetProperty("WLANMACAddress", (value))
}

// GetWLANMACAddress gets the value of WLANMACAddress for the instance
func (instance *MDM_DevDetail_Ext01) GetPropertyWLANMACAddress() (value string, err error) {
	retValue, err := instance.GetProperty("WLANMACAddress")
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
