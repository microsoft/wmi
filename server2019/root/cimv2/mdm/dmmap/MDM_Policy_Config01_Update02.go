// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.mdm.dmmap
//
// ////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Config01_Update02 struct
type MDM_Policy_Config01_Update02 struct {
	*cim.WmiInstance

	//
	ActiveHoursEnd int32

	//
	ActiveHoursMaxRange int32

	//
	ActiveHoursStart int32

	//
	AllowAutoUpdate int32

	//
	AllowAutoWindowsUpdateDownloadOverMeteredNetwork int32

	//
	AllowMUUpdateService int32

	//
	AllowNonMicrosoftSignedUpdate int32

	//
	AllowTemporaryEnterpriseFeatureControl int32

	//
	AllowUpdateService int32

	//
	AutomaticMaintenanceWakeUp int32

	//
	AutoRestartDeadlinePeriodInDays int32

	//
	AutoRestartDeadlinePeriodInDaysForFeatureUpdates int32

	//
	AutoRestartNotificationSchedule int32

	//
	AutoRestartRequiredNotificationDismissal int32

	//
	BranchReadinessLevel int32

	//
	ConfigureDeadlineForFeatureUpdates int32

	//
	ConfigureDeadlineForQualityUpdates int32

	//
	ConfigureDeadlineGracePeriod int32

	//
	ConfigureDeadlineGracePeriodForFeatureUpdates int32

	//
	ConfigureDeadlineNoAutoReboot int32

	//
	ConfigureFeatureUpdateUninstallPeriod int32

	//
	DeferFeatureUpdatesPeriodInDays int32

	//
	DeferQualityUpdatesPeriodInDays int32

	//
	DeferUpdatePeriod int32

	//
	DeferUpgradePeriod int32

	//
	DetectionFrequency int32

	//
	DisableDualScan int32

	//
	DisableWUfBSafeguards int32

	//
	DoNotEnforceEnterpriseTLSCertPinningForUpdateDetection int32

	//
	EngagedRestartDeadline int32

	//
	EngagedRestartDeadlineForFeatureUpdates int32

	//
	EngagedRestartSnoozeSchedule int32

	//
	EngagedRestartSnoozeScheduleForFeatureUpdates int32

	//
	EngagedRestartTransitionSchedule int32

	//
	EngagedRestartTransitionScheduleForFeatureUpdates int32

	//
	ExcludeWUDriversInQualityUpdate int32

	//
	FillEmptyContentUrls int32

	//
	IgnoreMOAppDownloadLimit int32

	//
	IgnoreMOUpdateDownloadLimit int32

	//
	InstanceID string

	//
	ManagePreviewBuilds int32

	//
	ParentID string

	//
	PauseDeferrals int32

	//
	PauseFeatureUpdates int32

	//
	PauseFeatureUpdatesStartTime string

	//
	PauseQualityUpdates int32

	//
	PauseQualityUpdatesStartTime string

	//
	PhoneUpdateRestrictions int32

	//
	ProductVersion string

	//
	RequireDeferUpgrade int32

	//
	RequireUpdateApproval int32

	//
	ScheduledInstallDay int32

	//
	ScheduledInstallEveryWeek int32

	//
	ScheduledInstallFirstWeek int32

	//
	ScheduledInstallFourthWeek int32

	//
	ScheduledInstallSecondWeek int32

	//
	ScheduledInstallThirdWeek int32

	//
	ScheduledInstallTime int32

	//
	ScheduleImminentRestartWarning int32

	//
	ScheduleRestartWarning int32

	//
	SetAutoRestartNotificationDisable int32

	//
	SetDisablePauseUXAccess int32

	//
	SetDisableUXWUAccess int32

	//
	SetEDURestart int32

	//
	SetPolicyDrivenUpdateSourceForDriverUpdates int32

	//
	SetPolicyDrivenUpdateSourceForFeatureUpdates int32

	//
	SetPolicyDrivenUpdateSourceForOtherUpdates int32

	//
	SetPolicyDrivenUpdateSourceForQualityUpdates int32

	//
	SetProxyBehaviorForUpdateDetection int32

	//
	TargetReleaseVersion string

	//
	UpdateNotificationLevel int32

	//
	UpdateServiceUrl string

	//
	UpdateServiceUrlAlternate string
}

func NewMDM_Policy_Config01_Update02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_Update02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Update02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_Update02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_Update02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Update02{
		WmiInstance: tmp,
	}
	return
}

// SetActiveHoursEnd sets the value of ActiveHoursEnd for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyActiveHoursEnd(value int32) (err error) {
	return instance.SetProperty("ActiveHoursEnd", (value))
}

// GetActiveHoursEnd gets the value of ActiveHoursEnd for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyActiveHoursEnd() (value int32, err error) {
	retValue, err := instance.GetProperty("ActiveHoursEnd")
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

	value = int32(valuetmp)

	return
}

// SetActiveHoursMaxRange sets the value of ActiveHoursMaxRange for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyActiveHoursMaxRange(value int32) (err error) {
	return instance.SetProperty("ActiveHoursMaxRange", (value))
}

