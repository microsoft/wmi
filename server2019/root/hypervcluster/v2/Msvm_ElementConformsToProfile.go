// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ElementConformsToProfile struct
type Msvm_ElementConformsToProfile struct {
	*CIM_ElementConformsToProfile
}

func NewMsvm_ElementConformsToProfileEx1(instance *cim.WmiInstance) (newInstance *Msvm_ElementConformsToProfile, err error) {
	tmp, err := NewCIM_ElementConformsToProfileEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ElementConformsToProfile{
		CIM_ElementConformsToProfile: tmp,
	}
	return
}

func NewMsvm_ElementConformsToProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ElementConformsToProfile, err error) {
	tmp, err := NewCIM_ElementConformsToProfileEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ElementConformsToProfile{
		CIM_ElementConformsToProfile: tmp,
	}
	return
}
