// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// MSFT_SmbServerNetworkInterface struct
type MSFT_SmbServerNetworkInterface struct {
	*cim.WmiInstance

	//
	FriendlyName string

	//
	InterfaceIndex uint32

	//
	IpAddress string

	//
	LinkSpeed uint64

	//
	RdmaCapable bool

	//
	RssCapable bool

	//
	ScopeName string
}

func NewMSFT_SmbServerNetworkInterfaceEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbServerNetworkInterface, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbServerNetworkInterface{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbServerNetworkInterfaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbServerNetworkInterface, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbServerNetworkInterface{
		WmiInstance: tmp,
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyFriendlyName() (value string, err error) {
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
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", (value))
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyInterfaceIndex() (value uint32, err error) {
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

// SetIpAddress sets the value of IpAddress for the instance
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyIpAddress(value string) (err error) {
	return instance.SetProperty("IpAddress", (value))
}

// GetIpAddress gets the value of IpAddress for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyIpAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IpAddress")
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

// SetLinkSpeed sets the value of LinkSpeed for the instance
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("LinkSpeed", (value))
}

// GetLinkSpeed gets the value of LinkSpeed for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyLinkSpeed() (value uint64, err error) {
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
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyRdmaCapable(value bool) (err error) {
	return instance.SetProperty("RdmaCapable", (value))
}

// GetRdmaCapable gets the value of RdmaCapable for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyRdmaCapable() (value bool, err error) {
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
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyRssCapable(value bool) (err error) {
	return instance.SetProperty("RssCapable", (value))
}

// GetRssCapable gets the value of RssCapable for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyRssCapable() (value bool, err error) {
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

// SetScopeName sets the value of ScopeName for the instance
func (instance *MSFT_SmbServerNetworkInterface) SetPropertyScopeName(value string) (err error) {
	return instance.SetProperty("ScopeName", (value))
}

// GetScopeName gets the value of ScopeName for the instance
func (instance *MSFT_SmbServerNetworkInterface) GetPropertyScopeName() (value string, err error) {
	retValue, err := instance.GetProperty("ScopeName")
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
