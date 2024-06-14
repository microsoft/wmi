// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Defender
//////////////////////////////////////////////
package defender

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
	AllowDatagramProcessingOnWinServer bool

	//
	AllowNetworkProtectionDownLevel bool

	//
	AllowNetworkProtectionOnWinServer bool

	//
	AllowSwitchToAsyncInspection bool

	//
	ApplyDisableNetworkScanningToIOAV bool

	//
	AttackSurfaceReductionOnlyExclusions []string

	//
	AttackSurfaceReductionRules_Actions []uint8

	//
	AttackSurfaceReductionRules_Ids []string

	//
	AttackSurfaceReductionRules_RuleSpecificExclusions []string

	//
	AttackSurfaceReductionRules_RuleSpecificExclusions_Id []string

	//
	BruteForceProtectionAggressiveness uint8

	//
	BruteForceProtectionConfiguredState uint8

	//
	BruteForceProtectionExclusions []string

	//
	BruteForceProtectionLocalNetworkBlocking bool

	//
	BruteForceProtectionMaxBlockTime uint32

	//
	BruteForceProtectionSkipLearningPeriod bool

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
	DefinitionUpdatesChannel uint8

	//
	DisableArchiveScanning bool

	//
	DisableAutoExclusions bool

	//
	DisableBehaviorMonitoring bool

	//
	DisableBlockAtFirstSeen bool

	//
	DisableCacheMaintenance bool

	//
	DisableCatchupFullScan bool

	//
	DisableCatchupQuickScan bool

	//
	DisableCoreServiceECSIntegration bool

	//
	DisableCoreServiceTelemetry bool

	//
	DisableCpuThrottleOnIdleScans bool

	//
	DisableDatagramProcessing bool

	//
	DisableDnsOverTcpParsing bool

	//
	DisableDnsParsing bool

	//
	DisableEmailScanning bool

	//
	DisableFtpParsing bool

	//
	DisableGradualRelease bool

	//
	DisableHttpParsing bool

	//
	DisableInboundConnectionFiltering bool

	//
	DisableIOAVProtection bool

	//
	DisableNetworkProtectionPerfTelemetry bool

	//
	DisablePrivacyMode bool

	//
	DisableQuicParsing bool

	//
	DisableRdpParsing bool

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
	DisableSmtpParsing bool

	//
	DisableSshParsing bool

	//
	DisableTamperProtection bool

	//
	DisableTlsParsing bool

	//
	EnableControlledFolderAccess uint8

	//
	EnableConvertWarnToBlock bool

	//
	EnableDnsSinkhole bool

	//
	EnableEcsConfiguration bool

	//
	EnableFileHashComputation bool

	//
	EnableFullScanOnBatteryPower bool

	//
	EnableLowCpuPriority bool

	//
	EnableNetworkProtection uint8

	//
	EnableUdpReceiveOffload bool

	//
	EnableUdpSegmentationOffload bool

	//
	EngineUpdatesChannel uint8

	//
	ExclusionExtension []string

	//
	ExclusionIpAddress []string

	//
	ExclusionPath []string

	//
	ExclusionProcess []string

	//
	ForceUseProxyOnly bool

	//
	HideExclusionsFromLocalUsers bool

	//
	HighThreatDefaultAction uint8

	//
	IntelTDTEnabled bool

	//
	LowThreatDefaultAction uint8

	//
	MAPSReporting uint8

	//
	MeteredConnectionUpdates bool

	//
	ModerateThreatDefaultAction uint8

	//
	NetworkProtectionReputationMode uint32

	//
	OobeEnableRtpAndSigUpdate bool

	//
	PerformanceModeStatus uint8

	//
	PlatformUpdatesChannel uint8

	//
	ProxyBypass []string

	//
	ProxyPacUrl string

	//
	ProxyServer string

	//
	PUAProtection uint8

	//
	QuarantinePurgeItemsAfterDelay uint32

	//
	QuickScanIncludeExclusions uint8

	//
	RandomizeScheduleTaskTimes bool

	//
	RealTimeScanDirection uint8

	//
	RemediationScheduleDay uint8

	//
	RemediationScheduleTime string

	//
	RemoteEncryptionProtectionAggressiveness uint8

	//
	RemoteEncryptionProtectionConfiguredState uint8

	//
	RemoteEncryptionProtectionExclusions []string

	//
	RemoteEncryptionProtectionMaxBlockTime uint32

	//
	RemoveScanningThreadPoolCap bool

	//
	ReportDynamicSignatureDroppedEvent bool

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
	ScanScheduleOffset uint32

	//
	ScanScheduleQuickScanTime string

	//
	ScanScheduleTime string

	//
	SchedulerRandomizationTime uint32

	//
	ServiceHealthReportInterval uint32

	//
	SevereThreatDefaultAction uint8

	//
	SharedSignaturesPath string

	//
	SharedSignaturesPathUpdateAtScheduledTimeOnly bool

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
	ThrottleForScheduledScanOnly bool

	//
	TrustLabelProtectionStatus uint32

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

// SetAllowDatagramProcessingOnWinServer sets the value of AllowDatagramProcessingOnWinServer for the instance
func (instance *MSFT_MpPreference) SetPropertyAllowDatagramProcessingOnWinServer(value bool) (err error) {
	return instance.SetProperty("AllowDatagramProcessingOnWinServer", (value))
}

// GetAllowDatagramProcessingOnWinServer gets the value of AllowDatagramProcessingOnWinServer for the instance
func (instance *MSFT_MpPreference) GetPropertyAllowDatagramProcessingOnWinServer() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowDatagramProcessingOnWinServer")
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

// SetAllowNetworkProtectionDownLevel sets the value of AllowNetworkProtectionDownLevel for the instance
func (instance *MSFT_MpPreference) SetPropertyAllowNetworkProtectionDownLevel(value bool) (err error) {
	return instance.SetProperty("AllowNetworkProtectionDownLevel", (value))
}

// GetAllowNetworkProtectionDownLevel gets the value of AllowNetworkProtectionDownLevel for the instance
func (instance *MSFT_MpPreference) GetPropertyAllowNetworkProtectionDownLevel() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowNetworkProtectionDownLevel")
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

// SetAllowSwitchToAsyncInspection sets the value of AllowSwitchToAsyncInspection for the instance
func (instance *MSFT_MpPreference) SetPropertyAllowSwitchToAsyncInspection(value bool) (err error) {
	return instance.SetProperty("AllowSwitchToAsyncInspection", (value))
}

// GetAllowSwitchToAsyncInspection gets the value of AllowSwitchToAsyncInspection for the instance
func (instance *MSFT_MpPreference) GetPropertyAllowSwitchToAsyncInspection() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowSwitchToAsyncInspection")
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

// SetApplyDisableNetworkScanningToIOAV sets the value of ApplyDisableNetworkScanningToIOAV for the instance
func (instance *MSFT_MpPreference) SetPropertyApplyDisableNetworkScanningToIOAV(value bool) (err error) {
	return instance.SetProperty("ApplyDisableNetworkScanningToIOAV", (value))
}

