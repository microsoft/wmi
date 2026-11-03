// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_Counters_ReFSDedupPerfCounters struct
type Win32_PerfFormattedData_Counters_ReFSDedupPerfCounters struct {
	*Win32_PerfFormattedData

	//
	ClustersduplicatedPersecond uint32

	//
	ClustersprocessedPersecond uint32

	//
	TableinsertioncountPersecond uint32
}

func NewWin32_PerfFormattedData_Counters_ReFSDedupPerfCountersEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_ReFSDedupPerfCounters, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_ReFSDedupPerfCounters{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_ReFSDedupPerfCountersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_ReFSDedupPerfCounters, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_ReFSDedupPerfCounters{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetClustersduplicatedPersecond sets the value of ClustersduplicatedPersecond for the instance
func (instance *Win32_PerfFormattedData_Counters_ReFSDedupPerfCounters) SetPropertyClustersduplicatedPersecond(value uint32) (err error) {
	return instance.SetProperty("ClustersduplicatedPersecond", (value))
}

// GetClustersduplicatedPersecond gets the value of ClustersduplicatedPersecond for the instance
func (instance *Win32_PerfFormattedData_Counters_ReFSDedupPerfCounters) GetPropertyClustersduplicatedPersecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClustersduplicatedPersecond")
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

// SetClustersprocessedPersecond sets the value of ClustersprocessedPersecond for the instance
func (instance *Win32_PerfFormattedData_Counters_ReFSDedupPerfCounters) SetPropertyClustersprocessedPersecond(value uint32) (err error) {
	return instance.SetProperty("ClustersprocessedPersecond", (value))
}

// GetClustersprocessedPersecond gets the value of ClustersprocessedPersecond for the instance
func (instance *Win32_PerfFormattedData_Counters_ReFSDedupPerfCounters) GetPropertyClustersprocessedPersecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClustersprocessedPersecond")
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

// SetTableinsertioncountPersecond sets the value of TableinsertioncountPersecond for the instance
func (instance *Win32_PerfFormattedData_Counters_ReFSDedupPerfCounters) SetPropertyTableinsertioncountPersecond(value uint32) (err error) {
	return instance.SetProperty("TableinsertioncountPersecond", (value))
}

// GetTableinsertioncountPersecond gets the value of TableinsertioncountPersecond for the instance
func (instance *Win32_PerfFormattedData_Counters_ReFSDedupPerfCounters) GetPropertyTableinsertioncountPersecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("TableinsertioncountPersecond")
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
