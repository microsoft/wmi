// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSParallel_DeviceBytesTransferred struct
type MSParallel_DeviceBytesTransferred struct {
	*MSParallel

	//
	Active bool

	//
	BoundedEcpReadCount int64

	//
	BoundedEcpWriteCount int64

	//
	ByteReadCount int64

	//
	ChannelNibbleReadCount int64

	//
	Flags1 uint32

	//
	Flags2 uint32

	//
	HwEcpReadCount int64

	//
	HwEcpWriteCount int64

	//
	HwEppReadCount int64

	//
	HwEppWriteCount int64

	//
	InstanceName string

	//
	NibbleReadCount int64

	//
	spare []uint32

	//
	SppWriteCount int64

	//
	SwEcpReadCount int64

	//
	SwEcpWriteCount int64

	//
	SwEppReadCount int64

	//
	SwEppWriteCount int64
}

func NewMSParallel_DeviceBytesTransferredEx1(instance *cim.WmiInstance) (newInstance *MSParallel_DeviceBytesTransferred, err error) {
	tmp, err := NewMSParallelEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSParallel_DeviceBytesTransferred{
		MSParallel: tmp,
	}
	return
}

func NewMSParallel_DeviceBytesTransferredEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSParallel_DeviceBytesTransferred, err error) {
	tmp, err := NewMSParallelEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSParallel_DeviceBytesTransferred{
		MSParallel: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetBoundedEcpReadCount sets the value of BoundedEcpReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyBoundedEcpReadCount(value int64) (err error) {
	return instance.SetProperty("BoundedEcpReadCount", (value))
}

// GetBoundedEcpReadCount gets the value of BoundedEcpReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyBoundedEcpReadCount() (value int64, err error) {
	retValue, err := instance.GetProperty("BoundedEcpReadCount")
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

// SetBoundedEcpWriteCount sets the value of BoundedEcpWriteCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyBoundedEcpWriteCount(value int64) (err error) {
	return instance.SetProperty("BoundedEcpWriteCount", (value))
}

// GetBoundedEcpWriteCount gets the value of BoundedEcpWriteCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyBoundedEcpWriteCount() (value int64, err error) {
	retValue, err := instance.GetProperty("BoundedEcpWriteCount")
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

// SetByteReadCount sets the value of ByteReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyByteReadCount(value int64) (err error) {
	return instance.SetProperty("ByteReadCount", (value))
}

// GetByteReadCount gets the value of ByteReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyByteReadCount() (value int64, err error) {
	retValue, err := instance.GetProperty("ByteReadCount")
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

// SetChannelNibbleReadCount sets the value of ChannelNibbleReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyChannelNibbleReadCount(value int64) (err error) {
	return instance.SetProperty("ChannelNibbleReadCount", (value))
}

// GetChannelNibbleReadCount gets the value of ChannelNibbleReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyChannelNibbleReadCount() (value int64, err error) {
	retValue, err := instance.GetProperty("ChannelNibbleReadCount")
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

// SetFlags1 sets the value of Flags1 for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyFlags1(value uint32) (err error) {
	return instance.SetProperty("Flags1", (value))
}

// GetFlags1 gets the value of Flags1 for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyFlags1() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags1")
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

// SetFlags2 sets the value of Flags2 for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyFlags2(value uint32) (err error) {
	return instance.SetProperty("Flags2", (value))
}

// GetFlags2 gets the value of Flags2 for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyFlags2() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags2")
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

// SetHwEcpReadCount sets the value of HwEcpReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyHwEcpReadCount(value int64) (err error) {
	return instance.SetProperty("HwEcpReadCount", (value))
}

// GetHwEcpReadCount gets the value of HwEcpReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyHwEcpReadCount() (value int64, err error) {
	retValue, err := instance.GetProperty("HwEcpReadCount")
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

// SetHwEcpWriteCount sets the value of HwEcpWriteCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyHwEcpWriteCount(value int64) (err error) {
	return instance.SetProperty("HwEcpWriteCount", (value))
}

// GetHwEcpWriteCount gets the value of HwEcpWriteCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyHwEcpWriteCount() (value int64, err error) {
	retValue, err := instance.GetProperty("HwEcpWriteCount")
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

// SetHwEppReadCount sets the value of HwEppReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyHwEppReadCount(value int64) (err error) {
	return instance.SetProperty("HwEppReadCount", (value))
}

// GetHwEppReadCount gets the value of HwEppReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyHwEppReadCount() (value int64, err error) {
	retValue, err := instance.GetProperty("HwEppReadCount")
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

// SetHwEppWriteCount sets the value of HwEppWriteCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyHwEppWriteCount(value int64) (err error) {
	return instance.SetProperty("HwEppWriteCount", (value))
}

// GetHwEppWriteCount gets the value of HwEppWriteCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyHwEppWriteCount() (value int64, err error) {
	retValue, err := instance.GetProperty("HwEppWriteCount")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetNibbleReadCount sets the value of NibbleReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyNibbleReadCount(value int64) (err error) {
	return instance.SetProperty("NibbleReadCount", (value))
}

// GetNibbleReadCount gets the value of NibbleReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyNibbleReadCount() (value int64, err error) {
	retValue, err := instance.GetProperty("NibbleReadCount")
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

// Setspare sets the value of spare for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertyspare(value []uint32) (err error) {
	return instance.SetProperty("spare", (value))
}

// Getspare gets the value of spare for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertyspare() (value []uint32, err error) {
	retValue, err := instance.GetProperty("spare")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetSppWriteCount sets the value of SppWriteCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertySppWriteCount(value int64) (err error) {
	return instance.SetProperty("SppWriteCount", (value))
}

// GetSppWriteCount gets the value of SppWriteCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertySppWriteCount() (value int64, err error) {
	retValue, err := instance.GetProperty("SppWriteCount")
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

// SetSwEcpReadCount sets the value of SwEcpReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertySwEcpReadCount(value int64) (err error) {
	return instance.SetProperty("SwEcpReadCount", (value))
}

// GetSwEcpReadCount gets the value of SwEcpReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertySwEcpReadCount() (value int64, err error) {
	retValue, err := instance.GetProperty("SwEcpReadCount")
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

// SetSwEcpWriteCount sets the value of SwEcpWriteCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertySwEcpWriteCount(value int64) (err error) {
	return instance.SetProperty("SwEcpWriteCount", (value))
}

// GetSwEcpWriteCount gets the value of SwEcpWriteCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertySwEcpWriteCount() (value int64, err error) {
	retValue, err := instance.GetProperty("SwEcpWriteCount")
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

// SetSwEppReadCount sets the value of SwEppReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertySwEppReadCount(value int64) (err error) {
	return instance.SetProperty("SwEppReadCount", (value))
}

// GetSwEppReadCount gets the value of SwEppReadCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertySwEppReadCount() (value int64, err error) {
	retValue, err := instance.GetProperty("SwEppReadCount")
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

// SetSwEppWriteCount sets the value of SwEppWriteCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) SetPropertySwEppWriteCount(value int64) (err error) {
	return instance.SetProperty("SwEppWriteCount", (value))
}

// GetSwEppWriteCount gets the value of SwEppWriteCount for the instance
func (instance *MSParallel_DeviceBytesTransferred) GetPropertySwEppWriteCount() (value int64, err error) {
	retValue, err := instance.GetProperty("SwEppWriteCount")
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
