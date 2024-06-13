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

// ads_certificationauthority struct
type ads_certificationauthority struct {
	*ds_top

	//
	DS_authorityRevocationList []Uint8Array

	//
	DS_cACertificate []Uint8Array

	//
	DS_cACertificateDN string

	//
	DS_cAConnect string

	//
	DS_cAUsages []string

	//
	DS_cAWEBURL string

	//
	DS_certificateRevocationList Uint8Array

	//
	DS_certificateTemplates []string

	//
	DS_cRLObject string

	//
	DS_crossCertificatePair []Uint8Array

	//
	DS_currentParentCA []string

	//
	DS_deltaRevocationList []Uint8Array

	//
	DS_dNSHostName string

	//
	DS_domainID string

	//
	DS_domainPolicyObject string

	//
	DS_enrollmentProviders string

	//
	DS_parentCA string

	//
	DS_parentCACertificateChain Uint8Array

	//
	DS_pendingCACertificates Uint8Array

	//
	DS_pendingParentCA []string

	//
	DS_previousCACertificates Uint8Array

	//
	DS_previousParentCA []string

	//
	DS_searchGuide []Uint8Array

	//
	DS_signatureAlgorithms string

	//
	DS_supportedApplicationContext []Uint8Array

	//
	DS_teletexTerminalIdentifier []Uint8Array
}

func Newads_certificationauthorityEx1(instance *cim.WmiInstance) (newInstance *ads_certificationauthority, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_certificationauthority{
		ds_top: tmp,
	}
	return
}

func Newads_certificationauthorityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_certificationauthority, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_certificationauthority{
		ds_top: tmp,
	}
	return
}

// SetDS_authorityRevocationList sets the value of DS_authorityRevocationList for the instance
func (instance *ads_certificationauthority) SetPropertyDS_authorityRevocationList(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_authorityRevocationList", (value))
}

// GetDS_authorityRevocationList gets the value of DS_authorityRevocationList for the instance
func (instance *ads_certificationauthority) GetPropertyDS_authorityRevocationList() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_authorityRevocationList")
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

// SetDS_cACertificate sets the value of DS_cACertificate for the instance
func (instance *ads_certificationauthority) SetPropertyDS_cACertificate(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_cACertificate", (value))
}