// GetActiveHoursMaxRange gets the value of ActiveHoursMaxRange for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyActiveHoursMaxRange() (value int32, err error) {
	retValue, err := instance.GetProperty("ActiveHoursMaxRange")
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

	value = int32(valuetmp)

	return
}

// SetActiveHoursStart sets the value of ActiveHoursStart for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyActiveHoursStart(value int32) (err error) {
	return instance.SetProperty("ActiveHoursStart", (value))
}

// GetActiveHoursStart gets the value of ActiveHoursStart for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyActiveHoursStart() (value int32, err error) {
	retValue, err := instance.GetProperty("ActiveHoursStart")
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

	value = int32(valuetmp)

	return
}

// SetAllowAutoUpdate sets the value of AllowAutoUpdate for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyAllowAutoUpdate(value int32) (err error) {
	return instance.SetProperty("AllowAutoUpdate", (value))
}

// GetAllowAutoUpdate gets the value of AllowAutoUpdate for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyAllowAutoUpdate() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAutoUpdate")
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

	value = int32(valuetmp)

	return
}

// SetAllowAutoWindowsUpdateDownloadOverMeteredNetwork sets the value of AllowAutoWindowsUpdateDownloadOverMeteredNetwork for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyAllowAutoWindowsUpdateDownloadOverMeteredNetwork(value int32) (err error) {
	return instance.SetProperty("AllowAutoWindowsUpdateDownloadOverMeteredNetwork", (value))
}

// GetAllowAutoWindowsUpdateDownloadOverMeteredNetwork gets the value of AllowAutoWindowsUpdateDownloadOverMeteredNetwork for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyAllowAutoWindowsUpdateDownloadOverMeteredNetwork() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAutoWindowsUpdateDownloadOverMeteredNetwork")
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

	value = int32(valuetmp)

	return
}

// SetAllowMUUpdateService sets the value of AllowMUUpdateService for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyAllowMUUpdateService(value int32) (err error) {
	return instance.SetProperty("AllowMUUpdateService", (value))
}

// GetAllowMUUpdateService gets the value of AllowMUUpdateService for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyAllowMUUpdateService() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowMUUpdateService")
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

	value = int32(valuetmp)

	return
}

// SetAllowNonMicrosoftSignedUpdate sets the value of AllowNonMicrosoftSignedUpdate for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyAllowNonMicrosoftSignedUpdate(value int32) (err error) {
	return instance.SetProperty("AllowNonMicrosoftSignedUpdate", (value))
}

// GetAllowNonMicrosoftSignedUpdate gets the value of AllowNonMicrosoftSignedUpdate for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyAllowNonMicrosoftSignedUpdate() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowNonMicrosoftSignedUpdate")
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

	value = int32(valuetmp)

	return
}

// SetAllowTemporaryEnterpriseFeatureControl sets the value of AllowTemporaryEnterpriseFeatureControl for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyAllowTemporaryEnterpriseFeatureControl(value int32) (err error) {
	return instance.SetProperty("AllowTemporaryEnterpriseFeatureControl", (value))
}

// GetAllowTemporaryEnterpriseFeatureControl gets the value of AllowTemporaryEnterpriseFeatureControl for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyAllowTemporaryEnterpriseFeatureControl() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowTemporaryEnterpriseFeatureControl")
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

	value = int32(valuetmp)

	return
}

// SetAllowUpdateService sets the value of AllowUpdateService for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyAllowUpdateService(value int32) (err error) {
	return instance.SetProperty("AllowUpdateService", (value))
}

// GetAllowUpdateService gets the value of AllowUpdateService for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyAllowUpdateService() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowUpdateService")
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

	value = int32(valuetmp)

	return
}

// SetAutomaticMaintenanceWakeUp sets the value of AutomaticMaintenanceWakeUp for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyAutomaticMaintenanceWakeUp(value int32) (err error) {
	return instance.SetProperty("AutomaticMaintenanceWakeUp", (value))
}

// GetAutomaticMaintenanceWakeUp gets the value of AutomaticMaintenanceWakeUp for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyAutomaticMaintenanceWakeUp() (value int32, err error) {
	retValue, err := instance.GetProperty("AutomaticMaintenanceWakeUp")
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

	value = int32(valuetmp)

	return
}

// SetAutoRestartDeadlinePeriodInDays sets the value of AutoRestartDeadlinePeriodInDays for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyAutoRestartDeadlinePeriodInDays(value int32) (err error) {
	return instance.SetProperty("AutoRestartDeadlinePeriodInDays", (value))
}

// GetAutoRestartDeadlinePeriodInDays gets the value of AutoRestartDeadlinePeriodInDays for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyAutoRestartDeadlinePeriodInDays() (value int32, err error) {
	retValue, err := instance.GetProperty("AutoRestartDeadlinePeriodInDays")
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

	value = int32(valuetmp)

	return
}

// SetAutoRestartDeadlinePeriodInDaysForFeatureUpdates sets the value of AutoRestartDeadlinePeriodInDaysForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyAutoRestartDeadlinePeriodInDaysForFeatureUpdates(value int32) (err error) {
	return instance.SetProperty("AutoRestartDeadlinePeriodInDaysForFeatureUpdates", (value))
}

