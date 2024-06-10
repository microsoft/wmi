// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.DeliveryOptimization
//
// ////////////////////////////////////////////
package deliveryoptimization

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
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
	SetHoursToLimitDownloadBackground string

	//
	SetHoursToLimitDownloadBackgroundProvider DeliveryOptimizationExtendedConfig_SetHoursToLimitDownloadBackgroundProvider

	//
	SetHoursToLimitDownloadForeground string

	//
	SetHoursToLimitDownloadForegroundProvider DeliveryOptimizationExtendedConfig_SetHoursToLimitDownloadForegroundProvider

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
	return instance.SetProperty("BatteryPctToSeed", (value))
}

// GetBatteryPctToSeed gets the value of BatteryPctToSeed for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyBatteryPctToSeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("BatteryPctToSeed")
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

// SetBatteryPctToSeedProvider sets the value of BatteryPctToSeedProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyBatteryPctToSeedProvider(value DeliveryOptimizationExtendedConfig_BatteryPctToSeedProvider) (err error) {
	return instance.SetProperty("BatteryPctToSeedProvider", (value))
}

// GetBatteryPctToSeedProvider gets the value of BatteryPctToSeedProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyBatteryPctToSeedProvider() (value DeliveryOptimizationExtendedConfig_BatteryPctToSeedProvider, err error) {
	retValue, err := instance.GetProperty("BatteryPctToSeedProvider")
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

	value = DeliveryOptimizationExtendedConfig_BatteryPctToSeedProvider(valuetmp)

	return
}

// SetMinTotalDiskSize sets the value of MinTotalDiskSize for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyMinTotalDiskSize(value uint32) (err error) {
	return instance.SetProperty("MinTotalDiskSize", (value))
}

// GetMinTotalDiskSize gets the value of MinTotalDiskSize for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyMinTotalDiskSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinTotalDiskSize")
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

// SetMinTotalDiskSizeProvider sets the value of MinTotalDiskSizeProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyMinTotalDiskSizeProvider(value DeliveryOptimizationExtendedConfig_MinTotalDiskSizeProvider) (err error) {
	return instance.SetProperty("MinTotalDiskSizeProvider", (value))
}

// GetMinTotalDiskSizeProvider gets the value of MinTotalDiskSizeProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyMinTotalDiskSizeProvider() (value DeliveryOptimizationExtendedConfig_MinTotalDiskSizeProvider, err error) {
	retValue, err := instance.GetProperty("MinTotalDiskSizeProvider")
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

	value = DeliveryOptimizationExtendedConfig_MinTotalDiskSizeProvider(valuetmp)

	return
}

// SetMinTotalRAM sets the value of MinTotalRAM for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyMinTotalRAM(value uint32) (err error) {
	return instance.SetProperty("MinTotalRAM", (value))
}

// GetMinTotalRAM gets the value of MinTotalRAM for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyMinTotalRAM() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinTotalRAM")
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

// SetMinTotalRAMProvider sets the value of MinTotalRAMProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyMinTotalRAMProvider(value DeliveryOptimizationExtendedConfig_MinTotalRAMProvider) (err error) {
	return instance.SetProperty("MinTotalRAMProvider", (value))
}

// GetMinTotalRAMProvider gets the value of MinTotalRAMProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyMinTotalRAMProvider() (value DeliveryOptimizationExtendedConfig_MinTotalRAMProvider, err error) {
	retValue, err := instance.GetProperty("MinTotalRAMProvider")
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

	value = DeliveryOptimizationExtendedConfig_MinTotalRAMProvider(valuetmp)

	return
}

// SetSetHoursToLimitDownloadBackground sets the value of SetHoursToLimitDownloadBackground for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertySetHoursToLimitDownloadBackground(value string) (err error) {
	return instance.SetProperty("SetHoursToLimitDownloadBackground", (value))
}

// GetSetHoursToLimitDownloadBackground gets the value of SetHoursToLimitDownloadBackground for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertySetHoursToLimitDownloadBackground() (value string, err error) {
	retValue, err := instance.GetProperty("SetHoursToLimitDownloadBackground")
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

// SetSetHoursToLimitDownloadBackgroundProvider sets the value of SetHoursToLimitDownloadBackgroundProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertySetHoursToLimitDownloadBackgroundProvider(value DeliveryOptimizationExtendedConfig_SetHoursToLimitDownloadBackgroundProvider) (err error) {
	return instance.SetProperty("SetHoursToLimitDownloadBackgroundProvider", (value))
}

// GetSetHoursToLimitDownloadBackgroundProvider gets the value of SetHoursToLimitDownloadBackgroundProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertySetHoursToLimitDownloadBackgroundProvider() (value DeliveryOptimizationExtendedConfig_SetHoursToLimitDownloadBackgroundProvider, err error) {
	retValue, err := instance.GetProperty("SetHoursToLimitDownloadBackgroundProvider")
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

	value = DeliveryOptimizationExtendedConfig_SetHoursToLimitDownloadBackgroundProvider(valuetmp)

	return
}

