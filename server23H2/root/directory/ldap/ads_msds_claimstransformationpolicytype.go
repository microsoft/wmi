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

// ads_msds_claimstransformationpolicytype struct
type ads_msds_claimstransformationpolicytype struct {
	*ds_top

	//
	DS_msDS_TransformationRules string

	//
	DS_msDS_TransformationRulesCompiled Uint8Array
}

func Newads_msds_claimstransformationpolicytypeEx1(instance *cim.WmiInstance) (newInstance *ads_msds_claimstransformationpolicytype, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_claimstransformationpolicytype{
		ds_top: tmp,
	}
	return
}

func Newads_msds_claimstransformationpolicytypeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_claimstransformationpolicytype, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_claimstransformationpolicytype{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_TransformationRules sets the value of DS_msDS_TransformationRules for the instance
func (instance *ads_msds_claimstransformationpolicytype) SetPropertyDS_msDS_TransformationRules(value string) (err error) {
	return instance.SetProperty("DS_msDS_TransformationRules", (value))
}

// GetDS_msDS_TransformationRules gets the value of DS_msDS_TransformationRules for the instance
func (instance *ads_msds_claimstransformationpolicytype) GetPropertyDS_msDS_TransformationRules() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_TransformationRules")
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

// SetDS_msDS_TransformationRulesCompiled sets the value of DS_msDS_TransformationRulesCompiled for the instance
func (instance *ads_msds_claimstransformationpolicytype) SetPropertyDS_msDS_TransformationRulesCompiled(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_TransformationRulesCompiled", (value))
}

// GetDS_msDS_TransformationRulesCompiled gets the value of DS_msDS_TransformationRulesCompiled for the instance
func (instance *ads_msds_claimstransformationpolicytype) GetPropertyDS_msDS_TransformationRulesCompiled() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_TransformationRulesCompiled")
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
