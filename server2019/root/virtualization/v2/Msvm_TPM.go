// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_TPM struct
type Msvm_TPM struct {
	*CIM_TPM
}

func NewMsvm_TPMEx1(instance *cim.WmiInstance) (newInstance *Msvm_TPM, err error) {
	tmp, err := NewCIM_TPMEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_TPM{
		CIM_TPM: tmp,
	}
	return
}

func NewMsvm_TPMEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_TPM, err error) {
	tmp, err := NewCIM_TPMEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_TPM{
		CIM_TPM: tmp,
	}
	return
}

func (instance *Msvm_TPM) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_TPM) GetRelatedSecuritySettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_SecuritySettingData")
}
