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

// ads_ipsecpolicy struct
type ads_ipsecpolicy struct {
	*ds_ipsecbase

	//
	DS_ipsecISAKMPReference string

	//
	DS_ipsecNFAReference []string
}

func Newads_ipsecpolicyEx1(instance *cim.WmiInstance) (newInstance *ads_ipsecpolicy, err error) {
	tmp, err := Newds_ipsecbaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ipsecpolicy{
		ds_ipsecbase: tmp,
	}
	return
}

func Newads_ipsecpolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ipsecpolicy, err error) {
	tmp, err := Newds_ipsecbaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ipsecpolicy{
		ds_ipsecbase: tmp,
	}
	return
}

// SetDS_ipsecISAKMPReference sets the value of DS_ipsecISAKMPReference for the instance
func (instance *ads_ipsecpolicy) SetPropertyDS_ipsecISAKMPReference(value string) (err error) {
	return instance.SetProperty("DS_ipsecISAKMPReference", (value))
}

// GetDS_ipsecISAKMPReference gets the value of DS_ipsecISAKMPReference for the instance
func (instance *ads_ipsecpolicy) GetPropertyDS_ipsecISAKMPReference() (value string, err error) {
	retValue, err := instance.GetProperty("DS_ipsecISAKMPReference")
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

// SetDS_ipsecNFAReference sets the value of DS_ipsecNFAReference for the instance
func (instance *ads_ipsecpolicy) SetPropertyDS_ipsecNFAReference(value []string) (err error) {
	return instance.SetProperty("DS_ipsecNFAReference", (value))
}

// GetDS_ipsecNFAReference gets the value of DS_ipsecNFAReference for the instance
func (instance *ads_ipsecpolicy) GetPropertyDS_ipsecNFAReference() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_ipsecNFAReference")
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
