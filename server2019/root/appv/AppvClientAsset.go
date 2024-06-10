// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Appv
//////////////////////////////////////////////
package appv

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AppvClientAsset struct
type AppvClientAsset struct {
	*cim.WmiInstance

	//
	ChannelCode string

	//
	CM_DSLID string

	//
	InstallDate string

	//
	InstalledLocation string

	//
	Language string

	//
	OsComponent string

	//
	PackageId string

	//
	PackageVersionId string

	//
	ProductID string

	//
	ProductName string

	//
	ProductVersion string

	//
	Publisher string

	//
	RegisteredUser string

	//
	ServicePack string

	//
	SoftwareCode string

	//
	UpgradeCode string

	//
	VersionMajor string

	//
	VersionMinor string
}

func NewAppvClientAssetEx1(instance *cim.WmiInstance) (newInstance *AppvClientAsset, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &AppvClientAsset{
		WmiInstance: tmp,
	}
	return
}

func NewAppvClientAssetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AppvClientAsset, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AppvClientAsset{
		WmiInstance: tmp,
	}
	return
}

// SetChannelCode sets the value of ChannelCode for the instance
func (instance *AppvClientAsset) SetPropertyChannelCode(value string) (err error) {
	return instance.SetProperty("ChannelCode", (value))
}

