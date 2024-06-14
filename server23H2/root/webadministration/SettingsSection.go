// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SettingsSection struct
type SettingsSection struct {
	*ConfigurationSection

	//
	HttpWebRequest HttpWebRequestSettings

	//
	Ipv6 Ipv6Settings

	//
	PerformanceCounters PerformanceCountersSettings

	//
	ServicePointManager ServicePointManagerSettings

	//
	Socket SocketSettings

	//
	WebProxyScript WebProxyScriptSettings
}

func NewSettingsSectionEx1(instance *cim.WmiInstance) (newInstance *SettingsSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SettingsSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewSettingsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SettingsSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SettingsSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetHttpWebRequest sets the value of HttpWebRequest for the instance
func (instance *SettingsSection) SetPropertyHttpWebRequest(value HttpWebRequestSettings) (err error) {
	return instance.SetProperty("HttpWebRequest", (value))
}

// GetHttpWebRequest gets the value of HttpWebRequest for the instance
func (instance *SettingsSection) GetPropertyHttpWebRequest() (value HttpWebRequestSettings, err error) {
	retValue, err := instance.GetProperty("HttpWebRequest")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(HttpWebRequestSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " HttpWebRequestSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = HttpWebRequestSettings(valuetmp)

	return
}

// SetIpv6 sets the value of Ipv6 for the instance
func (instance *SettingsSection) SetPropertyIpv6(value Ipv6Settings) (err error) {
	return instance.SetProperty("Ipv6", (value))
}

// GetIpv6 gets the value of Ipv6 for the instance
func (instance *SettingsSection) GetPropertyIpv6() (value Ipv6Settings, err error) {
	retValue, err := instance.GetProperty("Ipv6")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Ipv6Settings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Ipv6Settings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Ipv6Settings(valuetmp)

	return
}

// SetPerformanceCounters sets the value of PerformanceCounters for the instance
func (instance *SettingsSection) SetPropertyPerformanceCounters(value PerformanceCountersSettings) (err error) {
	return instance.SetProperty("PerformanceCounters", (value))
}

// GetPerformanceCounters gets the value of PerformanceCounters for the instance
func (instance *SettingsSection) GetPropertyPerformanceCounters() (value PerformanceCountersSettings, err error) {
	retValue, err := instance.GetProperty("PerformanceCounters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(PerformanceCountersSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " PerformanceCountersSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = PerformanceCountersSettings(valuetmp)

	return
}

// SetServicePointManager sets the value of ServicePointManager for the instance
func (instance *SettingsSection) SetPropertyServicePointManager(value ServicePointManagerSettings) (err error) {
	return instance.SetProperty("ServicePointManager", (value))
}

// GetServicePointManager gets the value of ServicePointManager for the instance
func (instance *SettingsSection) GetPropertyServicePointManager() (value ServicePointManagerSettings, err error) {
	retValue, err := instance.GetProperty("ServicePointManager")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ServicePointManagerSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ServicePointManagerSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ServicePointManagerSettings(valuetmp)

	return
}

// SetSocket sets the value of Socket for the instance
func (instance *SettingsSection) SetPropertySocket(value SocketSettings) (err error) {
	return instance.SetProperty("Socket", (value))
}

// GetSocket gets the value of Socket for the instance
func (instance *SettingsSection) GetPropertySocket() (value SocketSettings, err error) {
	retValue, err := instance.GetProperty("Socket")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(SocketSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " SocketSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = SocketSettings(valuetmp)

	return
}

// SetWebProxyScript sets the value of WebProxyScript for the instance
func (instance *SettingsSection) SetPropertyWebProxyScript(value WebProxyScriptSettings) (err error) {
	return instance.SetProperty("WebProxyScript", (value))
}

// GetWebProxyScript gets the value of WebProxyScript for the instance
func (instance *SettingsSection) GetPropertyWebProxyScript() (value WebProxyScriptSettings, err error) {
	retValue, err := instance.GetProperty("WebProxyScript")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(WebProxyScriptSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " WebProxyScriptSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WebProxyScriptSettings(valuetmp)

	return
}
