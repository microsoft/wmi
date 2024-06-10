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

// CIM_SecurityAssociationEndpoint struct
type CIM_SecurityAssociationEndpoint struct {
	*CIM_ProtocolEndpoint

	//
	IdleDurationSeconds uint64

	//
	LifetimeKilobytes uint64

	//
	LifetimeSeconds uint64

	//
	PacketLoggingActive bool

	//
	RefreshThresholdKbytesPercentage uint8

	//
	RefreshThresholdSecondsPercentage uint8
}

func NewCIM_SecurityAssociationEndpointEx1(instance *cim.WmiInstance) (newInstance *CIM_SecurityAssociationEndpoint, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SecurityAssociationEndpoint{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

func NewCIM_SecurityAssociationEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SecurityAssociationEndpoint, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SecurityAssociationEndpoint{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

// SetIdleDurationSeconds sets the value of IdleDurationSeconds for the instance
func (instance *CIM_SecurityAssociationEndpoint) SetPropertyIdleDurationSeconds(value uint64) (err error) {
	return instance.SetProperty("IdleDurationSeconds", (value))
}

// GetIdleDurationSeconds gets the value of IdleDurationSeconds for the instance
func (instance *CIM_SecurityAssociationEndpoint) GetPropertyIdleDurationSeconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("IdleDurationSeconds")
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

// SetLifetimeKilobytes sets the value of LifetimeKilobytes for the instance
func (instance *CIM_SecurityAssociationEndpoint) SetPropertyLifetimeKilobytes(value uint64) (err error) {
	return instance.SetProperty("LifetimeKilobytes", (value))
}

// GetLifetimeKilobytes gets the value of LifetimeKilobytes for the instance
func (instance *CIM_SecurityAssociationEndpoint) GetPropertyLifetimeKilobytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("LifetimeKilobytes")
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

// SetLifetimeSeconds sets the value of LifetimeSeconds for the instance
func (instance *CIM_SecurityAssociationEndpoint) SetPropertyLifetimeSeconds(value uint64) (err error) {
	return instance.SetProperty("LifetimeSeconds", (value))
}

// GetLifetimeSeconds gets the value of LifetimeSeconds for the instance
func (instance *CIM_SecurityAssociationEndpoint) GetPropertyLifetimeSeconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("LifetimeSeconds")
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

// SetPacketLoggingActive sets the value of PacketLoggingActive for the instance
func (instance *CIM_SecurityAssociationEndpoint) SetPropertyPacketLoggingActive(value bool) (err error) {
	return instance.SetProperty("PacketLoggingActive", (value))
}

// GetPacketLoggingActive gets the value of PacketLoggingActive for the instance
func (instance *CIM_SecurityAssociationEndpoint) GetPropertyPacketLoggingActive() (value bool, err error) {
	retValue, err := instance.GetProperty("PacketLoggingActive")
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

// SetRefreshThresholdKbytesPercentage sets the value of RefreshThresholdKbytesPercentage for the instance
func (instance *CIM_SecurityAssociationEndpoint) SetPropertyRefreshThresholdKbytesPercentage(value uint8) (err error) {
	return instance.SetProperty("RefreshThresholdKbytesPercentage", (value))
}

// GetRefreshThresholdKbytesPercentage gets the value of RefreshThresholdKbytesPercentage for the instance
func (instance *CIM_SecurityAssociationEndpoint) GetPropertyRefreshThresholdKbytesPercentage() (value uint8, err error) {
	retValue, err := instance.GetProperty("RefreshThresholdKbytesPercentage")
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

// SetRefreshThresholdSecondsPercentage sets the value of RefreshThresholdSecondsPercentage for the instance
func (instance *CIM_SecurityAssociationEndpoint) SetPropertyRefreshThresholdSecondsPercentage(value uint8) (err error) {
	return instance.SetProperty("RefreshThresholdSecondsPercentage", (value))
}

// GetRefreshThresholdSecondsPercentage gets the value of RefreshThresholdSecondsPercentage for the instance
func (instance *CIM_SecurityAssociationEndpoint) GetPropertyRefreshThresholdSecondsPercentage() (value uint8, err error) {
	retValue, err := instance.GetProperty("RefreshThresholdSecondsPercentage")
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
