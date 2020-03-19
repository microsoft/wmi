// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_TSDiscoveredLicenseServer struct
type Win32_TSDiscoveredLicenseServer struct {
	*cim.WmiInstance

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

func NewWin32_TSDiscoveredLicenseServerEx1(instance *cim.WmiInstance) (newInstance *Win32_TSDiscoveredLicenseServer, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_TSDiscoveredLicenseServer{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_TSDiscoveredLicenseServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSDiscoveredLicenseServer, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSDiscoveredLicenseServer{
		WmiInstance: tmp,
	}
	return
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
