// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_DSCMetaConfiguration struct
type MSFT_DSCMetaConfiguration struct {
	*cim.WmiInstance

	//
	ActionAfterReboot string

	//
	AgentId string

	//
	AllowModuleOverwrite bool

	//
	CertificateID string

	//
	ConfigurationDownloadManagers []OMI_ConfigurationDownloadManager

	//
	ConfigurationID string

	//
	ConfigurationMode string

	//
	ConfigurationModeFrequencyMins uint32

	//
	Credential MSFT_Credential

	//
	DebugMode []string

	//
	DownloadManagerCustomData []MSFT_KeyValuePair

	//
	DownloadManagerName string

	//
	LCMCompatibleVersions []string

	//
	LCMState string

	//
	LCMStateDetail string

	//
	LCMVersion string

	//
	MaximumDownloadSizeMB uint32

	//
	PartialConfigurations []MSFT_PartialConfiguration

	//
	RebootNodeIfNeeded bool

	//
	RefreshFrequencyMins uint32

	//
	RefreshMode string

	//
	ReportManagers []OMI_ReportManager

	//
	ResourceModuleManagers []OMI_ResourceModuleManager

	//
	SignatureValidationPolicy string

	//
	SignatureValidations []MSFT_SignatureValidation

	//
	StatusRetentionTimeInDays uint32
}

func NewMSFT_DSCMetaConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSFT_DSCMetaConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCMetaConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DSCMetaConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DSCMetaConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCMetaConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetActionAfterReboot sets the value of ActionAfterReboot for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyActionAfterReboot(value string) (err error) {
	return instance.SetProperty("ActionAfterReboot", (value))
}

// GetActionAfterReboot gets the value of ActionAfterReboot for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyActionAfterReboot() (value string, err error) {
	retValue, err := instance.GetProperty("ActionAfterReboot")
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

// SetAgentId sets the value of AgentId for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyAgentId(value string) (err error) {
	return instance.SetProperty("AgentId", (value))
}

// GetAgentId gets the value of AgentId for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyAgentId() (value string, err error) {
	retValue, err := instance.GetProperty("AgentId")
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

// SetAllowModuleOverwrite sets the value of AllowModuleOverwrite for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyAllowModuleOverwrite(value bool) (err error) {
	return instance.SetProperty("AllowModuleOverwrite", (value))
}

// GetAllowModuleOverwrite gets the value of AllowModuleOverwrite for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyAllowModuleOverwrite() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowModuleOverwrite")
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

// SetCertificateID sets the value of CertificateID for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyCertificateID(value string) (err error) {
	return instance.SetProperty("CertificateID", (value))
}

// GetCertificateID gets the value of CertificateID for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyCertificateID() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateID")
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

// SetConfigurationDownloadManagers sets the value of ConfigurationDownloadManagers for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyConfigurationDownloadManagers(value []OMI_ConfigurationDownloadManager) (err error) {
	return instance.SetProperty("ConfigurationDownloadManagers", (value))
}

// GetConfigurationDownloadManagers gets the value of ConfigurationDownloadManagers for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyConfigurationDownloadManagers() (value []OMI_ConfigurationDownloadManager, err error) {
	retValue, err := instance.GetProperty("ConfigurationDownloadManagers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(OMI_ConfigurationDownloadManager)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " OMI_ConfigurationDownloadManager is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, OMI_ConfigurationDownloadManager(valuetmp))
	}

	return
}

// SetConfigurationID sets the value of ConfigurationID for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyConfigurationID(value string) (err error) {
	return instance.SetProperty("ConfigurationID", (value))
}

// GetConfigurationID gets the value of ConfigurationID for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyConfigurationID() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigurationID")
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

// SetConfigurationMode sets the value of ConfigurationMode for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyConfigurationMode(value string) (err error) {
	return instance.SetProperty("ConfigurationMode", (value))
}

// GetConfigurationMode gets the value of ConfigurationMode for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyConfigurationMode() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigurationMode")
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

