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

// PS_VpnIPAddressAssignment struct
type PS_VpnIPAddressAssignment struct {
	*cim.WmiInstance
}

func NewPS_VpnIPAddressAssignmentEx1(instance *cim.WmiInstance) (newInstance *PS_VpnIPAddressAssignment, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_VpnIPAddressAssignment{
		WmiInstance: tmp,
	}
	return
}

func NewPS_VpnIPAddressAssignmentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_VpnIPAddressAssignment, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_VpnIPAddressAssignment{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="IPAddressRange" type="string []"></param>
// <param name="IPAssignmentMethod" type="string "></param>
// <param name="IPv6Prefix" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnIPAddressAssignment "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnIPAddressAssignment) SetByVpnIPAddressAssignment( /* IN */ PassThru bool,
	/* IN */ IPAddressRange []string,
	/* IN */ ComputerName string,
	/* IN */ IPAssignmentMethod string,
	/* IN */ IPv6Prefix string,
	/* OUT */ cmdletOutput VpnIPAddressAssignment) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByVpnIPAddressAssignment", PassThru, IPAddressRange, ComputerName, IPAssignmentMethod, IPv6Prefix)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="EntrypointName" type="string "></param>
// <param name="IPv6Prefix" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnIPAddressAssignment "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnIPAddressAssignment) SetByVpnIPv6PrefixEntrypoint( /* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ EntrypointName string,
	/* IN */ IPv6Prefix string,
	/* OUT */ cmdletOutput VpnIPAddressAssignment) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByVpnIPv6PrefixEntrypoint", ComputerName, PassThru, EntrypointName, IPv6Prefix)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
