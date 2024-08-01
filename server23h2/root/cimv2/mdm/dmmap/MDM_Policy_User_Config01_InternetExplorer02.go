// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_User_Config01_InternetExplorer02 struct
type MDM_Policy_User_Config01_InternetExplorer02 struct {
	*cim.WmiInstance

	//
	AddSearchProvider string

	//
	AllowActiveXFiltering string

	//
	AllowAddOnList string

	//
	AllowAutoComplete string

	//
	AllowCertificateAddressMismatchWarning string

	//
	AllowDeletingBrowsingHistoryOnExit string

	//
	AllowEnhancedProtectedMode string

	//
	AllowEnhancedSuggestionsInAddressBar string

	//
	AllowEnterpriseModeFromToolsMenu string

	//
	AllowEnterpriseModeSiteList string

	//
	AllowInternetExplorer7PolicyList string

	//
	AllowInternetExplorerStandardsMode string

	//
	AllowInternetZoneTemplate string

	//
	AllowIntranetZoneTemplate string

	//
	AllowLocalMachineZoneTemplate string

	//
	AllowLockedDownInternetZoneTemplate string

	//
	AllowLockedDownIntranetZoneTemplate string

	//
	AllowLockedDownLocalMachineZoneTemplate string

	//
	AllowLockedDownRestrictedSitesZoneTemplate string

	//
	AllowOneWordEntry string

	//
	AllowSaveTargetAsInIEMode string

	//
	AllowSiteToZoneAssignmentList string

	//
	AllowsLockedDownTrustedSitesZoneTemplate string

	//
	AllowSoftwareWhenSignatureIsInvalid string

	//
	AllowsRestrictedSitesZoneTemplate string

	//
	AllowSuggestedSites string

	//
	AllowTrustedSitesZoneTemplate string

	//
	CheckServerCertificateRevocation string

	//
	CheckSignaturesOnDownloadedPrograms string

	//
	ConfigureEdgeRedirectChannel string

	//
	ConsistentMimeHandlingInternetExplorerProcesses string

	//
	DisableActiveXVersionListAutoDownload string

	//
	DisableBypassOfSmartScreenWarnings string

	//
	DisableBypassOfSmartScreenWarningsAboutUncommonFiles string

	//
	DisableCompatView string

	//
	DisableConfiguringHistory string

	//
	DisableCrashDetection string

	//
	DisableCustomerExperienceImprovementProgramParticipation string

	//
	DisableDeletingUserVisitedWebsites string

	//
	DisableEnclosureDownloading string

	//
	DisableEncryptionSupport string

	//
	DisableFeedsBackgroundSync string

	//
	DisableFirstRunWizard string

	//
	DisableFlipAheadFeature string

	//
	DisableGeolocation string

	//
	DisableHomePageChange string

	//
	DisableIgnoringCertificateErrors string

	//
	DisableInPrivateBrowsing string

	//
	DisableInternetExplorerApp string

	//
	DisableProcessesInEnhancedProtectedMode string

	//
	DisableProxyChange string

	//
	DisableSearchProviderChange string

	//
	DisableSecondaryHomePageChange string

	//
	DisableSecuritySettingsCheck string

	//
	DisableWebAddressAutoComplete string

	//
	DoNotAllowActiveXControlsInProtectedMode string

	//
	DoNotBlockOutdatedActiveXControls string

	//
	DoNotBlockOutdatedActiveXControlsOnSpecificDomains string

	//
	EnableExtendedIEModeHotkeys string

	//
	IncludeAllLocalSites string

	//
	IncludeAllNetworkPaths string

	//
	InstanceID string

	//
	InternetZoneAllowAccessToDataSources string

	//
	InternetZoneAllowAutomaticPromptingForActiveXControls string

	//
	InternetZoneAllowAutomaticPromptingForFileDownloads string

	//
	InternetZoneAllowCopyPasteViaScript string

	//
	InternetZoneAllowDragAndDropCopyAndPasteFiles string

	//
	InternetZoneAllowFontDownloads string

	//
	InternetZoneAllowLessPrivilegedSites string

	//
	InternetZoneAllowLoadingOfXAMLFiles string

	//
	InternetZoneAllowNETFrameworkReliantComponents string

	//
	InternetZoneAllowOnlyApprovedDomainsToUseActiveXControls string

	//
	InternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl string

	//
	InternetZoneAllowScriptingOfInternetExplorerWebBrowserControls string

	//
	InternetZoneAllowScriptInitiatedWindows string

	//
	InternetZoneAllowScriptlets string

	//
	InternetZoneAllowSmartScreenIE string

	//
	InternetZoneAllowUpdatesToStatusBarViaScript string

	//
	InternetZoneAllowUserDataPersistence string

	//
	InternetZoneAllowVBScriptToRunInInternetExplorer string

	//
	InternetZoneDoNotRunAntimalwareAgainstActiveXControls string

	//
	InternetZoneDownloadSignedActiveXControls string

	//
	InternetZoneDownloadUnsignedActiveXControls string

	//
	InternetZoneEnableCrossSiteScriptingFilter string

	//
	InternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows string

	//
	InternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows string

	//
	InternetZoneEnableMIMESniffing string

	//
	InternetZoneEnableProtectedMode string

	//
	InternetZoneIncludeLocalPathWhenUploadingFilesToServer string

	//
	InternetZoneInitializeAndScriptActiveXControls string

	//
	InternetZoneJavaPermissions string

	//
	InternetZoneLaunchingApplicationsAndFilesInIFRAME string

	//
	InternetZoneLogonOptions string

	//
	InternetZoneNavigateWindowsAndFrames string

	//
	InternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode string

	//
	InternetZoneShowSecurityWarningForPotentiallyUnsafeFiles string

	//
	InternetZoneUsePopupBlocker string

	//
	IntranetZoneAllowAccessToDataSources string

	//
	IntranetZoneAllowAutomaticPromptingForActiveXControls string

	//
	IntranetZoneAllowAutomaticPromptingForFileDownloads string

	//
	IntranetZoneAllowFontDownloads string

	//
	IntranetZoneAllowLessPrivilegedSites string

	//
	IntranetZoneAllowNETFrameworkReliantComponents string

	//
	IntranetZoneAllowScriptlets string

	//
	IntranetZoneAllowSmartScreenIE string

	//
	IntranetZoneAllowUserDataPersistence string

	//
	IntranetZoneDoNotRunAntimalwareAgainstActiveXControls string

	//
	IntranetZoneInitializeAndScriptActiveXControls string

	//
	IntranetZoneJavaPermissions string

	//
	IntranetZoneNavigateWindowsAndFrames string

	//
	JScriptReplacement string

	//
	KeepIntranetSitesInInternetExplorer string

	//
	LocalMachineZoneAllowAccessToDataSources string

	//
	LocalMachineZoneAllowAutomaticPromptingForActiveXControls string

	//
	LocalMachineZoneAllowAutomaticPromptingForFileDownloads string

	//
	LocalMachineZoneAllowFontDownloads string

	//
	LocalMachineZoneAllowLessPrivilegedSites string

	//
	LocalMachineZoneAllowNETFrameworkReliantComponents string

	//
	LocalMachineZoneAllowScriptlets string

	//
	LocalMachineZoneAllowSmartScreenIE string

	//
	LocalMachineZoneAllowUserDataPersistence string

	//
	LocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls string

	//
	LocalMachineZoneInitializeAndScriptActiveXControls string

	//
	LocalMachineZoneJavaPermissions string

	//
	LocalMachineZoneNavigateWindowsAndFrames string

	//
	LockedDownInternetZoneAllowAccessToDataSources string

	//
	LockedDownInternetZoneAllowAutomaticPromptingForActiveXControls string

	//
	LockedDownInternetZoneAllowAutomaticPromptingForFileDownloads string

	//
	LockedDownInternetZoneAllowFontDownloads string

	//
	LockedDownInternetZoneAllowLessPrivilegedSites string

	//
	LockedDownInternetZoneAllowNETFrameworkReliantComponents string

	//
	LockedDownInternetZoneAllowScriptlets string

	//
	LockedDownInternetZoneAllowSmartScreenIE string

	//
	LockedDownInternetZoneAllowUserDataPersistence string

	//
	LockedDownInternetZoneInitializeAndScriptActiveXControls string

	//
	LockedDownInternetZoneJavaPermissions string

	//
	LockedDownInternetZoneNavigateWindowsAndFrames string

	//
	LockedDownIntranetJavaPermissions string

	//
	LockedDownIntranetZoneAllowAccessToDataSources string

	//
	LockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls string

	//
	LockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads string

	//
	LockedDownIntranetZoneAllowFontDownloads string

	//
	LockedDownIntranetZoneAllowLessPrivilegedSites string

	//
	LockedDownIntranetZoneAllowNETFrameworkReliantComponents string

	//
	LockedDownIntranetZoneAllowScriptlets string

	//
	LockedDownIntranetZoneAllowSmartScreenIE string

	//
	LockedDownIntranetZoneAllowUserDataPersistence string

	//
	LockedDownIntranetZoneInitializeAndScriptActiveXControls string

	//
	LockedDownIntranetZoneNavigateWindowsAndFrames string

	//
	LockedDownLocalMachineZoneAllowAccessToDataSources string

	//
	LockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls string

	//
	LockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads string

	//
	LockedDownLocalMachineZoneAllowFontDownloads string

	//
	LockedDownLocalMachineZoneAllowLessPrivilegedSites string

	//
	LockedDownLocalMachineZoneAllowNETFrameworkReliantComponents string

	//
	LockedDownLocalMachineZoneAllowScriptlets string

	//
	LockedDownLocalMachineZoneAllowSmartScreenIE string

	//
	LockedDownLocalMachineZoneAllowUserDataPersistence string

	//
	LockedDownLocalMachineZoneInitializeAndScriptActiveXControls string

	//
	LockedDownLocalMachineZoneJavaPermissions string

	//
	LockedDownLocalMachineZoneNavigateWindowsAndFrames string

	//
	LockedDownRestrictedSitesZoneAllowAccessToDataSources string

	//
	LockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls string

	//
	LockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads string

	//
	LockedDownRestrictedSitesZoneAllowFontDownloads string

	//
	LockedDownRestrictedSitesZoneAllowLessPrivilegedSites string

	//
	LockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents string

	//
	LockedDownRestrictedSitesZoneAllowScriptlets string

	//
	LockedDownRestrictedSitesZoneAllowSmartScreenIE string

	//
	LockedDownRestrictedSitesZoneAllowUserDataPersistence string

	//
	LockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls string

	//
	LockedDownRestrictedSitesZoneJavaPermissions string

	//
	LockedDownRestrictedSitesZoneNavigateWindowsAndFrames string

	//
	LockedDownTrustedSitesZoneAllowAccessToDataSources string

	//
	LockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls string

	//
	LockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads string

	//
	LockedDownTrustedSitesZoneAllowFontDownloads string

	//
	LockedDownTrustedSitesZoneAllowLessPrivilegedSites string

	//
	LockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents string

	//
	LockedDownTrustedSitesZoneAllowScriptlets string

	//
	LockedDownTrustedSitesZoneAllowSmartScreenIE string

	//
	LockedDownTrustedSitesZoneAllowUserDataPersistence string

	//
	LockedDownTrustedSitesZoneInitializeAndScriptActiveXControls string

	//
	LockedDownTrustedSitesZoneJavaPermissions string

	//
	LockedDownTrustedSitesZoneNavigateWindowsAndFrames string

	//
	MimeSniffingSafetyFeatureInternetExplorerProcesses string

	//
	MKProtocolSecurityRestrictionInternetExplorerProcesses string

	//
	NewTabDefaultPage string

	//
	NotificationBarInternetExplorerProcesses string

	//
	ParentID string

	//
	PreventManagingSmartScreenFilter string

	//
	PreventPerUserInstallationOfActiveXControls string

	//
	ProtectionFromZoneElevationInternetExplorerProcesses string

	//
	RemoveRunThisTimeButtonForOutdatedActiveXControls string

	//
	RestrictActiveXInstallInternetExplorerProcesses string

	//
	RestrictedSitesZoneAllowAccessToDataSources string

	//
	RestrictedSitesZoneAllowActiveScripting string

	//
	RestrictedSitesZoneAllowAutomaticPromptingForActiveXControls string

	//
	RestrictedSitesZoneAllowAutomaticPromptingForFileDownloads string

	//
	RestrictedSitesZoneAllowBinaryAndScriptBehaviors string

	//
	RestrictedSitesZoneAllowCopyPasteViaScript string

	//
	RestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles string

	//
	RestrictedSitesZoneAllowFileDownloads string

	//
	RestrictedSitesZoneAllowFontDownloads string

	//
	RestrictedSitesZoneAllowLessPrivilegedSites string

	//
	RestrictedSitesZoneAllowLoadingOfXAMLFiles string

	//
	RestrictedSitesZoneAllowMETAREFRESH string

	//
	RestrictedSitesZoneAllowNETFrameworkReliantComponents string

	//
	RestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls string

	//
	RestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl string

	//
	RestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls string

	//
	RestrictedSitesZoneAllowScriptInitiatedWindows string

	//
	RestrictedSitesZoneAllowScriptlets string

	//
	RestrictedSitesZoneAllowSmartScreenIE string

	//
	RestrictedSitesZoneAllowUpdatesToStatusBarViaScript string

	//
	RestrictedSitesZoneAllowUserDataPersistence string

	//
	RestrictedSitesZoneAllowVBScriptToRunInInternetExplorer string

	//
	RestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls string

	//
	RestrictedSitesZoneDownloadSignedActiveXControls string

	//
	RestrictedSitesZoneDownloadUnsignedActiveXControls string

	//
	RestrictedSitesZoneEnableCrossSiteScriptingFilter string

	//
	RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows string

	//
	RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows string

	//
	RestrictedSitesZoneEnableMIMESniffing string

	//
	RestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer string

	//
	RestrictedSitesZoneInitializeAndScriptActiveXControls string

	//
	RestrictedSitesZoneJavaPermissions string

	//
	RestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME string

	//
	RestrictedSitesZoneLogonOptions string

	//
	RestrictedSitesZoneNavigateWindowsAndFrames string

	//
	RestrictedSitesZoneRunActiveXControlsAndPlugins string

	//
	RestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode string

	//
	RestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting string

	//
	RestrictedSitesZoneScriptingOfJavaApplets string

	//
	RestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles string

	//
	RestrictedSitesZoneTurnOnProtectedMode string

	//
	RestrictedSitesZoneUsePopupBlocker string

	//
	RestrictFileDownloadInternetExplorerProcesses string

	//
	ScriptedWindowSecurityRestrictionsInternetExplorerProcesses string

	//
	SearchProviderList string

	//
	SendSitesNotInEnterpriseSiteListToEdge string

	//
	SpecifyUseOfActiveXInstallerService string

	//
	TrustedSitesZoneAllowAccessToDataSources string

	//
	TrustedSitesZoneAllowAutomaticPromptingForActiveXControls string

	//
	TrustedSitesZoneAllowAutomaticPromptingForFileDownloads string

	//
	TrustedSitesZoneAllowFontDownloads string

	//
	TrustedSitesZoneAllowLessPrivilegedSites string

	//
	TrustedSitesZoneAllowNETFrameworkReliantComponents string

	//
	TrustedSitesZoneAllowScriptlets string

	//
	TrustedSitesZoneAllowSmartScreenIE string

	//
	TrustedSitesZoneAllowUserDataPersistence string

	//
	TrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls string

	//
	TrustedSitesZoneInitializeAndScriptActiveXControls string

	//
	TrustedSitesZoneJavaPermissions string

	//
	TrustedSitesZoneNavigateWindowsAndFrames string
}

