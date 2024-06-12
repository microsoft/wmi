// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapter_PacketDirectCapabilities struct
type MSFT_NetAdapter_PacketDirectCapabilities struct {
	*cim.WmiInstance

	//
	DrainNotificationSupported bool

	//
	MaximumModerationInterval uint32

	//
	MaximumNumberOfRxQueues uint32

	//
	MaximumNumberOfRxQueuesForDefaultVPort uint32

	//
	MaximumNumberOfRxQueuesPerNonDefaultVPort uint32

	//
	MaximumNumberOfTxQueues uint32

	//
	MaximumNumberOfTxQueuesForDefaultVPort uint32

	//
	MaximumNumberOfTxQueuesPerNonDefaultVPort uint32

	//
	MaximumRxPartialBufferCount uint32

	//
	MaximumRxQueueSize uint32

	//
	MaximumTxPartialBufferCount uint32

	//
	MaximumTxQueueSize uint32

	//
	MinimumModerationInterval uint32

	//
	ModerationByCountSupported bool

	//
	ModerationByIntervalSupported bool

	//
	ModerationIntervalGranularity uint32
}

func NewMSFT_NetAdapter_PacketDirectCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapter_PacketDirectCapabilities, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_PacketDirectCapabilities{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapter_PacketDirectCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapter_PacketDirectCapabilities, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_PacketDirectCapabilities{
		WmiInstance: tmp,
	}
	return
}

// SetDrainNotificationSupported sets the value of DrainNotificationSupported for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyDrainNotificationSupported(value bool) (err error) {
	return instance.SetProperty("DrainNotificationSupported", (value))
}

// GetDrainNotificationSupported gets the value of DrainNotificationSupported for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyDrainNotificationSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("DrainNotificationSupported")
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

// SetMaximumModerationInterval sets the value of MaximumModerationInterval for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyMaximumModerationInterval(value uint32) (err error) {
	return instance.SetProperty("MaximumModerationInterval", (value))
}

// GetMaximumModerationInterval gets the value of MaximumModerationInterval for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyMaximumModerationInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumModerationInterval")
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

// SetMaximumNumberOfRxQueues sets the value of MaximumNumberOfRxQueues for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyMaximumNumberOfRxQueues(value uint32) (err error) {
	return instance.SetProperty("MaximumNumberOfRxQueues", (value))
}

// GetMaximumNumberOfRxQueues gets the value of MaximumNumberOfRxQueues for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyMaximumNumberOfRxQueues() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumNumberOfRxQueues")
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

// SetMaximumNumberOfRxQueuesForDefaultVPort sets the value of MaximumNumberOfRxQueuesForDefaultVPort for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyMaximumNumberOfRxQueuesForDefaultVPort(value uint32) (err error) {
	return instance.SetProperty("MaximumNumberOfRxQueuesForDefaultVPort", (value))
}

// GetMaximumNumberOfRxQueuesForDefaultVPort gets the value of MaximumNumberOfRxQueuesForDefaultVPort for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyMaximumNumberOfRxQueuesForDefaultVPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumNumberOfRxQueuesForDefaultVPort")
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

// SetMaximumNumberOfRxQueuesPerNonDefaultVPort sets the value of MaximumNumberOfRxQueuesPerNonDefaultVPort for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyMaximumNumberOfRxQueuesPerNonDefaultVPort(value uint32) (err error) {
	return instance.SetProperty("MaximumNumberOfRxQueuesPerNonDefaultVPort", (value))
}

// GetMaximumNumberOfRxQueuesPerNonDefaultVPort gets the value of MaximumNumberOfRxQueuesPerNonDefaultVPort for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyMaximumNumberOfRxQueuesPerNonDefaultVPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumNumberOfRxQueuesPerNonDefaultVPort")
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

// SetMaximumNumberOfTxQueues sets the value of MaximumNumberOfTxQueues for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyMaximumNumberOfTxQueues(value uint32) (err error) {
	return instance.SetProperty("MaximumNumberOfTxQueues", (value))
}

// GetMaximumNumberOfTxQueues gets the value of MaximumNumberOfTxQueues for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyMaximumNumberOfTxQueues() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumNumberOfTxQueues")
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

