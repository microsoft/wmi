// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// MSFT_VirtualDiskToVirtualDisk struct
type MSFT_VirtualDiskToVirtualDisk struct {
	*cim.WmiInstance

	// CopyMethodology specifies what copy methodology the copy engine uses to create and/or maintain the target element. Values are:
	///0 - 'Not Specified': The method of maintaining the copy is not specified.
	///3 - 'Full Copy': This indicates that a full copy of the source object is (or will be) generated .
	///4 - 'Incremental-Copy': Only changed data from source element is copied to target element.
	///5 - 'Differential-Copy': Only the new writes to the source element are copied to the target element.
	///6 - 'Copy-On-Write': Affected data is copied on the first write to the source or to the target elements.
	///7 - 'Copy-On-Access': Affected data is copied on the first access to the source element.
	///8 - 'Delta-Update': Difference based replication where after the initial copy, only updates to source are copied to target.
	///9 - 'Snap-And-Clone': The service creates a snapshot of the source element first, then uses the snapshot as the source of the copy operation to the target element.
	CopyMethodology VirtualDiskToVirtualDisk_CopyMethodology

	// CopyPriority allows the priority of background copy engine I/O to be managed relative to host I/O operations during a sequential background copy operation. Values are:
	///1 - 'Low': copy engine I/O lower priority than host I/O.
	///2 - 'Same': copy engine I/O has the same priority as host I/O.
	///3 - 'High': copy engine I/O has higher priority than host I/O.
	CopyPriority VirtualDiskToVirtualDisk_CopyPriority

	// CopyState describes the state of the association with respect to replication activity. Values are:
	///2 - 'Initialized': The link to enable replication is established and source/replica elements are associated, but the copy operation has not started.
	///3 - 'Unsynchronized': Not all the source element data has been copied to the target element.
	///4 - 'Synchronized': For the Mirror, Snapshot, or Clone replication, the target represents a copy of the source.
	///5 - 'Broken': The relationship is non-functional due to errors in the source, the target, the path between the two or space constraints.
	///6 - 'Fractured': Target is split from the source.
	///7 - 'Split': The target element was gracefully (or systematically) split from its source element -- consistency is guaranteed.
	///8 - 'Inactive': Copy operation has stopped, writes to source element will not be sent to target element.
	///9 - 'Suspended': Data flow between the source and target elements has stopped. Writes to source element are held until the association is resumed.
	///10 - 'Failedover': Reads and writes to/from the target element. Source element is not reachable.
	///11 - 'Prepared': Initialization completed and the copy operation started; however, the data flow has not started.
	///12 - 'Aborted': The copy operation is aborted with the Abort operation. Use the Resync Replica operation to restart the copy operation.
	///13 - 'Skewed': The target has been modified and is no longer synchronized with the source element or the point-in-time view.
	///14 - 'Mixed': Applies to the CopyState of GroupSynchronized. It indicates the StorageSynchronized associations of the elements in the groups have different CopyState values.
	CopyState VirtualDiskToVirtualDisk_CopyState

	// CopyType describes the Replication Policy. Values are:
	///2 - 'Async': create and maintain an asynchronous copy of the source.
	///3 - 'Sync': create and maintain a synchronized copy of the source.
	///4 - 'UnSyncAssoc': create an unsynchronized copy and maintain an association to the source.
	///5 - 'UnSyncUnAssoc': create an unsynchronized copy with a temporary association that is deleted upon completion of the copy operation.
	CopyType VirtualDiskToVirtualDisk_CopyType

	// Specifies the percent of the work completed to reach synchronization. Must be set to NULL if implementation is not capable of providing this information.
	PercentSynced uint16

	// ProgressStatus describes the status of the association with respect to Replication activity. Values are:
	///2 - 'Completed': The request is completed. Copy operation is idle.
	///3 - 'Dormant': Indicates that the copy operation is inactive suspended or quiesced.
	///4 - 'Initializing': In the process of establishing source/replica association and the copy operation has not started.
	///5 - 'Preparing': preparation-in-progress.
	///6 - 'Synchronizing': sync-in-progress.
	///7 - 'Resyncing': resync-in-progress.
	///8 - 'Restoring': restore-in-progress.
	///9 - 'Fracturing': fracture-in-progress.
	///10 - 'Splitting': split-in-progress.
	///11 - 'Failing over': in the process of switching source and target.
	///12 - 'Failing back': Undoing the result of failover.
	///13 - 'Detaching': detach-in-progress.
	///14 - 'Aborting': abort-in-progress.
	///15 - 'Mixed': Applies to groups with element pairs with different statuses. Generally, the individual statuses need to be examined.
	///16 - 'Suspending': The copy operation is in the process of being suspended.
	///17 - 'Requires fracture': The requested operation has completed, however, the synchronization relationship needs to be fractured before further copy operations can be issued.
	///18 - 'Requires resync': The requested operation has completed, however, the synchronization relationship needs to be resynced before further copy operations can be issued.
	///19 - 'Requires activate': The requested operation has completed, however, the synchronization relationship needs to be activated before further copy operations can be issued.
	///20 - 'Pending': The flow of data has stopped momentarily due to limited bandwidth or busy system.
	ProgressStatus VirtualDiskToVirtualDisk_ProgressStatus

	// ReplicaType provides information on how the Replica is being maintained. Values are:
	///2 - 'Full Copy': This indicates that a full copy of the source object is (or will be) generated .
	///3 - 'Before Delta': This indicates that the source object will be maintained as a delta data from the replica.
	///4 - 'After Delta': This indicates that the replica will be maintained as delta data from the source object.
	///5 - 'Log': This indicates that the replica object is being maintained as a log of changes to the source.
	///0 - 'Not Specified': The method of maintaining the copy is not specified.
	ReplicaType VirtualDiskToVirtualDisk_ReplicaType

	// RequestedCopyState is an integer enumeration that indicates the last requested or desired state for the association. The actual state of the association is represented by CopyState. Note that when CopyState reaches the requested state, this property will be set to 'Not Applicable.
	RequestedCopyState uint16

	//
	SourceVirtualDisk MSFT_VirtualDisk

	// Boolean indicating whether synchronization is maintained.
	SyncMaintained bool

	// Mode describes whether the target elements will be updated synchronously or asynchronously. If NULL, implementation decides the mode.
	SyncMode VirtualDiskToVirtualDisk_SyncMode

	// SyncState describes the state of the association with respect to Replication activity. Values are:
	///2 - 'Initialized': The link to enable replication is established and source/replica elements are associated, but the Copy engine has not started.
	///3 - 'PrepareInProgress': Preparation for Replication is in progress and the Copy engine has started.
	///4 - 'Prepared': All necessary preparation has completed.
	///5 - 'ResyncInProgress': Synchronization or Resynchronization is in progress. This may be the initial 'copy' or subsequent changes being copied.
	///6 - 'Synchronized': An Async or Sync replication is currently synchronized. When this value is set, SyncMaintained will be true.
	///7 - 'FractureInProgress': An operation to fracture an Async or Sync replication is in progress.
	///8 - 'QuiesceInProgress': A quiesce operation is in progress.
	///9 - 'Quiesced': The replication has been quiesced and is ready for a change.
	///10 - 'RestoreInProgress': An operation is in progress to copy the Synced object to the System object.
	///11 - 'Idle': The 'normal' state for an UnSyncAssoc replica.
	///12 - 'Broken': The relationship is non-functional due to errors in the source, the target, the path between the two or space constraints.
	///13 - 'Fractured': An Async or Sync replication is fractured.
	///14 - 'Frozen': All blocks copied from source to an UnSyncAssoc replica and the copy engine is stopped.
	///15 - 'CopyInProgress': A deferred background copy operation is in progress to copy the source to the replica target for an UnSyncAssoc association.
	///
	SyncState VirtualDiskToVirtualDisk_SyncState

	// The point in time that the virtual disks were synchronized.
	SyncTime string

	// SyncType describes the intended outcome of the replication. Values are:
	///6 - 'Mirror': create and maintain a copy of the source.
	///7 - 'Snapshot': create a point-in-time, virtual copy of the source.
	///8 - 'Clone': create a point-in-time, full copy the source.
	SyncType VirtualDiskToVirtualDisk_SyncType

	//
	TargetVirtualDisk MSFT_VirtualDisk
}