// GetApplyDisableNetworkScanningToIOAV gets the value of ApplyDisableNetworkScanningToIOAV for the instance
func (instance *MSFT_MpPreference) GetPropertyApplyDisableNetworkScanningToIOAV() (value bool, err error) {
	retValue, err := instance.GetProperty("ApplyDisableNetworkScanningToIOAV")
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

// SetAttackSurfaceReductionRules_RuleSpecificExclusions sets the value of AttackSurfaceReductionRules_RuleSpecificExclusions for the instance
func (instance *MSFT_MpPreference) SetPropertyAttackSurfaceReductionRules_RuleSpecificExclusions(value []string) (err error) {
	return instance.SetProperty("AttackSurfaceReductionRules_RuleSpecificExclusions", (value))
}

// GetAttackSurfaceReductionRules_RuleSpecificExclusions gets the value of AttackSurfaceReductionRules_RuleSpecificExclusions for the instance
func (instance *MSFT_MpPreference) GetPropertyAttackSurfaceReductionRules_RuleSpecificExclusions() (value []string, err error) {
	retValue, err := instance.GetProperty("AttackSurfaceReductionRules_RuleSpecificExclusions")
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

// SetAttackSurfaceReductionRules_RuleSpecificExclusions_Id sets the value of AttackSurfaceReductionRules_RuleSpecificExclusions_Id for the instance
func (instance *MSFT_MpPreference) SetPropertyAttackSurfaceReductionRules_RuleSpecificExclusions_Id(value []string) (err error) {
	return instance.SetProperty("AttackSurfaceReductionRules_RuleSpecificExclusions_Id", (value))
}

// GetAttackSurfaceReductionRules_RuleSpecificExclusions_Id gets the value of AttackSurfaceReductionRules_RuleSpecificExclusions_Id for the instance
func (instance *MSFT_MpPreference) GetPropertyAttackSurfaceReductionRules_RuleSpecificExclusions_Id() (value []string, err error) {
	retValue, err := instance.GetProperty("AttackSurfaceReductionRules_RuleSpecificExclusions_Id")
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

// SetBruteForceProtectionAggressiveness sets the value of BruteForceProtectionAggressiveness for the instance
func (instance *MSFT_MpPreference) SetPropertyBruteForceProtectionAggressiveness(value uint8) (err error) {
	return instance.SetProperty("BruteForceProtectionAggressiveness", (value))
}

// GetBruteForceProtectionAggressiveness gets the value of BruteForceProtectionAggressiveness for the instance
func (instance *MSFT_MpPreference) GetPropertyBruteForceProtectionAggressiveness() (value uint8, err error) {
	retValue, err := instance.GetProperty("BruteForceProtectionAggressiveness")
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

// SetBruteForceProtectionConfiguredState sets the value of BruteForceProtectionConfiguredState for the instance
func (instance *MSFT_MpPreference) SetPropertyBruteForceProtectionConfiguredState(value uint8) (err error) {
	return instance.SetProperty("BruteForceProtectionConfiguredState", (value))
}

// GetBruteForceProtectionConfiguredState gets the value of BruteForceProtectionConfiguredState for the instance
func (instance *MSFT_MpPreference) GetPropertyBruteForceProtectionConfiguredState() (value uint8, err error) {
	retValue, err := instance.GetProperty("BruteForceProtectionConfiguredState")
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

// SetBruteForceProtectionExclusions sets the value of BruteForceProtectionExclusions for the instance
func (instance *MSFT_MpPreference) SetPropertyBruteForceProtectionExclusions(value []string) (err error) {
	return instance.SetProperty("BruteForceProtectionExclusions", (value))
}

// GetBruteForceProtectionExclusions gets the value of BruteForceProtectionExclusions for the instance
func (instance *MSFT_MpPreference) GetPropertyBruteForceProtectionExclusions() (value []string, err error) {
	retValue, err := instance.GetProperty("BruteForceProtectionExclusions")
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

// SetBruteForceProtectionLocalNetworkBlocking sets the value of BruteForceProtectionLocalNetworkBlocking for the instance
func (instance *MSFT_MpPreference) SetPropertyBruteForceProtectionLocalNetworkBlocking(value bool) (err error) {
	return instance.SetProperty("BruteForceProtectionLocalNetworkBlocking", (value))
}

// GetBruteForceProtectionLocalNetworkBlocking gets the value of BruteForceProtectionLocalNetworkBlocking for the instance
func (instance *MSFT_MpPreference) GetPropertyBruteForceProtectionLocalNetworkBlocking() (value bool, err error) {
	retValue, err := instance.GetProperty("BruteForceProtectionLocalNetworkBlocking")
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

// SetBruteForceProtectionMaxBlockTime sets the value of BruteForceProtectionMaxBlockTime for the instance
func (instance *MSFT_MpPreference) SetPropertyBruteForceProtectionMaxBlockTime(value uint32) (err error) {
	return instance.SetProperty("BruteForceProtectionMaxBlockTime", (value))
}

// GetBruteForceProtectionMaxBlockTime gets the value of BruteForceProtectionMaxBlockTime for the instance
func (instance *MSFT_MpPreference) GetPropertyBruteForceProtectionMaxBlockTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("BruteForceProtectionMaxBlockTime")
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

// SetBruteForceProtectionSkipLearningPeriod sets the value of BruteForceProtectionSkipLearningPeriod for the instance
func (instance *MSFT_MpPreference) SetPropertyBruteForceProtectionSkipLearningPeriod(value bool) (err error) {
	return instance.SetProperty("BruteForceProtectionSkipLearningPeriod", (value))
}

// GetBruteForceProtectionSkipLearningPeriod gets the value of BruteForceProtectionSkipLearningPeriod for the instance
func (instance *MSFT_MpPreference) GetPropertyBruteForceProtectionSkipLearningPeriod() (value bool, err error) {
	retValue, err := instance.GetProperty("BruteForceProtectionSkipLearningPeriod")
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

// SetDefinitionUpdatesChannel sets the value of DefinitionUpdatesChannel for the instance
func (instance *MSFT_MpPreference) SetPropertyDefinitionUpdatesChannel(value uint8) (err error) {
	return instance.SetProperty("DefinitionUpdatesChannel", (value))
}

// GetDefinitionUpdatesChannel gets the value of DefinitionUpdatesChannel for the instance
func (instance *MSFT_MpPreference) GetPropertyDefinitionUpdatesChannel() (value uint8, err error) {
	retValue, err := instance.GetProperty("DefinitionUpdatesChannel")
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

// SetDisableCacheMaintenance sets the value of DisableCacheMaintenance for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableCacheMaintenance(value bool) (err error) {
	return instance.SetProperty("DisableCacheMaintenance", (value))
}

// GetDisableCacheMaintenance gets the value of DisableCacheMaintenance for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableCacheMaintenance() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableCacheMaintenance")
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

// SetDisableCoreServiceECSIntegration sets the value of DisableCoreServiceECSIntegration for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableCoreServiceECSIntegration(value bool) (err error) {
	return instance.SetProperty("DisableCoreServiceECSIntegration", (value))
}

// GetDisableCoreServiceECSIntegration gets the value of DisableCoreServiceECSIntegration for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableCoreServiceECSIntegration() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableCoreServiceECSIntegration")
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

// SetDisableCoreServiceTelemetry sets the value of DisableCoreServiceTelemetry for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableCoreServiceTelemetry(value bool) (err error) {
	return instance.SetProperty("DisableCoreServiceTelemetry", (value))
}

// GetDisableCoreServiceTelemetry gets the value of DisableCoreServiceTelemetry for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableCoreServiceTelemetry() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableCoreServiceTelemetry")
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

// SetDisableDnsOverTcpParsing sets the value of DisableDnsOverTcpParsing for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableDnsOverTcpParsing(value bool) (err error) {
	return instance.SetProperty("DisableDnsOverTcpParsing", (value))
}

// GetDisableDnsOverTcpParsing gets the value of DisableDnsOverTcpParsing for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableDnsOverTcpParsing() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableDnsOverTcpParsing")
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

// SetDisableDnsParsing sets the value of DisableDnsParsing for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableDnsParsing(value bool) (err error) {
	return instance.SetProperty("DisableDnsParsing", (value))
}

// GetDisableDnsParsing gets the value of DisableDnsParsing for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableDnsParsing() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableDnsParsing")
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

// SetDisableFtpParsing sets the value of DisableFtpParsing for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableFtpParsing(value bool) (err error) {
	return instance.SetProperty("DisableFtpParsing", (value))
}

// GetDisableFtpParsing gets the value of DisableFtpParsing for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableFtpParsing() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableFtpParsing")
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

// SetDisableGradualRelease sets the value of DisableGradualRelease for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableGradualRelease(value bool) (err error) {
	return instance.SetProperty("DisableGradualRelease", (value))
}

// GetDisableGradualRelease gets the value of DisableGradualRelease for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableGradualRelease() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableGradualRelease")
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

// SetDisableHttpParsing sets the value of DisableHttpParsing for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableHttpParsing(value bool) (err error) {
	return instance.SetProperty("DisableHttpParsing", (value))
}

// GetDisableHttpParsing gets the value of DisableHttpParsing for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableHttpParsing() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableHttpParsing")
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

// SetDisableInboundConnectionFiltering sets the value of DisableInboundConnectionFiltering for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableInboundConnectionFiltering(value bool) (err error) {
	return instance.SetProperty("DisableInboundConnectionFiltering", (value))
}

// GetDisableInboundConnectionFiltering gets the value of DisableInboundConnectionFiltering for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableInboundConnectionFiltering() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableInboundConnectionFiltering")
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

// SetDisableNetworkProtectionPerfTelemetry sets the value of DisableNetworkProtectionPerfTelemetry for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableNetworkProtectionPerfTelemetry(value bool) (err error) {
	return instance.SetProperty("DisableNetworkProtectionPerfTelemetry", (value))
}

// GetDisableNetworkProtectionPerfTelemetry gets the value of DisableNetworkProtectionPerfTelemetry for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableNetworkProtectionPerfTelemetry() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableNetworkProtectionPerfTelemetry")
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

// SetDisableQuicParsing sets the value of DisableQuicParsing for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableQuicParsing(value bool) (err error) {
	return instance.SetProperty("DisableQuicParsing", (value))
}

// GetDisableQuicParsing gets the value of DisableQuicParsing for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableQuicParsing() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableQuicParsing")
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

// SetDisableRdpParsing sets the value of DisableRdpParsing for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableRdpParsing(value bool) (err error) {
	return instance.SetProperty("DisableRdpParsing", (value))
}

// GetDisableRdpParsing gets the value of DisableRdpParsing for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableRdpParsing() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableRdpParsing")
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

// SetDisableSmtpParsing sets the value of DisableSmtpParsing for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableSmtpParsing(value bool) (err error) {
	return instance.SetProperty("DisableSmtpParsing", (value))
}

// GetDisableSmtpParsing gets the value of DisableSmtpParsing for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableSmtpParsing() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableSmtpParsing")
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