// SetMaximumNumberOfTxQueuesForDefaultVPort sets the value of MaximumNumberOfTxQueuesForDefaultVPort for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyMaximumNumberOfTxQueuesForDefaultVPort(value uint32) (err error) {
	return instance.SetProperty("MaximumNumberOfTxQueuesForDefaultVPort", (value))
}

// GetMaximumNumberOfTxQueuesForDefaultVPort gets the value of MaximumNumberOfTxQueuesForDefaultVPort for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyMaximumNumberOfTxQueuesForDefaultVPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumNumberOfTxQueuesForDefaultVPort")
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

// SetMaximumNumberOfTxQueuesPerNonDefaultVPort sets the value of MaximumNumberOfTxQueuesPerNonDefaultVPort for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyMaximumNumberOfTxQueuesPerNonDefaultVPort(value uint32) (err error) {
	return instance.SetProperty("MaximumNumberOfTxQueuesPerNonDefaultVPort", (value))
}

// GetMaximumNumberOfTxQueuesPerNonDefaultVPort gets the value of MaximumNumberOfTxQueuesPerNonDefaultVPort for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyMaximumNumberOfTxQueuesPerNonDefaultVPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumNumberOfTxQueuesPerNonDefaultVPort")
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

// SetMaximumRxPartialBufferCount sets the value of MaximumRxPartialBufferCount for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyMaximumRxPartialBufferCount(value uint32) (err error) {
	return instance.SetProperty("MaximumRxPartialBufferCount", (value))
}

// GetMaximumRxPartialBufferCount gets the value of MaximumRxPartialBufferCount for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyMaximumRxPartialBufferCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumRxPartialBufferCount")
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

// SetMaximumRxQueueSize sets the value of MaximumRxQueueSize for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyMaximumRxQueueSize(value uint32) (err error) {
	return instance.SetProperty("MaximumRxQueueSize", (value))
}

// GetMaximumRxQueueSize gets the value of MaximumRxQueueSize for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyMaximumRxQueueSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumRxQueueSize")
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

// SetMaximumTxPartialBufferCount sets the value of MaximumTxPartialBufferCount for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyMaximumTxPartialBufferCount(value uint32) (err error) {
	return instance.SetProperty("MaximumTxPartialBufferCount", (value))
}

// GetMaximumTxPartialBufferCount gets the value of MaximumTxPartialBufferCount for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyMaximumTxPartialBufferCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumTxPartialBufferCount")
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

// SetMaximumTxQueueSize sets the value of MaximumTxQueueSize for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyMaximumTxQueueSize(value uint32) (err error) {
	return instance.SetProperty("MaximumTxQueueSize", (value))
}

// GetMaximumTxQueueSize gets the value of MaximumTxQueueSize for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyMaximumTxQueueSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumTxQueueSize")
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

// SetMinimumModerationInterval sets the value of MinimumModerationInterval for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyMinimumModerationInterval(value uint32) (err error) {
	return instance.SetProperty("MinimumModerationInterval", (value))
}

// GetMinimumModerationInterval gets the value of MinimumModerationInterval for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyMinimumModerationInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinimumModerationInterval")
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

// SetModerationByCountSupported sets the value of ModerationByCountSupported for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyModerationByCountSupported(value bool) (err error) {
	return instance.SetProperty("ModerationByCountSupported", (value))
}

// GetModerationByCountSupported gets the value of ModerationByCountSupported for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyModerationByCountSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("ModerationByCountSupported")
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

// SetModerationByIntervalSupported sets the value of ModerationByIntervalSupported for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyModerationByIntervalSupported(value bool) (err error) {
	return instance.SetProperty("ModerationByIntervalSupported", (value))
}

// GetModerationByIntervalSupported gets the value of ModerationByIntervalSupported for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyModerationByIntervalSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("ModerationByIntervalSupported")
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

// SetModerationIntervalGranularity sets the value of ModerationIntervalGranularity for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) SetPropertyModerationIntervalGranularity(value uint32) (err error) {
	return instance.SetProperty("ModerationIntervalGranularity", (value))
}

// GetModerationIntervalGranularity gets the value of ModerationIntervalGranularity for the instance
func (instance *MSFT_NetAdapter_PacketDirectCapabilities) GetPropertyModerationIntervalGranularity() (value uint32, err error) {
	retValue, err := instance.GetProperty("ModerationIntervalGranularity")
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
