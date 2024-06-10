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

// HttpCachePolicySettings struct
type HttpCachePolicySettings struct {
	*EmbeddedObject

	//
	MaximumAge string

	//
	MaximumStale string

	//
	MinimumFresh string

	//
	PolicyLevel int32
}

func NewHttpCachePolicySettingsEx1(instance *cim.WmiInstance) (newInstance *HttpCachePolicySettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HttpCachePolicySettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewHttpCachePolicySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HttpCachePolicySettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HttpCachePolicySettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetMaximumAge sets the value of MaximumAge for the instance
func (instance *HttpCachePolicySettings) SetPropertyMaximumAge(value string) (err error) {
	return instance.SetProperty("MaximumAge", (value))
}

// GetMaximumAge gets the value of MaximumAge for the instance
func (instance *HttpCachePolicySettings) GetPropertyMaximumAge() (value string, err error) {
	retValue, err := instance.GetProperty("MaximumAge")
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

// SetMaximumStale sets the value of MaximumStale for the instance
func (instance *HttpCachePolicySettings) SetPropertyMaximumStale(value string) (err error) {
	return instance.SetProperty("MaximumStale", (value))
}

// GetMaximumStale gets the value of MaximumStale for the instance
func (instance *HttpCachePolicySettings) GetPropertyMaximumStale() (value string, err error) {
	retValue, err := instance.GetProperty("MaximumStale")
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

// SetMinimumFresh sets the value of MinimumFresh for the instance
func (instance *HttpCachePolicySettings) SetPropertyMinimumFresh(value string) (err error) {
	return instance.SetProperty("MinimumFresh", (value))
}

// GetMinimumFresh gets the value of MinimumFresh for the instance
func (instance *HttpCachePolicySettings) GetPropertyMinimumFresh() (value string, err error) {
	retValue, err := instance.GetProperty("MinimumFresh")
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

// SetPolicyLevel sets the value of PolicyLevel for the instance
func (instance *HttpCachePolicySettings) SetPropertyPolicyLevel(value int32) (err error) {
	return instance.SetProperty("PolicyLevel", (value))
}

// GetPolicyLevel gets the value of PolicyLevel for the instance
func (instance *HttpCachePolicySettings) GetPropertyPolicyLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("PolicyLevel")
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
