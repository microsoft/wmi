// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MLNX_NetAdapterFlowControlSettingData struct
type MLNX_NetAdapterFlowControlSettingData struct {
	*MLNX_NetAdapterSettingData

	//
	PFCRx string

	//
	PFCTx string

	//
	RroceDscpMarkPriorityFlowControl_0 uint8

	//
	RroceDscpMarkPriorityFlowControl_1 uint8

	//
	RroceDscpMarkPriorityFlowControl_2 uint8

	//
	RroceDscpMarkPriorityFlowControl_3 uint8

	//
	RroceDscpMarkPriorityFlowControl_4 uint8

	//
	RroceDscpMarkPriorityFlowControl_5 uint8

	//
	RroceDscpMarkPriorityFlowControl_6 uint8

	//
	RroceDscpMarkPriorityFlowControl_7 uint8

	//
	RxUntaggedMapToLossless uint8

	//
	TxUntagPriorityTag uint8
}

func NewMLNX_NetAdapterFlowControlSettingDataEx1(instance *cim.WmiInstance) (newInstance *MLNX_NetAdapterFlowControlSettingData, err error) {
	tmp, err := NewMLNX_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_NetAdapterFlowControlSettingData{
		MLNX_NetAdapterSettingData: tmp,
	}
	return
}

func NewMLNX_NetAdapterFlowControlSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_NetAdapterFlowControlSettingData, err error) {
	tmp, err := NewMLNX_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_NetAdapterFlowControlSettingData{
		MLNX_NetAdapterSettingData: tmp,
	}
	return
}

// SetPFCRx sets the value of PFCRx for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) SetPropertyPFCRx(value string) (err error) {
	return instance.SetProperty("PFCRx", (value))
}

// GetPFCRx gets the value of PFCRx for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) GetPropertyPFCRx() (value string, err error) {
	retValue, err := instance.GetProperty("PFCRx")
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

// SetPFCTx sets the value of PFCTx for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) SetPropertyPFCTx(value string) (err error) {
	return instance.SetProperty("PFCTx", (value))
}

// GetPFCTx gets the value of PFCTx for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) GetPropertyPFCTx() (value string, err error) {
	retValue, err := instance.GetProperty("PFCTx")
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

// SetRroceDscpMarkPriorityFlowControl_0 sets the value of RroceDscpMarkPriorityFlowControl_0 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) SetPropertyRroceDscpMarkPriorityFlowControl_0(value uint8) (err error) {
	return instance.SetProperty("RroceDscpMarkPriorityFlowControl_0", (value))
}

// GetRroceDscpMarkPriorityFlowControl_0 gets the value of RroceDscpMarkPriorityFlowControl_0 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) GetPropertyRroceDscpMarkPriorityFlowControl_0() (value uint8, err error) {
	retValue, err := instance.GetProperty("RroceDscpMarkPriorityFlowControl_0")
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

// SetRroceDscpMarkPriorityFlowControl_1 sets the value of RroceDscpMarkPriorityFlowControl_1 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) SetPropertyRroceDscpMarkPriorityFlowControl_1(value uint8) (err error) {
	return instance.SetProperty("RroceDscpMarkPriorityFlowControl_1", (value))
}

// GetRroceDscpMarkPriorityFlowControl_1 gets the value of RroceDscpMarkPriorityFlowControl_1 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) GetPropertyRroceDscpMarkPriorityFlowControl_1() (value uint8, err error) {
	retValue, err := instance.GetProperty("RroceDscpMarkPriorityFlowControl_1")
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

// SetRroceDscpMarkPriorityFlowControl_2 sets the value of RroceDscpMarkPriorityFlowControl_2 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) SetPropertyRroceDscpMarkPriorityFlowControl_2(value uint8) (err error) {
	return instance.SetProperty("RroceDscpMarkPriorityFlowControl_2", (value))
}

// GetRroceDscpMarkPriorityFlowControl_2 gets the value of RroceDscpMarkPriorityFlowControl_2 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) GetPropertyRroceDscpMarkPriorityFlowControl_2() (value uint8, err error) {
	retValue, err := instance.GetProperty("RroceDscpMarkPriorityFlowControl_2")
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

