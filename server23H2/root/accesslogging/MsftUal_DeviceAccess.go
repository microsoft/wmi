// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.AccessLogging
//////////////////////////////////////////////
package accesslogging

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MsftUal_DeviceAccess struct
type MsftUal_DeviceAccess struct {
	*cim.WmiInstance

	// The incremental counter of device accesses for a particular client device.
	ActivityCount uint32

	// The date and time when a client IP address is first seen by a role or product.
	FirstSeen string

	// The IP address of a client device that accompanies the UAL payload from installed roles and products.
	IPAddress string

	// The date and time when a client IP address is last seen by a role or product.
	LastSeen string

	// The name of the software parent product, or product line, that is providing User Access Logging data. This is also associated with a RoleName, and a RoleGuid.
	ProductName string

	// The UAL assigned, or registered, GUID representing the server role, or installed product. When a role or product reports its UAL data, this GUID accompanies the payload.
	RoleGuid string

	// The name of the role, component, or sub-product that is providing User Access Logging data. This is also associated with a ProductName, and a RoleGuid.
	RoleName string

	// A unique GUID for a tenant client of an installed role or product which accompanies the UAL data payload, if applicable.
	TenantIdentifier string
}

func NewMsftUal_DeviceAccessEx1(instance *cim.WmiInstance) (newInstance *MsftUal_DeviceAccess, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MsftUal_DeviceAccess{
		WmiInstance: tmp,
	}
	return
}

func NewMsftUal_DeviceAccessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftUal_DeviceAccess, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftUal_DeviceAccess{
		WmiInstance: tmp,
	}
	return
}

// SetActivityCount sets the value of ActivityCount for the instance
func (instance *MsftUal_DeviceAccess) SetPropertyActivityCount(value uint32) (err error) {
	return instance.SetProperty("ActivityCount", (value))
}

// GetActivityCount gets the value of ActivityCount for the instance
func (instance *MsftUal_DeviceAccess) GetPropertyActivityCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActivityCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetFirstSeen sets the value of FirstSeen for the instance
func (instance *MsftUal_DeviceAccess) SetPropertyFirstSeen(value string) (err error) {
	return instance.SetProperty("FirstSeen", (value))
}

// GetFirstSeen gets the value of FirstSeen for the instance
func (instance *MsftUal_DeviceAccess) GetPropertyFirstSeen() (value string, err error) {
	retValue, err := instance.GetProperty("FirstSeen")
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

// SetIPAddress sets the value of IPAddress for the instance
func (instance *MsftUal_DeviceAccess) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", (value))
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *MsftUal_DeviceAccess) GetPropertyIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddress")
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

// SetLastSeen sets the value of LastSeen for the instance
func (instance *MsftUal_DeviceAccess) SetPropertyLastSeen(value string) (err error) {
	return instance.SetProperty("LastSeen", (value))
}

// GetLastSeen gets the value of LastSeen for the instance
func (instance *MsftUal_DeviceAccess) GetPropertyLastSeen() (value string, err error) {
	retValue, err := instance.GetProperty("LastSeen")
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

// SetProductName sets the value of ProductName for the instance
func (instance *MsftUal_DeviceAccess) SetPropertyProductName(value string) (err error) {
	return instance.SetProperty("ProductName", (value))
}

// GetProductName gets the value of ProductName for the instance
func (instance *MsftUal_DeviceAccess) GetPropertyProductName() (value string, err error) {
	retValue, err := instance.GetProperty("ProductName")
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

// SetRoleGuid sets the value of RoleGuid for the instance
func (instance *MsftUal_DeviceAccess) SetPropertyRoleGuid(value string) (err error) {
	return instance.SetProperty("RoleGuid", (value))
}

// GetRoleGuid gets the value of RoleGuid for the instance
func (instance *MsftUal_DeviceAccess) GetPropertyRoleGuid() (value string, err error) {
	retValue, err := instance.GetProperty("RoleGuid")
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

// SetRoleName sets the value of RoleName for the instance
func (instance *MsftUal_DeviceAccess) SetPropertyRoleName(value string) (err error) {
	return instance.SetProperty("RoleName", (value))
}

// GetRoleName gets the value of RoleName for the instance
func (instance *MsftUal_DeviceAccess) GetPropertyRoleName() (value string, err error) {
	retValue, err := instance.GetProperty("RoleName")
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

// SetTenantIdentifier sets the value of TenantIdentifier for the instance
func (instance *MsftUal_DeviceAccess) SetPropertyTenantIdentifier(value string) (err error) {
	return instance.SetProperty("TenantIdentifier", (value))
}

// GetTenantIdentifier gets the value of TenantIdentifier for the instance
func (instance *MsftUal_DeviceAccess) GetPropertyTenantIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("TenantIdentifier")
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
