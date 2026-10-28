// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Lenovo_GetNumberOfAvailableSlots struct
type Lenovo_GetNumberOfAvailableSlots struct {
	*Lenovo_CertificateStorageElement

	//
	Active bool

	//
	InstanceName string

	// Get number of available slots
	NumberOfAvailableSlots uint32
}

func NewLenovo_GetNumberOfAvailableSlotsEx1(instance *cim.WmiInstance) (newInstance *Lenovo_GetNumberOfAvailableSlots, err error) {
	tmp, err := NewLenovo_CertificateStorageElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Lenovo_GetNumberOfAvailableSlots{
		Lenovo_CertificateStorageElement: tmp,
	}
	return
}

func NewLenovo_GetNumberOfAvailableSlotsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Lenovo_GetNumberOfAvailableSlots, err error) {
	tmp, err := NewLenovo_CertificateStorageElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Lenovo_GetNumberOfAvailableSlots{
		Lenovo_CertificateStorageElement: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *Lenovo_GetNumberOfAvailableSlots) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *Lenovo_GetNumberOfAvailableSlots) GetPropertyActive() (value bool, err error) {
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *Lenovo_GetNumberOfAvailableSlots) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *Lenovo_GetNumberOfAvailableSlots) GetPropertyInstanceName() (value string, err error) {
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

// SetNumberOfAvailableSlots sets the value of NumberOfAvailableSlots for the instance
func (instance *Lenovo_GetNumberOfAvailableSlots) SetPropertyNumberOfAvailableSlots(value uint32) (err error) {
	return instance.SetProperty("NumberOfAvailableSlots", (value))
}

// GetNumberOfAvailableSlots gets the value of NumberOfAvailableSlots for the instance
func (instance *Lenovo_GetNumberOfAvailableSlots) GetPropertyNumberOfAvailableSlots() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfAvailableSlots")
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
