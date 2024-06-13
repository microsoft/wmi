// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_WmiIPSecOffloadV1_Supported struct
type MSNdis_WmiIPSecOffloadV1_Supported struct {
	*MSNdis

	//
	AhEspCombined uint32

	//
	Encapsulation uint32

	//
	Flags uint32

	//
	IPv4Options uint32

	//
	TransportTunnelCombined uint32
}

func NewMSNdis_WmiIPSecOffloadV1_SupportedEx1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiIPSecOffloadV1_Supported, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiIPSecOffloadV1_Supported{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiIPSecOffloadV1_SupportedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiIPSecOffloadV1_Supported, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiIPSecOffloadV1_Supported{
		MSNdis: tmp,
	}
	return
}

// SetAhEspCombined sets the value of AhEspCombined for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_Supported) SetPropertyAhEspCombined(value uint32) (err error) {
	return instance.SetProperty("AhEspCombined", (value))
}

// GetAhEspCombined gets the value of AhEspCombined for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_Supported) GetPropertyAhEspCombined() (value uint32, err error) {
	retValue, err := instance.GetProperty("AhEspCombined")
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

// SetEncapsulation sets the value of Encapsulation for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_Supported) SetPropertyEncapsulation(value uint32) (err error) {
	return instance.SetProperty("Encapsulation", (value))
}

// GetEncapsulation gets the value of Encapsulation for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_Supported) GetPropertyEncapsulation() (value uint32, err error) {
	retValue, err := instance.GetProperty("Encapsulation")
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

// SetFlags sets the value of Flags for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_Supported) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_Supported) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetIPv4Options sets the value of IPv4Options for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_Supported) SetPropertyIPv4Options(value uint32) (err error) {
	return instance.SetProperty("IPv4Options", (value))
}

// GetIPv4Options gets the value of IPv4Options for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_Supported) GetPropertyIPv4Options() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4Options")
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

// SetTransportTunnelCombined sets the value of TransportTunnelCombined for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_Supported) SetPropertyTransportTunnelCombined(value uint32) (err error) {
	return instance.SetProperty("TransportTunnelCombined", (value))
}

// GetTransportTunnelCombined gets the value of TransportTunnelCombined for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_Supported) GetPropertyTransportTunnelCombined() (value uint32, err error) {
	retValue, err := instance.GetProperty("TransportTunnelCombined")
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
