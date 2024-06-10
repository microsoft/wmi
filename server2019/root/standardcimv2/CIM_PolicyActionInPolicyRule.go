// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_PolicyActionInPolicyRule struct
type CIM_PolicyActionInPolicyRule struct {
	*CIM_PolicyActionStructure
}

func NewCIM_PolicyActionInPolicyRuleEx1(instance *cim.WmiInstance) (newInstance *CIM_PolicyActionInPolicyRule, err error) {
	tmp, err := NewCIM_PolicyActionStructureEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PolicyActionInPolicyRule{
		CIM_PolicyActionStructure: tmp,
	}
	return
}

func NewCIM_PolicyActionInPolicyRuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PolicyActionInPolicyRule, err error) {
	tmp, err := NewCIM_PolicyActionStructureEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PolicyActionInPolicyRule{
		CIM_PolicyActionStructure: tmp,
	}
	return
}
