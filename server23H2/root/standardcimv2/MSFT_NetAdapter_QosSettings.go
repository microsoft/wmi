// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapter_QosSettings struct
type MSFT_NetAdapter_QosSettings struct {
	*cim.WmiInstance

	//
	BandwidthAssignmentTable []uint8

	//
	ClassificationEnabled bool

	//
	ClassificationTable []MSFT_NetAdapter_QosClassificationElement

	//
	FlowControlEnabled bool

	//
	NumberOfClassificationElements uint32

	//
	PriorityAssignmentTable []uint8

	//
	PriorityFlowControlEnableArray []bool

	//
	TransmissionSelectionEnabled bool

	//
	TsaAssignmentTable []uint8
}

func NewMSFT_NetAdapter_QosSettingsEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapter_QosSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_QosSettings{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapter_QosSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapter_QosSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_QosSettings{
		WmiInstance: tmp,
	}
	return
}

// SetBandwidthAssignmentTable sets the value of BandwidthAssignmentTable for the instance
func (instance *MSFT_NetAdapter_QosSettings) SetPropertyBandwidthAssignmentTable(value []uint8) (err error) {
	return instance.SetProperty("BandwidthAssignmentTable", (value))
}

// GetBandwidthAssignmentTable gets the value of BandwidthAssignmentTable for the instance
func (instance *MSFT_NetAdapter_QosSettings) GetPropertyBandwidthAssignmentTable() (value []uint8, err error) {
	retValue, err := instance.GetProperty("BandwidthAssignmentTable")
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

// SetClassificationEnabled sets the value of ClassificationEnabled for the instance
func (instance *MSFT_NetAdapter_QosSettings) SetPropertyClassificationEnabled(value bool) (err error) {
	return instance.SetProperty("ClassificationEnabled", (value))
}

// GetClassificationEnabled gets the value of ClassificationEnabled for the instance
func (instance *MSFT_NetAdapter_QosSettings) GetPropertyClassificationEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("ClassificationEnabled")
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

// SetClassificationTable sets the value of ClassificationTable for the instance
func (instance *MSFT_NetAdapter_QosSettings) SetPropertyClassificationTable(value []MSFT_NetAdapter_QosClassificationElement) (err error) {
	return instance.SetProperty("ClassificationTable", (value))
}

// GetClassificationTable gets the value of ClassificationTable for the instance
func (instance *MSFT_NetAdapter_QosSettings) GetPropertyClassificationTable() (value []MSFT_NetAdapter_QosClassificationElement, err error) {
	retValue, err := instance.GetProperty("ClassificationTable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_NetAdapter_QosClassificationElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapter_QosClassificationElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_NetAdapter_QosClassificationElement(valuetmp))
	}

	return
}

// SetFlowControlEnabled sets the value of FlowControlEnabled for the instance
func (instance *MSFT_NetAdapter_QosSettings) SetPropertyFlowControlEnabled(value bool) (err error) {
	return instance.SetProperty("FlowControlEnabled", (value))
}

// GetFlowControlEnabled gets the value of FlowControlEnabled for the instance
func (instance *MSFT_NetAdapter_QosSettings) GetPropertyFlowControlEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("FlowControlEnabled")
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

// SetNumberOfClassificationElements sets the value of NumberOfClassificationElements for the instance
func (instance *MSFT_NetAdapter_QosSettings) SetPropertyNumberOfClassificationElements(value uint32) (err error) {
	return instance.SetProperty("NumberOfClassificationElements", (value))
}

// GetNumberOfClassificationElements gets the value of NumberOfClassificationElements for the instance
func (instance *MSFT_NetAdapter_QosSettings) GetPropertyNumberOfClassificationElements() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfClassificationElements")
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

// SetPriorityAssignmentTable sets the value of PriorityAssignmentTable for the instance
func (instance *MSFT_NetAdapter_QosSettings) SetPropertyPriorityAssignmentTable(value []uint8) (err error) {
	return instance.SetProperty("PriorityAssignmentTable", (value))
}

// GetPriorityAssignmentTable gets the value of PriorityAssignmentTable for the instance
func (instance *MSFT_NetAdapter_QosSettings) GetPropertyPriorityAssignmentTable() (value []uint8, err error) {
	retValue, err := instance.GetProperty("PriorityAssignmentTable")
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

// SetPriorityFlowControlEnableArray sets the value of PriorityFlowControlEnableArray for the instance
func (instance *MSFT_NetAdapter_QosSettings) SetPropertyPriorityFlowControlEnableArray(value []bool) (err error) {
	return instance.SetProperty("PriorityFlowControlEnableArray", (value))
}

// GetPriorityFlowControlEnableArray gets the value of PriorityFlowControlEnableArray for the instance
func (instance *MSFT_NetAdapter_QosSettings) GetPropertyPriorityFlowControlEnableArray() (value []bool, err error) {
	retValue, err := instance.GetProperty("PriorityFlowControlEnableArray")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(bool)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, bool(valuetmp))
	}

	return
}

// SetTransmissionSelectionEnabled sets the value of TransmissionSelectionEnabled for the instance
func (instance *MSFT_NetAdapter_QosSettings) SetPropertyTransmissionSelectionEnabled(value bool) (err error) {
	return instance.SetProperty("TransmissionSelectionEnabled", (value))
}

// GetTransmissionSelectionEnabled gets the value of TransmissionSelectionEnabled for the instance
func (instance *MSFT_NetAdapter_QosSettings) GetPropertyTransmissionSelectionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("TransmissionSelectionEnabled")
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

// SetTsaAssignmentTable sets the value of TsaAssignmentTable for the instance
func (instance *MSFT_NetAdapter_QosSettings) SetPropertyTsaAssignmentTable(value []uint8) (err error) {
	return instance.SetProperty("TsaAssignmentTable", (value))
}

// GetTsaAssignmentTable gets the value of TsaAssignmentTable for the instance
func (instance *MSFT_NetAdapter_QosSettings) GetPropertyTsaAssignmentTable() (value []uint8, err error) {
	retValue, err := instance.GetProperty("TsaAssignmentTable")
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
