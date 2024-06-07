// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MDM_Policy_Result01_Start02 struct
type MDM_Policy_Result01_Start02 struct { 
	*cim.WmiInstance

	// 
	AllowPinnedFolderDocuments int32

	// 
	AllowPinnedFolderDownloads int32

	// 
	AllowPinnedFolderFileExplorer int32

	// 
	AllowPinnedFolderHomeGroup int32

	// 
	AllowPinnedFolderMusic int32

	// 
	AllowPinnedFolderNetwork int32

	// 
	AllowPinnedFolderPersonalFolder int32

	// 
	AllowPinnedFolderPictures int32

	// 
	AllowPinnedFolderSettings int32

	// 
	AllowPinnedFolderVideos int32

	// 
	DisableContextMenus int32

	// 
	ForceStartSize int32

	// 
	HideAppList int32

	// 
	HideChangeAccountSettings int32

	// 
	HideFrequentlyUsedApps int32

	// 
	HideHibernate int32

	// 
	HideLock int32

	// 
	HidePowerButton int32

	// 
	HideRecentJumplists int32

	// 
	HideRecentlyAddedApps int32

	// 
	HideRestart int32

	// 
	HideShutDown int32

	// 
	HideSignOut int32

	// 
	HideSleep int32

	// 
	HideSwitchAccount int32

	// 
	HideUserTile int32

	// 
	ImportEdgeAssets string

	// 
	InstanceID string

	// 
	NoPinningToTaskbar int32

	// 
	ParentID string

	// 
	ShowOrHideMostUsedApps int32

	// 
	StartLayout string
}

	func NewMDM_Policy_Result01_Start02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_Start02, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_Policy_Result01_Start02 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_Policy_Result01_Start02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_Policy_Result01_Start02, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_Policy_Result01_Start02 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAllowPinnedFolderDocuments sets the value of AllowPinnedFolderDocuments for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyAllowPinnedFolderDocuments(value int32) (err error) { 
    return instance.SetProperty("AllowPinnedFolderDocuments", (value))
}

// GetAllowPinnedFolderDocuments gets the value of AllowPinnedFolderDocuments for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyAllowPinnedFolderDocuments()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPinnedFolderDocuments")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetAllowPinnedFolderDownloads sets the value of AllowPinnedFolderDownloads for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyAllowPinnedFolderDownloads(value int32) (err error) { 
    return instance.SetProperty("AllowPinnedFolderDownloads", (value))
}

// GetAllowPinnedFolderDownloads gets the value of AllowPinnedFolderDownloads for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyAllowPinnedFolderDownloads()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPinnedFolderDownloads")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetAllowPinnedFolderFileExplorer sets the value of AllowPinnedFolderFileExplorer for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyAllowPinnedFolderFileExplorer(value int32) (err error) { 
    return instance.SetProperty("AllowPinnedFolderFileExplorer", (value))
}

// GetAllowPinnedFolderFileExplorer gets the value of AllowPinnedFolderFileExplorer for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyAllowPinnedFolderFileExplorer()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPinnedFolderFileExplorer")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetAllowPinnedFolderHomeGroup sets the value of AllowPinnedFolderHomeGroup for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyAllowPinnedFolderHomeGroup(value int32) (err error) { 
    return instance.SetProperty("AllowPinnedFolderHomeGroup", (value))
}

// GetAllowPinnedFolderHomeGroup gets the value of AllowPinnedFolderHomeGroup for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyAllowPinnedFolderHomeGroup()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPinnedFolderHomeGroup")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetAllowPinnedFolderMusic sets the value of AllowPinnedFolderMusic for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyAllowPinnedFolderMusic(value int32) (err error) { 
    return instance.SetProperty("AllowPinnedFolderMusic", (value))
}

// GetAllowPinnedFolderMusic gets the value of AllowPinnedFolderMusic for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyAllowPinnedFolderMusic()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPinnedFolderMusic")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetAllowPinnedFolderNetwork sets the value of AllowPinnedFolderNetwork for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyAllowPinnedFolderNetwork(value int32) (err error) { 
    return instance.SetProperty("AllowPinnedFolderNetwork", (value))
}

// GetAllowPinnedFolderNetwork gets the value of AllowPinnedFolderNetwork for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyAllowPinnedFolderNetwork()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPinnedFolderNetwork")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetAllowPinnedFolderPersonalFolder sets the value of AllowPinnedFolderPersonalFolder for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyAllowPinnedFolderPersonalFolder(value int32) (err error) { 
    return instance.SetProperty("AllowPinnedFolderPersonalFolder", (value))
}

