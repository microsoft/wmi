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

// XhtmlConformanceSection struct
type XhtmlConformanceSection struct {
	*ConfigurationSection

	//
	Mode int32
}

func NewXhtmlConformanceSectionEx1(instance *cim.WmiInstance) (newInstance *XhtmlConformanceSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &XhtmlConformanceSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewXhtmlConformanceSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *XhtmlConformanceSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &XhtmlConformanceSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetMode sets the value of Mode for the instance
func (instance *XhtmlConformanceSection) SetPropertyMode(value int32) (err error) {
	return instance.SetProperty("Mode", (value))
}

// GetMode gets the value of Mode for the instance
func (instance *XhtmlConformanceSection) GetPropertyMode() (value int32, err error) {
	retValue, err := instance.GetProperty("Mode")
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
