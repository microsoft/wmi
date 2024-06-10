// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source MemoryHierarchyError_Type
//
// ////////////////////////////////////////////
package wmi

// MemoryHierarchyError_Type
type MemoryHierarchyError_Type int

const (
	// MCA_WARNING_CACHE enum
	MemoryHierarchyError_Type_MCA_WARNING_CACHE MemoryHierarchyError_Type = 1
	// MCA_ERROR_CACHE enum
	MemoryHierarchyError_Type_MCA_ERROR_CACHE MemoryHierarchyError_Type = 2
	// MCA_WARNING_TLB enum
	MemoryHierarchyError_Type_MCA_WARNING_TLB MemoryHierarchyError_Type = 3
	// MCA_ERROR_TLB enum
	MemoryHierarchyError_Type_MCA_ERROR_TLB MemoryHierarchyError_Type = 4
	// MCA_WARNING_CPU_BUS enum
	MemoryHierarchyError_Type_MCA_WARNING_CPU_BUS MemoryHierarchyError_Type = 5
	// MCA_ERROR_CPU_BUS enum
	MemoryHierarchyError_Type_MCA_ERROR_CPU_BUS MemoryHierarchyError_Type = 6
	// MCA_WARNING_REGISTER_FILE enum
	MemoryHierarchyError_Type_MCA_WARNING_REGISTER_FILE MemoryHierarchyError_Type = 7
	// MCA_ERROR_REGISTER_FILE enum
	MemoryHierarchyError_Type_MCA_ERROR_REGISTER_FILE MemoryHierarchyError_Type = 8
	// MCA_WARNING_MAS enum
	MemoryHierarchyError_Type_MCA_WARNING_MAS MemoryHierarchyError_Type = 9
	// MCA_ERROR_MAS enum
	MemoryHierarchyError_Type_MCA_ERROR_MAS MemoryHierarchyError_Type = 10
	// MCA_WARNING_MEM_UNKNOWN enum
	MemoryHierarchyError_Type_MCA_WARNING_MEM_UNKNOWN MemoryHierarchyError_Type = 11
	// MCA_ERROR_MEM_UNKNOWN enum
	MemoryHierarchyError_Type_MCA_ERROR_MEM_UNKNOWN MemoryHierarchyError_Type = 12
	// MCA_WARNING_MEM_1_2 enum
	MemoryHierarchyError_Type_MCA_WARNING_MEM_1_2 MemoryHierarchyError_Type = 13
	// MCA_ERROR_MEM_1_2 enum
	MemoryHierarchyError_Type_MCA_ERROR_MEM_1_2 MemoryHierarchyError_Type = 14
	// MCA_WARNING_MEM_1_2_5 enum
	MemoryHierarchyError_Type_MCA_WARNING_MEM_1_2_5 MemoryHierarchyError_Type = 15
	// MCA_ERROR_MEM_1_2_5 enum
	MemoryHierarchyError_Type_MCA_ERROR_MEM_1_2_5 MemoryHierarchyError_Type = 16
	// MCA_WARNING_MEM_1_2_5_4 enum
	MemoryHierarchyError_Type_MCA_WARNING_MEM_1_2_5_4 MemoryHierarchyError_Type = 17
	// MCA_ERROR_MEM_1_2_5_4 enum
	MemoryHierarchyError_Type_MCA_ERROR_MEM_1_2_5_4 MemoryHierarchyError_Type = 18
	// MCA_WARNING_SYSTEM_EVENT enum
	MemoryHierarchyError_Type_MCA_WARNING_SYSTEM_EVENT MemoryHierarchyError_Type = 19
	// MCA_ERROR_SYSTEM_EVENT enum
	MemoryHierarchyError_Type_MCA_ERROR_SYSTEM_EVENT MemoryHierarchyError_Type = 20
	// MCA_WARNING_PCI_BUS_PARITY enum
	MemoryHierarchyError_Type_MCA_WARNING_PCI_BUS_PARITY MemoryHierarchyError_Type = 21
	// MCA_ERROR_PCI_BUS_PARITY enum
	MemoryHierarchyError_Type_MCA_ERROR_PCI_BUS_PARITY MemoryHierarchyError_Type = 22
	// MCA_WARNING_PCI_BUS_PARITY_NO_INFO enum
	MemoryHierarchyError_Type_MCA_WARNING_PCI_BUS_PARITY_NO_INFO MemoryHierarchyError_Type = 23
	// MCA_ERROR_PCI_BUS_PARITY_NO_INFO enum
	MemoryHierarchyError_Type_MCA_ERROR_PCI_BUS_PARITY_NO_INFO MemoryHierarchyError_Type = 24
	// MCA_WARNING_PCI_BUS_SERR enum
	MemoryHierarchyError_Type_MCA_WARNING_PCI_BUS_SERR MemoryHierarchyError_Type = 25
	// MCA_ERROR_PCI_BUS_SERR enum
	MemoryHierarchyError_Type_MCA_ERROR_PCI_BUS_SERR MemoryHierarchyError_Type = 26
	// MCA_WARNING_PCI_BUS_SERR_NO_INFO enum
	MemoryHierarchyError_Type_MCA_WARNING_PCI_BUS_SERR_NO_INFO MemoryHierarchyError_Type = 27
	// MCA_ERROR_PCI_BUS_SERR_NO_INFO enum
	MemoryHierarchyError_Type_MCA_ERROR_PCI_BUS_SERR_NO_INFO MemoryHierarchyError_Type = 28
	// MCA_WARNING_PCI_BUS_MASTER_ABORT enum
	MemoryHierarchyError_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT MemoryHierarchyError_Type = 29
	// MCA_ERROR_PCI_BUS_MASTER_ABORT enum
	MemoryHierarchyError_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT MemoryHierarchyError_Type = 30
	// MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO enum
	MemoryHierarchyError_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO MemoryHierarchyError_Type = 31
	// MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO enum
	MemoryHierarchyError_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO MemoryHierarchyError_Type = 32
	// MCA_WARNING_PCI_BUS_TIMEOUT enum
	MemoryHierarchyError_Type_MCA_WARNING_PCI_BUS_TIMEOUT MemoryHierarchyError_Type = 33
	// MCA_ERROR_PCI_BUS_TIMEOUT enum
	MemoryHierarchyError_Type_MCA_ERROR_PCI_BUS_TIMEOUT MemoryHierarchyError_Type = 34
	// MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO enum
	MemoryHierarchyError_Type_MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO MemoryHierarchyError_Type = 35
	// MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO enum
	MemoryHierarchyError_Type_MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO MemoryHierarchyError_Type = 36
	// MCA_WARNING_PCI_BUS_UNKNOWN enum
	MemoryHierarchyError_Type_MCA_WARNING_PCI_BUS_UNKNOWN MemoryHierarchyError_Type = 37
	// MCA_ERROR_PCI_BUS_UNKNOWN enum
	MemoryHierarchyError_Type_MCA_ERROR_PCI_BUS_UNKNOWN MemoryHierarchyError_Type = 38
	// MCA_WARNING_PCI_DEVICE enum
	MemoryHierarchyError_Type_MCA_WARNING_PCI_DEVICE MemoryHierarchyError_Type = 39
	// MCA_ERROR_PCI_DEVICE enum
	MemoryHierarchyError_Type_MCA_ERROR_PCI_DEVICE MemoryHierarchyError_Type = 40
	// MCA_WARNING_SMBIOS enum
	MemoryHierarchyError_Type_MCA_WARNING_SMBIOS MemoryHierarchyError_Type = 41
	// MCA_ERROR_SMBIOS enum
	MemoryHierarchyError_Type_MCA_ERROR_SMBIOS MemoryHierarchyError_Type = 42
	// MCA_WARNING_PLATFORM_SPECIFIC enum
	MemoryHierarchyError_Type_MCA_WARNING_PLATFORM_SPECIFIC MemoryHierarchyError_Type = 43
	// MCA_ERROR_PLATFORM_SPECIFIC enum
	MemoryHierarchyError_Type_MCA_ERROR_PLATFORM_SPECIFIC MemoryHierarchyError_Type = 44
	// MCA_WARNING_UNKNOWN enum
	MemoryHierarchyError_Type_MCA_WARNING_UNKNOWN MemoryHierarchyError_Type = 45
	// MCA_ERROR_UNKNOWN enum
	MemoryHierarchyError_Type_MCA_ERROR_UNKNOWN MemoryHierarchyError_Type = 46
	// MCA_WARNING_UNKNOWN_NO_CPU enum
	MemoryHierarchyError_Type_MCA_WARNING_UNKNOWN_NO_CPU MemoryHierarchyError_Type = 47
	// MCA_ERROR_UNKNOWN_NO_CPU enum
	MemoryHierarchyError_Type_MCA_ERROR_UNKNOWN_NO_CPU MemoryHierarchyError_Type = 48
	// MCA_WARNING_CMC_THRESHOLD_EXCEEDED enum
	MemoryHierarchyError_Type_MCA_WARNING_CMC_THRESHOLD_EXCEEDED MemoryHierarchyError_Type = 49
	// MCA_WARNING_CPE_THRESHOLD_EXCEEDED enum
	MemoryHierarchyError_Type_MCA_WARNING_CPE_THRESHOLD_EXCEEDED MemoryHierarchyError_Type = 50
	// MCA_WARNING_CPU_THERMAL_THROTTLED enum
	MemoryHierarchyError_Type_MCA_WARNING_CPU_THERMAL_THROTTLED MemoryHierarchyError_Type = 51
	// MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED enum
	MemoryHierarchyError_Type_MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED MemoryHierarchyError_Type = 52
	// MCA_WARNING_CPU enum
	MemoryHierarchyError_Type_MCA_WARNING_CPU MemoryHierarchyError_Type = 53
	// MCA_ERROR_CPU enum
	MemoryHierarchyError_Type_MCA_ERROR_CPU MemoryHierarchyError_Type = 54
	// MCA_MEMORYHIERARCHY_ERROR enum
	MemoryHierarchyError_Type_MCA_MEMORYHIERARCHY_ERROR MemoryHierarchyError_Type = 55
	// MCA_TLB_ERROR enum
	MemoryHierarchyError_Type_MCA_TLB_ERROR MemoryHierarchyError_Type = 56
	// MCA_BUS_ERROR enum
	MemoryHierarchyError_Type_MCA_BUS_ERROR MemoryHierarchyError_Type = 57
	// MCA_BUS_TIMEOUT_ERROR enum
	MemoryHierarchyError_Type_MCA_BUS_TIMEOUT_ERROR MemoryHierarchyError_Type = 58
	// MCA_INTERNALTIMER_ERROR enum
	MemoryHierarchyError_Type_MCA_INTERNALTIMER_ERROR MemoryHierarchyError_Type = 59
	// MCA_MICROCODE_ROM_PARITY_ERROR enum
	MemoryHierarchyError_Type_MCA_MICROCODE_ROM_PARITY_ERROR MemoryHierarchyError_Type = 60
	// MCA_EXTERNAL_ERROR enum
	MemoryHierarchyError_Type_MCA_EXTERNAL_ERROR MemoryHierarchyError_Type = 61
	// MCA_FRC_ERROR enum
	MemoryHierarchyError_Type_MCA_FRC_ERROR MemoryHierarchyError_Type = 62
)
