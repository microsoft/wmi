// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DiskIo_V2_TypeGroup2 struct
type DiskIo_V2_TypeGroup2 struct {
	*DiskIo_V2

	//
	Irp uint32
}

func NewDiskIo_V2_TypeGroup2Ex1(instance *cim.WmiInstance) (newInstance *DiskIo_V2_TypeGroup2, err error) {
	tmp, err := NewDiskIo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &DiskIo_V2_TypeGroup2{
		DiskIo_V2: tmp,
	}
	return
}

func NewDiskIo_V2_TypeGroup2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DiskIo_V2_TypeGroup2, err error) {
	tmp, err := NewDiskIo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DiskIo_V2_TypeGroup2{
		DiskIo_V2: tmp,
	}
	return
}

// SetIrp sets the value of Irp for the instance
func (instance *DiskIo_V2_TypeGroup2) SetPropertyIrp(value uint32) (err error) {
	return instance.SetProperty("Irp", (value))
}

// GetIrp gets the value of Irp for the instance
func (instance *DiskIo_V2_TypeGroup2) GetPropertyIrp() (value uint32, err error) {
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
