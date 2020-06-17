// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"fmt"
	//"log"
	"time"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"

	"github.com/microsoft/wmi/pkg/virtualization/core/job"
	"github.com/microsoft/wmi/pkg/virtualization/core/resource"
	"github.com/microsoft/wmi/pkg/virtualization/core/resource/resourceallocation"
	"github.com/microsoft/wmi/pkg/virtualization/core/resource/resourcepool"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/controller"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/disk"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/drive"
	"github.com/microsoft/wmi/pkg/virtualization/network/switchport"
	na "github.com/microsoft/wmi/pkg/virtualization/network/virtualnetworkadapter"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualMachine struct {
	*v2.Msvm_ComputerSystem
}

type VirtualMachineState int32

// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/virtual/msvm-computersystem?redirectedfrom=MSDN
const (
	Unknown            VirtualMachineState = 0
	Other              VirtualMachineState = 1
	Running            VirtualMachineState = 2
	Off                VirtualMachineState = 3
	Stopping           VirtualMachineState = 4
	Saved              VirtualMachineState = 6
	Paused             VirtualMachineState = 9
	Starting           VirtualMachineState = 10
	Reset              VirtualMachineState = 11
	Saving             VirtualMachineState = 32773
	Pausing            VirtualMachineState = 32776
	Resuming           VirtualMachineState = 32777
	FastSaved          VirtualMachineState = 32779
	FastSaving         VirtualMachineState = 32780
	ForceShutdown      VirtualMachineState = 32781
	ForceReboot        VirtualMachineState = 32782
	Hibernated         VirtualMachineState = 32783
	ComponentServicing VirtualMachineState = 32784
	RunningCritical    VirtualMachineState = 32785
	OffCritical        VirtualMachineState = 32786
	StoppingCritial    VirtualMachineState = 32787
	SavedCritical      VirtualMachineState = 32788
	PausedCritical     VirtualMachineState = 32789
	StartingCritical   VirtualMachineState = 32790
	ResetCritical      VirtualMachineState = 32791
	SavingCritical     VirtualMachineState = 32792
	PausingCritical    VirtualMachineState = 32793
	ResumingCritical   VirtualMachineState = 32794
	FastSaveCritical   VirtualMachineState = 32795
	FastSavingCritical VirtualMachineState = 32796
)

// NewVirtualMachine
func NewVirtualMachine(instance *wmi.WmiInstance) (*VirtualMachine, error) {
	wmivm, err := v2.NewMsvm_ComputerSystemEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualMachine{wmivm}, nil
}

// GetVirtualMachine gets an existing virtual machine
func GetVirtualMachine(whost *host.WmiHost, vmName string) (vm *VirtualMachine, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_ComputerSystem", "ElementName", vmName)
	wmivm, err := v2.NewMsvm_ComputerSystemEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	vm = &VirtualMachine{wmivm}
	return
}

func (vm *VirtualMachine) Name() (name string) {
	name, _ = vm.GetPropertyElementName()
	return
}

// GetVirtualMachineState get the virtual machine state
func (vm *VirtualMachine) State() (state VirtualMachineState, err error) {
	err = vm.Refresh()
	if err != nil {
		return
	}

	retValue, err := vm.GetProperty("EnabledState")
	if err != nil {
		return
	}
	intstate, ok := retValue.(int32)
	if !ok {
		return Unknown, errors.Wrapf(errors.Failed, "Failed to get the state of the VM [%+v]", retValue)
	}
	return VirtualMachineState(intstate), nil
}

// Stop Virtual Machine
func (vm *VirtualMachine) Stop(force bool) error {
	return vm.ChangeState(Off, v2.ConcreteJob_JobType_Power_Off_Virtual_Machine)
}

// Start Virtual Machine
func (vm *VirtualMachine) Start() error {
	err := vm.ChangeState(Running, v2.ConcreteJob_JobType_Start_Virtual_Machine)
	if err != nil {
		return err
	}
	return vm.WaitForState(Running, 30)
}

