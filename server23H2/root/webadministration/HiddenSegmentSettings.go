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

// HiddenSegmentSettings struct
type HiddenSegmentSettings struct {
	*EmbeddedObject

	//
	ApplyToWebDAV bool

	//
	HiddenSegments []SegmentElement
}

func NewHiddenSegmentSettingsEx1(instance *cim.WmiInstance) (newInstance *HiddenSegmentSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HiddenSegmentSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewHiddenSegmentSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HiddenSegmentSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HiddenSegmentSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetApplyToWebDAV sets the value of ApplyToWebDAV for the instance
func (instance *HiddenSegmentSettings) SetPropertyApplyToWebDAV(value bool) (err error) {
	return instance.SetProperty("ApplyToWebDAV", (value))
}

// GetApplyToWebDAV gets the value of ApplyToWebDAV for the instance
func (instance *HiddenSegmentSettings) GetPropertyApplyToWebDAV() (value bool, err error) {
	retValue, err := instance.GetProperty("ApplyToWebDAV")
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

// SetHiddenSegments sets the value of HiddenSegments for the instance
func (instance *HiddenSegmentSettings) SetPropertyHiddenSegments(value []SegmentElement) (err error) {
	return instance.SetProperty("HiddenSegments", (value))
}

// GetHiddenSegments gets the value of HiddenSegments for the instance
func (instance *HiddenSegmentSettings) GetPropertyHiddenSegments() (value []SegmentElement, err error) {
	retValue, err := instance.GetProperty("HiddenSegments")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(SegmentElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " SegmentElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SegmentElement(valuetmp))
	}

	return
}
