// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// BgpPeerConfig struct
type BgpPeerConfig struct {
	*cim.WmiInstance

	//
	ConnectivityStatus uint32

	//
	EgressPolicyList []string

	//
	GracefulRestart bool

	//
	GracefulRestartTime uint16

	//
	HoldTimeSec uint16

	//
	IdleHoldTimeSec uint16

	//
	IngressPolicyList []string

	//
	LocalASN uint32

	//
	LocalIPAddress string

	//
	MaxAllowedPrefix uint32

	//
	OperationMode uint32

	//
	PeerASN uint32

	//
	PeeringMode uint32

	//
	PeerIPAddress string

	//
	PeerName string

	//
	RouteReflectorClient bool

	//
	RoutingDomain string

	//
	Weight uint16
}

func NewBgpPeerConfigEx1(instance *cim.WmiInstance) (newInstance *BgpPeerConfig, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BgpPeerConfig{
		WmiInstance: tmp,
	}
	return
}

func NewBgpPeerConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BgpPeerConfig, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BgpPeerConfig{
		WmiInstance: tmp,
	}
	return
}

// SetConnectivityStatus sets the value of ConnectivityStatus for the instance
func (instance *BgpPeerConfig) SetPropertyConnectivityStatus(value uint32) (err error) {
	return instance.SetProperty("ConnectivityStatus", (value))
}

// GetConnectivityStatus gets the value of ConnectivityStatus for the instance
func (instance *BgpPeerConfig) GetPropertyConnectivityStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectivityStatus")
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

// SetEgressPolicyList sets the value of EgressPolicyList for the instance
func (instance *BgpPeerConfig) SetPropertyEgressPolicyList(value []string) (err error) {
	return instance.SetProperty("EgressPolicyList", (value))
}

