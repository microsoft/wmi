// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// ClusBfltHddBindingRecord struct
type ClusBfltHddBindingRecord struct {
	*cim.WmiInstance

	// Id.
	BindingId string

	// cDirtyPages.
	cDirtyPages uint32

	// cDirtySlots.
	cDirtySlots uint32

	// cPages.
	cPages uint32

	// cPagesL2.
	cPagesL2 uint32

	// cRefs.
	cRefs uint32

	// Device Guid.
	DeviceGuid string

	// DeviceSize.
	DeviceSize uint64

	// Flags.
	Flags uint32
}

func NewClusBfltHddBindingRecordEx1(instance *cim.WmiInstance) (newInstance *ClusBfltHddBindingRecord, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ClusBfltHddBindingRecord{
		WmiInstance: tmp,
	}
	return
}

func NewClusBfltHddBindingRecordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClusBfltHddBindingRecord, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClusBfltHddBindingRecord{
		WmiInstance: tmp,
	}
	return
}

// SetBindingId sets the value of BindingId for the instance
func (instance *ClusBfltHddBindingRecord) SetPropertyBindingId(value string) (err error) {
	return instance.SetProperty("BindingId", (value))
}

// GetBindingId gets the value of BindingId for the instance
func (instance *ClusBfltHddBindingRecord) GetPropertyBindingId() (value string, err error) {
	retValue, err := instance.GetProperty("BindingId")
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

// SetcDirtyPages sets the value of cDirtyPages for the instance
func (instance *ClusBfltHddBindingRecord) SetPropertycDirtyPages(value uint32) (err error) {
	return instance.SetProperty("cDirtyPages", (value))
}

// GetcDirtyPages gets the value of cDirtyPages for the instance
func (instance *ClusBfltHddBindingRecord) GetPropertycDirtyPages() (value uint32, err error) {
	retValue, err := instance.GetProperty("cDirtyPages")
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

// SetcDirtySlots sets the value of cDirtySlots for the instance
func (instance *ClusBfltHddBindingRecord) SetPropertycDirtySlots(value uint32) (err error) {
	return instance.SetProperty("cDirtySlots", (value))
}

// GetcDirtySlots gets the value of cDirtySlots for the instance
func (instance *ClusBfltHddBindingRecord) GetPropertycDirtySlots() (value uint32, err error) {
	retValue, err := instance.GetProperty("cDirtySlots")
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

// SetcPages sets the value of cPages for the instance
func (instance *ClusBfltHddBindingRecord) SetPropertycPages(value uint32) (err error) {
	return instance.SetProperty("cPages", (value))
}

// GetcPages gets the value of cPages for the instance
func (instance *ClusBfltHddBindingRecord) GetPropertycPages() (value uint32, err error) {
	retValue, err := instance.GetProperty("cPages")
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

// SetcPagesL2 sets the value of cPagesL2 for the instance
func (instance *ClusBfltHddBindingRecord) SetPropertycPagesL2(value uint32) (err error) {
	return instance.SetProperty("cPagesL2", (value))
}

// GetcPagesL2 gets the value of cPagesL2 for the instance
func (instance *ClusBfltHddBindingRecord) GetPropertycPagesL2() (value uint32, err error) {
	retValue, err := instance.GetProperty("cPagesL2")
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

// SetcRefs sets the value of cRefs for the instance
func (instance *ClusBfltHddBindingRecord) SetPropertycRefs(value uint32) (err error) {
	return instance.SetProperty("cRefs", (value))
}

// GetcRefs gets the value of cRefs for the instance
func (instance *ClusBfltHddBindingRecord) GetPropertycRefs() (value uint32, err error) {
	retValue, err := instance.GetProperty("cRefs")
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

// SetDeviceGuid sets the value of DeviceGuid for the instance
func (instance *ClusBfltHddBindingRecord) SetPropertyDeviceGuid(value string) (err error) {
	return instance.SetProperty("DeviceGuid", (value))
}

// GetDeviceGuid gets the value of DeviceGuid for the instance
func (instance *ClusBfltHddBindingRecord) GetPropertyDeviceGuid() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceGuid")
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

// SetDeviceSize sets the value of DeviceSize for the instance
func (instance *ClusBfltHddBindingRecord) SetPropertyDeviceSize(value uint64) (err error) {
	return instance.SetProperty("DeviceSize", (value))
}

// GetDeviceSize gets the value of DeviceSize for the instance
func (instance *ClusBfltHddBindingRecord) GetPropertyDeviceSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("DeviceSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFlags sets the value of Flags for the instance
func (instance *ClusBfltHddBindingRecord) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *ClusBfltHddBindingRecord) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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
