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

// HttpTracingSection struct
type HttpTracingSection struct {
	*ConfigurationSectionWithCollection

	//
	TraceUrls TraceUrlSettings
}

func NewHttpTracingSectionEx1(instance *cim.WmiInstance) (newInstance *HttpTracingSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HttpTracingSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewHttpTracingSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HttpTracingSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HttpTracingSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetTraceUrls sets the value of TraceUrls for the instance
func (instance *HttpTracingSection) SetPropertyTraceUrls(value TraceUrlSettings) (err error) {
	return instance.SetProperty("TraceUrls", (value))
}

// GetTraceUrls gets the value of TraceUrls for the instance
func (instance *HttpTracingSection) GetPropertyTraceUrls() (value TraceUrlSettings, err error) {
	retValue, err := instance.GetProperty("TraceUrls")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(TraceUrlSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " TraceUrlSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = TraceUrlSettings(valuetmp)

	return
}