// GetDS_cACertificate gets the value of DS_cACertificate for the instance
func (instance *ads_certificationauthority) GetPropertyDS_cACertificate() (value []Uint8Array, err error) {
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
func (instance *ads_certificationauthority) SetPropertyDS_cACertificateDN(value string) (err error) {
	return instance.SetProperty("DS_cACertificateDN", (value))
}

// GetDS_cACertificateDN gets the value of DS_cACertificateDN for the instance
func (instance *ads_certificationauthority) GetPropertyDS_cACertificateDN() (value string, err error) {
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

// SetDS_cAConnect sets the value of DS_cAConnect for the instance
func (instance *ads_certificationauthority) SetPropertyDS_cAConnect(value string) (err error) {
	return instance.SetProperty("DS_cAConnect", (value))
}

// GetDS_cAConnect gets the value of DS_cAConnect for the instance
func (instance *ads_certificationauthority) GetPropertyDS_cAConnect() (value string, err error) {
	retValue, err := instance.GetProperty("DS_cAConnect")
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

// SetDS_cAUsages sets the value of DS_cAUsages for the instance
func (instance *ads_certificationauthority) SetPropertyDS_cAUsages(value []string) (err error) {
	return instance.SetProperty("DS_cAUsages", (value))
}

// GetDS_cAUsages gets the value of DS_cAUsages for the instance
func (instance *ads_certificationauthority) GetPropertyDS_cAUsages() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_cAUsages")
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

// SetDS_cAWEBURL sets the value of DS_cAWEBURL for the instance
func (instance *ads_certificationauthority) SetPropertyDS_cAWEBURL(value string) (err error) {
	return instance.SetProperty("DS_cAWEBURL", (value))
}

// GetDS_cAWEBURL gets the value of DS_cAWEBURL for the instance
func (instance *ads_certificationauthority) GetPropertyDS_cAWEBURL() (value string, err error) {
	retValue, err := instance.GetProperty("DS_cAWEBURL")
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

// SetDS_certificateRevocationList sets the value of DS_certificateRevocationList for the instance
func (instance *ads_certificationauthority) SetPropertyDS_certificateRevocationList(value Uint8Array) (err error) {
	return instance.SetProperty("DS_certificateRevocationList", (value))
}

// GetDS_certificateRevocationList gets the value of DS_certificateRevocationList for the instance
func (instance *ads_certificationauthority) GetPropertyDS_certificateRevocationList() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_certificateRevocationList")
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

// SetDS_certificateTemplates sets the value of DS_certificateTemplates for the instance
func (instance *ads_certificationauthority) SetPropertyDS_certificateTemplates(value []string) (err error) {
	return instance.SetProperty("DS_certificateTemplates", (value))
}

// GetDS_certificateTemplates gets the value of DS_certificateTemplates for the instance
func (instance *ads_certificationauthority) GetPropertyDS_certificateTemplates() (value []string, err error) {
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

// SetDS_cRLObject sets the value of DS_cRLObject for the instance
func (instance *ads_certificationauthority) SetPropertyDS_cRLObject(value string) (err error) {
	return instance.SetProperty("DS_cRLObject", (value))
}

// GetDS_cRLObject gets the value of DS_cRLObject for the instance
func (instance *ads_certificationauthority) GetPropertyDS_cRLObject() (value string, err error) {
	retValue, err := instance.GetProperty("DS_cRLObject")
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

// SetDS_crossCertificatePair sets the value of DS_crossCertificatePair for the instance
func (instance *ads_certificationauthority) SetPropertyDS_crossCertificatePair(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_crossCertificatePair", (value))
}

// GetDS_crossCertificatePair gets the value of DS_crossCertificatePair for the instance
func (instance *ads_certificationauthority) GetPropertyDS_crossCertificatePair() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_crossCertificatePair")
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

// SetDS_currentParentCA sets the value of DS_currentParentCA for the instance
func (instance *ads_certificationauthority) SetPropertyDS_currentParentCA(value []string) (err error) {
	return instance.SetProperty("DS_currentParentCA", (value))
}

// GetDS_currentParentCA gets the value of DS_currentParentCA for the instance
func (instance *ads_certificationauthority) GetPropertyDS_currentParentCA() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_currentParentCA")
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

// SetDS_deltaRevocationList sets the value of DS_deltaRevocationList for the instance
func (instance *ads_certificationauthority) SetPropertyDS_deltaRevocationList(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_deltaRevocationList", (value))
}

// GetDS_deltaRevocationList gets the value of DS_deltaRevocationList for the instance
func (instance *ads_certificationauthority) GetPropertyDS_deltaRevocationList() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_deltaRevocationList")
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

// SetDS_dNSHostName sets the value of DS_dNSHostName for the instance
func (instance *ads_certificationauthority) SetPropertyDS_dNSHostName(value string) (err error) {
	return instance.SetProperty("DS_dNSHostName", (value))
}

// GetDS_dNSHostName gets the value of DS_dNSHostName for the instance
func (instance *ads_certificationauthority) GetPropertyDS_dNSHostName() (value string, err error) {
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

// SetDS_domainID sets the value of DS_domainID for the instance
func (instance *ads_certificationauthority) SetPropertyDS_domainID(value string) (err error) {
	return instance.SetProperty("DS_domainID", (value))
}

// GetDS_domainID gets the value of DS_domainID for the instance
func (instance *ads_certificationauthority) GetPropertyDS_domainID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_domainID")
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

// SetDS_domainPolicyObject sets the value of DS_domainPolicyObject for the instance
func (instance *ads_certificationauthority) SetPropertyDS_domainPolicyObject(value string) (err error) {
	return instance.SetProperty("DS_domainPolicyObject", (value))
}

// GetDS_domainPolicyObject gets the value of DS_domainPolicyObject for the instance
func (instance *ads_certificationauthority) GetPropertyDS_domainPolicyObject() (value string, err error) {
	retValue, err := instance.GetProperty("DS_domainPolicyObject")
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
func (instance *ads_certificationauthority) SetPropertyDS_enrollmentProviders(value string) (err error) {
	return instance.SetProperty("DS_enrollmentProviders", (value))
}

// GetDS_enrollmentProviders gets the value of DS_enrollmentProviders for the instance
func (instance *ads_certificationauthority) GetPropertyDS_enrollmentProviders() (value string, err error) {
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

// SetDS_parentCA sets the value of DS_parentCA for the instance
func (instance *ads_certificationauthority) SetPropertyDS_parentCA(value string) (err error) {
	return instance.SetProperty("DS_parentCA", (value))
}

// GetDS_parentCA gets the value of DS_parentCA for the instance
func (instance *ads_certificationauthority) GetPropertyDS_parentCA() (value string, err error) {
	retValue, err := instance.GetProperty("DS_parentCA")
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

// SetDS_parentCACertificateChain sets the value of DS_parentCACertificateChain for the instance
func (instance *ads_certificationauthority) SetPropertyDS_parentCACertificateChain(value Uint8Array) (err error) {
	return instance.SetProperty("DS_parentCACertificateChain", (value))
}

// GetDS_parentCACertificateChain gets the value of DS_parentCACertificateChain for the instance
func (instance *ads_certificationauthority) GetPropertyDS_parentCACertificateChain() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_parentCACertificateChain")
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

// SetDS_pendingCACertificates sets the value of DS_pendingCACertificates for the instance
func (instance *ads_certificationauthority) SetPropertyDS_pendingCACertificates(value Uint8Array) (err error) {
	return instance.SetProperty("DS_pendingCACertificates", (value))
}

// GetDS_pendingCACertificates gets the value of DS_pendingCACertificates for the instance
func (instance *ads_certificationauthority) GetPropertyDS_pendingCACertificates() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_pendingCACertificates")
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

// SetDS_pendingParentCA sets the value of DS_pendingParentCA for the instance
func (instance *ads_certificationauthority) SetPropertyDS_pendingParentCA(value []string) (err error) {
	return instance.SetProperty("DS_pendingParentCA", (value))
}

// GetDS_pendingParentCA gets the value of DS_pendingParentCA for the instance
func (instance *ads_certificationauthority) GetPropertyDS_pendingParentCA() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_pendingParentCA")
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

// SetDS_previousCACertificates sets the value of DS_previousCACertificates for the instance
func (instance *ads_certificationauthority) SetPropertyDS_previousCACertificates(value Uint8Array) (err error) {
	return instance.SetProperty("DS_previousCACertificates", (value))
}

// GetDS_previousCACertificates gets the value of DS_previousCACertificates for the instance
func (instance *ads_certificationauthority) GetPropertyDS_previousCACertificates() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_previousCACertificates")
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

// SetDS_previousParentCA sets the value of DS_previousParentCA for the instance
func (instance *ads_certificationauthority) SetPropertyDS_previousParentCA(value []string) (err error) {
	return instance.SetProperty("DS_previousParentCA", (value))
}

// GetDS_previousParentCA gets the value of DS_previousParentCA for the instance
func (instance *ads_certificationauthority) GetPropertyDS_previousParentCA() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_previousParentCA")
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

// SetDS_searchGuide sets the value of DS_searchGuide for the instance
func (instance *ads_certificationauthority) SetPropertyDS_searchGuide(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_searchGuide", (value))
}

// GetDS_searchGuide gets the value of DS_searchGuide for the instance
func (instance *ads_certificationauthority) GetPropertyDS_searchGuide() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_searchGuide")
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

// SetDS_signatureAlgorithms sets the value of DS_signatureAlgorithms for the instance
func (instance *ads_certificationauthority) SetPropertyDS_signatureAlgorithms(value string) (err error) {
	return instance.SetProperty("DS_signatureAlgorithms", (value))
}

// GetDS_signatureAlgorithms gets the value of DS_signatureAlgorithms for the instance
func (instance *ads_certificationauthority) GetPropertyDS_signatureAlgorithms() (value string, err error) {
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

// SetDS_supportedApplicationContext sets the value of DS_supportedApplicationContext for the instance
func (instance *ads_certificationauthority) SetPropertyDS_supportedApplicationContext(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_supportedApplicationContext", (value))
}

// GetDS_supportedApplicationContext gets the value of DS_supportedApplicationContext for the instance
func (instance *ads_certificationauthority) GetPropertyDS_supportedApplicationContext() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_supportedApplicationContext")
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

// SetDS_teletexTerminalIdentifier sets the value of DS_teletexTerminalIdentifier for the instance
func (instance *ads_certificationauthority) SetPropertyDS_teletexTerminalIdentifier(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_teletexTerminalIdentifier", (value))
}

// GetDS_teletexTerminalIdentifier gets the value of DS_teletexTerminalIdentifier for the instance
func (instance *ads_certificationauthority) GetPropertyDS_teletexTerminalIdentifier() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_teletexTerminalIdentifier")
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
