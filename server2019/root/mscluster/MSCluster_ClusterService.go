// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_ClusterService struct
type MSCluster_ClusterService struct {
	cim.WmiInstance

	//
	ElementName string
}

// SetElementName sets the value of ElementName for the instance
func (instance *MSCluster_ClusterService) SetPropertyElementName(value string) (err error) {
	return instance.SetProperty("ElementName", value)
}

// GetElementName gets the value of ElementName for the instance
func (instance *MSCluster_ClusterService) GetPropertyElementName() (value string, err error) {
	retValue, err := instance.GetProperty("ElementName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="WhatIf" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) UpdateFunctionalLevel( /* IN */ WhatIf bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UpdateFunctionalLevel", WhatIf)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="isReady" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) ClusterIsReadyForUpgrade( /* OUT */ isReady bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ClusterIsReadyForUpgrade")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) EnableHealth() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("EnableHealth")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) DisableHealth() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DisableHealth")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Providers" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) AddHealthProviders( /* IN */ Providers []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddHealthProviders", Providers)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Providers" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) RemoveHealthProviders( /* IN */ Providers []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveHealthProviders", Providers)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="FaultDomains" type="uint32 "></param>
// <param name="LocalDiskSize" type="uint32 "></param>
// <param name="ReservedNodes" type="float64 "></param>
// <param name="ReserveSpareNode" type="bool "></param>
// <param name="UpdateDomains" type="uint32 "></param>
// <param name="Version" type="uint32 "></param>
// <param name="VmCpuReservation" type="uint32 "></param>
// <param name="VmFlags" type="uint32 "></param>
// <param name="VmMemory" type="uint32 "></param>
// <param name="VmVirtualCoreCount" type="uint32 "></param>

// <param name="MaxNumOfVMsInCluster" type="uint32 "></param>
// <param name="MaxNumOfVMsInNode" type="uint32 "></param>
// <param name="PlacementScoreFlags" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) GetPlacementScore( /* IN */ VmMemory uint32,
	/* IN */ VmVirtualCoreCount uint32,
	/* IN */ VmCpuReservation uint32,
	/* IN */ VmFlags uint32,
	/* IN */ ReservedNodes float64,
	/* OUT */ MaxNumOfVMsInCluster uint32,
	/* OUT */ MaxNumOfVMsInNode uint32,
	/* OUT */ PlacementScoreFlags uint32,
	/* OPTIONAL IN */ Version uint32,
	/* OPTIONAL IN */ LocalDiskSize uint32,
	/* OPTIONAL IN */ UpdateDomains uint32,
	/* OPTIONAL IN */ FaultDomains uint32,
	/* OPTIONAL IN */ ReserveSpareNode bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetPlacementScore", VmMemory, VmVirtualCoreCount, VmCpuReservation, VmFlags, ReservedNodes, Version, LocalDiskSize, UpdateDomains, FaultDomains, ReserveSpareNode)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AvailabilitySetName" type="string "></param>
// <param name="LocalDiskSize" type="uint32 "></param>
// <param name="ReservationId" type="string "></param>
// <param name="TimeSpan" type="uint32 "></param>
// <param name="Version" type="uint32 "></param>
// <param name="VmCpuReservation" type="uint32 "></param>
// <param name="VmFlags" type="uint32 "></param>
// <param name="VmMemory" type="uint32 "></param>
// <param name="VmVirtualCoreCount" type="uint32 "></param>

// <param name="NodeId" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) CreateVmReservation( /* IN */ VmMemory uint32,
	/* IN */ VmVirtualCoreCount uint32,
	/* IN */ VmCpuReservation uint32,
	/* IN */ VmFlags uint32,
	/* IN */ TimeSpan uint32,
	/* IN */ ReservationId string,
	/* OUT */ NodeId uint32,
	/* OPTIONAL IN */ Version uint32,
	/* OPTIONAL IN */ LocalDiskSize uint32,
	/* OPTIONAL IN */ AvailabilitySetName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateVmReservation", VmMemory, VmVirtualCoreCount, VmCpuReservation, VmFlags, TimeSpan, ReservationId, Version, LocalDiskSize, AvailabilitySetName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AvailabilitySetName" type="string "></param>
// <param name="FaultDomain" type="uint32 "></param>
// <param name="LocalDiskSize" type="uint32 "></param>
// <param name="ReservationId" type="string "></param>
// <param name="Reserved" type="uint64 "></param>
// <param name="TimeSpan" type="uint32 "></param>
// <param name="UpdateDomain" type="uint32 "></param>
// <param name="Version" type="uint32 "></param>
// <param name="VmCpuReservation" type="uint32 "></param>
// <param name="VmFlags" type="uint32 "></param>
// <param name="VmMemory" type="uint32 "></param>
// <param name="VmVirtualCoreCount" type="uint32 "></param>

