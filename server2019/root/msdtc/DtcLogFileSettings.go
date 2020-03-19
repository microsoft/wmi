// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// DtcLogFileSettings struct
type DtcLogFileSettings struct {
	*cim.WmiInstance

	//
	MaxSizeInMB uint32

	//
	Path string

	//
	SizeInMB uint32
}

func NewDtcLogFileSettingsEx1(instance *cim.WmiInstance) (newInstance *DtcLogFileSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DtcLogFileSettings{
		WmiInstance: tmp,
	}
	return
}

func NewDtcLogFileSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DtcLogFileSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DtcLogFileSettings{
		WmiInstance: tmp,
	}
	return
}

// SetMaxSizeInMB sets the value of MaxSizeInMB for the instance
func (instance *DtcLogFileSettings) SetPropertyMaxSizeInMB(value uint32) (err error) {
	return instance.SetProperty("MaxSizeInMB", value)
}

// GetMaxSizeInMB gets the value of MaxSizeInMB for the instance
func (instance *DtcLogFileSettings) GetPropertyMaxSizeInMB() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxSizeInMB")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPath sets the value of Path for the instance
func (instance *DtcLogFileSettings) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", value)
}

// GetPath gets the value of Path for the instance
func (instance *DtcLogFileSettings) GetPropertyPath() (value string, err error) {
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

// SetSizeInMB sets the value of SizeInMB for the instance
func (instance *DtcLogFileSettings) SetPropertySizeInMB(value uint32) (err error) {
	return instance.SetProperty("SizeInMB", value)
}

// GetSizeInMB gets the value of SizeInMB for the instance
func (instance *DtcLogFileSettings) GetPropertySizeInMB() (value uint32, err error) {
	retValue, err := instance.GetProperty("SizeInMB")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
