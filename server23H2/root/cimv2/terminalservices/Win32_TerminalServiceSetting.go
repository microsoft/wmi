// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_TerminalServiceSetting struct
type Win32_TerminalServiceSetting struct {
	*CIM_Setting

	//
	ActiveDesktop uint32

	//
	AllowTSConnections uint32

	//
	DeleteTempFolders uint32

	//
	DirectConnectLicenseServers string

	//
	DisableForcibleLogoff uint32

	//
	EnableAutomaticReconnection uint32

	//
	EnableDFSS uint32

	//
	EnableDiskFSS uint32

	//
	EnableNetworkFSS uint32

	//
	EnableRemoteDesktopMSI uint32

	//
	FallbackPrintDriverType uint32

	//
	GetCapabilitiesID uint32

	//
	HomeDirectory string

	//
	LicensingDescription string

	//
	LicensingName string

	//
	LicensingType uint32

	//
	LimitedUserSessions uint32

	//
	Logons string

	//
	NetworkFSSCatchAllWeight uint32

	//
	NetworkFSSLocalSystemWeight uint32

	//
	NetworkFSSUserSessionWeight uint32

	//
	PolicySourceAllowTSConnections uint32

	//
	PolicySourceConfiguredLicenseServers uint32

	//
	PolicySourceDeleteTempFolders uint32

	//
	PolicySourceDirectConnectLicenseServers uint32

	//
	PolicySourceEnableAutomaticReconnection uint32

	//
	PolicySourceEnableDFSS uint32

	//
	PolicySourceEnableRemoteDesktopMSI uint32

	//
	PolicySourceFallbackPrintDriverType uint32

	//
	PolicySourceHomeDirectory uint32

	//
	PolicySourceLicensingType uint32

	//
	PolicySourceProfilePath uint32

	//
	PolicySourceRedirectSmartCards uint32

	//
	PolicySourceSingleSession uint32

	//
	PolicySourceTimeZoneRedirection uint32

	//
	PolicySourceUseRDEasyPrintDriver uint32

	//
	PolicySourceUseTempFolders uint32

	//
	PossibleLicensingTypes uint32

	//
	ProfilePath string

	//
	RedirectSmartCards uint32

	//
	ServerName string

	//
	SessionBrokerDrainMode uint32

	//
	SingleSession uint32

	//
	TerminalServerMode uint32

	//
	TimeZoneRedirection uint32

	//
	UseRDEasyPrintDriver uint32

	//
	UserPermission uint32

	//
	UseTempFolders uint32
}

func NewWin32_TerminalServiceSettingEx1(instance *cim.WmiInstance) (newInstance *Win32_TerminalServiceSetting, err error) {
	tmp, err := NewCIM_SettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TerminalServiceSetting{
		CIM_Setting: tmp,
	}
	return
}

func NewWin32_TerminalServiceSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TerminalServiceSetting, err error) {
	tmp, err := NewCIM_SettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TerminalServiceSetting{
		CIM_Setting: tmp,
	}
	return
}

// SetActiveDesktop sets the value of ActiveDesktop for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyActiveDesktop(value uint32) (err error) {
	return instance.SetProperty("ActiveDesktop", (value))
}

// GetActiveDesktop gets the value of ActiveDesktop for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyActiveDesktop() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveDesktop")
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

// SetAllowTSConnections sets the value of AllowTSConnections for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyAllowTSConnections(value uint32) (err error) {
	return instance.SetProperty("AllowTSConnections", (value))
}

// GetAllowTSConnections gets the value of AllowTSConnections for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyAllowTSConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllowTSConnections")
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

// SetDeleteTempFolders sets the value of DeleteTempFolders for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyDeleteTempFolders(value uint32) (err error) {
	return instance.SetProperty("DeleteTempFolders", (value))
}

// GetDeleteTempFolders gets the value of DeleteTempFolders for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyDeleteTempFolders() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeleteTempFolders")
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

// SetDirectConnectLicenseServers sets the value of DirectConnectLicenseServers for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyDirectConnectLicenseServers(value string) (err error) {
	return instance.SetProperty("DirectConnectLicenseServers", (value))
}

// GetDirectConnectLicenseServers gets the value of DirectConnectLicenseServers for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyDirectConnectLicenseServers() (value string, err error) {
	retValue, err := instance.GetProperty("DirectConnectLicenseServers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDisableForcibleLogoff sets the value of DisableForcibleLogoff for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyDisableForcibleLogoff(value uint32) (err error) {
	return instance.SetProperty("DisableForcibleLogoff", (value))
}

