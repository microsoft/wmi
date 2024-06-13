// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source SystemEventError_Type
//////////////////////////////////////////////
package wmi

// SystemEventError_Type
type SystemEventError_Type int

const (
	// MCA_WARNING_CACHE enum
	SystemEventError_Type_MCA_WARNING_CACHE SystemEventError_Type = 1
	// MCA_ERROR_CACHE enum
	SystemEventError_Type_MCA_ERROR_CACHE SystemEventError_Type = 2
	// MCA_WARNING_TLB enum
	SystemEventError_Type_MCA_WARNING_TLB SystemEventError_Type = 3
	// MCA_ERROR_TLB enum
	SystemEventError_Type_MCA_ERROR_TLB SystemEventError_Type = 4
	// MCA_WARNING_CPU_BUS enum
	SystemEventError_Type_MCA_WARNING_CPU_BUS SystemEventError_Type = 5
	// MCA_ERROR_CPU_BUS enum
	SystemEventError_Type_MCA_ERROR_CPU_BUS SystemEventError_Type = 6
	// MCA_WARNING_REGISTER_FILE enum
	SystemEventError_Type_MCA_WARNING_REGISTER_FILE SystemEventError_Type = 7
	// MCA_ERROR_REGISTER_FILE enum
	SystemEventError_Type_MCA_ERROR_REGISTER_FILE SystemEventError_Type = 8
	// MCA_WARNING_MAS enum
	SystemEventError_Type_MCA_WARNING_MAS SystemEventError_Type = 9
	// MCA_ERROR_MAS enum
	SystemEventError_Type_MCA_ERROR_MAS SystemEventError_Type = 10
	// MCA_WARNING_MEM_UNKNOWN enum
	SystemEventError_Type_MCA_WARNING_MEM_UNKNOWN SystemEventError_Type = 11
	// MCA_ERROR_MEM_UNKNOWN enum
	SystemEventError_Type_MCA_ERROR_MEM_UNKNOWN SystemEventError_Type = 12
	// MCA_WARNING_MEM_1_2 enum
	SystemEventError_Type_MCA_WARNING_MEM_1_2 SystemEventError_Type = 13
	// MCA_ERROR_MEM_1_2 enum
	SystemEventError_Type_MCA_ERROR_MEM_1_2 SystemEventError_Type = 14
	// MCA_WARNING_MEM_1_2_5 enum
	SystemEventError_Type_MCA_WARNING_MEM_1_2_5 SystemEventError_Type = 15
	// MCA_ERROR_MEM_1_2_5 enum
	SystemEventError_Type_MCA_ERROR_MEM_1_2_5 SystemEventError_Type = 16
	// MCA_WARNING_MEM_1_2_5_4 enum
	SystemEventError_Type_MCA_WARNING_MEM_1_2_5_4 SystemEventError_Type = 17
	// MCA_ERROR_MEM_1_2_5_4 enum
	SystemEventError_Type_MCA_ERROR_MEM_1_2_5_4 SystemEventError_Type = 18
	// MCA_WARNING_SYSTEM_EVENT enum
	SystemEventError_Type_MCA_WARNING_SYSTEM_EVENT SystemEventError_Type = 19
	// MCA_ERROR_SYSTEM_EVENT enum
	SystemEventError_Type_MCA_ERROR_SYSTEM_EVENT SystemEventError_Type = 20
	// MCA_WARNING_PCI_BUS_PARITY enum
	SystemEventError_Type_MCA_WARNING_PCI_BUS_PARITY SystemEventError_Type = 21
	// MCA_ERROR_PCI_BUS_PARITY enum
	SystemEventError_Type_MCA_ERROR_PCI_BUS_PARITY SystemEventError_Type = 22
	// MCA_WARNING_PCI_BUS_PARITY_NO_INFO enum
	SystemEventError_Type_MCA_WARNING_PCI_BUS_PARITY_NO_INFO SystemEventError_Type = 23
	// MCA_ERROR_PCI_BUS_PARITY_NO_INFO enum
	SystemEventError_Type_MCA_ERROR_PCI_BUS_PARITY_NO_INFO SystemEventError_Type = 24
	// MCA_WARNING_PCI_BUS_SERR enum
	SystemEventError_Type_MCA_WARNING_PCI_BUS_SERR SystemEventError_Type = 25
	// MCA_ERROR_PCI_BUS_SERR enum
	SystemEventError_Type_MCA_ERROR_PCI_BUS_SERR SystemEventError_Type = 26
	// MCA_WARNING_PCI_BUS_SERR_NO_INFO enum
	SystemEventError_Type_MCA_WARNING_PCI_BUS_SERR_NO_INFO SystemEventError_Type = 27
	// MCA_ERROR_PCI_BUS_SERR_NO_INFO enum
	SystemEventError_Type_MCA_ERROR_PCI_BUS_SERR_NO_INFO SystemEventError_Type = 28
	// MCA_WARNING_PCI_BUS_MASTER_ABORT enum
	SystemEventError_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT SystemEventError_Type = 29
	// MCA_ERROR_PCI_BUS_MASTER_ABORT enum
	SystemEventError_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT SystemEventError_Type = 30
	// MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO enum
	SystemEventError_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO SystemEventError_Type = 31
	// MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO enum
	SystemEventError_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO SystemEventError_Type = 32
	// MCA_WARNING_PCI_BUS_TIMEOUT enum
	SystemEventError_Type_MCA_WARNING_PCI_BUS_TIMEOUT SystemEventError_Type = 33
	// MCA_ERROR_PCI_BUS_TIMEOUT enum
	SystemEventError_Type_MCA_ERROR_PCI_BUS_TIMEOUT SystemEventError_Type = 34
	// MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO enum
	SystemEventError_Type_MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO SystemEventError_Type = 35
	// MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO enum
	SystemEventError_Type_MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO SystemEventError_Type = 36
	// MCA_WARNING_PCI_BUS_UNKNOWN enum
	SystemEventError_Type_MCA_WARNING_PCI_BUS_UNKNOWN SystemEventError_Type = 37
	// MCA_ERROR_PCI_BUS_UNKNOWN enum
	SystemEventError_Type_MCA_ERROR_PCI_BUS_UNKNOWN SystemEventError_Type = 38
	// MCA_WARNING_PCI_DEVICE enum
	SystemEventError_Type_MCA_WARNING_PCI_DEVICE SystemEventError_Type = 39
	// MCA_ERROR_PCI_DEVICE enum
	SystemEventError_Type_MCA_ERROR_PCI_DEVICE SystemEventError_Type = 40
	// MCA_WARNING_SMBIOS enum
	SystemEventError_Type_MCA_WARNING_SMBIOS SystemEventError_Type = 41
	// MCA_ERROR_SMBIOS enum
	SystemEventError_Type_MCA_ERROR_SMBIOS SystemEventError_Type = 42
	// MCA_WARNING_PLATFORM_SPECIFIC enum
	SystemEventError_Type_MCA_WARNING_PLATFORM_SPECIFIC SystemEventError_Type = 43
	// MCA_ERROR_PLATFORM_SPECIFIC enum
	SystemEventError_Type_MCA_ERROR_PLATFORM_SPECIFIC SystemEventError_Type = 44
	// MCA_WARNING_UNKNOWN enum
	SystemEventError_Type_MCA_WARNING_UNKNOWN SystemEventError_Type = 45
	// MCA_ERROR_UNKNOWN enum
	SystemEventError_Type_MCA_ERROR_UNKNOWN SystemEventError_Type = 46
	// MCA_WARNING_UNKNOWN_NO_CPU enum
	SystemEventError_Type_MCA_WARNING_UNKNOWN_NO_CPU SystemEventError_Type = 47
	// MCA_ERROR_UNKNOWN_NO_CPU enum
	SystemEventError_Type_MCA_ERROR_UNKNOWN_NO_CPU SystemEventError_Type = 48
	// MCA_WARNING_CMC_THRESHOLD_EXCEEDED enum
	SystemEventError_Type_MCA_WARNING_CMC_THRESHOLD_EXCEEDED SystemEventError_Type = 49
	// MCA_WARNING_CPE_THRESHOLD_EXCEEDED enum
	SystemEventError_Type_MCA_WARNING_CPE_THRESHOLD_EXCEEDED SystemEventError_Type = 50
	// MCA_WARNING_CPU_THERMAL_THROTTLED enum
	SystemEventError_Type_MCA_WARNING_CPU_THERMAL_THROTTLED SystemEventError_Type = 51
	// MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED enum
	SystemEventError_Type_MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED SystemEventError_Type = 52
	// MCA_WARNING_CPU enum
	SystemEventError_Type_MCA_WARNING_CPU SystemEventError_Type = 53
	// MCA_ERROR_CPU enum
	SystemEventError_Type_MCA_ERROR_CPU SystemEventError_Type = 54
	// MCA_MEMORYHIERARCHY_ERROR enum
	SystemEventError_Type_MCA_MEMORYHIERARCHY_ERROR SystemEventError_Type = 55
	// MCA_TLB_ERROR enum
	SystemEventError_Type_MCA_TLB_ERROR SystemEventError_Type = 56
	// MCA_BUS_ERROR enum
	SystemEventError_Type_MCA_BUS_ERROR SystemEventError_Type = 57
	// MCA_BUS_TIMEOUT_ERROR enum
	SystemEventError_Type_MCA_BUS_TIMEOUT_ERROR SystemEventError_Type = 58
	// MCA_INTERNALTIMER_ERROR enum
	SystemEventError_Type_MCA_INTERNALTIMER_ERROR SystemEventError_Type = 59
	// MCA_MICROCODE_ROM_PARITY_ERROR enum
	SystemEventError_Type_MCA_MICROCODE_ROM_PARITY_ERROR SystemEventError_Type = 60
	// MCA_EXTERNAL_ERROR enum
	SystemEventError_Type_MCA_EXTERNAL_ERROR SystemEventError_Type = 61
	// MCA_FRC_ERROR enum
	SystemEventError_Type_MCA_FRC_ERROR SystemEventError_Type = 62
)
