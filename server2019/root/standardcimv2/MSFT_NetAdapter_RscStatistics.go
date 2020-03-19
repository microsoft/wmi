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

// MSFT_NetAdapter_RscStatistics struct
type MSFT_NetAdapter_RscStatistics struct {
	*cim.WmiInstance

	//
	CoalescedBytes uint64

	//
	CoalescedPackets uint64

	//
	CoalescingEvents uint64

	//
	CoalescingExceptions uint64
}

func NewMSFT_NetAdapter_RscStatisticsEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapter_RscStatistics, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_RscStatistics{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapter_RscStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapter_RscStatistics, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_RscStatistics{
		WmiInstance: tmp,
	}
	return
}

// SetCoalescedBytes sets the value of CoalescedBytes for the instance
func (instance *MSFT_NetAdapter_RscStatistics) SetPropertyCoalescedBytes(value uint64) (err error) {
	return instance.SetProperty("CoalescedBytes", value)
}

// GetCoalescedBytes gets the value of CoalescedBytes for the instance
func (instance *MSFT_NetAdapter_RscStatistics) GetPropertyCoalescedBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("CoalescedBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCoalescedPackets sets the value of CoalescedPackets for the instance
func (instance *MSFT_NetAdapter_RscStatistics) SetPropertyCoalescedPackets(value uint64) (err error) {
	return instance.SetProperty("CoalescedPackets", value)
}

// GetCoalescedPackets gets the value of CoalescedPackets for the instance
func (instance *MSFT_NetAdapter_RscStatistics) GetPropertyCoalescedPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("CoalescedPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCoalescingEvents sets the value of CoalescingEvents for the instance
func (instance *MSFT_NetAdapter_RscStatistics) SetPropertyCoalescingEvents(value uint64) (err error) {
	return instance.SetProperty("CoalescingEvents", value)
}

// GetCoalescingEvents gets the value of CoalescingEvents for the instance
func (instance *MSFT_NetAdapter_RscStatistics) GetPropertyCoalescingEvents() (value uint64, err error) {
	retValue, err := instance.GetProperty("CoalescingEvents")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCoalescingExceptions sets the value of CoalescingExceptions for the instance
func (instance *MSFT_NetAdapter_RscStatistics) SetPropertyCoalescingExceptions(value uint64) (err error) {
	return instance.SetProperty("CoalescingExceptions", value)
}

// GetCoalescingExceptions gets the value of CoalescingExceptions for the instance
func (instance *MSFT_NetAdapter_RscStatistics) GetPropertyCoalescingExceptions() (value uint64, err error) {
	retValue, err := instance.GetProperty("CoalescingExceptions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