// GetDisableForcibleLogoff gets the value of DisableForcibleLogoff for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyDisableForcibleLogoff() (value uint32, err error) {
	retValue, err := instance.GetProperty("DisableForcibleLogoff")
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

// SetEnableAutomaticReconnection sets the value of EnableAutomaticReconnection for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyEnableAutomaticReconnection(value uint32) (err error) {
	return instance.SetProperty("EnableAutomaticReconnection", (value))
}

// GetEnableAutomaticReconnection gets the value of EnableAutomaticReconnection for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyEnableAutomaticReconnection() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnableAutomaticReconnection")
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

// SetEnableDFSS sets the value of EnableDFSS for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyEnableDFSS(value uint32) (err error) {
	return instance.SetProperty("EnableDFSS", (value))
}

// GetEnableDFSS gets the value of EnableDFSS for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyEnableDFSS() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnableDFSS")
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

// SetEnableDiskFSS sets the value of EnableDiskFSS for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyEnableDiskFSS(value uint32) (err error) {
	return instance.SetProperty("EnableDiskFSS", (value))
}

// GetEnableDiskFSS gets the value of EnableDiskFSS for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyEnableDiskFSS() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnableDiskFSS")
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

// SetEnableNetworkFSS sets the value of EnableNetworkFSS for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyEnableNetworkFSS(value uint32) (err error) {
	return instance.SetProperty("EnableNetworkFSS", (value))
}

// GetEnableNetworkFSS gets the value of EnableNetworkFSS for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyEnableNetworkFSS() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnableNetworkFSS")
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

// SetEnableRemoteDesktopMSI sets the value of EnableRemoteDesktopMSI for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyEnableRemoteDesktopMSI(value uint32) (err error) {
	return instance.SetProperty("EnableRemoteDesktopMSI", (value))
}

// GetEnableRemoteDesktopMSI gets the value of EnableRemoteDesktopMSI for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyEnableRemoteDesktopMSI() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnableRemoteDesktopMSI")
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

// SetFallbackPrintDriverType sets the value of FallbackPrintDriverType for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyFallbackPrintDriverType(value uint32) (err error) {
	return instance.SetProperty("FallbackPrintDriverType", (value))
}

// GetFallbackPrintDriverType gets the value of FallbackPrintDriverType for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyFallbackPrintDriverType() (value uint32, err error) {
	retValue, err := instance.GetProperty("FallbackPrintDriverType")
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

// SetGetCapabilitiesID sets the value of GetCapabilitiesID for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyGetCapabilitiesID(value uint32) (err error) {
	return instance.SetProperty("GetCapabilitiesID", (value))
}

// GetGetCapabilitiesID gets the value of GetCapabilitiesID for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyGetCapabilitiesID() (value uint32, err error) {
	retValue, err := instance.GetProperty("GetCapabilitiesID")
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

// SetHomeDirectory sets the value of HomeDirectory for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyHomeDirectory(value string) (err error) {
	return instance.SetProperty("HomeDirectory", (value))
}

// GetHomeDirectory gets the value of HomeDirectory for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyHomeDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("HomeDirectory")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLicensingDescription sets the value of LicensingDescription for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyLicensingDescription(value string) (err error) {
	return instance.SetProperty("LicensingDescription", (value))
}

// GetLicensingDescription gets the value of LicensingDescription for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyLicensingDescription() (value string, err error) {
	retValue, err := instance.GetProperty("LicensingDescription")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLicensingName sets the value of LicensingName for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyLicensingName(value string) (err error) {
	return instance.SetProperty("LicensingName", (value))
}

// GetLicensingName gets the value of LicensingName for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyLicensingName() (value string, err error) {
	retValue, err := instance.GetProperty("LicensingName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLicensingType sets the value of LicensingType for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyLicensingType(value uint32) (err error) {
	return instance.SetProperty("LicensingType", (value))
}

// GetLicensingType gets the value of LicensingType for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyLicensingType() (value uint32, err error) {
	retValue, err := instance.GetProperty("LicensingType")
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

// SetLimitedUserSessions sets the value of LimitedUserSessions for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyLimitedUserSessions(value uint32) (err error) {
	return instance.SetProperty("LimitedUserSessions", (value))
}

// GetLimitedUserSessions gets the value of LimitedUserSessions for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyLimitedUserSessions() (value uint32, err error) {
	retValue, err := instance.GetProperty("LimitedUserSessions")
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

