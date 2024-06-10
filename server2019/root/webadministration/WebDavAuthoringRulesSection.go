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

// WebDavAuthoringRulesSection struct
type WebDavAuthoringRulesSection struct {
	*ConfigurationSectionWithCollection

	//
	AllowNonMimeMapFiles bool

	//
	AuthoringRules []WebDavMimeTypeElement

	//
	DefaultAccess int32

	//
	DefaultMimeType string
}

func NewWebDavAuthoringRulesSectionEx1(instance *cim.WmiInstance) (newInstance *WebDavAuthoringRulesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebDavAuthoringRulesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewWebDavAuthoringRulesSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebDavAuthoringRulesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebDavAuthoringRulesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetAllowNonMimeMapFiles sets the value of AllowNonMimeMapFiles for the instance
func (instance *WebDavAuthoringRulesSection) SetPropertyAllowNonMimeMapFiles(value bool) (err error) {
	return instance.SetProperty("AllowNonMimeMapFiles", (value))
}

// GetAllowNonMimeMapFiles gets the value of AllowNonMimeMapFiles for the instance
func (instance *WebDavAuthoringRulesSection) GetPropertyAllowNonMimeMapFiles() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowNonMimeMapFiles")
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

// SetAuthoringRules sets the value of AuthoringRules for the instance
func (instance *WebDavAuthoringRulesSection) SetPropertyAuthoringRules(value []WebDavMimeTypeElement) (err error) {
	return instance.SetProperty("AuthoringRules", (value))
}

// GetAuthoringRules gets the value of AuthoringRules for the instance
func (instance *WebDavAuthoringRulesSection) GetPropertyAuthoringRules() (value []WebDavMimeTypeElement, err error) {
	retValue, err := instance.GetProperty("AuthoringRules")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(WebDavMimeTypeElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " WebDavMimeTypeElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, WebDavMimeTypeElement(valuetmp))
	}

	return
}

// SetDefaultAccess sets the value of DefaultAccess for the instance
func (instance *WebDavAuthoringRulesSection) SetPropertyDefaultAccess(value int32) (err error) {
	return instance.SetProperty("DefaultAccess", (value))
}

// GetDefaultAccess gets the value of DefaultAccess for the instance
func (instance *WebDavAuthoringRulesSection) GetPropertyDefaultAccess() (value int32, err error) {
	retValue, err := instance.GetProperty("DefaultAccess")
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

// SetDefaultMimeType sets the value of DefaultMimeType for the instance
func (instance *WebDavAuthoringRulesSection) SetPropertyDefaultMimeType(value string) (err error) {
	return instance.SetProperty("DefaultMimeType", (value))
}

// GetDefaultMimeType gets the value of DefaultMimeType for the instance
func (instance *WebDavAuthoringRulesSection) GetPropertyDefaultMimeType() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultMimeType")
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
