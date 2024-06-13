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

// ads_mskds_provserverconfiguration struct
type ads_mskds_provserverconfiguration struct {
	*ds_top

	//
	DS_msKds_KDFAlgorithmID string

	//
	DS_msKds_KDFParam Uint8Array

	//
	DS_msKds_PrivateKeyLength int32

	//
	DS_msKds_PublicKeyLength int32

	//
	DS_msKds_SecretAgreementAlgorithmID string

	//
	DS_msKds_SecretAgreementParam Uint8Array

	//
	DS_msKds_Version int32
}

func Newads_mskds_provserverconfigurationEx1(instance *cim.WmiInstance) (newInstance *ads_mskds_provserverconfiguration, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mskds_provserverconfiguration{
		ds_top: tmp,
	}
	return
}

func Newads_mskds_provserverconfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mskds_provserverconfiguration, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mskds_provserverconfiguration{
		ds_top: tmp,
	}
	return
}

// SetDS_msKds_KDFAlgorithmID sets the value of DS_msKds_KDFAlgorithmID for the instance
func (instance *ads_mskds_provserverconfiguration) SetPropertyDS_msKds_KDFAlgorithmID(value string) (err error) {
	return instance.SetProperty("DS_msKds_KDFAlgorithmID", (value))
}

// GetDS_msKds_KDFAlgorithmID gets the value of DS_msKds_KDFAlgorithmID for the instance
func (instance *ads_mskds_provserverconfiguration) GetPropertyDS_msKds_KDFAlgorithmID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msKds_KDFAlgorithmID")
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

// SetDS_msKds_KDFParam sets the value of DS_msKds_KDFParam for the instance
func (instance *ads_mskds_provserverconfiguration) SetPropertyDS_msKds_KDFParam(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msKds_KDFParam", (value))
}

// GetDS_msKds_KDFParam gets the value of DS_msKds_KDFParam for the instance
func (instance *ads_mskds_provserverconfiguration) GetPropertyDS_msKds_KDFParam() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msKds_KDFParam")
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

// SetDS_msKds_PrivateKeyLength sets the value of DS_msKds_PrivateKeyLength for the instance
func (instance *ads_mskds_provserverconfiguration) SetPropertyDS_msKds_PrivateKeyLength(value int32) (err error) {
	return instance.SetProperty("DS_msKds_PrivateKeyLength", (value))
}

// GetDS_msKds_PrivateKeyLength gets the value of DS_msKds_PrivateKeyLength for the instance
func (instance *ads_mskds_provserverconfiguration) GetPropertyDS_msKds_PrivateKeyLength() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msKds_PrivateKeyLength")
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

// SetDS_msKds_PublicKeyLength sets the value of DS_msKds_PublicKeyLength for the instance
func (instance *ads_mskds_provserverconfiguration) SetPropertyDS_msKds_PublicKeyLength(value int32) (err error) {
	return instance.SetProperty("DS_msKds_PublicKeyLength", (value))
}

// GetDS_msKds_PublicKeyLength gets the value of DS_msKds_PublicKeyLength for the instance
func (instance *ads_mskds_provserverconfiguration) GetPropertyDS_msKds_PublicKeyLength() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msKds_PublicKeyLength")
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

// SetDS_msKds_SecretAgreementAlgorithmID sets the value of DS_msKds_SecretAgreementAlgorithmID for the instance
func (instance *ads_mskds_provserverconfiguration) SetPropertyDS_msKds_SecretAgreementAlgorithmID(value string) (err error) {
	return instance.SetProperty("DS_msKds_SecretAgreementAlgorithmID", (value))
}

// GetDS_msKds_SecretAgreementAlgorithmID gets the value of DS_msKds_SecretAgreementAlgorithmID for the instance
func (instance *ads_mskds_provserverconfiguration) GetPropertyDS_msKds_SecretAgreementAlgorithmID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msKds_SecretAgreementAlgorithmID")
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

// SetDS_msKds_SecretAgreementParam sets the value of DS_msKds_SecretAgreementParam for the instance
func (instance *ads_mskds_provserverconfiguration) SetPropertyDS_msKds_SecretAgreementParam(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msKds_SecretAgreementParam", (value))
}

// GetDS_msKds_SecretAgreementParam gets the value of DS_msKds_SecretAgreementParam for the instance
func (instance *ads_mskds_provserverconfiguration) GetPropertyDS_msKds_SecretAgreementParam() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msKds_SecretAgreementParam")
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

// SetDS_msKds_Version sets the value of DS_msKds_Version for the instance
func (instance *ads_mskds_provserverconfiguration) SetPropertyDS_msKds_Version(value int32) (err error) {
	return instance.SetProperty("DS_msKds_Version", (value))
}

// GetDS_msKds_Version gets the value of DS_msKds_Version for the instance
func (instance *ads_mskds_provserverconfiguration) GetPropertyDS_msKds_Version() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msKds_Version")
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
