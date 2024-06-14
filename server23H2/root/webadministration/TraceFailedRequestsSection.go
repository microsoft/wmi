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

// TraceFailedRequestsSection struct
type TraceFailedRequestsSection struct {
	*ConfigurationSectionWithCollection

	//
	TraceFailedRequests []TraceUrl
}

func NewTraceFailedRequestsSectionEx1(instance *cim.WmiInstance) (newInstance *TraceFailedRequestsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TraceFailedRequestsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewTraceFailedRequestsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TraceFailedRequestsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TraceFailedRequestsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetTraceFailedRequests sets the value of TraceFailedRequests for the instance
func (instance *TraceFailedRequestsSection) SetPropertyTraceFailedRequests(value []TraceUrl) (err error) {
	return instance.SetProperty("TraceFailedRequests", (value))
}

// GetTraceFailedRequests gets the value of TraceFailedRequests for the instance
func (instance *TraceFailedRequestsSection) GetPropertyTraceFailedRequests() (value []TraceUrl, err error) {
	retValue, err := instance.GetProperty("TraceFailedRequests")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(TraceUrl)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " TraceUrl is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, TraceUrl(valuetmp))
	}

	return
}
