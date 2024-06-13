// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_ObjectHeader struct
type MSNdis_ObjectHeader struct {
	*MSNdis

	//
	Revision uint8

	//
	Size uint16

	//
	Type uint8
}

func NewMSNdis_ObjectHeaderEx1(instance *cim.WmiInstance) (newInstance *MSNdis_ObjectHeader, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ObjectHeader{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_ObjectHeaderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_ObjectHeader, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ObjectHeader{
		MSNdis: tmp,
	}
	return
}

// SetRevision sets the value of Revision for the instance
func (instance *MSNdis_ObjectHeader) SetPropertyRevision(value uint8) (err error) {
	return instance.SetProperty("Revision", (value))
}

// GetRevision gets the value of Revision for the instance
func (instance *MSNdis_ObjectHeader) GetPropertyRevision() (value uint8, err error) {
	retValue, err := instance.GetProperty("Revision")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSize sets the value of Size for the instance
func (instance *MSNdis_ObjectHeader) SetPropertySize(value uint16) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSNdis_ObjectHeader) GetPropertySize() (value uint16, err error) {
	retValue, err := instance.GetProperty("Size")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetType sets the value of Type for the instance
func (instance *MSNdis_ObjectHeader) SetPropertyType(value uint8) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSNdis_ObjectHeader) GetPropertyType() (value uint8, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
