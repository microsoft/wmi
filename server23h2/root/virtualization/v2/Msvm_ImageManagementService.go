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

// Msvm_ImageManagementService struct
type Msvm_ImageManagementService struct {
	*CIM_Service
}

func NewMsvm_ImageManagementServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_ImageManagementService, err error) {
	tmp, err := NewCIM_ServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ImageManagementService{
		CIM_Service: tmp,
	}
	return
}

func NewMsvm_ImageManagementServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ImageManagementService, err error) {
	tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ImageManagementService{
		CIM_Service: tmp,
	}
	return
}

//

// <param name="VirtualDiskSettingData" type="string ">String containing an embedded instance of class Msvm_VirtualHardDiskSettingData that is used to define attributes of the virtual hard disk to be set.</param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) SetVirtualHardDiskSettingData( /* IN */ VirtualDiskSettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("SetVirtualHardDiskSettingData", Action, PercentComplete, Timeout, VirtualDiskSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="VirtualDiskSettingData" type="string ">String containing an embedded instance of class Msvm_VirtualHardDiskSettingData that is used to define attributes of the virtual hard disk to be created.</param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) CreateVirtualHardDisk( /* IN */ VirtualDiskSettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("CreateVirtualHardDisk", Action, PercentComplete, Timeout, VirtualDiskSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ChildPath" type="string "></param>
// <param name="IgnoreIDMismatch" type="bool "></param>
// <param name="LeafPath" type="string "></param>
// <param name="ParentPath" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) SetParentVirtualHardDisk( /* IN */ ChildPath string,
	/* IN */ ParentPath string,
	/* IN */ LeafPath string,
	/* IN */ IgnoreIDMismatch bool,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("SetParentVirtualHardDisk", Action, PercentComplete, Timeout, ChildPath, ParentPath, LeafPath, IgnoreIDMismatch)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Path" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) CreateVirtualFloppyDisk( /* IN */ Path string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("CreateVirtualFloppyDisk", Action, PercentComplete, Timeout, Path)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DestinationPath" type="string "></param>
