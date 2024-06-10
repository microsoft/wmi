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

// FileIo_V2_Create struct
type FileIo_V2_Create struct {
	*FileIo_V2

	//
	CreateOptions uint32

	//
	FileAttributes uint32

	//
	FileObject uint32

	//
	IrpPtr uint32

	//
	OpenPath string

	//
	ShareAccess uint32

	//
	TTID uint32
}

func NewFileIo_V2_CreateEx1(instance *cim.WmiInstance) (newInstance *FileIo_V2_Create, err error) {
	tmp, err := NewFileIo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &FileIo_V2_Create{
		FileIo_V2: tmp,
	}
	return
}

func NewFileIo_V2_CreateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FileIo_V2_Create, err error) {
	tmp, err := NewFileIo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FileIo_V2_Create{
		FileIo_V2: tmp,
	}
	return
}

// SetCreateOptions sets the value of CreateOptions for the instance
func (instance *FileIo_V2_Create) SetPropertyCreateOptions(value uint32) (err error) {
	return instance.SetProperty("CreateOptions", (value))
}

// GetCreateOptions gets the value of CreateOptions for the instance
func (instance *FileIo_V2_Create) GetPropertyCreateOptions() (value uint32, err error) {
	retValue, err := instance.GetProperty("CreateOptions")
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

// SetFileAttributes sets the value of FileAttributes for the instance
func (instance *FileIo_V2_Create) SetPropertyFileAttributes(value uint32) (err error) {
	return instance.SetProperty("FileAttributes", (value))
}

// GetFileAttributes gets the value of FileAttributes for the instance
func (instance *FileIo_V2_Create) GetPropertyFileAttributes() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileAttributes")
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
func (instance *FileIo_V2_Create) SetPropertyFileObject(value uint32) (err error) {
	return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *FileIo_V2_Create) GetPropertyFileObject() (value uint32, err error) {
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
func (instance *FileIo_V2_Create) SetPropertyIrpPtr(value uint32) (err error) {
	return instance.SetProperty("IrpPtr", (value))
}

// GetIrpPtr gets the value of IrpPtr for the instance
func (instance *FileIo_V2_Create) GetPropertyIrpPtr() (value uint32, err error) {
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

// SetOpenPath sets the value of OpenPath for the instance
func (instance *FileIo_V2_Create) SetPropertyOpenPath(value string) (err error) {
	return instance.SetProperty("OpenPath", (value))
}

// GetOpenPath gets the value of OpenPath for the instance
func (instance *FileIo_V2_Create) GetPropertyOpenPath() (value string, err error) {
	retValue, err := instance.GetProperty("OpenPath")
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

// SetShareAccess sets the value of ShareAccess for the instance
func (instance *FileIo_V2_Create) SetPropertyShareAccess(value uint32) (err error) {
	return instance.SetProperty("ShareAccess", (value))
}

// GetShareAccess gets the value of ShareAccess for the instance
func (instance *FileIo_V2_Create) GetPropertyShareAccess() (value uint32, err error) {
	retValue, err := instance.GetProperty("ShareAccess")
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
func (instance *FileIo_V2_Create) SetPropertyTTID(value uint32) (err error) {
	return instance.SetProperty("TTID", (value))
}

// GetTTID gets the value of TTID for the instance
func (instance *FileIo_V2_Create) GetPropertyTTID() (value uint32, err error) {
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
