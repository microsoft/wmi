// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM struct
type Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM struct {
	Win32_PerfRawData

	//
	AverageReplicationLatency uint64

	//
	AverageReplicationSize uint64

	//
	CompressionEfficiency uint64

	//
	LastReplicationSize uint64

	//
	NetworkBytesRecv uint64

	//
	NetworkBytesSent uint64

	//
	ReplicationCount uint32

	//
	ReplicationLatency uint64

	//
	ResynchronizedBytes uint64
}

// SetAverageReplicationLatency sets the value of AverageReplicationLatency for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) SetPropertyAverageReplicationLatency(value uint64) (err error) {
	return instance.SetProperty("AverageReplicationLatency", value)
}

// GetAverageReplicationLatency gets the value of AverageReplicationLatency for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) GetPropertyAverageReplicationLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageReplicationLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageReplicationSize sets the value of AverageReplicationSize for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) SetPropertyAverageReplicationSize(value uint64) (err error) {
	return instance.SetProperty("AverageReplicationSize", value)
}

// GetAverageReplicationSize gets the value of AverageReplicationSize for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) GetPropertyAverageReplicationSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageReplicationSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompressionEfficiency sets the value of CompressionEfficiency for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) SetPropertyCompressionEfficiency(value uint64) (err error) {
	return instance.SetProperty("CompressionEfficiency", value)
}

// GetCompressionEfficiency gets the value of CompressionEfficiency for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) GetPropertyCompressionEfficiency() (value uint64, err error) {
	retValue, err := instance.GetProperty("CompressionEfficiency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastReplicationSize sets the value of LastReplicationSize for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) SetPropertyLastReplicationSize(value uint64) (err error) {
	return instance.SetProperty("LastReplicationSize", value)
}

// GetLastReplicationSize gets the value of LastReplicationSize for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) GetPropertyLastReplicationSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("LastReplicationSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkBytesRecv sets the value of NetworkBytesRecv for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) SetPropertyNetworkBytesRecv(value uint64) (err error) {
	return instance.SetProperty("NetworkBytesRecv", value)
}

// GetNetworkBytesRecv gets the value of NetworkBytesRecv for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) GetPropertyNetworkBytesRecv() (value uint64, err error) {
	retValue, err := instance.GetProperty("NetworkBytesRecv")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkBytesSent sets the value of NetworkBytesSent for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) SetPropertyNetworkBytesSent(value uint64) (err error) {
	return instance.SetProperty("NetworkBytesSent", value)
}

// GetNetworkBytesSent gets the value of NetworkBytesSent for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) GetPropertyNetworkBytesSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("NetworkBytesSent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationCount sets the value of ReplicationCount for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) SetPropertyReplicationCount(value uint32) (err error) {
	return instance.SetProperty("ReplicationCount", value)
}

// GetReplicationCount gets the value of ReplicationCount for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) GetPropertyReplicationCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplicationCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationLatency sets the value of ReplicationLatency for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) SetPropertyReplicationLatency(value uint64) (err error) {
	return instance.SetProperty("ReplicationLatency", value)
}

// GetReplicationLatency gets the value of ReplicationLatency for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) GetPropertyReplicationLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReplicationLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResynchronizedBytes sets the value of ResynchronizedBytes for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) SetPropertyResynchronizedBytes(value uint64) (err error) {
	return instance.SetProperty("ResynchronizedBytes", value)
}

// GetResynchronizedBytes gets the value of ResynchronizedBytes for the instance
func (instance *Win32_PerfRawData_HyperVReplicaStats_HyperVReplicaVM) GetPropertyResynchronizedBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResynchronizedBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
