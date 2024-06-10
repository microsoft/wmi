// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// HNet_ConnectionProperties struct
type HNet_ConnectionProperties struct {
	*cim.WmiInstance

	//
	Connection HNet_Connection

	//
	IsBridge bool

	//
	IsBridgeMember bool

	//
	IsFirewalled bool

	//
	IsIcsPrivate bool

	//
	IsIcsPublic bool
}

func NewHNet_ConnectionPropertiesEx1(instance *cim.WmiInstance) (newInstance *HNet_ConnectionProperties, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HNet_ConnectionProperties{
		WmiInstance: tmp,
	}
	return
}

func NewHNet_ConnectionPropertiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HNet_ConnectionProperties, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HNet_ConnectionProperties{
		WmiInstance: tmp,
	}
	return
}

// SetConnection sets the value of Connection for the instance
func (instance *HNet_ConnectionProperties) SetPropertyConnection(value HNet_Connection) (err error) {
	return instance.SetProperty("Connection", (value))
}

// GetConnection gets the value of Connection for the instance
func (instance *HNet_ConnectionProperties) GetPropertyConnection() (value HNet_Connection, err error) {
	retValue, err := instance.GetProperty("Connection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(HNet_Connection)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " HNet_Connection is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = HNet_Connection(valuetmp)

	return
}

// SetIsBridge sets the value of IsBridge for the instance
func (instance *HNet_ConnectionProperties) SetPropertyIsBridge(value bool) (err error) {
	return instance.SetProperty("IsBridge", (value))
}

// GetIsBridge gets the value of IsBridge for the instance
func (instance *HNet_ConnectionProperties) GetPropertyIsBridge() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBridge")
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

// SetIsBridgeMember sets the value of IsBridgeMember for the instance
func (instance *HNet_ConnectionProperties) SetPropertyIsBridgeMember(value bool) (err error) {
	return instance.SetProperty("IsBridgeMember", (value))
}

// GetIsBridgeMember gets the value of IsBridgeMember for the instance
func (instance *HNet_ConnectionProperties) GetPropertyIsBridgeMember() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBridgeMember")
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

// SetIsFirewalled sets the value of IsFirewalled for the instance
func (instance *HNet_ConnectionProperties) SetPropertyIsFirewalled(value bool) (err error) {
	return instance.SetProperty("IsFirewalled", (value))
}

// GetIsFirewalled gets the value of IsFirewalled for the instance
func (instance *HNet_ConnectionProperties) GetPropertyIsFirewalled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsFirewalled")
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

// SetIsIcsPrivate sets the value of IsIcsPrivate for the instance
func (instance *HNet_ConnectionProperties) SetPropertyIsIcsPrivate(value bool) (err error) {
	return instance.SetProperty("IsIcsPrivate", (value))
}

// GetIsIcsPrivate gets the value of IsIcsPrivate for the instance
func (instance *HNet_ConnectionProperties) GetPropertyIsIcsPrivate() (value bool, err error) {
	retValue, err := instance.GetProperty("IsIcsPrivate")
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

// SetIsIcsPublic sets the value of IsIcsPublic for the instance
func (instance *HNet_ConnectionProperties) SetPropertyIsIcsPublic(value bool) (err error) {
	return instance.SetProperty("IsIcsPublic", (value))
}

// GetIsIcsPublic gets the value of IsIcsPublic for the instance
func (instance *HNet_ConnectionProperties) GetPropertyIsIcsPublic() (value bool, err error) {
	retValue, err := instance.GetProperty("IsIcsPublic")
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
