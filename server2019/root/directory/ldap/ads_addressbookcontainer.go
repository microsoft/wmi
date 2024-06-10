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

// ads_addressbookcontainer struct
type ads_addressbookcontainer struct {
	*ds_top

	//
	DS_purportedSearch string
}

func Newads_addressbookcontainerEx1(instance *cim.WmiInstance) (newInstance *ads_addressbookcontainer, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_addressbookcontainer{
		ds_top: tmp,
	}
	return
}

func Newads_addressbookcontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_addressbookcontainer, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_addressbookcontainer{
		ds_top: tmp,
	}
	return
}

// SetDS_purportedSearch sets the value of DS_purportedSearch for the instance
func (instance *ads_addressbookcontainer) SetPropertyDS_purportedSearch(value string) (err error) {
	return instance.SetProperty("DS_purportedSearch", (value))
}

// GetDS_purportedSearch gets the value of DS_purportedSearch for the instance
func (instance *ads_addressbookcontainer) GetPropertyDS_purportedSearch() (value string, err error) {
	retValue, err := instance.GetProperty("DS_purportedSearch")
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
