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

// ads_mspki_enterprise_oid struct
type ads_mspki_enterprise_oid struct {
	*ds_top

	//
	DS_msDS_OIDToGroupLink string

	//
	DS_msPKI_Cert_Template_OID string

	//
	DS_msPKI_OID_Attribute int32

	//
	DS_msPKI_OID_CPS []string

	//
	DS_msPKI_OID_User_Notice []string

	//
	DS_msPKI_OIDLocalizedName []string
}

func Newads_mspki_enterprise_oidEx1(instance *cim.WmiInstance) (newInstance *ads_mspki_enterprise_oid, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mspki_enterprise_oid{
		ds_top: tmp,
	}
	return
}

func Newads_mspki_enterprise_oidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mspki_enterprise_oid, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mspki_enterprise_oid{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_OIDToGroupLink sets the value of DS_msDS_OIDToGroupLink for the instance
func (instance *ads_mspki_enterprise_oid) SetPropertyDS_msDS_OIDToGroupLink(value string) (err error) {
	return instance.SetProperty("DS_msDS_OIDToGroupLink", (value))
}

// GetDS_msDS_OIDToGroupLink gets the value of DS_msDS_OIDToGroupLink for the instance
func (instance *ads_mspki_enterprise_oid) GetPropertyDS_msDS_OIDToGroupLink() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_OIDToGroupLink")
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

// SetDS_msPKI_Cert_Template_OID sets the value of DS_msPKI_Cert_Template_OID for the instance
func (instance *ads_mspki_enterprise_oid) SetPropertyDS_msPKI_Cert_Template_OID(value string) (err error) {
	return instance.SetProperty("DS_msPKI_Cert_Template_OID", (value))
}

// GetDS_msPKI_Cert_Template_OID gets the value of DS_msPKI_Cert_Template_OID for the instance
func (instance *ads_mspki_enterprise_oid) GetPropertyDS_msPKI_Cert_Template_OID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msPKI_Cert_Template_OID")
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

// SetDS_msPKI_OID_Attribute sets the value of DS_msPKI_OID_Attribute for the instance
func (instance *ads_mspki_enterprise_oid) SetPropertyDS_msPKI_OID_Attribute(value int32) (err error) {
	return instance.SetProperty("DS_msPKI_OID_Attribute", (value))
}

// GetDS_msPKI_OID_Attribute gets the value of DS_msPKI_OID_Attribute for the instance
func (instance *ads_mspki_enterprise_oid) GetPropertyDS_msPKI_OID_Attribute() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msPKI_OID_Attribute")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetDS_msPKI_OID_CPS sets the value of DS_msPKI_OID_CPS for the instance
func (instance *ads_mspki_enterprise_oid) SetPropertyDS_msPKI_OID_CPS(value []string) (err error) {
	return instance.SetProperty("DS_msPKI_OID_CPS", (value))
}

// GetDS_msPKI_OID_CPS gets the value of DS_msPKI_OID_CPS for the instance
func (instance *ads_mspki_enterprise_oid) GetPropertyDS_msPKI_OID_CPS() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msPKI_OID_CPS")
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

// SetDS_msPKI_OID_User_Notice sets the value of DS_msPKI_OID_User_Notice for the instance
func (instance *ads_mspki_enterprise_oid) SetPropertyDS_msPKI_OID_User_Notice(value []string) (err error) {
	return instance.SetProperty("DS_msPKI_OID_User_Notice", (value))
}

// GetDS_msPKI_OID_User_Notice gets the value of DS_msPKI_OID_User_Notice for the instance
func (instance *ads_mspki_enterprise_oid) GetPropertyDS_msPKI_OID_User_Notice() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msPKI_OID_User_Notice")
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

// SetDS_msPKI_OIDLocalizedName sets the value of DS_msPKI_OIDLocalizedName for the instance
func (instance *ads_mspki_enterprise_oid) SetPropertyDS_msPKI_OIDLocalizedName(value []string) (err error) {
	return instance.SetProperty("DS_msPKI_OIDLocalizedName", (value))
}

// GetDS_msPKI_OIDLocalizedName gets the value of DS_msPKI_OIDLocalizedName for the instance
func (instance *ads_mspki_enterprise_oid) GetPropertyDS_msPKI_OIDLocalizedName() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msPKI_OIDLocalizedName")
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
