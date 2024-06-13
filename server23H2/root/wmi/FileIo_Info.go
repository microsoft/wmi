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

// FileIo_Info struct
type FileIo_Info struct {
	*FileIo

	//
	ExtraInfo uint32

	//
	FileKey uint32

	//
	FileObject uint32

	//
	InfoClass uint32

	//
	IrpPtr uint32

	//
	TTID uint32
}

func NewFileIo_InfoEx1(instance *cim.WmiInstance) (newInstance *FileIo_Info, err error) {
	tmp, err := NewFileIoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FileIo_Info{
		FileIo: tmp,
	}
	return
}

func NewFileIo_InfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FileIo_Info, err error) {
	tmp, err := NewFileIoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FileIo_Info{
		FileIo: tmp,
	}
	return
}

// SetExtraInfo sets the value of ExtraInfo for the instance
func (instance *FileIo_Info) SetPropertyExtraInfo(value uint32) (err error) {
	return instance.SetProperty("ExtraInfo", (value))
}

// GetExtraInfo gets the value of ExtraInfo for the instance
func (instance *FileIo_Info) GetPropertyExtraInfo() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExtraInfo")
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

// SetFileKey sets the value of FileKey for the instance
func (instance *FileIo_Info) SetPropertyFileKey(value uint32) (err error) {
	return instance.SetProperty("FileKey", (value))
}

// GetFileKey gets the value of FileKey for the instance
func (instance *FileIo_Info) GetPropertyFileKey() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileKey")
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
func (instance *FileIo_Info) SetPropertyFileObject(value uint32) (err error) {
	return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *FileIo_Info) GetPropertyFileObject() (value uint32, err error) {
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

// SetInfoClass sets the value of InfoClass for the instance
func (instance *FileIo_Info) SetPropertyInfoClass(value uint32) (err error) {
	return instance.SetProperty("InfoClass", (value))
}

// GetInfoClass gets the value of InfoClass for the instance
func (instance *FileIo_Info) GetPropertyInfoClass() (value uint32, err error) {
	retValue, err := instance.GetProperty("InfoClass")
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

// SetIrpPtr sets the value of IrpPtr for the instance
func (instance *FileIo_Info) SetPropertyIrpPtr(value uint32) (err error) {
	return instance.SetProperty("IrpPtr", (value))
}

// GetIrpPtr gets the value of IrpPtr for the instance
func (instance *FileIo_Info) GetPropertyIrpPtr() (value uint32, err error) {
	retValue, err := instance.GetProperty("IrpPtr")
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

// SetTTID sets the value of TTID for the instance
func (instance *FileIo_Info) SetPropertyTTID(value uint32) (err error) {
	return instance.SetProperty("TTID", (value))
}

// GetTTID gets the value of TTID for the instance
func (instance *FileIo_Info) GetPropertyTTID() (value uint32, err error) {
	retValue, err := instance.GetProperty("TTID")
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
