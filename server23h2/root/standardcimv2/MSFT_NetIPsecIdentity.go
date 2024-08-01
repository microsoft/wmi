// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetIPsecIdentity struct
type MSFT_NetIPsecIdentity struct {
	*cim.WmiInstance

	//
	AuthenticationMethod uint32

	//
	Flags uint32

	//
	Identity string

	//
	ImpersonationType uint32
}

func NewMSFT_NetIPsecIdentityEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIPsecIdentity, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPsecIdentity{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetIPsecIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIPsecIdentity, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPsecIdentity{
		WmiInstance: tmp,
	}
	return
}

// SetAuthenticationMethod sets the value of AuthenticationMethod for the instance
func (instance *MSFT_NetIPsecIdentity) SetPropertyAuthenticationMethod(value uint32) (err error) {
	return instance.SetProperty("AuthenticationMethod", (value))
}

// GetAuthenticationMethod gets the value of AuthenticationMethod for the instance
func (instance *MSFT_NetIPsecIdentity) GetPropertyAuthenticationMethod() (value uint32, err error) {
	retValue, err := instance.GetProperty("AuthenticationMethod")
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

// SetFlags sets the value of Flags for the instance
func (instance *MSFT_NetIPsecIdentity) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSFT_NetIPsecIdentity) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetIdentity sets the value of Identity for the instance
func (instance *MSFT_NetIPsecIdentity) SetPropertyIdentity(value string) (err error) {
	return instance.SetProperty("Identity", (value))
}

// GetIdentity gets the value of Identity for the instance
func (instance *MSFT_NetIPsecIdentity) GetPropertyIdentity() (value string, err error) {
	retValue, err := instance.GetProperty("Identity")
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

// SetImpersonationType sets the value of ImpersonationType for the instance
func (instance *MSFT_NetIPsecIdentity) SetPropertyImpersonationType(value uint32) (err error) {
	return instance.SetProperty("ImpersonationType", (value))
}

// GetImpersonationType gets the value of ImpersonationType for the instance
func (instance *MSFT_NetIPsecIdentity) GetPropertyImpersonationType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ImpersonationType")
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
