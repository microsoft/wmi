// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Gateway
//////////////////////////////////////////////
package gateway

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// GatewayTunnelStatistics struct
type GatewayTunnelStatistics struct {
	*cim.WmiInstance

	//
	BytesIn uint64

	//
	BytesOut uint64

	//
	ConnectedDuration uint32

	//
	ConnectionGuid string

	//
	IsConnected bool

	//
	PacketsIn uint64

	//
	PacketsOut uint64

	//
	RoutingDomainName string

	//
	TunnelName string
}

func NewGatewayTunnelStatisticsEx1(instance *cim.WmiInstance) (newInstance *GatewayTunnelStatistics, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &GatewayTunnelStatistics{
		WmiInstance: tmp,
	}
	return
}

func NewGatewayTunnelStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *GatewayTunnelStatistics, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &GatewayTunnelStatistics{
		WmiInstance: tmp,
	}
	return
}

// SetBytesIn sets the value of BytesIn for the instance
func (instance *GatewayTunnelStatistics) SetPropertyBytesIn(value uint64) (err error) {
	return instance.SetProperty("BytesIn", (value))
}

// GetBytesIn gets the value of BytesIn for the instance
func (instance *GatewayTunnelStatistics) GetPropertyBytesIn() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesIn")
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

// SetBytesOut sets the value of BytesOut for the instance
func (instance *GatewayTunnelStatistics) SetPropertyBytesOut(value uint64) (err error) {
	return instance.SetProperty("BytesOut", (value))
}

// GetBytesOut gets the value of BytesOut for the instance
func (instance *GatewayTunnelStatistics) GetPropertyBytesOut() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesOut")
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

// SetConnectedDuration sets the value of ConnectedDuration for the instance
func (instance *GatewayTunnelStatistics) SetPropertyConnectedDuration(value uint32) (err error) {
	return instance.SetProperty("ConnectedDuration", (value))
}

// GetConnectedDuration gets the value of ConnectedDuration for the instance
func (instance *GatewayTunnelStatistics) GetPropertyConnectedDuration() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectedDuration")
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

// SetConnectionGuid sets the value of ConnectionGuid for the instance
func (instance *GatewayTunnelStatistics) SetPropertyConnectionGuid(value string) (err error) {
	return instance.SetProperty("ConnectionGuid", (value))
}

// GetConnectionGuid gets the value of ConnectionGuid for the instance
func (instance *GatewayTunnelStatistics) GetPropertyConnectionGuid() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionGuid")
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

// SetIsConnected sets the value of IsConnected for the instance
func (instance *GatewayTunnelStatistics) SetPropertyIsConnected(value bool) (err error) {
	return instance.SetProperty("IsConnected", (value))
}

// GetIsConnected gets the value of IsConnected for the instance
func (instance *GatewayTunnelStatistics) GetPropertyIsConnected() (value bool, err error) {
	retValue, err := instance.GetProperty("IsConnected")
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

// SetPacketsIn sets the value of PacketsIn for the instance
func (instance *GatewayTunnelStatistics) SetPropertyPacketsIn(value uint64) (err error) {
	return instance.SetProperty("PacketsIn", (value))
}

// GetPacketsIn gets the value of PacketsIn for the instance
func (instance *GatewayTunnelStatistics) GetPropertyPacketsIn() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsIn")
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

// SetPacketsOut sets the value of PacketsOut for the instance
func (instance *GatewayTunnelStatistics) SetPropertyPacketsOut(value uint64) (err error) {
	return instance.SetProperty("PacketsOut", (value))
}

// GetPacketsOut gets the value of PacketsOut for the instance
func (instance *GatewayTunnelStatistics) GetPropertyPacketsOut() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsOut")
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

// SetRoutingDomainName sets the value of RoutingDomainName for the instance
func (instance *GatewayTunnelStatistics) SetPropertyRoutingDomainName(value string) (err error) {
	return instance.SetProperty("RoutingDomainName", (value))
}

// GetRoutingDomainName gets the value of RoutingDomainName for the instance
func (instance *GatewayTunnelStatistics) GetPropertyRoutingDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomainName")
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

// SetTunnelName sets the value of TunnelName for the instance
func (instance *GatewayTunnelStatistics) SetPropertyTunnelName(value string) (err error) {
	return instance.SetProperty("TunnelName", (value))
}

// GetTunnelName gets the value of TunnelName for the instance
func (instance *GatewayTunnelStatistics) GetPropertyTunnelName() (value string, err error) {
	retValue, err := instance.GetProperty("TunnelName")
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

//

// <param name="RoutingDomainName" type="string "></param>
// <param name="TunnelName" type="string "></param>

// <param name="cmdletOutput" type="GatewayTunnelStatistics []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *GatewayTunnelStatistics) Get( /* IN */ TunnelName string,
	/* IN */ RoutingDomainName string,
	/* OUT */ cmdletOutput []GatewayTunnelStatistics) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", TunnelName, RoutingDomainName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
