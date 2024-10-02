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

// DiskIo_V0_TypeGroup1 struct
type DiskIo_V0_TypeGroup1 struct {
	*DiskIo_V0

	//
	ByteOffset uint64

	//
	DiskNumber uint32

	//
	FileObject uint32

	//
	IrpFlags uint32

	//
	Reserved uint32

	//
	TransferSize uint32
}

func NewDiskIo_V0_TypeGroup1Ex1(instance *cim.WmiInstance) (newInstance *DiskIo_V0_TypeGroup1, err error) {
	tmp, err := NewDiskIo_V0Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &DiskIo_V0_TypeGroup1{
		DiskIo_V0: tmp,
	}
	return
}

func NewDiskIo_V0_TypeGroup1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DiskIo_V0_TypeGroup1, err error) {
	tmp, err := NewDiskIo_V0Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DiskIo_V0_TypeGroup1{
		DiskIo_V0: tmp,
	}
	return
}

// SetByteOffset sets the value of ByteOffset for the instance
func (instance *DiskIo_V0_TypeGroup1) SetPropertyByteOffset(value uint64) (err error) {
	return instance.SetProperty("ByteOffset", (value))
}

// GetByteOffset gets the value of ByteOffset for the instance
func (instance *DiskIo_V0_TypeGroup1) GetPropertyByteOffset() (value uint64, err error) {
	retValue, err := instance.GetProperty("ByteOffset")
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

// SetDiskNumber sets the value of DiskNumber for the instance
func (instance *DiskIo_V0_TypeGroup1) SetPropertyDiskNumber(value uint32) (err error) {
	return instance.SetProperty("DiskNumber", (value))
}

// GetDiskNumber gets the value of DiskNumber for the instance
func (instance *DiskIo_V0_TypeGroup1) GetPropertyDiskNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DiskNumber")
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

// SetFileObject sets the value of FileObject for the instance
func (instance *DiskIo_V0_TypeGroup1) SetPropertyFileObject(value uint32) (err error) {
	return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *DiskIo_V0_TypeGroup1) GetPropertyFileObject() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileObject")
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

// SetIrpFlags sets the value of IrpFlags for the instance
func (instance *DiskIo_V0_TypeGroup1) SetPropertyIrpFlags(value uint32) (err error) {
	return instance.SetProperty("IrpFlags", (value))
}

// GetIrpFlags gets the value of IrpFlags for the instance
func (instance *DiskIo_V0_TypeGroup1) GetPropertyIrpFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("IrpFlags")
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

// SetReserved sets the value of Reserved for the instance
func (instance *DiskIo_V0_TypeGroup1) SetPropertyReserved(value uint32) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *DiskIo_V0_TypeGroup1) GetPropertyReserved() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved")
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

// SetTransferSize sets the value of TransferSize for the instance
func (instance *DiskIo_V0_TypeGroup1) SetPropertyTransferSize(value uint32) (err error) {
	return instance.SetProperty("TransferSize", (value))
}

// GetTransferSize gets the value of TransferSize for the instance
func (instance *DiskIo_V0_TypeGroup1) GetPropertyTransferSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("TransferSize")
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
