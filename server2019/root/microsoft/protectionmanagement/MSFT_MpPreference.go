// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.protectionManagement
//////////////////////////////////////////////
package protectionmanagement

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MpPreference struct
type MSFT_MpPreference struct {
	*cim.WmiInstance

	//
	AllowNetworkProtectionOnWinServer bool

	//
	AttackSurfaceReductionOnlyExclusions []string

	//
	AttackSurfaceReductionRules_Actions []uint8

	//
	AttackSurfaceReductionRules_Ids []string

	//
	CheckForSignaturesBeforeRunningScan bool

	//
	CloudBlockLevel uint8

	//
	CloudExtendedTimeout uint32

	//
	ComputerID string

	//
	ControlledFolderAccessAllowedApplications []string

	//
	ControlledFolderAccessProtectedFolders []string

	//
	DisableArchiveScanning bool

	//
	DisableAutoExclusions bool

	//
	DisableBehaviorMonitoring bool

	//
	DisableBlockAtFirstSeen bool

	//
	DisableCatchupFullScan bool

	//
	DisableCatchupQuickScan bool

	//
	DisableCpuThrottleOnIdleScans bool

	//
	DisableDatagramProcessing bool

	//
	DisableEmailScanning bool

	//
	DisableIntrusionPreventionSystem bool

	//
	DisableIOAVProtection bool

	//
	DisablePrivacyMode bool

	//
	DisableRealtimeMonitoring bool

	//
	DisableRemovableDriveScanning bool

	//
	DisableRestorePoint bool

	//
	DisableScanningMappedNetworkDrivesForFullScan bool

	//
	DisableScanningNetworkFiles bool

	//
	DisableScriptScanning bool

	//
	EnableControlledFolderAccess uint8

	//
	EnableFileHashComputation bool

	//
	EnableLowCpuPriority bool

	//
	EnableNetworkProtection uint8

	//
	ExclusionExtension []string

	//
	ExclusionIpAddress []string

	//
	ExclusionPath []string

	//
	ExclusionProcess []string

	//
	HighThreatDefaultAction uint8

	//
	LowThreatDefaultAction uint8

	//
	MAPSReporting uint8

	//
	MeteredConnectionUpdates bool

	//
	ModerateThreatDefaultAction uint8

	//
	PUAProtection uint8

	//
	QuarantinePurgeItemsAfterDelay uint32

	//
	RandomizeScheduleTaskTimes bool

	//
	RealTimeScanDirection uint8

	//
	RemediationScheduleDay uint8

	//
	RemediationScheduleTime string

	//
	ReportingAdditionalActionTimeOut uint32

	//
	ReportingCriticalFailureTimeOut uint32

	//
	ReportingNonCriticalTimeOut uint32

	//
	ScanAvgCPULoadFactor uint8

	//
	ScanOnlyIfIdleEnabled bool

	//
	ScanParameters uint8

	//
	ScanPurgeItemsAfterDelay uint32

	//
	ScanScheduleDay uint8

	//
	ScanScheduleQuickScanTime string

	//
	ScanScheduleTime string

	//
	SevereThreatDefaultAction uint8

	//
	SharedSignaturesPath string

	//
	SignatureAuGracePeriod uint32

	//
	SignatureBlobFileSharesSources string

	//
	SignatureBlobUpdateInterval uint32

	//
	SignatureDefinitionUpdateFileSharesSources string

	//
	SignatureDisableUpdateOnStartupWithoutEngine bool

	//
	SignatureFallbackOrder string

	//
	SignatureFirstAuGracePeriod uint32

	//
	SignatureScheduleDay uint8

	//
	SignatureScheduleTime string

	//
	SignatureUpdateCatchupInterval uint32

	//
	SignatureUpdateInterval uint32

	//
	SubmitSamplesConsent uint8

	//
	ThreatIDDefaultAction_Actions []uint8

	//
	ThreatIDDefaultAction_Ids []int64

	//
	UILockdown bool

	//
	UnknownThreatDefaultAction uint8
}