// ChangeState changes the state of the Virtual Machine
func (vm *VirtualMachine) ChangeState(state VirtualMachineState, jobType v2.ConcreteJob_JobType) (err error) {
	cstate, err := vm.State()
	if err != nil {
		return err
	}
	// If the state is already satisfied, just return
	if cstate == state {
		return nil
	}
	result, err := vm.InvokeMethodWithReturn("RequestStateChange", int32(state))
	if err != nil {
		return err
	}
	return job.WaitForJobCompletion(vm.WmiInstance, result, jobType)
}

// WaitForState
func (vm *VirtualMachine) WaitForState(state VirtualMachineState, timeoutSeconds int32) (err error) {
	start := time.Now()
	// Run the loop, only if the job is actually running
	for {
		curState, err1 := vm.State()
		if err1 != nil {
			return err1
		}

		if curState == state {
			// Break for any valid state
			// TODO: WaitForSomeState
			return nil
		}
		time.Sleep(100 * time.Millisecond)

		// If we have waited enough time, break
		if time.Since(start) > (time.Duration(timeoutSeconds) * time.Second) {
			err = errors.Wrapf(errors.Timedout, "WaitForState timeout")
			break
		}
	}

	return
}

func (vm *VirtualMachine) GetVirtualSystemSettingData() (*VirtualSystemSettingData, error) {
	inst, err := vm.GetRelated("Msvm_VirtualSystemSettingData")
	if err != nil {
		return nil, err
	}
	return NewVirtualSystemSettingData(inst)
}

func (vm *VirtualMachine) GetVirtualNetworkAdapterByName(name string) (vna *na.VirtualNetworkAdapter, err error) {
	settings, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return
	}
	defer settings.Close()
	vna, err = settings.GetVirtualNetworkAdapterByName(name)
	return
}

func (vm *VirtualMachine) NewSyntheticDiskDrive(controllernumber, controllerlocation int32) (synDrive *drive.SyntheticDiskDrive, err error) {
	driverp, err := resourcepool.GetPrimordialResourcePool(vm.GetWmiHost(), v2.ResourcePool_ResourceType_Disk_Drive)
	if err != nil {
		return
	}
	defer driverp.Close()

	rasd, err := driverp.GetDefaultResourceAllocationSettingData()
	if err != nil {
		return
	}

	synDrive, err = drive.NewSyntheticDiskDrive(rasd.WmiInstance)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			synDrive.Close()
			synDrive = nil
		}
	}()

	// Only support SCSI
	controllers, err := vm.GetSCSIControllers()
	if err != nil {
		return
	}
	defer controllers.Close()
	// 1. Find the correct controller to use vased on the controllernumber
	if len(controllers) == 0 {
		err = errors.Wrapf(errors.NotFound, "VirtualMachine [%s] doesnt have SCSI Controller", vm.Name())
		return
	}
	if int(controllernumber) > len(controllers) {
		err = errors.Wrapf(errors.NotFound,
			"VirtualMachine [%s] doesnt have SCSI Controller with bus location [%d]", vm.Name(), controllernumber)
		return
	}

	if controllernumber == -1 {
		controllernumber = 0
	}
	scsicontroller, err := controller.NewSCSIControllerSettings(controllers[controllernumber].WmiInstance)
	if err != nil {
		return
	}

	synDrive.SetPropertyParent(scsicontroller.InstancePath())
	if controllerlocation == -1 {
		controllerlocation, err = scsicontroller.GetFreeLocation()
		if err != nil {
			err = errors.Wrapf(errors.NotFound, "Unable to find free location in SCSI Controller")
			return
		}
		// Find a free location
	}
	synDrive.SetPropertyAddressOnParent(fmt.Sprintf("%d", controllerlocation))

	return
}

