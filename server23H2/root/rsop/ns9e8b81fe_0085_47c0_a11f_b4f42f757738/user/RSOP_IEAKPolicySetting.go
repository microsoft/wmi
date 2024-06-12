// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS9E8B81FE_0085_47C0_A11F_B4F42F757738.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEAKPolicySetting struct
type RSOP_IEAKPolicySetting struct {
	*RSOP_PolicySetting

	//
	categories int32

	//
	channels int32

	//
	customFavorites int32

	//
	customizeAnimatedBitmaps bool

	//
	customizeLogoBitmaps bool

	//
	customLinks int32

	//
	deleteAdminCreatedFavoritesOnly bool

	//
	deleteExistingChannels bool

	//
	deleteExistingFavorites bool

	//
	deleteExistingToolbarButtons bool

	//
	enableDesktopChannelBarByDefault bool

	//
	enableTrustedPublisherLockdown bool

	//
	homePageURL string

	//
	importAuthenticodeSecurityInfo bool

	//
	importContentRatingsSettings bool

	//
	importedZoneCount uint32

	//
	importProgramSettings bool

	//
	importSecurityZoneSettings bool

	//
	largeAnimatedBitmapName string

	//
	largeAnimatedBitmapPath string

	//
	largeCustomLogoBitmapName string

	//
	largeCustomLogoBitmapPath string

	//
	onlineHelpPageURL string

	//
	placeFavoritesAtTopOfList bool

	//
	preferenceMode bool

	//
	searchBarURL string

	//
	smallAnimatedBitmapName string

	//
	smallAnimatedBitmapPath string

	//
	smallCustomLogoBitmapName string

	//
	smallCustomLogoBitmapPath string

	//
	titleBarCustomText string

	//
	titleBarText string

	//
	toolbarBackgroundBitmapPath string

	//
	toolbarButtons int32

	//
	userAgentText string
}

func NewRSOP_IEAKPolicySettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEAKPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEAKPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_IEAKPolicySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEAKPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEAKPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

// Setcategories sets the value of categories for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertycategories(value int32) (err error) {
	return instance.SetProperty("categories", (value))
}

// Getcategories gets the value of categories for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertycategories() (value int32, err error) {
	retValue, err := instance.GetProperty("categories")
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

// Setchannels sets the value of channels for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertychannels(value int32) (err error) {
	return instance.SetProperty("channels", (value))
}

// Getchannels gets the value of channels for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertychannels() (value int32, err error) {
	retValue, err := instance.GetProperty("channels")
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

// SetcustomFavorites sets the value of customFavorites for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertycustomFavorites(value int32) (err error) {
	return instance.SetProperty("customFavorites", (value))
}

// GetcustomFavorites gets the value of customFavorites for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertycustomFavorites() (value int32, err error) {
	retValue, err := instance.GetProperty("customFavorites")
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

// SetcustomizeAnimatedBitmaps sets the value of customizeAnimatedBitmaps for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertycustomizeAnimatedBitmaps(value bool) (err error) {
	return instance.SetProperty("customizeAnimatedBitmaps", (value))
}

// GetcustomizeAnimatedBitmaps gets the value of customizeAnimatedBitmaps for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertycustomizeAnimatedBitmaps() (value bool, err error) {
	retValue, err := instance.GetProperty("customizeAnimatedBitmaps")
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

// SetcustomizeLogoBitmaps sets the value of customizeLogoBitmaps for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertycustomizeLogoBitmaps(value bool) (err error) {
	return instance.SetProperty("customizeLogoBitmaps", (value))
}

// GetcustomizeLogoBitmaps gets the value of customizeLogoBitmaps for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertycustomizeLogoBitmaps() (value bool, err error) {
	retValue, err := instance.GetProperty("customizeLogoBitmaps")
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

// SetcustomLinks sets the value of customLinks for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertycustomLinks(value int32) (err error) {
	return instance.SetProperty("customLinks", (value))
}

// GetcustomLinks gets the value of customLinks for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertycustomLinks() (value int32, err error) {
	retValue, err := instance.GetProperty("customLinks")
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

