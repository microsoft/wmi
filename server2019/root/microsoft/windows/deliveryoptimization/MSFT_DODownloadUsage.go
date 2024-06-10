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

// MSFT_DODownloadUsage struct
type MSFT_DODownloadUsage struct {
	*MSFT_DOUsage

	//
	BackgroundRatePct uint8

	//
	ForegroundRatePct uint8

	//
	MonthlyBkRateBps uint64

	//
	MonthlyCacheHostBytes uint64

	//
	MonthlyCdnBytes uint64

	//
	MonthlyFrRateBps uint64

	//
	MonthlyGroupBytes uint64

	//
	MonthlyLinkLocalBytes uint64

	//
	NormalDownloads uint32

	//
	NormalDownloadsPending uint32

	//
	PriorityDownloads uint32

	//
	PriorityDownloadsPending uint32
}

func NewMSFT_DODownloadUsageEx1(instance *cim.WmiInstance) (newInstance *MSFT_DODownloadUsage, err error) {
	tmp, err := NewMSFT_DOUsageEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DODownloadUsage{
		MSFT_DOUsage: tmp,
	}
	return
}

func NewMSFT_DODownloadUsageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DODownloadUsage, err error) {
	tmp, err := NewMSFT_DOUsageEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DODownloadUsage{
		MSFT_DOUsage: tmp,
	}
	return
}

// SetBackgroundRatePct sets the value of BackgroundRatePct for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyBackgroundRatePct(value uint8) (err error) {
	return instance.SetProperty("BackgroundRatePct", value)
}

// GetBackgroundRatePct gets the value of BackgroundRatePct for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyBackgroundRatePct() (value uint8, err error) {
	retValue, err := instance.GetProperty("BackgroundRatePct")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetForegroundRatePct sets the value of ForegroundRatePct for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyForegroundRatePct(value uint8) (err error) {
	return instance.SetProperty("ForegroundRatePct", value)
}

// GetForegroundRatePct gets the value of ForegroundRatePct for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyForegroundRatePct() (value uint8, err error) {
	retValue, err := instance.GetProperty("ForegroundRatePct")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonthlyBkRateBps sets the value of MonthlyBkRateBps for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyMonthlyBkRateBps(value uint64) (err error) {
	return instance.SetProperty("MonthlyBkRateBps", value)
}

// GetMonthlyBkRateBps gets the value of MonthlyBkRateBps for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyMonthlyBkRateBps() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyBkRateBps")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonthlyCacheHostBytes sets the value of MonthlyCacheHostBytes for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyMonthlyCacheHostBytes(value uint64) (err error) {
	return instance.SetProperty("MonthlyCacheHostBytes", value)
}

// GetMonthlyCacheHostBytes gets the value of MonthlyCacheHostBytes for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyMonthlyCacheHostBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyCacheHostBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonthlyCdnBytes sets the value of MonthlyCdnBytes for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyMonthlyCdnBytes(value uint64) (err error) {
	return instance.SetProperty("MonthlyCdnBytes", value)
}

// GetMonthlyCdnBytes gets the value of MonthlyCdnBytes for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyMonthlyCdnBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyCdnBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonthlyFrRateBps sets the value of MonthlyFrRateBps for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyMonthlyFrRateBps(value uint64) (err error) {
	return instance.SetProperty("MonthlyFrRateBps", value)
}

// GetMonthlyFrRateBps gets the value of MonthlyFrRateBps for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyMonthlyFrRateBps() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyFrRateBps")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonthlyGroupBytes sets the value of MonthlyGroupBytes for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyMonthlyGroupBytes(value uint64) (err error) {
	return instance.SetProperty("MonthlyGroupBytes", value)
}

// GetMonthlyGroupBytes gets the value of MonthlyGroupBytes for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyMonthlyGroupBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyGroupBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonthlyLinkLocalBytes sets the value of MonthlyLinkLocalBytes for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyMonthlyLinkLocalBytes(value uint64) (err error) {
	return instance.SetProperty("MonthlyLinkLocalBytes", value)
}

// GetMonthlyLinkLocalBytes gets the value of MonthlyLinkLocalBytes for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyMonthlyLinkLocalBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyLinkLocalBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNormalDownloads sets the value of NormalDownloads for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyNormalDownloads(value uint32) (err error) {
	return instance.SetProperty("NormalDownloads", value)
}

// GetNormalDownloads gets the value of NormalDownloads for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyNormalDownloads() (value uint32, err error) {
	retValue, err := instance.GetProperty("NormalDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNormalDownloadsPending sets the value of NormalDownloadsPending for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyNormalDownloadsPending(value uint32) (err error) {
	return instance.SetProperty("NormalDownloadsPending", value)
}

// GetNormalDownloadsPending gets the value of NormalDownloadsPending for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyNormalDownloadsPending() (value uint32, err error) {
	retValue, err := instance.GetProperty("NormalDownloadsPending")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriorityDownloads sets the value of PriorityDownloads for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyPriorityDownloads(value uint32) (err error) {
	return instance.SetProperty("PriorityDownloads", value)
}

// GetPriorityDownloads gets the value of PriorityDownloads for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyPriorityDownloads() (value uint32, err error) {
	retValue, err := instance.GetProperty("PriorityDownloads")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriorityDownloadsPending sets the value of PriorityDownloadsPending for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyPriorityDownloadsPending(value uint32) (err error) {
	return instance.SetProperty("PriorityDownloadsPending", value)
}

// GetPriorityDownloadsPending gets the value of PriorityDownloadsPending for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyPriorityDownloadsPending() (value uint32, err error) {
	retValue, err := instance.GetProperty("PriorityDownloadsPending")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