// SetLogons sets the value of Logons for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyLogons(value string) (err error) {
	return instance.SetProperty("Logons", (value))
}

// GetLogons gets the value of Logons for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyLogons() (value string, err error) {
	retValue, err := instance.GetProperty("Logons")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetNetworkFSSCatchAllWeight sets the value of NetworkFSSCatchAllWeight for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyNetworkFSSCatchAllWeight(value uint32) (err error) {
	return instance.SetProperty("NetworkFSSCatchAllWeight", (value))
}

// GetNetworkFSSCatchAllWeight gets the value of NetworkFSSCatchAllWeight for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyNetworkFSSCatchAllWeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("NetworkFSSCatchAllWeight")
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

// SetNetworkFSSLocalSystemWeight sets the value of NetworkFSSLocalSystemWeight for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyNetworkFSSLocalSystemWeight(value uint32) (err error) {
	return instance.SetProperty("NetworkFSSLocalSystemWeight", (value))
}

// GetNetworkFSSLocalSystemWeight gets the value of NetworkFSSLocalSystemWeight for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyNetworkFSSLocalSystemWeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("NetworkFSSLocalSystemWeight")
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

// SetNetworkFSSUserSessionWeight sets the value of NetworkFSSUserSessionWeight for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyNetworkFSSUserSessionWeight(value uint32) (err error) {
	return instance.SetProperty("NetworkFSSUserSessionWeight", (value))
}

// GetNetworkFSSUserSessionWeight gets the value of NetworkFSSUserSessionWeight for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyNetworkFSSUserSessionWeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("NetworkFSSUserSessionWeight")
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

// SetPolicySourceAllowTSConnections sets the value of PolicySourceAllowTSConnections for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceAllowTSConnections(value uint32) (err error) {
	return instance.SetProperty("PolicySourceAllowTSConnections", (value))
}

// GetPolicySourceAllowTSConnections gets the value of PolicySourceAllowTSConnections for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceAllowTSConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceAllowTSConnections")
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

// SetPolicySourceConfiguredLicenseServers sets the value of PolicySourceConfiguredLicenseServers for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceConfiguredLicenseServers(value uint32) (err error) {
	return instance.SetProperty("PolicySourceConfiguredLicenseServers", (value))
}

// GetPolicySourceConfiguredLicenseServers gets the value of PolicySourceConfiguredLicenseServers for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceConfiguredLicenseServers() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceConfiguredLicenseServers")
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

// SetPolicySourceDeleteTempFolders sets the value of PolicySourceDeleteTempFolders for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceDeleteTempFolders(value uint32) (err error) {
	return instance.SetProperty("PolicySourceDeleteTempFolders", (value))
}

// GetPolicySourceDeleteTempFolders gets the value of PolicySourceDeleteTempFolders for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceDeleteTempFolders() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceDeleteTempFolders")
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

// SetPolicySourceDirectConnectLicenseServers sets the value of PolicySourceDirectConnectLicenseServers for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceDirectConnectLicenseServers(value uint32) (err error) {
	return instance.SetProperty("PolicySourceDirectConnectLicenseServers", (value))
}

// GetPolicySourceDirectConnectLicenseServers gets the value of PolicySourceDirectConnectLicenseServers for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceDirectConnectLicenseServers() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceDirectConnectLicenseServers")
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

// SetPolicySourceEnableAutomaticReconnection sets the value of PolicySourceEnableAutomaticReconnection for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceEnableAutomaticReconnection(value uint32) (err error) {
	return instance.SetProperty("PolicySourceEnableAutomaticReconnection", (value))
}

// GetPolicySourceEnableAutomaticReconnection gets the value of PolicySourceEnableAutomaticReconnection for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceEnableAutomaticReconnection() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceEnableAutomaticReconnection")
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

// SetPolicySourceEnableDFSS sets the value of PolicySourceEnableDFSS for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceEnableDFSS(value uint32) (err error) {
	return instance.SetProperty("PolicySourceEnableDFSS", (value))
}

// GetPolicySourceEnableDFSS gets the value of PolicySourceEnableDFSS for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceEnableDFSS() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceEnableDFSS")
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

// SetPolicySourceEnableRemoteDesktopMSI sets the value of PolicySourceEnableRemoteDesktopMSI for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceEnableRemoteDesktopMSI(value uint32) (err error) {
	return instance.SetProperty("PolicySourceEnableRemoteDesktopMSI", (value))
}

