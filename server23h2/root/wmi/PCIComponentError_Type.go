// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source PCIComponentError_Type
//////////////////////////////////////////////
package wmi

// PCIComponentError_Type
type PCIComponentError_Type int

const (
	// MCA_WARNING_CACHE enum
	PCIComponentError_Type_MCA_WARNING_CACHE PCIComponentError_Type = 1
	// MCA_ERROR_CACHE enum
	PCIComponentError_Type_MCA_ERROR_CACHE PCIComponentError_Type = 2
	// MCA_WARNING_TLB enum
	PCIComponentError_Type_MCA_WARNING_TLB PCIComponentError_Type = 3
	// MCA_ERROR_TLB enum
	PCIComponentError_Type_MCA_ERROR_TLB PCIComponentError_Type = 4
	// MCA_WARNING_CPU_BUS enum
	PCIComponentError_Type_MCA_WARNING_CPU_BUS PCIComponentError_Type = 5
	// MCA_ERROR_CPU_BUS enum
	PCIComponentError_Type_MCA_ERROR_CPU_BUS PCIComponentError_Type = 6
	// MCA_WARNING_REGISTER_FILE enum
	PCIComponentError_Type_MCA_WARNING_REGISTER_FILE PCIComponentError_Type = 7
	// MCA_ERROR_REGISTER_FILE enum
	PCIComponentError_Type_MCA_ERROR_REGISTER_FILE PCIComponentError_Type = 8
	// MCA_WARNING_MAS enum
	PCIComponentError_Type_MCA_WARNING_MAS PCIComponentError_Type = 9
	// MCA_ERROR_MAS enum
	PCIComponentError_Type_MCA_ERROR_MAS PCIComponentError_Type = 10
	// MCA_WARNING_MEM_UNKNOWN enum
	PCIComponentError_Type_MCA_WARNING_MEM_UNKNOWN PCIComponentError_Type = 11
	// MCA_ERROR_MEM_UNKNOWN enum
	PCIComponentError_Type_MCA_ERROR_MEM_UNKNOWN PCIComponentError_Type = 12
	// MCA_WARNING_MEM_1_2 enum
	PCIComponentError_Type_MCA_WARNING_MEM_1_2 PCIComponentError_Type = 13
	// MCA_ERROR_MEM_1_2 enum
	PCIComponentError_Type_MCA_ERROR_MEM_1_2 PCIComponentError_Type = 14
	// MCA_WARNING_MEM_1_2_5 enum
	PCIComponentError_Type_MCA_WARNING_MEM_1_2_5 PCIComponentError_Type = 15
	// MCA_ERROR_MEM_1_2_5 enum
	PCIComponentError_Type_MCA_ERROR_MEM_1_2_5 PCIComponentError_Type = 16
	// MCA_WARNING_MEM_1_2_5_4 enum
	PCIComponentError_Type_MCA_WARNING_MEM_1_2_5_4 PCIComponentError_Type = 17
	// MCA_ERROR_MEM_1_2_5_4 enum
	PCIComponentError_Type_MCA_ERROR_MEM_1_2_5_4 PCIComponentError_Type = 18
	// MCA_WARNING_SYSTEM_EVENT enum
	PCIComponentError_Type_MCA_WARNING_SYSTEM_EVENT PCIComponentError_Type = 19
	// MCA_ERROR_SYSTEM_EVENT enum
	PCIComponentError_Type_MCA_ERROR_SYSTEM_EVENT PCIComponentError_Type = 20
	// MCA_WARNING_PCI_BUS_PARITY enum
	PCIComponentError_Type_MCA_WARNING_PCI_BUS_PARITY PCIComponentError_Type = 21
	// MCA_ERROR_PCI_BUS_PARITY enum
	PCIComponentError_Type_MCA_ERROR_PCI_BUS_PARITY PCIComponentError_Type = 22
	// MCA_WARNING_PCI_BUS_PARITY_NO_INFO enum
	PCIComponentError_Type_MCA_WARNING_PCI_BUS_PARITY_NO_INFO PCIComponentError_Type = 23
	// MCA_ERROR_PCI_BUS_PARITY_NO_INFO enum
	PCIComponentError_Type_MCA_ERROR_PCI_BUS_PARITY_NO_INFO PCIComponentError_Type = 24
	// MCA_WARNING_PCI_BUS_SERR enum
	PCIComponentError_Type_MCA_WARNING_PCI_BUS_SERR PCIComponentError_Type = 25
	// MCA_ERROR_PCI_BUS_SERR enum
	PCIComponentError_Type_MCA_ERROR_PCI_BUS_SERR PCIComponentError_Type = 26
	// MCA_WARNING_PCI_BUS_SERR_NO_INFO enum
	PCIComponentError_Type_MCA_WARNING_PCI_BUS_SERR_NO_INFO PCIComponentError_Type = 27
	// MCA_ERROR_PCI_BUS_SERR_NO_INFO enum
	PCIComponentError_Type_MCA_ERROR_PCI_BUS_SERR_NO_INFO PCIComponentError_Type = 28
	// MCA_WARNING_PCI_BUS_MASTER_ABORT enum
	PCIComponentError_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT PCIComponentError_Type = 29
	// MCA_ERROR_PCI_BUS_MASTER_ABORT enum
	PCIComponentError_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT PCIComponentError_Type = 30
	// MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO enum
	PCIComponentError_Type_MCA_WARNING_PCI_BUS_MASTER_ABORT_NO_INFO PCIComponentError_Type = 31
	// MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO enum
	PCIComponentError_Type_MCA_ERROR_PCI_BUS_MASTER_ABORT_NO_INFO PCIComponentError_Type = 32
	// MCA_WARNING_PCI_BUS_TIMEOUT enum
	PCIComponentError_Type_MCA_WARNING_PCI_BUS_TIMEOUT PCIComponentError_Type = 33
	// MCA_ERROR_PCI_BUS_TIMEOUT enum
	PCIComponentError_Type_MCA_ERROR_PCI_BUS_TIMEOUT PCIComponentError_Type = 34
	// MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO enum
	PCIComponentError_Type_MCA_WARNING_PCI_BUS_TIMEOUT_NO_INFO PCIComponentError_Type = 35
	// MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO enum
	PCIComponentError_Type_MCA_ERROR_PCI_BUS_TIMEOUT_NO_INFO PCIComponentError_Type = 36
	// MCA_WARNING_PCI_BUS_UNKNOWN enum
	PCIComponentError_Type_MCA_WARNING_PCI_BUS_UNKNOWN PCIComponentError_Type = 37
	// MCA_ERROR_PCI_BUS_UNKNOWN enum
	PCIComponentError_Type_MCA_ERROR_PCI_BUS_UNKNOWN PCIComponentError_Type = 38
	// MCA_WARNING_PCI_DEVICE enum
	PCIComponentError_Type_MCA_WARNING_PCI_DEVICE PCIComponentError_Type = 39
	// MCA_ERROR_PCI_DEVICE enum
	PCIComponentError_Type_MCA_ERROR_PCI_DEVICE PCIComponentError_Type = 40
	// MCA_WARNING_SMBIOS enum
	PCIComponentError_Type_MCA_WARNING_SMBIOS PCIComponentError_Type = 41
	// MCA_ERROR_SMBIOS enum
	PCIComponentError_Type_MCA_ERROR_SMBIOS PCIComponentError_Type = 42
	// MCA_WARNING_PLATFORM_SPECIFIC enum
	PCIComponentError_Type_MCA_WARNING_PLATFORM_SPECIFIC PCIComponentError_Type = 43
	// MCA_ERROR_PLATFORM_SPECIFIC enum
	PCIComponentError_Type_MCA_ERROR_PLATFORM_SPECIFIC PCIComponentError_Type = 44
	// MCA_WARNING_UNKNOWN enum
	PCIComponentError_Type_MCA_WARNING_UNKNOWN PCIComponentError_Type = 45
	// MCA_ERROR_UNKNOWN enum
	PCIComponentError_Type_MCA_ERROR_UNKNOWN PCIComponentError_Type = 46
	// MCA_WARNING_UNKNOWN_NO_CPU enum
	PCIComponentError_Type_MCA_WARNING_UNKNOWN_NO_CPU PCIComponentError_Type = 47
	// MCA_ERROR_UNKNOWN_NO_CPU enum
	PCIComponentError_Type_MCA_ERROR_UNKNOWN_NO_CPU PCIComponentError_Type = 48
	// MCA_WARNING_CMC_THRESHOLD_EXCEEDED enum
	PCIComponentError_Type_MCA_WARNING_CMC_THRESHOLD_EXCEEDED PCIComponentError_Type = 49
	// MCA_WARNING_CPE_THRESHOLD_EXCEEDED enum
	PCIComponentError_Type_MCA_WARNING_CPE_THRESHOLD_EXCEEDED PCIComponentError_Type = 50
	// MCA_WARNING_CPU_THERMAL_THROTTLED enum
	PCIComponentError_Type_MCA_WARNING_CPU_THERMAL_THROTTLED PCIComponentError_Type = 51
	// MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED enum
	PCIComponentError_Type_MCA_INFO_CPU_THERMAL_THROTTLING_REMOVED PCIComponentError_Type = 52
	// MCA_WARNING_CPU enum
	PCIComponentError_Type_MCA_WARNING_CPU PCIComponentError_Type = 53
	// MCA_ERROR_CPU enum
	PCIComponentError_Type_MCA_ERROR_CPU PCIComponentError_Type = 54
	// MCA_MEMORYHIERARCHY_ERROR enum
	PCIComponentError_Type_MCA_MEMORYHIERARCHY_ERROR PCIComponentError_Type = 55
	// MCA_TLB_ERROR enum
	PCIComponentError_Type_MCA_TLB_ERROR PCIComponentError_Type = 56
	// MCA_BUS_ERROR enum
	PCIComponentError_Type_MCA_BUS_ERROR PCIComponentError_Type = 57
	// MCA_BUS_TIMEOUT_ERROR enum
	PCIComponentError_Type_MCA_BUS_TIMEOUT_ERROR PCIComponentError_Type = 58
	// MCA_INTERNALTIMER_ERROR enum
	PCIComponentError_Type_MCA_INTERNALTIMER_ERROR PCIComponentError_Type = 59
	// MCA_MICROCODE_ROM_PARITY_ERROR enum
	PCIComponentError_Type_MCA_MICROCODE_ROM_PARITY_ERROR PCIComponentError_Type = 60
	// MCA_EXTERNAL_ERROR enum
	PCIComponentError_Type_MCA_EXTERNAL_ERROR PCIComponentError_Type = 61
	// MCA_FRC_ERROR enum
	PCIComponentError_Type_MCA_FRC_ERROR PCIComponentError_Type = 62
)