// <param name="NodeId" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) CreateVmReservationEx( /* IN */ VmMemory uint32,
	/* IN */ VmVirtualCoreCount uint32,
	/* IN */ VmCpuReservation uint32,
	/* IN */ VmFlags uint32,
	/* IN */ TimeSpan uint32,
	/* IN */ ReservationId string,
	/* OUT */ NodeId uint32,
	/* OPTIONAL IN */ Version uint32,
	/* OPTIONAL IN */ LocalDiskSize uint32,
	/* OPTIONAL IN */ AvailabilitySetName string,
	/* OPTIONAL IN */ FaultDomain uint32,
	/* OPTIONAL IN */ UpdateDomain uint32,
	/* OPTIONAL IN */ Reserved uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateVmReservationEx", VmMemory, VmVirtualCoreCount, VmCpuReservation, VmFlags, TimeSpan, ReservationId, Version, LocalDiskSize, AvailabilitySetName, FaultDomain, UpdateDomain, Reserved)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="ReservationId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) RemoveVmReservation( /* IN */ ReservationId string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveVmReservation", ReservationId, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Version" type="uint32 "></param>
// <param name="VmCpuReservation" type="uint32 "></param>
// <param name="VmFlags" type="uint32 "></param>
// <param name="VmMemory" type="uint32 "></param>
// <param name="VmResourceName" type="string "></param>
// <param name="VmVirtualCoreCount" type="uint32 "></param>

// <param name="NodeId" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) ChangeVMSettings( /* IN */ VmResourceName string,
	/* IN */ VmMemory uint32,
	/* IN */ VmVirtualCoreCount uint32,
	/* IN */ VmCpuReservation uint32,
	/* IN */ VmFlags uint32,
	/* OUT */ NodeId uint32,
	/* OPTIONAL IN */ Version uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ChangeVMSettings", VmResourceName, VmMemory, VmVirtualCoreCount, VmCpuReservation, VmFlags, Version)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AccountKey" type="string "></param>
// <param name="AccountName" type="string "></param>
// <param name="CloudWitnessName" type="string "></param>
// <param name="EndpointInfo" type="string "></param>
// <param name="SASToken" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) CreateCloudWitness( /* IN */ AccountName string,
	/* IN */ AccountKey string,
	/* IN */ SASToken string,
	/* IN */ EndpointInfo string,
	/* IN */ CloudWitnessName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CreateCloudWitness", AccountName, AccountKey, SASToken, EndpointInfo, CloudWitnessName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="AccountKey" type="string "></param>
// <param name="SASToken" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) UpdateCloudWitnessKey( /* IN */ AccountKey string,
	/* IN */ SASToken string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UpdateCloudWitnessKey", AccountKey, SASToken)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="CurrentNodeName" type="string "></param>
// <param name="ReplacementNodeName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) NodeReplacement( /* IN */ CurrentNodeName string,
	/* IN */ ReplacementNodeName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("NodeReplacement", CurrentNodeName, ReplacementNodeName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) CreateClusterSet() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CreateClusterSet")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Force" type="bool "></param>
// <param name="RemoveFileServer" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) RemoveClusterSet( /* IN */ Force bool,
	/* IN */ RemoveFileServer bool,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveClusterSet", Force, RemoveFileServer, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Force" type="bool "></param>
// <param name="RemoveFileServer" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) RemoveClusterSetWorker( /* IN */ Force bool,
	/* IN */ RemoveFileServer bool,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveClusterSetWorker", Force, RemoveFileServer, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="CertData" type="uint8 []"></param>
// <param name="Flags" type="uint32 "></param>
// <param name="Id" type="uint32 "></param>
// <param name="InfraSOFSName" type="string "></param>
// <param name="KeyData" type="uint8 []"></param>
// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterService) SetupClusterSetWorker( /* IN */ Name string,
	/* IN */ InfraSOFSName string,
	/* IN */ Id uint32,
	/* IN */ CertData []uint8,
	/* IN */ KeyData []uint8,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetupClusterSetWorker", Name, InfraSOFSName, Id, CertData, KeyData, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
// <param name="SupportedVersions" type="MSCluster_NodeSupportedVersion []"></param>
func (instance *MSCluster_ClusterService) GetNodeSupportedVersions( /* OUT */ SupportedVersions []MSCluster_NodeSupportedVersion) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetNodeSupportedVersions")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
