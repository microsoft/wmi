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

// DenyByConcurrentRequestsSettings struct
type DenyByConcurrentRequestsSettings struct {
	*EmbeddedObject

	//
	Enabled bool

	//
	MaxConcurrentRequests uint32
}

func NewDenyByConcurrentRequestsSettingsEx1(instance *cim.WmiInstance) (newInstance *DenyByConcurrentRequestsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DenyByConcurrentRequestsSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewDenyByConcurrentRequestsSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DenyByConcurrentRequestsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DenyByConcurrentRequestsSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *DenyByConcurrentRequestsSettings) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *DenyByConcurrentRequestsSettings) GetPropertyEnabled() (value bool, err error) {
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

// SetMaxConcurrentRequests sets the value of MaxConcurrentRequests for the instance
func (instance *DenyByConcurrentRequestsSettings) SetPropertyMaxConcurrentRequests(value uint32) (err error) {
	return instance.SetProperty("MaxConcurrentRequests", (value))
}

// GetMaxConcurrentRequests gets the value of MaxConcurrentRequests for the instance
func (instance *DenyByConcurrentRequestsSettings) GetPropertyMaxConcurrentRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxConcurrentRequests")
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
