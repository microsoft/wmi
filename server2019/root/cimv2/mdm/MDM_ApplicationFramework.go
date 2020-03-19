// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_ApplicationFramework struct
type MDM_ApplicationFramework struct {
	*cim.WmiInstance

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

func NewMDM_ApplicationFrameworkEx1(instance *cim.WmiInstance) (newInstance *MDM_ApplicationFramework, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_ApplicationFramework{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_ApplicationFrameworkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_ApplicationFramework, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_ApplicationFramework{
		WmiInstance: tmp,
	}
	return
}

// SetMinimumPackageVersion sets the value of MinimumPackageVersion for the instance
func (instance *MDM_ApplicationFramework) SetPropertyMinimumPackageVersion(value string) (err error) {
	return instance.SetProperty("MinimumPackageVersion", value)
}

// GetMinimumPackageVersion gets the value of MinimumPackageVersion for the instance
func (instance *MDM_ApplicationFramework) GetPropertyMinimumPackageVersion() (value string, err error) {
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
func (instance *MDM_ApplicationFramework) SetPropertyPackageArchitecture(value string) (err error) {
	return instance.SetProperty("PackageArchitecture", value)
}

// GetPackageArchitecture gets the value of PackageArchitecture for the instance
func (instance *MDM_ApplicationFramework) GetPropertyPackageArchitecture() (value string, err error) {
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
func (instance *MDM_ApplicationFramework) SetPropertyPackageFullName(value string) (err error) {
	return instance.SetProperty("PackageFullName", value)
}

// GetPackageFullName gets the value of PackageFullName for the instance
func (instance *MDM_ApplicationFramework) GetPropertyPackageFullName() (value string, err error) {
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
func (instance *MDM_ApplicationFramework) SetPropertyPackageName(value string) (err error) {
	return instance.SetProperty("PackageName", value)
}

// GetPackageName gets the value of PackageName for the instance
func (instance *MDM_ApplicationFramework) GetPropertyPackageName() (value string, err error) {
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
func (instance *MDM_ApplicationFramework) SetPropertyPackagePublisher(value string) (err error) {
	return instance.SetProperty("PackagePublisher", value)
}

// GetPackagePublisher gets the value of PackagePublisher for the instance
func (instance *MDM_ApplicationFramework) GetPropertyPackagePublisher() (value string, err error) {
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
func (instance *MDM_ApplicationFramework) SetPropertyPackageVersion(value string) (err error) {
	return instance.SetProperty("PackageVersion", value)
}

// GetPackageVersion gets the value of PackageVersion for the instance
func (instance *MDM_ApplicationFramework) GetPropertyPackageVersion() (value string, err error) {
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
func (instance *MDM_ApplicationFramework) SetPropertyUserSID(value string) (err error) {
	return instance.SetProperty("UserSID", value)
}

// GetUserSID gets the value of UserSID for the instance
func (instance *MDM_ApplicationFramework) GetPropertyUserSID() (value string, err error) {
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
