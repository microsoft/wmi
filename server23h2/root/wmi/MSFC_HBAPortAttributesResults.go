// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFC_HBAPortAttributesResults struct
type MSFC_HBAPortAttributesResults struct {
	*cim.WmiInstance

	//
	FabricName []uint8

	//
	NodeWWN []uint8

	//
	NumberofDiscoveredPorts uint32

	//
	PortActiveFc4Types []uint8

	//
	PortFcId uint32

	//
	PortMaxFrameSize uint32

	//
	PortSpeed uint32

	//
	PortState uint32

	//
	PortSupportedClassofService uint32

	//
	PortSupportedFc4Types []uint8

	//
	PortSupportedSpeed uint32

	//
	PortType uint32

	//
	PortWWN []uint8
}

func NewMSFC_HBAPortAttributesResultsEx1(instance *cim.WmiInstance) (newInstance *MSFC_HBAPortAttributesResults, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_HBAPortAttributesResults{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_HBAPortAttributesResultsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_HBAPortAttributesResults, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_HBAPortAttributesResults{
		WmiInstance: tmp,
	}
	return
}

// SetFabricName sets the value of FabricName for the instance
func (instance *MSFC_HBAPortAttributesResults) SetPropertyFabricName(value []uint8) (err error) {
	return instance.SetProperty("FabricName", (value))
}

// GetFabricName gets the value of FabricName for the instance
func (instance *MSFC_HBAPortAttributesResults) GetPropertyFabricName() (value []uint8, err error) {
	retValue, err := instance.GetProperty("FabricName")
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

// SetNodeWWN sets the value of NodeWWN for the instance
func (instance *MSFC_HBAPortAttributesResults) SetPropertyNodeWWN(value []uint8) (err error) {
	return instance.SetProperty("NodeWWN", (value))
}

// GetNodeWWN gets the value of NodeWWN for the instance
func (instance *MSFC_HBAPortAttributesResults) GetPropertyNodeWWN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("NodeWWN")
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

// SetNumberofDiscoveredPorts sets the value of NumberofDiscoveredPorts for the instance
func (instance *MSFC_HBAPortAttributesResults) SetPropertyNumberofDiscoveredPorts(value uint32) (err error) {
	return instance.SetProperty("NumberofDiscoveredPorts", (value))
}

// GetNumberofDiscoveredPorts gets the value of NumberofDiscoveredPorts for the instance
func (instance *MSFC_HBAPortAttributesResults) GetPropertyNumberofDiscoveredPorts() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofDiscoveredPorts")
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

// SetPortActiveFc4Types sets the value of PortActiveFc4Types for the instance
func (instance *MSFC_HBAPortAttributesResults) SetPropertyPortActiveFc4Types(value []uint8) (err error) {
	return instance.SetProperty("PortActiveFc4Types", (value))
}

// GetPortActiveFc4Types gets the value of PortActiveFc4Types for the instance
func (instance *MSFC_HBAPortAttributesResults) GetPropertyPortActiveFc4Types() (value []uint8, err error) {
	retValue, err := instance.GetProperty("PortActiveFc4Types")
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

// SetPortFcId sets the value of PortFcId for the instance
func (instance *MSFC_HBAPortAttributesResults) SetPropertyPortFcId(value uint32) (err error) {
	return instance.SetProperty("PortFcId", (value))
}

// GetPortFcId gets the value of PortFcId for the instance
func (instance *MSFC_HBAPortAttributesResults) GetPropertyPortFcId() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortFcId")
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

// SetPortMaxFrameSize sets the value of PortMaxFrameSize for the instance
func (instance *MSFC_HBAPortAttributesResults) SetPropertyPortMaxFrameSize(value uint32) (err error) {
	return instance.SetProperty("PortMaxFrameSize", (value))
}

// GetPortMaxFrameSize gets the value of PortMaxFrameSize for the instance
func (instance *MSFC_HBAPortAttributesResults) GetPropertyPortMaxFrameSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortMaxFrameSize")
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

// SetPortSpeed sets the value of PortSpeed for the instance
func (instance *MSFC_HBAPortAttributesResults) SetPropertyPortSpeed(value uint32) (err error) {
	return instance.SetProperty("PortSpeed", (value))
}

// GetPortSpeed gets the value of PortSpeed for the instance
func (instance *MSFC_HBAPortAttributesResults) GetPropertyPortSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortSpeed")
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

// SetPortState sets the value of PortState for the instance
func (instance *MSFC_HBAPortAttributesResults) SetPropertyPortState(value uint32) (err error) {
	return instance.SetProperty("PortState", (value))
}

// GetPortState gets the value of PortState for the instance
func (instance *MSFC_HBAPortAttributesResults) GetPropertyPortState() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortState")
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

// SetPortSupportedClassofService sets the value of PortSupportedClassofService for the instance
func (instance *MSFC_HBAPortAttributesResults) SetPropertyPortSupportedClassofService(value uint32) (err error) {
	return instance.SetProperty("PortSupportedClassofService", (value))
}

// GetPortSupportedClassofService gets the value of PortSupportedClassofService for the instance
func (instance *MSFC_HBAPortAttributesResults) GetPropertyPortSupportedClassofService() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortSupportedClassofService")
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

// SetPortSupportedFc4Types sets the value of PortSupportedFc4Types for the instance
func (instance *MSFC_HBAPortAttributesResults) SetPropertyPortSupportedFc4Types(value []uint8) (err error) {
	return instance.SetProperty("PortSupportedFc4Types", (value))
}

// GetPortSupportedFc4Types gets the value of PortSupportedFc4Types for the instance
func (instance *MSFC_HBAPortAttributesResults) GetPropertyPortSupportedFc4Types() (value []uint8, err error) {
	retValue, err := instance.GetProperty("PortSupportedFc4Types")
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

// SetPortSupportedSpeed sets the value of PortSupportedSpeed for the instance
func (instance *MSFC_HBAPortAttributesResults) SetPropertyPortSupportedSpeed(value uint32) (err error) {
	return instance.SetProperty("PortSupportedSpeed", (value))
}

// GetPortSupportedSpeed gets the value of PortSupportedSpeed for the instance
func (instance *MSFC_HBAPortAttributesResults) GetPropertyPortSupportedSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortSupportedSpeed")
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

// SetPortType sets the value of PortType for the instance
func (instance *MSFC_HBAPortAttributesResults) SetPropertyPortType(value uint32) (err error) {
	return instance.SetProperty("PortType", (value))
}

// GetPortType gets the value of PortType for the instance
func (instance *MSFC_HBAPortAttributesResults) GetPropertyPortType() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortType")
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

// SetPortWWN sets the value of PortWWN for the instance
func (instance *MSFC_HBAPortAttributesResults) SetPropertyPortWWN(value []uint8) (err error) {
	return instance.SetProperty("PortWWN", (value))
}

// GetPortWWN gets the value of PortWWN for the instance
func (instance *MSFC_HBAPortAttributesResults) GetPropertyPortWWN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("PortWWN")
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