// SetDisableSshParsing sets the value of DisableSshParsing for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableSshParsing(value bool) (err error) {
	return instance.SetProperty("DisableSshParsing", (value))
}

// GetDisableSshParsing gets the value of DisableSshParsing for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableSshParsing() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableSshParsing")
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

// SetDisableTamperProtection sets the value of DisableTamperProtection for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableTamperProtection(value bool) (err error) {
	return instance.SetProperty("DisableTamperProtection", (value))
}

// GetDisableTamperProtection gets the value of DisableTamperProtection for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableTamperProtection() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableTamperProtection")
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

// SetDisableTlsParsing sets the value of DisableTlsParsing for the instance
func (instance *MSFT_MpPreference) SetPropertyDisableTlsParsing(value bool) (err error) {
	return instance.SetProperty("DisableTlsParsing", (value))
}

// GetDisableTlsParsing gets the value of DisableTlsParsing for the instance
func (instance *MSFT_MpPreference) GetPropertyDisableTlsParsing() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableTlsParsing")
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

// SetEnableConvertWarnToBlock sets the value of EnableConvertWarnToBlock for the instance
func (instance *MSFT_MpPreference) SetPropertyEnableConvertWarnToBlock(value bool) (err error) {
	return instance.SetProperty("EnableConvertWarnToBlock", (value))
}

// GetEnableConvertWarnToBlock gets the value of EnableConvertWarnToBlock for the instance
func (instance *MSFT_MpPreference) GetPropertyEnableConvertWarnToBlock() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableConvertWarnToBlock")
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

// SetEnableDnsSinkhole sets the value of EnableDnsSinkhole for the instance
func (instance *MSFT_MpPreference) SetPropertyEnableDnsSinkhole(value bool) (err error) {
	return instance.SetProperty("EnableDnsSinkhole", (value))
}

// GetEnableDnsSinkhole gets the value of EnableDnsSinkhole for the instance
func (instance *MSFT_MpPreference) GetPropertyEnableDnsSinkhole() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableDnsSinkhole")
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

// SetEnableEcsConfiguration sets the value of EnableEcsConfiguration for the instance
func (instance *MSFT_MpPreference) SetPropertyEnableEcsConfiguration(value bool) (err error) {
	return instance.SetProperty("EnableEcsConfiguration", (value))
}

// GetEnableEcsConfiguration gets the value of EnableEcsConfiguration for the instance
func (instance *MSFT_MpPreference) GetPropertyEnableEcsConfiguration() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableEcsConfiguration")
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

// SetEnableFullScanOnBatteryPower sets the value of EnableFullScanOnBatteryPower for the instance
func (instance *MSFT_MpPreference) SetPropertyEnableFullScanOnBatteryPower(value bool) (err error) {
	return instance.SetProperty("EnableFullScanOnBatteryPower", (value))
}

// GetEnableFullScanOnBatteryPower gets the value of EnableFullScanOnBatteryPower for the instance
func (instance *MSFT_MpPreference) GetPropertyEnableFullScanOnBatteryPower() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableFullScanOnBatteryPower")
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

// SetEnableUdpReceiveOffload sets the value of EnableUdpReceiveOffload for the instance
func (instance *MSFT_MpPreference) SetPropertyEnableUdpReceiveOffload(value bool) (err error) {
	return instance.SetProperty("EnableUdpReceiveOffload", (value))
}

// GetEnableUdpReceiveOffload gets the value of EnableUdpReceiveOffload for the instance
func (instance *MSFT_MpPreference) GetPropertyEnableUdpReceiveOffload() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableUdpReceiveOffload")
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

// SetEnableUdpSegmentationOffload sets the value of EnableUdpSegmentationOffload for the instance
func (instance *MSFT_MpPreference) SetPropertyEnableUdpSegmentationOffload(value bool) (err error) {
	return instance.SetProperty("EnableUdpSegmentationOffload", (value))
}

// GetEnableUdpSegmentationOffload gets the value of EnableUdpSegmentationOffload for the instance
func (instance *MSFT_MpPreference) GetPropertyEnableUdpSegmentationOffload() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableUdpSegmentationOffload")
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

// SetEngineUpdatesChannel sets the value of EngineUpdatesChannel for the instance
func (instance *MSFT_MpPreference) SetPropertyEngineUpdatesChannel(value uint8) (err error) {
	return instance.SetProperty("EngineUpdatesChannel", (value))
}