func (vm *VirtualMachine) NewEthernetPortAllocationSettingData(vna *na.VirtualNetworkAdapter) (epas *switchport.EthernetPortAllocationSettingData, err error) {
	vhdrp, err := resourcepool.GetPrimordialResourcePool(vm.GetWmiHost(), v2.ResourcePool_ResourceType_Ethernet_Connection)
	if err != nil {
		return
	}
	defer vhdrp.Close()
	rasd, err := vhdrp.GetDefaultResourceAllocationSettingData()
	if err != nil {
		return
	}
	err = rasd.SetPropertyParent(vna.InstancePath())
	if err != nil {
		rasd.Close()
		return
	}

	return switchport.NewEthernetPortAllocationSettingData(rasd.WmiInstance)
}

func (vm *VirtualMachine) NewSCSIController() (resource *resourceallocation.ResourceAllocationSettingData, err error) {
	rp, err := resourcepool.GetPrimordialResourcePool(vm.GetWmiHost(), v2.ResourcePool_ResourceType_Parallel_SCSI_HBA)
	if err != nil {
		return
	}
	defer rp.Close()
	rasd, err := rp.GetDefaultResourceAllocationSettingData()
	if err != nil {
		return
	}

	resource, err = resourceallocation.NewResourceAllocationSettingData(rasd.WmiInstance)
	return
}

func (vm *VirtualMachine) NewTPM() (resource *resourceallocation.ResourceAllocationSettingData, err error) {
	rp, err := resourcepool.GetPrimordialResourcePool(vm.GetWmiHost(), v2.ResourcePool_ResourceType_Other)
	if err != nil {
		return
	}
	defer rp.Close()
	rasd, err := rp.GetDefaultResourceAllocationSettingData()
	if err != nil {
		return
	}

	resource, err = resourceallocation.NewResourceAllocationSettingData(rasd.WmiInstance)
	return
}

func (vm *VirtualMachine) NewSyntheticNetworkAdapter(name string) (adapter *na.VirtualNetworkAdapter, err error) {
	vhdrp, err := resourcepool.GetPrimordialResourcePool(vm.GetWmiHost(), v2.ResourcePool_ResourceType_Ethernet_Adapter)
	if err != nil {
		return
	}
	defer vhdrp.Close()
	rasd, err := vhdrp.GetDefaultResourceAllocationSettingData()
	if err != nil {
		return
	}

	adapter, err = na.NewVirtualNetworkAdapter(rasd.WmiInstance)
	if err != nil {
		rasd.Close()
		return
	}

	defer func() {
		if err != nil {
			adapter.Close()
			adapter = nil
		}
	}()

	syntheticNetworkAdapter, err := na.NewSyntheticNetworkAdapter(rasd.WmiInstance)
	if err != nil {
		return
	}

	err = syntheticNetworkAdapter.InitializeIdentifier()
	if err != nil {
		return
	}

	err = adapter.SetPropertyElementName(name)
	return
}

func (vm *VirtualMachine) NewVirtualHardDisk(path string) (vhd *disk.VirtualHardDisk, err error) {
	vhdrp, err := resourcepool.GetPrimordialResourcePool(vm.GetWmiHost(), v2.ResourcePool_ResourceType_Logical_Disk)
	if err != nil {
		return
	}
	defer vhdrp.Close()
	rasd, err := vhdrp.GetDefaultResourceAllocationSettingData()
	if err != nil {
		return
	}

	vhd, err = disk.NewVirtualHardDisk(rasd.WmiInstance)
	if err != nil {
		return
	}
	err = vhd.SetPropertyHostResource([]string{path})
	if err != nil {
		return
	}

	return
}

func (vm *VirtualMachine) GetSCSIControllers() (col resourceallocation.ResourceAllocationSettingDataCollection, err error) {
	col, err = vm.GetResourceAllocationSettingData(v2.ResourcePool_ResourceType_Parallel_SCSI_HBA)
	return
}

func (vm *VirtualMachine) GetVirtualHardDisks() (col resourceallocation.ResourceAllocationSettingDataCollection, err error) {
	col, err = vm.GetResourceAllocationSettingData(v2.ResourcePool_ResourceType_Logical_Disk)
	return
}

