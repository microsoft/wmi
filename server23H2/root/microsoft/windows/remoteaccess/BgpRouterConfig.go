// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// BgpRouterConfig struct
type BgpRouterConfig struct {
	*cim.WmiInstance

	//
	BgpIdentifier string

	//
	ClientToClientReflection uint32

	//
	ClusterId uint32

	//
	CompareMEDAcrossASN bool

	//
	DefaultGatewayRouting bool

	//
	IPv6Routing uint32

	//
	LocalASN uint32

	//
	LocalIPv6Address string

	//
	PeerName []string

	//
	PolicyName []string

	//
	RouteReflector uint32

	//
	RoutingDomain string

	//
	TransitRouting uint32
}

func NewBgpRouterConfigEx1(instance *cim.WmiInstance) (newInstance *BgpRouterConfig, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BgpRouterConfig{
		WmiInstance: tmp,
	}
	return
}

func NewBgpRouterConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BgpRouterConfig, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BgpRouterConfig{
		WmiInstance: tmp,
	}
	return
}

// SetBgpIdentifier sets the value of BgpIdentifier for the instance
func (instance *BgpRouterConfig) SetPropertyBgpIdentifier(value string) (err error) {
	return instance.SetProperty("BgpIdentifier", (value))
}

// GetBgpIdentifier gets the value of BgpIdentifier for the instance
func (instance *BgpRouterConfig) GetPropertyBgpIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("BgpIdentifier")
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

// SetClientToClientReflection sets the value of ClientToClientReflection for the instance
func (instance *BgpRouterConfig) SetPropertyClientToClientReflection(value uint32) (err error) {
	return instance.SetProperty("ClientToClientReflection", (value))
}

// GetClientToClientReflection gets the value of ClientToClientReflection for the instance
func (instance *BgpRouterConfig) GetPropertyClientToClientReflection() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClientToClientReflection")
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

// SetClusterId sets the value of ClusterId for the instance
func (instance *BgpRouterConfig) SetPropertyClusterId(value uint32) (err error) {
	return instance.SetProperty("ClusterId", (value))
}

// GetClusterId gets the value of ClusterId for the instance
func (instance *BgpRouterConfig) GetPropertyClusterId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusterId")
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

// SetCompareMEDAcrossASN sets the value of CompareMEDAcrossASN for the instance
func (instance *BgpRouterConfig) SetPropertyCompareMEDAcrossASN(value bool) (err error) {
	return instance.SetProperty("CompareMEDAcrossASN", (value))
}

// GetCompareMEDAcrossASN gets the value of CompareMEDAcrossASN for the instance
func (instance *BgpRouterConfig) GetPropertyCompareMEDAcrossASN() (value bool, err error) {
	retValue, err := instance.GetProperty("CompareMEDAcrossASN")
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

// SetDefaultGatewayRouting sets the value of DefaultGatewayRouting for the instance
func (instance *BgpRouterConfig) SetPropertyDefaultGatewayRouting(value bool) (err error) {
	return instance.SetProperty("DefaultGatewayRouting", (value))
}

// GetDefaultGatewayRouting gets the value of DefaultGatewayRouting for the instance
func (instance *BgpRouterConfig) GetPropertyDefaultGatewayRouting() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultGatewayRouting")
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

// SetIPv6Routing sets the value of IPv6Routing for the instance
func (instance *BgpRouterConfig) SetPropertyIPv6Routing(value uint32) (err error) {
	return instance.SetProperty("IPv6Routing", (value))
}

// GetIPv6Routing gets the value of IPv6Routing for the instance
func (instance *BgpRouterConfig) GetPropertyIPv6Routing() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6Routing")
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

// SetLocalASN sets the value of LocalASN for the instance
func (instance *BgpRouterConfig) SetPropertyLocalASN(value uint32) (err error) {
	return instance.SetProperty("LocalASN", (value))
}

// GetLocalASN gets the value of LocalASN for the instance
func (instance *BgpRouterConfig) GetPropertyLocalASN() (value uint32, err error) {
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

// SetLocalIPv6Address sets the value of LocalIPv6Address for the instance
func (instance *BgpRouterConfig) SetPropertyLocalIPv6Address(value string) (err error) {
	return instance.SetProperty("LocalIPv6Address", (value))
}

// GetLocalIPv6Address gets the value of LocalIPv6Address for the instance
func (instance *BgpRouterConfig) GetPropertyLocalIPv6Address() (value string, err error) {
	retValue, err := instance.GetProperty("LocalIPv6Address")
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
func (instance *BgpRouterConfig) SetPropertyPeerName(value []string) (err error) {
	return instance.SetProperty("PeerName", (value))
}

// GetPeerName gets the value of PeerName for the instance
func (instance *BgpRouterConfig) GetPropertyPeerName() (value []string, err error) {
	retValue, err := instance.GetProperty("PeerName")
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

// SetPolicyName sets the value of PolicyName for the instance
func (instance *BgpRouterConfig) SetPropertyPolicyName(value []string) (err error) {
	return instance.SetProperty("PolicyName", (value))
}

// GetPolicyName gets the value of PolicyName for the instance
func (instance *BgpRouterConfig) GetPropertyPolicyName() (value []string, err error) {
	retValue, err := instance.GetProperty("PolicyName")
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

// SetRouteReflector sets the value of RouteReflector for the instance
func (instance *BgpRouterConfig) SetPropertyRouteReflector(value uint32) (err error) {
	return instance.SetProperty("RouteReflector", (value))
}

// GetRouteReflector gets the value of RouteReflector for the instance
func (instance *BgpRouterConfig) GetPropertyRouteReflector() (value uint32, err error) {
	retValue, err := instance.GetProperty("RouteReflector")
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

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *BgpRouterConfig) SetPropertyRoutingDomain(value string) (err error) {
	return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *BgpRouterConfig) GetPropertyRoutingDomain() (value string, err error) {
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

// SetTransitRouting sets the value of TransitRouting for the instance
func (instance *BgpRouterConfig) SetPropertyTransitRouting(value uint32) (err error) {
	return instance.SetProperty("TransitRouting", (value))
}

// GetTransitRouting gets the value of TransitRouting for the instance
func (instance *BgpRouterConfig) GetPropertyTransitRouting() (value uint32, err error) {
	retValue, err := instance.GetProperty("TransitRouting")
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
