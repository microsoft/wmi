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

// ads_msmqmigrateduser struct
type ads_msmqmigrateduser struct {
	*ds_top

	//
	DS_mSMQDigests []Uint8Array

	//
	DS_mSMQDigestsMig []Uint8Array

	//
	DS_mSMQSignCertificates Uint8Array

	//
	DS_mSMQSignCertificatesMig Uint8Array

	//
	DS_mSMQUserSid Uint8Array

	//
	DS_objectSid Uint8Array
}

func Newads_msmqmigrateduserEx1(instance *cim.WmiInstance) (newInstance *ads_msmqmigrateduser, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msmqmigrateduser{
		ds_top: tmp,
	}
	return
}

func Newads_msmqmigrateduserEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msmqmigrateduser, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msmqmigrateduser{
		ds_top: tmp,
	}
	return
}

// SetDS_mSMQDigests sets the value of DS_mSMQDigests for the instance
func (instance *ads_msmqmigrateduser) SetPropertyDS_mSMQDigests(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQDigests", (value))
}

// GetDS_mSMQDigests gets the value of DS_mSMQDigests for the instance
func (instance *ads_msmqmigrateduser) GetPropertyDS_mSMQDigests() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQDigests")
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

// SetDS_mSMQDigestsMig sets the value of DS_mSMQDigestsMig for the instance
func (instance *ads_msmqmigrateduser) SetPropertyDS_mSMQDigestsMig(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQDigestsMig", (value))
}

// GetDS_mSMQDigestsMig gets the value of DS_mSMQDigestsMig for the instance
func (instance *ads_msmqmigrateduser) GetPropertyDS_mSMQDigestsMig() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQDigestsMig")
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

// SetDS_mSMQSignCertificates sets the value of DS_mSMQSignCertificates for the instance
func (instance *ads_msmqmigrateduser) SetPropertyDS_mSMQSignCertificates(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQSignCertificates", (value))
}

// GetDS_mSMQSignCertificates gets the value of DS_mSMQSignCertificates for the instance
func (instance *ads_msmqmigrateduser) GetPropertyDS_mSMQSignCertificates() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQSignCertificates")
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

// SetDS_mSMQSignCertificatesMig sets the value of DS_mSMQSignCertificatesMig for the instance
func (instance *ads_msmqmigrateduser) SetPropertyDS_mSMQSignCertificatesMig(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQSignCertificatesMig", (value))
}

// GetDS_mSMQSignCertificatesMig gets the value of DS_mSMQSignCertificatesMig for the instance
func (instance *ads_msmqmigrateduser) GetPropertyDS_mSMQSignCertificatesMig() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQSignCertificatesMig")
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

// SetDS_mSMQUserSid sets the value of DS_mSMQUserSid for the instance
func (instance *ads_msmqmigrateduser) SetPropertyDS_mSMQUserSid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQUserSid", (value))
}

// GetDS_mSMQUserSid gets the value of DS_mSMQUserSid for the instance
func (instance *ads_msmqmigrateduser) GetPropertyDS_mSMQUserSid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQUserSid")
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

// SetDS_objectSid sets the value of DS_objectSid for the instance
func (instance *ads_msmqmigrateduser) SetPropertyDS_objectSid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_objectSid", (value))
}

// GetDS_objectSid gets the value of DS_objectSid for the instance
func (instance *ads_msmqmigrateduser) GetPropertyDS_objectSid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_objectSid")
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
