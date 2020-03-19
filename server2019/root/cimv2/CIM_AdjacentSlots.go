// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_AdjacentSlots struct
type CIM_AdjacentSlots struct {
	*cim.WmiInstance

	//
	DistanceBetweenSlots float32

	//
	SharedSlots bool

	//
	SlotA CIM_Slot

	//
	SlotB CIM_Slot
}

func NewCIM_AdjacentSlotsEx1(instance *cim.WmiInstance) (newInstance *CIM_AdjacentSlots, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_AdjacentSlots{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_AdjacentSlotsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_AdjacentSlots, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_AdjacentSlots{
		WmiInstance: tmp,
	}
	return
}

// SetDistanceBetweenSlots sets the value of DistanceBetweenSlots for the instance
func (instance *CIM_AdjacentSlots) SetPropertyDistanceBetweenSlots(value float32) (err error) {
	return instance.SetProperty("DistanceBetweenSlots", value)
}

// GetDistanceBetweenSlots gets the value of DistanceBetweenSlots for the instance
func (instance *CIM_AdjacentSlots) GetPropertyDistanceBetweenSlots() (value float32, err error) {
	retValue, err := instance.GetProperty("DistanceBetweenSlots")
	if err != nil {
		return
	}
	value, ok := retValue.(float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSharedSlots sets the value of SharedSlots for the instance
func (instance *CIM_AdjacentSlots) SetPropertySharedSlots(value bool) (err error) {
	return instance.SetProperty("SharedSlots", value)
}

// GetSharedSlots gets the value of SharedSlots for the instance
func (instance *CIM_AdjacentSlots) GetPropertySharedSlots() (value bool, err error) {
	retValue, err := instance.GetProperty("SharedSlots")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSlotA sets the value of SlotA for the instance
func (instance *CIM_AdjacentSlots) SetPropertySlotA(value CIM_Slot) (err error) {
	return instance.SetProperty("SlotA", value)
}

// GetSlotA gets the value of SlotA for the instance
func (instance *CIM_AdjacentSlots) GetPropertySlotA() (value CIM_Slot, err error) {
	retValue, err := instance.GetProperty("SlotA")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Slot)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSlotB sets the value of SlotB for the instance
func (instance *CIM_AdjacentSlots) SetPropertySlotB(value CIM_Slot) (err error) {
	return instance.SetProperty("SlotB", value)
}

// GetSlotB gets the value of SlotB for the instance
func (instance *CIM_AdjacentSlots) GetPropertySlotB() (value CIM_Slot, err error) {
	retValue, err := instance.GetProperty("SlotB")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Slot)
	if !ok {
		// TODO: Set an error
	}
	return
}
