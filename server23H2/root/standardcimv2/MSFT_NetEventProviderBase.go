// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetEventProviderBase struct
type MSFT_NetEventProviderBase struct {
	*CIM_LogicalElement

	//
	Guid string

	//
	Level uint8

	//
	MatchAllKeyword uint64

	//
	MatchAnyKeyword uint64

	//
	SessionGuid string

	//
	SessionName string
}

func NewMSFT_NetEventProviderBaseEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetEventProviderBase, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventProviderBase{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewMSFT_NetEventProviderBaseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetEventProviderBase, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventProviderBase{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetGuid sets the value of Guid for the instance
func (instance *MSFT_NetEventProviderBase) SetPropertyGuid(value string) (err error) {
	return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *MSFT_NetEventProviderBase) GetPropertyGuid() (value string, err error) {
	retValue, err := instance.GetProperty("Guid")
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

// SetLevel sets the value of Level for the instance
func (instance *MSFT_NetEventProviderBase) SetPropertyLevel(value uint8) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *MSFT_NetEventProviderBase) GetPropertyLevel() (value uint8, err error) {
	retValue, err := instance.GetProperty("Level")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetMatchAllKeyword sets the value of MatchAllKeyword for the instance
func (instance *MSFT_NetEventProviderBase) SetPropertyMatchAllKeyword(value uint64) (err error) {
	return instance.SetProperty("MatchAllKeyword", (value))
}

// GetMatchAllKeyword gets the value of MatchAllKeyword for the instance
func (instance *MSFT_NetEventProviderBase) GetPropertyMatchAllKeyword() (value uint64, err error) {
	retValue, err := instance.GetProperty("MatchAllKeyword")
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

// SetMatchAnyKeyword sets the value of MatchAnyKeyword for the instance
func (instance *MSFT_NetEventProviderBase) SetPropertyMatchAnyKeyword(value uint64) (err error) {
	return instance.SetProperty("MatchAnyKeyword", (value))
}

// GetMatchAnyKeyword gets the value of MatchAnyKeyword for the instance
func (instance *MSFT_NetEventProviderBase) GetPropertyMatchAnyKeyword() (value uint64, err error) {
	retValue, err := instance.GetProperty("MatchAnyKeyword")
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

// SetSessionGuid sets the value of SessionGuid for the instance
func (instance *MSFT_NetEventProviderBase) SetPropertySessionGuid(value string) (err error) {
	return instance.SetProperty("SessionGuid", (value))
}

// GetSessionGuid gets the value of SessionGuid for the instance
func (instance *MSFT_NetEventProviderBase) GetPropertySessionGuid() (value string, err error) {
	retValue, err := instance.GetProperty("SessionGuid")
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

// SetSessionName sets the value of SessionName for the instance
func (instance *MSFT_NetEventProviderBase) SetPropertySessionName(value string) (err error) {
	return instance.SetProperty("SessionName", (value))
}

// GetSessionName gets the value of SessionName for the instance
func (instance *MSFT_NetEventProviderBase) GetPropertySessionName() (value string, err error) {
	retValue, err := instance.GetProperty("SessionName")
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