// SetRroceDscpMarkPriorityFlowControl_3 sets the value of RroceDscpMarkPriorityFlowControl_3 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) SetPropertyRroceDscpMarkPriorityFlowControl_3(value uint8) (err error) {
	return instance.SetProperty("RroceDscpMarkPriorityFlowControl_3", (value))
}

// GetRroceDscpMarkPriorityFlowControl_3 gets the value of RroceDscpMarkPriorityFlowControl_3 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) GetPropertyRroceDscpMarkPriorityFlowControl_3() (value uint8, err error) {
	retValue, err := instance.GetProperty("RroceDscpMarkPriorityFlowControl_3")
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

// SetRroceDscpMarkPriorityFlowControl_4 sets the value of RroceDscpMarkPriorityFlowControl_4 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) SetPropertyRroceDscpMarkPriorityFlowControl_4(value uint8) (err error) {
	return instance.SetProperty("RroceDscpMarkPriorityFlowControl_4", (value))
}

// GetRroceDscpMarkPriorityFlowControl_4 gets the value of RroceDscpMarkPriorityFlowControl_4 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) GetPropertyRroceDscpMarkPriorityFlowControl_4() (value uint8, err error) {
	retValue, err := instance.GetProperty("RroceDscpMarkPriorityFlowControl_4")
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

// SetRroceDscpMarkPriorityFlowControl_5 sets the value of RroceDscpMarkPriorityFlowControl_5 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) SetPropertyRroceDscpMarkPriorityFlowControl_5(value uint8) (err error) {
	return instance.SetProperty("RroceDscpMarkPriorityFlowControl_5", (value))
}

// GetRroceDscpMarkPriorityFlowControl_5 gets the value of RroceDscpMarkPriorityFlowControl_5 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) GetPropertyRroceDscpMarkPriorityFlowControl_5() (value uint8, err error) {
	retValue, err := instance.GetProperty("RroceDscpMarkPriorityFlowControl_5")
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

// SetRroceDscpMarkPriorityFlowControl_6 sets the value of RroceDscpMarkPriorityFlowControl_6 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) SetPropertyRroceDscpMarkPriorityFlowControl_6(value uint8) (err error) {
	return instance.SetProperty("RroceDscpMarkPriorityFlowControl_6", (value))
}

// GetRroceDscpMarkPriorityFlowControl_6 gets the value of RroceDscpMarkPriorityFlowControl_6 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) GetPropertyRroceDscpMarkPriorityFlowControl_6() (value uint8, err error) {
	retValue, err := instance.GetProperty("RroceDscpMarkPriorityFlowControl_6")
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

// SetRroceDscpMarkPriorityFlowControl_7 sets the value of RroceDscpMarkPriorityFlowControl_7 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) SetPropertyRroceDscpMarkPriorityFlowControl_7(value uint8) (err error) {
	return instance.SetProperty("RroceDscpMarkPriorityFlowControl_7", (value))
}

// GetRroceDscpMarkPriorityFlowControl_7 gets the value of RroceDscpMarkPriorityFlowControl_7 for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) GetPropertyRroceDscpMarkPriorityFlowControl_7() (value uint8, err error) {
	retValue, err := instance.GetProperty("RroceDscpMarkPriorityFlowControl_7")
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

// SetRxUntaggedMapToLossless sets the value of RxUntaggedMapToLossless for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) SetPropertyRxUntaggedMapToLossless(value uint8) (err error) {
	return instance.SetProperty("RxUntaggedMapToLossless", (value))
}

// GetRxUntaggedMapToLossless gets the value of RxUntaggedMapToLossless for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) GetPropertyRxUntaggedMapToLossless() (value uint8, err error) {
	retValue, err := instance.GetProperty("RxUntaggedMapToLossless")
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

// SetTxUntagPriorityTag sets the value of TxUntagPriorityTag for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) SetPropertyTxUntagPriorityTag(value uint8) (err error) {
	return instance.SetProperty("TxUntagPriorityTag", (value))
}

// GetTxUntagPriorityTag gets the value of TxUntagPriorityTag for the instance
func (instance *MLNX_NetAdapterFlowControlSettingData) GetPropertyTxUntagPriorityTag() (value uint8, err error) {
	retValue, err := instance.GetProperty("TxUntagPriorityTag")
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
