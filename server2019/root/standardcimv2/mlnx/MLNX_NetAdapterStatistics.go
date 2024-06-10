// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_NetAdapterStatistics struct
type MLNX_NetAdapterStatistics struct {
	*CIM_NetworkPortStatistics
}

func NewMLNX_NetAdapterStatisticsEx1(instance *cim.WmiInstance) (newInstance *MLNX_NetAdapterStatistics, err error) {
	tmp, err := NewCIM_NetworkPortStatisticsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_NetAdapterStatistics{
		CIM_NetworkPortStatistics: tmp,
	}
	return
}

func NewMLNX_NetAdapterStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_NetAdapterStatistics, err error) {
	tmp, err := NewCIM_NetworkPortStatisticsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_NetAdapterStatistics{
		CIM_NetworkPortStatistics: tmp,
	}
	return
}
