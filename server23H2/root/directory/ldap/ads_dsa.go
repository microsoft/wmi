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

// ads_dsa struct
type ads_dsa struct {
	*ads_applicationentity

	//
	DS_knowledgeInformation []string
}

func Newads_dsaEx1(instance *cim.WmiInstance) (newInstance *ads_dsa, err error) {
	tmp, err := Newads_applicationentityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_dsa{
		ads_applicationentity: tmp,
	}
	return
}

func Newads_dsaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_dsa, err error) {
	tmp, err := Newads_applicationentityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_dsa{
		ads_applicationentity: tmp,
	}
	return
}

// SetDS_knowledgeInformation sets the value of DS_knowledgeInformation for the instance
func (instance *ads_dsa) SetPropertyDS_knowledgeInformation(value []string) (err error) {
	return instance.SetProperty("DS_knowledgeInformation", (value))
}

// GetDS_knowledgeInformation gets the value of DS_knowledgeInformation for the instance
func (instance *ads_dsa) GetPropertyDS_knowledgeInformation() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_knowledgeInformation")
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