func NewMDM_Policy_User_Config01_InternetExplorer02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_User_Config01_InternetExplorer02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_User_Config01_InternetExplorer02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_User_Config01_InternetExplorer02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_User_Config01_InternetExplorer02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_User_Config01_InternetExplorer02{
		WmiInstance: tmp,
	}
	return
}

// SetAddSearchProvider sets the value of AddSearchProvider for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAddSearchProvider(value string) (err error) {
	return instance.SetProperty("AddSearchProvider", (value))
}

// GetAddSearchProvider gets the value of AddSearchProvider for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAddSearchProvider() (value string, err error) {
	retValue, err := instance.GetProperty("AddSearchProvider")
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

// SetAllowActiveXFiltering sets the value of AllowActiveXFiltering for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowActiveXFiltering(value string) (err error) {
	return instance.SetProperty("AllowActiveXFiltering", (value))
}

// GetAllowActiveXFiltering gets the value of AllowActiveXFiltering for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowActiveXFiltering() (value string, err error) {
	retValue, err := instance.GetProperty("AllowActiveXFiltering")
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

// SetAllowAddOnList sets the value of AllowAddOnList for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowAddOnList(value string) (err error) {
	return instance.SetProperty("AllowAddOnList", (value))
}

// GetAllowAddOnList gets the value of AllowAddOnList for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowAddOnList() (value string, err error) {
	retValue, err := instance.GetProperty("AllowAddOnList")
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

// SetAllowAutoComplete sets the value of AllowAutoComplete for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowAutoComplete(value string) (err error) {
	return instance.SetProperty("AllowAutoComplete", (value))
}

// GetAllowAutoComplete gets the value of AllowAutoComplete for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowAutoComplete() (value string, err error) {
	retValue, err := instance.GetProperty("AllowAutoComplete")
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

// SetAllowCertificateAddressMismatchWarning sets the value of AllowCertificateAddressMismatchWarning for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowCertificateAddressMismatchWarning(value string) (err error) {
	return instance.SetProperty("AllowCertificateAddressMismatchWarning", (value))
}

// GetAllowCertificateAddressMismatchWarning gets the value of AllowCertificateAddressMismatchWarning for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowCertificateAddressMismatchWarning() (value string, err error) {
	retValue, err := instance.GetProperty("AllowCertificateAddressMismatchWarning")
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

// SetAllowDeletingBrowsingHistoryOnExit sets the value of AllowDeletingBrowsingHistoryOnExit for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowDeletingBrowsingHistoryOnExit(value string) (err error) {
	return instance.SetProperty("AllowDeletingBrowsingHistoryOnExit", (value))
}

// GetAllowDeletingBrowsingHistoryOnExit gets the value of AllowDeletingBrowsingHistoryOnExit for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowDeletingBrowsingHistoryOnExit() (value string, err error) {
	retValue, err := instance.GetProperty("AllowDeletingBrowsingHistoryOnExit")
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

// SetAllowEnhancedProtectedMode sets the value of AllowEnhancedProtectedMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowEnhancedProtectedMode(value string) (err error) {
	return instance.SetProperty("AllowEnhancedProtectedMode", (value))
}

// GetAllowEnhancedProtectedMode gets the value of AllowEnhancedProtectedMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowEnhancedProtectedMode() (value string, err error) {
	retValue, err := instance.GetProperty("AllowEnhancedProtectedMode")
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

// SetAllowEnhancedSuggestionsInAddressBar sets the value of AllowEnhancedSuggestionsInAddressBar for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowEnhancedSuggestionsInAddressBar(value string) (err error) {
	return instance.SetProperty("AllowEnhancedSuggestionsInAddressBar", (value))
}

// GetAllowEnhancedSuggestionsInAddressBar gets the value of AllowEnhancedSuggestionsInAddressBar for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowEnhancedSuggestionsInAddressBar() (value string, err error) {
	retValue, err := instance.GetProperty("AllowEnhancedSuggestionsInAddressBar")
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

// SetAllowEnterpriseModeFromToolsMenu sets the value of AllowEnterpriseModeFromToolsMenu for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowEnterpriseModeFromToolsMenu(value string) (err error) {
	return instance.SetProperty("AllowEnterpriseModeFromToolsMenu", (value))
}

// GetAllowEnterpriseModeFromToolsMenu gets the value of AllowEnterpriseModeFromToolsMenu for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowEnterpriseModeFromToolsMenu() (value string, err error) {
	retValue, err := instance.GetProperty("AllowEnterpriseModeFromToolsMenu")
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

// SetAllowEnterpriseModeSiteList sets the value of AllowEnterpriseModeSiteList for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowEnterpriseModeSiteList(value string) (err error) {
	return instance.SetProperty("AllowEnterpriseModeSiteList", (value))
}

// GetAllowEnterpriseModeSiteList gets the value of AllowEnterpriseModeSiteList for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowEnterpriseModeSiteList() (value string, err error) {
	retValue, err := instance.GetProperty("AllowEnterpriseModeSiteList")
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

// SetAllowInternetExplorer7PolicyList sets the value of AllowInternetExplorer7PolicyList for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowInternetExplorer7PolicyList(value string) (err error) {
	return instance.SetProperty("AllowInternetExplorer7PolicyList", (value))
}

// GetAllowInternetExplorer7PolicyList gets the value of AllowInternetExplorer7PolicyList for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowInternetExplorer7PolicyList() (value string, err error) {
	retValue, err := instance.GetProperty("AllowInternetExplorer7PolicyList")
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

// SetAllowInternetExplorerStandardsMode sets the value of AllowInternetExplorerStandardsMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowInternetExplorerStandardsMode(value string) (err error) {
	return instance.SetProperty("AllowInternetExplorerStandardsMode", (value))
}

// GetAllowInternetExplorerStandardsMode gets the value of AllowInternetExplorerStandardsMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowInternetExplorerStandardsMode() (value string, err error) {
	retValue, err := instance.GetProperty("AllowInternetExplorerStandardsMode")
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

// SetAllowInternetZoneTemplate sets the value of AllowInternetZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowInternetZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowInternetZoneTemplate", (value))
}

// GetAllowInternetZoneTemplate gets the value of AllowInternetZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowInternetZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowInternetZoneTemplate")
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

// SetAllowIntranetZoneTemplate sets the value of AllowIntranetZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowIntranetZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowIntranetZoneTemplate", (value))
}

// GetAllowIntranetZoneTemplate gets the value of AllowIntranetZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowIntranetZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowIntranetZoneTemplate")
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

// SetAllowLocalMachineZoneTemplate sets the value of AllowLocalMachineZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowLocalMachineZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowLocalMachineZoneTemplate", (value))
}

// GetAllowLocalMachineZoneTemplate gets the value of AllowLocalMachineZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowLocalMachineZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowLocalMachineZoneTemplate")
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

// SetAllowLockedDownInternetZoneTemplate sets the value of AllowLockedDownInternetZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowLockedDownInternetZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowLockedDownInternetZoneTemplate", (value))
}

// GetAllowLockedDownInternetZoneTemplate gets the value of AllowLockedDownInternetZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowLockedDownInternetZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowLockedDownInternetZoneTemplate")
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

// SetAllowLockedDownIntranetZoneTemplate sets the value of AllowLockedDownIntranetZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowLockedDownIntranetZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowLockedDownIntranetZoneTemplate", (value))
}

// GetAllowLockedDownIntranetZoneTemplate gets the value of AllowLockedDownIntranetZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowLockedDownIntranetZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowLockedDownIntranetZoneTemplate")
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

// SetAllowLockedDownLocalMachineZoneTemplate sets the value of AllowLockedDownLocalMachineZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowLockedDownLocalMachineZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowLockedDownLocalMachineZoneTemplate", (value))
}

// GetAllowLockedDownLocalMachineZoneTemplate gets the value of AllowLockedDownLocalMachineZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowLockedDownLocalMachineZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowLockedDownLocalMachineZoneTemplate")
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

// SetAllowLockedDownRestrictedSitesZoneTemplate sets the value of AllowLockedDownRestrictedSitesZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowLockedDownRestrictedSitesZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowLockedDownRestrictedSitesZoneTemplate", (value))
}

// GetAllowLockedDownRestrictedSitesZoneTemplate gets the value of AllowLockedDownRestrictedSitesZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowLockedDownRestrictedSitesZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowLockedDownRestrictedSitesZoneTemplate")
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

// SetAllowOneWordEntry sets the value of AllowOneWordEntry for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowOneWordEntry(value string) (err error) {
	return instance.SetProperty("AllowOneWordEntry", (value))
}

// GetAllowOneWordEntry gets the value of AllowOneWordEntry for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowOneWordEntry() (value string, err error) {
	retValue, err := instance.GetProperty("AllowOneWordEntry")
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

// SetAllowSaveTargetAsInIEMode sets the value of AllowSaveTargetAsInIEMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowSaveTargetAsInIEMode(value string) (err error) {
	return instance.SetProperty("AllowSaveTargetAsInIEMode", (value))
}

// GetAllowSaveTargetAsInIEMode gets the value of AllowSaveTargetAsInIEMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowSaveTargetAsInIEMode() (value string, err error) {
	retValue, err := instance.GetProperty("AllowSaveTargetAsInIEMode")
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

// SetAllowSiteToZoneAssignmentList sets the value of AllowSiteToZoneAssignmentList for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowSiteToZoneAssignmentList(value string) (err error) {
	return instance.SetProperty("AllowSiteToZoneAssignmentList", (value))
}

// GetAllowSiteToZoneAssignmentList gets the value of AllowSiteToZoneAssignmentList for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowSiteToZoneAssignmentList() (value string, err error) {
	retValue, err := instance.GetProperty("AllowSiteToZoneAssignmentList")
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

// SetAllowsLockedDownTrustedSitesZoneTemplate sets the value of AllowsLockedDownTrustedSitesZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowsLockedDownTrustedSitesZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowsLockedDownTrustedSitesZoneTemplate", (value))
}

