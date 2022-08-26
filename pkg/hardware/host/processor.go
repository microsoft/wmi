// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package host

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"

	"github.com/microsoft/wmi/server2019/root/cimv2"
)

type TotalProcessor struct {
	Cores             uint32
	LogicalProcessors uint32
	Manufacturer      string
	Virtualization    bool
}

type ProcessorInfo struct {
	Manufacturer string
	//CPUType      *cimv2.Win32_Processor
}

type Processor struct {
	*cimv2.Win32_Processor
}

// NewPhysicalMemory
func NewProcessor(instance *wmi.WmiInstance) (*Processor, error) {
	wmivm, err := cimv2.NewWin32_ProcessorEx1(instance)
	if err != nil {
		return nil, err
	}
	return &Processor{wmivm}, nil
}

// GetTotalProcessor
func GetTotalProcessor(whost *host.WmiHost) (proc *TotalProcessor, err error) {
	query := query.NewWmiQuery("Win32_Processor")
	processors, err := instance.GetWmiInstancesFromHost(whost, string(constant.CimV2), query)
	if err != nil {
		return
	}
	defer processors.Close()

	totalCores := uint32(0)
	totalLogicalProcessors := uint32(0)
	var manufac string
	var virtualizationFlag bool

	for _, tmp := range processors {
		procInstance, err1 := cimv2.NewWin32_ProcessorEx1(tmp)
		if err1 != nil {
			err = err1
			return
		}

		cores, err1 := procInstance.GetProperty("NumberOfCores")
		if err1 != nil {
			err = err1
			return
		}
		totalCores = totalCores + uint32(cores.(int32))
		lp, err1 := procInstance.GetProperty("NumberOfLogicalProcessors")
		if err1 != nil {
			err = err1
			return
		}
		totalLogicalProcessors = totalLogicalProcessors + uint32(lp.(int32))

		manuf, err1 := procInstance.GetProperty("Manufacturer")
		if err1 != nil {
			err = err1
			return
		}
		manufac = manuf.(string)

		virtualizationFlag, err1 := procInstance.GetProperty("VirtualizationFirmwareEnabled")
		if err1 != nil {
			err = err1
			return
		}
		virtualizationFlag = virtualizationFlag.(bool)
	}
	//procInfo, err := GetProcessorInfo(whost)

	return &TotalProcessor{
		Cores:             totalCores,
		LogicalProcessors: totalLogicalProcessors,
		Manufacturer:      manufac,
		Virtualization:    virtualizationFlag,
	}, nil
}

// GetProcessorInfo
func GetProcessorInfo(whost *host.WmiHost) (proc *ProcessorInfo, err error) {
	query := query.NewWmiQuery("Win32_Processor")
	procInfo, err := instance.GetWmiInstanceEx(whost, string(constant.CimV2), query)
	if err != nil {
		return
	}
	defer procInfo.Close()

	procInstance, err1 := cimv2.NewWin32_ProcessorEx1(procInfo)
	if err1 != nil {
		err = err1
		return
	}

	manuf, err1 := procInstance.GetProperty("Manufacturer")
	if err1 != nil {
		err = err1
		return
	}

	/*procType, err1 := procInstance.GetProperty("ProcessorType")
	  if err1 != nil {
	  	err = err1
	  	return
	  }*/

	return &ProcessorInfo{
		Manufacturer: manuf.(string),
		//CPUType:      procType.(int32),
	}, nil
}