func (vm *VirtualMachine) GetVirtualHardDrives() (col resourceallocation.ResourceAllocationSettingDataCollection, err error) {
	col, err = vm.GetResourceAllocationSettingData(v2.ResourcePool_ResourceType_Disk_Drive)
	return
}

func (vm *VirtualMachine) GetVirtualHardDiskByLocation(controllerNumber, controllerLocation int) (vhd *disk.VirtualHardDisk, err error) {
	col, err := vm.GetVirtualHardDisks()
	if err != nil {
		return
	}
	defer col.Close()
	for _, inst := range col {
		tmpvhd, err1 := disk.NewVirtualHardDisk(inst.WmiInstance)
		if err1 != nil {
			err = err1
			return
		}
		tmpdrive, err1 := tmpvhd.GetDrive()
		if err1 != nil {
			err = err1
			return
		}
		defer tmpdrive.Close()
		vhddrive, err1 := drive.NewVirtualDrive(tmpdrive.WmiInstance)
		if err1 != nil {
			err = err1
			return
		}
		tmpcontrollerLocation, err1 := vhddrive.GetControllerLocation()
		if err1 != nil {
			err = err1
			return
		}
		tmpcontrollerNumber, err1 := vhddrive.GetControllerNumber()
		if err1 != nil {
			err = err1
			return
		}

		if tmpcontrollerNumber == fmt.Sprintf("%d", controllerNumber) &&
			tmpcontrollerLocation == fmt.Sprintf("%d", controllerLocation) {
			// Found the match
			vhdclone, err1 := tmpvhd.Clone()
			if err1 != nil {
				err = err1
				return
			}
			vhd, err = disk.NewVirtualHardDisk(vhdclone)
			return
		}
	}
	err = errors.Wrapf(errors.NotFound,
		"Vhd with controllerNumber[%d] controllerLocation[%d] not found in Vm [%s]",
		controllerNumber, controllerLocation, vm.Name())
	return

}
func (vm *VirtualMachine) GetVirtualHardDiskByPath(path string) (vhd *disk.VirtualHardDisk, err error) {
	col, err := vm.GetVirtualHardDisks()
	if err != nil {
		return
	}
	defer col.Close()
	for _, inst := range col {
		tmpvhd, err1 := disk.NewVirtualHardDisk(inst.WmiInstance)
		if err1 != nil {
			err = err1
			return
		}
		vhdpath, err1 := tmpvhd.GetPath()
		if err != nil {
			err = err1
			return
		}
		if vhdpath == path {
			vhdclone, err1 := tmpvhd.Clone()
			if err1 != nil {
				err = err1
				return
			}
			vhd, err = disk.NewVirtualHardDisk(vhdclone)
			return
		}
	}
	err = errors.Wrapf(errors.NotFound, "Vhd with path [%s] not found in Vm [%s]", path, vm.Name())
	return
}

func (vm *VirtualMachine) GetResourceAllocationSettingData(rtype v2.ResourcePool_ResourceType) (col resourceallocation.ResourceAllocationSettingDataCollection, err error) {
	settings, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return
	}
	defer settings.Close()

	rasdcol, err := settings.GetAllRelated("CIM_ResourceAllocationSettingData")
	if err != nil {
		return
	}
	defer rasdcol.Close()

	resType := resource.GetResourceTypeValue(rtype)
	for _, ins := range rasdcol {
		rasd, err1 := resourceallocation.NewResourceAllocationSettingData(ins)
		if err1 != nil {
			err = err1
			return
		}
		curresType, err1 := rasd.GetResourceType()
		if err1 != nil {
			err = err1
			return
		}

		if !curresType.Equals(resType) {
			continue
		}
		instance, err1 := rasd.Clone()
		if err1 != nil {
			err = err1
			return
		}
		rasdfound, err1 := resourceallocation.NewResourceAllocationSettingData(instance)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, rasdfound)
	}
	if len(col) == 0 {
		err = errors.Wrapf(errors.NotFound, "GetResourceAllocationSettingData [%s] ", resType)
	}
	return
}
