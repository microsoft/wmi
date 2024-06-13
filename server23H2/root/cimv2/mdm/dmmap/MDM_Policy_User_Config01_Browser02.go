// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
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

// MDM_Policy_User_Config01_Browser02 struct
type MDM_Policy_User_Config01_Browser02 struct {
	*cim.WmiInstance

	//
	AllowAddressBarDropdown int32

	//
	AllowAutofill int32

	//
	AllowConfigurationUpdateForBooksLibrary int32

	//
	AllowCookies int32

	//
	AllowDeveloperTools int32

	//
	AllowDoNotTrack int32

	//
	AllowExtensions int32

	//
	AllowFlash int32

	//
	AllowFlashClickToRun int32

	//
	AllowFullScreenMode int32

	//
	AllowInPrivate int32

	//
	AllowMicrosoftCompatibilityList int32

	//
	AllowPasswordManager int32

	//
	AllowPopups int32

	//
	AllowPrelaunch int32

	//
	AllowPrinting int32

	//
	AllowSavingHistory int32

	//
	AllowSearchEngineCustomization int32

	//
	AllowSearchSuggestionsinAddressBar int32

	//
	AllowSideloadingOfExtensions int32

	//
	AllowSmartScreen int32

	//
	AllowTabPreloading int32

	//
	AllowWebContentOnNewTabPage int32

	//
	AlwaysEnableBooksLibrary int32

	//
	ClearBrowsingDataOnExit int32

	//
	ConfigureAdditionalSearchEngines string

	//
	ConfigureFavoritesBar int32

	//
	ConfigureHomeButton int32

	//
	ConfigureKioskMode int32

	//
	ConfigureKioskResetAfterIdleTimeout int32

	//
	ConfigureOpenMicrosoftEdgeWith int32

	//
	ConfigureTelemetryForMicrosoft365Analytics int32

	//
	DisableLockdownOfStartPages int32

	//
	EnableExtendedBooksTelemetry int32

	//
	EnterpriseModeSiteList string

	//
	EnterpriseSiteListServiceUrl string

	//
	HomePages string

	//
	InstanceID string

	//
	LockdownFavorites int32

	//
	ParentID string

	//
	PreventAccessToAboutFlagsInMicrosoftEdge int32

	//
	PreventCertErrorOverrides int32

	//
	PreventFirstRunPage int32

	//
	PreventLiveTileDataCollection int32

	//
	PreventSmartScreenPromptOverride int32

	//
	PreventSmartScreenPromptOverrideForFiles int32

	//
	PreventTurningOffRequiredExtensions string

	//
	PreventUsingLocalHostIPAddressForWebRTC int32

	//
	ProvisionFavorites string

	//
	SendIntranetTraffictoInternetExplorer int32

	//
	SetDefaultSearchEngine string

	//
	SetHomeButtonURL string

	//
	SetNewTabPageURL string

	//
	ShowMessageWhenOpeningSitesInInternetExplorer int32

	//
	SyncFavoritesBetweenIEAndMicrosoftEdge int32

	//
	UnlockHomeButton int32

	//
	UseSharedFolderForBooks int32
}

func NewMDM_Policy_User_Config01_Browser02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_User_Config01_Browser02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_User_Config01_Browser02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_User_Config01_Browser02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_User_Config01_Browser02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_User_Config01_Browser02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowAddressBarDropdown sets the value of AllowAddressBarDropdown for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowAddressBarDropdown(value int32) (err error) {
	return instance.SetProperty("AllowAddressBarDropdown", (value))
}

// GetAllowAddressBarDropdown gets the value of AllowAddressBarDropdown for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowAddressBarDropdown() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAddressBarDropdown")
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

// SetAllowAutofill sets the value of AllowAutofill for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowAutofill(value int32) (err error) {
	return instance.SetProperty("AllowAutofill", (value))
}

// GetAllowAutofill gets the value of AllowAutofill for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowAutofill() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAutofill")
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

// SetAllowConfigurationUpdateForBooksLibrary sets the value of AllowConfigurationUpdateForBooksLibrary for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowConfigurationUpdateForBooksLibrary(value int32) (err error) {
	return instance.SetProperty("AllowConfigurationUpdateForBooksLibrary", (value))
}