// GetAllowsLockedDownTrustedSitesZoneTemplate gets the value of AllowsLockedDownTrustedSitesZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowsLockedDownTrustedSitesZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowsLockedDownTrustedSitesZoneTemplate")
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

// SetAllowSoftwareWhenSignatureIsInvalid sets the value of AllowSoftwareWhenSignatureIsInvalid for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowSoftwareWhenSignatureIsInvalid(value string) (err error) {
	return instance.SetProperty("AllowSoftwareWhenSignatureIsInvalid", (value))
}

// GetAllowSoftwareWhenSignatureIsInvalid gets the value of AllowSoftwareWhenSignatureIsInvalid for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowSoftwareWhenSignatureIsInvalid() (value string, err error) {
	retValue, err := instance.GetProperty("AllowSoftwareWhenSignatureIsInvalid")
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

// SetAllowsRestrictedSitesZoneTemplate sets the value of AllowsRestrictedSitesZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowsRestrictedSitesZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowsRestrictedSitesZoneTemplate", (value))
}

// GetAllowsRestrictedSitesZoneTemplate gets the value of AllowsRestrictedSitesZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowsRestrictedSitesZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowsRestrictedSitesZoneTemplate")
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

// SetAllowSuggestedSites sets the value of AllowSuggestedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowSuggestedSites(value string) (err error) {
	return instance.SetProperty("AllowSuggestedSites", (value))
}

// GetAllowSuggestedSites gets the value of AllowSuggestedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowSuggestedSites() (value string, err error) {
	retValue, err := instance.GetProperty("AllowSuggestedSites")
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

// SetAllowTrustedSitesZoneTemplate sets the value of AllowTrustedSitesZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyAllowTrustedSitesZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowTrustedSitesZoneTemplate", (value))
}

// GetAllowTrustedSitesZoneTemplate gets the value of AllowTrustedSitesZoneTemplate for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyAllowTrustedSitesZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowTrustedSitesZoneTemplate")
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

// SetCheckServerCertificateRevocation sets the value of CheckServerCertificateRevocation for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyCheckServerCertificateRevocation(value string) (err error) {
	return instance.SetProperty("CheckServerCertificateRevocation", (value))
}

// GetCheckServerCertificateRevocation gets the value of CheckServerCertificateRevocation for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyCheckServerCertificateRevocation() (value string, err error) {
	retValue, err := instance.GetProperty("CheckServerCertificateRevocation")
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

// SetCheckSignaturesOnDownloadedPrograms sets the value of CheckSignaturesOnDownloadedPrograms for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyCheckSignaturesOnDownloadedPrograms(value string) (err error) {
	return instance.SetProperty("CheckSignaturesOnDownloadedPrograms", (value))
}

// GetCheckSignaturesOnDownloadedPrograms gets the value of CheckSignaturesOnDownloadedPrograms for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyCheckSignaturesOnDownloadedPrograms() (value string, err error) {
	retValue, err := instance.GetProperty("CheckSignaturesOnDownloadedPrograms")
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

// SetConfigureEdgeRedirectChannel sets the value of ConfigureEdgeRedirectChannel for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyConfigureEdgeRedirectChannel(value string) (err error) {
	return instance.SetProperty("ConfigureEdgeRedirectChannel", (value))
}

// GetConfigureEdgeRedirectChannel gets the value of ConfigureEdgeRedirectChannel for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyConfigureEdgeRedirectChannel() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigureEdgeRedirectChannel")
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

// SetConsistentMimeHandlingInternetExplorerProcesses sets the value of ConsistentMimeHandlingInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyConsistentMimeHandlingInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("ConsistentMimeHandlingInternetExplorerProcesses", (value))
}

// GetConsistentMimeHandlingInternetExplorerProcesses gets the value of ConsistentMimeHandlingInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyConsistentMimeHandlingInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("ConsistentMimeHandlingInternetExplorerProcesses")
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

// SetDisableActiveXVersionListAutoDownload sets the value of DisableActiveXVersionListAutoDownload for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableActiveXVersionListAutoDownload(value string) (err error) {
	return instance.SetProperty("DisableActiveXVersionListAutoDownload", (value))
}

// GetDisableActiveXVersionListAutoDownload gets the value of DisableActiveXVersionListAutoDownload for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableActiveXVersionListAutoDownload() (value string, err error) {
	retValue, err := instance.GetProperty("DisableActiveXVersionListAutoDownload")
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

// SetDisableBypassOfSmartScreenWarnings sets the value of DisableBypassOfSmartScreenWarnings for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableBypassOfSmartScreenWarnings(value string) (err error) {
	return instance.SetProperty("DisableBypassOfSmartScreenWarnings", (value))
}

// GetDisableBypassOfSmartScreenWarnings gets the value of DisableBypassOfSmartScreenWarnings for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableBypassOfSmartScreenWarnings() (value string, err error) {
	retValue, err := instance.GetProperty("DisableBypassOfSmartScreenWarnings")
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

// SetDisableBypassOfSmartScreenWarningsAboutUncommonFiles sets the value of DisableBypassOfSmartScreenWarningsAboutUncommonFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableBypassOfSmartScreenWarningsAboutUncommonFiles(value string) (err error) {
	return instance.SetProperty("DisableBypassOfSmartScreenWarningsAboutUncommonFiles", (value))
}

// GetDisableBypassOfSmartScreenWarningsAboutUncommonFiles gets the value of DisableBypassOfSmartScreenWarningsAboutUncommonFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableBypassOfSmartScreenWarningsAboutUncommonFiles() (value string, err error) {
	retValue, err := instance.GetProperty("DisableBypassOfSmartScreenWarningsAboutUncommonFiles")
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

// SetDisableCompatView sets the value of DisableCompatView for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableCompatView(value string) (err error) {
	return instance.SetProperty("DisableCompatView", (value))
}

// GetDisableCompatView gets the value of DisableCompatView for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableCompatView() (value string, err error) {
	retValue, err := instance.GetProperty("DisableCompatView")
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

// SetDisableConfiguringHistory sets the value of DisableConfiguringHistory for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableConfiguringHistory(value string) (err error) {
	return instance.SetProperty("DisableConfiguringHistory", (value))
}

// GetDisableConfiguringHistory gets the value of DisableConfiguringHistory for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableConfiguringHistory() (value string, err error) {
	retValue, err := instance.GetProperty("DisableConfiguringHistory")
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

// SetDisableCrashDetection sets the value of DisableCrashDetection for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableCrashDetection(value string) (err error) {
	return instance.SetProperty("DisableCrashDetection", (value))
}

// GetDisableCrashDetection gets the value of DisableCrashDetection for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableCrashDetection() (value string, err error) {
	retValue, err := instance.GetProperty("DisableCrashDetection")
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

// SetDisableCustomerExperienceImprovementProgramParticipation sets the value of DisableCustomerExperienceImprovementProgramParticipation for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableCustomerExperienceImprovementProgramParticipation(value string) (err error) {
	return instance.SetProperty("DisableCustomerExperienceImprovementProgramParticipation", (value))
}

// GetDisableCustomerExperienceImprovementProgramParticipation gets the value of DisableCustomerExperienceImprovementProgramParticipation for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableCustomerExperienceImprovementProgramParticipation() (value string, err error) {
	retValue, err := instance.GetProperty("DisableCustomerExperienceImprovementProgramParticipation")
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

// SetDisableDeletingUserVisitedWebsites sets the value of DisableDeletingUserVisitedWebsites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableDeletingUserVisitedWebsites(value string) (err error) {
	return instance.SetProperty("DisableDeletingUserVisitedWebsites", (value))
}

// GetDisableDeletingUserVisitedWebsites gets the value of DisableDeletingUserVisitedWebsites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableDeletingUserVisitedWebsites() (value string, err error) {
	retValue, err := instance.GetProperty("DisableDeletingUserVisitedWebsites")
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

// SetDisableEnclosureDownloading sets the value of DisableEnclosureDownloading for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableEnclosureDownloading(value string) (err error) {
	return instance.SetProperty("DisableEnclosureDownloading", (value))
}

// GetDisableEnclosureDownloading gets the value of DisableEnclosureDownloading for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableEnclosureDownloading() (value string, err error) {
	retValue, err := instance.GetProperty("DisableEnclosureDownloading")
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

// SetDisableEncryptionSupport sets the value of DisableEncryptionSupport for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableEncryptionSupport(value string) (err error) {
	return instance.SetProperty("DisableEncryptionSupport", (value))
}

// GetDisableEncryptionSupport gets the value of DisableEncryptionSupport for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableEncryptionSupport() (value string, err error) {
	retValue, err := instance.GetProperty("DisableEncryptionSupport")
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

// SetDisableFeedsBackgroundSync sets the value of DisableFeedsBackgroundSync for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableFeedsBackgroundSync(value string) (err error) {
	return instance.SetProperty("DisableFeedsBackgroundSync", (value))
}

// GetDisableFeedsBackgroundSync gets the value of DisableFeedsBackgroundSync for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableFeedsBackgroundSync() (value string, err error) {
	retValue, err := instance.GetProperty("DisableFeedsBackgroundSync")
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

// SetDisableFirstRunWizard sets the value of DisableFirstRunWizard for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableFirstRunWizard(value string) (err error) {
	return instance.SetProperty("DisableFirstRunWizard", (value))
}

// GetDisableFirstRunWizard gets the value of DisableFirstRunWizard for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableFirstRunWizard() (value string, err error) {
	retValue, err := instance.GetProperty("DisableFirstRunWizard")
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

// SetDisableFlipAheadFeature sets the value of DisableFlipAheadFeature for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableFlipAheadFeature(value string) (err error) {
	return instance.SetProperty("DisableFlipAheadFeature", (value))
}

// GetDisableFlipAheadFeature gets the value of DisableFlipAheadFeature for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableFlipAheadFeature() (value string, err error) {
	retValue, err := instance.GetProperty("DisableFlipAheadFeature")
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

// SetDisableGeolocation sets the value of DisableGeolocation for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableGeolocation(value string) (err error) {
	return instance.SetProperty("DisableGeolocation", (value))
}

// GetDisableGeolocation gets the value of DisableGeolocation for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableGeolocation() (value string, err error) {
	retValue, err := instance.GetProperty("DisableGeolocation")
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

// SetDisableHomePageChange sets the value of DisableHomePageChange for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableHomePageChange(value string) (err error) {
	return instance.SetProperty("DisableHomePageChange", (value))
}

// GetDisableHomePageChange gets the value of DisableHomePageChange for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableHomePageChange() (value string, err error) {
	retValue, err := instance.GetProperty("DisableHomePageChange")
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

// SetDisableIgnoringCertificateErrors sets the value of DisableIgnoringCertificateErrors for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableIgnoringCertificateErrors(value string) (err error) {
	return instance.SetProperty("DisableIgnoringCertificateErrors", (value))
}

// GetDisableIgnoringCertificateErrors gets the value of DisableIgnoringCertificateErrors for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableIgnoringCertificateErrors() (value string, err error) {
	retValue, err := instance.GetProperty("DisableIgnoringCertificateErrors")
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

// SetDisableInPrivateBrowsing sets the value of DisableInPrivateBrowsing for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableInPrivateBrowsing(value string) (err error) {
	return instance.SetProperty("DisableInPrivateBrowsing", (value))
}

// GetDisableInPrivateBrowsing gets the value of DisableInPrivateBrowsing for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableInPrivateBrowsing() (value string, err error) {
	retValue, err := instance.GetProperty("DisableInPrivateBrowsing")
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

// SetDisableInternetExplorerApp sets the value of DisableInternetExplorerApp for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableInternetExplorerApp(value string) (err error) {
	return instance.SetProperty("DisableInternetExplorerApp", (value))
}

// GetDisableInternetExplorerApp gets the value of DisableInternetExplorerApp for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableInternetExplorerApp() (value string, err error) {
	retValue, err := instance.GetProperty("DisableInternetExplorerApp")
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

// SetDisableProcessesInEnhancedProtectedMode sets the value of DisableProcessesInEnhancedProtectedMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableProcessesInEnhancedProtectedMode(value string) (err error) {
	return instance.SetProperty("DisableProcessesInEnhancedProtectedMode", (value))
}

// GetDisableProcessesInEnhancedProtectedMode gets the value of DisableProcessesInEnhancedProtectedMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableProcessesInEnhancedProtectedMode() (value string, err error) {
	retValue, err := instance.GetProperty("DisableProcessesInEnhancedProtectedMode")
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

// SetDisableProxyChange sets the value of DisableProxyChange for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableProxyChange(value string) (err error) {
	return instance.SetProperty("DisableProxyChange", (value))
}

// GetDisableProxyChange gets the value of DisableProxyChange for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableProxyChange() (value string, err error) {
	retValue, err := instance.GetProperty("DisableProxyChange")
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

