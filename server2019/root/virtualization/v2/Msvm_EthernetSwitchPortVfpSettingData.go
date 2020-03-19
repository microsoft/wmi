// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_EthernetSwitchPortVfpSettingData struct
type Msvm_EthernetSwitchPortVfpSettingData struct {
	*Msvm_EthernetSwitchPortFeatureSettingData
}

func NewMsvm_EthernetSwitchPortVfpSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortVfpSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortVfpSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchPortVfpSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchPortVfpSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortVfpSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

func (instance *Msvm_EthernetSwitchPortVfpSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
