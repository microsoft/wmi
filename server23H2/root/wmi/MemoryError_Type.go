// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source MemoryError_Type
//////////////////////////////////////////////
package wmi

// MemoryError_Type
type MemoryError_Type int

const (
	// MCA_WARNING_CACHE enum
	MemoryError_Type_MCA_WARNING_CACHE MemoryError_Type = 1
	// MCA_ERROR_CACHE enum
	MemoryError_Type_MCA_ERROR_CACHE MemoryError_Type = 2
	// MCA_WARNING_TLB enum
	MemoryError_Type_MCA_WARNING_TLB MemoryError_Type = 3
	// MCA_ERROR_TLB enum
	MemoryError_Type_MCA_ERROR_TLB MemoryError_Type = 4
	// MCA_WARNING_CPU_BUS enum
	MemoryError_Type_MCA_WARNING_CPU_BUS MemoryError_Type = 5
	// MCA_ERROR_CPU_BUS enum
	MemoryError_Type_MCA_ERROR_CPU_BUS MemoryError_Type = 6
	// MCA_WARNING_REGISTER_FILE enum
	MemoryError_Type_MCA_WARNING_REGISTER_FILE MemoryError_Type = 7
	// MCA_ERROR_REGISTER_FILE enum
	MemoryError_Type_MCA_ERROR_REGISTER_FILE MemoryError_Type = 8
	// MCA_WARNING_MAS enum
	MemoryError_Type_MCA_WARNING_MAS MemoryError_Type = 9
	// MCA_ERROR_MAS enum
	MemoryError_Type_MCA_ERROR_MAS MemoryError_Type = 10
	// MCA_WARNING_MEM_UNKNOWN enum
	MemoryError_Type_MCA_WARNING_MEM_UNKNOWN MemoryError_Type = 11
	// MCA_ERROR_MEM_UNKNOWN enum
	MemoryError_Type_MCA_ERROR_MEM_UNKNOWN MemoryError_Type = 12
	// MCA_WARNING_MEM_1_2 enum
	MemoryError_Type_MCA_WARNING_MEM_1_2 MemoryError_Type = 13
	// MCA_ERROR_MEM_1_2 enum
	MemoryError_Type_MCA_ERROR_MEM_1_2 MemoryError_Type = 14
	// MCA_WARNING_MEM_1_2_5 enum
	MemoryError_Type_MCA_WARNING_MEM_1_2_5 MemoryError_Type = 15
	// MCA_ERROR_MEM_1_2_5 enum
	MemoryError_Type_MCA_ERROR_MEM_1_2_5 MemoryError_Type = 16
	// MCA_WARNING_MEM_1_2_5_4 enum
	MemoryError_Type_MCA_WARNING_MEM_1_2_5_4 MemoryError_Type = 17
	// MCA_ERROR_MEM_1_2_5_4 enum
	MemoryError_Type_MCA_ERROR_MEM_1_2_5_4 MemoryError_Type = 18
	// MCA_WARNING_SYSTEM_EVENT enum
	MemoryError_Type_MCA_WARNING_SYSTEM_EVENT MemoryError_Type = 19
	// MCA_ERROR_SYSTEM_EVENT enum
	MemoryError_Type_MCA_ERROR_SYSTEM_EVENT MemoryError_Type = 20
	// MCA_WARNING_PCI_BUS_PARITY enum
	MemoryError_Type_MCA_WARNING_PCI_BUS_PARITY MemoryError_Type = 21
	// MCA_ERROR_PCI_BUS_PARITY enum
	MemoryError_Type_MCA_ERROR_PCI_BUS_PARITY MemoryError_Type = 22
	// MCA_WARNING_PCI_BUS_PARITY_NO_INFO enum
	MemoryError_Type_MCA_WARNING_PCI_BUS_PARITY_NO_INFO MemoryError_Type = 23
	// MCA_ERROR_PCI_BUS_PARITY_NO_INFO enum
	MemoryError_Type_MCA_ERROR_PCI_BUS_PARITY_NO_INFO MemoryError_Type = 24
	// MCA_WARNING_PCI_BUS_SERR enum
	MemoryError_Type_MCA_WARNING_PCI_BUS_SERR MemoryError_Type = 25
	// MCA_ERROR_PCI_BUS_SERR enum
	MemoryError_Type_MCA_ERROR_PCI_BUS_SERR MemoryError_Type = 26
	// MCA_WARNING_PCI_BUS_SERR_NO_INFO enum
	MemoryError_Type_MCA_WARNING_PCI_BUS_SERR_NO_INFO MemoryError_Type = 27
	// MCA_ERROR_PCI_BUS_SERR_NO_INFO enum
	MemoryError_Type_MCA_ERROR_PCI_BUS_SERR_NO_INFO MemoryError_Type = 28
	// MCA_WARNING_PCI_BUS_MASTER_ABORT enum
	MemoryError_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT MemoryError_Type = 29
	// MCA_ERROR_PCI_BUS_MASTER_ABORT enum
	MemoryError_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT MemoryError_Type = 30
	// MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO enum
	MemoryError_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO MemoryError_Type = 31
	// MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO enum
	MemoryError_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO MemoryError_Type = 32
	// MCA_WARNING_PCI_BUS_TIMEOUT enum
	MemoryError_Type_MCA_WARNING_PCI_BUS_TIMEOUT MemoryError_Type = 33
	// MCA_ERROR_PCI_BUS_TIMEOUT enum
	MemoryError_Type_MCA_ERROR_PCI_BUS_TIMEOUT MemoryError_Type = 34
	// MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO enum
	MemoryError_Type_MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO MemoryError_Type = 35
	// MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO enum
	MemoryError_Type_MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO MemoryError_Type = 36
	// MCA_WARNING_PCI_BUS_UNKNOWN enum
	MemoryError_Type_MCA_WARNING_PCI_BUS_UNKNOWN MemoryError_Type = 37
	// MCA_ERROR_PCI_BUS_UNKNOWN enum
	MemoryError_Type_MCA_ERROR_PCI_BUS_UNKNOWN MemoryError_Type = 38
	// MCA_WARNING_PCI_DEVICE enum
	MemoryError_Type_MCA_WARNING_PCI_DEVICE MemoryError_Type = 39
	// MCA_ERROR_PCI_DEVICE enum
	MemoryError_Type_MCA_ERROR_PCI_DEVICE MemoryError_Type = 40
	// MCA_WARNING_SMBIOS enum
	MemoryError_Type_MCA_WARNING_SMBIOS MemoryError_Type = 41
	// MCA_ERROR_SMBIOS enum
	MemoryError_Type_MCA_ERROR_SMBIOS MemoryError_Type = 42
	// MCA_WARNING_PLATFORM_SPECIFIC enum
	MemoryError_Type_MCA_WARNING_PLATFORM_SPECIFIC MemoryError_Type = 43
	// MCA_ERROR_PLATFORM_SPECIFIC enum
	MemoryError_Type_MCA_ERROR_PLATFORM_SPECIFIC MemoryError_Type = 44
	// MCA_WARNING_UNKNOWN enum
	MemoryError_Type_MCA_WARNING_UNKNOWN MemoryError_Type = 45
	// MCA_ERROR_UNKNOWN enum
	MemoryError_Type_MCA_ERROR_UNKNOWN MemoryError_Type = 46
	// MCA_WARNING_UNKNOWN_NO_CPU enum
	MemoryError_Type_MCA_WARNING_UNKNOWN_NO_CPU MemoryError_Type = 47
	// MCA_ERROR_UNKNOWN_NO_CPU enum
	MemoryError_Type_MCA_ERROR_UNKNOWN_NO_CPU MemoryError_Type = 48
	// MCA_WARNING_CMC_THRESHOLD_EXCEEDED enum
	MemoryError_Type_MCA_WARNING_CMC_THRESHOLD_EXCEEDED MemoryError_Type = 49
	// MCA_WARNING_CPE_THRESHOLD_EXCEEDED enum
	MemoryError_Type_MCA_WARNING_CPE_THRESHOLD_EXCEEDED MemoryError_Type = 50
	// MCA_WARNING_CPU_THERMAL_THROTTLED enum
	MemoryError_Type_MCA_WARNING_CPU_THERMAL_THROTTLED MemoryError_Type = 51
	// MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED enum
	MemoryError_Type_MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED MemoryError_Type = 52
	// MCA_WARNING_CPU enum
	MemoryError_Type_MCA_WARNING_CPU MemoryError_Type = 53
	// MCA_ERROR_CPU enum
	MemoryError_Type_MCA_ERROR_CPU MemoryError_Type = 54
	// MCA_MEMORYHIERARCHY_ERROR enum
	MemoryError_Type_MCA_MEMORYHIERARCHY_ERROR MemoryError_Type = 55
	// MCA_TLB_ERROR enum
	MemoryError_Type_MCA_TLB_ERROR MemoryError_Type = 56
	// MCA_BUS_ERROR enum
	MemoryError_Type_MCA_BUS_ERROR MemoryError_Type = 57
	// MCA_BUS_TIMEOUT_ERROR enum
	MemoryError_Type_MCA_BUS_TIMEOUT_ERROR MemoryError_Type = 58
	// MCA_INTERNALTIMER_ERROR enum
	MemoryError_Type_MCA_INTERNALTIMER_ERROR MemoryError_Type = 59
	// MCA_MICROCODE_ROM_PARITY_ERROR enum
	MemoryError_Type_MCA_MICROCODE_ROM_PARITY_ERROR MemoryError_Type = 60
	// MCA_EXTERNAL_ERROR enum
	MemoryError_Type_MCA_EXTERNAL_ERROR MemoryError_Type = 61
	// MCA_FRC_ERROR enum
	MemoryError_Type_MCA_FRC_ERROR MemoryError_Type = 62
)
