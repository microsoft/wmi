// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// OdbcLoggingSection struct
type OdbcLoggingSection struct {
	*ConfigurationSection

	//
	DataSource string

	//
	Password string

	//
	TableName string

	//
	UserName string
}

func NewOdbcLoggingSectionEx1(instance *cim.WmiInstance) (newInstance *OdbcLoggingSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &OdbcLoggingSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewOdbcLoggingSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *OdbcLoggingSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &OdbcLoggingSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetDataSource sets the value of DataSource for the instance
func (instance *OdbcLoggingSection) SetPropertyDataSource(value string) (err error) {
	return instance.SetProperty("DataSource", (value))
}

// GetDataSource gets the value of DataSource for the instance
func (instance *OdbcLoggingSection) GetPropertyDataSource() (value string, err error) {
	retValue, err := instance.GetProperty("DataSource")
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

// SetPassword sets the value of Password for the instance
func (instance *OdbcLoggingSection) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *OdbcLoggingSection) GetPropertyPassword() (value string, err error) {
	retValue, err := instance.GetProperty("Password")
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

// SetTableName sets the value of TableName for the instance
func (instance *OdbcLoggingSection) SetPropertyTableName(value string) (err error) {
	return instance.SetProperty("TableName", (value))
}

// GetTableName gets the value of TableName for the instance
func (instance *OdbcLoggingSection) GetPropertyTableName() (value string, err error) {
	retValue, err := instance.GetProperty("TableName")
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

// SetUserName sets the value of UserName for the instance
func (instance *OdbcLoggingSection) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *OdbcLoggingSection) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
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
