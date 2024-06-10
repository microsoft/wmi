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

// ads_remotestorageservicepoint struct
type ads_remotestorageservicepoint struct {
	*ads_serviceadministrationpoint

	//
	DS_remoteStorageGUID string
}

func Newads_remotestorageservicepointEx1(instance *cim.WmiInstance) (newInstance *ads_remotestorageservicepoint, err error) {
	tmp, err := Newads_serviceadministrationpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_remotestorageservicepoint{
		ads_serviceadministrationpoint: tmp,
	}
	return
}

func Newads_remotestorageservicepointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_remotestorageservicepoint, err error) {
	tmp, err := Newads_serviceadministrationpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_remotestorageservicepoint{
		ads_serviceadministrationpoint: tmp,
	}
	return
}

// SetDS_remoteStorageGUID sets the value of DS_remoteStorageGUID for the instance
func (instance *ads_remotestorageservicepoint) SetPropertyDS_remoteStorageGUID(value string) (err error) {
	return instance.SetProperty("DS_remoteStorageGUID", (value))
}

// GetDS_remoteStorageGUID gets the value of DS_remoteStorageGUID for the instance
func (instance *ads_remotestorageservicepoint) GetPropertyDS_remoteStorageGUID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_remoteStorageGUID")
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
