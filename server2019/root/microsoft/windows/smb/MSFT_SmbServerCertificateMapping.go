// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SmbServerCertificateMapping struct
type MSFT_SmbServerCertificateMapping struct {
	*cim.WmiInstance

	//
	DisplayName string

	//
	Name string

	//
	Subject string

	//
	Thumbprint string

	//
	Type SmbServerCertificateMapping_Type
}

func NewMSFT_SmbServerCertificateMappingEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbServerCertificateMapping, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbServerCertificateMapping{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbServerCertificateMappingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbServerCertificateMapping, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbServerCertificateMapping{
		WmiInstance: tmp,
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubject sets the value of Subject for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertySubject(value string) (err error) {
	return instance.SetProperty("Subject", value)
}

// GetSubject gets the value of Subject for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertySubject() (value string, err error) {
	retValue, err := instance.GetProperty("Subject")
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
func (instance *MSFT_SmbServerCertificateMapping) SetPropertyThumbprint(value string) (err error) {
	return instance.SetProperty("Thumbprint", value)
}

// GetThumbprint gets the value of Thumbprint for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertyThumbprint() (value string, err error) {
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

// SetType sets the value of Type for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertyType(value SmbServerCertificateMapping_Type) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertyType() (value SmbServerCertificateMapping_Type, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbServerCertificateMapping_Type)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="DisplayName" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="Subject" type="string "></param>
// <param name="Thumbprint" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="CreatedMapping" type="MSFT_SmbServerCertificateMapping "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbServerCertificateMapping) AddServerCertificateMapping( /* IN */ Name string,
	/* IN */ Subject string,
	/* IN */ Thumbprint string,
	/* IN */ DisplayName string,
	/* IN */ Type uint32,
	/* OUT */ CreatedMapping MSFT_SmbServerCertificateMapping) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddServerCertificateMapping", Name, Subject, Thumbprint, DisplayName, Type)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
