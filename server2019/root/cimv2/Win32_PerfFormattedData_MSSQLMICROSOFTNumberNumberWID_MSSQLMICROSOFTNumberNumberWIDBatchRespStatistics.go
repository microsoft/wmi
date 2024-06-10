// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2
//
// ////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics struct {
	*Win32_PerfFormattedData

	//
	Batches000000msAnd000001ms uint64

	//
	Batches000001msAnd000002ms uint64

	//
	Batches000002msAnd000005ms uint64

	//
	Batches000005msAnd000010ms uint64

	//
	Batches000010msAnd000020ms uint64

	//
	Batches000020msAnd000050ms uint64

	//
	Batches000050msAnd000100ms uint64

	//
	Batches000100msAnd000200ms uint64

	//
	Batches000200msAnd000500ms uint64

	//
	Batches000500msAnd001000ms uint64

	//
	Batches001000msAnd002000ms uint64

	//
	Batches002000msAnd005000ms uint64

	//
	Batches005000msAnd010000ms uint64

	//
	Batches010000msAnd020000ms uint64

	//
	Batches020000msAnd050000ms uint64

	//
	Batches050000msAnd100000ms uint64

	//
	Batches100000ms uint64
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatisticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetBatches000000msAnd000001ms sets the value of Batches000000msAnd000001ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches000000msAnd000001ms(value uint64) (err error) {
	return instance.SetProperty("Batches000000msAnd000001ms", (value))
}

// GetBatches000000msAnd000001ms gets the value of Batches000000msAnd000001ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches000000msAnd000001ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches000000msAnd000001ms")
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

// SetBatches000001msAnd000002ms sets the value of Batches000001msAnd000002ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches000001msAnd000002ms(value uint64) (err error) {
	return instance.SetProperty("Batches000001msAnd000002ms", (value))
}

// GetBatches000001msAnd000002ms gets the value of Batches000001msAnd000002ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches000001msAnd000002ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches000001msAnd000002ms")
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

// SetBatches000002msAnd000005ms sets the value of Batches000002msAnd000005ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches000002msAnd000005ms(value uint64) (err error) {
	return instance.SetProperty("Batches000002msAnd000005ms", (value))
}

// GetBatches000002msAnd000005ms gets the value of Batches000002msAnd000005ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches000002msAnd000005ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches000002msAnd000005ms")
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

// SetBatches000005msAnd000010ms sets the value of Batches000005msAnd000010ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches000005msAnd000010ms(value uint64) (err error) {
	return instance.SetProperty("Batches000005msAnd000010ms", (value))
}

// GetBatches000005msAnd000010ms gets the value of Batches000005msAnd000010ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches000005msAnd000010ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches000005msAnd000010ms")
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

// SetBatches000010msAnd000020ms sets the value of Batches000010msAnd000020ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches000010msAnd000020ms(value uint64) (err error) {
	return instance.SetProperty("Batches000010msAnd000020ms", (value))
}

// GetBatches000010msAnd000020ms gets the value of Batches000010msAnd000020ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches000010msAnd000020ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches000010msAnd000020ms")
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

// SetBatches000020msAnd000050ms sets the value of Batches000020msAnd000050ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches000020msAnd000050ms(value uint64) (err error) {
	return instance.SetProperty("Batches000020msAnd000050ms", (value))
}

// GetBatches000020msAnd000050ms gets the value of Batches000020msAnd000050ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches000020msAnd000050ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches000020msAnd000050ms")
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

// SetBatches000050msAnd000100ms sets the value of Batches000050msAnd000100ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches000050msAnd000100ms(value uint64) (err error) {
	return instance.SetProperty("Batches000050msAnd000100ms", (value))
}

// GetBatches000050msAnd000100ms gets the value of Batches000050msAnd000100ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches000050msAnd000100ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches000050msAnd000100ms")
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

// SetBatches000100msAnd000200ms sets the value of Batches000100msAnd000200ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches000100msAnd000200ms(value uint64) (err error) {
	return instance.SetProperty("Batches000100msAnd000200ms", (value))
}

// GetBatches000100msAnd000200ms gets the value of Batches000100msAnd000200ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches000100msAnd000200ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches000100msAnd000200ms")
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

// SetBatches000200msAnd000500ms sets the value of Batches000200msAnd000500ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches000200msAnd000500ms(value uint64) (err error) {
	return instance.SetProperty("Batches000200msAnd000500ms", (value))
}

// GetBatches000200msAnd000500ms gets the value of Batches000200msAnd000500ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches000200msAnd000500ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches000200msAnd000500ms")
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

// SetBatches000500msAnd001000ms sets the value of Batches000500msAnd001000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches000500msAnd001000ms(value uint64) (err error) {
	return instance.SetProperty("Batches000500msAnd001000ms", (value))
}

// GetBatches000500msAnd001000ms gets the value of Batches000500msAnd001000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches000500msAnd001000ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches000500msAnd001000ms")
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

// SetBatches001000msAnd002000ms sets the value of Batches001000msAnd002000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches001000msAnd002000ms(value uint64) (err error) {
	return instance.SetProperty("Batches001000msAnd002000ms", (value))
}

// GetBatches001000msAnd002000ms gets the value of Batches001000msAnd002000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches001000msAnd002000ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches001000msAnd002000ms")
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

// SetBatches002000msAnd005000ms sets the value of Batches002000msAnd005000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches002000msAnd005000ms(value uint64) (err error) {
	return instance.SetProperty("Batches002000msAnd005000ms", (value))
}

// GetBatches002000msAnd005000ms gets the value of Batches002000msAnd005000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches002000msAnd005000ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches002000msAnd005000ms")
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

// SetBatches005000msAnd010000ms sets the value of Batches005000msAnd010000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches005000msAnd010000ms(value uint64) (err error) {
	return instance.SetProperty("Batches005000msAnd010000ms", (value))
}

// GetBatches005000msAnd010000ms gets the value of Batches005000msAnd010000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches005000msAnd010000ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches005000msAnd010000ms")
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

// SetBatches010000msAnd020000ms sets the value of Batches010000msAnd020000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches010000msAnd020000ms(value uint64) (err error) {
	return instance.SetProperty("Batches010000msAnd020000ms", (value))
}

// GetBatches010000msAnd020000ms gets the value of Batches010000msAnd020000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches010000msAnd020000ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches010000msAnd020000ms")
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

// SetBatches020000msAnd050000ms sets the value of Batches020000msAnd050000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches020000msAnd050000ms(value uint64) (err error) {
	return instance.SetProperty("Batches020000msAnd050000ms", (value))
}

// GetBatches020000msAnd050000ms gets the value of Batches020000msAnd050000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches020000msAnd050000ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches020000msAnd050000ms")
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

// SetBatches050000msAnd100000ms sets the value of Batches050000msAnd100000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches050000msAnd100000ms(value uint64) (err error) {
	return instance.SetProperty("Batches050000msAnd100000ms", (value))
}

// GetBatches050000msAnd100000ms gets the value of Batches050000msAnd100000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches050000msAnd100000ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches050000msAnd100000ms")
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

// SetBatches100000ms sets the value of Batches100000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) SetPropertyBatches100000ms(value uint64) (err error) {
	return instance.SetProperty("Batches100000ms", (value))
}

// GetBatches100000ms gets the value of Batches100000ms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBatchRespStatistics) GetPropertyBatches100000ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("Batches100000ms")
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
