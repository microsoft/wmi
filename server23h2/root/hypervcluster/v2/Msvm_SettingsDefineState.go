// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_SettingsDefineState struct
type Msvm_SettingsDefineState struct {
	*CIM_SettingsDefineState
}

func NewMsvm_SettingsDefineStateEx1(instance *cim.WmiInstance) (newInstance *Msvm_SettingsDefineState, err error) {
	tmp, err := NewCIM_SettingsDefineStateEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SettingsDefineState{
		CIM_SettingsDefineState: tmp,
	}
	return
}

func NewMsvm_SettingsDefineStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SettingsDefineState, err error) {
	tmp, err := NewCIM_SettingsDefineStateEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SettingsDefineState{
		CIM_SettingsDefineState: tmp,
	}
	return
}