// GetAutoRestartDeadlinePeriodInDaysForFeatureUpdates gets the value of AutoRestartDeadlinePeriodInDaysForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyAutoRestartDeadlinePeriodInDaysForFeatureUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("AutoRestartDeadlinePeriodInDaysForFeatureUpdates")
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

	value = int32(valuetmp)

	return
}

// SetAutoRestartNotificationSchedule sets the value of AutoRestartNotificationSchedule for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyAutoRestartNotificationSchedule(value int32) (err error) {
	return instance.SetProperty("AutoRestartNotificationSchedule", (value))
}

// GetAutoRestartNotificationSchedule gets the value of AutoRestartNotificationSchedule for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyAutoRestartNotificationSchedule() (value int32, err error) {
	retValue, err := instance.GetProperty("AutoRestartNotificationSchedule")
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

	value = int32(valuetmp)

	return
}

// SetAutoRestartRequiredNotificationDismissal sets the value of AutoRestartRequiredNotificationDismissal for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyAutoRestartRequiredNotificationDismissal(value int32) (err error) {
	return instance.SetProperty("AutoRestartRequiredNotificationDismissal", (value))
}

// GetAutoRestartRequiredNotificationDismissal gets the value of AutoRestartRequiredNotificationDismissal for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyAutoRestartRequiredNotificationDismissal() (value int32, err error) {
	retValue, err := instance.GetProperty("AutoRestartRequiredNotificationDismissal")
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

	value = int32(valuetmp)

	return
}

// SetBranchReadinessLevel sets the value of BranchReadinessLevel for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyBranchReadinessLevel(value int32) (err error) {
	return instance.SetProperty("BranchReadinessLevel", (value))
}

// GetBranchReadinessLevel gets the value of BranchReadinessLevel for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyBranchReadinessLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("BranchReadinessLevel")
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

	value = int32(valuetmp)

	return
}

// SetConfigureDeadlineForFeatureUpdates sets the value of ConfigureDeadlineForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyConfigureDeadlineForFeatureUpdates(value int32) (err error) {
	return instance.SetProperty("ConfigureDeadlineForFeatureUpdates", (value))
}

// GetConfigureDeadlineForFeatureUpdates gets the value of ConfigureDeadlineForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyConfigureDeadlineForFeatureUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureDeadlineForFeatureUpdates")
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

	value = int32(valuetmp)

	return
}

// SetConfigureDeadlineForQualityUpdates sets the value of ConfigureDeadlineForQualityUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyConfigureDeadlineForQualityUpdates(value int32) (err error) {
	return instance.SetProperty("ConfigureDeadlineForQualityUpdates", (value))
}

// GetConfigureDeadlineForQualityUpdates gets the value of ConfigureDeadlineForQualityUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyConfigureDeadlineForQualityUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureDeadlineForQualityUpdates")
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

	value = int32(valuetmp)

	return
}

// SetConfigureDeadlineGracePeriod sets the value of ConfigureDeadlineGracePeriod for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyConfigureDeadlineGracePeriod(value int32) (err error) {
	return instance.SetProperty("ConfigureDeadlineGracePeriod", (value))
}

// GetConfigureDeadlineGracePeriod gets the value of ConfigureDeadlineGracePeriod for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyConfigureDeadlineGracePeriod() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureDeadlineGracePeriod")
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

	value = int32(valuetmp)

	return
}

// SetConfigureDeadlineGracePeriodForFeatureUpdates sets the value of ConfigureDeadlineGracePeriodForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyConfigureDeadlineGracePeriodForFeatureUpdates(value int32) (err error) {
	return instance.SetProperty("ConfigureDeadlineGracePeriodForFeatureUpdates", (value))
}

// GetConfigureDeadlineGracePeriodForFeatureUpdates gets the value of ConfigureDeadlineGracePeriodForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyConfigureDeadlineGracePeriodForFeatureUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureDeadlineGracePeriodForFeatureUpdates")
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

	value = int32(valuetmp)

	return
}

// SetConfigureDeadlineNoAutoReboot sets the value of ConfigureDeadlineNoAutoReboot for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyConfigureDeadlineNoAutoReboot(value int32) (err error) {
	return instance.SetProperty("ConfigureDeadlineNoAutoReboot", (value))
}

// GetConfigureDeadlineNoAutoReboot gets the value of ConfigureDeadlineNoAutoReboot for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyConfigureDeadlineNoAutoReboot() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureDeadlineNoAutoReboot")
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

	value = int32(valuetmp)

	return
}

// SetConfigureFeatureUpdateUninstallPeriod sets the value of ConfigureFeatureUpdateUninstallPeriod for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyConfigureFeatureUpdateUninstallPeriod(value int32) (err error) {
	return instance.SetProperty("ConfigureFeatureUpdateUninstallPeriod", (value))
}

// GetConfigureFeatureUpdateUninstallPeriod gets the value of ConfigureFeatureUpdateUninstallPeriod for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyConfigureFeatureUpdateUninstallPeriod() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureFeatureUpdateUninstallPeriod")
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

	value = int32(valuetmp)

	return
}

