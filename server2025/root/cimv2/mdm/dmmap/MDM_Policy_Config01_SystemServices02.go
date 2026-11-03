// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
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

// MDM_Policy_Config01_SystemServices02 struct
type MDM_Policy_Config01_SystemServices02 struct {
	*cim.WmiInstance

	//
	ConfigureComputerBrowserServiceStartupMode int32

	//
	ConfigureHomeGroupListenerServiceStartupMode int32

	//
	ConfigureHomeGroupProviderServiceStartupMode int32

	//
	ConfigureIISAdminServiceStartupMode int32

	//
	ConfigureInfraredMonitorServiceStartupMode int32

	//
	ConfigureInternetConnectionSharingServiceStartupMode int32

	//
	ConfigureLxssManagerServiceStartupMode int32

	//
	ConfigureMicrosoftFTPServiceStartupMode int32

	//
	ConfigureRemoteProcedureCallLocatorServiceStartupMode int32

	//
	ConfigureRoutingAndRemoteAccessServiceStartupMode int32

	//
	ConfigureSimpleTCPIPServicesStartupMode int32

	//
	ConfigureSpecialAdministrationConsoleHelperServiceStartupMode int32

	//
	ConfigureSSDPDiscoveryServiceStartupMode int32

	//
	ConfigureUPnPDeviceHostServiceStartupMode int32

	//
	ConfigureWebManagementServiceStartupMode int32

	//
	ConfigureWindowsMediaPlayerNetworkSharingServiceStartupMode int32

	//
	ConfigureWindowsMobileHotspotServiceStartupMode int32

	//
	ConfigureWorldWideWebPublishingServiceStartupMode int32

	//
	ConfigureXboxAccessoryManagementServiceStartupMode int32

	//
	ConfigureXboxLiveAuthManagerServiceStartupMode int32

	//
	ConfigureXboxLiveGameSaveServiceStartupMode int32

	//
	ConfigureXboxLiveNetworkingServiceStartupMode int32

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_Policy_Config01_SystemServices02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_SystemServices02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_SystemServices02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_SystemServices02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_SystemServices02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_SystemServices02{
		WmiInstance: tmp,
	}
	return
}

// SetConfigureComputerBrowserServiceStartupMode sets the value of ConfigureComputerBrowserServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureComputerBrowserServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureComputerBrowserServiceStartupMode", (value))
}

// GetConfigureComputerBrowserServiceStartupMode gets the value of ConfigureComputerBrowserServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureComputerBrowserServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureComputerBrowserServiceStartupMode")
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

// SetConfigureHomeGroupListenerServiceStartupMode sets the value of ConfigureHomeGroupListenerServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureHomeGroupListenerServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureHomeGroupListenerServiceStartupMode", (value))
}

// GetConfigureHomeGroupListenerServiceStartupMode gets the value of ConfigureHomeGroupListenerServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureHomeGroupListenerServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureHomeGroupListenerServiceStartupMode")
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

// SetConfigureHomeGroupProviderServiceStartupMode sets the value of ConfigureHomeGroupProviderServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureHomeGroupProviderServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureHomeGroupProviderServiceStartupMode", (value))
}

// GetConfigureHomeGroupProviderServiceStartupMode gets the value of ConfigureHomeGroupProviderServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureHomeGroupProviderServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureHomeGroupProviderServiceStartupMode")
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

// SetConfigureIISAdminServiceStartupMode sets the value of ConfigureIISAdminServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureIISAdminServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureIISAdminServiceStartupMode", (value))
}

// GetConfigureIISAdminServiceStartupMode gets the value of ConfigureIISAdminServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureIISAdminServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureIISAdminServiceStartupMode")
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

// SetConfigureInfraredMonitorServiceStartupMode sets the value of ConfigureInfraredMonitorServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureInfraredMonitorServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureInfraredMonitorServiceStartupMode", (value))
}

// GetConfigureInfraredMonitorServiceStartupMode gets the value of ConfigureInfraredMonitorServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureInfraredMonitorServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureInfraredMonitorServiceStartupMode")
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