// SetDisableSearchProviderChange sets the value of DisableSearchProviderChange for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableSearchProviderChange(value string) (err error) {
	return instance.SetProperty("DisableSearchProviderChange", (value))
}

// GetDisableSearchProviderChange gets the value of DisableSearchProviderChange for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableSearchProviderChange() (value string, err error) {
	retValue, err := instance.GetProperty("DisableSearchProviderChange")
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

// SetDisableSecondaryHomePageChange sets the value of DisableSecondaryHomePageChange for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableSecondaryHomePageChange(value string) (err error) {
	return instance.SetProperty("DisableSecondaryHomePageChange", (value))
}

// GetDisableSecondaryHomePageChange gets the value of DisableSecondaryHomePageChange for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableSecondaryHomePageChange() (value string, err error) {
	retValue, err := instance.GetProperty("DisableSecondaryHomePageChange")
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

// SetDisableSecuritySettingsCheck sets the value of DisableSecuritySettingsCheck for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableSecuritySettingsCheck(value string) (err error) {
	return instance.SetProperty("DisableSecuritySettingsCheck", (value))
}

// GetDisableSecuritySettingsCheck gets the value of DisableSecuritySettingsCheck for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableSecuritySettingsCheck() (value string, err error) {
	retValue, err := instance.GetProperty("DisableSecuritySettingsCheck")
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

// SetDisableWebAddressAutoComplete sets the value of DisableWebAddressAutoComplete for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDisableWebAddressAutoComplete(value string) (err error) {
	return instance.SetProperty("DisableWebAddressAutoComplete", (value))
}

// GetDisableWebAddressAutoComplete gets the value of DisableWebAddressAutoComplete for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDisableWebAddressAutoComplete() (value string, err error) {
	retValue, err := instance.GetProperty("DisableWebAddressAutoComplete")
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

// SetDoNotAllowActiveXControlsInProtectedMode sets the value of DoNotAllowActiveXControlsInProtectedMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDoNotAllowActiveXControlsInProtectedMode(value string) (err error) {
	return instance.SetProperty("DoNotAllowActiveXControlsInProtectedMode", (value))
}

// GetDoNotAllowActiveXControlsInProtectedMode gets the value of DoNotAllowActiveXControlsInProtectedMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDoNotAllowActiveXControlsInProtectedMode() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotAllowActiveXControlsInProtectedMode")
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

// SetDoNotBlockOutdatedActiveXControls sets the value of DoNotBlockOutdatedActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDoNotBlockOutdatedActiveXControls(value string) (err error) {
	return instance.SetProperty("DoNotBlockOutdatedActiveXControls", (value))
}

// GetDoNotBlockOutdatedActiveXControls gets the value of DoNotBlockOutdatedActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDoNotBlockOutdatedActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotBlockOutdatedActiveXControls")
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

// SetDoNotBlockOutdatedActiveXControlsOnSpecificDomains sets the value of DoNotBlockOutdatedActiveXControlsOnSpecificDomains for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyDoNotBlockOutdatedActiveXControlsOnSpecificDomains(value string) (err error) {
	return instance.SetProperty("DoNotBlockOutdatedActiveXControlsOnSpecificDomains", (value))
}

// GetDoNotBlockOutdatedActiveXControlsOnSpecificDomains gets the value of DoNotBlockOutdatedActiveXControlsOnSpecificDomains for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyDoNotBlockOutdatedActiveXControlsOnSpecificDomains() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotBlockOutdatedActiveXControlsOnSpecificDomains")
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

// SetEnableExtendedIEModeHotkeys sets the value of EnableExtendedIEModeHotkeys for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyEnableExtendedIEModeHotkeys(value string) (err error) {
	return instance.SetProperty("EnableExtendedIEModeHotkeys", (value))
}

// GetEnableExtendedIEModeHotkeys gets the value of EnableExtendedIEModeHotkeys for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyEnableExtendedIEModeHotkeys() (value string, err error) {
	retValue, err := instance.GetProperty("EnableExtendedIEModeHotkeys")
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

// SetIncludeAllLocalSites sets the value of IncludeAllLocalSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIncludeAllLocalSites(value string) (err error) {
	return instance.SetProperty("IncludeAllLocalSites", (value))
}

// GetIncludeAllLocalSites gets the value of IncludeAllLocalSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIncludeAllLocalSites() (value string, err error) {
	retValue, err := instance.GetProperty("IncludeAllLocalSites")
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

// SetIncludeAllNetworkPaths sets the value of IncludeAllNetworkPaths for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIncludeAllNetworkPaths(value string) (err error) {
	return instance.SetProperty("IncludeAllNetworkPaths", (value))
}

// GetIncludeAllNetworkPaths gets the value of IncludeAllNetworkPaths for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIncludeAllNetworkPaths() (value string, err error) {
	retValue, err := instance.GetProperty("IncludeAllNetworkPaths")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInstanceID() (value string, err error) {
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

// SetInternetZoneAllowAccessToDataSources sets the value of InternetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowAccessToDataSources", (value))
}

// GetInternetZoneAllowAccessToDataSources gets the value of InternetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowAccessToDataSources")
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

// SetInternetZoneAllowAutomaticPromptingForActiveXControls sets the value of InternetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowAutomaticPromptingForActiveXControls", (value))
}

// GetInternetZoneAllowAutomaticPromptingForActiveXControls gets the value of InternetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowAutomaticPromptingForActiveXControls")
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

// SetInternetZoneAllowAutomaticPromptingForFileDownloads sets the value of InternetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowAutomaticPromptingForFileDownloads", (value))
}

// GetInternetZoneAllowAutomaticPromptingForFileDownloads gets the value of InternetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowAutomaticPromptingForFileDownloads")
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

// SetInternetZoneAllowCopyPasteViaScript sets the value of InternetZoneAllowCopyPasteViaScript for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowCopyPasteViaScript(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowCopyPasteViaScript", (value))
}

// GetInternetZoneAllowCopyPasteViaScript gets the value of InternetZoneAllowCopyPasteViaScript for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowCopyPasteViaScript() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowCopyPasteViaScript")
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

// SetInternetZoneAllowDragAndDropCopyAndPasteFiles sets the value of InternetZoneAllowDragAndDropCopyAndPasteFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowDragAndDropCopyAndPasteFiles(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowDragAndDropCopyAndPasteFiles", (value))
}

// GetInternetZoneAllowDragAndDropCopyAndPasteFiles gets the value of InternetZoneAllowDragAndDropCopyAndPasteFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowDragAndDropCopyAndPasteFiles() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowDragAndDropCopyAndPasteFiles")
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

// SetInternetZoneAllowFontDownloads sets the value of InternetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowFontDownloads", (value))
}

// GetInternetZoneAllowFontDownloads gets the value of InternetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowFontDownloads")
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

// SetInternetZoneAllowLessPrivilegedSites sets the value of InternetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowLessPrivilegedSites", (value))
}

// GetInternetZoneAllowLessPrivilegedSites gets the value of InternetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowLessPrivilegedSites")
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

// SetInternetZoneAllowLoadingOfXAMLFiles sets the value of InternetZoneAllowLoadingOfXAMLFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowLoadingOfXAMLFiles(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowLoadingOfXAMLFiles", (value))
}

// GetInternetZoneAllowLoadingOfXAMLFiles gets the value of InternetZoneAllowLoadingOfXAMLFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowLoadingOfXAMLFiles() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowLoadingOfXAMLFiles")
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

// SetInternetZoneAllowNETFrameworkReliantComponents sets the value of InternetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowNETFrameworkReliantComponents", (value))
}

// GetInternetZoneAllowNETFrameworkReliantComponents gets the value of InternetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowNETFrameworkReliantComponents")
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

// SetInternetZoneAllowOnlyApprovedDomainsToUseActiveXControls sets the value of InternetZoneAllowOnlyApprovedDomainsToUseActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowOnlyApprovedDomainsToUseActiveXControls(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowOnlyApprovedDomainsToUseActiveXControls", (value))
}

// GetInternetZoneAllowOnlyApprovedDomainsToUseActiveXControls gets the value of InternetZoneAllowOnlyApprovedDomainsToUseActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowOnlyApprovedDomainsToUseActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowOnlyApprovedDomainsToUseActiveXControls")
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

// SetInternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl sets the value of InternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl", (value))
}

// GetInternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl gets the value of InternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl")
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

// SetInternetZoneAllowScriptingOfInternetExplorerWebBrowserControls sets the value of InternetZoneAllowScriptingOfInternetExplorerWebBrowserControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowScriptingOfInternetExplorerWebBrowserControls(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowScriptingOfInternetExplorerWebBrowserControls", (value))
}

// GetInternetZoneAllowScriptingOfInternetExplorerWebBrowserControls gets the value of InternetZoneAllowScriptingOfInternetExplorerWebBrowserControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowScriptingOfInternetExplorerWebBrowserControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowScriptingOfInternetExplorerWebBrowserControls")
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

// SetInternetZoneAllowScriptInitiatedWindows sets the value of InternetZoneAllowScriptInitiatedWindows for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowScriptInitiatedWindows(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowScriptInitiatedWindows", (value))
}

// GetInternetZoneAllowScriptInitiatedWindows gets the value of InternetZoneAllowScriptInitiatedWindows for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowScriptInitiatedWindows() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowScriptInitiatedWindows")
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

// SetInternetZoneAllowScriptlets sets the value of InternetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowScriptlets", (value))
}

// GetInternetZoneAllowScriptlets gets the value of InternetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowScriptlets")
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

// SetInternetZoneAllowSmartScreenIE sets the value of InternetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowSmartScreenIE", (value))
}

// GetInternetZoneAllowSmartScreenIE gets the value of InternetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowSmartScreenIE")
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

// SetInternetZoneAllowUpdatesToStatusBarViaScript sets the value of InternetZoneAllowUpdatesToStatusBarViaScript for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowUpdatesToStatusBarViaScript(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowUpdatesToStatusBarViaScript", (value))
}

// GetInternetZoneAllowUpdatesToStatusBarViaScript gets the value of InternetZoneAllowUpdatesToStatusBarViaScript for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowUpdatesToStatusBarViaScript() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowUpdatesToStatusBarViaScript")
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

// SetInternetZoneAllowUserDataPersistence sets the value of InternetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowUserDataPersistence", (value))
}

// GetInternetZoneAllowUserDataPersistence gets the value of InternetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowUserDataPersistence")
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

// SetInternetZoneAllowVBScriptToRunInInternetExplorer sets the value of InternetZoneAllowVBScriptToRunInInternetExplorer for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneAllowVBScriptToRunInInternetExplorer(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowVBScriptToRunInInternetExplorer", (value))
}

// GetInternetZoneAllowVBScriptToRunInInternetExplorer gets the value of InternetZoneAllowVBScriptToRunInInternetExplorer for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneAllowVBScriptToRunInInternetExplorer() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowVBScriptToRunInInternetExplorer")
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

// SetInternetZoneDoNotRunAntimalwareAgainstActiveXControls sets the value of InternetZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneDoNotRunAntimalwareAgainstActiveXControls(value string) (err error) {
	return instance.SetProperty("InternetZoneDoNotRunAntimalwareAgainstActiveXControls", (value))
}

// GetInternetZoneDoNotRunAntimalwareAgainstActiveXControls gets the value of InternetZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneDoNotRunAntimalwareAgainstActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneDoNotRunAntimalwareAgainstActiveXControls")
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

// SetInternetZoneDownloadSignedActiveXControls sets the value of InternetZoneDownloadSignedActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneDownloadSignedActiveXControls(value string) (err error) {
	return instance.SetProperty("InternetZoneDownloadSignedActiveXControls", (value))
}

// GetInternetZoneDownloadSignedActiveXControls gets the value of InternetZoneDownloadSignedActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneDownloadSignedActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneDownloadSignedActiveXControls")
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

// SetInternetZoneDownloadUnsignedActiveXControls sets the value of InternetZoneDownloadUnsignedActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneDownloadUnsignedActiveXControls(value string) (err error) {
	return instance.SetProperty("InternetZoneDownloadUnsignedActiveXControls", (value))
}

// GetInternetZoneDownloadUnsignedActiveXControls gets the value of InternetZoneDownloadUnsignedActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneDownloadUnsignedActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneDownloadUnsignedActiveXControls")
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

// SetInternetZoneEnableCrossSiteScriptingFilter sets the value of InternetZoneEnableCrossSiteScriptingFilter for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneEnableCrossSiteScriptingFilter(value string) (err error) {
	return instance.SetProperty("InternetZoneEnableCrossSiteScriptingFilter", (value))
}

// GetInternetZoneEnableCrossSiteScriptingFilter gets the value of InternetZoneEnableCrossSiteScriptingFilter for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneEnableCrossSiteScriptingFilter() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneEnableCrossSiteScriptingFilter")
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

// SetInternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows sets the value of InternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows(value string) (err error) {
	return instance.SetProperty("InternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows", (value))
}

// GetInternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows gets the value of InternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows")
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

// SetInternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows sets the value of InternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows(value string) (err error) {
	return instance.SetProperty("InternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows", (value))
}

// GetInternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows gets the value of InternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows")
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

// SetInternetZoneEnableMIMESniffing sets the value of InternetZoneEnableMIMESniffing for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneEnableMIMESniffing(value string) (err error) {
	return instance.SetProperty("InternetZoneEnableMIMESniffing", (value))
}

// GetInternetZoneEnableMIMESniffing gets the value of InternetZoneEnableMIMESniffing for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneEnableMIMESniffing() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneEnableMIMESniffing")
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

// SetInternetZoneEnableProtectedMode sets the value of InternetZoneEnableProtectedMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneEnableProtectedMode(value string) (err error) {
	return instance.SetProperty("InternetZoneEnableProtectedMode", (value))
}

// GetInternetZoneEnableProtectedMode gets the value of InternetZoneEnableProtectedMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneEnableProtectedMode() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneEnableProtectedMode")
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

// SetInternetZoneIncludeLocalPathWhenUploadingFilesToServer sets the value of InternetZoneIncludeLocalPathWhenUploadingFilesToServer for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneIncludeLocalPathWhenUploadingFilesToServer(value string) (err error) {
	return instance.SetProperty("InternetZoneIncludeLocalPathWhenUploadingFilesToServer", (value))
}

// GetInternetZoneIncludeLocalPathWhenUploadingFilesToServer gets the value of InternetZoneIncludeLocalPathWhenUploadingFilesToServer for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneIncludeLocalPathWhenUploadingFilesToServer() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneIncludeLocalPathWhenUploadingFilesToServer")
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

// SetInternetZoneInitializeAndScriptActiveXControls sets the value of InternetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("InternetZoneInitializeAndScriptActiveXControls", (value))
}

// GetInternetZoneInitializeAndScriptActiveXControls gets the value of InternetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneInitializeAndScriptActiveXControls")
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

// SetInternetZoneJavaPermissions sets the value of InternetZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("InternetZoneJavaPermissions", (value))
}

// GetInternetZoneJavaPermissions gets the value of InternetZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneJavaPermissions")
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

// SetInternetZoneLaunchingApplicationsAndFilesInIFRAME sets the value of InternetZoneLaunchingApplicationsAndFilesInIFRAME for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneLaunchingApplicationsAndFilesInIFRAME(value string) (err error) {
	return instance.SetProperty("InternetZoneLaunchingApplicationsAndFilesInIFRAME", (value))
}

// GetInternetZoneLaunchingApplicationsAndFilesInIFRAME gets the value of InternetZoneLaunchingApplicationsAndFilesInIFRAME for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneLaunchingApplicationsAndFilesInIFRAME() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneLaunchingApplicationsAndFilesInIFRAME")
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

// SetInternetZoneLogonOptions sets the value of InternetZoneLogonOptions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneLogonOptions(value string) (err error) {
	return instance.SetProperty("InternetZoneLogonOptions", (value))
}

// GetInternetZoneLogonOptions gets the value of InternetZoneLogonOptions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneLogonOptions() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneLogonOptions")
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

// SetInternetZoneNavigateWindowsAndFrames sets the value of InternetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("InternetZoneNavigateWindowsAndFrames", (value))
}

// GetInternetZoneNavigateWindowsAndFrames gets the value of InternetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneNavigateWindowsAndFrames")
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

// SetInternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode sets the value of InternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode(value string) (err error) {
	return instance.SetProperty("InternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode", (value))
}

// GetInternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode gets the value of InternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode")
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

// SetInternetZoneShowSecurityWarningForPotentiallyUnsafeFiles sets the value of InternetZoneShowSecurityWarningForPotentiallyUnsafeFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneShowSecurityWarningForPotentiallyUnsafeFiles(value string) (err error) {
	return instance.SetProperty("InternetZoneShowSecurityWarningForPotentiallyUnsafeFiles", (value))
}

// GetInternetZoneShowSecurityWarningForPotentiallyUnsafeFiles gets the value of InternetZoneShowSecurityWarningForPotentiallyUnsafeFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneShowSecurityWarningForPotentiallyUnsafeFiles() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneShowSecurityWarningForPotentiallyUnsafeFiles")
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

// SetInternetZoneUsePopupBlocker sets the value of InternetZoneUsePopupBlocker for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyInternetZoneUsePopupBlocker(value string) (err error) {
	return instance.SetProperty("InternetZoneUsePopupBlocker", (value))
}

// GetInternetZoneUsePopupBlocker gets the value of InternetZoneUsePopupBlocker for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyInternetZoneUsePopupBlocker() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneUsePopupBlocker")
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

// SetIntranetZoneAllowAccessToDataSources sets the value of IntranetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIntranetZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowAccessToDataSources", (value))
}

// GetIntranetZoneAllowAccessToDataSources gets the value of IntranetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIntranetZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowAccessToDataSources")
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

// SetIntranetZoneAllowAutomaticPromptingForActiveXControls sets the value of IntranetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIntranetZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowAutomaticPromptingForActiveXControls", (value))
}

// GetIntranetZoneAllowAutomaticPromptingForActiveXControls gets the value of IntranetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIntranetZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowAutomaticPromptingForActiveXControls")
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

// SetIntranetZoneAllowAutomaticPromptingForFileDownloads sets the value of IntranetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIntranetZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowAutomaticPromptingForFileDownloads", (value))
}

// GetIntranetZoneAllowAutomaticPromptingForFileDownloads gets the value of IntranetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIntranetZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowAutomaticPromptingForFileDownloads")
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

// SetIntranetZoneAllowFontDownloads sets the value of IntranetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIntranetZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowFontDownloads", (value))
}

// GetIntranetZoneAllowFontDownloads gets the value of IntranetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIntranetZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowFontDownloads")
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

// SetIntranetZoneAllowLessPrivilegedSites sets the value of IntranetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIntranetZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowLessPrivilegedSites", (value))
}

// GetIntranetZoneAllowLessPrivilegedSites gets the value of IntranetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIntranetZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowLessPrivilegedSites")
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

// SetIntranetZoneAllowNETFrameworkReliantComponents sets the value of IntranetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIntranetZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowNETFrameworkReliantComponents", (value))
}

// GetIntranetZoneAllowNETFrameworkReliantComponents gets the value of IntranetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIntranetZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowNETFrameworkReliantComponents")
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

// SetIntranetZoneAllowScriptlets sets the value of IntranetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIntranetZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowScriptlets", (value))
}

// GetIntranetZoneAllowScriptlets gets the value of IntranetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIntranetZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowScriptlets")
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

// SetIntranetZoneAllowSmartScreenIE sets the value of IntranetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIntranetZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowSmartScreenIE", (value))
}

// GetIntranetZoneAllowSmartScreenIE gets the value of IntranetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIntranetZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowSmartScreenIE")
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

// SetIntranetZoneAllowUserDataPersistence sets the value of IntranetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIntranetZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowUserDataPersistence", (value))
}

// GetIntranetZoneAllowUserDataPersistence gets the value of IntranetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIntranetZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowUserDataPersistence")
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

// SetIntranetZoneDoNotRunAntimalwareAgainstActiveXControls sets the value of IntranetZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIntranetZoneDoNotRunAntimalwareAgainstActiveXControls(value string) (err error) {
	return instance.SetProperty("IntranetZoneDoNotRunAntimalwareAgainstActiveXControls", (value))
}

// GetIntranetZoneDoNotRunAntimalwareAgainstActiveXControls gets the value of IntranetZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIntranetZoneDoNotRunAntimalwareAgainstActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneDoNotRunAntimalwareAgainstActiveXControls")
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

// SetIntranetZoneInitializeAndScriptActiveXControls sets the value of IntranetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIntranetZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("IntranetZoneInitializeAndScriptActiveXControls", (value))
}

// GetIntranetZoneInitializeAndScriptActiveXControls gets the value of IntranetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIntranetZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneInitializeAndScriptActiveXControls")
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

// SetIntranetZoneJavaPermissions sets the value of IntranetZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIntranetZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("IntranetZoneJavaPermissions", (value))
}

// GetIntranetZoneJavaPermissions gets the value of IntranetZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIntranetZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneJavaPermissions")
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

// SetIntranetZoneNavigateWindowsAndFrames sets the value of IntranetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyIntranetZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("IntranetZoneNavigateWindowsAndFrames", (value))
}

// GetIntranetZoneNavigateWindowsAndFrames gets the value of IntranetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyIntranetZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneNavigateWindowsAndFrames")
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

// SetJScriptReplacement sets the value of JScriptReplacement for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyJScriptReplacement(value string) (err error) {
	return instance.SetProperty("JScriptReplacement", (value))
}

// GetJScriptReplacement gets the value of JScriptReplacement for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyJScriptReplacement() (value string, err error) {
	retValue, err := instance.GetProperty("JScriptReplacement")
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

// SetKeepIntranetSitesInInternetExplorer sets the value of KeepIntranetSitesInInternetExplorer for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyKeepIntranetSitesInInternetExplorer(value string) (err error) {
	return instance.SetProperty("KeepIntranetSitesInInternetExplorer", (value))
}

// GetKeepIntranetSitesInInternetExplorer gets the value of KeepIntranetSitesInInternetExplorer for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyKeepIntranetSitesInInternetExplorer() (value string, err error) {
	retValue, err := instance.GetProperty("KeepIntranetSitesInInternetExplorer")
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

// SetLocalMachineZoneAllowAccessToDataSources sets the value of LocalMachineZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLocalMachineZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowAccessToDataSources", (value))
}

// GetLocalMachineZoneAllowAccessToDataSources gets the value of LocalMachineZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLocalMachineZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowAccessToDataSources")
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

// SetLocalMachineZoneAllowAutomaticPromptingForActiveXControls sets the value of LocalMachineZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLocalMachineZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowAutomaticPromptingForActiveXControls", (value))
}

// GetLocalMachineZoneAllowAutomaticPromptingForActiveXControls gets the value of LocalMachineZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLocalMachineZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowAutomaticPromptingForActiveXControls")
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

// SetLocalMachineZoneAllowAutomaticPromptingForFileDownloads sets the value of LocalMachineZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLocalMachineZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowAutomaticPromptingForFileDownloads", (value))
}

// GetLocalMachineZoneAllowAutomaticPromptingForFileDownloads gets the value of LocalMachineZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLocalMachineZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowAutomaticPromptingForFileDownloads")
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

// SetLocalMachineZoneAllowFontDownloads sets the value of LocalMachineZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLocalMachineZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowFontDownloads", (value))
}

// GetLocalMachineZoneAllowFontDownloads gets the value of LocalMachineZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLocalMachineZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowFontDownloads")
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

// SetLocalMachineZoneAllowLessPrivilegedSites sets the value of LocalMachineZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLocalMachineZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowLessPrivilegedSites", (value))
}

// GetLocalMachineZoneAllowLessPrivilegedSites gets the value of LocalMachineZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLocalMachineZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowLessPrivilegedSites")
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

// SetLocalMachineZoneAllowNETFrameworkReliantComponents sets the value of LocalMachineZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLocalMachineZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowNETFrameworkReliantComponents", (value))
}

// GetLocalMachineZoneAllowNETFrameworkReliantComponents gets the value of LocalMachineZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLocalMachineZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowNETFrameworkReliantComponents")
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

// SetLocalMachineZoneAllowScriptlets sets the value of LocalMachineZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLocalMachineZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowScriptlets", (value))
}

// GetLocalMachineZoneAllowScriptlets gets the value of LocalMachineZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLocalMachineZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowScriptlets")
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

// SetLocalMachineZoneAllowSmartScreenIE sets the value of LocalMachineZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLocalMachineZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowSmartScreenIE", (value))
}

// GetLocalMachineZoneAllowSmartScreenIE gets the value of LocalMachineZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLocalMachineZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowSmartScreenIE")
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

// SetLocalMachineZoneAllowUserDataPersistence sets the value of LocalMachineZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLocalMachineZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowUserDataPersistence", (value))
}

// GetLocalMachineZoneAllowUserDataPersistence gets the value of LocalMachineZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLocalMachineZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowUserDataPersistence")
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

// SetLocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls sets the value of LocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls", (value))
}

// GetLocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls gets the value of LocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls")
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

// SetLocalMachineZoneInitializeAndScriptActiveXControls sets the value of LocalMachineZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLocalMachineZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneInitializeAndScriptActiveXControls", (value))
}

// GetLocalMachineZoneInitializeAndScriptActiveXControls gets the value of LocalMachineZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLocalMachineZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneInitializeAndScriptActiveXControls")
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

