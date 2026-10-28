// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// TeeDriverWmiInfo struct
type TeeDriverWmiInfo struct {
	*cim.WmiInstance

	//
	Active bool

	// AsyncEvents model support.
	AsyncEvents uint32

	// Client Dma support.
	ClientDmaSupport uint32

	// Current D0ix state.
	D0ix uint32

	// Current Dma state.
	DmaEnabled uint32

	// FwNotfication protocol support.
	FwNotfication uint32

	//
	InstanceName string

	// Current pg state.
	Pg bool

	// Current state of Pgi_Alivness.
	Pgi_Alivness uint32

	// RTD3 flows support.
	RTD3 uint32
}

func NewTeeDriverWmiInfoEx1(instance *cim.WmiInstance) (newInstance *TeeDriverWmiInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &TeeDriverWmiInfo{
		WmiInstance: tmp,
	}
	return
}

func NewTeeDriverWmiInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TeeDriverWmiInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TeeDriverWmiInfo{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *TeeDriverWmiInfo) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *TeeDriverWmiInfo) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetAsyncEvents sets the value of AsyncEvents for the instance
func (instance *TeeDriverWmiInfo) SetPropertyAsyncEvents(value uint32) (err error) {
	return instance.SetProperty("AsyncEvents", (value))
}

// GetAsyncEvents gets the value of AsyncEvents for the instance
func (instance *TeeDriverWmiInfo) GetPropertyAsyncEvents() (value uint32, err error) {
	retValue, err := instance.GetProperty("AsyncEvents")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetClientDmaSupport sets the value of ClientDmaSupport for the instance
func (instance *TeeDriverWmiInfo) SetPropertyClientDmaSupport(value uint32) (err error) {
	return instance.SetProperty("ClientDmaSupport", (value))
}

// GetClientDmaSupport gets the value of ClientDmaSupport for the instance
func (instance *TeeDriverWmiInfo) GetPropertyClientDmaSupport() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClientDmaSupport")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetD0ix sets the value of D0ix for the instance
func (instance *TeeDriverWmiInfo) SetPropertyD0ix(value uint32) (err error) {
	return instance.SetProperty("D0ix", (value))
}

// GetD0ix gets the value of D0ix for the instance
func (instance *TeeDriverWmiInfo) GetPropertyD0ix() (value uint32, err error) {
	retValue, err := instance.GetProperty("D0ix")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetDmaEnabled sets the value of DmaEnabled for the instance
func (instance *TeeDriverWmiInfo) SetPropertyDmaEnabled(value uint32) (err error) {
	return instance.SetProperty("DmaEnabled", (value))
}

// GetDmaEnabled gets the value of DmaEnabled for the instance
func (instance *TeeDriverWmiInfo) GetPropertyDmaEnabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("DmaEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetFwNotfication sets the value of FwNotfication for the instance
func (instance *TeeDriverWmiInfo) SetPropertyFwNotfication(value uint32) (err error) {
	return instance.SetProperty("FwNotfication", (value))
}

// GetFwNotfication gets the value of FwNotfication for the instance
func (instance *TeeDriverWmiInfo) GetPropertyFwNotfication() (value uint32, err error) {
	retValue, err := instance.GetProperty("FwNotfication")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *TeeDriverWmiInfo) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *TeeDriverWmiInfo) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetPg sets the value of Pg for the instance
func (instance *TeeDriverWmiInfo) SetPropertyPg(value bool) (err error) {
	return instance.SetProperty("Pg", (value))
}

// GetPg gets the value of Pg for the instance
func (instance *TeeDriverWmiInfo) GetPropertyPg() (value bool, err error) {
	retValue, err := instance.GetProperty("Pg")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetPgi_Alivness sets the value of Pgi_Alivness for the instance
func (instance *TeeDriverWmiInfo) SetPropertyPgi_Alivness(value uint32) (err error) {
	return instance.SetProperty("Pgi_Alivness", (value))
}

// GetPgi_Alivness gets the value of Pgi_Alivness for the instance
func (instance *TeeDriverWmiInfo) GetPropertyPgi_Alivness() (value uint32, err error) {
	retValue, err := instance.GetProperty("Pgi_Alivness")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetRTD3 sets the value of RTD3 for the instance
func (instance *TeeDriverWmiInfo) SetPropertyRTD3(value uint32) (err error) {
	return instance.SetProperty("RTD3", (value))
}

// GetRTD3 gets the value of RTD3 for the instance
func (instance *TeeDriverWmiInfo) GetPropertyRTD3() (value uint32, err error) {
	retValue, err := instance.GetProperty("RTD3")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