// GetAllowPinnedFolderPersonalFolder gets the value of AllowPinnedFolderPersonalFolder for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyAllowPinnedFolderPersonalFolder()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPinnedFolderPersonalFolder")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetAllowPinnedFolderPictures sets the value of AllowPinnedFolderPictures for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyAllowPinnedFolderPictures(value int32) (err error) { 
    return instance.SetProperty("AllowPinnedFolderPictures", (value))
}

// GetAllowPinnedFolderPictures gets the value of AllowPinnedFolderPictures for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyAllowPinnedFolderPictures()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPinnedFolderPictures")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetAllowPinnedFolderSettings sets the value of AllowPinnedFolderSettings for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyAllowPinnedFolderSettings(value int32) (err error) { 
    return instance.SetProperty("AllowPinnedFolderSettings", (value))
}

// GetAllowPinnedFolderSettings gets the value of AllowPinnedFolderSettings for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyAllowPinnedFolderSettings()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPinnedFolderSettings")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetAllowPinnedFolderVideos sets the value of AllowPinnedFolderVideos for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyAllowPinnedFolderVideos(value int32) (err error) { 
    return instance.SetProperty("AllowPinnedFolderVideos", (value))
}

// GetAllowPinnedFolderVideos gets the value of AllowPinnedFolderVideos for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyAllowPinnedFolderVideos()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPinnedFolderVideos")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetDisableContextMenus sets the value of DisableContextMenus for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyDisableContextMenus(value int32) (err error) { 
    return instance.SetProperty("DisableContextMenus", (value))
}

// GetDisableContextMenus gets the value of DisableContextMenus for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyDisableContextMenus()(value int32, err error) { 
    retValue, err := instance.GetProperty("DisableContextMenus")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetForceStartSize sets the value of ForceStartSize for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyForceStartSize(value int32) (err error) { 
    return instance.SetProperty("ForceStartSize", (value))
}

// GetForceStartSize gets the value of ForceStartSize for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyForceStartSize()(value int32, err error) { 
    retValue, err := instance.GetProperty("ForceStartSize")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHideAppList sets the value of HideAppList for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHideAppList(value int32) (err error) { 
    return instance.SetProperty("HideAppList", (value))
}

// GetHideAppList gets the value of HideAppList for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHideAppList()(value int32, err error) { 
    retValue, err := instance.GetProperty("HideAppList")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHideChangeAccountSettings sets the value of HideChangeAccountSettings for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHideChangeAccountSettings(value int32) (err error) { 
    return instance.SetProperty("HideChangeAccountSettings", (value))
}

// GetHideChangeAccountSettings gets the value of HideChangeAccountSettings for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHideChangeAccountSettings()(value int32, err error) { 
    retValue, err := instance.GetProperty("HideChangeAccountSettings")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHideFrequentlyUsedApps sets the value of HideFrequentlyUsedApps for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHideFrequentlyUsedApps(value int32) (err error) { 
    return instance.SetProperty("HideFrequentlyUsedApps", (value))
}

// GetHideFrequentlyUsedApps gets the value of HideFrequentlyUsedApps for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHideFrequentlyUsedApps()(value int32, err error) { 
    retValue, err := instance.GetProperty("HideFrequentlyUsedApps")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHideHibernate sets the value of HideHibernate for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHideHibernate(value int32) (err error) { 
    return instance.SetProperty("HideHibernate", (value))
}

// GetHideHibernate gets the value of HideHibernate for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHideHibernate()(value int32, err error) { 
    retValue, err := instance.GetProperty("HideHibernate")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHideLock sets the value of HideLock for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHideLock(value int32) (err error) { 
    return instance.SetProperty("HideLock", (value))
}

// GetHideLock gets the value of HideLock for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHideLock()(value int32, err error) { 
    retValue, err := instance.GetProperty("HideLock")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHidePowerButton sets the value of HidePowerButton for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHidePowerButton(value int32) (err error) { 
    return instance.SetProperty("HidePowerButton", (value))
}

// GetHidePowerButton gets the value of HidePowerButton for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHidePowerButton()(value int32, err error) { 
    retValue, err := instance.GetProperty("HidePowerButton")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHideRecentJumplists sets the value of HideRecentJumplists for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHideRecentJumplists(value int32) (err error) { 
    return instance.SetProperty("HideRecentJumplists", (value))
}

