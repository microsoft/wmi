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

// MDM_Policy_Result01_InternetExplorer02 struct
type MDM_Policy_Result01_InternetExplorer02 struct {
	*cim.WmiInstance

	//
	AddSearchProvider string

	//
	AllowActiveXFiltering string

	//
	AllowAddOnList string

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
	AllowFallbackToSSL3 string

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
	ConsistentMimeHandlingInternetExplorerProcesses string

	//
	DisableAdobeFlash string

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
	DisableIgnoringCertificateErrors string

	//
	DisableInPrivateBrowsing string

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
	DisableUpdateCheck string

	//
	DisableWebAddressAutoComplete string

	//
	DoNotAllowActiveXControlsInProtectedMode string

	//
	DoNotAllowUsersToAddSites string

	//
	DoNotAllowUsersToChangePolicies string

	//
	DoNotBlockOutdatedActiveXControls string

	//
	DoNotBlockOutdatedActiveXControlsOnSpecificDomains string

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
	SecurityZonesUseOnlyMachineSettings string

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

func NewMDM_Policy_Result01_InternetExplorer02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_InternetExplorer02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_InternetExplorer02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_InternetExplorer02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_InternetExplorer02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_InternetExplorer02{
		WmiInstance: tmp,
	}
	return
}

// SetAddSearchProvider sets the value of AddSearchProvider for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAddSearchProvider(value string) (err error) {
	return instance.SetProperty("AddSearchProvider", value)
}

// GetAddSearchProvider gets the value of AddSearchProvider for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAddSearchProvider() (value string, err error) {
	retValue, err := instance.GetProperty("AddSearchProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowActiveXFiltering sets the value of AllowActiveXFiltering for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowActiveXFiltering(value string) (err error) {
	return instance.SetProperty("AllowActiveXFiltering", value)
}

// GetAllowActiveXFiltering gets the value of AllowActiveXFiltering for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowActiveXFiltering() (value string, err error) {
	retValue, err := instance.GetProperty("AllowActiveXFiltering")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowAddOnList sets the value of AllowAddOnList for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowAddOnList(value string) (err error) {
	return instance.SetProperty("AllowAddOnList", value)
}

// GetAllowAddOnList gets the value of AllowAddOnList for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowAddOnList() (value string, err error) {
	retValue, err := instance.GetProperty("AllowAddOnList")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowCertificateAddressMismatchWarning sets the value of AllowCertificateAddressMismatchWarning for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowCertificateAddressMismatchWarning(value string) (err error) {
	return instance.SetProperty("AllowCertificateAddressMismatchWarning", value)
}

// GetAllowCertificateAddressMismatchWarning gets the value of AllowCertificateAddressMismatchWarning for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowCertificateAddressMismatchWarning() (value string, err error) {
	retValue, err := instance.GetProperty("AllowCertificateAddressMismatchWarning")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowDeletingBrowsingHistoryOnExit sets the value of AllowDeletingBrowsingHistoryOnExit for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowDeletingBrowsingHistoryOnExit(value string) (err error) {
	return instance.SetProperty("AllowDeletingBrowsingHistoryOnExit", value)
}

// GetAllowDeletingBrowsingHistoryOnExit gets the value of AllowDeletingBrowsingHistoryOnExit for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowDeletingBrowsingHistoryOnExit() (value string, err error) {
	retValue, err := instance.GetProperty("AllowDeletingBrowsingHistoryOnExit")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowEnhancedProtectedMode sets the value of AllowEnhancedProtectedMode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowEnhancedProtectedMode(value string) (err error) {
	return instance.SetProperty("AllowEnhancedProtectedMode", value)
}

// GetAllowEnhancedProtectedMode gets the value of AllowEnhancedProtectedMode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowEnhancedProtectedMode() (value string, err error) {
	retValue, err := instance.GetProperty("AllowEnhancedProtectedMode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowEnhancedSuggestionsInAddressBar sets the value of AllowEnhancedSuggestionsInAddressBar for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowEnhancedSuggestionsInAddressBar(value string) (err error) {
	return instance.SetProperty("AllowEnhancedSuggestionsInAddressBar", value)
}

// GetAllowEnhancedSuggestionsInAddressBar gets the value of AllowEnhancedSuggestionsInAddressBar for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowEnhancedSuggestionsInAddressBar() (value string, err error) {
	retValue, err := instance.GetProperty("AllowEnhancedSuggestionsInAddressBar")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowEnterpriseModeFromToolsMenu sets the value of AllowEnterpriseModeFromToolsMenu for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowEnterpriseModeFromToolsMenu(value string) (err error) {
	return instance.SetProperty("AllowEnterpriseModeFromToolsMenu", value)
}

// GetAllowEnterpriseModeFromToolsMenu gets the value of AllowEnterpriseModeFromToolsMenu for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowEnterpriseModeFromToolsMenu() (value string, err error) {
	retValue, err := instance.GetProperty("AllowEnterpriseModeFromToolsMenu")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowEnterpriseModeSiteList sets the value of AllowEnterpriseModeSiteList for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowEnterpriseModeSiteList(value string) (err error) {
	return instance.SetProperty("AllowEnterpriseModeSiteList", value)
}

// GetAllowEnterpriseModeSiteList gets the value of AllowEnterpriseModeSiteList for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowEnterpriseModeSiteList() (value string, err error) {
	retValue, err := instance.GetProperty("AllowEnterpriseModeSiteList")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowFallbackToSSL3 sets the value of AllowFallbackToSSL3 for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowFallbackToSSL3(value string) (err error) {
	return instance.SetProperty("AllowFallbackToSSL3", value)
}

// GetAllowFallbackToSSL3 gets the value of AllowFallbackToSSL3 for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowFallbackToSSL3() (value string, err error) {
	retValue, err := instance.GetProperty("AllowFallbackToSSL3")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowInternetExplorer7PolicyList sets the value of AllowInternetExplorer7PolicyList for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowInternetExplorer7PolicyList(value string) (err error) {
	return instance.SetProperty("AllowInternetExplorer7PolicyList", value)
}

// GetAllowInternetExplorer7PolicyList gets the value of AllowInternetExplorer7PolicyList for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowInternetExplorer7PolicyList() (value string, err error) {
	retValue, err := instance.GetProperty("AllowInternetExplorer7PolicyList")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowInternetExplorerStandardsMode sets the value of AllowInternetExplorerStandardsMode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowInternetExplorerStandardsMode(value string) (err error) {
	return instance.SetProperty("AllowInternetExplorerStandardsMode", value)
}

// GetAllowInternetExplorerStandardsMode gets the value of AllowInternetExplorerStandardsMode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowInternetExplorerStandardsMode() (value string, err error) {
	retValue, err := instance.GetProperty("AllowInternetExplorerStandardsMode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowInternetZoneTemplate sets the value of AllowInternetZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowInternetZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowInternetZoneTemplate", value)
}

// GetAllowInternetZoneTemplate gets the value of AllowInternetZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowInternetZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowInternetZoneTemplate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowIntranetZoneTemplate sets the value of AllowIntranetZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowIntranetZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowIntranetZoneTemplate", value)
}

// GetAllowIntranetZoneTemplate gets the value of AllowIntranetZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowIntranetZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowIntranetZoneTemplate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowLocalMachineZoneTemplate sets the value of AllowLocalMachineZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowLocalMachineZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowLocalMachineZoneTemplate", value)
}

// GetAllowLocalMachineZoneTemplate gets the value of AllowLocalMachineZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowLocalMachineZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowLocalMachineZoneTemplate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowLockedDownInternetZoneTemplate sets the value of AllowLockedDownInternetZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowLockedDownInternetZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowLockedDownInternetZoneTemplate", value)
}

// GetAllowLockedDownInternetZoneTemplate gets the value of AllowLockedDownInternetZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowLockedDownInternetZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowLockedDownInternetZoneTemplate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowLockedDownIntranetZoneTemplate sets the value of AllowLockedDownIntranetZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowLockedDownIntranetZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowLockedDownIntranetZoneTemplate", value)
}

// GetAllowLockedDownIntranetZoneTemplate gets the value of AllowLockedDownIntranetZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowLockedDownIntranetZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowLockedDownIntranetZoneTemplate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowLockedDownLocalMachineZoneTemplate sets the value of AllowLockedDownLocalMachineZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowLockedDownLocalMachineZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowLockedDownLocalMachineZoneTemplate", value)
}

// GetAllowLockedDownLocalMachineZoneTemplate gets the value of AllowLockedDownLocalMachineZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowLockedDownLocalMachineZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowLockedDownLocalMachineZoneTemplate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowLockedDownRestrictedSitesZoneTemplate sets the value of AllowLockedDownRestrictedSitesZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowLockedDownRestrictedSitesZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowLockedDownRestrictedSitesZoneTemplate", value)
}

// GetAllowLockedDownRestrictedSitesZoneTemplate gets the value of AllowLockedDownRestrictedSitesZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowLockedDownRestrictedSitesZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowLockedDownRestrictedSitesZoneTemplate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowOneWordEntry sets the value of AllowOneWordEntry for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowOneWordEntry(value string) (err error) {
	return instance.SetProperty("AllowOneWordEntry", value)
}

// GetAllowOneWordEntry gets the value of AllowOneWordEntry for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowOneWordEntry() (value string, err error) {
	retValue, err := instance.GetProperty("AllowOneWordEntry")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowSiteToZoneAssignmentList sets the value of AllowSiteToZoneAssignmentList for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowSiteToZoneAssignmentList(value string) (err error) {
	return instance.SetProperty("AllowSiteToZoneAssignmentList", value)
}

// GetAllowSiteToZoneAssignmentList gets the value of AllowSiteToZoneAssignmentList for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowSiteToZoneAssignmentList() (value string, err error) {
	retValue, err := instance.GetProperty("AllowSiteToZoneAssignmentList")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowsLockedDownTrustedSitesZoneTemplate sets the value of AllowsLockedDownTrustedSitesZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowsLockedDownTrustedSitesZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowsLockedDownTrustedSitesZoneTemplate", value)
}

// GetAllowsLockedDownTrustedSitesZoneTemplate gets the value of AllowsLockedDownTrustedSitesZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowsLockedDownTrustedSitesZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowsLockedDownTrustedSitesZoneTemplate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowSoftwareWhenSignatureIsInvalid sets the value of AllowSoftwareWhenSignatureIsInvalid for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowSoftwareWhenSignatureIsInvalid(value string) (err error) {
	return instance.SetProperty("AllowSoftwareWhenSignatureIsInvalid", value)
}