// SetdeleteAdminCreatedFavoritesOnly sets the value of deleteAdminCreatedFavoritesOnly for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertydeleteAdminCreatedFavoritesOnly(value bool) (err error) {
	return instance.SetProperty("deleteAdminCreatedFavoritesOnly", (value))
}

// GetdeleteAdminCreatedFavoritesOnly gets the value of deleteAdminCreatedFavoritesOnly for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertydeleteAdminCreatedFavoritesOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("deleteAdminCreatedFavoritesOnly")
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

// SetdeleteExistingChannels sets the value of deleteExistingChannels for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertydeleteExistingChannels(value bool) (err error) {
	return instance.SetProperty("deleteExistingChannels", (value))
}

// GetdeleteExistingChannels gets the value of deleteExistingChannels for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertydeleteExistingChannels() (value bool, err error) {
	retValue, err := instance.GetProperty("deleteExistingChannels")
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

// SetdeleteExistingFavorites sets the value of deleteExistingFavorites for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertydeleteExistingFavorites(value bool) (err error) {
	return instance.SetProperty("deleteExistingFavorites", (value))
}

// GetdeleteExistingFavorites gets the value of deleteExistingFavorites for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertydeleteExistingFavorites() (value bool, err error) {
	retValue, err := instance.GetProperty("deleteExistingFavorites")
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

// SetdeleteExistingToolbarButtons sets the value of deleteExistingToolbarButtons for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertydeleteExistingToolbarButtons(value bool) (err error) {
	return instance.SetProperty("deleteExistingToolbarButtons", (value))
}

// GetdeleteExistingToolbarButtons gets the value of deleteExistingToolbarButtons for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertydeleteExistingToolbarButtons() (value bool, err error) {
	retValue, err := instance.GetProperty("deleteExistingToolbarButtons")
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

// SetenableDesktopChannelBarByDefault sets the value of enableDesktopChannelBarByDefault for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertyenableDesktopChannelBarByDefault(value bool) (err error) {
	return instance.SetProperty("enableDesktopChannelBarByDefault", (value))
}

// GetenableDesktopChannelBarByDefault gets the value of enableDesktopChannelBarByDefault for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertyenableDesktopChannelBarByDefault() (value bool, err error) {
	retValue, err := instance.GetProperty("enableDesktopChannelBarByDefault")
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

// SetenableTrustedPublisherLockdown sets the value of enableTrustedPublisherLockdown for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertyenableTrustedPublisherLockdown(value bool) (err error) {
	return instance.SetProperty("enableTrustedPublisherLockdown", (value))
}

// GetenableTrustedPublisherLockdown gets the value of enableTrustedPublisherLockdown for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertyenableTrustedPublisherLockdown() (value bool, err error) {
	retValue, err := instance.GetProperty("enableTrustedPublisherLockdown")
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

// SethomePageURL sets the value of homePageURL for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertyhomePageURL(value string) (err error) {
	return instance.SetProperty("homePageURL", (value))
}

// GethomePageURL gets the value of homePageURL for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertyhomePageURL() (value string, err error) {
	retValue, err := instance.GetProperty("homePageURL")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetimportAuthenticodeSecurityInfo sets the value of importAuthenticodeSecurityInfo for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertyimportAuthenticodeSecurityInfo(value bool) (err error) {
	return instance.SetProperty("importAuthenticodeSecurityInfo", (value))
}

// GetimportAuthenticodeSecurityInfo gets the value of importAuthenticodeSecurityInfo for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertyimportAuthenticodeSecurityInfo() (value bool, err error) {
	retValue, err := instance.GetProperty("importAuthenticodeSecurityInfo")
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

// SetimportContentRatingsSettings sets the value of importContentRatingsSettings for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertyimportContentRatingsSettings(value bool) (err error) {
	return instance.SetProperty("importContentRatingsSettings", (value))
}

// GetimportContentRatingsSettings gets the value of importContentRatingsSettings for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertyimportContentRatingsSettings() (value bool, err error) {
	retValue, err := instance.GetProperty("importContentRatingsSettings")
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

// SetimportedZoneCount sets the value of importedZoneCount for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertyimportedZoneCount(value uint32) (err error) {
	return instance.SetProperty("importedZoneCount", (value))
}

