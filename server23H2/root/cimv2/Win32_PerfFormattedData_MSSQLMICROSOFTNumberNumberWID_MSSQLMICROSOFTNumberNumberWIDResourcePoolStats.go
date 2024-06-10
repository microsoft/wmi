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

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats struct {
	*Win32_PerfFormattedData

	//
	ActivememorygrantamountKB uint64

	//
	Activememorygrantscount uint64

	//
	AvgDiskReadIOms uint64

	//
	AvgDiskWriteIOms uint64

	//
	CachememorytargetKB uint64

	//
	CompilememorytargetKB uint64

	//
	CPUcontroleffectPercent uint64

	//
	CPUusagePercent uint64

	//
	CPUusagetargetPercent uint64

	//
	DiskReadBytesPersec uint64

	//
	DiskReadIOPersec uint64

	//
	DiskReadIOThrottledPersec uint64

	//
	DiskWriteBytesPersec uint64

	//
	DiskWriteIOPersec uint64

	//
	DiskWriteIOThrottledPersec uint64

	//
	MaxmemoryKB uint64

	//
	MemorygrantsPersec uint64

	//
	MemorygranttimeoutsPersec uint64

	//
	Pendingmemorygrantscount uint64

	//
	QueryexecmemorytargetKB uint64

	//
	TargetmemoryKB uint64

	//
	UsedmemoryKB uint64
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStatsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStatsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetActivememorygrantamountKB sets the value of ActivememorygrantamountKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyActivememorygrantamountKB(value uint64) (err error) {
	return instance.SetProperty("ActivememorygrantamountKB", (value))
}

// GetActivememorygrantamountKB gets the value of ActivememorygrantamountKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyActivememorygrantamountKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("ActivememorygrantamountKB")
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

// SetActivememorygrantscount sets the value of Activememorygrantscount for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyActivememorygrantscount(value uint64) (err error) {
	return instance.SetProperty("Activememorygrantscount", (value))
}

// GetActivememorygrantscount gets the value of Activememorygrantscount for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyActivememorygrantscount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Activememorygrantscount")
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

// SetAvgDiskReadIOms sets the value of AvgDiskReadIOms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyAvgDiskReadIOms(value uint64) (err error) {
	return instance.SetProperty("AvgDiskReadIOms", (value))
}

// GetAvgDiskReadIOms gets the value of AvgDiskReadIOms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyAvgDiskReadIOms() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgDiskReadIOms")
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

// SetAvgDiskWriteIOms sets the value of AvgDiskWriteIOms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyAvgDiskWriteIOms(value uint64) (err error) {
	return instance.SetProperty("AvgDiskWriteIOms", (value))
}

// GetAvgDiskWriteIOms gets the value of AvgDiskWriteIOms for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyAvgDiskWriteIOms() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgDiskWriteIOms")
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

// SetCachememorytargetKB sets the value of CachememorytargetKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyCachememorytargetKB(value uint64) (err error) {
	return instance.SetProperty("CachememorytargetKB", (value))
}

// GetCachememorytargetKB gets the value of CachememorytargetKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyCachememorytargetKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("CachememorytargetKB")
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

// SetCompilememorytargetKB sets the value of CompilememorytargetKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyCompilememorytargetKB(value uint64) (err error) {
	return instance.SetProperty("CompilememorytargetKB", (value))
}

// GetCompilememorytargetKB gets the value of CompilememorytargetKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyCompilememorytargetKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("CompilememorytargetKB")
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

// SetCPUcontroleffectPercent sets the value of CPUcontroleffectPercent for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyCPUcontroleffectPercent(value uint64) (err error) {
	return instance.SetProperty("CPUcontroleffectPercent", (value))
}

// GetCPUcontroleffectPercent gets the value of CPUcontroleffectPercent for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyCPUcontroleffectPercent() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUcontroleffectPercent")
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

// SetCPUusagePercent sets the value of CPUusagePercent for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyCPUusagePercent(value uint64) (err error) {
	return instance.SetProperty("CPUusagePercent", (value))
}

// GetCPUusagePercent gets the value of CPUusagePercent for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyCPUusagePercent() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUusagePercent")
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

// SetCPUusagetargetPercent sets the value of CPUusagetargetPercent for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyCPUusagetargetPercent(value uint64) (err error) {
	return instance.SetProperty("CPUusagetargetPercent", (value))
}

// GetCPUusagetargetPercent gets the value of CPUusagetargetPercent for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyCPUusagetargetPercent() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUusagetargetPercent")
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

