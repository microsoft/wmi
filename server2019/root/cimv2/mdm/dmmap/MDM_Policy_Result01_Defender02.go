// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_Result01_Defender02 struct
type MDM_Policy_Result01_Defender02 struct {
	*cim.WmiInstance

	//
	AllowArchiveScanning int32

	//
	AllowBehaviorMonitoring int32

	//
	AllowCloudProtection int32

	//
	AllowEmailScanning int32

	//
	AllowFullScanOnMappedNetworkDrives int32

	//
	AllowFullScanRemovableDriveScanning int32

	//
	AllowIntrusionPreventionSystem int32

	//
	AllowIOAVProtection int32

	//
	AllowOnAccessProtection int32

	//
	AllowRealtimeMonitoring int32

	//
	AllowScanningNetworkFiles int32

	//
	AllowScriptScanning int32

	//
	AllowUserUIAccess int32

	//
	AttackSurfaceReductionOnlyExclusions string

	//
	AttackSurfaceReductionRules string

	//
	AvgCPULoadFactor int32

	//
	CheckForSignaturesBeforeRunningScan int32

	//
	CloudBlockLevel int32

	//
	CloudExtendedTimeout int32

	//
	ControlledFolderAccessAllowedApplications string

	//
	ControlledFolderAccessProtectedFolders string

	//
	DaysToRetainCleanedMalware int32

	//
	DisableCatchupFullScan int32

	//
	DisableCatchupQuickScan int32

	//
	EnableControlledFolderAccess int32

	//
	EnableLowCPUPriority int32

	//
	EnableNetworkProtection int32

	//
	ExcludedExtensions string

	//
	ExcludedPaths string

	//
	ExcludedProcesses string

	//
	InstanceID string

	//
	ParentID string

	//
	PUAProtection int32

	//
	RealTimeScanDirection int32

	//
	ScanParameter int32

	//
	ScheduleQuickScanTime int32

	//
	ScheduleScanDay int32

	//
	ScheduleScanTime int32

	//
	SecurityIntelligenceLocation string

	//
	SignatureUpdateFallbackOrder string

	//
	SignatureUpdateFileSharesSources string

	//
	SignatureUpdateInterval int32

	//
	SubmitSamplesConsent int32

	//
	ThreatSeverityDefaultAction string
}

func NewMDM_Policy_Result01_Defender02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_Defender02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Defender02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_Defender02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_Defender02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Defender02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowArchiveScanning sets the value of AllowArchiveScanning for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAllowArchiveScanning(value int32) (err error) {
	return instance.SetProperty("AllowArchiveScanning", value)
}

// GetAllowArchiveScanning gets the value of AllowArchiveScanning for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAllowArchiveScanning() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowArchiveScanning")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowBehaviorMonitoring sets the value of AllowBehaviorMonitoring for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAllowBehaviorMonitoring(value int32) (err error) {
	return instance.SetProperty("AllowBehaviorMonitoring", value)
}

// GetAllowBehaviorMonitoring gets the value of AllowBehaviorMonitoring for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAllowBehaviorMonitoring() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowBehaviorMonitoring")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowCloudProtection sets the value of AllowCloudProtection for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAllowCloudProtection(value int32) (err error) {
	return instance.SetProperty("AllowCloudProtection", value)
}

// GetAllowCloudProtection gets the value of AllowCloudProtection for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAllowCloudProtection() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowCloudProtection")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowEmailScanning sets the value of AllowEmailScanning for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAllowEmailScanning(value int32) (err error) {
	return instance.SetProperty("AllowEmailScanning", value)
}

// GetAllowEmailScanning gets the value of AllowEmailScanning for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAllowEmailScanning() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowEmailScanning")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowFullScanOnMappedNetworkDrives sets the value of AllowFullScanOnMappedNetworkDrives for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAllowFullScanOnMappedNetworkDrives(value int32) (err error) {
	return instance.SetProperty("AllowFullScanOnMappedNetworkDrives", value)
}

// GetAllowFullScanOnMappedNetworkDrives gets the value of AllowFullScanOnMappedNetworkDrives for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAllowFullScanOnMappedNetworkDrives() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowFullScanOnMappedNetworkDrives")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowFullScanRemovableDriveScanning sets the value of AllowFullScanRemovableDriveScanning for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAllowFullScanRemovableDriveScanning(value int32) (err error) {
	return instance.SetProperty("AllowFullScanRemovableDriveScanning", value)
}

// GetAllowFullScanRemovableDriveScanning gets the value of AllowFullScanRemovableDriveScanning for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAllowFullScanRemovableDriveScanning() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowFullScanRemovableDriveScanning")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowIntrusionPreventionSystem sets the value of AllowIntrusionPreventionSystem for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAllowIntrusionPreventionSystem(value int32) (err error) {
	return instance.SetProperty("AllowIntrusionPreventionSystem", value)
}

