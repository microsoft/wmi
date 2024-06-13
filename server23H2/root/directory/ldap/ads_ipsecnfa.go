// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_ipsecnfa struct
type ads_ipsecnfa struct {
	*ds_ipsecbase

	//
	DS_ipsecFilterReference []string

	//
	DS_ipsecNegotiationPolicyReference string
}

func Newads_ipsecnfaEx1(instance *cim.WmiInstance) (newInstance *ads_ipsecnfa, err error) {
	tmp, err := Newds_ipsecbaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ipsecnfa{
		ds_ipsecbase: tmp,
	}
	return
}

func Newads_ipsecnfaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ipsecnfa, err error) {
	tmp, err := Newds_ipsecbaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ipsecnfa{
		ds_ipsecbase: tmp,
	}
	return
}

// SetDS_ipsecFilterReference sets the value of DS_ipsecFilterReference for the instance
func (instance *ads_ipsecnfa) SetPropertyDS_ipsecFilterReference(value []string) (err error) {
	return instance.SetProperty("DS_ipsecFilterReference", (value))
}

// GetDS_ipsecFilterReference gets the value of DS_ipsecFilterReference for the instance
func (instance *ads_ipsecnfa) GetPropertyDS_ipsecFilterReference() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_ipsecFilterReference")
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

// SetDS_ipsecNegotiationPolicyReference sets the value of DS_ipsecNegotiationPolicyReference for the instance
func (instance *ads_ipsecnfa) SetPropertyDS_ipsecNegotiationPolicyReference(value string) (err error) {
	return instance.SetProperty("DS_ipsecNegotiationPolicyReference", (value))
}

// GetDS_ipsecNegotiationPolicyReference gets the value of DS_ipsecNegotiationPolicyReference for the instance
func (instance *ads_ipsecnfa) GetPropertyDS_ipsecNegotiationPolicyReference() (value string, err error) {
	retValue, err := instance.GetProperty("DS_ipsecNegotiationPolicyReference")
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
