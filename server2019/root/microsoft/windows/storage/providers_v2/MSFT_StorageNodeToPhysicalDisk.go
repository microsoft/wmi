// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageNodeToPhysicalDisk struct
type MSFT_StorageNodeToPhysicalDisk struct {
	*cim.WmiInstance

	// The operating system's number for the disk on this StorageNode. Disk 0 is typically the boot device. Disk numbers may not necessarily remain the same across reboot, and are not necessarily the same on different nodes.
	DiskNumber uint32

	// Denotes the health status of the PhysicalDisk on this StorageNode.
	HealthStatus StorageNodeToPhysicalDisk_HealthStatus

	// Indicates whether the physical disk uses MPIO.
	IsMpioEnabled bool

	// Indicates whether the physical disk is physically connected to this storage node.
	IsPhysicallyConnected bool

	// The MPIO load balance policy being used by the disk.
	LoadBalancePolicy StorageNodeToPhysicalDisk_LoadBalancePolicy

	// Denotes the operational status of the PhysicalDisk.
	OperationalStatus []StorageNodeToPhysicalDisk_OperationalStatus

	// Collection of MPIO path IDs, reported by the MPIO DSM, when applicable.
	PathId []string

	// The current state of MPIO paths between the node and physical disk.
	PathState []StorageNodeToPhysicalDisk_PathState

	//
	PhysicalDisk MSFT_PhysicalDisk

	//
	StorageNode MSFT_StorageNode
}

func NewMSFT_StorageNodeToPhysicalDiskEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageNodeToPhysicalDisk, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageNodeToPhysicalDisk{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_StorageNodeToPhysicalDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageNodeToPhysicalDisk, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageNodeToPhysicalDisk{
		WmiInstance: tmp,
	}
	return
}

// SetDiskNumber sets the value of DiskNumber for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyDiskNumber(value uint32) (err error) {
	return instance.SetProperty("DiskNumber", value)
}

// GetDiskNumber gets the value of DiskNumber for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyDiskNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DiskNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyHealthStatus(value StorageNodeToPhysicalDisk_HealthStatus) (err error) {
	return instance.SetProperty("HealthStatus", value)
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyHealthStatus() (value StorageNodeToPhysicalDisk_HealthStatus, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(StorageNodeToPhysicalDisk_HealthStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsMpioEnabled sets the value of IsMpioEnabled for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyIsMpioEnabled(value bool) (err error) {
	return instance.SetProperty("IsMpioEnabled", value)
}

// GetIsMpioEnabled gets the value of IsMpioEnabled for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyIsMpioEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsMpioEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsPhysicallyConnected sets the value of IsPhysicallyConnected for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyIsPhysicallyConnected(value bool) (err error) {
	return instance.SetProperty("IsPhysicallyConnected", value)
}

// GetIsPhysicallyConnected gets the value of IsPhysicallyConnected for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyIsPhysicallyConnected() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPhysicallyConnected")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoadBalancePolicy sets the value of LoadBalancePolicy for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyLoadBalancePolicy(value StorageNodeToPhysicalDisk_LoadBalancePolicy) (err error) {
	return instance.SetProperty("LoadBalancePolicy", value)
}

// GetLoadBalancePolicy gets the value of LoadBalancePolicy for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyLoadBalancePolicy() (value StorageNodeToPhysicalDisk_LoadBalancePolicy, err error) {
	retValue, err := instance.GetProperty("LoadBalancePolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(StorageNodeToPhysicalDisk_LoadBalancePolicy)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyOperationalStatus(value []StorageNodeToPhysicalDisk_OperationalStatus) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyOperationalStatus() (value []StorageNodeToPhysicalDisk_OperationalStatus, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]StorageNodeToPhysicalDisk_OperationalStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPathId sets the value of PathId for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyPathId(value []string) (err error) {
	return instance.SetProperty("PathId", value)
}

// GetPathId gets the value of PathId for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyPathId() (value []string, err error) {
	retValue, err := instance.GetProperty("PathId")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPathState sets the value of PathState for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyPathState(value []StorageNodeToPhysicalDisk_PathState) (err error) {
	return instance.SetProperty("PathState", value)
}

// GetPathState gets the value of PathState for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyPathState() (value []StorageNodeToPhysicalDisk_PathState, err error) {
	retValue, err := instance.GetProperty("PathState")
	if err != nil {
		return
	}
	value, ok := retValue.([]StorageNodeToPhysicalDisk_PathState)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhysicalDisk sets the value of PhysicalDisk for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyPhysicalDisk(value MSFT_PhysicalDisk) (err error) {
	return instance.SetProperty("PhysicalDisk", value)
}

// GetPhysicalDisk gets the value of PhysicalDisk for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyPhysicalDisk() (value MSFT_PhysicalDisk, err error) {
	retValue, err := instance.GetProperty("PhysicalDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_PhysicalDisk)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageNode sets the value of StorageNode for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyStorageNode(value MSFT_StorageNode) (err error) {
	return instance.SetProperty("StorageNode", value)
}

// GetStorageNode gets the value of StorageNode for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyStorageNode() (value MSFT_StorageNode, err error) {
	retValue, err := instance.GetProperty("StorageNode")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageNode)
	if !ok {
		// TODO: Set an error
	}
	return
}
