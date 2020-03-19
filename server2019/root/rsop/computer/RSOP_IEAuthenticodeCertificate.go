// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_IEAuthenticodeCertificate struct
type RSOP_IEAuthenticodeCertificate struct {
	*cim.WmiInstance

	//
	certIndex int32

	//
	expirationDate string

	//
	friendlyName string

	//
	intendedPurposes string

	//
	issuerName string

	//
	rsopID string

	//
	rsopPrecedence int32

	//
	subjectName string

	//
	tabIndex int32
}

func NewRSOP_IEAuthenticodeCertificateEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEAuthenticodeCertificate, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEAuthenticodeCertificate{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEAuthenticodeCertificateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEAuthenticodeCertificate, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEAuthenticodeCertificate{
		WmiInstance: tmp,
	}
	return
}

// SetcertIndex sets the value of certIndex for the instance
func (instance *RSOP_IEAuthenticodeCertificate) SetPropertycertIndex(value int32) (err error) {
	return instance.SetProperty("certIndex", value)
}

// GetcertIndex gets the value of certIndex for the instance
func (instance *RSOP_IEAuthenticodeCertificate) GetPropertycertIndex() (value int32, err error) {
	retValue, err := instance.GetProperty("certIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetexpirationDate sets the value of expirationDate for the instance
func (instance *RSOP_IEAuthenticodeCertificate) SetPropertyexpirationDate(value string) (err error) {
	return instance.SetProperty("expirationDate", value)
}

// GetexpirationDate gets the value of expirationDate for the instance
func (instance *RSOP_IEAuthenticodeCertificate) GetPropertyexpirationDate() (value string, err error) {
	retValue, err := instance.GetProperty("expirationDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetfriendlyName sets the value of friendlyName for the instance
func (instance *RSOP_IEAuthenticodeCertificate) SetPropertyfriendlyName(value string) (err error) {
	return instance.SetProperty("friendlyName", value)
}

// GetfriendlyName gets the value of friendlyName for the instance
func (instance *RSOP_IEAuthenticodeCertificate) GetPropertyfriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("friendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetintendedPurposes sets the value of intendedPurposes for the instance
func (instance *RSOP_IEAuthenticodeCertificate) SetPropertyintendedPurposes(value string) (err error) {
	return instance.SetProperty("intendedPurposes", value)
}

// GetintendedPurposes gets the value of intendedPurposes for the instance
func (instance *RSOP_IEAuthenticodeCertificate) GetPropertyintendedPurposes() (value string, err error) {
	retValue, err := instance.GetProperty("intendedPurposes")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetissuerName sets the value of issuerName for the instance
func (instance *RSOP_IEAuthenticodeCertificate) SetPropertyissuerName(value string) (err error) {
	return instance.SetProperty("issuerName", value)
}

// GetissuerName gets the value of issuerName for the instance
func (instance *RSOP_IEAuthenticodeCertificate) GetPropertyissuerName() (value string, err error) {
	retValue, err := instance.GetProperty("issuerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IEAuthenticodeCertificate) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", value)
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IEAuthenticodeCertificate) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IEAuthenticodeCertificate) SetPropertyrsopPrecedence(value int32) (err error) {
	return instance.SetProperty("rsopPrecedence", value)
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IEAuthenticodeCertificate) GetPropertyrsopPrecedence() (value int32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetsubjectName sets the value of subjectName for the instance
func (instance *RSOP_IEAuthenticodeCertificate) SetPropertysubjectName(value string) (err error) {
	return instance.SetProperty("subjectName", value)
}

// GetsubjectName gets the value of subjectName for the instance
func (instance *RSOP_IEAuthenticodeCertificate) GetPropertysubjectName() (value string, err error) {
	retValue, err := instance.GetProperty("subjectName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SettabIndex sets the value of tabIndex for the instance
func (instance *RSOP_IEAuthenticodeCertificate) SetPropertytabIndex(value int32) (err error) {
	return instance.SetProperty("tabIndex", value)
}

// GettabIndex gets the value of tabIndex for the instance
func (instance *RSOP_IEAuthenticodeCertificate) GetPropertytabIndex() (value int32, err error) {
	retValue, err := instance.GetProperty("tabIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
