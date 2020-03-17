// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_DNSProtocolEndpoint struct
type CIM_DNSProtocolEndpoint struct {
	CIM_ProtocolEndpoint

	// 650
	DHCPOptionsToUse []DNSProtocolEndpoint_DHCPOptionsToUse

	// 649
	Hostname string
}

// SetDHCPOptionsToUse sets the value of DHCPOptionsToUse for the instance
func (instance *CIM_DNSProtocolEndpoint) SetPropertyDHCPOptionsToUse(value []DNSProtocolEndpoint_DHCPOptionsToUse) (err error) {
	return instance.SetProperty("DHCPOptionsToUse", value)
}

// GetDHCPOptionsToUse gets the value of DHCPOptionsToUse for the instance
func (instance *CIM_DNSProtocolEndpoint) GetPropertyDHCPOptionsToUse() (value []DNSProtocolEndpoint_DHCPOptionsToUse, err error) {
	retValue, err := instance.GetProperty("DHCPOptionsToUse")
	if err != nil {
		return
	}
	value, ok := retValue.([]DNSProtocolEndpoint_DHCPOptionsToUse)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostname sets the value of Hostname for the instance
func (instance *CIM_DNSProtocolEndpoint) SetPropertyHostname(value string) (err error) {
	return instance.SetProperty("Hostname", value)
}

// GetHostname gets the value of Hostname for the instance
func (instance *CIM_DNSProtocolEndpoint) GetPropertyHostname() (value string, err error) {
	retValue, err := instance.GetProperty("Hostname")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
