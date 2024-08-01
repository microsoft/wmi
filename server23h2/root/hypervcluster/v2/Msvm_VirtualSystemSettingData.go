// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_VirtualSystemSettingData struct
type Msvm_VirtualSystemSettingData struct {
	*CIM_VirtualSystemSettingData

	//
	AdditionalRecoveryInformation string

	//
	AllowFullSCSICommandSet bool

	//
	AllowReducedFcRedundancy bool

	//
	Architecture VirtualSystemSettingData_Architecture

	//
	AutomaticCriticalErrorAction VirtualSystemSettingData_AutomaticCriticalErrorAction

	//
	AutomaticCriticalErrorActionTimeout string

	//
	AutomaticSnapshotsEnabled bool

	//
	BaseBoardSerialNumber string

	//
	BIOSGUID string

	//
	BIOSNumLock bool

	//
	BIOSSerialNumber string

	//
	BootOrder []uint16

	//
	BootPciExpress bool

	//
	BootPciExpressInstanceFilter string

	//
	BootSourceOrder []string

	//
	ChassisAssetTag string

	//
	ChassisSerialNumber string

	//
	ClusterWideNodeCapabilitiesValidationMode VirtualSystemSettingData_ClusterWideNodeCapabilitiesValidationMode

	//
	ConsoleMode VirtualSystemSettingData_ConsoleMode

	//
	DebugChannelId uint32

	//
	DebugPort uint32

	//
	DebugPortEnabled VirtualSystemSettingData_DebugPortEnabled

	//
	EnableHibernation bool

	//
	EnhancedSessionTransportType uint16

	//
	FirmwareFile string

	//
	FirmwareParameters []uint8

	//
	GuestControlledCacheTypes bool

	//
	GuestFeatureSet uint64

	// Filepath of a directory where information about the guest runtime state is stored.
	GuestStateDataRoot string

	// Filepath of a file where information about the guest runtime state is stored. A relative path appends to the value of the GuestStateDataRoot property.
	GuestStateFile string

	//
	GuestStateIsolationEnabled bool

	//
	GuestStateIsolationMode uint16

	//
	GuestStateIsolationType uint16

	//
	HighMmioGapBase uint64

	//
	HighMmioGapSize uint64

	//
	IncrementalBackupEnabled bool

	//
	IsAutomaticSnapshot bool

	//
	IsSaved bool

	//
	LockOnDisconnect bool

	//
	LowMmioGapSize uint64

	//
	MemoryHostingJobObjectName string

	//
	NetworkBootPreferredProtocol VirtualSystemSettingData_NetworkBootPreferredProtocol

	//
	NumaNodeTopologyArray []string

	//
	Parent string

	//
	PauseAfterBootFailure bool

	//
	SecureBootEnabled bool

	//
	SecureBootTemplateId string

	//
	TurnOffOnGuestRestart bool

	//
	UserSnapshotType uint16

	//
	Version string

	//
	VirtualNumaEnabled bool

	//
	VirtualSlitType VirtualSystemSettingData_VirtualSlitType

	//
	VirtualSystemSubType VirtualSystemSettingData_VirtualSystemSubType

	//
	VMBusMessageRedirection uint16

	//
	Vtl2AddressRangeBase uint64

	//
	Vtl2AddressRangeSize uint64

	//
	Vtl2AddressSpaceConfigurationMode VirtualSystemSettingData_Vtl2AddressSpaceConfigurationMode

	//
	Vtl2MmioAddressRangeSize uint64

	//
	WatchdogEnabled bool

	//
	WorkerJobObjectName string
}

func NewMsvm_VirtualSystemSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemSettingData, err error) {
	tmp, err := NewCIM_VirtualSystemSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemSettingData{
		CIM_VirtualSystemSettingData: tmp,
	}
	return
}

func NewMsvm_VirtualSystemSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemSettingData, err error) {
	tmp, err := NewCIM_VirtualSystemSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemSettingData{
		CIM_VirtualSystemSettingData: tmp,
	}
	return
}

