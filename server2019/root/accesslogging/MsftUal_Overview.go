// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// MsftUal_Overview struct
type MsftUal_Overview struct {
	*cim.WmiInstance

	// The date and time when a role component or product, becomes operational through UAL.
	FirstSeen string

	// The UAL assigned, or registered, GUID representing the server role, or installed product. When a role or product reports its UAL data, this GUID accompanies the payload.
	GUID string

	// The date and time when a role component, or product, was last operational through UAL.
	LastSeen string

	// The name of the software parent product, or product line, that is providing User Access Logging data. This is also associated with a RoleName, and a RoleGuid.
	ProductName string

	// The name of the role, component, or sub-product that is providing User Access Logging data. This is also associated with a ProductName, and a RoleGuid.
	RoleName string
}

func NewMsftUal_OverviewEx1(instance *cim.WmiInstance) (newInstance *MsftUal_Overview, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MsftUal_Overview{
		WmiInstance: tmp,
	}
	return
}

func NewMsftUal_OverviewEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftUal_Overview, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftUal_Overview{
		WmiInstance: tmp,
	}
	return
}

// SetFirstSeen sets the value of FirstSeen for the instance
func (instance *MsftUal_Overview) SetPropertyFirstSeen(value string) (err error) {
	return instance.SetProperty("FirstSeen", (value))
}

// GetFirstSeen gets the value of FirstSeen for the instance
func (instance *MsftUal_Overview) GetPropertyFirstSeen() (value string, err error) {
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

// SetGUID sets the value of GUID for the instance
func (instance *MsftUal_Overview) SetPropertyGUID(value string) (err error) {
	return instance.SetProperty("GUID", (value))
}

// GetGUID gets the value of GUID for the instance
func (instance *MsftUal_Overview) GetPropertyGUID() (value string, err error) {
	retValue, err := instance.GetProperty("GUID")
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
func (instance *MsftUal_Overview) SetPropertyLastSeen(value string) (err error) {
	return instance.SetProperty("LastSeen", (value))
}

// GetLastSeen gets the value of LastSeen for the instance
func (instance *MsftUal_Overview) GetPropertyLastSeen() (value string, err error) {
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
func (instance *MsftUal_Overview) SetPropertyProductName(value string) (err error) {
	return instance.SetProperty("ProductName", (value))
}

// GetProductName gets the value of ProductName for the instance
func (instance *MsftUal_Overview) GetPropertyProductName() (value string, err error) {
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

// SetRoleName sets the value of RoleName for the instance
func (instance *MsftUal_Overview) SetPropertyRoleName(value string) (err error) {
	return instance.SetProperty("RoleName", (value))
}

// GetRoleName gets the value of RoleName for the instance
func (instance *MsftUal_Overview) GetPropertyRoleName() (value string, err error) {
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