// GetAllowSoftwareWhenSignatureIsInvalid gets the value of AllowSoftwareWhenSignatureIsInvalid for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowSoftwareWhenSignatureIsInvalid() (value string, err error) {
	retValue, err := instance.GetProperty("AllowSoftwareWhenSignatureIsInvalid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowsRestrictedSitesZoneTemplate sets the value of AllowsRestrictedSitesZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowsRestrictedSitesZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowsRestrictedSitesZoneTemplate", value)
}

// GetAllowsRestrictedSitesZoneTemplate gets the value of AllowsRestrictedSitesZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowsRestrictedSitesZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowsRestrictedSitesZoneTemplate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowSuggestedSites sets the value of AllowSuggestedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowSuggestedSites(value string) (err error) {
	return instance.SetProperty("AllowSuggestedSites", value)
}

// GetAllowSuggestedSites gets the value of AllowSuggestedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowSuggestedSites() (value string, err error) {
	retValue, err := instance.GetProperty("AllowSuggestedSites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowTrustedSitesZoneTemplate sets the value of AllowTrustedSitesZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyAllowTrustedSitesZoneTemplate(value string) (err error) {
	return instance.SetProperty("AllowTrustedSitesZoneTemplate", value)
}

// GetAllowTrustedSitesZoneTemplate gets the value of AllowTrustedSitesZoneTemplate for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyAllowTrustedSitesZoneTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("AllowTrustedSitesZoneTemplate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCheckServerCertificateRevocation sets the value of CheckServerCertificateRevocation for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyCheckServerCertificateRevocation(value string) (err error) {
	return instance.SetProperty("CheckServerCertificateRevocation", value)
}

// GetCheckServerCertificateRevocation gets the value of CheckServerCertificateRevocation for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyCheckServerCertificateRevocation() (value string, err error) {
	retValue, err := instance.GetProperty("CheckServerCertificateRevocation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCheckSignaturesOnDownloadedPrograms sets the value of CheckSignaturesOnDownloadedPrograms for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyCheckSignaturesOnDownloadedPrograms(value string) (err error) {
	return instance.SetProperty("CheckSignaturesOnDownloadedPrograms", value)
}

// GetCheckSignaturesOnDownloadedPrograms gets the value of CheckSignaturesOnDownloadedPrograms for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyCheckSignaturesOnDownloadedPrograms() (value string, err error) {
	retValue, err := instance.GetProperty("CheckSignaturesOnDownloadedPrograms")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConsistentMimeHandlingInternetExplorerProcesses sets the value of ConsistentMimeHandlingInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyConsistentMimeHandlingInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("ConsistentMimeHandlingInternetExplorerProcesses", value)
}

// GetConsistentMimeHandlingInternetExplorerProcesses gets the value of ConsistentMimeHandlingInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyConsistentMimeHandlingInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("ConsistentMimeHandlingInternetExplorerProcesses")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableAdobeFlash sets the value of DisableAdobeFlash for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableAdobeFlash(value string) (err error) {
	return instance.SetProperty("DisableAdobeFlash", value)
}

// GetDisableAdobeFlash gets the value of DisableAdobeFlash for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableAdobeFlash() (value string, err error) {
	retValue, err := instance.GetProperty("DisableAdobeFlash")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableBypassOfSmartScreenWarnings sets the value of DisableBypassOfSmartScreenWarnings for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableBypassOfSmartScreenWarnings(value string) (err error) {
	return instance.SetProperty("DisableBypassOfSmartScreenWarnings", value)
}

// GetDisableBypassOfSmartScreenWarnings gets the value of DisableBypassOfSmartScreenWarnings for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableBypassOfSmartScreenWarnings() (value string, err error) {
	retValue, err := instance.GetProperty("DisableBypassOfSmartScreenWarnings")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableBypassOfSmartScreenWarningsAboutUncommonFiles sets the value of DisableBypassOfSmartScreenWarningsAboutUncommonFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableBypassOfSmartScreenWarningsAboutUncommonFiles(value string) (err error) {
	return instance.SetProperty("DisableBypassOfSmartScreenWarningsAboutUncommonFiles", value)
}

// GetDisableBypassOfSmartScreenWarningsAboutUncommonFiles gets the value of DisableBypassOfSmartScreenWarningsAboutUncommonFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableBypassOfSmartScreenWarningsAboutUncommonFiles() (value string, err error) {
	retValue, err := instance.GetProperty("DisableBypassOfSmartScreenWarningsAboutUncommonFiles")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableCompatView sets the value of DisableCompatView for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableCompatView(value string) (err error) {
	return instance.SetProperty("DisableCompatView", value)
}

// GetDisableCompatView gets the value of DisableCompatView for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableCompatView() (value string, err error) {
	retValue, err := instance.GetProperty("DisableCompatView")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableConfiguringHistory sets the value of DisableConfiguringHistory for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableConfiguringHistory(value string) (err error) {
	return instance.SetProperty("DisableConfiguringHistory", value)
}

// GetDisableConfiguringHistory gets the value of DisableConfiguringHistory for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableConfiguringHistory() (value string, err error) {
	retValue, err := instance.GetProperty("DisableConfiguringHistory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableCrashDetection sets the value of DisableCrashDetection for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableCrashDetection(value string) (err error) {
	return instance.SetProperty("DisableCrashDetection", value)
}

// GetDisableCrashDetection gets the value of DisableCrashDetection for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableCrashDetection() (value string, err error) {
	retValue, err := instance.GetProperty("DisableCrashDetection")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableCustomerExperienceImprovementProgramParticipation sets the value of DisableCustomerExperienceImprovementProgramParticipation for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableCustomerExperienceImprovementProgramParticipation(value string) (err error) {
	return instance.SetProperty("DisableCustomerExperienceImprovementProgramParticipation", value)
}

// GetDisableCustomerExperienceImprovementProgramParticipation gets the value of DisableCustomerExperienceImprovementProgramParticipation for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableCustomerExperienceImprovementProgramParticipation() (value string, err error) {
	retValue, err := instance.GetProperty("DisableCustomerExperienceImprovementProgramParticipation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableDeletingUserVisitedWebsites sets the value of DisableDeletingUserVisitedWebsites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableDeletingUserVisitedWebsites(value string) (err error) {
	return instance.SetProperty("DisableDeletingUserVisitedWebsites", value)
}

// GetDisableDeletingUserVisitedWebsites gets the value of DisableDeletingUserVisitedWebsites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableDeletingUserVisitedWebsites() (value string, err error) {
	retValue, err := instance.GetProperty("DisableDeletingUserVisitedWebsites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableEnclosureDownloading sets the value of DisableEnclosureDownloading for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableEnclosureDownloading(value string) (err error) {
	return instance.SetProperty("DisableEnclosureDownloading", value)
}

// GetDisableEnclosureDownloading gets the value of DisableEnclosureDownloading for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableEnclosureDownloading() (value string, err error) {
	retValue, err := instance.GetProperty("DisableEnclosureDownloading")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableEncryptionSupport sets the value of DisableEncryptionSupport for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableEncryptionSupport(value string) (err error) {
	return instance.SetProperty("DisableEncryptionSupport", value)
}

// GetDisableEncryptionSupport gets the value of DisableEncryptionSupport for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableEncryptionSupport() (value string, err error) {
	retValue, err := instance.GetProperty("DisableEncryptionSupport")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableFeedsBackgroundSync sets the value of DisableFeedsBackgroundSync for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableFeedsBackgroundSync(value string) (err error) {
	return instance.SetProperty("DisableFeedsBackgroundSync", value)
}

// GetDisableFeedsBackgroundSync gets the value of DisableFeedsBackgroundSync for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableFeedsBackgroundSync() (value string, err error) {
	retValue, err := instance.GetProperty("DisableFeedsBackgroundSync")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableFirstRunWizard sets the value of DisableFirstRunWizard for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableFirstRunWizard(value string) (err error) {
	return instance.SetProperty("DisableFirstRunWizard", value)
}

// GetDisableFirstRunWizard gets the value of DisableFirstRunWizard for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableFirstRunWizard() (value string, err error) {
	retValue, err := instance.GetProperty("DisableFirstRunWizard")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableFlipAheadFeature sets the value of DisableFlipAheadFeature for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableFlipAheadFeature(value string) (err error) {
	return instance.SetProperty("DisableFlipAheadFeature", value)
}

// GetDisableFlipAheadFeature gets the value of DisableFlipAheadFeature for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableFlipAheadFeature() (value string, err error) {
	retValue, err := instance.GetProperty("DisableFlipAheadFeature")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableGeolocation sets the value of DisableGeolocation for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableGeolocation(value string) (err error) {
	return instance.SetProperty("DisableGeolocation", value)
}

// GetDisableGeolocation gets the value of DisableGeolocation for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableGeolocation() (value string, err error) {
	retValue, err := instance.GetProperty("DisableGeolocation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableIgnoringCertificateErrors sets the value of DisableIgnoringCertificateErrors for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableIgnoringCertificateErrors(value string) (err error) {
	return instance.SetProperty("DisableIgnoringCertificateErrors", value)
}

// GetDisableIgnoringCertificateErrors gets the value of DisableIgnoringCertificateErrors for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableIgnoringCertificateErrors() (value string, err error) {
	retValue, err := instance.GetProperty("DisableIgnoringCertificateErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableInPrivateBrowsing sets the value of DisableInPrivateBrowsing for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableInPrivateBrowsing(value string) (err error) {
	return instance.SetProperty("DisableInPrivateBrowsing", value)
}

// GetDisableInPrivateBrowsing gets the value of DisableInPrivateBrowsing for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableInPrivateBrowsing() (value string, err error) {
	retValue, err := instance.GetProperty("DisableInPrivateBrowsing")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableProcessesInEnhancedProtectedMode sets the value of DisableProcessesInEnhancedProtectedMode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableProcessesInEnhancedProtectedMode(value string) (err error) {
	return instance.SetProperty("DisableProcessesInEnhancedProtectedMode", value)
}

// GetDisableProcessesInEnhancedProtectedMode gets the value of DisableProcessesInEnhancedProtectedMode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableProcessesInEnhancedProtectedMode() (value string, err error) {
	retValue, err := instance.GetProperty("DisableProcessesInEnhancedProtectedMode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableProxyChange sets the value of DisableProxyChange for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableProxyChange(value string) (err error) {
	return instance.SetProperty("DisableProxyChange", value)
}