// SetAdditionalRecoveryInformation sets the value of AdditionalRecoveryInformation for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyAdditionalRecoveryInformation(value string) (err error) {
	return instance.SetProperty("AdditionalRecoveryInformation", (value))
}

// GetAdditionalRecoveryInformation gets the value of AdditionalRecoveryInformation for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyAdditionalRecoveryInformation() (value string, err error) {
	retValue, err := instance.GetProperty("AdditionalRecoveryInformation")
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

// SetAllowFullSCSICommandSet sets the value of AllowFullSCSICommandSet for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyAllowFullSCSICommandSet(value bool) (err error) {
	return instance.SetProperty("AllowFullSCSICommandSet", (value))
}

// GetAllowFullSCSICommandSet gets the value of AllowFullSCSICommandSet for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyAllowFullSCSICommandSet() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowFullSCSICommandSet")
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

// SetAllowReducedFcRedundancy sets the value of AllowReducedFcRedundancy for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyAllowReducedFcRedundancy(value bool) (err error) {
	return instance.SetProperty("AllowReducedFcRedundancy", (value))
}

// GetAllowReducedFcRedundancy gets the value of AllowReducedFcRedundancy for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyAllowReducedFcRedundancy() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowReducedFcRedundancy")
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

// SetArchitecture sets the value of Architecture for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyArchitecture(value VirtualSystemSettingData_Architecture) (err error) {
	return instance.SetProperty("Architecture", (value))
}

// GetArchitecture gets the value of Architecture for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyArchitecture() (value VirtualSystemSettingData_Architecture, err error) {
	retValue, err := instance.GetProperty("Architecture")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualSystemSettingData_Architecture(valuetmp)

	return
}

// SetAutomaticCriticalErrorAction sets the value of AutomaticCriticalErrorAction for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyAutomaticCriticalErrorAction(value VirtualSystemSettingData_AutomaticCriticalErrorAction) (err error) {
	return instance.SetProperty("AutomaticCriticalErrorAction", (value))
}

// GetAutomaticCriticalErrorAction gets the value of AutomaticCriticalErrorAction for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyAutomaticCriticalErrorAction() (value VirtualSystemSettingData_AutomaticCriticalErrorAction, err error) {
	retValue, err := instance.GetProperty("AutomaticCriticalErrorAction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualSystemSettingData_AutomaticCriticalErrorAction(valuetmp)

	return
}

// SetAutomaticCriticalErrorActionTimeout sets the value of AutomaticCriticalErrorActionTimeout for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyAutomaticCriticalErrorActionTimeout(value string) (err error) {
	return instance.SetProperty("AutomaticCriticalErrorActionTimeout", (value))
}

// GetAutomaticCriticalErrorActionTimeout gets the value of AutomaticCriticalErrorActionTimeout for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyAutomaticCriticalErrorActionTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("AutomaticCriticalErrorActionTimeout")
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

// SetAutomaticSnapshotsEnabled sets the value of AutomaticSnapshotsEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyAutomaticSnapshotsEnabled(value bool) (err error) {
	return instance.SetProperty("AutomaticSnapshotsEnabled", (value))
}

// GetAutomaticSnapshotsEnabled gets the value of AutomaticSnapshotsEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyAutomaticSnapshotsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AutomaticSnapshotsEnabled")
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

// SetBaseBoardSerialNumber sets the value of BaseBoardSerialNumber for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyBaseBoardSerialNumber(value string) (err error) {
	return instance.SetProperty("BaseBoardSerialNumber", (value))
}

// GetBaseBoardSerialNumber gets the value of BaseBoardSerialNumber for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyBaseBoardSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("BaseBoardSerialNumber")
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

// SetBIOSGUID sets the value of BIOSGUID for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyBIOSGUID(value string) (err error) {
	return instance.SetProperty("BIOSGUID", (value))
}

// GetBIOSGUID gets the value of BIOSGUID for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyBIOSGUID() (value string, err error) {
	retValue, err := instance.GetProperty("BIOSGUID")
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

// SetBIOSNumLock sets the value of BIOSNumLock for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyBIOSNumLock(value bool) (err error) {
	return instance.SetProperty("BIOSNumLock", (value))
}

