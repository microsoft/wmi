// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.DeliveryOptimization
//////////////////////////////////////////////
package deliveryoptimization

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DeliveryOptimizationExtendedConfig struct
type MSFT_DeliveryOptimizationExtendedConfig struct {
	*MSFT_DOBaseStatus

	//
	BatteryPctToSeed uint32

	//
	BatteryPctToSeedProvider DeliveryOptimizationExtendedConfig_BatteryPctToSeedProvider

	//
	MinTotalDiskSize uint32

	//
	MinTotalDiskSizeProvider DeliveryOptimizationExtendedConfig_MinTotalDiskSizeProvider

	//
	MinTotalRAM uint32

	//
	MinTotalRAMProvider DeliveryOptimizationExtendedConfig_MinTotalRAMProvider

	//
	VpnKeywords string

	//
	VpnKeywordsProvider DeliveryOptimizationExtendedConfig_VpnKeywordsProvider

	//
	VpnPeerCachingAllowed bool

	//
	VpnPeerCachingAllowedProvider DeliveryOptimizationExtendedConfig_VpnPeerCachingAllowedProvider

	//
	WorkingDirectory string

	//
	WorkingDirectoryProvider DeliveryOptimizationExtendedConfig_WorkingDirectoryProvider
}

func NewMSFT_DeliveryOptimizationExtendedConfigEx1(instance *cim.WmiInstance) (newInstance *MSFT_DeliveryOptimizationExtendedConfig, err error) {
	tmp, err := NewMSFT_DOBaseStatusEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DeliveryOptimizationExtendedConfig{
		MSFT_DOBaseStatus: tmp,
	}
	return
}

func NewMSFT_DeliveryOptimizationExtendedConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DeliveryOptimizationExtendedConfig, err error) {
	tmp, err := NewMSFT_DOBaseStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DeliveryOptimizationExtendedConfig{
		MSFT_DOBaseStatus: tmp,
	}
	return
}

// SetBatteryPctToSeed sets the value of BatteryPctToSeed for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyBatteryPctToSeed(value uint32) (err error) {
	return instance.SetProperty("BatteryPctToSeed", value)
}

// GetBatteryPctToSeed gets the value of BatteryPctToSeed for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyBatteryPctToSeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("BatteryPctToSeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBatteryPctToSeedProvider sets the value of BatteryPctToSeedProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyBatteryPctToSeedProvider(value DeliveryOptimizationExtendedConfig_BatteryPctToSeedProvider) (err error) {
	return instance.SetProperty("BatteryPctToSeedProvider", value)
}

// GetBatteryPctToSeedProvider gets the value of BatteryPctToSeedProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyBatteryPctToSeedProvider() (value DeliveryOptimizationExtendedConfig_BatteryPctToSeedProvider, err error) {
	retValue, err := instance.GetProperty("BatteryPctToSeedProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationExtendedConfig_BatteryPctToSeedProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinTotalDiskSize sets the value of MinTotalDiskSize for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyMinTotalDiskSize(value uint32) (err error) {
	return instance.SetProperty("MinTotalDiskSize", value)
}

// GetMinTotalDiskSize gets the value of MinTotalDiskSize for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyMinTotalDiskSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinTotalDiskSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinTotalDiskSizeProvider sets the value of MinTotalDiskSizeProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyMinTotalDiskSizeProvider(value DeliveryOptimizationExtendedConfig_MinTotalDiskSizeProvider) (err error) {
	return instance.SetProperty("MinTotalDiskSizeProvider", value)
}

// GetMinTotalDiskSizeProvider gets the value of MinTotalDiskSizeProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyMinTotalDiskSizeProvider() (value DeliveryOptimizationExtendedConfig_MinTotalDiskSizeProvider, err error) {
	retValue, err := instance.GetProperty("MinTotalDiskSizeProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationExtendedConfig_MinTotalDiskSizeProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinTotalRAM sets the value of MinTotalRAM for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyMinTotalRAM(value uint32) (err error) {
	return instance.SetProperty("MinTotalRAM", value)
}

// GetMinTotalRAM gets the value of MinTotalRAM for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyMinTotalRAM() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinTotalRAM")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinTotalRAMProvider sets the value of MinTotalRAMProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyMinTotalRAMProvider(value DeliveryOptimizationExtendedConfig_MinTotalRAMProvider) (err error) {
	return instance.SetProperty("MinTotalRAMProvider", value)
}

// GetMinTotalRAMProvider gets the value of MinTotalRAMProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyMinTotalRAMProvider() (value DeliveryOptimizationExtendedConfig_MinTotalRAMProvider, err error) {
	retValue, err := instance.GetProperty("MinTotalRAMProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationExtendedConfig_MinTotalRAMProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVpnKeywords sets the value of VpnKeywords for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyVpnKeywords(value string) (err error) {
	return instance.SetProperty("VpnKeywords", value)
}

// GetVpnKeywords gets the value of VpnKeywords for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyVpnKeywords() (value string, err error) {
	retValue, err := instance.GetProperty("VpnKeywords")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVpnKeywordsProvider sets the value of VpnKeywordsProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyVpnKeywordsProvider(value DeliveryOptimizationExtendedConfig_VpnKeywordsProvider) (err error) {
	return instance.SetProperty("VpnKeywordsProvider", value)
}

// GetVpnKeywordsProvider gets the value of VpnKeywordsProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyVpnKeywordsProvider() (value DeliveryOptimizationExtendedConfig_VpnKeywordsProvider, err error) {
	retValue, err := instance.GetProperty("VpnKeywordsProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationExtendedConfig_VpnKeywordsProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVpnPeerCachingAllowed sets the value of VpnPeerCachingAllowed for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyVpnPeerCachingAllowed(value bool) (err error) {
	return instance.SetProperty("VpnPeerCachingAllowed", value)
}

// GetVpnPeerCachingAllowed gets the value of VpnPeerCachingAllowed for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyVpnPeerCachingAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("VpnPeerCachingAllowed")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVpnPeerCachingAllowedProvider sets the value of VpnPeerCachingAllowedProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyVpnPeerCachingAllowedProvider(value DeliveryOptimizationExtendedConfig_VpnPeerCachingAllowedProvider) (err error) {
	return instance.SetProperty("VpnPeerCachingAllowedProvider", value)
}

// GetVpnPeerCachingAllowedProvider gets the value of VpnPeerCachingAllowedProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyVpnPeerCachingAllowedProvider() (value DeliveryOptimizationExtendedConfig_VpnPeerCachingAllowedProvider, err error) {
	retValue, err := instance.GetProperty("VpnPeerCachingAllowedProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationExtendedConfig_VpnPeerCachingAllowedProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkingDirectory sets the value of WorkingDirectory for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyWorkingDirectory(value string) (err error) {
	return instance.SetProperty("WorkingDirectory", value)
}

// GetWorkingDirectory gets the value of WorkingDirectory for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyWorkingDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("WorkingDirectory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkingDirectoryProvider sets the value of WorkingDirectoryProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyWorkingDirectoryProvider(value DeliveryOptimizationExtendedConfig_WorkingDirectoryProvider) (err error) {
	return instance.SetProperty("WorkingDirectoryProvider", value)
}

// GetWorkingDirectoryProvider gets the value of WorkingDirectoryProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyWorkingDirectoryProvider() (value DeliveryOptimizationExtendedConfig_WorkingDirectoryProvider, err error) {
	retValue, err := instance.GetProperty("WorkingDirectoryProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationExtendedConfig_WorkingDirectoryProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}
