// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_IBPortStatistics struct
type CIM_IBPortStatistics struct {
	*CIM_NetworkPortStatistics

	//
	ExcessiveBufferOverrunErrors uint16

	//
	LinkDownedCounter uint16

	//
	LinkErrorRecoveryCounter uint16

	//
	LocalLinkIntegrityErrors uint16

	//
	PortRcvConstraintErrors uint16

	//
	PortRcvErrors uint16

	//
	PortRcvRemotePhysicalErrors uint16

	//
	PortRcvSwitchRelayErrors uint16

	//
	PortXmitConstraintErrors uint16

	//
	PortXmitDiscards uint16

	//
	SymbolErrorCounter uint16

	//
	VL15Dropped uint16
}

func NewCIM_IBPortStatisticsEx1(instance *cim.WmiInstance) (newInstance *CIM_IBPortStatistics, err error) {
	tmp, err := NewCIM_NetworkPortStatisticsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_IBPortStatistics{
		CIM_NetworkPortStatistics: tmp,
	}
	return
}

func NewCIM_IBPortStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_IBPortStatistics, err error) {
	tmp, err := NewCIM_NetworkPortStatisticsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_IBPortStatistics{
		CIM_NetworkPortStatistics: tmp,
	}
	return
}

// SetExcessiveBufferOverrunErrors sets the value of ExcessiveBufferOverrunErrors for the instance
func (instance *CIM_IBPortStatistics) SetPropertyExcessiveBufferOverrunErrors(value uint16) (err error) {
	return instance.SetProperty("ExcessiveBufferOverrunErrors", value)
}

// GetExcessiveBufferOverrunErrors gets the value of ExcessiveBufferOverrunErrors for the instance
func (instance *CIM_IBPortStatistics) GetPropertyExcessiveBufferOverrunErrors() (value uint16, err error) {
	retValue, err := instance.GetProperty("ExcessiveBufferOverrunErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLinkDownedCounter sets the value of LinkDownedCounter for the instance
func (instance *CIM_IBPortStatistics) SetPropertyLinkDownedCounter(value uint16) (err error) {
	return instance.SetProperty("LinkDownedCounter", value)
}

// GetLinkDownedCounter gets the value of LinkDownedCounter for the instance
func (instance *CIM_IBPortStatistics) GetPropertyLinkDownedCounter() (value uint16, err error) {
	retValue, err := instance.GetProperty("LinkDownedCounter")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLinkErrorRecoveryCounter sets the value of LinkErrorRecoveryCounter for the instance
func (instance *CIM_IBPortStatistics) SetPropertyLinkErrorRecoveryCounter(value uint16) (err error) {
	return instance.SetProperty("LinkErrorRecoveryCounter", value)
}

// GetLinkErrorRecoveryCounter gets the value of LinkErrorRecoveryCounter for the instance
func (instance *CIM_IBPortStatistics) GetPropertyLinkErrorRecoveryCounter() (value uint16, err error) {
	retValue, err := instance.GetProperty("LinkErrorRecoveryCounter")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalLinkIntegrityErrors sets the value of LocalLinkIntegrityErrors for the instance
func (instance *CIM_IBPortStatistics) SetPropertyLocalLinkIntegrityErrors(value uint16) (err error) {
	return instance.SetProperty("LocalLinkIntegrityErrors", value)
}

// GetLocalLinkIntegrityErrors gets the value of LocalLinkIntegrityErrors for the instance
func (instance *CIM_IBPortStatistics) GetPropertyLocalLinkIntegrityErrors() (value uint16, err error) {
	retValue, err := instance.GetProperty("LocalLinkIntegrityErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortRcvConstraintErrors sets the value of PortRcvConstraintErrors for the instance
func (instance *CIM_IBPortStatistics) SetPropertyPortRcvConstraintErrors(value uint16) (err error) {
	return instance.SetProperty("PortRcvConstraintErrors", value)
}

// GetPortRcvConstraintErrors gets the value of PortRcvConstraintErrors for the instance
func (instance *CIM_IBPortStatistics) GetPropertyPortRcvConstraintErrors() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortRcvConstraintErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortRcvErrors sets the value of PortRcvErrors for the instance
func (instance *CIM_IBPortStatistics) SetPropertyPortRcvErrors(value uint16) (err error) {
	return instance.SetProperty("PortRcvErrors", value)
}

// GetPortRcvErrors gets the value of PortRcvErrors for the instance
func (instance *CIM_IBPortStatistics) GetPropertyPortRcvErrors() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortRcvErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortRcvRemotePhysicalErrors sets the value of PortRcvRemotePhysicalErrors for the instance
func (instance *CIM_IBPortStatistics) SetPropertyPortRcvRemotePhysicalErrors(value uint16) (err error) {
	return instance.SetProperty("PortRcvRemotePhysicalErrors", value)
}

// GetPortRcvRemotePhysicalErrors gets the value of PortRcvRemotePhysicalErrors for the instance
func (instance *CIM_IBPortStatistics) GetPropertyPortRcvRemotePhysicalErrors() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortRcvRemotePhysicalErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortRcvSwitchRelayErrors sets the value of PortRcvSwitchRelayErrors for the instance
func (instance *CIM_IBPortStatistics) SetPropertyPortRcvSwitchRelayErrors(value uint16) (err error) {
	return instance.SetProperty("PortRcvSwitchRelayErrors", value)
}

// GetPortRcvSwitchRelayErrors gets the value of PortRcvSwitchRelayErrors for the instance
func (instance *CIM_IBPortStatistics) GetPropertyPortRcvSwitchRelayErrors() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortRcvSwitchRelayErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortXmitConstraintErrors sets the value of PortXmitConstraintErrors for the instance
func (instance *CIM_IBPortStatistics) SetPropertyPortXmitConstraintErrors(value uint16) (err error) {
	return instance.SetProperty("PortXmitConstraintErrors", value)
}

// GetPortXmitConstraintErrors gets the value of PortXmitConstraintErrors for the instance
func (instance *CIM_IBPortStatistics) GetPropertyPortXmitConstraintErrors() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortXmitConstraintErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortXmitDiscards sets the value of PortXmitDiscards for the instance
func (instance *CIM_IBPortStatistics) SetPropertyPortXmitDiscards(value uint16) (err error) {
	return instance.SetProperty("PortXmitDiscards", value)
}

// GetPortXmitDiscards gets the value of PortXmitDiscards for the instance
func (instance *CIM_IBPortStatistics) GetPropertyPortXmitDiscards() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortXmitDiscards")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSymbolErrorCounter sets the value of SymbolErrorCounter for the instance
func (instance *CIM_IBPortStatistics) SetPropertySymbolErrorCounter(value uint16) (err error) {
	return instance.SetProperty("SymbolErrorCounter", value)
}

// GetSymbolErrorCounter gets the value of SymbolErrorCounter for the instance
func (instance *CIM_IBPortStatistics) GetPropertySymbolErrorCounter() (value uint16, err error) {
	retValue, err := instance.GetProperty("SymbolErrorCounter")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVL15Dropped sets the value of VL15Dropped for the instance
func (instance *CIM_IBPortStatistics) SetPropertyVL15Dropped(value uint16) (err error) {
	return instance.SetProperty("VL15Dropped", value)
}

// GetVL15Dropped gets the value of VL15Dropped for the instance
func (instance *CIM_IBPortStatistics) GetPropertyVL15Dropped() (value uint16, err error) {
	retValue, err := instance.GetProperty("VL15Dropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
