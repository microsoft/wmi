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

// Msvm_FeatureSettingsDefineCapabilities struct
type Msvm_FeatureSettingsDefineCapabilities struct {
	*CIM_SettingsDefineCapabilities
}

func NewMsvm_FeatureSettingsDefineCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Msvm_FeatureSettingsDefineCapabilities, err error) {
	tmp, err := NewCIM_SettingsDefineCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_FeatureSettingsDefineCapabilities{
		CIM_SettingsDefineCapabilities: tmp,
	}
	return
}

func NewMsvm_FeatureSettingsDefineCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_FeatureSettingsDefineCapabilities, err error) {
	tmp, err := NewCIM_SettingsDefineCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_FeatureSettingsDefineCapabilities{
		CIM_SettingsDefineCapabilities: tmp,
	}
	return
}