// SetDeferFeatureUpdatesPeriodInDays sets the value of DeferFeatureUpdatesPeriodInDays for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyDeferFeatureUpdatesPeriodInDays(value int32) (err error) {
	return instance.SetProperty("DeferFeatureUpdatesPeriodInDays", (value))
}

// GetDeferFeatureUpdatesPeriodInDays gets the value of DeferFeatureUpdatesPeriodInDays for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyDeferFeatureUpdatesPeriodInDays() (value int32, err error) {
	retValue, err := instance.GetProperty("DeferFeatureUpdatesPeriodInDays")
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

	value = int32(valuetmp)

	return
}

// SetDeferQualityUpdatesPeriodInDays sets the value of DeferQualityUpdatesPeriodInDays for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyDeferQualityUpdatesPeriodInDays(value int32) (err error) {
	return instance.SetProperty("DeferQualityUpdatesPeriodInDays", (value))
}

// GetDeferQualityUpdatesPeriodInDays gets the value of DeferQualityUpdatesPeriodInDays for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyDeferQualityUpdatesPeriodInDays() (value int32, err error) {
	retValue, err := instance.GetProperty("DeferQualityUpdatesPeriodInDays")
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

	value = int32(valuetmp)

	return
}

// SetDeferUpdatePeriod sets the value of DeferUpdatePeriod for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyDeferUpdatePeriod(value int32) (err error) {
	return instance.SetProperty("DeferUpdatePeriod", (value))
}

// GetDeferUpdatePeriod gets the value of DeferUpdatePeriod for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyDeferUpdatePeriod() (value int32, err error) {
	retValue, err := instance.GetProperty("DeferUpdatePeriod")
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

	value = int32(valuetmp)

	return
}

// SetDeferUpgradePeriod sets the value of DeferUpgradePeriod for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyDeferUpgradePeriod(value int32) (err error) {
	return instance.SetProperty("DeferUpgradePeriod", (value))
}

// GetDeferUpgradePeriod gets the value of DeferUpgradePeriod for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyDeferUpgradePeriod() (value int32, err error) {
	retValue, err := instance.GetProperty("DeferUpgradePeriod")
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

	value = int32(valuetmp)

	return
}

// SetDetectionFrequency sets the value of DetectionFrequency for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyDetectionFrequency(value int32) (err error) {
	return instance.SetProperty("DetectionFrequency", (value))
}

// GetDetectionFrequency gets the value of DetectionFrequency for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyDetectionFrequency() (value int32, err error) {
	retValue, err := instance.GetProperty("DetectionFrequency")
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

	value = int32(valuetmp)

	return
}

// SetDisableDualScan sets the value of DisableDualScan for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyDisableDualScan(value int32) (err error) {
	return instance.SetProperty("DisableDualScan", (value))
}

// GetDisableDualScan gets the value of DisableDualScan for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyDisableDualScan() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableDualScan")
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

	value = int32(valuetmp)

	return
}

// SetDisableWUfBSafeguards sets the value of DisableWUfBSafeguards for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyDisableWUfBSafeguards(value int32) (err error) {
	return instance.SetProperty("DisableWUfBSafeguards", (value))
}

// GetDisableWUfBSafeguards gets the value of DisableWUfBSafeguards for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyDisableWUfBSafeguards() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableWUfBSafeguards")
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

	value = int32(valuetmp)

	return
}

// SetDoNotEnforceEnterpriseTLSCertPinningForUpdateDetection sets the value of DoNotEnforceEnterpriseTLSCertPinningForUpdateDetection for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyDoNotEnforceEnterpriseTLSCertPinningForUpdateDetection(value int32) (err error) {
	return instance.SetProperty("DoNotEnforceEnterpriseTLSCertPinningForUpdateDetection", (value))
}

// GetDoNotEnforceEnterpriseTLSCertPinningForUpdateDetection gets the value of DoNotEnforceEnterpriseTLSCertPinningForUpdateDetection for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyDoNotEnforceEnterpriseTLSCertPinningForUpdateDetection() (value int32, err error) {
	retValue, err := instance.GetProperty("DoNotEnforceEnterpriseTLSCertPinningForUpdateDetection")
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

	value = int32(valuetmp)

	return
}

// SetEngagedRestartDeadline sets the value of EngagedRestartDeadline for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyEngagedRestartDeadline(value int32) (err error) {
	return instance.SetProperty("EngagedRestartDeadline", (value))
}

// GetEngagedRestartDeadline gets the value of EngagedRestartDeadline for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyEngagedRestartDeadline() (value int32, err error) {
	retValue, err := instance.GetProperty("EngagedRestartDeadline")
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

	value = int32(valuetmp)

	return
}

// SetEngagedRestartDeadlineForFeatureUpdates sets the value of EngagedRestartDeadlineForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyEngagedRestartDeadlineForFeatureUpdates(value int32) (err error) {
	return instance.SetProperty("EngagedRestartDeadlineForFeatureUpdates", (value))
}

// GetEngagedRestartDeadlineForFeatureUpdates gets the value of EngagedRestartDeadlineForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyEngagedRestartDeadlineForFeatureUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("EngagedRestartDeadlineForFeatureUpdates")
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

	value = int32(valuetmp)

	return
}

