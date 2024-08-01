// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_DeviceRegistrationInfo struct
type MDM_DeviceRegistrationInfo struct {
	*cim.WmiInstance

	//
	CertificateThumbprint string

	//
	DeviceId string

	//
	UPN string
}

func NewMDM_DeviceRegistrationInfoEx1(instance *cim.WmiInstance) (newInstance *MDM_DeviceRegistrationInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceRegistrationInfo{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_DeviceRegistrationInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_DeviceRegistrationInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceRegistrationInfo{
		WmiInstance: tmp,
	}
	return
}

// SetCertificateThumbprint sets the value of CertificateThumbprint for the instance
func (instance *MDM_DeviceRegistrationInfo) SetPropertyCertificateThumbprint(value string) (err error) {
	return instance.SetProperty("CertificateThumbprint", (value))
}

// GetCertificateThumbprint gets the value of CertificateThumbprint for the instance
func (instance *MDM_DeviceRegistrationInfo) GetPropertyCertificateThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateThumbprint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDeviceId sets the value of DeviceId for the instance
func (instance *MDM_DeviceRegistrationInfo) SetPropertyDeviceId(value string) (err error) {
	return instance.SetProperty("DeviceId", (value))
}

// GetDeviceId gets the value of DeviceId for the instance
func (instance *MDM_DeviceRegistrationInfo) GetPropertyDeviceId() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetUPN sets the value of UPN for the instance
func (instance *MDM_DeviceRegistrationInfo) SetPropertyUPN(value string) (err error) {
	return instance.SetProperty("UPN", (value))
}

// GetUPN gets the value of UPN for the instance
func (instance *MDM_DeviceRegistrationInfo) GetPropertyUPN() (value string, err error) {
	retValue, err := instance.GetProperty("UPN")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
