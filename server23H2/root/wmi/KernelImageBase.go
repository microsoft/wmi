// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// KernelImageBase struct
type KernelImageBase struct {
	*Image_V2

	//
	ImageBase uint32
}

func NewKernelImageBaseEx1(instance *cim.WmiInstance) (newInstance *KernelImageBase, err error) {
	tmp, err := NewImage_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &KernelImageBase{
		Image_V2: tmp,
	}
	return
}

func NewKernelImageBaseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KernelImageBase, err error) {
	tmp, err := NewImage_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KernelImageBase{
		Image_V2: tmp,
	}
	return
}

// SetImageBase sets the value of ImageBase for the instance
func (instance *KernelImageBase) SetPropertyImageBase(value uint32) (err error) {
	return instance.SetProperty("ImageBase", (value))
}

// GetImageBase gets the value of ImageBase for the instance
func (instance *KernelImageBase) GetPropertyImageBase() (value uint32, err error) {
	retValue, err := instance.GetProperty("ImageBase")
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