// GetBIOSNumLock gets the value of BIOSNumLock for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyBIOSNumLock() (value bool, err error) {
	retValue, err := instance.GetProperty("BIOSNumLock")
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

// SetBIOSSerialNumber sets the value of BIOSSerialNumber for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyBIOSSerialNumber(value string) (err error) {
	return instance.SetProperty("BIOSSerialNumber", (value))
}

// GetBIOSSerialNumber gets the value of BIOSSerialNumber for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyBIOSSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("BIOSSerialNumber")
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

// SetBootOrder sets the value of BootOrder for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyBootOrder(value []uint16) (err error) {
	return instance.SetProperty("BootOrder", (value))
}

// GetBootOrder gets the value of BootOrder for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyBootOrder() (value []uint16, err error) {
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

// SetBootPciExpress sets the value of BootPciExpress for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyBootPciExpress(value bool) (err error) {
	return instance.SetProperty("BootPciExpress", (value))
}

// GetBootPciExpress gets the value of BootPciExpress for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyBootPciExpress() (value bool, err error) {
	retValue, err := instance.GetProperty("BootPciExpress")
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

// SetBootPciExpressInstanceFilter sets the value of BootPciExpressInstanceFilter for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyBootPciExpressInstanceFilter(value string) (err error) {
	return instance.SetProperty("BootPciExpressInstanceFilter", (value))
}

// GetBootPciExpressInstanceFilter gets the value of BootPciExpressInstanceFilter for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyBootPciExpressInstanceFilter() (value string, err error) {
	retValue, err := instance.GetProperty("BootPciExpressInstanceFilter")
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

// SetBootSourceOrder sets the value of BootSourceOrder for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyBootSourceOrder(value []string) (err error) {
	return instance.SetProperty("BootSourceOrder", (value))
}

// GetBootSourceOrder gets the value of BootSourceOrder for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyBootSourceOrder() (value []string, err error) {
	retValue, err := instance.GetProperty("BootSourceOrder")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetChassisAssetTag sets the value of ChassisAssetTag for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyChassisAssetTag(value string) (err error) {
	return instance.SetProperty("ChassisAssetTag", (value))
}

// GetChassisAssetTag gets the value of ChassisAssetTag for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyChassisAssetTag() (value string, err error) {
	retValue, err := instance.GetProperty("ChassisAssetTag")
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

// SetChassisSerialNumber sets the value of ChassisSerialNumber for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyChassisSerialNumber(value string) (err error) {
	return instance.SetProperty("ChassisSerialNumber", (value))
}

// GetChassisSerialNumber gets the value of ChassisSerialNumber for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyChassisSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("ChassisSerialNumber")
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

// SetClusterWideNodeCapabilitiesValidationMode sets the value of ClusterWideNodeCapabilitiesValidationMode for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyClusterWideNodeCapabilitiesValidationMode(value VirtualSystemSettingData_ClusterWideNodeCapabilitiesValidationMode) (err error) {
	return instance.SetProperty("ClusterWideNodeCapabilitiesValidationMode", (value))
}

// GetClusterWideNodeCapabilitiesValidationMode gets the value of ClusterWideNodeCapabilitiesValidationMode for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyClusterWideNodeCapabilitiesValidationMode() (value VirtualSystemSettingData_ClusterWideNodeCapabilitiesValidationMode, err error) {
	retValue, err := instance.GetProperty("ClusterWideNodeCapabilitiesValidationMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualSystemSettingData_ClusterWideNodeCapabilitiesValidationMode(valuetmp)

	return
}

// SetConsoleMode sets the value of ConsoleMode for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyConsoleMode(value VirtualSystemSettingData_ConsoleMode) (err error) {
	return instance.SetProperty("ConsoleMode", (value))
}

// GetConsoleMode gets the value of ConsoleMode for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyConsoleMode() (value VirtualSystemSettingData_ConsoleMode, err error) {
	retValue, err := instance.GetProperty("ConsoleMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualSystemSettingData_ConsoleMode(valuetmp)

	return
}

