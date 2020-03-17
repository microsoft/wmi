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

// MSFT_StorageQoSFlow struct
type MSFT_StorageQoSFlow struct {
	cim.WmiInstance

	//
	BandwidthLimit uint64

	//
	FilePath string

	//
	FlowId string

	//
	InitiatorBandwidth uint64

	//
	InitiatorId string

	//
	InitiatorIOPS uint64

	//
	InitiatorLatency uint64

	//
	InitiatorName string

	//
	InitiatorNodeName string

	//
	Interval uint64

	//
	Limit uint64

	//
	PolicyId string

	//
	Reservation uint64

	//
	Status uint16

	//
	StorageNodeBandwidth uint64

	//
	StorageNodeIOPS uint64

	//
	StorageNodeLatency uint64

	//
	StorageNodeName string

	//
	TimeStamp uint64

	//
	VolumeId string
}

// SetBandwidthLimit sets the value of BandwidthLimit for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyBandwidthLimit(value uint64) (err error) {
	return instance.SetProperty("BandwidthLimit", value)
}

// GetBandwidthLimit gets the value of BandwidthLimit for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyBandwidthLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("BandwidthLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFilePath sets the value of FilePath for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyFilePath(value string) (err error) {
	return instance.SetProperty("FilePath", value)
}

// GetFilePath gets the value of FilePath for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("FilePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlowId sets the value of FlowId for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyFlowId(value string) (err error) {
	return instance.SetProperty("FlowId", value)
}

// GetFlowId gets the value of FlowId for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyFlowId() (value string, err error) {
	retValue, err := instance.GetProperty("FlowId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInitiatorBandwidth sets the value of InitiatorBandwidth for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyInitiatorBandwidth(value uint64) (err error) {
	return instance.SetProperty("InitiatorBandwidth", value)
}

// GetInitiatorBandwidth gets the value of InitiatorBandwidth for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyInitiatorBandwidth() (value uint64, err error) {
	retValue, err := instance.GetProperty("InitiatorBandwidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInitiatorId sets the value of InitiatorId for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyInitiatorId(value string) (err error) {
	return instance.SetProperty("InitiatorId", value)
}

// GetInitiatorId gets the value of InitiatorId for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyInitiatorId() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInitiatorIOPS sets the value of InitiatorIOPS for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyInitiatorIOPS(value uint64) (err error) {
	return instance.SetProperty("InitiatorIOPS", value)
}

// GetInitiatorIOPS gets the value of InitiatorIOPS for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyInitiatorIOPS() (value uint64, err error) {
	retValue, err := instance.GetProperty("InitiatorIOPS")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInitiatorLatency sets the value of InitiatorLatency for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyInitiatorLatency(value uint64) (err error) {
	return instance.SetProperty("InitiatorLatency", value)
}

// GetInitiatorLatency gets the value of InitiatorLatency for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyInitiatorLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("InitiatorLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInitiatorName sets the value of InitiatorName for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyInitiatorName(value string) (err error) {
	return instance.SetProperty("InitiatorName", value)
}

// GetInitiatorName gets the value of InitiatorName for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyInitiatorName() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInitiatorNodeName sets the value of InitiatorNodeName for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyInitiatorNodeName(value string) (err error) {
	return instance.SetProperty("InitiatorNodeName", value)
}

// GetInitiatorNodeName gets the value of InitiatorNodeName for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyInitiatorNodeName() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorNodeName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterval sets the value of Interval for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyInterval(value uint64) (err error) {
	return instance.SetProperty("Interval", value)
}

// GetInterval gets the value of Interval for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyInterval() (value uint64, err error) {
	retValue, err := instance.GetProperty("Interval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLimit sets the value of Limit for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyLimit(value uint64) (err error) {
	return instance.SetProperty("Limit", value)
}

// GetLimit gets the value of Limit for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("Limit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyId sets the value of PolicyId for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyPolicyId(value string) (err error) {
	return instance.SetProperty("PolicyId", value)
}

// GetPolicyId gets the value of PolicyId for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyPolicyId() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReservation sets the value of Reservation for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyReservation(value uint64) (err error) {
	return instance.SetProperty("Reservation", value)
}

// GetReservation gets the value of Reservation for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyReservation() (value uint64, err error) {
	retValue, err := instance.GetProperty("Reservation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyStatus(value uint16) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageNodeBandwidth sets the value of StorageNodeBandwidth for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyStorageNodeBandwidth(value uint64) (err error) {
	return instance.SetProperty("StorageNodeBandwidth", value)
}

// GetStorageNodeBandwidth gets the value of StorageNodeBandwidth for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyStorageNodeBandwidth() (value uint64, err error) {
	retValue, err := instance.GetProperty("StorageNodeBandwidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageNodeIOPS sets the value of StorageNodeIOPS for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyStorageNodeIOPS(value uint64) (err error) {
	return instance.SetProperty("StorageNodeIOPS", value)
}

// GetStorageNodeIOPS gets the value of StorageNodeIOPS for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyStorageNodeIOPS() (value uint64, err error) {
	retValue, err := instance.GetProperty("StorageNodeIOPS")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageNodeLatency sets the value of StorageNodeLatency for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyStorageNodeLatency(value uint64) (err error) {
	return instance.SetProperty("StorageNodeLatency", value)
}

// GetStorageNodeLatency gets the value of StorageNodeLatency for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyStorageNodeLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("StorageNodeLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageNodeName sets the value of StorageNodeName for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyStorageNodeName(value string) (err error) {
	return instance.SetProperty("StorageNodeName", value)
}

// GetStorageNodeName gets the value of StorageNodeName for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyStorageNodeName() (value string, err error) {
	retValue, err := instance.GetProperty("StorageNodeName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeStamp sets the value of TimeStamp for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyTimeStamp(value uint64) (err error) {
	return instance.SetProperty("TimeStamp", value)
}

// GetTimeStamp gets the value of TimeStamp for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyTimeStamp() (value uint64, err error) {
	retValue, err := instance.GetProperty("TimeStamp")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolumeId sets the value of VolumeId for the instance
func (instance *MSFT_StorageQoSFlow) SetPropertyVolumeId(value string) (err error) {
	return instance.SetProperty("VolumeId", value)
}

// GetVolumeId gets the value of VolumeId for the instance
func (instance *MSFT_StorageQoSFlow) GetPropertyVolumeId() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