// GetAllowIntrusionPreventionSystem gets the value of AllowIntrusionPreventionSystem for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAllowIntrusionPreventionSystem() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowIntrusionPreventionSystem")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowIOAVProtection sets the value of AllowIOAVProtection for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAllowIOAVProtection(value int32) (err error) {
	return instance.SetProperty("AllowIOAVProtection", value)
}

// GetAllowIOAVProtection gets the value of AllowIOAVProtection for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAllowIOAVProtection() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowIOAVProtection")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowOnAccessProtection sets the value of AllowOnAccessProtection for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAllowOnAccessProtection(value int32) (err error) {
	return instance.SetProperty("AllowOnAccessProtection", value)
}

// GetAllowOnAccessProtection gets the value of AllowOnAccessProtection for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAllowOnAccessProtection() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowOnAccessProtection")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowRealtimeMonitoring sets the value of AllowRealtimeMonitoring for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAllowRealtimeMonitoring(value int32) (err error) {
	return instance.SetProperty("AllowRealtimeMonitoring", value)
}

// GetAllowRealtimeMonitoring gets the value of AllowRealtimeMonitoring for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAllowRealtimeMonitoring() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowRealtimeMonitoring")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowScanningNetworkFiles sets the value of AllowScanningNetworkFiles for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAllowScanningNetworkFiles(value int32) (err error) {
	return instance.SetProperty("AllowScanningNetworkFiles", value)
}

// GetAllowScanningNetworkFiles gets the value of AllowScanningNetworkFiles for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAllowScanningNetworkFiles() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowScanningNetworkFiles")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowScriptScanning sets the value of AllowScriptScanning for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAllowScriptScanning(value int32) (err error) {
	return instance.SetProperty("AllowScriptScanning", value)
}

// GetAllowScriptScanning gets the value of AllowScriptScanning for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAllowScriptScanning() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowScriptScanning")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowUserUIAccess sets the value of AllowUserUIAccess for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAllowUserUIAccess(value int32) (err error) {
	return instance.SetProperty("AllowUserUIAccess", value)
}

// GetAllowUserUIAccess gets the value of AllowUserUIAccess for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAllowUserUIAccess() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowUserUIAccess")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAttackSurfaceReductionOnlyExclusions sets the value of AttackSurfaceReductionOnlyExclusions for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAttackSurfaceReductionOnlyExclusions(value string) (err error) {
	return instance.SetProperty("AttackSurfaceReductionOnlyExclusions", value)
}

// GetAttackSurfaceReductionOnlyExclusions gets the value of AttackSurfaceReductionOnlyExclusions for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAttackSurfaceReductionOnlyExclusions() (value string, err error) {
	retValue, err := instance.GetProperty("AttackSurfaceReductionOnlyExclusions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAttackSurfaceReductionRules sets the value of AttackSurfaceReductionRules for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAttackSurfaceReductionRules(value string) (err error) {
	return instance.SetProperty("AttackSurfaceReductionRules", value)
}

// GetAttackSurfaceReductionRules gets the value of AttackSurfaceReductionRules for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAttackSurfaceReductionRules() (value string, err error) {
	retValue, err := instance.GetProperty("AttackSurfaceReductionRules")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgCPULoadFactor sets the value of AvgCPULoadFactor for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyAvgCPULoadFactor(value int32) (err error) {
	return instance.SetProperty("AvgCPULoadFactor", value)
}

// GetAvgCPULoadFactor gets the value of AvgCPULoadFactor for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyAvgCPULoadFactor() (value int32, err error) {
	retValue, err := instance.GetProperty("AvgCPULoadFactor")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCheckForSignaturesBeforeRunningScan sets the value of CheckForSignaturesBeforeRunningScan for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyCheckForSignaturesBeforeRunningScan(value int32) (err error) {
	return instance.SetProperty("CheckForSignaturesBeforeRunningScan", value)
}

// GetCheckForSignaturesBeforeRunningScan gets the value of CheckForSignaturesBeforeRunningScan for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyCheckForSignaturesBeforeRunningScan() (value int32, err error) {
	retValue, err := instance.GetProperty("CheckForSignaturesBeforeRunningScan")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCloudBlockLevel sets the value of CloudBlockLevel for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyCloudBlockLevel(value int32) (err error) {
	return instance.SetProperty("CloudBlockLevel", value)
}

// GetCloudBlockLevel gets the value of CloudBlockLevel for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyCloudBlockLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("CloudBlockLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCloudExtendedTimeout sets the value of CloudExtendedTimeout for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyCloudExtendedTimeout(value int32) (err error) {
	return instance.SetProperty("CloudExtendedTimeout", value)
}

