// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetBranchCacheClientSettingData struct
type MSFT_NetBranchCacheClientSettingData struct {
	*MSFT_NetBranchCacheSettingData

	//
	CurrentClientMode uint32

	//
	DistributedCachingIsEnabled bool

	//
	HostedCacheDiscoveryEnabled bool

	//
	HostedCacheServerList []string

	//
	HostedCacheVersion uint32

	//
	MinimumSmbLatencyInMilliseconds uint32

	//
	PreferredContentInformationVersion uint32

	//
	ServeDistributedCachingPeersOnBatteryPower bool
}

func NewMSFT_NetBranchCacheClientSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetBranchCacheClientSettingData, err error) {
	tmp, err := NewMSFT_NetBranchCacheSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheClientSettingData{
		MSFT_NetBranchCacheSettingData: tmp,
	}
	return
}

func NewMSFT_NetBranchCacheClientSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetBranchCacheClientSettingData, err error) {
	tmp, err := NewMSFT_NetBranchCacheSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheClientSettingData{
		MSFT_NetBranchCacheSettingData: tmp,
	}
	return
}

// SetCurrentClientMode sets the value of CurrentClientMode for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) SetPropertyCurrentClientMode(value uint32) (err error) {
	return instance.SetProperty("CurrentClientMode", value)
}

// GetCurrentClientMode gets the value of CurrentClientMode for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) GetPropertyCurrentClientMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentClientMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDistributedCachingIsEnabled sets the value of DistributedCachingIsEnabled for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) SetPropertyDistributedCachingIsEnabled(value bool) (err error) {
	return instance.SetProperty("DistributedCachingIsEnabled", value)
}

// GetDistributedCachingIsEnabled gets the value of DistributedCachingIsEnabled for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) GetPropertyDistributedCachingIsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DistributedCachingIsEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheDiscoveryEnabled sets the value of HostedCacheDiscoveryEnabled for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) SetPropertyHostedCacheDiscoveryEnabled(value bool) (err error) {
	return instance.SetProperty("HostedCacheDiscoveryEnabled", value)
}

// GetHostedCacheDiscoveryEnabled gets the value of HostedCacheDiscoveryEnabled for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) GetPropertyHostedCacheDiscoveryEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("HostedCacheDiscoveryEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheServerList sets the value of HostedCacheServerList for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) SetPropertyHostedCacheServerList(value []string) (err error) {
	return instance.SetProperty("HostedCacheServerList", value)
}

// GetHostedCacheServerList gets the value of HostedCacheServerList for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) GetPropertyHostedCacheServerList() (value []string, err error) {
	retValue, err := instance.GetProperty("HostedCacheServerList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheVersion sets the value of HostedCacheVersion for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) SetPropertyHostedCacheVersion(value uint32) (err error) {
	return instance.SetProperty("HostedCacheVersion", value)
}

// GetHostedCacheVersion gets the value of HostedCacheVersion for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) GetPropertyHostedCacheVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("HostedCacheVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinimumSmbLatencyInMilliseconds sets the value of MinimumSmbLatencyInMilliseconds for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) SetPropertyMinimumSmbLatencyInMilliseconds(value uint32) (err error) {
	return instance.SetProperty("MinimumSmbLatencyInMilliseconds", value)
}

// GetMinimumSmbLatencyInMilliseconds gets the value of MinimumSmbLatencyInMilliseconds for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) GetPropertyMinimumSmbLatencyInMilliseconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinimumSmbLatencyInMilliseconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreferredContentInformationVersion sets the value of PreferredContentInformationVersion for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) SetPropertyPreferredContentInformationVersion(value uint32) (err error) {
	return instance.SetProperty("PreferredContentInformationVersion", value)
}

// GetPreferredContentInformationVersion gets the value of PreferredContentInformationVersion for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) GetPropertyPreferredContentInformationVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("PreferredContentInformationVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServeDistributedCachingPeersOnBatteryPower sets the value of ServeDistributedCachingPeersOnBatteryPower for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) SetPropertyServeDistributedCachingPeersOnBatteryPower(value bool) (err error) {
	return instance.SetProperty("ServeDistributedCachingPeersOnBatteryPower", value)
}

// GetServeDistributedCachingPeersOnBatteryPower gets the value of ServeDistributedCachingPeersOnBatteryPower for the instance
func (instance *MSFT_NetBranchCacheClientSettingData) GetPropertyServeDistributedCachingPeersOnBatteryPower() (value bool, err error) {
	retValue, err := instance.GetProperty("ServeDistributedCachingPeersOnBatteryPower")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