// SetLocalMachineZoneJavaPermissions sets the value of LocalMachineZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLocalMachineZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneJavaPermissions", (value))
}

// GetLocalMachineZoneJavaPermissions gets the value of LocalMachineZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLocalMachineZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneJavaPermissions")
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

// SetLocalMachineZoneNavigateWindowsAndFrames sets the value of LocalMachineZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLocalMachineZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneNavigateWindowsAndFrames", (value))
}

// GetLocalMachineZoneNavigateWindowsAndFrames gets the value of LocalMachineZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLocalMachineZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneNavigateWindowsAndFrames")
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

// SetLockedDownInternetZoneAllowAccessToDataSources sets the value of LockedDownInternetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowAccessToDataSources", (value))
}

// GetLockedDownInternetZoneAllowAccessToDataSources gets the value of LockedDownInternetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowAccessToDataSources")
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

// SetLockedDownInternetZoneAllowAutomaticPromptingForActiveXControls sets the value of LockedDownInternetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowAutomaticPromptingForActiveXControls", (value))
}

// GetLockedDownInternetZoneAllowAutomaticPromptingForActiveXControls gets the value of LockedDownInternetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowAutomaticPromptingForActiveXControls")
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

// SetLockedDownInternetZoneAllowAutomaticPromptingForFileDownloads sets the value of LockedDownInternetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowAutomaticPromptingForFileDownloads", (value))
}

// GetLockedDownInternetZoneAllowAutomaticPromptingForFileDownloads gets the value of LockedDownInternetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowAutomaticPromptingForFileDownloads")
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

// SetLockedDownInternetZoneAllowFontDownloads sets the value of LockedDownInternetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowFontDownloads", (value))
}

// GetLockedDownInternetZoneAllowFontDownloads gets the value of LockedDownInternetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowFontDownloads")
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

// SetLockedDownInternetZoneAllowLessPrivilegedSites sets the value of LockedDownInternetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowLessPrivilegedSites", (value))
}

// GetLockedDownInternetZoneAllowLessPrivilegedSites gets the value of LockedDownInternetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowLessPrivilegedSites")
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

// SetLockedDownInternetZoneAllowNETFrameworkReliantComponents sets the value of LockedDownInternetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowNETFrameworkReliantComponents", (value))
}

// GetLockedDownInternetZoneAllowNETFrameworkReliantComponents gets the value of LockedDownInternetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowNETFrameworkReliantComponents")
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

// SetLockedDownInternetZoneAllowScriptlets sets the value of LockedDownInternetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowScriptlets", (value))
}

// GetLockedDownInternetZoneAllowScriptlets gets the value of LockedDownInternetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowScriptlets")
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

// SetLockedDownInternetZoneAllowSmartScreenIE sets the value of LockedDownInternetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowSmartScreenIE", (value))
}

// GetLockedDownInternetZoneAllowSmartScreenIE gets the value of LockedDownInternetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowSmartScreenIE")
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

// SetLockedDownInternetZoneAllowUserDataPersistence sets the value of LockedDownInternetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowUserDataPersistence", (value))
}

// GetLockedDownInternetZoneAllowUserDataPersistence gets the value of LockedDownInternetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowUserDataPersistence")
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

// SetLockedDownInternetZoneInitializeAndScriptActiveXControls sets the value of LockedDownInternetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownInternetZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneInitializeAndScriptActiveXControls", (value))
}

// GetLockedDownInternetZoneInitializeAndScriptActiveXControls gets the value of LockedDownInternetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownInternetZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneInitializeAndScriptActiveXControls")
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

// SetLockedDownInternetZoneJavaPermissions sets the value of LockedDownInternetZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownInternetZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneJavaPermissions", (value))
}

// GetLockedDownInternetZoneJavaPermissions gets the value of LockedDownInternetZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownInternetZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneJavaPermissions")
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

// SetLockedDownInternetZoneNavigateWindowsAndFrames sets the value of LockedDownInternetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownInternetZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneNavigateWindowsAndFrames", (value))
}

// GetLockedDownInternetZoneNavigateWindowsAndFrames gets the value of LockedDownInternetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownInternetZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneNavigateWindowsAndFrames")
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

// SetLockedDownIntranetJavaPermissions sets the value of LockedDownIntranetJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownIntranetJavaPermissions(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetJavaPermissions", (value))
}

// GetLockedDownIntranetJavaPermissions gets the value of LockedDownIntranetJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownIntranetJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetJavaPermissions")
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

// SetLockedDownIntranetZoneAllowAccessToDataSources sets the value of LockedDownIntranetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowAccessToDataSources", (value))
}

// GetLockedDownIntranetZoneAllowAccessToDataSources gets the value of LockedDownIntranetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowAccessToDataSources")
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

// SetLockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls sets the value of LockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls", (value))
}

// GetLockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls gets the value of LockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls")
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

// SetLockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads sets the value of LockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads", (value))
}

// GetLockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads gets the value of LockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads")
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

// SetLockedDownIntranetZoneAllowFontDownloads sets the value of LockedDownIntranetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowFontDownloads", (value))
}

// GetLockedDownIntranetZoneAllowFontDownloads gets the value of LockedDownIntranetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowFontDownloads")
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

// SetLockedDownIntranetZoneAllowLessPrivilegedSites sets the value of LockedDownIntranetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowLessPrivilegedSites", (value))
}

// GetLockedDownIntranetZoneAllowLessPrivilegedSites gets the value of LockedDownIntranetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowLessPrivilegedSites")
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

// SetLockedDownIntranetZoneAllowNETFrameworkReliantComponents sets the value of LockedDownIntranetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowNETFrameworkReliantComponents", (value))
}

// GetLockedDownIntranetZoneAllowNETFrameworkReliantComponents gets the value of LockedDownIntranetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowNETFrameworkReliantComponents")
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

// SetLockedDownIntranetZoneAllowScriptlets sets the value of LockedDownIntranetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowScriptlets", (value))
}

// GetLockedDownIntranetZoneAllowScriptlets gets the value of LockedDownIntranetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowScriptlets")
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

// SetLockedDownIntranetZoneAllowSmartScreenIE sets the value of LockedDownIntranetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowSmartScreenIE", (value))
}

// GetLockedDownIntranetZoneAllowSmartScreenIE gets the value of LockedDownIntranetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowSmartScreenIE")
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

// SetLockedDownIntranetZoneAllowUserDataPersistence sets the value of LockedDownIntranetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowUserDataPersistence", (value))
}

// GetLockedDownIntranetZoneAllowUserDataPersistence gets the value of LockedDownIntranetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowUserDataPersistence")
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

// SetLockedDownIntranetZoneInitializeAndScriptActiveXControls sets the value of LockedDownIntranetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownIntranetZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneInitializeAndScriptActiveXControls", (value))
}

// GetLockedDownIntranetZoneInitializeAndScriptActiveXControls gets the value of LockedDownIntranetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownIntranetZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneInitializeAndScriptActiveXControls")
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

// SetLockedDownIntranetZoneNavigateWindowsAndFrames sets the value of LockedDownIntranetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownIntranetZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneNavigateWindowsAndFrames", (value))
}

// GetLockedDownIntranetZoneNavigateWindowsAndFrames gets the value of LockedDownIntranetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownIntranetZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneNavigateWindowsAndFrames")
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

// SetLockedDownLocalMachineZoneAllowAccessToDataSources sets the value of LockedDownLocalMachineZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowAccessToDataSources", (value))
}

// GetLockedDownLocalMachineZoneAllowAccessToDataSources gets the value of LockedDownLocalMachineZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowAccessToDataSources")
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

// SetLockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls sets the value of LockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls", (value))
}

// GetLockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls gets the value of LockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls")
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

// SetLockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads sets the value of LockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads", (value))
}

// GetLockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads gets the value of LockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads")
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

// SetLockedDownLocalMachineZoneAllowFontDownloads sets the value of LockedDownLocalMachineZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowFontDownloads", (value))
}

// GetLockedDownLocalMachineZoneAllowFontDownloads gets the value of LockedDownLocalMachineZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowFontDownloads")
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

// SetLockedDownLocalMachineZoneAllowLessPrivilegedSites sets the value of LockedDownLocalMachineZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowLessPrivilegedSites", (value))
}

// GetLockedDownLocalMachineZoneAllowLessPrivilegedSites gets the value of LockedDownLocalMachineZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowLessPrivilegedSites")
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

// SetLockedDownLocalMachineZoneAllowNETFrameworkReliantComponents sets the value of LockedDownLocalMachineZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowNETFrameworkReliantComponents", (value))
}

// GetLockedDownLocalMachineZoneAllowNETFrameworkReliantComponents gets the value of LockedDownLocalMachineZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowNETFrameworkReliantComponents")
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

// SetLockedDownLocalMachineZoneAllowScriptlets sets the value of LockedDownLocalMachineZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowScriptlets", (value))
}

// GetLockedDownLocalMachineZoneAllowScriptlets gets the value of LockedDownLocalMachineZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowScriptlets")
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

// SetLockedDownLocalMachineZoneAllowSmartScreenIE sets the value of LockedDownLocalMachineZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowSmartScreenIE", (value))
}

// GetLockedDownLocalMachineZoneAllowSmartScreenIE gets the value of LockedDownLocalMachineZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowSmartScreenIE")
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

// SetLockedDownLocalMachineZoneAllowUserDataPersistence sets the value of LockedDownLocalMachineZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowUserDataPersistence", (value))
}

// GetLockedDownLocalMachineZoneAllowUserDataPersistence gets the value of LockedDownLocalMachineZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowUserDataPersistence")
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

// SetLockedDownLocalMachineZoneInitializeAndScriptActiveXControls sets the value of LockedDownLocalMachineZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneInitializeAndScriptActiveXControls", (value))
}

// GetLockedDownLocalMachineZoneInitializeAndScriptActiveXControls gets the value of LockedDownLocalMachineZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneInitializeAndScriptActiveXControls")
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

// SetLockedDownLocalMachineZoneJavaPermissions sets the value of LockedDownLocalMachineZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneJavaPermissions", (value))
}

// GetLockedDownLocalMachineZoneJavaPermissions gets the value of LockedDownLocalMachineZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneJavaPermissions")
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

// SetLockedDownLocalMachineZoneNavigateWindowsAndFrames sets the value of LockedDownLocalMachineZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneNavigateWindowsAndFrames", (value))
}

// GetLockedDownLocalMachineZoneNavigateWindowsAndFrames gets the value of LockedDownLocalMachineZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneNavigateWindowsAndFrames")
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

// SetLockedDownRestrictedSitesZoneAllowAccessToDataSources sets the value of LockedDownRestrictedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowAccessToDataSources", (value))
}

// GetLockedDownRestrictedSitesZoneAllowAccessToDataSources gets the value of LockedDownRestrictedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowAccessToDataSources")
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

// SetLockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls sets the value of LockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls", (value))
}

// GetLockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls gets the value of LockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls")
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

// SetLockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads sets the value of LockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads", (value))
}

// GetLockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads gets the value of LockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads")
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

// SetLockedDownRestrictedSitesZoneAllowFontDownloads sets the value of LockedDownRestrictedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowFontDownloads", (value))
}

// GetLockedDownRestrictedSitesZoneAllowFontDownloads gets the value of LockedDownRestrictedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowFontDownloads")
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

// SetLockedDownRestrictedSitesZoneAllowLessPrivilegedSites sets the value of LockedDownRestrictedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowLessPrivilegedSites", (value))
}

// GetLockedDownRestrictedSitesZoneAllowLessPrivilegedSites gets the value of LockedDownRestrictedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowLessPrivilegedSites")
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

// SetLockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents sets the value of LockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents", (value))
}

// GetLockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents gets the value of LockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents")
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

// SetLockedDownRestrictedSitesZoneAllowScriptlets sets the value of LockedDownRestrictedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowScriptlets", (value))
}

// GetLockedDownRestrictedSitesZoneAllowScriptlets gets the value of LockedDownRestrictedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowScriptlets")
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

// SetLockedDownRestrictedSitesZoneAllowSmartScreenIE sets the value of LockedDownRestrictedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowSmartScreenIE", (value))
}

// GetLockedDownRestrictedSitesZoneAllowSmartScreenIE gets the value of LockedDownRestrictedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowSmartScreenIE")
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

// SetLockedDownRestrictedSitesZoneAllowUserDataPersistence sets the value of LockedDownRestrictedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowUserDataPersistence", (value))
}

// GetLockedDownRestrictedSitesZoneAllowUserDataPersistence gets the value of LockedDownRestrictedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowUserDataPersistence")
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

// SetLockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls sets the value of LockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls", (value))
}

// GetLockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls gets the value of LockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls")
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

