// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.SDDC.Management
//
// ////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SDDC_VirtualMachine struct
type SDDC_VirtualMachine struct {
	*cim.WmiInstance

	//
	Alerts []SDDC_Alert

	//
	BootOrder []uint16

	//
	ComputerName string

	//
	CpuUsage float64

	//
	CreationTime string

	//
	EncryptStateAndVmMigrationTraffic bool

	//
	Generation uint16

	//
	GuestStateDataRoot string

	//
	GuestStateFile string

	//
	GuestStateIsolationType uint16

	//
	HeartBeat uint16

	//
	Host string

	//
	Id string

	//
	IntegratedServiceVersion string

	//
	IsClustered bool

	//
	IsDynamicMemoryEnabled bool

	//
	IsGuestStateIsolationEnabled bool

	//
	LatestSnapshot string

	//
	MemoryAssigned uint64

	//
	MemoryDemand uint64

	//
	MemoryStartUp uint64

	//
	MemoryStartUpUnits string

	//
	Name string

	//
	Os string

	//
	OsVersion string

	//
	ParentCheckpointId string

	//
	ProcessorCount uint16

	//
	ReplicationHealth uint16

	//
	ReplicationMode uint16

	//
	ReplicationState uint16

	//
	SecureBootEnabled bool

	//
	ShieldingRequested bool

	//
	SizeOfSystemFiles uint64

	//
	SnapshotDataRoot string

	//
	Snapshots []SDDC_VmSnapshot

	//
	State uint16

	//
	Status []uint16

	//
	TotalIops float64

	//
	TotalNetworkUsage float64

	//
	TotalThroughput float64

	//
	TpmEnabled bool

	//
	Uptime string

	//
	UserSnapshotType uint16

	//
	Version string

	//
	Vhds []SDDC_Vhd

	//
	VirtualSystemType string

	//
	VmIntegrationServices []SDDC_VmIntegrationService

	//
	VNics []SDDC_VmNetAdapter
}

func NewSDDC_VirtualMachineEx1(instance *cim.WmiInstance) (newInstance *SDDC_VirtualMachine, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_VirtualMachine{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_VirtualMachineEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_VirtualMachine, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_VirtualMachine{
		WmiInstance: tmp,
	}
	return
}

// SetAlerts sets the value of Alerts for the instance
func (instance *SDDC_VirtualMachine) SetPropertyAlerts(value []SDDC_Alert) (err error) {
	return instance.SetProperty("Alerts", (value))
}

// GetAlerts gets the value of Alerts for the instance
func (instance *SDDC_VirtualMachine) GetPropertyAlerts() (value []SDDC_Alert, err error) {
	retValue, err := instance.GetProperty("Alerts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(SDDC_Alert)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " SDDC_Alert is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SDDC_Alert(valuetmp))
	}

	return
}

// SetBootOrder sets the value of BootOrder for the instance
func (instance *SDDC_VirtualMachine) SetPropertyBootOrder(value []uint16) (err error) {
	return instance.SetProperty("BootOrder", (value))
}