// GetAllowConfigurationUpdateForBooksLibrary gets the value of AllowConfigurationUpdateForBooksLibrary for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowConfigurationUpdateForBooksLibrary() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowConfigurationUpdateForBooksLibrary")
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

// SetAllowCookies sets the value of AllowCookies for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowCookies(value int32) (err error) {
	return instance.SetProperty("AllowCookies", (value))
}

// GetAllowCookies gets the value of AllowCookies for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowCookies() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowCookies")
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

// SetAllowDeveloperTools sets the value of AllowDeveloperTools for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowDeveloperTools(value int32) (err error) {
	return instance.SetProperty("AllowDeveloperTools", (value))
}

// GetAllowDeveloperTools gets the value of AllowDeveloperTools for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowDeveloperTools() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowDeveloperTools")
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

// SetAllowDoNotTrack sets the value of AllowDoNotTrack for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowDoNotTrack(value int32) (err error) {
	return instance.SetProperty("AllowDoNotTrack", (value))
}

// GetAllowDoNotTrack gets the value of AllowDoNotTrack for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowDoNotTrack() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowDoNotTrack")
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

// SetAllowExtensions sets the value of AllowExtensions for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowExtensions(value int32) (err error) {
	return instance.SetProperty("AllowExtensions", (value))
}

// GetAllowExtensions gets the value of AllowExtensions for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowExtensions() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowExtensions")
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

// SetAllowFlash sets the value of AllowFlash for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowFlash(value int32) (err error) {
	return instance.SetProperty("AllowFlash", (value))
}

// GetAllowFlash gets the value of AllowFlash for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowFlash() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowFlash")
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

// SetAllowFlashClickToRun sets the value of AllowFlashClickToRun for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowFlashClickToRun(value int32) (err error) {
	return instance.SetProperty("AllowFlashClickToRun", (value))
}

// GetAllowFlashClickToRun gets the value of AllowFlashClickToRun for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowFlashClickToRun() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowFlashClickToRun")
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

// SetAllowFullScreenMode sets the value of AllowFullScreenMode for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowFullScreenMode(value int32) (err error) {
	return instance.SetProperty("AllowFullScreenMode", (value))
}

// GetAllowFullScreenMode gets the value of AllowFullScreenMode for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowFullScreenMode() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowFullScreenMode")
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

// SetAllowInPrivate sets the value of AllowInPrivate for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowInPrivate(value int32) (err error) {
	return instance.SetProperty("AllowInPrivate", (value))
}

// GetAllowInPrivate gets the value of AllowInPrivate for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowInPrivate() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowInPrivate")
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

// SetAllowMicrosoftCompatibilityList sets the value of AllowMicrosoftCompatibilityList for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowMicrosoftCompatibilityList(value int32) (err error) {
	return instance.SetProperty("AllowMicrosoftCompatibilityList", (value))
}

// GetAllowMicrosoftCompatibilityList gets the value of AllowMicrosoftCompatibilityList for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowMicrosoftCompatibilityList() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowMicrosoftCompatibilityList")
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

// SetAllowPasswordManager sets the value of AllowPasswordManager for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowPasswordManager(value int32) (err error) {
	return instance.SetProperty("AllowPasswordManager", (value))
}

// GetAllowPasswordManager gets the value of AllowPasswordManager for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowPasswordManager() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowPasswordManager")
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

// SetAllowPopups sets the value of AllowPopups for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowPopups(value int32) (err error) {
	return instance.SetProperty("AllowPopups", (value))
}

// GetAllowPopups gets the value of AllowPopups for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowPopups() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowPopups")
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

// SetAllowPrelaunch sets the value of AllowPrelaunch for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowPrelaunch(value int32) (err error) {
	return instance.SetProperty("AllowPrelaunch", (value))
}

// GetAllowPrelaunch gets the value of AllowPrelaunch for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowPrelaunch() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowPrelaunch")
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

// SetAllowPrinting sets the value of AllowPrinting for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowPrinting(value int32) (err error) {
	return instance.SetProperty("AllowPrinting", (value))
}

// GetAllowPrinting gets the value of AllowPrinting for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowPrinting() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowPrinting")
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

// SetAllowSavingHistory sets the value of AllowSavingHistory for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowSavingHistory(value int32) (err error) {
	return instance.SetProperty("AllowSavingHistory", (value))
}

// GetAllowSavingHistory gets the value of AllowSavingHistory for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowSavingHistory() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSavingHistory")
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

