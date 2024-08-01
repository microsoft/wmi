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

// WmiMonitorRawEEdidV1Block struct
type WmiMonitorRawEEdidV1Block struct {
	*MSMonitorClass

	//
	Active bool

	//
	Content []uint8

	//
	Id uint8

	//
	InstanceName string

	//
	Type uint8
}

func NewWmiMonitorRawEEdidV1BlockEx1(instance *cim.WmiInstance) (newInstance *WmiMonitorRawEEdidV1Block, err error) {
	tmp, err := NewMSMonitorClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorRawEEdidV1Block{
		MSMonitorClass: tmp,
	}
	return
}

func NewWmiMonitorRawEEdidV1BlockEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WmiMonitorRawEEdidV1Block, err error) {
	tmp, err := NewMSMonitorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WmiMonitorRawEEdidV1Block{
		MSMonitorClass: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *WmiMonitorRawEEdidV1Block) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WmiMonitorRawEEdidV1Block) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetContent sets the value of Content for the instance
func (instance *WmiMonitorRawEEdidV1Block) SetPropertyContent(value []uint8) (err error) {
	return instance.SetProperty("Content", (value))
}

// GetContent gets the value of Content for the instance
func (instance *WmiMonitorRawEEdidV1Block) GetPropertyContent() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Content")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetId sets the value of Id for the instance
func (instance *WmiMonitorRawEEdidV1Block) SetPropertyId(value uint8) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *WmiMonitorRawEEdidV1Block) GetPropertyId() (value uint8, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *WmiMonitorRawEEdidV1Block) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WmiMonitorRawEEdidV1Block) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetType sets the value of Type for the instance
func (instance *WmiMonitorRawEEdidV1Block) SetPropertyType(value uint8) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *WmiMonitorRawEEdidV1Block) GetPropertyType() (value uint8, err error) {
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
