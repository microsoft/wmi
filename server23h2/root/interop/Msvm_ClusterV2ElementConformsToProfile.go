// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Interop
//////////////////////////////////////////////
package interop

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ClusterV2ElementConformsToProfile struct
type Msvm_ClusterV2ElementConformsToProfile struct {
	*Msvm_ElementConformsToProfile
}

func NewMsvm_ClusterV2ElementConformsToProfileEx1(instance *cim.WmiInstance) (newInstance *Msvm_ClusterV2ElementConformsToProfile, err error) {
	tmp, err := NewMsvm_ElementConformsToProfileEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ClusterV2ElementConformsToProfile{
		Msvm_ElementConformsToProfile: tmp,
	}
	return
}

func NewMsvm_ClusterV2ElementConformsToProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ClusterV2ElementConformsToProfile, err error) {
	tmp, err := NewMsvm_ElementConformsToProfileEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ClusterV2ElementConformsToProfile{
		Msvm_ElementConformsToProfile: tmp,
	}
	return
}