// <param name="SourcePath" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) MergeVirtualHardDisk( /* IN */ SourcePath string,
	/* IN */ DestinationPath string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("MergeVirtualHardDisk", Action, PercentComplete, Timeout, SourcePath, DestinationPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Mode" type="uint16 "></param>
// <param name="Path" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) CompactVirtualHardDisk( /* IN */ Path string,
	/* IN */ Mode uint16,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("CompactVirtualHardDisk", Action, PercentComplete, Timeout, Path, Mode)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="MaxInternalSize" type="uint64 "></param>
// <param name="Path" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) ResizeVirtualHardDisk( /* IN */ Path string,
	/* IN */ MaxInternalSize uint64,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ResizeVirtualHardDisk", Action, PercentComplete, Timeout, Path, MaxInternalSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SourcePath" type="string "></param>
// <param name="VirtualDiskSettingData" type="string ">String containing an embedded instance of class Msvm_VirtualHardDiskSettingData that is used to define attributes of the virtual hard disk to be created.</param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) ConvertVirtualHardDisk( /* IN */ SourcePath string,
	/* IN */ VirtualDiskSettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ConvertVirtualHardDisk", Action, PercentComplete, Timeout, SourcePath, VirtualDiskSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Path" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SettingData" type="string "></param>
func (instance *Msvm_ImageManagementService) GetVirtualHardDiskSettingData( /* IN */ Path string,
	/* OUT */ SettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("GetVirtualHardDiskSettingData", Action, PercentComplete, Timeout, Path)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Path" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="State" type="string "></param>
func (instance *Msvm_ImageManagementService) GetVirtualHardDiskState( /* IN */ Path string,
	/* OUT */ State string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("GetVirtualHardDiskState", Action, PercentComplete, Timeout, Path)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AssignDriveLetter" type="bool "></param>
// <param name="Path" type="string "></param>
// <param name="ReadOnly" type="bool "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) AttachVirtualHardDisk( /* IN */ Path string,
	/* IN */ AssignDriveLetter bool,
	/* IN */ ReadOnly bool,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("AttachVirtualHardDisk", Action, PercentComplete, Timeout, Path, AssignDriveLetter, ReadOnly)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Path" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) ValidateVirtualHardDisk( /* IN */ Path string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ValidateVirtualHardDisk", Action, PercentComplete, Timeout, Path)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Path" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) ValidatePersistentReservationSupport( /* IN */ Path string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ValidatePersistentReservationSupport", Action, PercentComplete, Timeout, Path)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AdditionalInformation" type="uint32 []"></param>
// <param name="VHDSetPath" type="string "></param>

// <param name="Information" type="string "></param>
// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) GetVHDSetInformation( /* IN */ VHDSetPath string,
	/* IN */ AdditionalInformation []uint32,
	/* OUT */ Information string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("GetVHDSetInformation", Action, PercentComplete, Timeout, VHDSetPath, AdditionalInformation)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AdditionalInformation" type="uint32 []"></param>
// <param name="SnapshotIds" type="string []"></param>
// <param name="VHDSetPath" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SnapshotInformation" type="string []"></param>
func (instance *Msvm_ImageManagementService) GetVHDSnapshotInformation( /* IN */ VHDSetPath string,
	/* IN */ SnapshotIds []string,
	/* IN */ AdditionalInformation []uint32,
	/* OUT */ SnapshotInformation []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("GetVHDSnapshotInformation", Action, PercentComplete, Timeout, VHDSetPath, SnapshotIds, AdditionalInformation)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Information" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) SetVHDSnapshotInformation( /* IN */ Information string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("SetVHDSnapshotInformation", Action, PercentComplete, Timeout, Information)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PersistReferenceSnapshot" type="bool "></param>
// <param name="SnapshotId" type="string "></param>
// <param name="VHDSetPath" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) DeleteVHDSnapshot( /* IN */ VHDSetPath string,
	/* IN */ SnapshotId string,
	/* IN */ PersistReferenceSnapshot bool,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("DeleteVHDSnapshot", Action, PercentComplete, Timeout, VHDSetPath, SnapshotId, PersistReferenceSnapshot)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="VirtualHardDiskPath" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) ConvertVirtualHardDiskToVHDSet( /* IN */ VirtualHardDiskPath string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ConvertVirtualHardDiskToVHDSet", Action, PercentComplete, Timeout, VirtualHardDiskPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="VHDSetPath" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) OptimizeVHDSet( /* IN */ VHDSetPath string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("OptimizeVHDSet", Action, PercentComplete, Timeout, VHDSetPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ByteLength" type="uint64 "></param>
// <param name="ByteOffset" type="uint64 "></param>
// <param name="LimitId" type="string "></param>
// <param name="Path" type="string "></param>
// <param name="TargetSnapshotId" type="string "></param>

// <param name="ChangedByteLengths" type="uint64 []"></param>
// <param name="ChangedByteOffsets" type="uint64 []"></param>
// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ProcessedByteLength" type="uint64 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) GetVirtualDiskChanges( /* IN */ Path string,
	/* IN */ LimitId string,
	/* IN */ TargetSnapshotId string,
	/* IN */ ByteOffset uint64,
	/* IN */ ByteLength uint64,
	/* OUT */ ProcessedByteLength uint64,
	/* OUT */ ChangedByteOffsets []uint64,
	/* OUT */ ChangedByteLengths []uint64,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("GetVirtualDiskChanges", Action, PercentComplete, Timeout, Path, LimitId, TargetSnapshotId, ByteOffset, ByteLength)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CriterionType" type="uint16 "></param>
// <param name="SelectionCriterion" type="string "></param>

// <param name="Image" type="Msvm_MountedStorageImage "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ImageManagementService) FindMountedStorageImageInstance( /* IN */ SelectionCriterion string,
	/* IN */ CriterionType uint16,
	/* OUT */ Image Msvm_MountedStorageImage) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("FindMountedStorageImageInstance", SelectionCriterion, CriterionType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_ImageManagementService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}
