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

// ads_ipsecnegotiationpolicy struct
type ads_ipsecnegotiationpolicy struct {
	*ds_ipsecbase

	//
	DS_iPSECNegotiationPolicyAction string

	//
	DS_iPSECNegotiationPolicyType string
}

func Newads_ipsecnegotiationpolicyEx1(instance *cim.WmiInstance) (newInstance *ads_ipsecnegotiationpolicy, err error) {
	tmp, err := Newds_ipsecbaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ipsecnegotiationpolicy{
		ds_ipsecbase: tmp,
	}
	return
}

func Newads_ipsecnegotiationpolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ipsecnegotiationpolicy, err error) {
	tmp, err := Newds_ipsecbaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ipsecnegotiationpolicy{
		ds_ipsecbase: tmp,
	}
	return
}

// SetDS_iPSECNegotiationPolicyAction sets the value of DS_iPSECNegotiationPolicyAction for the instance
func (instance *ads_ipsecnegotiationpolicy) SetPropertyDS_iPSECNegotiationPolicyAction(value string) (err error) {
	return instance.SetProperty("DS_iPSECNegotiationPolicyAction", (value))
}

// GetDS_iPSECNegotiationPolicyAction gets the value of DS_iPSECNegotiationPolicyAction for the instance
func (instance *ads_ipsecnegotiationpolicy) GetPropertyDS_iPSECNegotiationPolicyAction() (value string, err error) {
	retValue, err := instance.GetProperty("DS_iPSECNegotiationPolicyAction")
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

// SetDS_iPSECNegotiationPolicyType sets the value of DS_iPSECNegotiationPolicyType for the instance
func (instance *ads_ipsecnegotiationpolicy) SetPropertyDS_iPSECNegotiationPolicyType(value string) (err error) {
	return instance.SetProperty("DS_iPSECNegotiationPolicyType", (value))
}

// GetDS_iPSECNegotiationPolicyType gets the value of DS_iPSECNegotiationPolicyType for the instance
func (instance *ads_ipsecnegotiationpolicy) GetPropertyDS_iPSECNegotiationPolicyType() (value string, err error) {
	retValue, err := instance.GetProperty("DS_iPSECNegotiationPolicyType")
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