// GetPolicySourceEnableRemoteDesktopMSI gets the value of PolicySourceEnableRemoteDesktopMSI for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceEnableRemoteDesktopMSI() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceEnableRemoteDesktopMSI")
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

// SetPolicySourceFallbackPrintDriverType sets the value of PolicySourceFallbackPrintDriverType for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceFallbackPrintDriverType(value uint32) (err error) {
	return instance.SetProperty("PolicySourceFallbackPrintDriverType", (value))
}

// GetPolicySourceFallbackPrintDriverType gets the value of PolicySourceFallbackPrintDriverType for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceFallbackPrintDriverType() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceFallbackPrintDriverType")
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

// SetPolicySourceHomeDirectory sets the value of PolicySourceHomeDirectory for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceHomeDirectory(value uint32) (err error) {
	return instance.SetProperty("PolicySourceHomeDirectory", (value))
}

// GetPolicySourceHomeDirectory gets the value of PolicySourceHomeDirectory for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceHomeDirectory() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceHomeDirectory")
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

// SetPolicySourceLicensingType sets the value of PolicySourceLicensingType for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceLicensingType(value uint32) (err error) {
	return instance.SetProperty("PolicySourceLicensingType", (value))
}

// GetPolicySourceLicensingType gets the value of PolicySourceLicensingType for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceLicensingType() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceLicensingType")
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

// SetPolicySourceProfilePath sets the value of PolicySourceProfilePath for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceProfilePath(value uint32) (err error) {
	return instance.SetProperty("PolicySourceProfilePath", (value))
}

// GetPolicySourceProfilePath gets the value of PolicySourceProfilePath for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceProfilePath() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceProfilePath")
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

// SetPolicySourceRedirectSmartCards sets the value of PolicySourceRedirectSmartCards for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceRedirectSmartCards(value uint32) (err error) {
	return instance.SetProperty("PolicySourceRedirectSmartCards", (value))
}

// GetPolicySourceRedirectSmartCards gets the value of PolicySourceRedirectSmartCards for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceRedirectSmartCards() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceRedirectSmartCards")
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

// SetPolicySourceSingleSession sets the value of PolicySourceSingleSession for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceSingleSession(value uint32) (err error) {
	return instance.SetProperty("PolicySourceSingleSession", (value))
}

// GetPolicySourceSingleSession gets the value of PolicySourceSingleSession for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceSingleSession() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceSingleSession")
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

// SetPolicySourceTimeZoneRedirection sets the value of PolicySourceTimeZoneRedirection for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceTimeZoneRedirection(value uint32) (err error) {
	return instance.SetProperty("PolicySourceTimeZoneRedirection", (value))
}

// GetPolicySourceTimeZoneRedirection gets the value of PolicySourceTimeZoneRedirection for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceTimeZoneRedirection() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceTimeZoneRedirection")
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

// SetPolicySourceUseRDEasyPrintDriver sets the value of PolicySourceUseRDEasyPrintDriver for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceUseRDEasyPrintDriver(value uint32) (err error) {
	return instance.SetProperty("PolicySourceUseRDEasyPrintDriver", (value))
}

// GetPolicySourceUseRDEasyPrintDriver gets the value of PolicySourceUseRDEasyPrintDriver for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceUseRDEasyPrintDriver() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceUseRDEasyPrintDriver")
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

// SetPolicySourceUseTempFolders sets the value of PolicySourceUseTempFolders for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPolicySourceUseTempFolders(value uint32) (err error) {
	return instance.SetProperty("PolicySourceUseTempFolders", (value))
}

// GetPolicySourceUseTempFolders gets the value of PolicySourceUseTempFolders for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPolicySourceUseTempFolders() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceUseTempFolders")
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

// SetPossibleLicensingTypes sets the value of PossibleLicensingTypes for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyPossibleLicensingTypes(value uint32) (err error) {
	return instance.SetProperty("PossibleLicensingTypes", (value))
}

// GetPossibleLicensingTypes gets the value of PossibleLicensingTypes for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyPossibleLicensingTypes() (value uint32, err error) {
	retValue, err := instance.GetProperty("PossibleLicensingTypes")
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

// SetProfilePath sets the value of ProfilePath for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyProfilePath(value string) (err error) {
	return instance.SetProperty("ProfilePath", (value))
}