// SetDebugChannelId sets the value of DebugChannelId for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyDebugChannelId(value uint32) (err error) {
	return instance.SetProperty("DebugChannelId", (value))
}

// GetDebugChannelId gets the value of DebugChannelId for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyDebugChannelId() (value uint32, err error) {
	retValue, err := instance.GetProperty("DebugChannelId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetDebugPort sets the value of DebugPort for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyDebugPort(value uint32) (err error) {
	return instance.SetProperty("DebugPort", (value))
}

// GetDebugPort gets the value of DebugPort for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyDebugPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("DebugPort")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetDebugPortEnabled sets the value of DebugPortEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyDebugPortEnabled(value VirtualSystemSettingData_DebugPortEnabled) (err error) {
	return instance.SetProperty("DebugPortEnabled", (value))
}

// GetDebugPortEnabled gets the value of DebugPortEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyDebugPortEnabled() (value VirtualSystemSettingData_DebugPortEnabled, err error) {
	retValue, err := instance.GetProperty("DebugPortEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualSystemSettingData_DebugPortEnabled(valuetmp)

	return
}

// SetEnableHibernation sets the value of EnableHibernation for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyEnableHibernation(value bool) (err error) {
	return instance.SetProperty("EnableHibernation", (value))
}

// GetEnableHibernation gets the value of EnableHibernation for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyEnableHibernation() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableHibernation")
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

// SetEnhancedSessionTransportType sets the value of EnhancedSessionTransportType for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyEnhancedSessionTransportType(value uint16) (err error) {
	return instance.SetProperty("EnhancedSessionTransportType", (value))
}

// GetEnhancedSessionTransportType gets the value of EnhancedSessionTransportType for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyEnhancedSessionTransportType() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnhancedSessionTransportType")
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

// SetFirmwareFile sets the value of FirmwareFile for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyFirmwareFile(value string) (err error) {
	return instance.SetProperty("FirmwareFile", (value))
}

// GetFirmwareFile gets the value of FirmwareFile for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyFirmwareFile() (value string, err error) {
	retValue, err := instance.GetProperty("FirmwareFile")
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

// SetFirmwareParameters sets the value of FirmwareParameters for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyFirmwareParameters(value []uint8) (err error) {
	return instance.SetProperty("FirmwareParameters", (value))
}

// GetFirmwareParameters gets the value of FirmwareParameters for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyFirmwareParameters() (value []uint8, err error) {
	retValue, err := instance.GetProperty("FirmwareParameters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetGuestControlledCacheTypes sets the value of GuestControlledCacheTypes for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyGuestControlledCacheTypes(value bool) (err error) {
	return instance.SetProperty("GuestControlledCacheTypes", (value))
}

// GetGuestControlledCacheTypes gets the value of GuestControlledCacheTypes for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyGuestControlledCacheTypes() (value bool, err error) {
	retValue, err := instance.GetProperty("GuestControlledCacheTypes")
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

// SetGuestFeatureSet sets the value of GuestFeatureSet for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyGuestFeatureSet(value uint64) (err error) {
	return instance.SetProperty("GuestFeatureSet", (value))
}

// GetGuestFeatureSet gets the value of GuestFeatureSet for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyGuestFeatureSet() (value uint64, err error) {
	retValue, err := instance.GetProperty("GuestFeatureSet")
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

// SetGuestStateDataRoot sets the value of GuestStateDataRoot for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyGuestStateDataRoot(value string) (err error) {
	return instance.SetProperty("GuestStateDataRoot", (value))
}

// GetGuestStateDataRoot gets the value of GuestStateDataRoot for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyGuestStateDataRoot() (value string, err error) {
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
func (instance *Msvm_VirtualSystemSettingData) SetPropertyGuestStateFile(value string) (err error) {
	return instance.SetProperty("GuestStateFile", (value))
}

// GetGuestStateFile gets the value of GuestStateFile for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyGuestStateFile() (value string, err error) {
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

// SetGuestStateIsolationEnabled sets the value of GuestStateIsolationEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyGuestStateIsolationEnabled(value bool) (err error) {
	return instance.SetProperty("GuestStateIsolationEnabled", (value))
}

