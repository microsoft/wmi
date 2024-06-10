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

// MSFT_NetIKEMMCryptoSet struct
type MSFT_NetIKEMMCryptoSet struct {
	*MSFT_NetIKECryptoSet

	//
	ForceDiffieHellman bool

	//
	MaxLifetimeMinutes uint32

	//
	MaxLifetimeSessions uint32
}

func NewMSFT_NetIKEMMCryptoSetEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIKEMMCryptoSet, err error) {
	tmp, err := NewMSFT_NetIKECryptoSetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKEMMCryptoSet{
		MSFT_NetIKECryptoSet: tmp,
	}
	return
}

func NewMSFT_NetIKEMMCryptoSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIKEMMCryptoSet, err error) {
	tmp, err := NewMSFT_NetIKECryptoSetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKEMMCryptoSet{
		MSFT_NetIKECryptoSet: tmp,
	}
	return
}

// SetForceDiffieHellman sets the value of ForceDiffieHellman for the instance
func (instance *MSFT_NetIKEMMCryptoSet) SetPropertyForceDiffieHellman(value bool) (err error) {
	return instance.SetProperty("ForceDiffieHellman", (value))
}

// GetForceDiffieHellman gets the value of ForceDiffieHellman for the instance
func (instance *MSFT_NetIKEMMCryptoSet) GetPropertyForceDiffieHellman() (value bool, err error) {
	retValue, err := instance.GetProperty("ForceDiffieHellman")
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

// SetMaxLifetimeMinutes sets the value of MaxLifetimeMinutes for the instance
func (instance *MSFT_NetIKEMMCryptoSet) SetPropertyMaxLifetimeMinutes(value uint32) (err error) {
	return instance.SetProperty("MaxLifetimeMinutes", (value))
}

// GetMaxLifetimeMinutes gets the value of MaxLifetimeMinutes for the instance
func (instance *MSFT_NetIKEMMCryptoSet) GetPropertyMaxLifetimeMinutes() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxLifetimeMinutes")
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

// SetMaxLifetimeSessions sets the value of MaxLifetimeSessions for the instance
func (instance *MSFT_NetIKEMMCryptoSet) SetPropertyMaxLifetimeSessions(value uint32) (err error) {
	return instance.SetProperty("MaxLifetimeSessions", (value))
}

// GetMaxLifetimeSessions gets the value of MaxLifetimeSessions for the instance
func (instance *MSFT_NetIKEMMCryptoSet) GetPropertyMaxLifetimeSessions() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxLifetimeSessions")
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

//

// <param name="NewName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIKEMMCryptoSet) Rename( /* IN */ NewName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Rename", NewName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewGPOSession" type="string "></param>
// <param name="NewID" type="string "></param>
// <param name="NewName" type="string "></param>
// <param name="NewPolicyStore" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIKEMMCryptoSet) CloneObject( /* IN */ NewName string,
	/* IN */ NewID string,
	/* IN */ NewPolicyStore string,
	/* IN */ NewGPOSession string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CloneObject", NewName, NewID, NewPolicyStore, NewGPOSession)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
