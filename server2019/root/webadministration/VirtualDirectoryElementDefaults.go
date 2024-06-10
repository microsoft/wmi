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

// VirtualDirectoryElementDefaults struct
type VirtualDirectoryElementDefaults struct {
	*EmbeddedObject

	//
	AllowSubDirConfig bool

	//
	LogonMethod int32

	//
	Password string

	//
	Path string

	//
	PhysicalPath string

	//
	UserName string
}

func NewVirtualDirectoryElementDefaultsEx1(instance *cim.WmiInstance) (newInstance *VirtualDirectoryElementDefaults, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VirtualDirectoryElementDefaults{
		EmbeddedObject: tmp,
	}
	return
}

func NewVirtualDirectoryElementDefaultsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VirtualDirectoryElementDefaults, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VirtualDirectoryElementDefaults{
		EmbeddedObject: tmp,
	}
	return
}

// SetAllowSubDirConfig sets the value of AllowSubDirConfig for the instance
func (instance *VirtualDirectoryElementDefaults) SetPropertyAllowSubDirConfig(value bool) (err error) {
	return instance.SetProperty("AllowSubDirConfig", (value))
}

// GetAllowSubDirConfig gets the value of AllowSubDirConfig for the instance
func (instance *VirtualDirectoryElementDefaults) GetPropertyAllowSubDirConfig() (value bool, err error) {
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

// SetLogonMethod sets the value of LogonMethod for the instance
func (instance *VirtualDirectoryElementDefaults) SetPropertyLogonMethod(value int32) (err error) {
	return instance.SetProperty("LogonMethod", (value))
}

// GetLogonMethod gets the value of LogonMethod for the instance
func (instance *VirtualDirectoryElementDefaults) GetPropertyLogonMethod() (value int32, err error) {
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
func (instance *VirtualDirectoryElementDefaults) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *VirtualDirectoryElementDefaults) GetPropertyPassword() (value string, err error) {
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
func (instance *VirtualDirectoryElementDefaults) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *VirtualDirectoryElementDefaults) GetPropertyPath() (value string, err error) {
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
func (instance *VirtualDirectoryElementDefaults) SetPropertyPhysicalPath(value string) (err error) {
	return instance.SetProperty("PhysicalPath", (value))
}

// GetPhysicalPath gets the value of PhysicalPath for the instance
func (instance *VirtualDirectoryElementDefaults) GetPropertyPhysicalPath() (value string, err error) {
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

// SetUserName sets the value of UserName for the instance
func (instance *VirtualDirectoryElementDefaults) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *VirtualDirectoryElementDefaults) GetPropertyUserName() (value string, err error) {
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
