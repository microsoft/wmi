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
}

type Processor struct {
	*cimv2.Win32_Processor
}

type ComputerSystem struct {
	*cimv2.Win32_ComputerSystem
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
	}

	return &TotalProcessor{
		Cores:             totalCores,
		LogicalProcessors: totalLogicalProcessors,
	}, nil
}

func GetProcessor(whost *host.WmiHost) (proc *Processor, err error) {
	queryProcessor := query.NewWmiQuery("Win32_Processor") //query for fields from Win32_processor class

	procInfo, err := instance.GetWmiInstanceEx(whost, string(constant.CimV2), queryProcessor)
	if err != nil {
		return nil, err
	}
	procInstance, err := cimv2.NewWin32_ProcessorEx1(procInfo)
	if err != nil {
		return nil, err
	}
	return &Processor{procInstance}, nil
}

func GetComputerSystem(whost *host.WmiHost) (cs *ComputerSystem, err error) {
	queryComputerSystem := query.NewWmiQuery("Win32_ComputerSystem") //query for fields from Win32_ComputerSystem class

	winComputerSystemInfo, err := instance.GetWmiInstanceEx(whost, string(constant.CimV2), queryComputerSystem)
	if err != nil {
		return nil, err
	}
	winCompSystemInstance, err := cimv2.NewWin32_ComputerSystemEx1(winComputerSystemInfo)
	if err != nil {
		return nil, err
	}
	return &ComputerSystem{winCompSystemInstance}, nil
}
