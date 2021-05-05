// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package memory

import (
	"strconv"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type MemorySettingData struct {
	*v2.Msvm_MemorySettingData
}

// NewMemorySettingData
func NewMemorySettingData(instance *wmi.WmiInstance) (*MemorySettingData, error) {
	wmivm, err := v2.NewMsvm_MemorySettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &MemorySettingData{wmivm}, nil
}

func (msd *MemorySettingData) SetSizeMB(size uint64) error {
	return msd.SetPropertyVirtualQuantity(size)
}

func (msd *MemorySettingData) GetSizeMB() (uint64, error) {
	virtualQuantity, err := msd.GetProperty("VirtualQuantity")
	if err != nil {
		return 0, err
	}
	// msd.GetPropertyVirtualQuantity was failing to cast the interface{} to uint64 directly
	virtualQuantityString, ok := virtualQuantity.(string)
	if !ok {
		return 0, err
	}
	return strconv.ParseUint(virtualQuantityString, 10, 64)
}

func (msd *MemorySettingData) GetMaximumMemoryMB() (uint64, error) {
	limit, err := msd.GetProperty("Limit")
	if err != nil {
		return 0, err
	}
	limitString, ok := limit.(string)
	if !ok {
		return 0, err
	}
	return strconv.ParseUint(limitString, 10, 64)
}

func (msd *MemorySettingData) GetMinimumMemoryMB() (uint64, error) {
	reservation, err := msd.GetProperty("Reservation")
	if err != nil {
		return 0, err
	}
	reservationString, ok := reservation.(string)
	if !ok {
		return 0, err
	}
	return strconv.ParseUint(reservationString, 10, 64)
}

func (msd *MemorySettingData) GetTargetMemoryBuffer() (uint32, error) {
	memoryBufferInterface, err := msd.GetProperty("TargetMemoryBuffer")
	if err != nil {
		return 0, err
	}
	memoryBuffer, ok := memoryBufferInterface.(int32)
	if !ok {
		return 0, err
	}
	return uint32(memoryBuffer), err
}

type DynamicMemoryConfiguration struct {
	DynamicMemoryEnabled bool
	MinimumMemoryMB      uint64
	MaximumMemoryMB      uint64
	TargetMemoryBuffer   uint32
}

func (msd *MemorySettingData) ConfigureDynamicMemory(config *DynamicMemoryConfiguration) error {
	if config == nil {
		return errors.Wrap(errors.InvalidInput, "Non-nil config must be provided")
	}

	err := msd.SetPropertyDynamicMemoryEnabled(config.DynamicMemoryEnabled)
	if err != nil {
		return err
	}

	memoryMB, err := msd.GetSizeMB()

	if config.DynamicMemoryEnabled {
		// first validate the fields
		// https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2012-R2-and-2012/hh831766(v=ws.11)
		if config.MinimumMemoryMB > memoryMB || config.MinimumMemoryMB < 32 {
			return errors.Wrapf(errors.InvalidInput, "Minimum memory must be between 32 MB and startup memory %d MB", memoryMB)
		}
		if config.MaximumMemoryMB < memoryMB {
			return errors.Wrapf(errors.InvalidInput, "Maximum memory must be greater than or equal to startup memory %d MB", memoryMB)
		}
		if config.TargetMemoryBuffer < 5 || config.TargetMemoryBuffer > 2000 {
			return errors.Wrap(errors.InvalidInput, "Memory buffer must be between 5 and 2000 percent")
		}

		// now set the new fields
		err = msd.SetPropertyLimit(config.MaximumMemoryMB)
		if err != nil {
			return err
		}
		err = msd.SetPropertyReservation(config.MinimumMemoryMB)
		if err != nil {
			return err
		}
		err = msd.SetPropertyTargetMemoryBuffer(config.TargetMemoryBuffer)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetDefaultMemorySettingData(whost *host.WmiHost) (*MemorySettingData, error) {
	inst, err := instance.CreateWmiInstance(whost, string(constant.Virtualization), "Msvm_MemorySettingData")
	if err != nil {
		return nil, err
	}
	return NewMemorySettingData(inst)
}
