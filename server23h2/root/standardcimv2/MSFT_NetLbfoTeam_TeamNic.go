// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetLbfoTeam_TeamNic struct
type MSFT_NetLbfoTeam_TeamNic struct {
	*CIM_Component
}

func NewMSFT_NetLbfoTeam_TeamNicEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetLbfoTeam_TeamNic, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLbfoTeam_TeamNic{
		CIM_Component: tmp,
	}
	return
}

func NewMSFT_NetLbfoTeam_TeamNicEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetLbfoTeam_TeamNic, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLbfoTeam_TeamNic{
		CIM_Component: tmp,
	}
	return
}