// GetDisableProxyChange gets the value of DisableProxyChange for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableProxyChange() (value string, err error) {
	retValue, err := instance.GetProperty("DisableProxyChange")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableSearchProviderChange sets the value of DisableSearchProviderChange for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableSearchProviderChange(value string) (err error) {
	return instance.SetProperty("DisableSearchProviderChange", value)
}

// GetDisableSearchProviderChange gets the value of DisableSearchProviderChange for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableSearchProviderChange() (value string, err error) {
	retValue, err := instance.GetProperty("DisableSearchProviderChange")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableSecondaryHomePageChange sets the value of DisableSecondaryHomePageChange for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableSecondaryHomePageChange(value string) (err error) {
	return instance.SetProperty("DisableSecondaryHomePageChange", value)
}

// GetDisableSecondaryHomePageChange gets the value of DisableSecondaryHomePageChange for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableSecondaryHomePageChange() (value string, err error) {
	retValue, err := instance.GetProperty("DisableSecondaryHomePageChange")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableSecuritySettingsCheck sets the value of DisableSecuritySettingsCheck for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableSecuritySettingsCheck(value string) (err error) {
	return instance.SetProperty("DisableSecuritySettingsCheck", value)
}

// GetDisableSecuritySettingsCheck gets the value of DisableSecuritySettingsCheck for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableSecuritySettingsCheck() (value string, err error) {
	retValue, err := instance.GetProperty("DisableSecuritySettingsCheck")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableUpdateCheck sets the value of DisableUpdateCheck for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableUpdateCheck(value string) (err error) {
	return instance.SetProperty("DisableUpdateCheck", value)
}

// GetDisableUpdateCheck gets the value of DisableUpdateCheck for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableUpdateCheck() (value string, err error) {
	retValue, err := instance.GetProperty("DisableUpdateCheck")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableWebAddressAutoComplete sets the value of DisableWebAddressAutoComplete for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDisableWebAddressAutoComplete(value string) (err error) {
	return instance.SetProperty("DisableWebAddressAutoComplete", value)
}

// GetDisableWebAddressAutoComplete gets the value of DisableWebAddressAutoComplete for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDisableWebAddressAutoComplete() (value string, err error) {
	retValue, err := instance.GetProperty("DisableWebAddressAutoComplete")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDoNotAllowActiveXControlsInProtectedMode sets the value of DoNotAllowActiveXControlsInProtectedMode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDoNotAllowActiveXControlsInProtectedMode(value string) (err error) {
	return instance.SetProperty("DoNotAllowActiveXControlsInProtectedMode", value)
}

// GetDoNotAllowActiveXControlsInProtectedMode gets the value of DoNotAllowActiveXControlsInProtectedMode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDoNotAllowActiveXControlsInProtectedMode() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotAllowActiveXControlsInProtectedMode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDoNotAllowUsersToAddSites sets the value of DoNotAllowUsersToAddSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDoNotAllowUsersToAddSites(value string) (err error) {
	return instance.SetProperty("DoNotAllowUsersToAddSites", value)
}

// GetDoNotAllowUsersToAddSites gets the value of DoNotAllowUsersToAddSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDoNotAllowUsersToAddSites() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotAllowUsersToAddSites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDoNotAllowUsersToChangePolicies sets the value of DoNotAllowUsersToChangePolicies for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDoNotAllowUsersToChangePolicies(value string) (err error) {
	return instance.SetProperty("DoNotAllowUsersToChangePolicies", value)
}

// GetDoNotAllowUsersToChangePolicies gets the value of DoNotAllowUsersToChangePolicies for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDoNotAllowUsersToChangePolicies() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotAllowUsersToChangePolicies")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDoNotBlockOutdatedActiveXControls sets the value of DoNotBlockOutdatedActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDoNotBlockOutdatedActiveXControls(value string) (err error) {
	return instance.SetProperty("DoNotBlockOutdatedActiveXControls", value)
}

// GetDoNotBlockOutdatedActiveXControls gets the value of DoNotBlockOutdatedActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDoNotBlockOutdatedActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotBlockOutdatedActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDoNotBlockOutdatedActiveXControlsOnSpecificDomains sets the value of DoNotBlockOutdatedActiveXControlsOnSpecificDomains for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyDoNotBlockOutdatedActiveXControlsOnSpecificDomains(value string) (err error) {
	return instance.SetProperty("DoNotBlockOutdatedActiveXControlsOnSpecificDomains", value)
}

// GetDoNotBlockOutdatedActiveXControlsOnSpecificDomains gets the value of DoNotBlockOutdatedActiveXControlsOnSpecificDomains for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyDoNotBlockOutdatedActiveXControlsOnSpecificDomains() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotBlockOutdatedActiveXControlsOnSpecificDomains")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIncludeAllLocalSites sets the value of IncludeAllLocalSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIncludeAllLocalSites(value string) (err error) {
	return instance.SetProperty("IncludeAllLocalSites", value)
}

// GetIncludeAllLocalSites gets the value of IncludeAllLocalSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIncludeAllLocalSites() (value string, err error) {
	retValue, err := instance.GetProperty("IncludeAllLocalSites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIncludeAllNetworkPaths sets the value of IncludeAllNetworkPaths for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIncludeAllNetworkPaths(value string) (err error) {
	return instance.SetProperty("IncludeAllNetworkPaths", value)
}

// GetIncludeAllNetworkPaths gets the value of IncludeAllNetworkPaths for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIncludeAllNetworkPaths() (value string, err error) {
	retValue, err := instance.GetProperty("IncludeAllNetworkPaths")
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
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInstanceID() (value string, err error) {
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

// SetInternetZoneAllowAccessToDataSources sets the value of InternetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowAccessToDataSources", value)
}

// GetInternetZoneAllowAccessToDataSources gets the value of InternetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowAccessToDataSources")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowAutomaticPromptingForActiveXControls sets the value of InternetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowAutomaticPromptingForActiveXControls", value)
}

// GetInternetZoneAllowAutomaticPromptingForActiveXControls gets the value of InternetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowAutomaticPromptingForActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowAutomaticPromptingForFileDownloads sets the value of InternetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowAutomaticPromptingForFileDownloads", value)
}

// GetInternetZoneAllowAutomaticPromptingForFileDownloads gets the value of InternetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowAutomaticPromptingForFileDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowCopyPasteViaScript sets the value of InternetZoneAllowCopyPasteViaScript for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowCopyPasteViaScript(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowCopyPasteViaScript", value)
}

// GetInternetZoneAllowCopyPasteViaScript gets the value of InternetZoneAllowCopyPasteViaScript for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowCopyPasteViaScript() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowCopyPasteViaScript")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowDragAndDropCopyAndPasteFiles sets the value of InternetZoneAllowDragAndDropCopyAndPasteFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowDragAndDropCopyAndPasteFiles(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowDragAndDropCopyAndPasteFiles", value)
}

// GetInternetZoneAllowDragAndDropCopyAndPasteFiles gets the value of InternetZoneAllowDragAndDropCopyAndPasteFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowDragAndDropCopyAndPasteFiles() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowDragAndDropCopyAndPasteFiles")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowFontDownloads sets the value of InternetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowFontDownloads", value)
}

// GetInternetZoneAllowFontDownloads gets the value of InternetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowFontDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowLessPrivilegedSites sets the value of InternetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowLessPrivilegedSites", value)
}

// GetInternetZoneAllowLessPrivilegedSites gets the value of InternetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowLessPrivilegedSites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowLoadingOfXAMLFiles sets the value of InternetZoneAllowLoadingOfXAMLFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowLoadingOfXAMLFiles(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowLoadingOfXAMLFiles", value)
}

// GetInternetZoneAllowLoadingOfXAMLFiles gets the value of InternetZoneAllowLoadingOfXAMLFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowLoadingOfXAMLFiles() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowLoadingOfXAMLFiles")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowNETFrameworkReliantComponents sets the value of InternetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowNETFrameworkReliantComponents", value)
}

// GetInternetZoneAllowNETFrameworkReliantComponents gets the value of InternetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowNETFrameworkReliantComponents")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowOnlyApprovedDomainsToUseActiveXControls sets the value of InternetZoneAllowOnlyApprovedDomainsToUseActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowOnlyApprovedDomainsToUseActiveXControls(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowOnlyApprovedDomainsToUseActiveXControls", value)
}

// GetInternetZoneAllowOnlyApprovedDomainsToUseActiveXControls gets the value of InternetZoneAllowOnlyApprovedDomainsToUseActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowOnlyApprovedDomainsToUseActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowOnlyApprovedDomainsToUseActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl sets the value of InternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl", value)
}

// GetInternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl gets the value of InternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowScriptingOfInternetExplorerWebBrowserControls sets the value of InternetZoneAllowScriptingOfInternetExplorerWebBrowserControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowScriptingOfInternetExplorerWebBrowserControls(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowScriptingOfInternetExplorerWebBrowserControls", value)
}

// GetInternetZoneAllowScriptingOfInternetExplorerWebBrowserControls gets the value of InternetZoneAllowScriptingOfInternetExplorerWebBrowserControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowScriptingOfInternetExplorerWebBrowserControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowScriptingOfInternetExplorerWebBrowserControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowScriptInitiatedWindows sets the value of InternetZoneAllowScriptInitiatedWindows for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowScriptInitiatedWindows(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowScriptInitiatedWindows", value)
}

// GetInternetZoneAllowScriptInitiatedWindows gets the value of InternetZoneAllowScriptInitiatedWindows for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowScriptInitiatedWindows() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowScriptInitiatedWindows")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowScriptlets sets the value of InternetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowScriptlets", value)
}

// GetInternetZoneAllowScriptlets gets the value of InternetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowScriptlets")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowSmartScreenIE sets the value of InternetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowSmartScreenIE", value)
}

// GetInternetZoneAllowSmartScreenIE gets the value of InternetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowSmartScreenIE")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowUpdatesToStatusBarViaScript sets the value of InternetZoneAllowUpdatesToStatusBarViaScript for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowUpdatesToStatusBarViaScript(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowUpdatesToStatusBarViaScript", value)
}

// GetInternetZoneAllowUpdatesToStatusBarViaScript gets the value of InternetZoneAllowUpdatesToStatusBarViaScript for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowUpdatesToStatusBarViaScript() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowUpdatesToStatusBarViaScript")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowUserDataPersistence sets the value of InternetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowUserDataPersistence", value)
}

