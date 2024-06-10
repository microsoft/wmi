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

// StaticContentSection struct
type StaticContentSection struct {
	*ConfigurationSectionWithCollection

	//
	ClientCache HttpClientCache

	//
	DefaultDocFooter string

	//
	EnableDocFooter bool

	//
	IsDocFooterFileName bool

	//
	StaticContent []MimeMapElement
}

func NewStaticContentSectionEx1(instance *cim.WmiInstance) (newInstance *StaticContentSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &StaticContentSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewStaticContentSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *StaticContentSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &StaticContentSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetClientCache sets the value of ClientCache for the instance
func (instance *StaticContentSection) SetPropertyClientCache(value HttpClientCache) (err error) {
	return instance.SetProperty("ClientCache", (value))
}

// GetClientCache gets the value of ClientCache for the instance
func (instance *StaticContentSection) GetPropertyClientCache() (value HttpClientCache, err error) {
	retValue, err := instance.GetProperty("ClientCache")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(HttpClientCache)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " HttpClientCache is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = HttpClientCache(valuetmp)

	return
}

// SetDefaultDocFooter sets the value of DefaultDocFooter for the instance
func (instance *StaticContentSection) SetPropertyDefaultDocFooter(value string) (err error) {
	return instance.SetProperty("DefaultDocFooter", (value))
}

// GetDefaultDocFooter gets the value of DefaultDocFooter for the instance
func (instance *StaticContentSection) GetPropertyDefaultDocFooter() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultDocFooter")
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

// SetEnableDocFooter sets the value of EnableDocFooter for the instance
func (instance *StaticContentSection) SetPropertyEnableDocFooter(value bool) (err error) {
	return instance.SetProperty("EnableDocFooter", (value))
}

// GetEnableDocFooter gets the value of EnableDocFooter for the instance
func (instance *StaticContentSection) GetPropertyEnableDocFooter() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableDocFooter")
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

// SetIsDocFooterFileName sets the value of IsDocFooterFileName for the instance
func (instance *StaticContentSection) SetPropertyIsDocFooterFileName(value bool) (err error) {
	return instance.SetProperty("IsDocFooterFileName", (value))
}

// GetIsDocFooterFileName gets the value of IsDocFooterFileName for the instance
func (instance *StaticContentSection) GetPropertyIsDocFooterFileName() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDocFooterFileName")
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

// SetStaticContent sets the value of StaticContent for the instance
func (instance *StaticContentSection) SetPropertyStaticContent(value []MimeMapElement) (err error) {
	return instance.SetProperty("StaticContent", (value))
}

// GetStaticContent gets the value of StaticContent for the instance
func (instance *StaticContentSection) GetPropertyStaticContent() (value []MimeMapElement, err error) {
	retValue, err := instance.GetProperty("StaticContent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MimeMapElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MimeMapElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MimeMapElement(valuetmp))
	}

	return
}