// GetProfilePath gets the value of ProfilePath for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyProfilePath() (value string, err error) {
	retValue, err := instance.GetProperty("ProfilePath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetRedirectSmartCards sets the value of RedirectSmartCards for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyRedirectSmartCards(value uint32) (err error) {
	return instance.SetProperty("RedirectSmartCards", (value))
}

// GetRedirectSmartCards gets the value of RedirectSmartCards for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyRedirectSmartCards() (value uint32, err error) {
	retValue, err := instance.GetProperty("RedirectSmartCards")
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

// SetServerName sets the value of ServerName for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyServerName(value string) (err error) {
	return instance.SetProperty("ServerName", (value))
}

// GetServerName gets the value of ServerName for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyServerName() (value string, err error) {
	retValue, err := instance.GetProperty("ServerName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSessionBrokerDrainMode sets the value of SessionBrokerDrainMode for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertySessionBrokerDrainMode(value uint32) (err error) {
	return instance.SetProperty("SessionBrokerDrainMode", (value))
}

// GetSessionBrokerDrainMode gets the value of SessionBrokerDrainMode for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertySessionBrokerDrainMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionBrokerDrainMode")
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

// SetSingleSession sets the value of SingleSession for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertySingleSession(value uint32) (err error) {
	return instance.SetProperty("SingleSession", (value))
}

// GetSingleSession gets the value of SingleSession for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertySingleSession() (value uint32, err error) {
	retValue, err := instance.GetProperty("SingleSession")
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

// SetTerminalServerMode sets the value of TerminalServerMode for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyTerminalServerMode(value uint32) (err error) {
	return instance.SetProperty("TerminalServerMode", (value))
}

// GetTerminalServerMode gets the value of TerminalServerMode for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyTerminalServerMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("TerminalServerMode")
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

// SetTimeZoneRedirection sets the value of TimeZoneRedirection for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyTimeZoneRedirection(value uint32) (err error) {
	return instance.SetProperty("TimeZoneRedirection", (value))
}

// GetTimeZoneRedirection gets the value of TimeZoneRedirection for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyTimeZoneRedirection() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimeZoneRedirection")
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

// SetUseRDEasyPrintDriver sets the value of UseRDEasyPrintDriver for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyUseRDEasyPrintDriver(value uint32) (err error) {
	return instance.SetProperty("UseRDEasyPrintDriver", (value))
}

// GetUseRDEasyPrintDriver gets the value of UseRDEasyPrintDriver for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyUseRDEasyPrintDriver() (value uint32, err error) {
	retValue, err := instance.GetProperty("UseRDEasyPrintDriver")
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

// SetUserPermission sets the value of UserPermission for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyUserPermission(value uint32) (err error) {
	return instance.SetProperty("UserPermission", (value))
}

// GetUserPermission gets the value of UserPermission for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyUserPermission() (value uint32, err error) {
	retValue, err := instance.GetProperty("UserPermission")
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

// SetUseTempFolders sets the value of UseTempFolders for the instance
func (instance *Win32_TerminalServiceSetting) SetPropertyUseTempFolders(value uint32) (err error) {
	return instance.SetProperty("UseTempFolders", (value))
}

