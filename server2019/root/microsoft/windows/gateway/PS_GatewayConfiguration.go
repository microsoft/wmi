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

// PS_GatewayConfiguration struct
type PS_GatewayConfiguration struct {
	*cim.WmiInstance
}

func NewPS_GatewayConfigurationEx1(instance *cim.WmiInstance) (newInstance *PS_GatewayConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_GatewayConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewPS_GatewayConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_GatewayConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_GatewayConfiguration{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="RoutingDomainName" type="string "></param>

// <param name="cmdletOutput" type="GatewayConfigurationObject "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_GatewayConfiguration) Get( /* IN */ RoutingDomainName string,
	/* OUT */ cmdletOutput GatewayConfigurationObject) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", RoutingDomainName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RoutingDomainName" type="string "></param>
// <param name="Standby" type="bool "></param>

// <param name="cmdletOutput" type="GatewayConfigurationObject "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_GatewayConfiguration) Set( /* IN */ RoutingDomainName string,
	/* IN */ Standby bool,
	/* OUT */ cmdletOutput GatewayConfigurationObject) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", RoutingDomainName, Standby)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
