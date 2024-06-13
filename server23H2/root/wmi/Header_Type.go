// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source Header_Type
//////////////////////////////////////////////
package wmi

// Header_Type
type Header_Type int

const (
	// MCA_WARNING_CACHE enum
	Header_Type_MCA_WARNING_CACHE Header_Type = 1
	// MCA_ERROR_CACHE enum
	Header_Type_MCA_ERROR_CACHE Header_Type = 2
	// MCA_WARNING_TLB enum
	Header_Type_MCA_WARNING_TLB Header_Type = 3
	// MCA_ERROR_TLB enum
	Header_Type_MCA_ERROR_TLB Header_Type = 4
	// MCA_WARNING_CPU_BUS enum
	Header_Type_MCA_WARNING_CPU_BUS Header_Type = 5
	// MCA_ERROR_CPU_BUS enum
	Header_Type_MCA_ERROR_CPU_BUS Header_Type = 6
	// MCA_WARNING_REGISTER_FILE enum
	Header_Type_MCA_WARNING_REGISTER_FILE Header_Type = 7
	// MCA_ERROR_REGISTER_FILE enum
	Header_Type_MCA_ERROR_REGISTER_FILE Header_Type = 8
	// MCA_WARNING_MAS enum
	Header_Type_MCA_WARNING_MAS Header_Type = 9
	// MCA_ERROR_MAS enum
	Header_Type_MCA_ERROR_MAS Header_Type = 10
	// MCA_WARNING_MEM_UNKNOWN enum
	Header_Type_MCA_WARNING_MEM_UNKNOWN Header_Type = 11
	// MCA_ERROR_MEM_UNKNOWN enum
	Header_Type_MCA_ERROR_MEM_UNKNOWN Header_Type = 12
	// MCA_WARNING_MEM_1_2 enum
	Header_Type_MCA_WARNING_MEM_1_2 Header_Type = 13
	// MCA_ERROR_MEM_1_2 enum
	Header_Type_MCA_ERROR_MEM_1_2 Header_Type = 14
	// MCA_WARNING_MEM_1_2_5 enum
	Header_Type_MCA_WARNING_MEM_1_2_5 Header_Type = 15
	// MCA_ERROR_MEM_1_2_5 enum
	Header_Type_MCA_ERROR_MEM_1_2_5 Header_Type = 16
	// MCA_WARNING_MEM_1_2_5_4 enum
	Header_Type_MCA_WARNING_MEM_1_2_5_4 Header_Type = 17
	// MCA_ERROR_MEM_1_2_5_4 enum
	Header_Type_MCA_ERROR_MEM_1_2_5_4 Header_Type = 18
	// MCA_WARNING_SYSTEM_EVENT enum
	Header_Type_MCA_WARNING_SYSTEM_EVENT Header_Type = 19
	// MCA_ERROR_SYSTEM_EVENT enum
	Header_Type_MCA_ERROR_SYSTEM_EVENT Header_Type = 20
	// MCA_WARNING_PCI_BUS_PARITY enum
	Header_Type_MCA_WARNING_PCI_BUS_PARITY Header_Type = 21
	// MCA_ERROR_PCI_BUS_PARITY enum
	Header_Type_MCA_ERROR_PCI_BUS_PARITY Header_Type = 22
	// MCA_WARNING_PCI_BUS_PARITY_NO_INFO enum
	Header_Type_MCA_WARNING_PCI_BUS_PARITY_NO_INFO Header_Type = 23
	// MCA_ERROR_PCI_BUS_PARITY_NO_INFO enum
	Header_Type_MCA_ERROR_PCI_BUS_PARITY_NO_INFO Header_Type = 24
	// MCA_WARNING_PCI_BUS_SERR enum
	Header_Type_MCA_WARNING_PCI_BUS_SERR Header_Type = 25
	// MCA_ERROR_PCI_BUS_SERR enum
	Header_Type_MCA_ERROR_PCI_BUS_SERR Header_Type = 26
	// MCA_WARNING_PCI_BUS_SERR_NO_INFO enum
	Header_Type_MCA_WARNING_PCI_BUS_SERR_NO_INFO Header_Type = 27
	// MCA_ERROR_PCI_BUS_SERR_NO_INFO enum
	Header_Type_MCA_ERROR_PCI_BUS_SERR_NO_INFO Header_Type = 28
	// MCA_WARNING_PCI_BUS_MASTER_ABORT enum
	Header_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT Header_Type = 29
	// MCA_ERROR_PCI_BUS_MASTER_ABORT enum
	Header_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT Header_Type = 30
	// MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO enum
	Header_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO Header_Type = 31
	// MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO enum
	Header_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO Header_Type = 32
	// MCA_WARNING_PCI_BUS_TIMEOUT enum
	Header_Type_MCA_WARNING_PCI_BUS_TIMEOUT Header_Type = 33
	// MCA_ERROR_PCI_BUS_TIMEOUT enum
	Header_Type_MCA_ERROR_PCI_BUS_TIMEOUT Header_Type = 34
	// MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO enum
	Header_Type_MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO Header_Type = 35
	// MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO enum
	Header_Type_MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO Header_Type = 36
	// MCA_WARNING_PCI_BUS_UNKNOWN enum
	Header_Type_MCA_WARNING_PCI_BUS_UNKNOWN Header_Type = 37
	// MCA_ERROR_PCI_BUS_UNKNOWN enum
	Header_Type_MCA_ERROR_PCI_BUS_UNKNOWN Header_Type = 38
	// MCA_WARNING_PCI_DEVICE enum
	Header_Type_MCA_WARNING_PCI_DEVICE Header_Type = 39
	// MCA_ERROR_PCI_DEVICE enum
	Header_Type_MCA_ERROR_PCI_DEVICE Header_Type = 40
	// MCA_WARNING_SMBIOS enum
	Header_Type_MCA_WARNING_SMBIOS Header_Type = 41
	// MCA_ERROR_SMBIOS enum
	Header_Type_MCA_ERROR_SMBIOS Header_Type = 42
	// MCA_WARNING_PLATFORM_SPECIFIC enum
	Header_Type_MCA_WARNING_PLATFORM_SPECIFIC Header_Type = 43
	// MCA_ERROR_PLATFORM_SPECIFIC enum
	Header_Type_MCA_ERROR_PLATFORM_SPECIFIC Header_Type = 44
	// MCA_WARNING_UNKNOWN enum
	Header_Type_MCA_WARNING_UNKNOWN Header_Type = 45
	// MCA_ERROR_UNKNOWN enum
	Header_Type_MCA_ERROR_UNKNOWN Header_Type = 46
	// MCA_WARNING_UNKNOWN_NO_CPU enum
	Header_Type_MCA_WARNING_UNKNOWN_NO_CPU Header_Type = 47
	// MCA_ERROR_UNKNOWN_NO_CPU enum
	Header_Type_MCA_ERROR_UNKNOWN_NO_CPU Header_Type = 48
	// MCA_WARNING_CMC_THRESHOLD_EXCEEDED enum
	Header_Type_MCA_WARNING_CMC_THRESHOLD_EXCEEDED Header_Type = 49
	// MCA_WARNING_CPE_THRESHOLD_EXCEEDED enum
	Header_Type_MCA_WARNING_CPE_THRESHOLD_EXCEEDED Header_Type = 50
	// MCA_WARNING_CPU_THERMAL_THROTTLED enum
	Header_Type_MCA_WARNING_CPU_THERMAL_THROTTLED Header_Type = 51
	// MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED enum
	Header_Type_MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED Header_Type = 52
	// MCA_WARNING_CPU enum
	Header_Type_MCA_WARNING_CPU Header_Type = 53
	// MCA_ERROR_CPU enum
	Header_Type_MCA_ERROR_CPU Header_Type = 54
	// MCA_MEMORYHIERARCHY_ERROR enum
	Header_Type_MCA_MEMORYHIERARCHY_ERROR Header_Type = 55
	// MCA_TLB_ERROR enum
	Header_Type_MCA_TLB_ERROR Header_Type = 56
	// MCA_BUS_ERROR enum
	Header_Type_MCA_BUS_ERROR Header_Type = 57
	// MCA_BUS_TIMEOUT_ERROR enum
	Header_Type_MCA_BUS_TIMEOUT_ERROR Header_Type = 58
	// MCA_INTERNALTIMER_ERROR enum
	Header_Type_MCA_INTERNALTIMER_ERROR Header_Type = 59
	// MCA_MICROCODE_ROM_PARITY_ERROR enum
	Header_Type_MCA_MICROCODE_ROM_PARITY_ERROR Header_Type = 60
	// MCA_EXTERNAL_ERROR enum
	Header_Type_MCA_EXTERNAL_ERROR Header_Type = 61
	// MCA_FRC_ERROR enum
	Header_Type_MCA_FRC_ERROR Header_Type = 62
)
