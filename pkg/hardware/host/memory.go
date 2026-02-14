// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package host

import (
	"strconv"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/cimv2"
)

type Memory struct {
	SizeBytes uint64
}

type PhysicalMemory struct {
	*cimv2.Win32_PhysicalMemory
}

// NewPhysicalMemory
func NewPhysicalMemory(instance *wmi.WmiInstance) (*PhysicalMemory, error) {
	wmivm, err := cimv2.NewWin32_PhysicalMemoryEx1(instance)
	if err != nil {
		return nil, err
	}
	return &PhysicalMemory{wmivm}, nil
}

// GetTotalPhysicalMemory
func GetTotalPhysicalMemory(whost *host.WmiHost) (mem *Memory, err error) {
	query := query.NewWmiQuery("Win32_PhysicalMemory")
	memories, err := instance.GetWmiInstancesFromHost(whost, string(constant.CimV2), query)
	if err != nil {
		return
	}
	defer memories.Close()

	totalMemoryBytes := uint64(0)

	for _, tmp := range memories {
		memInstance, err1 := NewPhysicalMemory(tmp)
		if err1 != nil {
			err = err1
			return
		}

		memsizeVar, err1 := memInstance.GetProperty("Capacity")
		if err1 != nil {
			err = err1
			return
		}
		if memsizeVar == nil {
			err = errors.Wrapf(errors.NotFound, "Capacity property is nil")
			return
		}
		memsize, err1 := strconv.ParseUint(memsizeVar.(string), 10, 64)
		if err1 != nil {
			err = err1
			return
		}
		totalMemoryBytes = totalMemoryBytes + memsize
	}

	return &Memory{
		SizeBytes: totalMemoryBytes,
	}, nil
}

// GetFreePhysicalMemory
func GetFreePhysicalMemory(whost *host.WmiHost) (mem *Memory, err error) {
	query := query.NewWmiQuery("Win32_OperatingSystem")
	osinstance, err := instance.GetWmiInstanceEx(whost, string(constant.CimV2), query)
	if err != nil {
		return
	}
	defer osinstance.Close()

	tmpInstance, err1 := cimv2.NewWin32_OperatingSystemEx1(osinstance)
	if err1 != nil {
		err = err1
		return
	}

	memsizeVar, err1 := tmpInstance.GetProperty("FreePhysicalMemory")
	if err1 != nil {
		err = err1
		return
	}
	if memsizeVar == nil {
		err = errors.Wrapf(errors.NotFound, "FreePhysicalMemory property is nil")
		return
	}
	memsize, err1 := strconv.ParseUint(memsizeVar.(string), 10, 64)
	if err1 != nil {
		err = err1
		return
	}

	return &Memory{
		SizeBytes: memsize * 1024,
	}, nil
}
