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

// ds_domain struct
type ds_domain struct {
	*ds_top

	//
	DS_dc string
}

func Newds_domainEx1(instance *cim.WmiInstance) (newInstance *ds_domain, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_domain{
		ds_top: tmp,
	}
	return
}

func Newds_domainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_domain, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_domain{
		ds_top: tmp,
	}
	return
}

// SetDS_dc sets the value of DS_dc for the instance
func (instance *ds_domain) SetPropertyDS_dc(value string) (err error) {
	return instance.SetProperty("DS_dc", (value))
}

// GetDS_dc gets the value of DS_dc for the instance
func (instance *ds_domain) GetPropertyDS_dc() (value string, err error) {
	retValue, err := instance.GetProperty("DS_dc")
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
