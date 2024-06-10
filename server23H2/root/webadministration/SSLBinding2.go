// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SSLBinding2 struct
type SSLBinding2 struct {
	*Object

	//
	BindingOwnerID string

	//
	CertificateCheckMode int32

	//
	CertificateHash string

	//
	CTLIdentifier string

	//
	CTLStoreName string

	//
	HostName string

	//
	IPAddress string

	//
	Port uint32

	//
	RevocationFreshnessTime string

	//
	RevocationURLRetrievalTimeout string

	//
	SslAlwaysNegoClientCert bool

	//
	SslUseDsMapper bool

	//
	StoreName string
}

func NewSSLBinding2Ex1(instance *cim.WmiInstance) (newInstance *SSLBinding2, err error) {
	tmp, err := NewObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SSLBinding2{
		Object: tmp,
	}
	return
}

func NewSSLBinding2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SSLBinding2, err error) {
	tmp, err := NewObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SSLBinding2{
		Object: tmp,
	}
	return
}

// SetBindingOwnerID sets the value of BindingOwnerID for the instance
func (instance *SSLBinding2) SetPropertyBindingOwnerID(value string) (err error) {
	return instance.SetProperty("BindingOwnerID", (value))
}

// GetBindingOwnerID gets the value of BindingOwnerID for the instance
func (instance *SSLBinding2) GetPropertyBindingOwnerID() (value string, err error) {
	retValue, err := instance.GetProperty("BindingOwnerID")
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

// SetCertificateCheckMode sets the value of CertificateCheckMode for the instance
func (instance *SSLBinding2) SetPropertyCertificateCheckMode(value int32) (err error) {
	return instance.SetProperty("CertificateCheckMode", (value))
}

// GetCertificateCheckMode gets the value of CertificateCheckMode for the instance
func (instance *SSLBinding2) GetPropertyCertificateCheckMode() (value int32, err error) {
	retValue, err := instance.GetProperty("CertificateCheckMode")
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

// SetCertificateHash sets the value of CertificateHash for the instance
func (instance *SSLBinding2) SetPropertyCertificateHash(value string) (err error) {
	return instance.SetProperty("CertificateHash", (value))
}

// GetCertificateHash gets the value of CertificateHash for the instance
func (instance *SSLBinding2) GetPropertyCertificateHash() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateHash")
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

// SetCTLIdentifier sets the value of CTLIdentifier for the instance
func (instance *SSLBinding2) SetPropertyCTLIdentifier(value string) (err error) {
	return instance.SetProperty("CTLIdentifier", (value))
}

// GetCTLIdentifier gets the value of CTLIdentifier for the instance
func (instance *SSLBinding2) GetPropertyCTLIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("CTLIdentifier")
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

// SetCTLStoreName sets the value of CTLStoreName for the instance
func (instance *SSLBinding2) SetPropertyCTLStoreName(value string) (err error) {
	return instance.SetProperty("CTLStoreName", (value))
}

// GetCTLStoreName gets the value of CTLStoreName for the instance
func (instance *SSLBinding2) GetPropertyCTLStoreName() (value string, err error) {
	retValue, err := instance.GetProperty("CTLStoreName")
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

// SetHostName sets the value of HostName for the instance
func (instance *SSLBinding2) SetPropertyHostName(value string) (err error) {
	return instance.SetProperty("HostName", (value))
}

// GetHostName gets the value of HostName for the instance
func (instance *SSLBinding2) GetPropertyHostName() (value string, err error) {
	retValue, err := instance.GetProperty("HostName")
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

// SetIPAddress sets the value of IPAddress for the instance
func (instance *SSLBinding2) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", (value))
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *SSLBinding2) GetPropertyIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddress")
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

// SetPort sets the value of Port for the instance
func (instance *SSLBinding2) SetPropertyPort(value uint32) (err error) {
	return instance.SetProperty("Port", (value))
}

// GetPort gets the value of Port for the instance
func (instance *SSLBinding2) GetPropertyPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("Port")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetRevocationFreshnessTime sets the value of RevocationFreshnessTime for the instance
func (instance *SSLBinding2) SetPropertyRevocationFreshnessTime(value string) (err error) {
	return instance.SetProperty("RevocationFreshnessTime", (value))
}

// GetRevocationFreshnessTime gets the value of RevocationFreshnessTime for the instance
func (instance *SSLBinding2) GetPropertyRevocationFreshnessTime() (value string, err error) {
	retValue, err := instance.GetProperty("RevocationFreshnessTime")
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

// SetRevocationURLRetrievalTimeout sets the value of RevocationURLRetrievalTimeout for the instance
func (instance *SSLBinding2) SetPropertyRevocationURLRetrievalTimeout(value string) (err error) {
	return instance.SetProperty("RevocationURLRetrievalTimeout", (value))
}

// GetRevocationURLRetrievalTimeout gets the value of RevocationURLRetrievalTimeout for the instance
func (instance *SSLBinding2) GetPropertyRevocationURLRetrievalTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("RevocationURLRetrievalTimeout")
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

// SetSslAlwaysNegoClientCert sets the value of SslAlwaysNegoClientCert for the instance
func (instance *SSLBinding2) SetPropertySslAlwaysNegoClientCert(value bool) (err error) {
	return instance.SetProperty("SslAlwaysNegoClientCert", (value))
}

// GetSslAlwaysNegoClientCert gets the value of SslAlwaysNegoClientCert for the instance
func (instance *SSLBinding2) GetPropertySslAlwaysNegoClientCert() (value bool, err error) {
	retValue, err := instance.GetProperty("SslAlwaysNegoClientCert")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetSslUseDsMapper sets the value of SslUseDsMapper for the instance
func (instance *SSLBinding2) SetPropertySslUseDsMapper(value bool) (err error) {
	return instance.SetProperty("SslUseDsMapper", (value))
}

// GetSslUseDsMapper gets the value of SslUseDsMapper for the instance
func (instance *SSLBinding2) GetPropertySslUseDsMapper() (value bool, err error) {
	retValue, err := instance.GetProperty("SslUseDsMapper")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetStoreName sets the value of StoreName for the instance
func (instance *SSLBinding2) SetPropertyStoreName(value string) (err error) {
	return instance.SetProperty("StoreName", (value))
}

// GetStoreName gets the value of StoreName for the instance
func (instance *SSLBinding2) GetPropertyStoreName() (value string, err error) {
	retValue, err := instance.GetProperty("StoreName")
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

//

// <param name="CertificateHash" type="string "></param>
// <param name="HostName" type="string "></param>
// <param name="IPAddress" type="string "></param>
// <param name="Port" type="uint32 "></param>
// <param name="StoreName" type="string "></param>
func (instance *SSLBinding2) Create( /* IN */ IPAddress string,
	/* IN */ Port uint32,
	/* IN */ CertificateHash string,
	/* IN */ StoreName string,
	/* IN */ HostName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("Create", IPAddress, Port, CertificateHash, StoreName, HostName)
	if err != nil {
		return
	}
	return

}
