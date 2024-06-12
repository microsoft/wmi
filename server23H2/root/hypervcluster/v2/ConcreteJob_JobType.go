// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source ConcreteJob_JobType
//////////////////////////////////////////////
package v2

// ConcreteJob_JobType
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
	ConcreteJob_JobType_Pause_Virtual_Machine ConcreteJob_JobType = 26
	// Resume_Virtual_Machine enum
	ConcreteJob_JobType_Resume_Virtual_Machine ConcreteJob_JobType = 27
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
