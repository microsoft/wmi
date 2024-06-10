// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ServerManagerServerComponent struct
type MSFT_ServerManagerServerComponent struct {
	*cim.WmiInstance

	//
	BestPracticeModels []string

	//
	ConfigurationStatus uint8

	//
	Deploys []string

	//
	Description string

	//
	Descriptor interface{}

	//
	DisplayName string

	//
	EventQuery string

	//
	FeatureType uint8

	//
	Installed uint8

	//
	InstallWithParentByDefault bool

	//
	MajorVersion int32

	//
	MinorVersion int32

	//
	MutualExclusions []string

	//
	NonAncestorDependencies []string

	//
	NumericId int32

	//
	OptionalCompanions []MSFT_OptionalCompanion

	//
	ParentName string

	//
	PostInstallDescription string

	//
	PostUninstallDescription string

	//
	SubFeatures []string

	//
	SystemServices []MSFT_ServiceToMonitor

	//
	UniqueName string
}

func NewMSFT_ServerManagerServerComponentEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerManagerServerComponent, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerServerComponent{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerManagerServerComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerManagerServerComponent, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerServerComponent{
		WmiInstance: tmp,
	}
	return
}

// SetBestPracticeModels sets the value of BestPracticeModels for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyBestPracticeModels(value []string) (err error) {
	return instance.SetProperty("BestPracticeModels", (value))
}

// GetBestPracticeModels gets the value of BestPracticeModels for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyBestPracticeModels() (value []string, err error) {
	retValue, err := instance.GetProperty("BestPracticeModels")
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

// SetConfigurationStatus sets the value of ConfigurationStatus for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyConfigurationStatus(value uint8) (err error) {
	return instance.SetProperty("ConfigurationStatus", (value))
}

// GetConfigurationStatus gets the value of ConfigurationStatus for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyConfigurationStatus() (value uint8, err error) {
	retValue, err := instance.GetProperty("ConfigurationStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetDeploys sets the value of Deploys for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyDeploys(value []string) (err error) {
	return instance.SetProperty("Deploys", (value))
}

// GetDeploys gets the value of Deploys for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyDeploys() (value []string, err error) {
	retValue, err := instance.GetProperty("Deploys")
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

// SetDescription sets the value of Description for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetDescriptor sets the value of Descriptor for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyDescriptor(value interface{}) (err error) {
	return instance.SetProperty("Descriptor", (value))
}

// GetDescriptor gets the value of Descriptor for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyDescriptor() (value interface{}, err error) {
	retValue, err := instance.GetProperty("Descriptor")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
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

// SetEventQuery sets the value of EventQuery for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyEventQuery(value string) (err error) {
	return instance.SetProperty("EventQuery", (value))
}

// GetEventQuery gets the value of EventQuery for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyEventQuery() (value string, err error) {
	retValue, err := instance.GetProperty("EventQuery")
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

// SetFeatureType sets the value of FeatureType for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyFeatureType(value uint8) (err error) {
	return instance.SetProperty("FeatureType", (value))
}

// GetFeatureType gets the value of FeatureType for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyFeatureType() (value uint8, err error) {
	retValue, err := instance.GetProperty("FeatureType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetInstalled sets the value of Installed for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyInstalled(value uint8) (err error) {
	return instance.SetProperty("Installed", (value))
}

// GetInstalled gets the value of Installed for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyInstalled() (value uint8, err error) {
	retValue, err := instance.GetProperty("Installed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetInstallWithParentByDefault sets the value of InstallWithParentByDefault for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyInstallWithParentByDefault(value bool) (err error) {
	return instance.SetProperty("InstallWithParentByDefault", (value))
}

// GetInstallWithParentByDefault gets the value of InstallWithParentByDefault for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyInstallWithParentByDefault() (value bool, err error) {
	retValue, err := instance.GetProperty("InstallWithParentByDefault")
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

// SetMajorVersion sets the value of MajorVersion for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyMajorVersion(value int32) (err error) {
	return instance.SetProperty("MajorVersion", (value))
}

// GetMajorVersion gets the value of MajorVersion for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyMajorVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("MajorVersion")
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