// GetEngineUpdatesChannel gets the value of EngineUpdatesChannel for the instance
func (instance *MSFT_MpPreference) GetPropertyEngineUpdatesChannel() (value uint8, err error) {
	retValue, err := instance.GetProperty("EngineUpdatesChannel")
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

// SetForceUseProxyOnly sets the value of ForceUseProxyOnly for the instance
func (instance *MSFT_MpPreference) SetPropertyForceUseProxyOnly(value bool) (err error) {
	return instance.SetProperty("ForceUseProxyOnly", (value))
}

// GetForceUseProxyOnly gets the value of ForceUseProxyOnly for the instance
func (instance *MSFT_MpPreference) GetPropertyForceUseProxyOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("ForceUseProxyOnly")
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

// SetHideExclusionsFromLocalUsers sets the value of HideExclusionsFromLocalUsers for the instance
func (instance *MSFT_MpPreference) SetPropertyHideExclusionsFromLocalUsers(value bool) (err error) {
	return instance.SetProperty("HideExclusionsFromLocalUsers", (value))
}

// GetHideExclusionsFromLocalUsers gets the value of HideExclusionsFromLocalUsers for the instance
func (instance *MSFT_MpPreference) GetPropertyHideExclusionsFromLocalUsers() (value bool, err error) {
	retValue, err := instance.GetProperty("HideExclusionsFromLocalUsers")
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

// SetIntelTDTEnabled sets the value of IntelTDTEnabled for the instance
func (instance *MSFT_MpPreference) SetPropertyIntelTDTEnabled(value bool) (err error) {
	return instance.SetProperty("IntelTDTEnabled", (value))
}

// GetIntelTDTEnabled gets the value of IntelTDTEnabled for the instance
func (instance *MSFT_MpPreference) GetPropertyIntelTDTEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IntelTDTEnabled")
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

// SetNetworkProtectionReputationMode sets the value of NetworkProtectionReputationMode for the instance
func (instance *MSFT_MpPreference) SetPropertyNetworkProtectionReputationMode(value uint32) (err error) {
	return instance.SetProperty("NetworkProtectionReputationMode", (value))
}

// GetNetworkProtectionReputationMode gets the value of NetworkProtectionReputationMode for the instance
func (instance *MSFT_MpPreference) GetPropertyNetworkProtectionReputationMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("NetworkProtectionReputationMode")
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

// SetOobeEnableRtpAndSigUpdate sets the value of OobeEnableRtpAndSigUpdate for the instance
func (instance *MSFT_MpPreference) SetPropertyOobeEnableRtpAndSigUpdate(value bool) (err error) {
	return instance.SetProperty("OobeEnableRtpAndSigUpdate", (value))
}

// GetOobeEnableRtpAndSigUpdate gets the value of OobeEnableRtpAndSigUpdate for the instance
func (instance *MSFT_MpPreference) GetPropertyOobeEnableRtpAndSigUpdate() (value bool, err error) {
	retValue, err := instance.GetProperty("OobeEnableRtpAndSigUpdate")
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

// SetPerformanceModeStatus sets the value of PerformanceModeStatus for the instance
func (instance *MSFT_MpPreference) SetPropertyPerformanceModeStatus(value uint8) (err error) {
	return instance.SetProperty("PerformanceModeStatus", (value))
}

// GetPerformanceModeStatus gets the value of PerformanceModeStatus for the instance
func (instance *MSFT_MpPreference) GetPropertyPerformanceModeStatus() (value uint8, err error) {
	retValue, err := instance.GetProperty("PerformanceModeStatus")
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

// SetPlatformUpdatesChannel sets the value of PlatformUpdatesChannel for the instance
func (instance *MSFT_MpPreference) SetPropertyPlatformUpdatesChannel(value uint8) (err error) {
	return instance.SetProperty("PlatformUpdatesChannel", (value))
}

// GetPlatformUpdatesChannel gets the value of PlatformUpdatesChannel for the instance
func (instance *MSFT_MpPreference) GetPropertyPlatformUpdatesChannel() (value uint8, err error) {
	retValue, err := instance.GetProperty("PlatformUpdatesChannel")
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

// SetProxyBypass sets the value of ProxyBypass for the instance
func (instance *MSFT_MpPreference) SetPropertyProxyBypass(value []string) (err error) {
	return instance.SetProperty("ProxyBypass", (value))
}

// GetProxyBypass gets the value of ProxyBypass for the instance
func (instance *MSFT_MpPreference) GetPropertyProxyBypass() (value []string, err error) {
	retValue, err := instance.GetProperty("ProxyBypass")
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

// SetProxyPacUrl sets the value of ProxyPacUrl for the instance
func (instance *MSFT_MpPreference) SetPropertyProxyPacUrl(value string) (err error) {
	return instance.SetProperty("ProxyPacUrl", (value))
}

// GetProxyPacUrl gets the value of ProxyPacUrl for the instance
func (instance *MSFT_MpPreference) GetPropertyProxyPacUrl() (value string, err error) {
	retValue, err := instance.GetProperty("ProxyPacUrl")
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

// SetProxyServer sets the value of ProxyServer for the instance
func (instance *MSFT_MpPreference) SetPropertyProxyServer(value string) (err error) {
	return instance.SetProperty("ProxyServer", (value))
}

// GetProxyServer gets the value of ProxyServer for the instance
func (instance *MSFT_MpPreference) GetPropertyProxyServer() (value string, err error) {
	retValue, err := instance.GetProperty("ProxyServer")
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

// SetQuickScanIncludeExclusions sets the value of QuickScanIncludeExclusions for the instance
func (instance *MSFT_MpPreference) SetPropertyQuickScanIncludeExclusions(value uint8) (err error) {
	return instance.SetProperty("QuickScanIncludeExclusions", (value))
}

// GetQuickScanIncludeExclusions gets the value of QuickScanIncludeExclusions for the instance
func (instance *MSFT_MpPreference) GetPropertyQuickScanIncludeExclusions() (value uint8, err error) {
	retValue, err := instance.GetProperty("QuickScanIncludeExclusions")
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

// SetRemoteEncryptionProtectionAggressiveness sets the value of RemoteEncryptionProtectionAggressiveness for the instance
func (instance *MSFT_MpPreference) SetPropertyRemoteEncryptionProtectionAggressiveness(value uint8) (err error) {
	return instance.SetProperty("RemoteEncryptionProtectionAggressiveness", (value))
}

// GetRemoteEncryptionProtectionAggressiveness gets the value of RemoteEncryptionProtectionAggressiveness for the instance
func (instance *MSFT_MpPreference) GetPropertyRemoteEncryptionProtectionAggressiveness() (value uint8, err error) {
	retValue, err := instance.GetProperty("RemoteEncryptionProtectionAggressiveness")
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

// SetRemoteEncryptionProtectionConfiguredState sets the value of RemoteEncryptionProtectionConfiguredState for the instance
func (instance *MSFT_MpPreference) SetPropertyRemoteEncryptionProtectionConfiguredState(value uint8) (err error) {
	return instance.SetProperty("RemoteEncryptionProtectionConfiguredState", (value))
}

// GetRemoteEncryptionProtectionConfiguredState gets the value of RemoteEncryptionProtectionConfiguredState for the instance
func (instance *MSFT_MpPreference) GetPropertyRemoteEncryptionProtectionConfiguredState() (value uint8, err error) {
	retValue, err := instance.GetProperty("RemoteEncryptionProtectionConfiguredState")
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

// SetRemoteEncryptionProtectionExclusions sets the value of RemoteEncryptionProtectionExclusions for the instance
func (instance *MSFT_MpPreference) SetPropertyRemoteEncryptionProtectionExclusions(value []string) (err error) {
	return instance.SetProperty("RemoteEncryptionProtectionExclusions", (value))
}

// GetRemoteEncryptionProtectionExclusions gets the value of RemoteEncryptionProtectionExclusions for the instance
func (instance *MSFT_MpPreference) GetPropertyRemoteEncryptionProtectionExclusions() (value []string, err error) {
	retValue, err := instance.GetProperty("RemoteEncryptionProtectionExclusions")
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

// SetRemoteEncryptionProtectionMaxBlockTime sets the value of RemoteEncryptionProtectionMaxBlockTime for the instance
func (instance *MSFT_MpPreference) SetPropertyRemoteEncryptionProtectionMaxBlockTime(value uint32) (err error) {
	return instance.SetProperty("RemoteEncryptionProtectionMaxBlockTime", (value))
}

// GetRemoteEncryptionProtectionMaxBlockTime gets the value of RemoteEncryptionProtectionMaxBlockTime for the instance
func (instance *MSFT_MpPreference) GetPropertyRemoteEncryptionProtectionMaxBlockTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("RemoteEncryptionProtectionMaxBlockTime")
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

// SetRemoveScanningThreadPoolCap sets the value of RemoveScanningThreadPoolCap for the instance
func (instance *MSFT_MpPreference) SetPropertyRemoveScanningThreadPoolCap(value bool) (err error) {
	return instance.SetProperty("RemoveScanningThreadPoolCap", (value))
}

// GetRemoveScanningThreadPoolCap gets the value of RemoveScanningThreadPoolCap for the instance
func (instance *MSFT_MpPreference) GetPropertyRemoveScanningThreadPoolCap() (value bool, err error) {
	retValue, err := instance.GetProperty("RemoveScanningThreadPoolCap")
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

// SetReportDynamicSignatureDroppedEvent sets the value of ReportDynamicSignatureDroppedEvent for the instance
func (instance *MSFT_MpPreference) SetPropertyReportDynamicSignatureDroppedEvent(value bool) (err error) {
	return instance.SetProperty("ReportDynamicSignatureDroppedEvent", (value))
}

// GetReportDynamicSignatureDroppedEvent gets the value of ReportDynamicSignatureDroppedEvent for the instance
func (instance *MSFT_MpPreference) GetPropertyReportDynamicSignatureDroppedEvent() (value bool, err error) {
	retValue, err := instance.GetProperty("ReportDynamicSignatureDroppedEvent")
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

// SetScanScheduleOffset sets the value of ScanScheduleOffset for the instance
func (instance *MSFT_MpPreference) SetPropertyScanScheduleOffset(value uint32) (err error) {
	return instance.SetProperty("ScanScheduleOffset", (value))
}

// GetScanScheduleOffset gets the value of ScanScheduleOffset for the instance
func (instance *MSFT_MpPreference) GetPropertyScanScheduleOffset() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScanScheduleOffset")
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

// SetSchedulerRandomizationTime sets the value of SchedulerRandomizationTime for the instance
func (instance *MSFT_MpPreference) SetPropertySchedulerRandomizationTime(value uint32) (err error) {
	return instance.SetProperty("SchedulerRandomizationTime", (value))
}

// GetSchedulerRandomizationTime gets the value of SchedulerRandomizationTime for the instance
func (instance *MSFT_MpPreference) GetPropertySchedulerRandomizationTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("SchedulerRandomizationTime")
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

// SetServiceHealthReportInterval sets the value of ServiceHealthReportInterval for the instance
func (instance *MSFT_MpPreference) SetPropertyServiceHealthReportInterval(value uint32) (err error) {
	return instance.SetProperty("ServiceHealthReportInterval", (value))
}

// GetServiceHealthReportInterval gets the value of ServiceHealthReportInterval for the instance
func (instance *MSFT_MpPreference) GetPropertyServiceHealthReportInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("ServiceHealthReportInterval")
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

// SetSharedSignaturesPathUpdateAtScheduledTimeOnly sets the value of SharedSignaturesPathUpdateAtScheduledTimeOnly for the instance
func (instance *MSFT_MpPreference) SetPropertySharedSignaturesPathUpdateAtScheduledTimeOnly(value bool) (err error) {
	return instance.SetProperty("SharedSignaturesPathUpdateAtScheduledTimeOnly", (value))
}

// GetSharedSignaturesPathUpdateAtScheduledTimeOnly gets the value of SharedSignaturesPathUpdateAtScheduledTimeOnly for the instance
func (instance *MSFT_MpPreference) GetPropertySharedSignaturesPathUpdateAtScheduledTimeOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("SharedSignaturesPathUpdateAtScheduledTimeOnly")
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

// SetThrottleForScheduledScanOnly sets the value of ThrottleForScheduledScanOnly for the instance
func (instance *MSFT_MpPreference) SetPropertyThrottleForScheduledScanOnly(value bool) (err error) {
	return instance.SetProperty("ThrottleForScheduledScanOnly", (value))
}

// GetThrottleForScheduledScanOnly gets the value of ThrottleForScheduledScanOnly for the instance
func (instance *MSFT_MpPreference) GetPropertyThrottleForScheduledScanOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("ThrottleForScheduledScanOnly")
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

// SetTrustLabelProtectionStatus sets the value of TrustLabelProtectionStatus for the instance
func (instance *MSFT_MpPreference) SetPropertyTrustLabelProtectionStatus(value uint32) (err error) {
	return instance.SetProperty("TrustLabelProtectionStatus", (value))
}

// GetTrustLabelProtectionStatus gets the value of TrustLabelProtectionStatus for the instance
func (instance *MSFT_MpPreference) GetPropertyTrustLabelProtectionStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("TrustLabelProtectionStatus")
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

// <param name="AllowDatagramProcessingOnWinServer" type="bool "></param>
// <param name="AllowNetworkProtectionDownLevel" type="bool "></param>
// <param name="AllowNetworkProtectionOnWinServer" type="bool "></param>
// <param name="AllowSwitchToAsyncInspection" type="bool "></param>
// <param name="ApplyDisableNetworkScanningToIOAV" type="bool "></param>
// <param name="AttackSurfaceReductionOnlyExclusions" type="string []"></param>
// <param name="AttackSurfaceReductionRules_Actions" type="uint8 []"></param>
// <param name="AttackSurfaceReductionRules_Ids" type="string []"></param>
// <param name="AttackSurfaceReductionRules_RuleSpecificExclusions" type="string []"></param>
// <param name="AttackSurfaceReductionRules_RuleSpecificExclusions_Id" type="string []"></param>
// <param name="BruteForceProtectionAggressiveness" type="uint8 "></param>
// <param name="BruteForceProtectionConfiguredState" type="uint8 "></param>
// <param name="BruteForceProtectionExclusions" type="string []"></param>
// <param name="BruteForceProtectionLocalNetworkBlocking" type="bool "></param>
// <param name="BruteForceProtectionMaxBlockTime" type="uint32 "></param>
// <param name="BruteForceProtectionSkipLearningPeriod" type="bool "></param>
// <param name="CheckForSignaturesBeforeRunningScan" type="bool "></param>
// <param name="CloudBlockLevel" type="uint8 "></param>
// <param name="CloudExtendedTimeout" type="uint32 "></param>
// <param name="ControlledFolderAccessAllowedApplications" type="string []"></param>
// <param name="ControlledFolderAccessProtectedFolders" type="string []"></param>
// <param name="DefinitionUpdatesChannel" type="uint8 "></param>
// <param name="DisableArchiveScanning" type="bool "></param>
// <param name="DisableAutoExclusions" type="bool "></param>
// <param name="DisableBehaviorMonitoring" type="bool "></param>
// <param name="DisableBlockAtFirstSeen" type="bool "></param>
// <param name="DisableCacheMaintenance" type="bool "></param>
// <param name="DisableCatchupFullScan" type="bool "></param>
// <param name="DisableCatchupQuickScan" type="bool "></param>
// <param name="DisableCoreServiceECSIntegration" type="bool "></param>
// <param name="DisableCoreServiceTelemetry" type="bool "></param>
// <param name="DisableCpuThrottleOnIdleScans" type="bool "></param>
// <param name="DisableDatagramProcessing" type="bool "></param>
// <param name="DisableDnsOverTcpParsing" type="bool "></param>
// <param name="DisableDnsParsing" type="bool "></param>
// <param name="DisableEmailScanning" type="bool "></param>
// <param name="DisableFtpParsing" type="bool "></param>
// <param name="DisableGradualRelease" type="bool "></param>
// <param name="DisableHttpParsing" type="bool "></param>
// <param name="DisableInboundConnectionFiltering" type="bool "></param>
// <param name="DisableIntrusionPreventionSystem" type="bool "></param>
// <param name="DisableIOAVProtection" type="bool "></param>
// <param name="DisableNetworkProtectionPerfTelemetry" type="bool "></param>
// <param name="DisablePrivacyMode" type="bool "></param>
// <param name="DisableQuicParsing" type="bool "></param>
// <param name="DisableRdpParsing" type="bool "></param>
// <param name="DisableRealtimeMonitoring" type="bool "></param>
// <param name="DisableRemovableDriveScanning" type="bool "></param>
// <param name="DisableRestorePoint" type="bool "></param>
// <param name="DisableScanningMappedNetworkDrivesForFullScan" type="bool "></param>
// <param name="DisableScanningNetworkFiles" type="bool "></param>
// <param name="DisableScriptScanning" type="bool "></param>
// <param name="DisableSmtpParsing" type="bool "></param>
// <param name="DisableSshParsing" type="bool "></param>
// <param name="DisableTamperProtection" type="bool "></param>
// <param name="DisableTDTFeature" type="bool "></param>
// <param name="DisableTlsParsing" type="bool "></param>
// <param name="EnableControlledFolderAccess" type="uint8 "></param>
// <param name="EnableConvertWarnToBlock" type="bool "></param>
// <param name="EnableDnsSinkhole" type="bool "></param>
// <param name="EnableEcsConfiguration" type="bool "></param>
// <param name="EnableFileHashComputation" type="bool "></param>
// <param name="EnableFullScanOnBatteryPower" type="bool "></param>
// <param name="EnableLowCpuPriority" type="bool "></param>
// <param name="EnableNetworkProtection" type="uint8 "></param>
// <param name="EnableUdpReceiveOffload" type="bool "></param>
// <param name="EnableUdpSegmentationOffload" type="bool "></param>
// <param name="EngineUpdatesChannel" type="uint8 "></param>
// <param name="ExclusionExtension" type="string []"></param>
// <param name="ExclusionIpAddress" type="string []"></param>
// <param name="ExclusionPath" type="string []"></param>
// <param name="ExclusionProcess" type="string []"></param>
// <param name="Force" type="bool "></param>
// <param name="ForceUseProxyOnly" type="bool "></param>
// <param name="HighThreatDefaultAction" type="uint8 "></param>
// <param name="IntelTDTEnabled" type="bool "></param>
// <param name="LowThreatDefaultAction" type="uint8 "></param>
// <param name="MAPSReporting" type="uint8 "></param>
// <param name="MeteredConnectionUpdates" type="bool "></param>
// <param name="ModerateThreatDefaultAction" type="uint8 "></param>
// <param name="NetworkProtectionReputationMode" type="uint32 "></param>
// <param name="OobeEnableRtpAndSigUpdate" type="bool "></param>
// <param name="PerformanceModeStatus" type="uint8 "></param>
// <param name="PlatformUpdatesChannel" type="uint8 "></param>
// <param name="ProxyBypass" type="string []"></param>
// <param name="ProxyPacUrl" type="string "></param>
// <param name="ProxyServer" type="string "></param>
// <param name="PUAProtection" type="uint8 "></param>
// <param name="QuarantinePurgeItemsAfterDelay" type="uint32 "></param>
// <param name="QuickScanIncludeExclusions" type="uint8 "></param>
// <param name="RandomizeScheduleTaskTimes" type="bool "></param>
// <param name="RealTimeScanDirection" type="uint8 "></param>
// <param name="RemediationScheduleDay" type="uint8 "></param>
// <param name="RemediationScheduleTime" type="string "></param>
// <param name="RemoteEncryptionProtectionAggressiveness" type="uint8 "></param>
// <param name="RemoteEncryptionProtectionConfiguredState" type="uint8 "></param>
// <param name="RemoteEncryptionProtectionExclusions" type="string []"></param>
// <param name="RemoteEncryptionProtectionMaxBlockTime" type="uint32 "></param>
// <param name="RemoveScanningThreadPoolCap" type="bool "></param>
// <param name="ReportDynamicSignatureDroppedEvent" type="bool "></param>
// <param name="ReportingAdditionalActionTimeOut" type="uint32 "></param>
// <param name="ReportingCriticalFailureTimeOut" type="uint32 "></param>
// <param name="ReportingNonCriticalTimeOut" type="uint32 "></param>
// <param name="ScanAvgCPULoadFactor" type="uint8 "></param>
// <param name="ScanOnlyIfIdleEnabled" type="bool "></param>
// <param name="ScanParameters" type="uint8 "></param>
// <param name="ScanPurgeItemsAfterDelay" type="uint32 "></param>
// <param name="ScanScheduleDay" type="uint8 "></param>
// <param name="ScanScheduleOffset" type="uint32 "></param>
// <param name="ScanScheduleQuickScanTime" type="string "></param>
// <param name="ScanScheduleTime" type="string "></param>
// <param name="SchedulerRandomizationTime" type="uint32 "></param>
// <param name="ServiceHealthReportInterval" type="uint32 "></param>
// <param name="SevereThreatDefaultAction" type="uint8 "></param>
// <param name="SharedSignaturesPath" type="string "></param>
// <param name="SharedSignaturesPathUpdateAtScheduledTimeOnly" type="bool "></param>
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
// <param name="ThrottleForScheduledScanOnly" type="bool "></param>
// <param name="TrustLabelProtectionStatus" type="uint32 "></param>
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
	/* IN */ RemoteEncryptionProtectionConfiguredState uint8,
	/* IN */ RemoteEncryptionProtectionMaxBlockTime uint32,
	/* IN */ RemoteEncryptionProtectionAggressiveness uint8,
	/* IN */ RemoteEncryptionProtectionExclusions []string,
	/* IN */ BruteForceProtectionConfiguredState uint8,
	/* IN */ BruteForceProtectionMaxBlockTime uint32,
	/* IN */ BruteForceProtectionAggressiveness uint8,
	/* IN */ BruteForceProtectionExclusions []string,
	/* IN */ BruteForceProtectionLocalNetworkBlocking bool,
	/* IN */ BruteForceProtectionSkipLearningPeriod bool,
	/* IN */ ReportingAdditionalActionTimeOut uint32,
	/* IN */ ReportingCriticalFailureTimeOut uint32,
	/* IN */ ReportingNonCriticalTimeOut uint32,
	/* IN */ ServiceHealthReportInterval uint32,
	/* IN */ ReportDynamicSignatureDroppedEvent bool,
	/* IN */ ScanAvgCPULoadFactor uint8,
	/* IN */ CheckForSignaturesBeforeRunningScan bool,
	/* IN */ ScanPurgeItemsAfterDelay uint32,
	/* IN */ ScanOnlyIfIdleEnabled bool,
	/* IN */ ScanParameters uint8,
	/* IN */ ScanScheduleDay uint8,
	/* IN */ ScanScheduleQuickScanTime string,
	/* IN */ ScanScheduleTime string,
	/* IN */ ThrottleForScheduledScanOnly bool,
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
	/* IN */ SchedulerRandomizationTime uint32,
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
	/* IN */ ApplyDisableNetworkScanningToIOAV bool,
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
	/* IN */ SharedSignaturesPathUpdateAtScheduledTimeOnly bool,
	/* IN */ EnableLowCpuPriority bool,
	/* IN */ EnableFileHashComputation bool,
	/* IN */ MeteredConnectionUpdates bool,
	/* IN */ AllowNetworkProtectionOnWinServer bool,
	/* IN */ DisableDatagramProcessing bool,
	/* IN */ EnableConvertWarnToBlock bool,
	/* IN */ DisableCpuThrottleOnIdleScans bool,
	/* IN */ EnableFullScanOnBatteryPower bool,
	/* IN */ ProxyPacUrl string,
	/* IN */ ProxyServer string,
	/* IN */ ProxyBypass []string,
	/* IN */ ForceUseProxyOnly bool,
	/* IN */ DisableTlsParsing bool,
	/* IN */ DisableHttpParsing bool,
	/* IN */ DisableDnsParsing bool,
	/* IN */ DisableDnsOverTcpParsing bool,
	/* IN */ DisableSshParsing bool,
	/* IN */ PlatformUpdatesChannel uint8,
	/* IN */ EngineUpdatesChannel uint8,
	/* IN */ DefinitionUpdatesChannel uint8,
	/* IN */ DisableGradualRelease bool,
	/* IN */ AllowNetworkProtectionDownLevel bool,
	/* IN */ AllowDatagramProcessingOnWinServer bool,
	/* IN */ EnableDnsSinkhole bool,
	/* IN */ DisableInboundConnectionFiltering bool,
	/* IN */ DisableRdpParsing bool,
	/* IN */ DisableNetworkProtectionPerfTelemetry bool,
	/* IN */ TrustLabelProtectionStatus uint32,
	/* IN */ DisableFtpParsing bool,
	/* IN */ AllowSwitchToAsyncInspection bool,
	/* IN */ ScanScheduleOffset uint32,
	/* IN */ DisableTDTFeature bool,
	/* IN */ DisableTamperProtection bool,
	/* IN */ DisableSmtpParsing bool,
	/* IN */ DisableQuicParsing bool,
	/* IN */ NetworkProtectionReputationMode uint32,
	/* IN */ IntelTDTEnabled bool,
	/* IN */ AttackSurfaceReductionRules_RuleSpecificExclusions_Id []string,
	/* IN */ AttackSurfaceReductionRules_RuleSpecificExclusions []string,
	/* IN */ OobeEnableRtpAndSigUpdate bool,
	/* IN */ PerformanceModeStatus uint8,
	/* IN */ QuickScanIncludeExclusions uint8,
	/* IN */ RemoveScanningThreadPoolCap bool,
	/* IN */ DisableCacheMaintenance bool,
	/* IN */ DisableCoreServiceECSIntegration bool,
	/* IN */ DisableCoreServiceTelemetry bool,
	/* IN */ EnableUdpSegmentationOffload bool,
	/* IN */ EnableUdpReceiveOffload bool,
	/* IN */ EnableEcsConfiguration bool,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set", DisableAutoExclusions, ExclusionPath, ExclusionExtension, ExclusionProcess, ExclusionIpAddress, QuarantinePurgeItemsAfterDelay, RealTimeScanDirection, RemediationScheduleDay, RemediationScheduleTime, RemoteEncryptionProtectionConfiguredState, RemoteEncryptionProtectionMaxBlockTime, RemoteEncryptionProtectionAggressiveness, RemoteEncryptionProtectionExclusions, BruteForceProtectionConfiguredState, BruteForceProtectionMaxBlockTime, BruteForceProtectionAggressiveness, BruteForceProtectionExclusions, BruteForceProtectionLocalNetworkBlocking, BruteForceProtectionSkipLearningPeriod, ReportingAdditionalActionTimeOut, ReportingCriticalFailureTimeOut, ReportingNonCriticalTimeOut, ServiceHealthReportInterval, ReportDynamicSignatureDroppedEvent, ScanAvgCPULoadFactor, CheckForSignaturesBeforeRunningScan, ScanPurgeItemsAfterDelay, ScanOnlyIfIdleEnabled, ScanParameters, ScanScheduleDay, ScanScheduleQuickScanTime, ScanScheduleTime, ThrottleForScheduledScanOnly, SignatureFirstAuGracePeriod, SignatureAuGracePeriod, SignatureDefinitionUpdateFileSharesSources, SignatureDisableUpdateOnStartupWithoutEngine, SignatureFallbackOrder, SignatureScheduleDay, SignatureScheduleTime, SignatureUpdateCatchupInterval, SignatureBlobFileSharesSources, SignatureUpdateInterval, SignatureBlobUpdateInterval, MAPSReporting, SubmitSamplesConsent, DisablePrivacyMode, RandomizeScheduleTaskTimes, SchedulerRandomizationTime, DisableBehaviorMonitoring, DisableIntrusionPreventionSystem, DisableIOAVProtection, DisableRealtimeMonitoring, DisableScriptScanning, DisableArchiveScanning, DisableCatchupFullScan, DisableCatchupQuickScan, DisableEmailScanning, DisableRemovableDriveScanning, DisableRestorePoint, DisableScanningMappedNetworkDrivesForFullScan, DisableScanningNetworkFiles, ApplyDisableNetworkScanningToIOAV, UILockdown, ThreatIDDefaultAction_Ids, ThreatIDDefaultAction_Actions, UnknownThreatDefaultAction, LowThreatDefaultAction, ModerateThreatDefaultAction, HighThreatDefaultAction, SevereThreatDefaultAction, PUAProtection, DisableBlockAtFirstSeen, CloudBlockLevel, CloudExtendedTimeout, EnableNetworkProtection, EnableControlledFolderAccess, AttackSurfaceReductionOnlyExclusions, AttackSurfaceReductionRules_Ids, AttackSurfaceReductionRules_Actions, ControlledFolderAccessAllowedApplications, ControlledFolderAccessProtectedFolders, SharedSignaturesPath, SharedSignaturesPathUpdateAtScheduledTimeOnly, EnableLowCpuPriority, EnableFileHashComputation, MeteredConnectionUpdates, AllowNetworkProtectionOnWinServer, DisableDatagramProcessing, EnableConvertWarnToBlock, DisableCpuThrottleOnIdleScans, EnableFullScanOnBatteryPower, ProxyPacUrl, ProxyServer, ProxyBypass, ForceUseProxyOnly, DisableTlsParsing, DisableHttpParsing, DisableDnsParsing, DisableDnsOverTcpParsing, DisableSshParsing, PlatformUpdatesChannel, EngineUpdatesChannel, DefinitionUpdatesChannel, DisableGradualRelease, AllowNetworkProtectionDownLevel, AllowDatagramProcessingOnWinServer, EnableDnsSinkhole, DisableInboundConnectionFiltering, DisableRdpParsing, DisableNetworkProtectionPerfTelemetry, TrustLabelProtectionStatus, DisableFtpParsing, AllowSwitchToAsyncInspection, ScanScheduleOffset, DisableTDTFeature, DisableTamperProtection, DisableSmtpParsing, DisableQuicParsing, NetworkProtectionReputationMode, IntelTDTEnabled, AttackSurfaceReductionRules_RuleSpecificExclusions_Id, AttackSurfaceReductionRules_RuleSpecificExclusions, OobeEnableRtpAndSigUpdate, PerformanceModeStatus, QuickScanIncludeExclusions, RemoveScanningThreadPoolCap, DisableCacheMaintenance, DisableCoreServiceECSIntegration, DisableCoreServiceTelemetry, EnableUdpSegmentationOffload, EnableUdpReceiveOffload, EnableEcsConfiguration, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="AllowDatagramProcessingOnWinServer" type="bool "></param>
// <param name="AllowNetworkProtectionDownLevel" type="bool "></param>
// <param name="AllowNetworkProtectionOnWinServer" type="bool "></param>
// <param name="AllowSwitchToAsyncInspection" type="bool "></param>
// <param name="ApplyDisableNetworkScanningToIOAV" type="bool "></param>
// <param name="AttackSurfaceReductionOnlyExclusions" type="string []"></param>
// <param name="AttackSurfaceReductionRules_Actions" type="uint8 []"></param>
// <param name="AttackSurfaceReductionRules_Ids" type="string []"></param>
// <param name="AttackSurfaceReductionRules_RuleSpecificExclusions" type="string "></param>
// <param name="AttackSurfaceReductionRules_RuleSpecificExclusions_Id" type="string "></param>
// <param name="BruteForceProtectionAggressiveness" type="bool "></param>
// <param name="BruteForceProtectionConfiguredState" type="bool "></param>
// <param name="BruteForceProtectionExclusions" type="string []"></param>
// <param name="BruteForceProtectionLocalNetworkBlocking" type="bool "></param>
// <param name="BruteForceProtectionMaxBlockTime" type="bool "></param>
// <param name="BruteForceProtectionSkipLearningPeriod" type="bool "></param>
// <param name="CheckForSignaturesBeforeRunningScan" type="bool "></param>
// <param name="CloudBlockLevel" type="bool "></param>
// <param name="CloudExtendedTimeout" type="bool "></param>
// <param name="ControlledFolderAccessAllowedApplications" type="string []"></param>
// <param name="ControlledFolderAccessProtectedFolders" type="string []"></param>
// <param name="DefinitionUpdatesChannel" type="bool "></param>
// <param name="DisableArchiveScanning" type="bool "></param>
// <param name="DisableAutoExclusions" type="bool "></param>
// <param name="DisableBehaviorMonitoring" type="bool "></param>
// <param name="DisableBlockAtFirstSeen" type="bool "></param>
// <param name="DisableCacheMaintenance" type="bool "></param>
// <param name="DisableCatchupFullScan" type="bool "></param>
// <param name="DisableCatchupQuickScan" type="bool "></param>
// <param name="DisableCoreServiceECSIntegration" type="bool "></param>
// <param name="DisableCoreServiceTelemetry" type="bool "></param>
// <param name="DisableCpuThrottleOnIdleScans" type="bool "></param>
// <param name="DisableDatagramProcessing" type="bool "></param>
// <param name="DisableDnsOverTcpParsing" type="bool "></param>
// <param name="DisableDnsParsing" type="bool "></param>
// <param name="DisableEmailScanning" type="bool "></param>
// <param name="DisableFtpParsing" type="bool "></param>
// <param name="DisableGradualRelease" type="bool "></param>
// <param name="DisableHttpParsing" type="bool "></param>
// <param name="DisableInboundConnectionFiltering" type="bool "></param>
// <param name="DisableIntrusionPreventionSystem" type="bool "></param>
// <param name="DisableIOAVProtection" type="bool "></param>
// <param name="DisableNetworkProtectionPerfTelemetry" type="bool "></param>
// <param name="DisablePrivacyMode" type="bool "></param>
// <param name="DisableQuicParsing" type="bool "></param>
// <param name="DisableRdpParsing" type="bool "></param>
// <param name="DisableRealtimeMonitoring" type="bool "></param>
// <param name="DisableRemovableDriveScanning" type="bool "></param>
// <param name="DisableRestorePoint" type="bool "></param>
// <param name="DisableScanningMappedNetworkDrivesForFullScan" type="bool "></param>
// <param name="DisableScanningNetworkFiles" type="bool "></param>
// <param name="DisableScriptScanning" type="bool "></param>
// <param name="DisableSmtpParsing" type="bool "></param>
// <param name="DisableSshParsing" type="bool "></param>
// <param name="DisableTDTFeature" type="bool "></param>
// <param name="DisableTlsParsing" type="bool "></param>
// <param name="EnableControlledFolderAccess" type="bool "></param>
// <param name="EnableConvertWarnToBlock" type="bool "></param>
// <param name="EnableDnsSinkhole" type="bool "></param>
// <param name="EnableEcsConfiguration" type="bool "></param>
// <param name="EnableFileHashComputation" type="bool "></param>
// <param name="EnableFullScanOnBatteryPower" type="bool "></param>
// <param name="EnableLowCpuPriority" type="bool "></param>
// <param name="EnableNetworkProtection" type="bool "></param>
// <param name="EnableUdpReceiveOffload" type="bool "></param>
// <param name="EnableUdpSegmentationOffload" type="bool "></param>
// <param name="EngineUpdatesChannel" type="bool "></param>
// <param name="ExclusionExtension" type="string []"></param>
// <param name="ExclusionIpAddress" type="string []"></param>
// <param name="ExclusionPath" type="string []"></param>
// <param name="ExclusionProcess" type="string []"></param>
// <param name="Force" type="bool "></param>
// <param name="ForceUseProxyOnly" type="bool "></param>
// <param name="HighThreatDefaultAction" type="bool "></param>
// <param name="IntelTDTEnabled" type="bool "></param>
// <param name="LowThreatDefaultAction" type="bool "></param>
// <param name="MAPSReporting" type="bool "></param>
// <param name="MeteredConnectionUpdates" type="bool "></param>
// <param name="ModerateThreatDefaultAction" type="bool "></param>
// <param name="NetworkProtectionReputationMode" type="bool "></param>
// <param name="OobeEnableRtpAndSigUpdate" type="bool "></param>
// <param name="PerformanceModeStatus" type="bool "></param>
// <param name="PlatformUpdatesChannel" type="bool "></param>
// <param name="ProxyBypass" type="bool "></param>
// <param name="ProxyPacUrl" type="bool "></param>
// <param name="ProxyServer" type="bool "></param>
// <param name="PUAProtection" type="bool "></param>
// <param name="QuarantinePurgeItemsAfterDelay" type="bool "></param>
// <param name="QuickScanIncludeExclusions" type="bool "></param>
// <param name="RandomizeScheduleTaskTimes" type="bool "></param>
// <param name="RealTimeScanDirection" type="bool "></param>
// <param name="RemediationScheduleDay" type="bool "></param>
// <param name="RemediationScheduleTime" type="bool "></param>
// <param name="RemoteEncryptionProtectionAggressiveness" type="bool "></param>
// <param name="RemoteEncryptionProtectionConfiguredState" type="bool "></param>
// <param name="RemoteEncryptionProtectionExclusions" type="string []"></param>
// <param name="RemoteEncryptionProtectionMaxBlockTime" type="bool "></param>
// <param name="RemoveScanningThreadPoolCap" type="bool "></param>
// <param name="ReportDynamicSignatureDroppedEvent" type="bool "></param>
// <param name="ReportingAdditionalActionTimeOut" type="bool "></param>
// <param name="ReportingCriticalFailureTimeOut" type="bool "></param>
// <param name="ReportingNonCriticalTimeOut" type="bool "></param>
// <param name="ScanAvgCPULoadFactor" type="bool "></param>
// <param name="ScanOnlyIfIdleEnabled" type="bool "></param>
// <param name="ScanParameters" type="bool "></param>
// <param name="ScanPurgeItemsAfterDelay" type="bool "></param>
// <param name="ScanScheduleDay" type="bool "></param>
// <param name="ScanScheduleOffset" type="bool "></param>
// <param name="ScanScheduleQuickScanTime" type="bool "></param>
// <param name="ScanScheduleTime" type="bool "></param>
// <param name="SchedulerRandomizationTime" type="bool "></param>
// <param name="ServiceHealthReportInterval" type="bool "></param>
// <param name="SevereThreatDefaultAction" type="bool "></param>
// <param name="SharedSignaturesPath" type="bool "></param>
// <param name="SharedSignaturesPathUpdateAtScheduledTimeOnly" type="bool "></param>
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
// <param name="ThrottleForScheduledScanOnly" type="bool "></param>
// <param name="TrustLabelProtectionStatus" type="bool "></param>
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
	/* IN */ RemoteEncryptionProtectionConfiguredState bool,
	/* IN */ RemoteEncryptionProtectionMaxBlockTime bool,
	/* IN */ RemoteEncryptionProtectionAggressiveness bool,
	/* IN */ RemoteEncryptionProtectionExclusions []string,
	/* IN */ BruteForceProtectionConfiguredState bool,
	/* IN */ BruteForceProtectionMaxBlockTime bool,
	/* IN */ BruteForceProtectionAggressiveness bool,
	/* IN */ BruteForceProtectionExclusions []string,
	/* IN */ BruteForceProtectionLocalNetworkBlocking bool,
	/* IN */ BruteForceProtectionSkipLearningPeriod bool,
	/* IN */ ReportingAdditionalActionTimeOut bool,
	/* IN */ ReportingCriticalFailureTimeOut bool,
	/* IN */ ReportingNonCriticalTimeOut bool,
	/* IN */ ServiceHealthReportInterval bool,
	/* IN */ ReportDynamicSignatureDroppedEvent bool,
	/* IN */ ScanAvgCPULoadFactor bool,
	/* IN */ CheckForSignaturesBeforeRunningScan bool,
	/* IN */ ScanPurgeItemsAfterDelay bool,
	/* IN */ ScanOnlyIfIdleEnabled bool,
	/* IN */ ScanParameters bool,
	/* IN */ ScanScheduleDay bool,
	/* IN */ ScanScheduleQuickScanTime bool,
	/* IN */ ScanScheduleTime bool,
	/* IN */ ThrottleForScheduledScanOnly bool,
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
	/* IN */ SchedulerRandomizationTime bool,
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
	/* IN */ ApplyDisableNetworkScanningToIOAV bool,
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
	/* IN */ SharedSignaturesPath bool,
	/* IN */ SharedSignaturesPathUpdateAtScheduledTimeOnly bool,
	/* IN */ EnableLowCpuPriority bool,
	/* IN */ EnableFileHashComputation bool,
	/* IN */ MeteredConnectionUpdates bool,
	/* IN */ AllowNetworkProtectionOnWinServer bool,
	/* IN */ DisableDatagramProcessing bool,
	/* IN */ EnableConvertWarnToBlock bool,
	/* IN */ DisableCpuThrottleOnIdleScans bool,
	/* IN */ EnableFullScanOnBatteryPower bool,
	/* IN */ ProxyPacUrl bool,
	/* IN */ ProxyServer bool,
	/* IN */ ProxyBypass bool,
	/* IN */ ForceUseProxyOnly bool,
	/* IN */ DisableTlsParsing bool,
	/* IN */ DisableHttpParsing bool,
	/* IN */ DisableDnsParsing bool,
	/* IN */ DisableDnsOverTcpParsing bool,
	/* IN */ DisableSshParsing bool,
	/* IN */ PlatformUpdatesChannel bool,
	/* IN */ EngineUpdatesChannel bool,
	/* IN */ DefinitionUpdatesChannel bool,
	/* IN */ DisableGradualRelease bool,
	/* IN */ AllowNetworkProtectionDownLevel bool,
	/* IN */ AllowDatagramProcessingOnWinServer bool,
	/* IN */ EnableDnsSinkhole bool,
	/* IN */ DisableInboundConnectionFiltering bool,
	/* IN */ DisableRdpParsing bool,
	/* IN */ DisableNetworkProtectionPerfTelemetry bool,
	/* IN */ TrustLabelProtectionStatus bool,
	/* IN */ DisableFtpParsing bool,
	/* IN */ AllowSwitchToAsyncInspection bool,
	/* IN */ ScanScheduleOffset bool,
	/* IN */ DisableTDTFeature bool,
	/* IN */ DisableSmtpParsing bool,
	/* IN */ DisableQuicParsing bool,
	/* IN */ NetworkProtectionReputationMode bool,
	/* IN */ IntelTDTEnabled bool,
	/* IN */ AttackSurfaceReductionRules_RuleSpecificExclusions_Id string,
	/* IN */ AttackSurfaceReductionRules_RuleSpecificExclusions string,
	/* IN */ OobeEnableRtpAndSigUpdate bool,
	/* IN */ PerformanceModeStatus bool,
	/* IN */ QuickScanIncludeExclusions bool,
	/* IN */ RemoveScanningThreadPoolCap bool,
	/* IN */ DisableCacheMaintenance bool,
	/* IN */ DisableCoreServiceECSIntegration bool,
	/* IN */ DisableCoreServiceTelemetry bool,
	/* IN */ EnableUdpSegmentationOffload bool,
	/* IN */ EnableUdpReceiveOffload bool,
	/* IN */ EnableEcsConfiguration bool,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove", DisableAutoExclusions, ExclusionPath, ExclusionExtension, ExclusionProcess, ExclusionIpAddress, QuarantinePurgeItemsAfterDelay, RealTimeScanDirection, RemediationScheduleDay, RemediationScheduleTime, RemoteEncryptionProtectionConfiguredState, RemoteEncryptionProtectionMaxBlockTime, RemoteEncryptionProtectionAggressiveness, RemoteEncryptionProtectionExclusions, BruteForceProtectionConfiguredState, BruteForceProtectionMaxBlockTime, BruteForceProtectionAggressiveness, BruteForceProtectionExclusions, BruteForceProtectionLocalNetworkBlocking, BruteForceProtectionSkipLearningPeriod, ReportingAdditionalActionTimeOut, ReportingCriticalFailureTimeOut, ReportingNonCriticalTimeOut, ServiceHealthReportInterval, ReportDynamicSignatureDroppedEvent, ScanAvgCPULoadFactor, CheckForSignaturesBeforeRunningScan, ScanPurgeItemsAfterDelay, ScanOnlyIfIdleEnabled, ScanParameters, ScanScheduleDay, ScanScheduleQuickScanTime, ScanScheduleTime, ThrottleForScheduledScanOnly, SignatureFirstAuGracePeriod, SignatureAuGracePeriod, SignatureDefinitionUpdateFileSharesSources, SignatureDisableUpdateOnStartupWithoutEngine, SignatureFallbackOrder, SignatureScheduleDay, SignatureScheduleTime, SignatureUpdateCatchupInterval, SignatureBlobFileSharesSources, SignatureUpdateInterval, SignatureBlobUpdateInterval, MAPSReporting, SubmitSamplesConsent, DisablePrivacyMode, RandomizeScheduleTaskTimes, SchedulerRandomizationTime, DisableBehaviorMonitoring, DisableIntrusionPreventionSystem, DisableIOAVProtection, DisableRealtimeMonitoring, DisableScriptScanning, DisableArchiveScanning, DisableCatchupFullScan, DisableCatchupQuickScan, DisableEmailScanning, DisableRemovableDriveScanning, DisableRestorePoint, DisableScanningMappedNetworkDrivesForFullScan, DisableScanningNetworkFiles, ApplyDisableNetworkScanningToIOAV, UILockdown, ThreatIDDefaultAction_Ids, ThreatIDDefaultAction_Actions, UnknownThreatDefaultAction, LowThreatDefaultAction, ModerateThreatDefaultAction, HighThreatDefaultAction, SevereThreatDefaultAction, PUAProtection, DisableBlockAtFirstSeen, CloudBlockLevel, CloudExtendedTimeout, EnableNetworkProtection, EnableControlledFolderAccess, AttackSurfaceReductionOnlyExclusions, AttackSurfaceReductionRules_Ids, AttackSurfaceReductionRules_Actions, ControlledFolderAccessAllowedApplications, ControlledFolderAccessProtectedFolders, SharedSignaturesPath, SharedSignaturesPathUpdateAtScheduledTimeOnly, EnableLowCpuPriority, EnableFileHashComputation, MeteredConnectionUpdates, AllowNetworkProtectionOnWinServer, DisableDatagramProcessing, EnableConvertWarnToBlock, DisableCpuThrottleOnIdleScans, EnableFullScanOnBatteryPower, ProxyPacUrl, ProxyServer, ProxyBypass, ForceUseProxyOnly, DisableTlsParsing, DisableHttpParsing, DisableDnsParsing, DisableDnsOverTcpParsing, DisableSshParsing, PlatformUpdatesChannel, EngineUpdatesChannel, DefinitionUpdatesChannel, DisableGradualRelease, AllowNetworkProtectionDownLevel, AllowDatagramProcessingOnWinServer, EnableDnsSinkhole, DisableInboundConnectionFiltering, DisableRdpParsing, DisableNetworkProtectionPerfTelemetry, TrustLabelProtectionStatus, DisableFtpParsing, AllowSwitchToAsyncInspection, ScanScheduleOffset, DisableTDTFeature, DisableSmtpParsing, DisableQuicParsing, NetworkProtectionReputationMode, IntelTDTEnabled, AttackSurfaceReductionRules_RuleSpecificExclusions_Id, AttackSurfaceReductionRules_RuleSpecificExclusions, OobeEnableRtpAndSigUpdate, PerformanceModeStatus, QuickScanIncludeExclusions, RemoveScanningThreadPoolCap, DisableCacheMaintenance, DisableCoreServiceECSIntegration, DisableCoreServiceTelemetry, EnableUdpSegmentationOffload, EnableUdpReceiveOffload, EnableEcsConfiguration, Force)
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
// <param name="AttackSurfaceReductionRules_RuleSpecificExclusions" type="string []"></param>
// <param name="AttackSurfaceReductionRules_RuleSpecificExclusions_Id" type="string []"></param>
// <param name="BruteForceProtectionExclusions" type="string []"></param>
// <param name="ControlledFolderAccessAllowedApplications" type="string []"></param>
// <param name="ControlledFolderAccessProtectedFolders" type="string []"></param>
// <param name="ExclusionExtension" type="string []"></param>
// <param name="ExclusionIpAddress" type="string []"></param>
// <param name="ExclusionPath" type="string []"></param>
// <param name="ExclusionProcess" type="string []"></param>
// <param name="Force" type="bool "></param>
// <param name="RemoteEncryptionProtectionExclusions" type="string []"></param>
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
	/* IN */ AttackSurfaceReductionRules_RuleSpecificExclusions_Id []string,
	/* IN */ AttackSurfaceReductionRules_RuleSpecificExclusions []string,
	/* IN */ RemoteEncryptionProtectionExclusions []string,
	/* IN */ BruteForceProtectionExclusions []string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Add", ExclusionPath, ExclusionExtension, ExclusionProcess, ExclusionIpAddress, ThreatIDDefaultAction_Ids, ThreatIDDefaultAction_Actions, AttackSurfaceReductionOnlyExclusions, AttackSurfaceReductionRules_Ids, AttackSurfaceReductionRules_Actions, ControlledFolderAccessAllowedApplications, ControlledFolderAccessProtectedFolders, SharedSignaturesPath, AttackSurfaceReductionRules_RuleSpecificExclusions_Id, AttackSurfaceReductionRules_RuleSpecificExclusions, RemoteEncryptionProtectionExclusions, BruteForceProtectionExclusions, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