// GetHideRecentJumplists gets the value of HideRecentJumplists for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHideRecentJumplists()(value int32, err error) { 
    retValue, err := instance.GetProperty("HideRecentJumplists")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHideRecentlyAddedApps sets the value of HideRecentlyAddedApps for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHideRecentlyAddedApps(value int32) (err error) { 
    return instance.SetProperty("HideRecentlyAddedApps", (value))
}

// GetHideRecentlyAddedApps gets the value of HideRecentlyAddedApps for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHideRecentlyAddedApps()(value int32, err error) { 
    retValue, err := instance.GetProperty("HideRecentlyAddedApps")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHideRestart sets the value of HideRestart for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHideRestart(value int32) (err error) { 
    return instance.SetProperty("HideRestart", (value))
}

// GetHideRestart gets the value of HideRestart for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHideRestart()(value int32, err error) { 
    retValue, err := instance.GetProperty("HideRestart")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHideShutDown sets the value of HideShutDown for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHideShutDown(value int32) (err error) { 
    return instance.SetProperty("HideShutDown", (value))
}

// GetHideShutDown gets the value of HideShutDown for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHideShutDown()(value int32, err error) { 
    retValue, err := instance.GetProperty("HideShutDown")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHideSignOut sets the value of HideSignOut for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHideSignOut(value int32) (err error) { 
    return instance.SetProperty("HideSignOut", (value))
}

// GetHideSignOut gets the value of HideSignOut for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHideSignOut()(value int32, err error) { 
    retValue, err := instance.GetProperty("HideSignOut")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHideSleep sets the value of HideSleep for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHideSleep(value int32) (err error) { 
    return instance.SetProperty("HideSleep", (value))
}

// GetHideSleep gets the value of HideSleep for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHideSleep()(value int32, err error) { 
    retValue, err := instance.GetProperty("HideSleep")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHideSwitchAccount sets the value of HideSwitchAccount for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHideSwitchAccount(value int32) (err error) { 
    return instance.SetProperty("HideSwitchAccount", (value))
}

// GetHideSwitchAccount gets the value of HideSwitchAccount for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHideSwitchAccount()(value int32, err error) { 
    retValue, err := instance.GetProperty("HideSwitchAccount")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetHideUserTile sets the value of HideUserTile for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyHideUserTile(value int32) (err error) { 
    return instance.SetProperty("HideUserTile", (value))
}

// GetHideUserTile gets the value of HideUserTile for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyHideUserTile()(value int32, err error) { 
    retValue, err := instance.GetProperty("HideUserTile")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetImportEdgeAssets sets the value of ImportEdgeAssets for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyImportEdgeAssets(value string) (err error) { 
    return instance.SetProperty("ImportEdgeAssets", (value))
}

// GetImportEdgeAssets gets the value of ImportEdgeAssets for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyImportEdgeAssets()(value string, err error) { 
    retValue, err := instance.GetProperty("ImportEdgeAssets")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyInstanceID()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetNoPinningToTaskbar sets the value of NoPinningToTaskbar for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyNoPinningToTaskbar(value int32) (err error) { 
    return instance.SetProperty("NoPinningToTaskbar", (value))
}

// GetNoPinningToTaskbar gets the value of NoPinningToTaskbar for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyNoPinningToTaskbar()(value int32, err error) { 
    retValue, err := instance.GetProperty("NoPinningToTaskbar")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyParentID()(value string, err error) { 
    retValue, err := instance.GetProperty("ParentID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetShowOrHideMostUsedApps sets the value of ShowOrHideMostUsedApps for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyShowOrHideMostUsedApps(value int32) (err error) { 
    return instance.SetProperty("ShowOrHideMostUsedApps", (value))
}

// GetShowOrHideMostUsedApps gets the value of ShowOrHideMostUsedApps for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyShowOrHideMostUsedApps()(value int32, err error) { 
    retValue, err := instance.GetProperty("ShowOrHideMostUsedApps")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetStartLayout sets the value of StartLayout for the instance
func (instance *MDM_Policy_Result01_Start02) SetPropertyStartLayout(value string) (err error) { 
    return instance.SetProperty("StartLayout", (value))
}

// GetStartLayout gets the value of StartLayout for the instance
func (instance *MDM_Policy_Result01_Start02) GetPropertyStartLayout()(value string, err error) { 
    retValue, err := instance.GetProperty("StartLayout")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

