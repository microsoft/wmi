// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_DeviceRegistrationInfo struct
type MDM_DeviceRegistrationInfo struct {
	cim.WmiInstance

	//
	CertificateThumbprint string

	//
	DeviceId string

	//
	UPN string
}

// SetCertificateThumbprint sets the value of CertificateThumbprint for the instance
func (instance *MDM_DeviceRegistrationInfo) SetPropertyCertificateThumbprint(value string) (err error) {
	return instance.SetProperty("CertificateThumbprint", value)
}

// GetCertificateThumbprint gets the value of CertificateThumbprint for the instance
func (instance *MDM_DeviceRegistrationInfo) GetPropertyCertificateThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateThumbprint")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceId sets the value of DeviceId for the instance
func (instance *MDM_DeviceRegistrationInfo) SetPropertyDeviceId(value string) (err error) {
	return instance.SetProperty("DeviceId", value)
}

// GetDeviceId gets the value of DeviceId for the instance
func (instance *MDM_DeviceRegistrationInfo) GetPropertyDeviceId() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUPN sets the value of UPN for the instance
func (instance *MDM_DeviceRegistrationInfo) SetPropertyUPN(value string) (err error) {
	return instance.SetProperty("UPN", value)
}

// GetUPN gets the value of UPN for the instance
func (instance *MDM_DeviceRegistrationInfo) GetPropertyUPN() (value string, err error) {
	retValue, err := instance.GetProperty("UPN")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