// GetCloudExtendedTimeout gets the value of CloudExtendedTimeout for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyCloudExtendedTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("CloudExtendedTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetControlledFolderAccessAllowedApplications sets the value of ControlledFolderAccessAllowedApplications for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyControlledFolderAccessAllowedApplications(value string) (err error) {
	return instance.SetProperty("ControlledFolderAccessAllowedApplications", value)
}

// GetControlledFolderAccessAllowedApplications gets the value of ControlledFolderAccessAllowedApplications for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyControlledFolderAccessAllowedApplications() (value string, err error) {
	retValue, err := instance.GetProperty("ControlledFolderAccessAllowedApplications")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetControlledFolderAccessProtectedFolders sets the value of ControlledFolderAccessProtectedFolders for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyControlledFolderAccessProtectedFolders(value string) (err error) {
	return instance.SetProperty("ControlledFolderAccessProtectedFolders", value)
}

// GetControlledFolderAccessProtectedFolders gets the value of ControlledFolderAccessProtectedFolders for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyControlledFolderAccessProtectedFolders() (value string, err error) {
	retValue, err := instance.GetProperty("ControlledFolderAccessProtectedFolders")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDaysToRetainCleanedMalware sets the value of DaysToRetainCleanedMalware for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyDaysToRetainCleanedMalware(value int32) (err error) {
	return instance.SetProperty("DaysToRetainCleanedMalware", value)
}

// GetDaysToRetainCleanedMalware gets the value of DaysToRetainCleanedMalware for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyDaysToRetainCleanedMalware() (value int32, err error) {
	retValue, err := instance.GetProperty("DaysToRetainCleanedMalware")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableCatchupFullScan sets the value of DisableCatchupFullScan for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyDisableCatchupFullScan(value int32) (err error) {
	return instance.SetProperty("DisableCatchupFullScan", value)
}

// GetDisableCatchupFullScan gets the value of DisableCatchupFullScan for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyDisableCatchupFullScan() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableCatchupFullScan")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableCatchupQuickScan sets the value of DisableCatchupQuickScan for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyDisableCatchupQuickScan(value int32) (err error) {
	return instance.SetProperty("DisableCatchupQuickScan", value)
}

// GetDisableCatchupQuickScan gets the value of DisableCatchupQuickScan for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyDisableCatchupQuickScan() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableCatchupQuickScan")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableControlledFolderAccess sets the value of EnableControlledFolderAccess for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyEnableControlledFolderAccess(value int32) (err error) {
	return instance.SetProperty("EnableControlledFolderAccess", value)
}

// GetEnableControlledFolderAccess gets the value of EnableControlledFolderAccess for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyEnableControlledFolderAccess() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableControlledFolderAccess")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableLowCPUPriority sets the value of EnableLowCPUPriority for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyEnableLowCPUPriority(value int32) (err error) {
	return instance.SetProperty("EnableLowCPUPriority", value)
}

// GetEnableLowCPUPriority gets the value of EnableLowCPUPriority for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyEnableLowCPUPriority() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableLowCPUPriority")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableNetworkProtection sets the value of EnableNetworkProtection for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyEnableNetworkProtection(value int32) (err error) {
	return instance.SetProperty("EnableNetworkProtection", value)
}

// GetEnableNetworkProtection gets the value of EnableNetworkProtection for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyEnableNetworkProtection() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableNetworkProtection")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExcludedExtensions sets the value of ExcludedExtensions for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyExcludedExtensions(value string) (err error) {
	return instance.SetProperty("ExcludedExtensions", value)
}

// GetExcludedExtensions gets the value of ExcludedExtensions for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyExcludedExtensions() (value string, err error) {
	retValue, err := instance.GetProperty("ExcludedExtensions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExcludedPaths sets the value of ExcludedPaths for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyExcludedPaths(value string) (err error) {
	return instance.SetProperty("ExcludedPaths", value)
}

// GetExcludedPaths gets the value of ExcludedPaths for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyExcludedPaths() (value string, err error) {
	retValue, err := instance.GetProperty("ExcludedPaths")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExcludedProcesses sets the value of ExcludedProcesses for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyExcludedProcesses(value string) (err error) {
	return instance.SetProperty("ExcludedProcesses", value)
}

// GetExcludedProcesses gets the value of ExcludedProcesses for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyExcludedProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("ExcludedProcesses")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPUAProtection sets the value of PUAProtection for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyPUAProtection(value int32) (err error) {
	return instance.SetProperty("PUAProtection", value)
}

