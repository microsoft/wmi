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

// Header_PartitionInfoExtension_TypeGroup struct
type Header_PartitionInfoExtension_TypeGroup struct {
	*EventTraceEvent

	//
	EventVersion uint16

	//
	ParentId interface{}

	//
	PartitionId interface{}

	//
	PartitionType uint32

	//
	QpcOffsetFromRoot int64

	//
	Reserved uint16
}

func NewHeader_PartitionInfoExtension_TypeGroupEx1(instance *cim.WmiInstance) (newInstance *Header_PartitionInfoExtension_TypeGroup, err error) {
	tmp, err := NewEventTraceEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Header_PartitionInfoExtension_TypeGroup{
		EventTraceEvent: tmp,
	}
	return
}

func NewHeader_PartitionInfoExtension_TypeGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Header_PartitionInfoExtension_TypeGroup, err error) {
	tmp, err := NewEventTraceEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Header_PartitionInfoExtension_TypeGroup{
		EventTraceEvent: tmp,
	}
	return
}

// SetEventVersion sets the value of EventVersion for the instance
func (instance *Header_PartitionInfoExtension_TypeGroup) SetPropertyEventVersion(value uint16) (err error) {
	return instance.SetProperty("EventVersion", (value))
}

// GetEventVersion gets the value of EventVersion for the instance
func (instance *Header_PartitionInfoExtension_TypeGroup) GetPropertyEventVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("EventVersion")
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

// SetParentId sets the value of ParentId for the instance
func (instance *Header_PartitionInfoExtension_TypeGroup) SetPropertyParentId(value interface{}) (err error) {
	return instance.SetProperty("ParentId", (value))
}

// GetParentId gets the value of ParentId for the instance
func (instance *Header_PartitionInfoExtension_TypeGroup) GetPropertyParentId() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ParentId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetPartitionId sets the value of PartitionId for the instance
func (instance *Header_PartitionInfoExtension_TypeGroup) SetPropertyPartitionId(value interface{}) (err error) {
	return instance.SetProperty("PartitionId", (value))
}

// GetPartitionId gets the value of PartitionId for the instance
func (instance *Header_PartitionInfoExtension_TypeGroup) GetPropertyPartitionId() (value interface{}, err error) {
	retValue, err := instance.GetProperty("PartitionId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetPartitionType sets the value of PartitionType for the instance
func (instance *Header_PartitionInfoExtension_TypeGroup) SetPropertyPartitionType(value uint32) (err error) {
	return instance.SetProperty("PartitionType", (value))
}

// GetPartitionType gets the value of PartitionType for the instance
func (instance *Header_PartitionInfoExtension_TypeGroup) GetPropertyPartitionType() (value uint32, err error) {
	retValue, err := instance.GetProperty("PartitionType")
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

// SetQpcOffsetFromRoot sets the value of QpcOffsetFromRoot for the instance
func (instance *Header_PartitionInfoExtension_TypeGroup) SetPropertyQpcOffsetFromRoot(value int64) (err error) {
	return instance.SetProperty("QpcOffsetFromRoot", (value))
}

// GetQpcOffsetFromRoot gets the value of QpcOffsetFromRoot for the instance
func (instance *Header_PartitionInfoExtension_TypeGroup) GetPropertyQpcOffsetFromRoot() (value int64, err error) {
	retValue, err := instance.GetProperty("QpcOffsetFromRoot")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetReserved sets the value of Reserved for the instance
func (instance *Header_PartitionInfoExtension_TypeGroup) SetPropertyReserved(value uint16) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *Header_PartitionInfoExtension_TypeGroup) GetPropertyReserved() (value uint16, err error) {
	retValue, err := instance.GetProperty("Reserved")
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