// SetConfigureInternetConnectionSharingServiceStartupMode sets the value of ConfigureInternetConnectionSharingServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureInternetConnectionSharingServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureInternetConnectionSharingServiceStartupMode", (value))
}

// GetConfigureInternetConnectionSharingServiceStartupMode gets the value of ConfigureInternetConnectionSharingServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureInternetConnectionSharingServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureInternetConnectionSharingServiceStartupMode")
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

// SetConfigureLxssManagerServiceStartupMode sets the value of ConfigureLxssManagerServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureLxssManagerServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureLxssManagerServiceStartupMode", (value))
}

// GetConfigureLxssManagerServiceStartupMode gets the value of ConfigureLxssManagerServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureLxssManagerServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureLxssManagerServiceStartupMode")
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

// SetConfigureMicrosoftFTPServiceStartupMode sets the value of ConfigureMicrosoftFTPServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureMicrosoftFTPServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureMicrosoftFTPServiceStartupMode", (value))
}

// GetConfigureMicrosoftFTPServiceStartupMode gets the value of ConfigureMicrosoftFTPServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureMicrosoftFTPServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureMicrosoftFTPServiceStartupMode")
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

// SetConfigureRemoteProcedureCallLocatorServiceStartupMode sets the value of ConfigureRemoteProcedureCallLocatorServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureRemoteProcedureCallLocatorServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureRemoteProcedureCallLocatorServiceStartupMode", (value))
}

// GetConfigureRemoteProcedureCallLocatorServiceStartupMode gets the value of ConfigureRemoteProcedureCallLocatorServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureRemoteProcedureCallLocatorServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureRemoteProcedureCallLocatorServiceStartupMode")
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

// SetConfigureRoutingAndRemoteAccessServiceStartupMode sets the value of ConfigureRoutingAndRemoteAccessServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureRoutingAndRemoteAccessServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureRoutingAndRemoteAccessServiceStartupMode", (value))
}

// GetConfigureRoutingAndRemoteAccessServiceStartupMode gets the value of ConfigureRoutingAndRemoteAccessServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureRoutingAndRemoteAccessServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureRoutingAndRemoteAccessServiceStartupMode")
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

// SetConfigureSimpleTCPIPServicesStartupMode sets the value of ConfigureSimpleTCPIPServicesStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureSimpleTCPIPServicesStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureSimpleTCPIPServicesStartupMode", (value))
}

// GetConfigureSimpleTCPIPServicesStartupMode gets the value of ConfigureSimpleTCPIPServicesStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureSimpleTCPIPServicesStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureSimpleTCPIPServicesStartupMode")
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

// SetConfigureSpecialAdministrationConsoleHelperServiceStartupMode sets the value of ConfigureSpecialAdministrationConsoleHelperServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureSpecialAdministrationConsoleHelperServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureSpecialAdministrationConsoleHelperServiceStartupMode", (value))
}

// GetConfigureSpecialAdministrationConsoleHelperServiceStartupMode gets the value of ConfigureSpecialAdministrationConsoleHelperServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureSpecialAdministrationConsoleHelperServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureSpecialAdministrationConsoleHelperServiceStartupMode")
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

// SetConfigureSSDPDiscoveryServiceStartupMode sets the value of ConfigureSSDPDiscoveryServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureSSDPDiscoveryServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureSSDPDiscoveryServiceStartupMode", (value))
}

// GetConfigureSSDPDiscoveryServiceStartupMode gets the value of ConfigureSSDPDiscoveryServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureSSDPDiscoveryServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureSSDPDiscoveryServiceStartupMode")
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

// SetConfigureUPnPDeviceHostServiceStartupMode sets the value of ConfigureUPnPDeviceHostServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureUPnPDeviceHostServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureUPnPDeviceHostServiceStartupMode", (value))
}

