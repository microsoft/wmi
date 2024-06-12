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

// Image_Load struct
type Image_Load struct {
	*Image

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
	Reserved0 uint16

	//
	Reserved1 uint32

	//
	Reserved2 uint32

	//
	Reserved3 uint32

	//
	Reserved4 uint32

	//
	SignatureLevel uint8

	//
	SignatureType uint8

	//
	TimeDateStamp uint32
}

func NewImage_LoadEx1(instance *cim.WmiInstance) (newInstance *Image_Load, err error) {
	tmp, err := NewImageEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Image_Load{
		Image: tmp,
	}
	return
}

func NewImage_LoadEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Image_Load, err error) {
	tmp, err := NewImageEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Image_Load{
		Image: tmp,
	}
	return
}

// SetDefaultBase sets the value of DefaultBase for the instance
func (instance *Image_Load) SetPropertyDefaultBase(value uint32) (err error) {
	return instance.SetProperty("DefaultBase", (value))
}

// GetDefaultBase gets the value of DefaultBase for the instance
func (instance *Image_Load) GetPropertyDefaultBase() (value uint32, err error) {
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
func (instance *Image_Load) SetPropertyFileName(value string) (err error) {
	return instance.SetProperty("FileName", (value))
}

// GetFileName gets the value of FileName for the instance
func (instance *Image_Load) GetPropertyFileName() (value string, err error) {
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
func (instance *Image_Load) SetPropertyImageBase(value uint32) (err error) {
	return instance.SetProperty("ImageBase", (value))
}

// GetImageBase gets the value of ImageBase for the instance
func (instance *Image_Load) GetPropertyImageBase() (value uint32, err error) {
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
func (instance *Image_Load) SetPropertyImageChecksum(value uint32) (err error) {
	return instance.SetProperty("ImageChecksum", (value))
}

// GetImageChecksum gets the value of ImageChecksum for the instance
func (instance *Image_Load) GetPropertyImageChecksum() (value uint32, err error) {
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
func (instance *Image_Load) SetPropertyImageSize(value uint32) (err error) {
	return instance.SetProperty("ImageSize", (value))
}

// GetImageSize gets the value of ImageSize for the instance
func (instance *Image_Load) GetPropertyImageSize() (value uint32, err error) {
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
func (instance *Image_Load) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Image_Load) GetPropertyProcessId() (value uint32, err error) {
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
func (instance *Image_Load) SetPropertyReserved0(value uint16) (err error) {
	return instance.SetProperty("Reserved0", (value))
}

// GetReserved0 gets the value of Reserved0 for the instance
func (instance *Image_Load) GetPropertyReserved0() (value uint16, err error) {
	retValue, err := instance.GetProperty("Reserved0")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetReserved1 sets the value of Reserved1 for the instance
func (instance *Image_Load) SetPropertyReserved1(value uint32) (err error) {
	return instance.SetProperty("Reserved1", (value))
}

// GetReserved1 gets the value of Reserved1 for the instance
func (instance *Image_Load) GetPropertyReserved1() (value uint32, err error) {
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
func (instance *Image_Load) SetPropertyReserved2(value uint32) (err error) {
	return instance.SetProperty("Reserved2", (value))
}

// GetReserved2 gets the value of Reserved2 for the instance
func (instance *Image_Load) GetPropertyReserved2() (value uint32, err error) {
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
func (instance *Image_Load) SetPropertyReserved3(value uint32) (err error) {
	return instance.SetProperty("Reserved3", (value))
}

// GetReserved3 gets the value of Reserved3 for the instance
func (instance *Image_Load) GetPropertyReserved3() (value uint32, err error) {
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
func (instance *Image_Load) SetPropertyReserved4(value uint32) (err error) {
	return instance.SetProperty("Reserved4", (value))
}

// GetReserved4 gets the value of Reserved4 for the instance
func (instance *Image_Load) GetPropertyReserved4() (value uint32, err error) {
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

// SetSignatureLevel sets the value of SignatureLevel for the instance
func (instance *Image_Load) SetPropertySignatureLevel(value uint8) (err error) {
	return instance.SetProperty("SignatureLevel", (value))
}

// GetSignatureLevel gets the value of SignatureLevel for the instance
func (instance *Image_Load) GetPropertySignatureLevel() (value uint8, err error) {
	retValue, err := instance.GetProperty("SignatureLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSignatureType sets the value of SignatureType for the instance
func (instance *Image_Load) SetPropertySignatureType(value uint8) (err error) {
	return instance.SetProperty("SignatureType", (value))
}

// GetSignatureType gets the value of SignatureType for the instance
func (instance *Image_Load) GetPropertySignatureType() (value uint8, err error) {
	retValue, err := instance.GetProperty("SignatureType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetTimeDateStamp sets the value of TimeDateStamp for the instance
func (instance *Image_Load) SetPropertyTimeDateStamp(value uint32) (err error) {
	return instance.SetProperty("TimeDateStamp", (value))
}

// GetTimeDateStamp gets the value of TimeDateStamp for the instance
func (instance *Image_Load) GetPropertyTimeDateStamp() (value uint32, err error) {
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
