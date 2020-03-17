// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_IEConnectionWinINetSettings struct
type RSOP_IEConnectionWinINetSettings struct {
	cim.WmiInstance

	//
	connectionName string

	//
	internetPerConnOptionListData []uint8

	//
	internetPerConnOptionListDataSize uint32

	//
	rsopID string

	//
	rsopPrecedence int32
}

// SetconnectionName sets the value of connectionName for the instance
func (instance *RSOP_IEConnectionWinINetSettings) SetPropertyconnectionName(value string) (err error) {
	return instance.SetProperty("connectionName", value)
}

// GetconnectionName gets the value of connectionName for the instance
func (instance *RSOP_IEConnectionWinINetSettings) GetPropertyconnectionName() (value string, err error) {
	retValue, err := instance.GetProperty("connectionName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetinternetPerConnOptionListData sets the value of internetPerConnOptionListData for the instance
func (instance *RSOP_IEConnectionWinINetSettings) SetPropertyinternetPerConnOptionListData(value []uint8) (err error) {
	return instance.SetProperty("internetPerConnOptionListData", value)
}

// GetinternetPerConnOptionListData gets the value of internetPerConnOptionListData for the instance
func (instance *RSOP_IEConnectionWinINetSettings) GetPropertyinternetPerConnOptionListData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("internetPerConnOptionListData")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetinternetPerConnOptionListDataSize sets the value of internetPerConnOptionListDataSize for the instance
func (instance *RSOP_IEConnectionWinINetSettings) SetPropertyinternetPerConnOptionListDataSize(value uint32) (err error) {
	return instance.SetProperty("internetPerConnOptionListDataSize", value)
}

// GetinternetPerConnOptionListDataSize gets the value of internetPerConnOptionListDataSize for the instance
func (instance *RSOP_IEConnectionWinINetSettings) GetPropertyinternetPerConnOptionListDataSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("internetPerConnOptionListDataSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IEConnectionWinINetSettings) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", value)
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IEConnectionWinINetSettings) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IEConnectionWinINetSettings) SetPropertyrsopPrecedence(value int32) (err error) {
	return instance.SetProperty("rsopPrecedence", value)
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IEConnectionWinINetSettings) GetPropertyrsopPrecedence() (value int32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
