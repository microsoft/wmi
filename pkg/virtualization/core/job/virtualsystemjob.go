// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package job

import (
	"fmt"
	"strings"
	"time"

	moccodes "github.com/microsoft/moc/pkg/errors/codes"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

// ConcreteJob_JobType copied from /server2019/root/virtualization/v2/ConcreteJob_JobType.go
// Remade here with corrections to enum values for Pause and Resume, as the autogenerated file was out of date.
type ConcreteJob_JobType int

const (
	// Unknown enum
	ConcreteJob_JobType_Unknown ConcreteJob_JobType = 0
	// Define_Virtual_Machine enum
	ConcreteJob_JobType_Define_Virtual_Machine ConcreteJob_JobType = 1
	// Modify_Virtual_Machine enum
	ConcreteJob_JobType_Modify_Virtual_Machine ConcreteJob_JobType = 2
	// Destroy_Virtual_Machine enum
	ConcreteJob_JobType_Destroy_Virtual_Machine ConcreteJob_JobType = 3
	// Modify_Management_Service_Settings enum
	ConcreteJob_JobType_Modify_Management_Service_Settings ConcreteJob_JobType = 4
	// Initialize_Virtual_Machine enum
	ConcreteJob_JobType_Initialize_Virtual_Machine ConcreteJob_JobType = 10
	// Waiting_to_Start_Virtual_Machine enum
	ConcreteJob_JobType_Waiting_to_Start_Virtual_Machine ConcreteJob_JobType = 11
	// Start_Virtual_Machine enum
	ConcreteJob_JobType_Start_Virtual_Machine ConcreteJob_JobType = 12
	// Power_Off_Virtual_Machine enum
	ConcreteJob_JobType_Power_Off_Virtual_Machine ConcreteJob_JobType = 13
	// Save_Virtual_Machine enum
	ConcreteJob_JobType_Save_Virtual_Machine ConcreteJob_JobType = 14
	// Restore_Virtual_Machine enum
	ConcreteJob_JobType_Restore_Virtual_Machine ConcreteJob_JobType = 15
	// Shut_Down_Virtual_Machine enum
	ConcreteJob_JobType_Shut_Down_Virtual_Machine ConcreteJob_JobType = 16
	// Pause_Virtual_Machine enum
	ConcreteJob_JobType_Pause_Virtual_Machine ConcreteJob_JobType = 20
	// Resume_Virtual_Machine enum
	ConcreteJob_JobType_Resume_Virtual_Machine ConcreteJob_JobType = 21
	// Reset_Virtual_Machine enum
	ConcreteJob_JobType_Reset_Virtual_Machine ConcreteJob_JobType = 28
	// Reboot_Virtual_Machine enum
	ConcreteJob_JobType_Reboot_Virtual_Machine ConcreteJob_JobType = 29
	// Add_Virtual_Machine_Resources enum
	ConcreteJob_JobType_Add_Virtual_Machine_Resources ConcreteJob_JobType = 30
	// Modify_Virtual_Machine_Resources enum
	ConcreteJob_JobType_Modify_Virtual_Machine_Resources ConcreteJob_JobType = 31
	// Remove_Virtual_Machine_Resources enum
	ConcreteJob_JobType_Remove_Virtual_Machine_Resources ConcreteJob_JobType = 32
	// Request_Initial_Virtual_Machine_Memory enum
	ConcreteJob_JobType_Request_Initial_Virtual_Machine_Memory ConcreteJob_JobType = 40
	// Add_Memory_to_Virtual_Machine enum
	ConcreteJob_JobType_Add_Memory_to_Virtual_Machine ConcreteJob_JobType = 41
	// Remove_Memory_from_Virtual_Machine enum
	ConcreteJob_JobType_Remove_Memory_from_Virtual_Machine ConcreteJob_JobType = 42
	// Merging_VHD_Disks enum
	ConcreteJob_JobType_Merging_VHD_Disks ConcreteJob_JobType = 50
	// Create_VSS_Snapshot_inside_Virtual_Machine enum
	ConcreteJob_JobType_Create_VSS_Snapshot_inside_Virtual_Machine ConcreteJob_JobType = 51
	// Get_Import_Setting_Data enum
	ConcreteJob_JobType_Get_Import_Setting_Data ConcreteJob_JobType = 60
	// Import_Virtual_Machine enum
	ConcreteJob_JobType_Import_Virtual_Machine ConcreteJob_JobType = 61
	// Export_Virtual_Machine enum
	ConcreteJob_JobType_Export_Virtual_Machine ConcreteJob_JobType = 62
	// Register_Configuration enum
	ConcreteJob_JobType_Register_Configuration ConcreteJob_JobType = 63
	// Unregister_Configuration enum
	ConcreteJob_JobType_Unregister_Configuration ConcreteJob_JobType = 64
	// Snapshot_Virtual_Machine enum
	ConcreteJob_JobType_Snapshot_Virtual_Machine ConcreteJob_JobType = 70
	// Apply_Virtual_Machine_Snapshot enum
	ConcreteJob_JobType_Apply_Virtual_Machine_Snapshot ConcreteJob_JobType = 71
	// Delete_Virtual_Machine_Snapshot enum
	ConcreteJob_JobType_Delete_Virtual_Machine_Snapshot ConcreteJob_JobType = 72
	// Clear_Virtual_Machine_Snapshot_State enum
	ConcreteJob_JobType_Clear_Virtual_Machine_Snapshot_State ConcreteJob_JobType = 73
	// Add_Resources_to_Resource_Pool enum
	ConcreteJob_JobType_Add_Resources_to_Resource_Pool ConcreteJob_JobType = 80
	// Remove_Resources_from_Resource_Pool enum
	ConcreteJob_JobType_Remove_Resources_from_Resource_Pool ConcreteJob_JobType = 81
	// Modify_Replication_Server_Settings enum
	ConcreteJob_JobType_Modify_Replication_Server_Settings ConcreteJob_JobType = 90
	// Create_Replication_Relationship enum
	ConcreteJob_JobType_Create_Replication_Relationship ConcreteJob_JobType = 91
	// Modify_Replication_Relationship_Settings enum
	ConcreteJob_JobType_Modify_Replication_Relationship_Settings ConcreteJob_JobType = 92
	// Remove_Replication_Relationship enum
	ConcreteJob_JobType_Remove_Replication_Relationship ConcreteJob_JobType = 93
	// Start_Inband_Initial_Replication enum
	ConcreteJob_JobType_Start_Inband_Initial_Replication ConcreteJob_JobType = 94
	// Import_Replication enum
	ConcreteJob_JobType_Import_Replication ConcreteJob_JobType = 95
	// Replicate_State_Change enum
	ConcreteJob_JobType_Replicate_State_Change ConcreteJob_JobType = 96
	// Initiate_Failover enum
	ConcreteJob_JobType_Initiate_Failover ConcreteJob_JobType = 97
	// Revert_Failover enum
	ConcreteJob_JobType_Revert_Failover ConcreteJob_JobType = 98
	// Commit_Failover enum
	ConcreteJob_JobType_Commit_Failover ConcreteJob_JobType = 99
	// Inititate_Synced_Replication enum
	ConcreteJob_JobType_Inititate_Synced_Replication ConcreteJob_JobType = 100
	// Cancel_Synced_Replication enum
	ConcreteJob_JobType_Cancel_Synced_Replication ConcreteJob_JobType = 101
	// Initiate_Test_Replica enum
	ConcreteJob_JobType_Initiate_Test_Replica ConcreteJob_JobType = 102
	// Remove_Test_Replica enum
	ConcreteJob_JobType_Remove_Test_Replica ConcreteJob_JobType = 103
	// Reverse_Replication enum
	ConcreteJob_JobType_Reverse_Replication ConcreteJob_JobType = 104
	// Replication_Sending_Delta enum
	ConcreteJob_JobType_Replication_Sending_Delta ConcreteJob_JobType = 105
	// Replication_Receiving_Delta enum
	ConcreteJob_JobType_Replication_Receiving_Delta ConcreteJob_JobType = 106
	// Resynchronizing enum
	ConcreteJob_JobType_Resynchronizing ConcreteJob_JobType = 107
	// Apply_change_log enum
	ConcreteJob_JobType_Apply_change_log ConcreteJob_JobType = 108
	// Stop_Initial_Replication enum
	ConcreteJob_JobType_Stop_Initial_Replication ConcreteJob_JobType = 109
	// Stop_Resynchronizing enum
	ConcreteJob_JobType_Stop_Resynchronizing ConcreteJob_JobType = 110
	// Get_Replica_statistics enum
	ConcreteJob_JobType_Get_Replica_statistics ConcreteJob_JobType = 111
	// Prepare_for_Consistency_Checker enum
	ConcreteJob_JobType_Prepare_for_Consistency_Checker ConcreteJob_JobType = 112
	// Consistency_Checker enum
	ConcreteJob_JobType_Consistency_Checker ConcreteJob_JobType = 113
	// Stop_Consistency_Checker enum
	ConcreteJob_JobType_Stop_Consistency_Checker ConcreteJob_JobType = 114
	// Test_Replication_Connection enum
	ConcreteJob_JobType_Test_Replication_Connection ConcreteJob_JobType = 115
	// Sending_Initial_Replica enum
	ConcreteJob_JobType_Sending_Initial_Replica ConcreteJob_JobType = 116
	// Start_Resync_Initial_Replication enum
	ConcreteJob_JobType_Start_Resync_Initial_Replication ConcreteJob_JobType = 117
	// Start_Export_Initial_Replication enum
	ConcreteJob_JobType_Start_Export_Initial_Replication ConcreteJob_JobType = 118
	// Reset_Replica_Statistics enum
	ConcreteJob_JobType_Reset_Replica_Statistics ConcreteJob_JobType = 119
	// Apply_Registered_Deltas enum
	ConcreteJob_JobType_Apply_Registered_Deltas ConcreteJob_JobType = 120
	// Resynchronizing_Extended_Replication enum
	ConcreteJob_JobType_Resynchronizing_Extended_Replication ConcreteJob_JobType = 121
	// Reading_Test_Replica_Configuration enum
	ConcreteJob_JobType_Reading_Test_Replica_Configuration ConcreteJob_JobType = 122
	// Change_the_replication_mode_to_primary enum
	ConcreteJob_JobType_Change_the_replication_mode_to_primary ConcreteJob_JobType = 123
	// Initiate_Failback enum
	ConcreteJob_JobType_Initiate_Failback ConcreteJob_JobType = 124
	// Update_Disk_Set enum
	ConcreteJob_JobType_Update_Disk_Set ConcreteJob_JobType = 125
	// Define_Ethernet_Switch enum
	ConcreteJob_JobType_Define_Ethernet_Switch ConcreteJob_JobType = 130
	// Modify_Ethernet_Switch_Settings enum
	ConcreteJob_JobType_Modify_Ethernet_Switch_Settings ConcreteJob_JobType = 131
	// Destroy_Ethernet_Switch enum
	ConcreteJob_JobType_Destroy_Ethernet_Switch ConcreteJob_JobType = 132
	// Add_Ethernet_Switch_Resources enum
	ConcreteJob_JobType_Add_Ethernet_Switch_Resources ConcreteJob_JobType = 133
	// Modify_Ethernet_Switch_Resources enum
	ConcreteJob_JobType_Modify_Ethernet_Switch_Resources ConcreteJob_JobType = 134
	// Remove_Ethernet_Switch_Resources enum
	ConcreteJob_JobType_Remove_Ethernet_Switch_Resources ConcreteJob_JobType = 135
	// Validate_Planned_Virtual_Machine enum
	ConcreteJob_JobType_Validate_Planned_Virtual_Machine ConcreteJob_JobType = 140
	// Realizing_Virtual_Machine enum
	ConcreteJob_JobType_Realizing_Virtual_Machine ConcreteJob_JobType = 141
	// Creating_a_Resource_Pool enum
	ConcreteJob_JobType_Creating_a_Resource_Pool ConcreteJob_JobType = 150
	// Changing_the_Parent_Resources_of_a_Resource_Pool enum
	ConcreteJob_JobType_Changing_the_Parent_Resources_of_a_Resource_Pool ConcreteJob_JobType = 151
	// Changing_the_Non_alloction_Settings_of_a_Resource_Pool enum
	ConcreteJob_JobType_Changing_the_Non_alloction_Settings_of_a_Resource_Pool ConcreteJob_JobType = 152
	// Deleting_a_Resource_Pool enum
	ConcreteJob_JobType_Deleting_a_Resource_Pool ConcreteJob_JobType = 153
	// Enable_RemoteFx_GPU enum
	ConcreteJob_JobType_Enable_RemoteFx_GPU ConcreteJob_JobType = 160
	// Disable_RemoteFx_GPU enum
	ConcreteJob_JobType_Disable_RemoteFx_GPU ConcreteJob_JobType = 161
	// Modify_3D_Service_Settings enum
	ConcreteJob_JobType_Modify_3D_Service_Settings ConcreteJob_JobType = 162
	// Backup_Virtual_Machine enum
	ConcreteJob_JobType_Backup_Virtual_Machine ConcreteJob_JobType = 170
	// Guest_Service_Interface enum
	ConcreteJob_JobType_Guest_Service_Interface ConcreteJob_JobType = 180
	// Query_Guest_Cluster_Information enum
	ConcreteJob_JobType_Query_Guest_Cluster_Information ConcreteJob_JobType = 181
	// Define_Collection enum
	ConcreteJob_JobType_Define_Collection ConcreteJob_JobType = 190
	// Destroy_Collection enum
	ConcreteJob_JobType_Destroy_Collection ConcreteJob_JobType = 191
	// Rename_Collection enum
	ConcreteJob_JobType_Rename_Collection ConcreteJob_JobType = 192
	// Add_Member_to_Collection enum
	ConcreteJob_JobType_Add_Member_to_Collection ConcreteJob_JobType = 193
	// Remove_Member_from_Collection enum
	ConcreteJob_JobType_Remove_Member_from_Collection ConcreteJob_JobType = 194
	// Add_Setting_to_Collection enum
	ConcreteJob_JobType_Add_Setting_to_Collection ConcreteJob_JobType = 195
	// Remove_Setting_from_Collection enum
	ConcreteJob_JobType_Remove_Setting_from_Collection ConcreteJob_JobType = 196
	// Modify_Setting_on_Collection enum
	ConcreteJob_JobType_Modify_Setting_on_Collection ConcreteJob_JobType = 197
	// Snapshot_Collection enum
	ConcreteJob_JobType_Snapshot_Collection ConcreteJob_JobType = 198
	// Convert_Snapshot_to_Reference_Point enum
	ConcreteJob_JobType_Convert_Snapshot_to_Reference_Point ConcreteJob_JobType = 200
	// Create_Reference_Point enum
	ConcreteJob_JobType_Create_Reference_Point ConcreteJob_JobType = 201
	// Delete_Reference_Point enum
	ConcreteJob_JobType_Delete_Reference_Point ConcreteJob_JobType = 202
	// Export_Reference_Point enum
	ConcreteJob_JobType_Export_Reference_Point ConcreteJob_JobType = 203
	// Remove_Associated_Data_from_Reference_Point enum
	ConcreteJob_JobType_Remove_Associated_Data_from_Reference_Point ConcreteJob_JobType = 204
	// Create_Reference_Point_on_Collection enum
	ConcreteJob_JobType_Create_Reference_Point_on_Collection ConcreteJob_JobType = 205
	// Export_Reference_Point_on_Collection enum
	ConcreteJob_JobType_Export_Reference_Point_on_Collection ConcreteJob_JobType = 206
	// Remove_Associated_Data_from_Reference_Point_on_Collection enum
	ConcreteJob_JobType_Remove_Associated_Data_from_Reference_Point_on_Collection ConcreteJob_JobType = 207
	// Delete_Reference_Point_on_Collection enum
	ConcreteJob_JobType_Delete_Reference_Point_on_Collection ConcreteJob_JobType = 208
	// Import_Reference_Point_metadata enum
	ConcreteJob_JobType_Import_Reference_Point_metadata ConcreteJob_JobType = 209
	// Mount_or_Dismount_Assignable_Device enum
	ConcreteJob_JobType_Mount_or_Dismount_Assignable_Device ConcreteJob_JobType = 260
)

type VirtualSystemJob struct {
	*v2.Msvm_ConcreteJob
}

// NewVirtualSystemJob
func NewVirtualSystemJob(instance *wmi.WmiInstance) (*VirtualSystemJob, error) {
	j, err := v2.NewMsvm_ConcreteJobEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualSystemJob{j}, nil
}

func (vmjob *VirtualSystemJob) String() string {
	jtype, err := vmjob.JobType()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("Type[%s]", jtype)
}

// GetJobType gets the value of JobType for the instance
func (vmjob *VirtualSystemJob) JobType() (value ConcreteJob_JobType, err error) {
	retValue, err := vmjob.GetProperty("JobType")
	if err != nil {
		return
	}
	valueint, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	value = ConcreteJob_JobType(valueint)
	return
}

// WaitForPercentComplete waits for the percentComplete or timeout
func (vmjob *VirtualSystemJob) WaitForPercentComplete(percentComplete uint16, timeoutSeconds int16) error {
	start := time.Now()
	// Run the loop, only if the job is actually running
	for !vmjob.IsComplete() {
		pComplete, err := vmjob.PercentComplete()
		if err != nil {
			return err
		}
		//fmt.Printf("WaitForPercentComplete [%d]\n ", pComplete)
		// Break if have achieved the target
		if pComplete >= percentComplete {
			break
		}

		time.Sleep(100 * time.Millisecond)
		if timeoutSeconds < 0 {
			continue
		}
		// If we have waited enough time, break
		if time.Since(start) > (time.Duration(timeoutSeconds) * time.Second) {
			state, err2 := vmjob.GetPropertyJobState()
			if err2 != nil {
				state = 0
			}
			exception := vmjob.GetException()
			return errors.Wrapf(errors.Timedout, "WaitForPercentComplete timeout. Current state: [%v], Exception: [%v]", state, exception)
		}
	}

	return vmjob.GetException()
}

// WaitForAction waits for the task based on the action type, percent complete and timeoutSeconds
func (vmjob *VirtualSystemJob) WaitForAction(action wmi.UserAction, percentComplete uint16, timeoutSeconds int16) error {
	switch action {
	case wmi.Wait:
		return vmjob.WaitForPercentComplete(percentComplete, timeoutSeconds)
	case wmi.Cancel:
		return vmjob.WaitForPercentComplete(percentComplete, timeoutSeconds)
		// Fixme
		// vm.Cancel()
	case wmi.None:
		fallthrough
	case wmi.Default:
		fallthrough
	case wmi.Async:
		break
	}
	return nil
}

// PercentComplete
func (vmjob *VirtualSystemJob) PercentComplete() (uint16, error) {
	err := vmjob.Refresh()
	if err != nil {
		return 0, err
	}
	retValue, err := vmjob.GetProperty("PercentComplete")
	if err != nil {
		return 0, err
	}
	return uint16(retValue.(int32)), nil
}

func (vmjob *VirtualSystemJob) IsComplete() bool {
	err := vmjob.Refresh()
	if err != nil {

	}
	state, err := vmjob.GetPropertyJobState()
	if err != nil {
		return false
	}
	switch state {
	case v2.ConcreteJob_JobState_New:
		fallthrough
	case v2.ConcreteJob_JobState_Starting:
		fallthrough
	case v2.ConcreteJob_JobState_Running:
		fallthrough
	case v2.ConcreteJob_JobState_Suspended:
		fallthrough
	case v2.ConcreteJob_JobState_Shutting_Down:
		return false
	case v2.ConcreteJob_JobState_Completed:
		fallthrough
	case v2.ConcreteJob_JobState_Terminated:
		fallthrough
	case v2.ConcreteJob_JobState_Killed:
		fallthrough
	case v2.ConcreteJob_JobState_Exception:
		return true
	}
	return false
}

func (vmjob *VirtualSystemJob) GetException() error {
	vmjob.Refresh()
	state, err := vmjob.GetPropertyJobState()
	if err != nil {
		return err
	}
	switch state {
	case v2.ConcreteJob_JobState_Terminated:
		fallthrough
	case v2.ConcreteJob_JobState_Killed:
		fallthrough
	case v2.ConcreteJob_JobState_Exception:
		errorCode, _ := vmjob.GetPropertyErrorCode()
		errorDescription, _ := vmjob.GetPropertyErrorDescription()
		errorSummaryDescription, _ := vmjob.GetPropertyErrorSummaryDescription()
		if errorCode == 0 {
			if strings.Contains(errorSummaryDescription, errors.OutOfMemoryErrorString) {
				return errors.Wrapf(errors.NewWMIError(errorCode),
					"ErrorCode[%d] ErrorDescription[%s] ErrorSummaryDescription [%s]",
					moccodes.OutOfCapacity, errorDescription, errorSummaryDescription)
			}
		}

		return errors.Wrapf(errors.NewWMIError(errorCode),
			"ErrorCode[%d] ErrorDescription[%s] ErrorSummaryDescription [%s]",
			errorCode, errorDescription, errorSummaryDescription)
	}
	return nil
}

func WaitForJobCompletionEx(result int32, currentJob *VirtualSystemJob, timeoutSeconds int16) error {
	if result == 0 {
		return nil
	} else if result == 4096 {
		return currentJob.WaitForAction(wmi.Wait, 100, timeoutSeconds)
	} else {
		return errors.Wrapf(errors.Failed, "Unable to Wait for Job on Result[%d] ", result)
	}

}

func WaitForJobCompletionEx2(instance *wmi.WmiInstance, result int32, jobName string) error {
	return nil
}

func WaitForJobCompletion(instance *wmi.WmiInstance, result int32, jobType ConcreteJob_JobType, timeoutSeconds int16) error {
	if result == 0 {
		return nil
	} else if result == 4096 {
		vmjob, err := getJob(instance, jobType)
		if err != nil {
			// Job is scheduled, but we were not able to find the job
			return err
		}
		defer vmjob.Close()
		return vmjob.WaitForAction(wmi.Wait, 100, timeoutSeconds)
	} else {
		return errors.Wrapf(errors.Failed, "Unable to Wait for Job on Resource Pool Result[%d] JobType[%d]", result, jobType)
	}
}

func getJob(instance *wmi.WmiInstance, jobType ConcreteJob_JobType) (*VirtualSystemJob, error) {
	jobString := fmt.Sprintf("%d", jobType)
	query := query.NewWmiQuery("Msvm_ConcreteJob", "JobType", jobString)
	jobs, err := instance.GetAllRelatedWithQuery(query)
	if err != nil {
		return nil, err
	}
	if len(jobs) == 0 {
		return nil, errors.Wrapf(errors.NotFound, "Unable to find related Job with type [%d]", jobType)
	}
	defer jobs.Close()
	// FIXME: Find the correct Job, when multiple jobs are returned
	jobInstance, err := jobs[0].Clone()
	if err != nil {
		return nil, err
	}

	return NewVirtualSystemJob(jobInstance)
}

func getCIMJob(instance *wmi.WmiInstance, jobName string) (*VirtualSystemJob, error) {
	query := query.NewWmiQuery("CIM_ConcreteJob", "Name", jobName)
	jobs, err := instance.GetAllRelatedWithQuery(query)
	if err != nil {
		return nil, err
	}
	if len(jobs) == 0 {
		return nil, errors.Wrapf(errors.NotFound, "Unable to find related Job with name [%s]", jobName)
	}
	defer jobs.Close()
	// FIXME: Find the correct Job, when multiple jobs are returned
	jobInstance, err := jobs[0].Clone()
	if err != nil {
		return nil, err
	}

	return NewVirtualSystemJob(jobInstance)
}
