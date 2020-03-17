// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_NTDomain struct
type Win32_NTDomain struct {
	CIM_System

	//
	ClientSiteName string

	//
	DcSiteName string

	//
	DnsForestName string

	//
	DomainControllerAddress string

	//
	DomainControllerAddressType int32

	//
	DomainControllerName string

	//
	DomainGuid string

	//
	DomainName string

	//
	DSDirectoryServiceFlag bool

	//
	DSDnsControllerFlag bool

	//
	DSDnsDomainFlag bool

	//
	DSDnsForestFlag bool

	//
	DSGlobalCatalogFlag bool

	//
	DSKerberosDistributionCenterFlag bool

	//
	DSPrimaryDomainControllerFlag bool

	//
	DSTimeServiceFlag bool

	//
	DSWritableFlag bool
}

// SetClientSiteName sets the value of ClientSiteName for the instance
func (instance *Win32_NTDomain) SetPropertyClientSiteName(value string) (err error) {
	return instance.SetProperty("ClientSiteName", value)
}

// GetClientSiteName gets the value of ClientSiteName for the instance
func (instance *Win32_NTDomain) GetPropertyClientSiteName() (value string, err error) {
	retValue, err := instance.GetProperty("ClientSiteName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDcSiteName sets the value of DcSiteName for the instance
func (instance *Win32_NTDomain) SetPropertyDcSiteName(value string) (err error) {
	return instance.SetProperty("DcSiteName", value)
}

// GetDcSiteName gets the value of DcSiteName for the instance
func (instance *Win32_NTDomain) GetPropertyDcSiteName() (value string, err error) {
	retValue, err := instance.GetProperty("DcSiteName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsForestName sets the value of DnsForestName for the instance
func (instance *Win32_NTDomain) SetPropertyDnsForestName(value string) (err error) {
	return instance.SetProperty("DnsForestName", value)
}

// GetDnsForestName gets the value of DnsForestName for the instance
func (instance *Win32_NTDomain) GetPropertyDnsForestName() (value string, err error) {
	retValue, err := instance.GetProperty("DnsForestName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomainControllerAddress sets the value of DomainControllerAddress for the instance
func (instance *Win32_NTDomain) SetPropertyDomainControllerAddress(value string) (err error) {
	return instance.SetProperty("DomainControllerAddress", value)
}

// GetDomainControllerAddress gets the value of DomainControllerAddress for the instance
func (instance *Win32_NTDomain) GetPropertyDomainControllerAddress() (value string, err error) {
	retValue, err := instance.GetProperty("DomainControllerAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomainControllerAddressType sets the value of DomainControllerAddressType for the instance
func (instance *Win32_NTDomain) SetPropertyDomainControllerAddressType(value int32) (err error) {
	return instance.SetProperty("DomainControllerAddressType", value)
}

// GetDomainControllerAddressType gets the value of DomainControllerAddressType for the instance
func (instance *Win32_NTDomain) GetPropertyDomainControllerAddressType() (value int32, err error) {
	retValue, err := instance.GetProperty("DomainControllerAddressType")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomainControllerName sets the value of DomainControllerName for the instance
func (instance *Win32_NTDomain) SetPropertyDomainControllerName(value string) (err error) {
	return instance.SetProperty("DomainControllerName", value)
}

// GetDomainControllerName gets the value of DomainControllerName for the instance
func (instance *Win32_NTDomain) GetPropertyDomainControllerName() (value string, err error) {
	retValue, err := instance.GetProperty("DomainControllerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomainGuid sets the value of DomainGuid for the instance
func (instance *Win32_NTDomain) SetPropertyDomainGuid(value string) (err error) {
	return instance.SetProperty("DomainGuid", value)
}

// GetDomainGuid gets the value of DomainGuid for the instance
func (instance *Win32_NTDomain) GetPropertyDomainGuid() (value string, err error) {
	retValue, err := instance.GetProperty("DomainGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomainName sets the value of DomainName for the instance
func (instance *Win32_NTDomain) SetPropertyDomainName(value string) (err error) {
	return instance.SetProperty("DomainName", value)
}

// GetDomainName gets the value of DomainName for the instance
func (instance *Win32_NTDomain) GetPropertyDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("DomainName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDSDirectoryServiceFlag sets the value of DSDirectoryServiceFlag for the instance
func (instance *Win32_NTDomain) SetPropertyDSDirectoryServiceFlag(value bool) (err error) {
	return instance.SetProperty("DSDirectoryServiceFlag", value)
}

// GetDSDirectoryServiceFlag gets the value of DSDirectoryServiceFlag for the instance
func (instance *Win32_NTDomain) GetPropertyDSDirectoryServiceFlag() (value bool, err error) {
	retValue, err := instance.GetProperty("DSDirectoryServiceFlag")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDSDnsControllerFlag sets the value of DSDnsControllerFlag for the instance
func (instance *Win32_NTDomain) SetPropertyDSDnsControllerFlag(value bool) (err error) {
	return instance.SetProperty("DSDnsControllerFlag", value)
}

// GetDSDnsControllerFlag gets the value of DSDnsControllerFlag for the instance
func (instance *Win32_NTDomain) GetPropertyDSDnsControllerFlag() (value bool, err error) {
	retValue, err := instance.GetProperty("DSDnsControllerFlag")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDSDnsDomainFlag sets the value of DSDnsDomainFlag for the instance
func (instance *Win32_NTDomain) SetPropertyDSDnsDomainFlag(value bool) (err error) {
	return instance.SetProperty("DSDnsDomainFlag", value)
}

// GetDSDnsDomainFlag gets the value of DSDnsDomainFlag for the instance
func (instance *Win32_NTDomain) GetPropertyDSDnsDomainFlag() (value bool, err error) {
	retValue, err := instance.GetProperty("DSDnsDomainFlag")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDSDnsForestFlag sets the value of DSDnsForestFlag for the instance
func (instance *Win32_NTDomain) SetPropertyDSDnsForestFlag(value bool) (err error) {
	return instance.SetProperty("DSDnsForestFlag", value)
}

// GetDSDnsForestFlag gets the value of DSDnsForestFlag for the instance
func (instance *Win32_NTDomain) GetPropertyDSDnsForestFlag() (value bool, err error) {
	retValue, err := instance.GetProperty("DSDnsForestFlag")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDSGlobalCatalogFlag sets the value of DSGlobalCatalogFlag for the instance
func (instance *Win32_NTDomain) SetPropertyDSGlobalCatalogFlag(value bool) (err error) {
	return instance.SetProperty("DSGlobalCatalogFlag", value)
}

// GetDSGlobalCatalogFlag gets the value of DSGlobalCatalogFlag for the instance
func (instance *Win32_NTDomain) GetPropertyDSGlobalCatalogFlag() (value bool, err error) {
	retValue, err := instance.GetProperty("DSGlobalCatalogFlag")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDSKerberosDistributionCenterFlag sets the value of DSKerberosDistributionCenterFlag for the instance
func (instance *Win32_NTDomain) SetPropertyDSKerberosDistributionCenterFlag(value bool) (err error) {
	return instance.SetProperty("DSKerberosDistributionCenterFlag", value)
}

// GetDSKerberosDistributionCenterFlag gets the value of DSKerberosDistributionCenterFlag for the instance
func (instance *Win32_NTDomain) GetPropertyDSKerberosDistributionCenterFlag() (value bool, err error) {
	retValue, err := instance.GetProperty("DSKerberosDistributionCenterFlag")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDSPrimaryDomainControllerFlag sets the value of DSPrimaryDomainControllerFlag for the instance
func (instance *Win32_NTDomain) SetPropertyDSPrimaryDomainControllerFlag(value bool) (err error) {
	return instance.SetProperty("DSPrimaryDomainControllerFlag", value)
}

// GetDSPrimaryDomainControllerFlag gets the value of DSPrimaryDomainControllerFlag for the instance
func (instance *Win32_NTDomain) GetPropertyDSPrimaryDomainControllerFlag() (value bool, err error) {
	retValue, err := instance.GetProperty("DSPrimaryDomainControllerFlag")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDSTimeServiceFlag sets the value of DSTimeServiceFlag for the instance
func (instance *Win32_NTDomain) SetPropertyDSTimeServiceFlag(value bool) (err error) {
	return instance.SetProperty("DSTimeServiceFlag", value)
}

// GetDSTimeServiceFlag gets the value of DSTimeServiceFlag for the instance
func (instance *Win32_NTDomain) GetPropertyDSTimeServiceFlag() (value bool, err error) {
	retValue, err := instance.GetProperty("DSTimeServiceFlag")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDSWritableFlag sets the value of DSWritableFlag for the instance
func (instance *Win32_NTDomain) SetPropertyDSWritableFlag(value bool) (err error) {
	return instance.SetProperty("DSWritableFlag", value)
}

// GetDSWritableFlag gets the value of DSWritableFlag for the instance
func (instance *Win32_NTDomain) GetPropertyDSWritableFlag() (value bool, err error) {
	retValue, err := instance.GetProperty("DSWritableFlag")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
