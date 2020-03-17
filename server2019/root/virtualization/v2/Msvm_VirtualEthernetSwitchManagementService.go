// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VirtualEthernetSwitchManagementService struct
type Msvm_VirtualEthernetSwitchManagementService struct {
	CIM_VirtualSystemManagementService
}

//

// <param name="AffectedConfiguration" type="CIM_SettingData ">Reference to the affected Ethernet switch port or Ethernet Switch configuration.</param>
// <param name="FeatureSettings" type="string []">Array of strings each containing one embedded instance of class Msvm_FeatureSettingData that describes the feature settings to be added to the switch port configuration.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned. In this case, the instances of class Msvm_FeatureSettingData representing the added feature settings are available via association Msvm_EthernetPortSettingDataComponent from the instance of class Msvm_EthernetPortAllocationSettingData representing the affected switch port.</param>
// <param name="ResultingFeatureSettings" type="Msvm_FeatureSettingData []">Array of references to instances of class Msvm_FeatureSettingData representing the added feature settings.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualEthernetSwitchManagementService) AddFeatureSettings( /* IN */ AffectedConfiguration CIM_SettingData,
	/* IN */ FeatureSettings []string,
	/* OUT */ ResultingFeatureSettings []Msvm_FeatureSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("AddFeatureSettings", Action, PercentComplete, Timeout, AffectedConfiguration, FeatureSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FeatureSettings" type="string []">Array of strings each containing an embedded instance of class Msvm_FeatureSettingData that describes modifications to the current feature settings of an existing Ethernet switch port. All instances must have a valid InstanceID in order to identify the feature settings to be modified.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job be returned. In this case, the instances of class Msvm_FeatureSettingData representing the modified settings are available via association Msvm_EthernetPortSettingDataComponent from the instance of class Msvm_EthernetPortAllocationSettingData representing the affected switch port.</param>
// <param name="ResultingFeatureSettings" type="Msvm_FeatureSettingData []">Array of references to instances of class Msvm_FeatureSettingData representing the modified feature settings.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualEthernetSwitchManagementService) ModifyFeatureSettings( /* IN */ FeatureSettings []string,
	/* OUT */ ResultingFeatureSettings []Msvm_FeatureSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyFeatureSettings", Action, PercentComplete, Timeout, FeatureSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FeatureSettings" type="Msvm_FeatureSettingData []">Array of references to instances of class Msvm_FeatureSettingData that are to be removed.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job my be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualEthernetSwitchManagementService) RemoveFeatureSettings( /* IN */ FeatureSettings []Msvm_FeatureSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RemoveFeatureSettings", Action, PercentComplete, Timeout, FeatureSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_VirtualEthernetSwitchManagementService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_VirtualEthernetSwitchManagementService) GetRelatedVirtualEthernetSwitchManagementCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitchManagementCapabilities")
}
