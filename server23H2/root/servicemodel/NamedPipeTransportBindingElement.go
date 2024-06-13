// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// NamedPipeTransportBindingElement struct
type NamedPipeTransportBindingElement struct {
	*ConnectionOrientedTransportBindingElement

	// The connection pool settings.
	ConnectionPoolSettings NamedPipeConnectionPoolSettings
}

func NewNamedPipeTransportBindingElementEx1(instance *cim.WmiInstance) (newInstance *NamedPipeTransportBindingElement, err error) {
	tmp, err := NewConnectionOrientedTransportBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NamedPipeTransportBindingElement{
		ConnectionOrientedTransportBindingElement: tmp,
	}
	return
}

func NewNamedPipeTransportBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NamedPipeTransportBindingElement, err error) {
	tmp, err := NewConnectionOrientedTransportBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NamedPipeTransportBindingElement{
		ConnectionOrientedTransportBindingElement: tmp,
	}
	return
}

// SetConnectionPoolSettings sets the value of ConnectionPoolSettings for the instance
func (instance *NamedPipeTransportBindingElement) SetPropertyConnectionPoolSettings(value NamedPipeConnectionPoolSettings) (err error) {
	return instance.SetProperty("ConnectionPoolSettings", (value))
}

// GetConnectionPoolSettings gets the value of ConnectionPoolSettings for the instance
func (instance *NamedPipeTransportBindingElement) GetPropertyConnectionPoolSettings() (value NamedPipeConnectionPoolSettings, err error) {
	retValue, err := instance.GetProperty("ConnectionPoolSettings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(NamedPipeConnectionPoolSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " NamedPipeConnectionPoolSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = NamedPipeConnectionPoolSettings(valuetmp)

	return
}
