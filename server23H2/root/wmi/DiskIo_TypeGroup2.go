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

// DiskIo_TypeGroup2 struct
type DiskIo_TypeGroup2 struct {
	*DiskIo

	//
	Irp uint32

	//
	IssuingThreadId uint32
}

func NewDiskIo_TypeGroup2Ex1(instance *cim.WmiInstance) (newInstance *DiskIo_TypeGroup2, err error) {
	tmp, err := NewDiskIoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DiskIo_TypeGroup2{
		DiskIo: tmp,
	}
	return
}

func NewDiskIo_TypeGroup2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DiskIo_TypeGroup2, err error) {
	tmp, err := NewDiskIoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DiskIo_TypeGroup2{
		DiskIo: tmp,
	}
	return
}

// SetIrp sets the value of Irp for the instance
func (instance *DiskIo_TypeGroup2) SetPropertyIrp(value uint32) (err error) {
	return instance.SetProperty("Irp", (value))
}

// GetIrp gets the value of Irp for the instance
func (instance *DiskIo_TypeGroup2) GetPropertyIrp() (value uint32, err error) {
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

// SetIssuingThreadId sets the value of IssuingThreadId for the instance
func (instance *DiskIo_TypeGroup2) SetPropertyIssuingThreadId(value uint32) (err error) {
	return instance.SetProperty("IssuingThreadId", (value))
}

// GetIssuingThreadId gets the value of IssuingThreadId for the instance
func (instance *DiskIo_TypeGroup2) GetPropertyIssuingThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("IssuingThreadId")
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
