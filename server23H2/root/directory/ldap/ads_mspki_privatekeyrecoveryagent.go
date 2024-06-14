// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_mspki_privatekeyrecoveryagent struct
type ads_mspki_privatekeyrecoveryagent struct {
	*ds_top

	//
	DS_userCertificate []Uint8Array
}

func Newads_mspki_privatekeyrecoveryagentEx1(instance *cim.WmiInstance) (newInstance *ads_mspki_privatekeyrecoveryagent, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mspki_privatekeyrecoveryagent{
		ds_top: tmp,
	}
	return
}

func Newads_mspki_privatekeyrecoveryagentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mspki_privatekeyrecoveryagent, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mspki_privatekeyrecoveryagent{
		ds_top: tmp,
	}
	return
}

// SetDS_userCertificate sets the value of DS_userCertificate for the instance
func (instance *ads_mspki_privatekeyrecoveryagent) SetPropertyDS_userCertificate(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_userCertificate", (value))
}

// GetDS_userCertificate gets the value of DS_userCertificate for the instance
func (instance *ads_mspki_privatekeyrecoveryagent) GetPropertyDS_userCertificate() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_userCertificate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}
