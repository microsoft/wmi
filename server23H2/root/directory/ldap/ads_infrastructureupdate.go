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

// ads_infrastructureupdate struct
type ads_infrastructureupdate struct {
	*ds_top

	//
	DS_dNReferenceUpdate []string
}

func Newads_infrastructureupdateEx1(instance *cim.WmiInstance) (newInstance *ads_infrastructureupdate, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_infrastructureupdate{
		ds_top: tmp,
	}
	return
}

func Newads_infrastructureupdateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_infrastructureupdate, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_infrastructureupdate{
		ds_top: tmp,
	}
	return
}

// SetDS_dNReferenceUpdate sets the value of DS_dNReferenceUpdate for the instance
func (instance *ads_infrastructureupdate) SetPropertyDS_dNReferenceUpdate(value []string) (err error) {
	return instance.SetProperty("DS_dNReferenceUpdate", (value))
}

// GetDS_dNReferenceUpdate gets the value of DS_dNReferenceUpdate for the instance
func (instance *ads_infrastructureupdate) GetPropertyDS_dNReferenceUpdate() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_dNReferenceUpdate")
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