// SetSetHoursToLimitDownloadForeground sets the value of SetHoursToLimitDownloadForeground for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertySetHoursToLimitDownloadForeground(value string) (err error) {
	return instance.SetProperty("SetHoursToLimitDownloadForeground", (value))
}

// GetSetHoursToLimitDownloadForeground gets the value of SetHoursToLimitDownloadForeground for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertySetHoursToLimitDownloadForeground() (value string, err error) {
	retValue, err := instance.GetProperty("SetHoursToLimitDownloadForeground")
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

// SetSetHoursToLimitDownloadForegroundProvider sets the value of SetHoursToLimitDownloadForegroundProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertySetHoursToLimitDownloadForegroundProvider(value DeliveryOptimizationExtendedConfig_SetHoursToLimitDownloadForegroundProvider) (err error) {
	return instance.SetProperty("SetHoursToLimitDownloadForegroundProvider", (value))
}

// GetSetHoursToLimitDownloadForegroundProvider gets the value of SetHoursToLimitDownloadForegroundProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertySetHoursToLimitDownloadForegroundProvider() (value DeliveryOptimizationExtendedConfig_SetHoursToLimitDownloadForegroundProvider, err error) {
	retValue, err := instance.GetProperty("SetHoursToLimitDownloadForegroundProvider")
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

	value = DeliveryOptimizationExtendedConfig_SetHoursToLimitDownloadForegroundProvider(valuetmp)

	return
}

// SetVpnKeywords sets the value of VpnKeywords for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyVpnKeywords(value string) (err error) {
	return instance.SetProperty("VpnKeywords", (value))
}

// GetVpnKeywords gets the value of VpnKeywords for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyVpnKeywords() (value string, err error) {
	retValue, err := instance.GetProperty("VpnKeywords")
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

// SetVpnKeywordsProvider sets the value of VpnKeywordsProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyVpnKeywordsProvider(value DeliveryOptimizationExtendedConfig_VpnKeywordsProvider) (err error) {
	return instance.SetProperty("VpnKeywordsProvider", (value))
}

// GetVpnKeywordsProvider gets the value of VpnKeywordsProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyVpnKeywordsProvider() (value DeliveryOptimizationExtendedConfig_VpnKeywordsProvider, err error) {
	retValue, err := instance.GetProperty("VpnKeywordsProvider")
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

	value = DeliveryOptimizationExtendedConfig_VpnKeywordsProvider(valuetmp)

	return
}

// SetVpnPeerCachingAllowed sets the value of VpnPeerCachingAllowed for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyVpnPeerCachingAllowed(value bool) (err error) {
	return instance.SetProperty("VpnPeerCachingAllowed", (value))
}

// GetVpnPeerCachingAllowed gets the value of VpnPeerCachingAllowed for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyVpnPeerCachingAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("VpnPeerCachingAllowed")
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

// SetVpnPeerCachingAllowedProvider sets the value of VpnPeerCachingAllowedProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyVpnPeerCachingAllowedProvider(value DeliveryOptimizationExtendedConfig_VpnPeerCachingAllowedProvider) (err error) {
	return instance.SetProperty("VpnPeerCachingAllowedProvider", (value))
}

// GetVpnPeerCachingAllowedProvider gets the value of VpnPeerCachingAllowedProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyVpnPeerCachingAllowedProvider() (value DeliveryOptimizationExtendedConfig_VpnPeerCachingAllowedProvider, err error) {
	retValue, err := instance.GetProperty("VpnPeerCachingAllowedProvider")
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

	value = DeliveryOptimizationExtendedConfig_VpnPeerCachingAllowedProvider(valuetmp)

	return
}

// SetWorkingDirectory sets the value of WorkingDirectory for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyWorkingDirectory(value string) (err error) {
	return instance.SetProperty("WorkingDirectory", (value))
}

// GetWorkingDirectory gets the value of WorkingDirectory for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyWorkingDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("WorkingDirectory")
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

// SetWorkingDirectoryProvider sets the value of WorkingDirectoryProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) SetPropertyWorkingDirectoryProvider(value DeliveryOptimizationExtendedConfig_WorkingDirectoryProvider) (err error) {
	return instance.SetProperty("WorkingDirectoryProvider", (value))
}

// GetWorkingDirectoryProvider gets the value of WorkingDirectoryProvider for the instance
func (instance *MSFT_DeliveryOptimizationExtendedConfig) GetPropertyWorkingDirectoryProvider() (value DeliveryOptimizationExtendedConfig_WorkingDirectoryProvider, err error) {
	retValue, err := instance.GetProperty("WorkingDirectoryProvider")
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

	value = DeliveryOptimizationExtendedConfig_WorkingDirectoryProvider(valuetmp)

	return
}
