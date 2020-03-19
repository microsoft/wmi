// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapter_RdmaStatistics struct
type MSFT_NetAdapter_RdmaStatistics struct {
	*cim.WmiInstance

	//
	AcceptedConnections uint64

	//
	ActiveConnections uint64

	//
	CompletionQueueErrors uint64

	//
	ConnectionErrors uint64

	//
	FailedConnectionAttempts uint64

	//
	InboundBytes uint64

	//
	InboundFrames uint64

	//
	InitiatedConnections uint64

	//
	OutboundBytes uint64

	//
	OutboundFrames uint64
}

func NewMSFT_NetAdapter_RdmaStatisticsEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapter_RdmaStatistics, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_RdmaStatistics{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapter_RdmaStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapter_RdmaStatistics, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_RdmaStatistics{
		WmiInstance: tmp,
	}
	return
}

// SetAcceptedConnections sets the value of AcceptedConnections for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) SetPropertyAcceptedConnections(value uint64) (err error) {
	return instance.SetProperty("AcceptedConnections", value)
}

// GetAcceptedConnections gets the value of AcceptedConnections for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) GetPropertyAcceptedConnections() (value uint64, err error) {
	retValue, err := instance.GetProperty("AcceptedConnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActiveConnections sets the value of ActiveConnections for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) SetPropertyActiveConnections(value uint64) (err error) {
	return instance.SetProperty("ActiveConnections", value)
}

// GetActiveConnections gets the value of ActiveConnections for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) GetPropertyActiveConnections() (value uint64, err error) {
	retValue, err := instance.GetProperty("ActiveConnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompletionQueueErrors sets the value of CompletionQueueErrors for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) SetPropertyCompletionQueueErrors(value uint64) (err error) {
	return instance.SetProperty("CompletionQueueErrors", value)
}

// GetCompletionQueueErrors gets the value of CompletionQueueErrors for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) GetPropertyCompletionQueueErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("CompletionQueueErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionErrors sets the value of ConnectionErrors for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) SetPropertyConnectionErrors(value uint64) (err error) {
	return instance.SetProperty("ConnectionErrors", value)
}

// GetConnectionErrors gets the value of ConnectionErrors for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) GetPropertyConnectionErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("ConnectionErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailedConnectionAttempts sets the value of FailedConnectionAttempts for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) SetPropertyFailedConnectionAttempts(value uint64) (err error) {
	return instance.SetProperty("FailedConnectionAttempts", value)
}

// GetFailedConnectionAttempts gets the value of FailedConnectionAttempts for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) GetPropertyFailedConnectionAttempts() (value uint64, err error) {
	retValue, err := instance.GetProperty("FailedConnectionAttempts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInboundBytes sets the value of InboundBytes for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) SetPropertyInboundBytes(value uint64) (err error) {
	return instance.SetProperty("InboundBytes", value)
}

// GetInboundBytes gets the value of InboundBytes for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) GetPropertyInboundBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("InboundBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInboundFrames sets the value of InboundFrames for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) SetPropertyInboundFrames(value uint64) (err error) {
	return instance.SetProperty("InboundFrames", value)
}

// GetInboundFrames gets the value of InboundFrames for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) GetPropertyInboundFrames() (value uint64, err error) {
	retValue, err := instance.GetProperty("InboundFrames")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInitiatedConnections sets the value of InitiatedConnections for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) SetPropertyInitiatedConnections(value uint64) (err error) {
	return instance.SetProperty("InitiatedConnections", value)
}

// GetInitiatedConnections gets the value of InitiatedConnections for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) GetPropertyInitiatedConnections() (value uint64, err error) {
	retValue, err := instance.GetProperty("InitiatedConnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOutboundBytes sets the value of OutboundBytes for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) SetPropertyOutboundBytes(value uint64) (err error) {
	return instance.SetProperty("OutboundBytes", value)
}

// GetOutboundBytes gets the value of OutboundBytes for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) GetPropertyOutboundBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutboundBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOutboundFrames sets the value of OutboundFrames for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) SetPropertyOutboundFrames(value uint64) (err error) {
	return instance.SetProperty("OutboundFrames", value)
}

// GetOutboundFrames gets the value of OutboundFrames for the instance
func (instance *MSFT_NetAdapter_RdmaStatistics) GetPropertyOutboundFrames() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutboundFrames")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
