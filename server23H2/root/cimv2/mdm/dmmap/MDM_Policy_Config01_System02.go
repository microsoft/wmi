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

// MDM_Policy_Config01_System02 struct
type MDM_Policy_Config01_System02 struct {
	*cim.WmiInstance

	//
	AllowBuildPreview int32

	//
	AllowCommercialDataPipeline int32

	//
	AllowDesktopAnalyticsProcessing int32

	//
	AllowDeviceNameInDiagnosticData int32

	//
	AllowEmbeddedMode int32

	//
	AllowExperimentation int32

	//
	AllowFontProviders int32

	//
	AllowLocation int32

	//
	AllowMicrosoftManagedDesktopProcessing int32

	//
	AllowStorageCard int32

	//
	AllowTelemetry int32

	//
	AllowUpdateComplianceProcessing int32

	//
	AllowUserToResetPhone int32

	//
	AllowWUfBCloudProcessing int32

	//
	BootStartDriverInitialization string

	//
	ConfigureMicrosoft365UploadEndpoint string

	//
	ConfigureTelemetryOptInChangeNotification int32

	//
	ConfigureTelemetryOptInSettingsUx int32

	//
	DisableDeviceDelete int32

	//
	DisableDiagnosticDataViewer int32

	//
	DisableDirectXDatabaseUpdate int32

	//
	DisableEnterpriseAuthProxy int32

	//
	DisableOneDriveFileSync int32

	//
	DisableOneSettingsDownloads int32

	//
	DisableSystemRestore string

	//
	EnableOneSettingsAuditing int32

	//
	FeedbackHubAlwaysSaveDiagnosticsLocally int32

	//
	InstanceID string

	//
	LimitDiagnosticLogCollection int32

	//
	LimitDumpCollection int32

	//
	LimitEnhancedDiagnosticDataWindowsAnalytics int32

	//
	ParentID string

	//
	TelemetryProxy string

	//
	TurnOffFileHistory int32
}

func NewMDM_Policy_Config01_System02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_System02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_System02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_System02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_System02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_System02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowBuildPreview sets the value of AllowBuildPreview for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowBuildPreview(value int32) (err error) {
	return instance.SetProperty("AllowBuildPreview", (value))
}

// GetAllowBuildPreview gets the value of AllowBuildPreview for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowBuildPreview() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowBuildPreview")
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

// SetAllowCommercialDataPipeline sets the value of AllowCommercialDataPipeline for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowCommercialDataPipeline(value int32) (err error) {
	return instance.SetProperty("AllowCommercialDataPipeline", (value))
}

// GetAllowCommercialDataPipeline gets the value of AllowCommercialDataPipeline for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowCommercialDataPipeline() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowCommercialDataPipeline")
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

// SetAllowDesktopAnalyticsProcessing sets the value of AllowDesktopAnalyticsProcessing for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowDesktopAnalyticsProcessing(value int32) (err error) {
	return instance.SetProperty("AllowDesktopAnalyticsProcessing", (value))
}

// GetAllowDesktopAnalyticsProcessing gets the value of AllowDesktopAnalyticsProcessing for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowDesktopAnalyticsProcessing() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowDesktopAnalyticsProcessing")
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

// SetAllowDeviceNameInDiagnosticData sets the value of AllowDeviceNameInDiagnosticData for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowDeviceNameInDiagnosticData(value int32) (err error) {
	return instance.SetProperty("AllowDeviceNameInDiagnosticData", (value))
}

// GetAllowDeviceNameInDiagnosticData gets the value of AllowDeviceNameInDiagnosticData for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowDeviceNameInDiagnosticData() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowDeviceNameInDiagnosticData")
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

// SetAllowEmbeddedMode sets the value of AllowEmbeddedMode for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowEmbeddedMode(value int32) (err error) {
	return instance.SetProperty("AllowEmbeddedMode", (value))
}