func NewMSFT_MpPreferenceEx1(instance *cim.WmiInstance) (newInstance *MSFT_MpPreference, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_MpPreference{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_MpPreferenceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MpPreference, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpPreference{
		WmiInstance: tmp,
	}
	return
}

// SetAllowNetworkProtectionOnWinServer sets the value of AllowNetworkProtectionOnWinServer for the instance
func (instance *MSFT_MpPreference) SetPropertyAllowNetworkProtectionOnWinServer(value bool) (err error) {
	return instance.SetProperty("AllowNetworkProtectionOnWinServer", (value))
}

// GetAllowNetworkProtectionOnWinServer gets the value of AllowNetworkProtectionOnWinServer for the instance
func (instance *MSFT_MpPreference) GetPropertyAllowNetworkProtectionOnWinServer() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowNetworkProtectionOnWinServer")
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

// SetAttackSurfaceReductionOnlyExclusions sets the value of AttackSurfaceReductionOnlyExclusions for the instance
func (instance *MSFT_MpPreference) SetPropertyAttackSurfaceReductionOnlyExclusions(value []string) (err error) {
	return instance.SetProperty("AttackSurfaceReductionOnlyExclusions", (value))
}

// GetAttackSurfaceReductionOnlyExclusions gets the value of AttackSurfaceReductionOnlyExclusions for the instance
func (instance *MSFT_MpPreference) GetPropertyAttackSurfaceReductionOnlyExclusions() (value []string, err error) {
	retValue, err := instance.GetProperty("AttackSurfaceReductionOnlyExclusions")
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

// SetAttackSurfaceReductionRules_Actions sets the value of AttackSurfaceReductionRules_Actions for the instance
func (instance *MSFT_MpPreference) SetPropertyAttackSurfaceReductionRules_Actions(value []uint8) (err error) {
	return instance.SetProperty("AttackSurfaceReductionRules_Actions", (value))
}

// GetAttackSurfaceReductionRules_Actions gets the value of AttackSurfaceReductionRules_Actions for the instance
func (instance *MSFT_MpPreference) GetPropertyAttackSurfaceReductionRules_Actions() (value []uint8, err error) {
	retValue, err := instance.GetProperty("AttackSurfaceReductionRules_Actions")
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

// SetAttackSurfaceReductionRules_Ids sets the value of AttackSurfaceReductionRules_Ids for the instance
func (instance *MSFT_MpPreference) SetPropertyAttackSurfaceReductionRules_Ids(value []string) (err error) {
	return instance.SetProperty("AttackSurfaceReductionRules_Ids", (value))
}

// GetAttackSurfaceReductionRules_Ids gets the value of AttackSurfaceReductionRules_Ids for the instance
func (instance *MSFT_MpPreference) GetPropertyAttackSurfaceReductionRules_Ids() (value []string, err error) {
	retValue, err := instance.GetProperty("AttackSurfaceReductionRules_Ids")
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

// SetCheckForSignaturesBeforeRunningScan sets the value of CheckForSignaturesBeforeRunningScan for the instance
func (instance *MSFT_MpPreference) SetPropertyCheckForSignaturesBeforeRunningScan(value bool) (err error) {
	return instance.SetProperty("CheckForSignaturesBeforeRunningScan", (value))
}

// GetCheckForSignaturesBeforeRunningScan gets the value of CheckForSignaturesBeforeRunningScan for the instance
func (instance *MSFT_MpPreference) GetPropertyCheckForSignaturesBeforeRunningScan() (value bool, err error) {
	retValue, err := instance.GetProperty("CheckForSignaturesBeforeRunningScan")
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

// SetCloudBlockLevel sets the value of CloudBlockLevel for the instance
func (instance *MSFT_MpPreference) SetPropertyCloudBlockLevel(value uint8) (err error) {
	return instance.SetProperty("CloudBlockLevel", (value))
}

// GetCloudBlockLevel gets the value of CloudBlockLevel for the instance
func (instance *MSFT_MpPreference) GetPropertyCloudBlockLevel() (value uint8, err error) {
	retValue, err := instance.GetProperty("CloudBlockLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetCloudExtendedTimeout sets the value of CloudExtendedTimeout for the instance
func (instance *MSFT_MpPreference) SetPropertyCloudExtendedTimeout(value uint32) (err error) {
	return instance.SetProperty("CloudExtendedTimeout", (value))
}

// GetCloudExtendedTimeout gets the value of CloudExtendedTimeout for the instance
func (instance *MSFT_MpPreference) GetPropertyCloudExtendedTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("CloudExtendedTimeout")
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

// SetComputerID sets the value of ComputerID for the instance
func (instance *MSFT_MpPreference) SetPropertyComputerID(value string) (err error) {
	return instance.SetProperty("ComputerID", (value))
}

// GetComputerID gets the value of ComputerID for the instance
func (instance *MSFT_MpPreference) GetPropertyComputerID() (value string, err error) {
	retValue, err := instance.GetProperty("ComputerID")
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

// SetControlledFolderAccessAllowedApplications sets the value of ControlledFolderAccessAllowedApplications for the instance
func (instance *MSFT_MpPreference) SetPropertyControlledFolderAccessAllowedApplications(value []string) (err error) {
	return instance.SetProperty("ControlledFolderAccessAllowedApplications", (value))
}

// GetControlledFolderAccessAllowedApplications gets the value of ControlledFolderAccessAllowedApplications for the instance
func (instance *MSFT_MpPreference) GetPropertyControlledFolderAccessAllowedApplications() (value []string, err error) {
	retValue, err := instance.GetProperty("ControlledFolderAccessAllowedApplications")
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

// SetControlledFolderAccessProtectedFolders sets the value of ControlledFolderAccessProtectedFolders for the instance
func (instance *MSFT_MpPreference) SetPropertyControlledFolderAccessProtectedFolders(value []string) (err error) {
	return instance.SetProperty("ControlledFolderAccessProtectedFolders", (value))
}

// GetControlledFolderAccessProtectedFolders gets the value of ControlledFolderAccessProtectedFolders for the instance
func (instance *MSFT_MpPreference) GetPropertyControlledFolderAccessProtectedFolders() (value []string, err error) {
	retValue, err := instance.GetProperty("ControlledFolderAccessProtectedFolders")
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

// SetDisableArchiveScanning sets the value of DisableArchiveScanning for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableArchiveScanning(value bool) (err error) {
	return instance.SetProperty("DisableArchiveScanning", (value))
}

// GetDisableArchiveScanning gets the value of DisableArchiveScanning for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableArchiveScanning() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableArchiveScanning")
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

// SetDisableAutoExclusions sets the value of DisableAutoExclusions for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableAutoExclusions(value bool) (err error) {
	return instance.SetProperty("DisableAutoExclusions", (value))
}

// GetDisableAutoExclusions gets the value of DisableAutoExclusions for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableAutoExclusions() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableAutoExclusions")
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

// SetDisableBehaviorMonitoring sets the value of DisableBehaviorMonitoring for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableBehaviorMonitoring(value bool) (err error) {
	return instance.SetProperty("DisableBehaviorMonitoring", (value))
}

// GetDisableBehaviorMonitoring gets the value of DisableBehaviorMonitoring for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableBehaviorMonitoring() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableBehaviorMonitoring")
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

// SetDisableBlockAtFirstSeen sets the value of DisableBlockAtFirstSeen for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableBlockAtFirstSeen(value bool) (err error) {
	return instance.SetProperty("DisableBlockAtFirstSeen", (value))
}

// GetDisableBlockAtFirstSeen gets the value of DisableBlockAtFirstSeen for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableBlockAtFirstSeen() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableBlockAtFirstSeen")
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

// SetDisableCatchupFullScan sets the value of DisableCatchupFullScan for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableCatchupFullScan(value bool) (err error) {
	return instance.SetProperty("DisableCatchupFullScan", (value))
}

// GetDisableCatchupFullScan gets the value of DisableCatchupFullScan for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableCatchupFullScan() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableCatchupFullScan")
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

// SetDisableCatchupQuickScan sets the value of DisableCatchupQuickScan for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableCatchupQuickScan(value bool) (err error) {
	return instance.SetProperty("DisableCatchupQuickScan", (value))
}

// GetDisableCatchupQuickScan gets the value of DisableCatchupQuickScan for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableCatchupQuickScan() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableCatchupQuickScan")
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

// SetDisableCpuThrottleOnIdleScans sets the value of DisableCpuThrottleOnIdleScans for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableCpuThrottleOnIdleScans(value bool) (err error) {
	return instance.SetProperty("DisableCpuThrottleOnIdleScans", (value))
}

// GetDisableCpuThrottleOnIdleScans gets the value of DisableCpuThrottleOnIdleScans for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableCpuThrottleOnIdleScans() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableCpuThrottleOnIdleScans")
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

// SetDisableDatagramProcessing sets the value of DisableDatagramProcessing for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableDatagramProcessing(value bool) (err error) {
	return instance.SetProperty("DisableDatagramProcessing", (value))
}

// GetDisableDatagramProcessing gets the value of DisableDatagramProcessing for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableDatagramProcessing() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableDatagramProcessing")
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

// SetDisableEmailScanning sets the value of DisableEmailScanning for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableEmailScanning(value bool) (err error) {
	return instance.SetProperty("DisableEmailScanning", (value))
}

// GetDisableEmailScanning gets the value of DisableEmailScanning for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableEmailScanning() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableEmailScanning")
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

// SetDisableIntrusionPreventionSystem sets the value of DisableIntrusionPreventionSystem for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableIntrusionPreventionSystem(value bool) (err error) {
	return instance.SetProperty("DisableIntrusionPreventionSystem", (value))
}

// GetDisableIntrusionPreventionSystem gets the value of DisableIntrusionPreventionSystem for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableIntrusionPreventionSystem() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableIntrusionPreventionSystem")
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

// SetDisableIOAVProtection sets the value of DisableIOAVProtection for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableIOAVProtection(value bool) (err error) {
	return instance.SetProperty("DisableIOAVProtection", (value))
}

// GetDisableIOAVProtection gets the value of DisableIOAVProtection for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableIOAVProtection() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableIOAVProtection")
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

// SetDisablePrivacyMode sets the value of DisablePrivacyMode for the instance
func (instance *MSFT_MpPreference) SetPropertyDisablePrivacyMode(value bool) (err error) {
	return instance.SetProperty("DisablePrivacyMode", (value))
}

// GetDisablePrivacyMode gets the value of DisablePrivacyMode for the instance
func (instance *MSFT_MpPreference) GetPropertyDisablePrivacyMode() (value bool, err error) {
	retValue, err := instance.GetProperty("DisablePrivacyMode")
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

// SetDisableRealtimeMonitoring sets the value of DisableRealtimeMonitoring for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableRealtimeMonitoring(value bool) (err error) {
	return instance.SetProperty("DisableRealtimeMonitoring", (value))
}

// GetDisableRealtimeMonitoring gets the value of DisableRealtimeMonitoring for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableRealtimeMonitoring() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableRealtimeMonitoring")
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

// SetDisableRemovableDriveScanning sets the value of DisableRemovableDriveScanning for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableRemovableDriveScanning(value bool) (err error) {
	return instance.SetProperty("DisableRemovableDriveScanning", (value))
}

// GetDisableRemovableDriveScanning gets the value of DisableRemovableDriveScanning for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableRemovableDriveScanning() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableRemovableDriveScanning")
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

// SetDisableRestorePoint sets the value of DisableRestorePoint for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableRestorePoint(value bool) (err error) {
	return instance.SetProperty("DisableRestorePoint", (value))
}

// GetDisableRestorePoint gets the value of DisableRestorePoint for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableRestorePoint() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableRestorePoint")
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

// SetDisableScanningMappedNetworkDrivesForFullScan sets the value of DisableScanningMappedNetworkDrivesForFullScan for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableScanningMappedNetworkDrivesForFullScan(value bool) (err error) {
	return instance.SetProperty("DisableScanningMappedNetworkDrivesForFullScan", (value))
}

// GetDisableScanningMappedNetworkDrivesForFullScan gets the value of DisableScanningMappedNetworkDrivesForFullScan for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableScanningMappedNetworkDrivesForFullScan() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableScanningMappedNetworkDrivesForFullScan")
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

// SetDisableScanningNetworkFiles sets the value of DisableScanningNetworkFiles for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableScanningNetworkFiles(value bool) (err error) {
	return instance.SetProperty("DisableScanningNetworkFiles", (value))
}

// GetDisableScanningNetworkFiles gets the value of DisableScanningNetworkFiles for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableScanningNetworkFiles() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableScanningNetworkFiles")
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

// SetDisableScriptScanning sets the value of DisableScriptScanning for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableScriptScanning(value bool) (err error) {
	return instance.SetProperty("DisableScriptScanning", (value))
}

// GetDisableScriptScanning gets the value of DisableScriptScanning for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableScriptScanning() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableScriptScanning")
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

