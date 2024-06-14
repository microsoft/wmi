// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DeviceFiltersSection struct
type DeviceFiltersSection struct {
	*ConfigurationSectionWithCollection

	//
	DeviceFilters []DeviceFilter
}

func NewDeviceFiltersSectionEx1(instance *cim.WmiInstance) (newInstance *DeviceFiltersSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DeviceFiltersSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewDeviceFiltersSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DeviceFiltersSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DeviceFiltersSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetDeviceFilters sets the value of DeviceFilters for the instance
func (instance *DeviceFiltersSection) SetPropertyDeviceFilters(value []DeviceFilter) (err error) {
	return instance.SetProperty("DeviceFilters", (value))
}

// GetDeviceFilters gets the value of DeviceFilters for the instance
func (instance *DeviceFiltersSection) GetPropertyDeviceFilters() (value []DeviceFilter, err error) {
	retValue, err := instance.GetProperty("DeviceFilters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(DeviceFilter)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " DeviceFilter is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DeviceFilter(valuetmp))
	}

	return
}
