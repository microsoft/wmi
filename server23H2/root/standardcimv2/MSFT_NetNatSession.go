// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetNatSession struct
type MSFT_NetNatSession struct {
	*MSFT_NetSettingData

	//
	CreationTime string

	//
	ExternalDestinationAddress string

	//
	ExternalDestinationPort uint16

	//
	ExternalSourceAddress string

	//
	ExternalSourcePort uint16

	//
	InternalDestinationAddress string

	//
	InternalDestinationPort uint16

	//
	InternalRoutingDomainId string

	//
	InternalSourceAddress string

	//
	InternalSourcePort uint16

	//
	NatName string

	//
	Protocol uint32
}

func NewMSFT_NetNatSessionEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetNatSession, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNatSession{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetNatSessionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetNatSession, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNatSession{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetCreationTime sets the value of CreationTime for the instance
func (instance *MSFT_NetNatSession) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", (value))
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *MSFT_NetNatSession) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
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

// SetExternalDestinationAddress sets the value of ExternalDestinationAddress for the instance
func (instance *MSFT_NetNatSession) SetPropertyExternalDestinationAddress(value string) (err error) {
	return instance.SetProperty("ExternalDestinationAddress", (value))
}

// GetExternalDestinationAddress gets the value of ExternalDestinationAddress for the instance
func (instance *MSFT_NetNatSession) GetPropertyExternalDestinationAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ExternalDestinationAddress")
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

// SetExternalDestinationPort sets the value of ExternalDestinationPort for the instance
func (instance *MSFT_NetNatSession) SetPropertyExternalDestinationPort(value uint16) (err error) {
	return instance.SetProperty("ExternalDestinationPort", (value))
}

// GetExternalDestinationPort gets the value of ExternalDestinationPort for the instance
func (instance *MSFT_NetNatSession) GetPropertyExternalDestinationPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("ExternalDestinationPort")
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

// SetExternalSourceAddress sets the value of ExternalSourceAddress for the instance
func (instance *MSFT_NetNatSession) SetPropertyExternalSourceAddress(value string) (err error) {
	return instance.SetProperty("ExternalSourceAddress", (value))
}

// GetExternalSourceAddress gets the value of ExternalSourceAddress for the instance
func (instance *MSFT_NetNatSession) GetPropertyExternalSourceAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ExternalSourceAddress")
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

// SetExternalSourcePort sets the value of ExternalSourcePort for the instance
func (instance *MSFT_NetNatSession) SetPropertyExternalSourcePort(value uint16) (err error) {
	return instance.SetProperty("ExternalSourcePort", (value))
}

// GetExternalSourcePort gets the value of ExternalSourcePort for the instance
func (instance *MSFT_NetNatSession) GetPropertyExternalSourcePort() (value uint16, err error) {
	retValue, err := instance.GetProperty("ExternalSourcePort")
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

// SetInternalDestinationAddress sets the value of InternalDestinationAddress for the instance
func (instance *MSFT_NetNatSession) SetPropertyInternalDestinationAddress(value string) (err error) {
	return instance.SetProperty("InternalDestinationAddress", (value))
}

// GetInternalDestinationAddress gets the value of InternalDestinationAddress for the instance
func (instance *MSFT_NetNatSession) GetPropertyInternalDestinationAddress() (value string, err error) {
	retValue, err := instance.GetProperty("InternalDestinationAddress")
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

// SetInternalDestinationPort sets the value of InternalDestinationPort for the instance
func (instance *MSFT_NetNatSession) SetPropertyInternalDestinationPort(value uint16) (err error) {
	return instance.SetProperty("InternalDestinationPort", (value))
}

// GetInternalDestinationPort gets the value of InternalDestinationPort for the instance
func (instance *MSFT_NetNatSession) GetPropertyInternalDestinationPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("InternalDestinationPort")
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
func (instance *MSFT_NetNatSession) SetPropertyInternalRoutingDomainId(value string) (err error) {
	return instance.SetProperty("InternalRoutingDomainId", (value))
}

// GetInternalRoutingDomainId gets the value of InternalRoutingDomainId for the instance
func (instance *MSFT_NetNatSession) GetPropertyInternalRoutingDomainId() (value string, err error) {
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

// SetInternalSourceAddress sets the value of InternalSourceAddress for the instance
func (instance *MSFT_NetNatSession) SetPropertyInternalSourceAddress(value string) (err error) {
	return instance.SetProperty("InternalSourceAddress", (value))
}

// GetInternalSourceAddress gets the value of InternalSourceAddress for the instance
func (instance *MSFT_NetNatSession) GetPropertyInternalSourceAddress() (value string, err error) {
	retValue, err := instance.GetProperty("InternalSourceAddress")
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

// SetInternalSourcePort sets the value of InternalSourcePort for the instance
func (instance *MSFT_NetNatSession) SetPropertyInternalSourcePort(value uint16) (err error) {
	return instance.SetProperty("InternalSourcePort", (value))
}

// GetInternalSourcePort gets the value of InternalSourcePort for the instance
func (instance *MSFT_NetNatSession) GetPropertyInternalSourcePort() (value uint16, err error) {
	retValue, err := instance.GetProperty("InternalSourcePort")
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

// SetNatName sets the value of NatName for the instance
func (instance *MSFT_NetNatSession) SetPropertyNatName(value string) (err error) {
	return instance.SetProperty("NatName", (value))
}

// GetNatName gets the value of NatName for the instance
func (instance *MSFT_NetNatSession) GetPropertyNatName() (value string, err error) {
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
func (instance *MSFT_NetNatSession) SetPropertyProtocol(value uint32) (err error) {
	return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *MSFT_NetNatSession) GetPropertyProtocol() (value uint32, err error) {
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
