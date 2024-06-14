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

// FileIo_V1_Name struct
type FileIo_V1_Name struct {
	*FileIo_V1

	//
	FileName string

	//
	FileObject uint32
}

func NewFileIo_V1_NameEx1(instance *cim.WmiInstance) (newInstance *FileIo_V1_Name, err error) {
	tmp, err := NewFileIo_V1Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &FileIo_V1_Name{
		FileIo_V1: tmp,
	}
	return
}

func NewFileIo_V1_NameEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FileIo_V1_Name, err error) {
	tmp, err := NewFileIo_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FileIo_V1_Name{
		FileIo_V1: tmp,
	}
	return
}

// SetFileName sets the value of FileName for the instance
func (instance *FileIo_V1_Name) SetPropertyFileName(value string) (err error) {
	return instance.SetProperty("FileName", (value))
}

// GetFileName gets the value of FileName for the instance
func (instance *FileIo_V1_Name) GetPropertyFileName() (value string, err error) {
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
func (instance *FileIo_V1_Name) SetPropertyFileObject(value uint32) (err error) {
	return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *FileIo_V1_Name) GetPropertyFileObject() (value uint32, err error) {
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
