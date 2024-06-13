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

// ads_msauthz_centralaccesspolicy struct
type ads_msauthz_centralaccesspolicy struct {
	*ds_top

	//
	DS_msAuthz_CentralAccessPolicyID Uint8Array

	//
	DS_msAuthz_MemberRulesInCentralAccessPolicy []string
}

func Newads_msauthz_centralaccesspolicyEx1(instance *cim.WmiInstance) (newInstance *ads_msauthz_centralaccesspolicy, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msauthz_centralaccesspolicy{
		ds_top: tmp,
	}
	return
}

func Newads_msauthz_centralaccesspolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msauthz_centralaccesspolicy, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msauthz_centralaccesspolicy{
		ds_top: tmp,
	}
	return
}

// SetDS_msAuthz_CentralAccessPolicyID sets the value of DS_msAuthz_CentralAccessPolicyID for the instance
func (instance *ads_msauthz_centralaccesspolicy) SetPropertyDS_msAuthz_CentralAccessPolicyID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msAuthz_CentralAccessPolicyID", (value))
}

// GetDS_msAuthz_CentralAccessPolicyID gets the value of DS_msAuthz_CentralAccessPolicyID for the instance
func (instance *ads_msauthz_centralaccesspolicy) GetPropertyDS_msAuthz_CentralAccessPolicyID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msAuthz_CentralAccessPolicyID")
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

// SetDS_msAuthz_MemberRulesInCentralAccessPolicy sets the value of DS_msAuthz_MemberRulesInCentralAccessPolicy for the instance
func (instance *ads_msauthz_centralaccesspolicy) SetPropertyDS_msAuthz_MemberRulesInCentralAccessPolicy(value []string) (err error) {
	return instance.SetProperty("DS_msAuthz_MemberRulesInCentralAccessPolicy", (value))
}

// GetDS_msAuthz_MemberRulesInCentralAccessPolicy gets the value of DS_msAuthz_MemberRulesInCentralAccessPolicy for the instance
func (instance *ads_msauthz_centralaccesspolicy) GetPropertyDS_msAuthz_MemberRulesInCentralAccessPolicy() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msAuthz_MemberRulesInCentralAccessPolicy")
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
