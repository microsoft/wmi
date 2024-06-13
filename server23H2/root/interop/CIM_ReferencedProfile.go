// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Interop
//////////////////////////////////////////////
package interop

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ReferencedProfile struct
type CIM_ReferencedProfile struct {
	*CIM_Dependency
}

func NewCIM_ReferencedProfileEx1(instance *cim.WmiInstance) (newInstance *CIM_ReferencedProfile, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ReferencedProfile{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_ReferencedProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ReferencedProfile, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ReferencedProfile{
		CIM_Dependency: tmp,
	}
	return
}
