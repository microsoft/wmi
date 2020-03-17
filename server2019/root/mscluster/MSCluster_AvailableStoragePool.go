// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_AvailableStoragePool struct
type MSCluster_AvailableStoragePool struct {
	MSCluster_LogicalElement

	//
	Attributes uint64

	//
	ConnectedNodes []string

	//
	HealthStatus uint32

	//
	Id string

	//
	QuorumStatus uint32

	//
	TotalSize uint64

	//
	Usage uint64
}

// SetAttributes sets the value of Attributes for the instance
func (instance *MSCluster_AvailableStoragePool) SetPropertyAttributes(value uint64) (err error) {
	return instance.SetProperty("Attributes", value)
}

// GetAttributes gets the value of Attributes for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyAttributes() (value uint64, err error) {
	retValue, err := instance.GetProperty("Attributes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectedNodes sets the value of ConnectedNodes for the instance
func (instance *MSCluster_AvailableStoragePool) SetPropertyConnectedNodes(value []string) (err error) {
	return instance.SetProperty("ConnectedNodes", value)
}

// GetConnectedNodes gets the value of ConnectedNodes for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyConnectedNodes() (value []string, err error) {
	retValue, err := instance.GetProperty("ConnectedNodes")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSCluster_AvailableStoragePool) SetPropertyHealthStatus(value uint32) (err error) {
	return instance.SetProperty("HealthStatus", value)
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyHealthStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
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
func (instance *MSCluster_AvailableStoragePool) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuorumStatus sets the value of QuorumStatus for the instance
func (instance *MSCluster_AvailableStoragePool) SetPropertyQuorumStatus(value uint32) (err error) {
	return instance.SetProperty("QuorumStatus", value)
}

// GetQuorumStatus gets the value of QuorumStatus for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyQuorumStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuorumStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalSize sets the value of TotalSize for the instance
func (instance *MSCluster_AvailableStoragePool) SetPropertyTotalSize(value uint64) (err error) {
	return instance.SetProperty("TotalSize", value)
}

// GetTotalSize gets the value of TotalSize for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyTotalSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUsage sets the value of Usage for the instance
func (instance *MSCluster_AvailableStoragePool) SetPropertyUsage(value uint64) (err error) {
	return instance.SetProperty("Usage", value)
}

// GetUsage gets the value of Usage for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyUsage() (value uint64, err error) {
	retValue, err := instance.GetProperty("Usage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

//
func (instance *MSCluster_AvailableStoragePool) AddToCluster() (err error) {
	_, err = instance.InvokeMethodWithReturn("AddToCluster")
	if err != nil {
		return
	}
	return

}