// GetConfigureUPnPDeviceHostServiceStartupMode gets the value of ConfigureUPnPDeviceHostServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureUPnPDeviceHostServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureUPnPDeviceHostServiceStartupMode")
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

// SetConfigureWebManagementServiceStartupMode sets the value of ConfigureWebManagementServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureWebManagementServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureWebManagementServiceStartupMode", (value))
}

// GetConfigureWebManagementServiceStartupMode gets the value of ConfigureWebManagementServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureWebManagementServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureWebManagementServiceStartupMode")
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

// SetConfigureWindowsMediaPlayerNetworkSharingServiceStartupMode sets the value of ConfigureWindowsMediaPlayerNetworkSharingServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureWindowsMediaPlayerNetworkSharingServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureWindowsMediaPlayerNetworkSharingServiceStartupMode", (value))
}

// GetConfigureWindowsMediaPlayerNetworkSharingServiceStartupMode gets the value of ConfigureWindowsMediaPlayerNetworkSharingServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureWindowsMediaPlayerNetworkSharingServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureWindowsMediaPlayerNetworkSharingServiceStartupMode")
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

// SetConfigureWindowsMobileHotspotServiceStartupMode sets the value of ConfigureWindowsMobileHotspotServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureWindowsMobileHotspotServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureWindowsMobileHotspotServiceStartupMode", (value))
}

// GetConfigureWindowsMobileHotspotServiceStartupMode gets the value of ConfigureWindowsMobileHotspotServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureWindowsMobileHotspotServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureWindowsMobileHotspotServiceStartupMode")
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

// SetConfigureWorldWideWebPublishingServiceStartupMode sets the value of ConfigureWorldWideWebPublishingServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureWorldWideWebPublishingServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureWorldWideWebPublishingServiceStartupMode", (value))
}

// GetConfigureWorldWideWebPublishingServiceStartupMode gets the value of ConfigureWorldWideWebPublishingServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureWorldWideWebPublishingServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureWorldWideWebPublishingServiceStartupMode")
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

// SetConfigureXboxAccessoryManagementServiceStartupMode sets the value of ConfigureXboxAccessoryManagementServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureXboxAccessoryManagementServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureXboxAccessoryManagementServiceStartupMode", (value))
}

// GetConfigureXboxAccessoryManagementServiceStartupMode gets the value of ConfigureXboxAccessoryManagementServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureXboxAccessoryManagementServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureXboxAccessoryManagementServiceStartupMode")
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

// SetConfigureXboxLiveAuthManagerServiceStartupMode sets the value of ConfigureXboxLiveAuthManagerServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureXboxLiveAuthManagerServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureXboxLiveAuthManagerServiceStartupMode", (value))
}

// GetConfigureXboxLiveAuthManagerServiceStartupMode gets the value of ConfigureXboxLiveAuthManagerServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureXboxLiveAuthManagerServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureXboxLiveAuthManagerServiceStartupMode")
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

// SetConfigureXboxLiveGameSaveServiceStartupMode sets the value of ConfigureXboxLiveGameSaveServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureXboxLiveGameSaveServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureXboxLiveGameSaveServiceStartupMode", (value))
}

// GetConfigureXboxLiveGameSaveServiceStartupMode gets the value of ConfigureXboxLiveGameSaveServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureXboxLiveGameSaveServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureXboxLiveGameSaveServiceStartupMode")
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

// SetConfigureXboxLiveNetworkingServiceStartupMode sets the value of ConfigureXboxLiveNetworkingServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyConfigureXboxLiveNetworkingServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureXboxLiveNetworkingServiceStartupMode", (value))
}

// GetConfigureXboxLiveNetworkingServiceStartupMode gets the value of ConfigureXboxLiveNetworkingServiceStartupMode for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyConfigureXboxLiveNetworkingServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureXboxLiveNetworkingServiceStartupMode")
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
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyInstanceID() (value string, err error) {
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_SystemServices02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_SystemServices02) GetPropertyParentID() (value string, err error) {
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