// GetInternetZoneAllowUserDataPersistence gets the value of InternetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowUserDataPersistence")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneAllowVBScriptToRunInInternetExplorer sets the value of InternetZoneAllowVBScriptToRunInInternetExplorer for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneAllowVBScriptToRunInInternetExplorer(value string) (err error) {
	return instance.SetProperty("InternetZoneAllowVBScriptToRunInInternetExplorer", value)
}

// GetInternetZoneAllowVBScriptToRunInInternetExplorer gets the value of InternetZoneAllowVBScriptToRunInInternetExplorer for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneAllowVBScriptToRunInInternetExplorer() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneAllowVBScriptToRunInInternetExplorer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneDoNotRunAntimalwareAgainstActiveXControls sets the value of InternetZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneDoNotRunAntimalwareAgainstActiveXControls(value string) (err error) {
	return instance.SetProperty("InternetZoneDoNotRunAntimalwareAgainstActiveXControls", value)
}

// GetInternetZoneDoNotRunAntimalwareAgainstActiveXControls gets the value of InternetZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneDoNotRunAntimalwareAgainstActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneDoNotRunAntimalwareAgainstActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneDownloadSignedActiveXControls sets the value of InternetZoneDownloadSignedActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneDownloadSignedActiveXControls(value string) (err error) {
	return instance.SetProperty("InternetZoneDownloadSignedActiveXControls", value)
}

// GetInternetZoneDownloadSignedActiveXControls gets the value of InternetZoneDownloadSignedActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneDownloadSignedActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneDownloadSignedActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneDownloadUnsignedActiveXControls sets the value of InternetZoneDownloadUnsignedActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneDownloadUnsignedActiveXControls(value string) (err error) {
	return instance.SetProperty("InternetZoneDownloadUnsignedActiveXControls", value)
}

// GetInternetZoneDownloadUnsignedActiveXControls gets the value of InternetZoneDownloadUnsignedActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneDownloadUnsignedActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneDownloadUnsignedActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneEnableCrossSiteScriptingFilter sets the value of InternetZoneEnableCrossSiteScriptingFilter for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneEnableCrossSiteScriptingFilter(value string) (err error) {
	return instance.SetProperty("InternetZoneEnableCrossSiteScriptingFilter", value)
}

// GetInternetZoneEnableCrossSiteScriptingFilter gets the value of InternetZoneEnableCrossSiteScriptingFilter for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneEnableCrossSiteScriptingFilter() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneEnableCrossSiteScriptingFilter")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows sets the value of InternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows(value string) (err error) {
	return instance.SetProperty("InternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows", value)
}

// GetInternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows gets the value of InternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows sets the value of InternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows(value string) (err error) {
	return instance.SetProperty("InternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows", value)
}

// GetInternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows gets the value of InternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneEnableMIMESniffing sets the value of InternetZoneEnableMIMESniffing for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneEnableMIMESniffing(value string) (err error) {
	return instance.SetProperty("InternetZoneEnableMIMESniffing", value)
}

// GetInternetZoneEnableMIMESniffing gets the value of InternetZoneEnableMIMESniffing for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneEnableMIMESniffing() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneEnableMIMESniffing")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneEnableProtectedMode sets the value of InternetZoneEnableProtectedMode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneEnableProtectedMode(value string) (err error) {
	return instance.SetProperty("InternetZoneEnableProtectedMode", value)
}

// GetInternetZoneEnableProtectedMode gets the value of InternetZoneEnableProtectedMode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneEnableProtectedMode() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneEnableProtectedMode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneIncludeLocalPathWhenUploadingFilesToServer sets the value of InternetZoneIncludeLocalPathWhenUploadingFilesToServer for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneIncludeLocalPathWhenUploadingFilesToServer(value string) (err error) {
	return instance.SetProperty("InternetZoneIncludeLocalPathWhenUploadingFilesToServer", value)
}

// GetInternetZoneIncludeLocalPathWhenUploadingFilesToServer gets the value of InternetZoneIncludeLocalPathWhenUploadingFilesToServer for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneIncludeLocalPathWhenUploadingFilesToServer() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneIncludeLocalPathWhenUploadingFilesToServer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneInitializeAndScriptActiveXControls sets the value of InternetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("InternetZoneInitializeAndScriptActiveXControls", value)
}

// GetInternetZoneInitializeAndScriptActiveXControls gets the value of InternetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneInitializeAndScriptActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneJavaPermissions sets the value of InternetZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("InternetZoneJavaPermissions", value)
}

// GetInternetZoneJavaPermissions gets the value of InternetZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneJavaPermissions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneLaunchingApplicationsAndFilesInIFRAME sets the value of InternetZoneLaunchingApplicationsAndFilesInIFRAME for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneLaunchingApplicationsAndFilesInIFRAME(value string) (err error) {
	return instance.SetProperty("InternetZoneLaunchingApplicationsAndFilesInIFRAME", value)
}

// GetInternetZoneLaunchingApplicationsAndFilesInIFRAME gets the value of InternetZoneLaunchingApplicationsAndFilesInIFRAME for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneLaunchingApplicationsAndFilesInIFRAME() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneLaunchingApplicationsAndFilesInIFRAME")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneLogonOptions sets the value of InternetZoneLogonOptions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneLogonOptions(value string) (err error) {
	return instance.SetProperty("InternetZoneLogonOptions", value)
}

// GetInternetZoneLogonOptions gets the value of InternetZoneLogonOptions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneLogonOptions() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneLogonOptions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneNavigateWindowsAndFrames sets the value of InternetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("InternetZoneNavigateWindowsAndFrames", value)
}

// GetInternetZoneNavigateWindowsAndFrames gets the value of InternetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneNavigateWindowsAndFrames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode sets the value of InternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode(value string) (err error) {
	return instance.SetProperty("InternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode", value)
}

// GetInternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode gets the value of InternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneShowSecurityWarningForPotentiallyUnsafeFiles sets the value of InternetZoneShowSecurityWarningForPotentiallyUnsafeFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneShowSecurityWarningForPotentiallyUnsafeFiles(value string) (err error) {
	return instance.SetProperty("InternetZoneShowSecurityWarningForPotentiallyUnsafeFiles", value)
}

// GetInternetZoneShowSecurityWarningForPotentiallyUnsafeFiles gets the value of InternetZoneShowSecurityWarningForPotentiallyUnsafeFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneShowSecurityWarningForPotentiallyUnsafeFiles() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneShowSecurityWarningForPotentiallyUnsafeFiles")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetZoneUsePopupBlocker sets the value of InternetZoneUsePopupBlocker for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyInternetZoneUsePopupBlocker(value string) (err error) {
	return instance.SetProperty("InternetZoneUsePopupBlocker", value)
}

// GetInternetZoneUsePopupBlocker gets the value of InternetZoneUsePopupBlocker for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyInternetZoneUsePopupBlocker() (value string, err error) {
	retValue, err := instance.GetProperty("InternetZoneUsePopupBlocker")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntranetZoneAllowAccessToDataSources sets the value of IntranetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIntranetZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowAccessToDataSources", value)
}

// GetIntranetZoneAllowAccessToDataSources gets the value of IntranetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIntranetZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowAccessToDataSources")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntranetZoneAllowAutomaticPromptingForActiveXControls sets the value of IntranetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIntranetZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowAutomaticPromptingForActiveXControls", value)
}

// GetIntranetZoneAllowAutomaticPromptingForActiveXControls gets the value of IntranetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIntranetZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowAutomaticPromptingForActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntranetZoneAllowAutomaticPromptingForFileDownloads sets the value of IntranetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIntranetZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowAutomaticPromptingForFileDownloads", value)
}

// GetIntranetZoneAllowAutomaticPromptingForFileDownloads gets the value of IntranetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIntranetZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowAutomaticPromptingForFileDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntranetZoneAllowFontDownloads sets the value of IntranetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIntranetZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowFontDownloads", value)
}

// GetIntranetZoneAllowFontDownloads gets the value of IntranetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIntranetZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowFontDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntranetZoneAllowLessPrivilegedSites sets the value of IntranetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIntranetZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowLessPrivilegedSites", value)
}

// GetIntranetZoneAllowLessPrivilegedSites gets the value of IntranetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIntranetZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowLessPrivilegedSites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntranetZoneAllowNETFrameworkReliantComponents sets the value of IntranetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIntranetZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowNETFrameworkReliantComponents", value)
}

// GetIntranetZoneAllowNETFrameworkReliantComponents gets the value of IntranetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIntranetZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowNETFrameworkReliantComponents")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntranetZoneAllowScriptlets sets the value of IntranetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIntranetZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowScriptlets", value)
}

// GetIntranetZoneAllowScriptlets gets the value of IntranetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIntranetZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowScriptlets")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntranetZoneAllowSmartScreenIE sets the value of IntranetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIntranetZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowSmartScreenIE", value)
}

// GetIntranetZoneAllowSmartScreenIE gets the value of IntranetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIntranetZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowSmartScreenIE")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntranetZoneAllowUserDataPersistence sets the value of IntranetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIntranetZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("IntranetZoneAllowUserDataPersistence", value)
}

// GetIntranetZoneAllowUserDataPersistence gets the value of IntranetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIntranetZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneAllowUserDataPersistence")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntranetZoneDoNotRunAntimalwareAgainstActiveXControls sets the value of IntranetZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIntranetZoneDoNotRunAntimalwareAgainstActiveXControls(value string) (err error) {
	return instance.SetProperty("IntranetZoneDoNotRunAntimalwareAgainstActiveXControls", value)
}

// GetIntranetZoneDoNotRunAntimalwareAgainstActiveXControls gets the value of IntranetZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIntranetZoneDoNotRunAntimalwareAgainstActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneDoNotRunAntimalwareAgainstActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntranetZoneInitializeAndScriptActiveXControls sets the value of IntranetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIntranetZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("IntranetZoneInitializeAndScriptActiveXControls", value)
}

// GetIntranetZoneInitializeAndScriptActiveXControls gets the value of IntranetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIntranetZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneInitializeAndScriptActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntranetZoneJavaPermissions sets the value of IntranetZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIntranetZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("IntranetZoneJavaPermissions", value)
}

