// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// MDM_DeviceUpdateCenter_Enrollment01 struct
type MDM_DeviceUpdateCenter_Enrollment01 struct {
	*cim.WmiInstance

	//
	CustomPackageId string

	//
	DeviceModelId string

	//
	InstanceID string

	//
	OemPartnerRing string

	//
	ParentID string

	//
	PublisherId string
}

func NewMDM_DeviceUpdateCenter_Enrollment01Ex1(instance *cim.WmiInstance) (newInstance *MDM_DeviceUpdateCenter_Enrollment01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceUpdateCenter_Enrollment01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_DeviceUpdateCenter_Enrollment01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_DeviceUpdateCenter_Enrollment01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceUpdateCenter_Enrollment01{
		WmiInstance: tmp,
	}
	return
}

// SetCustomPackageId sets the value of CustomPackageId for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) SetPropertyCustomPackageId(value string) (err error) {
	return instance.SetProperty("CustomPackageId", (value))
}

// GetCustomPackageId gets the value of CustomPackageId for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) GetPropertyCustomPackageId() (value string, err error) {
	retValue, err := instance.GetProperty("CustomPackageId")
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

// SetDeviceModelId sets the value of DeviceModelId for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) SetPropertyDeviceModelId(value string) (err error) {
	return instance.SetProperty("DeviceModelId", (value))
}

// GetDeviceModelId gets the value of DeviceModelId for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) GetPropertyDeviceModelId() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceModelId")
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
func (instance *MDM_DeviceUpdateCenter_Enrollment01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) GetPropertyInstanceID() (value string, err error) {
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

// SetOemPartnerRing sets the value of OemPartnerRing for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) SetPropertyOemPartnerRing(value string) (err error) {
	return instance.SetProperty("OemPartnerRing", (value))
}

// GetOemPartnerRing gets the value of OemPartnerRing for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) GetPropertyOemPartnerRing() (value string, err error) {
	retValue, err := instance.GetProperty("OemPartnerRing")
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
func (instance *MDM_DeviceUpdateCenter_Enrollment01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) GetPropertyParentID() (value string, err error) {
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

// SetPublisherId sets the value of PublisherId for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) SetPropertyPublisherId(value string) (err error) {
	return instance.SetProperty("PublisherId", (value))
}

// GetPublisherId gets the value of PublisherId for the instance
func (instance *MDM_DeviceUpdateCenter_Enrollment01) GetPropertyPublisherId() (value string, err error) {
	retValue, err := instance.GetProperty("PublisherId")
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
