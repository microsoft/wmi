// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_StatusDot11RoamingCompletion struct
type MSNdis_StatusDot11RoamingCompletion struct {
	*WMIEvent

	//
	Active bool

	//
	InstanceName string

	//
	NdisStatusDot11RoamingCompletionIndication []uint8

	//
	NumberElements uint32
}

func NewMSNdis_StatusDot11RoamingCompletionEx1(instance *cim.WmiInstance) (newInstance *MSNdis_StatusDot11RoamingCompletion, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_StatusDot11RoamingCompletion{
		WMIEvent: tmp,
	}
	return
}

func NewMSNdis_StatusDot11RoamingCompletionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_StatusDot11RoamingCompletion, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_StatusDot11RoamingCompletion{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_StatusDot11RoamingCompletion) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_StatusDot11RoamingCompletion) GetPropertyActive() (value bool, err error) {
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSNdis_StatusDot11RoamingCompletion) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_StatusDot11RoamingCompletion) GetPropertyInstanceName() (value string, err error) {
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

// SetNdisStatusDot11RoamingCompletionIndication sets the value of NdisStatusDot11RoamingCompletionIndication for the instance
func (instance *MSNdis_StatusDot11RoamingCompletion) SetPropertyNdisStatusDot11RoamingCompletionIndication(value []uint8) (err error) {
	return instance.SetProperty("NdisStatusDot11RoamingCompletionIndication", (value))
}

// GetNdisStatusDot11RoamingCompletionIndication gets the value of NdisStatusDot11RoamingCompletionIndication for the instance
func (instance *MSNdis_StatusDot11RoamingCompletion) GetPropertyNdisStatusDot11RoamingCompletionIndication() (value []uint8, err error) {
	retValue, err := instance.GetProperty("NdisStatusDot11RoamingCompletionIndication")
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

// SetNumberElements sets the value of NumberElements for the instance
func (instance *MSNdis_StatusDot11RoamingCompletion) SetPropertyNumberElements(value uint32) (err error) {
	return instance.SetProperty("NumberElements", (value))
}

// GetNumberElements gets the value of NumberElements for the instance
func (instance *MSNdis_StatusDot11RoamingCompletion) GetPropertyNumberElements() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberElements")
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