// GetEgressPolicyList gets the value of EgressPolicyList for the instance
func (instance *BgpPeerConfig) GetPropertyEgressPolicyList() (value []string, err error) {
	retValue, err := instance.GetProperty("EgressPolicyList")
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

// SetGracefulRestart sets the value of GracefulRestart for the instance
func (instance *BgpPeerConfig) SetPropertyGracefulRestart(value bool) (err error) {
	return instance.SetProperty("GracefulRestart", (value))
}

// GetGracefulRestart gets the value of GracefulRestart for the instance
func (instance *BgpPeerConfig) GetPropertyGracefulRestart() (value bool, err error) {
	retValue, err := instance.GetProperty("GracefulRestart")
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

// SetGracefulRestartTime sets the value of GracefulRestartTime for the instance
func (instance *BgpPeerConfig) SetPropertyGracefulRestartTime(value uint16) (err error) {
	return instance.SetProperty("GracefulRestartTime", (value))
}

// GetGracefulRestartTime gets the value of GracefulRestartTime for the instance
func (instance *BgpPeerConfig) GetPropertyGracefulRestartTime() (value uint16, err error) {
	retValue, err := instance.GetProperty("GracefulRestartTime")
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

// SetHoldTimeSec sets the value of HoldTimeSec for the instance
func (instance *BgpPeerConfig) SetPropertyHoldTimeSec(value uint16) (err error) {
	return instance.SetProperty("HoldTimeSec", (value))
}

// GetHoldTimeSec gets the value of HoldTimeSec for the instance
func (instance *BgpPeerConfig) GetPropertyHoldTimeSec() (value uint16, err error) {
	retValue, err := instance.GetProperty("HoldTimeSec")
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

// SetIdleHoldTimeSec sets the value of IdleHoldTimeSec for the instance
func (instance *BgpPeerConfig) SetPropertyIdleHoldTimeSec(value uint16) (err error) {
	return instance.SetProperty("IdleHoldTimeSec", (value))
}

// GetIdleHoldTimeSec gets the value of IdleHoldTimeSec for the instance
func (instance *BgpPeerConfig) GetPropertyIdleHoldTimeSec() (value uint16, err error) {
	retValue, err := instance.GetProperty("IdleHoldTimeSec")
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

// SetIngressPolicyList sets the value of IngressPolicyList for the instance
func (instance *BgpPeerConfig) SetPropertyIngressPolicyList(value []string) (err error) {
	return instance.SetProperty("IngressPolicyList", (value))
}

// GetIngressPolicyList gets the value of IngressPolicyList for the instance
func (instance *BgpPeerConfig) GetPropertyIngressPolicyList() (value []string, err error) {
	retValue, err := instance.GetProperty("IngressPolicyList")
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

// SetLocalASN sets the value of LocalASN for the instance
func (instance *BgpPeerConfig) SetPropertyLocalASN(value uint32) (err error) {
	return instance.SetProperty("LocalASN", (value))
}

// GetLocalASN gets the value of LocalASN for the instance
func (instance *BgpPeerConfig) GetPropertyLocalASN() (value uint32, err error) {
	retValue, err := instance.GetProperty("LocalASN")
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

// SetLocalIPAddress sets the value of LocalIPAddress for the instance
func (instance *BgpPeerConfig) SetPropertyLocalIPAddress(value string) (err error) {
	return instance.SetProperty("LocalIPAddress", (value))
}

// GetLocalIPAddress gets the value of LocalIPAddress for the instance
func (instance *BgpPeerConfig) GetPropertyLocalIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("LocalIPAddress")
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

// SetMaxAllowedPrefix sets the value of MaxAllowedPrefix for the instance
func (instance *BgpPeerConfig) SetPropertyMaxAllowedPrefix(value uint32) (err error) {
	return instance.SetProperty("MaxAllowedPrefix", (value))
}

// GetMaxAllowedPrefix gets the value of MaxAllowedPrefix for the instance
func (instance *BgpPeerConfig) GetPropertyMaxAllowedPrefix() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxAllowedPrefix")
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

// SetOperationMode sets the value of OperationMode for the instance
func (instance *BgpPeerConfig) SetPropertyOperationMode(value uint32) (err error) {
	return instance.SetProperty("OperationMode", (value))
}

// GetOperationMode gets the value of OperationMode for the instance
func (instance *BgpPeerConfig) GetPropertyOperationMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("OperationMode")
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

// SetPeerASN sets the value of PeerASN for the instance
func (instance *BgpPeerConfig) SetPropertyPeerASN(value uint32) (err error) {
	return instance.SetProperty("PeerASN", (value))
}

// GetPeerASN gets the value of PeerASN for the instance
func (instance *BgpPeerConfig) GetPropertyPeerASN() (value uint32, err error) {
	retValue, err := instance.GetProperty("PeerASN")
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

// SetPeeringMode sets the value of PeeringMode for the instance
func (instance *BgpPeerConfig) SetPropertyPeeringMode(value uint32) (err error) {
	return instance.SetProperty("PeeringMode", (value))
}

// GetPeeringMode gets the value of PeeringMode for the instance
func (instance *BgpPeerConfig) GetPropertyPeeringMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("PeeringMode")
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

// SetPeerIPAddress sets the value of PeerIPAddress for the instance
func (instance *BgpPeerConfig) SetPropertyPeerIPAddress(value string) (err error) {
	return instance.SetProperty("PeerIPAddress", (value))
}

// GetPeerIPAddress gets the value of PeerIPAddress for the instance
func (instance *BgpPeerConfig) GetPropertyPeerIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("PeerIPAddress")
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

// SetPeerName sets the value of PeerName for the instance
func (instance *BgpPeerConfig) SetPropertyPeerName(value string) (err error) {
	return instance.SetProperty("PeerName", (value))
}

// GetPeerName gets the value of PeerName for the instance
func (instance *BgpPeerConfig) GetPropertyPeerName() (value string, err error) {
	retValue, err := instance.GetProperty("PeerName")
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

// SetRouteReflectorClient sets the value of RouteReflectorClient for the instance
func (instance *BgpPeerConfig) SetPropertyRouteReflectorClient(value bool) (err error) {
	return instance.SetProperty("RouteReflectorClient", (value))
}

// GetRouteReflectorClient gets the value of RouteReflectorClient for the instance
func (instance *BgpPeerConfig) GetPropertyRouteReflectorClient() (value bool, err error) {
	retValue, err := instance.GetProperty("RouteReflectorClient")
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

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *BgpPeerConfig) SetPropertyRoutingDomain(value string) (err error) {
	return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *BgpPeerConfig) GetPropertyRoutingDomain() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomain")
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

// SetWeight sets the value of Weight for the instance
func (instance *BgpPeerConfig) SetPropertyWeight(value uint16) (err error) {
	return instance.SetProperty("Weight", (value))
}

// GetWeight gets the value of Weight for the instance
func (instance *BgpPeerConfig) GetPropertyWeight() (value uint16, err error) {
	retValue, err := instance.GetProperty("Weight")
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
