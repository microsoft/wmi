// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetNat struct
type MSFT_NetNat struct {
	*MSFT_NetSettingData

	//
	Active uint8

	//
	ExternalIPInterfaceAddressPrefix string

	//
	IcmpQueryTimeout uint32

	//
	InternalIPInterfaceAddressPrefix string

	//
	InternalRoutingDomainId string

	//
	Name string

	//
	Store uint32

	//
	TcpEstablishedConnectionTimeout uint32

	//
	TcpFilteringBehavior uint8

	//
	TcpTransientConnectionTimeout uint32

	//
	UdpFilteringBehavior uint8

	//
	UdpIdleSessionTimeout uint32

	//
	UdpInboundRefresh uint8
}

func NewMSFT_NetNatEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetNat, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNat{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetNatEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetNat, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNat{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFT_NetNat) SetPropertyActive(value uint8) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFT_NetNat) GetPropertyActive() (value uint8, err error) {
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

// SetExternalIPInterfaceAddressPrefix sets the value of ExternalIPInterfaceAddressPrefix for the instance
func (instance *MSFT_NetNat) SetPropertyExternalIPInterfaceAddressPrefix(value string) (err error) {
	return instance.SetProperty("ExternalIPInterfaceAddressPrefix", (value))
}

// GetExternalIPInterfaceAddressPrefix gets the value of ExternalIPInterfaceAddressPrefix for the instance
func (instance *MSFT_NetNat) GetPropertyExternalIPInterfaceAddressPrefix() (value string, err error) {
	retValue, err := instance.GetProperty("ExternalIPInterfaceAddressPrefix")
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

// SetIcmpQueryTimeout sets the value of IcmpQueryTimeout for the instance
func (instance *MSFT_NetNat) SetPropertyIcmpQueryTimeout(value uint32) (err error) {
	return instance.SetProperty("IcmpQueryTimeout", (value))
}

// GetIcmpQueryTimeout gets the value of IcmpQueryTimeout for the instance
func (instance *MSFT_NetNat) GetPropertyIcmpQueryTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("IcmpQueryTimeout")
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

// SetInternalIPInterfaceAddressPrefix sets the value of InternalIPInterfaceAddressPrefix for the instance
func (instance *MSFT_NetNat) SetPropertyInternalIPInterfaceAddressPrefix(value string) (err error) {
	return instance.SetProperty("InternalIPInterfaceAddressPrefix", (value))
}

// GetInternalIPInterfaceAddressPrefix gets the value of InternalIPInterfaceAddressPrefix for the instance
func (instance *MSFT_NetNat) GetPropertyInternalIPInterfaceAddressPrefix() (value string, err error) {
	retValue, err := instance.GetProperty("InternalIPInterfaceAddressPrefix")
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

// SetInternalRoutingDomainId sets the value of InternalRoutingDomainId for the instance
func (instance *MSFT_NetNat) SetPropertyInternalRoutingDomainId(value string) (err error) {
	return instance.SetProperty("InternalRoutingDomainId", (value))
}

// GetInternalRoutingDomainId gets the value of InternalRoutingDomainId for the instance
func (instance *MSFT_NetNat) GetPropertyInternalRoutingDomainId() (value string, err error) {
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

// SetName sets the value of Name for the instance
func (instance *MSFT_NetNat) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_NetNat) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetStore sets the value of Store for the instance
func (instance *MSFT_NetNat) SetPropertyStore(value uint32) (err error) {
	return instance.SetProperty("Store", (value))
}

// GetStore gets the value of Store for the instance
func (instance *MSFT_NetNat) GetPropertyStore() (value uint32, err error) {
	retValue, err := instance.GetProperty("Store")
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

// SetTcpEstablishedConnectionTimeout sets the value of TcpEstablishedConnectionTimeout for the instance
func (instance *MSFT_NetNat) SetPropertyTcpEstablishedConnectionTimeout(value uint32) (err error) {
	return instance.SetProperty("TcpEstablishedConnectionTimeout", (value))
}

// GetTcpEstablishedConnectionTimeout gets the value of TcpEstablishedConnectionTimeout for the instance
func (instance *MSFT_NetNat) GetPropertyTcpEstablishedConnectionTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpEstablishedConnectionTimeout")
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

// SetTcpFilteringBehavior sets the value of TcpFilteringBehavior for the instance
func (instance *MSFT_NetNat) SetPropertyTcpFilteringBehavior(value uint8) (err error) {
	return instance.SetProperty("TcpFilteringBehavior", (value))
}

// GetTcpFilteringBehavior gets the value of TcpFilteringBehavior for the instance
func (instance *MSFT_NetNat) GetPropertyTcpFilteringBehavior() (value uint8, err error) {
	retValue, err := instance.GetProperty("TcpFilteringBehavior")
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

// SetTcpTransientConnectionTimeout sets the value of TcpTransientConnectionTimeout for the instance
func (instance *MSFT_NetNat) SetPropertyTcpTransientConnectionTimeout(value uint32) (err error) {
	return instance.SetProperty("TcpTransientConnectionTimeout", (value))
}

// GetTcpTransientConnectionTimeout gets the value of TcpTransientConnectionTimeout for the instance
func (instance *MSFT_NetNat) GetPropertyTcpTransientConnectionTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpTransientConnectionTimeout")
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

// SetUdpFilteringBehavior sets the value of UdpFilteringBehavior for the instance
func (instance *MSFT_NetNat) SetPropertyUdpFilteringBehavior(value uint8) (err error) {
	return instance.SetProperty("UdpFilteringBehavior", (value))
}

// GetUdpFilteringBehavior gets the value of UdpFilteringBehavior for the instance
func (instance *MSFT_NetNat) GetPropertyUdpFilteringBehavior() (value uint8, err error) {
	retValue, err := instance.GetProperty("UdpFilteringBehavior")
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

// SetUdpIdleSessionTimeout sets the value of UdpIdleSessionTimeout for the instance
func (instance *MSFT_NetNat) SetPropertyUdpIdleSessionTimeout(value uint32) (err error) {
	return instance.SetProperty("UdpIdleSessionTimeout", (value))
}

// GetUdpIdleSessionTimeout gets the value of UdpIdleSessionTimeout for the instance
func (instance *MSFT_NetNat) GetPropertyUdpIdleSessionTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("UdpIdleSessionTimeout")
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

// SetUdpInboundRefresh sets the value of UdpInboundRefresh for the instance
func (instance *MSFT_NetNat) SetPropertyUdpInboundRefresh(value uint8) (err error) {
	return instance.SetProperty("UdpInboundRefresh", (value))
}

// GetUdpInboundRefresh gets the value of UdpInboundRefresh for the instance
func (instance *MSFT_NetNat) GetPropertyUdpInboundRefresh() (value uint8, err error) {
	retValue, err := instance.GetProperty("UdpInboundRefresh")
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
