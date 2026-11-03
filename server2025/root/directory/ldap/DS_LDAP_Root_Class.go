// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DS_LDAP_Root_Class struct
type DS_LDAP_Root_Class struct {
	*cim.WmiInstance

	//
	ADSIPath string
}

func NewDS_LDAP_Root_ClassEx1(instance *cim.WmiInstance) (newInstance *DS_LDAP_Root_Class, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DS_LDAP_Root_Class{
		WmiInstance: tmp,
	}
	return
}

func NewDS_LDAP_Root_ClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DS_LDAP_Root_Class, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DS_LDAP_Root_Class{
		WmiInstance: tmp,
	}
	return
}

// SetADSIPath sets the value of ADSIPath for the instance
func (instance *DS_LDAP_Root_Class) SetPropertyADSIPath(value string) (err error) {
	return instance.SetProperty("ADSIPath", (value))
}

// GetADSIPath gets the value of ADSIPath for the instance
func (instance *DS_LDAP_Root_Class) GetPropertyADSIPath() (value string, err error) {
	retValue, err := instance.GetProperty("ADSIPath")
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
