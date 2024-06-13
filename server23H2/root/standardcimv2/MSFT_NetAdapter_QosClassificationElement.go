// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapter_QosClassificationElement struct
type MSFT_NetAdapter_QosClassificationElement struct {
	*cim.WmiInstance

	//
	Priority uint8

	//
	ProtocolSelector uint16

	//
	ProtocolSpecificValue uint16
}

func NewMSFT_NetAdapter_QosClassificationElementEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapter_QosClassificationElement, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_QosClassificationElement{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapter_QosClassificationElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapter_QosClassificationElement, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_QosClassificationElement{
		WmiInstance: tmp,
	}
	return
}

// SetPriority sets the value of Priority for the instance
func (instance *MSFT_NetAdapter_QosClassificationElement) SetPropertyPriority(value uint8) (err error) {
	return instance.SetProperty("Priority", (value))
}

// GetPriority gets the value of Priority for the instance
func (instance *MSFT_NetAdapter_QosClassificationElement) GetPropertyPriority() (value uint8, err error) {
	retValue, err := instance.GetProperty("Priority")
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

// SetProtocolSelector sets the value of ProtocolSelector for the instance
func (instance *MSFT_NetAdapter_QosClassificationElement) SetPropertyProtocolSelector(value uint16) (err error) {
	return instance.SetProperty("ProtocolSelector", (value))
}

// GetProtocolSelector gets the value of ProtocolSelector for the instance
func (instance *MSFT_NetAdapter_QosClassificationElement) GetPropertyProtocolSelector() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProtocolSelector")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetProtocolSpecificValue sets the value of ProtocolSpecificValue for the instance
func (instance *MSFT_NetAdapter_QosClassificationElement) SetPropertyProtocolSpecificValue(value uint16) (err error) {
	return instance.SetProperty("ProtocolSpecificValue", (value))
}

// GetProtocolSpecificValue gets the value of ProtocolSpecificValue for the instance
func (instance *MSFT_NetAdapter_QosClassificationElement) GetPropertyProtocolSpecificValue() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProtocolSpecificValue")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