// SetConfigurationModeFrequencyMins sets the value of ConfigurationModeFrequencyMins for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyConfigurationModeFrequencyMins(value uint32) (err error) {
	return instance.SetProperty("ConfigurationModeFrequencyMins", (value))
}

// GetConfigurationModeFrequencyMins gets the value of ConfigurationModeFrequencyMins for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyConfigurationModeFrequencyMins() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConfigurationModeFrequencyMins")
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

// SetCredential sets the value of Credential for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyCredential(value MSFT_Credential) (err error) {
	return instance.SetProperty("Credential", (value))
}

// GetCredential gets the value of Credential for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyCredential() (value MSFT_Credential, err error) {
	retValue, err := instance.GetProperty("Credential")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_Credential)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_Credential is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_Credential(valuetmp)

	return
}

// SetDebugMode sets the value of DebugMode for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyDebugMode(value []string) (err error) {
	return instance.SetProperty("DebugMode", (value))
}

// GetDebugMode gets the value of DebugMode for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyDebugMode() (value []string, err error) {
	retValue, err := instance.GetProperty("DebugMode")
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

// SetDownloadManagerCustomData sets the value of DownloadManagerCustomData for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyDownloadManagerCustomData(value []MSFT_KeyValuePair) (err error) {
	return instance.SetProperty("DownloadManagerCustomData", (value))
}

// GetDownloadManagerCustomData gets the value of DownloadManagerCustomData for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyDownloadManagerCustomData() (value []MSFT_KeyValuePair, err error) {
	retValue, err := instance.GetProperty("DownloadManagerCustomData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_KeyValuePair)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_KeyValuePair is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_KeyValuePair(valuetmp))
	}

	return
}

// SetDownloadManagerName sets the value of DownloadManagerName for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyDownloadManagerName(value string) (err error) {
	return instance.SetProperty("DownloadManagerName", (value))
}

// GetDownloadManagerName gets the value of DownloadManagerName for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyDownloadManagerName() (value string, err error) {
	retValue, err := instance.GetProperty("DownloadManagerName")
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

// SetLCMCompatibleVersions sets the value of LCMCompatibleVersions for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyLCMCompatibleVersions(value []string) (err error) {
	return instance.SetProperty("LCMCompatibleVersions", (value))
}

// GetLCMCompatibleVersions gets the value of LCMCompatibleVersions for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyLCMCompatibleVersions() (value []string, err error) {
	retValue, err := instance.GetProperty("LCMCompatibleVersions")
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

// SetLCMState sets the value of LCMState for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyLCMState(value string) (err error) {
	return instance.SetProperty("LCMState", (value))
}

// GetLCMState gets the value of LCMState for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyLCMState() (value string, err error) {
	retValue, err := instance.GetProperty("LCMState")
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

// SetLCMStateDetail sets the value of LCMStateDetail for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyLCMStateDetail(value string) (err error) {
	return instance.SetProperty("LCMStateDetail", (value))
}

// GetLCMStateDetail gets the value of LCMStateDetail for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyLCMStateDetail() (value string, err error) {
	retValue, err := instance.GetProperty("LCMStateDetail")
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

// SetLCMVersion sets the value of LCMVersion for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyLCMVersion(value string) (err error) {
	return instance.SetProperty("LCMVersion", (value))
}

// GetLCMVersion gets the value of LCMVersion for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyLCMVersion() (value string, err error) {
	retValue, err := instance.GetProperty("LCMVersion")
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

// SetMaximumDownloadSizeMB sets the value of MaximumDownloadSizeMB for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyMaximumDownloadSizeMB(value uint32) (err error) {
	return instance.SetProperty("MaximumDownloadSizeMB", (value))
}

// GetMaximumDownloadSizeMB gets the value of MaximumDownloadSizeMB for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyMaximumDownloadSizeMB() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumDownloadSizeMB")
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

// SetPartialConfigurations sets the value of PartialConfigurations for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyPartialConfigurations(value []MSFT_PartialConfiguration) (err error) {
	return instance.SetProperty("PartialConfigurations", (value))
}

