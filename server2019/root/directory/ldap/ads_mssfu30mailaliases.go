// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_mssfu30mailaliases struct
type ads_mssfu30mailaliases struct {
	*ds_top

	//
	DS_msSFU30Aliases []string

	//
	DS_msSFU30Name string

	//
	DS_msSFU30NisDomain string

	//
	DS_nisMapName string
}

func Newads_mssfu30mailaliasesEx1(instance *cim.WmiInstance) (newInstance *ads_mssfu30mailaliases, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mssfu30mailaliases{
		ds_top: tmp,
	}
	return
}

func Newads_mssfu30mailaliasesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mssfu30mailaliases, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mssfu30mailaliases{
		ds_top: tmp,
	}
	return
}

// SetDS_msSFU30Aliases sets the value of DS_msSFU30Aliases for the instance
func (instance *ads_mssfu30mailaliases) SetPropertyDS_msSFU30Aliases(value []string) (err error) {
	return instance.SetProperty("DS_msSFU30Aliases", (value))
}

// GetDS_msSFU30Aliases gets the value of DS_msSFU30Aliases for the instance
func (instance *ads_mssfu30mailaliases) GetPropertyDS_msSFU30Aliases() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30Aliases")
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

// SetDS_msSFU30Name sets the value of DS_msSFU30Name for the instance
func (instance *ads_mssfu30mailaliases) SetPropertyDS_msSFU30Name(value string) (err error) {
	return instance.SetProperty("DS_msSFU30Name", (value))
}

// GetDS_msSFU30Name gets the value of DS_msSFU30Name for the instance
func (instance *ads_mssfu30mailaliases) GetPropertyDS_msSFU30Name() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30Name")
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

// SetDS_msSFU30NisDomain sets the value of DS_msSFU30NisDomain for the instance
func (instance *ads_mssfu30mailaliases) SetPropertyDS_msSFU30NisDomain(value string) (err error) {
	return instance.SetProperty("DS_msSFU30NisDomain", (value))
}

// GetDS_msSFU30NisDomain gets the value of DS_msSFU30NisDomain for the instance
func (instance *ads_mssfu30mailaliases) GetPropertyDS_msSFU30NisDomain() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30NisDomain")
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

// SetDS_nisMapName sets the value of DS_nisMapName for the instance
func (instance *ads_mssfu30mailaliases) SetPropertyDS_nisMapName(value string) (err error) {
	return instance.SetProperty("DS_nisMapName", (value))
}

// GetDS_nisMapName gets the value of DS_nisMapName for the instance
func (instance *ads_mssfu30mailaliases) GetPropertyDS_nisMapName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_nisMapName")
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
