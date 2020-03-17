// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageNodeToPhysicalDisk struct
type MSFT_StorageNodeToPhysicalDisk struct {
	cim.WmiInstance

	//
	DiskNumber uint32

	//
	HealthStatus uint16

	//
	IsMpioEnabled bool

	//
	IsPhysicallyConnected bool

	//
	LoadBalancePolicy uint16

	//
	OperationalStatus []uint16

	//
	PathId []string

	//
	PathState []uint16

	//
	PhysicalDisk MSFT_PhysicalDisk

	//
	StorageNode MSFT_StorageNode
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
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyHealthStatus(value uint16) (err error) {
	return instance.SetProperty("HealthStatus", value)
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyHealthStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
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
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyLoadBalancePolicy(value uint16) (err error) {
	return instance.SetProperty("LoadBalancePolicy", value)
}

// GetLoadBalancePolicy gets the value of LoadBalancePolicy for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyLoadBalancePolicy() (value uint16, err error) {
	retValue, err := instance.GetProperty("LoadBalancePolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyOperationalStatus(value []uint16) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyOperationalStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
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
func (instance *MSFT_StorageNodeToPhysicalDisk) SetPropertyPathState(value []uint16) (err error) {
	return instance.SetProperty("PathState", value)
}

// GetPathState gets the value of PathState for the instance
func (instance *MSFT_StorageNodeToPhysicalDisk) GetPropertyPathState() (value []uint16, err error) {
	retValue, err := instance.GetProperty("PathState")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
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
