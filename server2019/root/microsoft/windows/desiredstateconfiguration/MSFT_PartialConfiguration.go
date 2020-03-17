// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

// MSFT_PartialConfiguration struct
type MSFT_PartialConfiguration struct {
	OMI_MetaConfigurationResource

	//
	ConfigurationSource []string

	//
	DependsOn []string

	//
	Description string

	//
	ExclusiveResources []string

	//
	RefreshMode string

	//
	ResourceModuleSource []string
}

// SetConfigurationSource sets the value of ConfigurationSource for the instance
func (instance *MSFT_PartialConfiguration) SetPropertyConfigurationSource(value []string) (err error) {
	return instance.SetProperty("ConfigurationSource", value)
}

// GetConfigurationSource gets the value of ConfigurationSource for the instance
func (instance *MSFT_PartialConfiguration) GetPropertyConfigurationSource() (value []string, err error) {
	retValue, err := instance.GetProperty("ConfigurationSource")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDependsOn sets the value of DependsOn for the instance
func (instance *MSFT_PartialConfiguration) SetPropertyDependsOn(value []string) (err error) {
	return instance.SetProperty("DependsOn", value)
}

// GetDependsOn gets the value of DependsOn for the instance
func (instance *MSFT_PartialConfiguration) GetPropertyDependsOn() (value []string, err error) {
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

// SetDescription sets the value of Description for the instance
func (instance *MSFT_PartialConfiguration) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_PartialConfiguration) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExclusiveResources sets the value of ExclusiveResources for the instance
func (instance *MSFT_PartialConfiguration) SetPropertyExclusiveResources(value []string) (err error) {
	return instance.SetProperty("ExclusiveResources", value)
}

// GetExclusiveResources gets the value of ExclusiveResources for the instance
func (instance *MSFT_PartialConfiguration) GetPropertyExclusiveResources() (value []string, err error) {
	retValue, err := instance.GetProperty("ExclusiveResources")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRefreshMode sets the value of RefreshMode for the instance
func (instance *MSFT_PartialConfiguration) SetPropertyRefreshMode(value string) (err error) {
	return instance.SetProperty("RefreshMode", value)
}

// GetRefreshMode gets the value of RefreshMode for the instance
func (instance *MSFT_PartialConfiguration) GetPropertyRefreshMode() (value string, err error) {
	retValue, err := instance.GetProperty("RefreshMode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceModuleSource sets the value of ResourceModuleSource for the instance
func (instance *MSFT_PartialConfiguration) SetPropertyResourceModuleSource(value []string) (err error) {
	return instance.SetProperty("ResourceModuleSource", value)
}

// GetResourceModuleSource gets the value of ResourceModuleSource for the instance
func (instance *MSFT_PartialConfiguration) GetPropertyResourceModuleSource() (value []string, err error) {
	retValue, err := instance.GetProperty("ResourceModuleSource")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
