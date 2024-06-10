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

// MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo struct
type MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo struct {
	*cim.WmiInstance

	//
	AuthMethod IKEPresharedKeyAuthenticationInfo_AuthMethod

	//
	Id []uint8

	//
	IdType IKEPresharedKeyAuthenticationInfo_IdType

	//
	key []uint8

	//
	SecurityFlags uint64
}

func NewMSiSCSIInitiator_IKEPresharedKeyAuthenticationInfoEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSIInitiator_IKEPresharedKeyAuthenticationInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo{
		WmiInstance: tmp,
	}
	return
}

// SetAuthMethod sets the value of AuthMethod for the instance
func (instance *MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo) SetPropertyAuthMethod(value IKEPresharedKeyAuthenticationInfo_AuthMethod) (err error) {
	return instance.SetProperty("AuthMethod", (value))
}

// GetAuthMethod gets the value of AuthMethod for the instance
func (instance *MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo) GetPropertyAuthMethod() (value IKEPresharedKeyAuthenticationInfo_AuthMethod, err error) {
	retValue, err := instance.GetProperty("AuthMethod")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = IKEPresharedKeyAuthenticationInfo_AuthMethod(valuetmp)

	return
}

// SetId sets the value of Id for the instance
func (instance *MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo) SetPropertyId(value []uint8) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo) GetPropertyId() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetIdType sets the value of IdType for the instance
func (instance *MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo) SetPropertyIdType(value IKEPresharedKeyAuthenticationInfo_IdType) (err error) {
	return instance.SetProperty("IdType", (value))
}

// GetIdType gets the value of IdType for the instance
func (instance *MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo) GetPropertyIdType() (value IKEPresharedKeyAuthenticationInfo_IdType, err error) {
	retValue, err := instance.GetProperty("IdType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = IKEPresharedKeyAuthenticationInfo_IdType(valuetmp)

	return
}

// Setkey sets the value of key for the instance
func (instance *MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo) SetPropertykey(value []uint8) (err error) {
	return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo) GetPropertykey() (value []uint8, err error) {
	retValue, err := instance.GetProperty("key")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetSecurityFlags sets the value of SecurityFlags for the instance
func (instance *MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo) SetPropertySecurityFlags(value uint64) (err error) {
	return instance.SetProperty("SecurityFlags", (value))
}

// GetSecurityFlags gets the value of SecurityFlags for the instance
func (instance *MSiSCSIInitiator_IKEPresharedKeyAuthenticationInfo) GetPropertySecurityFlags() (value uint64, err error) {
	retValue, err := instance.GetProperty("SecurityFlags")
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