// SetMinorVersion sets the value of MinorVersion for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyMinorVersion(value int32) (err error) {
	return instance.SetProperty("MinorVersion", (value))
}

// GetMinorVersion gets the value of MinorVersion for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyMinorVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("MinorVersion")
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

// SetMutualExclusions sets the value of MutualExclusions for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyMutualExclusions(value []string) (err error) {
	return instance.SetProperty("MutualExclusions", (value))
}

// GetMutualExclusions gets the value of MutualExclusions for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyMutualExclusions() (value []string, err error) {
	retValue, err := instance.GetProperty("MutualExclusions")
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

// SetNonAncestorDependencies sets the value of NonAncestorDependencies for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyNonAncestorDependencies(value []string) (err error) {
	return instance.SetProperty("NonAncestorDependencies", (value))
}

// GetNonAncestorDependencies gets the value of NonAncestorDependencies for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyNonAncestorDependencies() (value []string, err error) {
	retValue, err := instance.GetProperty("NonAncestorDependencies")
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

// SetNumericId sets the value of NumericId for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyNumericId(value int32) (err error) {
	return instance.SetProperty("NumericId", (value))
}

// GetNumericId gets the value of NumericId for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyNumericId() (value int32, err error) {
	retValue, err := instance.GetProperty("NumericId")
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

// SetOptionalCompanions sets the value of OptionalCompanions for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyOptionalCompanions(value []MSFT_OptionalCompanion) (err error) {
	return instance.SetProperty("OptionalCompanions", (value))
}

// GetOptionalCompanions gets the value of OptionalCompanions for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyOptionalCompanions() (value []MSFT_OptionalCompanion, err error) {
	retValue, err := instance.GetProperty("OptionalCompanions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_OptionalCompanion)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_OptionalCompanion is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_OptionalCompanion(valuetmp))
	}

	return
}

// SetParentName sets the value of ParentName for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyParentName(value string) (err error) {
	return instance.SetProperty("ParentName", (value))
}

// GetParentName gets the value of ParentName for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyParentName() (value string, err error) {
	retValue, err := instance.GetProperty("ParentName")
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

// SetPostInstallDescription sets the value of PostInstallDescription for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyPostInstallDescription(value string) (err error) {
	return instance.SetProperty("PostInstallDescription", (value))
}

// GetPostInstallDescription gets the value of PostInstallDescription for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyPostInstallDescription() (value string, err error) {
	retValue, err := instance.GetProperty("PostInstallDescription")
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

// SetPostUninstallDescription sets the value of PostUninstallDescription for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyPostUninstallDescription(value string) (err error) {
	return instance.SetProperty("PostUninstallDescription", (value))
}

// GetPostUninstallDescription gets the value of PostUninstallDescription for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyPostUninstallDescription() (value string, err error) {
	retValue, err := instance.GetProperty("PostUninstallDescription")
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

// SetSubFeatures sets the value of SubFeatures for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertySubFeatures(value []string) (err error) {
	return instance.SetProperty("SubFeatures", (value))
}

// GetSubFeatures gets the value of SubFeatures for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertySubFeatures() (value []string, err error) {
	retValue, err := instance.GetProperty("SubFeatures")
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

// SetSystemServices sets the value of SystemServices for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertySystemServices(value []MSFT_ServiceToMonitor) (err error) {
	return instance.SetProperty("SystemServices", (value))
}

// GetSystemServices gets the value of SystemServices for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertySystemServices() (value []MSFT_ServiceToMonitor, err error) {
	retValue, err := instance.GetProperty("SystemServices")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_ServiceToMonitor)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_ServiceToMonitor is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_ServiceToMonitor(valuetmp))
	}

	return
}

// SetUniqueName sets the value of UniqueName for the instance
func (instance *MSFT_ServerManagerServerComponent) SetPropertyUniqueName(value string) (err error) {
	return instance.SetProperty("UniqueName", (value))
}

// GetUniqueName gets the value of UniqueName for the instance
func (instance *MSFT_ServerManagerServerComponent) GetPropertyUniqueName() (value string, err error) {
	retValue, err := instance.GetProperty("UniqueName")
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
