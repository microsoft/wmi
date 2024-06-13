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

// ads_remotemailrecipient struct
type ads_remotemailrecipient struct {
	*ds_top

	//
	DS_garbageCollPeriod int32

	//
	DS_info string

	//
	DS_labeledURI []string

	//
	DS_legacyExchangeDN string

	//
	DS_managedBy string

	//
	DS_msDS_ExternalDirectoryObjectId string

	//
	DS_msDS_GeoCoordinatesAltitude int64

	//
	DS_msDS_GeoCoordinatesLatitude int64

	//
	DS_msDS_GeoCoordinatesLongitude int64

	//
	DS_msDS_PhoneticDisplayName string

	//
	DS_msExchAssistantName string

	//
	DS_msExchLabeledURI []string

	//
	DS_remoteSource string

	//
	DS_remoteSourceType int32

	//
	DS_secretary []string

	//
	DS_showInAddressBook []string

	//
	DS_telephoneNumber string

	//
	DS_textEncodedORAddress string

	//
	DS_userCert Uint8Array

	//
	DS_userCertificate []Uint8Array

	//
	DS_userSMIMECertificate []Uint8Array
}

func Newads_remotemailrecipientEx1(instance *cim.WmiInstance) (newInstance *ads_remotemailrecipient, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_remotemailrecipient{
		ds_top: tmp,
	}
	return
}

func Newads_remotemailrecipientEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_remotemailrecipient, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_remotemailrecipient{
		ds_top: tmp,
	}
	return
}

// SetDS_garbageCollPeriod sets the value of DS_garbageCollPeriod for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_garbageCollPeriod(value int32) (err error) {
	return instance.SetProperty("DS_garbageCollPeriod", (value))
}

// GetDS_garbageCollPeriod gets the value of DS_garbageCollPeriod for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_garbageCollPeriod() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_garbageCollPeriod")
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

// SetDS_info sets the value of DS_info for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_info(value string) (err error) {
	return instance.SetProperty("DS_info", (value))
}

// GetDS_info gets the value of DS_info for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_info() (value string, err error) {
	retValue, err := instance.GetProperty("DS_info")
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

// SetDS_labeledURI sets the value of DS_labeledURI for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_labeledURI(value []string) (err error) {
	return instance.SetProperty("DS_labeledURI", (value))
}

// GetDS_labeledURI gets the value of DS_labeledURI for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_labeledURI() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_labeledURI")
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

// SetDS_legacyExchangeDN sets the value of DS_legacyExchangeDN for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_legacyExchangeDN(value string) (err error) {
	return instance.SetProperty("DS_legacyExchangeDN", (value))
}

// GetDS_legacyExchangeDN gets the value of DS_legacyExchangeDN for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_legacyExchangeDN() (value string, err error) {
	retValue, err := instance.GetProperty("DS_legacyExchangeDN")
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_managedBy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_managedBy")
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

// SetDS_msDS_ExternalDirectoryObjectId sets the value of DS_msDS_ExternalDirectoryObjectId for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_msDS_ExternalDirectoryObjectId(value string) (err error) {
	return instance.SetProperty("DS_msDS_ExternalDirectoryObjectId", (value))
}

// GetDS_msDS_ExternalDirectoryObjectId gets the value of DS_msDS_ExternalDirectoryObjectId for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_msDS_ExternalDirectoryObjectId() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ExternalDirectoryObjectId")
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

// SetDS_msDS_GeoCoordinatesAltitude sets the value of DS_msDS_GeoCoordinatesAltitude for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_msDS_GeoCoordinatesAltitude(value int64) (err error) {
	return instance.SetProperty("DS_msDS_GeoCoordinatesAltitude", (value))
}

// GetDS_msDS_GeoCoordinatesAltitude gets the value of DS_msDS_GeoCoordinatesAltitude for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_msDS_GeoCoordinatesAltitude() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_GeoCoordinatesAltitude")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_msDS_GeoCoordinatesLatitude sets the value of DS_msDS_GeoCoordinatesLatitude for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_msDS_GeoCoordinatesLatitude(value int64) (err error) {
	return instance.SetProperty("DS_msDS_GeoCoordinatesLatitude", (value))
}

// GetDS_msDS_GeoCoordinatesLatitude gets the value of DS_msDS_GeoCoordinatesLatitude for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_msDS_GeoCoordinatesLatitude() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_GeoCoordinatesLatitude")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_msDS_GeoCoordinatesLongitude sets the value of DS_msDS_GeoCoordinatesLongitude for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_msDS_GeoCoordinatesLongitude(value int64) (err error) {
	return instance.SetProperty("DS_msDS_GeoCoordinatesLongitude", (value))
}