// GetAllowEmbeddedMode gets the value of AllowEmbeddedMode for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowEmbeddedMode() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowEmbeddedMode")
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

// SetAllowExperimentation sets the value of AllowExperimentation for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowExperimentation(value int32) (err error) {
	return instance.SetProperty("AllowExperimentation", (value))
}

// GetAllowExperimentation gets the value of AllowExperimentation for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowExperimentation() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowExperimentation")
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

// SetAllowFontProviders sets the value of AllowFontProviders for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowFontProviders(value int32) (err error) {
	return instance.SetProperty("AllowFontProviders", (value))
}

// GetAllowFontProviders gets the value of AllowFontProviders for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowFontProviders() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowFontProviders")
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

// SetAllowLocation sets the value of AllowLocation for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowLocation(value int32) (err error) {
	return instance.SetProperty("AllowLocation", (value))
}

// GetAllowLocation gets the value of AllowLocation for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowLocation() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowLocation")
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

// SetAllowMicrosoftManagedDesktopProcessing sets the value of AllowMicrosoftManagedDesktopProcessing for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowMicrosoftManagedDesktopProcessing(value int32) (err error) {
	return instance.SetProperty("AllowMicrosoftManagedDesktopProcessing", (value))
}

// GetAllowMicrosoftManagedDesktopProcessing gets the value of AllowMicrosoftManagedDesktopProcessing for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowMicrosoftManagedDesktopProcessing() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowMicrosoftManagedDesktopProcessing")
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

// SetAllowStorageCard sets the value of AllowStorageCard for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowStorageCard(value int32) (err error) {
	return instance.SetProperty("AllowStorageCard", (value))
}

// GetAllowStorageCard gets the value of AllowStorageCard for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowStorageCard() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowStorageCard")
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

// SetAllowTelemetry sets the value of AllowTelemetry for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowTelemetry(value int32) (err error) {
	return instance.SetProperty("AllowTelemetry", (value))
}

// GetAllowTelemetry gets the value of AllowTelemetry for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowTelemetry() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowTelemetry")
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

// SetAllowUpdateComplianceProcessing sets the value of AllowUpdateComplianceProcessing for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowUpdateComplianceProcessing(value int32) (err error) {
	return instance.SetProperty("AllowUpdateComplianceProcessing", (value))
}

// GetAllowUpdateComplianceProcessing gets the value of AllowUpdateComplianceProcessing for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowUpdateComplianceProcessing() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowUpdateComplianceProcessing")
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

// SetAllowUserToResetPhone sets the value of AllowUserToResetPhone for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowUserToResetPhone(value int32) (err error) {
	return instance.SetProperty("AllowUserToResetPhone", (value))
}

// GetAllowUserToResetPhone gets the value of AllowUserToResetPhone for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowUserToResetPhone() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowUserToResetPhone")
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

// SetAllowWUfBCloudProcessing sets the value of AllowWUfBCloudProcessing for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyAllowWUfBCloudProcessing(value int32) (err error) {
	return instance.SetProperty("AllowWUfBCloudProcessing", (value))
}

// GetAllowWUfBCloudProcessing gets the value of AllowWUfBCloudProcessing for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyAllowWUfBCloudProcessing() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWUfBCloudProcessing")
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

// SetBootStartDriverInitialization sets the value of BootStartDriverInitialization for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyBootStartDriverInitialization(value string) (err error) {
	return instance.SetProperty("BootStartDriverInitialization", (value))
}

// GetBootStartDriverInitialization gets the value of BootStartDriverInitialization for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyBootStartDriverInitialization() (value string, err error) {
	retValue, err := instance.GetProperty("BootStartDriverInitialization")
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

// SetConfigureMicrosoft365UploadEndpoint sets the value of ConfigureMicrosoft365UploadEndpoint for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyConfigureMicrosoft365UploadEndpoint(value string) (err error) {
	return instance.SetProperty("ConfigureMicrosoft365UploadEndpoint", (value))
}

