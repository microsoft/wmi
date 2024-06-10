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

// DiskIo_V2_TypeGroup3 struct
type DiskIo_V2_TypeGroup3 struct {
	*DiskIo_V2

	//
	DiskNumber uint32

	//
	HighResResponseTime uint64

	//
	Irp uint32

	//
	IrpFlags uint32
}

func NewDiskIo_V2_TypeGroup3Ex1(instance *cim.WmiInstance) (newInstance *DiskIo_V2_TypeGroup3, err error) {
	tmp, err := NewDiskIo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &DiskIo_V2_TypeGroup3{
		DiskIo_V2: tmp,
	}
	return
}

func NewDiskIo_V2_TypeGroup3Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DiskIo_V2_TypeGroup3, err error) {
	tmp, err := NewDiskIo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DiskIo_V2_TypeGroup3{
		DiskIo_V2: tmp,
	}
	return
}

// SetDiskNumber sets the value of DiskNumber for the instance
func (instance *DiskIo_V2_TypeGroup3) SetPropertyDiskNumber(value uint32) (err error) {
	return instance.SetProperty("DiskNumber", (value))
}

// GetDiskNumber gets the value of DiskNumber for the instance
func (instance *DiskIo_V2_TypeGroup3) GetPropertyDiskNumber() (value uint32, err error) {
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

// SetHighResResponseTime sets the value of HighResResponseTime for the instance
func (instance *DiskIo_V2_TypeGroup3) SetPropertyHighResResponseTime(value uint64) (err error) {
	return instance.SetProperty("HighResResponseTime", (value))
}

// GetHighResResponseTime gets the value of HighResResponseTime for the instance
func (instance *DiskIo_V2_TypeGroup3) GetPropertyHighResResponseTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("HighResResponseTime")
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

// SetIrp sets the value of Irp for the instance
func (instance *DiskIo_V2_TypeGroup3) SetPropertyIrp(value uint32) (err error) {
	return instance.SetProperty("Irp", (value))
}

// GetIrp gets the value of Irp for the instance
func (instance *DiskIo_V2_TypeGroup3) GetPropertyIrp() (value uint32, err error) {
	retValue, err := instance.GetProperty("Irp")
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
func (instance *DiskIo_V2_TypeGroup3) SetPropertyIrpFlags(value uint32) (err error) {
	return instance.SetProperty("IrpFlags", (value))
}

// GetIrpFlags gets the value of IrpFlags for the instance
func (instance *DiskIo_V2_TypeGroup3) GetPropertyIrpFlags() (value uint32, err error) {
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
