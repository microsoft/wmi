// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package disk

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualHardDiskType uint16

const (
	VirtualHardDiskType_NONE   = 0
	VirtualHardDiskType_LEGACY = 1
	VirtualHardDiskType_FLAT   = 2
	VirtualHardDiskType_SPARSE = 3
)

type VirtualHardDiskFormat uint16

const (
	VirtualHardDiskFormat_NONE = 0
	VirtualHardDiskFormat_ISO  = 1
	VirtualHardDiskFormat_1    = 2
	VirtualHardDiskFormat_2    = 3
)

type VirtualHardDiskSettingData struct {
	*v2.Msvm_VirtualHardDiskSettingData
}

// NewVirtualHardDiskSettingData
func NewVirtualHardDiskSettingData(instance *wmi.WmiInstance) (*VirtualHardDiskSettingData, error) {
	wmivm, err := v2.NewMsvm_VirtualHardDiskSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualHardDiskSettingData{wmivm}, nil
}

func GetDefaultVirtualHardDiskSettingData(whost *host.WmiHost) (*VirtualHardDiskSettingData, error) {
	inst, err := instance.CreateWmiInstance(whost, string(constant.Virtualization), "Msvm_VirtualHardDiskSettingData")
	if err != nil {
		return nil, err
	}
	return NewVirtualHardDiskSettingData(inst)
}

func GetVirtualHardDiskSettingData(whost *host.WmiHost, path string,
	logicalSectorSize, physicalSectorSize, blockSize uint32,
	diskSize uint64, dynamic bool) (vhdsetting *VirtualHardDiskSettingData, err error) {

	vhdsetting, err = GetDefaultVirtualHardDiskSettingData(whost)
	if err != nil {
		return nil, err
	}

	vhdsetting.SetPropertyPath(path)
	vhdsetting.SetPropertyFormat(uint16(VirtualHardDiskFormat_2))
	vhdsetting.SetPropertyBlockSize(blockSize)
	vhdsetting.SetPropertyLogicalSectorSize(logicalSectorSize)
	vhdsetting.SetPropertyPhysicalSectorSize(physicalSectorSize)
	vhdsetting.SetPropertyMaxInternalSize(diskSize)
	if dynamic {
		vhdsetting.SetPropertyType(uint16(VirtualHardDiskType_SPARSE))
	} else {
		// Fixed
		vhdsetting.SetPropertyType(uint16(VirtualHardDiskType_FLAT))
	}

	return
}
