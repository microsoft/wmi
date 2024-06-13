// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetIPv6Protocol struct
type MSFT_NetIPv6Protocol struct {
	*MSFT_NetBaseIPProtocol

	//
	MaxDadAttempts uint32

	//
	MaxPreferredLifetime string

	//
	MaxRandomTime string

	//
	MaxValidLifetime string

	//
	RegenerateTime string

	//
	UseTemporaryAddresses uint32
}

func NewMSFT_NetIPv6ProtocolEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIPv6Protocol, err error) {
	tmp, err := NewMSFT_NetBaseIPProtocolEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPv6Protocol{
		MSFT_NetBaseIPProtocol: tmp,
	}
	return
}

func NewMSFT_NetIPv6ProtocolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIPv6Protocol, err error) {
	tmp, err := NewMSFT_NetBaseIPProtocolEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPv6Protocol{
		MSFT_NetBaseIPProtocol: tmp,
	}
	return
}

// SetMaxDadAttempts sets the value of MaxDadAttempts for the instance
func (instance *MSFT_NetIPv6Protocol) SetPropertyMaxDadAttempts(value uint32) (err error) {
	return instance.SetProperty("MaxDadAttempts", (value))
}

// GetMaxDadAttempts gets the value of MaxDadAttempts for the instance
func (instance *MSFT_NetIPv6Protocol) GetPropertyMaxDadAttempts() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxDadAttempts")
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

// SetMaxPreferredLifetime sets the value of MaxPreferredLifetime for the instance
func (instance *MSFT_NetIPv6Protocol) SetPropertyMaxPreferredLifetime(value string) (err error) {
	return instance.SetProperty("MaxPreferredLifetime", (value))
}

// GetMaxPreferredLifetime gets the value of MaxPreferredLifetime for the instance
func (instance *MSFT_NetIPv6Protocol) GetPropertyMaxPreferredLifetime() (value string, err error) {
	retValue, err := instance.GetProperty("MaxPreferredLifetime")
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

// SetMaxRandomTime sets the value of MaxRandomTime for the instance
func (instance *MSFT_NetIPv6Protocol) SetPropertyMaxRandomTime(value string) (err error) {
	return instance.SetProperty("MaxRandomTime", (value))
}

// GetMaxRandomTime gets the value of MaxRandomTime for the instance
func (instance *MSFT_NetIPv6Protocol) GetPropertyMaxRandomTime() (value string, err error) {
	retValue, err := instance.GetProperty("MaxRandomTime")
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

// SetMaxValidLifetime sets the value of MaxValidLifetime for the instance
func (instance *MSFT_NetIPv6Protocol) SetPropertyMaxValidLifetime(value string) (err error) {
	return instance.SetProperty("MaxValidLifetime", (value))
}

// GetMaxValidLifetime gets the value of MaxValidLifetime for the instance
func (instance *MSFT_NetIPv6Protocol) GetPropertyMaxValidLifetime() (value string, err error) {
	retValue, err := instance.GetProperty("MaxValidLifetime")
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

// SetRegenerateTime sets the value of RegenerateTime for the instance
func (instance *MSFT_NetIPv6Protocol) SetPropertyRegenerateTime(value string) (err error) {
	return instance.SetProperty("RegenerateTime", (value))
}

// GetRegenerateTime gets the value of RegenerateTime for the instance
func (instance *MSFT_NetIPv6Protocol) GetPropertyRegenerateTime() (value string, err error) {
	retValue, err := instance.GetProperty("RegenerateTime")
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

// SetUseTemporaryAddresses sets the value of UseTemporaryAddresses for the instance
func (instance *MSFT_NetIPv6Protocol) SetPropertyUseTemporaryAddresses(value uint32) (err error) {
	return instance.SetProperty("UseTemporaryAddresses", (value))
}

// GetUseTemporaryAddresses gets the value of UseTemporaryAddresses for the instance
func (instance *MSFT_NetIPv6Protocol) GetPropertyUseTemporaryAddresses() (value uint32, err error) {
	retValue, err := instance.GetProperty("UseTemporaryAddresses")
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