// SetLockedDownRestrictedSitesZoneJavaPermissions sets the value of LockedDownRestrictedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneJavaPermissions", (value))
}

// GetLockedDownRestrictedSitesZoneJavaPermissions gets the value of LockedDownRestrictedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneJavaPermissions")
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

// SetLockedDownRestrictedSitesZoneNavigateWindowsAndFrames sets the value of LockedDownRestrictedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneNavigateWindowsAndFrames", (value))
}

// GetLockedDownRestrictedSitesZoneNavigateWindowsAndFrames gets the value of LockedDownRestrictedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneNavigateWindowsAndFrames")
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

// SetLockedDownTrustedSitesZoneAllowAccessToDataSources sets the value of LockedDownTrustedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowAccessToDataSources", (value))
}

// GetLockedDownTrustedSitesZoneAllowAccessToDataSources gets the value of LockedDownTrustedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowAccessToDataSources")
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

// SetLockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls sets the value of LockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls", (value))
}

// GetLockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls gets the value of LockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls")
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

// SetLockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads sets the value of LockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads", (value))
}

// GetLockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads gets the value of LockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads")
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

// SetLockedDownTrustedSitesZoneAllowFontDownloads sets the value of LockedDownTrustedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowFontDownloads", (value))
}

// GetLockedDownTrustedSitesZoneAllowFontDownloads gets the value of LockedDownTrustedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowFontDownloads")
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

// SetLockedDownTrustedSitesZoneAllowLessPrivilegedSites sets the value of LockedDownTrustedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowLessPrivilegedSites", (value))
}

// GetLockedDownTrustedSitesZoneAllowLessPrivilegedSites gets the value of LockedDownTrustedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowLessPrivilegedSites")
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

// SetLockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents sets the value of LockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents", (value))
}

// GetLockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents gets the value of LockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents")
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

// SetLockedDownTrustedSitesZoneAllowScriptlets sets the value of LockedDownTrustedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowScriptlets", (value))
}

// GetLockedDownTrustedSitesZoneAllowScriptlets gets the value of LockedDownTrustedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowScriptlets")
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

// SetLockedDownTrustedSitesZoneAllowSmartScreenIE sets the value of LockedDownTrustedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowSmartScreenIE", (value))
}

// GetLockedDownTrustedSitesZoneAllowSmartScreenIE gets the value of LockedDownTrustedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowSmartScreenIE")
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

// SetLockedDownTrustedSitesZoneAllowUserDataPersistence sets the value of LockedDownTrustedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowUserDataPersistence", (value))
}

// GetLockedDownTrustedSitesZoneAllowUserDataPersistence gets the value of LockedDownTrustedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowUserDataPersistence")
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

// SetLockedDownTrustedSitesZoneInitializeAndScriptActiveXControls sets the value of LockedDownTrustedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneInitializeAndScriptActiveXControls", (value))
}

// GetLockedDownTrustedSitesZoneInitializeAndScriptActiveXControls gets the value of LockedDownTrustedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneInitializeAndScriptActiveXControls")
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

// SetLockedDownTrustedSitesZoneJavaPermissions sets the value of LockedDownTrustedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneJavaPermissions", (value))
}

// GetLockedDownTrustedSitesZoneJavaPermissions gets the value of LockedDownTrustedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneJavaPermissions")
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

// SetLockedDownTrustedSitesZoneNavigateWindowsAndFrames sets the value of LockedDownTrustedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneNavigateWindowsAndFrames", (value))
}

// GetLockedDownTrustedSitesZoneNavigateWindowsAndFrames gets the value of LockedDownTrustedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneNavigateWindowsAndFrames")
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

// SetMimeSniffingSafetyFeatureInternetExplorerProcesses sets the value of MimeSniffingSafetyFeatureInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyMimeSniffingSafetyFeatureInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("MimeSniffingSafetyFeatureInternetExplorerProcesses", (value))
}

// GetMimeSniffingSafetyFeatureInternetExplorerProcesses gets the value of MimeSniffingSafetyFeatureInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyMimeSniffingSafetyFeatureInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("MimeSniffingSafetyFeatureInternetExplorerProcesses")
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

// SetMKProtocolSecurityRestrictionInternetExplorerProcesses sets the value of MKProtocolSecurityRestrictionInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyMKProtocolSecurityRestrictionInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("MKProtocolSecurityRestrictionInternetExplorerProcesses", (value))
}

// GetMKProtocolSecurityRestrictionInternetExplorerProcesses gets the value of MKProtocolSecurityRestrictionInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyMKProtocolSecurityRestrictionInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("MKProtocolSecurityRestrictionInternetExplorerProcesses")
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

// SetNewTabDefaultPage sets the value of NewTabDefaultPage for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyNewTabDefaultPage(value string) (err error) {
	return instance.SetProperty("NewTabDefaultPage", (value))
}

// GetNewTabDefaultPage gets the value of NewTabDefaultPage for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyNewTabDefaultPage() (value string, err error) {
	retValue, err := instance.GetProperty("NewTabDefaultPage")
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

// SetNotificationBarInternetExplorerProcesses sets the value of NotificationBarInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyNotificationBarInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("NotificationBarInternetExplorerProcesses", (value))
}

// GetNotificationBarInternetExplorerProcesses gets the value of NotificationBarInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyNotificationBarInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("NotificationBarInternetExplorerProcesses")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyParentID() (value string, err error) {
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

// SetPreventManagingSmartScreenFilter sets the value of PreventManagingSmartScreenFilter for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyPreventManagingSmartScreenFilter(value string) (err error) {
	return instance.SetProperty("PreventManagingSmartScreenFilter", (value))
}

// GetPreventManagingSmartScreenFilter gets the value of PreventManagingSmartScreenFilter for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyPreventManagingSmartScreenFilter() (value string, err error) {
	retValue, err := instance.GetProperty("PreventManagingSmartScreenFilter")
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

// SetPreventPerUserInstallationOfActiveXControls sets the value of PreventPerUserInstallationOfActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyPreventPerUserInstallationOfActiveXControls(value string) (err error) {
	return instance.SetProperty("PreventPerUserInstallationOfActiveXControls", (value))
}

// GetPreventPerUserInstallationOfActiveXControls gets the value of PreventPerUserInstallationOfActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyPreventPerUserInstallationOfActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("PreventPerUserInstallationOfActiveXControls")
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

// SetProtectionFromZoneElevationInternetExplorerProcesses sets the value of ProtectionFromZoneElevationInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyProtectionFromZoneElevationInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("ProtectionFromZoneElevationInternetExplorerProcesses", (value))
}

// GetProtectionFromZoneElevationInternetExplorerProcesses gets the value of ProtectionFromZoneElevationInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyProtectionFromZoneElevationInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("ProtectionFromZoneElevationInternetExplorerProcesses")
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

// SetRemoveRunThisTimeButtonForOutdatedActiveXControls sets the value of RemoveRunThisTimeButtonForOutdatedActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRemoveRunThisTimeButtonForOutdatedActiveXControls(value string) (err error) {
	return instance.SetProperty("RemoveRunThisTimeButtonForOutdatedActiveXControls", (value))
}

// GetRemoveRunThisTimeButtonForOutdatedActiveXControls gets the value of RemoveRunThisTimeButtonForOutdatedActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRemoveRunThisTimeButtonForOutdatedActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RemoveRunThisTimeButtonForOutdatedActiveXControls")
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

// SetRestrictActiveXInstallInternetExplorerProcesses sets the value of RestrictActiveXInstallInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictActiveXInstallInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("RestrictActiveXInstallInternetExplorerProcesses", (value))
}

// GetRestrictActiveXInstallInternetExplorerProcesses gets the value of RestrictActiveXInstallInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictActiveXInstallInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictActiveXInstallInternetExplorerProcesses")
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

// SetRestrictedSitesZoneAllowAccessToDataSources sets the value of RestrictedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowAccessToDataSources", (value))
}

// GetRestrictedSitesZoneAllowAccessToDataSources gets the value of RestrictedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowAccessToDataSources")
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

// SetRestrictedSitesZoneAllowActiveScripting sets the value of RestrictedSitesZoneAllowActiveScripting for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowActiveScripting(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowActiveScripting", (value))
}

// GetRestrictedSitesZoneAllowActiveScripting gets the value of RestrictedSitesZoneAllowActiveScripting for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowActiveScripting() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowActiveScripting")
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

// SetRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls sets the value of RestrictedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowAutomaticPromptingForActiveXControls", (value))
}

// GetRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls gets the value of RestrictedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowAutomaticPromptingForActiveXControls")
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

// SetRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads sets the value of RestrictedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowAutomaticPromptingForFileDownloads", (value))
}

// GetRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads gets the value of RestrictedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowAutomaticPromptingForFileDownloads")
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

// SetRestrictedSitesZoneAllowBinaryAndScriptBehaviors sets the value of RestrictedSitesZoneAllowBinaryAndScriptBehaviors for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowBinaryAndScriptBehaviors(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowBinaryAndScriptBehaviors", (value))
}

// GetRestrictedSitesZoneAllowBinaryAndScriptBehaviors gets the value of RestrictedSitesZoneAllowBinaryAndScriptBehaviors for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowBinaryAndScriptBehaviors() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowBinaryAndScriptBehaviors")
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

// SetRestrictedSitesZoneAllowCopyPasteViaScript sets the value of RestrictedSitesZoneAllowCopyPasteViaScript for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowCopyPasteViaScript(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowCopyPasteViaScript", (value))
}

// GetRestrictedSitesZoneAllowCopyPasteViaScript gets the value of RestrictedSitesZoneAllowCopyPasteViaScript for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowCopyPasteViaScript() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowCopyPasteViaScript")
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

// SetRestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles sets the value of RestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles", (value))
}

// GetRestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles gets the value of RestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles")
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

// SetRestrictedSitesZoneAllowFileDownloads sets the value of RestrictedSitesZoneAllowFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowFileDownloads(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowFileDownloads", (value))
}

// GetRestrictedSitesZoneAllowFileDownloads gets the value of RestrictedSitesZoneAllowFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowFileDownloads")
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

// SetRestrictedSitesZoneAllowFontDownloads sets the value of RestrictedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowFontDownloads", (value))
}

// GetRestrictedSitesZoneAllowFontDownloads gets the value of RestrictedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowFontDownloads")
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

// SetRestrictedSitesZoneAllowLessPrivilegedSites sets the value of RestrictedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowLessPrivilegedSites", (value))
}

// GetRestrictedSitesZoneAllowLessPrivilegedSites gets the value of RestrictedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowLessPrivilegedSites")
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

// SetRestrictedSitesZoneAllowLoadingOfXAMLFiles sets the value of RestrictedSitesZoneAllowLoadingOfXAMLFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowLoadingOfXAMLFiles(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowLoadingOfXAMLFiles", (value))
}

// GetRestrictedSitesZoneAllowLoadingOfXAMLFiles gets the value of RestrictedSitesZoneAllowLoadingOfXAMLFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowLoadingOfXAMLFiles() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowLoadingOfXAMLFiles")
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

// SetRestrictedSitesZoneAllowMETAREFRESH sets the value of RestrictedSitesZoneAllowMETAREFRESH for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowMETAREFRESH(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowMETAREFRESH", (value))
}

// GetRestrictedSitesZoneAllowMETAREFRESH gets the value of RestrictedSitesZoneAllowMETAREFRESH for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowMETAREFRESH() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowMETAREFRESH")
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

// SetRestrictedSitesZoneAllowNETFrameworkReliantComponents sets the value of RestrictedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowNETFrameworkReliantComponents", (value))
}

// GetRestrictedSitesZoneAllowNETFrameworkReliantComponents gets the value of RestrictedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowNETFrameworkReliantComponents")
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

// SetRestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls sets the value of RestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls", (value))
}

// GetRestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls gets the value of RestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls")
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

// SetRestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl sets the value of RestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl", (value))
}

// GetRestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl gets the value of RestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl")
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

// SetRestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls sets the value of RestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls", (value))
}

// GetRestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls gets the value of RestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls")
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

// SetRestrictedSitesZoneAllowScriptInitiatedWindows sets the value of RestrictedSitesZoneAllowScriptInitiatedWindows for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowScriptInitiatedWindows(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowScriptInitiatedWindows", (value))
}

// GetRestrictedSitesZoneAllowScriptInitiatedWindows gets the value of RestrictedSitesZoneAllowScriptInitiatedWindows for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowScriptInitiatedWindows() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowScriptInitiatedWindows")
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

// SetRestrictedSitesZoneAllowScriptlets sets the value of RestrictedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowScriptlets", (value))
}

// GetRestrictedSitesZoneAllowScriptlets gets the value of RestrictedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowScriptlets")
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

// SetRestrictedSitesZoneAllowSmartScreenIE sets the value of RestrictedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowSmartScreenIE", (value))
}

