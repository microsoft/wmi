// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"encoding/xml"
	"fmt"
	"strings"

	//"log"
	"time"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"

	"reflect"

	job "github.com/microsoft/wmi/pkg/virtualization/core/job"
	"github.com/microsoft/wmi/pkg/virtualization/core/memory"
	"github.com/microsoft/wmi/pkg/virtualization/core/pcie"
	"github.com/microsoft/wmi/pkg/virtualization/core/processor"
	"github.com/microsoft/wmi/pkg/virtualization/core/resource"
	"github.com/microsoft/wmi/pkg/virtualization/core/resource/resourceallocation"
	"github.com/microsoft/wmi/pkg/virtualization/core/resource/resourcepool"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/controller"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/disk"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/drive"
	"github.com/microsoft/wmi/pkg/virtualization/network/switchport"
	na "github.com/microsoft/wmi/pkg/virtualization/network/virtualnetworkadapter"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualHardDiskType int32

const (
	VirtualHardDiskType_OS_VIRTUALHARDDISK       VirtualHardDiskType = 0
	VirtualHardDiskType_DATADISK_VIRTUALHARDDISK VirtualHardDiskType = 1
)

type VirtualMachine struct {
	*v2.Msvm_ComputerSystem
}

type MsvmSecuritySettingData struct {
	*v2.Msvm_SecuritySettingData
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

const (
	StateChangeTimeoutSeconds = 300
)

const (
	IDEController  = "IDE Controller"
	SCSIController = "SCSI Controller"
)

type HyperVGeneration string

const (
	HyperVGeneration_V1 = "Microsoft:Hyper-V:SubType:1"
	HyperVGeneration_V2 = "Microsoft:Hyper-V:SubType:2"
)

type INSTANCE struct {
	XMLName   xml.Name `xml:"INSTANCE"`
	Text      string   `xml:",chardata"`
	CLASSNAME string   `xml:"CLASSNAME,attr"`
	PROPERTY  []struct {
		Text  string `xml:",chardata"`
		NAME  string `xml:"NAME,attr"`
		TYPE  string `xml:"TYPE,attr"`
		VALUE string `xml:"VALUE"`
	} `xml:"PROPERTY"`
}

// NewVirtualMachine
func NewVirtualMachine(instance *wmi.WmiInstance) (*VirtualMachine, error) {
	wmivm, err := v2.NewMsvm_ComputerSystemEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualMachine{wmivm}, nil
}

// GetVirtualMachineByVMName gets an existing virtual machine
func GetVirtualMachineByVMName(whost *host.WmiHost, vmName string) (vm *VirtualMachine, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_ComputerSystem", "ElementName", vmName)
	wmivm, err := v2.NewMsvm_ComputerSystemEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	vm = &VirtualMachine{wmivm}
	return
}

// GetVirtualMachineByVMId gets an existing virtual machine
func GetVirtualMachineByVMId(whost *host.WmiHost, vmID string) (vm *VirtualMachine, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_ComputerSystem", "Name", vmID)
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

func (vm *VirtualMachine) ID() (name string) {
	name, _ = vm.GetPropertyName()
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

func (vm *VirtualMachine) StatusDescriptions() ([]string, error) {
	err := vm.Refresh()
	if err != nil {
		return nil, err
	}

	retValue, err := vm.GetProperty("StatusDescriptions")
	if err != nil {
		return nil, err
	}

	var statuses []string
	for _, status := range retValue.([]interface{}) {
		statuses = append(statuses, status.(string))
	}

	return statuses, nil
}

func (vm *VirtualMachine) Status() (string, error) {
	err := vm.Refresh()
	if err != nil {
		return "", err
	}

	retValue, err := vm.GetProperty("Status")
	if err != nil {
		return "", err
	}
	status, ok := retValue.(string)
	if !ok {
		return "", errors.Wrapf(errors.Failed, "Failed to get status for this VM [%+v]", retValue)
	}

	return status, nil
}

// Stop Virtual Machine
func (vm *VirtualMachine) Stop(force bool) error {
	if force {
		err := vm.ChangeState(Off, job.ConcreteJob_JobType_Power_Off_Virtual_Machine, -1)
		if err != nil {
			return err
		}
	} else {
		err := vm.ChangeState(Stopping, job.ConcreteJob_JobType_Shut_Down_Virtual_Machine, -1)
		if err != nil {
			return err
		}
	}
	return vm.WaitForState(Off, StateChangeTimeoutSeconds)
}

// Start Virtual Machine
func (vm *VirtualMachine) Start() error {
	err := vm.ChangeState(Running, job.ConcreteJob_JobType_Start_Virtual_Machine, -1)
	if err != nil {
		return err
	}
	return vm.WaitForState(Running, StateChangeTimeoutSeconds)
}

// Resume Virtual Machine
func (vm *VirtualMachine) Resume() error {
	err := vm.ChangeState(Running, job.ConcreteJob_JobType_Resume_Virtual_Machine, -1)
	if err != nil {
		return err
	}
	return vm.WaitForState(Running, StateChangeTimeoutSeconds)
}

// Restore Virtual Machine
func (vm *VirtualMachine) Restore() error {
	err := vm.ChangeState(Running, job.ConcreteJob_JobType_Restore_Virtual_Machine, -1)
	if err != nil {
		return err
	}
	return vm.WaitForState(Running, StateChangeTimeoutSeconds)
}

// Pause Virtual Machine
func (vm *VirtualMachine) Pause() error {
	err := vm.ChangeState(Paused, job.ConcreteJob_JobType_Pause_Virtual_Machine, -1)
	if err != nil {
		return err
	}
	return vm.WaitForState(Paused, StateChangeTimeoutSeconds)
}

// Save Virtual Machine
func (vm *VirtualMachine) Save() error {
	err := vm.ChangeState(Saved, job.ConcreteJob_JobType_Save_Virtual_Machine, -1)
	if err != nil {
		return err
	}
	return vm.WaitForState(Saved, StateChangeTimeoutSeconds)
}

// ChangeState changes the state of the Virtual Machine
func (vm *VirtualMachine) ChangeState(state VirtualMachineState, jobType job.ConcreteJob_JobType, timeoutSeconds int16) (err error) {
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
	return job.WaitForJobCompletion(vm.WmiInstance, result, jobType, timeoutSeconds)
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
			vmState, err2 := vm.State()
			if err2 != nil {
				vmState = Unknown
			}
			status, err2 := vm.Status()
			if err2 != nil {
				status = fmt.Sprintf("Unknown (error retreiving the status [%+v])", err2)
			}
			statusDescriptions, err2 := vm.StatusDescriptions()
			if err2 != nil {
				statusDescriptions = []string{fmt.Sprintf("Unknown (error retreiving the status descriptions [%+v])", err2)}
			}
			err = errors.Wrapf(errors.Timedout, "WaitForState timeout. Current state: [%v], status: [%v], status descriptions: [%v]", vmState, status, statusDescriptions)
			return err
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

func (vm *VirtualMachine) GetSecuritySettingData() (value *MsvmSecuritySettingData, err error) {
	inst, err := vm.GetRelated("Msvm_Tpm")

	// If the TPM is not found, then it is not configured or enabled
	if inst == nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	tpmwmi, err := v2.NewMsvm_TPMEx1(inst)
	if err != nil {
		return nil, err
	}

	cimSettings, err := tpmwmi.GetRelatedSecuritySettingData()
	if err != nil {
		return nil, err
	}

	securitySettings, err := v2.NewMsvm_SecuritySettingDataEx1(cimSettings)
	if err != nil {
		return nil, err
	}

	return &MsvmSecuritySettingData{securitySettings}, nil
}

func (vm *VirtualMachine) GetOSConfiguration() (computerName string, isWindows bool) {
	inst, err := vm.GetRelated("Msvm_KvpExchangeComponent")
	if err != nil {
		return
	}

	kvp, err := v2.NewMsvm_KvpExchangeComponentEx1(inst)
	if err != nil {
		return
	}

	guestIntrinsicItems, err := kvp.GetPropertyGuestIntrinsicExchangeItems()
	if err != nil {
		return
	}

	guestKvPairs := make(map[string]string)
	for _, item := range guestIntrinsicItems {
		var guestProperty INSTANCE
		err = xml.Unmarshal([]byte(item), &guestProperty)
		if err != nil {
			return
		}

		var key, value string
		for _, property := range guestProperty.PROPERTY {
			if property.NAME == "Name" {
				key = property.VALUE
			}

			if property.NAME == "Data" {
				value = property.VALUE
			}
		}

		if key != "" {
			guestKvPairs[key] = value
		}
	}

	computerName = guestKvPairs["FullyQualifiedDomainName"]
	osName := guestKvPairs["OSName"]
	if strings.Contains(strings.ToLower(osName), "window") {
		isWindows = true
	}

	return
}

func (vm *VirtualMachine) GetVirtualMachineGeneration() (HyperVGeneration, error) {

	settings, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return "", err
	}
	defer settings.Close()
	retValue, err := settings.GetProperty("VirtualSystemSubType")
	if err != nil {
		return "", err
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return "", err
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return "", err
	}

	value := HyperVGeneration(valuetmp)

	return value, nil
}

func (vm *VirtualMachine) GetVirtualNetworkAdapters() (col na.VirtualNetworkAdapterCollection, err error) {
	settings, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return
	}
	defer settings.Close()
	col, err = settings.GetVirtualNetworkAdapters()
	return
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

func (vm *VirtualMachine) NewSyntheticDiskDrive(controllernumber, controllerlocation int32, diskType VirtualHardDiskType) (synDrive *drive.SyntheticDiskDrive, err error) {
	driverp, err := resourcepool.GetPrimordialResourcePool(vm.GetWmiHost(), v2.ResourcePool_ResourceType_Disk_Drive)

	generation, err := vm.GetVirtualMachineGeneration()
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

	var controllers resourceallocation.ResourceAllocationSettingDataCollection
	var controllerType string
	if generation == HyperVGeneration_V1 && diskType == VirtualHardDiskType_OS_VIRTUALHARDDISK {
		controllers, err = vm.GetIDEControllers()
		controllerType = IDEController
		if err != nil {
			return
		}

	} else {
		controllers, err = vm.GetSCSIControllers()
		controllerType = SCSIController
		if err != nil {
			return
		}
	}

	defer controllers.Close()
	// 1. Find the correct controller to use vased on the controllernumber
	if len(controllers) == 0 {
		err = errors.Wrapf(errors.NotFound, "VirtualMachine [%s] doesnt have [%s]", vm.Name(), controllerType)
		return
	}
	if int(controllernumber) > len(controllers) {
		err = errors.Wrapf(errors.NotFound,
			"VirtualMachine [%s] doesnt have [%s] with bus location [%d]", vm.Name(), controllerType, controllernumber)
		return
	}

	if controllernumber == -1 {
		controllernumber = 0
	}

	if generation == HyperVGeneration_V1 && diskType == VirtualHardDiskType_OS_VIRTUALHARDDISK {
		idecontroller, err := controller.NewIDEControllerSettings(controllers[controllernumber].WmiInstance)
		if err != nil {
			return nil, err
		}

		synDrive.SetPropertyParent(idecontroller.InstancePath())
		if controllerlocation == -1 {
			controllerlocation, err = idecontroller.GetFreeLocation()
			if err != nil {
				err = errors.Wrapf(errors.NotFound, "Unable to find free location in IDE Controller")
				return nil, err
			}
			// Find a free location
		}
		synDrive.SetPropertyAddressOnParent(fmt.Sprintf("%d", controllerlocation))
	} else {
		scsicontroller, err := controller.NewSCSIControllerSettings(controllers[controllernumber].WmiInstance)
		if err != nil {
			return nil, err
		}

		synDrive.SetPropertyParent(scsicontroller.InstancePath())
		if controllerlocation == -1 {
			controllerlocation, err = scsicontroller.GetFreeLocation()
			if err != nil {
				err = errors.Wrapf(errors.NotFound, "Unable to find free location in SCSI Controller")
				return nil, err
			}
			// Find a free location
		}
		synDrive.SetPropertyAddressOnParent(fmt.Sprintf("%d", controllerlocation))

	}
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

func (vm *VirtualMachine) NewDvdDrive() (dvd *drive.DvdDrive, err error) {
	dvdrp, err := resourcepool.GetPrimordialResourcePool(vm.GetWmiHost(), v2.ResourcePool_ResourceType_DVD_drive)

	generation, err := vm.GetVirtualMachineGeneration()
	if err != nil {
		return
	}
	defer dvdrp.Close()
	rasd, err := dvdrp.GetDefaultResourceAllocationSettingData()
	if err != nil {
		return
	}

	dvd, err = drive.NewDvdDrive(rasd.WmiInstance)

	defer func() {
		if err != nil {
			dvd.Close()
			dvd = nil
		}
	}()

	if generation == HyperVGeneration_V1 {
		controllers, err := vm.GetIDEControllers()
		if err != nil {
			return nil, err
		}
		defer controllers.Close()
		// 1. Find the correct controller to use vased on the controllernumber
		if len(controllers) == 0 {
			err = errors.Wrapf(errors.NotFound, "VirtualMachine [%s] doesnt have IDE Controller", vm.Name())
			return nil, err
		}
		// for generation 1 VM's there are 2 IDE controllers and the dvd drive is always attached to the second IDE Controller
		if len(controllers) == 2 {
			idecontroller, err := controller.NewIDEControllerSettings(controllers[1].WmiInstance)
			if err != nil {
				return nil, err
			}

			dvd.SetPropertyParent(idecontroller.InstancePath())

			controllerlocation, err := idecontroller.GetFreeLocation()
			if err != nil {
				err = errors.Wrapf(errors.NotFound, "Unable to find free location in IDE Controller")
				return nil, err
			}

			dvd.SetPropertyAddressOnParent(fmt.Sprintf("%d", controllerlocation))
		} else {
			err = errors.Wrapf(errors.NotFound, "VirtualMachine [%s] doesnt have 2 IDE Controllers", vm.Name())
			return nil, err
		}
	}
	return
}

func (vm *VirtualMachine) NewLogicalDisk() (ld *disk.LogicalDisk, err error) {
	ldrp, err := resourcepool.GetPrimordialResourcePool(vm.GetWmiHost(), v2.ResourcePool_ResourceType_Logical_Disk)
	if err != nil {
		return
	}
	defer ldrp.Close()
	rasd, err := ldrp.GetDefaultResourceAllocationSettingData()
	if err != nil {
		return
	}

	ld, err = disk.NewLogicalDisk(rasd.WmiInstance)
	return
}

func (vm *VirtualMachine) NewPcieDevice(hostPcieDeviceWmiPath string) (newPcieDevice *pcie.PciExpressSettingData, err error) {
	whost := vm.GetWmiHost()

	newPcieDevice, err = pcie.GetDefaultPciExpressSettingData(whost)
	if err != nil {
		return
	}

	err = newPcieDevice.SetPropertyHostResource([]string{hostPcieDeviceWmiPath})
	return
}

func (vm *VirtualMachine) GetPcieDevice(hostResource string) (pcieDevice *pcie.PcieDevice, err error) {
	settings, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return
	}
	defer settings.Close()

	pcieDevice, err = settings.GetPcieDevice(hostResource)
	return
}

func (vm *VirtualMachine) GetSCSIControllers() (col resourceallocation.ResourceAllocationSettingDataCollection, err error) {
	col, err = vm.GetResourceAllocationSettingData(v2.ResourcePool_ResourceType_Parallel_SCSI_HBA)
	return
}

func (vm *VirtualMachine) GetIDEControllers() (col resourceallocation.ResourceAllocationSettingDataCollection, err error) {
	col, err = vm.GetResourceAllocationSettingData(v2.ResourcePool_ResourceType_IDE_Controller)
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

func (vm *VirtualMachine) GetAttachedVirtualHardDisks() (vhdPaths []string, err error) {
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

		vhdclone, err1 := tmpvhd.Clone()
		if err1 != nil {
			err = err1
			return
		}

		retVhd, err1 := disk.NewVirtualHardDisk(vhdclone)
		if err1 != nil {
			err = err1
			return
		}

		vhdpath, err1 := retVhd.GetPropertyHostResource()
		if err1 != nil || len(vhdpath) == 0 {
			err = fmt.Errorf("Unable to read HostResource field from disk WMI %s", err1)
			return
		}
		vhdPaths = append(vhdPaths, vhdpath[0])
	}

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

func (vm *VirtualMachine) GetMemory() (vmmemory *memory.MemorySettingData, err error) {
	settings, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return
	}
	defer settings.Close()
	vmmemory, err = settings.GetMemorySetting()
	return
}

func (vm *VirtualMachine) GetProcessor() (vmprocessor *processor.ProcessorSettingData, err error) {
	settings, err := vm.GetVirtualSystemSettingData()
	if err != nil {
		return
	}
	defer settings.Close()
	vmprocessor, err = settings.GetProcessorSetting()
	return
}
