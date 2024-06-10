// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.GatewayHealthMonitor
//
// ////////////////////////////////////////////
package gatewayhealthmonitor

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_GatewayHealthMonitor struct
type MSFT_GatewayHealthMonitor struct {
	*cim.WmiInstance
}

func NewMSFT_GatewayHealthMonitorEx1(instance *cim.WmiInstance) (newInstance *MSFT_GatewayHealthMonitor, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_GatewayHealthMonitor{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_GatewayHealthMonitorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_GatewayHealthMonitor, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_GatewayHealthMonitor{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="InterfaceList" type="string []"></param>
// <param name="ServiceList" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_GatewayHealthMonitor) StartGatewayMonitoring( /* IN */ InterfaceList []string,
	/* IN */ ServiceList []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StartGatewayMonitoring", InterfaceList, ServiceList)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_GatewayHealthMonitor) StopGatewayMonitoring() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StopGatewayMonitoring")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
