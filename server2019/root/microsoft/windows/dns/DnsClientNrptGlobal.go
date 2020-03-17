// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Dns
//////////////////////////////////////////////
package dns

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// DnsClientNrptGlobal struct
type DnsClientNrptGlobal struct {
	cim.WmiInstance

	//
	EnableDAForAllNetworks string

	//
	QueryPolicy string

	//
	SecureNameQueryFallback string
}

// SetEnableDAForAllNetworks sets the value of EnableDAForAllNetworks for the instance
func (instance *DnsClientNrptGlobal) SetPropertyEnableDAForAllNetworks(value string) (err error) {
	return instance.SetProperty("EnableDAForAllNetworks", value)
}

// GetEnableDAForAllNetworks gets the value of EnableDAForAllNetworks for the instance
func (instance *DnsClientNrptGlobal) GetPropertyEnableDAForAllNetworks() (value string, err error) {
	retValue, err := instance.GetProperty("EnableDAForAllNetworks")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQueryPolicy sets the value of QueryPolicy for the instance
func (instance *DnsClientNrptGlobal) SetPropertyQueryPolicy(value string) (err error) {
	return instance.SetProperty("QueryPolicy", value)
}

// GetQueryPolicy gets the value of QueryPolicy for the instance
func (instance *DnsClientNrptGlobal) GetPropertyQueryPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("QueryPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecureNameQueryFallback sets the value of SecureNameQueryFallback for the instance
func (instance *DnsClientNrptGlobal) SetPropertySecureNameQueryFallback(value string) (err error) {
	return instance.SetProperty("SecureNameQueryFallback", value)
}

// GetSecureNameQueryFallback gets the value of SecureNameQueryFallback for the instance
func (instance *DnsClientNrptGlobal) GetPropertySecureNameQueryFallback() (value string, err error) {
	retValue, err := instance.GetProperty("SecureNameQueryFallback")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
