// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.HardwareManagement
//
// ////////////////////////////////////////////
package hardwaremanagement

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_PCSVLogRecord struct
type MSFT_PCSVLogRecord struct {
	*CIM_LogRecord

	//
	RawData []uint8
}

func NewMSFT_PCSVLogRecordEx1(instance *cim.WmiInstance) (newInstance *MSFT_PCSVLogRecord, err error) {
	tmp, err := NewCIM_LogRecordEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_PCSVLogRecord{
		CIM_LogRecord: tmp,
	}
	return
}

func NewMSFT_PCSVLogRecordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_PCSVLogRecord, err error) {
	tmp, err := NewCIM_LogRecordEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_PCSVLogRecord{
		CIM_LogRecord: tmp,
	}
	return
}

// SetRawData sets the value of RawData for the instance
func (instance *MSFT_PCSVLogRecord) SetPropertyRawData(value []uint8) (err error) {
	return instance.SetProperty("RawData", (value))
}

// GetRawData gets the value of RawData for the instance
func (instance *MSFT_PCSVLogRecord) GetPropertyRawData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("RawData")
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