func NewMSFT_VirtualDiskToVirtualDiskEx1(instance *cim.WmiInstance) (newInstance *MSFT_VirtualDiskToVirtualDisk, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_VirtualDiskToVirtualDisk{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_VirtualDiskToVirtualDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_VirtualDiskToVirtualDisk, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_VirtualDiskToVirtualDisk{
		WmiInstance: tmp,
	}
	return
}

// SetCopyMethodology sets the value of CopyMethodology for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertyCopyMethodology(value VirtualDiskToVirtualDisk_CopyMethodology) (err error) {
	return instance.SetProperty("CopyMethodology", (value))
}

// GetCopyMethodology gets the value of CopyMethodology for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertyCopyMethodology() (value VirtualDiskToVirtualDisk_CopyMethodology, err error) {
	retValue, err := instance.GetProperty("CopyMethodology")
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

	value = VirtualDiskToVirtualDisk_CopyMethodology(valuetmp)

	return
}

// SetCopyPriority sets the value of CopyPriority for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertyCopyPriority(value VirtualDiskToVirtualDisk_CopyPriority) (err error) {
	return instance.SetProperty("CopyPriority", (value))
}

// GetCopyPriority gets the value of CopyPriority for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertyCopyPriority() (value VirtualDiskToVirtualDisk_CopyPriority, err error) {
	retValue, err := instance.GetProperty("CopyPriority")
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

	value = VirtualDiskToVirtualDisk_CopyPriority(valuetmp)

	return
}

