// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VideoHead struct
type Msvm_VideoHead struct {
	*CIM_VideoHead
}

func NewMsvm_VideoHeadEx1(instance *cim.WmiInstance) (newInstance *Msvm_VideoHead, err error) {
	tmp, err := NewCIM_VideoHeadEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VideoHead{
		CIM_VideoHead: tmp,
	}
	return
}

func NewMsvm_VideoHeadEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VideoHead, err error) {
	tmp, err := NewCIM_VideoHeadEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VideoHead{
		CIM_VideoHead: tmp,
	}
	return
}