// GetGuestStateIsolationEnabled gets the value of GuestStateIsolationEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyGuestStateIsolationEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("GuestStateIsolationEnabled")
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

// SetGuestStateIsolationMode sets the value of GuestStateIsolationMode for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyGuestStateIsolationMode(value uint16) (err error) {
	return instance.SetProperty("GuestStateIsolationMode", (value))
}

// GetGuestStateIsolationMode gets the value of GuestStateIsolationMode for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyGuestStateIsolationMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("GuestStateIsolationMode")
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

// SetGuestStateIsolationType sets the value of GuestStateIsolationType for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyGuestStateIsolationType(value uint16) (err error) {
	return instance.SetProperty("GuestStateIsolationType", (value))
}

// GetGuestStateIsolationType gets the value of GuestStateIsolationType for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyGuestStateIsolationType() (value uint16, err error) {
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

// SetHighMmioGapBase sets the value of HighMmioGapBase for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyHighMmioGapBase(value uint64) (err error) {
	return instance.SetProperty("HighMmioGapBase", (value))
}

// GetHighMmioGapBase gets the value of HighMmioGapBase for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyHighMmioGapBase() (value uint64, err error) {
	retValue, err := instance.GetProperty("HighMmioGapBase")
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

// SetHighMmioGapSize sets the value of HighMmioGapSize for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyHighMmioGapSize(value uint64) (err error) {
	return instance.SetProperty("HighMmioGapSize", (value))
}

// GetHighMmioGapSize gets the value of HighMmioGapSize for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyHighMmioGapSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("HighMmioGapSize")
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

// SetIncrementalBackupEnabled sets the value of IncrementalBackupEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyIncrementalBackupEnabled(value bool) (err error) {
	return instance.SetProperty("IncrementalBackupEnabled", (value))
}

// GetIncrementalBackupEnabled gets the value of IncrementalBackupEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyIncrementalBackupEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IncrementalBackupEnabled")
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

// SetIsAutomaticSnapshot sets the value of IsAutomaticSnapshot for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyIsAutomaticSnapshot(value bool) (err error) {
	return instance.SetProperty("IsAutomaticSnapshot", (value))
}

// GetIsAutomaticSnapshot gets the value of IsAutomaticSnapshot for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyIsAutomaticSnapshot() (value bool, err error) {
	retValue, err := instance.GetProperty("IsAutomaticSnapshot")
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

// SetIsSaved sets the value of IsSaved for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyIsSaved(value bool) (err error) {
	return instance.SetProperty("IsSaved", (value))
}

// GetIsSaved gets the value of IsSaved for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyIsSaved() (value bool, err error) {
	retValue, err := instance.GetProperty("IsSaved")
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

// SetLockOnDisconnect sets the value of LockOnDisconnect for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyLockOnDisconnect(value bool) (err error) {
	return instance.SetProperty("LockOnDisconnect", (value))
}

// GetLockOnDisconnect gets the value of LockOnDisconnect for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyLockOnDisconnect() (value bool, err error) {
	retValue, err := instance.GetProperty("LockOnDisconnect")
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

// SetLowMmioGapSize sets the value of LowMmioGapSize for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyLowMmioGapSize(value uint64) (err error) {
	return instance.SetProperty("LowMmioGapSize", (value))
}

// GetLowMmioGapSize gets the value of LowMmioGapSize for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyLowMmioGapSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("LowMmioGapSize")
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

// SetMemoryHostingJobObjectName sets the value of MemoryHostingJobObjectName for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyMemoryHostingJobObjectName(value string) (err error) {
	return instance.SetProperty("MemoryHostingJobObjectName", (value))
}

// GetMemoryHostingJobObjectName gets the value of MemoryHostingJobObjectName for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyMemoryHostingJobObjectName() (value string, err error) {
	retValue, err := instance.GetProperty("MemoryHostingJobObjectName")
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