// GetimportedZoneCount gets the value of importedZoneCount for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertyimportedZoneCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("importedZoneCount")
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

// SetimportProgramSettings sets the value of importProgramSettings for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertyimportProgramSettings(value bool) (err error) {
	return instance.SetProperty("importProgramSettings", (value))
}

// GetimportProgramSettings gets the value of importProgramSettings for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertyimportProgramSettings() (value bool, err error) {
	retValue, err := instance.GetProperty("importProgramSettings")
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

// SetimportSecurityZoneSettings sets the value of importSecurityZoneSettings for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertyimportSecurityZoneSettings(value bool) (err error) {
	return instance.SetProperty("importSecurityZoneSettings", (value))
}

// GetimportSecurityZoneSettings gets the value of importSecurityZoneSettings for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertyimportSecurityZoneSettings() (value bool, err error) {
	retValue, err := instance.GetProperty("importSecurityZoneSettings")
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

// SetlargeAnimatedBitmapName sets the value of largeAnimatedBitmapName for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertylargeAnimatedBitmapName(value string) (err error) {
	return instance.SetProperty("largeAnimatedBitmapName", (value))
}

// GetlargeAnimatedBitmapName gets the value of largeAnimatedBitmapName for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertylargeAnimatedBitmapName() (value string, err error) {
	retValue, err := instance.GetProperty("largeAnimatedBitmapName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetlargeAnimatedBitmapPath sets the value of largeAnimatedBitmapPath for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertylargeAnimatedBitmapPath(value string) (err error) {
	return instance.SetProperty("largeAnimatedBitmapPath", (value))
}

// GetlargeAnimatedBitmapPath gets the value of largeAnimatedBitmapPath for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertylargeAnimatedBitmapPath() (value string, err error) {
	retValue, err := instance.GetProperty("largeAnimatedBitmapPath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetlargeCustomLogoBitmapName sets the value of largeCustomLogoBitmapName for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertylargeCustomLogoBitmapName(value string) (err error) {
	return instance.SetProperty("largeCustomLogoBitmapName", (value))
}

// GetlargeCustomLogoBitmapName gets the value of largeCustomLogoBitmapName for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertylargeCustomLogoBitmapName() (value string, err error) {
	retValue, err := instance.GetProperty("largeCustomLogoBitmapName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetlargeCustomLogoBitmapPath sets the value of largeCustomLogoBitmapPath for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertylargeCustomLogoBitmapPath(value string) (err error) {
	return instance.SetProperty("largeCustomLogoBitmapPath", (value))
}

// GetlargeCustomLogoBitmapPath gets the value of largeCustomLogoBitmapPath for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertylargeCustomLogoBitmapPath() (value string, err error) {
	retValue, err := instance.GetProperty("largeCustomLogoBitmapPath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetonlineHelpPageURL sets the value of onlineHelpPageURL for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertyonlineHelpPageURL(value string) (err error) {
	return instance.SetProperty("onlineHelpPageURL", (value))
}

// GetonlineHelpPageURL gets the value of onlineHelpPageURL for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertyonlineHelpPageURL() (value string, err error) {
	retValue, err := instance.GetProperty("onlineHelpPageURL")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetplaceFavoritesAtTopOfList sets the value of placeFavoritesAtTopOfList for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertyplaceFavoritesAtTopOfList(value bool) (err error) {
	return instance.SetProperty("placeFavoritesAtTopOfList", (value))
}

// GetplaceFavoritesAtTopOfList gets the value of placeFavoritesAtTopOfList for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertyplaceFavoritesAtTopOfList() (value bool, err error) {
	retValue, err := instance.GetProperty("placeFavoritesAtTopOfList")
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

// SetpreferenceMode sets the value of preferenceMode for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertypreferenceMode(value bool) (err error) {
	return instance.SetProperty("preferenceMode", (value))
}

// GetpreferenceMode gets the value of preferenceMode for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertypreferenceMode() (value bool, err error) {
	retValue, err := instance.GetProperty("preferenceMode")
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

