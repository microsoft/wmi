// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics struct {
	*Win32_PerfFormattedData

	//
	AvgLengthofBatchedWrites uint64

	//
	AvgTimeBetweenBatchesms uint64

	//
	AvgTimetoWriteBatchms uint64

	//
	TransmissionObjGetsPerSec uint64

	//
	TransmissionObjSetDirtyPerSec uint64

	//
	TransmissionObjWritesPerSec uint64
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatisticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetAvgLengthofBatchedWrites sets the value of AvgLengthofBatchedWrites for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics) SetPropertyAvgLengthofBatchedWrites(value uint64) (err error) {
	return instance.SetProperty("AvgLengthofBatchedWrites", (value))
}

// GetAvgLengthofBatchedWrites gets the value of AvgLengthofBatchedWrites for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics) GetPropertyAvgLengthofBatchedWrites() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgLengthofBatchedWrites")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAvgTimeBetweenBatchesms sets the value of AvgTimeBetweenBatchesms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics) SetPropertyAvgTimeBetweenBatchesms(value uint64) (err error) {
	return instance.SetProperty("AvgTimeBetweenBatchesms", (value))
}

// GetAvgTimeBetweenBatchesms gets the value of AvgTimeBetweenBatchesms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics) GetPropertyAvgTimeBetweenBatchesms() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgTimeBetweenBatchesms")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAvgTimetoWriteBatchms sets the value of AvgTimetoWriteBatchms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics) SetPropertyAvgTimetoWriteBatchms(value uint64) (err error) {
	return instance.SetProperty("AvgTimetoWriteBatchms", (value))
}

// GetAvgTimetoWriteBatchms gets the value of AvgTimetoWriteBatchms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics) GetPropertyAvgTimetoWriteBatchms() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgTimetoWriteBatchms")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTransmissionObjGetsPerSec sets the value of TransmissionObjGetsPerSec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics) SetPropertyTransmissionObjGetsPerSec(value uint64) (err error) {
	return instance.SetProperty("TransmissionObjGetsPerSec", (value))
}

// GetTransmissionObjGetsPerSec gets the value of TransmissionObjGetsPerSec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics) GetPropertyTransmissionObjGetsPerSec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransmissionObjGetsPerSec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTransmissionObjSetDirtyPerSec sets the value of TransmissionObjSetDirtyPerSec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics) SetPropertyTransmissionObjSetDirtyPerSec(value uint64) (err error) {
	return instance.SetProperty("TransmissionObjSetDirtyPerSec", (value))
}

// GetTransmissionObjSetDirtyPerSec gets the value of TransmissionObjSetDirtyPerSec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics) GetPropertyTransmissionObjSetDirtyPerSec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransmissionObjSetDirtyPerSec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTransmissionObjWritesPerSec sets the value of TransmissionObjWritesPerSec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics) SetPropertyTransmissionObjWritesPerSec(value uint64) (err error) {
	return instance.SetProperty("TransmissionObjWritesPerSec", (value))
}

// GetTransmissionObjWritesPerSec gets the value of TransmissionObjWritesPerSec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerTOStatistics) GetPropertyTransmissionObjWritesPerSec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransmissionObjWritesPerSec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