// SetNetworkBootPreferredProtocol sets the value of NetworkBootPreferredProtocol for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyNetworkBootPreferredProtocol(value VirtualSystemSettingData_NetworkBootPreferredProtocol) (err error) {
	return instance.SetProperty("NetworkBootPreferredProtocol", (value))
}

// GetNetworkBootPreferredProtocol gets the value of NetworkBootPreferredProtocol for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyNetworkBootPreferredProtocol() (value VirtualSystemSettingData_NetworkBootPreferredProtocol, err error) {
	retValue, err := instance.GetProperty("NetworkBootPreferredProtocol")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualSystemSettingData_NetworkBootPreferredProtocol(valuetmp)

	return
}

// SetNumaNodeTopologyArray sets the value of NumaNodeTopologyArray for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyNumaNodeTopologyArray(value []string) (err error) {
	return instance.SetProperty("NumaNodeTopologyArray", (value))
}

// GetNumaNodeTopologyArray gets the value of NumaNodeTopologyArray for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyNumaNodeTopologyArray() (value []string, err error) {
	retValue, err := instance.GetProperty("NumaNodeTopologyArray")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetParent sets the value of Parent for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyParent(value string) (err error) {
	return instance.SetProperty("Parent", (value))
}

// GetParent gets the value of Parent for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyParent() (value string, err error) {
	retValue, err := instance.GetProperty("Parent")
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

// SetPauseAfterBootFailure sets the value of PauseAfterBootFailure for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyPauseAfterBootFailure(value bool) (err error) {
	return instance.SetProperty("PauseAfterBootFailure", (value))
}

// GetPauseAfterBootFailure gets the value of PauseAfterBootFailure for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyPauseAfterBootFailure() (value bool, err error) {
	retValue, err := instance.GetProperty("PauseAfterBootFailure")
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

// SetSecureBootEnabled sets the value of SecureBootEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertySecureBootEnabled(value bool) (err error) {
	return instance.SetProperty("SecureBootEnabled", (value))
}

// GetSecureBootEnabled gets the value of SecureBootEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertySecureBootEnabled() (value bool, err error) {
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

// SetSecureBootTemplateId sets the value of SecureBootTemplateId for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertySecureBootTemplateId(value string) (err error) {
	return instance.SetProperty("SecureBootTemplateId", (value))
}

// GetSecureBootTemplateId gets the value of SecureBootTemplateId for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertySecureBootTemplateId() (value string, err error) {
	retValue, err := instance.GetProperty("SecureBootTemplateId")
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

// SetTurnOffOnGuestRestart sets the value of TurnOffOnGuestRestart for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyTurnOffOnGuestRestart(value bool) (err error) {
	return instance.SetProperty("TurnOffOnGuestRestart", (value))
}

// GetTurnOffOnGuestRestart gets the value of TurnOffOnGuestRestart for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyTurnOffOnGuestRestart() (value bool, err error) {
	retValue, err := instance.GetProperty("TurnOffOnGuestRestart")
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

// SetUserSnapshotType sets the value of UserSnapshotType for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyUserSnapshotType(value uint16) (err error) {
	return instance.SetProperty("UserSnapshotType", (value))
}

// GetUserSnapshotType gets the value of UserSnapshotType for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyUserSnapshotType() (value uint16, err error) {
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
func (instance *Msvm_VirtualSystemSettingData) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyVersion() (value string, err error) {
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

// SetVirtualNumaEnabled sets the value of VirtualNumaEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyVirtualNumaEnabled(value bool) (err error) {
	return instance.SetProperty("VirtualNumaEnabled", (value))
}

// GetVirtualNumaEnabled gets the value of VirtualNumaEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyVirtualNumaEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("VirtualNumaEnabled")
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

// SetVirtualSlitType sets the value of VirtualSlitType for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyVirtualSlitType(value VirtualSystemSettingData_VirtualSlitType) (err error) {
	return instance.SetProperty("VirtualSlitType", (value))
}

// GetVirtualSlitType gets the value of VirtualSlitType for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyVirtualSlitType() (value VirtualSystemSettingData_VirtualSlitType, err error) {
	retValue, err := instance.GetProperty("VirtualSlitType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualSystemSettingData_VirtualSlitType(valuetmp)

	return
}

