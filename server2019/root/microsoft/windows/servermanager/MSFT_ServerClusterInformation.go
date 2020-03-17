// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServerClusterInformation struct
type MSFT_ServerClusterInformation struct {
	cim.WmiInstance

	//
	GroupType int32

	//
	Name string

	//
	ObjectType uint8
}

// SetGroupType sets the value of GroupType for the instance
func (instance *MSFT_ServerClusterInformation) SetPropertyGroupType(value int32) (err error) {
	return instance.SetProperty("GroupType", value)
}

// GetGroupType gets the value of GroupType for the instance
func (instance *MSFT_ServerClusterInformation) GetPropertyGroupType() (value int32, err error) {
	retValue, err := instance.GetProperty("GroupType")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_ServerClusterInformation) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ServerClusterInformation) GetPropertyName() (value string, err error) {
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

// SetObjectType sets the value of ObjectType for the instance
func (instance *MSFT_ServerClusterInformation) SetPropertyObjectType(value uint8) (err error) {
	return instance.SetProperty("ObjectType", value)
}

// GetObjectType gets the value of ObjectType for the instance
func (instance *MSFT_ServerClusterInformation) GetPropertyObjectType() (value uint8, err error) {
	retValue, err := instance.GetProperty("ObjectType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
