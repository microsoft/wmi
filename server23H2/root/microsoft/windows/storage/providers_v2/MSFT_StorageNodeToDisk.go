// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
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

// MSFT_StorageNodeToDisk struct
type MSFT_StorageNodeToDisk struct {
	*cim.WmiInstance

	//
	Disk MSFT_Disk

	// The operating system's number for the disk. Disk 0 is typically the boot device. Disk numbers may not necessarily remain the same across reboots.
	DiskNumber uint32

	// The health status of the Volume.
	///0 - 'Healthy': The disk is functioning normally.
	///1 - 'Warning': The disk is still functioning, but has detected errors or issues that require administrator intervention.
	///2 - 'Unhealthy': The volume is not functioning, due to errors or failures. The volume needs immediate attention from an administrator.
	HealthStatus StorageNodeToDisk_HealthStatus

	//
	IsOffline bool

	//
	IsReadOnly bool

	// If IsOffline is TRUE, this property informs the user of the specific reason for the disk being offline.
	///1 - 'Policy': The user requested the disk to be offline.
	///2 - 'Redundant Path': The disk is used for multi-path I/O.
	///3 - 'Snapshot': The disk is a snapshot disk.
	///4 - 'Collision': There was a signature or identifier collision with another disk.
	///5 - 'Resource Exhaustion': There were insufficient resources to bring the disk online.
	///6 - 'Critical Write Failures': There were critical write failures on the disk.
	///7 - 'Data Integrity Scan Required': A data integrity scan is required.
	OfflineReason StorageNodeToDisk_OfflineReason

	// An array of values that denote the current operational status of the volume.
	///0 - 'Unknown': The operational status is unknown.
	///1 - 'Other': A vendor-specific OperationalStatus has been specified by setting the OtherOperationalStatusDescription property.
	///2 - 'OK': The disk is responding to commands and is in a normal operating state.
	///3 - 'Degraded': The disk is responding to commands, but is not running in an optimal operating state.
	///4 - 'Stressed': The disk is functioning, but needs attention. For example, the disk might be overloaded or overheated.
	///5 - 'Predictive Failure': The disk is functioning, but a failure is likely to occur in the near future.
	///6 - 'Error': An error has occurred.
	///7 - 'Non-Recoverable Error': A non-recoverable error has occurred.
	///8 - 'Starting': The disk is in the process of starting.
	///9 - 'Stopping': The disk is in the process of stopping.
	///10 - 'Stopped': The disk was stopped or shut down in a clean and orderly fashion.
	///11 - 'In Service': The disk is being configured, maintained, cleaned, or otherwise administered.
	///12 - 'No Contact': The storage provider has knowledge of the disk, but has never been able to establish communication with it.
	///13 - 'Lost Communication': The storage provider has knowledge of the disk and has contacted it successfully in the past, but the disk is currently unreachable.
	///14 - 'Aborted': Similar to Stopped, except that the disk stopped abruptly and may require configuration or maintenance.
	///15 - 'Dormant': The disk is reachable, but it is inactive.
	///16 - 'Supporting Entity in Error': This status value does not necessarily indicate trouble with the disk, but it does indicate that another device or connection that the disk depends on may need attention.
	///17 - 'Completed': The disk has completed an operation. This status value should be combined with OK, Error, or Degraded, depending on the outcome of the operation.
	///0xD010 - 'Online': In Windows-based storage subsystems, this indicates that the object is online.
	///0xD011 - 'Not Ready': In Windows-based storage subsystems, this indicates that the object is not ready.
	///0xD012 - 'No Media': In Windows-based storage subsystems, this indicates that the object has no media present.
	///0xD013 - 'Offline': In Windows-based storage subsystems, this indicates that the object is offline.
	///0xD014 - 'Failed': In Windows-based storage subsystems, this indicates that the object is in a failed state.
	OperationalStatus []StorageNodeToDisk_OperationalStatus

	//
	StorageNode MSFT_StorageNode
}

func NewMSFT_StorageNodeToDiskEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageNodeToDisk, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageNodeToDisk{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_StorageNodeToDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageNodeToDisk, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageNodeToDisk{
		WmiInstance: tmp,
	}
	return
}

// SetDisk sets the value of Disk for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyDisk(value MSFT_Disk) (err error) {
	return instance.SetProperty("Disk", (value))
}

// GetDisk gets the value of Disk for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyDisk() (value MSFT_Disk, err error) {
	retValue, err := instance.GetProperty("Disk")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_Disk)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_Disk is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_Disk(valuetmp)

	return
}

// SetDiskNumber sets the value of DiskNumber for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyDiskNumber(value uint32) (err error) {
	return instance.SetProperty("DiskNumber", (value))
}

// GetDiskNumber gets the value of DiskNumber for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyDiskNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DiskNumber")
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

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyHealthStatus(value StorageNodeToDisk_HealthStatus) (err error) {
	return instance.SetProperty("HealthStatus", (value))
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyHealthStatus() (value StorageNodeToDisk_HealthStatus, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
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

	value = StorageNodeToDisk_HealthStatus(valuetmp)

	return
}

// SetIsOffline sets the value of IsOffline for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyIsOffline(value bool) (err error) {
	return instance.SetProperty("IsOffline", (value))
}

// GetIsOffline gets the value of IsOffline for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyIsOffline() (value bool, err error) {
	retValue, err := instance.GetProperty("IsOffline")
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

// SetIsReadOnly sets the value of IsReadOnly for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyIsReadOnly(value bool) (err error) {
	return instance.SetProperty("IsReadOnly", (value))
}

// GetIsReadOnly gets the value of IsReadOnly for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyIsReadOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("IsReadOnly")
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

// SetOfflineReason sets the value of OfflineReason for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyOfflineReason(value StorageNodeToDisk_OfflineReason) (err error) {
	return instance.SetProperty("OfflineReason", (value))
}

// GetOfflineReason gets the value of OfflineReason for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyOfflineReason() (value StorageNodeToDisk_OfflineReason, err error) {
	retValue, err := instance.GetProperty("OfflineReason")
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

	value = StorageNodeToDisk_OfflineReason(valuetmp)

	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyOperationalStatus(value []StorageNodeToDisk_OperationalStatus) (err error) {
	return instance.SetProperty("OperationalStatus", (value))
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyOperationalStatus() (value []StorageNodeToDisk_OperationalStatus, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, StorageNodeToDisk_OperationalStatus(valuetmp))
	}

	return
}

// SetStorageNode sets the value of StorageNode for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyStorageNode(value MSFT_StorageNode) (err error) {
	return instance.SetProperty("StorageNode", (value))
}

// GetStorageNode gets the value of StorageNode for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyStorageNode() (value MSFT_StorageNode, err error) {
	retValue, err := instance.GetProperty("StorageNode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_StorageNode)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_StorageNode is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_StorageNode(valuetmp)

	return
}
