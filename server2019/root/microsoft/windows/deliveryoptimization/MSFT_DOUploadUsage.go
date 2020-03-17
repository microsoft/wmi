// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DeliveryOptimization
//////////////////////////////////////////////
package deliveryoptimization

// MSFT_DOUploadUsage struct
type MSFT_DOUploadUsage struct {
	MSFT_DOUsage

	//
	MonthlyUploadRestriction DOUploadUsage_MonthlyUploadRestriction

	//
	UploadRatePct uint8

	//
	Uploads uint32
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