// SetEnableControlledFolderAccess sets the value of EnableControlledFolderAccess for the instance
func (instance *MSFT_MpPreference) SetPropertyEnableControlledFolderAccess(value uint8) (err error) {
	return instance.SetProperty("EnableControlledFolderAccess", (value))
}

// GetEnableControlledFolderAccess gets the value of EnableControlledFolderAccess for the instance
func (instance *MSFT_MpPreference) GetPropertyEnableControlledFolderAccess() (value uint8, err error) {
	retValue, err := instance.GetProperty("EnableControlledFolderAccess")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetEnableFileHashComputation sets the value of EnableFileHashComputation for the instance
func (instance *MSFT_MpPreference) SetPropertyEnableFileHashComputation(value bool) (err error) {
	return instance.SetProperty("EnableFileHashComputation", (value))
}

// GetEnableFileHashComputation gets the value of EnableFileHashComputation for the instance
func (instance *MSFT_MpPreference) GetPropertyEnableFileHashComputation() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableFileHashComputation")
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

// SetEnableLowCpuPriority sets the value of EnableLowCpuPriority for the instance
func (instance *MSFT_MpPreference) SetPropertyEnableLowCpuPriority(value bool) (err error) {
	return instance.SetProperty("EnableLowCpuPriority", (value))
}

// GetEnableLowCpuPriority gets the value of EnableLowCpuPriority for the instance
func (instance *MSFT_MpPreference) GetPropertyEnableLowCpuPriority() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableLowCpuPriority")
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

// SetEnableNetworkProtection sets the value of EnableNetworkProtection for the instance
func (instance *MSFT_MpPreference) SetPropertyEnableNetworkProtection(value uint8) (err error) {
	return instance.SetProperty("EnableNetworkProtection", (value))
}

// GetEnableNetworkProtection gets the value of EnableNetworkProtection for the instance
func (instance *MSFT_MpPreference) GetPropertyEnableNetworkProtection() (value uint8, err error) {
	retValue, err := instance.GetProperty("EnableNetworkProtection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetExclusionExtension sets the value of ExclusionExtension for the instance
func (instance *MSFT_MpPreference) SetPropertyExclusionExtension(value []string) (err error) {
	return instance.SetProperty("ExclusionExtension", (value))
}

// GetExclusionExtension gets the value of ExclusionExtension for the instance
func (instance *MSFT_MpPreference) GetPropertyExclusionExtension() (value []string, err error) {
	retValue, err := instance.GetProperty("ExclusionExtension")
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

// SetExclusionIpAddress sets the value of ExclusionIpAddress for the instance
func (instance *MSFT_MpPreference) SetPropertyExclusionIpAddress(value []string) (err error) {
	return instance.SetProperty("ExclusionIpAddress", (value))
}

// GetExclusionIpAddress gets the value of ExclusionIpAddress for the instance
func (instance *MSFT_MpPreference) GetPropertyExclusionIpAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("ExclusionIpAddress")
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

// SetExclusionPath sets the value of ExclusionPath for the instance
func (instance *MSFT_MpPreference) SetPropertyExclusionPath(value []string) (err error) {
	return instance.SetProperty("ExclusionPath", (value))
}

// GetExclusionPath gets the value of ExclusionPath for the instance
func (instance *MSFT_MpPreference) GetPropertyExclusionPath() (value []string, err error) {
	retValue, err := instance.GetProperty("ExclusionPath")
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

// SetExclusionProcess sets the value of ExclusionProcess for the instance
func (instance *MSFT_MpPreference) SetPropertyExclusionProcess(value []string) (err error) {
	return instance.SetProperty("ExclusionProcess", (value))
}

// GetExclusionProcess gets the value of ExclusionProcess for the instance
func (instance *MSFT_MpPreference) GetPropertyExclusionProcess() (value []string, err error) {
	retValue, err := instance.GetProperty("ExclusionProcess")
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

// SetHighThreatDefaultAction sets the value of HighThreatDefaultAction for the instance
func (instance *MSFT_MpPreference) SetPropertyHighThreatDefaultAction(value uint8) (err error) {
	return instance.SetProperty("HighThreatDefaultAction", (value))
}

// GetHighThreatDefaultAction gets the value of HighThreatDefaultAction for the instance
func (instance *MSFT_MpPreference) GetPropertyHighThreatDefaultAction() (value uint8, err error) {
	retValue, err := instance.GetProperty("HighThreatDefaultAction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetLowThreatDefaultAction sets the value of LowThreatDefaultAction for the instance
func (instance *MSFT_MpPreference) SetPropertyLowThreatDefaultAction(value uint8) (err error) {
	return instance.SetProperty("LowThreatDefaultAction", (value))
}

// GetLowThreatDefaultAction gets the value of LowThreatDefaultAction for the instance
func (instance *MSFT_MpPreference) GetPropertyLowThreatDefaultAction() (value uint8, err error) {
	retValue, err := instance.GetProperty("LowThreatDefaultAction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetMAPSReporting sets the value of MAPSReporting for the instance
func (instance *MSFT_MpPreference) SetPropertyMAPSReporting(value uint8) (err error) {
	return instance.SetProperty("MAPSReporting", (value))
}

// GetMAPSReporting gets the value of MAPSReporting for the instance
func (instance *MSFT_MpPreference) GetPropertyMAPSReporting() (value uint8, err error) {
	retValue, err := instance.GetProperty("MAPSReporting")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetMeteredConnectionUpdates sets the value of MeteredConnectionUpdates for the instance
func (instance *MSFT_MpPreference) SetPropertyMeteredConnectionUpdates(value bool) (err error) {
	return instance.SetProperty("MeteredConnectionUpdates", (value))
}

// GetMeteredConnectionUpdates gets the value of MeteredConnectionUpdates for the instance
func (instance *MSFT_MpPreference) GetPropertyMeteredConnectionUpdates() (value bool, err error) {
	retValue, err := instance.GetProperty("MeteredConnectionUpdates")
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

// SetModerateThreatDefaultAction sets the value of ModerateThreatDefaultAction for the instance
func (instance *MSFT_MpPreference) SetPropertyModerateThreatDefaultAction(value uint8) (err error) {
	return instance.SetProperty("ModerateThreatDefaultAction", (value))
}

// GetModerateThreatDefaultAction gets the value of ModerateThreatDefaultAction for the instance
func (instance *MSFT_MpPreference) GetPropertyModerateThreatDefaultAction() (value uint8, err error) {
	retValue, err := instance.GetProperty("ModerateThreatDefaultAction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetPUAProtection sets the value of PUAProtection for the instance
func (instance *MSFT_MpPreference) SetPropertyPUAProtection(value uint8) (err error) {
	return instance.SetProperty("PUAProtection", (value))
}

// GetPUAProtection gets the value of PUAProtection for the instance
func (instance *MSFT_MpPreference) GetPropertyPUAProtection() (value uint8, err error) {
	retValue, err := instance.GetProperty("PUAProtection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetQuarantinePurgeItemsAfterDelay sets the value of QuarantinePurgeItemsAfterDelay for the instance
func (instance *MSFT_MpPreference) SetPropertyQuarantinePurgeItemsAfterDelay(value uint32) (err error) {
	return instance.SetProperty("QuarantinePurgeItemsAfterDelay", (value))
}

// GetQuarantinePurgeItemsAfterDelay gets the value of QuarantinePurgeItemsAfterDelay for the instance
func (instance *MSFT_MpPreference) GetPropertyQuarantinePurgeItemsAfterDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuarantinePurgeItemsAfterDelay")
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

// SetRandomizeScheduleTaskTimes sets the value of RandomizeScheduleTaskTimes for the instance
func (instance *MSFT_MpPreference) SetPropertyRandomizeScheduleTaskTimes(value bool) (err error) {
	return instance.SetProperty("RandomizeScheduleTaskTimes", (value))
}

// GetRandomizeScheduleTaskTimes gets the value of RandomizeScheduleTaskTimes for the instance
func (instance *MSFT_MpPreference) GetPropertyRandomizeScheduleTaskTimes() (value bool, err error) {
	retValue, err := instance.GetProperty("RandomizeScheduleTaskTimes")
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

// SetRealTimeScanDirection sets the value of RealTimeScanDirection for the instance
func (instance *MSFT_MpPreference) SetPropertyRealTimeScanDirection(value uint8) (err error) {
	return instance.SetProperty("RealTimeScanDirection", (value))
}

// GetRealTimeScanDirection gets the value of RealTimeScanDirection for the instance
func (instance *MSFT_MpPreference) GetPropertyRealTimeScanDirection() (value uint8, err error) {
	retValue, err := instance.GetProperty("RealTimeScanDirection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetRemediationScheduleDay sets the value of RemediationScheduleDay for the instance
func (instance *MSFT_MpPreference) SetPropertyRemediationScheduleDay(value uint8) (err error) {
	return instance.SetProperty("RemediationScheduleDay", (value))
}

// GetRemediationScheduleDay gets the value of RemediationScheduleDay for the instance
func (instance *MSFT_MpPreference) GetPropertyRemediationScheduleDay() (value uint8, err error) {
	retValue, err := instance.GetProperty("RemediationScheduleDay")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetRemediationScheduleTime sets the value of RemediationScheduleTime for the instance
func (instance *MSFT_MpPreference) SetPropertyRemediationScheduleTime(value string) (err error) {
	return instance.SetProperty("RemediationScheduleTime", (value))
}

// GetRemediationScheduleTime gets the value of RemediationScheduleTime for the instance
func (instance *MSFT_MpPreference) GetPropertyRemediationScheduleTime() (value string, err error) {
	retValue, err := instance.GetProperty("RemediationScheduleTime")
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

// SetReportingAdditionalActionTimeOut sets the value of ReportingAdditionalActionTimeOut for the instance
func (instance *MSFT_MpPreference) SetPropertyReportingAdditionalActionTimeOut(value uint32) (err error) {
	return instance.SetProperty("ReportingAdditionalActionTimeOut", (value))
}

// GetReportingAdditionalActionTimeOut gets the value of ReportingAdditionalActionTimeOut for the instance
func (instance *MSFT_MpPreference) GetPropertyReportingAdditionalActionTimeOut() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReportingAdditionalActionTimeOut")
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

// SetReportingCriticalFailureTimeOut sets the value of ReportingCriticalFailureTimeOut for the instance
func (instance *MSFT_MpPreference) SetPropertyReportingCriticalFailureTimeOut(value uint32) (err error) {
	return instance.SetProperty("ReportingCriticalFailureTimeOut", (value))
}

// GetReportingCriticalFailureTimeOut gets the value of ReportingCriticalFailureTimeOut for the instance
func (instance *MSFT_MpPreference) GetPropertyReportingCriticalFailureTimeOut() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReportingCriticalFailureTimeOut")
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

// SetReportingNonCriticalTimeOut sets the value of ReportingNonCriticalTimeOut for the instance
func (instance *MSFT_MpPreference) SetPropertyReportingNonCriticalTimeOut(value uint32) (err error) {
	return instance.SetProperty("ReportingNonCriticalTimeOut", (value))
}

// GetReportingNonCriticalTimeOut gets the value of ReportingNonCriticalTimeOut for the instance
func (instance *MSFT_MpPreference) GetPropertyReportingNonCriticalTimeOut() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReportingNonCriticalTimeOut")
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

// SetScanAvgCPULoadFactor sets the value of ScanAvgCPULoadFactor for the instance
func (instance *MSFT_MpPreference) SetPropertyScanAvgCPULoadFactor(value uint8) (err error) {
	return instance.SetProperty("ScanAvgCPULoadFactor", (value))
}

// GetScanAvgCPULoadFactor gets the value of ScanAvgCPULoadFactor for the instance
func (instance *MSFT_MpPreference) GetPropertyScanAvgCPULoadFactor() (value uint8, err error) {
	retValue, err := instance.GetProperty("ScanAvgCPULoadFactor")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetScanOnlyIfIdleEnabled sets the value of ScanOnlyIfIdleEnabled for the instance
func (instance *MSFT_MpPreference) SetPropertyScanOnlyIfIdleEnabled(value bool) (err error) {
	return instance.SetProperty("ScanOnlyIfIdleEnabled", (value))
}

// GetScanOnlyIfIdleEnabled gets the value of ScanOnlyIfIdleEnabled for the instance
func (instance *MSFT_MpPreference) GetPropertyScanOnlyIfIdleEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("ScanOnlyIfIdleEnabled")
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

// SetScanParameters sets the value of ScanParameters for the instance
func (instance *MSFT_MpPreference) SetPropertyScanParameters(value uint8) (err error) {
	return instance.SetProperty("ScanParameters", (value))
}

// GetScanParameters gets the value of ScanParameters for the instance
func (instance *MSFT_MpPreference) GetPropertyScanParameters() (value uint8, err error) {
	retValue, err := instance.GetProperty("ScanParameters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetScanPurgeItemsAfterDelay sets the value of ScanPurgeItemsAfterDelay for the instance
func (instance *MSFT_MpPreference) SetPropertyScanPurgeItemsAfterDelay(value uint32) (err error) {
	return instance.SetProperty("ScanPurgeItemsAfterDelay", (value))
}

// GetScanPurgeItemsAfterDelay gets the value of ScanPurgeItemsAfterDelay for the instance
func (instance *MSFT_MpPreference) GetPropertyScanPurgeItemsAfterDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScanPurgeItemsAfterDelay")
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

// SetScanScheduleDay sets the value of ScanScheduleDay for the instance
func (instance *MSFT_MpPreference) SetPropertyScanScheduleDay(value uint8) (err error) {
	return instance.SetProperty("ScanScheduleDay", (value))
}

// GetScanScheduleDay gets the value of ScanScheduleDay for the instance
func (instance *MSFT_MpPreference) GetPropertyScanScheduleDay() (value uint8, err error) {
	retValue, err := instance.GetProperty("ScanScheduleDay")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetScanScheduleQuickScanTime sets the value of ScanScheduleQuickScanTime for the instance
func (instance *MSFT_MpPreference) SetPropertyScanScheduleQuickScanTime(value string) (err error) {
	return instance.SetProperty("ScanScheduleQuickScanTime", (value))
}

// GetScanScheduleQuickScanTime gets the value of ScanScheduleQuickScanTime for the instance
func (instance *MSFT_MpPreference) GetPropertyScanScheduleQuickScanTime() (value string, err error) {
	retValue, err := instance.GetProperty("ScanScheduleQuickScanTime")
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

// SetScanScheduleTime sets the value of ScanScheduleTime for the instance
func (instance *MSFT_MpPreference) SetPropertyScanScheduleTime(value string) (err error) {
	return instance.SetProperty("ScanScheduleTime", (value))
}

// GetScanScheduleTime gets the value of ScanScheduleTime for the instance
func (instance *MSFT_MpPreference) GetPropertyScanScheduleTime() (value string, err error) {
	retValue, err := instance.GetProperty("ScanScheduleTime")
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

// SetSevereThreatDefaultAction sets the value of SevereThreatDefaultAction for the instance
func (instance *MSFT_MpPreference) SetPropertySevereThreatDefaultAction(value uint8) (err error) {
	return instance.SetProperty("SevereThreatDefaultAction", (value))
}

// GetSevereThreatDefaultAction gets the value of SevereThreatDefaultAction for the instance
func (instance *MSFT_MpPreference) GetPropertySevereThreatDefaultAction() (value uint8, err error) {
	retValue, err := instance.GetProperty("SevereThreatDefaultAction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSharedSignaturesPath sets the value of SharedSignaturesPath for the instance
func (instance *MSFT_MpPreference) SetPropertySharedSignaturesPath(value string) (err error) {
	return instance.SetProperty("SharedSignaturesPath", (value))
}

// GetSharedSignaturesPath gets the value of SharedSignaturesPath for the instance
func (instance *MSFT_MpPreference) GetPropertySharedSignaturesPath() (value string, err error) {
	retValue, err := instance.GetProperty("SharedSignaturesPath")
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

// SetSignatureAuGracePeriod sets the value of SignatureAuGracePeriod for the instance
func (instance *MSFT_MpPreference) SetPropertySignatureAuGracePeriod(value uint32) (err error) {
	return instance.SetProperty("SignatureAuGracePeriod", (value))
}

// GetSignatureAuGracePeriod gets the value of SignatureAuGracePeriod for the instance
func (instance *MSFT_MpPreference) GetPropertySignatureAuGracePeriod() (value uint32, err error) {
	retValue, err := instance.GetProperty("SignatureAuGracePeriod")
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

// SetSignatureBlobFileSharesSources sets the value of SignatureBlobFileSharesSources for the instance
func (instance *MSFT_MpPreference) SetPropertySignatureBlobFileSharesSources(value string) (err error) {
	return instance.SetProperty("SignatureBlobFileSharesSources", (value))
}

// GetSignatureBlobFileSharesSources gets the value of SignatureBlobFileSharesSources for the instance
func (instance *MSFT_MpPreference) GetPropertySignatureBlobFileSharesSources() (value string, err error) {
	retValue, err := instance.GetProperty("SignatureBlobFileSharesSources")
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

// SetSignatureBlobUpdateInterval sets the value of SignatureBlobUpdateInterval for the instance
func (instance *MSFT_MpPreference) SetPropertySignatureBlobUpdateInterval(value uint32) (err error) {
	return instance.SetProperty("SignatureBlobUpdateInterval", (value))
}

// GetSignatureBlobUpdateInterval gets the value of SignatureBlobUpdateInterval for the instance
func (instance *MSFT_MpPreference) GetPropertySignatureBlobUpdateInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("SignatureBlobUpdateInterval")
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

// SetSignatureDefinitionUpdateFileSharesSources sets the value of SignatureDefinitionUpdateFileSharesSources for the instance
func (instance *MSFT_MpPreference) SetPropertySignatureDefinitionUpdateFileSharesSources(value string) (err error) {
	return instance.SetProperty("SignatureDefinitionUpdateFileSharesSources", (value))
}

// GetSignatureDefinitionUpdateFileSharesSources gets the value of SignatureDefinitionUpdateFileSharesSources for the instance
func (instance *MSFT_MpPreference) GetPropertySignatureDefinitionUpdateFileSharesSources() (value string, err error) {
	retValue, err := instance.GetProperty("SignatureDefinitionUpdateFileSharesSources")
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

// SetSignatureDisableUpdateOnStartupWithoutEngine sets the value of SignatureDisableUpdateOnStartupWithoutEngine for the instance
func (instance *MSFT_MpPreference) SetPropertySignatureDisableUpdateOnStartupWithoutEngine(value bool) (err error) {
	return instance.SetProperty("SignatureDisableUpdateOnStartupWithoutEngine", (value))
}

// GetSignatureDisableUpdateOnStartupWithoutEngine gets the value of SignatureDisableUpdateOnStartupWithoutEngine for the instance
func (instance *MSFT_MpPreference) GetPropertySignatureDisableUpdateOnStartupWithoutEngine() (value bool, err error) {
	retValue, err := instance.GetProperty("SignatureDisableUpdateOnStartupWithoutEngine")
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

// SetSignatureFallbackOrder sets the value of SignatureFallbackOrder for the instance
func (instance *MSFT_MpPreference) SetPropertySignatureFallbackOrder(value string) (err error) {
	return instance.SetProperty("SignatureFallbackOrder", (value))
}

// GetSignatureFallbackOrder gets the value of SignatureFallbackOrder for the instance
func (instance *MSFT_MpPreference) GetPropertySignatureFallbackOrder() (value string, err error) {
	retValue, err := instance.GetProperty("SignatureFallbackOrder")
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

// SetSignatureFirstAuGracePeriod sets the value of SignatureFirstAuGracePeriod for the instance
func (instance *MSFT_MpPreference) SetPropertySignatureFirstAuGracePeriod(value uint32) (err error) {
	return instance.SetProperty("SignatureFirstAuGracePeriod", (value))
}

// GetSignatureFirstAuGracePeriod gets the value of SignatureFirstAuGracePeriod for the instance
func (instance *MSFT_MpPreference) GetPropertySignatureFirstAuGracePeriod() (value uint32, err error) {
	retValue, err := instance.GetProperty("SignatureFirstAuGracePeriod")
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

// SetSignatureScheduleDay sets the value of SignatureScheduleDay for the instance
func (instance *MSFT_MpPreference) SetPropertySignatureScheduleDay(value uint8) (err error) {
	return instance.SetProperty("SignatureScheduleDay", (value))
}

// GetSignatureScheduleDay gets the value of SignatureScheduleDay for the instance
func (instance *MSFT_MpPreference) GetPropertySignatureScheduleDay() (value uint8, err error) {
	retValue, err := instance.GetProperty("SignatureScheduleDay")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSignatureScheduleTime sets the value of SignatureScheduleTime for the instance
func (instance *MSFT_MpPreference) SetPropertySignatureScheduleTime(value string) (err error) {
	return instance.SetProperty("SignatureScheduleTime", (value))
}

// GetSignatureScheduleTime gets the value of SignatureScheduleTime for the instance
func (instance *MSFT_MpPreference) GetPropertySignatureScheduleTime() (value string, err error) {
	retValue, err := instance.GetProperty("SignatureScheduleTime")
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

// SetSignatureUpdateCatchupInterval sets the value of SignatureUpdateCatchupInterval for the instance
func (instance *MSFT_MpPreference) SetPropertySignatureUpdateCatchupInterval(value uint32) (err error) {
	return instance.SetProperty("SignatureUpdateCatchupInterval", (value))
}

// GetSignatureUpdateCatchupInterval gets the value of SignatureUpdateCatchupInterval for the instance
func (instance *MSFT_MpPreference) GetPropertySignatureUpdateCatchupInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("SignatureUpdateCatchupInterval")
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

// SetSignatureUpdateInterval sets the value of SignatureUpdateInterval for the instance
func (instance *MSFT_MpPreference) SetPropertySignatureUpdateInterval(value uint32) (err error) {
	return instance.SetProperty("SignatureUpdateInterval", (value))
}

// GetSignatureUpdateInterval gets the value of SignatureUpdateInterval for the instance
func (instance *MSFT_MpPreference) GetPropertySignatureUpdateInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("SignatureUpdateInterval")
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

// SetSubmitSamplesConsent sets the value of SubmitSamplesConsent for the instance
func (instance *MSFT_MpPreference) SetPropertySubmitSamplesConsent(value uint8) (err error) {
	return instance.SetProperty("SubmitSamplesConsent", (value))
}

// GetSubmitSamplesConsent gets the value of SubmitSamplesConsent for the instance
func (instance *MSFT_MpPreference) GetPropertySubmitSamplesConsent() (value uint8, err error) {
	retValue, err := instance.GetProperty("SubmitSamplesConsent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetThreatIDDefaultAction_Actions sets the value of ThreatIDDefaultAction_Actions for the instance
func (instance *MSFT_MpPreference) SetPropertyThreatIDDefaultAction_Actions(value []uint8) (err error) {
	return instance.SetProperty("ThreatIDDefaultAction_Actions", (value))
}

// GetThreatIDDefaultAction_Actions gets the value of ThreatIDDefaultAction_Actions for the instance
func (instance *MSFT_MpPreference) GetPropertyThreatIDDefaultAction_Actions() (value []uint8, err error) {
	retValue, err := instance.GetProperty("ThreatIDDefaultAction_Actions")
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

// SetThreatIDDefaultAction_Ids sets the value of ThreatIDDefaultAction_Ids for the instance
func (instance *MSFT_MpPreference) SetPropertyThreatIDDefaultAction_Ids(value []int64) (err error) {
	return instance.SetProperty("ThreatIDDefaultAction_Ids", (value))
}

// GetThreatIDDefaultAction_Ids gets the value of ThreatIDDefaultAction_Ids for the instance
func (instance *MSFT_MpPreference) GetPropertyThreatIDDefaultAction_Ids() (value []int64, err error) {
	retValue, err := instance.GetProperty("ThreatIDDefaultAction_Ids")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int64)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, int64(valuetmp))
	}

	return
}

// SetUILockdown sets the value of UILockdown for the instance
func (instance *MSFT_MpPreference) SetPropertyUILockdown(value bool) (err error) {
	return instance.SetProperty("UILockdown", (value))
}

// GetUILockdown gets the value of UILockdown for the instance
func (instance *MSFT_MpPreference) GetPropertyUILockdown() (value bool, err error) {
	retValue, err := instance.GetProperty("UILockdown")
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

// SetUnknownThreatDefaultAction sets the value of UnknownThreatDefaultAction for the instance
func (instance *MSFT_MpPreference) SetPropertyUnknownThreatDefaultAction(value uint8) (err error) {
	return instance.SetProperty("UnknownThreatDefaultAction", (value))
}

// GetUnknownThreatDefaultAction gets the value of UnknownThreatDefaultAction for the instance
func (instance *MSFT_MpPreference) GetPropertyUnknownThreatDefaultAction() (value uint8, err error) {
	retValue, err := instance.GetProperty("UnknownThreatDefaultAction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

//

// <param name="AllowNetworkProtectionOnWinServer" type="bool "></param>
// <param name="AttackSurfaceReductionOnlyExclusions" type="string []"></param>
// <param name="AttackSurfaceReductionRules_Actions" type="uint8 []"></param>
// <param name="AttackSurfaceReductionRules_Ids" type="string []"></param>
// <param name="CheckForSignaturesBeforeRunningScan" type="bool "></param>
// <param name="CloudBlockLevel" type="uint8 "></param>
// <param name="CloudExtendedTimeout" type="uint32 "></param>
// <param name="ControlledFolderAccessAllowedApplications" type="string []"></param>
// <param name="ControlledFolderAccessProtectedFolders" type="string []"></param>
// <param name="DisableArchiveScanning" type="bool "></param>
// <param name="DisableAutoExclusions" type="bool "></param>
// <param name="DisableBehaviorMonitoring" type="bool "></param>
// <param name="DisableBlockAtFirstSeen" type="bool "></param>
// <param name="DisableCatchupFullScan" type="bool "></param>
// <param name="DisableCatchupQuickScan" type="bool "></param>
// <param name="DisableCpuThrottleOnIdleScans" type="bool "></param>
// <param name="DisableDatagramProcessing" type="bool "></param>
// <param name="DisableEmailScanning" type="bool "></param>
// <param name="DisableIntrusionPreventionSystem" type="bool "></param>
// <param name="DisableIOAVProtection" type="bool "></param>
// <param name="DisablePrivacyMode" type="bool "></param>
// <param name="DisableRealtimeMonitoring" type="bool "></param>
// <param name="DisableRemovableDriveScanning" type="bool "></param>
// <param name="DisableRestorePoint" type="bool "></param>
// <param name="DisableScanningMappedNetworkDrivesForFullScan" type="bool "></param>
// <param name="DisableScanningNetworkFiles" type="bool "></param>
// <param name="DisableScriptScanning" type="bool "></param>
// <param name="EnableControlledFolderAccess" type="uint8 "></param>
// <param name="EnableFileHashComputation" type="bool "></param>
// <param name="EnableLowCpuPriority" type="bool "></param>
// <param name="EnableNetworkProtection" type="uint8 "></param>
// <param name="ExclusionExtension" type="string []"></param>
// <param name="ExclusionIpAddress" type="string []"></param>
// <param name="ExclusionPath" type="string []"></param>
// <param name="ExclusionProcess" type="string []"></param>
// <param name="Force" type="bool "></param>
// <param name="HighThreatDefaultAction" type="uint8 "></param>
// <param name="LowThreatDefaultAction" type="uint8 "></param>
// <param name="MAPSReporting" type="uint8 "></param>
// <param name="MeteredConnectionUpdates" type="bool "></param>
// <param name="ModerateThreatDefaultAction" type="uint8 "></param>
// <param name="PUAProtection" type="uint8 "></param>
// <param name="QuarantinePurgeItemsAfterDelay" type="uint32 "></param>
// <param name="RandomizeScheduleTaskTimes" type="bool "></param>
// <param name="RealTimeScanDirection" type="uint8 "></param>
// <param name="RemediationScheduleDay" type="uint8 "></param>
// <param name="RemediationScheduleTime" type="string "></param>
// <param name="ReportingAdditionalActionTimeOut" type="uint32 "></param>
// <param name="ReportingCriticalFailureTimeOut" type="uint32 "></param>
// <param name="ReportingNonCriticalTimeOut" type="uint32 "></param>
// <param name="ScanAvgCPULoadFactor" type="uint8 "></param>
// <param name="ScanOnlyIfIdleEnabled" type="bool "></param>
// <param name="ScanParameters" type="uint8 "></param>
// <param name="ScanPurgeItemsAfterDelay" type="uint32 "></param>
// <param name="ScanScheduleDay" type="uint8 "></param>
// <param name="ScanScheduleQuickScanTime" type="string "></param>
// <param name="ScanScheduleTime" type="string "></param>
// <param name="SevereThreatDefaultAction" type="uint8 "></param>
// <param name="SharedSignaturesPath" type="string "></param>
// <param name="SignatureAuGracePeriod" type="uint32 "></param>
// <param name="SignatureBlobFileSharesSources" type="string "></param>
// <param name="SignatureBlobUpdateInterval" type="uint32 "></param>
// <param name="SignatureDefinitionUpdateFileSharesSources" type="string "></param>
// <param name="SignatureDisableUpdateOnStartupWithoutEngine" type="bool "></param>
// <param name="SignatureFallbackOrder" type="string "></param>
// <param name="SignatureFirstAuGracePeriod" type="uint32 "></param>
// <param name="SignatureScheduleDay" type="uint8 "></param>
// <param name="SignatureScheduleTime" type="string "></param>
// <param name="SignatureUpdateCatchupInterval" type="uint32 "></param>
// <param name="SignatureUpdateInterval" type="uint32 "></param>
// <param name="SubmitSamplesConsent" type="uint8 "></param>
// <param name="ThreatIDDefaultAction_Actions" type="uint8 []"></param>
// <param name="ThreatIDDefaultAction_Ids" type="int64 []"></param>
// <param name="UILockdown" type="bool "></param>
// <param name="UnknownThreatDefaultAction" type="uint8 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MpPreference) Set( /* IN */ DisableAutoExclusions bool,
	/* IN */ ExclusionPath []string,
	/* IN */ ExclusionExtension []string,
	/* IN */ ExclusionProcess []string,
	/* IN */ ExclusionIpAddress []string,
	/* IN */ QuarantinePurgeItemsAfterDelay uint32,
	/* IN */ RealTimeScanDirection uint8,
	/* IN */ RemediationScheduleDay uint8,
	/* IN */ RemediationScheduleTime string,
	/* IN */ ReportingAdditionalActionTimeOut uint32,
	/* IN */ ReportingCriticalFailureTimeOut uint32,
	/* IN */ ReportingNonCriticalTimeOut uint32,
	/* IN */ ScanAvgCPULoadFactor uint8,
	/* IN */ CheckForSignaturesBeforeRunningScan bool,
	/* IN */ ScanPurgeItemsAfterDelay uint32,
	/* IN */ ScanOnlyIfIdleEnabled bool,
	/* IN */ ScanParameters uint8,
	/* IN */ ScanScheduleDay uint8,
	/* IN */ ScanScheduleQuickScanTime string,
	/* IN */ ScanScheduleTime string,
	/* IN */ SignatureFirstAuGracePeriod uint32,
	/* IN */ SignatureAuGracePeriod uint32,
	/* IN */ SignatureDefinitionUpdateFileSharesSources string,
	/* IN */ SignatureDisableUpdateOnStartupWithoutEngine bool,
	/* IN */ SignatureFallbackOrder string,
	/* IN */ SignatureScheduleDay uint8,
	/* IN */ SignatureScheduleTime string,
	/* IN */ SignatureUpdateCatchupInterval uint32,
	/* IN */ SignatureBlobFileSharesSources string,
	/* IN */ SignatureUpdateInterval uint32,
	/* IN */ SignatureBlobUpdateInterval uint32,
	/* IN */ MAPSReporting uint8,
	/* IN */ SubmitSamplesConsent uint8,
	/* IN */ DisablePrivacyMode bool,
	/* IN */ RandomizeScheduleTaskTimes bool,
	/* IN */ DisableBehaviorMonitoring bool,
	/* IN */ DisableIntrusionPreventionSystem bool,
	/* IN */ DisableIOAVProtection bool,
	/* IN */ DisableRealtimeMonitoring bool,
	/* IN */ DisableScriptScanning bool,
	/* IN */ DisableArchiveScanning bool,
	/* IN */ DisableCatchupFullScan bool,
	/* IN */ DisableCatchupQuickScan bool,
	/* IN */ DisableEmailScanning bool,
	/* IN */ DisableRemovableDriveScanning bool,
	/* IN */ DisableRestorePoint bool,
	/* IN */ DisableScanningMappedNetworkDrivesForFullScan bool,
	/* IN */ DisableScanningNetworkFiles bool,
	/* IN */ UILockdown bool,
	/* IN */ ThreatIDDefaultAction_Ids []int64,
	/* IN */ ThreatIDDefaultAction_Actions []uint8,
	/* IN */ UnknownThreatDefaultAction uint8,
	/* IN */ LowThreatDefaultAction uint8,
	/* IN */ ModerateThreatDefaultAction uint8,
	/* IN */ HighThreatDefaultAction uint8,
	/* IN */ SevereThreatDefaultAction uint8,
	/* IN */ PUAProtection uint8,
	/* IN */ DisableBlockAtFirstSeen bool,
	/* IN */ CloudBlockLevel uint8,
	/* IN */ CloudExtendedTimeout uint32,
	/* IN */ EnableNetworkProtection uint8,
	/* IN */ EnableControlledFolderAccess uint8,
	/* IN */ AttackSurfaceReductionOnlyExclusions []string,
	/* IN */ AttackSurfaceReductionRules_Ids []string,
	/* IN */ AttackSurfaceReductionRules_Actions []uint8,
	/* IN */ ControlledFolderAccessAllowedApplications []string,
	/* IN */ ControlledFolderAccessProtectedFolders []string,
	/* IN */ SharedSignaturesPath string,
	/* IN */ EnableLowCpuPriority bool,
	/* IN */ EnableFileHashComputation bool,
	/* IN */ MeteredConnectionUpdates bool,
	/* IN */ AllowNetworkProtectionOnWinServer bool,
	/* IN */ DisableDatagramProcessing bool,
	/* IN */ DisableCpuThrottleOnIdleScans bool,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set", DisableAutoExclusions, ExclusionPath, ExclusionExtension, ExclusionProcess, ExclusionIpAddress, QuarantinePurgeItemsAfterDelay, RealTimeScanDirection, RemediationScheduleDay, RemediationScheduleTime, ReportingAdditionalActionTimeOut, ReportingCriticalFailureTimeOut, ReportingNonCriticalTimeOut, ScanAvgCPULoadFactor, CheckForSignaturesBeforeRunningScan, ScanPurgeItemsAfterDelay, ScanOnlyIfIdleEnabled, ScanParameters, ScanScheduleDay, ScanScheduleQuickScanTime, ScanScheduleTime, SignatureFirstAuGracePeriod, SignatureAuGracePeriod, SignatureDefinitionUpdateFileSharesSources, SignatureDisableUpdateOnStartupWithoutEngine, SignatureFallbackOrder, SignatureScheduleDay, SignatureScheduleTime, SignatureUpdateCatchupInterval, SignatureBlobFileSharesSources, SignatureUpdateInterval, SignatureBlobUpdateInterval, MAPSReporting, SubmitSamplesConsent, DisablePrivacyMode, RandomizeScheduleTaskTimes, DisableBehaviorMonitoring, DisableIntrusionPreventionSystem, DisableIOAVProtection, DisableRealtimeMonitoring, DisableScriptScanning, DisableArchiveScanning, DisableCatchupFullScan, DisableCatchupQuickScan, DisableEmailScanning, DisableRemovableDriveScanning, DisableRestorePoint, DisableScanningMappedNetworkDrivesForFullScan, DisableScanningNetworkFiles, UILockdown, ThreatIDDefaultAction_Ids, ThreatIDDefaultAction_Actions, UnknownThreatDefaultAction, LowThreatDefaultAction, ModerateThreatDefaultAction, HighThreatDefaultAction, SevereThreatDefaultAction, PUAProtection, DisableBlockAtFirstSeen, CloudBlockLevel, CloudExtendedTimeout, EnableNetworkProtection, EnableControlledFolderAccess, AttackSurfaceReductionOnlyExclusions, AttackSurfaceReductionRules_Ids, AttackSurfaceReductionRules_Actions, ControlledFolderAccessAllowedApplications, ControlledFolderAccessProtectedFolders, SharedSignaturesPath, EnableLowCpuPriority, EnableFileHashComputation, MeteredConnectionUpdates, AllowNetworkProtectionOnWinServer, DisableDatagramProcessing, DisableCpuThrottleOnIdleScans, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="AllowNetworkProtectionOnWinServer" type="bool "></param>
// <param name="AttackSurfaceReductionOnlyExclusions" type="string []"></param>
// <param name="AttackSurfaceReductionRules_Actions" type="uint8 []"></param>
// <param name="AttackSurfaceReductionRules_Ids" type="string []"></param>
// <param name="CheckForSignaturesBeforeRunningScan" type="bool "></param>
// <param name="CloudBlockLevel" type="bool "></param>
// <param name="CloudExtendedTimeout" type="bool "></param>
// <param name="ControlledFolderAccessAllowedApplications" type="string []"></param>
// <param name="ControlledFolderAccessProtectedFolders" type="string []"></param>
// <param name="DisableArchiveScanning" type="bool "></param>
// <param name="DisableAutoExclusions" type="bool "></param>
// <param name="DisableBehaviorMonitoring" type="bool "></param>
// <param name="DisableBlockAtFirstSeen" type="bool "></param>
// <param name="DisableCatchupFullScan" type="bool "></param>
// <param name="DisableCatchupQuickScan" type="bool "></param>
// <param name="DisableCpuThrottleOnIdleScans" type="bool "></param>
// <param name="DisableDatagramProcessing" type="bool "></param>
// <param name="DisableEmailScanning" type="bool "></param>
// <param name="DisableIntrusionPreventionSystem" type="bool "></param>
// <param name="DisableIOAVProtection" type="bool "></param>
// <param name="DisablePrivacyMode" type="bool "></param>
// <param name="DisableRealtimeMonitoring" type="bool "></param>
// <param name="DisableRemovableDriveScanning" type="bool "></param>
// <param name="DisableRestorePoint" type="bool "></param>
// <param name="DisableScanningMappedNetworkDrivesForFullScan" type="bool "></param>
// <param name="DisableScanningNetworkFiles" type="bool "></param>
// <param name="DisableScriptScanning" type="bool "></param>
// <param name="EnableControlledFolderAccess" type="bool "></param>
// <param name="EnableFileHashComputation" type="bool "></param>
// <param name="EnableLowCpuPriority" type="bool "></param>
// <param name="EnableNetworkProtection" type="bool "></param>
// <param name="ExclusionExtension" type="string []"></param>
// <param name="ExclusionIpAddress" type="string []"></param>
// <param name="ExclusionPath" type="string []"></param>
// <param name="ExclusionProcess" type="string []"></param>
// <param name="Force" type="bool "></param>
// <param name="HighThreatDefaultAction" type="bool "></param>
// <param name="LowThreatDefaultAction" type="bool "></param>
// <param name="MAPSReporting" type="bool "></param>
// <param name="MeteredConnectionUpdates" type="bool "></param>
// <param name="ModerateThreatDefaultAction" type="bool "></param>
// <param name="PUAProtection" type="bool "></param>
// <param name="QuarantinePurgeItemsAfterDelay" type="bool "></param>
// <param name="RandomizeScheduleTaskTimes" type="bool "></param>
// <param name="RealTimeScanDirection" type="bool "></param>
// <param name="RemediationScheduleDay" type="bool "></param>
// <param name="RemediationScheduleTime" type="bool "></param>
// <param name="ReportingAdditionalActionTimeOut" type="bool "></param>
// <param name="ReportingCriticalFailureTimeOut" type="bool "></param>
// <param name="ReportingNonCriticalTimeOut" type="bool "></param>
// <param name="ScanAvgCPULoadFactor" type="bool "></param>
// <param name="ScanOnlyIfIdleEnabled" type="bool "></param>
// <param name="ScanParameters" type="bool "></param>
// <param name="ScanPurgeItemsAfterDelay" type="bool "></param>
// <param name="ScanScheduleDay" type="bool "></param>
// <param name="ScanScheduleQuickScanTime" type="bool "></param>
// <param name="ScanScheduleTime" type="bool "></param>
// <param name="SevereThreatDefaultAction" type="bool "></param>
// <param name="SharedSignaturesPath" type="string "></param>
// <param name="SignatureAuGracePeriod" type="bool "></param>
// <param name="SignatureBlobFileSharesSources" type="bool "></param>
// <param name="SignatureBlobUpdateInterval" type="bool "></param>
// <param name="SignatureDefinitionUpdateFileSharesSources" type="bool "></param>
// <param name="SignatureDisableUpdateOnStartupWithoutEngine" type="bool "></param>
// <param name="SignatureFallbackOrder" type="bool "></param>
// <param name="SignatureFirstAuGracePeriod" type="bool "></param>
// <param name="SignatureScheduleDay" type="bool "></param>
// <param name="SignatureScheduleTime" type="bool "></param>
// <param name="SignatureUpdateCatchupInterval" type="bool "></param>
// <param name="SignatureUpdateInterval" type="bool "></param>
// <param name="SubmitSamplesConsent" type="bool "></param>
// <param name="ThreatIDDefaultAction_Actions" type="uint8 []"></param>
// <param name="ThreatIDDefaultAction_Ids" type="int64 []"></param>
// <param name="UILockdown" type="bool "></param>
// <param name="UnknownThreatDefaultAction" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MpPreference) Remove( /* IN */ DisableAutoExclusions bool,
	/* IN */ ExclusionPath []string,
	/* IN */ ExclusionExtension []string,
	/* IN */ ExclusionProcess []string,
	/* IN */ ExclusionIpAddress []string,
	/* IN */ QuarantinePurgeItemsAfterDelay bool,
	/* IN */ RealTimeScanDirection bool,
	/* IN */ RemediationScheduleDay bool,
	/* IN */ RemediationScheduleTime bool,
	/* IN */ ReportingAdditionalActionTimeOut bool,
	/* IN */ ReportingCriticalFailureTimeOut bool,
	/* IN */ ReportingNonCriticalTimeOut bool,
	/* IN */ ScanAvgCPULoadFactor bool,
	/* IN */ CheckForSignaturesBeforeRunningScan bool,
	/* IN */ ScanPurgeItemsAfterDelay bool,
	/* IN */ ScanOnlyIfIdleEnabled bool,
	/* IN */ ScanParameters bool,
	/* IN */ ScanScheduleDay bool,
	/* IN */ ScanScheduleQuickScanTime bool,
	/* IN */ ScanScheduleTime bool,
	/* IN */ SignatureFirstAuGracePeriod bool,
	/* IN */ SignatureAuGracePeriod bool,
	/* IN */ SignatureDefinitionUpdateFileSharesSources bool,
	/* IN */ SignatureDisableUpdateOnStartupWithoutEngine bool,
	/* IN */ SignatureFallbackOrder bool,
	/* IN */ SignatureScheduleDay bool,
	/* IN */ SignatureScheduleTime bool,
	/* IN */ SignatureUpdateCatchupInterval bool,
	/* IN */ SignatureBlobFileSharesSources bool,
	/* IN */ SignatureUpdateInterval bool,
	/* IN */ SignatureBlobUpdateInterval bool,
	/* IN */ MAPSReporting bool,
	/* IN */ SubmitSamplesConsent bool,
	/* IN */ DisablePrivacyMode bool,
	/* IN */ RandomizeScheduleTaskTimes bool,
	/* IN */ DisableBehaviorMonitoring bool,
	/* IN */ DisableIntrusionPreventionSystem bool,
	/* IN */ DisableIOAVProtection bool,
	/* IN */ DisableRealtimeMonitoring bool,
	/* IN */ DisableScriptScanning bool,
	/* IN */ DisableArchiveScanning bool,
	/* IN */ DisableCatchupFullScan bool,
	/* IN */ DisableCatchupQuickScan bool,
	/* IN */ DisableEmailScanning bool,
	/* IN */ DisableRemovableDriveScanning bool,
	/* IN */ DisableRestorePoint bool,
	/* IN */ DisableScanningMappedNetworkDrivesForFullScan bool,
	/* IN */ DisableScanningNetworkFiles bool,
	/* IN */ UILockdown bool,
	/* IN */ ThreatIDDefaultAction_Ids []int64,
	/* IN */ ThreatIDDefaultAction_Actions []uint8,
	/* IN */ UnknownThreatDefaultAction bool,
	/* IN */ LowThreatDefaultAction bool,
	/* IN */ ModerateThreatDefaultAction bool,
	/* IN */ HighThreatDefaultAction bool,
	/* IN */ SevereThreatDefaultAction bool,
	/* IN */ PUAProtection bool,
	/* IN */ DisableBlockAtFirstSeen bool,
	/* IN */ CloudBlockLevel bool,
	/* IN */ CloudExtendedTimeout bool,
	/* IN */ EnableNetworkProtection bool,
	/* IN */ EnableControlledFolderAccess bool,
	/* IN */ AttackSurfaceReductionOnlyExclusions []string,
	/* IN */ AttackSurfaceReductionRules_Ids []string,
	/* IN */ AttackSurfaceReductionRules_Actions []uint8,
	/* IN */ ControlledFolderAccessAllowedApplications []string,
	/* IN */ ControlledFolderAccessProtectedFolders []string,
	/* IN */ SharedSignaturesPath string,
	/* IN */ EnableLowCpuPriority bool,
	/* IN */ EnableFileHashComputation bool,
	/* IN */ MeteredConnectionUpdates bool,
	/* IN */ AllowNetworkProtectionOnWinServer bool,
	/* IN */ DisableDatagramProcessing bool,
	/* IN */ DisableCpuThrottleOnIdleScans bool,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove", DisableAutoExclusions, ExclusionPath, ExclusionExtension, ExclusionProcess, ExclusionIpAddress, QuarantinePurgeItemsAfterDelay, RealTimeScanDirection, RemediationScheduleDay, RemediationScheduleTime, ReportingAdditionalActionTimeOut, ReportingCriticalFailureTimeOut, ReportingNonCriticalTimeOut, ScanAvgCPULoadFactor, CheckForSignaturesBeforeRunningScan, ScanPurgeItemsAfterDelay, ScanOnlyIfIdleEnabled, ScanParameters, ScanScheduleDay, ScanScheduleQuickScanTime, ScanScheduleTime, SignatureFirstAuGracePeriod, SignatureAuGracePeriod, SignatureDefinitionUpdateFileSharesSources, SignatureDisableUpdateOnStartupWithoutEngine, SignatureFallbackOrder, SignatureScheduleDay, SignatureScheduleTime, SignatureUpdateCatchupInterval, SignatureBlobFileSharesSources, SignatureUpdateInterval, SignatureBlobUpdateInterval, MAPSReporting, SubmitSamplesConsent, DisablePrivacyMode, RandomizeScheduleTaskTimes, DisableBehaviorMonitoring, DisableIntrusionPreventionSystem, DisableIOAVProtection, DisableRealtimeMonitoring, DisableScriptScanning, DisableArchiveScanning, DisableCatchupFullScan, DisableCatchupQuickScan, DisableEmailScanning, DisableRemovableDriveScanning, DisableRestorePoint, DisableScanningMappedNetworkDrivesForFullScan, DisableScanningNetworkFiles, UILockdown, ThreatIDDefaultAction_Ids, ThreatIDDefaultAction_Actions, UnknownThreatDefaultAction, LowThreatDefaultAction, ModerateThreatDefaultAction, HighThreatDefaultAction, SevereThreatDefaultAction, PUAProtection, DisableBlockAtFirstSeen, CloudBlockLevel, CloudExtendedTimeout, EnableNetworkProtection, EnableControlledFolderAccess, AttackSurfaceReductionOnlyExclusions, AttackSurfaceReductionRules_Ids, AttackSurfaceReductionRules_Actions, ControlledFolderAccessAllowedApplications, ControlledFolderAccessProtectedFolders, SharedSignaturesPath, EnableLowCpuPriority, EnableFileHashComputation, MeteredConnectionUpdates, AllowNetworkProtectionOnWinServer, DisableDatagramProcessing, DisableCpuThrottleOnIdleScans, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="AttackSurfaceReductionOnlyExclusions" type="string []"></param>
// <param name="AttackSurfaceReductionRules_Actions" type="uint8 []"></param>
// <param name="AttackSurfaceReductionRules_Ids" type="string []"></param>
// <param name="ControlledFolderAccessAllowedApplications" type="string []"></param>
// <param name="ControlledFolderAccessProtectedFolders" type="string []"></param>
// <param name="ExclusionExtension" type="string []"></param>
// <param name="ExclusionIpAddress" type="string []"></param>
// <param name="ExclusionPath" type="string []"></param>
// <param name="ExclusionProcess" type="string []"></param>
// <param name="Force" type="bool "></param>
// <param name="SharedSignaturesPath" type="string "></param>
// <param name="ThreatIDDefaultAction_Actions" type="uint8 []"></param>
// <param name="ThreatIDDefaultAction_Ids" type="int64 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MpPreference) Add( /* IN */ ExclusionPath []string,
	/* IN */ ExclusionExtension []string,
	/* IN */ ExclusionProcess []string,
	/* IN */ ExclusionIpAddress []string,
	/* IN */ ThreatIDDefaultAction_Ids []int64,
	/* IN */ ThreatIDDefaultAction_Actions []uint8,
	/* IN */ AttackSurfaceReductionOnlyExclusions []string,
	/* IN */ AttackSurfaceReductionRules_Ids []string,
	/* IN */ AttackSurfaceReductionRules_Actions []uint8,
	/* IN */ ControlledFolderAccessAllowedApplications []string,
	/* IN */ ControlledFolderAccessProtectedFolders []string,
	/* IN */ SharedSignaturesPath string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Add", ExclusionPath, ExclusionExtension, ExclusionProcess, ExclusionIpAddress, ThreatIDDefaultAction_Ids, ThreatIDDefaultAction_Actions, AttackSurfaceReductionOnlyExclusions, AttackSurfaceReductionRules_Ids, AttackSurfaceReductionRules_Actions, ControlledFolderAccessAllowedApplications, ControlledFolderAccessProtectedFolders, SharedSignaturesPath, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
