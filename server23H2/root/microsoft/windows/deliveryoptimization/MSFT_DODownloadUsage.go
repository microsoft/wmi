// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.DeliveryOptimization
//////////////////////////////////////////////
package deliveryoptimization

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
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
	return instance.SetProperty("BackgroundRatePct", (value))
}

// GetBackgroundRatePct gets the value of BackgroundRatePct for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyBackgroundRatePct() (value uint8, err error) {
	retValue, err := instance.GetProperty("BackgroundRatePct")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetForegroundRatePct sets the value of ForegroundRatePct for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyForegroundRatePct(value uint8) (err error) {
	return instance.SetProperty("ForegroundRatePct", (value))
}

// GetForegroundRatePct gets the value of ForegroundRatePct for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyForegroundRatePct() (value uint8, err error) {
	retValue, err := instance.GetProperty("ForegroundRatePct")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetMonthlyBkRateBps sets the value of MonthlyBkRateBps for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyMonthlyBkRateBps(value uint64) (err error) {
	return instance.SetProperty("MonthlyBkRateBps", (value))
}

// GetMonthlyBkRateBps gets the value of MonthlyBkRateBps for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyMonthlyBkRateBps() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyBkRateBps")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMonthlyCacheHostBytes sets the value of MonthlyCacheHostBytes for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyMonthlyCacheHostBytes(value uint64) (err error) {
	return instance.SetProperty("MonthlyCacheHostBytes", (value))
}

// GetMonthlyCacheHostBytes gets the value of MonthlyCacheHostBytes for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyMonthlyCacheHostBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyCacheHostBytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMonthlyCdnBytes sets the value of MonthlyCdnBytes for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyMonthlyCdnBytes(value uint64) (err error) {
	return instance.SetProperty("MonthlyCdnBytes", (value))
}

// GetMonthlyCdnBytes gets the value of MonthlyCdnBytes for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyMonthlyCdnBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyCdnBytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMonthlyFrRateBps sets the value of MonthlyFrRateBps for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyMonthlyFrRateBps(value uint64) (err error) {
	return instance.SetProperty("MonthlyFrRateBps", (value))
}

// GetMonthlyFrRateBps gets the value of MonthlyFrRateBps for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyMonthlyFrRateBps() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyFrRateBps")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetNormalDownloads sets the value of NormalDownloads for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyNormalDownloads(value uint32) (err error) {
	return instance.SetProperty("NormalDownloads", (value))
}

// GetNormalDownloads gets the value of NormalDownloads for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyNormalDownloads() (value uint32, err error) {
	retValue, err := instance.GetProperty("NormalDownloads")
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

// SetNormalDownloadsPending sets the value of NormalDownloadsPending for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyNormalDownloadsPending(value uint32) (err error) {
	return instance.SetProperty("NormalDownloadsPending", (value))
}

// GetNormalDownloadsPending gets the value of NormalDownloadsPending for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyNormalDownloadsPending() (value uint32, err error) {
	retValue, err := instance.GetProperty("NormalDownloadsPending")
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

// SetPriorityDownloads sets the value of PriorityDownloads for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyPriorityDownloads(value uint32) (err error) {
	return instance.SetProperty("PriorityDownloads", (value))
}

// GetPriorityDownloads gets the value of PriorityDownloads for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyPriorityDownloads() (value uint32, err error) {
	retValue, err := instance.GetProperty("PriorityDownloads")
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

// SetPriorityDownloadsPending sets the value of PriorityDownloadsPending for the instance
func (instance *MSFT_DODownloadUsage) SetPropertyPriorityDownloadsPending(value uint32) (err error) {
	return instance.SetProperty("PriorityDownloadsPending", (value))
}

// GetPriorityDownloadsPending gets the value of PriorityDownloadsPending for the instance
func (instance *MSFT_DODownloadUsage) GetPropertyPriorityDownloadsPending() (value uint32, err error) {
	retValue, err := instance.GetProperty("PriorityDownloadsPending")
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
