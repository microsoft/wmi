// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetIPv4Protocol struct
type MSFT_NetIPv4Protocol struct {
	*MSFT_NetBaseIPProtocol

	//
	MinimumMtu uint32
}

func NewMSFT_NetIPv4ProtocolEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIPv4Protocol, err error) {
	tmp, err := NewMSFT_NetBaseIPProtocolEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPv4Protocol{
		MSFT_NetBaseIPProtocol: tmp,
	}
	return
}

func NewMSFT_NetIPv4ProtocolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIPv4Protocol, err error) {
	tmp, err := NewMSFT_NetBaseIPProtocolEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPv4Protocol{
		MSFT_NetBaseIPProtocol: tmp,
	}
	return
}

// SetMinimumMtu sets the value of MinimumMtu for the instance
func (instance *MSFT_NetIPv4Protocol) SetPropertyMinimumMtu(value uint32) (err error) {
	return instance.SetProperty("MinimumMtu", (value))
}

// GetMinimumMtu gets the value of MinimumMtu for the instance
func (instance *MSFT_NetIPv4Protocol) GetPropertyMinimumMtu() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinimumMtu")
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