// GetChannelCode gets the value of ChannelCode for the instance
func (instance *AppvClientAsset) GetPropertyChannelCode() (value string, err error) {
	retValue, err := instance.GetProperty("ChannelCode")
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

// SetCM_DSLID sets the value of CM_DSLID for the instance
func (instance *AppvClientAsset) SetPropertyCM_DSLID(value string) (err error) {
	return instance.SetProperty("CM_DSLID", (value))
}

// GetCM_DSLID gets the value of CM_DSLID for the instance
func (instance *AppvClientAsset) GetPropertyCM_DSLID() (value string, err error) {
	retValue, err := instance.GetProperty("CM_DSLID")
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

// SetInstallDate sets the value of InstallDate for the instance
func (instance *AppvClientAsset) SetPropertyInstallDate(value string) (err error) {
	return instance.SetProperty("InstallDate", (value))
}

// GetInstallDate gets the value of InstallDate for the instance
func (instance *AppvClientAsset) GetPropertyInstallDate() (value string, err error) {
	retValue, err := instance.GetProperty("InstallDate")
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

// SetInstalledLocation sets the value of InstalledLocation for the instance
func (instance *AppvClientAsset) SetPropertyInstalledLocation(value string) (err error) {
	return instance.SetProperty("InstalledLocation", (value))
}

// GetInstalledLocation gets the value of InstalledLocation for the instance
func (instance *AppvClientAsset) GetPropertyInstalledLocation() (value string, err error) {
	retValue, err := instance.GetProperty("InstalledLocation")
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

// SetLanguage sets the value of Language for the instance
func (instance *AppvClientAsset) SetPropertyLanguage(value string) (err error) {
	return instance.SetProperty("Language", (value))
}

// GetLanguage gets the value of Language for the instance
func (instance *AppvClientAsset) GetPropertyLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("Language")
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

// SetOsComponent sets the value of OsComponent for the instance
func (instance *AppvClientAsset) SetPropertyOsComponent(value string) (err error) {
	return instance.SetProperty("OsComponent", (value))
}

// GetOsComponent gets the value of OsComponent for the instance
func (instance *AppvClientAsset) GetPropertyOsComponent() (value string, err error) {
	retValue, err := instance.GetProperty("OsComponent")
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

// SetPackageId sets the value of PackageId for the instance
func (instance *AppvClientAsset) SetPropertyPackageId(value string) (err error) {
	return instance.SetProperty("PackageId", (value))
}

// GetPackageId gets the value of PackageId for the instance
func (instance *AppvClientAsset) GetPropertyPackageId() (value string, err error) {
	retValue, err := instance.GetProperty("PackageId")
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

// SetPackageVersionId sets the value of PackageVersionId for the instance
func (instance *AppvClientAsset) SetPropertyPackageVersionId(value string) (err error) {
	return instance.SetProperty("PackageVersionId", (value))
}

// GetPackageVersionId gets the value of PackageVersionId for the instance
func (instance *AppvClientAsset) GetPropertyPackageVersionId() (value string, err error) {
	retValue, err := instance.GetProperty("PackageVersionId")
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

// SetProductID sets the value of ProductID for the instance
func (instance *AppvClientAsset) SetPropertyProductID(value string) (err error) {
	return instance.SetProperty("ProductID", (value))
}

// GetProductID gets the value of ProductID for the instance
func (instance *AppvClientAsset) GetPropertyProductID() (value string, err error) {
	retValue, err := instance.GetProperty("ProductID")
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
func (instance *AppvClientAsset) SetPropertyProductName(value string) (err error) {
	return instance.SetProperty("ProductName", (value))
}

// GetProductName gets the value of ProductName for the instance
func (instance *AppvClientAsset) GetPropertyProductName() (value string, err error) {
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

// SetProductVersion sets the value of ProductVersion for the instance
func (instance *AppvClientAsset) SetPropertyProductVersion(value string) (err error) {
	return instance.SetProperty("ProductVersion", (value))
}

// GetProductVersion gets the value of ProductVersion for the instance
func (instance *AppvClientAsset) GetPropertyProductVersion() (value string, err error) {
	retValue, err := instance.GetProperty("ProductVersion")
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

// SetPublisher sets the value of Publisher for the instance
func (instance *AppvClientAsset) SetPropertyPublisher(value string) (err error) {
	return instance.SetProperty("Publisher", (value))
}

// GetPublisher gets the value of Publisher for the instance
func (instance *AppvClientAsset) GetPropertyPublisher() (value string, err error) {
	retValue, err := instance.GetProperty("Publisher")
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

// SetRegisteredUser sets the value of RegisteredUser for the instance
func (instance *AppvClientAsset) SetPropertyRegisteredUser(value string) (err error) {
	return instance.SetProperty("RegisteredUser", (value))
}

// GetRegisteredUser gets the value of RegisteredUser for the instance
func (instance *AppvClientAsset) GetPropertyRegisteredUser() (value string, err error) {
	retValue, err := instance.GetProperty("RegisteredUser")
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

// SetServicePack sets the value of ServicePack for the instance
func (instance *AppvClientAsset) SetPropertyServicePack(value string) (err error) {
	return instance.SetProperty("ServicePack", (value))
}

// GetServicePack gets the value of ServicePack for the instance
func (instance *AppvClientAsset) GetPropertyServicePack() (value string, err error) {
	retValue, err := instance.GetProperty("ServicePack")
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

// SetSoftwareCode sets the value of SoftwareCode for the instance
func (instance *AppvClientAsset) SetPropertySoftwareCode(value string) (err error) {
	return instance.SetProperty("SoftwareCode", (value))
}

// GetSoftwareCode gets the value of SoftwareCode for the instance
func (instance *AppvClientAsset) GetPropertySoftwareCode() (value string, err error) {
	retValue, err := instance.GetProperty("SoftwareCode")
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

// SetUpgradeCode sets the value of UpgradeCode for the instance
func (instance *AppvClientAsset) SetPropertyUpgradeCode(value string) (err error) {
	return instance.SetProperty("UpgradeCode", (value))
}

// GetUpgradeCode gets the value of UpgradeCode for the instance
func (instance *AppvClientAsset) GetPropertyUpgradeCode() (value string, err error) {
	retValue, err := instance.GetProperty("UpgradeCode")
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

// SetVersionMajor sets the value of VersionMajor for the instance
func (instance *AppvClientAsset) SetPropertyVersionMajor(value string) (err error) {
	return instance.SetProperty("VersionMajor", (value))
}

// GetVersionMajor gets the value of VersionMajor for the instance
func (instance *AppvClientAsset) GetPropertyVersionMajor() (value string, err error) {
	retValue, err := instance.GetProperty("VersionMajor")
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

// SetVersionMinor sets the value of VersionMinor for the instance
func (instance *AppvClientAsset) SetPropertyVersionMinor(value string) (err error) {
	return instance.SetProperty("VersionMinor", (value))
}

// GetVersionMinor gets the value of VersionMinor for the instance
func (instance *AppvClientAsset) GetPropertyVersionMinor() (value string, err error) {
	retValue, err := instance.GetProperty("VersionMinor")
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