// SetEngagedRestartSnoozeSchedule sets the value of EngagedRestartSnoozeSchedule for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyEngagedRestartSnoozeSchedule(value int32) (err error) {
	return instance.SetProperty("EngagedRestartSnoozeSchedule", (value))
}

// GetEngagedRestartSnoozeSchedule gets the value of EngagedRestartSnoozeSchedule for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyEngagedRestartSnoozeSchedule() (value int32, err error) {
	retValue, err := instance.GetProperty("EngagedRestartSnoozeSchedule")
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

	value = int32(valuetmp)

	return
}

// SetEngagedRestartSnoozeScheduleForFeatureUpdates sets the value of EngagedRestartSnoozeScheduleForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyEngagedRestartSnoozeScheduleForFeatureUpdates(value int32) (err error) {
	return instance.SetProperty("EngagedRestartSnoozeScheduleForFeatureUpdates", (value))
}

// GetEngagedRestartSnoozeScheduleForFeatureUpdates gets the value of EngagedRestartSnoozeScheduleForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyEngagedRestartSnoozeScheduleForFeatureUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("EngagedRestartSnoozeScheduleForFeatureUpdates")
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

	value = int32(valuetmp)

	return
}

// SetEngagedRestartTransitionSchedule sets the value of EngagedRestartTransitionSchedule for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyEngagedRestartTransitionSchedule(value int32) (err error) {
	return instance.SetProperty("EngagedRestartTransitionSchedule", (value))
}

// GetEngagedRestartTransitionSchedule gets the value of EngagedRestartTransitionSchedule for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyEngagedRestartTransitionSchedule() (value int32, err error) {
	retValue, err := instance.GetProperty("EngagedRestartTransitionSchedule")
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

	value = int32(valuetmp)

	return
}

// SetEngagedRestartTransitionScheduleForFeatureUpdates sets the value of EngagedRestartTransitionScheduleForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyEngagedRestartTransitionScheduleForFeatureUpdates(value int32) (err error) {
	return instance.SetProperty("EngagedRestartTransitionScheduleForFeatureUpdates", (value))
}

// GetEngagedRestartTransitionScheduleForFeatureUpdates gets the value of EngagedRestartTransitionScheduleForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyEngagedRestartTransitionScheduleForFeatureUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("EngagedRestartTransitionScheduleForFeatureUpdates")
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

	value = int32(valuetmp)

	return
}

// SetExcludeWUDriversInQualityUpdate sets the value of ExcludeWUDriversInQualityUpdate for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyExcludeWUDriversInQualityUpdate(value int32) (err error) {
	return instance.SetProperty("ExcludeWUDriversInQualityUpdate", (value))
}

// GetExcludeWUDriversInQualityUpdate gets the value of ExcludeWUDriversInQualityUpdate for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyExcludeWUDriversInQualityUpdate() (value int32, err error) {
	retValue, err := instance.GetProperty("ExcludeWUDriversInQualityUpdate")
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

	value = int32(valuetmp)

	return
}

// SetFillEmptyContentUrls sets the value of FillEmptyContentUrls for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyFillEmptyContentUrls(value int32) (err error) {
	return instance.SetProperty("FillEmptyContentUrls", (value))
}

// GetFillEmptyContentUrls gets the value of FillEmptyContentUrls for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyFillEmptyContentUrls() (value int32, err error) {
	retValue, err := instance.GetProperty("FillEmptyContentUrls")
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

	value = int32(valuetmp)

	return
}

// SetIgnoreMOAppDownloadLimit sets the value of IgnoreMOAppDownloadLimit for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyIgnoreMOAppDownloadLimit(value int32) (err error) {
	return instance.SetProperty("IgnoreMOAppDownloadLimit", (value))
}

// GetIgnoreMOAppDownloadLimit gets the value of IgnoreMOAppDownloadLimit for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyIgnoreMOAppDownloadLimit() (value int32, err error) {
	retValue, err := instance.GetProperty("IgnoreMOAppDownloadLimit")
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

	value = int32(valuetmp)

	return
}

// SetIgnoreMOUpdateDownloadLimit sets the value of IgnoreMOUpdateDownloadLimit for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyIgnoreMOUpdateDownloadLimit(value int32) (err error) {
	return instance.SetProperty("IgnoreMOUpdateDownloadLimit", (value))
}

// GetIgnoreMOUpdateDownloadLimit gets the value of IgnoreMOUpdateDownloadLimit for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyIgnoreMOUpdateDownloadLimit() (value int32, err error) {
	retValue, err := instance.GetProperty("IgnoreMOUpdateDownloadLimit")
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

	value = int32(valuetmp)

	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetManagePreviewBuilds sets the value of ManagePreviewBuilds for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyManagePreviewBuilds(value int32) (err error) {
	return instance.SetProperty("ManagePreviewBuilds", (value))
}

