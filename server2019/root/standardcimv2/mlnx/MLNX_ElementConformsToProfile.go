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

// MLNX_ElementConformsToProfile struct
type MLNX_ElementConformsToProfile struct {
	*CIM_ElementConformsToProfile
}

func NewMLNX_ElementConformsToProfileEx1(instance *cim.WmiInstance) (newInstance *MLNX_ElementConformsToProfile, err error) {
	tmp, err := NewCIM_ElementConformsToProfileEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_ElementConformsToProfile{
		CIM_ElementConformsToProfile: tmp,
	}
	return
}

func NewMLNX_ElementConformsToProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_ElementConformsToProfile, err error) {
	tmp, err := NewCIM_ElementConformsToProfileEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_ElementConformsToProfile{
		CIM_ElementConformsToProfile: tmp,
	}
	return
}
