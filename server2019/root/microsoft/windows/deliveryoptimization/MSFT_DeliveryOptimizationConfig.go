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

// MSFT_DeliveryOptimizationConfig struct
type MSFT_DeliveryOptimizationConfig struct {
	*MSFT_DOBaseStatus

	//
	DownBackLimitBps uint32

	//
	DownBackLimitBpsProvider DeliveryOptimizationConfig_DownBackLimitBpsProvider

	//
	DownBackLimitPct uint8

	//
	DownBackLimitPctProvider DeliveryOptimizationConfig_DownBackLimitPctProvider

	//
	DownForeLimitBps uint32

	//
	DownForeLimitBpsProvider DeliveryOptimizationConfig_DownForeLimitBpsProvider

	//
	DownForeLimitPct uint8

	//
	DownForeLimitPctProvider DeliveryOptimizationConfig_DownForeLimitPctProvider

	// 13
	DownloadMode DeliveryOptimizationConfig_DownloadMode

	//
	DownloadModeProvider DeliveryOptimizationConfig_DownloadModeProvider

	//
	MaxUploadRatePct uint8

	//
	MaxUploadRateProvider DeliveryOptimizationConfig_MaxUploadRateProvider

	//
	UpLimitMonthlyGB float64

	//
	UpLimitMonthlyGBProvider DeliveryOptimizationConfig_UpLimitMonthlyGBProvider
}

func NewMSFT_DeliveryOptimizationConfigEx1(instance *cim.WmiInstance) (newInstance *MSFT_DeliveryOptimizationConfig, err error) {
	tmp, err := NewMSFT_DOBaseStatusEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DeliveryOptimizationConfig{
		MSFT_DOBaseStatus: tmp,
	}
	return
}

func NewMSFT_DeliveryOptimizationConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DeliveryOptimizationConfig, err error) {
	tmp, err := NewMSFT_DOBaseStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DeliveryOptimizationConfig{
		MSFT_DOBaseStatus: tmp,
	}
	return
}

// SetDownBackLimitBps sets the value of DownBackLimitBps for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyDownBackLimitBps(value uint32) (err error) {
	return instance.SetProperty("DownBackLimitBps", value)
}

// GetDownBackLimitBps gets the value of DownBackLimitBps for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyDownBackLimitBps() (value uint32, err error) {
	retValue, err := instance.GetProperty("DownBackLimitBps")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDownBackLimitBpsProvider sets the value of DownBackLimitBpsProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyDownBackLimitBpsProvider(value DeliveryOptimizationConfig_DownBackLimitBpsProvider) (err error) {
	return instance.SetProperty("DownBackLimitBpsProvider", value)
}

// GetDownBackLimitBpsProvider gets the value of DownBackLimitBpsProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyDownBackLimitBpsProvider() (value DeliveryOptimizationConfig_DownBackLimitBpsProvider, err error) {
	retValue, err := instance.GetProperty("DownBackLimitBpsProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationConfig_DownBackLimitBpsProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDownBackLimitPct sets the value of DownBackLimitPct for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyDownBackLimitPct(value uint8) (err error) {
	return instance.SetProperty("DownBackLimitPct", value)
}

// GetDownBackLimitPct gets the value of DownBackLimitPct for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyDownBackLimitPct() (value uint8, err error) {
	retValue, err := instance.GetProperty("DownBackLimitPct")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDownBackLimitPctProvider sets the value of DownBackLimitPctProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyDownBackLimitPctProvider(value DeliveryOptimizationConfig_DownBackLimitPctProvider) (err error) {
	return instance.SetProperty("DownBackLimitPctProvider", value)
}

// GetDownBackLimitPctProvider gets the value of DownBackLimitPctProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyDownBackLimitPctProvider() (value DeliveryOptimizationConfig_DownBackLimitPctProvider, err error) {
	retValue, err := instance.GetProperty("DownBackLimitPctProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationConfig_DownBackLimitPctProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDownForeLimitBps sets the value of DownForeLimitBps for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyDownForeLimitBps(value uint32) (err error) {
	return instance.SetProperty("DownForeLimitBps", value)
}

// GetDownForeLimitBps gets the value of DownForeLimitBps for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyDownForeLimitBps() (value uint32, err error) {
	retValue, err := instance.GetProperty("DownForeLimitBps")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDownForeLimitBpsProvider sets the value of DownForeLimitBpsProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyDownForeLimitBpsProvider(value DeliveryOptimizationConfig_DownForeLimitBpsProvider) (err error) {
	return instance.SetProperty("DownForeLimitBpsProvider", value)
}