// SetAllowSearchEngineCustomization sets the value of AllowSearchEngineCustomization for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowSearchEngineCustomization(value int32) (err error) {
	return instance.SetProperty("AllowSearchEngineCustomization", (value))
}

// GetAllowSearchEngineCustomization gets the value of AllowSearchEngineCustomization for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowSearchEngineCustomization() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSearchEngineCustomization")
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

// SetAllowSearchSuggestionsinAddressBar sets the value of AllowSearchSuggestionsinAddressBar for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowSearchSuggestionsinAddressBar(value int32) (err error) {
	return instance.SetProperty("AllowSearchSuggestionsinAddressBar", (value))
}

// GetAllowSearchSuggestionsinAddressBar gets the value of AllowSearchSuggestionsinAddressBar for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowSearchSuggestionsinAddressBar() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSearchSuggestionsinAddressBar")
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

// SetAllowSideloadingOfExtensions sets the value of AllowSideloadingOfExtensions for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowSideloadingOfExtensions(value int32) (err error) {
	return instance.SetProperty("AllowSideloadingOfExtensions", (value))
}

// GetAllowSideloadingOfExtensions gets the value of AllowSideloadingOfExtensions for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowSideloadingOfExtensions() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSideloadingOfExtensions")
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

// SetAllowSmartScreen sets the value of AllowSmartScreen for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowSmartScreen(value int32) (err error) {
	return instance.SetProperty("AllowSmartScreen", (value))
}

// GetAllowSmartScreen gets the value of AllowSmartScreen for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowSmartScreen() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSmartScreen")
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

// SetAllowTabPreloading sets the value of AllowTabPreloading for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowTabPreloading(value int32) (err error) {
	return instance.SetProperty("AllowTabPreloading", (value))
}

// GetAllowTabPreloading gets the value of AllowTabPreloading for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowTabPreloading() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowTabPreloading")
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

// SetAllowWebContentOnNewTabPage sets the value of AllowWebContentOnNewTabPage for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAllowWebContentOnNewTabPage(value int32) (err error) {
	return instance.SetProperty("AllowWebContentOnNewTabPage", (value))
}

// GetAllowWebContentOnNewTabPage gets the value of AllowWebContentOnNewTabPage for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAllowWebContentOnNewTabPage() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWebContentOnNewTabPage")
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

// SetAlwaysEnableBooksLibrary sets the value of AlwaysEnableBooksLibrary for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyAlwaysEnableBooksLibrary(value int32) (err error) {
	return instance.SetProperty("AlwaysEnableBooksLibrary", (value))
}

// GetAlwaysEnableBooksLibrary gets the value of AlwaysEnableBooksLibrary for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyAlwaysEnableBooksLibrary() (value int32, err error) {
	retValue, err := instance.GetProperty("AlwaysEnableBooksLibrary")
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

// SetClearBrowsingDataOnExit sets the value of ClearBrowsingDataOnExit for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyClearBrowsingDataOnExit(value int32) (err error) {
	return instance.SetProperty("ClearBrowsingDataOnExit", (value))
}

// GetClearBrowsingDataOnExit gets the value of ClearBrowsingDataOnExit for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyClearBrowsingDataOnExit() (value int32, err error) {
	retValue, err := instance.GetProperty("ClearBrowsingDataOnExit")
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

// SetConfigureAdditionalSearchEngines sets the value of ConfigureAdditionalSearchEngines for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyConfigureAdditionalSearchEngines(value string) (err error) {
	return instance.SetProperty("ConfigureAdditionalSearchEngines", (value))
}

// GetConfigureAdditionalSearchEngines gets the value of ConfigureAdditionalSearchEngines for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyConfigureAdditionalSearchEngines() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigureAdditionalSearchEngines")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetConfigureFavoritesBar sets the value of ConfigureFavoritesBar for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyConfigureFavoritesBar(value int32) (err error) {
	return instance.SetProperty("ConfigureFavoritesBar", (value))
}

// GetConfigureFavoritesBar gets the value of ConfigureFavoritesBar for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyConfigureFavoritesBar() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureFavoritesBar")
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

// SetConfigureHomeButton sets the value of ConfigureHomeButton for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyConfigureHomeButton(value int32) (err error) {
	return instance.SetProperty("ConfigureHomeButton", (value))
}