// SetDiskReadBytesPersec sets the value of DiskReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyDiskReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("DiskReadBytesPersec", (value))
}

// GetDiskReadBytesPersec gets the value of DiskReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyDiskReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DiskReadBytesPersec")
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

// SetDiskReadIOPersec sets the value of DiskReadIOPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyDiskReadIOPersec(value uint64) (err error) {
	return instance.SetProperty("DiskReadIOPersec", (value))
}

// GetDiskReadIOPersec gets the value of DiskReadIOPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyDiskReadIOPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DiskReadIOPersec")
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

// SetDiskReadIOThrottledPersec sets the value of DiskReadIOThrottledPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyDiskReadIOThrottledPersec(value uint64) (err error) {
	return instance.SetProperty("DiskReadIOThrottledPersec", (value))
}

// GetDiskReadIOThrottledPersec gets the value of DiskReadIOThrottledPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyDiskReadIOThrottledPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DiskReadIOThrottledPersec")
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

// SetDiskWriteBytesPersec sets the value of DiskWriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyDiskWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("DiskWriteBytesPersec", (value))
}

// GetDiskWriteBytesPersec gets the value of DiskWriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyDiskWriteBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DiskWriteBytesPersec")
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

// SetDiskWriteIOPersec sets the value of DiskWriteIOPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyDiskWriteIOPersec(value uint64) (err error) {
	return instance.SetProperty("DiskWriteIOPersec", (value))
}

// GetDiskWriteIOPersec gets the value of DiskWriteIOPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyDiskWriteIOPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DiskWriteIOPersec")
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

// SetDiskWriteIOThrottledPersec sets the value of DiskWriteIOThrottledPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyDiskWriteIOThrottledPersec(value uint64) (err error) {
	return instance.SetProperty("DiskWriteIOThrottledPersec", (value))
}

// GetDiskWriteIOThrottledPersec gets the value of DiskWriteIOThrottledPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyDiskWriteIOThrottledPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DiskWriteIOThrottledPersec")
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

// SetMaxmemoryKB sets the value of MaxmemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyMaxmemoryKB(value uint64) (err error) {
	return instance.SetProperty("MaxmemoryKB", (value))
}

// GetMaxmemoryKB gets the value of MaxmemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyMaxmemoryKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxmemoryKB")
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

// SetMemorygrantsPersec sets the value of MemorygrantsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyMemorygrantsPersec(value uint64) (err error) {
	return instance.SetProperty("MemorygrantsPersec", (value))
}

// GetMemorygrantsPersec gets the value of MemorygrantsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyMemorygrantsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemorygrantsPersec")
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

// SetMemorygranttimeoutsPersec sets the value of MemorygranttimeoutsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyMemorygranttimeoutsPersec(value uint64) (err error) {
	return instance.SetProperty("MemorygranttimeoutsPersec", (value))
}

// GetMemorygranttimeoutsPersec gets the value of MemorygranttimeoutsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyMemorygranttimeoutsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemorygranttimeoutsPersec")
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

// SetPendingmemorygrantscount sets the value of Pendingmemorygrantscount for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyPendingmemorygrantscount(value uint64) (err error) {
	return instance.SetProperty("Pendingmemorygrantscount", (value))
}

// GetPendingmemorygrantscount gets the value of Pendingmemorygrantscount for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyPendingmemorygrantscount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Pendingmemorygrantscount")
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

// SetQueryexecmemorytargetKB sets the value of QueryexecmemorytargetKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyQueryexecmemorytargetKB(value uint64) (err error) {
	return instance.SetProperty("QueryexecmemorytargetKB", (value))
}

// GetQueryexecmemorytargetKB gets the value of QueryexecmemorytargetKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyQueryexecmemorytargetKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("QueryexecmemorytargetKB")
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

// SetTargetmemoryKB sets the value of TargetmemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyTargetmemoryKB(value uint64) (err error) {
	return instance.SetProperty("TargetmemoryKB", (value))
}

// GetTargetmemoryKB gets the value of TargetmemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyTargetmemoryKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("TargetmemoryKB")
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

// SetUsedmemoryKB sets the value of UsedmemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) SetPropertyUsedmemoryKB(value uint64) (err error) {
	return instance.SetProperty("UsedmemoryKB", (value))
}

// GetUsedmemoryKB gets the value of UsedmemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDResourcePoolStats) GetPropertyUsedmemoryKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("UsedmemoryKB")
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
