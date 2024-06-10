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

// SiteLimits struct
type SiteLimits struct {
	*EmbeddedObject

	//
	ConnectionTimeout string

	//
	MaxBandwidth uint32

	//
	MaxConnections uint32

	//
	MaxUrlSegments uint32
}

func NewSiteLimitsEx1(instance *cim.WmiInstance) (newInstance *SiteLimits, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SiteLimits{
		EmbeddedObject: tmp,
	}
	return
}

func NewSiteLimitsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SiteLimits, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SiteLimits{
		EmbeddedObject: tmp,
	}
	return
}

// SetConnectionTimeout sets the value of ConnectionTimeout for the instance
func (instance *SiteLimits) SetPropertyConnectionTimeout(value string) (err error) {
	return instance.SetProperty("ConnectionTimeout", (value))
}

// GetConnectionTimeout gets the value of ConnectionTimeout for the instance
func (instance *SiteLimits) GetPropertyConnectionTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionTimeout")
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

// SetMaxBandwidth sets the value of MaxBandwidth for the instance
func (instance *SiteLimits) SetPropertyMaxBandwidth(value uint32) (err error) {
	return instance.SetProperty("MaxBandwidth", (value))
}

// GetMaxBandwidth gets the value of MaxBandwidth for the instance
func (instance *SiteLimits) GetPropertyMaxBandwidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxBandwidth")
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

// SetMaxConnections sets the value of MaxConnections for the instance
func (instance *SiteLimits) SetPropertyMaxConnections(value uint32) (err error) {
	return instance.SetProperty("MaxConnections", (value))
}

// GetMaxConnections gets the value of MaxConnections for the instance
func (instance *SiteLimits) GetPropertyMaxConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxConnections")
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

// SetMaxUrlSegments sets the value of MaxUrlSegments for the instance
func (instance *SiteLimits) SetPropertyMaxUrlSegments(value uint32) (err error) {
	return instance.SetProperty("MaxUrlSegments", (value))
}

// GetMaxUrlSegments gets the value of MaxUrlSegments for the instance
func (instance *SiteLimits) GetPropertyMaxUrlSegments() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxUrlSegments")
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
