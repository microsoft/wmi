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

// MDM_ApplicationMinVersion struct
type MDM_ApplicationMinVersion struct {
	cim.WmiInstance

	//
	MinimumPackageVersion string

	//
	PackageArchitecture string

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

// SetMinimumPackageVersion sets the value of MinimumPackageVersion for the instance
func (instance *MDM_ApplicationMinVersion) SetPropertyMinimumPackageVersion(value string) (err error) {
	return instance.SetProperty("MinimumPackageVersion", value)
}

// GetMinimumPackageVersion gets the value of MinimumPackageVersion for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyMinimumPackageVersion() (value string, err error) {
	retValue, err := instance.GetProperty("MinimumPackageVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageArchitecture sets the value of PackageArchitecture for the instance
func (instance *MDM_ApplicationMinVersion) SetPropertyPackageArchitecture(value string) (err error) {
	return instance.SetProperty("PackageArchitecture", value)
}

// GetPackageArchitecture gets the value of PackageArchitecture for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyPackageArchitecture() (value string, err error) {
	retValue, err := instance.GetProperty("PackageArchitecture")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageFullName sets the value of PackageFullName for the instance
func (instance *MDM_ApplicationMinVersion) SetPropertyPackageFullName(value string) (err error) {
	return instance.SetProperty("PackageFullName", value)
}

// GetPackageFullName gets the value of PackageFullName for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyPackageFullName() (value string, err error) {
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
func (instance *MDM_ApplicationMinVersion) SetPropertyPackageName(value string) (err error) {
	return instance.SetProperty("PackageName", value)
}

// GetPackageName gets the value of PackageName for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyPackageName() (value string, err error) {
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
func (instance *MDM_ApplicationMinVersion) SetPropertyPackagePublisher(value string) (err error) {
	return instance.SetProperty("PackagePublisher", value)
}

// GetPackagePublisher gets the value of PackagePublisher for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyPackagePublisher() (value string, err error) {
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
func (instance *MDM_ApplicationMinVersion) SetPropertyPackageVersion(value string) (err error) {
	return instance.SetProperty("PackageVersion", value)
}

// GetPackageVersion gets the value of PackageVersion for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyPackageVersion() (value string, err error) {
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
func (instance *MDM_ApplicationMinVersion) SetPropertyUserSID(value string) (err error) {
	return instance.SetProperty("UserSID", value)
}

// GetUserSID gets the value of UserSID for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyUserSID() (value string, err error) {
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
