// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Application struct
type MDM_Application struct {
	cim.WmiInstance

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

// SetDependencies sets the value of Dependencies for the instance
func (instance *MDM_Application) SetPropertyDependencies(value string) (err error) {
	return instance.SetProperty("Dependencies", value)
}

// GetDependencies gets the value of Dependencies for the instance
func (instance *MDM_Application) GetPropertyDependencies() (value string, err error) {
	retValue, err := instance.GetProperty("Dependencies")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstallPath sets the value of InstallPath for the instance
func (instance *MDM_Application) SetPropertyInstallPath(value string) (err error) {
	return instance.SetProperty("InstallPath", value)
}

// GetInstallPath gets the value of InstallPath for the instance
func (instance *MDM_Application) GetPropertyInstallPath() (value string, err error) {
	retValue, err := instance.GetProperty("InstallPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsBundle sets the value of IsBundle for the instance
func (instance *MDM_Application) SetPropertyIsBundle(value bool) (err error) {
	return instance.SetProperty("IsBundle", value)
}

// GetIsBundle gets the value of IsBundle for the instance
func (instance *MDM_Application) GetPropertyIsBundle() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBundle")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsDevelopmentMode sets the value of IsDevelopmentMode for the instance
func (instance *MDM_Application) SetPropertyIsDevelopmentMode(value bool) (err error) {
	return instance.SetProperty("IsDevelopmentMode", value)
}

// GetIsDevelopmentMode gets the value of IsDevelopmentMode for the instance
func (instance *MDM_Application) GetPropertyIsDevelopmentMode() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDevelopmentMode")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsFramework sets the value of IsFramework for the instance
func (instance *MDM_Application) SetPropertyIsFramework(value bool) (err error) {
	return instance.SetProperty("IsFramework", value)
}

// GetIsFramework gets the value of IsFramework for the instance
func (instance *MDM_Application) GetPropertyIsFramework() (value bool, err error) {
	retValue, err := instance.GetProperty("IsFramework")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsResourcePackage sets the value of IsResourcePackage for the instance
func (instance *MDM_Application) SetPropertyIsResourcePackage(value bool) (err error) {
	return instance.SetProperty("IsResourcePackage", value)
}

// GetIsResourcePackage gets the value of IsResourcePackage for the instance
func (instance *MDM_Application) GetPropertyIsResourcePackage() (value bool, err error) {
	retValue, err := instance.GetProperty("IsResourcePackage")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageFullName sets the value of PackageFullName for the instance
func (instance *MDM_Application) SetPropertyPackageFullName(value string) (err error) {
	return instance.SetProperty("PackageFullName", value)
}

// GetPackageFullName gets the value of PackageFullName for the instance
func (instance *MDM_Application) GetPropertyPackageFullName() (value string, err error) {
	retValue, err := instance.GetProperty("PackageFullName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageName sets the value of PackageName for the instance
func (instance *MDM_Application) SetPropertyPackageName(value string) (err error) {
	return instance.SetProperty("PackageName", value)
}

// GetPackageName gets the value of PackageName for the instance
func (instance *MDM_Application) GetPropertyPackageName() (value string, err error) {
	retValue, err := instance.GetProperty("PackageName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackagePublisher sets the value of PackagePublisher for the instance
func (instance *MDM_Application) SetPropertyPackagePublisher(value string) (err error) {
	return instance.SetProperty("PackagePublisher", value)
}

// GetPackagePublisher gets the value of PackagePublisher for the instance
func (instance *MDM_Application) GetPropertyPackagePublisher() (value string, err error) {
	retValue, err := instance.GetProperty("PackagePublisher")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageVersion sets the value of PackageVersion for the instance
func (instance *MDM_Application) SetPropertyPackageVersion(value string) (err error) {
	return instance.SetProperty("PackageVersion", value)
}

// GetPackageVersion gets the value of PackageVersion for the instance
func (instance *MDM_Application) GetPropertyPackageVersion() (value string, err error) {
	retValue, err := instance.GetProperty("PackageVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserSID sets the value of UserSID for the instance
func (instance *MDM_Application) SetPropertyUserSID(value string) (err error) {
	return instance.SetProperty("UserSID", value)
}

// GetUserSID gets the value of UserSID for the instance
func (instance *MDM_Application) GetPropertyUserSID() (value string, err error) {
	retValue, err := instance.GetProperty("UserSID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
