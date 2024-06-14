// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Registry_ChangeNotification struct
type Registry_ChangeNotification struct {
	*Registry

	//
	KeyHandle uint32

	//
	Notification uint32

	//
	Primary uint8

	//
	Type uint8

	//
	WatchSubtree uint8
}

func NewRegistry_ChangeNotificationEx1(instance *cim.WmiInstance) (newInstance *Registry_ChangeNotification, err error) {
	tmp, err := NewRegistryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Registry_ChangeNotification{
		Registry: tmp,
	}
	return
}

func NewRegistry_ChangeNotificationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Registry_ChangeNotification, err error) {
	tmp, err := NewRegistryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Registry_ChangeNotification{
		Registry: tmp,
	}
	return
}

// SetKeyHandle sets the value of KeyHandle for the instance
func (instance *Registry_ChangeNotification) SetPropertyKeyHandle(value uint32) (err error) {
	return instance.SetProperty("KeyHandle", (value))
}

// GetKeyHandle gets the value of KeyHandle for the instance
func (instance *Registry_ChangeNotification) GetPropertyKeyHandle() (value uint32, err error) {
	retValue, err := instance.GetProperty("KeyHandle")
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

// SetNotification sets the value of Notification for the instance
func (instance *Registry_ChangeNotification) SetPropertyNotification(value uint32) (err error) {
	return instance.SetProperty("Notification", (value))
}

// GetNotification gets the value of Notification for the instance
func (instance *Registry_ChangeNotification) GetPropertyNotification() (value uint32, err error) {
	retValue, err := instance.GetProperty("Notification")
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

// SetPrimary sets the value of Primary for the instance
func (instance *Registry_ChangeNotification) SetPropertyPrimary(value uint8) (err error) {
	return instance.SetProperty("Primary", (value))
}

// GetPrimary gets the value of Primary for the instance
func (instance *Registry_ChangeNotification) GetPropertyPrimary() (value uint8, err error) {
	retValue, err := instance.GetProperty("Primary")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetType sets the value of Type for the instance
func (instance *Registry_ChangeNotification) SetPropertyType(value uint8) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *Registry_ChangeNotification) GetPropertyType() (value uint8, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetWatchSubtree sets the value of WatchSubtree for the instance
func (instance *Registry_ChangeNotification) SetPropertyWatchSubtree(value uint8) (err error) {
	return instance.SetProperty("WatchSubtree", (value))
}

// GetWatchSubtree gets the value of WatchSubtree for the instance
func (instance *Registry_ChangeNotification) GetPropertyWatchSubtree() (value uint8, err error) {
	retValue, err := instance.GetProperty("WatchSubtree")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
