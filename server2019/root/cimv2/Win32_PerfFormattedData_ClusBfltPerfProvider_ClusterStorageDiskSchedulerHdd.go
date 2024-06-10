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

// Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd struct
type Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd struct {
	*Win32_PerfFormattedData

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

func NewWin32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHddEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHddEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetBusyNormalPriRate sets the value of BusyNormalPriRate for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) SetPropertyBusyNormalPriRate(value uint64) (err error) {
	return instance.SetProperty("BusyNormalPriRate", (value))
}

// GetBusyNormalPriRate gets the value of BusyNormalPriRate for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) GetPropertyBusyNormalPriRate() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) SetPropertyBusyRate(value uint64) (err error) {
	return instance.SetProperty("BusyRate", (value))
}

// GetBusyRate gets the value of BusyRate for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) GetPropertyBusyRate() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) SetPropertyExecutionNormalPriRate(value uint64) (err error) {
	return instance.SetProperty("ExecutionNormalPriRate", (value))
}

// GetExecutionNormalPriRate gets the value of ExecutionNormalPriRate for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) GetPropertyExecutionNormalPriRate() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) SetPropertyNormalizedTransfersPersec(value uint64) (err error) {
	return instance.SetProperty("NormalizedTransfersPersec", (value))
}

// GetNormalizedTransfersPersec gets the value of NormalizedTransfersPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) GetPropertyNormalizedTransfersPersec() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) SetPropertySlowDisk(value uint64) (err error) {
	return instance.SetProperty("SlowDisk", (value))
}

// GetSlowDisk gets the value of SlowDisk for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) GetPropertySlowDisk() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) SetPropertyTimeBetweenIOCompAvgus(value uint64) (err error) {
	return instance.SetProperty("TimeBetweenIOCompAvgus", (value))
}

// GetTimeBetweenIOCompAvgus gets the value of TimeBetweenIOCompAvgus for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) GetPropertyTimeBetweenIOCompAvgus() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) SetPropertyTimeBetweenIOCompHighus(value uint64) (err error) {
	return instance.SetProperty("TimeBetweenIOCompHighus", (value))
}

// GetTimeBetweenIOCompHighus gets the value of TimeBetweenIOCompHighus for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) GetPropertyTimeBetweenIOCompHighus() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) SetPropertyTimeBetweenIOCompLowus(value uint64) (err error) {
	return instance.SetProperty("TimeBetweenIOCompLowus", (value))
}

// GetTimeBetweenIOCompLowus gets the value of TimeBetweenIOCompLowus for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageDiskSchedulerHdd) GetPropertyTimeBetweenIOCompLowus() (value uint64, err error) {
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
