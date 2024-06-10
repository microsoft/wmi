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

// AppvClientPackage struct
type AppvClientPackage struct {
	*cim.WmiInstance

	//
	Assets []string

	//
	DeploymentMachineData string

	//
	DeploymentUserData string

	//
	GlobalPending bool

	//
	HasAssetIntelligence bool

	//
	InUse bool

	//
	IsPublishedGlobally bool

	//
	IsPublishedToUser bool

	//
	Name string

	//
	PackageId string

	//
	PackageSize uint64

	//
	Path string

	//
	PercentLoaded uint16

	//
	UserConfigurationData string

	//
	UserPending bool

	//
	Version string

	//
	VersionId string
}

func NewAppvClientPackageEx1(instance *cim.WmiInstance) (newInstance *AppvClientPackage, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &AppvClientPackage{
		WmiInstance: tmp,
	}
	return
}

func NewAppvClientPackageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AppvClientPackage, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AppvClientPackage{
		WmiInstance: tmp,
	}
	return
}

// SetAssets sets the value of Assets for the instance
func (instance *AppvClientPackage) SetPropertyAssets(value []string) (err error) {
	return instance.SetProperty("Assets", (value))
}

// GetAssets gets the value of Assets for the instance
func (instance *AppvClientPackage) GetPropertyAssets() (value []string, err error) {
	retValue, err := instance.GetProperty("Assets")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDeploymentMachineData sets the value of DeploymentMachineData for the instance
func (instance *AppvClientPackage) SetPropertyDeploymentMachineData(value string) (err error) {
	return instance.SetProperty("DeploymentMachineData", (value))
}

// GetDeploymentMachineData gets the value of DeploymentMachineData for the instance
func (instance *AppvClientPackage) GetPropertyDeploymentMachineData() (value string, err error) {
	retValue, err := instance.GetProperty("DeploymentMachineData")
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

// SetDeploymentUserData sets the value of DeploymentUserData for the instance
func (instance *AppvClientPackage) SetPropertyDeploymentUserData(value string) (err error) {
	return instance.SetProperty("DeploymentUserData", (value))
}

// GetDeploymentUserData gets the value of DeploymentUserData for the instance
func (instance *AppvClientPackage) GetPropertyDeploymentUserData() (value string, err error) {
	retValue, err := instance.GetProperty("DeploymentUserData")
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

// SetGlobalPending sets the value of GlobalPending for the instance
func (instance *AppvClientPackage) SetPropertyGlobalPending(value bool) (err error) {
	return instance.SetProperty("GlobalPending", (value))
}

// GetGlobalPending gets the value of GlobalPending for the instance
func (instance *AppvClientPackage) GetPropertyGlobalPending() (value bool, err error) {
	retValue, err := instance.GetProperty("GlobalPending")
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

// SetHasAssetIntelligence sets the value of HasAssetIntelligence for the instance
func (instance *AppvClientPackage) SetPropertyHasAssetIntelligence(value bool) (err error) {
	return instance.SetProperty("HasAssetIntelligence", (value))
}

// GetHasAssetIntelligence gets the value of HasAssetIntelligence for the instance
func (instance *AppvClientPackage) GetPropertyHasAssetIntelligence() (value bool, err error) {
	retValue, err := instance.GetProperty("HasAssetIntelligence")
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

// SetInUse sets the value of InUse for the instance
func (instance *AppvClientPackage) SetPropertyInUse(value bool) (err error) {
	return instance.SetProperty("InUse", (value))
}

// GetInUse gets the value of InUse for the instance
func (instance *AppvClientPackage) GetPropertyInUse() (value bool, err error) {
	retValue, err := instance.GetProperty("InUse")
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

// SetIsPublishedGlobally sets the value of IsPublishedGlobally for the instance
func (instance *AppvClientPackage) SetPropertyIsPublishedGlobally(value bool) (err error) {
	return instance.SetProperty("IsPublishedGlobally", (value))
}

// GetIsPublishedGlobally gets the value of IsPublishedGlobally for the instance
func (instance *AppvClientPackage) GetPropertyIsPublishedGlobally() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPublishedGlobally")
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

// SetIsPublishedToUser sets the value of IsPublishedToUser for the instance
func (instance *AppvClientPackage) SetPropertyIsPublishedToUser(value bool) (err error) {
	return instance.SetProperty("IsPublishedToUser", (value))
}

// GetIsPublishedToUser gets the value of IsPublishedToUser for the instance
func (instance *AppvClientPackage) GetPropertyIsPublishedToUser() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPublishedToUser")
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

// SetName sets the value of Name for the instance
func (instance *AppvClientPackage) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *AppvClientPackage) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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
func (instance *AppvClientPackage) SetPropertyPackageId(value string) (err error) {
	return instance.SetProperty("PackageId", (value))
}

// GetPackageId gets the value of PackageId for the instance
func (instance *AppvClientPackage) GetPropertyPackageId() (value string, err error) {
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

// SetPackageSize sets the value of PackageSize for the instance
func (instance *AppvClientPackage) SetPropertyPackageSize(value uint64) (err error) {
	return instance.SetProperty("PackageSize", (value))
}

// GetPackageSize gets the value of PackageSize for the instance
func (instance *AppvClientPackage) GetPropertyPackageSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("PackageSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPath sets the value of Path for the instance
func (instance *AppvClientPackage) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *AppvClientPackage) GetPropertyPath() (value string, err error) {
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

// SetPercentLoaded sets the value of PercentLoaded for the instance
func (instance *AppvClientPackage) SetPropertyPercentLoaded(value uint16) (err error) {
	return instance.SetProperty("PercentLoaded", (value))
}

// GetPercentLoaded gets the value of PercentLoaded for the instance
func (instance *AppvClientPackage) GetPropertyPercentLoaded() (value uint16, err error) {
	retValue, err := instance.GetProperty("PercentLoaded")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetUserConfigurationData sets the value of UserConfigurationData for the instance
func (instance *AppvClientPackage) SetPropertyUserConfigurationData(value string) (err error) {
	return instance.SetProperty("UserConfigurationData", (value))
}

// GetUserConfigurationData gets the value of UserConfigurationData for the instance
func (instance *AppvClientPackage) GetPropertyUserConfigurationData() (value string, err error) {
	retValue, err := instance.GetProperty("UserConfigurationData")
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

// SetUserPending sets the value of UserPending for the instance
func (instance *AppvClientPackage) SetPropertyUserPending(value bool) (err error) {
	return instance.SetProperty("UserPending", (value))
}

// GetUserPending gets the value of UserPending for the instance
func (instance *AppvClientPackage) GetPropertyUserPending() (value bool, err error) {
	retValue, err := instance.GetProperty("UserPending")
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

// SetVersion sets the value of Version for the instance
func (instance *AppvClientPackage) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *AppvClientPackage) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
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

// SetVersionId sets the value of VersionId for the instance
func (instance *AppvClientPackage) SetPropertyVersionId(value string) (err error) {
	return instance.SetProperty("VersionId", (value))
}

// GetVersionId gets the value of VersionId for the instance
func (instance *AppvClientPackage) GetPropertyVersionId() (value string, err error) {
	retValue, err := instance.GetProperty("VersionId")
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
