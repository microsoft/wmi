// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// OMI_BaseResource struct
type OMI_BaseResource struct {
	*cim.WmiInstance

	//
	ConfigurationName string

	//
	DependsOn []string

	//
	ModuleName string

	//
	ModuleVersion string

	//
	PsDscRunAsCredential MSFT_Credential

	//
	ResourceId string

	//
	SourceInfo string
}

func NewOMI_BaseResourceEx1(instance *cim.WmiInstance) (newInstance *OMI_BaseResource, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &OMI_BaseResource{
		WmiInstance: tmp,
	}
	return
}

func NewOMI_BaseResourceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *OMI_BaseResource, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &OMI_BaseResource{
		WmiInstance: tmp,
	}
	return
}

// SetConfigurationName sets the value of ConfigurationName for the instance
func (instance *OMI_BaseResource) SetPropertyConfigurationName(value string) (err error) {
	return instance.SetProperty("ConfigurationName", value)
}

// GetConfigurationName gets the value of ConfigurationName for the instance
func (instance *OMI_BaseResource) GetPropertyConfigurationName() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigurationName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDependsOn sets the value of DependsOn for the instance
func (instance *OMI_BaseResource) SetPropertyDependsOn(value []string) (err error) {
	return instance.SetProperty("DependsOn", value)
}

// GetDependsOn gets the value of DependsOn for the instance
func (instance *OMI_BaseResource) GetPropertyDependsOn() (value []string, err error) {
	retValue, err := instance.GetProperty("DependsOn")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetModuleName sets the value of ModuleName for the instance
func (instance *OMI_BaseResource) SetPropertyModuleName(value string) (err error) {
	return instance.SetProperty("ModuleName", value)
}

// GetModuleName gets the value of ModuleName for the instance
func (instance *OMI_BaseResource) GetPropertyModuleName() (value string, err error) {
	retValue, err := instance.GetProperty("ModuleName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetModuleVersion sets the value of ModuleVersion for the instance
func (instance *OMI_BaseResource) SetPropertyModuleVersion(value string) (err error) {
	return instance.SetProperty("ModuleVersion", value)
}

// GetModuleVersion gets the value of ModuleVersion for the instance
func (instance *OMI_BaseResource) GetPropertyModuleVersion() (value string, err error) {
	retValue, err := instance.GetProperty("ModuleVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPsDscRunAsCredential sets the value of PsDscRunAsCredential for the instance
func (instance *OMI_BaseResource) SetPropertyPsDscRunAsCredential(value MSFT_Credential) (err error) {
	return instance.SetProperty("PsDscRunAsCredential", value)
}

// GetPsDscRunAsCredential gets the value of PsDscRunAsCredential for the instance
func (instance *OMI_BaseResource) GetPropertyPsDscRunAsCredential() (value MSFT_Credential, err error) {
	retValue, err := instance.GetProperty("PsDscRunAsCredential")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Credential)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceId sets the value of ResourceId for the instance
func (instance *OMI_BaseResource) SetPropertyResourceId(value string) (err error) {
	return instance.SetProperty("ResourceId", value)
}

// GetResourceId gets the value of ResourceId for the instance
func (instance *OMI_BaseResource) GetPropertyResourceId() (value string, err error) {
	retValue, err := instance.GetProperty("ResourceId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceInfo sets the value of SourceInfo for the instance
func (instance *OMI_BaseResource) SetPropertySourceInfo(value string) (err error) {
	return instance.SetProperty("SourceInfo", value)
}

// GetSourceInfo gets the value of SourceInfo for the instance
func (instance *OMI_BaseResource) GetPropertySourceInfo() (value string, err error) {
	retValue, err := instance.GetProperty("SourceInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
