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

// MLNX_RegisteredProfile struct
type MLNX_RegisteredProfile struct {
	*CIM_RegisteredProfile
}

func NewMLNX_RegisteredProfileEx1(instance *cim.WmiInstance) (newInstance *MLNX_RegisteredProfile, err error) {
	tmp, err := NewCIM_RegisteredProfileEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_RegisteredProfile{
		CIM_RegisteredProfile: tmp,
	}
	return
}

func NewMLNX_RegisteredProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_RegisteredProfile, err error) {
	tmp, err := NewCIM_RegisteredProfileEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_RegisteredProfile{
		CIM_RegisteredProfile: tmp,
	}
	return
}
