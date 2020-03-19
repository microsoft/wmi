// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MTDisk struct
type MSFT_MTDisk struct {
	*CIM_ManagedElement

	//
	ActiveTime []float32

	//
	AverageResponseTime float32

	//
	Capacity uint64

	//
	CurrentIndex uint16

	//
	DiskNumber uint32

	//
	IntervalSeconds uint16

	//
	Name string

	//
	ReadTransferRate []float32

	//
	Volumes []MSFT_MTDiskVolume

	//
	WriteTransferRate []float32
}

func NewMSFT_MTDiskEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTDisk, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTDisk{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_MTDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MTDisk, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTDisk{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetActiveTime sets the value of ActiveTime for the instance
func (instance *MSFT_MTDisk) SetPropertyActiveTime(value []float32) (err error) {
	return instance.SetProperty("ActiveTime", value)
}

// GetActiveTime gets the value of ActiveTime for the instance
func (instance *MSFT_MTDisk) GetPropertyActiveTime() (value []float32, err error) {
	retValue, err := instance.GetProperty("ActiveTime")
	if err != nil {
		return
	}
	value, ok := retValue.([]float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageResponseTime sets the value of AverageResponseTime for the instance
func (instance *MSFT_MTDisk) SetPropertyAverageResponseTime(value float32) (err error) {
	return instance.SetProperty("AverageResponseTime", value)
}

// GetAverageResponseTime gets the value of AverageResponseTime for the instance
func (instance *MSFT_MTDisk) GetPropertyAverageResponseTime() (value float32, err error) {
	retValue, err := instance.GetProperty("AverageResponseTime")
	if err != nil {
		return
	}
	value, ok := retValue.(float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCapacity sets the value of Capacity for the instance
func (instance *MSFT_MTDisk) SetPropertyCapacity(value uint64) (err error) {
	return instance.SetProperty("Capacity", value)
}

// GetCapacity gets the value of Capacity for the instance
func (instance *MSFT_MTDisk) GetPropertyCapacity() (value uint64, err error) {
	retValue, err := instance.GetProperty("Capacity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentIndex sets the value of CurrentIndex for the instance
func (instance *MSFT_MTDisk) SetPropertyCurrentIndex(value uint16) (err error) {
	return instance.SetProperty("CurrentIndex", value)
}

// GetCurrentIndex gets the value of CurrentIndex for the instance
func (instance *MSFT_MTDisk) GetPropertyCurrentIndex() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiskNumber sets the value of DiskNumber for the instance
func (instance *MSFT_MTDisk) SetPropertyDiskNumber(value uint32) (err error) {
	return instance.SetProperty("DiskNumber", value)
}

// GetDiskNumber gets the value of DiskNumber for the instance
func (instance *MSFT_MTDisk) GetPropertyDiskNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DiskNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntervalSeconds sets the value of IntervalSeconds for the instance
func (instance *MSFT_MTDisk) SetPropertyIntervalSeconds(value uint16) (err error) {
	return instance.SetProperty("IntervalSeconds", value)
}

// GetIntervalSeconds gets the value of IntervalSeconds for the instance
func (instance *MSFT_MTDisk) GetPropertyIntervalSeconds() (value uint16, err error) {
	retValue, err := instance.GetProperty("IntervalSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_MTDisk) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_MTDisk) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadTransferRate sets the value of ReadTransferRate for the instance
func (instance *MSFT_MTDisk) SetPropertyReadTransferRate(value []float32) (err error) {
	return instance.SetProperty("ReadTransferRate", value)
}

// GetReadTransferRate gets the value of ReadTransferRate for the instance
func (instance *MSFT_MTDisk) GetPropertyReadTransferRate() (value []float32, err error) {
	retValue, err := instance.GetProperty("ReadTransferRate")
	if err != nil {
		return
	}
	value, ok := retValue.([]float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolumes sets the value of Volumes for the instance
func (instance *MSFT_MTDisk) SetPropertyVolumes(value []MSFT_MTDiskVolume) (err error) {
	return instance.SetProperty("Volumes", value)
}

// GetVolumes gets the value of Volumes for the instance
func (instance *MSFT_MTDisk) GetPropertyVolumes() (value []MSFT_MTDiskVolume, err error) {
	retValue, err := instance.GetProperty("Volumes")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_MTDiskVolume)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteTransferRate sets the value of WriteTransferRate for the instance
func (instance *MSFT_MTDisk) SetPropertyWriteTransferRate(value []float32) (err error) {
	return instance.SetProperty("WriteTransferRate", value)
}

// GetWriteTransferRate gets the value of WriteTransferRate for the instance
func (instance *MSFT_MTDisk) GetPropertyWriteTransferRate() (value []float32, err error) {
	retValue, err := instance.GetProperty("WriteTransferRate")
	if err != nil {
		return
	}
	value, ok := retValue.([]float32)
	if !ok {
		// TODO: Set an error
	}
	return
}
