// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_RoamingUserHealthConfiguration struct
type Win32_RoamingUserHealthConfiguration struct {
	cim.WmiInstance

	// Configure how the Win32_UserProfile::HealthStatus property should reflect the use of temporary profiles.
	HealthStatusForTempProfiles RoamingUserHealthConfiguration_HealthStatusForTempProfiles

	// This is the time threshold, in hours, after which the profile health is reported as Caution when the profile has not been downloaded yet
	LastProfileDownloadIntervalCautionInHours uint16

	// This is the time threshold, in hours, after which the profile health is reported as Unhealthy when the profile has not been uploaded yet
	LastProfileDownloadIntervalUnhealthyInHours uint16

	// This is the time threshold, in hours, after which the profile health is reported as Caution when the profile has not been uploaded yet
	LastProfileUploadIntervalCautionInHours uint16

	// This is the time threshold, in hours, after which the profile health is reported as Unhealthy when the profile has not been downloaded yet
	LastProfileUploadIntervalUnhealthyInHours uint16
}

// SetHealthStatusForTempProfiles sets the value of HealthStatusForTempProfiles for the instance
func (instance *Win32_RoamingUserHealthConfiguration) SetPropertyHealthStatusForTempProfiles(value RoamingUserHealthConfiguration_HealthStatusForTempProfiles) (err error) {
	return instance.SetProperty("HealthStatusForTempProfiles", value)
}

// GetHealthStatusForTempProfiles gets the value of HealthStatusForTempProfiles for the instance
func (instance *Win32_RoamingUserHealthConfiguration) GetPropertyHealthStatusForTempProfiles() (value RoamingUserHealthConfiguration_HealthStatusForTempProfiles, err error) {
	retValue, err := instance.GetProperty("HealthStatusForTempProfiles")
	if err != nil {
		return
	}
	value, ok := retValue.(RoamingUserHealthConfiguration_HealthStatusForTempProfiles)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastProfileDownloadIntervalCautionInHours sets the value of LastProfileDownloadIntervalCautionInHours for the instance
func (instance *Win32_RoamingUserHealthConfiguration) SetPropertyLastProfileDownloadIntervalCautionInHours(value uint16) (err error) {
	return instance.SetProperty("LastProfileDownloadIntervalCautionInHours", value)
}

// GetLastProfileDownloadIntervalCautionInHours gets the value of LastProfileDownloadIntervalCautionInHours for the instance
func (instance *Win32_RoamingUserHealthConfiguration) GetPropertyLastProfileDownloadIntervalCautionInHours() (value uint16, err error) {
	retValue, err := instance.GetProperty("LastProfileDownloadIntervalCautionInHours")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastProfileDownloadIntervalUnhealthyInHours sets the value of LastProfileDownloadIntervalUnhealthyInHours for the instance
func (instance *Win32_RoamingUserHealthConfiguration) SetPropertyLastProfileDownloadIntervalUnhealthyInHours(value uint16) (err error) {
	return instance.SetProperty("LastProfileDownloadIntervalUnhealthyInHours", value)
}

// GetLastProfileDownloadIntervalUnhealthyInHours gets the value of LastProfileDownloadIntervalUnhealthyInHours for the instance
func (instance *Win32_RoamingUserHealthConfiguration) GetPropertyLastProfileDownloadIntervalUnhealthyInHours() (value uint16, err error) {
	retValue, err := instance.GetProperty("LastProfileDownloadIntervalUnhealthyInHours")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastProfileUploadIntervalCautionInHours sets the value of LastProfileUploadIntervalCautionInHours for the instance
func (instance *Win32_RoamingUserHealthConfiguration) SetPropertyLastProfileUploadIntervalCautionInHours(value uint16) (err error) {
	return instance.SetProperty("LastProfileUploadIntervalCautionInHours", value)
}

// GetLastProfileUploadIntervalCautionInHours gets the value of LastProfileUploadIntervalCautionInHours for the instance
func (instance *Win32_RoamingUserHealthConfiguration) GetPropertyLastProfileUploadIntervalCautionInHours() (value uint16, err error) {
	retValue, err := instance.GetProperty("LastProfileUploadIntervalCautionInHours")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastProfileUploadIntervalUnhealthyInHours sets the value of LastProfileUploadIntervalUnhealthyInHours for the instance
func (instance *Win32_RoamingUserHealthConfiguration) SetPropertyLastProfileUploadIntervalUnhealthyInHours(value uint16) (err error) {
	return instance.SetProperty("LastProfileUploadIntervalUnhealthyInHours", value)
}

// GetLastProfileUploadIntervalUnhealthyInHours gets the value of LastProfileUploadIntervalUnhealthyInHours for the instance
func (instance *Win32_RoamingUserHealthConfiguration) GetPropertyLastProfileUploadIntervalUnhealthyInHours() (value uint16, err error) {
	retValue, err := instance.GetProperty("LastProfileUploadIntervalUnhealthyInHours")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