// SetCopyState sets the value of CopyState for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertyCopyState(value VirtualDiskToVirtualDisk_CopyState) (err error) {
	return instance.SetProperty("CopyState", (value))
}

// GetCopyState gets the value of CopyState for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertyCopyState() (value VirtualDiskToVirtualDisk_CopyState, err error) {
	retValue, err := instance.GetProperty("CopyState")
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

	value = VirtualDiskToVirtualDisk_CopyState(valuetmp)

	return
}

// SetCopyType sets the value of CopyType for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertyCopyType(value VirtualDiskToVirtualDisk_CopyType) (err error) {
	return instance.SetProperty("CopyType", (value))
}

// GetCopyType gets the value of CopyType for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertyCopyType() (value VirtualDiskToVirtualDisk_CopyType, err error) {
	retValue, err := instance.GetProperty("CopyType")
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

	value = VirtualDiskToVirtualDisk_CopyType(valuetmp)

	return
}

// SetPercentSynced sets the value of PercentSynced for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertyPercentSynced(value uint16) (err error) {
	return instance.SetProperty("PercentSynced", (value))
}

// GetPercentSynced gets the value of PercentSynced for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertyPercentSynced() (value uint16, err error) {
	retValue, err := instance.GetProperty("PercentSynced")
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

// SetProgressStatus sets the value of ProgressStatus for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertyProgressStatus(value VirtualDiskToVirtualDisk_ProgressStatus) (err error) {
	return instance.SetProperty("ProgressStatus", (value))
}

// GetProgressStatus gets the value of ProgressStatus for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertyProgressStatus() (value VirtualDiskToVirtualDisk_ProgressStatus, err error) {
	retValue, err := instance.GetProperty("ProgressStatus")
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

	value = VirtualDiskToVirtualDisk_ProgressStatus(valuetmp)

	return
}

// SetReplicaType sets the value of ReplicaType for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertyReplicaType(value VirtualDiskToVirtualDisk_ReplicaType) (err error) {
	return instance.SetProperty("ReplicaType", (value))
}

