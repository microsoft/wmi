// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Personalization struct
type MDM_Personalization struct {
	*cim.WmiInstance

	//
	DesktopImageStatus int32

	//
	DesktopImageUrl string

	//
	InstanceID string

	//
	LockScreenImageStatus int32

	//
	LockScreenImageUrl string

	//
	ParentID string
}

func NewMDM_PersonalizationEx1(instance *cim.WmiInstance) (newInstance *MDM_Personalization, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Personalization{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_PersonalizationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Personalization, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Personalization{
		WmiInstance: tmp,
	}
	return
}

// SetDesktopImageStatus sets the value of DesktopImageStatus for the instance
func (instance *MDM_Personalization) SetPropertyDesktopImageStatus(value int32) (err error) {
	return instance.SetProperty("DesktopImageStatus", value)
}

// GetDesktopImageStatus gets the value of DesktopImageStatus for the instance
func (instance *MDM_Personalization) GetPropertyDesktopImageStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("DesktopImageStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDesktopImageUrl sets the value of DesktopImageUrl for the instance
func (instance *MDM_Personalization) SetPropertyDesktopImageUrl(value string) (err error) {
	return instance.SetProperty("DesktopImageUrl", value)
}

// GetDesktopImageUrl gets the value of DesktopImageUrl for the instance
func (instance *MDM_Personalization) GetPropertyDesktopImageUrl() (value string, err error) {
	retValue, err := instance.GetProperty("DesktopImageUrl")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Personalization) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Personalization) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockScreenImageStatus sets the value of LockScreenImageStatus for the instance
func (instance *MDM_Personalization) SetPropertyLockScreenImageStatus(value int32) (err error) {
	return instance.SetProperty("LockScreenImageStatus", value)
}

// GetLockScreenImageStatus gets the value of LockScreenImageStatus for the instance
func (instance *MDM_Personalization) GetPropertyLockScreenImageStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("LockScreenImageStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockScreenImageUrl sets the value of LockScreenImageUrl for the instance
func (instance *MDM_Personalization) SetPropertyLockScreenImageUrl(value string) (err error) {
	return instance.SetProperty("LockScreenImageUrl", value)
}

// GetLockScreenImageUrl gets the value of LockScreenImageUrl for the instance
func (instance *MDM_Personalization) GetPropertyLockScreenImageUrl() (value string, err error) {
	retValue, err := instance.GetProperty("LockScreenImageUrl")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Personalization) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Personalization) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
