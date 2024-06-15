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

// ads_mssfu30nismapconfig struct
type ads_mssfu30nismapconfig struct {
	*ds_top

	//
	DS_msSFU30FieldSeparator string

	//
	DS_msSFU30IntraFieldSeparator string

	//
	DS_msSFU30KeyAttributes []string

	//
	DS_msSFU30MapFilter string

	//
	DS_msSFU30NSMAPFieldPosition string

	//
	DS_msSFU30ResultAttributes []string

	//
	DS_msSFU30SearchAttributes []string
}

func Newads_mssfu30nismapconfigEx1(instance *cim.WmiInstance) (newInstance *ads_mssfu30nismapconfig, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mssfu30nismapconfig{
		ds_top: tmp,
	}
	return
}

func Newads_mssfu30nismapconfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mssfu30nismapconfig, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mssfu30nismapconfig{
		ds_top: tmp,
	}
	return
}

// SetDS_msSFU30FieldSeparator sets the value of DS_msSFU30FieldSeparator for the instance
func (instance *ads_mssfu30nismapconfig) SetPropertyDS_msSFU30FieldSeparator(value string) (err error) {
	return instance.SetProperty("DS_msSFU30FieldSeparator", (value))
}

// GetDS_msSFU30FieldSeparator gets the value of DS_msSFU30FieldSeparator for the instance
func (instance *ads_mssfu30nismapconfig) GetPropertyDS_msSFU30FieldSeparator() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30FieldSeparator")
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

// SetDS_msSFU30IntraFieldSeparator sets the value of DS_msSFU30IntraFieldSeparator for the instance
func (instance *ads_mssfu30nismapconfig) SetPropertyDS_msSFU30IntraFieldSeparator(value string) (err error) {
	return instance.SetProperty("DS_msSFU30IntraFieldSeparator", (value))
}

// GetDS_msSFU30IntraFieldSeparator gets the value of DS_msSFU30IntraFieldSeparator for the instance
func (instance *ads_mssfu30nismapconfig) GetPropertyDS_msSFU30IntraFieldSeparator() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30IntraFieldSeparator")
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

// SetDS_msSFU30KeyAttributes sets the value of DS_msSFU30KeyAttributes for the instance
func (instance *ads_mssfu30nismapconfig) SetPropertyDS_msSFU30KeyAttributes(value []string) (err error) {
	return instance.SetProperty("DS_msSFU30KeyAttributes", (value))
}

// GetDS_msSFU30KeyAttributes gets the value of DS_msSFU30KeyAttributes for the instance
func (instance *ads_mssfu30nismapconfig) GetPropertyDS_msSFU30KeyAttributes() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30KeyAttributes")
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

// SetDS_msSFU30MapFilter sets the value of DS_msSFU30MapFilter for the instance
func (instance *ads_mssfu30nismapconfig) SetPropertyDS_msSFU30MapFilter(value string) (err error) {
	return instance.SetProperty("DS_msSFU30MapFilter", (value))
}

// GetDS_msSFU30MapFilter gets the value of DS_msSFU30MapFilter for the instance
func (instance *ads_mssfu30nismapconfig) GetPropertyDS_msSFU30MapFilter() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30MapFilter")
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

// SetDS_msSFU30NSMAPFieldPosition sets the value of DS_msSFU30NSMAPFieldPosition for the instance
func (instance *ads_mssfu30nismapconfig) SetPropertyDS_msSFU30NSMAPFieldPosition(value string) (err error) {
	return instance.SetProperty("DS_msSFU30NSMAPFieldPosition", (value))
}

// GetDS_msSFU30NSMAPFieldPosition gets the value of DS_msSFU30NSMAPFieldPosition for the instance
func (instance *ads_mssfu30nismapconfig) GetPropertyDS_msSFU30NSMAPFieldPosition() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30NSMAPFieldPosition")
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

// SetDS_msSFU30ResultAttributes sets the value of DS_msSFU30ResultAttributes for the instance
func (instance *ads_mssfu30nismapconfig) SetPropertyDS_msSFU30ResultAttributes(value []string) (err error) {
	return instance.SetProperty("DS_msSFU30ResultAttributes", (value))
}

// GetDS_msSFU30ResultAttributes gets the value of DS_msSFU30ResultAttributes for the instance
func (instance *ads_mssfu30nismapconfig) GetPropertyDS_msSFU30ResultAttributes() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30ResultAttributes")
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

// SetDS_msSFU30SearchAttributes sets the value of DS_msSFU30SearchAttributes for the instance
func (instance *ads_mssfu30nismapconfig) SetPropertyDS_msSFU30SearchAttributes(value []string) (err error) {
	return instance.SetProperty("DS_msSFU30SearchAttributes", (value))
}

// GetDS_msSFU30SearchAttributes gets the value of DS_msSFU30SearchAttributes for the instance
func (instance *ads_mssfu30nismapconfig) GetPropertyDS_msSFU30SearchAttributes() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30SearchAttributes")
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
