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

// MSFT_NetNatExternalAddress struct
type MSFT_NetNatExternalAddress struct {
	*MSFT_NetSettingData

	//
	Active uint8

	//
	ExternalAddressID uint32

	//
	IPAddress string

	//
	NatName string

	//
	PortEnd uint16

	//
	PortStart uint16
}

func NewMSFT_NetNatExternalAddressEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetNatExternalAddress, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNatExternalAddress{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetNatExternalAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetNatExternalAddress, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNatExternalAddress{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFT_NetNatExternalAddress) SetPropertyActive(value uint8) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFT_NetNatExternalAddress) GetPropertyActive() (value uint8, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetExternalAddressID sets the value of ExternalAddressID for the instance
func (instance *MSFT_NetNatExternalAddress) SetPropertyExternalAddressID(value uint32) (err error) {
	return instance.SetProperty("ExternalAddressID", (value))
}

// GetExternalAddressID gets the value of ExternalAddressID for the instance
func (instance *MSFT_NetNatExternalAddress) GetPropertyExternalAddressID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExternalAddressID")
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

// SetIPAddress sets the value of IPAddress for the instance
func (instance *MSFT_NetNatExternalAddress) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", (value))
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *MSFT_NetNatExternalAddress) GetPropertyIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddress")
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

// SetNatName sets the value of NatName for the instance
func (instance *MSFT_NetNatExternalAddress) SetPropertyNatName(value string) (err error) {
	return instance.SetProperty("NatName", (value))
}

// GetNatName gets the value of NatName for the instance
func (instance *MSFT_NetNatExternalAddress) GetPropertyNatName() (value string, err error) {
	retValue, err := instance.GetProperty("NatName")
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

// SetPortEnd sets the value of PortEnd for the instance
func (instance *MSFT_NetNatExternalAddress) SetPropertyPortEnd(value uint16) (err error) {
	return instance.SetProperty("PortEnd", (value))
}

// GetPortEnd gets the value of PortEnd for the instance
func (instance *MSFT_NetNatExternalAddress) GetPropertyPortEnd() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortEnd")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetPortStart sets the value of PortStart for the instance
func (instance *MSFT_NetNatExternalAddress) SetPropertyPortStart(value uint16) (err error) {
	return instance.SetProperty("PortStart", (value))
}

// GetPortStart gets the value of PortStart for the instance
func (instance *MSFT_NetNatExternalAddress) GetPropertyPortStart() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortStart")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
