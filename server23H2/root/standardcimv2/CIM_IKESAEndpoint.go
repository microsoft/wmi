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

// CIM_IKESAEndpoint struct
type CIM_IKESAEndpoint struct {
	*CIM_SecurityAssociationEndpoint

	//
	AuthenticationMethod uint16

	//
	CipherAlgorithm uint16

	//
	GroupId uint16

	//
	HashAlgorithm uint16

	//
	InitiatorCookie uint64

	//
	OtherAuthenticationMethod string

	//
	OtherCipherAlgorithm string

	//
	OtherHashAlgorithm string

	//
	ResponderCookie uint64

	//
	VendorID string
}

func NewCIM_IKESAEndpointEx1(instance *cim.WmiInstance) (newInstance *CIM_IKESAEndpoint, err error) {
	tmp, err := NewCIM_SecurityAssociationEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_IKESAEndpoint{
		CIM_SecurityAssociationEndpoint: tmp,
	}
	return
}

func NewCIM_IKESAEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_IKESAEndpoint, err error) {
	tmp, err := NewCIM_SecurityAssociationEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_IKESAEndpoint{
		CIM_SecurityAssociationEndpoint: tmp,
	}
	return
}

// SetAuthenticationMethod sets the value of AuthenticationMethod for the instance
func (instance *CIM_IKESAEndpoint) SetPropertyAuthenticationMethod(value uint16) (err error) {
	return instance.SetProperty("AuthenticationMethod", (value))
}

// GetAuthenticationMethod gets the value of AuthenticationMethod for the instance
func (instance *CIM_IKESAEndpoint) GetPropertyAuthenticationMethod() (value uint16, err error) {
	retValue, err := instance.GetProperty("AuthenticationMethod")
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

// SetCipherAlgorithm sets the value of CipherAlgorithm for the instance
func (instance *CIM_IKESAEndpoint) SetPropertyCipherAlgorithm(value uint16) (err error) {
	return instance.SetProperty("CipherAlgorithm", (value))
}

// GetCipherAlgorithm gets the value of CipherAlgorithm for the instance
func (instance *CIM_IKESAEndpoint) GetPropertyCipherAlgorithm() (value uint16, err error) {
	retValue, err := instance.GetProperty("CipherAlgorithm")
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

// SetGroupId sets the value of GroupId for the instance
func (instance *CIM_IKESAEndpoint) SetPropertyGroupId(value uint16) (err error) {
	return instance.SetProperty("GroupId", (value))
}

// GetGroupId gets the value of GroupId for the instance
func (instance *CIM_IKESAEndpoint) GetPropertyGroupId() (value uint16, err error) {
	retValue, err := instance.GetProperty("GroupId")
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

// SetHashAlgorithm sets the value of HashAlgorithm for the instance
func (instance *CIM_IKESAEndpoint) SetPropertyHashAlgorithm(value uint16) (err error) {
	return instance.SetProperty("HashAlgorithm", (value))
}

// GetHashAlgorithm gets the value of HashAlgorithm for the instance
func (instance *CIM_IKESAEndpoint) GetPropertyHashAlgorithm() (value uint16, err error) {
	retValue, err := instance.GetProperty("HashAlgorithm")
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

// SetInitiatorCookie sets the value of InitiatorCookie for the instance
func (instance *CIM_IKESAEndpoint) SetPropertyInitiatorCookie(value uint64) (err error) {
	return instance.SetProperty("InitiatorCookie", (value))
}

// GetInitiatorCookie gets the value of InitiatorCookie for the instance
func (instance *CIM_IKESAEndpoint) GetPropertyInitiatorCookie() (value uint64, err error) {
	retValue, err := instance.GetProperty("InitiatorCookie")
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

// SetOtherAuthenticationMethod sets the value of OtherAuthenticationMethod for the instance
func (instance *CIM_IKESAEndpoint) SetPropertyOtherAuthenticationMethod(value string) (err error) {
	return instance.SetProperty("OtherAuthenticationMethod", (value))
}

// GetOtherAuthenticationMethod gets the value of OtherAuthenticationMethod for the instance
func (instance *CIM_IKESAEndpoint) GetPropertyOtherAuthenticationMethod() (value string, err error) {
	retValue, err := instance.GetProperty("OtherAuthenticationMethod")
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

// SetOtherCipherAlgorithm sets the value of OtherCipherAlgorithm for the instance
func (instance *CIM_IKESAEndpoint) SetPropertyOtherCipherAlgorithm(value string) (err error) {
	return instance.SetProperty("OtherCipherAlgorithm", (value))
}

// GetOtherCipherAlgorithm gets the value of OtherCipherAlgorithm for the instance
func (instance *CIM_IKESAEndpoint) GetPropertyOtherCipherAlgorithm() (value string, err error) {
	retValue, err := instance.GetProperty("OtherCipherAlgorithm")
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

// SetOtherHashAlgorithm sets the value of OtherHashAlgorithm for the instance
func (instance *CIM_IKESAEndpoint) SetPropertyOtherHashAlgorithm(value string) (err error) {
	return instance.SetProperty("OtherHashAlgorithm", (value))
}

// GetOtherHashAlgorithm gets the value of OtherHashAlgorithm for the instance
func (instance *CIM_IKESAEndpoint) GetPropertyOtherHashAlgorithm() (value string, err error) {
	retValue, err := instance.GetProperty("OtherHashAlgorithm")
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

// SetResponderCookie sets the value of ResponderCookie for the instance
func (instance *CIM_IKESAEndpoint) SetPropertyResponderCookie(value uint64) (err error) {
	return instance.SetProperty("ResponderCookie", (value))
}

// GetResponderCookie gets the value of ResponderCookie for the instance
func (instance *CIM_IKESAEndpoint) GetPropertyResponderCookie() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResponderCookie")
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

// SetVendorID sets the value of VendorID for the instance
func (instance *CIM_IKESAEndpoint) SetPropertyVendorID(value string) (err error) {
	return instance.SetProperty("VendorID", (value))
}

// GetVendorID gets the value of VendorID for the instance
func (instance *CIM_IKESAEndpoint) GetPropertyVendorID() (value string, err error) {
	retValue, err := instance.GetProperty("VendorID")
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
