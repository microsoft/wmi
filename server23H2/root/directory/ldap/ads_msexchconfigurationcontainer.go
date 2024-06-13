// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_msexchconfigurationcontainer struct
type ads_msexchconfigurationcontainer struct {
	*ads_container

	//
	DS_addressBookRoots []string

	//
	DS_addressBookRoots2 []string

	//
	DS_globalAddressList []string

	//
	DS_globalAddressList2 []string

	//
	DS_templateRoots []string

	//
	DS_templateRoots2 []string
}

func Newads_msexchconfigurationcontainerEx1(instance *cim.WmiInstance) (newInstance *ads_msexchconfigurationcontainer, err error) {
	tmp, err := Newads_containerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msexchconfigurationcontainer{
		ads_container: tmp,
	}
	return
}

func Newads_msexchconfigurationcontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msexchconfigurationcontainer, err error) {
	tmp, err := Newads_containerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msexchconfigurationcontainer{
		ads_container: tmp,
	}
	return
}

// SetDS_addressBookRoots sets the value of DS_addressBookRoots for the instance
func (instance *ads_msexchconfigurationcontainer) SetPropertyDS_addressBookRoots(value []string) (err error) {
	return instance.SetProperty("DS_addressBookRoots", (value))
}

// GetDS_addressBookRoots gets the value of DS_addressBookRoots for the instance
func (instance *ads_msexchconfigurationcontainer) GetPropertyDS_addressBookRoots() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_addressBookRoots")
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

// SetDS_addressBookRoots2 sets the value of DS_addressBookRoots2 for the instance
func (instance *ads_msexchconfigurationcontainer) SetPropertyDS_addressBookRoots2(value []string) (err error) {
	return instance.SetProperty("DS_addressBookRoots2", (value))
}

// GetDS_addressBookRoots2 gets the value of DS_addressBookRoots2 for the instance
func (instance *ads_msexchconfigurationcontainer) GetPropertyDS_addressBookRoots2() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_addressBookRoots2")
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

// SetDS_globalAddressList sets the value of DS_globalAddressList for the instance
func (instance *ads_msexchconfigurationcontainer) SetPropertyDS_globalAddressList(value []string) (err error) {
	return instance.SetProperty("DS_globalAddressList", (value))
}

// GetDS_globalAddressList gets the value of DS_globalAddressList for the instance
func (instance *ads_msexchconfigurationcontainer) GetPropertyDS_globalAddressList() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_globalAddressList")
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

// SetDS_globalAddressList2 sets the value of DS_globalAddressList2 for the instance
func (instance *ads_msexchconfigurationcontainer) SetPropertyDS_globalAddressList2(value []string) (err error) {
	return instance.SetProperty("DS_globalAddressList2", (value))
}

// GetDS_globalAddressList2 gets the value of DS_globalAddressList2 for the instance
func (instance *ads_msexchconfigurationcontainer) GetPropertyDS_globalAddressList2() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_globalAddressList2")
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

// SetDS_templateRoots sets the value of DS_templateRoots for the instance
func (instance *ads_msexchconfigurationcontainer) SetPropertyDS_templateRoots(value []string) (err error) {
	return instance.SetProperty("DS_templateRoots", (value))
}

// GetDS_templateRoots gets the value of DS_templateRoots for the instance
func (instance *ads_msexchconfigurationcontainer) GetPropertyDS_templateRoots() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_templateRoots")
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

// SetDS_templateRoots2 sets the value of DS_templateRoots2 for the instance
func (instance *ads_msexchconfigurationcontainer) SetPropertyDS_templateRoots2(value []string) (err error) {
	return instance.SetProperty("DS_templateRoots2", (value))
}

// GetDS_templateRoots2 gets the value of DS_templateRoots2 for the instance
func (instance *ads_msexchconfigurationcontainer) GetPropertyDS_templateRoots2() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_templateRoots2")
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
