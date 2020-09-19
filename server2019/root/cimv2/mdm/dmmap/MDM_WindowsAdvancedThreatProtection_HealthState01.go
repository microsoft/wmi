// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// MDM_WindowsAdvancedThreatProtection_HealthState01 struct
type MDM_WindowsAdvancedThreatProtection_HealthState01 struct {
	*cim.WmiInstance

	//
	InstanceID string

	//
	LastConnected string

	//
	OnboardingState int32

	//
	OrgId string

	//
	ParentID string

	//
	SenseIsRunning bool
}

func NewMDM_WindowsAdvancedThreatProtection_HealthState01Ex1(instance *cim.WmiInstance) (newInstance *MDM_WindowsAdvancedThreatProtection_HealthState01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_WindowsAdvancedThreatProtection_HealthState01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_WindowsAdvancedThreatProtection_HealthState01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_WindowsAdvancedThreatProtection_HealthState01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_WindowsAdvancedThreatProtection_HealthState01{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_HealthState01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_HealthState01) GetPropertyInstanceID() (value string, err error) {
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

// SetLastConnected sets the value of LastConnected for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_HealthState01) SetPropertyLastConnected(value string) (err error) {
	return instance.SetProperty("LastConnected", (value))
}

// GetLastConnected gets the value of LastConnected for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_HealthState01) GetPropertyLastConnected() (value string, err error) {
	retValue, err := instance.GetProperty("LastConnected")
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

// SetOnboardingState sets the value of OnboardingState for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_HealthState01) SetPropertyOnboardingState(value int32) (err error) {
	return instance.SetProperty("OnboardingState", (value))
}

// GetOnboardingState gets the value of OnboardingState for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_HealthState01) GetPropertyOnboardingState() (value int32, err error) {
	retValue, err := instance.GetProperty("OnboardingState")
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

// SetOrgId sets the value of OrgId for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_HealthState01) SetPropertyOrgId(value string) (err error) {
	return instance.SetProperty("OrgId", (value))
}

// GetOrgId gets the value of OrgId for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_HealthState01) GetPropertyOrgId() (value string, err error) {
	retValue, err := instance.GetProperty("OrgId")
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
func (instance *MDM_WindowsAdvancedThreatProtection_HealthState01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_HealthState01) GetPropertyParentID() (value string, err error) {
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

// SetSenseIsRunning sets the value of SenseIsRunning for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_HealthState01) SetPropertySenseIsRunning(value bool) (err error) {
	return instance.SetProperty("SenseIsRunning", (value))
}

// GetSenseIsRunning gets the value of SenseIsRunning for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_HealthState01) GetPropertySenseIsRunning() (value bool, err error) {
	retValue, err := instance.GetProperty("SenseIsRunning")
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
