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

// ads_pkienrollmentservice struct
type ads_pkienrollmentservice struct {
	*ds_top

	//
	DS_cACertificate []Uint8Array

	//
	DS_cACertificateDN string

	//
	DS_certificateTemplates []string

	//
	DS_dNSHostName string

	//
	DS_enrollmentProviders string

	//
	DS_msPKI_Enrollment_Servers []string

	//
	DS_msPKI_Site_Name string

	//
	DS_signatureAlgorithms string
}

func Newads_pkienrollmentserviceEx1(instance *cim.WmiInstance) (newInstance *ads_pkienrollmentservice, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_pkienrollmentservice{
		ds_top: tmp,
	}
	return
}

func Newads_pkienrollmentserviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_pkienrollmentservice, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_pkienrollmentservice{
		ds_top: tmp,
	}
	return
}

// SetDS_cACertificate sets the value of DS_cACertificate for the instance
func (instance *ads_pkienrollmentservice) SetPropertyDS_cACertificate(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_cACertificate", (value))
}

// GetDS_cACertificate gets the value of DS_cACertificate for the instance
func (instance *ads_pkienrollmentservice) GetPropertyDS_cACertificate() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_cACertificate")
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

// SetDS_cACertificateDN sets the value of DS_cACertificateDN for the instance
func (instance *ads_pkienrollmentservice) SetPropertyDS_cACertificateDN(value string) (err error) {
	return instance.SetProperty("DS_cACertificateDN", (value))
}

// GetDS_cACertificateDN gets the value of DS_cACertificateDN for the instance
func (instance *ads_pkienrollmentservice) GetPropertyDS_cACertificateDN() (value string, err error) {
	retValue, err := instance.GetProperty("DS_cACertificateDN")
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

// SetDS_certificateTemplates sets the value of DS_certificateTemplates for the instance
func (instance *ads_pkienrollmentservice) SetPropertyDS_certificateTemplates(value []string) (err error) {
	return instance.SetProperty("DS_certificateTemplates", (value))
}

// GetDS_certificateTemplates gets the value of DS_certificateTemplates for the instance
func (instance *ads_pkienrollmentservice) GetPropertyDS_certificateTemplates() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_certificateTemplates")
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

// SetDS_dNSHostName sets the value of DS_dNSHostName for the instance
func (instance *ads_pkienrollmentservice) SetPropertyDS_dNSHostName(value string) (err error) {
	return instance.SetProperty("DS_dNSHostName", (value))
}

// GetDS_dNSHostName gets the value of DS_dNSHostName for the instance
func (instance *ads_pkienrollmentservice) GetPropertyDS_dNSHostName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_dNSHostName")
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

// SetDS_enrollmentProviders sets the value of DS_enrollmentProviders for the instance
func (instance *ads_pkienrollmentservice) SetPropertyDS_enrollmentProviders(value string) (err error) {
	return instance.SetProperty("DS_enrollmentProviders", (value))
}

// GetDS_enrollmentProviders gets the value of DS_enrollmentProviders for the instance
func (instance *ads_pkienrollmentservice) GetPropertyDS_enrollmentProviders() (value string, err error) {
	retValue, err := instance.GetProperty("DS_enrollmentProviders")
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

// SetDS_msPKI_Enrollment_Servers sets the value of DS_msPKI_Enrollment_Servers for the instance
func (instance *ads_pkienrollmentservice) SetPropertyDS_msPKI_Enrollment_Servers(value []string) (err error) {
	return instance.SetProperty("DS_msPKI_Enrollment_Servers", (value))
}

// GetDS_msPKI_Enrollment_Servers gets the value of DS_msPKI_Enrollment_Servers for the instance
func (instance *ads_pkienrollmentservice) GetPropertyDS_msPKI_Enrollment_Servers() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msPKI_Enrollment_Servers")
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

// SetDS_msPKI_Site_Name sets the value of DS_msPKI_Site_Name for the instance
func (instance *ads_pkienrollmentservice) SetPropertyDS_msPKI_Site_Name(value string) (err error) {
	return instance.SetProperty("DS_msPKI_Site_Name", (value))
}

// GetDS_msPKI_Site_Name gets the value of DS_msPKI_Site_Name for the instance
func (instance *ads_pkienrollmentservice) GetPropertyDS_msPKI_Site_Name() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msPKI_Site_Name")
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

// SetDS_signatureAlgorithms sets the value of DS_signatureAlgorithms for the instance
func (instance *ads_pkienrollmentservice) SetPropertyDS_signatureAlgorithms(value string) (err error) {
	return instance.SetProperty("DS_signatureAlgorithms", (value))
}

// GetDS_signatureAlgorithms gets the value of DS_signatureAlgorithms for the instance
func (instance *ads_pkienrollmentservice) GetPropertyDS_signatureAlgorithms() (value string, err error) {
	retValue, err := instance.GetProperty("DS_signatureAlgorithms")
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
