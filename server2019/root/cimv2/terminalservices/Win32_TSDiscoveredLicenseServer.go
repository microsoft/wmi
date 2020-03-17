// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_TSDiscoveredLicenseServer struct
type Win32_TSDiscoveredLicenseServer struct {
	cim.WmiInstance

	//
	HowDiscovered uint32

	//
	IsAdminOnLS uint32

	//
	IsLSAvailable uint32

	//
	IssuingCALs uint32

	//
	LicenseServer string
}

// SetHowDiscovered sets the value of HowDiscovered for the instance
func (instance *Win32_TSDiscoveredLicenseServer) SetPropertyHowDiscovered(value uint32) (err error) {
	return instance.SetProperty("HowDiscovered", value)
}

// GetHowDiscovered gets the value of HowDiscovered for the instance
func (instance *Win32_TSDiscoveredLicenseServer) GetPropertyHowDiscovered() (value uint32, err error) {
	retValue, err := instance.GetProperty("HowDiscovered")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsAdminOnLS sets the value of IsAdminOnLS for the instance
func (instance *Win32_TSDiscoveredLicenseServer) SetPropertyIsAdminOnLS(value uint32) (err error) {
	return instance.SetProperty("IsAdminOnLS", value)
}

// GetIsAdminOnLS gets the value of IsAdminOnLS for the instance
func (instance *Win32_TSDiscoveredLicenseServer) GetPropertyIsAdminOnLS() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsAdminOnLS")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsLSAvailable sets the value of IsLSAvailable for the instance
func (instance *Win32_TSDiscoveredLicenseServer) SetPropertyIsLSAvailable(value uint32) (err error) {
	return instance.SetProperty("IsLSAvailable", value)
}

// GetIsLSAvailable gets the value of IsLSAvailable for the instance
func (instance *Win32_TSDiscoveredLicenseServer) GetPropertyIsLSAvailable() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsLSAvailable")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIssuingCALs sets the value of IssuingCALs for the instance
func (instance *Win32_TSDiscoveredLicenseServer) SetPropertyIssuingCALs(value uint32) (err error) {
	return instance.SetProperty("IssuingCALs", value)
}

// GetIssuingCALs gets the value of IssuingCALs for the instance
func (instance *Win32_TSDiscoveredLicenseServer) GetPropertyIssuingCALs() (value uint32, err error) {
	retValue, err := instance.GetProperty("IssuingCALs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLicenseServer sets the value of LicenseServer for the instance
func (instance *Win32_TSDiscoveredLicenseServer) SetPropertyLicenseServer(value string) (err error) {
	return instance.SetProperty("LicenseServer", value)
}

// GetLicenseServer gets the value of LicenseServer for the instance
func (instance *Win32_TSDiscoveredLicenseServer) GetPropertyLicenseServer() (value string, err error) {
	retValue, err := instance.GetProperty("LicenseServer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
