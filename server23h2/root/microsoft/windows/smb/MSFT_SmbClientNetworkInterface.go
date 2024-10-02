// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_SmbClientNetworkInterface struct
type MSFT_SmbClientNetworkInterface struct {
	*cim.WmiInstance

	//
	FriendlyName string

	//
	InterfaceIndex uint32

	//
	IpAddresses []string

	//
	LinkSpeed uint64

	//
	RdmaCapable bool

	//
	RssCapable bool
}

func NewMSFT_SmbClientNetworkInterfaceEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbClientNetworkInterface, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbClientNetworkInterface{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbClientNetworkInterfaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbClientNetworkInterface, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbClientNetworkInterface{
		WmiInstance: tmp,
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_SmbClientNetworkInterface) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_SmbClientNetworkInterface) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
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

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_SmbClientNetworkInterface) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", (value))
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_SmbClientNetworkInterface) GetPropertyInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceIndex")
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

// SetIpAddresses sets the value of IpAddresses for the instance
func (instance *MSFT_SmbClientNetworkInterface) SetPropertyIpAddresses(value []string) (err error) {
	return instance.SetProperty("IpAddresses", (value))
}

// GetIpAddresses gets the value of IpAddresses for the instance
func (instance *MSFT_SmbClientNetworkInterface) GetPropertyIpAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IpAddresses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetLinkSpeed sets the value of LinkSpeed for the instance
func (instance *MSFT_SmbClientNetworkInterface) SetPropertyLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("LinkSpeed", (value))
}

// GetLinkSpeed gets the value of LinkSpeed for the instance
func (instance *MSFT_SmbClientNetworkInterface) GetPropertyLinkSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("LinkSpeed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRdmaCapable sets the value of RdmaCapable for the instance
func (instance *MSFT_SmbClientNetworkInterface) SetPropertyRdmaCapable(value bool) (err error) {
	return instance.SetProperty("RdmaCapable", (value))
}

// GetRdmaCapable gets the value of RdmaCapable for the instance
func (instance *MSFT_SmbClientNetworkInterface) GetPropertyRdmaCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("RdmaCapable")
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

// SetRssCapable sets the value of RssCapable for the instance
func (instance *MSFT_SmbClientNetworkInterface) SetPropertyRssCapable(value bool) (err error) {
	return instance.SetProperty("RssCapable", (value))
}

// GetRssCapable gets the value of RssCapable for the instance
func (instance *MSFT_SmbClientNetworkInterface) GetPropertyRssCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("RssCapable")
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
