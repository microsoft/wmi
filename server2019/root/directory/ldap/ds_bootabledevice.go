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

// ds_bootabledevice struct
type ds_bootabledevice struct {
	*ds_top

	//
	DS_bootFile []string

	//
	DS_bootParameter []string
}

func Newds_bootabledeviceEx1(instance *cim.WmiInstance) (newInstance *ds_bootabledevice, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_bootabledevice{
		ds_top: tmp,
	}
	return
}

func Newds_bootabledeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_bootabledevice, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_bootabledevice{
		ds_top: tmp,
	}
	return
}

// SetDS_bootFile sets the value of DS_bootFile for the instance
func (instance *ds_bootabledevice) SetPropertyDS_bootFile(value []string) (err error) {
	return instance.SetProperty("DS_bootFile", (value))
}

// GetDS_bootFile gets the value of DS_bootFile for the instance
func (instance *ds_bootabledevice) GetPropertyDS_bootFile() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_bootFile")
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

// SetDS_bootParameter sets the value of DS_bootParameter for the instance
func (instance *ds_bootabledevice) SetPropertyDS_bootParameter(value []string) (err error) {
	return instance.SetProperty("DS_bootParameter", (value))
}

// GetDS_bootParameter gets the value of DS_bootParameter for the instance
func (instance *ds_bootabledevice) GetPropertyDS_bootParameter() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_bootParameter")
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
