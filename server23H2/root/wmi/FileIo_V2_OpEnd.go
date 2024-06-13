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

// FileIo_V2_OpEnd struct
type FileIo_V2_OpEnd struct {
	*FileIo_V2

	//
	ExtraInfo uint32

	//
	IrpPtr uint32

	//
	NtStatus uint32
}

func NewFileIo_V2_OpEndEx1(instance *cim.WmiInstance) (newInstance *FileIo_V2_OpEnd, err error) {
	tmp, err := NewFileIo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &FileIo_V2_OpEnd{
		FileIo_V2: tmp,
	}
	return
}

func NewFileIo_V2_OpEndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FileIo_V2_OpEnd, err error) {
	tmp, err := NewFileIo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FileIo_V2_OpEnd{
		FileIo_V2: tmp,
	}
	return
}

// SetExtraInfo sets the value of ExtraInfo for the instance
func (instance *FileIo_V2_OpEnd) SetPropertyExtraInfo(value uint32) (err error) {
	return instance.SetProperty("ExtraInfo", (value))
}

// GetExtraInfo gets the value of ExtraInfo for the instance
func (instance *FileIo_V2_OpEnd) GetPropertyExtraInfo() (value uint32, err error) {
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

// SetIrpPtr sets the value of IrpPtr for the instance
func (instance *FileIo_V2_OpEnd) SetPropertyIrpPtr(value uint32) (err error) {
	return instance.SetProperty("IrpPtr", (value))
}

// GetIrpPtr gets the value of IrpPtr for the instance
func (instance *FileIo_V2_OpEnd) GetPropertyIrpPtr() (value uint32, err error) {
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

// SetNtStatus sets the value of NtStatus for the instance
func (instance *FileIo_V2_OpEnd) SetPropertyNtStatus(value uint32) (err error) {
	return instance.SetProperty("NtStatus", (value))
}

// GetNtStatus gets the value of NtStatus for the instance
func (instance *FileIo_V2_OpEnd) GetPropertyNtStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("NtStatus")
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
