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

// MSFT_NetBaseIPProtocol struct
type MSFT_NetBaseIPProtocol struct {
	*CIM_ProtocolEndpoint

	//
	AddressMaskReply uint8

	//
	DeadGatewayDetection uint8

	//
	DefaultHopLimit uint32

	//
	DhcpMediaSense uint8

	//
	GroupForwardedFragments uint8

	//
	IcmpRedirects uint8

	//
	MediaSenseEventLog uint8

	//
	MldLevel uint32

	//
	MldVersion uint32

	//
	MulticastForwarding uint8

	//
	NeighborCacheLimit uint32

	//
	RandomizeIdentifiers uint8

	//
	ReassemblyLimit uint32

	//
	RouteCacheLimit uint32

	//
	SourceRoutingBehavior uint32
}

func NewMSFT_NetBaseIPProtocolEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetBaseIPProtocol, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBaseIPProtocol{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

func NewMSFT_NetBaseIPProtocolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetBaseIPProtocol, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBaseIPProtocol{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

// SetAddressMaskReply sets the value of AddressMaskReply for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyAddressMaskReply(value uint8) (err error) {
	return instance.SetProperty("AddressMaskReply", (value))
}

// GetAddressMaskReply gets the value of AddressMaskReply for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyAddressMaskReply() (value uint8, err error) {
	retValue, err := instance.GetProperty("AddressMaskReply")
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

// SetDeadGatewayDetection sets the value of DeadGatewayDetection for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyDeadGatewayDetection(value uint8) (err error) {
	return instance.SetProperty("DeadGatewayDetection", (value))
}

// GetDeadGatewayDetection gets the value of DeadGatewayDetection for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyDeadGatewayDetection() (value uint8, err error) {
	retValue, err := instance.GetProperty("DeadGatewayDetection")
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

// SetDefaultHopLimit sets the value of DefaultHopLimit for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyDefaultHopLimit(value uint32) (err error) {
	return instance.SetProperty("DefaultHopLimit", (value))
}

// GetDefaultHopLimit gets the value of DefaultHopLimit for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyDefaultHopLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultHopLimit")
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

// SetDhcpMediaSense sets the value of DhcpMediaSense for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyDhcpMediaSense(value uint8) (err error) {
	return instance.SetProperty("DhcpMediaSense", (value))
}

// GetDhcpMediaSense gets the value of DhcpMediaSense for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyDhcpMediaSense() (value uint8, err error) {
	retValue, err := instance.GetProperty("DhcpMediaSense")
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

// SetGroupForwardedFragments sets the value of GroupForwardedFragments for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyGroupForwardedFragments(value uint8) (err error) {
	return instance.SetProperty("GroupForwardedFragments", (value))
}

// GetGroupForwardedFragments gets the value of GroupForwardedFragments for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyGroupForwardedFragments() (value uint8, err error) {
	retValue, err := instance.GetProperty("GroupForwardedFragments")
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

// SetIcmpRedirects sets the value of IcmpRedirects for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyIcmpRedirects(value uint8) (err error) {
	return instance.SetProperty("IcmpRedirects", (value))
}

// GetIcmpRedirects gets the value of IcmpRedirects for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyIcmpRedirects() (value uint8, err error) {
	retValue, err := instance.GetProperty("IcmpRedirects")
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

// SetMediaSenseEventLog sets the value of MediaSenseEventLog for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyMediaSenseEventLog(value uint8) (err error) {
	return instance.SetProperty("MediaSenseEventLog", (value))
}

// GetMediaSenseEventLog gets the value of MediaSenseEventLog for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyMediaSenseEventLog() (value uint8, err error) {
	retValue, err := instance.GetProperty("MediaSenseEventLog")
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

// SetMldLevel sets the value of MldLevel for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyMldLevel(value uint32) (err error) {
	return instance.SetProperty("MldLevel", (value))
}

// GetMldLevel gets the value of MldLevel for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyMldLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("MldLevel")
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

// SetMldVersion sets the value of MldVersion for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyMldVersion(value uint32) (err error) {
	return instance.SetProperty("MldVersion", (value))
}

// GetMldVersion gets the value of MldVersion for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyMldVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("MldVersion")
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

// SetMulticastForwarding sets the value of MulticastForwarding for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyMulticastForwarding(value uint8) (err error) {
	return instance.SetProperty("MulticastForwarding", (value))
}

// GetMulticastForwarding gets the value of MulticastForwarding for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyMulticastForwarding() (value uint8, err error) {
	retValue, err := instance.GetProperty("MulticastForwarding")
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

// SetNeighborCacheLimit sets the value of NeighborCacheLimit for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyNeighborCacheLimit(value uint32) (err error) {
	return instance.SetProperty("NeighborCacheLimit", (value))
}

// GetNeighborCacheLimit gets the value of NeighborCacheLimit for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyNeighborCacheLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("NeighborCacheLimit")
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

// SetRandomizeIdentifiers sets the value of RandomizeIdentifiers for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyRandomizeIdentifiers(value uint8) (err error) {
	return instance.SetProperty("RandomizeIdentifiers", (value))
}

// GetRandomizeIdentifiers gets the value of RandomizeIdentifiers for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyRandomizeIdentifiers() (value uint8, err error) {
	retValue, err := instance.GetProperty("RandomizeIdentifiers")
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

// SetReassemblyLimit sets the value of ReassemblyLimit for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyReassemblyLimit(value uint32) (err error) {
	return instance.SetProperty("ReassemblyLimit", (value))
}

// GetReassemblyLimit gets the value of ReassemblyLimit for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyReassemblyLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReassemblyLimit")
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

// SetRouteCacheLimit sets the value of RouteCacheLimit for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertyRouteCacheLimit(value uint32) (err error) {
	return instance.SetProperty("RouteCacheLimit", (value))
}

// GetRouteCacheLimit gets the value of RouteCacheLimit for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertyRouteCacheLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("RouteCacheLimit")
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

// SetSourceRoutingBehavior sets the value of SourceRoutingBehavior for the instance
func (instance *MSFT_NetBaseIPProtocol) SetPropertySourceRoutingBehavior(value uint32) (err error) {
	return instance.SetProperty("SourceRoutingBehavior", (value))
}

// GetSourceRoutingBehavior gets the value of SourceRoutingBehavior for the instance
func (instance *MSFT_NetBaseIPProtocol) GetPropertySourceRoutingBehavior() (value uint32, err error) {
	retValue, err := instance.GetProperty("SourceRoutingBehavior")
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
