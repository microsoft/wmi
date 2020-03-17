// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_UserProfile struct
type Win32_UserProfile struct {
	cim.WmiInstance

	// A Win32_FolderRedirectionHealth object that represents the health of the user's redirected AppData\Roaming folder.
	AppDataRoaming Win32_FolderRedirectionHealth

	// A Win32_FolderRedirectionHealth object that represents the health of the user's redirected Contacts folder.
	Contacts Win32_FolderRedirectionHealth

	// A Win32_FolderRedirectionHealth object that represents the health of the user's redirected Desktop folder.
	Desktop Win32_FolderRedirectionHealth

	// A Win32_FolderRedirectionHealth object that represents the health of the user's redirected Documents folder.
	Documents Win32_FolderRedirectionHealth

	// A Win32_FolderRedirectionHealth object that represents the health of the user's redirected Downloads folder.
	Downloads Win32_FolderRedirectionHealth

	// A Win32_FolderRedirectionHealth object that represents the health of the user's redirected Favorites folder.
	Favorites Win32_FolderRedirectionHealth

	// The health status of this profile, based on the values that were set in the Win32_RoamingUserHealthConfiguration properties.
	HealthStatus UserProfile_HealthStatus

	// If the profile is a roaming profile, this property is a DATETIME value that indicates the last time an attempt was made to download the profile from the server, even if it was unsuccessful. If the profile is a local profile, this property is zero.
	LastAttemptedProfileDownloadTime string

	// If the profile is a roaming profile, this property is a DATETIME value that indicates the last time an attempt was made to upload the profile to the server, even if it was unsuccessful.
	LastAttemptedProfileUploadTime string

	// If this profile is a roaming profile, this property is a DATETIME value that indicates the last time the profile's registry hive was uploaded to the server.
	LastBackgroundRegistryUploadTime string

	//
	LastDownloadTime string

	//
	LastUploadTime string

	//
	LastUseTime string

	// A Win32_FolderRedirectionHealth object that represents the health of the user's redirected Links folder.
	Links Win32_FolderRedirectionHealth

	//
	Loaded bool

	//
	LocalPath string

	// A Win32_FolderRedirectionHealth object that represents the health of the user's redirected Music folder.
	Music Win32_FolderRedirectionHealth

	// A Win32_FolderRedirectionHealth object that represents the health of the user's redirected Pictures folder.
	Pictures Win32_FolderRedirectionHealth

	//
	RefCount uint32

	//
	RoamingConfigured bool

	//
	RoamingPath string

	//
	RoamingPreference bool

	// A Win32_FolderRedirectionHealth object that represents the health of the user's redirected Saved Games folder.
	SavedGames Win32_FolderRedirectionHealth

	// A Win32_FolderRedirectionHealth object that represents the health of the user's redirected Searches folder.
	Searches Win32_FolderRedirectionHealth

	//
	SID string

	//
	Special bool

	// A Win32_FolderRedirectionHealth object that represents the health of the user's redirected Start Menu folder.
	StartMenu Win32_FolderRedirectionHealth

	//
	Status uint32

	// A Win32_FolderRedirectionHealth object that represents the health of the user's redirected Videos folder.
	Videos Win32_FolderRedirectionHealth
}

// SetAppDataRoaming sets the value of AppDataRoaming for the instance
func (instance *Win32_UserProfile) SetPropertyAppDataRoaming(value Win32_FolderRedirectionHealth) (err error) {
	return instance.SetProperty("AppDataRoaming", value)
}