// GetReplicaType gets the value of ReplicaType for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertyReplicaType() (value VirtualDiskToVirtualDisk_ReplicaType, err error) {
	retValue, err := instance.GetProperty("ReplicaType")
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

	value = VirtualDiskToVirtualDisk_ReplicaType(valuetmp)

	return
}

// SetRequestedCopyState sets the value of RequestedCopyState for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertyRequestedCopyState(value uint16) (err error) {
	return instance.SetProperty("RequestedCopyState", (value))
}

// GetRequestedCopyState gets the value of RequestedCopyState for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertyRequestedCopyState() (value uint16, err error) {
	retValue, err := instance.GetProperty("RequestedCopyState")
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

// SetSourceVirtualDisk sets the value of SourceVirtualDisk for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertySourceVirtualDisk(value MSFT_VirtualDisk) (err error) {
	return instance.SetProperty("SourceVirtualDisk", (value))
}

// GetSourceVirtualDisk gets the value of SourceVirtualDisk for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertySourceVirtualDisk() (value MSFT_VirtualDisk, err error) {
	retValue, err := instance.GetProperty("SourceVirtualDisk")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_VirtualDisk)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_VirtualDisk is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_VirtualDisk(valuetmp)

	return
}

// SetSyncMaintained sets the value of SyncMaintained for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertySyncMaintained(value bool) (err error) {
	return instance.SetProperty("SyncMaintained", (value))
}

// GetSyncMaintained gets the value of SyncMaintained for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertySyncMaintained() (value bool, err error) {
	retValue, err := instance.GetProperty("SyncMaintained")
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

// SetSyncMode sets the value of SyncMode for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertySyncMode(value VirtualDiskToVirtualDisk_SyncMode) (err error) {
	return instance.SetProperty("SyncMode", (value))
}

// GetSyncMode gets the value of SyncMode for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertySyncMode() (value VirtualDiskToVirtualDisk_SyncMode, err error) {
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

	value = VirtualDiskToVirtualDisk_SyncMode(valuetmp)

	return
}

// SetSyncState sets the value of SyncState for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertySyncState(value VirtualDiskToVirtualDisk_SyncState) (err error) {
	return instance.SetProperty("SyncState", (value))
}

// GetSyncState gets the value of SyncState for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertySyncState() (value VirtualDiskToVirtualDisk_SyncState, err error) {
	retValue, err := instance.GetProperty("SyncState")
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

	value = VirtualDiskToVirtualDisk_SyncState(valuetmp)

	return
}

// SetSyncTime sets the value of SyncTime for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertySyncTime(value string) (err error) {
	return instance.SetProperty("SyncTime", (value))
}

// GetSyncTime gets the value of SyncTime for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertySyncTime() (value string, err error) {
	retValue, err := instance.GetProperty("SyncTime")
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

// SetSyncType sets the value of SyncType for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertySyncType(value VirtualDiskToVirtualDisk_SyncType) (err error) {
	return instance.SetProperty("SyncType", (value))
}

// GetSyncType gets the value of SyncType for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertySyncType() (value VirtualDiskToVirtualDisk_SyncType, err error) {
	retValue, err := instance.GetProperty("SyncType")
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

	value = VirtualDiskToVirtualDisk_SyncType(valuetmp)

	return
}

// SetTargetVirtualDisk sets the value of TargetVirtualDisk for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertyTargetVirtualDisk(value MSFT_VirtualDisk) (err error) {
	return instance.SetProperty("TargetVirtualDisk", (value))
}

// GetTargetVirtualDisk gets the value of TargetVirtualDisk for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertyTargetVirtualDisk() (value MSFT_VirtualDisk, err error) {
	retValue, err := instance.GetProperty("TargetVirtualDisk")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_VirtualDisk)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_VirtualDisk is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_VirtualDisk(valuetmp)

	return
}
