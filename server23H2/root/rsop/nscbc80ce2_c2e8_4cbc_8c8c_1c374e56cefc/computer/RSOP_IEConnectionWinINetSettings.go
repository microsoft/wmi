// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NSCBC80CE2_C2E8_4CBC_8C8C_1C374E56CEFC.Computer
//
// ////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEConnectionWinINetSettings struct
type RSOP_IEConnectionWinINetSettings struct {
	*cim.WmiInstance

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

func NewRSOP_IEConnectionWinINetSettingsEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEConnectionWinINetSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionWinINetSettings{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEConnectionWinINetSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEConnectionWinINetSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionWinINetSettings{
		WmiInstance: tmp,
	}
	return
}

// SetconnectionName sets the value of connectionName for the instance
func (instance *RSOP_IEConnectionWinINetSettings) SetPropertyconnectionName(value string) (err error) {
	return instance.SetProperty("connectionName", (value))
}

// GetconnectionName gets the value of connectionName for the instance
func (instance *RSOP_IEConnectionWinINetSettings) GetPropertyconnectionName() (value string, err error) {
	retValue, err := instance.GetProperty("connectionName")
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

// SetinternetPerConnOptionListData sets the value of internetPerConnOptionListData for the instance
func (instance *RSOP_IEConnectionWinINetSettings) SetPropertyinternetPerConnOptionListData(value []uint8) (err error) {
	return instance.SetProperty("internetPerConnOptionListData", (value))
}

// GetinternetPerConnOptionListData gets the value of internetPerConnOptionListData for the instance
func (instance *RSOP_IEConnectionWinINetSettings) GetPropertyinternetPerConnOptionListData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("internetPerConnOptionListData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetinternetPerConnOptionListDataSize sets the value of internetPerConnOptionListDataSize for the instance
func (instance *RSOP_IEConnectionWinINetSettings) SetPropertyinternetPerConnOptionListDataSize(value uint32) (err error) {
	return instance.SetProperty("internetPerConnOptionListDataSize", (value))
}

// GetinternetPerConnOptionListDataSize gets the value of internetPerConnOptionListDataSize for the instance
func (instance *RSOP_IEConnectionWinINetSettings) GetPropertyinternetPerConnOptionListDataSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("internetPerConnOptionListDataSize")
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

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IEConnectionWinINetSettings) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", (value))
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IEConnectionWinINetSettings) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
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

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IEConnectionWinINetSettings) SetPropertyrsopPrecedence(value int32) (err error) {
	return instance.SetProperty("rsopPrecedence", (value))
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IEConnectionWinINetSettings) GetPropertyrsopPrecedence() (value int32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
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

	value = int32(valuetmp)

	return
}