// GetManagePreviewBuilds gets the value of ManagePreviewBuilds for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyManagePreviewBuilds() (value int32, err error) {
	retValue, err := instance.GetProperty("ManagePreviewBuilds")
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

	value = int32(valuetmp)

	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetPauseDeferrals sets the value of PauseDeferrals for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyPauseDeferrals(value int32) (err error) {
	return instance.SetProperty("PauseDeferrals", (value))
}

// GetPauseDeferrals gets the value of PauseDeferrals for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyPauseDeferrals() (value int32, err error) {
	retValue, err := instance.GetProperty("PauseDeferrals")
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

	value = int32(valuetmp)

	return
}

// SetPauseFeatureUpdates sets the value of PauseFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyPauseFeatureUpdates(value int32) (err error) {
	return instance.SetProperty("PauseFeatureUpdates", (value))
}

// GetPauseFeatureUpdates gets the value of PauseFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyPauseFeatureUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("PauseFeatureUpdates")
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

	value = int32(valuetmp)

	return
}

// SetPauseFeatureUpdatesStartTime sets the value of PauseFeatureUpdatesStartTime for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyPauseFeatureUpdatesStartTime(value string) (err error) {
	return instance.SetProperty("PauseFeatureUpdatesStartTime", (value))
}

// GetPauseFeatureUpdatesStartTime gets the value of PauseFeatureUpdatesStartTime for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyPauseFeatureUpdatesStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("PauseFeatureUpdatesStartTime")
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

// SetPauseQualityUpdates sets the value of PauseQualityUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyPauseQualityUpdates(value int32) (err error) {
	return instance.SetProperty("PauseQualityUpdates", (value))
}

// GetPauseQualityUpdates gets the value of PauseQualityUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyPauseQualityUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("PauseQualityUpdates")
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

	value = int32(valuetmp)

	return
}

// SetPauseQualityUpdatesStartTime sets the value of PauseQualityUpdatesStartTime for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyPauseQualityUpdatesStartTime(value string) (err error) {
	return instance.SetProperty("PauseQualityUpdatesStartTime", (value))
}

// GetPauseQualityUpdatesStartTime gets the value of PauseQualityUpdatesStartTime for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyPauseQualityUpdatesStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("PauseQualityUpdatesStartTime")
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

// SetPhoneUpdateRestrictions sets the value of PhoneUpdateRestrictions for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyPhoneUpdateRestrictions(value int32) (err error) {
	return instance.SetProperty("PhoneUpdateRestrictions", (value))
}

// GetPhoneUpdateRestrictions gets the value of PhoneUpdateRestrictions for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyPhoneUpdateRestrictions() (value int32, err error) {
	retValue, err := instance.GetProperty("PhoneUpdateRestrictions")
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

	value = int32(valuetmp)

	return
}

// SetProductVersion sets the value of ProductVersion for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyProductVersion(value string) (err error) {
	return instance.SetProperty("ProductVersion", (value))
}

// GetProductVersion gets the value of ProductVersion for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyProductVersion() (value string, err error) {
	retValue, err := instance.GetProperty("ProductVersion")
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

// SetRequireDeferUpgrade sets the value of RequireDeferUpgrade for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyRequireDeferUpgrade(value int32) (err error) {
	return instance.SetProperty("RequireDeferUpgrade", (value))
}

// GetRequireDeferUpgrade gets the value of RequireDeferUpgrade for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyRequireDeferUpgrade() (value int32, err error) {
	retValue, err := instance.GetProperty("RequireDeferUpgrade")
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

	value = int32(valuetmp)

	return
}

// SetRequireUpdateApproval sets the value of RequireUpdateApproval for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyRequireUpdateApproval(value int32) (err error) {
	return instance.SetProperty("RequireUpdateApproval", (value))
}

// GetRequireUpdateApproval gets the value of RequireUpdateApproval for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyRequireUpdateApproval() (value int32, err error) {
	retValue, err := instance.GetProperty("RequireUpdateApproval")
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

	value = int32(valuetmp)

	return
}

// SetScheduledInstallDay sets the value of ScheduledInstallDay for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyScheduledInstallDay(value int32) (err error) {
	return instance.SetProperty("ScheduledInstallDay", (value))
}

// GetScheduledInstallDay gets the value of ScheduledInstallDay for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyScheduledInstallDay() (value int32, err error) {
	retValue, err := instance.GetProperty("ScheduledInstallDay")
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

	value = int32(valuetmp)

	return
}

// SetScheduledInstallEveryWeek sets the value of ScheduledInstallEveryWeek for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyScheduledInstallEveryWeek(value int32) (err error) {
	return instance.SetProperty("ScheduledInstallEveryWeek", (value))
}

// GetScheduledInstallEveryWeek gets the value of ScheduledInstallEveryWeek for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyScheduledInstallEveryWeek() (value int32, err error) {
	retValue, err := instance.GetProperty("ScheduledInstallEveryWeek")
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

	value = int32(valuetmp)

	return
}

// SetScheduledInstallFirstWeek sets the value of ScheduledInstallFirstWeek for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyScheduledInstallFirstWeek(value int32) (err error) {
	return instance.SetProperty("ScheduledInstallFirstWeek", (value))
}

