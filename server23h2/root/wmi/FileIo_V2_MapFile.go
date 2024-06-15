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

// FileIo_V2_MapFile struct
type FileIo_V2_MapFile struct {
	*FileIo_V2

	//
	FileObject uint32

	//
	MiscInfo uint64

	//
	ProcessId uint32

	//
	ViewBase uint32

	//
	ViewSize interface{}
}

func NewFileIo_V2_MapFileEx1(instance *cim.WmiInstance) (newInstance *FileIo_V2_MapFile, err error) {
	tmp, err := NewFileIo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &FileIo_V2_MapFile{
		FileIo_V2: tmp,
	}
	return
}

func NewFileIo_V2_MapFileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FileIo_V2_MapFile, err error) {
	tmp, err := NewFileIo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FileIo_V2_MapFile{
		FileIo_V2: tmp,
	}
	return
}

// SetFileObject sets the value of FileObject for the instance
func (instance *FileIo_V2_MapFile) SetPropertyFileObject(value uint32) (err error) {
	return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *FileIo_V2_MapFile) GetPropertyFileObject() (value uint32, err error) {
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

// SetMiscInfo sets the value of MiscInfo for the instance
func (instance *FileIo_V2_MapFile) SetPropertyMiscInfo(value uint64) (err error) {
	return instance.SetProperty("MiscInfo", (value))
}

// GetMiscInfo gets the value of MiscInfo for the instance
func (instance *FileIo_V2_MapFile) GetPropertyMiscInfo() (value uint64, err error) {
	retValue, err := instance.GetProperty("MiscInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetProcessId sets the value of ProcessId for the instance
func (instance *FileIo_V2_MapFile) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *FileIo_V2_MapFile) GetPropertyProcessId() (value uint32, err error) {
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

// SetViewBase sets the value of ViewBase for the instance
func (instance *FileIo_V2_MapFile) SetPropertyViewBase(value uint32) (err error) {
	return instance.SetProperty("ViewBase", (value))
}

// GetViewBase gets the value of ViewBase for the instance
func (instance *FileIo_V2_MapFile) GetPropertyViewBase() (value uint32, err error) {
	retValue, err := instance.GetProperty("ViewBase")
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

// SetViewSize sets the value of ViewSize for the instance
func (instance *FileIo_V2_MapFile) SetPropertyViewSize(value interface{}) (err error) {
	return instance.SetProperty("ViewSize", (value))
}

// GetViewSize gets the value of ViewSize for the instance
func (instance *FileIo_V2_MapFile) GetPropertyViewSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ViewSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
