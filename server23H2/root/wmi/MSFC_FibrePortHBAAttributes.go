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

// MSFC_FibrePortHBAAttributes struct
type MSFC_FibrePortHBAAttributes struct {
	*cim.WmiInstance

	//
	Active bool

	//
	Attributes MSFC_HBAPortAttributesResults

	//
	HBAStatus uint32

	//
	InstanceName string

	//
	UniquePortId uint64
}

func NewMSFC_FibrePortHBAAttributesEx1(instance *cim.WmiInstance) (newInstance *MSFC_FibrePortHBAAttributes, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_FibrePortHBAAttributes{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_FibrePortHBAAttributesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_FibrePortHBAAttributes, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_FibrePortHBAAttributes{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFC_FibrePortHBAAttributes) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFC_FibrePortHBAAttributes) GetPropertyActive() (value bool, err error) {
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

// SetAttributes sets the value of Attributes for the instance
func (instance *MSFC_FibrePortHBAAttributes) SetPropertyAttributes(value MSFC_HBAPortAttributesResults) (err error) {
	return instance.SetProperty("Attributes", (value))
}

// GetAttributes gets the value of Attributes for the instance
func (instance *MSFC_FibrePortHBAAttributes) GetPropertyAttributes() (value MSFC_HBAPortAttributesResults, err error) {
	retValue, err := instance.GetProperty("Attributes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFC_HBAPortAttributesResults)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFC_HBAPortAttributesResults is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFC_HBAPortAttributesResults(valuetmp)

	return
}

// SetHBAStatus sets the value of HBAStatus for the instance
func (instance *MSFC_FibrePortHBAAttributes) SetPropertyHBAStatus(value uint32) (err error) {
	return instance.SetProperty("HBAStatus", (value))
}

// GetHBAStatus gets the value of HBAStatus for the instance
func (instance *MSFC_FibrePortHBAAttributes) GetPropertyHBAStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("HBAStatus")
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
func (instance *MSFC_FibrePortHBAAttributes) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFC_FibrePortHBAAttributes) GetPropertyInstanceName() (value string, err error) {
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

// SetUniquePortId sets the value of UniquePortId for the instance
func (instance *MSFC_FibrePortHBAAttributes) SetPropertyUniquePortId(value uint64) (err error) {
	return instance.SetProperty("UniquePortId", (value))
}

// GetUniquePortId gets the value of UniquePortId for the instance
func (instance *MSFC_FibrePortHBAAttributes) GetPropertyUniquePortId() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniquePortId")
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
