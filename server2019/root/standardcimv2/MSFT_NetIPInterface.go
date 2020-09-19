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

// MSFT_NetIPInterface struct
type MSFT_NetIPInterface struct {
	*CIM_LANEndpoint

	//
	AddressFamily uint16

	//
	AdvertiseDefaultRoute uint8

	//
	AdvertisedRouterLifetime string

	//
	Advertising uint8

	//
	AutomaticMetric uint8

	//
	BaseReachableTime uint32

	//
	ClampMss uint8

	//
	CompartmentId uint32

	//
	ConnectionState uint8

	//
	CurrentHopLimit uint32

	//
	DadRetransmitTime uint32

	//
	DadTransmits uint32

	//
	Dhcp uint8

	//
	DirectedMacWolPattern uint8

	//
	EcnMarking uint8

	//
	ForceArpNdWolPattern uint8

	//
	Forwarding uint8

	//
	IgnoreDefaultRoutes uint8

	//
	InterfaceAlias string

	//
	InterfaceIndex uint32

	//
	InterfaceMetric uint32

	//
	IsolationId uint32

	//
	LowestIfNetLuid uint64

	//
	ManagedAddressConfiguration uint8

	//
	NeighborDiscoverySupported uint8

	//
	NeighborUnreachabilityDetection uint8

	//
	NlMtu uint32

	//
	OtherStatefulConfiguration uint8

	//
	ReachableTime uint32

	//
	RetransmitTime uint32

	//
	RouterDiscovery uint8

	//
	Store uint8

	//
	WeakHostReceive uint8

	//
	WeakHostSend uint8
}

func NewMSFT_NetIPInterfaceEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIPInterface, err error) {
	tmp, err := NewCIM_LANEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPInterface{
		CIM_LANEndpoint: tmp,
	}
	return
}

func NewMSFT_NetIPInterfaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIPInterface, err error) {
	tmp, err := NewCIM_LANEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPInterface{
		CIM_LANEndpoint: tmp,
	}
	return
}

// SetAddressFamily sets the value of AddressFamily for the instance
func (instance *MSFT_NetIPInterface) SetPropertyAddressFamily(value uint16) (err error) {
	return instance.SetProperty("AddressFamily", (value))
}

