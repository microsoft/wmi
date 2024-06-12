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

// ads_samserver struct
type ads_samserver struct {
	*ds_securityobject

	//
	DS_samDomainUpdates Uint8Array
}

func Newads_samserverEx1(instance *cim.WmiInstance) (newInstance *ads_samserver, err error) {
	tmp, err := Newds_securityobjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_samserver{
		ds_securityobject: tmp,
	}
	return
}

func Newads_samserverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_samserver, err error) {
	tmp, err := Newds_securityobjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_samserver{
		ds_securityobject: tmp,
	}
	return
}

// SetDS_samDomainUpdates sets the value of DS_samDomainUpdates for the instance
func (instance *ads_samserver) SetPropertyDS_samDomainUpdates(value Uint8Array) (err error) {
	return instance.SetProperty("DS_samDomainUpdates", (value))
}

// GetDS_samDomainUpdates gets the value of DS_samDomainUpdates for the instance
func (instance *ads_samserver) GetPropertyDS_samDomainUpdates() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_samDomainUpdates")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}