// GetConfigureHomeButton gets the value of ConfigureHomeButton for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyConfigureHomeButton() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureHomeButton")
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

// SetConfigureKioskMode sets the value of ConfigureKioskMode for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyConfigureKioskMode(value int32) (err error) {
	return instance.SetProperty("ConfigureKioskMode", (value))
}

// GetConfigureKioskMode gets the value of ConfigureKioskMode for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyConfigureKioskMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureKioskMode")
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

// SetConfigureKioskResetAfterIdleTimeout sets the value of ConfigureKioskResetAfterIdleTimeout for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyConfigureKioskResetAfterIdleTimeout(value int32) (err error) {
	return instance.SetProperty("ConfigureKioskResetAfterIdleTimeout", (value))
}

// GetConfigureKioskResetAfterIdleTimeout gets the value of ConfigureKioskResetAfterIdleTimeout for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyConfigureKioskResetAfterIdleTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureKioskResetAfterIdleTimeout")
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

// SetConfigureOpenMicrosoftEdgeWith sets the value of ConfigureOpenMicrosoftEdgeWith for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyConfigureOpenMicrosoftEdgeWith(value int32) (err error) {
	return instance.SetProperty("ConfigureOpenMicrosoftEdgeWith", (value))
}

// GetConfigureOpenMicrosoftEdgeWith gets the value of ConfigureOpenMicrosoftEdgeWith for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyConfigureOpenMicrosoftEdgeWith() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureOpenMicrosoftEdgeWith")
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

// SetConfigureTelemetryForMicrosoft365Analytics sets the value of ConfigureTelemetryForMicrosoft365Analytics for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyConfigureTelemetryForMicrosoft365Analytics(value int32) (err error) {
	return instance.SetProperty("ConfigureTelemetryForMicrosoft365Analytics", (value))
}

// GetConfigureTelemetryForMicrosoft365Analytics gets the value of ConfigureTelemetryForMicrosoft365Analytics for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyConfigureTelemetryForMicrosoft365Analytics() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureTelemetryForMicrosoft365Analytics")
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

// SetDisableLockdownOfStartPages sets the value of DisableLockdownOfStartPages for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyDisableLockdownOfStartPages(value int32) (err error) {
	return instance.SetProperty("DisableLockdownOfStartPages", (value))
}

// GetDisableLockdownOfStartPages gets the value of DisableLockdownOfStartPages for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyDisableLockdownOfStartPages() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableLockdownOfStartPages")
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

// SetEnableExtendedBooksTelemetry sets the value of EnableExtendedBooksTelemetry for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyEnableExtendedBooksTelemetry(value int32) (err error) {
	return instance.SetProperty("EnableExtendedBooksTelemetry", (value))
}

// GetEnableExtendedBooksTelemetry gets the value of EnableExtendedBooksTelemetry for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyEnableExtendedBooksTelemetry() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableExtendedBooksTelemetry")
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

// SetEnterpriseModeSiteList sets the value of EnterpriseModeSiteList for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyEnterpriseModeSiteList(value string) (err error) {
	return instance.SetProperty("EnterpriseModeSiteList", (value))
}

// GetEnterpriseModeSiteList gets the value of EnterpriseModeSiteList for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyEnterpriseModeSiteList() (value string, err error) {
	retValue, err := instance.GetProperty("EnterpriseModeSiteList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetEnterpriseSiteListServiceUrl sets the value of EnterpriseSiteListServiceUrl for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyEnterpriseSiteListServiceUrl(value string) (err error) {
	return instance.SetProperty("EnterpriseSiteListServiceUrl", (value))
}

// GetEnterpriseSiteListServiceUrl gets the value of EnterpriseSiteListServiceUrl for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyEnterpriseSiteListServiceUrl() (value string, err error) {
	retValue, err := instance.GetProperty("EnterpriseSiteListServiceUrl")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetHomePages sets the value of HomePages for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyHomePages(value string) (err error) {
	return instance.SetProperty("HomePages", (value))
}

// GetHomePages gets the value of HomePages for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyHomePages() (value string, err error) {
	retValue, err := instance.GetProperty("HomePages")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
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
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyInstanceID() (value string, err error) {
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

// SetLockdownFavorites sets the value of LockdownFavorites for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyLockdownFavorites(value int32) (err error) {
	return instance.SetProperty("LockdownFavorites", (value))
}

