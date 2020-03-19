// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_NetworkAdapter struct
type CIM_NetworkAdapter struct {
	*CIM_LogicalDevice

	//
	AutoSense bool

	//
	MaxSpeed uint64

	//
	NetworkAddresses []string

	//
	PermanentAddress string

	//
	Speed uint64
}

func NewCIM_NetworkAdapterEx1(instance *cim.WmiInstance) (newInstance *CIM_NetworkAdapter, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_NetworkAdapter{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewCIM_NetworkAdapterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_NetworkAdapter, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_NetworkAdapter{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetAutoSense sets the value of AutoSense for the instance
func (instance *CIM_NetworkAdapter) SetPropertyAutoSense(value bool) (err error) {
	return instance.SetProperty("AutoSense", value)
}

// GetAutoSense gets the value of AutoSense for the instance
func (instance *CIM_NetworkAdapter) GetPropertyAutoSense() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoSense")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxSpeed sets the value of MaxSpeed for the instance
func (instance *CIM_NetworkAdapter) SetPropertyMaxSpeed(value uint64) (err error) {
	return instance.SetProperty("MaxSpeed", value)
}

// GetMaxSpeed gets the value of MaxSpeed for the instance
func (instance *CIM_NetworkAdapter) GetPropertyMaxSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkAddresses sets the value of NetworkAddresses for the instance
func (instance *CIM_NetworkAdapter) SetPropertyNetworkAddresses(value []string) (err error) {
	return instance.SetProperty("NetworkAddresses", value)
}

// GetNetworkAddresses gets the value of NetworkAddresses for the instance
func (instance *CIM_NetworkAdapter) GetPropertyNetworkAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("NetworkAddresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPermanentAddress sets the value of PermanentAddress for the instance
func (instance *CIM_NetworkAdapter) SetPropertyPermanentAddress(value string) (err error) {
	return instance.SetProperty("PermanentAddress", value)
}

// GetPermanentAddress gets the value of PermanentAddress for the instance
func (instance *CIM_NetworkAdapter) GetPropertyPermanentAddress() (value string, err error) {
	retValue, err := instance.GetProperty("PermanentAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpeed sets the value of Speed for the instance
func (instance *CIM_NetworkAdapter) SetPropertySpeed(value uint64) (err error) {
	return instance.SetProperty("Speed", value)
}

// GetSpeed gets the value of Speed for the instance
func (instance *CIM_NetworkAdapter) GetPropertySpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("Speed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
