// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Hgs
//////////////////////////////////////////////
package hgs

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_HgsGuardian struct
type MSFT_HgsGuardian struct {
	*cim.WmiInstance

	//
	EncryptionCertificate []uint8

	//
	EncryptionCertificateSignature string

	//
	EncryptionCertificateSignatureAlgorithm string

	//
	GuardianMetadataVersion uint32

	//
	HasPrivateSigningKey bool

	//
	Name string

	//
	SigningCertificate []uint8
}

func NewMSFT_HgsGuardianEx1(instance *cim.WmiInstance) (newInstance *MSFT_HgsGuardian, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_HgsGuardian{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_HgsGuardianEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_HgsGuardian, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_HgsGuardian{
		WmiInstance: tmp,
	}
	return
}

// SetEncryptionCertificate sets the value of EncryptionCertificate for the instance
func (instance *MSFT_HgsGuardian) SetPropertyEncryptionCertificate(value []uint8) (err error) {
	return instance.SetProperty("EncryptionCertificate", value)
}

// GetEncryptionCertificate gets the value of EncryptionCertificate for the instance
func (instance *MSFT_HgsGuardian) GetPropertyEncryptionCertificate() (value []uint8, err error) {
	retValue, err := instance.GetProperty("EncryptionCertificate")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryptionCertificateSignature sets the value of EncryptionCertificateSignature for the instance
func (instance *MSFT_HgsGuardian) SetPropertyEncryptionCertificateSignature(value string) (err error) {
	return instance.SetProperty("EncryptionCertificateSignature", value)
}

// GetEncryptionCertificateSignature gets the value of EncryptionCertificateSignature for the instance
func (instance *MSFT_HgsGuardian) GetPropertyEncryptionCertificateSignature() (value string, err error) {
	retValue, err := instance.GetProperty("EncryptionCertificateSignature")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryptionCertificateSignatureAlgorithm sets the value of EncryptionCertificateSignatureAlgorithm for the instance
func (instance *MSFT_HgsGuardian) SetPropertyEncryptionCertificateSignatureAlgorithm(value string) (err error) {
	return instance.SetProperty("EncryptionCertificateSignatureAlgorithm", value)
}

// GetEncryptionCertificateSignatureAlgorithm gets the value of EncryptionCertificateSignatureAlgorithm for the instance
func (instance *MSFT_HgsGuardian) GetPropertyEncryptionCertificateSignatureAlgorithm() (value string, err error) {
	retValue, err := instance.GetProperty("EncryptionCertificateSignatureAlgorithm")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuardianMetadataVersion sets the value of GuardianMetadataVersion for the instance
func (instance *MSFT_HgsGuardian) SetPropertyGuardianMetadataVersion(value uint32) (err error) {
	return instance.SetProperty("GuardianMetadataVersion", value)
}

// GetGuardianMetadataVersion gets the value of GuardianMetadataVersion for the instance
func (instance *MSFT_HgsGuardian) GetPropertyGuardianMetadataVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("GuardianMetadataVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHasPrivateSigningKey sets the value of HasPrivateSigningKey for the instance
func (instance *MSFT_HgsGuardian) SetPropertyHasPrivateSigningKey(value bool) (err error) {
	return instance.SetProperty("HasPrivateSigningKey", value)
}

// GetHasPrivateSigningKey gets the value of HasPrivateSigningKey for the instance
func (instance *MSFT_HgsGuardian) GetPropertyHasPrivateSigningKey() (value bool, err error) {
	retValue, err := instance.GetProperty("HasPrivateSigningKey")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_HgsGuardian) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_HgsGuardian) GetPropertyName() (value string, err error) {
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

// SetSigningCertificate sets the value of SigningCertificate for the instance
func (instance *MSFT_HgsGuardian) SetPropertySigningCertificate(value []uint8) (err error) {
	return instance.SetProperty("SigningCertificate", value)
}

// GetSigningCertificate gets the value of SigningCertificate for the instance
func (instance *MSFT_HgsGuardian) GetPropertySigningCertificate() (value []uint8, err error) {
	retValue, err := instance.GetProperty("SigningCertificate")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="AllowExpired" type="bool "></param>
// <param name="AllowUntrustedRoot" type="bool "></param>
// <param name="EncryptionCertificate" type="string "></param>
// <param name="EncryptionCertificatePassword" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="SigningCertificate" type="string "></param>
// <param name="SigningCertificatePassword" type="string "></param>

// <param name="cmdletOutput" type="MSFT_HgsGuardian "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsGuardian) NewByAcceptCertificates( /* IN */ Name string,
	/* IN */ EncryptionCertificate string,
	/* IN */ SigningCertificate string,
	/* IN */ SigningCertificatePassword string,
	/* IN */ EncryptionCertificatePassword string,
	/* IN */ AllowExpired bool,
	/* IN */ AllowUntrustedRoot bool,
	/* OUT */ cmdletOutput MSFT_HgsGuardian) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewByAcceptCertificates", Name, EncryptionCertificate, SigningCertificate, SigningCertificatePassword, EncryptionCertificatePassword, AllowExpired, AllowUntrustedRoot)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllowExpired" type="bool "></param>
// <param name="AllowUntrustedRoot" type="bool "></param>
// <param name="EncryptionCertificateThumbprint" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="SigningCertificateThumbprint" type="string "></param>

// <param name="cmdletOutput" type="MSFT_HgsGuardian "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsGuardian) NewByCertificateThumbprints( /* IN */ Name string,
	/* IN */ SigningCertificateThumbprint string,
	/* IN */ EncryptionCertificateThumbprint string,
	/* IN */ AllowExpired bool,
	/* IN */ AllowUntrustedRoot bool,
	/* OUT */ cmdletOutput MSFT_HgsGuardian) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewByCertificateThumbprints", Name, SigningCertificateThumbprint, EncryptionCertificateThumbprint, AllowExpired, AllowUntrustedRoot)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="GenerateCertificates" type="bool "></param>
// <param name="Name" type="string "></param>

// <param name="cmdletOutput" type="MSFT_HgsGuardian "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsGuardian) NewByGenerateCertificates( /* IN */ Name string,
	/* IN */ GenerateCertificates bool,
	/* OUT */ cmdletOutput MSFT_HgsGuardian) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewByGenerateCertificates", Name, GenerateCertificates)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllowExpired" type="bool "></param>
// <param name="AllowUntrustedRoot" type="bool "></param>
// <param name="Name" type="string "></param>
// <param name="Path" type="string "></param>

// <param name="cmdletOutput" type="MSFT_HgsGuardian "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsGuardian) Import( /* IN */ Path string,
	/* IN */ Name string,
	/* IN */ AllowExpired bool,
	/* IN */ AllowUntrustedRoot bool,
	/* OUT */ cmdletOutput MSFT_HgsGuardian) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Import", Path, Name, AllowExpired, AllowUntrustedRoot)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InputObject" type="MSFT_HgsGuardian "></param>
// <param name="Path" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsGuardian) Export( /* IN */ Path string,
	/* IN */ InputObject MSFT_HgsGuardian) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Export", Path, InputObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsGuardian) Remove( /* IN */ Name string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove", Name)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
