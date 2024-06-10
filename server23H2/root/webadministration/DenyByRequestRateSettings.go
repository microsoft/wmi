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

// DenyByRequestRateSettings struct
type DenyByRequestRateSettings struct {
	*EmbeddedObject

	//
	Enabled bool

	//
	MaxRequests uint32

	//
	RequestIntervalInMilliseconds uint32
}

func NewDenyByRequestRateSettingsEx1(instance *cim.WmiInstance) (newInstance *DenyByRequestRateSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DenyByRequestRateSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewDenyByRequestRateSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DenyByRequestRateSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DenyByRequestRateSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *DenyByRequestRateSettings) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *DenyByRequestRateSettings) GetPropertyEnabled() (value bool, err error) {
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

// SetMaxRequests sets the value of MaxRequests for the instance
func (instance *DenyByRequestRateSettings) SetPropertyMaxRequests(value uint32) (err error) {
	return instance.SetProperty("MaxRequests", (value))
}

// GetMaxRequests gets the value of MaxRequests for the instance
func (instance *DenyByRequestRateSettings) GetPropertyMaxRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxRequests")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetRequestIntervalInMilliseconds sets the value of RequestIntervalInMilliseconds for the instance
func (instance *DenyByRequestRateSettings) SetPropertyRequestIntervalInMilliseconds(value uint32) (err error) {
	return instance.SetProperty("RequestIntervalInMilliseconds", (value))
}

// GetRequestIntervalInMilliseconds gets the value of RequestIntervalInMilliseconds for the instance
func (instance *DenyByRequestRateSettings) GetPropertyRequestIntervalInMilliseconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("RequestIntervalInMilliseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