// GetPartialConfigurations gets the value of PartialConfigurations for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyPartialConfigurations() (value []MSFT_PartialConfiguration, err error) {
	retValue, err := instance.GetProperty("PartialConfigurations")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_PartialConfiguration)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_PartialConfiguration is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_PartialConfiguration(valuetmp))
	}

	return
}

// SetRebootNodeIfNeeded sets the value of RebootNodeIfNeeded for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyRebootNodeIfNeeded(value bool) (err error) {
	return instance.SetProperty("RebootNodeIfNeeded", (value))
}

// GetRebootNodeIfNeeded gets the value of RebootNodeIfNeeded for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyRebootNodeIfNeeded() (value bool, err error) {
	retValue, err := instance.GetProperty("RebootNodeIfNeeded")
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

// SetRefreshFrequencyMins sets the value of RefreshFrequencyMins for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyRefreshFrequencyMins(value uint32) (err error) {
	return instance.SetProperty("RefreshFrequencyMins", (value))
}

// GetRefreshFrequencyMins gets the value of RefreshFrequencyMins for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyRefreshFrequencyMins() (value uint32, err error) {
	retValue, err := instance.GetProperty("RefreshFrequencyMins")
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

// SetRefreshMode sets the value of RefreshMode for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyRefreshMode(value string) (err error) {
	return instance.SetProperty("RefreshMode", (value))
}

// GetRefreshMode gets the value of RefreshMode for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyRefreshMode() (value string, err error) {
	retValue, err := instance.GetProperty("RefreshMode")
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

// SetReportManagers sets the value of ReportManagers for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyReportManagers(value []OMI_ReportManager) (err error) {
	return instance.SetProperty("ReportManagers", (value))
}

// GetReportManagers gets the value of ReportManagers for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyReportManagers() (value []OMI_ReportManager, err error) {
	retValue, err := instance.GetProperty("ReportManagers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(OMI_ReportManager)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " OMI_ReportManager is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, OMI_ReportManager(valuetmp))
	}

	return
}

// SetResourceModuleManagers sets the value of ResourceModuleManagers for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyResourceModuleManagers(value []OMI_ResourceModuleManager) (err error) {
	return instance.SetProperty("ResourceModuleManagers", (value))
}

// GetResourceModuleManagers gets the value of ResourceModuleManagers for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyResourceModuleManagers() (value []OMI_ResourceModuleManager, err error) {
	retValue, err := instance.GetProperty("ResourceModuleManagers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(OMI_ResourceModuleManager)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " OMI_ResourceModuleManager is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, OMI_ResourceModuleManager(valuetmp))
	}

	return
}

// SetSignatureValidationPolicy sets the value of SignatureValidationPolicy for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertySignatureValidationPolicy(value string) (err error) {
	return instance.SetProperty("SignatureValidationPolicy", (value))
}

// GetSignatureValidationPolicy gets the value of SignatureValidationPolicy for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertySignatureValidationPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("SignatureValidationPolicy")
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

// SetSignatureValidations sets the value of SignatureValidations for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertySignatureValidations(value []MSFT_SignatureValidation) (err error) {
	return instance.SetProperty("SignatureValidations", (value))
}

// GetSignatureValidations gets the value of SignatureValidations for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertySignatureValidations() (value []MSFT_SignatureValidation, err error) {
	retValue, err := instance.GetProperty("SignatureValidations")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_SignatureValidation)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_SignatureValidation is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_SignatureValidation(valuetmp))
	}

	return
}

// SetStatusRetentionTimeInDays sets the value of StatusRetentionTimeInDays for the instance
func (instance *MSFT_DSCMetaConfiguration) SetPropertyStatusRetentionTimeInDays(value uint32) (err error) {
	return instance.SetProperty("StatusRetentionTimeInDays", (value))
}

// GetStatusRetentionTimeInDays gets the value of StatusRetentionTimeInDays for the instance
func (instance *MSFT_DSCMetaConfiguration) GetPropertyStatusRetentionTimeInDays() (value uint32, err error) {
	retValue, err := instance.GetProperty("StatusRetentionTimeInDays")
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
