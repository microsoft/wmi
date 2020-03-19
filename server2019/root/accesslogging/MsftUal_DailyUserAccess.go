// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.AccessLogging
//////////////////////////////////////////////
package accesslogging

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MsftUal_DailyUserAccess struct
type MsftUal_DailyUserAccess struct {
	*cim.WmiInstance

	// The number of accesses of a role, or installed product, on the local server from a unique client user.
	AccessCount uint32

	// The date that a client user accessed a role, or installed product, on the local server.
	AccessDate string

	// The name of the software parent product, or product line, that is providing User Access Logging data. This is also associated with a RoleName, and a RoleGuid.
	ProductName string

	// The UAL assigned, or registered, GUID representing the server role, or installed product. When a role or product reports its UAL data, this GUID accompanies the payload.
	RoleGuid string

	// The name of the role, component, or sub-product that is providing User Access Logging data. This is also associated with a ProductName, and a RoleGuid.
	RoleName string

	// The client user name that accompanies the UAL payload from installed roles and products, if applicable.
	UserName string
}

func NewMsftUal_DailyUserAccessEx1(instance *cim.WmiInstance) (newInstance *MsftUal_DailyUserAccess, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MsftUal_DailyUserAccess{
		WmiInstance: tmp,
	}
	return
}

func NewMsftUal_DailyUserAccessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftUal_DailyUserAccess, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftUal_DailyUserAccess{
		WmiInstance: tmp,
	}
	return
}

// SetAccessCount sets the value of AccessCount for the instance
func (instance *MsftUal_DailyUserAccess) SetPropertyAccessCount(value uint32) (err error) {
	return instance.SetProperty("AccessCount", value)
}

// GetAccessCount gets the value of AccessCount for the instance
func (instance *MsftUal_DailyUserAccess) GetPropertyAccessCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("AccessCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAccessDate sets the value of AccessDate for the instance
func (instance *MsftUal_DailyUserAccess) SetPropertyAccessDate(value string) (err error) {
	return instance.SetProperty("AccessDate", value)
}

// GetAccessDate gets the value of AccessDate for the instance
func (instance *MsftUal_DailyUserAccess) GetPropertyAccessDate() (value string, err error) {
	retValue, err := instance.GetProperty("AccessDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProductName sets the value of ProductName for the instance
func (instance *MsftUal_DailyUserAccess) SetPropertyProductName(value string) (err error) {
	return instance.SetProperty("ProductName", value)
}

// GetProductName gets the value of ProductName for the instance
func (instance *MsftUal_DailyUserAccess) GetPropertyProductName() (value string, err error) {
	retValue, err := instance.GetProperty("ProductName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoleGuid sets the value of RoleGuid for the instance
func (instance *MsftUal_DailyUserAccess) SetPropertyRoleGuid(value string) (err error) {
	return instance.SetProperty("RoleGuid", value)
}

// GetRoleGuid gets the value of RoleGuid for the instance
func (instance *MsftUal_DailyUserAccess) GetPropertyRoleGuid() (value string, err error) {
	retValue, err := instance.GetProperty("RoleGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoleName sets the value of RoleName for the instance
func (instance *MsftUal_DailyUserAccess) SetPropertyRoleName(value string) (err error) {
	return instance.SetProperty("RoleName", value)
}

// GetRoleName gets the value of RoleName for the instance
func (instance *MsftUal_DailyUserAccess) GetPropertyRoleName() (value string, err error) {
	retValue, err := instance.GetProperty("RoleName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserName sets the value of UserName for the instance
func (instance *MsftUal_DailyUserAccess) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", value)
}

// GetUserName gets the value of UserName for the instance
func (instance *MsftUal_DailyUserAccess) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
