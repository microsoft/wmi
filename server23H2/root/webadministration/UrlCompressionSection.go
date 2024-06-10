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

// UrlCompressionSection struct
type UrlCompressionSection struct {
	*ConfigurationSection

	//
	DoDynamicCompression bool

	//
	DoStaticCompression bool

	//
	DynamicCompressionBeforeCache bool
}

func NewUrlCompressionSectionEx1(instance *cim.WmiInstance) (newInstance *UrlCompressionSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &UrlCompressionSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewUrlCompressionSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *UrlCompressionSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &UrlCompressionSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetDoDynamicCompression sets the value of DoDynamicCompression for the instance
func (instance *UrlCompressionSection) SetPropertyDoDynamicCompression(value bool) (err error) {
	return instance.SetProperty("DoDynamicCompression", (value))
}

// GetDoDynamicCompression gets the value of DoDynamicCompression for the instance
func (instance *UrlCompressionSection) GetPropertyDoDynamicCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("DoDynamicCompression")
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

// SetDoStaticCompression sets the value of DoStaticCompression for the instance
func (instance *UrlCompressionSection) SetPropertyDoStaticCompression(value bool) (err error) {
	return instance.SetProperty("DoStaticCompression", (value))
}

// GetDoStaticCompression gets the value of DoStaticCompression for the instance
func (instance *UrlCompressionSection) GetPropertyDoStaticCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("DoStaticCompression")
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

// SetDynamicCompressionBeforeCache sets the value of DynamicCompressionBeforeCache for the instance
func (instance *UrlCompressionSection) SetPropertyDynamicCompressionBeforeCache(value bool) (err error) {
	return instance.SetProperty("DynamicCompressionBeforeCache", (value))
}

// GetDynamicCompressionBeforeCache gets the value of DynamicCompressionBeforeCache for the instance
func (instance *UrlCompressionSection) GetPropertyDynamicCompressionBeforeCache() (value bool, err error) {
	retValue, err := instance.GetProperty("DynamicCompressionBeforeCache")
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
