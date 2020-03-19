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

// MSFT_NetBranchCacheStatus struct
type MSFT_NetBranchCacheStatus struct {
	*CIM_LogicalElement

	//
	BranchCacheIsEnabled bool

	//
	BranchCacheServiceStartType uint32

	//
	BranchCacheServiceStatus uint32

	//
	ClientConfiguration MSFT_NetBranchCacheClientSettingData

	//
	ContentServerConfiguration MSFT_NetBranchCacheContentServerSettingData

	//
	DataCache MSFT_NetBranchCacheDataCache

	//
	HashCache MSFT_NetBranchCacheHashCache

	//
	HostedCacheServerConfiguration MSFT_NetBranchCacheHostedCacheServerSettingData

	//
	NetworkConfiguration MSFT_NetBranchCacheNetworkSettingData
}

func NewMSFT_NetBranchCacheStatusEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetBranchCacheStatus, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheStatus{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewMSFT_NetBranchCacheStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetBranchCacheStatus, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheStatus{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetBranchCacheIsEnabled sets the value of BranchCacheIsEnabled for the instance
func (instance *MSFT_NetBranchCacheStatus) SetPropertyBranchCacheIsEnabled(value bool) (err error) {
	return instance.SetProperty("BranchCacheIsEnabled", value)
}

// GetBranchCacheIsEnabled gets the value of BranchCacheIsEnabled for the instance
func (instance *MSFT_NetBranchCacheStatus) GetPropertyBranchCacheIsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("BranchCacheIsEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBranchCacheServiceStartType sets the value of BranchCacheServiceStartType for the instance
func (instance *MSFT_NetBranchCacheStatus) SetPropertyBranchCacheServiceStartType(value uint32) (err error) {
	return instance.SetProperty("BranchCacheServiceStartType", value)
}

// GetBranchCacheServiceStartType gets the value of BranchCacheServiceStartType for the instance
func (instance *MSFT_NetBranchCacheStatus) GetPropertyBranchCacheServiceStartType() (value uint32, err error) {
	retValue, err := instance.GetProperty("BranchCacheServiceStartType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBranchCacheServiceStatus sets the value of BranchCacheServiceStatus for the instance
func (instance *MSFT_NetBranchCacheStatus) SetPropertyBranchCacheServiceStatus(value uint32) (err error) {
	return instance.SetProperty("BranchCacheServiceStatus", value)
}

// GetBranchCacheServiceStatus gets the value of BranchCacheServiceStatus for the instance
func (instance *MSFT_NetBranchCacheStatus) GetPropertyBranchCacheServiceStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("BranchCacheServiceStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClientConfiguration sets the value of ClientConfiguration for the instance
func (instance *MSFT_NetBranchCacheStatus) SetPropertyClientConfiguration(value MSFT_NetBranchCacheClientSettingData) (err error) {
	return instance.SetProperty("ClientConfiguration", value)
}

// GetClientConfiguration gets the value of ClientConfiguration for the instance
func (instance *MSFT_NetBranchCacheStatus) GetPropertyClientConfiguration() (value MSFT_NetBranchCacheClientSettingData, err error) {
	retValue, err := instance.GetProperty("ClientConfiguration")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetBranchCacheClientSettingData)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContentServerConfiguration sets the value of ContentServerConfiguration for the instance
func (instance *MSFT_NetBranchCacheStatus) SetPropertyContentServerConfiguration(value MSFT_NetBranchCacheContentServerSettingData) (err error) {
	return instance.SetProperty("ContentServerConfiguration", value)
}

// GetContentServerConfiguration gets the value of ContentServerConfiguration for the instance
func (instance *MSFT_NetBranchCacheStatus) GetPropertyContentServerConfiguration() (value MSFT_NetBranchCacheContentServerSettingData, err error) {
	retValue, err := instance.GetProperty("ContentServerConfiguration")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetBranchCacheContentServerSettingData)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDataCache sets the value of DataCache for the instance
func (instance *MSFT_NetBranchCacheStatus) SetPropertyDataCache(value MSFT_NetBranchCacheDataCache) (err error) {
	return instance.SetProperty("DataCache", value)
}

// GetDataCache gets the value of DataCache for the instance
func (instance *MSFT_NetBranchCacheStatus) GetPropertyDataCache() (value MSFT_NetBranchCacheDataCache, err error) {
	retValue, err := instance.GetProperty("DataCache")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetBranchCacheDataCache)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHashCache sets the value of HashCache for the instance
func (instance *MSFT_NetBranchCacheStatus) SetPropertyHashCache(value MSFT_NetBranchCacheHashCache) (err error) {
	return instance.SetProperty("HashCache", value)
}

// GetHashCache gets the value of HashCache for the instance
func (instance *MSFT_NetBranchCacheStatus) GetPropertyHashCache() (value MSFT_NetBranchCacheHashCache, err error) {
	retValue, err := instance.GetProperty("HashCache")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetBranchCacheHashCache)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheServerConfiguration sets the value of HostedCacheServerConfiguration for the instance
func (instance *MSFT_NetBranchCacheStatus) SetPropertyHostedCacheServerConfiguration(value MSFT_NetBranchCacheHostedCacheServerSettingData) (err error) {
	return instance.SetProperty("HostedCacheServerConfiguration", value)
}

// GetHostedCacheServerConfiguration gets the value of HostedCacheServerConfiguration for the instance
func (instance *MSFT_NetBranchCacheStatus) GetPropertyHostedCacheServerConfiguration() (value MSFT_NetBranchCacheHostedCacheServerSettingData, err error) {
	retValue, err := instance.GetProperty("HostedCacheServerConfiguration")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetBranchCacheHostedCacheServerSettingData)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkConfiguration sets the value of NetworkConfiguration for the instance
func (instance *MSFT_NetBranchCacheStatus) SetPropertyNetworkConfiguration(value MSFT_NetBranchCacheNetworkSettingData) (err error) {
	return instance.SetProperty("NetworkConfiguration", value)
}

// GetNetworkConfiguration gets the value of NetworkConfiguration for the instance
func (instance *MSFT_NetBranchCacheStatus) GetPropertyNetworkConfiguration() (value MSFT_NetBranchCacheNetworkSettingData, err error) {
	retValue, err := instance.GetProperty("NetworkConfiguration")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetBranchCacheNetworkSettingData)
	if !ok {
		// TODO: Set an error
	}
	return
}
