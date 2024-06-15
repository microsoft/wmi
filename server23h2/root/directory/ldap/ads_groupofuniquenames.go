// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_groupofuniquenames struct
type ads_groupofuniquenames struct {
	*ds_top

	//
	DS_businessCategory []string

	//
	DS_o []string

	//
	DS_ou []string

	//
	DS_owner string

	//
	DS_seeAlso []string

	//
	DS_uniqueMember []string
}

func Newads_groupofuniquenamesEx1(instance *cim.WmiInstance) (newInstance *ads_groupofuniquenames, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_groupofuniquenames{
		ds_top: tmp,
	}
	return
}

func Newads_groupofuniquenamesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_groupofuniquenames, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_groupofuniquenames{
		ds_top: tmp,
	}
	return
}

// SetDS_businessCategory sets the value of DS_businessCategory for the instance
func (instance *ads_groupofuniquenames) SetPropertyDS_businessCategory(value []string) (err error) {
	return instance.SetProperty("DS_businessCategory", (value))
}

// GetDS_businessCategory gets the value of DS_businessCategory for the instance
func (instance *ads_groupofuniquenames) GetPropertyDS_businessCategory() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_businessCategory")
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

// SetDS_o sets the value of DS_o for the instance
func (instance *ads_groupofuniquenames) SetPropertyDS_o(value []string) (err error) {
	return instance.SetProperty("DS_o", (value))
}

// GetDS_o gets the value of DS_o for the instance
func (instance *ads_groupofuniquenames) GetPropertyDS_o() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_o")
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

// SetDS_ou sets the value of DS_ou for the instance
func (instance *ads_groupofuniquenames) SetPropertyDS_ou(value []string) (err error) {
	return instance.SetProperty("DS_ou", (value))
}

// GetDS_ou gets the value of DS_ou for the instance
func (instance *ads_groupofuniquenames) GetPropertyDS_ou() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_ou")
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

// SetDS_owner sets the value of DS_owner for the instance
func (instance *ads_groupofuniquenames) SetPropertyDS_owner(value string) (err error) {
	return instance.SetProperty("DS_owner", (value))
}

// GetDS_owner gets the value of DS_owner for the instance
func (instance *ads_groupofuniquenames) GetPropertyDS_owner() (value string, err error) {
	retValue, err := instance.GetProperty("DS_owner")
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

// SetDS_seeAlso sets the value of DS_seeAlso for the instance
func (instance *ads_groupofuniquenames) SetPropertyDS_seeAlso(value []string) (err error) {
	return instance.SetProperty("DS_seeAlso", (value))
}

// GetDS_seeAlso gets the value of DS_seeAlso for the instance
func (instance *ads_groupofuniquenames) GetPropertyDS_seeAlso() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_seeAlso")
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

// SetDS_uniqueMember sets the value of DS_uniqueMember for the instance
func (instance *ads_groupofuniquenames) SetPropertyDS_uniqueMember(value []string) (err error) {
	return instance.SetProperty("DS_uniqueMember", (value))
}

// GetDS_uniqueMember gets the value of DS_uniqueMember for the instance
func (instance *ads_groupofuniquenames) GetPropertyDS_uniqueMember() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_uniqueMember")
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
