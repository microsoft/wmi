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
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_BgpStatistics struct
type PS_BgpStatistics struct {
	*cim.WmiInstance
}

func NewPS_BgpStatisticsEx1(instance *cim.WmiInstance) (newInstance *PS_BgpStatistics, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_BgpStatistics{
		WmiInstance: tmp,
	}
	return
}

func NewPS_BgpStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_BgpStatistics, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_BgpStatistics{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="PeerName" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="BgpStatistics []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpStatistics) Get( /* IN */ RoutingDomain string,
	/* IN */ PeerName []string,
	/* OUT */ cmdletOutput []BgpStatistics) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", RoutingDomain, PeerName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllRoutingDomains" type="bool "></param>

// <param name="cmdletOutput" type="BgpStatistics []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpStatistics) GetForAllRoutingDomains( /* IN */ AllRoutingDomains bool,
	/* OUT */ cmdletOutput []BgpStatistics) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetForAllRoutingDomains", AllRoutingDomains)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
