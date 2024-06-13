// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Application struct
type MDM_Application struct {
	*cim.WmiInstance

	//
	Dependencies string

	//
	InstallPath string

	//
	IsBundle bool

	//
	IsDevelopmentMode bool

	//
	IsFramework bool

	//
	IsResourcePackage bool

	//
	PackageFullName string

	//
	PackageName string

	//
	PackagePublisher string

	//
	PackageVersion string

	//
	UserSID string
}

func NewMDM_ApplicationEx1(instance *cim.WmiInstance) (newInstance *MDM_Application, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Application{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_ApplicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Application, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Application{
		WmiInstance: tmp,
	}
	return
}

// SetDependencies sets the value of Dependencies for the instance
func (instance *MDM_Application) SetPropertyDependencies(value string) (err error) {
	return instance.SetProperty("Dependencies", (value))
}

// GetDependencies gets the value of Dependencies for the instance
func (instance *MDM_Application) GetPropertyDependencies() (value string, err error) {
	retValue, err := instance.GetProperty("Dependencies")
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

// SetInstallPath sets the value of InstallPath for the instance
func (instance *MDM_Application) SetPropertyInstallPath(value string) (err error) {
	return instance.SetProperty("InstallPath", (value))
}

// GetInstallPath gets the value of InstallPath for the instance
func (instance *MDM_Application) GetPropertyInstallPath() (value string, err error) {
	retValue, err := instance.GetProperty("InstallPath")
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

// SetIsBundle sets the value of IsBundle for the instance
func (instance *MDM_Application) SetPropertyIsBundle(value bool) (err error) {
	return instance.SetProperty("IsBundle", (value))
}

// GetIsBundle gets the value of IsBundle for the instance
func (instance *MDM_Application) GetPropertyIsBundle() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBundle")
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

// SetIsDevelopmentMode sets the value of IsDevelopmentMode for the instance
func (instance *MDM_Application) SetPropertyIsDevelopmentMode(value bool) (err error) {
	return instance.SetProperty("IsDevelopmentMode", (value))
}

// GetIsDevelopmentMode gets the value of IsDevelopmentMode for the instance
func (instance *MDM_Application) GetPropertyIsDevelopmentMode() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDevelopmentMode")
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

// SetIsFramework sets the value of IsFramework for the instance
func (instance *MDM_Application) SetPropertyIsFramework(value bool) (err error) {
	return instance.SetProperty("IsFramework", (value))
}

// GetIsFramework gets the value of IsFramework for the instance
func (instance *MDM_Application) GetPropertyIsFramework() (value bool, err error) {
	retValue, err := instance.GetProperty("IsFramework")
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

// SetIsResourcePackage sets the value of IsResourcePackage for the instance
func (instance *MDM_Application) SetPropertyIsResourcePackage(value bool) (err error) {
	return instance.SetProperty("IsResourcePackage", (value))
}

// GetIsResourcePackage gets the value of IsResourcePackage for the instance
func (instance *MDM_Application) GetPropertyIsResourcePackage() (value bool, err error) {
	retValue, err := instance.GetProperty("IsResourcePackage")
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

// SetPackageFullName sets the value of PackageFullName for the instance
func (instance *MDM_Application) SetPropertyPackageFullName(value string) (err error) {
	return instance.SetProperty("PackageFullName", (value))
}

// GetPackageFullName gets the value of PackageFullName for the instance
func (instance *MDM_Application) GetPropertyPackageFullName() (value string, err error) {
	retValue, err := instance.GetProperty("PackageFullName")
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

// SetPackageName sets the value of PackageName for the instance
func (instance *MDM_Application) SetPropertyPackageName(value string) (err error) {
	return instance.SetProperty("PackageName", (value))
}

// GetPackageName gets the value of PackageName for the instance
func (instance *MDM_Application) GetPropertyPackageName() (value string, err error) {
	retValue, err := instance.GetProperty("PackageName")
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

// SetPackagePublisher sets the value of PackagePublisher for the instance
func (instance *MDM_Application) SetPropertyPackagePublisher(value string) (err error) {
	return instance.SetProperty("PackagePublisher", (value))
}

// GetPackagePublisher gets the value of PackagePublisher for the instance
func (instance *MDM_Application) GetPropertyPackagePublisher() (value string, err error) {
	retValue, err := instance.GetProperty("PackagePublisher")
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

// SetPackageVersion sets the value of PackageVersion for the instance
func (instance *MDM_Application) SetPropertyPackageVersion(value string) (err error) {
	return instance.SetProperty("PackageVersion", (value))
}

// GetPackageVersion gets the value of PackageVersion for the instance
func (instance *MDM_Application) GetPropertyPackageVersion() (value string, err error) {
	retValue, err := instance.GetProperty("PackageVersion")
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

// SetUserSID sets the value of UserSID for the instance
func (instance *MDM_Application) SetPropertyUserSID(value string) (err error) {
	return instance.SetProperty("UserSID", (value))
}

// GetUserSID gets the value of UserSID for the instance
func (instance *MDM_Application) GetPropertyUserSID() (value string, err error) {
	retValue, err := instance.GetProperty("UserSID")
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