// GetConfigureMicrosoft365UploadEndpoint gets the value of ConfigureMicrosoft365UploadEndpoint for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyConfigureMicrosoft365UploadEndpoint() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigureMicrosoft365UploadEndpoint")
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

// SetConfigureTelemetryOptInChangeNotification sets the value of ConfigureTelemetryOptInChangeNotification for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyConfigureTelemetryOptInChangeNotification(value int32) (err error) {
	return instance.SetProperty("ConfigureTelemetryOptInChangeNotification", (value))
}

// GetConfigureTelemetryOptInChangeNotification gets the value of ConfigureTelemetryOptInChangeNotification for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyConfigureTelemetryOptInChangeNotification() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureTelemetryOptInChangeNotification")
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

// SetConfigureTelemetryOptInSettingsUx sets the value of ConfigureTelemetryOptInSettingsUx for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyConfigureTelemetryOptInSettingsUx(value int32) (err error) {
	return instance.SetProperty("ConfigureTelemetryOptInSettingsUx", (value))
}

// GetConfigureTelemetryOptInSettingsUx gets the value of ConfigureTelemetryOptInSettingsUx for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyConfigureTelemetryOptInSettingsUx() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureTelemetryOptInSettingsUx")
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

// SetDisableDeviceDelete sets the value of DisableDeviceDelete for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyDisableDeviceDelete(value int32) (err error) {
	return instance.SetProperty("DisableDeviceDelete", (value))
}

// GetDisableDeviceDelete gets the value of DisableDeviceDelete for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyDisableDeviceDelete() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableDeviceDelete")
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

// SetDisableDiagnosticDataViewer sets the value of DisableDiagnosticDataViewer for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyDisableDiagnosticDataViewer(value int32) (err error) {
	return instance.SetProperty("DisableDiagnosticDataViewer", (value))
}

// GetDisableDiagnosticDataViewer gets the value of DisableDiagnosticDataViewer for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyDisableDiagnosticDataViewer() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableDiagnosticDataViewer")
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

// SetDisableDirectXDatabaseUpdate sets the value of DisableDirectXDatabaseUpdate for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyDisableDirectXDatabaseUpdate(value int32) (err error) {
	return instance.SetProperty("DisableDirectXDatabaseUpdate", (value))
}

// GetDisableDirectXDatabaseUpdate gets the value of DisableDirectXDatabaseUpdate for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyDisableDirectXDatabaseUpdate() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableDirectXDatabaseUpdate")
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

// SetDisableEnterpriseAuthProxy sets the value of DisableEnterpriseAuthProxy for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyDisableEnterpriseAuthProxy(value int32) (err error) {
	return instance.SetProperty("DisableEnterpriseAuthProxy", (value))
}

// GetDisableEnterpriseAuthProxy gets the value of DisableEnterpriseAuthProxy for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyDisableEnterpriseAuthProxy() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableEnterpriseAuthProxy")
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

// SetDisableOneDriveFileSync sets the value of DisableOneDriveFileSync for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyDisableOneDriveFileSync(value int32) (err error) {
	return instance.SetProperty("DisableOneDriveFileSync", (value))
}

// GetDisableOneDriveFileSync gets the value of DisableOneDriveFileSync for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyDisableOneDriveFileSync() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableOneDriveFileSync")
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

// SetDisableOneSettingsDownloads sets the value of DisableOneSettingsDownloads for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyDisableOneSettingsDownloads(value int32) (err error) {
	return instance.SetProperty("DisableOneSettingsDownloads", (value))
}

// GetDisableOneSettingsDownloads gets the value of DisableOneSettingsDownloads for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyDisableOneSettingsDownloads() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableOneSettingsDownloads")
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

// SetDisableSystemRestore sets the value of DisableSystemRestore for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyDisableSystemRestore(value string) (err error) {
	return instance.SetProperty("DisableSystemRestore", (value))
}

