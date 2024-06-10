// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.ManagementTools
//
// ////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MTMemorySummary struct
type MSFT_MTMemorySummary struct {
	*CIM_ManagedElement

	//
	Available uint64

	//
	Cached uint64

	//
	Capacity uint64

	//
	CommitLimit uint64

	//
	Committed uint64

	//
	CurrentIndex uint16

	//
	DynamicMemoryEnabled bool

	//
	DynamicMemoryMax uint64

	//
	FormFactor uint16

	//
	Free uint64

	//
	HardwareReserved uint64

	//
	Installed uint64

	//
	IntervalSeconds uint16

	//
	InUse uint64

	//
	Modified uint64

	//
	Name string

	//
	NonPagedPool uint64

	//
	PagedPool uint64

	//
	PageSize uint32

	//
	Speed uint32

	//
	Standby uint64

	//
	Total uint64

	//
	TotalSlots uint16

	//
	Type uint16

	//
	UsedSlots uint16

	//
	Utilization []float32
}

func NewMSFT_MTMemorySummaryEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTMemorySummary, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTMemorySummary{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_MTMemorySummaryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MTMemorySummary, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTMemorySummary{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetAvailable sets the value of Available for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyAvailable(value uint64) (err error) {
	return instance.SetProperty("Available", (value))
}

// GetAvailable gets the value of Available for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyAvailable() (value uint64, err error) {
	retValue, err := instance.GetProperty("Available")
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

// SetCached sets the value of Cached for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyCached(value uint64) (err error) {
	return instance.SetProperty("Cached", (value))
}

// GetCached gets the value of Cached for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyCached() (value uint64, err error) {
	retValue, err := instance.GetProperty("Cached")
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

// SetCapacity sets the value of Capacity for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyCapacity(value uint64) (err error) {
	return instance.SetProperty("Capacity", (value))
}

// GetCapacity gets the value of Capacity for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyCapacity() (value uint64, err error) {
	retValue, err := instance.GetProperty("Capacity")
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

// SetCommitLimit sets the value of CommitLimit for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyCommitLimit(value uint64) (err error) {
	return instance.SetProperty("CommitLimit", (value))
}

// GetCommitLimit gets the value of CommitLimit for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyCommitLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("CommitLimit")
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

// SetCommitted sets the value of Committed for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyCommitted(value uint64) (err error) {
	return instance.SetProperty("Committed", (value))
}

// GetCommitted gets the value of Committed for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyCommitted() (value uint64, err error) {
	retValue, err := instance.GetProperty("Committed")
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

// SetCurrentIndex sets the value of CurrentIndex for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyCurrentIndex(value uint16) (err error) {
	return instance.SetProperty("CurrentIndex", (value))
}

// GetCurrentIndex gets the value of CurrentIndex for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyCurrentIndex() (value uint16, err error) {
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

// SetDynamicMemoryEnabled sets the value of DynamicMemoryEnabled for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyDynamicMemoryEnabled(value bool) (err error) {
	return instance.SetProperty("DynamicMemoryEnabled", (value))
}

// GetDynamicMemoryEnabled gets the value of DynamicMemoryEnabled for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyDynamicMemoryEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DynamicMemoryEnabled")
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

// SetDynamicMemoryMax sets the value of DynamicMemoryMax for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyDynamicMemoryMax(value uint64) (err error) {
	return instance.SetProperty("DynamicMemoryMax", (value))
}

// GetDynamicMemoryMax gets the value of DynamicMemoryMax for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyDynamicMemoryMax() (value uint64, err error) {
	retValue, err := instance.GetProperty("DynamicMemoryMax")
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

// SetFormFactor sets the value of FormFactor for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyFormFactor(value uint16) (err error) {
	return instance.SetProperty("FormFactor", (value))
}

// GetFormFactor gets the value of FormFactor for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyFormFactor() (value uint16, err error) {
	retValue, err := instance.GetProperty("FormFactor")
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

// SetFree sets the value of Free for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyFree(value uint64) (err error) {
	return instance.SetProperty("Free", (value))
}

// GetFree gets the value of Free for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyFree() (value uint64, err error) {
	retValue, err := instance.GetProperty("Free")
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