// GetPUAProtection gets the value of PUAProtection for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyPUAProtection() (value int32, err error) {
	retValue, err := instance.GetProperty("PUAProtection")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRealTimeScanDirection sets the value of RealTimeScanDirection for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyRealTimeScanDirection(value int32) (err error) {
	return instance.SetProperty("RealTimeScanDirection", value)
}

// GetRealTimeScanDirection gets the value of RealTimeScanDirection for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyRealTimeScanDirection() (value int32, err error) {
	retValue, err := instance.GetProperty("RealTimeScanDirection")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScanParameter sets the value of ScanParameter for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyScanParameter(value int32) (err error) {
	return instance.SetProperty("ScanParameter", value)
}

// GetScanParameter gets the value of ScanParameter for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyScanParameter() (value int32, err error) {
	retValue, err := instance.GetProperty("ScanParameter")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScheduleQuickScanTime sets the value of ScheduleQuickScanTime for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyScheduleQuickScanTime(value int32) (err error) {
	return instance.SetProperty("ScheduleQuickScanTime", value)
}

// GetScheduleQuickScanTime gets the value of ScheduleQuickScanTime for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyScheduleQuickScanTime() (value int32, err error) {
	retValue, err := instance.GetProperty("ScheduleQuickScanTime")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScheduleScanDay sets the value of ScheduleScanDay for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyScheduleScanDay(value int32) (err error) {
	return instance.SetProperty("ScheduleScanDay", value)
}

// GetScheduleScanDay gets the value of ScheduleScanDay for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyScheduleScanDay() (value int32, err error) {
	retValue, err := instance.GetProperty("ScheduleScanDay")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScheduleScanTime sets the value of ScheduleScanTime for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyScheduleScanTime(value int32) (err error) {
	return instance.SetProperty("ScheduleScanTime", value)
}

// GetScheduleScanTime gets the value of ScheduleScanTime for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyScheduleScanTime() (value int32, err error) {
	retValue, err := instance.GetProperty("ScheduleScanTime")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecurityIntelligenceLocation sets the value of SecurityIntelligenceLocation for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertySecurityIntelligenceLocation(value string) (err error) {
	return instance.SetProperty("SecurityIntelligenceLocation", value)
}

// GetSecurityIntelligenceLocation gets the value of SecurityIntelligenceLocation for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertySecurityIntelligenceLocation() (value string, err error) {
	retValue, err := instance.GetProperty("SecurityIntelligenceLocation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSignatureUpdateFallbackOrder sets the value of SignatureUpdateFallbackOrder for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertySignatureUpdateFallbackOrder(value string) (err error) {
	return instance.SetProperty("SignatureUpdateFallbackOrder", value)
}

// GetSignatureUpdateFallbackOrder gets the value of SignatureUpdateFallbackOrder for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertySignatureUpdateFallbackOrder() (value string, err error) {
	retValue, err := instance.GetProperty("SignatureUpdateFallbackOrder")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSignatureUpdateFileSharesSources sets the value of SignatureUpdateFileSharesSources for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertySignatureUpdateFileSharesSources(value string) (err error) {
	return instance.SetProperty("SignatureUpdateFileSharesSources", value)
}

// GetSignatureUpdateFileSharesSources gets the value of SignatureUpdateFileSharesSources for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertySignatureUpdateFileSharesSources() (value string, err error) {
	retValue, err := instance.GetProperty("SignatureUpdateFileSharesSources")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSignatureUpdateInterval sets the value of SignatureUpdateInterval for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertySignatureUpdateInterval(value int32) (err error) {
	return instance.SetProperty("SignatureUpdateInterval", value)
}

// GetSignatureUpdateInterval gets the value of SignatureUpdateInterval for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertySignatureUpdateInterval() (value int32, err error) {
	retValue, err := instance.GetProperty("SignatureUpdateInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubmitSamplesConsent sets the value of SubmitSamplesConsent for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertySubmitSamplesConsent(value int32) (err error) {
	return instance.SetProperty("SubmitSamplesConsent", value)
}

// GetSubmitSamplesConsent gets the value of SubmitSamplesConsent for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertySubmitSamplesConsent() (value int32, err error) {
	retValue, err := instance.GetProperty("SubmitSamplesConsent")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreatSeverityDefaultAction sets the value of ThreatSeverityDefaultAction for the instance
func (instance *MDM_Policy_Result01_Defender02) SetPropertyThreatSeverityDefaultAction(value string) (err error) {
	return instance.SetProperty("ThreatSeverityDefaultAction", value)
}

// GetThreatSeverityDefaultAction gets the value of ThreatSeverityDefaultAction for the instance
func (instance *MDM_Policy_Result01_Defender02) GetPropertyThreatSeverityDefaultAction() (value string, err error) {
	retValue, err := instance.GetProperty("ThreatSeverityDefaultAction")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
