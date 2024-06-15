// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_BasedOn struct
type Msvm_BasedOn struct {
	*CIM_BasedOn
}

func NewMsvm_BasedOnEx1(instance *cim.WmiInstance) (newInstance *Msvm_BasedOn, err error) {
	tmp, err := NewCIM_BasedOnEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_BasedOn{
		CIM_BasedOn: tmp,
	}
	return
}

func NewMsvm_BasedOnEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_BasedOn, err error) {
	tmp, err := NewCIM_BasedOnEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_BasedOn{
		CIM_BasedOn: tmp,
	}
	return
}
