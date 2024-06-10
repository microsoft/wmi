// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// GlobalizationSection struct
type GlobalizationSection struct {
	*ConfigurationSection

	//
	Culture string

	//
	EnableBestFitResponseEncoding bool

	//
	EnableClientBasedCulture bool

	//
	FileEncoding string

	//
	RequestEncoding string

	//
	ResourceProviderFactoryType string

	//
	ResponseEncoding string

	//
	ResponseHeaderEncoding string

	//
	UiCulture string
}

func NewGlobalizationSectionEx1(instance *cim.WmiInstance) (newInstance *GlobalizationSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &GlobalizationSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewGlobalizationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *GlobalizationSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &GlobalizationSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetCulture sets the value of Culture for the instance
func (instance *GlobalizationSection) SetPropertyCulture(value string) (err error) {
	return instance.SetProperty("Culture", (value))
}

// GetCulture gets the value of Culture for the instance
func (instance *GlobalizationSection) GetPropertyCulture() (value string, err error) {
	retValue, err := instance.GetProperty("Culture")
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

// SetEnableBestFitResponseEncoding sets the value of EnableBestFitResponseEncoding for the instance
func (instance *GlobalizationSection) SetPropertyEnableBestFitResponseEncoding(value bool) (err error) {
	return instance.SetProperty("EnableBestFitResponseEncoding", (value))
}

// GetEnableBestFitResponseEncoding gets the value of EnableBestFitResponseEncoding for the instance
func (instance *GlobalizationSection) GetPropertyEnableBestFitResponseEncoding() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableBestFitResponseEncoding")
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

// SetEnableClientBasedCulture sets the value of EnableClientBasedCulture for the instance
func (instance *GlobalizationSection) SetPropertyEnableClientBasedCulture(value bool) (err error) {
	return instance.SetProperty("EnableClientBasedCulture", (value))
}

// GetEnableClientBasedCulture gets the value of EnableClientBasedCulture for the instance
func (instance *GlobalizationSection) GetPropertyEnableClientBasedCulture() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableClientBasedCulture")
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

// SetFileEncoding sets the value of FileEncoding for the instance
func (instance *GlobalizationSection) SetPropertyFileEncoding(value string) (err error) {
	return instance.SetProperty("FileEncoding", (value))
}

// GetFileEncoding gets the value of FileEncoding for the instance
func (instance *GlobalizationSection) GetPropertyFileEncoding() (value string, err error) {
	retValue, err := instance.GetProperty("FileEncoding")
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

// SetRequestEncoding sets the value of RequestEncoding for the instance
func (instance *GlobalizationSection) SetPropertyRequestEncoding(value string) (err error) {
	return instance.SetProperty("RequestEncoding", (value))
}

// GetRequestEncoding gets the value of RequestEncoding for the instance
func (instance *GlobalizationSection) GetPropertyRequestEncoding() (value string, err error) {
	retValue, err := instance.GetProperty("RequestEncoding")
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

// SetResourceProviderFactoryType sets the value of ResourceProviderFactoryType for the instance
func (instance *GlobalizationSection) SetPropertyResourceProviderFactoryType(value string) (err error) {
	return instance.SetProperty("ResourceProviderFactoryType", (value))
}

// GetResourceProviderFactoryType gets the value of ResourceProviderFactoryType for the instance
func (instance *GlobalizationSection) GetPropertyResourceProviderFactoryType() (value string, err error) {
	retValue, err := instance.GetProperty("ResourceProviderFactoryType")
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

// SetResponseEncoding sets the value of ResponseEncoding for the instance
func (instance *GlobalizationSection) SetPropertyResponseEncoding(value string) (err error) {
	return instance.SetProperty("ResponseEncoding", (value))
}

// GetResponseEncoding gets the value of ResponseEncoding for the instance
func (instance *GlobalizationSection) GetPropertyResponseEncoding() (value string, err error) {
	retValue, err := instance.GetProperty("ResponseEncoding")
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

// SetResponseHeaderEncoding sets the value of ResponseHeaderEncoding for the instance
func (instance *GlobalizationSection) SetPropertyResponseHeaderEncoding(value string) (err error) {
	return instance.SetProperty("ResponseHeaderEncoding", (value))
}

// GetResponseHeaderEncoding gets the value of ResponseHeaderEncoding for the instance
func (instance *GlobalizationSection) GetPropertyResponseHeaderEncoding() (value string, err error) {
	retValue, err := instance.GetProperty("ResponseHeaderEncoding")
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

// SetUiCulture sets the value of UiCulture for the instance
func (instance *GlobalizationSection) SetPropertyUiCulture(value string) (err error) {
	return instance.SetProperty("UiCulture", (value))
}

// GetUiCulture gets the value of UiCulture for the instance
func (instance *GlobalizationSection) GetPropertyUiCulture() (value string, err error) {
	retValue, err := instance.GetProperty("UiCulture")
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