// GetDisableSystemRestore gets the value of DisableSystemRestore for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyDisableSystemRestore() (value string, err error) {
	retValue, err := instance.GetProperty("DisableSystemRestore")
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

// SetEnableOneSettingsAuditing sets the value of EnableOneSettingsAuditing for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyEnableOneSettingsAuditing(value int32) (err error) {
	return instance.SetProperty("EnableOneSettingsAuditing", (value))
}

// GetEnableOneSettingsAuditing gets the value of EnableOneSettingsAuditing for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyEnableOneSettingsAuditing() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableOneSettingsAuditing")
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

// SetFeedbackHubAlwaysSaveDiagnosticsLocally sets the value of FeedbackHubAlwaysSaveDiagnosticsLocally for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyFeedbackHubAlwaysSaveDiagnosticsLocally(value int32) (err error) {
	return instance.SetProperty("FeedbackHubAlwaysSaveDiagnosticsLocally", (value))
}

// GetFeedbackHubAlwaysSaveDiagnosticsLocally gets the value of FeedbackHubAlwaysSaveDiagnosticsLocally for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyFeedbackHubAlwaysSaveDiagnosticsLocally() (value int32, err error) {
	retValue, err := instance.GetProperty("FeedbackHubAlwaysSaveDiagnosticsLocally")
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
func (instance *MDM_Policy_Config01_System02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyInstanceID() (value string, err error) {
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

// SetLimitDiagnosticLogCollection sets the value of LimitDiagnosticLogCollection for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyLimitDiagnosticLogCollection(value int32) (err error) {
	return instance.SetProperty("LimitDiagnosticLogCollection", (value))
}

// GetLimitDiagnosticLogCollection gets the value of LimitDiagnosticLogCollection for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyLimitDiagnosticLogCollection() (value int32, err error) {
	retValue, err := instance.GetProperty("LimitDiagnosticLogCollection")
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

// SetLimitDumpCollection sets the value of LimitDumpCollection for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyLimitDumpCollection(value int32) (err error) {
	return instance.SetProperty("LimitDumpCollection", (value))
}

// GetLimitDumpCollection gets the value of LimitDumpCollection for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyLimitDumpCollection() (value int32, err error) {
	retValue, err := instance.GetProperty("LimitDumpCollection")
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

// SetLimitEnhancedDiagnosticDataWindowsAnalytics sets the value of LimitEnhancedDiagnosticDataWindowsAnalytics for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyLimitEnhancedDiagnosticDataWindowsAnalytics(value int32) (err error) {
	return instance.SetProperty("LimitEnhancedDiagnosticDataWindowsAnalytics", (value))
}

// GetLimitEnhancedDiagnosticDataWindowsAnalytics gets the value of LimitEnhancedDiagnosticDataWindowsAnalytics for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyLimitEnhancedDiagnosticDataWindowsAnalytics() (value int32, err error) {
	retValue, err := instance.GetProperty("LimitEnhancedDiagnosticDataWindowsAnalytics")
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
func (instance *MDM_Policy_Config01_System02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyParentID() (value string, err error) {
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

// SetTelemetryProxy sets the value of TelemetryProxy for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyTelemetryProxy(value string) (err error) {
	return instance.SetProperty("TelemetryProxy", (value))
}

// GetTelemetryProxy gets the value of TelemetryProxy for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyTelemetryProxy() (value string, err error) {
	retValue, err := instance.GetProperty("TelemetryProxy")
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

// SetTurnOffFileHistory sets the value of TurnOffFileHistory for the instance
func (instance *MDM_Policy_Config01_System02) SetPropertyTurnOffFileHistory(value int32) (err error) {
	return instance.SetProperty("TurnOffFileHistory", (value))
}

// GetTurnOffFileHistory gets the value of TurnOffFileHistory for the instance
func (instance *MDM_Policy_Config01_System02) GetPropertyTurnOffFileHistory() (value int32, err error) {
	retValue, err := instance.GetProperty("TurnOffFileHistory")
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
