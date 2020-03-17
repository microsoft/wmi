// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source StorageNodeToPhysicalDisk_LoadBalancePolicy
//////////////////////////////////////////////
package providers_v2

// StorageNodeToPhysicalDisk_LoadBalancePolicy
type StorageNodeToPhysicalDisk_LoadBalancePolicy int

const (
	// Unknown enum
	StorageNodeToPhysicalDisk_LoadBalancePolicy_Unknown StorageNodeToPhysicalDisk_LoadBalancePolicy = 0
	// Fail_Over enum
	StorageNodeToPhysicalDisk_LoadBalancePolicy_Fail_Over StorageNodeToPhysicalDisk_LoadBalancePolicy = 1
	// Round_Robin enum
	StorageNodeToPhysicalDisk_LoadBalancePolicy_Round_Robin StorageNodeToPhysicalDisk_LoadBalancePolicy = 2
	// Round_Robin_with_Subset enum
	StorageNodeToPhysicalDisk_LoadBalancePolicy_Round_Robin_with_Subset StorageNodeToPhysicalDisk_LoadBalancePolicy = 3
	// Least_Queue_Depth enum
	StorageNodeToPhysicalDisk_LoadBalancePolicy_Least_Queue_Depth StorageNodeToPhysicalDisk_LoadBalancePolicy = 4
	// Weighted_Paths enum
	StorageNodeToPhysicalDisk_LoadBalancePolicy_Weighted_Paths StorageNodeToPhysicalDisk_LoadBalancePolicy = 5
	// Least_Blocks enum
	StorageNodeToPhysicalDisk_LoadBalancePolicy_Least_Blocks StorageNodeToPhysicalDisk_LoadBalancePolicy = 6
	// Vendor_Specific enum
	StorageNodeToPhysicalDisk_LoadBalancePolicy_Vendor_Specific StorageNodeToPhysicalDisk_LoadBalancePolicy = 7
)
