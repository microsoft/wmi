// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NS5254145F_8B72_4CC6_B513_4966D1F5A5CB.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_FileBlocked struct
type RSOP_FileBlocked struct {
	*RSOP_SecuritySettingsBlocked

	//
	Mode FileBlocked_Mode

	//
	OriginalPath string

	//
	Path string

	//
	SDDLString string
}

func NewRSOP_FileBlockedEx1(instance *cim.WmiInstance) (newInstance *RSOP_FileBlocked, err error) {
	tmp, err := NewRSOP_SecuritySettingsBlockedEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_FileBlocked{
		RSOP_SecuritySettingsBlocked: tmp,
	}
	return
}

func NewRSOP_FileBlockedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_FileBlocked, err error) {
	tmp, err := NewRSOP_SecuritySettingsBlockedEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_FileBlocked{
		RSOP_SecuritySettingsBlocked: tmp,
	}
	return
}

// SetMode sets the value of Mode for the instance
func (instance *RSOP_FileBlocked) SetPropertyMode(value FileBlocked_Mode) (err error) {
	return instance.SetProperty("Mode", (value))
}

// GetMode gets the value of Mode for the instance
func (instance *RSOP_FileBlocked) GetPropertyMode() (value FileBlocked_Mode, err error) {
	retValue, err := instance.GetProperty("Mode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FileBlocked_Mode(valuetmp)

	return
}

// SetOriginalPath sets the value of OriginalPath for the instance
func (instance *RSOP_FileBlocked) SetPropertyOriginalPath(value string) (err error) {
	return instance.SetProperty("OriginalPath", (value))
}

// GetOriginalPath gets the value of OriginalPath for the instance
func (instance *RSOP_FileBlocked) GetPropertyOriginalPath() (value string, err error) {
	retValue, err := instance.GetProperty("OriginalPath")
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

// SetPath sets the value of Path for the instance
func (instance *RSOP_FileBlocked) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *RSOP_FileBlocked) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetSDDLString sets the value of SDDLString for the instance
func (instance *RSOP_FileBlocked) SetPropertySDDLString(value string) (err error) {
	return instance.SetProperty("SDDLString", (value))
}

// GetSDDLString gets the value of SDDLString for the instance
func (instance *RSOP_FileBlocked) GetPropertySDDLString() (value string, err error) {
	retValue, err := instance.GetProperty("SDDLString")
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
