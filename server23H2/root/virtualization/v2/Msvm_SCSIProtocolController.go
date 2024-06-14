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

// Msvm_SCSIProtocolController struct
type Msvm_SCSIProtocolController struct {
	*CIM_SCSIProtocolController
}

func NewMsvm_SCSIProtocolControllerEx1(instance *cim.WmiInstance) (newInstance *Msvm_SCSIProtocolController, err error) {
	tmp, err := NewCIM_SCSIProtocolControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SCSIProtocolController{
		CIM_SCSIProtocolController: tmp,
	}
	return
}

func NewMsvm_SCSIProtocolControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SCSIProtocolController, err error) {
	tmp, err := NewCIM_SCSIProtocolControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SCSIProtocolController{
		CIM_SCSIProtocolController: tmp,
	}
	return
}
