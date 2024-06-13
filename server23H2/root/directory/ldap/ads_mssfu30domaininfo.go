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

// ads_mssfu30domaininfo struct
type ads_mssfu30domaininfo struct {
	*ds_top

	//
	DS_msSFU30CryptMethod string

	//
	DS_msSFU30Domains []string

	//
	DS_msSFU30IsValidContainer int32

	//
	DS_msSFU30MasterServerName string

	//
	DS_msSFU30MaxGidNumber int32

	//
	DS_msSFU30MaxUidNumber int32

	//
	DS_msSFU30OrderNumber string

	//
	DS_msSFU30SearchContainer string

	//
	DS_msSFU30YpServers []string
}

func Newads_mssfu30domaininfoEx1(instance *cim.WmiInstance) (newInstance *ads_mssfu30domaininfo, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mssfu30domaininfo{
		ds_top: tmp,
	}
	return
}

func Newads_mssfu30domaininfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mssfu30domaininfo, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mssfu30domaininfo{
		ds_top: tmp,
	}
	return
}

// SetDS_msSFU30CryptMethod sets the value of DS_msSFU30CryptMethod for the instance
func (instance *ads_mssfu30domaininfo) SetPropertyDS_msSFU30CryptMethod(value string) (err error) {
	return instance.SetProperty("DS_msSFU30CryptMethod", (value))
}

// GetDS_msSFU30CryptMethod gets the value of DS_msSFU30CryptMethod for the instance
func (instance *ads_mssfu30domaininfo) GetPropertyDS_msSFU30CryptMethod() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30CryptMethod")
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

// SetDS_msSFU30Domains sets the value of DS_msSFU30Domains for the instance
func (instance *ads_mssfu30domaininfo) SetPropertyDS_msSFU30Domains(value []string) (err error) {
	return instance.SetProperty("DS_msSFU30Domains", (value))
}

// GetDS_msSFU30Domains gets the value of DS_msSFU30Domains for the instance
func (instance *ads_mssfu30domaininfo) GetPropertyDS_msSFU30Domains() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30Domains")
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

// SetDS_msSFU30IsValidContainer sets the value of DS_msSFU30IsValidContainer for the instance
func (instance *ads_mssfu30domaininfo) SetPropertyDS_msSFU30IsValidContainer(value int32) (err error) {
	return instance.SetProperty("DS_msSFU30IsValidContainer", (value))
}

// GetDS_msSFU30IsValidContainer gets the value of DS_msSFU30IsValidContainer for the instance
func (instance *ads_mssfu30domaininfo) GetPropertyDS_msSFU30IsValidContainer() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30IsValidContainer")
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

// SetDS_msSFU30MasterServerName sets the value of DS_msSFU30MasterServerName for the instance
func (instance *ads_mssfu30domaininfo) SetPropertyDS_msSFU30MasterServerName(value string) (err error) {
	return instance.SetProperty("DS_msSFU30MasterServerName", (value))
}

// GetDS_msSFU30MasterServerName gets the value of DS_msSFU30MasterServerName for the instance
func (instance *ads_mssfu30domaininfo) GetPropertyDS_msSFU30MasterServerName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30MasterServerName")
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

// SetDS_msSFU30MaxGidNumber sets the value of DS_msSFU30MaxGidNumber for the instance
func (instance *ads_mssfu30domaininfo) SetPropertyDS_msSFU30MaxGidNumber(value int32) (err error) {
	return instance.SetProperty("DS_msSFU30MaxGidNumber", (value))
}

// GetDS_msSFU30MaxGidNumber gets the value of DS_msSFU30MaxGidNumber for the instance
func (instance *ads_mssfu30domaininfo) GetPropertyDS_msSFU30MaxGidNumber() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30MaxGidNumber")
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

// SetDS_msSFU30MaxUidNumber sets the value of DS_msSFU30MaxUidNumber for the instance
func (instance *ads_mssfu30domaininfo) SetPropertyDS_msSFU30MaxUidNumber(value int32) (err error) {
	return instance.SetProperty("DS_msSFU30MaxUidNumber", (value))
}

// GetDS_msSFU30MaxUidNumber gets the value of DS_msSFU30MaxUidNumber for the instance
func (instance *ads_mssfu30domaininfo) GetPropertyDS_msSFU30MaxUidNumber() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30MaxUidNumber")
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

// SetDS_msSFU30OrderNumber sets the value of DS_msSFU30OrderNumber for the instance
func (instance *ads_mssfu30domaininfo) SetPropertyDS_msSFU30OrderNumber(value string) (err error) {
	return instance.SetProperty("DS_msSFU30OrderNumber", (value))
}

// GetDS_msSFU30OrderNumber gets the value of DS_msSFU30OrderNumber for the instance
func (instance *ads_mssfu30domaininfo) GetPropertyDS_msSFU30OrderNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30OrderNumber")
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

// SetDS_msSFU30SearchContainer sets the value of DS_msSFU30SearchContainer for the instance
func (instance *ads_mssfu30domaininfo) SetPropertyDS_msSFU30SearchContainer(value string) (err error) {
	return instance.SetProperty("DS_msSFU30SearchContainer", (value))
}

// GetDS_msSFU30SearchContainer gets the value of DS_msSFU30SearchContainer for the instance
func (instance *ads_mssfu30domaininfo) GetPropertyDS_msSFU30SearchContainer() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30SearchContainer")
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

// SetDS_msSFU30YpServers sets the value of DS_msSFU30YpServers for the instance
func (instance *ads_mssfu30domaininfo) SetPropertyDS_msSFU30YpServers(value []string) (err error) {
	return instance.SetProperty("DS_msSFU30YpServers", (value))
}

// GetDS_msSFU30YpServers gets the value of DS_msSFU30YpServers for the instance
func (instance *ads_mssfu30domaininfo) GetPropertyDS_msSFU30YpServers() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30YpServers")
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
