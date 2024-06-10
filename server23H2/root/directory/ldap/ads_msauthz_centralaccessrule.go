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

// ads_msauthz_centralaccessrule struct
type ads_msauthz_centralaccessrule struct {
	*ds_top

	//
	DS_Enabled bool

	//
	DS_msAuthz_EffectiveSecurityPolicy string

	//
	DS_msAuthz_LastEffectiveSecurityPolicy string

	//
	DS_msAuthz_MemberRulesInCentralAccessPolicyBL []string

	//
	DS_msAuthz_ProposedSecurityPolicy string

	//
	DS_msAuthz_ResourceCondition string
}

func Newads_msauthz_centralaccessruleEx1(instance *cim.WmiInstance) (newInstance *ads_msauthz_centralaccessrule, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msauthz_centralaccessrule{
		ds_top: tmp,
	}
	return
}

func Newads_msauthz_centralaccessruleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msauthz_centralaccessrule, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msauthz_centralaccessrule{
		ds_top: tmp,
	}
	return
}

// SetDS_Enabled sets the value of DS_Enabled for the instance
func (instance *ads_msauthz_centralaccessrule) SetPropertyDS_Enabled(value bool) (err error) {
	return instance.SetProperty("DS_Enabled", (value))
}

// GetDS_Enabled gets the value of DS_Enabled for the instance
func (instance *ads_msauthz_centralaccessrule) GetPropertyDS_Enabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_Enabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDS_msAuthz_EffectiveSecurityPolicy sets the value of DS_msAuthz_EffectiveSecurityPolicy for the instance
func (instance *ads_msauthz_centralaccessrule) SetPropertyDS_msAuthz_EffectiveSecurityPolicy(value string) (err error) {
	return instance.SetProperty("DS_msAuthz_EffectiveSecurityPolicy", (value))
}

// GetDS_msAuthz_EffectiveSecurityPolicy gets the value of DS_msAuthz_EffectiveSecurityPolicy for the instance
func (instance *ads_msauthz_centralaccessrule) GetPropertyDS_msAuthz_EffectiveSecurityPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msAuthz_EffectiveSecurityPolicy")
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

// SetDS_msAuthz_LastEffectiveSecurityPolicy sets the value of DS_msAuthz_LastEffectiveSecurityPolicy for the instance
func (instance *ads_msauthz_centralaccessrule) SetPropertyDS_msAuthz_LastEffectiveSecurityPolicy(value string) (err error) {
	return instance.SetProperty("DS_msAuthz_LastEffectiveSecurityPolicy", (value))
}

// GetDS_msAuthz_LastEffectiveSecurityPolicy gets the value of DS_msAuthz_LastEffectiveSecurityPolicy for the instance
func (instance *ads_msauthz_centralaccessrule) GetPropertyDS_msAuthz_LastEffectiveSecurityPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msAuthz_LastEffectiveSecurityPolicy")
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

// SetDS_msAuthz_MemberRulesInCentralAccessPolicyBL sets the value of DS_msAuthz_MemberRulesInCentralAccessPolicyBL for the instance
func (instance *ads_msauthz_centralaccessrule) SetPropertyDS_msAuthz_MemberRulesInCentralAccessPolicyBL(value []string) (err error) {
	return instance.SetProperty("DS_msAuthz_MemberRulesInCentralAccessPolicyBL", (value))
}

// GetDS_msAuthz_MemberRulesInCentralAccessPolicyBL gets the value of DS_msAuthz_MemberRulesInCentralAccessPolicyBL for the instance
func (instance *ads_msauthz_centralaccessrule) GetPropertyDS_msAuthz_MemberRulesInCentralAccessPolicyBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msAuthz_MemberRulesInCentralAccessPolicyBL")
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

// SetDS_msAuthz_ProposedSecurityPolicy sets the value of DS_msAuthz_ProposedSecurityPolicy for the instance
func (instance *ads_msauthz_centralaccessrule) SetPropertyDS_msAuthz_ProposedSecurityPolicy(value string) (err error) {
	return instance.SetProperty("DS_msAuthz_ProposedSecurityPolicy", (value))
}

// GetDS_msAuthz_ProposedSecurityPolicy gets the value of DS_msAuthz_ProposedSecurityPolicy for the instance
func (instance *ads_msauthz_centralaccessrule) GetPropertyDS_msAuthz_ProposedSecurityPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msAuthz_ProposedSecurityPolicy")
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

// SetDS_msAuthz_ResourceCondition sets the value of DS_msAuthz_ResourceCondition for the instance
func (instance *ads_msauthz_centralaccessrule) SetPropertyDS_msAuthz_ResourceCondition(value string) (err error) {
	return instance.SetProperty("DS_msAuthz_ResourceCondition", (value))
}

// GetDS_msAuthz_ResourceCondition gets the value of DS_msAuthz_ResourceCondition for the instance
func (instance *ads_msauthz_centralaccessrule) GetPropertyDS_msAuthz_ResourceCondition() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msAuthz_ResourceCondition")
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
