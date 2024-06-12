// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ds_domainrelatedobject struct
type ds_domainrelatedobject struct {
	*ds_top

	//
	DS_associatedDomain []string
}

func Newds_domainrelatedobjectEx1(instance *cim.WmiInstance) (newInstance *ds_domainrelatedobject, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_domainrelatedobject{
		ds_top: tmp,
	}
	return
}

func Newds_domainrelatedobjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_domainrelatedobject, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_domainrelatedobject{
		ds_top: tmp,
	}
	return
}

// SetDS_associatedDomain sets the value of DS_associatedDomain for the instance
func (instance *ds_domainrelatedobject) SetPropertyDS_associatedDomain(value []string) (err error) {
	return instance.SetProperty("DS_associatedDomain", (value))
}

// GetDS_associatedDomain gets the value of DS_associatedDomain for the instance
func (instance *ds_domainrelatedobject) GetPropertyDS_associatedDomain() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_associatedDomain")
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
