// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MTLogicalProcessor struct
type MSFT_MTLogicalProcessor struct {
	*CIM_ManagedElement

	//
	CpuId uint16

	//
	CurrentIndex uint16

	//
	IntervalSeconds uint16

	//
	NodeId uint16

	//
	Parking bool

	//
	Privileged []float32

	//
	Utilization []float32
}

func NewMSFT_MTLogicalProcessorEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTLogicalProcessor, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTLogicalProcessor{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_MTLogicalProcessorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MTLogicalProcessor, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTLogicalProcessor{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetCpuId sets the value of CpuId for the instance
func (instance *MSFT_MTLogicalProcessor) SetPropertyCpuId(value uint16) (err error) {
	return instance.SetProperty("CpuId", (value))
}

// GetCpuId gets the value of CpuId for the instance
func (instance *MSFT_MTLogicalProcessor) GetPropertyCpuId() (value uint16, err error) {
	retValue, err := instance.GetProperty("CpuId")
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

// SetCurrentIndex sets the value of CurrentIndex for the instance
func (instance *MSFT_MTLogicalProcessor) SetPropertyCurrentIndex(value uint16) (err error) {
	return instance.SetProperty("CurrentIndex", (value))
}

// GetCurrentIndex gets the value of CurrentIndex for the instance
func (instance *MSFT_MTLogicalProcessor) GetPropertyCurrentIndex() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentIndex")
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

// SetIntervalSeconds sets the value of IntervalSeconds for the instance
func (instance *MSFT_MTLogicalProcessor) SetPropertyIntervalSeconds(value uint16) (err error) {
	return instance.SetProperty("IntervalSeconds", (value))
}

// GetIntervalSeconds gets the value of IntervalSeconds for the instance
func (instance *MSFT_MTLogicalProcessor) GetPropertyIntervalSeconds() (value uint16, err error) {
	retValue, err := instance.GetProperty("IntervalSeconds")
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

// SetNodeId sets the value of NodeId for the instance
func (instance *MSFT_MTLogicalProcessor) SetPropertyNodeId(value uint16) (err error) {
	return instance.SetProperty("NodeId", (value))
}

// GetNodeId gets the value of NodeId for the instance
func (instance *MSFT_MTLogicalProcessor) GetPropertyNodeId() (value uint16, err error) {
	retValue, err := instance.GetProperty("NodeId")
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

// SetParking sets the value of Parking for the instance
func (instance *MSFT_MTLogicalProcessor) SetPropertyParking(value bool) (err error) {
	return instance.SetProperty("Parking", (value))
}

// GetParking gets the value of Parking for the instance
func (instance *MSFT_MTLogicalProcessor) GetPropertyParking() (value bool, err error) {
	retValue, err := instance.GetProperty("Parking")
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

// SetPrivileged sets the value of Privileged for the instance
func (instance *MSFT_MTLogicalProcessor) SetPropertyPrivileged(value []float32) (err error) {
	return instance.SetProperty("Privileged", (value))
}

// GetPrivileged gets the value of Privileged for the instance
func (instance *MSFT_MTLogicalProcessor) GetPropertyPrivileged() (value []float32, err error) {
	retValue, err := instance.GetProperty("Privileged")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(float32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " float32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, float32(valuetmp))
	}

	return
}

// SetUtilization sets the value of Utilization for the instance
func (instance *MSFT_MTLogicalProcessor) SetPropertyUtilization(value []float32) (err error) {
	return instance.SetProperty("Utilization", (value))
}

// GetUtilization gets the value of Utilization for the instance
func (instance *MSFT_MTLogicalProcessor) GetPropertyUtilization() (value []float32, err error) {
	retValue, err := instance.GetProperty("Utilization")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(float32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " float32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, float32(valuetmp))
	}

	return
}
