// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd struct
type Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd struct {
	*Win32_PerfRawData

	//
	BusyNormalPriRate uint64

	//
	BusyRate uint64

	//
	ExecutionNormalPriRate uint64

	//
	NormalizedTransfersPersec uint64

	//
	SlowDisk uint64

	//
	TimeBetweenIOCompAvgus uint64

	//
	TimeBetweenIOCompHighus uint64

	//
	TimeBetweenIOCompLowus uint64
}

func NewWin32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsdEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetBusyNormalPriRate sets the value of BusyNormalPriRate for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) SetPropertyBusyNormalPriRate(value uint64) (err error) {
	return instance.SetProperty("BusyNormalPriRate", (value))
}

// GetBusyNormalPriRate gets the value of BusyNormalPriRate for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) GetPropertyBusyNormalPriRate() (value uint64, err error) {
	retValue, err := instance.GetProperty("BusyNormalPriRate")
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

// SetBusyRate sets the value of BusyRate for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) SetPropertyBusyRate(value uint64) (err error) {
	return instance.SetProperty("BusyRate", (value))
}

// GetBusyRate gets the value of BusyRate for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) GetPropertyBusyRate() (value uint64, err error) {
	retValue, err := instance.GetProperty("BusyRate")
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

// SetExecutionNormalPriRate sets the value of ExecutionNormalPriRate for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) SetPropertyExecutionNormalPriRate(value uint64) (err error) {
	return instance.SetProperty("ExecutionNormalPriRate", (value))
}

// GetExecutionNormalPriRate gets the value of ExecutionNormalPriRate for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) GetPropertyExecutionNormalPriRate() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExecutionNormalPriRate")
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

// SetNormalizedTransfersPersec sets the value of NormalizedTransfersPersec for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) SetPropertyNormalizedTransfersPersec(value uint64) (err error) {
	return instance.SetProperty("NormalizedTransfersPersec", (value))
}

// GetNormalizedTransfersPersec gets the value of NormalizedTransfersPersec for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) GetPropertyNormalizedTransfersPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NormalizedTransfersPersec")
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

// SetSlowDisk sets the value of SlowDisk for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) SetPropertySlowDisk(value uint64) (err error) {
	return instance.SetProperty("SlowDisk", (value))
}

// GetSlowDisk gets the value of SlowDisk for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) GetPropertySlowDisk() (value uint64, err error) {
	retValue, err := instance.GetProperty("SlowDisk")
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

// SetTimeBetweenIOCompAvgus sets the value of TimeBetweenIOCompAvgus for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) SetPropertyTimeBetweenIOCompAvgus(value uint64) (err error) {
	return instance.SetProperty("TimeBetweenIOCompAvgus", (value))
}

// GetTimeBetweenIOCompAvgus gets the value of TimeBetweenIOCompAvgus for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) GetPropertyTimeBetweenIOCompAvgus() (value uint64, err error) {
	retValue, err := instance.GetProperty("TimeBetweenIOCompAvgus")
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

// SetTimeBetweenIOCompHighus sets the value of TimeBetweenIOCompHighus for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) SetPropertyTimeBetweenIOCompHighus(value uint64) (err error) {
	return instance.SetProperty("TimeBetweenIOCompHighus", (value))
}

// GetTimeBetweenIOCompHighus gets the value of TimeBetweenIOCompHighus for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) GetPropertyTimeBetweenIOCompHighus() (value uint64, err error) {
	retValue, err := instance.GetProperty("TimeBetweenIOCompHighus")
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

// SetTimeBetweenIOCompLowus sets the value of TimeBetweenIOCompLowus for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) SetPropertyTimeBetweenIOCompLowus(value uint64) (err error) {
	return instance.SetProperty("TimeBetweenIOCompLowus", (value))
}

// GetTimeBetweenIOCompLowus gets the value of TimeBetweenIOCompLowus for the instance
func (instance *Win32_PerfRawData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerSsd) GetPropertyTimeBetweenIOCompLowus() (value uint64, err error) {
	retValue, err := instance.GetProperty("TimeBetweenIOCompLowus")
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
