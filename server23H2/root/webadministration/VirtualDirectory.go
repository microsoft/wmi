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

// VirtualDirectory struct
type VirtualDirectory struct {
	*ConfiguredObject

	//
	AllowSubDirConfig bool

	//
	ApplicationPath string

	//
	LogonMethod int32

	//
	Password string

	//
	Path string

	//
	PhysicalPath string

	//
	SiteName string

	//
	UserName string
}

func NewVirtualDirectoryEx1(instance *cim.WmiInstance) (newInstance *VirtualDirectory, err error) {
	tmp, err := NewConfiguredObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VirtualDirectory{
		ConfiguredObject: tmp,
	}
	return
}

func NewVirtualDirectoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VirtualDirectory, err error) {
	tmp, err := NewConfiguredObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VirtualDirectory{
		ConfiguredObject: tmp,
	}
	return
}

// SetAllowSubDirConfig sets the value of AllowSubDirConfig for the instance
func (instance *VirtualDirectory) SetPropertyAllowSubDirConfig(value bool) (err error) {
	return instance.SetProperty("AllowSubDirConfig", (value))
}

// GetAllowSubDirConfig gets the value of AllowSubDirConfig for the instance
func (instance *VirtualDirectory) GetPropertyAllowSubDirConfig() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowSubDirConfig")
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

// SetApplicationPath sets the value of ApplicationPath for the instance
func (instance *VirtualDirectory) SetPropertyApplicationPath(value string) (err error) {
	return instance.SetProperty("ApplicationPath", (value))
}

// GetApplicationPath gets the value of ApplicationPath for the instance
func (instance *VirtualDirectory) GetPropertyApplicationPath() (value string, err error) {
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

// SetLogonMethod sets the value of LogonMethod for the instance
func (instance *VirtualDirectory) SetPropertyLogonMethod(value int32) (err error) {
	return instance.SetProperty("LogonMethod", (value))
}

// GetLogonMethod gets the value of LogonMethod for the instance
func (instance *VirtualDirectory) GetPropertyLogonMethod() (value int32, err error) {
	retValue, err := instance.GetProperty("LogonMethod")
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

// SetPassword sets the value of Password for the instance
func (instance *VirtualDirectory) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *VirtualDirectory) GetPropertyPassword() (value string, err error) {
	retValue, err := instance.GetProperty("Password")
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

// SetPath sets the value of Path for the instance
func (instance *VirtualDirectory) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *VirtualDirectory) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetPhysicalPath sets the value of PhysicalPath for the instance
func (instance *VirtualDirectory) SetPropertyPhysicalPath(value string) (err error) {
	return instance.SetProperty("PhysicalPath", (value))
}

// GetPhysicalPath gets the value of PhysicalPath for the instance
func (instance *VirtualDirectory) GetPropertyPhysicalPath() (value string, err error) {
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

// SetSiteName sets the value of SiteName for the instance
func (instance *VirtualDirectory) SetPropertySiteName(value string) (err error) {
	return instance.SetProperty("SiteName", (value))
}

// GetSiteName gets the value of SiteName for the instance
func (instance *VirtualDirectory) GetPropertySiteName() (value string, err error) {
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

// SetUserName sets the value of UserName for the instance
func (instance *VirtualDirectory) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *VirtualDirectory) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
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

//

// <param name="ApplicationPath" type="string "></param>
// <param name="PhysicalPath" type="string "></param>
// <param name="SiteName" type="string "></param>
// <param name="VirtualDirectoryPath" type="string "></param>
func (instance *VirtualDirectory) Create( /* IN */ VirtualDirectoryPath string,
	/* IN */ ApplicationPath string,
	/* OPTIONAL IN */ PhysicalPath string,
	/* IN */ SiteName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("Create", VirtualDirectoryPath, ApplicationPath, PhysicalPath, SiteName)
	if err != nil {
		return
	}
	return

}

//

// <param name="PropertyName" type="string "></param>
func (instance *VirtualDirectory) RevertToParent( /* OPTIONAL IN */ PropertyName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("RevertToParent", PropertyName)
	if err != nil {
		return
	}
	return

}
