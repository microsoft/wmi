// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_ApplicationManagementCategory struct
type RSOP_ApplicationManagementCategory struct {
	cim.WmiInstance

	//
	CategoryId string

	//
	CreationTime string

	//
	Name string
}

// SetCategoryId sets the value of CategoryId for the instance
func (instance *RSOP_ApplicationManagementCategory) SetPropertyCategoryId(value string) (err error) {
	return instance.SetProperty("CategoryId", value)
}

// GetCategoryId gets the value of CategoryId for the instance
func (instance *RSOP_ApplicationManagementCategory) GetPropertyCategoryId() (value string, err error) {
	retValue, err := instance.GetProperty("CategoryId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCreationTime sets the value of CreationTime for the instance
func (instance *RSOP_ApplicationManagementCategory) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", value)
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *RSOP_ApplicationManagementCategory) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *RSOP_ApplicationManagementCategory) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *RSOP_ApplicationManagementCategory) GetPropertyName() (value string, err error) {
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
