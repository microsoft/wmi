// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.Gateway
//
// ////////////////////////////////////////////
package gateway

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_GatewayTunnelPacketTrace struct
type PS_GatewayTunnelPacketTrace struct {
	*cim.WmiInstance
}

func NewPS_GatewayTunnelPacketTraceEx1(instance *cim.WmiInstance) (newInstance *PS_GatewayTunnelPacketTrace, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_GatewayTunnelPacketTrace{
		WmiInstance: tmp,
	}
	return
}

func NewPS_GatewayTunnelPacketTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_GatewayTunnelPacketTrace, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_GatewayTunnelPacketTrace{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="TunnelName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_GatewayTunnelPacketTrace) Disable( /* IN */ TunnelName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable", TunnelName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="TunnelName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_GatewayTunnelPacketTrace) Enable( /* IN */ TunnelName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable", TunnelName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