// GetDS_msDS_GeoCoordinatesLongitude gets the value of DS_msDS_GeoCoordinatesLongitude for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_msDS_GeoCoordinatesLongitude() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_GeoCoordinatesLongitude")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_msDS_PhoneticDisplayName sets the value of DS_msDS_PhoneticDisplayName for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_msDS_PhoneticDisplayName(value string) (err error) {
	return instance.SetProperty("DS_msDS_PhoneticDisplayName", (value))
}

// GetDS_msDS_PhoneticDisplayName gets the value of DS_msDS_PhoneticDisplayName for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_msDS_PhoneticDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_PhoneticDisplayName")
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

// SetDS_msExchAssistantName sets the value of DS_msExchAssistantName for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_msExchAssistantName(value string) (err error) {
	return instance.SetProperty("DS_msExchAssistantName", (value))
}

// GetDS_msExchAssistantName gets the value of DS_msExchAssistantName for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_msExchAssistantName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msExchAssistantName")
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

// SetDS_msExchLabeledURI sets the value of DS_msExchLabeledURI for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_msExchLabeledURI(value []string) (err error) {
	return instance.SetProperty("DS_msExchLabeledURI", (value))
}

// GetDS_msExchLabeledURI gets the value of DS_msExchLabeledURI for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_msExchLabeledURI() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msExchLabeledURI")
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

// SetDS_remoteSource sets the value of DS_remoteSource for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_remoteSource(value string) (err error) {
	return instance.SetProperty("DS_remoteSource", (value))
}

// GetDS_remoteSource gets the value of DS_remoteSource for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_remoteSource() (value string, err error) {
	retValue, err := instance.GetProperty("DS_remoteSource")
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

// SetDS_remoteSourceType sets the value of DS_remoteSourceType for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_remoteSourceType(value int32) (err error) {
	return instance.SetProperty("DS_remoteSourceType", (value))
}

// GetDS_remoteSourceType gets the value of DS_remoteSourceType for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_remoteSourceType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_remoteSourceType")
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

// SetDS_secretary sets the value of DS_secretary for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_secretary(value []string) (err error) {
	return instance.SetProperty("DS_secretary", (value))
}

// GetDS_secretary gets the value of DS_secretary for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_secretary() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_secretary")
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

// SetDS_showInAddressBook sets the value of DS_showInAddressBook for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_showInAddressBook(value []string) (err error) {
	return instance.SetProperty("DS_showInAddressBook", (value))
}

// GetDS_showInAddressBook gets the value of DS_showInAddressBook for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_showInAddressBook() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_showInAddressBook")
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

// SetDS_telephoneNumber sets the value of DS_telephoneNumber for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_telephoneNumber(value string) (err error) {
	return instance.SetProperty("DS_telephoneNumber", (value))
}

// GetDS_telephoneNumber gets the value of DS_telephoneNumber for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_telephoneNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DS_telephoneNumber")
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

// SetDS_textEncodedORAddress sets the value of DS_textEncodedORAddress for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_textEncodedORAddress(value string) (err error) {
	return instance.SetProperty("DS_textEncodedORAddress", (value))
}

// GetDS_textEncodedORAddress gets the value of DS_textEncodedORAddress for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_textEncodedORAddress() (value string, err error) {
	retValue, err := instance.GetProperty("DS_textEncodedORAddress")
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

// SetDS_userCert sets the value of DS_userCert for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_userCert(value Uint8Array) (err error) {
	return instance.SetProperty("DS_userCert", (value))
}

// GetDS_userCert gets the value of DS_userCert for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_userCert() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_userCert")
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

// SetDS_userCertificate sets the value of DS_userCertificate for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_userCertificate(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_userCertificate", (value))
}

// GetDS_userCertificate gets the value of DS_userCertificate for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_userCertificate() (value []Uint8Array, err error) {
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

// SetDS_userSMIMECertificate sets the value of DS_userSMIMECertificate for the instance
func (instance *ads_remotemailrecipient) SetPropertyDS_userSMIMECertificate(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_userSMIMECertificate", (value))
}

// GetDS_userSMIMECertificate gets the value of DS_userSMIMECertificate for the instance
func (instance *ads_remotemailrecipient) GetPropertyDS_userSMIMECertificate() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_userSMIMECertificate")
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