// GetDownForeLimitBpsProvider gets the value of DownForeLimitBpsProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyDownForeLimitBpsProvider() (value DeliveryOptimizationConfig_DownForeLimitBpsProvider, err error) {
	retValue, err := instance.GetProperty("DownForeLimitBpsProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationConfig_DownForeLimitBpsProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDownForeLimitPct sets the value of DownForeLimitPct for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyDownForeLimitPct(value uint8) (err error) {
	return instance.SetProperty("DownForeLimitPct", value)
}

// GetDownForeLimitPct gets the value of DownForeLimitPct for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyDownForeLimitPct() (value uint8, err error) {
	retValue, err := instance.GetProperty("DownForeLimitPct")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDownForeLimitPctProvider sets the value of DownForeLimitPctProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyDownForeLimitPctProvider(value DeliveryOptimizationConfig_DownForeLimitPctProvider) (err error) {
	return instance.SetProperty("DownForeLimitPctProvider", value)
}

// GetDownForeLimitPctProvider gets the value of DownForeLimitPctProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyDownForeLimitPctProvider() (value DeliveryOptimizationConfig_DownForeLimitPctProvider, err error) {
	retValue, err := instance.GetProperty("DownForeLimitPctProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationConfig_DownForeLimitPctProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDownloadMode sets the value of DownloadMode for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyDownloadMode(value DeliveryOptimizationConfig_DownloadMode) (err error) {
	return instance.SetProperty("DownloadMode", value)
}

// GetDownloadMode gets the value of DownloadMode for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyDownloadMode() (value DeliveryOptimizationConfig_DownloadMode, err error) {
	retValue, err := instance.GetProperty("DownloadMode")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationConfig_DownloadMode)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDownloadModeProvider sets the value of DownloadModeProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyDownloadModeProvider(value DeliveryOptimizationConfig_DownloadModeProvider) (err error) {
	return instance.SetProperty("DownloadModeProvider", value)
}

// GetDownloadModeProvider gets the value of DownloadModeProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyDownloadModeProvider() (value DeliveryOptimizationConfig_DownloadModeProvider, err error) {
	retValue, err := instance.GetProperty("DownloadModeProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationConfig_DownloadModeProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxUploadRatePct sets the value of MaxUploadRatePct for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyMaxUploadRatePct(value uint8) (err error) {
	return instance.SetProperty("MaxUploadRatePct", value)
}

// GetMaxUploadRatePct gets the value of MaxUploadRatePct for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyMaxUploadRatePct() (value uint8, err error) {
	retValue, err := instance.GetProperty("MaxUploadRatePct")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxUploadRateProvider sets the value of MaxUploadRateProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyMaxUploadRateProvider(value DeliveryOptimizationConfig_MaxUploadRateProvider) (err error) {
	return instance.SetProperty("MaxUploadRateProvider", value)
}

// GetMaxUploadRateProvider gets the value of MaxUploadRateProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyMaxUploadRateProvider() (value DeliveryOptimizationConfig_MaxUploadRateProvider, err error) {
	retValue, err := instance.GetProperty("MaxUploadRateProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationConfig_MaxUploadRateProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpLimitMonthlyGB sets the value of UpLimitMonthlyGB for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyUpLimitMonthlyGB(value float64) (err error) {
	return instance.SetProperty("UpLimitMonthlyGB", value)
}

// GetUpLimitMonthlyGB gets the value of UpLimitMonthlyGB for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyUpLimitMonthlyGB() (value float64, err error) {
	retValue, err := instance.GetProperty("UpLimitMonthlyGB")
	if err != nil {
		return
	}
	value, ok := retValue.(float64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpLimitMonthlyGBProvider sets the value of UpLimitMonthlyGBProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) SetPropertyUpLimitMonthlyGBProvider(value DeliveryOptimizationConfig_UpLimitMonthlyGBProvider) (err error) {
	return instance.SetProperty("UpLimitMonthlyGBProvider", value)
}

// GetUpLimitMonthlyGBProvider gets the value of UpLimitMonthlyGBProvider for the instance
func (instance *MSFT_DeliveryOptimizationConfig) GetPropertyUpLimitMonthlyGBProvider() (value DeliveryOptimizationConfig_UpLimitMonthlyGBProvider, err error) {
	retValue, err := instance.GetProperty("UpLimitMonthlyGBProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationConfig_UpLimitMonthlyGBProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}