// GetAppDataRoaming gets the value of AppDataRoaming for the instance
func (instance *Win32_UserProfile) GetPropertyAppDataRoaming() (value Win32_FolderRedirectionHealth, err error) {
	retValue, err := instance.GetProperty("AppDataRoaming")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_FolderRedirectionHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContacts sets the value of Contacts for the instance
func (instance *Win32_UserProfile) SetPropertyContacts(value Win32_FolderRedirectionHealth) (err error) {
	return instance.SetProperty("Contacts", value)
}

// GetContacts gets the value of Contacts for the instance
func (instance *Win32_UserProfile) GetPropertyContacts() (value Win32_FolderRedirectionHealth, err error) {
	retValue, err := instance.GetProperty("Contacts")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_FolderRedirectionHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDesktop sets the value of Desktop for the instance
func (instance *Win32_UserProfile) SetPropertyDesktop(value Win32_FolderRedirectionHealth) (err error) {
	return instance.SetProperty("Desktop", value)
}

// GetDesktop gets the value of Desktop for the instance
func (instance *Win32_UserProfile) GetPropertyDesktop() (value Win32_FolderRedirectionHealth, err error) {
	retValue, err := instance.GetProperty("Desktop")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_FolderRedirectionHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDocuments sets the value of Documents for the instance
func (instance *Win32_UserProfile) SetPropertyDocuments(value Win32_FolderRedirectionHealth) (err error) {
	return instance.SetProperty("Documents", value)
}

// GetDocuments gets the value of Documents for the instance
func (instance *Win32_UserProfile) GetPropertyDocuments() (value Win32_FolderRedirectionHealth, err error) {
	retValue, err := instance.GetProperty("Documents")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_FolderRedirectionHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDownloads sets the value of Downloads for the instance
func (instance *Win32_UserProfile) SetPropertyDownloads(value Win32_FolderRedirectionHealth) (err error) {
	return instance.SetProperty("Downloads", value)
}

// GetDownloads gets the value of Downloads for the instance
func (instance *Win32_UserProfile) GetPropertyDownloads() (value Win32_FolderRedirectionHealth, err error) {
	retValue, err := instance.GetProperty("Downloads")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_FolderRedirectionHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFavorites sets the value of Favorites for the instance
func (instance *Win32_UserProfile) SetPropertyFavorites(value Win32_FolderRedirectionHealth) (err error) {
	return instance.SetProperty("Favorites", value)
}

// GetFavorites gets the value of Favorites for the instance
func (instance *Win32_UserProfile) GetPropertyFavorites() (value Win32_FolderRedirectionHealth, err error) {
	retValue, err := instance.GetProperty("Favorites")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_FolderRedirectionHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *Win32_UserProfile) SetPropertyHealthStatus(value UserProfile_HealthStatus) (err error) {
	return instance.SetProperty("HealthStatus", value)
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *Win32_UserProfile) GetPropertyHealthStatus() (value UserProfile_HealthStatus, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(UserProfile_HealthStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastAttemptedProfileDownloadTime sets the value of LastAttemptedProfileDownloadTime for the instance
func (instance *Win32_UserProfile) SetPropertyLastAttemptedProfileDownloadTime(value string) (err error) {
	return instance.SetProperty("LastAttemptedProfileDownloadTime", value)
}

// GetLastAttemptedProfileDownloadTime gets the value of LastAttemptedProfileDownloadTime for the instance
func (instance *Win32_UserProfile) GetPropertyLastAttemptedProfileDownloadTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastAttemptedProfileDownloadTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastAttemptedProfileUploadTime sets the value of LastAttemptedProfileUploadTime for the instance
func (instance *Win32_UserProfile) SetPropertyLastAttemptedProfileUploadTime(value string) (err error) {
	return instance.SetProperty("LastAttemptedProfileUploadTime", value)
}

// GetLastAttemptedProfileUploadTime gets the value of LastAttemptedProfileUploadTime for the instance
func (instance *Win32_UserProfile) GetPropertyLastAttemptedProfileUploadTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastAttemptedProfileUploadTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastBackgroundRegistryUploadTime sets the value of LastBackgroundRegistryUploadTime for the instance
func (instance *Win32_UserProfile) SetPropertyLastBackgroundRegistryUploadTime(value string) (err error) {
	return instance.SetProperty("LastBackgroundRegistryUploadTime", value)
}

// GetLastBackgroundRegistryUploadTime gets the value of LastBackgroundRegistryUploadTime for the instance
func (instance *Win32_UserProfile) GetPropertyLastBackgroundRegistryUploadTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastBackgroundRegistryUploadTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastDownloadTime sets the value of LastDownloadTime for the instance
func (instance *Win32_UserProfile) SetPropertyLastDownloadTime(value string) (err error) {
	return instance.SetProperty("LastDownloadTime", value)
}

// GetLastDownloadTime gets the value of LastDownloadTime for the instance
func (instance *Win32_UserProfile) GetPropertyLastDownloadTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastDownloadTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastUploadTime sets the value of LastUploadTime for the instance
func (instance *Win32_UserProfile) SetPropertyLastUploadTime(value string) (err error) {
	return instance.SetProperty("LastUploadTime", value)
}

// GetLastUploadTime gets the value of LastUploadTime for the instance
func (instance *Win32_UserProfile) GetPropertyLastUploadTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastUploadTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastUseTime sets the value of LastUseTime for the instance
func (instance *Win32_UserProfile) SetPropertyLastUseTime(value string) (err error) {
	return instance.SetProperty("LastUseTime", value)
}

// GetLastUseTime gets the value of LastUseTime for the instance
func (instance *Win32_UserProfile) GetPropertyLastUseTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastUseTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLinks sets the value of Links for the instance
func (instance *Win32_UserProfile) SetPropertyLinks(value Win32_FolderRedirectionHealth) (err error) {
	return instance.SetProperty("Links", value)
}

// GetLinks gets the value of Links for the instance
func (instance *Win32_UserProfile) GetPropertyLinks() (value Win32_FolderRedirectionHealth, err error) {
	retValue, err := instance.GetProperty("Links")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_FolderRedirectionHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoaded sets the value of Loaded for the instance
func (instance *Win32_UserProfile) SetPropertyLoaded(value bool) (err error) {
	return instance.SetProperty("Loaded", value)
}

// GetLoaded gets the value of Loaded for the instance
func (instance *Win32_UserProfile) GetPropertyLoaded() (value bool, err error) {
	retValue, err := instance.GetProperty("Loaded")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalPath sets the value of LocalPath for the instance
func (instance *Win32_UserProfile) SetPropertyLocalPath(value string) (err error) {
	return instance.SetProperty("LocalPath", value)
}

// GetLocalPath gets the value of LocalPath for the instance
func (instance *Win32_UserProfile) GetPropertyLocalPath() (value string, err error) {
	retValue, err := instance.GetProperty("LocalPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMusic sets the value of Music for the instance
func (instance *Win32_UserProfile) SetPropertyMusic(value Win32_FolderRedirectionHealth) (err error) {
	return instance.SetProperty("Music", value)
}

// GetMusic gets the value of Music for the instance
func (instance *Win32_UserProfile) GetPropertyMusic() (value Win32_FolderRedirectionHealth, err error) {
	retValue, err := instance.GetProperty("Music")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_FolderRedirectionHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPictures sets the value of Pictures for the instance
func (instance *Win32_UserProfile) SetPropertyPictures(value Win32_FolderRedirectionHealth) (err error) {
	return instance.SetProperty("Pictures", value)
}

// GetPictures gets the value of Pictures for the instance
func (instance *Win32_UserProfile) GetPropertyPictures() (value Win32_FolderRedirectionHealth, err error) {
	retValue, err := instance.GetProperty("Pictures")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_FolderRedirectionHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRefCount sets the value of RefCount for the instance
func (instance *Win32_UserProfile) SetPropertyRefCount(value uint32) (err error) {
	return instance.SetProperty("RefCount", value)
}

// GetRefCount gets the value of RefCount for the instance
func (instance *Win32_UserProfile) GetPropertyRefCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("RefCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoamingConfigured sets the value of RoamingConfigured for the instance
func (instance *Win32_UserProfile) SetPropertyRoamingConfigured(value bool) (err error) {
	return instance.SetProperty("RoamingConfigured", value)
}

// GetRoamingConfigured gets the value of RoamingConfigured for the instance
func (instance *Win32_UserProfile) GetPropertyRoamingConfigured() (value bool, err error) {
	retValue, err := instance.GetProperty("RoamingConfigured")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoamingPath sets the value of RoamingPath for the instance
func (instance *Win32_UserProfile) SetPropertyRoamingPath(value string) (err error) {
	return instance.SetProperty("RoamingPath", value)
}

// GetRoamingPath gets the value of RoamingPath for the instance
func (instance *Win32_UserProfile) GetPropertyRoamingPath() (value string, err error) {
	retValue, err := instance.GetProperty("RoamingPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoamingPreference sets the value of RoamingPreference for the instance
func (instance *Win32_UserProfile) SetPropertyRoamingPreference(value bool) (err error) {
	return instance.SetProperty("RoamingPreference", value)
}

// GetRoamingPreference gets the value of RoamingPreference for the instance
func (instance *Win32_UserProfile) GetPropertyRoamingPreference() (value bool, err error) {
	retValue, err := instance.GetProperty("RoamingPreference")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSavedGames sets the value of SavedGames for the instance
func (instance *Win32_UserProfile) SetPropertySavedGames(value Win32_FolderRedirectionHealth) (err error) {
	return instance.SetProperty("SavedGames", value)
}

// GetSavedGames gets the value of SavedGames for the instance
func (instance *Win32_UserProfile) GetPropertySavedGames() (value Win32_FolderRedirectionHealth, err error) {
	retValue, err := instance.GetProperty("SavedGames")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_FolderRedirectionHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSearches sets the value of Searches for the instance
func (instance *Win32_UserProfile) SetPropertySearches(value Win32_FolderRedirectionHealth) (err error) {
	return instance.SetProperty("Searches", value)
}

// GetSearches gets the value of Searches for the instance
func (instance *Win32_UserProfile) GetPropertySearches() (value Win32_FolderRedirectionHealth, err error) {
	retValue, err := instance.GetProperty("Searches")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_FolderRedirectionHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSID sets the value of SID for the instance
func (instance *Win32_UserProfile) SetPropertySID(value string) (err error) {
	return instance.SetProperty("SID", value)
}

// GetSID gets the value of SID for the instance
func (instance *Win32_UserProfile) GetPropertySID() (value string, err error) {
	retValue, err := instance.GetProperty("SID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpecial sets the value of Special for the instance
func (instance *Win32_UserProfile) SetPropertySpecial(value bool) (err error) {
	return instance.SetProperty("Special", value)
}

// GetSpecial gets the value of Special for the instance
func (instance *Win32_UserProfile) GetPropertySpecial() (value bool, err error) {
	retValue, err := instance.GetProperty("Special")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartMenu sets the value of StartMenu for the instance
func (instance *Win32_UserProfile) SetPropertyStartMenu(value Win32_FolderRedirectionHealth) (err error) {
	return instance.SetProperty("StartMenu", value)
}

// GetStartMenu gets the value of StartMenu for the instance
func (instance *Win32_UserProfile) GetPropertyStartMenu() (value Win32_FolderRedirectionHealth, err error) {
	retValue, err := instance.GetProperty("StartMenu")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_FolderRedirectionHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *Win32_UserProfile) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *Win32_UserProfile) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVideos sets the value of Videos for the instance
func (instance *Win32_UserProfile) SetPropertyVideos(value Win32_FolderRedirectionHealth) (err error) {
	return instance.SetProperty("Videos", value)
}

// GetVideos gets the value of Videos for the instance
func (instance *Win32_UserProfile) GetPropertyVideos() (value Win32_FolderRedirectionHealth, err error) {
	retValue, err := instance.GetProperty("Videos")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_FolderRedirectionHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Flags" type="uint32 "></param>
// <param name="NewOwnerSID" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_UserProfile) ChangeOwner( /* IN */ NewOwnerSID string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ChangeOwner", NewOwnerSID, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