// GetBootOrder gets the value of BootOrder for the instance
func (instance *SDDC_VirtualMachine) GetPropertyBootOrder() (value []uint16, err error) {
	retValue, err := instance.GetProperty("BootOrder")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetComputerName sets the value of ComputerName for the instance
func (instance *SDDC_VirtualMachine) SetPropertyComputerName(value string) (err error) {
	return instance.SetProperty("ComputerName", (value))
}

// GetComputerName gets the value of ComputerName for the instance
func (instance *SDDC_VirtualMachine) GetPropertyComputerName() (value string, err error) {
	retValue, err := instance.GetProperty("ComputerName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetCpuUsage sets the value of CpuUsage for the instance
func (instance *SDDC_VirtualMachine) SetPropertyCpuUsage(value float64) (err error) {
	return instance.SetProperty("CpuUsage", (value))
}

// GetCpuUsage gets the value of CpuUsage for the instance
func (instance *SDDC_VirtualMachine) GetPropertyCpuUsage() (value float64, err error) {
	retValue, err := instance.GetProperty("CpuUsage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float64(valuetmp)

	return
}

// SetCreationTime sets the value of CreationTime for the instance
func (instance *SDDC_VirtualMachine) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", (value))
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *SDDC_VirtualMachine) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetEncryptStateAndVmMigrationTraffic sets the value of EncryptStateAndVmMigrationTraffic for the instance
func (instance *SDDC_VirtualMachine) SetPropertyEncryptStateAndVmMigrationTraffic(value bool) (err error) {
	return instance.SetProperty("EncryptStateAndVmMigrationTraffic", (value))
}

// GetEncryptStateAndVmMigrationTraffic gets the value of EncryptStateAndVmMigrationTraffic for the instance
func (instance *SDDC_VirtualMachine) GetPropertyEncryptStateAndVmMigrationTraffic() (value bool, err error) {
	retValue, err := instance.GetProperty("EncryptStateAndVmMigrationTraffic")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetGeneration sets the value of Generation for the instance
func (instance *SDDC_VirtualMachine) SetPropertyGeneration(value uint16) (err error) {
	return instance.SetProperty("Generation", (value))
}

// GetGeneration gets the value of Generation for the instance
func (instance *SDDC_VirtualMachine) GetPropertyGeneration() (value uint16, err error) {
	retValue, err := instance.GetProperty("Generation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetGuestStateDataRoot sets the value of GuestStateDataRoot for the instance
func (instance *SDDC_VirtualMachine) SetPropertyGuestStateDataRoot(value string) (err error) {
	return instance.SetProperty("GuestStateDataRoot", (value))
}

// GetGuestStateDataRoot gets the value of GuestStateDataRoot for the instance
func (instance *SDDC_VirtualMachine) GetPropertyGuestStateDataRoot() (value string, err error) {
	retValue, err := instance.GetProperty("GuestStateDataRoot")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetGuestStateFile sets the value of GuestStateFile for the instance
func (instance *SDDC_VirtualMachine) SetPropertyGuestStateFile(value string) (err error) {
	return instance.SetProperty("GuestStateFile", (value))
}

// GetGuestStateFile gets the value of GuestStateFile for the instance
func (instance *SDDC_VirtualMachine) GetPropertyGuestStateFile() (value string, err error) {
	retValue, err := instance.GetProperty("GuestStateFile")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetGuestStateIsolationType sets the value of GuestStateIsolationType for the instance
func (instance *SDDC_VirtualMachine) SetPropertyGuestStateIsolationType(value uint16) (err error) {
	return instance.SetProperty("GuestStateIsolationType", (value))
}

// GetGuestStateIsolationType gets the value of GuestStateIsolationType for the instance
func (instance *SDDC_VirtualMachine) GetPropertyGuestStateIsolationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("GuestStateIsolationType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetHeartBeat sets the value of HeartBeat for the instance
func (instance *SDDC_VirtualMachine) SetPropertyHeartBeat(value uint16) (err error) {
	return instance.SetProperty("HeartBeat", (value))
}

// GetHeartBeat gets the value of HeartBeat for the instance
func (instance *SDDC_VirtualMachine) GetPropertyHeartBeat() (value uint16, err error) {
	retValue, err := instance.GetProperty("HeartBeat")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetHost sets the value of Host for the instance
func (instance *SDDC_VirtualMachine) SetPropertyHost(value string) (err error) {
	return instance.SetProperty("Host", (value))
}

// GetHost gets the value of Host for the instance
func (instance *SDDC_VirtualMachine) GetPropertyHost() (value string, err error) {
	retValue, err := instance.GetProperty("Host")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetId sets the value of Id for the instance
func (instance *SDDC_VirtualMachine) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *SDDC_VirtualMachine) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetIntegratedServiceVersion sets the value of IntegratedServiceVersion for the instance
func (instance *SDDC_VirtualMachine) SetPropertyIntegratedServiceVersion(value string) (err error) {
	return instance.SetProperty("IntegratedServiceVersion", (value))
}

// GetIntegratedServiceVersion gets the value of IntegratedServiceVersion for the instance
func (instance *SDDC_VirtualMachine) GetPropertyIntegratedServiceVersion() (value string, err error) {
	retValue, err := instance.GetProperty("IntegratedServiceVersion")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetIsClustered sets the value of IsClustered for the instance
func (instance *SDDC_VirtualMachine) SetPropertyIsClustered(value bool) (err error) {
	return instance.SetProperty("IsClustered", (value))
}

// GetIsClustered gets the value of IsClustered for the instance
func (instance *SDDC_VirtualMachine) GetPropertyIsClustered() (value bool, err error) {
	retValue, err := instance.GetProperty("IsClustered")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetIsDynamicMemoryEnabled sets the value of IsDynamicMemoryEnabled for the instance
func (instance *SDDC_VirtualMachine) SetPropertyIsDynamicMemoryEnabled(value bool) (err error) {
	return instance.SetProperty("IsDynamicMemoryEnabled", (value))
}

// GetIsDynamicMemoryEnabled gets the value of IsDynamicMemoryEnabled for the instance
func (instance *SDDC_VirtualMachine) GetPropertyIsDynamicMemoryEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDynamicMemoryEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetIsGuestStateIsolationEnabled sets the value of IsGuestStateIsolationEnabled for the instance
func (instance *SDDC_VirtualMachine) SetPropertyIsGuestStateIsolationEnabled(value bool) (err error) {
	return instance.SetProperty("IsGuestStateIsolationEnabled", (value))
}

// GetIsGuestStateIsolationEnabled gets the value of IsGuestStateIsolationEnabled for the instance
func (instance *SDDC_VirtualMachine) GetPropertyIsGuestStateIsolationEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsGuestStateIsolationEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetLatestSnapshot sets the value of LatestSnapshot for the instance
func (instance *SDDC_VirtualMachine) SetPropertyLatestSnapshot(value string) (err error) {
	return instance.SetProperty("LatestSnapshot", (value))
}

// GetLatestSnapshot gets the value of LatestSnapshot for the instance
func (instance *SDDC_VirtualMachine) GetPropertyLatestSnapshot() (value string, err error) {
	retValue, err := instance.GetProperty("LatestSnapshot")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetMemoryAssigned sets the value of MemoryAssigned for the instance
func (instance *SDDC_VirtualMachine) SetPropertyMemoryAssigned(value uint64) (err error) {
	return instance.SetProperty("MemoryAssigned", (value))
}

// GetMemoryAssigned gets the value of MemoryAssigned for the instance
func (instance *SDDC_VirtualMachine) GetPropertyMemoryAssigned() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemoryAssigned")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMemoryDemand sets the value of MemoryDemand for the instance
func (instance *SDDC_VirtualMachine) SetPropertyMemoryDemand(value uint64) (err error) {
	return instance.SetProperty("MemoryDemand", (value))
}

// GetMemoryDemand gets the value of MemoryDemand for the instance
func (instance *SDDC_VirtualMachine) GetPropertyMemoryDemand() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemoryDemand")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMemoryStartUp sets the value of MemoryStartUp for the instance
func (instance *SDDC_VirtualMachine) SetPropertyMemoryStartUp(value uint64) (err error) {
	return instance.SetProperty("MemoryStartUp", (value))
}

// GetMemoryStartUp gets the value of MemoryStartUp for the instance
func (instance *SDDC_VirtualMachine) GetPropertyMemoryStartUp() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemoryStartUp")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMemoryStartUpUnits sets the value of MemoryStartUpUnits for the instance
func (instance *SDDC_VirtualMachine) SetPropertyMemoryStartUpUnits(value string) (err error) {
	return instance.SetProperty("MemoryStartUpUnits", (value))
}

// GetMemoryStartUpUnits gets the value of MemoryStartUpUnits for the instance
func (instance *SDDC_VirtualMachine) GetPropertyMemoryStartUpUnits() (value string, err error) {
	retValue, err := instance.GetProperty("MemoryStartUpUnits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *SDDC_VirtualMachine) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *SDDC_VirtualMachine) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetOs sets the value of Os for the instance
func (instance *SDDC_VirtualMachine) SetPropertyOs(value string) (err error) {
	return instance.SetProperty("Os", (value))
}

// GetOs gets the value of Os for the instance
func (instance *SDDC_VirtualMachine) GetPropertyOs() (value string, err error) {
	retValue, err := instance.GetProperty("Os")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetOsVersion sets the value of OsVersion for the instance
func (instance *SDDC_VirtualMachine) SetPropertyOsVersion(value string) (err error) {
	return instance.SetProperty("OsVersion", (value))
}

// GetOsVersion gets the value of OsVersion for the instance
func (instance *SDDC_VirtualMachine) GetPropertyOsVersion() (value string, err error) {
	retValue, err := instance.GetProperty("OsVersion")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetParentCheckpointId sets the value of ParentCheckpointId for the instance
func (instance *SDDC_VirtualMachine) SetPropertyParentCheckpointId(value string) (err error) {
	return instance.SetProperty("ParentCheckpointId", (value))
}

// GetParentCheckpointId gets the value of ParentCheckpointId for the instance
func (instance *SDDC_VirtualMachine) GetPropertyParentCheckpointId() (value string, err error) {
	retValue, err := instance.GetProperty("ParentCheckpointId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetProcessorCount sets the value of ProcessorCount for the instance
func (instance *SDDC_VirtualMachine) SetPropertyProcessorCount(value uint16) (err error) {
	return instance.SetProperty("ProcessorCount", (value))
}

// GetProcessorCount gets the value of ProcessorCount for the instance
func (instance *SDDC_VirtualMachine) GetPropertyProcessorCount() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProcessorCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetReplicationHealth sets the value of ReplicationHealth for the instance
func (instance *SDDC_VirtualMachine) SetPropertyReplicationHealth(value uint16) (err error) {
	return instance.SetProperty("ReplicationHealth", (value))
}

// GetReplicationHealth gets the value of ReplicationHealth for the instance
func (instance *SDDC_VirtualMachine) GetPropertyReplicationHealth() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationHealth")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetReplicationMode sets the value of ReplicationMode for the instance
func (instance *SDDC_VirtualMachine) SetPropertyReplicationMode(value uint16) (err error) {
	return instance.SetProperty("ReplicationMode", (value))
}

// GetReplicationMode gets the value of ReplicationMode for the instance
func (instance *SDDC_VirtualMachine) GetPropertyReplicationMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetReplicationState sets the value of ReplicationState for the instance
func (instance *SDDC_VirtualMachine) SetPropertyReplicationState(value uint16) (err error) {
	return instance.SetProperty("ReplicationState", (value))
}

// GetReplicationState gets the value of ReplicationState for the instance
func (instance *SDDC_VirtualMachine) GetPropertyReplicationState() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetSecureBootEnabled sets the value of SecureBootEnabled for the instance
func (instance *SDDC_VirtualMachine) SetPropertySecureBootEnabled(value bool) (err error) {
	return instance.SetProperty("SecureBootEnabled", (value))
}

// GetSecureBootEnabled gets the value of SecureBootEnabled for the instance
func (instance *SDDC_VirtualMachine) GetPropertySecureBootEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("SecureBootEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetShieldingRequested sets the value of ShieldingRequested for the instance
func (instance *SDDC_VirtualMachine) SetPropertyShieldingRequested(value bool) (err error) {
	return instance.SetProperty("ShieldingRequested", (value))
}

// GetShieldingRequested gets the value of ShieldingRequested for the instance
func (instance *SDDC_VirtualMachine) GetPropertyShieldingRequested() (value bool, err error) {
	retValue, err := instance.GetProperty("ShieldingRequested")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetSizeOfSystemFiles sets the value of SizeOfSystemFiles for the instance
func (instance *SDDC_VirtualMachine) SetPropertySizeOfSystemFiles(value uint64) (err error) {
	return instance.SetProperty("SizeOfSystemFiles", (value))
}

// GetSizeOfSystemFiles gets the value of SizeOfSystemFiles for the instance
func (instance *SDDC_VirtualMachine) GetPropertySizeOfSystemFiles() (value uint64, err error) {
	retValue, err := instance.GetProperty("SizeOfSystemFiles")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSnapshotDataRoot sets the value of SnapshotDataRoot for the instance
func (instance *SDDC_VirtualMachine) SetPropertySnapshotDataRoot(value string) (err error) {
	return instance.SetProperty("SnapshotDataRoot", (value))
}

// GetSnapshotDataRoot gets the value of SnapshotDataRoot for the instance
func (instance *SDDC_VirtualMachine) GetPropertySnapshotDataRoot() (value string, err error) {
	retValue, err := instance.GetProperty("SnapshotDataRoot")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSnapshots sets the value of Snapshots for the instance
func (instance *SDDC_VirtualMachine) SetPropertySnapshots(value []SDDC_VmSnapshot) (err error) {
	return instance.SetProperty("Snapshots", (value))
}

// GetSnapshots gets the value of Snapshots for the instance
func (instance *SDDC_VirtualMachine) GetPropertySnapshots() (value []SDDC_VmSnapshot, err error) {
	retValue, err := instance.GetProperty("Snapshots")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(SDDC_VmSnapshot)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " SDDC_VmSnapshot is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SDDC_VmSnapshot(valuetmp))
	}

	return
}

// SetState sets the value of State for the instance
func (instance *SDDC_VirtualMachine) SetPropertyState(value uint16) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *SDDC_VirtualMachine) GetPropertyState() (value uint16, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetStatus sets the value of Status for the instance
func (instance *SDDC_VirtualMachine) SetPropertyStatus(value []uint16) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *SDDC_VirtualMachine) GetPropertyStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetTotalIops sets the value of TotalIops for the instance
func (instance *SDDC_VirtualMachine) SetPropertyTotalIops(value float64) (err error) {
	return instance.SetProperty("TotalIops", (value))
}

// GetTotalIops gets the value of TotalIops for the instance
func (instance *SDDC_VirtualMachine) GetPropertyTotalIops() (value float64, err error) {
	retValue, err := instance.GetProperty("TotalIops")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float64(valuetmp)

	return
}

// SetTotalNetworkUsage sets the value of TotalNetworkUsage for the instance
func (instance *SDDC_VirtualMachine) SetPropertyTotalNetworkUsage(value float64) (err error) {
	return instance.SetProperty("TotalNetworkUsage", (value))
}

// GetTotalNetworkUsage gets the value of TotalNetworkUsage for the instance
func (instance *SDDC_VirtualMachine) GetPropertyTotalNetworkUsage() (value float64, err error) {
	retValue, err := instance.GetProperty("TotalNetworkUsage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float64(valuetmp)

	return
}

// SetTotalThroughput sets the value of TotalThroughput for the instance
func (instance *SDDC_VirtualMachine) SetPropertyTotalThroughput(value float64) (err error) {
	return instance.SetProperty("TotalThroughput", (value))
}

// GetTotalThroughput gets the value of TotalThroughput for the instance
func (instance *SDDC_VirtualMachine) GetPropertyTotalThroughput() (value float64, err error) {
	retValue, err := instance.GetProperty("TotalThroughput")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float64(valuetmp)

	return
}

// SetTpmEnabled sets the value of TpmEnabled for the instance
func (instance *SDDC_VirtualMachine) SetPropertyTpmEnabled(value bool) (err error) {
	return instance.SetProperty("TpmEnabled", (value))
}

// GetTpmEnabled gets the value of TpmEnabled for the instance
func (instance *SDDC_VirtualMachine) GetPropertyTpmEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("TpmEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUptime sets the value of Uptime for the instance
func (instance *SDDC_VirtualMachine) SetPropertyUptime(value string) (err error) {
	return instance.SetProperty("Uptime", (value))
}

// GetUptime gets the value of Uptime for the instance
func (instance *SDDC_VirtualMachine) GetPropertyUptime() (value string, err error) {
	retValue, err := instance.GetProperty("Uptime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetUserSnapshotType sets the value of UserSnapshotType for the instance
func (instance *SDDC_VirtualMachine) SetPropertyUserSnapshotType(value uint16) (err error) {
	return instance.SetProperty("UserSnapshotType", (value))
}

// GetUserSnapshotType gets the value of UserSnapshotType for the instance
func (instance *SDDC_VirtualMachine) GetPropertyUserSnapshotType() (value uint16, err error) {
	retValue, err := instance.GetProperty("UserSnapshotType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetVersion sets the value of Version for the instance
func (instance *SDDC_VirtualMachine) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *SDDC_VirtualMachine) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetVhds sets the value of Vhds for the instance
func (instance *SDDC_VirtualMachine) SetPropertyVhds(value []SDDC_Vhd) (err error) {
	return instance.SetProperty("Vhds", (value))
}

// GetVhds gets the value of Vhds for the instance
func (instance *SDDC_VirtualMachine) GetPropertyVhds() (value []SDDC_Vhd, err error) {
	retValue, err := instance.GetProperty("Vhds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(SDDC_Vhd)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " SDDC_Vhd is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SDDC_Vhd(valuetmp))
	}

	return
}

// SetVirtualSystemType sets the value of VirtualSystemType for the instance
func (instance *SDDC_VirtualMachine) SetPropertyVirtualSystemType(value string) (err error) {
	return instance.SetProperty("VirtualSystemType", (value))
}

// GetVirtualSystemType gets the value of VirtualSystemType for the instance
func (instance *SDDC_VirtualMachine) GetPropertyVirtualSystemType() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetVmIntegrationServices sets the value of VmIntegrationServices for the instance
func (instance *SDDC_VirtualMachine) SetPropertyVmIntegrationServices(value []SDDC_VmIntegrationService) (err error) {
	return instance.SetProperty("VmIntegrationServices", (value))
}

// GetVmIntegrationServices gets the value of VmIntegrationServices for the instance
func (instance *SDDC_VirtualMachine) GetPropertyVmIntegrationServices() (value []SDDC_VmIntegrationService, err error) {
	retValue, err := instance.GetProperty("VmIntegrationServices")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(SDDC_VmIntegrationService)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " SDDC_VmIntegrationService is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SDDC_VmIntegrationService(valuetmp))
	}

	return
}

// SetVNics sets the value of VNics for the instance
func (instance *SDDC_VirtualMachine) SetPropertyVNics(value []SDDC_VmNetAdapter) (err error) {
	return instance.SetProperty("VNics", (value))
}

// GetVNics gets the value of VNics for the instance
func (instance *SDDC_VirtualMachine) GetPropertyVNics() (value []SDDC_VmNetAdapter, err error) {
	retValue, err := instance.GetProperty("VNics")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(SDDC_VmNetAdapter)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " SDDC_VmNetAdapter is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SDDC_VmNetAdapter(valuetmp))
	}

	return
}

//

// <param name="SeriesName" type="string "></param>
// <param name="TimeFrame" type="uint16 "></param>

// <param name="Metric" type="SDDC_Metric "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_VirtualMachine) GetMetrics( /* IN */ SeriesName string,
	/* IN */ TimeFrame uint16,
	/* OUT */ Metric SDDC_Metric) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetMetrics", SeriesName, TimeFrame)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RefreshType" type="uint16 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_VirtualMachine) Refresh( /* IN */ RefreshType uint16) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Refresh", RefreshType)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