// GetIntranetZoneJavaPermissions gets the value of IntranetZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIntranetZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneJavaPermissions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntranetZoneNavigateWindowsAndFrames sets the value of IntranetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyIntranetZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("IntranetZoneNavigateWindowsAndFrames", value)
}

// GetIntranetZoneNavigateWindowsAndFrames gets the value of IntranetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyIntranetZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("IntranetZoneNavigateWindowsAndFrames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalMachineZoneAllowAccessToDataSources sets the value of LocalMachineZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLocalMachineZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowAccessToDataSources", value)
}

// GetLocalMachineZoneAllowAccessToDataSources gets the value of LocalMachineZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLocalMachineZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowAccessToDataSources")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalMachineZoneAllowAutomaticPromptingForActiveXControls sets the value of LocalMachineZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLocalMachineZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowAutomaticPromptingForActiveXControls", value)
}

// GetLocalMachineZoneAllowAutomaticPromptingForActiveXControls gets the value of LocalMachineZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLocalMachineZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowAutomaticPromptingForActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalMachineZoneAllowAutomaticPromptingForFileDownloads sets the value of LocalMachineZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLocalMachineZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowAutomaticPromptingForFileDownloads", value)
}

// GetLocalMachineZoneAllowAutomaticPromptingForFileDownloads gets the value of LocalMachineZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLocalMachineZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowAutomaticPromptingForFileDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalMachineZoneAllowFontDownloads sets the value of LocalMachineZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLocalMachineZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowFontDownloads", value)
}

// GetLocalMachineZoneAllowFontDownloads gets the value of LocalMachineZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLocalMachineZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowFontDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalMachineZoneAllowLessPrivilegedSites sets the value of LocalMachineZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLocalMachineZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowLessPrivilegedSites", value)
}

// GetLocalMachineZoneAllowLessPrivilegedSites gets the value of LocalMachineZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLocalMachineZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowLessPrivilegedSites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalMachineZoneAllowNETFrameworkReliantComponents sets the value of LocalMachineZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLocalMachineZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowNETFrameworkReliantComponents", value)
}

// GetLocalMachineZoneAllowNETFrameworkReliantComponents gets the value of LocalMachineZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLocalMachineZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowNETFrameworkReliantComponents")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalMachineZoneAllowScriptlets sets the value of LocalMachineZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLocalMachineZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowScriptlets", value)
}

// GetLocalMachineZoneAllowScriptlets gets the value of LocalMachineZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLocalMachineZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowScriptlets")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalMachineZoneAllowSmartScreenIE sets the value of LocalMachineZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLocalMachineZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowSmartScreenIE", value)
}

// GetLocalMachineZoneAllowSmartScreenIE gets the value of LocalMachineZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLocalMachineZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowSmartScreenIE")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalMachineZoneAllowUserDataPersistence sets the value of LocalMachineZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLocalMachineZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneAllowUserDataPersistence", value)
}

// GetLocalMachineZoneAllowUserDataPersistence gets the value of LocalMachineZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLocalMachineZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneAllowUserDataPersistence")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls sets the value of LocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls", value)
}

// GetLocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls gets the value of LocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneDoNotRunAntimalwareAgainstActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalMachineZoneInitializeAndScriptActiveXControls sets the value of LocalMachineZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLocalMachineZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneInitializeAndScriptActiveXControls", value)
}

// GetLocalMachineZoneInitializeAndScriptActiveXControls gets the value of LocalMachineZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLocalMachineZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneInitializeAndScriptActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalMachineZoneJavaPermissions sets the value of LocalMachineZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLocalMachineZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneJavaPermissions", value)
}

// GetLocalMachineZoneJavaPermissions gets the value of LocalMachineZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLocalMachineZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneJavaPermissions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalMachineZoneNavigateWindowsAndFrames sets the value of LocalMachineZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLocalMachineZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("LocalMachineZoneNavigateWindowsAndFrames", value)
}

// GetLocalMachineZoneNavigateWindowsAndFrames gets the value of LocalMachineZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLocalMachineZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("LocalMachineZoneNavigateWindowsAndFrames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownInternetZoneAllowAccessToDataSources sets the value of LockedDownInternetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowAccessToDataSources", value)
}

// GetLockedDownInternetZoneAllowAccessToDataSources gets the value of LockedDownInternetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowAccessToDataSources")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownInternetZoneAllowAutomaticPromptingForActiveXControls sets the value of LockedDownInternetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowAutomaticPromptingForActiveXControls", value)
}

// GetLockedDownInternetZoneAllowAutomaticPromptingForActiveXControls gets the value of LockedDownInternetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowAutomaticPromptingForActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownInternetZoneAllowAutomaticPromptingForFileDownloads sets the value of LockedDownInternetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowAutomaticPromptingForFileDownloads", value)
}

// GetLockedDownInternetZoneAllowAutomaticPromptingForFileDownloads gets the value of LockedDownInternetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowAutomaticPromptingForFileDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownInternetZoneAllowFontDownloads sets the value of LockedDownInternetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowFontDownloads", value)
}

// GetLockedDownInternetZoneAllowFontDownloads gets the value of LockedDownInternetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowFontDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownInternetZoneAllowLessPrivilegedSites sets the value of LockedDownInternetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowLessPrivilegedSites", value)
}

// GetLockedDownInternetZoneAllowLessPrivilegedSites gets the value of LockedDownInternetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowLessPrivilegedSites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownInternetZoneAllowNETFrameworkReliantComponents sets the value of LockedDownInternetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowNETFrameworkReliantComponents", value)
}

// GetLockedDownInternetZoneAllowNETFrameworkReliantComponents gets the value of LockedDownInternetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowNETFrameworkReliantComponents")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownInternetZoneAllowScriptlets sets the value of LockedDownInternetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowScriptlets", value)
}

// GetLockedDownInternetZoneAllowScriptlets gets the value of LockedDownInternetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowScriptlets")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownInternetZoneAllowSmartScreenIE sets the value of LockedDownInternetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowSmartScreenIE", value)
}

// GetLockedDownInternetZoneAllowSmartScreenIE gets the value of LockedDownInternetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowSmartScreenIE")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownInternetZoneAllowUserDataPersistence sets the value of LockedDownInternetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownInternetZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneAllowUserDataPersistence", value)
}

// GetLockedDownInternetZoneAllowUserDataPersistence gets the value of LockedDownInternetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownInternetZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneAllowUserDataPersistence")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownInternetZoneInitializeAndScriptActiveXControls sets the value of LockedDownInternetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownInternetZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneInitializeAndScriptActiveXControls", value)
}

// GetLockedDownInternetZoneInitializeAndScriptActiveXControls gets the value of LockedDownInternetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownInternetZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneInitializeAndScriptActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownInternetZoneJavaPermissions sets the value of LockedDownInternetZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownInternetZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneJavaPermissions", value)
}

// GetLockedDownInternetZoneJavaPermissions gets the value of LockedDownInternetZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownInternetZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneJavaPermissions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownInternetZoneNavigateWindowsAndFrames sets the value of LockedDownInternetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownInternetZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("LockedDownInternetZoneNavigateWindowsAndFrames", value)
}

// GetLockedDownInternetZoneNavigateWindowsAndFrames gets the value of LockedDownInternetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownInternetZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownInternetZoneNavigateWindowsAndFrames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownIntranetJavaPermissions sets the value of LockedDownIntranetJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownIntranetJavaPermissions(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetJavaPermissions", value)
}

// GetLockedDownIntranetJavaPermissions gets the value of LockedDownIntranetJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownIntranetJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetJavaPermissions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownIntranetZoneAllowAccessToDataSources sets the value of LockedDownIntranetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowAccessToDataSources", value)
}

// GetLockedDownIntranetZoneAllowAccessToDataSources gets the value of LockedDownIntranetZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowAccessToDataSources")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls sets the value of LockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls", value)
}

// GetLockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls gets the value of LockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowAutomaticPromptingForActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads sets the value of LockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads", value)
}

// GetLockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads gets the value of LockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowAutomaticPromptingForFileDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownIntranetZoneAllowFontDownloads sets the value of LockedDownIntranetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowFontDownloads", value)
}

// GetLockedDownIntranetZoneAllowFontDownloads gets the value of LockedDownIntranetZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowFontDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownIntranetZoneAllowLessPrivilegedSites sets the value of LockedDownIntranetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowLessPrivilegedSites", value)
}

// GetLockedDownIntranetZoneAllowLessPrivilegedSites gets the value of LockedDownIntranetZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowLessPrivilegedSites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownIntranetZoneAllowNETFrameworkReliantComponents sets the value of LockedDownIntranetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowNETFrameworkReliantComponents", value)
}

// GetLockedDownIntranetZoneAllowNETFrameworkReliantComponents gets the value of LockedDownIntranetZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowNETFrameworkReliantComponents")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownIntranetZoneAllowScriptlets sets the value of LockedDownIntranetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowScriptlets", value)
}

// GetLockedDownIntranetZoneAllowScriptlets gets the value of LockedDownIntranetZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowScriptlets")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownIntranetZoneAllowSmartScreenIE sets the value of LockedDownIntranetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowSmartScreenIE", value)
}

// GetLockedDownIntranetZoneAllowSmartScreenIE gets the value of LockedDownIntranetZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowSmartScreenIE")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownIntranetZoneAllowUserDataPersistence sets the value of LockedDownIntranetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownIntranetZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneAllowUserDataPersistence", value)
}

// GetLockedDownIntranetZoneAllowUserDataPersistence gets the value of LockedDownIntranetZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownIntranetZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneAllowUserDataPersistence")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownIntranetZoneInitializeAndScriptActiveXControls sets the value of LockedDownIntranetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownIntranetZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneInitializeAndScriptActiveXControls", value)
}

// GetLockedDownIntranetZoneInitializeAndScriptActiveXControls gets the value of LockedDownIntranetZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownIntranetZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneInitializeAndScriptActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownIntranetZoneNavigateWindowsAndFrames sets the value of LockedDownIntranetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownIntranetZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("LockedDownIntranetZoneNavigateWindowsAndFrames", value)
}

