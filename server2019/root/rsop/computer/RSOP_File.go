// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_File struct
type RSOP_File struct {
	*RSOP_SecuritySettings

	//
	Mode File_Mode

	//
	OriginalPath string

	//
	Path string

	//
	SDDLString string
}

func NewRSOP_FileEx1(instance *cim.WmiInstance) (newInstance *RSOP_File, err error) {
	tmp, err := NewRSOP_SecuritySettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_File{
		RSOP_SecuritySettings: tmp,
	}
	return
}

func NewRSOP_FileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_File, err error) {
	tmp, err := NewRSOP_SecuritySettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_File{
		RSOP_SecuritySettings: tmp,
	}
	return
}

// SetMode sets the value of Mode for the instance
func (instance *RSOP_File) SetPropertyMode(value File_Mode) (err error) {
	return instance.SetProperty("Mode", value)
}

// GetMode gets the value of Mode for the instance
func (instance *RSOP_File) GetPropertyMode() (value File_Mode, err error) {
	retValue, err := instance.GetProperty("Mode")
	if err != nil {
		return
	}
	value, ok := retValue.(File_Mode)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOriginalPath sets the value of OriginalPath for the instance
func (instance *RSOP_File) SetPropertyOriginalPath(value string) (err error) {
	return instance.SetProperty("OriginalPath", value)
}

// GetOriginalPath gets the value of OriginalPath for the instance
func (instance *RSOP_File) GetPropertyOriginalPath() (value string, err error) {
	retValue, err := instance.GetProperty("OriginalPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPath sets the value of Path for the instance
func (instance *RSOP_File) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", value)
}

// GetPath gets the value of Path for the instance
func (instance *RSOP_File) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSDDLString sets the value of SDDLString for the instance
func (instance *RSOP_File) SetPropertySDDLString(value string) (err error) {
	return instance.SetProperty("SDDLString", value)
}

// GetSDDLString gets the value of SDDLString for the instance
func (instance *RSOP_File) GetPropertySDDLString() (value string, err error) {
	retValue, err := instance.GetProperty("SDDLString")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}