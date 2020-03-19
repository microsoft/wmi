// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_EthernetSwitchPortBandwidthData struct
type Msvm_EthernetSwitchPortBandwidthData struct {
	*Msvm_EthernetPortData

	//
	CurrentBandwidthReservationPercentage uint32
}

func NewMsvm_EthernetSwitchPortBandwidthDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortBandwidthData, err error) {
	tmp, err := NewMsvm_EthernetPortDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortBandwidthData{
		Msvm_EthernetPortData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchPortBandwidthDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchPortBandwidthData, err error) {
	tmp, err := NewMsvm_EthernetPortDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortBandwidthData{
		Msvm_EthernetPortData: tmp,
	}
	return
}

// SetCurrentBandwidthReservationPercentage sets the value of CurrentBandwidthReservationPercentage for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthData) SetPropertyCurrentBandwidthReservationPercentage(value uint32) (err error) {
	return instance.SetProperty("CurrentBandwidthReservationPercentage", value)
}

// GetCurrentBandwidthReservationPercentage gets the value of CurrentBandwidthReservationPercentage for the instance
func (instance *Msvm_EthernetSwitchPortBandwidthData) GetPropertyCurrentBandwidthReservationPercentage() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentBandwidthReservationPercentage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchPortBandwidthData) GetRelatedEthernetSwitchPort() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchPort")
}