// GetAddressFamily gets the value of AddressFamily for the instance
func (instance *MSFT_NetIPInterface) GetPropertyAddressFamily() (value uint16, err error) {
	retValue, err := instance.GetProperty("AddressFamily")
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

// SetAdvertiseDefaultRoute sets the value of AdvertiseDefaultRoute for the instance
func (instance *MSFT_NetIPInterface) SetPropertyAdvertiseDefaultRoute(value uint8) (err error) {
	return instance.SetProperty("AdvertiseDefaultRoute", (value))
}

// GetAdvertiseDefaultRoute gets the value of AdvertiseDefaultRoute for the instance
func (instance *MSFT_NetIPInterface) GetPropertyAdvertiseDefaultRoute() (value uint8, err error) {
	retValue, err := instance.GetProperty("AdvertiseDefaultRoute")
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

// SetAdvertisedRouterLifetime sets the value of AdvertisedRouterLifetime for the instance
func (instance *MSFT_NetIPInterface) SetPropertyAdvertisedRouterLifetime(value string) (err error) {
	return instance.SetProperty("AdvertisedRouterLifetime", (value))
}

// GetAdvertisedRouterLifetime gets the value of AdvertisedRouterLifetime for the instance
func (instance *MSFT_NetIPInterface) GetPropertyAdvertisedRouterLifetime() (value string, err error) {
	retValue, err := instance.GetProperty("AdvertisedRouterLifetime")
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

// SetAdvertising sets the value of Advertising for the instance
func (instance *MSFT_NetIPInterface) SetPropertyAdvertising(value uint8) (err error) {
	return instance.SetProperty("Advertising", (value))
}

// GetAdvertising gets the value of Advertising for the instance
func (instance *MSFT_NetIPInterface) GetPropertyAdvertising() (value uint8, err error) {
	retValue, err := instance.GetProperty("Advertising")
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

// SetAutomaticMetric sets the value of AutomaticMetric for the instance
func (instance *MSFT_NetIPInterface) SetPropertyAutomaticMetric(value uint8) (err error) {
	return instance.SetProperty("AutomaticMetric", (value))
}

// GetAutomaticMetric gets the value of AutomaticMetric for the instance
func (instance *MSFT_NetIPInterface) GetPropertyAutomaticMetric() (value uint8, err error) {
	retValue, err := instance.GetProperty("AutomaticMetric")
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

// SetBaseReachableTime sets the value of BaseReachableTime for the instance
func (instance *MSFT_NetIPInterface) SetPropertyBaseReachableTime(value uint32) (err error) {
	return instance.SetProperty("BaseReachableTime", (value))
}

// GetBaseReachableTime gets the value of BaseReachableTime for the instance
func (instance *MSFT_NetIPInterface) GetPropertyBaseReachableTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("BaseReachableTime")
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

// SetClampMss sets the value of ClampMss for the instance
func (instance *MSFT_NetIPInterface) SetPropertyClampMss(value uint8) (err error) {
	return instance.SetProperty("ClampMss", (value))
}

// GetClampMss gets the value of ClampMss for the instance
func (instance *MSFT_NetIPInterface) GetPropertyClampMss() (value uint8, err error) {
	retValue, err := instance.GetProperty("ClampMss")
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

// SetCompartmentId sets the value of CompartmentId for the instance
func (instance *MSFT_NetIPInterface) SetPropertyCompartmentId(value uint32) (err error) {
	return instance.SetProperty("CompartmentId", (value))
}

// GetCompartmentId gets the value of CompartmentId for the instance
func (instance *MSFT_NetIPInterface) GetPropertyCompartmentId() (value uint32, err error) {
	retValue, err := instance.GetProperty("CompartmentId")
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

// SetConnectionState sets the value of ConnectionState for the instance
func (instance *MSFT_NetIPInterface) SetPropertyConnectionState(value uint8) (err error) {
	return instance.SetProperty("ConnectionState", (value))
}

// GetConnectionState gets the value of ConnectionState for the instance
func (instance *MSFT_NetIPInterface) GetPropertyConnectionState() (value uint8, err error) {
	retValue, err := instance.GetProperty("ConnectionState")
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

// SetCurrentHopLimit sets the value of CurrentHopLimit for the instance
func (instance *MSFT_NetIPInterface) SetPropertyCurrentHopLimit(value uint32) (err error) {
	return instance.SetProperty("CurrentHopLimit", (value))
}

// GetCurrentHopLimit gets the value of CurrentHopLimit for the instance
func (instance *MSFT_NetIPInterface) GetPropertyCurrentHopLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentHopLimit")
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

// SetDadRetransmitTime sets the value of DadRetransmitTime for the instance
func (instance *MSFT_NetIPInterface) SetPropertyDadRetransmitTime(value uint32) (err error) {
	return instance.SetProperty("DadRetransmitTime", (value))
}

// GetDadRetransmitTime gets the value of DadRetransmitTime for the instance
func (instance *MSFT_NetIPInterface) GetPropertyDadRetransmitTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("DadRetransmitTime")
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

// SetDadTransmits sets the value of DadTransmits for the instance
func (instance *MSFT_NetIPInterface) SetPropertyDadTransmits(value uint32) (err error) {
	return instance.SetProperty("DadTransmits", (value))
}

// GetDadTransmits gets the value of DadTransmits for the instance
func (instance *MSFT_NetIPInterface) GetPropertyDadTransmits() (value uint32, err error) {
	retValue, err := instance.GetProperty("DadTransmits")
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

// SetDhcp sets the value of Dhcp for the instance
func (instance *MSFT_NetIPInterface) SetPropertyDhcp(value uint8) (err error) {
	return instance.SetProperty("Dhcp", (value))
}

// GetDhcp gets the value of Dhcp for the instance
func (instance *MSFT_NetIPInterface) GetPropertyDhcp() (value uint8, err error) {
	retValue, err := instance.GetProperty("Dhcp")
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

// SetDirectedMacWolPattern sets the value of DirectedMacWolPattern for the instance
func (instance *MSFT_NetIPInterface) SetPropertyDirectedMacWolPattern(value uint8) (err error) {
	return instance.SetProperty("DirectedMacWolPattern", (value))
}

// GetDirectedMacWolPattern gets the value of DirectedMacWolPattern for the instance
func (instance *MSFT_NetIPInterface) GetPropertyDirectedMacWolPattern() (value uint8, err error) {
	retValue, err := instance.GetProperty("DirectedMacWolPattern")
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

// SetEcnMarking sets the value of EcnMarking for the instance
func (instance *MSFT_NetIPInterface) SetPropertyEcnMarking(value uint8) (err error) {
	return instance.SetProperty("EcnMarking", (value))
}

// GetEcnMarking gets the value of EcnMarking for the instance
func (instance *MSFT_NetIPInterface) GetPropertyEcnMarking() (value uint8, err error) {
	retValue, err := instance.GetProperty("EcnMarking")
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

// SetForceArpNdWolPattern sets the value of ForceArpNdWolPattern for the instance
func (instance *MSFT_NetIPInterface) SetPropertyForceArpNdWolPattern(value uint8) (err error) {
	return instance.SetProperty("ForceArpNdWolPattern", (value))
}

// GetForceArpNdWolPattern gets the value of ForceArpNdWolPattern for the instance
func (instance *MSFT_NetIPInterface) GetPropertyForceArpNdWolPattern() (value uint8, err error) {
	retValue, err := instance.GetProperty("ForceArpNdWolPattern")
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

// SetForwarding sets the value of Forwarding for the instance
func (instance *MSFT_NetIPInterface) SetPropertyForwarding(value uint8) (err error) {
	return instance.SetProperty("Forwarding", (value))
}

// GetForwarding gets the value of Forwarding for the instance
func (instance *MSFT_NetIPInterface) GetPropertyForwarding() (value uint8, err error) {
	retValue, err := instance.GetProperty("Forwarding")
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

// SetIgnoreDefaultRoutes sets the value of IgnoreDefaultRoutes for the instance
func (instance *MSFT_NetIPInterface) SetPropertyIgnoreDefaultRoutes(value uint8) (err error) {
	return instance.SetProperty("IgnoreDefaultRoutes", (value))
}

// GetIgnoreDefaultRoutes gets the value of IgnoreDefaultRoutes for the instance
func (instance *MSFT_NetIPInterface) GetPropertyIgnoreDefaultRoutes() (value uint8, err error) {
	retValue, err := instance.GetProperty("IgnoreDefaultRoutes")
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

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_NetIPInterface) SetPropertyInterfaceAlias(value string) (err error) {
	return instance.SetProperty("InterfaceAlias", (value))
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_NetIPInterface) GetPropertyInterfaceAlias() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceAlias")
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

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_NetIPInterface) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", (value))
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_NetIPInterface) GetPropertyInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceIndex")
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

// SetInterfaceMetric sets the value of InterfaceMetric for the instance
func (instance *MSFT_NetIPInterface) SetPropertyInterfaceMetric(value uint32) (err error) {
	return instance.SetProperty("InterfaceMetric", (value))
}

// GetInterfaceMetric gets the value of InterfaceMetric for the instance
func (instance *MSFT_NetIPInterface) GetPropertyInterfaceMetric() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceMetric")
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

// SetIsolationId sets the value of IsolationId for the instance
func (instance *MSFT_NetIPInterface) SetPropertyIsolationId(value uint32) (err error) {
	return instance.SetProperty("IsolationId", (value))
}

// GetIsolationId gets the value of IsolationId for the instance
func (instance *MSFT_NetIPInterface) GetPropertyIsolationId() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsolationId")
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

// SetLowestIfNetLuid sets the value of LowestIfNetLuid for the instance
func (instance *MSFT_NetIPInterface) SetPropertyLowestIfNetLuid(value uint64) (err error) {
	return instance.SetProperty("LowestIfNetLuid", (value))
}

// GetLowestIfNetLuid gets the value of LowestIfNetLuid for the instance
func (instance *MSFT_NetIPInterface) GetPropertyLowestIfNetLuid() (value uint64, err error) {
	retValue, err := instance.GetProperty("LowestIfNetLuid")
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

// SetManagedAddressConfiguration sets the value of ManagedAddressConfiguration for the instance
func (instance *MSFT_NetIPInterface) SetPropertyManagedAddressConfiguration(value uint8) (err error) {
	return instance.SetProperty("ManagedAddressConfiguration", (value))
}

// GetManagedAddressConfiguration gets the value of ManagedAddressConfiguration for the instance
func (instance *MSFT_NetIPInterface) GetPropertyManagedAddressConfiguration() (value uint8, err error) {
	retValue, err := instance.GetProperty("ManagedAddressConfiguration")
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

// SetNeighborDiscoverySupported sets the value of NeighborDiscoverySupported for the instance
func (instance *MSFT_NetIPInterface) SetPropertyNeighborDiscoverySupported(value uint8) (err error) {
	return instance.SetProperty("NeighborDiscoverySupported", (value))
}

// GetNeighborDiscoverySupported gets the value of NeighborDiscoverySupported for the instance
func (instance *MSFT_NetIPInterface) GetPropertyNeighborDiscoverySupported() (value uint8, err error) {
	retValue, err := instance.GetProperty("NeighborDiscoverySupported")
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

// SetNeighborUnreachabilityDetection sets the value of NeighborUnreachabilityDetection for the instance
func (instance *MSFT_NetIPInterface) SetPropertyNeighborUnreachabilityDetection(value uint8) (err error) {
	return instance.SetProperty("NeighborUnreachabilityDetection", (value))
}

// GetNeighborUnreachabilityDetection gets the value of NeighborUnreachabilityDetection for the instance
func (instance *MSFT_NetIPInterface) GetPropertyNeighborUnreachabilityDetection() (value uint8, err error) {
	retValue, err := instance.GetProperty("NeighborUnreachabilityDetection")
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

// SetNlMtu sets the value of NlMtu for the instance
func (instance *MSFT_NetIPInterface) SetPropertyNlMtu(value uint32) (err error) {
	return instance.SetProperty("NlMtu", (value))
}

// GetNlMtu gets the value of NlMtu for the instance
func (instance *MSFT_NetIPInterface) GetPropertyNlMtu() (value uint32, err error) {
	retValue, err := instance.GetProperty("NlMtu")
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

// SetOtherStatefulConfiguration sets the value of OtherStatefulConfiguration for the instance
func (instance *MSFT_NetIPInterface) SetPropertyOtherStatefulConfiguration(value uint8) (err error) {
	return instance.SetProperty("OtherStatefulConfiguration", (value))
}

// GetOtherStatefulConfiguration gets the value of OtherStatefulConfiguration for the instance
func (instance *MSFT_NetIPInterface) GetPropertyOtherStatefulConfiguration() (value uint8, err error) {
	retValue, err := instance.GetProperty("OtherStatefulConfiguration")
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

// SetReachableTime sets the value of ReachableTime for the instance
func (instance *MSFT_NetIPInterface) SetPropertyReachableTime(value uint32) (err error) {
	return instance.SetProperty("ReachableTime", (value))
}

// GetReachableTime gets the value of ReachableTime for the instance
func (instance *MSFT_NetIPInterface) GetPropertyReachableTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReachableTime")
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

// SetRetransmitTime sets the value of RetransmitTime for the instance
func (instance *MSFT_NetIPInterface) SetPropertyRetransmitTime(value uint32) (err error) {
	return instance.SetProperty("RetransmitTime", (value))
}

// GetRetransmitTime gets the value of RetransmitTime for the instance
func (instance *MSFT_NetIPInterface) GetPropertyRetransmitTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("RetransmitTime")
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

// SetRouterDiscovery sets the value of RouterDiscovery for the instance
func (instance *MSFT_NetIPInterface) SetPropertyRouterDiscovery(value uint8) (err error) {
	return instance.SetProperty("RouterDiscovery", (value))
}

// GetRouterDiscovery gets the value of RouterDiscovery for the instance
func (instance *MSFT_NetIPInterface) GetPropertyRouterDiscovery() (value uint8, err error) {
	retValue, err := instance.GetProperty("RouterDiscovery")
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

// SetStore sets the value of Store for the instance
func (instance *MSFT_NetIPInterface) SetPropertyStore(value uint8) (err error) {
	return instance.SetProperty("Store", (value))
}

// GetStore gets the value of Store for the instance
func (instance *MSFT_NetIPInterface) GetPropertyStore() (value uint8, err error) {
	retValue, err := instance.GetProperty("Store")
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

// SetWeakHostReceive sets the value of WeakHostReceive for the instance
func (instance *MSFT_NetIPInterface) SetPropertyWeakHostReceive(value uint8) (err error) {
	return instance.SetProperty("WeakHostReceive", (value))
}

// GetWeakHostReceive gets the value of WeakHostReceive for the instance
func (instance *MSFT_NetIPInterface) GetPropertyWeakHostReceive() (value uint8, err error) {
	retValue, err := instance.GetProperty("WeakHostReceive")
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

// SetWeakHostSend sets the value of WeakHostSend for the instance
func (instance *MSFT_NetIPInterface) SetPropertyWeakHostSend(value uint8) (err error) {
	return instance.SetProperty("WeakHostSend", (value))
}

// GetWeakHostSend gets the value of WeakHostSend for the instance
func (instance *MSFT_NetIPInterface) GetPropertyWeakHostSend() (value uint8, err error) {
	retValue, err := instance.GetProperty("WeakHostSend")
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
