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

// PS_VpnIPAddressRange struct
type PS_VpnIPAddressRange struct {
	*cim.WmiInstance
}

func NewPS_VpnIPAddressRangeEx1(instance *cim.WmiInstance) (newInstance *PS_VpnIPAddressRange, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_VpnIPAddressRange{
		WmiInstance: tmp,
	}
	return
}

func NewPS_VpnIPAddressRangeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_VpnIPAddressRange, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_VpnIPAddressRange{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="IPAddressRange" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="VpnIPAddressRange "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnIPAddressRange) Add( /* IN */ IPAddressRange []string,
	/* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ RoutingDomain string,
	/* OUT */ cmdletOutput VpnIPAddressRange) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", IPAddressRange, ComputerName, PassThru, RoutingDomain)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="IPAddress" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="VpnIPAddressRange "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnIPAddressRange) Remove( /* IN */ IPAddress string,
	/* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* IN */ RoutingDomain string,
	/* OUT */ cmdletOutput VpnIPAddressRange) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", IPAddress, ComputerName, PassThru, Force, RoutingDomain)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