// GetScheduledInstallFirstWeek gets the value of ScheduledInstallFirstWeek for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyScheduledInstallFirstWeek() (value int32, err error) {
	retValue, err := instance.GetProperty("ScheduledInstallFirstWeek")
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

	value = int32(valuetmp)

	return
}

// SetScheduledInstallFourthWeek sets the value of ScheduledInstallFourthWeek for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyScheduledInstallFourthWeek(value int32) (err error) {
	return instance.SetProperty("ScheduledInstallFourthWeek", (value))
}

// GetScheduledInstallFourthWeek gets the value of ScheduledInstallFourthWeek for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyScheduledInstallFourthWeek() (value int32, err error) {
	retValue, err := instance.GetProperty("ScheduledInstallFourthWeek")
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

	value = int32(valuetmp)

	return
}

// SetScheduledInstallSecondWeek sets the value of ScheduledInstallSecondWeek for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyScheduledInstallSecondWeek(value int32) (err error) {
	return instance.SetProperty("ScheduledInstallSecondWeek", (value))
}

// GetScheduledInstallSecondWeek gets the value of ScheduledInstallSecondWeek for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyScheduledInstallSecondWeek() (value int32, err error) {
	retValue, err := instance.GetProperty("ScheduledInstallSecondWeek")
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

	value = int32(valuetmp)

	return
}

// SetScheduledInstallThirdWeek sets the value of ScheduledInstallThirdWeek for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyScheduledInstallThirdWeek(value int32) (err error) {
	return instance.SetProperty("ScheduledInstallThirdWeek", (value))
}

// GetScheduledInstallThirdWeek gets the value of ScheduledInstallThirdWeek for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyScheduledInstallThirdWeek() (value int32, err error) {
	retValue, err := instance.GetProperty("ScheduledInstallThirdWeek")
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

	value = int32(valuetmp)

	return
}

// SetScheduledInstallTime sets the value of ScheduledInstallTime for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyScheduledInstallTime(value int32) (err error) {
	return instance.SetProperty("ScheduledInstallTime", (value))
}

// GetScheduledInstallTime gets the value of ScheduledInstallTime for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyScheduledInstallTime() (value int32, err error) {
	retValue, err := instance.GetProperty("ScheduledInstallTime")
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

	value = int32(valuetmp)

	return
}

// SetScheduleImminentRestartWarning sets the value of ScheduleImminentRestartWarning for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyScheduleImminentRestartWarning(value int32) (err error) {
	return instance.SetProperty("ScheduleImminentRestartWarning", (value))
}

// GetScheduleImminentRestartWarning gets the value of ScheduleImminentRestartWarning for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyScheduleImminentRestartWarning() (value int32, err error) {
	retValue, err := instance.GetProperty("ScheduleImminentRestartWarning")
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

	value = int32(valuetmp)

	return
}

// SetScheduleRestartWarning sets the value of ScheduleRestartWarning for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyScheduleRestartWarning(value int32) (err error) {
	return instance.SetProperty("ScheduleRestartWarning", (value))
}

// GetScheduleRestartWarning gets the value of ScheduleRestartWarning for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyScheduleRestartWarning() (value int32, err error) {
	retValue, err := instance.GetProperty("ScheduleRestartWarning")
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

	value = int32(valuetmp)

	return
}

// SetSetAutoRestartNotificationDisable sets the value of SetAutoRestartNotificationDisable for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertySetAutoRestartNotificationDisable(value int32) (err error) {
	return instance.SetProperty("SetAutoRestartNotificationDisable", (value))
}

// GetSetAutoRestartNotificationDisable gets the value of SetAutoRestartNotificationDisable for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertySetAutoRestartNotificationDisable() (value int32, err error) {
	retValue, err := instance.GetProperty("SetAutoRestartNotificationDisable")
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

	value = int32(valuetmp)

	return
}

// SetSetDisablePauseUXAccess sets the value of SetDisablePauseUXAccess for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertySetDisablePauseUXAccess(value int32) (err error) {
	return instance.SetProperty("SetDisablePauseUXAccess", (value))
}

// GetSetDisablePauseUXAccess gets the value of SetDisablePauseUXAccess for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertySetDisablePauseUXAccess() (value int32, err error) {
	retValue, err := instance.GetProperty("SetDisablePauseUXAccess")
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

	value = int32(valuetmp)

	return
}

// SetSetDisableUXWUAccess sets the value of SetDisableUXWUAccess for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertySetDisableUXWUAccess(value int32) (err error) {
	return instance.SetProperty("SetDisableUXWUAccess", (value))
}

// GetSetDisableUXWUAccess gets the value of SetDisableUXWUAccess for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertySetDisableUXWUAccess() (value int32, err error) {
	retValue, err := instance.GetProperty("SetDisableUXWUAccess")
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

	value = int32(valuetmp)

	return
}

// SetSetEDURestart sets the value of SetEDURestart for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertySetEDURestart(value int32) (err error) {
	return instance.SetProperty("SetEDURestart", (value))
}

// GetSetEDURestart gets the value of SetEDURestart for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertySetEDURestart() (value int32, err error) {
	retValue, err := instance.GetProperty("SetEDURestart")
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

	value = int32(valuetmp)

	return
}

