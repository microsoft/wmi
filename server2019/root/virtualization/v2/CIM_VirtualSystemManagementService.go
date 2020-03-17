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

// CIM_VirtualSystemManagementService struct
type CIM_VirtualSystemManagementService struct {
	CIM_Service
}

// Adds resources to a virtual system configuration
///.When applied to a "state" virtual system configuration, as a side effect resources are added to the active virtual system.

// <param name="AffectedConfiguration" type="CIM_VirtualSystemSettingData ">Reference to the affected virtual system configuration.</param>
// <param name="ResourceSettings" type="string []">Array of strings each containing one embedded instance of class CIM_ResourceAllocationSettingData that describes the virtual aspects of a virtual resource to be added to the virtual system.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned. In this case, the instances of class CIM_ResourceAllocationSettingData representing the added resource settings are available via association CIM_ConreteComponent from the instance of class CIM_VirtualSystemSettingData representing the affected virtual system configuration.</param>
// <param name="ResultingResourceSettings" type="CIM_ResourceAllocationSettingData []">Array of references to instances of class CIM_ResourceAllocationSettingData representing virtual aspects of the added virtual resources.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_VirtualSystemManagementService) AddResourceSettings( /* IN */ AffectedConfiguration CIM_VirtualSystemSettingData,
	/* IN */ ResourceSettings []string,
	/* OUT */ ResultingResourceSettings []CIM_ResourceAllocationSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("AddResourceSettings", Action, PercentComplete, Timeout, AffectedConfiguration, ResourceSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Defines a virtual system.
///Input that is not completely specified may be filled out with default values.

// <param name="ReferenceConfiguration" type="CIM_VirtualSystemSettingData ">Reference to an instance of class CIM_VirtualSystemSettingData object that is the top level object of a reference virtual system configuration. The reference configuration is used to complement the configuration of the new virtual system if parameters SystemSettings and ResourceSettings did not provide respective information.</param>
// <param name="ResourceSettings" type="string []">Array of strings each containing an embedded instance of class CIM_ResourceAllocationSettingData that describes the virtual aspects of a virtual resource to be created in the scope of the new virtual system.</param>
// <param name="SystemSettings" type="string ">String containing an embedded instance of class CIM_VirtualSystemSettingData that is used to define attributes of the virtual system to be created.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned. In this case, the instance of class CIM_ComputerSystem representing the new virtual systemis presented via association CIM_AffectedJobElementwith property AffectedElement refering to the new instance of class CIM_ComputerSystem and property ElementEffects set to 5 (Create).</param>
// <param name="ResultingSystem" type="CIM_ComputerSystem ">If a virtual computer system is successfully defined, a reference to an instance of class CIM_ComputerSystem that represents the newly defined virtual computer system is returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_VirtualSystemManagementService) DefineSystem( /* IN */ SystemSettings string,
	/* IN */ ResourceSettings []string,
	/* IN */ ReferenceConfiguration CIM_VirtualSystemSettingData,
	/* OUT */ ResultingSystem CIM_ComputerSystem,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("DefineSystem", Action, PercentComplete, Timeout, SystemSettings, ResourceSettings, ReferenceConfiguration)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Destroys a virtual system.
///The referenced virtual system is destroyed, including any elements scoped by it. Virtual resources are returned to their resource pools, which may imply the destruction of those resources (implementation dependent). If the virtual system is active when the operation is invoked, it is first deactivated and then destroyed. If snapshots were created from the virtual system, these are destroyed as well.

// <param name="AffectedSystem" type="CIM_ComputerSystem ">Reference to an instance of class CIM_ComputerSystem representing the virtual computer system that it to be destroyed.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_VirtualSystemManagementService) DestroySystem( /* IN */ AffectedSystem CIM_ComputerSystem,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("DestroySystem", Action, PercentComplete, Timeout, AffectedSystem)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Modifies virtual resource settings.
///When applied to parts of a "current" virtual system configuration, as a side effect resources of the active virtual system may be modified.

// <param name="ResourceSettings" type="string []">Array of strings each containing an embedded instance of class CIM_ResourceAllocationSettingData that describes modifications to the virtual aspects of an existing virtual resource. All instances must have a valid InstanceID in order to identify the virtual resource setting to be modified.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job be returned. In this case, the instances of class CIM_ResourceAllocationSettingData representing the modified resource settings are available via association CIM_ConreteComponent from the instance of class CIM_VirtualSystemSettingData representing the affected virtual system configuration.</param>
// <param name="ResultingResourceSettings" type="CIM_ResourceAllocationSettingData []">Array of references to instances of class Cim_ResourceAllocationSettingData representing virtual aspects of the modified virtual resources.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_VirtualSystemManagementService) ModifyResourceSettings( /* IN */ ResourceSettings []string,
	/* OUT */ ResultingResourceSettings []CIM_ResourceAllocationSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyResourceSettings", Action, PercentComplete, Timeout, ResourceSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Modifies virtual system settings.
///When applied to the system settings of a "current" virtual system configuration, as a side effect the virtual system instance may be modified.

// <param name="SystemSettings" type="string ">String containing an instance of class CIM_VirtualSystemSettingData that is used to modify the settings of the virtual system. The instance must have a valid InstanceID in order to identify the virtual system setting to be modified.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_VirtualSystemManagementService) ModifySystemSettings( /* IN */ SystemSettings string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifySystemSettings", Action, PercentComplete, Timeout, SystemSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Removes virtual resource settings from a virtual system configuration.
///When applied to parts of a "current" virtual system configuration, as a side effect resources of the active virtual system may be removed.

// <param name="ResourceSettings" type="CIM_ResourceAllocationSettingData []">Array of references to instances of class CIM_ResourceAllocationSettingData where each instance represents the settings of a virtual resource within a virtual system configuration that are to be removed.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job my be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_VirtualSystemManagementService) RemoveResourceSettings( /* IN */ ResourceSettings []CIM_ResourceAllocationSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RemoveResourceSettings", Action, PercentComplete, Timeout, ResourceSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
