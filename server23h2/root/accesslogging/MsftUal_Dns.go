// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// MsftUal_Dns struct
type MsftUal_Dns struct {
	*cim.WmiInstance

	// The host name of the client. This is associated with IPAddress.
	HostName string

	// The IP address of a DNS client record. This is associated with hostname.
	IPAddress string

	// The date and time when a DNS client record was last seen in DNS.
	LastSeen string

	// The name of the software parent product, or product line, that is providing User Access Logging data. This is also associated with a RoleName, and a RoleGuid.
	ProductName string

	// The UAL assigned, or registered, GUID representing the server role, or installed product. When a role or product reports its UAL data, this GUID accompanies the payload.
	RoleGuid string

	// The name of the role, component, or sub-product that is providing User Access Logging data. This is also associated with a ProductName, and a RoleGuid.
	RoleName string
}

func NewMsftUal_DnsEx1(instance *cim.WmiInstance) (newInstance *MsftUal_Dns, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MsftUal_Dns{
		WmiInstance: tmp,
	}
	return
}

func NewMsftUal_DnsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftUal_Dns, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftUal_Dns{
		WmiInstance: tmp,
	}
	return
}

// SetHostName sets the value of HostName for the instance
func (instance *MsftUal_Dns) SetPropertyHostName(value string) (err error) {
	return instance.SetProperty("HostName", (value))
}

// GetHostName gets the value of HostName for the instance
func (instance *MsftUal_Dns) GetPropertyHostName() (value string, err error) {
	retValue, err := instance.GetProperty("HostName")
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
func (instance *MsftUal_Dns) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", (value))
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *MsftUal_Dns) GetPropertyIPAddress() (value string, err error) {
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
func (instance *MsftUal_Dns) SetPropertyLastSeen(value string) (err error) {
	return instance.SetProperty("LastSeen", (value))
}

// GetLastSeen gets the value of LastSeen for the instance
func (instance *MsftUal_Dns) GetPropertyLastSeen() (value string, err error) {
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
func (instance *MsftUal_Dns) SetPropertyProductName(value string) (err error) {
	return instance.SetProperty("ProductName", (value))
}

// GetProductName gets the value of ProductName for the instance
func (instance *MsftUal_Dns) GetPropertyProductName() (value string, err error) {
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
func (instance *MsftUal_Dns) SetPropertyRoleGuid(value string) (err error) {
	return instance.SetProperty("RoleGuid", (value))
}

// GetRoleGuid gets the value of RoleGuid for the instance
func (instance *MsftUal_Dns) GetPropertyRoleGuid() (value string, err error) {
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
func (instance *MsftUal_Dns) SetPropertyRoleName(value string) (err error) {
	return instance.SetProperty("RoleName", (value))
}

// GetRoleName gets the value of RoleName for the instance
func (instance *MsftUal_Dns) GetPropertyRoleName() (value string, err error) {
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
