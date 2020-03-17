// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapter_RdmaMissingCounterInfo struct
type MSFT_NetAdapter_RdmaMissingCounterInfo struct {
	cim.WmiInstance

	//
	AcceptPerformanceCounterMissing bool

	//
	ActiveConnectionPerformanceCounterMissing bool

	//
	CompletionQueueErrorPerformanceCounterMissing bool

	//
	ConnectFailurePerformanceCounterMissing bool

	//
	ConnectionErrorPerformanceCounterMissing bool

	//
	ConnectPerformanceCounterMissing bool

	//
	RDMAInFramesPerformanceCounterMissing bool

	//
	RDMAInOctetsPerformanceCounterMissing bool

	//
	RDMAOutFramesPerformanceCounterMissing bool

	//
	RDMAOutOctetsPerformanceCounterMissing bool
}

// SetAcceptPerformanceCounterMissing sets the value of AcceptPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) SetPropertyAcceptPerformanceCounterMissing(value bool) (err error) {
	return instance.SetProperty("AcceptPerformanceCounterMissing", value)
}

// GetAcceptPerformanceCounterMissing gets the value of AcceptPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) GetPropertyAcceptPerformanceCounterMissing() (value bool, err error) {
	retValue, err := instance.GetProperty("AcceptPerformanceCounterMissing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActiveConnectionPerformanceCounterMissing sets the value of ActiveConnectionPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) SetPropertyActiveConnectionPerformanceCounterMissing(value bool) (err error) {
	return instance.SetProperty("ActiveConnectionPerformanceCounterMissing", value)
}

// GetActiveConnectionPerformanceCounterMissing gets the value of ActiveConnectionPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) GetPropertyActiveConnectionPerformanceCounterMissing() (value bool, err error) {
	retValue, err := instance.GetProperty("ActiveConnectionPerformanceCounterMissing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompletionQueueErrorPerformanceCounterMissing sets the value of CompletionQueueErrorPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) SetPropertyCompletionQueueErrorPerformanceCounterMissing(value bool) (err error) {
	return instance.SetProperty("CompletionQueueErrorPerformanceCounterMissing", value)
}

// GetCompletionQueueErrorPerformanceCounterMissing gets the value of CompletionQueueErrorPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) GetPropertyCompletionQueueErrorPerformanceCounterMissing() (value bool, err error) {
	retValue, err := instance.GetProperty("CompletionQueueErrorPerformanceCounterMissing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectFailurePerformanceCounterMissing sets the value of ConnectFailurePerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) SetPropertyConnectFailurePerformanceCounterMissing(value bool) (err error) {
	return instance.SetProperty("ConnectFailurePerformanceCounterMissing", value)
}

// GetConnectFailurePerformanceCounterMissing gets the value of ConnectFailurePerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) GetPropertyConnectFailurePerformanceCounterMissing() (value bool, err error) {
	retValue, err := instance.GetProperty("ConnectFailurePerformanceCounterMissing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionErrorPerformanceCounterMissing sets the value of ConnectionErrorPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) SetPropertyConnectionErrorPerformanceCounterMissing(value bool) (err error) {
	return instance.SetProperty("ConnectionErrorPerformanceCounterMissing", value)
}

// GetConnectionErrorPerformanceCounterMissing gets the value of ConnectionErrorPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) GetPropertyConnectionErrorPerformanceCounterMissing() (value bool, err error) {
	retValue, err := instance.GetProperty("ConnectionErrorPerformanceCounterMissing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectPerformanceCounterMissing sets the value of ConnectPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) SetPropertyConnectPerformanceCounterMissing(value bool) (err error) {
	return instance.SetProperty("ConnectPerformanceCounterMissing", value)
}

// GetConnectPerformanceCounterMissing gets the value of ConnectPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) GetPropertyConnectPerformanceCounterMissing() (value bool, err error) {
	retValue, err := instance.GetProperty("ConnectPerformanceCounterMissing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRDMAInFramesPerformanceCounterMissing sets the value of RDMAInFramesPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) SetPropertyRDMAInFramesPerformanceCounterMissing(value bool) (err error) {
	return instance.SetProperty("RDMAInFramesPerformanceCounterMissing", value)
}

// GetRDMAInFramesPerformanceCounterMissing gets the value of RDMAInFramesPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) GetPropertyRDMAInFramesPerformanceCounterMissing() (value bool, err error) {
	retValue, err := instance.GetProperty("RDMAInFramesPerformanceCounterMissing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRDMAInOctetsPerformanceCounterMissing sets the value of RDMAInOctetsPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) SetPropertyRDMAInOctetsPerformanceCounterMissing(value bool) (err error) {
	return instance.SetProperty("RDMAInOctetsPerformanceCounterMissing", value)
}

// GetRDMAInOctetsPerformanceCounterMissing gets the value of RDMAInOctetsPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) GetPropertyRDMAInOctetsPerformanceCounterMissing() (value bool, err error) {
	retValue, err := instance.GetProperty("RDMAInOctetsPerformanceCounterMissing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRDMAOutFramesPerformanceCounterMissing sets the value of RDMAOutFramesPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) SetPropertyRDMAOutFramesPerformanceCounterMissing(value bool) (err error) {
	return instance.SetProperty("RDMAOutFramesPerformanceCounterMissing", value)
}

// GetRDMAOutFramesPerformanceCounterMissing gets the value of RDMAOutFramesPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) GetPropertyRDMAOutFramesPerformanceCounterMissing() (value bool, err error) {
	retValue, err := instance.GetProperty("RDMAOutFramesPerformanceCounterMissing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRDMAOutOctetsPerformanceCounterMissing sets the value of RDMAOutOctetsPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) SetPropertyRDMAOutOctetsPerformanceCounterMissing(value bool) (err error) {
	return instance.SetProperty("RDMAOutOctetsPerformanceCounterMissing", value)
}

// GetRDMAOutOctetsPerformanceCounterMissing gets the value of RDMAOutOctetsPerformanceCounterMissing for the instance
func (instance *MSFT_NetAdapter_RdmaMissingCounterInfo) GetPropertyRDMAOutOctetsPerformanceCounterMissing() (value bool, err error) {
	retValue, err := instance.GetProperty("RDMAOutOctetsPerformanceCounterMissing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
