// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source BusError_Type
//
// ////////////////////////////////////////////
package wmi

// BusError_Type
type BusError_Type int

const (
	// MCA_WARNING_CACHE enum
	BusError_Type_MCA_WARNING_CACHE BusError_Type = 1
	// MCA_ERROR_CACHE enum
	BusError_Type_MCA_ERROR_CACHE BusError_Type = 2
	// MCA_WARNING_TLB enum
	BusError_Type_MCA_WARNING_TLB BusError_Type = 3
	// MCA_ERROR_TLB enum
	BusError_Type_MCA_ERROR_TLB BusError_Type = 4
	// MCA_WARNING_CPU_BUS enum
	BusError_Type_MCA_WARNING_CPU_BUS BusError_Type = 5
	// MCA_ERROR_CPU_BUS enum
	BusError_Type_MCA_ERROR_CPU_BUS BusError_Type = 6
	// MCA_WARNING_REGISTER_FILE enum
	BusError_Type_MCA_WARNING_REGISTER_FILE BusError_Type = 7
	// MCA_ERROR_REGISTER_FILE enum
	BusError_Type_MCA_ERROR_REGISTER_FILE BusError_Type = 8
	// MCA_WARNING_MAS enum
	BusError_Type_MCA_WARNING_MAS BusError_Type = 9
	// MCA_ERROR_MAS enum
	BusError_Type_MCA_ERROR_MAS BusError_Type = 10
	// MCA_WARNING_MEM_UNKNOWN enum
	BusError_Type_MCA_WARNING_MEM_UNKNOWN BusError_Type = 11
	// MCA_ERROR_MEM_UNKNOWN enum
	BusError_Type_MCA_ERROR_MEM_UNKNOWN BusError_Type = 12
	// MCA_WARNING_MEM_1_2 enum
	BusError_Type_MCA_WARNING_MEM_1_2 BusError_Type = 13
	// MCA_ERROR_MEM_1_2 enum
	BusError_Type_MCA_ERROR_MEM_1_2 BusError_Type = 14
	// MCA_WARNING_MEM_1_2_5 enum
	BusError_Type_MCA_WARNING_MEM_1_2_5 BusError_Type = 15
	// MCA_ERROR_MEM_1_2_5 enum
	BusError_Type_MCA_ERROR_MEM_1_2_5 BusError_Type = 16
	// MCA_WARNING_MEM_1_2_5_4 enum
	BusError_Type_MCA_WARNING_MEM_1_2_5_4 BusError_Type = 17
	// MCA_ERROR_MEM_1_2_5_4 enum
	BusError_Type_MCA_ERROR_MEM_1_2_5_4 BusError_Type = 18
	// MCA_WARNING_SYSTEM_EVENT enum
	BusError_Type_MCA_WARNING_SYSTEM_EVENT BusError_Type = 19
	// MCA_ERROR_SYSTEM_EVENT enum
	BusError_Type_MCA_ERROR_SYSTEM_EVENT BusError_Type = 20
	// MCA_WARNING_PCI_BUS_PARITY enum
	BusError_Type_MCA_WARNING_PCI_BUS_PARITY BusError_Type = 21
	// MCA_ERROR_PCI_BUS_PARITY enum
	BusError_Type_MCA_ERROR_PCI_BUS_PARITY BusError_Type = 22
	// MCA_WARNING_PCI_BUS_PARITY_NO_INFO enum
	BusError_Type_MCA_WARNING_PCI_BUS_PARITY_NO_INFO BusError_Type = 23
	// MCA_ERROR_PCI_BUS_PARITY_NO_INFO enum
	BusError_Type_MCA_ERROR_PCI_BUS_PARITY_NO_INFO BusError_Type = 24
	// MCA_WARNING_PCI_BUS_SERR enum
	BusError_Type_MCA_WARNING_PCI_BUS_SERR BusError_Type = 25
	// MCA_ERROR_PCI_BUS_SERR enum
	BusError_Type_MCA_ERROR_PCI_BUS_SERR BusError_Type = 26
	// MCA_WARNING_PCI_BUS_SERR_NO_INFO enum
	BusError_Type_MCA_WARNING_PCI_BUS_SERR_NO_INFO BusError_Type = 27
	// MCA_ERROR_PCI_BUS_SERR_NO_INFO enum
	BusError_Type_MCA_ERROR_PCI_BUS_SERR_NO_INFO BusError_Type = 28
	// MCA_WARNING_PCI_BUS_MASTER_ABORT enum
	BusError_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT BusError_Type = 29
	// MCA_ERROR_PCI_BUS_MASTER_ABORT enum
	BusError_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT BusError_Type = 30
	// MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO enum
	BusError_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO BusError_Type = 31
	// MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO enum
	BusError_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO BusError_Type = 32
	// MCA_WARNING_PCI_BUS_TIMEOUT enum
	BusError_Type_MCA_WARNING_PCI_BUS_TIMEOUT BusError_Type = 33
	// MCA_ERROR_PCI_BUS_TIMEOUT enum
	BusError_Type_MCA_ERROR_PCI_BUS_TIMEOUT BusError_Type = 34
	// MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO enum
	BusError_Type_MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO BusError_Type = 35
	// MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO enum
	BusError_Type_MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO BusError_Type = 36
	// MCA_WARNING_PCI_BUS_UNKNOWN enum
	BusError_Type_MCA_WARNING_PCI_BUS_UNKNOWN BusError_Type = 37
	// MCA_ERROR_PCI_BUS_UNKNOWN enum
	BusError_Type_MCA_ERROR_PCI_BUS_UNKNOWN BusError_Type = 38
	// MCA_WARNING_PCI_DEVICE enum
	BusError_Type_MCA_WARNING_PCI_DEVICE BusError_Type = 39
	// MCA_ERROR_PCI_DEVICE enum
	BusError_Type_MCA_ERROR_PCI_DEVICE BusError_Type = 40
	// MCA_WARNING_SMBIOS enum
	BusError_Type_MCA_WARNING_SMBIOS BusError_Type = 41
	// MCA_ERROR_SMBIOS enum
	BusError_Type_MCA_ERROR_SMBIOS BusError_Type = 42
	// MCA_WARNING_PLATFORM_SPECIFIC enum
	BusError_Type_MCA_WARNING_PLATFORM_SPECIFIC BusError_Type = 43
	// MCA_ERROR_PLATFORM_SPECIFIC enum
	BusError_Type_MCA_ERROR_PLATFORM_SPECIFIC BusError_Type = 44
	// MCA_WARNING_UNKNOWN enum
	BusError_Type_MCA_WARNING_UNKNOWN BusError_Type = 45
	// MCA_ERROR_UNKNOWN enum
	BusError_Type_MCA_ERROR_UNKNOWN BusError_Type = 46
	// MCA_WARNING_UNKNOWN_NO_CPU enum
	BusError_Type_MCA_WARNING_UNKNOWN_NO_CPU BusError_Type = 47
	// MCA_ERROR_UNKNOWN_NO_CPU enum
	BusError_Type_MCA_ERROR_UNKNOWN_NO_CPU BusError_Type = 48
	// MCA_WARNING_CMC_THRESHOLD_EXCEEDED enum
	BusError_Type_MCA_WARNING_CMC_THRESHOLD_EXCEEDED BusError_Type = 49
	// MCA_WARNING_CPE_THRESHOLD_EXCEEDED enum
	BusError_Type_MCA_WARNING_CPE_THRESHOLD_EXCEEDED BusError_Type = 50
	// MCA_WARNING_CPU_THERMAL_THROTTLED enum
	BusError_Type_MCA_WARNING_CPU_THERMAL_THROTTLED BusError_Type = 51
	// MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED enum
	BusError_Type_MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED BusError_Type = 52
	// MCA_WARNING_CPU enum
	BusError_Type_MCA_WARNING_CPU BusError_Type = 53
	// MCA_ERROR_CPU enum
	BusError_Type_MCA_ERROR_CPU BusError_Type = 54
	// MCA_MEMORYHIERARCHY_ERROR enum
	BusError_Type_MCA_MEMORYHIERARCHY_ERROR BusError_Type = 55
	// MCA_TLB_ERROR enum
	BusError_Type_MCA_TLB_ERROR BusError_Type = 56
	// MCA_BUS_ERROR enum
	BusError_Type_MCA_BUS_ERROR BusError_Type = 57
	// MCA_BUS_TIMEOUT_ERROR enum
	BusError_Type_MCA_BUS_TIMEOUT_ERROR BusError_Type = 58
	// MCA_INTERNALTIMER_ERROR enum
	BusError_Type_MCA_INTERNALTIMER_ERROR BusError_Type = 59
	// MCA_MICROCODE_ROM_PARITY_ERROR enum
	BusError_Type_MCA_MICROCODE_ROM_PARITY_ERROR BusError_Type = 60
	// MCA_EXTERNAL_ERROR enum
	BusError_Type_MCA_EXTERNAL_ERROR BusError_Type = 61
	// MCA_FRC_ERROR enum
	BusError_Type_MCA_FRC_ERROR BusError_Type = 62
)
