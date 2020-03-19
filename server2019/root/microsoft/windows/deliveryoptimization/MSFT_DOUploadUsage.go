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

// MSFT_DOUploadUsage struct
type MSFT_DOUploadUsage struct {
	*MSFT_DOUsage

	//
	MonthlyUploadRestriction DOUploadUsage_MonthlyUploadRestriction

	//
	UploadRatePct uint8

	//
	Uploads uint32
}

func NewMSFT_DOUploadUsageEx1(instance *cim.WmiInstance) (newInstance *MSFT_DOUploadUsage, err error) {
	tmp, err := NewMSFT_DOUsageEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DOUploadUsage{
		MSFT_DOUsage: tmp,
	}
	return
}

func NewMSFT_DOUploadUsageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DOUploadUsage, err error) {
	tmp, err := NewMSFT_DOUsageEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DOUploadUsage{
		MSFT_DOUsage: tmp,
	}
	return
}

// SetMonthlyUploadRestriction sets the value of MonthlyUploadRestriction for the instance
func (instance *MSFT_DOUploadUsage) SetPropertyMonthlyUploadRestriction(value DOUploadUsage_MonthlyUploadRestriction) (err error) {
	return instance.SetProperty("MonthlyUploadRestriction", value)
}

// GetMonthlyUploadRestriction gets the value of MonthlyUploadRestriction for the instance
func (instance *MSFT_DOUploadUsage) GetPropertyMonthlyUploadRestriction() (value DOUploadUsage_MonthlyUploadRestriction, err error) {
	retValue, err := instance.GetProperty("MonthlyUploadRestriction")
	if err != nil {
		return
	}
	value, ok := retValue.(DOUploadUsage_MonthlyUploadRestriction)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUploadRatePct sets the value of UploadRatePct for the instance
func (instance *MSFT_DOUploadUsage) SetPropertyUploadRatePct(value uint8) (err error) {
	return instance.SetProperty("UploadRatePct", value)
}

// GetUploadRatePct gets the value of UploadRatePct for the instance
func (instance *MSFT_DOUploadUsage) GetPropertyUploadRatePct() (value uint8, err error) {
	retValue, err := instance.GetProperty("UploadRatePct")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUploads sets the value of Uploads for the instance
func (instance *MSFT_DOUploadUsage) SetPropertyUploads(value uint32) (err error) {
	return instance.SetProperty("Uploads", value)
}

// GetUploads gets the value of Uploads for the instance
func (instance *MSFT_DOUploadUsage) GetPropertyUploads() (value uint32, err error) {
	retValue, err := instance.GetProperty("Uploads")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