// GetRestrictedSitesZoneAllowSmartScreenIE gets the value of RestrictedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowSmartScreenIE")
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

// SetRestrictedSitesZoneAllowUpdatesToStatusBarViaScript sets the value of RestrictedSitesZoneAllowUpdatesToStatusBarViaScript for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowUpdatesToStatusBarViaScript(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowUpdatesToStatusBarViaScript", (value))
}

// GetRestrictedSitesZoneAllowUpdatesToStatusBarViaScript gets the value of RestrictedSitesZoneAllowUpdatesToStatusBarViaScript for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowUpdatesToStatusBarViaScript() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowUpdatesToStatusBarViaScript")
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

// SetRestrictedSitesZoneAllowUserDataPersistence sets the value of RestrictedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowUserDataPersistence", (value))
}

// GetRestrictedSitesZoneAllowUserDataPersistence gets the value of RestrictedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowUserDataPersistence")
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

// SetRestrictedSitesZoneAllowVBScriptToRunInInternetExplorer sets the value of RestrictedSitesZoneAllowVBScriptToRunInInternetExplorer for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowVBScriptToRunInInternetExplorer(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowVBScriptToRunInInternetExplorer", (value))
}

// GetRestrictedSitesZoneAllowVBScriptToRunInInternetExplorer gets the value of RestrictedSitesZoneAllowVBScriptToRunInInternetExplorer for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowVBScriptToRunInInternetExplorer() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowVBScriptToRunInInternetExplorer")
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

// SetRestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls sets the value of RestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls", (value))
}

// GetRestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls gets the value of RestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls")
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

// SetRestrictedSitesZoneDownloadSignedActiveXControls sets the value of RestrictedSitesZoneDownloadSignedActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneDownloadSignedActiveXControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneDownloadSignedActiveXControls", (value))
}

// GetRestrictedSitesZoneDownloadSignedActiveXControls gets the value of RestrictedSitesZoneDownloadSignedActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneDownloadSignedActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneDownloadSignedActiveXControls")
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

// SetRestrictedSitesZoneDownloadUnsignedActiveXControls sets the value of RestrictedSitesZoneDownloadUnsignedActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneDownloadUnsignedActiveXControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneDownloadUnsignedActiveXControls", (value))
}

// GetRestrictedSitesZoneDownloadUnsignedActiveXControls gets the value of RestrictedSitesZoneDownloadUnsignedActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneDownloadUnsignedActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneDownloadUnsignedActiveXControls")
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

// SetRestrictedSitesZoneEnableCrossSiteScriptingFilter sets the value of RestrictedSitesZoneEnableCrossSiteScriptingFilter for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneEnableCrossSiteScriptingFilter(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneEnableCrossSiteScriptingFilter", (value))
}

// GetRestrictedSitesZoneEnableCrossSiteScriptingFilter gets the value of RestrictedSitesZoneEnableCrossSiteScriptingFilter for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneEnableCrossSiteScriptingFilter() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneEnableCrossSiteScriptingFilter")
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

// SetRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows sets the value of RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows", (value))
}

// GetRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows gets the value of RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows")
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

// SetRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows sets the value of RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows", (value))
}

// GetRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows gets the value of RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows")
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

// SetRestrictedSitesZoneEnableMIMESniffing sets the value of RestrictedSitesZoneEnableMIMESniffing for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneEnableMIMESniffing(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneEnableMIMESniffing", (value))
}

// GetRestrictedSitesZoneEnableMIMESniffing gets the value of RestrictedSitesZoneEnableMIMESniffing for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneEnableMIMESniffing() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneEnableMIMESniffing")
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

// SetRestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer sets the value of RestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer", (value))
}

// GetRestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer gets the value of RestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer")
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

// SetRestrictedSitesZoneInitializeAndScriptActiveXControls sets the value of RestrictedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneInitializeAndScriptActiveXControls", (value))
}

// GetRestrictedSitesZoneInitializeAndScriptActiveXControls gets the value of RestrictedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneInitializeAndScriptActiveXControls")
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

// SetRestrictedSitesZoneJavaPermissions sets the value of RestrictedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneJavaPermissions", (value))
}

// GetRestrictedSitesZoneJavaPermissions gets the value of RestrictedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneJavaPermissions")
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

// SetRestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME sets the value of RestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME", (value))
}

// GetRestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME gets the value of RestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME")
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

// SetRestrictedSitesZoneLogonOptions sets the value of RestrictedSitesZoneLogonOptions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneLogonOptions(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneLogonOptions", (value))
}

// GetRestrictedSitesZoneLogonOptions gets the value of RestrictedSitesZoneLogonOptions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneLogonOptions() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneLogonOptions")
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

// SetRestrictedSitesZoneNavigateWindowsAndFrames sets the value of RestrictedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneNavigateWindowsAndFrames", (value))
}

// GetRestrictedSitesZoneNavigateWindowsAndFrames gets the value of RestrictedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneNavigateWindowsAndFrames")
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

// SetRestrictedSitesZoneRunActiveXControlsAndPlugins sets the value of RestrictedSitesZoneRunActiveXControlsAndPlugins for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneRunActiveXControlsAndPlugins(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneRunActiveXControlsAndPlugins", (value))
}

// GetRestrictedSitesZoneRunActiveXControlsAndPlugins gets the value of RestrictedSitesZoneRunActiveXControlsAndPlugins for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneRunActiveXControlsAndPlugins() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneRunActiveXControlsAndPlugins")
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

// SetRestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode sets the value of RestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode", (value))
}

// GetRestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode gets the value of RestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode")
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

// SetRestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting sets the value of RestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting", (value))
}

// GetRestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting gets the value of RestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting")
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

// SetRestrictedSitesZoneScriptingOfJavaApplets sets the value of RestrictedSitesZoneScriptingOfJavaApplets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneScriptingOfJavaApplets(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneScriptingOfJavaApplets", (value))
}

// GetRestrictedSitesZoneScriptingOfJavaApplets gets the value of RestrictedSitesZoneScriptingOfJavaApplets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneScriptingOfJavaApplets() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneScriptingOfJavaApplets")
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

// SetRestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles sets the value of RestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles", (value))
}

// GetRestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles gets the value of RestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles")
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

// SetRestrictedSitesZoneTurnOnProtectedMode sets the value of RestrictedSitesZoneTurnOnProtectedMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneTurnOnProtectedMode(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneTurnOnProtectedMode", (value))
}

// GetRestrictedSitesZoneTurnOnProtectedMode gets the value of RestrictedSitesZoneTurnOnProtectedMode for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneTurnOnProtectedMode() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneTurnOnProtectedMode")
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

// SetRestrictedSitesZoneUsePopupBlocker sets the value of RestrictedSitesZoneUsePopupBlocker for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictedSitesZoneUsePopupBlocker(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneUsePopupBlocker", (value))
}

// GetRestrictedSitesZoneUsePopupBlocker gets the value of RestrictedSitesZoneUsePopupBlocker for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictedSitesZoneUsePopupBlocker() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneUsePopupBlocker")
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

// SetRestrictFileDownloadInternetExplorerProcesses sets the value of RestrictFileDownloadInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyRestrictFileDownloadInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("RestrictFileDownloadInternetExplorerProcesses", (value))
}

// GetRestrictFileDownloadInternetExplorerProcesses gets the value of RestrictFileDownloadInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyRestrictFileDownloadInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictFileDownloadInternetExplorerProcesses")
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

// SetScriptedWindowSecurityRestrictionsInternetExplorerProcesses sets the value of ScriptedWindowSecurityRestrictionsInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyScriptedWindowSecurityRestrictionsInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("ScriptedWindowSecurityRestrictionsInternetExplorerProcesses", (value))
}

// GetScriptedWindowSecurityRestrictionsInternetExplorerProcesses gets the value of ScriptedWindowSecurityRestrictionsInternetExplorerProcesses for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyScriptedWindowSecurityRestrictionsInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("ScriptedWindowSecurityRestrictionsInternetExplorerProcesses")
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

// SetSearchProviderList sets the value of SearchProviderList for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertySearchProviderList(value string) (err error) {
	return instance.SetProperty("SearchProviderList", (value))
}

// GetSearchProviderList gets the value of SearchProviderList for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertySearchProviderList() (value string, err error) {
	retValue, err := instance.GetProperty("SearchProviderList")
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

// SetSendSitesNotInEnterpriseSiteListToEdge sets the value of SendSitesNotInEnterpriseSiteListToEdge for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertySendSitesNotInEnterpriseSiteListToEdge(value string) (err error) {
	return instance.SetProperty("SendSitesNotInEnterpriseSiteListToEdge", (value))
}

// GetSendSitesNotInEnterpriseSiteListToEdge gets the value of SendSitesNotInEnterpriseSiteListToEdge for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertySendSitesNotInEnterpriseSiteListToEdge() (value string, err error) {
	retValue, err := instance.GetProperty("SendSitesNotInEnterpriseSiteListToEdge")
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

// SetSpecifyUseOfActiveXInstallerService sets the value of SpecifyUseOfActiveXInstallerService for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertySpecifyUseOfActiveXInstallerService(value string) (err error) {
	return instance.SetProperty("SpecifyUseOfActiveXInstallerService", (value))
}

// GetSpecifyUseOfActiveXInstallerService gets the value of SpecifyUseOfActiveXInstallerService for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertySpecifyUseOfActiveXInstallerService() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyUseOfActiveXInstallerService")
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

// SetTrustedSitesZoneAllowAccessToDataSources sets the value of TrustedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowAccessToDataSources", (value))
}

// GetTrustedSitesZoneAllowAccessToDataSources gets the value of TrustedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowAccessToDataSources")
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

// SetTrustedSitesZoneAllowAutomaticPromptingForActiveXControls sets the value of TrustedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowAutomaticPromptingForActiveXControls", (value))
}

// GetTrustedSitesZoneAllowAutomaticPromptingForActiveXControls gets the value of TrustedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowAutomaticPromptingForActiveXControls")
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

// SetTrustedSitesZoneAllowAutomaticPromptingForFileDownloads sets the value of TrustedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowAutomaticPromptingForFileDownloads", (value))
}

// GetTrustedSitesZoneAllowAutomaticPromptingForFileDownloads gets the value of TrustedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowAutomaticPromptingForFileDownloads")
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

// SetTrustedSitesZoneAllowFontDownloads sets the value of TrustedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowFontDownloads", (value))
}

// GetTrustedSitesZoneAllowFontDownloads gets the value of TrustedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowFontDownloads")
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

// SetTrustedSitesZoneAllowLessPrivilegedSites sets the value of TrustedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowLessPrivilegedSites", (value))
}

// GetTrustedSitesZoneAllowLessPrivilegedSites gets the value of TrustedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowLessPrivilegedSites")
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

// SetTrustedSitesZoneAllowNETFrameworkReliantComponents sets the value of TrustedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowNETFrameworkReliantComponents", (value))
}

// GetTrustedSitesZoneAllowNETFrameworkReliantComponents gets the value of TrustedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowNETFrameworkReliantComponents")
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

// SetTrustedSitesZoneAllowScriptlets sets the value of TrustedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowScriptlets", (value))
}

// GetTrustedSitesZoneAllowScriptlets gets the value of TrustedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowScriptlets")
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

// SetTrustedSitesZoneAllowSmartScreenIE sets the value of TrustedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowSmartScreenIE", (value))
}

// GetTrustedSitesZoneAllowSmartScreenIE gets the value of TrustedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowSmartScreenIE")
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

// SetTrustedSitesZoneAllowUserDataPersistence sets the value of TrustedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowUserDataPersistence", (value))
}

// GetTrustedSitesZoneAllowUserDataPersistence gets the value of TrustedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowUserDataPersistence")
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

// SetTrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls sets the value of TrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyTrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls", (value))
}

// GetTrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls gets the value of TrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyTrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls")
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

// SetTrustedSitesZoneInitializeAndScriptActiveXControls sets the value of TrustedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyTrustedSitesZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneInitializeAndScriptActiveXControls", (value))
}

// GetTrustedSitesZoneInitializeAndScriptActiveXControls gets the value of TrustedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyTrustedSitesZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneInitializeAndScriptActiveXControls")
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

// SetTrustedSitesZoneJavaPermissions sets the value of TrustedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyTrustedSitesZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneJavaPermissions", (value))
}

// GetTrustedSitesZoneJavaPermissions gets the value of TrustedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyTrustedSitesZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneJavaPermissions")
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

// SetTrustedSitesZoneNavigateWindowsAndFrames sets the value of TrustedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) SetPropertyTrustedSitesZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneNavigateWindowsAndFrames", (value))
}

// GetTrustedSitesZoneNavigateWindowsAndFrames gets the value of TrustedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_User_Config01_InternetExplorer02) GetPropertyTrustedSitesZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneNavigateWindowsAndFrames")
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
