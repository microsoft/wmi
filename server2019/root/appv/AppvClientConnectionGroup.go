// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Appv
//////////////////////////////////////////////
package appv

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// AppvClientConnectionGroup struct
type AppvClientConnectionGroup struct {
	*cim.WmiInstance

	//
	CustomData string

	//
	GlobalPending bool

	//
	GroupId string

	//
	InUse bool

	//
	IsEnabledGlobally bool

	//
	IsEnabledToUser bool

	//
	Name string

	//
	Packages []string

	//
	PercentLoaded uint16

	//
	Priority uint32

	//
	UserPending bool

	//
	VersionId string
}

func NewAppvClientConnectionGroupEx1(instance *cim.WmiInstance) (newInstance *AppvClientConnectionGroup, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &AppvClientConnectionGroup{
		WmiInstance: tmp,
	}
	return
}

func NewAppvClientConnectionGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AppvClientConnectionGroup, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AppvClientConnectionGroup{
		WmiInstance: tmp,
	}
	return
}

// SetCustomData sets the value of CustomData for the instance
func (instance *AppvClientConnectionGroup) SetPropertyCustomData(value string) (err error) {
	return instance.SetProperty("CustomData", value)
}

// GetCustomData gets the value of CustomData for the instance
func (instance *AppvClientConnectionGroup) GetPropertyCustomData() (value string, err error) {
	retValue, err := instance.GetProperty("CustomData")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGlobalPending sets the value of GlobalPending for the instance
func (instance *AppvClientConnectionGroup) SetPropertyGlobalPending(value bool) (err error) {
	return instance.SetProperty("GlobalPending", value)
}

// GetGlobalPending gets the value of GlobalPending for the instance
func (instance *AppvClientConnectionGroup) GetPropertyGlobalPending() (value bool, err error) {
	retValue, err := instance.GetProperty("GlobalPending")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGroupId sets the value of GroupId for the instance
func (instance *AppvClientConnectionGroup) SetPropertyGroupId(value string) (err error) {
	return instance.SetProperty("GroupId", value)
}

// GetGroupId gets the value of GroupId for the instance
func (instance *AppvClientConnectionGroup) GetPropertyGroupId() (value string, err error) {
	retValue, err := instance.GetProperty("GroupId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInUse sets the value of InUse for the instance
func (instance *AppvClientConnectionGroup) SetPropertyInUse(value bool) (err error) {
	return instance.SetProperty("InUse", value)
}

// GetInUse gets the value of InUse for the instance
func (instance *AppvClientConnectionGroup) GetPropertyInUse() (value bool, err error) {
	retValue, err := instance.GetProperty("InUse")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsEnabledGlobally sets the value of IsEnabledGlobally for the instance
func (instance *AppvClientConnectionGroup) SetPropertyIsEnabledGlobally(value bool) (err error) {
	return instance.SetProperty("IsEnabledGlobally", value)
}

// GetIsEnabledGlobally gets the value of IsEnabledGlobally for the instance
func (instance *AppvClientConnectionGroup) GetPropertyIsEnabledGlobally() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEnabledGlobally")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsEnabledToUser sets the value of IsEnabledToUser for the instance
func (instance *AppvClientConnectionGroup) SetPropertyIsEnabledToUser(value bool) (err error) {
	return instance.SetProperty("IsEnabledToUser", value)
}

// GetIsEnabledToUser gets the value of IsEnabledToUser for the instance
func (instance *AppvClientConnectionGroup) GetPropertyIsEnabledToUser() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEnabledToUser")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *AppvClientConnectionGroup) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *AppvClientConnectionGroup) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackages sets the value of Packages for the instance
func (instance *AppvClientConnectionGroup) SetPropertyPackages(value []string) (err error) {
	return instance.SetProperty("Packages", value)
}

// GetPackages gets the value of Packages for the instance
func (instance *AppvClientConnectionGroup) GetPropertyPackages() (value []string, err error) {
	retValue, err := instance.GetProperty("Packages")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentLoaded sets the value of PercentLoaded for the instance
func (instance *AppvClientConnectionGroup) SetPropertyPercentLoaded(value uint16) (err error) {
	return instance.SetProperty("PercentLoaded", value)
}

// GetPercentLoaded gets the value of PercentLoaded for the instance
func (instance *AppvClientConnectionGroup) GetPropertyPercentLoaded() (value uint16, err error) {
	retValue, err := instance.GetProperty("PercentLoaded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriority sets the value of Priority for the instance
func (instance *AppvClientConnectionGroup) SetPropertyPriority(value uint32) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *AppvClientConnectionGroup) GetPropertyPriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserPending sets the value of UserPending for the instance
func (instance *AppvClientConnectionGroup) SetPropertyUserPending(value bool) (err error) {
	return instance.SetProperty("UserPending", value)
}

// GetUserPending gets the value of UserPending for the instance
func (instance *AppvClientConnectionGroup) GetPropertyUserPending() (value bool, err error) {
	retValue, err := instance.GetProperty("UserPending")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersionId sets the value of VersionId for the instance
func (instance *AppvClientConnectionGroup) SetPropertyVersionId(value string) (err error) {
	return instance.SetProperty("VersionId", value)
}

// GetVersionId gets the value of VersionId for the instance
func (instance *AppvClientConnectionGroup) GetPropertyVersionId() (value string, err error) {
	retValue, err := instance.GetProperty("VersionId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
