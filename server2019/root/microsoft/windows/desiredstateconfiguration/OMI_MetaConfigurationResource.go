// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// OMI_MetaConfigurationResource struct
type OMI_MetaConfigurationResource struct {
	cim.WmiInstance

	//
	ResourceId string

	//
	SourceInfo string
}

// SetResourceId sets the value of ResourceId for the instance
func (instance *OMI_MetaConfigurationResource) SetPropertyResourceId(value string) (err error) {
	return instance.SetProperty("ResourceId", value)
}

// GetResourceId gets the value of ResourceId for the instance
func (instance *OMI_MetaConfigurationResource) GetPropertyResourceId() (value string, err error) {
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
func (instance *OMI_MetaConfigurationResource) SetPropertySourceInfo(value string) (err error) {
	return instance.SetProperty("SourceInfo", value)
}

// GetSourceInfo gets the value of SourceInfo for the instance
func (instance *OMI_MetaConfigurationResource) GetPropertySourceInfo() (value string, err error) {
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
