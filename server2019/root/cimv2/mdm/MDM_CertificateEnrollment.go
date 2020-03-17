// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_CertificateEnrollment struct
type MDM_CertificateEnrollment struct {
	cim.WmiInstance

	//
	ConfigurationParameters string

	//
	EnhancedKeyUsages string

	//
	Error uint32

	//
	ExpirationThreshold uint32

	//
	Issuers string

	//
	RequestID string

	//
	SerialNumber string

	//
	Status uint32

	//
	StoreLocation uint8

	//
	SubjectAlternativeNames string

	//
	SubjectName string

	//
	Thumbprint string

	//
	ValidFrom string

	//
	ValidTo string
}

// SetConfigurationParameters sets the value of ConfigurationParameters for the instance
func (instance *MDM_CertificateEnrollment) SetPropertyConfigurationParameters(value string) (err error) {
	return instance.SetProperty("ConfigurationParameters", value)
}

// GetConfigurationParameters gets the value of ConfigurationParameters for the instance
func (instance *MDM_CertificateEnrollment) GetPropertyConfigurationParameters() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigurationParameters")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnhancedKeyUsages sets the value of EnhancedKeyUsages for the instance
func (instance *MDM_CertificateEnrollment) SetPropertyEnhancedKeyUsages(value string) (err error) {
	return instance.SetProperty("EnhancedKeyUsages", value)
}

// GetEnhancedKeyUsages gets the value of EnhancedKeyUsages for the instance
func (instance *MDM_CertificateEnrollment) GetPropertyEnhancedKeyUsages() (value string, err error) {
	retValue, err := instance.GetProperty("EnhancedKeyUsages")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetError sets the value of Error for the instance
func (instance *MDM_CertificateEnrollment) SetPropertyError(value uint32) (err error) {
	return instance.SetProperty("Error", value)
}

// GetError gets the value of Error for the instance
func (instance *MDM_CertificateEnrollment) GetPropertyError() (value uint32, err error) {
	retValue, err := instance.GetProperty("Error")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExpirationThreshold sets the value of ExpirationThreshold for the instance
func (instance *MDM_CertificateEnrollment) SetPropertyExpirationThreshold(value uint32) (err error) {
	return instance.SetProperty("ExpirationThreshold", value)
}

// GetExpirationThreshold gets the value of ExpirationThreshold for the instance
func (instance *MDM_CertificateEnrollment) GetPropertyExpirationThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExpirationThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIssuers sets the value of Issuers for the instance
func (instance *MDM_CertificateEnrollment) SetPropertyIssuers(value string) (err error) {
	return instance.SetProperty("Issuers", value)
}

// GetIssuers gets the value of Issuers for the instance
func (instance *MDM_CertificateEnrollment) GetPropertyIssuers() (value string, err error) {
	retValue, err := instance.GetProperty("Issuers")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestID sets the value of RequestID for the instance
func (instance *MDM_CertificateEnrollment) SetPropertyRequestID(value string) (err error) {
	return instance.SetProperty("RequestID", value)
}

// GetRequestID gets the value of RequestID for the instance
func (instance *MDM_CertificateEnrollment) GetPropertyRequestID() (value string, err error) {
	retValue, err := instance.GetProperty("RequestID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *MDM_CertificateEnrollment) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", value)
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *MDM_CertificateEnrollment) GetPropertySerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MDM_CertificateEnrollment) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_CertificateEnrollment) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStoreLocation sets the value of StoreLocation for the instance
func (instance *MDM_CertificateEnrollment) SetPropertyStoreLocation(value uint8) (err error) {
	return instance.SetProperty("StoreLocation", value)
}

// GetStoreLocation gets the value of StoreLocation for the instance
func (instance *MDM_CertificateEnrollment) GetPropertyStoreLocation() (value uint8, err error) {
	retValue, err := instance.GetProperty("StoreLocation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubjectAlternativeNames sets the value of SubjectAlternativeNames for the instance
func (instance *MDM_CertificateEnrollment) SetPropertySubjectAlternativeNames(value string) (err error) {
	return instance.SetProperty("SubjectAlternativeNames", value)
}

// GetSubjectAlternativeNames gets the value of SubjectAlternativeNames for the instance
func (instance *MDM_CertificateEnrollment) GetPropertySubjectAlternativeNames() (value string, err error) {
	retValue, err := instance.GetProperty("SubjectAlternativeNames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubjectName sets the value of SubjectName for the instance
func (instance *MDM_CertificateEnrollment) SetPropertySubjectName(value string) (err error) {
	return instance.SetProperty("SubjectName", value)
}

// GetSubjectName gets the value of SubjectName for the instance
func (instance *MDM_CertificateEnrollment) GetPropertySubjectName() (value string, err error) {
	retValue, err := instance.GetProperty("SubjectName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThumbprint sets the value of Thumbprint for the instance
func (instance *MDM_CertificateEnrollment) SetPropertyThumbprint(value string) (err error) {
	return instance.SetProperty("Thumbprint", value)
}

// GetThumbprint gets the value of Thumbprint for the instance
func (instance *MDM_CertificateEnrollment) GetPropertyThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("Thumbprint")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValidFrom sets the value of ValidFrom for the instance
func (instance *MDM_CertificateEnrollment) SetPropertyValidFrom(value string) (err error) {
	return instance.SetProperty("ValidFrom", value)
}

// GetValidFrom gets the value of ValidFrom for the instance
func (instance *MDM_CertificateEnrollment) GetPropertyValidFrom() (value string, err error) {
	retValue, err := instance.GetProperty("ValidFrom")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValidTo sets the value of ValidTo for the instance
func (instance *MDM_CertificateEnrollment) SetPropertyValidTo(value string) (err error) {
	return instance.SetProperty("ValidTo", value)
}

// GetValidTo gets the value of ValidTo for the instance
func (instance *MDM_CertificateEnrollment) GetPropertyValidTo() (value string, err error) {
	retValue, err := instance.GetProperty("ValidTo")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
