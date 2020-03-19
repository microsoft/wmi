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

// Msvm_SyntheticDisplayController struct
type Msvm_SyntheticDisplayController struct {
	*CIM_DisplayController
}

func NewMsvm_SyntheticDisplayControllerEx1(instance *cim.WmiInstance) (newInstance *Msvm_SyntheticDisplayController, err error) {
	tmp, err := NewCIM_DisplayControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticDisplayController{
		CIM_DisplayController: tmp,
	}
	return
}

func NewMsvm_SyntheticDisplayControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SyntheticDisplayController, err error) {
	tmp, err := NewCIM_DisplayControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticDisplayController{
		CIM_DisplayController: tmp,
	}
	return
}

func (instance *Msvm_SyntheticDisplayController) GetRelatedVideoHead() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VideoHead")
}

func (instance *Msvm_SyntheticDisplayController) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_SyntheticDisplayController) GetRelatedSyntheticDisplayControllerSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_SyntheticDisplayControllerSettingData")
}

func (instance *Msvm_SyntheticDisplayController) GetRelatedResourcePool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePool")
}
