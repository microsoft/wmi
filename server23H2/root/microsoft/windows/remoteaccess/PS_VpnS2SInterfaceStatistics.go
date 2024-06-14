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

// PS_VpnS2SInterfaceStatistics struct
type PS_VpnS2SInterfaceStatistics struct {
	*cim.WmiInstance
}

func NewPS_VpnS2SInterfaceStatisticsEx1(instance *cim.WmiInstance) (newInstance *PS_VpnS2SInterfaceStatistics, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_VpnS2SInterfaceStatistics{
		WmiInstance: tmp,
	}
	return
}

func NewPS_VpnS2SInterfaceStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_VpnS2SInterfaceStatistics, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_VpnS2SInterfaceStatistics{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnS2SInterfaceStatistics) Clear( /* IN */ Name string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Clear", Name, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Clear" type="bool "></param>
// <param name="Force" type="bool "></param>

// <param name="cmdletOutput" type="VpnS2SInterfaceStatistics []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnS2SInterfaceStatistics) GetByAll( /* IN */ Clear bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput []VpnS2SInterfaceStatistics) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByAll", Clear, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Clear" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="Name" type="string "></param>

// <param name="cmdletOutput" type="VpnS2SInterfaceStatistics "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnS2SInterfaceStatistics) Get( /* IN */ Name string,
	/* IN */ Clear bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput VpnS2SInterfaceStatistics) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Name, Clear, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
