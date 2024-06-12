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

// MDM_DeviceUpdateCenter_DesiredUpdates01 struct
type MDM_DeviceUpdateCenter_DesiredUpdates01 struct {
	*cim.WmiInstance

	//
	InstanceID string

	//
	OcpVersion string

	//
	OsVersion string

	//
	ParentID string

	//
	SystemManifestVersion string
}

func NewMDM_DeviceUpdateCenter_DesiredUpdates01Ex1(instance *cim.WmiInstance) (newInstance *MDM_DeviceUpdateCenter_DesiredUpdates01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceUpdateCenter_DesiredUpdates01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_DeviceUpdateCenter_DesiredUpdates01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_DeviceUpdateCenter_DesiredUpdates01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceUpdateCenter_DesiredUpdates01{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_DeviceUpdateCenter_DesiredUpdates01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceUpdateCenter_DesiredUpdates01) GetPropertyInstanceID() (value string, err error) {
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

// SetOcpVersion sets the value of OcpVersion for the instance
func (instance *MDM_DeviceUpdateCenter_DesiredUpdates01) SetPropertyOcpVersion(value string) (err error) {
	return instance.SetProperty("OcpVersion", (value))
}

// GetOcpVersion gets the value of OcpVersion for the instance
func (instance *MDM_DeviceUpdateCenter_DesiredUpdates01) GetPropertyOcpVersion() (value string, err error) {
	retValue, err := instance.GetProperty("OcpVersion")
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

// SetOsVersion sets the value of OsVersion for the instance
func (instance *MDM_DeviceUpdateCenter_DesiredUpdates01) SetPropertyOsVersion(value string) (err error) {
	return instance.SetProperty("OsVersion", (value))
}

// GetOsVersion gets the value of OsVersion for the instance
func (instance *MDM_DeviceUpdateCenter_DesiredUpdates01) GetPropertyOsVersion() (value string, err error) {
	retValue, err := instance.GetProperty("OsVersion")
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
func (instance *MDM_DeviceUpdateCenter_DesiredUpdates01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceUpdateCenter_DesiredUpdates01) GetPropertyParentID() (value string, err error) {
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

// SetSystemManifestVersion sets the value of SystemManifestVersion for the instance
func (instance *MDM_DeviceUpdateCenter_DesiredUpdates01) SetPropertySystemManifestVersion(value string) (err error) {
	return instance.SetProperty("SystemManifestVersion", (value))
}

// GetSystemManifestVersion gets the value of SystemManifestVersion for the instance
func (instance *MDM_DeviceUpdateCenter_DesiredUpdates01) GetPropertySystemManifestVersion() (value string, err error) {
	retValue, err := instance.GetProperty("SystemManifestVersion")
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
