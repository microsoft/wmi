// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// OMI_MetaConfigurationResource struct
type OMI_MetaConfigurationResource struct {
	*cim.WmiInstance

	//
	ResourceId string

	//
	SourceInfo string
}

func NewOMI_MetaConfigurationResourceEx1(instance *cim.WmiInstance) (newInstance *OMI_MetaConfigurationResource, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &OMI_MetaConfigurationResource{
		WmiInstance: tmp,
	}
	return
}

func NewOMI_MetaConfigurationResourceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *OMI_MetaConfigurationResource, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &OMI_MetaConfigurationResource{
		WmiInstance: tmp,
	}
	return
}

// SetResourceId sets the value of ResourceId for the instance
func (instance *OMI_MetaConfigurationResource) SetPropertyResourceId(value string) (err error) {
	return instance.SetProperty("ResourceId", (value))
}

// GetResourceId gets the value of ResourceId for the instance
func (instance *OMI_MetaConfigurationResource) GetPropertyResourceId() (value string, err error) {
	retValue, err := instance.GetProperty("ResourceId")
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

// SetSourceInfo sets the value of SourceInfo for the instance
func (instance *OMI_MetaConfigurationResource) SetPropertySourceInfo(value string) (err error) {
	return instance.SetProperty("SourceInfo", (value))
}

// GetSourceInfo gets the value of SourceInfo for the instance
func (instance *OMI_MetaConfigurationResource) GetPropertySourceInfo() (value string, err error) {
	retValue, err := instance.GetProperty("SourceInfo")
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
