// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Hardware
//
// ////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_Account struct
type CIM_Account struct {
	*CIM_LogicalElement

	//
	CreationClassName string

	//
	Descriptions []string

	//
	Host []string

	//
	LocalityName []string

	//
	ObjectClass []string

	//
	OrganizationName []string

	//
	OU []string

	//
	SeeAlso []string

	//
	SystemCreationClassName string

	//
	SystemName string

	//
	UserCertificate []string

	//
	UserID string

	//
	UserPassword []string
}

func NewCIM_AccountEx1(instance *cim.WmiInstance) (newInstance *CIM_Account, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Account{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewCIM_AccountEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Account, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Account{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_Account) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", (value))
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_Account) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
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

// SetDescriptions sets the value of Descriptions for the instance
func (instance *CIM_Account) SetPropertyDescriptions(value []string) (err error) {
	return instance.SetProperty("Descriptions", (value))
}

// GetDescriptions gets the value of Descriptions for the instance
func (instance *CIM_Account) GetPropertyDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("Descriptions")
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

// SetHost sets the value of Host for the instance
func (instance *CIM_Account) SetPropertyHost(value []string) (err error) {
	return instance.SetProperty("Host", (value))
}

// GetHost gets the value of Host for the instance
func (instance *CIM_Account) GetPropertyHost() (value []string, err error) {
	retValue, err := instance.GetProperty("Host")
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

// SetLocalityName sets the value of LocalityName for the instance
func (instance *CIM_Account) SetPropertyLocalityName(value []string) (err error) {
	return instance.SetProperty("LocalityName", (value))
}

// GetLocalityName gets the value of LocalityName for the instance
func (instance *CIM_Account) GetPropertyLocalityName() (value []string, err error) {
	retValue, err := instance.GetProperty("LocalityName")
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

// SetObjectClass sets the value of ObjectClass for the instance
func (instance *CIM_Account) SetPropertyObjectClass(value []string) (err error) {
	return instance.SetProperty("ObjectClass", (value))
}

// GetObjectClass gets the value of ObjectClass for the instance
func (instance *CIM_Account) GetPropertyObjectClass() (value []string, err error) {
	retValue, err := instance.GetProperty("ObjectClass")
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

// SetOrganizationName sets the value of OrganizationName for the instance
func (instance *CIM_Account) SetPropertyOrganizationName(value []string) (err error) {
	return instance.SetProperty("OrganizationName", (value))
}

// GetOrganizationName gets the value of OrganizationName for the instance
func (instance *CIM_Account) GetPropertyOrganizationName() (value []string, err error) {
	retValue, err := instance.GetProperty("OrganizationName")
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

// SetOU sets the value of OU for the instance
func (instance *CIM_Account) SetPropertyOU(value []string) (err error) {
	return instance.SetProperty("OU", (value))
}

// GetOU gets the value of OU for the instance
func (instance *CIM_Account) GetPropertyOU() (value []string, err error) {
	retValue, err := instance.GetProperty("OU")
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

// SetSeeAlso sets the value of SeeAlso for the instance
func (instance *CIM_Account) SetPropertySeeAlso(value []string) (err error) {
	return instance.SetProperty("SeeAlso", (value))
}

// GetSeeAlso gets the value of SeeAlso for the instance
func (instance *CIM_Account) GetPropertySeeAlso() (value []string, err error) {
	retValue, err := instance.GetProperty("SeeAlso")
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

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *CIM_Account) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", (value))
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *CIM_Account) GetPropertySystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemCreationClassName")
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

// SetSystemName sets the value of SystemName for the instance
func (instance *CIM_Account) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", (value))
}

// GetSystemName gets the value of SystemName for the instance
func (instance *CIM_Account) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
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

// SetUserCertificate sets the value of UserCertificate for the instance
func (instance *CIM_Account) SetPropertyUserCertificate(value []string) (err error) {
	return instance.SetProperty("UserCertificate", (value))
}

// GetUserCertificate gets the value of UserCertificate for the instance
func (instance *CIM_Account) GetPropertyUserCertificate() (value []string, err error) {
	retValue, err := instance.GetProperty("UserCertificate")
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

// SetUserID sets the value of UserID for the instance
func (instance *CIM_Account) SetPropertyUserID(value string) (err error) {
	return instance.SetProperty("UserID", (value))
}

// GetUserID gets the value of UserID for the instance
func (instance *CIM_Account) GetPropertyUserID() (value string, err error) {
	retValue, err := instance.GetProperty("UserID")
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

// SetUserPassword sets the value of UserPassword for the instance
func (instance *CIM_Account) SetPropertyUserPassword(value []string) (err error) {
	return instance.SetProperty("UserPassword", (value))
}

// GetUserPassword gets the value of UserPassword for the instance
func (instance *CIM_Account) GetPropertyUserPassword() (value []string, err error) {
	retValue, err := instance.GetProperty("UserPassword")
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
