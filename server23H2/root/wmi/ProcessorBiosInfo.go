// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ProcessorBiosInfo struct
type ProcessorBiosInfo struct {
	*MSProcessorClass

	//
	Active bool

	//
	ApicId uint32

	//
	InstanceName string

	//
	NtNumber uint32

	//
	PBlk uint32

	//
	PBlkLen uint32

	//
	Pct AcpiPct

	//
	ProcessorId uint32

	//
	Pss AcpiPss
}

func NewProcessorBiosInfoEx1(instance *cim.WmiInstance) (newInstance *ProcessorBiosInfo, err error) {
	tmp, err := NewMSProcessorClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ProcessorBiosInfo{
		MSProcessorClass: tmp,
	}
	return
}

func NewProcessorBiosInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProcessorBiosInfo, err error) {
	tmp, err := NewMSProcessorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProcessorBiosInfo{
		MSProcessorClass: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *ProcessorBiosInfo) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *ProcessorBiosInfo) GetPropertyActive() (value bool, err error) {
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

// SetApicId sets the value of ApicId for the instance
func (instance *ProcessorBiosInfo) SetPropertyApicId(value uint32) (err error) {
	return instance.SetProperty("ApicId", (value))
}

// GetApicId gets the value of ApicId for the instance
func (instance *ProcessorBiosInfo) GetPropertyApicId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ApicId")
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
func (instance *ProcessorBiosInfo) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ProcessorBiosInfo) GetPropertyInstanceName() (value string, err error) {
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

// SetNtNumber sets the value of NtNumber for the instance
func (instance *ProcessorBiosInfo) SetPropertyNtNumber(value uint32) (err error) {
	return instance.SetProperty("NtNumber", (value))
}

// GetNtNumber gets the value of NtNumber for the instance
func (instance *ProcessorBiosInfo) GetPropertyNtNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("NtNumber")
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

// SetPBlk sets the value of PBlk for the instance
func (instance *ProcessorBiosInfo) SetPropertyPBlk(value uint32) (err error) {
	return instance.SetProperty("PBlk", (value))
}

// GetPBlk gets the value of PBlk for the instance
func (instance *ProcessorBiosInfo) GetPropertyPBlk() (value uint32, err error) {
	retValue, err := instance.GetProperty("PBlk")
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

// SetPBlkLen sets the value of PBlkLen for the instance
func (instance *ProcessorBiosInfo) SetPropertyPBlkLen(value uint32) (err error) {
	return instance.SetProperty("PBlkLen", (value))
}

// GetPBlkLen gets the value of PBlkLen for the instance
func (instance *ProcessorBiosInfo) GetPropertyPBlkLen() (value uint32, err error) {
	retValue, err := instance.GetProperty("PBlkLen")
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

// SetPct sets the value of Pct for the instance
func (instance *ProcessorBiosInfo) SetPropertyPct(value AcpiPct) (err error) {
	return instance.SetProperty("Pct", (value))
}

// GetPct gets the value of Pct for the instance
func (instance *ProcessorBiosInfo) GetPropertyPct() (value AcpiPct, err error) {
	retValue, err := instance.GetProperty("Pct")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AcpiPct)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AcpiPct is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AcpiPct(valuetmp)

	return
}

// SetProcessorId sets the value of ProcessorId for the instance
func (instance *ProcessorBiosInfo) SetPropertyProcessorId(value uint32) (err error) {
	return instance.SetProperty("ProcessorId", (value))
}

// GetProcessorId gets the value of ProcessorId for the instance
func (instance *ProcessorBiosInfo) GetPropertyProcessorId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessorId")
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

// SetPss sets the value of Pss for the instance
func (instance *ProcessorBiosInfo) SetPropertyPss(value AcpiPss) (err error) {
	return instance.SetProperty("Pss", (value))
}

// GetPss gets the value of Pss for the instance
func (instance *ProcessorBiosInfo) GetPropertyPss() (value AcpiPss, err error) {
	retValue, err := instance.GetProperty("Pss")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AcpiPss)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AcpiPss is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AcpiPss(valuetmp)

	return
}
