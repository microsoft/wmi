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

// Image_Load_V2 struct
type Image_Load_V2 struct {
	*Image_V2

	//
	DefaultBase uint32

	//
	FileName string

	//
	ImageBase uint32

	//
	ImageChecksum uint32

	//
	ImageSize uint32

	//
	ProcessId uint32

	//
	Reserved0 uint32

	//
	Reserved1 uint32

	//
	Reserved2 uint32

	//
	Reserved3 uint32

	//
	Reserved4 uint32

	//
	TimeDateStamp uint32
}

func NewImage_Load_V2Ex1(instance *cim.WmiInstance) (newInstance *Image_Load_V2, err error) {
	tmp, err := NewImage_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &Image_Load_V2{
		Image_V2: tmp,
	}
	return
}

func NewImage_Load_V2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Image_Load_V2, err error) {
	tmp, err := NewImage_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Image_Load_V2{
		Image_V2: tmp,
	}
	return
}

// SetDefaultBase sets the value of DefaultBase for the instance
func (instance *Image_Load_V2) SetPropertyDefaultBase(value uint32) (err error) {
	return instance.SetProperty("DefaultBase", (value))
}

// GetDefaultBase gets the value of DefaultBase for the instance
func (instance *Image_Load_V2) GetPropertyDefaultBase() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultBase")
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

// SetFileName sets the value of FileName for the instance
func (instance *Image_Load_V2) SetPropertyFileName(value string) (err error) {
	return instance.SetProperty("FileName", (value))
}

// GetFileName gets the value of FileName for the instance
func (instance *Image_Load_V2) GetPropertyFileName() (value string, err error) {
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
func (instance *Image_Load_V2) SetPropertyImageBase(value uint32) (err error) {
	return instance.SetProperty("ImageBase", (value))
}

// GetImageBase gets the value of ImageBase for the instance
func (instance *Image_Load_V2) GetPropertyImageBase() (value uint32, err error) {
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

// SetImageChecksum sets the value of ImageChecksum for the instance
func (instance *Image_Load_V2) SetPropertyImageChecksum(value uint32) (err error) {
	return instance.SetProperty("ImageChecksum", (value))
}

// GetImageChecksum gets the value of ImageChecksum for the instance
func (instance *Image_Load_V2) GetPropertyImageChecksum() (value uint32, err error) {
	retValue, err := instance.GetProperty("ImageChecksum")
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
func (instance *Image_Load_V2) SetPropertyImageSize(value uint32) (err error) {
	return instance.SetProperty("ImageSize", (value))
}

// GetImageSize gets the value of ImageSize for the instance
func (instance *Image_Load_V2) GetPropertyImageSize() (value uint32, err error) {
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
func (instance *Image_Load_V2) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Image_Load_V2) GetPropertyProcessId() (value uint32, err error) {
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

// SetReserved0 sets the value of Reserved0 for the instance
func (instance *Image_Load_V2) SetPropertyReserved0(value uint32) (err error) {
	return instance.SetProperty("Reserved0", (value))
}

// GetReserved0 gets the value of Reserved0 for the instance
func (instance *Image_Load_V2) GetPropertyReserved0() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved0")
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

// SetReserved1 sets the value of Reserved1 for the instance
func (instance *Image_Load_V2) SetPropertyReserved1(value uint32) (err error) {
	return instance.SetProperty("Reserved1", (value))
}

// GetReserved1 gets the value of Reserved1 for the instance
func (instance *Image_Load_V2) GetPropertyReserved1() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved1")
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

// SetReserved2 sets the value of Reserved2 for the instance
func (instance *Image_Load_V2) SetPropertyReserved2(value uint32) (err error) {
	return instance.SetProperty("Reserved2", (value))
}

// GetReserved2 gets the value of Reserved2 for the instance
func (instance *Image_Load_V2) GetPropertyReserved2() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved2")
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

// SetReserved3 sets the value of Reserved3 for the instance
func (instance *Image_Load_V2) SetPropertyReserved3(value uint32) (err error) {
	return instance.SetProperty("Reserved3", (value))
}

// GetReserved3 gets the value of Reserved3 for the instance
func (instance *Image_Load_V2) GetPropertyReserved3() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved3")
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

// SetReserved4 sets the value of Reserved4 for the instance
func (instance *Image_Load_V2) SetPropertyReserved4(value uint32) (err error) {
	return instance.SetProperty("Reserved4", (value))
}

// GetReserved4 gets the value of Reserved4 for the instance
func (instance *Image_Load_V2) GetPropertyReserved4() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved4")
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

// SetTimeDateStamp sets the value of TimeDateStamp for the instance
func (instance *Image_Load_V2) SetPropertyTimeDateStamp(value uint32) (err error) {
	return instance.SetProperty("TimeDateStamp", (value))
}

// GetTimeDateStamp gets the value of TimeDateStamp for the instance
func (instance *Image_Load_V2) GetPropertyTimeDateStamp() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimeDateStamp")
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
