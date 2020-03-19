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

// Msvm_EthernetSwitchPortFeatureSettingData struct
type Msvm_EthernetSwitchPortFeatureSettingData struct {
	*Msvm_FeatureSettingData
}

func NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortFeatureSettingData, err error) {
	tmp, err := NewMsvm_FeatureSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortFeatureSettingData{
		Msvm_FeatureSettingData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchPortFeatureSettingData, err error) {
	tmp, err := NewMsvm_FeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortFeatureSettingData{
		Msvm_FeatureSettingData: tmp,
	}
	return
}

func (instance *Msvm_EthernetSwitchPortFeatureSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
