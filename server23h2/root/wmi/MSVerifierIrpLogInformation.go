// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// MSVerifierIrpLogInformation struct
type MSVerifierIrpLogInformation struct {
	*cim.WmiInstance

	//
	Active bool

	// DeviceType
	DeviceType uint32

	//
	Entries []MSVerifierIrpLogEntry

	//
	EntryCount uint32

	//
	InstanceName string
}

func NewMSVerifierIrpLogInformationEx1(instance *cim.WmiInstance) (newInstance *MSVerifierIrpLogInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSVerifierIrpLogInformation{
		WmiInstance: tmp,
	}
	return
}

func NewMSVerifierIrpLogInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSVerifierIrpLogInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSVerifierIrpLogInformation{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSVerifierIrpLogInformation) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSVerifierIrpLogInformation) GetPropertyActive() (value bool, err error) {
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

// SetDeviceType sets the value of DeviceType for the instance
func (instance *MSVerifierIrpLogInformation) SetPropertyDeviceType(value uint32) (err error) {
	return instance.SetProperty("DeviceType", (value))
}

// GetDeviceType gets the value of DeviceType for the instance
func (instance *MSVerifierIrpLogInformation) GetPropertyDeviceType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceType")
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

// SetEntries sets the value of Entries for the instance
func (instance *MSVerifierIrpLogInformation) SetPropertyEntries(value []MSVerifierIrpLogEntry) (err error) {
	return instance.SetProperty("Entries", (value))
}

// GetEntries gets the value of Entries for the instance
func (instance *MSVerifierIrpLogInformation) GetPropertyEntries() (value []MSVerifierIrpLogEntry, err error) {
	retValue, err := instance.GetProperty("Entries")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSVerifierIrpLogEntry)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSVerifierIrpLogEntry is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSVerifierIrpLogEntry(valuetmp))
	}

	return
}

// SetEntryCount sets the value of EntryCount for the instance
func (instance *MSVerifierIrpLogInformation) SetPropertyEntryCount(value uint32) (err error) {
	return instance.SetProperty("EntryCount", (value))
}

// GetEntryCount gets the value of EntryCount for the instance
func (instance *MSVerifierIrpLogInformation) GetPropertyEntryCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("EntryCount")
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
func (instance *MSVerifierIrpLogInformation) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSVerifierIrpLogInformation) GetPropertyInstanceName() (value string, err error) {
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
