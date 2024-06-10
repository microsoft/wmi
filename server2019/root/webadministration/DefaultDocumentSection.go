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

// DefaultDocumentSection struct
type DefaultDocumentSection struct {
	*ConfigurationSectionWithCollection

	//
	Enabled bool

	//
	Files FileSettings
}

func NewDefaultDocumentSectionEx1(instance *cim.WmiInstance) (newInstance *DefaultDocumentSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DefaultDocumentSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewDefaultDocumentSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DefaultDocumentSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DefaultDocumentSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *DefaultDocumentSection) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *DefaultDocumentSection) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetFiles sets the value of Files for the instance
func (instance *DefaultDocumentSection) SetPropertyFiles(value FileSettings) (err error) {
	return instance.SetProperty("Files", (value))
}

// GetFiles gets the value of Files for the instance
func (instance *DefaultDocumentSection) GetPropertyFiles() (value FileSettings, err error) {
	retValue, err := instance.GetProperty("Files")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FileSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FileSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FileSettings(valuetmp)

	return
}
