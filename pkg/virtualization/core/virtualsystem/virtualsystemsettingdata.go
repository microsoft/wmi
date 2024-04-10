// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"fmt"
	"net"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/gpu"
	"github.com/microsoft/wmi/pkg/virtualization/core/memory"
	"github.com/microsoft/wmi/pkg/virtualization/core/pcie"
	"github.com/microsoft/wmi/pkg/virtualization/core/processor"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/disk"
	na "github.com/microsoft/wmi/pkg/virtualization/network/virtualnetworkadapter"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualSystemSettingData struct {
	*v2.Msvm_VirtualSystemSettingData
}

// NewVirtualSystemSettingData
func NewVirtualSystemSettingData(instance *wmi.WmiInstance) (*VirtualSystemSettingData, error) {
	wmivm, err := v2.NewMsvm_VirtualSystemSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualSystemSettingData{wmivm}, nil
}

func GetDefaultVirtualSystemSettingData(whost *host.WmiHost) (*VirtualSystemSettingData, error) {
	inst, err := instance.CreateWmiInstance(whost, string(constant.Virtualization), "Msvm_VirtualSystemSettingData")
	if err != nil {
		return nil, err
	}
	return NewVirtualSystemSettingData(inst)
}

func GetVirtualSystemSettingData(whost *host.WmiHost, name string) (*VirtualSystemSettingData, error) {
	vmsettings, err := GetDefaultVirtualSystemSettingData(whost)
	if err != nil {
		return nil, err
	}
	vmsettings.SetPropertyElementName(name)
	// vmsettings.SetPropertyVirtualSystemSubType
	// vmsettings.SetProperty("VirtualSystemSubType", 1) // 2nd Generation

	return vmsettings, err
}

func (vm *VirtualSystemSettingData) GetPcieDevices() (col pcie.PcieDeviceCollection, err error) {
	rasdcollection, err := vm.GetAllRelated("Msvm_PciExpressSettingData")
	if err != nil {
		return nil, err
	}
	defer rasdcollection.Close()

	col, err = pcie.NewPcieDeviceCollection(rasdcollection)
	return
}

func (vm *VirtualSystemSettingData) GetPcieDevice(hostResource string) (*pcie.PcieDevice, error) {
	pcieDeviceCollection, err := vm.GetPcieDevices()
	if err != nil {
		return nil, err
	}
	defer pcieDeviceCollection.Close()

	for _, pcieDevice := range pcieDeviceCollection {
		path, err := pcieDevice.GetPath()
		if err != nil {
			return nil, err
		}
		if path == hostResource {
			// Found the match
			return pcieDevice.CloneEx1()
		}
	}
	err = errors.Wrapf(errors.NotFound, "Unable to find PCIe device with host resource [%s]", hostResource)
	return nil, err
}

func (vm *VirtualSystemSettingData) GetGpuPartitionSettingCollection() (col gpu.GpuPartitionSettingCollection, err error) {
	rasdcollection, err := vm.GetAllRelated("Msvm_GpuPartitionSettingData")
	if err != nil {
		return nil, err
	}
	defer rasdcollection.Close()

	col, err = gpu.NewGpuPartitionSettingCollection(rasdcollection)
	return
}

func (vm *VirtualSystemSettingData) GetGpuPartitionSettingData(partitionSizeBytes uint64) (*gpu.GpuPartitionSettingData, error) {
	gpuPartitionSettingCollection, err := vm.GetGpuPartitionSettingCollection()
	if err != nil {
		return nil, err
	}
	defer gpuPartitionSettingCollection.Close()

	for _, gpuPartitionSetting := range gpuPartitionSettingCollection {
		minVRAM, err := gpuPartitionSetting.GetMinPartitionVRAM()
		if err != nil {
			return nil, err
		}
		if minVRAM == partitionSizeBytes {
			// Found the match
			return gpuPartitionSetting.CloneEx1()
		}
	}
	err = errors.Wrapf(errors.NotFound, "Unable to find GPU partition with minimum partition VRAM of [%d bytes]", partitionSizeBytes)
	return nil, err
}

func (vm *VirtualSystemSettingData) GetSyntheticVirtualNetworkAdapters() (col na.VirtualNetworkAdapterCollection, err error) {
	psds, err := vm.GetAllRelated("Msvm_SyntheticEthernetPortSettingData")
	if err != nil {
		return nil, err
	}

	col, err = na.NewVirtualNetworkAdapterCollection(psds)
	if err != nil {
		psds.Close()
	}
	return
}

