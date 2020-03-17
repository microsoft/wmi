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

// MSFT_StorageQoSPolicy struct
type MSFT_StorageQoSPolicy struct {
	cim.WmiInstance

	//
	BandwidthLimit uint64

	//
	Name string

	//
	ParentPolicy string

	//
	PolicyId string

	//
	PolicyType uint16

	//
	Status uint16

	//
	ThroughputLimit uint64

	//
	ThroughputReservation uint64
}

// SetBandwidthLimit sets the value of BandwidthLimit for the instance
func (instance *MSFT_StorageQoSPolicy) SetPropertyBandwidthLimit(value uint64) (err error) {
	return instance.SetProperty("BandwidthLimit", value)
}

// GetBandwidthLimit gets the value of BandwidthLimit for the instance
func (instance *MSFT_StorageQoSPolicy) GetPropertyBandwidthLimit() (value uint64, err error) {
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

// SetName sets the value of Name for the instance
func (instance *MSFT_StorageQoSPolicy) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_StorageQoSPolicy) GetPropertyName() (value string, err error) {
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

// SetParentPolicy sets the value of ParentPolicy for the instance
func (instance *MSFT_StorageQoSPolicy) SetPropertyParentPolicy(value string) (err error) {
	return instance.SetProperty("ParentPolicy", value)
}

// GetParentPolicy gets the value of ParentPolicy for the instance
func (instance *MSFT_StorageQoSPolicy) GetPropertyParentPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("ParentPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyId sets the value of PolicyId for the instance
func (instance *MSFT_StorageQoSPolicy) SetPropertyPolicyId(value string) (err error) {
	return instance.SetProperty("PolicyId", value)
}

// GetPolicyId gets the value of PolicyId for the instance
func (instance *MSFT_StorageQoSPolicy) GetPropertyPolicyId() (value string, err error) {
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

// SetPolicyType sets the value of PolicyType for the instance
func (instance *MSFT_StorageQoSPolicy) SetPropertyPolicyType(value uint16) (err error) {
	return instance.SetProperty("PolicyType", value)
}

// GetPolicyType gets the value of PolicyType for the instance
func (instance *MSFT_StorageQoSPolicy) GetPropertyPolicyType() (value uint16, err error) {
	retValue, err := instance.GetProperty("PolicyType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSFT_StorageQoSPolicy) SetPropertyStatus(value uint16) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_StorageQoSPolicy) GetPropertyStatus() (value uint16, err error) {
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

// SetThroughputLimit sets the value of ThroughputLimit for the instance
func (instance *MSFT_StorageQoSPolicy) SetPropertyThroughputLimit(value uint64) (err error) {
	return instance.SetProperty("ThroughputLimit", value)
}

// GetThroughputLimit gets the value of ThroughputLimit for the instance
func (instance *MSFT_StorageQoSPolicy) GetPropertyThroughputLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("ThroughputLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThroughputReservation sets the value of ThroughputReservation for the instance
func (instance *MSFT_StorageQoSPolicy) SetPropertyThroughputReservation(value uint64) (err error) {
	return instance.SetProperty("ThroughputReservation", value)
}

// GetThroughputReservation gets the value of ThroughputReservation for the instance
func (instance *MSFT_StorageQoSPolicy) GetPropertyThroughputReservation() (value uint64, err error) {
	retValue, err := instance.GetProperty("ThroughputReservation")
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

// <param name="BandwidthLimit" type="uint64 "></param>
// <param name="Limit" type="uint64 "></param>
// <param name="NewName" type="string "></param>
// <param name="Reservation" type="uint64 "></param>

// <param name="ReturnValue" type="int32 "></param>
func (instance *MSFT_StorageQoSPolicy) SetAttributes( /* IN */ NewName string,
	/* IN */ Limit uint64,
	/* IN */ Reservation uint64,
	/* IN */ BandwidthLimit uint64) (result int32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetAttributes", NewName, Limit, Reservation, BandwidthLimit)
	if err != nil {
		return
	}
	result = int32(retVal)
	return

}

//

// <param name="ReturnValue" type="int32 "></param>
func (instance *MSFT_StorageQoSPolicy) DeletePolicy() (result int32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DeletePolicy")
	if err != nil {
		return
	}
	result = int32(retVal)
	return

}
