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

// FileIo_SimpleOp struct
type FileIo_SimpleOp struct {
	*FileIo

	//
	FileKey uint32

	//
	FileObject uint32

	//
	IrpPtr uint32

	//
	TTID uint32
}

func NewFileIo_SimpleOpEx1(instance *cim.WmiInstance) (newInstance *FileIo_SimpleOp, err error) {
	tmp, err := NewFileIoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FileIo_SimpleOp{
		FileIo: tmp,
	}
	return
}

func NewFileIo_SimpleOpEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FileIo_SimpleOp, err error) {
	tmp, err := NewFileIoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FileIo_SimpleOp{
		FileIo: tmp,
	}
	return
}

// SetFileKey sets the value of FileKey for the instance
func (instance *FileIo_SimpleOp) SetPropertyFileKey(value uint32) (err error) {
	return instance.SetProperty("FileKey", (value))
}

// GetFileKey gets the value of FileKey for the instance
func (instance *FileIo_SimpleOp) GetPropertyFileKey() (value uint32, err error) {
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
func (instance *FileIo_SimpleOp) SetPropertyFileObject(value uint32) (err error) {
	return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *FileIo_SimpleOp) GetPropertyFileObject() (value uint32, err error) {
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

// SetIrpPtr sets the value of IrpPtr for the instance
func (instance *FileIo_SimpleOp) SetPropertyIrpPtr(value uint32) (err error) {
	return instance.SetProperty("IrpPtr", (value))
}

// GetIrpPtr gets the value of IrpPtr for the instance
func (instance *FileIo_SimpleOp) GetPropertyIrpPtr() (value uint32, err error) {
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
func (instance *FileIo_SimpleOp) SetPropertyTTID(value uint32) (err error) {
	return instance.SetProperty("TTID", (value))
}

// GetTTID gets the value of TTID for the instance
func (instance *FileIo_SimpleOp) GetPropertyTTID() (value uint32, err error) {
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
