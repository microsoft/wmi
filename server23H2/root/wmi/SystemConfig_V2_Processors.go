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

// SystemConfig_V2_Processors struct
type SystemConfig_V2_Processors struct {
	*SystemConfig_V2

	//
	FeatureSet uint32

	//
	ProcessorIdentifier []byte

	//
	ProcessorIndex uint32

	//
	ProcessorName []byte

	//
	ProcessorSpeed uint32

	//
	VendorIdentifier []byte
}

func NewSystemConfig_V2_ProcessorsEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_Processors, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_Processors{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_ProcessorsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_Processors, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_Processors{
		SystemConfig_V2: tmp,
	}
	return
}

// SetFeatureSet sets the value of FeatureSet for the instance
func (instance *SystemConfig_V2_Processors) SetPropertyFeatureSet(value uint32) (err error) {
	return instance.SetProperty("FeatureSet", (value))
}

// GetFeatureSet gets the value of FeatureSet for the instance
func (instance *SystemConfig_V2_Processors) GetPropertyFeatureSet() (value uint32, err error) {
	retValue, err := instance.GetProperty("FeatureSet")
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

// SetProcessorIdentifier sets the value of ProcessorIdentifier for the instance
func (instance *SystemConfig_V2_Processors) SetPropertyProcessorIdentifier(value []byte) (err error) {
	return instance.SetProperty("ProcessorIdentifier", (value))
}

// GetProcessorIdentifier gets the value of ProcessorIdentifier for the instance
func (instance *SystemConfig_V2_Processors) GetPropertyProcessorIdentifier() (value []byte, err error) {
	retValue, err := instance.GetProperty("ProcessorIdentifier")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(byte)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, byte(valuetmp))
	}

	return
}

// SetProcessorIndex sets the value of ProcessorIndex for the instance
func (instance *SystemConfig_V2_Processors) SetPropertyProcessorIndex(value uint32) (err error) {
	return instance.SetProperty("ProcessorIndex", (value))
}

// GetProcessorIndex gets the value of ProcessorIndex for the instance
func (instance *SystemConfig_V2_Processors) GetPropertyProcessorIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessorIndex")
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

// SetProcessorName sets the value of ProcessorName for the instance
func (instance *SystemConfig_V2_Processors) SetPropertyProcessorName(value []byte) (err error) {
	return instance.SetProperty("ProcessorName", (value))
}

// GetProcessorName gets the value of ProcessorName for the instance
func (instance *SystemConfig_V2_Processors) GetPropertyProcessorName() (value []byte, err error) {
	retValue, err := instance.GetProperty("ProcessorName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(byte)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, byte(valuetmp))
	}

	return
}

// SetProcessorSpeed sets the value of ProcessorSpeed for the instance
func (instance *SystemConfig_V2_Processors) SetPropertyProcessorSpeed(value uint32) (err error) {
	return instance.SetProperty("ProcessorSpeed", (value))
}

// GetProcessorSpeed gets the value of ProcessorSpeed for the instance
func (instance *SystemConfig_V2_Processors) GetPropertyProcessorSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessorSpeed")
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

// SetVendorIdentifier sets the value of VendorIdentifier for the instance
func (instance *SystemConfig_V2_Processors) SetPropertyVendorIdentifier(value []byte) (err error) {
	return instance.SetProperty("VendorIdentifier", (value))
}

// GetVendorIdentifier gets the value of VendorIdentifier for the instance
func (instance *SystemConfig_V2_Processors) GetPropertyVendorIdentifier() (value []byte, err error) {
	retValue, err := instance.GetProperty("VendorIdentifier")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(byte)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, byte(valuetmp))
	}

	return
}