// GetUseTempFolders gets the value of UseTempFolders for the instance
func (instance *Win32_TerminalServiceSetting) GetPropertyUseTempFolders() (value uint32, err error) {
	retValue, err := instance.GetProperty("UseTempFolders")
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

//

// <param name="TimeZoneRedirection" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) SetTimeZoneRedirection( /* IN */ TimeZoneRedirection uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetTimeZoneRedirection", TimeZoneRedirection)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="LicenseServerName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) AddDirectConnectLicenseServer( /* IN */ LicenseServerName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddDirectConnectLicenseServer", LicenseServerName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="LicenseServerName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) DeleteDirectConnectLicenseServer( /* IN */ LicenseServerName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DeleteDirectConnectLicenseServer", LicenseServerName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="LicenseServerList" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) UpdateDirectConnectLicenseServer( /* IN */ LicenseServerList string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UpdateDirectConnectLicenseServer", LicenseServerList)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="RegisteredLSList" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) GetRegisteredLicenseServerList( /* OUT */ RegisteredLSList []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetRegisteredLicenseServerList")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
// <param name="SpecifiedLSList" type="string []"></param>
func (instance *Win32_TerminalServiceSetting) GetSpecifiedLicenseServerList( /* OUT */ SpecifiedLSList []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSpecifiedLicenseServerList")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SpecifiedLSList" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) SetSpecifiedLicenseServerList( /* IN */ SpecifiedLSList []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetSpecifiedLicenseServerList", SpecifiedLSList)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) EmptySpecifiedLicenseServerList() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("EmptySpecifiedLicenseServerList")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="LicenseServerName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) SetPrimaryLicenseServer( /* IN */ LicenseServerName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPrimaryLicenseServer", LicenseServerName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="LicenseServerName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) AddLSToSpecifiedLicenseServerList( /* IN */ LicenseServerName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddLSToSpecifiedLicenseServerList", LicenseServerName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="LicenseServerName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) RemoveLSFromSpecifiedLicenseServerList( /* IN */ LicenseServerName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveLSFromSpecifiedLicenseServerList", LicenseServerName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="LicensingClientId" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="TenantAadToken" type="string "></param>
// <param name="TokenLength" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) GetLicensingAadInfo( /* OUT */ LicensingClientId string,
	/* OUT */ TokenLength uint32,
	/* OUT */ TenantAadToken string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetLicensingAadInfo")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="LicensingClientId" type="string "></param>
// <param name="TenantAadToken" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) SetLicensingAadInfo( /* IN */ LicensingClientId string,
	/* IN */ TenantAadToken string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetLicensingAadInfo", LicensingClientId, TenantAadToken)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Version" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) GetRDSVersion( /* OUT */ Version uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetRDSVersion")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="LicensingType" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) ChangeMode( /* IN */ LicensingType uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ChangeMode", LicensingType)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="AllowTSConnections" type="uint32 "></param>
// <param name="ModifyFirewallException" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) SetAllowTSConnections( /* IN */ AllowTSConnections uint32,
	/* IN */ ModifyFirewallException uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetAllowTSConnections", AllowTSConnections, ModifyFirewallException)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="SingleSession" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) SetSingleSession( /* IN */ SingleSession uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetSingleSession", SingleSession)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="DisableForcibleLogoff" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) SetDisableForcibleLogoff( /* IN */ DisableForcibleLogoff uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetDisableForcibleLogoff", DisableForcibleLogoff)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ProfilePath" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) SetProfilePath( /* IN */ ProfilePath string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetProfilePath", ProfilePath)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="HomeDirectory" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) SetHomeDirectory( /* IN */ HomeDirectory string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetHomeDirectory", HomeDirectory)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="PropertyName" type="string "></param>
// <param name="Value" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) SetPolicyPropertyName( /* IN */ PropertyName string,
	/* IN */ Value bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPolicyPropertyName", PropertyName, Value)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="FallbackPrintDriverType" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) SetFallbackPrintDriverType( /* IN */ FallbackPrintDriverType uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetFallbackPrintDriverType", FallbackPrintDriverType)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="LanaIdDescriptions" type="string []"></param>
// <param name="LanaIdList" type="uint32 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) GetTSLanaIds( /* OUT */ LanaIdList []uint32,
	/* OUT */ LanaIdDescriptions []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetTSLanaIds")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
// <param name="WinstaDriverNames" type="string []"></param>
func (instance *Win32_TerminalServiceSetting) GetWinstationDriverNames( /* OUT */ WinstaDriverNames []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetWinstationDriverNames")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="LanaId" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="WinstaDriverName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) CreateWinstation( /* IN */ Name string,
	/* IN */ WinstaDriverName string,
	/* IN */ LanaId uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CreateWinstation", Name, WinstaDriverName, LanaId)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ServerName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) PingLicenseServer( /* IN */ ServerName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("PingLicenseServer", ServerName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ServerName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="TStoLSConnectivityStatus" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) GetTStoLSConnectivityStatus( /* IN */ ServerName string,
	/* OUT */ TStoLSConnectivityStatus uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetTStoLSConnectivityStatus", ServerName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ServerName" type="string "></param>

// <param name="AccessAllowed" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) CanAccessLicenseServer( /* IN */ ServerName string,
	/* OUT */ AccessAllowed uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CanAccessLicenseServer", ServerName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Count" type="uint32 "></param>
// <param name="LicenseServersList" type="Win32_TSDiscoveredLicenseServer []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) FindLicenseServers( /* OUT */ LicenseServersList []Win32_TSDiscoveredLicenseServer,
	/* OUT */ Count uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("FindLicenseServers")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DaysLeft" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) GetGracePeriodDays( /* OUT */ DaysLeft uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetGracePeriodDays")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Domain" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TerminalServiceSetting) GetDomain( /* OUT */ Domain string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetDomain")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
