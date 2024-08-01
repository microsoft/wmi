// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetEventWFPCaptureProvider struct
type MSFT_NetEventWFPCaptureProvider struct {
	*MSFT_NetEventProviderBase

	//
	CaptureLayerSet uint64

	//
	DiscardedEvents bool

	//
	IPAddresses []string

	//
	TCPPorts []uint16

	//
	UDPPorts []uint16
}

func NewMSFT_NetEventWFPCaptureProviderEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetEventWFPCaptureProvider, err error) {
	tmp, err := NewMSFT_NetEventProviderBaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventWFPCaptureProvider{
		MSFT_NetEventProviderBase: tmp,
	}
	return
}

func NewMSFT_NetEventWFPCaptureProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetEventWFPCaptureProvider, err error) {
	tmp, err := NewMSFT_NetEventProviderBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventWFPCaptureProvider{
		MSFT_NetEventProviderBase: tmp,
	}
	return
}

// SetCaptureLayerSet sets the value of CaptureLayerSet for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) SetPropertyCaptureLayerSet(value uint64) (err error) {
	return instance.SetProperty("CaptureLayerSet", (value))
}

// GetCaptureLayerSet gets the value of CaptureLayerSet for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) GetPropertyCaptureLayerSet() (value uint64, err error) {
	retValue, err := instance.GetProperty("CaptureLayerSet")
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

// SetDiscardedEvents sets the value of DiscardedEvents for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) SetPropertyDiscardedEvents(value bool) (err error) {
	return instance.SetProperty("DiscardedEvents", (value))
}

// GetDiscardedEvents gets the value of DiscardedEvents for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) GetPropertyDiscardedEvents() (value bool, err error) {
	retValue, err := instance.GetProperty("DiscardedEvents")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetIPAddresses sets the value of IPAddresses for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) SetPropertyIPAddresses(value []string) (err error) {
	return instance.SetProperty("IPAddresses", (value))
}

// GetIPAddresses gets the value of IPAddresses for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) GetPropertyIPAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IPAddresses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetTCPPorts sets the value of TCPPorts for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) SetPropertyTCPPorts(value []uint16) (err error) {
	return instance.SetProperty("TCPPorts", (value))
}

// GetTCPPorts gets the value of TCPPorts for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) GetPropertyTCPPorts() (value []uint16, err error) {
	retValue, err := instance.GetProperty("TCPPorts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetUDPPorts sets the value of UDPPorts for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) SetPropertyUDPPorts(value []uint16) (err error) {
	return instance.SetProperty("UDPPorts", (value))
}

// GetUDPPorts gets the value of UDPPorts for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) GetPropertyUDPPorts() (value []uint16, err error) {
	retValue, err := instance.GetProperty("UDPPorts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}
