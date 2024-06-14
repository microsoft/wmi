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

// Image_V1_Load struct
type Image_V1_Load struct {
	*Image_V1

	//
	FileName string

	//
	ImageBase uint32

	//
	ImageSize uint32

	//
	ProcessId uint32
}

func NewImage_V1_LoadEx1(instance *cim.WmiInstance) (newInstance *Image_V1_Load, err error) {
	tmp, err := NewImage_V1Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &Image_V1_Load{
		Image_V1: tmp,
	}
	return
}

func NewImage_V1_LoadEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Image_V1_Load, err error) {
	tmp, err := NewImage_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Image_V1_Load{
		Image_V1: tmp,
	}
	return
}

// SetFileName sets the value of FileName for the instance
func (instance *Image_V1_Load) SetPropertyFileName(value string) (err error) {
	return instance.SetProperty("FileName", (value))
}

// GetFileName gets the value of FileName for the instance
func (instance *Image_V1_Load) GetPropertyFileName() (value string, err error) {
	retValue, err := instance.GetProperty("FileName")
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

// SetImageBase sets the value of ImageBase for the instance
func (instance *Image_V1_Load) SetPropertyImageBase(value uint32) (err error) {
	return instance.SetProperty("ImageBase", (value))
}

// GetImageBase gets the value of ImageBase for the instance
func (instance *Image_V1_Load) GetPropertyImageBase() (value uint32, err error) {
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

// SetImageSize sets the value of ImageSize for the instance
func (instance *Image_V1_Load) SetPropertyImageSize(value uint32) (err error) {
	return instance.SetProperty("ImageSize", (value))
}

// GetImageSize gets the value of ImageSize for the instance
func (instance *Image_V1_Load) GetPropertyImageSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("ImageSize")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *Image_V1_Load) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Image_V1_Load) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
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
