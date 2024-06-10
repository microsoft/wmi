// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSI_SecurityCapabilities struct
type MSiSCSI_SecurityCapabilities struct {
	*cim.WmiInstance

	//
	Active bool

	// TRUE if adapter supports certificates
	CertificatesSupported bool

	// **typedef** Array of encryption types. This field is a variable length array.
	EncryptionAvailable []SecurityCapabilities_EncryptionAvailable

	// Number of encryption types available.
	EncryptionAvailableCount uint32

	//
	InstanceName string

	// TRUE if the adapter can use IPSEC to protect iSCSI traffic.
	ProtectiScsiTraffic bool

	// TRUE if the adapter can use IPSEC to protect iSNS traffic.
	ProtectiSNSTraffic bool
}

func NewMSiSCSI_SecurityCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_SecurityCapabilities, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_SecurityCapabilities{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSI_SecurityCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_SecurityCapabilities, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_SecurityCapabilities{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_SecurityCapabilities) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_SecurityCapabilities) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetCertificatesSupported sets the value of CertificatesSupported for the instance
func (instance *MSiSCSI_SecurityCapabilities) SetPropertyCertificatesSupported(value bool) (err error) {
	return instance.SetProperty("CertificatesSupported", (value))
}

// GetCertificatesSupported gets the value of CertificatesSupported for the instance
func (instance *MSiSCSI_SecurityCapabilities) GetPropertyCertificatesSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("CertificatesSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetEncryptionAvailable sets the value of EncryptionAvailable for the instance
func (instance *MSiSCSI_SecurityCapabilities) SetPropertyEncryptionAvailable(value []SecurityCapabilities_EncryptionAvailable) (err error) {
	return instance.SetProperty("EncryptionAvailable", (value))
}

// GetEncryptionAvailable gets the value of EncryptionAvailable for the instance
func (instance *MSiSCSI_SecurityCapabilities) GetPropertyEncryptionAvailable() (value []SecurityCapabilities_EncryptionAvailable, err error) {
	retValue, err := instance.GetProperty("EncryptionAvailable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SecurityCapabilities_EncryptionAvailable(valuetmp))
	}

	return
}

// SetEncryptionAvailableCount sets the value of EncryptionAvailableCount for the instance
func (instance *MSiSCSI_SecurityCapabilities) SetPropertyEncryptionAvailableCount(value uint32) (err error) {
	return instance.SetProperty("EncryptionAvailableCount", (value))
}

// GetEncryptionAvailableCount gets the value of EncryptionAvailableCount for the instance
func (instance *MSiSCSI_SecurityCapabilities) GetPropertyEncryptionAvailableCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("EncryptionAvailableCount")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_SecurityCapabilities) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_SecurityCapabilities) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetProtectiScsiTraffic sets the value of ProtectiScsiTraffic for the instance
func (instance *MSiSCSI_SecurityCapabilities) SetPropertyProtectiScsiTraffic(value bool) (err error) {
	return instance.SetProperty("ProtectiScsiTraffic", (value))
}

// GetProtectiScsiTraffic gets the value of ProtectiScsiTraffic for the instance
func (instance *MSiSCSI_SecurityCapabilities) GetPropertyProtectiScsiTraffic() (value bool, err error) {
	retValue, err := instance.GetProperty("ProtectiScsiTraffic")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetProtectiSNSTraffic sets the value of ProtectiSNSTraffic for the instance
func (instance *MSiSCSI_SecurityCapabilities) SetPropertyProtectiSNSTraffic(value bool) (err error) {
	return instance.SetProperty("ProtectiSNSTraffic", (value))
}

// GetProtectiSNSTraffic gets the value of ProtectiSNSTraffic for the instance
func (instance *MSiSCSI_SecurityCapabilities) GetPropertyProtectiSNSTraffic() (value bool, err error) {
	retValue, err := instance.GetProperty("ProtectiSNSTraffic")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
