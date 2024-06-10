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

// FileIo_PathOperation struct
type FileIo_PathOperation struct {
	*FileIo

	//
	ExtraInfo uint32

	//
	FileKey uint32

	//
	FileName string

	//
	FileObject uint32

	//
	InfoClass uint32

	//
	IrpPtr uint32

	//
	TTID uint32
}

func NewFileIo_PathOperationEx1(instance *cim.WmiInstance) (newInstance *FileIo_PathOperation, err error) {
	tmp, err := NewFileIoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FileIo_PathOperation{
		FileIo: tmp,
	}
	return
}

func NewFileIo_PathOperationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FileIo_PathOperation, err error) {
	tmp, err := NewFileIoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FileIo_PathOperation{
		FileIo: tmp,
	}
	return
}

// SetExtraInfo sets the value of ExtraInfo for the instance
func (instance *FileIo_PathOperation) SetPropertyExtraInfo(value uint32) (err error) {
	return instance.SetProperty("ExtraInfo", (value))
}

// GetExtraInfo gets the value of ExtraInfo for the instance
func (instance *FileIo_PathOperation) GetPropertyExtraInfo() (value uint32, err error) {
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
func (instance *FileIo_PathOperation) SetPropertyFileKey(value uint32) (err error) {
	return instance.SetProperty("FileKey", (value))
}

// GetFileKey gets the value of FileKey for the instance
func (instance *FileIo_PathOperation) GetPropertyFileKey() (value uint32, err error) {
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

// SetFileName sets the value of FileName for the instance
func (instance *FileIo_PathOperation) SetPropertyFileName(value string) (err error) {
	return instance.SetProperty("FileName", (value))
}

// GetFileName gets the value of FileName for the instance
func (instance *FileIo_PathOperation) GetPropertyFileName() (value string, err error) {
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

// SetFileObject sets the value of FileObject for the instance
func (instance *FileIo_PathOperation) SetPropertyFileObject(value uint32) (err error) {
	return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *FileIo_PathOperation) GetPropertyFileObject() (value uint32, err error) {
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
func (instance *FileIo_PathOperation) SetPropertyInfoClass(value uint32) (err error) {
	return instance.SetProperty("InfoClass", (value))
}

// GetInfoClass gets the value of InfoClass for the instance
func (instance *FileIo_PathOperation) GetPropertyInfoClass() (value uint32, err error) {
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
func (instance *FileIo_PathOperation) SetPropertyIrpPtr(value uint32) (err error) {
	return instance.SetProperty("IrpPtr", (value))
}

// GetIrpPtr gets the value of IrpPtr for the instance
func (instance *FileIo_PathOperation) GetPropertyIrpPtr() (value uint32, err error) {
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
func (instance *FileIo_PathOperation) SetPropertyTTID(value uint32) (err error) {
	return instance.SetProperty("TTID", (value))
}

// GetTTID gets the value of TTID for the instance
func (instance *FileIo_PathOperation) GetPropertyTTID() (value uint32, err error) {
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
