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

// MSNdis_StatusMediaSpecificIndicationEx struct
type MSNdis_StatusMediaSpecificIndicationEx struct {
	*WMIEvent

	//
	Active bool

	//
	InstanceName string

	//
	NdisStatusMediaSpecificIndication []uint8

	//
	NumberElements uint32
}

func NewMSNdis_StatusMediaSpecificIndicationExEx1(instance *cim.WmiInstance) (newInstance *MSNdis_StatusMediaSpecificIndicationEx, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_StatusMediaSpecificIndicationEx{
		WMIEvent: tmp,
	}
	return
}

func NewMSNdis_StatusMediaSpecificIndicationExEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_StatusMediaSpecificIndicationEx, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_StatusMediaSpecificIndicationEx{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_StatusMediaSpecificIndicationEx) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_StatusMediaSpecificIndicationEx) GetPropertyActive() (value bool, err error) {
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
func (instance *MSNdis_StatusMediaSpecificIndicationEx) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_StatusMediaSpecificIndicationEx) GetPropertyInstanceName() (value string, err error) {
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

// SetNdisStatusMediaSpecificIndication sets the value of NdisStatusMediaSpecificIndication for the instance
func (instance *MSNdis_StatusMediaSpecificIndicationEx) SetPropertyNdisStatusMediaSpecificIndication(value []uint8) (err error) {
	return instance.SetProperty("NdisStatusMediaSpecificIndication", (value))
}

// GetNdisStatusMediaSpecificIndication gets the value of NdisStatusMediaSpecificIndication for the instance
func (instance *MSNdis_StatusMediaSpecificIndicationEx) GetPropertyNdisStatusMediaSpecificIndication() (value []uint8, err error) {
	retValue, err := instance.GetProperty("NdisStatusMediaSpecificIndication")
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
func (instance *MSNdis_StatusMediaSpecificIndicationEx) SetPropertyNumberElements(value uint32) (err error) {
	return instance.SetProperty("NumberElements", (value))
}

// GetNumberElements gets the value of NumberElements for the instance
func (instance *MSNdis_StatusMediaSpecificIndicationEx) GetPropertyNumberElements() (value uint32, err error) {
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