// SetHardwareReserved sets the value of HardwareReserved for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyHardwareReserved(value uint64) (err error) {
	return instance.SetProperty("HardwareReserved", (value))
}

// GetHardwareReserved gets the value of HardwareReserved for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyHardwareReserved() (value uint64, err error) {
	retValue, err := instance.GetProperty("HardwareReserved")
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

// SetInstalled sets the value of Installed for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyInstalled(value uint64) (err error) {
	return instance.SetProperty("Installed", (value))
}

// GetInstalled gets the value of Installed for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyInstalled() (value uint64, err error) {
	retValue, err := instance.GetProperty("Installed")
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

// SetIntervalSeconds sets the value of IntervalSeconds for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyIntervalSeconds(value uint16) (err error) {
	return instance.SetProperty("IntervalSeconds", (value))
}

// GetIntervalSeconds gets the value of IntervalSeconds for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyIntervalSeconds() (value uint16, err error) {
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

// SetInUse sets the value of InUse for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyInUse(value uint64) (err error) {
	return instance.SetProperty("InUse", (value))
}

// GetInUse gets the value of InUse for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyInUse() (value uint64, err error) {
	retValue, err := instance.GetProperty("InUse")
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

// SetModified sets the value of Modified for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyModified(value uint64) (err error) {
	return instance.SetProperty("Modified", (value))
}

// GetModified gets the value of Modified for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyModified() (value uint64, err error) {
	retValue, err := instance.GetProperty("Modified")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetNonPagedPool sets the value of NonPagedPool for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyNonPagedPool(value uint64) (err error) {
	return instance.SetProperty("NonPagedPool", (value))
}

// GetNonPagedPool gets the value of NonPagedPool for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyNonPagedPool() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonPagedPool")
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

// SetPagedPool sets the value of PagedPool for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyPagedPool(value uint64) (err error) {
	return instance.SetProperty("PagedPool", (value))
}

// GetPagedPool gets the value of PagedPool for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyPagedPool() (value uint64, err error) {
	retValue, err := instance.GetProperty("PagedPool")
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

// SetPageSize sets the value of PageSize for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyPageSize(value uint32) (err error) {
	return instance.SetProperty("PageSize", (value))
}

// GetPageSize gets the value of PageSize for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyPageSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PageSize")
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

// SetSpeed sets the value of Speed for the instance
func (instance *MSFT_MTMemorySummary) SetPropertySpeed(value uint32) (err error) {
	return instance.SetProperty("Speed", (value))
}

// GetSpeed gets the value of Speed for the instance
func (instance *MSFT_MTMemorySummary) GetPropertySpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("Speed")
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

// SetStandby sets the value of Standby for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyStandby(value uint64) (err error) {
	return instance.SetProperty("Standby", (value))
}

// GetStandby gets the value of Standby for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyStandby() (value uint64, err error) {
	retValue, err := instance.GetProperty("Standby")
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

// SetTotal sets the value of Total for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyTotal(value uint64) (err error) {
	return instance.SetProperty("Total", (value))
}

// GetTotal gets the value of Total for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("Total")
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

// SetTotalSlots sets the value of TotalSlots for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyTotalSlots(value uint16) (err error) {
	return instance.SetProperty("TotalSlots", (value))
}

// GetTotalSlots gets the value of TotalSlots for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyTotalSlots() (value uint16, err error) {
	retValue, err := instance.GetProperty("TotalSlots")
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

// SetType sets the value of Type for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyType(value uint16) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyType() (value uint16, err error) {
	retValue, err := instance.GetProperty("Type")
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

// SetUsedSlots sets the value of UsedSlots for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyUsedSlots(value uint16) (err error) {
	return instance.SetProperty("UsedSlots", (value))
}

// GetUsedSlots gets the value of UsedSlots for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyUsedSlots() (value uint16, err error) {
	retValue, err := instance.GetProperty("UsedSlots")
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

// SetUtilization sets the value of Utilization for the instance
func (instance *MSFT_MTMemorySummary) SetPropertyUtilization(value []float32) (err error) {
	return instance.SetProperty("Utilization", (value))
}

// GetUtilization gets the value of Utilization for the instance
func (instance *MSFT_MTMemorySummary) GetPropertyUtilization() (value []float32, err error) {
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