// SetSetPolicyDrivenUpdateSourceForDriverUpdates sets the value of SetPolicyDrivenUpdateSourceForDriverUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertySetPolicyDrivenUpdateSourceForDriverUpdates(value int32) (err error) {
	return instance.SetProperty("SetPolicyDrivenUpdateSourceForDriverUpdates", (value))
}

// GetSetPolicyDrivenUpdateSourceForDriverUpdates gets the value of SetPolicyDrivenUpdateSourceForDriverUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertySetPolicyDrivenUpdateSourceForDriverUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("SetPolicyDrivenUpdateSourceForDriverUpdates")
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

	value = int32(valuetmp)

	return
}

// SetSetPolicyDrivenUpdateSourceForFeatureUpdates sets the value of SetPolicyDrivenUpdateSourceForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertySetPolicyDrivenUpdateSourceForFeatureUpdates(value int32) (err error) {
	return instance.SetProperty("SetPolicyDrivenUpdateSourceForFeatureUpdates", (value))
}

// GetSetPolicyDrivenUpdateSourceForFeatureUpdates gets the value of SetPolicyDrivenUpdateSourceForFeatureUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertySetPolicyDrivenUpdateSourceForFeatureUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("SetPolicyDrivenUpdateSourceForFeatureUpdates")
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

	value = int32(valuetmp)

	return
}

// SetSetPolicyDrivenUpdateSourceForOtherUpdates sets the value of SetPolicyDrivenUpdateSourceForOtherUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertySetPolicyDrivenUpdateSourceForOtherUpdates(value int32) (err error) {
	return instance.SetProperty("SetPolicyDrivenUpdateSourceForOtherUpdates", (value))
}

// GetSetPolicyDrivenUpdateSourceForOtherUpdates gets the value of SetPolicyDrivenUpdateSourceForOtherUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertySetPolicyDrivenUpdateSourceForOtherUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("SetPolicyDrivenUpdateSourceForOtherUpdates")
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

	value = int32(valuetmp)

	return
}

// SetSetPolicyDrivenUpdateSourceForQualityUpdates sets the value of SetPolicyDrivenUpdateSourceForQualityUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertySetPolicyDrivenUpdateSourceForQualityUpdates(value int32) (err error) {
	return instance.SetProperty("SetPolicyDrivenUpdateSourceForQualityUpdates", (value))
}

// GetSetPolicyDrivenUpdateSourceForQualityUpdates gets the value of SetPolicyDrivenUpdateSourceForQualityUpdates for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertySetPolicyDrivenUpdateSourceForQualityUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("SetPolicyDrivenUpdateSourceForQualityUpdates")
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

	value = int32(valuetmp)

	return
}

// SetSetProxyBehaviorForUpdateDetection sets the value of SetProxyBehaviorForUpdateDetection for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertySetProxyBehaviorForUpdateDetection(value int32) (err error) {
	return instance.SetProperty("SetProxyBehaviorForUpdateDetection", (value))
}

// GetSetProxyBehaviorForUpdateDetection gets the value of SetProxyBehaviorForUpdateDetection for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertySetProxyBehaviorForUpdateDetection() (value int32, err error) {
	retValue, err := instance.GetProperty("SetProxyBehaviorForUpdateDetection")
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

	value = int32(valuetmp)

	return
}

// SetTargetReleaseVersion sets the value of TargetReleaseVersion for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyTargetReleaseVersion(value string) (err error) {
	return instance.SetProperty("TargetReleaseVersion", (value))
}

// GetTargetReleaseVersion gets the value of TargetReleaseVersion for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyTargetReleaseVersion() (value string, err error) {
	retValue, err := instance.GetProperty("TargetReleaseVersion")
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

// SetUpdateNotificationLevel sets the value of UpdateNotificationLevel for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyUpdateNotificationLevel(value int32) (err error) {
	return instance.SetProperty("UpdateNotificationLevel", (value))
}

// GetUpdateNotificationLevel gets the value of UpdateNotificationLevel for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyUpdateNotificationLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("UpdateNotificationLevel")
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

	value = int32(valuetmp)

	return
}

// SetUpdateServiceUrl sets the value of UpdateServiceUrl for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyUpdateServiceUrl(value string) (err error) {
	return instance.SetProperty("UpdateServiceUrl", (value))
}

// GetUpdateServiceUrl gets the value of UpdateServiceUrl for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyUpdateServiceUrl() (value string, err error) {
	retValue, err := instance.GetProperty("UpdateServiceUrl")
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

// SetUpdateServiceUrlAlternate sets the value of UpdateServiceUrlAlternate for the instance
func (instance *MDM_Policy_Config01_Update02) SetPropertyUpdateServiceUrlAlternate(value string) (err error) {
	return instance.SetProperty("UpdateServiceUrlAlternate", (value))
}

// GetUpdateServiceUrlAlternate gets the value of UpdateServiceUrlAlternate for the instance
func (instance *MDM_Policy_Config01_Update02) GetPropertyUpdateServiceUrlAlternate() (value string, err error) {
	retValue, err := instance.GetProperty("UpdateServiceUrlAlternate")
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
