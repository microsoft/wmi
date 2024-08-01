// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ReplicationSettings struct
type MSFT_ReplicationSettings struct {
	*cim.WmiInstance

	// A set of volumes where the replication journal for the ReplicationGroup is hosted.
	LogDevices []MSFT_Volume

	// Size of replication journal in units of bytes. Size must be in multiples of gigabytes.
	LogSizeInBytes uint64

	// Minimum number of synchronous replication partnerships that are in synchronous replication state for I/O to continue on source Replication Group.
	ReplicationQuorum uint16

	// Mode describes whether the target elements will be updated synchronously or asynchronously. If NULL, implementation decides the mode.
	SyncMode ReplicationSettings_SyncMode

	// TODO
	TargetElementSupplier uint16

	// TODO
	ThinProvisioningPolicy uint16
}

func NewMSFT_ReplicationSettingsEx1(instance *cim.WmiInstance) (newInstance *MSFT_ReplicationSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ReplicationSettings{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ReplicationSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ReplicationSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ReplicationSettings{
		WmiInstance: tmp,
	}
	return
}

// SetLogDevices sets the value of LogDevices for the instance
func (instance *MSFT_ReplicationSettings) SetPropertyLogDevices(value []MSFT_Volume) (err error) {
	return instance.SetProperty("LogDevices", (value))
}

// GetLogDevices gets the value of LogDevices for the instance
func (instance *MSFT_ReplicationSettings) GetPropertyLogDevices() (value []MSFT_Volume, err error) {
	retValue, err := instance.GetProperty("LogDevices")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_Volume)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_Volume is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_Volume(valuetmp))
	}

	return
}

// SetLogSizeInBytes sets the value of LogSizeInBytes for the instance
func (instance *MSFT_ReplicationSettings) SetPropertyLogSizeInBytes(value uint64) (err error) {
	return instance.SetProperty("LogSizeInBytes", (value))
}

// GetLogSizeInBytes gets the value of LogSizeInBytes for the instance
func (instance *MSFT_ReplicationSettings) GetPropertyLogSizeInBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogSizeInBytes")
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

// SetReplicationQuorum sets the value of ReplicationQuorum for the instance
func (instance *MSFT_ReplicationSettings) SetPropertyReplicationQuorum(value uint16) (err error) {
	return instance.SetProperty("ReplicationQuorum", (value))
}

// GetReplicationQuorum gets the value of ReplicationQuorum for the instance
func (instance *MSFT_ReplicationSettings) GetPropertyReplicationQuorum() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationQuorum")
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

// SetSyncMode sets the value of SyncMode for the instance
func (instance *MSFT_ReplicationSettings) SetPropertySyncMode(value ReplicationSettings_SyncMode) (err error) {
	return instance.SetProperty("SyncMode", (value))
}

// GetSyncMode gets the value of SyncMode for the instance
func (instance *MSFT_ReplicationSettings) GetPropertySyncMode() (value ReplicationSettings_SyncMode, err error) {
	retValue, err := instance.GetProperty("SyncMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ReplicationSettings_SyncMode(valuetmp)

	return
}

// SetTargetElementSupplier sets the value of TargetElementSupplier for the instance
func (instance *MSFT_ReplicationSettings) SetPropertyTargetElementSupplier(value uint16) (err error) {
	return instance.SetProperty("TargetElementSupplier", (value))
}

// GetTargetElementSupplier gets the value of TargetElementSupplier for the instance
func (instance *MSFT_ReplicationSettings) GetPropertyTargetElementSupplier() (value uint16, err error) {
	retValue, err := instance.GetProperty("TargetElementSupplier")
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

// SetThinProvisioningPolicy sets the value of ThinProvisioningPolicy for the instance
func (instance *MSFT_ReplicationSettings) SetPropertyThinProvisioningPolicy(value uint16) (err error) {
	return instance.SetProperty("ThinProvisioningPolicy", (value))
}

// GetThinProvisioningPolicy gets the value of ThinProvisioningPolicy for the instance
func (instance *MSFT_ReplicationSettings) GetPropertyThinProvisioningPolicy() (value uint16, err error) {
	retValue, err := instance.GetProperty("ThinProvisioningPolicy")
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
