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

// Msvm_ReferencedProfile struct
type Msvm_ReferencedProfile struct {
	*CIM_ReferencedProfile
}

func NewMsvm_ReferencedProfileEx1(instance *cim.WmiInstance) (newInstance *Msvm_ReferencedProfile, err error) {
	tmp, err := NewCIM_ReferencedProfileEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReferencedProfile{
		CIM_ReferencedProfile: tmp,
	}
	return
}

func NewMsvm_ReferencedProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ReferencedProfile, err error) {
	tmp, err := NewCIM_ReferencedProfileEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReferencedProfile{
		CIM_ReferencedProfile: tmp,
	}
	return
}
