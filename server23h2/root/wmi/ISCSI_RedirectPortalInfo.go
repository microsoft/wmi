// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ISCSI_RedirectPortalInfo struct
type ISCSI_RedirectPortalInfo struct {
	*cim.WmiInstance

	// Original Target IP Address given in the login
	OriginalIPAddr ISCSI_IP_Address

	// Original Target portal's socket number given in the login
	OriginalPort uint32

	// TRUE if login was redirected. RedirectedIPAddr and RedirectedPort are valid then.
	Redirected uint8

	// Redirected Target IP Address
	RedirectedIPAddr ISCSI_IP_Address

	// Redirected Target portal's socket number
	RedirectedPort uint32

	// TRUE if the redirection is temporary. FALSE otherwise
	TemporaryRedirect uint8

	// A uniquely generated connection ID. Do not confuse this with CID.
	UniqueConnectionId uint64
}

func NewISCSI_RedirectPortalInfoEx1(instance *cim.WmiInstance) (newInstance *ISCSI_RedirectPortalInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ISCSI_RedirectPortalInfo{
		WmiInstance: tmp,
	}
	return
}

func NewISCSI_RedirectPortalInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ISCSI_RedirectPortalInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ISCSI_RedirectPortalInfo{
		WmiInstance: tmp,
	}
	return
}

// SetOriginalIPAddr sets the value of OriginalIPAddr for the instance
func (instance *ISCSI_RedirectPortalInfo) SetPropertyOriginalIPAddr(value ISCSI_IP_Address) (err error) {
	return instance.SetProperty("OriginalIPAddr", (value))
}

// GetOriginalIPAddr gets the value of OriginalIPAddr for the instance
func (instance *ISCSI_RedirectPortalInfo) GetPropertyOriginalIPAddr() (value ISCSI_IP_Address, err error) {
	retValue, err := instance.GetProperty("OriginalIPAddr")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ISCSI_IP_Address)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ISCSI_IP_Address is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ISCSI_IP_Address(valuetmp)

	return
}

// SetOriginalPort sets the value of OriginalPort for the instance
func (instance *ISCSI_RedirectPortalInfo) SetPropertyOriginalPort(value uint32) (err error) {
	return instance.SetProperty("OriginalPort", (value))
}

// GetOriginalPort gets the value of OriginalPort for the instance
func (instance *ISCSI_RedirectPortalInfo) GetPropertyOriginalPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("OriginalPort")
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

// SetRedirected sets the value of Redirected for the instance
func (instance *ISCSI_RedirectPortalInfo) SetPropertyRedirected(value uint8) (err error) {
	return instance.SetProperty("Redirected", (value))
}

// GetRedirected gets the value of Redirected for the instance
func (instance *ISCSI_RedirectPortalInfo) GetPropertyRedirected() (value uint8, err error) {
	retValue, err := instance.GetProperty("Redirected")
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

// SetRedirectedIPAddr sets the value of RedirectedIPAddr for the instance
func (instance *ISCSI_RedirectPortalInfo) SetPropertyRedirectedIPAddr(value ISCSI_IP_Address) (err error) {
	return instance.SetProperty("RedirectedIPAddr", (value))
}

// GetRedirectedIPAddr gets the value of RedirectedIPAddr for the instance
func (instance *ISCSI_RedirectPortalInfo) GetPropertyRedirectedIPAddr() (value ISCSI_IP_Address, err error) {
	retValue, err := instance.GetProperty("RedirectedIPAddr")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ISCSI_IP_Address)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ISCSI_IP_Address is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ISCSI_IP_Address(valuetmp)

	return
}

// SetRedirectedPort sets the value of RedirectedPort for the instance
func (instance *ISCSI_RedirectPortalInfo) SetPropertyRedirectedPort(value uint32) (err error) {
	return instance.SetProperty("RedirectedPort", (value))
}

// GetRedirectedPort gets the value of RedirectedPort for the instance
func (instance *ISCSI_RedirectPortalInfo) GetPropertyRedirectedPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("RedirectedPort")
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

// SetTemporaryRedirect sets the value of TemporaryRedirect for the instance
func (instance *ISCSI_RedirectPortalInfo) SetPropertyTemporaryRedirect(value uint8) (err error) {
	return instance.SetProperty("TemporaryRedirect", (value))
}

// GetTemporaryRedirect gets the value of TemporaryRedirect for the instance
func (instance *ISCSI_RedirectPortalInfo) GetPropertyTemporaryRedirect() (value uint8, err error) {
	retValue, err := instance.GetProperty("TemporaryRedirect")
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

// SetUniqueConnectionId sets the value of UniqueConnectionId for the instance
func (instance *ISCSI_RedirectPortalInfo) SetPropertyUniqueConnectionId(value uint64) (err error) {
	return instance.SetProperty("UniqueConnectionId", (value))
}

// GetUniqueConnectionId gets the value of UniqueConnectionId for the instance
func (instance *ISCSI_RedirectPortalInfo) GetPropertyUniqueConnectionId() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniqueConnectionId")
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