// GetLockdownFavorites gets the value of LockdownFavorites for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyLockdownFavorites() (value int32, err error) {
	retValue, err := instance.GetProperty("LockdownFavorites")
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
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyParentID() (value string, err error) {
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

// SetPreventAccessToAboutFlagsInMicrosoftEdge sets the value of PreventAccessToAboutFlagsInMicrosoftEdge for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyPreventAccessToAboutFlagsInMicrosoftEdge(value int32) (err error) {
	return instance.SetProperty("PreventAccessToAboutFlagsInMicrosoftEdge", (value))
}

// GetPreventAccessToAboutFlagsInMicrosoftEdge gets the value of PreventAccessToAboutFlagsInMicrosoftEdge for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyPreventAccessToAboutFlagsInMicrosoftEdge() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventAccessToAboutFlagsInMicrosoftEdge")
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

// SetPreventCertErrorOverrides sets the value of PreventCertErrorOverrides for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyPreventCertErrorOverrides(value int32) (err error) {
	return instance.SetProperty("PreventCertErrorOverrides", (value))
}

// GetPreventCertErrorOverrides gets the value of PreventCertErrorOverrides for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyPreventCertErrorOverrides() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventCertErrorOverrides")
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

// SetPreventFirstRunPage sets the value of PreventFirstRunPage for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyPreventFirstRunPage(value int32) (err error) {
	return instance.SetProperty("PreventFirstRunPage", (value))
}

// GetPreventFirstRunPage gets the value of PreventFirstRunPage for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyPreventFirstRunPage() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventFirstRunPage")
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

// SetPreventLiveTileDataCollection sets the value of PreventLiveTileDataCollection for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyPreventLiveTileDataCollection(value int32) (err error) {
	return instance.SetProperty("PreventLiveTileDataCollection", (value))
}

// GetPreventLiveTileDataCollection gets the value of PreventLiveTileDataCollection for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyPreventLiveTileDataCollection() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventLiveTileDataCollection")
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

// SetPreventSmartScreenPromptOverride sets the value of PreventSmartScreenPromptOverride for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyPreventSmartScreenPromptOverride(value int32) (err error) {
	return instance.SetProperty("PreventSmartScreenPromptOverride", (value))
}

// GetPreventSmartScreenPromptOverride gets the value of PreventSmartScreenPromptOverride for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyPreventSmartScreenPromptOverride() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventSmartScreenPromptOverride")
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

// SetPreventSmartScreenPromptOverrideForFiles sets the value of PreventSmartScreenPromptOverrideForFiles for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyPreventSmartScreenPromptOverrideForFiles(value int32) (err error) {
	return instance.SetProperty("PreventSmartScreenPromptOverrideForFiles", (value))
}

// GetPreventSmartScreenPromptOverrideForFiles gets the value of PreventSmartScreenPromptOverrideForFiles for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyPreventSmartScreenPromptOverrideForFiles() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventSmartScreenPromptOverrideForFiles")
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

// SetPreventTurningOffRequiredExtensions sets the value of PreventTurningOffRequiredExtensions for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyPreventTurningOffRequiredExtensions(value string) (err error) {
	return instance.SetProperty("PreventTurningOffRequiredExtensions", (value))
}

// GetPreventTurningOffRequiredExtensions gets the value of PreventTurningOffRequiredExtensions for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyPreventTurningOffRequiredExtensions() (value string, err error) {
	retValue, err := instance.GetProperty("PreventTurningOffRequiredExtensions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetPreventUsingLocalHostIPAddressForWebRTC sets the value of PreventUsingLocalHostIPAddressForWebRTC for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyPreventUsingLocalHostIPAddressForWebRTC(value int32) (err error) {
	return instance.SetProperty("PreventUsingLocalHostIPAddressForWebRTC", (value))
}

// GetPreventUsingLocalHostIPAddressForWebRTC gets the value of PreventUsingLocalHostIPAddressForWebRTC for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyPreventUsingLocalHostIPAddressForWebRTC() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventUsingLocalHostIPAddressForWebRTC")
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

// SetProvisionFavorites sets the value of ProvisionFavorites for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyProvisionFavorites(value string) (err error) {
	return instance.SetProperty("ProvisionFavorites", (value))
}

