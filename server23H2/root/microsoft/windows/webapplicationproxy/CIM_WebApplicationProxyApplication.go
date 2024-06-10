// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.WebApplicationProxy
//
// ////////////////////////////////////////////
package webapplicationproxy

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_WebApplicationProxyApplication struct
type CIM_WebApplicationProxyApplication struct {
	*cim.WmiInstance
}

func NewCIM_WebApplicationProxyApplicationEx1(instance *cim.WmiInstance) (newInstance *CIM_WebApplicationProxyApplication, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_WebApplicationProxyApplication{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_WebApplicationProxyApplicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_WebApplicationProxyApplication, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_WebApplicationProxyApplication{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ID" type="string "></param>

// <param name="cmdletOutput" type="PublishedWebApp []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_WebApplicationProxyApplication) GetByID( /* IN */ ID string,
	/* OUT */ cmdletOutput []PublishedWebApp) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByID", ID)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Name" type="string "></param>

// <param name="cmdletOutput" type="PublishedWebApp []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_WebApplicationProxyApplication) GetByName( /* IN */ Name string,
	/* OUT */ cmdletOutput []PublishedWebApp) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByName", Name)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ADFSRelyingPartyName" type="string "></param>
// <param name="ADFSUserCertificateStore" type="string "></param>
// <param name="BackendServerAuthenticationSPN" type="string "></param>
// <param name="BackendServerCertificateValidation" type="string "></param>
// <param name="BackendServerUrl" type="string "></param>
// <param name="ClientCertificateAuthenticationBindingMode" type="string "></param>
// <param name="ClientCertificatePreauthenticationThumbprint" type="string "></param>
// <param name="DisableHttpOnlyCookieProtection" type="bool "></param>
// <param name="DisableTranslateUrlInRequestHeaders" type="bool "></param>
// <param name="DisableTranslateUrlInResponseHeaders" type="bool "></param>
// <param name="EnableHTTPRedirect" type="bool "></param>
// <param name="EnableSignOut" type="bool "></param>
// <param name="ExternalCertificateThumbprint" type="string "></param>
// <param name="ExternalPreauthentication" type="string "></param>
// <param name="ExternalUrl" type="string "></param>
// <param name="InactiveTransactionsTimeoutSec" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="PersistentAccessCookieExpirationTimeSec" type="uint32 "></param>
// <param name="UseOAuthAuthentication" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_WebApplicationProxyApplication) Add( /* IN */ Name string,
	/* IN */ ExternalPreauthentication string,
	/* IN */ ClientCertificateAuthenticationBindingMode string,
	/* IN */ BackendServerCertificateValidation string,
	/* IN */ ExternalUrl string,
	/* IN */ ExternalCertificateThumbprint string,
	/* IN */ EnableSignOut bool,
	/* IN */ InactiveTransactionsTimeoutSec uint32,
	/* IN */ ClientCertificatePreauthenticationThumbprint string,
	/* IN */ EnableHTTPRedirect bool,
	/* IN */ ADFSUserCertificateStore string,
	/* IN */ DisableHttpOnlyCookieProtection bool,
	/* IN */ PersistentAccessCookieExpirationTimeSec uint32,
	/* IN */ BackendServerUrl string,
	/* IN */ DisableTranslateUrlInRequestHeaders bool,
	/* IN */ DisableTranslateUrlInResponseHeaders bool,
	/* IN */ BackendServerAuthenticationSPN string,
	/* IN */ ADFSRelyingPartyName string,
	/* IN */ UseOAuthAuthentication bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Add", Name, ExternalPreauthentication, ClientCertificateAuthenticationBindingMode, BackendServerCertificateValidation, ExternalUrl, ExternalCertificateThumbprint, EnableSignOut, InactiveTransactionsTimeoutSec, ClientCertificatePreauthenticationThumbprint, EnableHTTPRedirect, ADFSUserCertificateStore, DisableHttpOnlyCookieProtection, PersistentAccessCookieExpirationTimeSec, BackendServerUrl, DisableTranslateUrlInRequestHeaders, DisableTranslateUrlInResponseHeaders, BackendServerAuthenticationSPN, ADFSRelyingPartyName, UseOAuthAuthentication)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ADFSUserCertificateStore" type="string "></param>
// <param name="BackendServerAuthenticationMode" type="string "></param>
// <param name="BackendServerAuthenticationSPN" type="string "></param>
// <param name="BackendServerCertificateValidation" type="string "></param>
// <param name="BackendServerUrl" type="string "></param>
// <param name="ClientCertificateAuthenticationBindingMode" type="string "></param>
// <param name="ClientCertificatePreauthenticationThumbprint" type="string "></param>
// <param name="DisableHttpOnlyCookieProtection" type="bool "></param>
// <param name="DisableTranslateUrlInRequestHeaders" type="bool "></param>
// <param name="DisableTranslateUrlInResponseHeaders" type="bool "></param>
// <param name="EnableHTTPRedirect" type="bool "></param>
// <param name="EnableSignOut" type="bool "></param>
// <param name="ExternalCertificateThumbprint" type="string "></param>
// <param name="ExternalUrl" type="string "></param>
// <param name="ID" type="string "></param>
// <param name="InactiveTransactionsTimeoutSec" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="PersistentAccessCookieExpirationTimeSec" type="uint32 "></param>
// <param name="UseOAuthAuthentication" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_WebApplicationProxyApplication) SetByID( /* IN */ ClientCertificateAuthenticationBindingMode string,
	/* IN */ BackendServerCertificateValidation string,
	/* IN */ ExternalUrl string,
	/* IN */ ExternalCertificateThumbprint string,
	/* IN */ BackendServerUrl string,
	/* IN */ DisableTranslateUrlInRequestHeaders bool,
	/* IN */ EnableHTTPRedirect bool,
	/* IN */ ADFSUserCertificateStore string,
	/* IN */ DisableHttpOnlyCookieProtection bool,
	/* IN */ PersistentAccessCookieExpirationTimeSec uint32,
	/* IN */ EnableSignOut bool,
	/* IN */ BackendServerAuthenticationMode string,
	/* IN */ DisableTranslateUrlInResponseHeaders bool,
	/* IN */ BackendServerAuthenticationSPN string,
	/* IN */ Name string,
	/* IN */ UseOAuthAuthentication bool,
	/* IN */ InactiveTransactionsTimeoutSec uint32,
	/* IN */ ClientCertificatePreauthenticationThumbprint string,
	/* IN */ ID string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByID", ClientCertificateAuthenticationBindingMode, BackendServerCertificateValidation, ExternalUrl, ExternalCertificateThumbprint, BackendServerUrl, DisableTranslateUrlInRequestHeaders, EnableHTTPRedirect, ADFSUserCertificateStore, DisableHttpOnlyCookieProtection, PersistentAccessCookieExpirationTimeSec, EnableSignOut, BackendServerAuthenticationMode, DisableTranslateUrlInResponseHeaders, BackendServerAuthenticationSPN, Name, UseOAuthAuthentication, InactiveTransactionsTimeoutSec, ClientCertificatePreauthenticationThumbprint, ID)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ID" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_WebApplicationProxyApplication) RemoveByID( /* IN */ ID []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveByID", ID)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_WebApplicationProxyApplication) RemoveByName( /* IN */ Name string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveByName", Name)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
