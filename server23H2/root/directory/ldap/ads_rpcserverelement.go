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

// ads_rpcserverelement struct
type ads_rpcserverelement struct {
	*ds_rpcentry

	//
	DS_rpcNsBindings []string

	//
	DS_rpcNsInterfaceID string

	//
	DS_rpcNsTransferSyntax string
}

func Newads_rpcserverelementEx1(instance *cim.WmiInstance) (newInstance *ads_rpcserverelement, err error) {
	tmp, err := Newds_rpcentryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_rpcserverelement{
		ds_rpcentry: tmp,
	}
	return
}

func Newads_rpcserverelementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_rpcserverelement, err error) {
	tmp, err := Newds_rpcentryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_rpcserverelement{
		ds_rpcentry: tmp,
	}
	return
}

// SetDS_rpcNsBindings sets the value of DS_rpcNsBindings for the instance
func (instance *ads_rpcserverelement) SetPropertyDS_rpcNsBindings(value []string) (err error) {
	return instance.SetProperty("DS_rpcNsBindings", (value))
}

// GetDS_rpcNsBindings gets the value of DS_rpcNsBindings for the instance
func (instance *ads_rpcserverelement) GetPropertyDS_rpcNsBindings() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_rpcNsBindings")
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

// SetDS_rpcNsInterfaceID sets the value of DS_rpcNsInterfaceID for the instance
func (instance *ads_rpcserverelement) SetPropertyDS_rpcNsInterfaceID(value string) (err error) {
	return instance.SetProperty("DS_rpcNsInterfaceID", (value))
}

// GetDS_rpcNsInterfaceID gets the value of DS_rpcNsInterfaceID for the instance
func (instance *ads_rpcserverelement) GetPropertyDS_rpcNsInterfaceID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_rpcNsInterfaceID")
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

// SetDS_rpcNsTransferSyntax sets the value of DS_rpcNsTransferSyntax for the instance
func (instance *ads_rpcserverelement) SetPropertyDS_rpcNsTransferSyntax(value string) (err error) {
	return instance.SetProperty("DS_rpcNsTransferSyntax", (value))
}

// GetDS_rpcNsTransferSyntax gets the value of DS_rpcNsTransferSyntax for the instance
func (instance *ads_rpcserverelement) GetPropertyDS_rpcNsTransferSyntax() (value string, err error) {
	retValue, err := instance.GetProperty("DS_rpcNsTransferSyntax")
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
