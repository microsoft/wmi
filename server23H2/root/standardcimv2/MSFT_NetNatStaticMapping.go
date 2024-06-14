// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetNatStaticMapping struct
type MSFT_NetNatStaticMapping struct {
	*MSFT_NetSettingData

	//
	Active uint8

	//
	ExternalIPAddress string

	//
	ExternalPort uint16

	//
	InternalIPAddress string

	//
	InternalPort uint16

	//
	InternalRoutingDomainId string

	//
	NatName string

	//
	Protocol uint32

	//
	RemoteExternalIPAddressPrefix string

	//
	StaticMappingID uint32
}

func NewMSFT_NetNatStaticMappingEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetNatStaticMapping, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNatStaticMapping{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetNatStaticMappingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetNatStaticMapping, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNatStaticMapping{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFT_NetNatStaticMapping) SetPropertyActive(value uint8) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFT_NetNatStaticMapping) GetPropertyActive() (value uint8, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetExternalIPAddress sets the value of ExternalIPAddress for the instance
func (instance *MSFT_NetNatStaticMapping) SetPropertyExternalIPAddress(value string) (err error) {
	return instance.SetProperty("ExternalIPAddress", (value))
}

// GetExternalIPAddress gets the value of ExternalIPAddress for the instance
func (instance *MSFT_NetNatStaticMapping) GetPropertyExternalIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ExternalIPAddress")
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

// SetExternalPort sets the value of ExternalPort for the instance
func (instance *MSFT_NetNatStaticMapping) SetPropertyExternalPort(value uint16) (err error) {
	return instance.SetProperty("ExternalPort", (value))
}

// GetExternalPort gets the value of ExternalPort for the instance
func (instance *MSFT_NetNatStaticMapping) GetPropertyExternalPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("ExternalPort")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetInternalIPAddress sets the value of InternalIPAddress for the instance
func (instance *MSFT_NetNatStaticMapping) SetPropertyInternalIPAddress(value string) (err error) {
	return instance.SetProperty("InternalIPAddress", (value))
}

// GetInternalIPAddress gets the value of InternalIPAddress for the instance
func (instance *MSFT_NetNatStaticMapping) GetPropertyInternalIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("InternalIPAddress")
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

// SetInternalPort sets the value of InternalPort for the instance
func (instance *MSFT_NetNatStaticMapping) SetPropertyInternalPort(value uint16) (err error) {
	return instance.SetProperty("InternalPort", (value))
}

// GetInternalPort gets the value of InternalPort for the instance
func (instance *MSFT_NetNatStaticMapping) GetPropertyInternalPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("InternalPort")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetInternalRoutingDomainId sets the value of InternalRoutingDomainId for the instance
func (instance *MSFT_NetNatStaticMapping) SetPropertyInternalRoutingDomainId(value string) (err error) {
	return instance.SetProperty("InternalRoutingDomainId", (value))
}

// GetInternalRoutingDomainId gets the value of InternalRoutingDomainId for the instance
func (instance *MSFT_NetNatStaticMapping) GetPropertyInternalRoutingDomainId() (value string, err error) {
	retValue, err := instance.GetProperty("InternalRoutingDomainId")
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

// SetNatName sets the value of NatName for the instance
func (instance *MSFT_NetNatStaticMapping) SetPropertyNatName(value string) (err error) {
	return instance.SetProperty("NatName", (value))
}

// GetNatName gets the value of NatName for the instance
func (instance *MSFT_NetNatStaticMapping) GetPropertyNatName() (value string, err error) {
	retValue, err := instance.GetProperty("NatName")
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

// SetProtocol sets the value of Protocol for the instance
func (instance *MSFT_NetNatStaticMapping) SetPropertyProtocol(value uint32) (err error) {
	return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *MSFT_NetNatStaticMapping) GetPropertyProtocol() (value uint32, err error) {
	retValue, err := instance.GetProperty("Protocol")
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

// SetRemoteExternalIPAddressPrefix sets the value of RemoteExternalIPAddressPrefix for the instance
func (instance *MSFT_NetNatStaticMapping) SetPropertyRemoteExternalIPAddressPrefix(value string) (err error) {
	return instance.SetProperty("RemoteExternalIPAddressPrefix", (value))
}

// GetRemoteExternalIPAddressPrefix gets the value of RemoteExternalIPAddressPrefix for the instance
func (instance *MSFT_NetNatStaticMapping) GetPropertyRemoteExternalIPAddressPrefix() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteExternalIPAddressPrefix")
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

// SetStaticMappingID sets the value of StaticMappingID for the instance
func (instance *MSFT_NetNatStaticMapping) SetPropertyStaticMappingID(value uint32) (err error) {
	return instance.SetProperty("StaticMappingID", (value))
}

// GetStaticMappingID gets the value of StaticMappingID for the instance
func (instance *MSFT_NetNatStaticMapping) GetPropertyStaticMappingID() (value uint32, err error) {
	retValue, err := instance.GetProperty("StaticMappingID")
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
