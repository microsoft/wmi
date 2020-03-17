// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

// MSFT_OffloadDataTransferSetting struct
type MSFT_OffloadDataTransferSetting struct {
	MSFT_StorageObject

	//
	NumberOfTokensInUse uint32

	//
	NumberOfTokensMax uint32

	//
	OptimalDataTokenSize uint32

	//
	SupportInterSubsystem bool
}

// SetNumberOfTokensInUse sets the value of NumberOfTokensInUse for the instance
func (instance *MSFT_OffloadDataTransferSetting) SetPropertyNumberOfTokensInUse(value uint32) (err error) {
	return instance.SetProperty("NumberOfTokensInUse", value)
}

// GetNumberOfTokensInUse gets the value of NumberOfTokensInUse for the instance
func (instance *MSFT_OffloadDataTransferSetting) GetPropertyNumberOfTokensInUse() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfTokensInUse")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfTokensMax sets the value of NumberOfTokensMax for the instance
func (instance *MSFT_OffloadDataTransferSetting) SetPropertyNumberOfTokensMax(value uint32) (err error) {
	return instance.SetProperty("NumberOfTokensMax", value)
}

// GetNumberOfTokensMax gets the value of NumberOfTokensMax for the instance
func (instance *MSFT_OffloadDataTransferSetting) GetPropertyNumberOfTokensMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfTokensMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOptimalDataTokenSize sets the value of OptimalDataTokenSize for the instance
func (instance *MSFT_OffloadDataTransferSetting) SetPropertyOptimalDataTokenSize(value uint32) (err error) {
	return instance.SetProperty("OptimalDataTokenSize", value)
}

// GetOptimalDataTokenSize gets the value of OptimalDataTokenSize for the instance
func (instance *MSFT_OffloadDataTransferSetting) GetPropertyOptimalDataTokenSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("OptimalDataTokenSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportInterSubsystem sets the value of SupportInterSubsystem for the instance
func (instance *MSFT_OffloadDataTransferSetting) SetPropertySupportInterSubsystem(value bool) (err error) {
	return instance.SetProperty("SupportInterSubsystem", value)
}

// GetSupportInterSubsystem gets the value of SupportInterSubsystem for the instance
func (instance *MSFT_OffloadDataTransferSetting) GetPropertySupportInterSubsystem() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportInterSubsystem")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
