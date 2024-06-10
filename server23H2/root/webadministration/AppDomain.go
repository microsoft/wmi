// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AppDomain struct
type AppDomain struct {
	*Object

	//
	ApplicationPath string

	//
	Id string

	//
	IsIdle bool

	//
	PhysicalPath string

	//
	ProcessId uint32

	//
	SiteName string
}

func NewAppDomainEx1(instance *cim.WmiInstance) (newInstance *AppDomain, err error) {
	tmp, err := NewObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AppDomain{
		Object: tmp,
	}
	return
}

func NewAppDomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AppDomain, err error) {
	tmp, err := NewObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AppDomain{
		Object: tmp,
	}
	return
}

// SetApplicationPath sets the value of ApplicationPath for the instance
func (instance *AppDomain) SetPropertyApplicationPath(value string) (err error) {
	return instance.SetProperty("ApplicationPath", (value))
}

// GetApplicationPath gets the value of ApplicationPath for the instance
func (instance *AppDomain) GetPropertyApplicationPath() (value string, err error) {
	retValue, err := instance.GetProperty("ApplicationPath")
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

// SetId sets the value of Id for the instance
func (instance *AppDomain) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *AppDomain) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetIsIdle sets the value of IsIdle for the instance
func (instance *AppDomain) SetPropertyIsIdle(value bool) (err error) {
	return instance.SetProperty("IsIdle", (value))
}

// GetIsIdle gets the value of IsIdle for the instance
func (instance *AppDomain) GetPropertyIsIdle() (value bool, err error) {
	retValue, err := instance.GetProperty("IsIdle")
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

// SetPhysicalPath sets the value of PhysicalPath for the instance
func (instance *AppDomain) SetPropertyPhysicalPath(value string) (err error) {
	return instance.SetProperty("PhysicalPath", (value))
}

// GetPhysicalPath gets the value of PhysicalPath for the instance
func (instance *AppDomain) GetPropertyPhysicalPath() (value string, err error) {
	retValue, err := instance.GetProperty("PhysicalPath")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *AppDomain) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *AppDomain) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
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

// SetSiteName sets the value of SiteName for the instance
func (instance *AppDomain) SetPropertySiteName(value string) (err error) {
	return instance.SetProperty("SiteName", (value))
}

// GetSiteName gets the value of SiteName for the instance
func (instance *AppDomain) GetPropertySiteName() (value string, err error) {
	retValue, err := instance.GetProperty("SiteName")
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

func (instance *AppDomain) Unload() (err error) {
	_, err = instance.InvokeMethodWithReturn("Unload")
	if err != nil {
		return
	}
	return

}