func (vm *VirtualSystemSettingData) GetEmulatedVirtualNetworkAdapters() (col na.VirtualNetworkAdapterCollection, err error) {
	psds, err := vm.GetAllRelated("Msvm_EmulatedEthernetPortSettingData")
	if err != nil {
		return nil, err
	}

	col, err = na.NewVirtualNetworkAdapterCollection(psds)
	if err != nil {
		psds.Close()
	}
	return
}

func (vm *VirtualSystemSettingData) GetVirtualNetworkAdapters() (col na.VirtualNetworkAdapterCollection, err error) {
	psds, err := vm.GetSyntheticVirtualNetworkAdapters()
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}
	defer func() {
		if err != nil {
			psds.Close()
		}
	}()
	psdse, err := vm.GetEmulatedVirtualNetworkAdapters()
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}
	defer func() {
		if err != nil {
			psdse.Close()
		}
	}()

	col = append(col, psds...)
	col = append(col, psdse...)
	return
}

// GetVirtualNetworkAdapterByMACAddress
func (vm *VirtualSystemSettingData) GetVirtualNetworkAdapterByMACAddress(macAddress string) (vna *na.VirtualNetworkAdapter, err error) {
	hw, err := net.ParseMAC(macAddress)
	if err != nil {
		err = errors.Wrapf(errors.InvalidInput, "MacAddress")
		return
	}
	nas, err := vm.GetVirtualNetworkAdapters()
	if err != nil {
		return
	}

	defer nas.Close()

	for _, networkAdapter := range nas {
		physicalAddress, err1 := networkAdapter.GetPropertyAddress()
		if err1 != nil {
			err = err1
			return
		}
		hwcurrent, err1 := net.ParseMAC(physicalAddress)
		if err1 != nil {
			err = err1
			return
		}
		if hw.String() == hwcurrent.String() {
			// Found the match
			// Assumption is only one adapter would match, as MAC is unique
			return networkAdapter.CloneEx1()
		}
	}
	err = errors.Wrapf(errors.NotFound, "Virtual Network Adapter with macaddress [%s]", macAddress)
	return
}

// GetVirtualNetworkAdapterByName
func (vm *VirtualSystemSettingData) GetVirtualNetworkAdapterByName(adapterName string) (vna *na.VirtualNetworkAdapter, err error) {
	nas, err := vm.GetVirtualNetworkAdapters()
	if err != nil {
		return
	}

	defer nas.Close()

	for _, networkAdapter := range nas {
		name, err1 := networkAdapter.GetPropertyElementName()
		if err1 != nil {
			err = err1
			return
		}
		if name == adapterName {
			// Found the match
			// Assumption is only one adapter would match, - FIXME - Duplicates
			return networkAdapter.CloneEx1()
		}
	}
	err = errors.Wrapf(errors.NotFound, "Virtual Network Adapter with name [%s]", adapterName)
	return
}

func (vm *VirtualSystemSettingData) GetVirtualHardDisks() (col disk.VirtualHardDiskCollection, err error) {
	rasdcollection, err := vm.getResourceAllocationSettingData(v2.ResourceAllocationSettingData_ResourceType_Disk_Drive)
	if err != nil {
		return nil, err
	}

	col, err = disk.NewVirtualHardDiskCollection(rasdcollection)
	if err != nil {
		rasdcollection.Close()
	}
	return
}

func (vm *VirtualSystemSettingData) getResourceAllocationSettingData(rtype v2.ResourceAllocationSettingData_ResourceType) (col wmi.WmiInstanceCollection, err error) {
	resourceType := fmt.Sprintf("%d", int32(rtype))
	query := query.NewWmiQuery("Cim_ResourceAllocationSettingData", "ResourceType", resourceType)
	col, err = vm.GetAllRelatedWithQuery(query)
	return
}

func (vm *VirtualSystemSettingData) GetMemorySetting() (out *memory.MemorySettingData, err error) {
	msd, err := vm.GetRelated("Msvm_MemorySettingData")
	if err != nil {
		return nil, err
	}

	return memory.NewMemorySettingData(msd)
}

func (vm *VirtualSystemSettingData) GetProcessorSetting() (out *processor.ProcessorSettingData, err error) {
	psd, err := vm.GetRelated("Msvm_ProcessorSettingData")
	if err != nil {
		return nil, err
	}

	return processor.NewProcessorSettingData(psd)
}

func (vm *VirtualSystemSettingData) SetHyperVGeneration(generation HyperVGeneration) (err error) {

	return vm.SetProperty("VirtualSystemSubType", string(generation))

}
