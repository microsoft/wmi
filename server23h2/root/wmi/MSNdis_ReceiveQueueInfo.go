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

// MSNdis_ReceiveQueueInfo struct
type MSNdis_ReceiveQueueInfo struct {
	*MSNdis

	//
	Flags uint32

	//
	Header MSNdis_ObjectHeader

	//
	LookaheadSize uint32

	//
	MSIXTableEntry uint32

	//
	NumSuggestedReceiveBuffers uint32

	//
	ProcessorAffinity MSNdis_GroupAffinity

	//
	QueueGroupId uint32

	//
	QueueId uint32

	//
	QueueName MSNdis_CountedString

	//
	QueueState uint32

	//
	QueueType uint32

	//
	VmName MSNdis_CountedString
}

func NewMSNdis_ReceiveQueueInfoEx1(instance *cim.WmiInstance) (newInstance *MSNdis_ReceiveQueueInfo, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveQueueInfo{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_ReceiveQueueInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_ReceiveQueueInfo, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveQueueInfo{
		MSNdis: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSNdis_ReceiveQueueInfo) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_ReceiveQueueInfo) GetPropertyFlags() (value uint32, err error) {
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

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_ReceiveQueueInfo) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_ReceiveQueueInfo) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}

// SetLookaheadSize sets the value of LookaheadSize for the instance
func (instance *MSNdis_ReceiveQueueInfo) SetPropertyLookaheadSize(value uint32) (err error) {
	return instance.SetProperty("LookaheadSize", (value))
}

// GetLookaheadSize gets the value of LookaheadSize for the instance
func (instance *MSNdis_ReceiveQueueInfo) GetPropertyLookaheadSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("LookaheadSize")
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

// SetMSIXTableEntry sets the value of MSIXTableEntry for the instance
func (instance *MSNdis_ReceiveQueueInfo) SetPropertyMSIXTableEntry(value uint32) (err error) {
	return instance.SetProperty("MSIXTableEntry", (value))
}

// GetMSIXTableEntry gets the value of MSIXTableEntry for the instance
func (instance *MSNdis_ReceiveQueueInfo) GetPropertyMSIXTableEntry() (value uint32, err error) {
	retValue, err := instance.GetProperty("MSIXTableEntry")
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

// SetNumSuggestedReceiveBuffers sets the value of NumSuggestedReceiveBuffers for the instance
func (instance *MSNdis_ReceiveQueueInfo) SetPropertyNumSuggestedReceiveBuffers(value uint32) (err error) {
	return instance.SetProperty("NumSuggestedReceiveBuffers", (value))
}

// GetNumSuggestedReceiveBuffers gets the value of NumSuggestedReceiveBuffers for the instance
func (instance *MSNdis_ReceiveQueueInfo) GetPropertyNumSuggestedReceiveBuffers() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumSuggestedReceiveBuffers")
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

// SetProcessorAffinity sets the value of ProcessorAffinity for the instance
func (instance *MSNdis_ReceiveQueueInfo) SetPropertyProcessorAffinity(value MSNdis_GroupAffinity) (err error) {
	return instance.SetProperty("ProcessorAffinity", (value))
}

// GetProcessorAffinity gets the value of ProcessorAffinity for the instance
func (instance *MSNdis_ReceiveQueueInfo) GetPropertyProcessorAffinity() (value MSNdis_GroupAffinity, err error) {
	retValue, err := instance.GetProperty("ProcessorAffinity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_GroupAffinity)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_GroupAffinity is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_GroupAffinity(valuetmp)

	return
}

// SetQueueGroupId sets the value of QueueGroupId for the instance
func (instance *MSNdis_ReceiveQueueInfo) SetPropertyQueueGroupId(value uint32) (err error) {
	return instance.SetProperty("QueueGroupId", (value))
}

// GetQueueGroupId gets the value of QueueGroupId for the instance
func (instance *MSNdis_ReceiveQueueInfo) GetPropertyQueueGroupId() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueueGroupId")
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

// SetQueueId sets the value of QueueId for the instance
func (instance *MSNdis_ReceiveQueueInfo) SetPropertyQueueId(value uint32) (err error) {
	return instance.SetProperty("QueueId", (value))
}

// GetQueueId gets the value of QueueId for the instance
func (instance *MSNdis_ReceiveQueueInfo) GetPropertyQueueId() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueueId")
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

// SetQueueName sets the value of QueueName for the instance
func (instance *MSNdis_ReceiveQueueInfo) SetPropertyQueueName(value MSNdis_CountedString) (err error) {
	return instance.SetProperty("QueueName", (value))
}

// GetQueueName gets the value of QueueName for the instance
func (instance *MSNdis_ReceiveQueueInfo) GetPropertyQueueName() (value MSNdis_CountedString, err error) {
	retValue, err := instance.GetProperty("QueueName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_CountedString)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_CountedString is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_CountedString(valuetmp)

	return
}

// SetQueueState sets the value of QueueState for the instance
func (instance *MSNdis_ReceiveQueueInfo) SetPropertyQueueState(value uint32) (err error) {
	return instance.SetProperty("QueueState", (value))
}

// GetQueueState gets the value of QueueState for the instance
func (instance *MSNdis_ReceiveQueueInfo) GetPropertyQueueState() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueueState")
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

// SetQueueType sets the value of QueueType for the instance
func (instance *MSNdis_ReceiveQueueInfo) SetPropertyQueueType(value uint32) (err error) {
	return instance.SetProperty("QueueType", (value))
}

// GetQueueType gets the value of QueueType for the instance
func (instance *MSNdis_ReceiveQueueInfo) GetPropertyQueueType() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueueType")
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

// SetVmName sets the value of VmName for the instance
func (instance *MSNdis_ReceiveQueueInfo) SetPropertyVmName(value MSNdis_CountedString) (err error) {
	return instance.SetProperty("VmName", (value))
}

// GetVmName gets the value of VmName for the instance
func (instance *MSNdis_ReceiveQueueInfo) GetPropertyVmName() (value MSNdis_CountedString, err error) {
	retValue, err := instance.GetProperty("VmName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_CountedString)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_CountedString is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_CountedString(valuetmp)

	return
}
