// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSF716F6F3_A4B9_400A_9367_120B37ED0381.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEConnectionDialUpCredentials struct
type RSOP_IEConnectionDialUpCredentials struct {
	*cim.WmiInstance

	//
	callbackID uint32

	//
	callbackNumber string

	//
	connectionName string

	//
	domain string

	//
	entryName string

	//
	password string

	//
	phoneNumber string

	//
	rasDialParamsData []uint8

	//
	rasDialParamsDataSize uint32

	//
	rsopID string

	//
	rsopPrecedence int32

	//
	subEntry uint32

	//
	userName string

	//
	windowsVersion uint32
}

func NewRSOP_IEConnectionDialUpCredentialsEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEConnectionDialUpCredentials, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionDialUpCredentials{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEConnectionDialUpCredentialsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEConnectionDialUpCredentials, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionDialUpCredentials{
		WmiInstance: tmp,
	}
	return
}

// SetcallbackID sets the value of callbackID for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertycallbackID(value uint32) (err error) {
	return instance.SetProperty("callbackID", (value))
}

// GetcallbackID gets the value of callbackID for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertycallbackID() (value uint32, err error) {
	retValue, err := instance.GetProperty("callbackID")
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

// SetcallbackNumber sets the value of callbackNumber for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertycallbackNumber(value string) (err error) {
	return instance.SetProperty("callbackNumber", (value))
}

// GetcallbackNumber gets the value of callbackNumber for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertycallbackNumber() (value string, err error) {
	retValue, err := instance.GetProperty("callbackNumber")
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

// SetconnectionName sets the value of connectionName for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertyconnectionName(value string) (err error) {
	return instance.SetProperty("connectionName", (value))
}

// GetconnectionName gets the value of connectionName for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertyconnectionName() (value string, err error) {
	retValue, err := instance.GetProperty("connectionName")
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

// Setdomain sets the value of domain for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertydomain(value string) (err error) {
	return instance.SetProperty("domain", (value))
}

// Getdomain gets the value of domain for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertydomain() (value string, err error) {
	retValue, err := instance.GetProperty("domain")
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

// SetentryName sets the value of entryName for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertyentryName(value string) (err error) {
	return instance.SetProperty("entryName", (value))
}

// GetentryName gets the value of entryName for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertyentryName() (value string, err error) {
	retValue, err := instance.GetProperty("entryName")
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

// Setpassword sets the value of password for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertypassword(value string) (err error) {
	return instance.SetProperty("password", (value))
}

// Getpassword gets the value of password for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertypassword() (value string, err error) {
	retValue, err := instance.GetProperty("password")
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

// SetphoneNumber sets the value of phoneNumber for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertyphoneNumber(value string) (err error) {
	return instance.SetProperty("phoneNumber", (value))
}

// GetphoneNumber gets the value of phoneNumber for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertyphoneNumber() (value string, err error) {
	retValue, err := instance.GetProperty("phoneNumber")
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

// SetrasDialParamsData sets the value of rasDialParamsData for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertyrasDialParamsData(value []uint8) (err error) {
	return instance.SetProperty("rasDialParamsData", (value))
}

// GetrasDialParamsData gets the value of rasDialParamsData for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertyrasDialParamsData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("rasDialParamsData")
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

// SetrasDialParamsDataSize sets the value of rasDialParamsDataSize for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertyrasDialParamsDataSize(value uint32) (err error) {
	return instance.SetProperty("rasDialParamsDataSize", (value))
}

// GetrasDialParamsDataSize gets the value of rasDialParamsDataSize for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertyrasDialParamsDataSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("rasDialParamsDataSize")
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

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", (value))
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
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

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertyrsopPrecedence(value int32) (err error) {
	return instance.SetProperty("rsopPrecedence", (value))
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertyrsopPrecedence() (value int32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
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

	value = int32(valuetmp)

	return
}

// SetsubEntry sets the value of subEntry for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertysubEntry(value uint32) (err error) {
	return instance.SetProperty("subEntry", (value))
}

// GetsubEntry gets the value of subEntry for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertysubEntry() (value uint32, err error) {
	retValue, err := instance.GetProperty("subEntry")
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

// SetuserName sets the value of userName for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertyuserName(value string) (err error) {
	return instance.SetProperty("userName", (value))
}

// GetuserName gets the value of userName for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertyuserName() (value string, err error) {
	retValue, err := instance.GetProperty("userName")
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

// SetwindowsVersion sets the value of windowsVersion for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) SetPropertywindowsVersion(value uint32) (err error) {
	return instance.SetProperty("windowsVersion", (value))
}

// GetwindowsVersion gets the value of windowsVersion for the instance
func (instance *RSOP_IEConnectionDialUpCredentials) GetPropertywindowsVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("windowsVersion")
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
