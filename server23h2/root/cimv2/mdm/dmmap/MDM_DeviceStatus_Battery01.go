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

// MDM_DeviceStatus_Battery01 struct
type MDM_DeviceStatus_Battery01 struct {
	*cim.WmiInstance

	//
	EstimatedChargeRemaining int32

	//
	EstimatedRuntime int32

	//
	InstanceID string

	//
	ParentID string

	//
	Status int32
}

func NewMDM_DeviceStatus_Battery01Ex1(instance *cim.WmiInstance) (newInstance *MDM_DeviceStatus_Battery01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceStatus_Battery01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_DeviceStatus_Battery01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_DeviceStatus_Battery01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceStatus_Battery01{
		WmiInstance: tmp,
	}
	return
}

// SetEstimatedChargeRemaining sets the value of EstimatedChargeRemaining for the instance
func (instance *MDM_DeviceStatus_Battery01) SetPropertyEstimatedChargeRemaining(value int32) (err error) {
	return instance.SetProperty("EstimatedChargeRemaining", (value))
}

// GetEstimatedChargeRemaining gets the value of EstimatedChargeRemaining for the instance
func (instance *MDM_DeviceStatus_Battery01) GetPropertyEstimatedChargeRemaining() (value int32, err error) {
	retValue, err := instance.GetProperty("EstimatedChargeRemaining")
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

// SetEstimatedRuntime sets the value of EstimatedRuntime for the instance
func (instance *MDM_DeviceStatus_Battery01) SetPropertyEstimatedRuntime(value int32) (err error) {
	return instance.SetProperty("EstimatedRuntime", (value))
}

// GetEstimatedRuntime gets the value of EstimatedRuntime for the instance
func (instance *MDM_DeviceStatus_Battery01) GetPropertyEstimatedRuntime() (value int32, err error) {
	retValue, err := instance.GetProperty("EstimatedRuntime")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_Battery01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_Battery01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_DeviceStatus_Battery01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceStatus_Battery01) GetPropertyParentID() (value string, err error) {
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

// SetStatus sets the value of Status for the instance
func (instance *MDM_DeviceStatus_Battery01) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_DeviceStatus_Battery01) GetPropertyStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Status")
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