// SetVirtualSystemSubType sets the value of VirtualSystemSubType for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyVirtualSystemSubType(value VirtualSystemSettingData_VirtualSystemSubType) (err error) {
	return instance.SetProperty("VirtualSystemSubType", (value))
}

// GetVirtualSystemSubType gets the value of VirtualSystemSubType for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyVirtualSystemSubType() (value VirtualSystemSettingData_VirtualSystemSubType, err error) {
	retValue, err := instance.GetProperty("VirtualSystemSubType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualSystemSettingData_VirtualSystemSubType(valuetmp)

	return
}

// SetVMBusMessageRedirection sets the value of VMBusMessageRedirection for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyVMBusMessageRedirection(value uint16) (err error) {
	return instance.SetProperty("VMBusMessageRedirection", (value))
}

// GetVMBusMessageRedirection gets the value of VMBusMessageRedirection for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyVMBusMessageRedirection() (value uint16, err error) {
	retValue, err := instance.GetProperty("VMBusMessageRedirection")
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

// SetVtl2AddressRangeBase sets the value of Vtl2AddressRangeBase for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyVtl2AddressRangeBase(value uint64) (err error) {
	return instance.SetProperty("Vtl2AddressRangeBase", (value))
}

// GetVtl2AddressRangeBase gets the value of Vtl2AddressRangeBase for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyVtl2AddressRangeBase() (value uint64, err error) {
	retValue, err := instance.GetProperty("Vtl2AddressRangeBase")
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

// SetVtl2AddressRangeSize sets the value of Vtl2AddressRangeSize for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyVtl2AddressRangeSize(value uint64) (err error) {
	return instance.SetProperty("Vtl2AddressRangeSize", (value))
}

// GetVtl2AddressRangeSize gets the value of Vtl2AddressRangeSize for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyVtl2AddressRangeSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Vtl2AddressRangeSize")
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

// SetVtl2AddressSpaceConfigurationMode sets the value of Vtl2AddressSpaceConfigurationMode for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyVtl2AddressSpaceConfigurationMode(value VirtualSystemSettingData_Vtl2AddressSpaceConfigurationMode) (err error) {
	return instance.SetProperty("Vtl2AddressSpaceConfigurationMode", (value))
}

// GetVtl2AddressSpaceConfigurationMode gets the value of Vtl2AddressSpaceConfigurationMode for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyVtl2AddressSpaceConfigurationMode() (value VirtualSystemSettingData_Vtl2AddressSpaceConfigurationMode, err error) {
	retValue, err := instance.GetProperty("Vtl2AddressSpaceConfigurationMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualSystemSettingData_Vtl2AddressSpaceConfigurationMode(valuetmp)

	return
}

// SetVtl2MmioAddressRangeSize sets the value of Vtl2MmioAddressRangeSize for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyVtl2MmioAddressRangeSize(value uint64) (err error) {
	return instance.SetProperty("Vtl2MmioAddressRangeSize", (value))
}

// GetVtl2MmioAddressRangeSize gets the value of Vtl2MmioAddressRangeSize for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyVtl2MmioAddressRangeSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Vtl2MmioAddressRangeSize")
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

// SetWatchdogEnabled sets the value of WatchdogEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyWatchdogEnabled(value bool) (err error) {
	return instance.SetProperty("WatchdogEnabled", (value))
}

// GetWatchdogEnabled gets the value of WatchdogEnabled for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyWatchdogEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("WatchdogEnabled")
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

// SetWorkerJobObjectName sets the value of WorkerJobObjectName for the instance
func (instance *Msvm_VirtualSystemSettingData) SetPropertyWorkerJobObjectName(value string) (err error) {
	return instance.SetProperty("WorkerJobObjectName", (value))
}

// GetWorkerJobObjectName gets the value of WorkerJobObjectName for the instance
func (instance *Msvm_VirtualSystemSettingData) GetPropertyWorkerJobObjectName() (value string, err error) {
	retValue, err := instance.GetProperty("WorkerJobObjectName")
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
