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

// MDM_WindowsAdvancedThreatProtection_DeviceTagging01 struct
type MDM_WindowsAdvancedThreatProtection_DeviceTagging01 struct {
	*cim.WmiInstance

	//
	Criticality int32

	//
	Group string

	//
	IdMethod int32

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_WindowsAdvancedThreatProtection_DeviceTagging01Ex1(instance *cim.WmiInstance) (newInstance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_WindowsAdvancedThreatProtection_DeviceTagging01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_WindowsAdvancedThreatProtection_DeviceTagging01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_WindowsAdvancedThreatProtection_DeviceTagging01{
		WmiInstance: tmp,
	}
	return
}

// SetCriticality sets the value of Criticality for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) SetPropertyCriticality(value int32) (err error) {
	return instance.SetProperty("Criticality", (value))
}

// GetCriticality gets the value of Criticality for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) GetPropertyCriticality() (value int32, err error) {
	retValue, err := instance.GetProperty("Criticality")
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

// SetGroup sets the value of Group for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) SetPropertyGroup(value string) (err error) {
	return instance.SetProperty("Group", (value))
}

// GetGroup gets the value of Group for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) GetPropertyGroup() (value string, err error) {
	retValue, err := instance.GetProperty("Group")
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

// SetIdMethod sets the value of IdMethod for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) SetPropertyIdMethod(value int32) (err error) {
	return instance.SetProperty("IdMethod", (value))
}

// GetIdMethod gets the value of IdMethod for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) GetPropertyIdMethod() (value int32, err error) {
	retValue, err := instance.GetProperty("IdMethod")
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
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) GetPropertyParentID() (value string, err error) {
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
