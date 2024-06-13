// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.EventTracingManagement
//////////////////////////////////////////////
package eventtracingmanagement

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_EtwTraceProvider struct
type MSFT_EtwTraceProvider struct {
	*CIM_LogicalElement

	//
	AutologgerName string

	//
	Guid string

	//
	Level uint8

	//
	MatchAllKeyword uint64

	//
	MatchAnyKeyword uint64

	//
	Property uint32

	//
	SessionName string
}

func NewMSFT_EtwTraceProviderEx1(instance *cim.WmiInstance) (newInstance *MSFT_EtwTraceProvider, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_EtwTraceProvider{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewMSFT_EtwTraceProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_EtwTraceProvider, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_EtwTraceProvider{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetAutologgerName sets the value of AutologgerName for the instance
func (instance *MSFT_EtwTraceProvider) SetPropertyAutologgerName(value string) (err error) {
	return instance.SetProperty("AutologgerName", (value))
}

// GetAutologgerName gets the value of AutologgerName for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertyAutologgerName() (value string, err error) {
	retValue, err := instance.GetProperty("AutologgerName")
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

// SetGuid sets the value of Guid for the instance
func (instance *MSFT_EtwTraceProvider) SetPropertyGuid(value string) (err error) {
	return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertyGuid() (value string, err error) {
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
func (instance *MSFT_EtwTraceProvider) SetPropertyLevel(value uint8) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertyLevel() (value uint8, err error) {
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
func (instance *MSFT_EtwTraceProvider) SetPropertyMatchAllKeyword(value uint64) (err error) {
	return instance.SetProperty("MatchAllKeyword", (value))
}

// GetMatchAllKeyword gets the value of MatchAllKeyword for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertyMatchAllKeyword() (value uint64, err error) {
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
func (instance *MSFT_EtwTraceProvider) SetPropertyMatchAnyKeyword(value uint64) (err error) {
	return instance.SetProperty("MatchAnyKeyword", (value))
}

// GetMatchAnyKeyword gets the value of MatchAnyKeyword for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertyMatchAnyKeyword() (value uint64, err error) {
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

// SetProperty sets the value of Property for the instance
func (instance *MSFT_EtwTraceProvider) SetPropertyProperty(value uint32) (err error) {
	return instance.SetProperty("Property", (value))
}

// GetProperty gets the value of Property for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertyProperty() (value uint32, err error) {
	retValue, err := instance.GetProperty("Property")
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

// SetSessionName sets the value of SessionName for the instance
func (instance *MSFT_EtwTraceProvider) SetPropertySessionName(value string) (err error) {
	return instance.SetProperty("SessionName", (value))
}

// GetSessionName gets the value of SessionName for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertySessionName() (value string, err error) {
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