// SetsearchBarURL sets the value of searchBarURL for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertysearchBarURL(value string) (err error) {
	return instance.SetProperty("searchBarURL", (value))
}

// GetsearchBarURL gets the value of searchBarURL for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertysearchBarURL() (value string, err error) {
	retValue, err := instance.GetProperty("searchBarURL")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetsmallAnimatedBitmapName sets the value of smallAnimatedBitmapName for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertysmallAnimatedBitmapName(value string) (err error) {
	return instance.SetProperty("smallAnimatedBitmapName", (value))
}

// GetsmallAnimatedBitmapName gets the value of smallAnimatedBitmapName for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertysmallAnimatedBitmapName() (value string, err error) {
	retValue, err := instance.GetProperty("smallAnimatedBitmapName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetsmallAnimatedBitmapPath sets the value of smallAnimatedBitmapPath for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertysmallAnimatedBitmapPath(value string) (err error) {
	return instance.SetProperty("smallAnimatedBitmapPath", (value))
}

// GetsmallAnimatedBitmapPath gets the value of smallAnimatedBitmapPath for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertysmallAnimatedBitmapPath() (value string, err error) {
	retValue, err := instance.GetProperty("smallAnimatedBitmapPath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetsmallCustomLogoBitmapName sets the value of smallCustomLogoBitmapName for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertysmallCustomLogoBitmapName(value string) (err error) {
	return instance.SetProperty("smallCustomLogoBitmapName", (value))
}

// GetsmallCustomLogoBitmapName gets the value of smallCustomLogoBitmapName for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertysmallCustomLogoBitmapName() (value string, err error) {
	retValue, err := instance.GetProperty("smallCustomLogoBitmapName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetsmallCustomLogoBitmapPath sets the value of smallCustomLogoBitmapPath for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertysmallCustomLogoBitmapPath(value string) (err error) {
	return instance.SetProperty("smallCustomLogoBitmapPath", (value))
}

// GetsmallCustomLogoBitmapPath gets the value of smallCustomLogoBitmapPath for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertysmallCustomLogoBitmapPath() (value string, err error) {
	retValue, err := instance.GetProperty("smallCustomLogoBitmapPath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SettitleBarCustomText sets the value of titleBarCustomText for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertytitleBarCustomText(value string) (err error) {
	return instance.SetProperty("titleBarCustomText", (value))
}

// GettitleBarCustomText gets the value of titleBarCustomText for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertytitleBarCustomText() (value string, err error) {
	retValue, err := instance.GetProperty("titleBarCustomText")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SettitleBarText sets the value of titleBarText for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertytitleBarText(value string) (err error) {
	return instance.SetProperty("titleBarText", (value))
}

// GettitleBarText gets the value of titleBarText for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertytitleBarText() (value string, err error) {
	retValue, err := instance.GetProperty("titleBarText")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SettoolbarBackgroundBitmapPath sets the value of toolbarBackgroundBitmapPath for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertytoolbarBackgroundBitmapPath(value string) (err error) {
	return instance.SetProperty("toolbarBackgroundBitmapPath", (value))
}

// GettoolbarBackgroundBitmapPath gets the value of toolbarBackgroundBitmapPath for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertytoolbarBackgroundBitmapPath() (value string, err error) {
	retValue, err := instance.GetProperty("toolbarBackgroundBitmapPath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SettoolbarButtons sets the value of toolbarButtons for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertytoolbarButtons(value int32) (err error) {
	return instance.SetProperty("toolbarButtons", (value))
}

// GettoolbarButtons gets the value of toolbarButtons for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertytoolbarButtons() (value int32, err error) {
	retValue, err := instance.GetProperty("toolbarButtons")
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

// SetuserAgentText sets the value of userAgentText for the instance
func (instance *RSOP_IEAKPolicySetting) SetPropertyuserAgentText(value string) (err error) {
	return instance.SetProperty("userAgentText", (value))
}

// GetuserAgentText gets the value of userAgentText for the instance
func (instance *RSOP_IEAKPolicySetting) GetPropertyuserAgentText() (value string, err error) {
	retValue, err := instance.GetProperty("userAgentText")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
