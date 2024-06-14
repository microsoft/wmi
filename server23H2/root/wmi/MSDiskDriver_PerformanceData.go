// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSDiskDriver_PerformanceData struct
type MSDiskDriver_PerformanceData struct {
	*MSDiskDriver

	//
	BytesRead int64

	//
	BytesWritten int64

	//
	IdleTime int64

	//
	QueryTime int64

	//
	QueueDepth uint32

	//
	ReadCount uint32

	//
	ReadTime int64

	//
	SplitCount uint32

	//
	StorageDeviceNumber uint32

	//
	StorageManagerName []uint16

	//
	WriteCount uint32

	//
	WriteTime int64
}

func NewMSDiskDriver_PerformanceDataEx1(instance *cim.WmiInstance) (newInstance *MSDiskDriver_PerformanceData, err error) {
	tmp, err := NewMSDiskDriverEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSDiskDriver_PerformanceData{
		MSDiskDriver: tmp,
	}
	return
}

func NewMSDiskDriver_PerformanceDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSDiskDriver_PerformanceData, err error) {
	tmp, err := NewMSDiskDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSDiskDriver_PerformanceData{
		MSDiskDriver: tmp,
	}
	return
}

// SetBytesRead sets the value of BytesRead for the instance
func (instance *MSDiskDriver_PerformanceData) SetPropertyBytesRead(value int64) (err error) {
	return instance.SetProperty("BytesRead", (value))
}

// GetBytesRead gets the value of BytesRead for the instance
func (instance *MSDiskDriver_PerformanceData) GetPropertyBytesRead() (value int64, err error) {
	retValue, err := instance.GetProperty("BytesRead")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetBytesWritten sets the value of BytesWritten for the instance
func (instance *MSDiskDriver_PerformanceData) SetPropertyBytesWritten(value int64) (err error) {
	return instance.SetProperty("BytesWritten", (value))
}

// GetBytesWritten gets the value of BytesWritten for the instance
func (instance *MSDiskDriver_PerformanceData) GetPropertyBytesWritten() (value int64, err error) {
	retValue, err := instance.GetProperty("BytesWritten")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetIdleTime sets the value of IdleTime for the instance
func (instance *MSDiskDriver_PerformanceData) SetPropertyIdleTime(value int64) (err error) {
	return instance.SetProperty("IdleTime", (value))
}

// GetIdleTime gets the value of IdleTime for the instance
func (instance *MSDiskDriver_PerformanceData) GetPropertyIdleTime() (value int64, err error) {
	retValue, err := instance.GetProperty("IdleTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetQueryTime sets the value of QueryTime for the instance
func (instance *MSDiskDriver_PerformanceData) SetPropertyQueryTime(value int64) (err error) {
	return instance.SetProperty("QueryTime", (value))
}

// GetQueryTime gets the value of QueryTime for the instance
func (instance *MSDiskDriver_PerformanceData) GetPropertyQueryTime() (value int64, err error) {
	retValue, err := instance.GetProperty("QueryTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetQueueDepth sets the value of QueueDepth for the instance
func (instance *MSDiskDriver_PerformanceData) SetPropertyQueueDepth(value uint32) (err error) {
	return instance.SetProperty("QueueDepth", (value))
}

// GetQueueDepth gets the value of QueueDepth for the instance
func (instance *MSDiskDriver_PerformanceData) GetPropertyQueueDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueueDepth")
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

// SetReadCount sets the value of ReadCount for the instance
func (instance *MSDiskDriver_PerformanceData) SetPropertyReadCount(value uint32) (err error) {
	return instance.SetProperty("ReadCount", (value))
}

// GetReadCount gets the value of ReadCount for the instance
func (instance *MSDiskDriver_PerformanceData) GetPropertyReadCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReadCount")
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

// SetReadTime sets the value of ReadTime for the instance
func (instance *MSDiskDriver_PerformanceData) SetPropertyReadTime(value int64) (err error) {
	return instance.SetProperty("ReadTime", (value))
}

// GetReadTime gets the value of ReadTime for the instance
func (instance *MSDiskDriver_PerformanceData) GetPropertyReadTime() (value int64, err error) {
	retValue, err := instance.GetProperty("ReadTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetSplitCount sets the value of SplitCount for the instance
func (instance *MSDiskDriver_PerformanceData) SetPropertySplitCount(value uint32) (err error) {
	return instance.SetProperty("SplitCount", (value))
}

// GetSplitCount gets the value of SplitCount for the instance
func (instance *MSDiskDriver_PerformanceData) GetPropertySplitCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("SplitCount")
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

// SetStorageDeviceNumber sets the value of StorageDeviceNumber for the instance
func (instance *MSDiskDriver_PerformanceData) SetPropertyStorageDeviceNumber(value uint32) (err error) {
	return instance.SetProperty("StorageDeviceNumber", (value))
}

// GetStorageDeviceNumber gets the value of StorageDeviceNumber for the instance
func (instance *MSDiskDriver_PerformanceData) GetPropertyStorageDeviceNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("StorageDeviceNumber")
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

// SetStorageManagerName sets the value of StorageManagerName for the instance
func (instance *MSDiskDriver_PerformanceData) SetPropertyStorageManagerName(value []uint16) (err error) {
	return instance.SetProperty("StorageManagerName", (value))
}

// GetStorageManagerName gets the value of StorageManagerName for the instance
func (instance *MSDiskDriver_PerformanceData) GetPropertyStorageManagerName() (value []uint16, err error) {
	retValue, err := instance.GetProperty("StorageManagerName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetWriteCount sets the value of WriteCount for the instance
func (instance *MSDiskDriver_PerformanceData) SetPropertyWriteCount(value uint32) (err error) {
	return instance.SetProperty("WriteCount", (value))
}

// GetWriteCount gets the value of WriteCount for the instance
func (instance *MSDiskDriver_PerformanceData) GetPropertyWriteCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("WriteCount")
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

// SetWriteTime sets the value of WriteTime for the instance
func (instance *MSDiskDriver_PerformanceData) SetPropertyWriteTime(value int64) (err error) {
	return instance.SetProperty("WriteTime", (value))
}

// GetWriteTime gets the value of WriteTime for the instance
func (instance *MSDiskDriver_PerformanceData) GetPropertyWriteTime() (value int64, err error) {
	retValue, err := instance.GetProperty("WriteTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}
