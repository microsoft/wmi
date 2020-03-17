// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Cluster.Scaleout
//////////////////////////////////////////////
package scaleout

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ClusterSet struct
type MSFT_ClusterSet struct {
	cim.WmiInstance

	//
	ClusterName string

	//
	EvacuationMoveThreshold uint32

	//
	Id uint64

	//
	MigrationExcludeNetworks string

	//
	MigrationNetworkOrder string

	//
	Name string

	//
	NamespaceRoot string

	//
	VMFailoverMode uint32
}

// SetClusterName sets the value of ClusterName for the instance
func (instance *MSFT_ClusterSet) SetPropertyClusterName(value string) (err error) {
	return instance.SetProperty("ClusterName", value)
}

// GetClusterName gets the value of ClusterName for the instance
func (instance *MSFT_ClusterSet) GetPropertyClusterName() (value string, err error) {
	retValue, err := instance.GetProperty("ClusterName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEvacuationMoveThreshold sets the value of EvacuationMoveThreshold for the instance
func (instance *MSFT_ClusterSet) SetPropertyEvacuationMoveThreshold(value uint32) (err error) {
	return instance.SetProperty("EvacuationMoveThreshold", value)
}

// GetEvacuationMoveThreshold gets the value of EvacuationMoveThreshold for the instance
func (instance *MSFT_ClusterSet) GetPropertyEvacuationMoveThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("EvacuationMoveThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *MSFT_ClusterSet) SetPropertyId(value uint64) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_ClusterSet) GetPropertyId() (value uint64, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMigrationExcludeNetworks sets the value of MigrationExcludeNetworks for the instance
func (instance *MSFT_ClusterSet) SetPropertyMigrationExcludeNetworks(value string) (err error) {
	return instance.SetProperty("MigrationExcludeNetworks", value)
}

// GetMigrationExcludeNetworks gets the value of MigrationExcludeNetworks for the instance
func (instance *MSFT_ClusterSet) GetPropertyMigrationExcludeNetworks() (value string, err error) {
	retValue, err := instance.GetProperty("MigrationExcludeNetworks")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMigrationNetworkOrder sets the value of MigrationNetworkOrder for the instance
func (instance *MSFT_ClusterSet) SetPropertyMigrationNetworkOrder(value string) (err error) {
	return instance.SetProperty("MigrationNetworkOrder", value)
}

// GetMigrationNetworkOrder gets the value of MigrationNetworkOrder for the instance
func (instance *MSFT_ClusterSet) GetPropertyMigrationNetworkOrder() (value string, err error) {
	retValue, err := instance.GetProperty("MigrationNetworkOrder")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_ClusterSet) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ClusterSet) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNamespaceRoot sets the value of NamespaceRoot for the instance
func (instance *MSFT_ClusterSet) SetPropertyNamespaceRoot(value string) (err error) {
	return instance.SetProperty("NamespaceRoot", value)
}

// GetNamespaceRoot gets the value of NamespaceRoot for the instance
func (instance *MSFT_ClusterSet) GetPropertyNamespaceRoot() (value string, err error) {
	retValue, err := instance.GetProperty("NamespaceRoot")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVMFailoverMode sets the value of VMFailoverMode for the instance
func (instance *MSFT_ClusterSet) SetPropertyVMFailoverMode(value uint32) (err error) {
	return instance.SetProperty("VMFailoverMode", value)
}

// GetVMFailoverMode gets the value of VMFailoverMode for the instance
func (instance *MSFT_ClusterSet) GetPropertyVMFailoverMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("VMFailoverMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Name" type="string "></param>

// <param name="Member" type="MSFT_ClusterSetMember "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSet) AddMember( /* IN */ Name string,
	/* IN */ Flags uint32,
	/* OUT */ Member MSFT_ClusterSetMember) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddMember", Name, Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="vmConfig" type="string "></param>
// <param name="vmHost" type="string "></param>
// <param name="vmId" type="string "></param>
// <param name="vmName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VM" type="MSFT_ClusterSetVM "></param>
func (instance *MSFT_ClusterSet) AddVm( /* IN */ vmConfig string,
	/* IN */ vmId string,
	/* IN */ vmName string,
	/* IN */ vmHost string,
	/* OUT */ VM MSFT_ClusterSetVM) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddVm", vmConfig, vmId, vmName, vmHost)
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
// <param name="PlacementCondition" type="string "></param>
// <param name="Version" type="uint32 "></param>
// <param name="VmCpuReservation" type="uint32 "></param>
// <param name="VmFlags" type="uint32 "></param>
// <param name="VmMemory" type="uint32 "></param>
// <param name="VmVirtualCoreCount" type="uint32 "></param>

// <param name="Node" type="MSFT_ClusterSetNode "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSet) GetOptimalNodeForVm( /* IN */ VmMemory uint32,
	/* IN */ VmVirtualCoreCount uint32,
	/* IN */ VmCpuReservation uint32,
	/* IN */ VmFlags uint32,
	/* IN */ Version uint32,
	/* IN */ LocalDiskSize uint32,
	/* IN */ PlacementCondition string,
	/* IN */ AvailabilitySetName string,
	/* OUT */ Node MSFT_ClusterSetNode) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetOptimalNodeForVm", VmMemory, VmVirtualCoreCount, VmCpuReservation, VmFlags, Version, LocalDiskSize, PlacementCondition, AvailabilitySetName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="operation" type="string "></param>
// <param name="param1" type="string "></param>
// <param name="param2" type="string "></param>
// <param name="param3" type="string "></param>

// <param name="outparam1" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSet) DoOp( /* IN */ operation string,
	/* IN */ param1 string,
	/* IN */ param2 string,
	/* IN */ param3 string,
	/* OUT */ outparam1 string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DoOp", operation, param1, param2, param3)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="NamespaceRoot" type="string "></param>
// <param name="VMFailoverMode" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSet) Set( /* IN */ NamespaceRoot string,
	/* IN */ VMFailoverMode uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set", NamespaceRoot, VMFailoverMode)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="DestinationClusterName" type="string "></param>
// <param name="DestinationRGName" type="string "></param>
// <param name="Flags" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="SourceClusterName" type="string "></param>
// <param name="SourceRGName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSet) AddSRPartnership( /* IN */ Name string,
	/* IN */ SourceClusterName string,
	/* IN */ SourceRGName string,
	/* IN */ DestinationClusterName string,
	/* IN */ DestinationRGName string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddSRPartnership", Name, SourceClusterName, SourceRGName, DestinationClusterName, DestinationRGName, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