// GetProvisionFavorites gets the value of ProvisionFavorites for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyProvisionFavorites() (value string, err error) {
	retValue, err := instance.GetProperty("ProvisionFavorites")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSendIntranetTraffictoInternetExplorer sets the value of SendIntranetTraffictoInternetExplorer for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertySendIntranetTraffictoInternetExplorer(value int32) (err error) {
	return instance.SetProperty("SendIntranetTraffictoInternetExplorer", (value))
}

// GetSendIntranetTraffictoInternetExplorer gets the value of SendIntranetTraffictoInternetExplorer for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertySendIntranetTraffictoInternetExplorer() (value int32, err error) {
	retValue, err := instance.GetProperty("SendIntranetTraffictoInternetExplorer")
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

// SetSetDefaultSearchEngine sets the value of SetDefaultSearchEngine for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertySetDefaultSearchEngine(value string) (err error) {
	return instance.SetProperty("SetDefaultSearchEngine", (value))
}

// GetSetDefaultSearchEngine gets the value of SetDefaultSearchEngine for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertySetDefaultSearchEngine() (value string, err error) {
	retValue, err := instance.GetProperty("SetDefaultSearchEngine")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSetHomeButtonURL sets the value of SetHomeButtonURL for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertySetHomeButtonURL(value string) (err error) {
	return instance.SetProperty("SetHomeButtonURL", (value))
}

// GetSetHomeButtonURL gets the value of SetHomeButtonURL for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertySetHomeButtonURL() (value string, err error) {
	retValue, err := instance.GetProperty("SetHomeButtonURL")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSetNewTabPageURL sets the value of SetNewTabPageURL for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertySetNewTabPageURL(value string) (err error) {
	return instance.SetProperty("SetNewTabPageURL", (value))
}

// GetSetNewTabPageURL gets the value of SetNewTabPageURL for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertySetNewTabPageURL() (value string, err error) {
	retValue, err := instance.GetProperty("SetNewTabPageURL")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetShowMessageWhenOpeningSitesInInternetExplorer sets the value of ShowMessageWhenOpeningSitesInInternetExplorer for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyShowMessageWhenOpeningSitesInInternetExplorer(value int32) (err error) {
	return instance.SetProperty("ShowMessageWhenOpeningSitesInInternetExplorer", (value))
}

// GetShowMessageWhenOpeningSitesInInternetExplorer gets the value of ShowMessageWhenOpeningSitesInInternetExplorer for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyShowMessageWhenOpeningSitesInInternetExplorer() (value int32, err error) {
	retValue, err := instance.GetProperty("ShowMessageWhenOpeningSitesInInternetExplorer")
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

// SetSyncFavoritesBetweenIEAndMicrosoftEdge sets the value of SyncFavoritesBetweenIEAndMicrosoftEdge for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertySyncFavoritesBetweenIEAndMicrosoftEdge(value int32) (err error) {
	return instance.SetProperty("SyncFavoritesBetweenIEAndMicrosoftEdge", (value))
}

// GetSyncFavoritesBetweenIEAndMicrosoftEdge gets the value of SyncFavoritesBetweenIEAndMicrosoftEdge for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertySyncFavoritesBetweenIEAndMicrosoftEdge() (value int32, err error) {
	retValue, err := instance.GetProperty("SyncFavoritesBetweenIEAndMicrosoftEdge")
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

// SetUnlockHomeButton sets the value of UnlockHomeButton for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyUnlockHomeButton(value int32) (err error) {
	return instance.SetProperty("UnlockHomeButton", (value))
}

// GetUnlockHomeButton gets the value of UnlockHomeButton for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyUnlockHomeButton() (value int32, err error) {
	retValue, err := instance.GetProperty("UnlockHomeButton")
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

// SetUseSharedFolderForBooks sets the value of UseSharedFolderForBooks for the instance
func (instance *MDM_Policy_User_Config01_Browser02) SetPropertyUseSharedFolderForBooks(value int32) (err error) {
	return instance.SetProperty("UseSharedFolderForBooks", (value))
}

// GetUseSharedFolderForBooks gets the value of UseSharedFolderForBooks for the instance
func (instance *MDM_Policy_User_Config01_Browser02) GetPropertyUseSharedFolderForBooks() (value int32, err error) {
	retValue, err := instance.GetProperty("UseSharedFolderForBooks")
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
