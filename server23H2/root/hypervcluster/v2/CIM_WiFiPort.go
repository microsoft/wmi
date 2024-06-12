// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_WiFiPort struct
type CIM_WiFiPort struct {
	*CIM_NetworkPort
}

func NewCIM_WiFiPortEx1(instance *cim.WmiInstance) (newInstance *CIM_WiFiPort, err error) {
	tmp, err := NewCIM_NetworkPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_WiFiPort{
		CIM_NetworkPort: tmp,
	}
	return
}

func NewCIM_WiFiPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_WiFiPort, err error) {
	tmp, err := NewCIM_NetworkPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_WiFiPort{
		CIM_NetworkPort: tmp,
	}
	return
}
