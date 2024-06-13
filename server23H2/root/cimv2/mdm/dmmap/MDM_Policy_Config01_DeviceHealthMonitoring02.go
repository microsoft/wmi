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

// MDM_Policy_Config01_DeviceHealthMonitoring02 struct
type MDM_Policy_Config01_DeviceHealthMonitoring02 struct {
	*cim.WmiInstance

	//
	AllowDeviceHealthMonitoring int32

	//
	ConfigDeviceHealthMonitoringScope string

	//
	ConfigDeviceHealthMonitoringServiceInstance string

	//
	ConfigDeviceHealthMonitoringUploadDestination string

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_Policy_Config01_DeviceHealthMonitoring02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_DeviceHealthMonitoring02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_DeviceHealthMonitoring02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_DeviceHealthMonitoring02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_DeviceHealthMonitoring02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_DeviceHealthMonitoring02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowDeviceHealthMonitoring sets the value of AllowDeviceHealthMonitoring for the instance
func (instance *MDM_Policy_Config01_DeviceHealthMonitoring02) SetPropertyAllowDeviceHealthMonitoring(value int32) (err error) {
	return instance.SetProperty("AllowDeviceHealthMonitoring", (value))
}

// GetAllowDeviceHealthMonitoring gets the value of AllowDeviceHealthMonitoring for the instance
func (instance *MDM_Policy_Config01_DeviceHealthMonitoring02) GetPropertyAllowDeviceHealthMonitoring() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowDeviceHealthMonitoring")
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

// SetConfigDeviceHealthMonitoringScope sets the value of ConfigDeviceHealthMonitoringScope for the instance
func (instance *MDM_Policy_Config01_DeviceHealthMonitoring02) SetPropertyConfigDeviceHealthMonitoringScope(value string) (err error) {
	return instance.SetProperty("ConfigDeviceHealthMonitoringScope", (value))
}

// GetConfigDeviceHealthMonitoringScope gets the value of ConfigDeviceHealthMonitoringScope for the instance
func (instance *MDM_Policy_Config01_DeviceHealthMonitoring02) GetPropertyConfigDeviceHealthMonitoringScope() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigDeviceHealthMonitoringScope")
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

// SetConfigDeviceHealthMonitoringServiceInstance sets the value of ConfigDeviceHealthMonitoringServiceInstance for the instance
func (instance *MDM_Policy_Config01_DeviceHealthMonitoring02) SetPropertyConfigDeviceHealthMonitoringServiceInstance(value string) (err error) {
	return instance.SetProperty("ConfigDeviceHealthMonitoringServiceInstance", (value))
}

// GetConfigDeviceHealthMonitoringServiceInstance gets the value of ConfigDeviceHealthMonitoringServiceInstance for the instance
func (instance *MDM_Policy_Config01_DeviceHealthMonitoring02) GetPropertyConfigDeviceHealthMonitoringServiceInstance() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigDeviceHealthMonitoringServiceInstance")
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

// SetConfigDeviceHealthMonitoringUploadDestination sets the value of ConfigDeviceHealthMonitoringUploadDestination for the instance
func (instance *MDM_Policy_Config01_DeviceHealthMonitoring02) SetPropertyConfigDeviceHealthMonitoringUploadDestination(value string) (err error) {
	return instance.SetProperty("ConfigDeviceHealthMonitoringUploadDestination", (value))
}

// GetConfigDeviceHealthMonitoringUploadDestination gets the value of ConfigDeviceHealthMonitoringUploadDestination for the instance
func (instance *MDM_Policy_Config01_DeviceHealthMonitoring02) GetPropertyConfigDeviceHealthMonitoringUploadDestination() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigDeviceHealthMonitoringUploadDestination")
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
func (instance *MDM_Policy_Config01_DeviceHealthMonitoring02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_DeviceHealthMonitoring02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_DeviceHealthMonitoring02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_DeviceHealthMonitoring02) GetPropertyParentID() (value string, err error) {
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