// GetLockedDownIntranetZoneNavigateWindowsAndFrames gets the value of LockedDownIntranetZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownIntranetZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownIntranetZoneNavigateWindowsAndFrames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownLocalMachineZoneAllowAccessToDataSources sets the value of LockedDownLocalMachineZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowAccessToDataSources", value)
}

// GetLockedDownLocalMachineZoneAllowAccessToDataSources gets the value of LockedDownLocalMachineZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowAccessToDataSources")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls sets the value of LockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls", value)
}

// GetLockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls gets the value of LockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowAutomaticPromptingForActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads sets the value of LockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads", value)
}

// GetLockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads gets the value of LockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowAutomaticPromptingForFileDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownLocalMachineZoneAllowFontDownloads sets the value of LockedDownLocalMachineZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowFontDownloads", value)
}

// GetLockedDownLocalMachineZoneAllowFontDownloads gets the value of LockedDownLocalMachineZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowFontDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownLocalMachineZoneAllowLessPrivilegedSites sets the value of LockedDownLocalMachineZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowLessPrivilegedSites", value)
}

// GetLockedDownLocalMachineZoneAllowLessPrivilegedSites gets the value of LockedDownLocalMachineZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowLessPrivilegedSites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownLocalMachineZoneAllowNETFrameworkReliantComponents sets the value of LockedDownLocalMachineZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowNETFrameworkReliantComponents", value)
}

// GetLockedDownLocalMachineZoneAllowNETFrameworkReliantComponents gets the value of LockedDownLocalMachineZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowNETFrameworkReliantComponents")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownLocalMachineZoneAllowScriptlets sets the value of LockedDownLocalMachineZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowScriptlets", value)
}

// GetLockedDownLocalMachineZoneAllowScriptlets gets the value of LockedDownLocalMachineZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowScriptlets")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownLocalMachineZoneAllowSmartScreenIE sets the value of LockedDownLocalMachineZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowSmartScreenIE", value)
}

// GetLockedDownLocalMachineZoneAllowSmartScreenIE gets the value of LockedDownLocalMachineZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowSmartScreenIE")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownLocalMachineZoneAllowUserDataPersistence sets the value of LockedDownLocalMachineZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneAllowUserDataPersistence", value)
}

// GetLockedDownLocalMachineZoneAllowUserDataPersistence gets the value of LockedDownLocalMachineZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneAllowUserDataPersistence")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownLocalMachineZoneInitializeAndScriptActiveXControls sets the value of LockedDownLocalMachineZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneInitializeAndScriptActiveXControls", value)
}

// GetLockedDownLocalMachineZoneInitializeAndScriptActiveXControls gets the value of LockedDownLocalMachineZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneInitializeAndScriptActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownLocalMachineZoneJavaPermissions sets the value of LockedDownLocalMachineZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneJavaPermissions", value)
}

// GetLockedDownLocalMachineZoneJavaPermissions gets the value of LockedDownLocalMachineZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneJavaPermissions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownLocalMachineZoneNavigateWindowsAndFrames sets the value of LockedDownLocalMachineZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownLocalMachineZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("LockedDownLocalMachineZoneNavigateWindowsAndFrames", value)
}

// GetLockedDownLocalMachineZoneNavigateWindowsAndFrames gets the value of LockedDownLocalMachineZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownLocalMachineZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownLocalMachineZoneNavigateWindowsAndFrames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownRestrictedSitesZoneAllowAccessToDataSources sets the value of LockedDownRestrictedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowAccessToDataSources", value)
}

// GetLockedDownRestrictedSitesZoneAllowAccessToDataSources gets the value of LockedDownRestrictedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowAccessToDataSources")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls sets the value of LockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls", value)
}

// GetLockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls gets the value of LockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads sets the value of LockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads", value)
}

// GetLockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads gets the value of LockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownRestrictedSitesZoneAllowFontDownloads sets the value of LockedDownRestrictedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowFontDownloads", value)
}

// GetLockedDownRestrictedSitesZoneAllowFontDownloads gets the value of LockedDownRestrictedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowFontDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownRestrictedSitesZoneAllowLessPrivilegedSites sets the value of LockedDownRestrictedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowLessPrivilegedSites", value)
}

// GetLockedDownRestrictedSitesZoneAllowLessPrivilegedSites gets the value of LockedDownRestrictedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowLessPrivilegedSites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents sets the value of LockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents", value)
}

// GetLockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents gets the value of LockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowNETFrameworkReliantComponents")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownRestrictedSitesZoneAllowScriptlets sets the value of LockedDownRestrictedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowScriptlets", value)
}

// GetLockedDownRestrictedSitesZoneAllowScriptlets gets the value of LockedDownRestrictedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowScriptlets")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownRestrictedSitesZoneAllowSmartScreenIE sets the value of LockedDownRestrictedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowSmartScreenIE", value)
}

// GetLockedDownRestrictedSitesZoneAllowSmartScreenIE gets the value of LockedDownRestrictedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowSmartScreenIE")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownRestrictedSitesZoneAllowUserDataPersistence sets the value of LockedDownRestrictedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneAllowUserDataPersistence", value)
}

// GetLockedDownRestrictedSitesZoneAllowUserDataPersistence gets the value of LockedDownRestrictedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneAllowUserDataPersistence")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls sets the value of LockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls", value)
}

// GetLockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls gets the value of LockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneInitializeAndScriptActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownRestrictedSitesZoneJavaPermissions sets the value of LockedDownRestrictedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneJavaPermissions", value)
}

// GetLockedDownRestrictedSitesZoneJavaPermissions gets the value of LockedDownRestrictedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneJavaPermissions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownRestrictedSitesZoneNavigateWindowsAndFrames sets the value of LockedDownRestrictedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownRestrictedSitesZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("LockedDownRestrictedSitesZoneNavigateWindowsAndFrames", value)
}

// GetLockedDownRestrictedSitesZoneNavigateWindowsAndFrames gets the value of LockedDownRestrictedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownRestrictedSitesZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownRestrictedSitesZoneNavigateWindowsAndFrames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownTrustedSitesZoneAllowAccessToDataSources sets the value of LockedDownTrustedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowAccessToDataSources", value)
}

// GetLockedDownTrustedSitesZoneAllowAccessToDataSources gets the value of LockedDownTrustedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowAccessToDataSources")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls sets the value of LockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls", value)
}

// GetLockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls gets the value of LockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowAutomaticPromptingForActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads sets the value of LockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads", value)
}

// GetLockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads gets the value of LockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowAutomaticPromptingForFileDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownTrustedSitesZoneAllowFontDownloads sets the value of LockedDownTrustedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowFontDownloads", value)
}

// GetLockedDownTrustedSitesZoneAllowFontDownloads gets the value of LockedDownTrustedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowFontDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownTrustedSitesZoneAllowLessPrivilegedSites sets the value of LockedDownTrustedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowLessPrivilegedSites", value)
}

// GetLockedDownTrustedSitesZoneAllowLessPrivilegedSites gets the value of LockedDownTrustedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowLessPrivilegedSites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents sets the value of LockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents", value)
}

// GetLockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents gets the value of LockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowNETFrameworkReliantComponents")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownTrustedSitesZoneAllowScriptlets sets the value of LockedDownTrustedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowScriptlets", value)
}

// GetLockedDownTrustedSitesZoneAllowScriptlets gets the value of LockedDownTrustedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowScriptlets")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownTrustedSitesZoneAllowSmartScreenIE sets the value of LockedDownTrustedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowSmartScreenIE", value)
}

// GetLockedDownTrustedSitesZoneAllowSmartScreenIE gets the value of LockedDownTrustedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowSmartScreenIE")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownTrustedSitesZoneAllowUserDataPersistence sets the value of LockedDownTrustedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneAllowUserDataPersistence", value)
}

// GetLockedDownTrustedSitesZoneAllowUserDataPersistence gets the value of LockedDownTrustedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneAllowUserDataPersistence")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownTrustedSitesZoneInitializeAndScriptActiveXControls sets the value of LockedDownTrustedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneInitializeAndScriptActiveXControls", value)
}

// GetLockedDownTrustedSitesZoneInitializeAndScriptActiveXControls gets the value of LockedDownTrustedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneInitializeAndScriptActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownTrustedSitesZoneJavaPermissions sets the value of LockedDownTrustedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneJavaPermissions", value)
}

// GetLockedDownTrustedSitesZoneJavaPermissions gets the value of LockedDownTrustedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneJavaPermissions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockedDownTrustedSitesZoneNavigateWindowsAndFrames sets the value of LockedDownTrustedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyLockedDownTrustedSitesZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("LockedDownTrustedSitesZoneNavigateWindowsAndFrames", value)
}

// GetLockedDownTrustedSitesZoneNavigateWindowsAndFrames gets the value of LockedDownTrustedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyLockedDownTrustedSitesZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("LockedDownTrustedSitesZoneNavigateWindowsAndFrames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMimeSniffingSafetyFeatureInternetExplorerProcesses sets the value of MimeSniffingSafetyFeatureInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyMimeSniffingSafetyFeatureInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("MimeSniffingSafetyFeatureInternetExplorerProcesses", value)
}

// GetMimeSniffingSafetyFeatureInternetExplorerProcesses gets the value of MimeSniffingSafetyFeatureInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyMimeSniffingSafetyFeatureInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("MimeSniffingSafetyFeatureInternetExplorerProcesses")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMKProtocolSecurityRestrictionInternetExplorerProcesses sets the value of MKProtocolSecurityRestrictionInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyMKProtocolSecurityRestrictionInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("MKProtocolSecurityRestrictionInternetExplorerProcesses", value)
}

// GetMKProtocolSecurityRestrictionInternetExplorerProcesses gets the value of MKProtocolSecurityRestrictionInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyMKProtocolSecurityRestrictionInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("MKProtocolSecurityRestrictionInternetExplorerProcesses")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNewTabDefaultPage sets the value of NewTabDefaultPage for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyNewTabDefaultPage(value string) (err error) {
	return instance.SetProperty("NewTabDefaultPage", value)
}

// GetNewTabDefaultPage gets the value of NewTabDefaultPage for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyNewTabDefaultPage() (value string, err error) {
	retValue, err := instance.GetProperty("NewTabDefaultPage")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotificationBarInternetExplorerProcesses sets the value of NotificationBarInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyNotificationBarInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("NotificationBarInternetExplorerProcesses", value)
}

