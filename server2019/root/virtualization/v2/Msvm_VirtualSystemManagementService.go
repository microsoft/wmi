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

// Msvm_VirtualSystemManagementService struct
type Msvm_VirtualSystemManagementService struct {
	*CIM_VirtualSystemManagementService
}

func NewMsvm_VirtualSystemManagementServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemManagementService, err error) {
	tmp, err := NewCIM_VirtualSystemManagementServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemManagementService{
		CIM_VirtualSystemManagementService: tmp,
	}
	return
}

func NewMsvm_VirtualSystemManagementServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemManagementService, err error) {
	tmp, err := NewCIM_VirtualSystemManagementServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemManagementService{
		CIM_VirtualSystemManagementService: tmp,
	}
	return
}

//

// <param name="ReferenceConfiguration" type="CIM_VirtualSystemSettingData ">Reference to an instance of class CIM_VirtualSystemSettingData object that is the top level object of a reference virtual system configuration. The reference configuration is used to complement the configuration of the new virtual system if parameters SystemSettings and ResourceSettings did not provide respective information.</param>
// <param name="ResourceSettings" type="string []">Array of strings each containing an embedded instance of class CIM_ResourceAllocationSettingData that describes the virtual aspects of a virtual resource to be created in the scope of the new virtual system.</param>
// <param name="SystemSettings" type="string ">String containing an embedded instance of class CIM_VirtualSystemSettingData that is used to define attributes of the virtual system to be created.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned. In this case, the instance of class CIM_ComputerSystem representing the new virtual systemis presented via association CIM_AffectedJobElementwith property AffectedElement refering to the new instance of class CIM_ComputerSystem and property ElementEffects set to 5 (Create).</param>
// <param name="ResultingSystem" type="CIM_ComputerSystem ">If a virtual computer system is successfully defined, a reference to an instance of class CIM_ComputerSystem that represents the newly defined virtual computer system is returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) DefinePlannedSystem( /* IN */ SystemSettings string,
	/* IN */ ResourceSettings []string,
	/* IN */ ReferenceConfiguration CIM_VirtualSystemSettingData,
	/* OUT */ ResultingSystem CIM_ComputerSystem,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("DefinePlannedSystem", Action, PercentComplete, Timeout, SystemSettings, ResourceSettings, ReferenceConfiguration)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PlannedSystem" type="Msvm_PlannedComputerSystem "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) ValidatePlannedSystem( /* IN */ PlannedSystem Msvm_PlannedComputerSystem,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ValidatePlannedSystem", Action, PercentComplete, Timeout, PlannedSystem)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="UpgradeSettingData" type="string ">String containing an instance of class CIM_SettingData that is used to upgrade the virtual system.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) UpgradeSystemVersion( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ UpgradeSettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("UpgradeSystemVersion", Action, PercentComplete, Timeout, ComputerSystem, UpgradeSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="GenerateNewSystemIdentifier" type="bool "></param>
// <param name="SnapshotFolder" type="string "></param>
// <param name="SystemDefinitionFile" type="string "></param>

// <param name="ImportedSystem" type="Msvm_PlannedComputerSystem "></param>
// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) ImportSystemDefinition( /* IN */ SystemDefinitionFile string,
	/* IN */ SnapshotFolder string,
	/* IN */ GenerateNewSystemIdentifier bool,
	/* OUT */ ImportedSystem Msvm_PlannedComputerSystem,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ImportSystemDefinition", Action, PercentComplete, Timeout, SystemDefinitionFile, SnapshotFolder, GenerateNewSystemIdentifier)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PlannedSystem" type="Msvm_PlannedComputerSystem "></param>
// <param name="SnapshotFolder" type="string "></param>

// <param name="ImportedSnapshots" type="Msvm_VirtualSystemSettingData []"></param>
// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) ImportSnapshotDefinitions( /* IN */ PlannedSystem Msvm_PlannedComputerSystem,
	/* IN */ SnapshotFolder string,
	/* OUT */ ImportedSnapshots []Msvm_VirtualSystemSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ImportSnapshotDefinitions", Action, PercentComplete, Timeout, PlannedSystem, SnapshotFolder)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PlannedSystem" type="Msvm_PlannedComputerSystem "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ResultingSystem" type="CIM_ComputerSystem "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) RealizePlannedSystem( /* IN */ PlannedSystem Msvm_PlannedComputerSystem,
	/* OUT */ ResultingSystem CIM_ComputerSystem,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RealizePlannedSystem", Action, PercentComplete, Timeout, PlannedSystem)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="ExportDirectory" type="string "></param>
// <param name="ExportSettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) ExportSystemDefinition( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ ExportDirectory string,
	/* IN */ ExportSettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ExportSystemDefinition", Action, PercentComplete, Timeout, ComputerSystem, ExportDirectory, ExportSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AffectedConfiguration" type="Msvm_EthernetPortAllocationSettingData ">Reference to the affected Ethernet connection.</param>
// <param name="FeatureSettings" type="string []">Array of strings each containing one embedded instance of class Msvm_EthernetSwitchPortFeatureSettingData that describes the feature settings to be added to the connection configuration.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned. In this case, the instances of class Msvm_EthernetSwitchPortFeatureSettingData representing the added feature settings are available via association Msvm_EthernetPortSettingDataComponent from the instance of class Msvm_EthernetPortAllocationSettingData representing the affected switch port.</param>
// <param name="ResultingFeatureSettings" type="Msvm_EthernetSwitchPortFeatureSettingData []">Array of references to instances of class Msvm_EthernetSwitchPortFeatureSettingData representing the added feature settings.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) AddFeatureSettings( /* IN */ AffectedConfiguration Msvm_EthernetPortAllocationSettingData,
	/* IN */ FeatureSettings []string,
	/* OUT */ ResultingFeatureSettings []Msvm_EthernetSwitchPortFeatureSettingData,
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

// <param name="FeatureSettings" type="string []">Array of strings each containing an embedded instance of class Msvm_EthernetSwitchPortFeatureSettingData that describes modifications to the current feature settings of an existing Ethernet connection. All instances must have a valid InstanceID in order to identify the feature settings to be modified.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job be returned. In this case, the instances of class Msvm_EthernetSwitchPortFeatureSettingData representing the modified settings are available via association Msvm_EthernetPortSettingDataComponent from the instance of class Msvm_EthernetPortAllocationSettingData representing the affected switch port.</param>
// <param name="ResultingFeatureSettings" type="Msvm_EthernetSwitchPortFeatureSettingData []">Array of references to instances of class Msvm_EthernetSwitchPortFeatureSettingData representing the modified feature settings.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) ModifyFeatureSettings( /* IN */ FeatureSettings []string,
	/* OUT */ ResultingFeatureSettings []Msvm_EthernetSwitchPortFeatureSettingData,
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

// <param name="FeatureSettings" type="Msvm_EthernetSwitchPortFeatureSettingData []">Array of references to instances of class Msvm_EthernetSwitchPortFeatureSettingData that are to be removed.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) RemoveFeatureSettings( /* IN */ FeatureSettings []Msvm_EthernetSwitchPortFeatureSettingData,
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

//

// <param name="AffectedConfiguration" type="CIM_VirtualSystemSettingData ">Reference to the affected virtual system configuration.</param>
// <param name="BootSourceSettings" type="string []">Array of strings each containing one embedded instance of class CIM_SettingData that describes the virtual aspects of a virtual resource to be added to the virtual system.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned. In this case, the instances of class CIM_SettingData representing the added boot source settings are available via association CIM_ConreteComponent from the instance of class CIM_VirtualSystemSettingData representing the affected virtual system configuration.</param>
// <param name="ResultingBootSourceSettings" type="CIM_SettingData []">Array of references to instances of class Msvm_BootSourceSettingData representing properties of the boot sources.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) AddBootSourceSettings( /* IN */ AffectedConfiguration CIM_VirtualSystemSettingData,
	/* IN */ BootSourceSettings []string,
	/* OUT */ ResultingBootSourceSettings []CIM_SettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("AddBootSourceSettings", Action, PercentComplete, Timeout, AffectedConfiguration, BootSourceSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AffectedConfiguration" type="CIM_VirtualSystemSettingData ">Reference to the affected virtual system configuration.</param>
// <param name="GuestServiceSettings" type="string []">Array of strings each containing an embedded instance of class CIM_SettingData that describes addition of the virtual aspects of a guest service. All instances must have a valid service ID in order to identify the guest service setting to be added.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job be returned. In this case, the instances of class CIM_SettingData representing the added guest service settings are available via association CIM_ConreteComponent from the instance of class CIM_VirtualSystemSettingData representing the affected virtual system configuration.</param>
// <param name="ResultingGuestServiceSettings" type="CIM_SettingData []">Array of references to instances of class Cim_SettingData representing virtual aspects of the modified guest services.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) AddGuestServiceSettings( /* IN */ AffectedConfiguration CIM_VirtualSystemSettingData,
	/* IN */ GuestServiceSettings []string,
	/* OUT */ ResultingGuestServiceSettings []CIM_SettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("AddGuestServiceSettings", Action, PercentComplete, Timeout, AffectedConfiguration, GuestServiceSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="GuestServiceSettings" type="string []">Array of strings each containing an embedded instance of class CIM_SettingData that describes modifications to the virtual aspects of an existing guest service. All instances must have a valid InstanceID in order to identify the guest service setting to be modified.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job be returned. In this case, the instances of class CIM_SettingData representing the modified guest service settings are available via association CIM_ConreteComponent from the instance of class CIM_VirtualSystemSettingData representing the affected virtual system configuration.</param>
// <param name="ResultingGuestServiceSettings" type="CIM_SettingData []">Array of references to instances of class Cim_SettingData representing virtual aspects of the modified guest services.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) ModifyGuestServiceSettings( /* IN */ GuestServiceSettings []string,
	/* OUT */ ResultingGuestServiceSettings []CIM_SettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyGuestServiceSettings", Action, PercentComplete, Timeout, GuestServiceSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="BootSourceSettings" type="CIM_SettingData []">Array of references to instances of class CIM_ResourceAllocationSettingData where each instance represents the settings of a boot source within a virtual system configuration that are to be removed.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job my be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) RemoveBootSourceSettings( /* IN */ BootSourceSettings []CIM_SettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RemoveBootSourceSettings", Action, PercentComplete, Timeout, BootSourceSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="GuestServiceSettings" type="CIM_SettingData []">Array of references to instances of class Cim_SettingData representing virtual aspects of the modified guest services.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job my be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) RemoveGuestServiceSettings( /* IN */ GuestServiceSettings []CIM_SettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RemoveGuestServiceSettings", Action, PercentComplete, Timeout, GuestServiceSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="HeightPixels" type="uint16 "></param>
// <param name="TargetSystem" type="CIM_VirtualSystemSettingData "></param>
// <param name="WidthPixels" type="uint16 "></param>

// <param name="ImageData" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) GetVirtualSystemThumbnailImage( /* IN */ TargetSystem CIM_VirtualSystemSettingData,
	/* IN */ WidthPixels uint16,
	/* IN */ HeightPixels uint16,
	/* OUT */ ImageData []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetVirtualSystemThumbnailImage", TargetSystem, WidthPixels, HeightPixels)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) ModifyServiceSettings( /* IN */ SettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyServiceSettings", Action, PercentComplete, Timeout, SettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RequestedInformation" type="uint32 []"></param>
// <param name="SettingData" type="CIM_VirtualSystemSettingData []"></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="SummaryInformation" type="Msvm_SummaryInformationBase []"></param>
func (instance *Msvm_VirtualSystemManagementService) GetSummaryInformation( /* IN */ SettingData []CIM_VirtualSystemSettingData,
	/* IN */ RequestedInformation []uint32,
	/* OUT */ SummaryInformation []Msvm_SummaryInformationBase) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSummaryInformation", SettingData, RequestedInformation)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DefinitionFiles" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="SummaryInformation" type="Msvm_SummaryInformationBase []"></param>
func (instance *Msvm_VirtualSystemManagementService) GetDefinitionFileSummaryInformation( /* IN */ DefinitionFiles []string,
	/* OUT */ SummaryInformation []Msvm_SummaryInformationBase) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetDefinitionFileSummaryInformation", DefinitionFiles)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DataItems" type="string []"></param>
// <param name="TargetSystem" type="CIM_ComputerSystem "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) AddKvpItems( /* IN */ TargetSystem CIM_ComputerSystem,
	/* IN */ DataItems []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("AddKvpItems", Action, PercentComplete, Timeout, TargetSystem, DataItems)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DataItems" type="string []"></param>
// <param name="TargetSystem" type="CIM_ComputerSystem "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) ModifyKvpItems( /* IN */ TargetSystem CIM_ComputerSystem,
	/* IN */ DataItems []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyKvpItems", Action, PercentComplete, Timeout, TargetSystem, DataItems)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DataItems" type="string []"></param>
// <param name="TargetSystem" type="CIM_ComputerSystem "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) RemoveKvpItems( /* IN */ TargetSystem CIM_ComputerSystem,
	/* IN */ DataItems []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RemoveKvpItems", Action, PercentComplete, Timeout, TargetSystem, DataItems)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Errors" type="string []"></param>

// <param name="ErrorMessage" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) FormatError( /* IN */ Errors []string,
	/* OUT */ ErrorMessage string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("FormatError", Errors)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) ModifyDiskMergeSettings( /* IN */ SettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyDiskMergeSettings", Action, PercentComplete, Timeout, SettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="NumberOfWwpns" type="uint32 "></param>

// <param name="GeneratedWwpn" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) GenerateWwpn( /* IN */ NumberOfWwpns uint32,
	/* OUT */ GeneratedWwpn []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GenerateWwpn", NumberOfWwpns)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FcPortSettings" type="string []">Array of strings each containing an embedded instance of class Msvm_SyntheticFcPortSettingData that describes settings for synthetic fibre channel ports for virtual machines.All instances must have a valid InstanceID in order to identify the feature settings to be modified.</param>
// <param name="SecretEncoding" type="VirtualSystemManagementService_SecretEncoding "></param>
// <param name="SharedSecret" type="uint8 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) AddFibreChannelChap( /* IN */ FcPortSettings []string,
	/* IN */ SecretEncoding VirtualSystemManagementService_SecretEncoding,
	/* IN */ SharedSecret []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddFibreChannelChap", FcPortSettings, SecretEncoding, SharedSecret)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="FcPortSettings" type="string []">Array of strings each containing an embedded instance of class Msvm_SyntheticFcPortSettingData that describes settings for synthetic fibre channel ports for virtual machines.All instances must have a valid InstanceID in order to identify the feature settings to be modified.</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) RemoveFibreChannelChap( /* IN */ FcPortSettings []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveFibreChannelChap", FcPortSettings)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="NetworkConfiguration" type="string []"></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) SetGuestNetworkAdapterConfiguration( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ NetworkConfiguration []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("SetGuestNetworkAdapterConfiguration", Action, PercentComplete, Timeout, ComputerSystem, NetworkConfiguration)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Vssd" type="CIM_VirtualSystemSettingData "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Size" type="uint64 "></param>
func (instance *Msvm_VirtualSystemManagementService) GetSizeOfSystemFiles( /* IN */ Vssd CIM_VirtualSystemSettingData,
	/* OUT */ Size uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSizeOfSystemFiles", Vssd)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CurrentWwpn" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) GetCurrentWwpnFromGenerator( /* OUT */ CurrentWwpn string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetCurrentWwpnFromGenerator")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IsolationId" type="uint32 "></param>
// <param name="IsSender" type="bool "></param>
// <param name="ReceiverIP" type="string "></param>
// <param name="ReceiverMac" type="string "></param>
// <param name="SenderIP" type="string "></param>
// <param name="SequenceNumber" type="uint32 "></param>
// <param name="TargetNetworkAdapter" type="Msvm_EthernetPortAllocationSettingData ">Reference to the Ethernet connection.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="RoundTripTime" type="uint32 ">The round trip time for the Ping request.</param>
func (instance *Msvm_VirtualSystemManagementService) TestNetworkConnection( /* IN */ TargetNetworkAdapter Msvm_EthernetPortAllocationSettingData,
	/* IN */ IsSender bool,
	/* IN */ SenderIP string,
	/* IN */ ReceiverIP string,
	/* IN */ ReceiverMac string,
	/* IN */ IsolationId uint32,
	/* IN */ SequenceNumber uint32,
	/* OUT */ RoundTripTime uint32,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("TestNetworkConnection", Action, PercentComplete, Timeout, TargetNetworkAdapter, IsSender, SenderIP, ReceiverIP, ReceiverMac, IsolationId, SequenceNumber)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DiagnosticSettings" type="string ">An embedded instance of class Msvm_NetworkConnectionDiagnosticSettingData that describes the settings used to diagnose the connectivity.</param>
// <param name="TargetNetworkAdapter" type="Msvm_EthernetPortAllocationSettingData ">Reference to the Ethernet connection.</param>

// <param name="DiagnosticInformation" type="string ">If successful, this object contains the output of the ping request. This is an embedded instance of Msvm_NetworkConnectionDiagnosticInformation.</param>
// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) DiagnoseNetworkConnection( /* IN */ TargetNetworkAdapter Msvm_EthernetPortAllocationSettingData,
	/* IN */ DiagnosticSettings string,
	/* OUT */ DiagnosticInformation string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("DiagnoseNetworkConnection", Action, PercentComplete, Timeout, TargetNetworkAdapter, DiagnosticSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ImcData" type="uint8 []"></param>
// <param name="TargetSystem" type="CIM_ComputerSystem "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) SetInitialMachineConfigurationData( /* IN */ TargetSystem CIM_ComputerSystem,
	/* IN */ ImcData []uint8,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("SetInitialMachineConfigurationData", Action, PercentComplete, Timeout, TargetSystem, ImcData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AffectedConfiguration" type="Msvm_VirtualSystemSettingData ">Reference to the affected virtual system configuration.</param>
// <param name="ComponentSettings" type="string []">Array of strings each containing one embedded instanceof class Msvm_SystemComponentSettingData that describes the virtual aspects of a virtual resource to be added to the virtual system.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job may be returned. In this case, the instances of class Msvm_SystemComponentSettingData representing the added component settings are available via association CIM_ConcreteComponent from the instance of class Msvm_VirtualSystemSettingData representing the affected virtual system configuration.</param>
// <param name="ResultingComponentSettings" type="Msvm_SystemComponentSettingData []">Array of references to instances of class Msvm_SystemComponentSettingData representing properties of the resulting components.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) AddSystemComponentSettings( /* IN */ AffectedConfiguration Msvm_VirtualSystemSettingData,
	/* IN */ ComponentSettings []string,
	/* OUT */ ResultingComponentSettings []Msvm_SystemComponentSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("AddSystemComponentSettings", Action, PercentComplete, Timeout, AffectedConfiguration, ComponentSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComponentSettings" type="string []">Array of strings each containing an embedded instance of class Msvm_SystemComponentSettingData that describes modifications to the virtual aspects of an existing system component.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job be returned. In this case, the instances of class Msvm_SystemComponentSettingData representing the modified settings are available via association CIM_ConcreteComponent from the instance of class CIM_VirtualSystemSettingData representing the affected virtual system configuration.</param>
// <param name="ResultingComponentSettings" type="Msvm_SystemComponentSettingData []">Array of references to instances of class Msvm_SystemComponentSettingData representing virtual aspects of the modified components.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) ModifySystemComponentSettings( /* IN */ ComponentSettings []string,
	/* OUT */ ResultingComponentSettings []Msvm_SystemComponentSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifySystemComponentSettings", Action, PercentComplete, Timeout, ComponentSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComponentSettings" type="Msvm_SystemComponentSettingData []">Array of references to instances of class Msvm_SystemComponentSettingData where each instance represents the settings of a boot source within a virtual system configuration that are to be removed.</param>

// <param name="Job" type="CIM_ConcreteJob ">If the operation is long running, then optionally a job my be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemManagementService) RemoveSystemComponentSettings( /* IN */ ComponentSettings []Msvm_SystemComponentSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RemoveSystemComponentSettings", Action, PercentComplete, Timeout, ComponentSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_VirtualSystemManagementService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_VirtualSystemManagementService) GetRelatedVirtualSystemManagementServiceSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemManagementServiceSettingData")
}

func (instance *Msvm_VirtualSystemManagementService) GetRelatedVirtualSystemManagementCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemManagementCapabilities")
}