// GetNotificationBarInternetExplorerProcesses gets the value of NotificationBarInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyNotificationBarInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("NotificationBarInternetExplorerProcesses")
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
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyParentID() (value string, err error) {
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

// SetPreventManagingSmartScreenFilter sets the value of PreventManagingSmartScreenFilter for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyPreventManagingSmartScreenFilter(value string) (err error) {
	return instance.SetProperty("PreventManagingSmartScreenFilter", value)
}

// GetPreventManagingSmartScreenFilter gets the value of PreventManagingSmartScreenFilter for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyPreventManagingSmartScreenFilter() (value string, err error) {
	retValue, err := instance.GetProperty("PreventManagingSmartScreenFilter")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreventPerUserInstallationOfActiveXControls sets the value of PreventPerUserInstallationOfActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyPreventPerUserInstallationOfActiveXControls(value string) (err error) {
	return instance.SetProperty("PreventPerUserInstallationOfActiveXControls", value)
}

// GetPreventPerUserInstallationOfActiveXControls gets the value of PreventPerUserInstallationOfActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyPreventPerUserInstallationOfActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("PreventPerUserInstallationOfActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtectionFromZoneElevationInternetExplorerProcesses sets the value of ProtectionFromZoneElevationInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyProtectionFromZoneElevationInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("ProtectionFromZoneElevationInternetExplorerProcesses", value)
}

// GetProtectionFromZoneElevationInternetExplorerProcesses gets the value of ProtectionFromZoneElevationInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyProtectionFromZoneElevationInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("ProtectionFromZoneElevationInternetExplorerProcesses")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoveRunThisTimeButtonForOutdatedActiveXControls sets the value of RemoveRunThisTimeButtonForOutdatedActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRemoveRunThisTimeButtonForOutdatedActiveXControls(value string) (err error) {
	return instance.SetProperty("RemoveRunThisTimeButtonForOutdatedActiveXControls", value)
}

// GetRemoveRunThisTimeButtonForOutdatedActiveXControls gets the value of RemoveRunThisTimeButtonForOutdatedActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRemoveRunThisTimeButtonForOutdatedActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RemoveRunThisTimeButtonForOutdatedActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictActiveXInstallInternetExplorerProcesses sets the value of RestrictActiveXInstallInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictActiveXInstallInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("RestrictActiveXInstallInternetExplorerProcesses", value)
}

// GetRestrictActiveXInstallInternetExplorerProcesses gets the value of RestrictActiveXInstallInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictActiveXInstallInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictActiveXInstallInternetExplorerProcesses")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowAccessToDataSources sets the value of RestrictedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowAccessToDataSources", value)
}

// GetRestrictedSitesZoneAllowAccessToDataSources gets the value of RestrictedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowAccessToDataSources")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowActiveScripting sets the value of RestrictedSitesZoneAllowActiveScripting for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowActiveScripting(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowActiveScripting", value)
}

// GetRestrictedSitesZoneAllowActiveScripting gets the value of RestrictedSitesZoneAllowActiveScripting for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowActiveScripting() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowActiveScripting")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls sets the value of RestrictedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowAutomaticPromptingForActiveXControls", value)
}

// GetRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls gets the value of RestrictedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowAutomaticPromptingForActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads sets the value of RestrictedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowAutomaticPromptingForFileDownloads", value)
}

// GetRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads gets the value of RestrictedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowAutomaticPromptingForFileDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowBinaryAndScriptBehaviors sets the value of RestrictedSitesZoneAllowBinaryAndScriptBehaviors for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowBinaryAndScriptBehaviors(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowBinaryAndScriptBehaviors", value)
}

// GetRestrictedSitesZoneAllowBinaryAndScriptBehaviors gets the value of RestrictedSitesZoneAllowBinaryAndScriptBehaviors for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowBinaryAndScriptBehaviors() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowBinaryAndScriptBehaviors")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowCopyPasteViaScript sets the value of RestrictedSitesZoneAllowCopyPasteViaScript for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowCopyPasteViaScript(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowCopyPasteViaScript", value)
}

// GetRestrictedSitesZoneAllowCopyPasteViaScript gets the value of RestrictedSitesZoneAllowCopyPasteViaScript for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowCopyPasteViaScript() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowCopyPasteViaScript")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles sets the value of RestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles", value)
}

// GetRestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles gets the value of RestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowDragAndDropCopyAndPasteFiles")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowFileDownloads sets the value of RestrictedSitesZoneAllowFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowFileDownloads(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowFileDownloads", value)
}

// GetRestrictedSitesZoneAllowFileDownloads gets the value of RestrictedSitesZoneAllowFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowFileDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowFontDownloads sets the value of RestrictedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowFontDownloads", value)
}

// GetRestrictedSitesZoneAllowFontDownloads gets the value of RestrictedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowFontDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowLessPrivilegedSites sets the value of RestrictedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowLessPrivilegedSites", value)
}

// GetRestrictedSitesZoneAllowLessPrivilegedSites gets the value of RestrictedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowLessPrivilegedSites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowLoadingOfXAMLFiles sets the value of RestrictedSitesZoneAllowLoadingOfXAMLFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowLoadingOfXAMLFiles(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowLoadingOfXAMLFiles", value)
}

// GetRestrictedSitesZoneAllowLoadingOfXAMLFiles gets the value of RestrictedSitesZoneAllowLoadingOfXAMLFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowLoadingOfXAMLFiles() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowLoadingOfXAMLFiles")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowMETAREFRESH sets the value of RestrictedSitesZoneAllowMETAREFRESH for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowMETAREFRESH(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowMETAREFRESH", value)
}

// GetRestrictedSitesZoneAllowMETAREFRESH gets the value of RestrictedSitesZoneAllowMETAREFRESH for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowMETAREFRESH() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowMETAREFRESH")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowNETFrameworkReliantComponents sets the value of RestrictedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowNETFrameworkReliantComponents", value)
}

// GetRestrictedSitesZoneAllowNETFrameworkReliantComponents gets the value of RestrictedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowNETFrameworkReliantComponents")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls sets the value of RestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls", value)
}

// GetRestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls gets the value of RestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowOnlyApprovedDomainsToUseActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl sets the value of RestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl", value)
}

// GetRestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl gets the value of RestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowOnlyApprovedDomainsToUseTDCActiveXControl")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls sets the value of RestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls", value)
}

// GetRestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls gets the value of RestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowScriptingOfInternetExplorerWebBrowserControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowScriptInitiatedWindows sets the value of RestrictedSitesZoneAllowScriptInitiatedWindows for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowScriptInitiatedWindows(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowScriptInitiatedWindows", value)
}

// GetRestrictedSitesZoneAllowScriptInitiatedWindows gets the value of RestrictedSitesZoneAllowScriptInitiatedWindows for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowScriptInitiatedWindows() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowScriptInitiatedWindows")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowScriptlets sets the value of RestrictedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowScriptlets", value)
}

// GetRestrictedSitesZoneAllowScriptlets gets the value of RestrictedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowScriptlets")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowSmartScreenIE sets the value of RestrictedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowSmartScreenIE", value)
}

// GetRestrictedSitesZoneAllowSmartScreenIE gets the value of RestrictedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowSmartScreenIE")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowUpdatesToStatusBarViaScript sets the value of RestrictedSitesZoneAllowUpdatesToStatusBarViaScript for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowUpdatesToStatusBarViaScript(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowUpdatesToStatusBarViaScript", value)
}

// GetRestrictedSitesZoneAllowUpdatesToStatusBarViaScript gets the value of RestrictedSitesZoneAllowUpdatesToStatusBarViaScript for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowUpdatesToStatusBarViaScript() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowUpdatesToStatusBarViaScript")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowUserDataPersistence sets the value of RestrictedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowUserDataPersistence", value)
}

// GetRestrictedSitesZoneAllowUserDataPersistence gets the value of RestrictedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowUserDataPersistence")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneAllowVBScriptToRunInInternetExplorer sets the value of RestrictedSitesZoneAllowVBScriptToRunInInternetExplorer for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneAllowVBScriptToRunInInternetExplorer(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneAllowVBScriptToRunInInternetExplorer", value)
}

// GetRestrictedSitesZoneAllowVBScriptToRunInInternetExplorer gets the value of RestrictedSitesZoneAllowVBScriptToRunInInternetExplorer for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneAllowVBScriptToRunInInternetExplorer() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneAllowVBScriptToRunInInternetExplorer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls sets the value of RestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls", value)
}

// GetRestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls gets the value of RestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneDoNotRunAntimalwareAgainstActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneDownloadSignedActiveXControls sets the value of RestrictedSitesZoneDownloadSignedActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneDownloadSignedActiveXControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneDownloadSignedActiveXControls", value)
}

// GetRestrictedSitesZoneDownloadSignedActiveXControls gets the value of RestrictedSitesZoneDownloadSignedActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneDownloadSignedActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneDownloadSignedActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneDownloadUnsignedActiveXControls sets the value of RestrictedSitesZoneDownloadUnsignedActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneDownloadUnsignedActiveXControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneDownloadUnsignedActiveXControls", value)
}

// GetRestrictedSitesZoneDownloadUnsignedActiveXControls gets the value of RestrictedSitesZoneDownloadUnsignedActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneDownloadUnsignedActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneDownloadUnsignedActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneEnableCrossSiteScriptingFilter sets the value of RestrictedSitesZoneEnableCrossSiteScriptingFilter for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneEnableCrossSiteScriptingFilter(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneEnableCrossSiteScriptingFilter", value)
}

// GetRestrictedSitesZoneEnableCrossSiteScriptingFilter gets the value of RestrictedSitesZoneEnableCrossSiteScriptingFilter for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneEnableCrossSiteScriptingFilter() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneEnableCrossSiteScriptingFilter")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows sets the value of RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows", value)
}

// GetRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows gets the value of RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsAcrossWindows")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows sets the value of RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows", value)
}

// GetRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows gets the value of RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneEnableDraggingOfContentFromDifferentDomainsWithinWindows")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneEnableMIMESniffing sets the value of RestrictedSitesZoneEnableMIMESniffing for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneEnableMIMESniffing(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneEnableMIMESniffing", value)
}

// GetRestrictedSitesZoneEnableMIMESniffing gets the value of RestrictedSitesZoneEnableMIMESniffing for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneEnableMIMESniffing() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneEnableMIMESniffing")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer sets the value of RestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer", value)
}

// GetRestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer gets the value of RestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneIncludeLocalPathWhenUploadingFilesToServer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneInitializeAndScriptActiveXControls sets the value of RestrictedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneInitializeAndScriptActiveXControls", value)
}

// GetRestrictedSitesZoneInitializeAndScriptActiveXControls gets the value of RestrictedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneInitializeAndScriptActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneJavaPermissions sets the value of RestrictedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneJavaPermissions", value)
}

// GetRestrictedSitesZoneJavaPermissions gets the value of RestrictedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneJavaPermissions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME sets the value of RestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME", value)
}

// GetRestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME gets the value of RestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneLaunchingApplicationsAndFilesInIFRAME")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneLogonOptions sets the value of RestrictedSitesZoneLogonOptions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneLogonOptions(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneLogonOptions", value)
}

// GetRestrictedSitesZoneLogonOptions gets the value of RestrictedSitesZoneLogonOptions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneLogonOptions() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneLogonOptions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneNavigateWindowsAndFrames sets the value of RestrictedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneNavigateWindowsAndFrames", value)
}

// GetRestrictedSitesZoneNavigateWindowsAndFrames gets the value of RestrictedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneNavigateWindowsAndFrames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneRunActiveXControlsAndPlugins sets the value of RestrictedSitesZoneRunActiveXControlsAndPlugins for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneRunActiveXControlsAndPlugins(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneRunActiveXControlsAndPlugins", value)
}

// GetRestrictedSitesZoneRunActiveXControlsAndPlugins gets the value of RestrictedSitesZoneRunActiveXControlsAndPlugins for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneRunActiveXControlsAndPlugins() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneRunActiveXControlsAndPlugins")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode sets the value of RestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode", value)
}

// GetRestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode gets the value of RestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneRunNETFrameworkReliantComponentsSignedWithAuthenticode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting sets the value of RestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting", value)
}

// GetRestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting gets the value of RestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneScriptActiveXControlsMarkedSafeForScripting")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneScriptingOfJavaApplets sets the value of RestrictedSitesZoneScriptingOfJavaApplets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneScriptingOfJavaApplets(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneScriptingOfJavaApplets", value)
}

// GetRestrictedSitesZoneScriptingOfJavaApplets gets the value of RestrictedSitesZoneScriptingOfJavaApplets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneScriptingOfJavaApplets() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneScriptingOfJavaApplets")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles sets the value of RestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles", value)
}

// GetRestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles gets the value of RestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneShowSecurityWarningForPotentiallyUnsafeFiles")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneTurnOnProtectedMode sets the value of RestrictedSitesZoneTurnOnProtectedMode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneTurnOnProtectedMode(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneTurnOnProtectedMode", value)
}

// GetRestrictedSitesZoneTurnOnProtectedMode gets the value of RestrictedSitesZoneTurnOnProtectedMode for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneTurnOnProtectedMode() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneTurnOnProtectedMode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSitesZoneUsePopupBlocker sets the value of RestrictedSitesZoneUsePopupBlocker for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictedSitesZoneUsePopupBlocker(value string) (err error) {
	return instance.SetProperty("RestrictedSitesZoneUsePopupBlocker", value)
}

// GetRestrictedSitesZoneUsePopupBlocker gets the value of RestrictedSitesZoneUsePopupBlocker for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictedSitesZoneUsePopupBlocker() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneUsePopupBlocker")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictFileDownloadInternetExplorerProcesses sets the value of RestrictFileDownloadInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyRestrictFileDownloadInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("RestrictFileDownloadInternetExplorerProcesses", value)
}

// GetRestrictFileDownloadInternetExplorerProcesses gets the value of RestrictFileDownloadInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyRestrictFileDownloadInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictFileDownloadInternetExplorerProcesses")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScriptedWindowSecurityRestrictionsInternetExplorerProcesses sets the value of ScriptedWindowSecurityRestrictionsInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyScriptedWindowSecurityRestrictionsInternetExplorerProcesses(value string) (err error) {
	return instance.SetProperty("ScriptedWindowSecurityRestrictionsInternetExplorerProcesses", value)
}

// GetScriptedWindowSecurityRestrictionsInternetExplorerProcesses gets the value of ScriptedWindowSecurityRestrictionsInternetExplorerProcesses for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyScriptedWindowSecurityRestrictionsInternetExplorerProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("ScriptedWindowSecurityRestrictionsInternetExplorerProcesses")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSearchProviderList sets the value of SearchProviderList for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertySearchProviderList(value string) (err error) {
	return instance.SetProperty("SearchProviderList", value)
}

// GetSearchProviderList gets the value of SearchProviderList for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertySearchProviderList() (value string, err error) {
	retValue, err := instance.GetProperty("SearchProviderList")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecurityZonesUseOnlyMachineSettings sets the value of SecurityZonesUseOnlyMachineSettings for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertySecurityZonesUseOnlyMachineSettings(value string) (err error) {
	return instance.SetProperty("SecurityZonesUseOnlyMachineSettings", value)
}

// GetSecurityZonesUseOnlyMachineSettings gets the value of SecurityZonesUseOnlyMachineSettings for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertySecurityZonesUseOnlyMachineSettings() (value string, err error) {
	retValue, err := instance.GetProperty("SecurityZonesUseOnlyMachineSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpecifyUseOfActiveXInstallerService sets the value of SpecifyUseOfActiveXInstallerService for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertySpecifyUseOfActiveXInstallerService(value string) (err error) {
	return instance.SetProperty("SpecifyUseOfActiveXInstallerService", value)
}

// GetSpecifyUseOfActiveXInstallerService gets the value of SpecifyUseOfActiveXInstallerService for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertySpecifyUseOfActiveXInstallerService() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyUseOfActiveXInstallerService")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedSitesZoneAllowAccessToDataSources sets the value of TrustedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowAccessToDataSources(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowAccessToDataSources", value)
}

// GetTrustedSitesZoneAllowAccessToDataSources gets the value of TrustedSitesZoneAllowAccessToDataSources for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowAccessToDataSources() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowAccessToDataSources")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedSitesZoneAllowAutomaticPromptingForActiveXControls sets the value of TrustedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowAutomaticPromptingForActiveXControls(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowAutomaticPromptingForActiveXControls", value)
}

// GetTrustedSitesZoneAllowAutomaticPromptingForActiveXControls gets the value of TrustedSitesZoneAllowAutomaticPromptingForActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowAutomaticPromptingForActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowAutomaticPromptingForActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedSitesZoneAllowAutomaticPromptingForFileDownloads sets the value of TrustedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowAutomaticPromptingForFileDownloads(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowAutomaticPromptingForFileDownloads", value)
}

// GetTrustedSitesZoneAllowAutomaticPromptingForFileDownloads gets the value of TrustedSitesZoneAllowAutomaticPromptingForFileDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowAutomaticPromptingForFileDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowAutomaticPromptingForFileDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedSitesZoneAllowFontDownloads sets the value of TrustedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowFontDownloads(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowFontDownloads", value)
}

// GetTrustedSitesZoneAllowFontDownloads gets the value of TrustedSitesZoneAllowFontDownloads for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowFontDownloads() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowFontDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedSitesZoneAllowLessPrivilegedSites sets the value of TrustedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowLessPrivilegedSites(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowLessPrivilegedSites", value)
}

// GetTrustedSitesZoneAllowLessPrivilegedSites gets the value of TrustedSitesZoneAllowLessPrivilegedSites for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowLessPrivilegedSites() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowLessPrivilegedSites")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedSitesZoneAllowNETFrameworkReliantComponents sets the value of TrustedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowNETFrameworkReliantComponents(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowNETFrameworkReliantComponents", value)
}

// GetTrustedSitesZoneAllowNETFrameworkReliantComponents gets the value of TrustedSitesZoneAllowNETFrameworkReliantComponents for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowNETFrameworkReliantComponents() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowNETFrameworkReliantComponents")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedSitesZoneAllowScriptlets sets the value of TrustedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowScriptlets(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowScriptlets", value)
}

// GetTrustedSitesZoneAllowScriptlets gets the value of TrustedSitesZoneAllowScriptlets for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowScriptlets() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowScriptlets")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedSitesZoneAllowSmartScreenIE sets the value of TrustedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowSmartScreenIE(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowSmartScreenIE", value)
}

// GetTrustedSitesZoneAllowSmartScreenIE gets the value of TrustedSitesZoneAllowSmartScreenIE for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowSmartScreenIE() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowSmartScreenIE")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedSitesZoneAllowUserDataPersistence sets the value of TrustedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyTrustedSitesZoneAllowUserDataPersistence(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneAllowUserDataPersistence", value)
}

// GetTrustedSitesZoneAllowUserDataPersistence gets the value of TrustedSitesZoneAllowUserDataPersistence for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyTrustedSitesZoneAllowUserDataPersistence() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneAllowUserDataPersistence")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls sets the value of TrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyTrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls", value)
}

// GetTrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls gets the value of TrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyTrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneDoNotRunAntimalwareAgainstActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedSitesZoneInitializeAndScriptActiveXControls sets the value of TrustedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyTrustedSitesZoneInitializeAndScriptActiveXControls(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneInitializeAndScriptActiveXControls", value)
}

// GetTrustedSitesZoneInitializeAndScriptActiveXControls gets the value of TrustedSitesZoneInitializeAndScriptActiveXControls for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyTrustedSitesZoneInitializeAndScriptActiveXControls() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneInitializeAndScriptActiveXControls")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedSitesZoneJavaPermissions sets the value of TrustedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyTrustedSitesZoneJavaPermissions(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneJavaPermissions", value)
}

// GetTrustedSitesZoneJavaPermissions gets the value of TrustedSitesZoneJavaPermissions for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyTrustedSitesZoneJavaPermissions() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneJavaPermissions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedSitesZoneNavigateWindowsAndFrames sets the value of TrustedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) SetPropertyTrustedSitesZoneNavigateWindowsAndFrames(value string) (err error) {
	return instance.SetProperty("TrustedSitesZoneNavigateWindowsAndFrames", value)
}

// GetTrustedSitesZoneNavigateWindowsAndFrames gets the value of TrustedSitesZoneNavigateWindowsAndFrames for the instance
func (instance *MDM_Policy_Result01_InternetExplorer02) GetPropertyTrustedSitesZoneNavigateWindowsAndFrames() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneNavigateWindowsAndFrames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
