// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ExternalEthernetPortCapabilities struct
type Msvm_ExternalEthernetPortCapabilities struct {
	CIM_Capabilities

	// A boolean value which indicates whether IOV is supported by the network adapter.If the value is TRUE, then IOV is supported by the network adapter and IOVSupportReasons will be empty. Otherwise the IOVSupportReasons property will have the reasons why IOV cannot be supported.
	IOVSupport bool

	// An array of strings that indicates the possible reasons why IOV is not supported. If the value of IOVSupport is TRUE this array will be empty.
	IOVSupportReasons []string
}

// SetIOVSupport sets the value of IOVSupport for the instance
func (instance *Msvm_ExternalEthernetPortCapabilities) SetPropertyIOVSupport(value bool) (err error) {
	return instance.SetProperty("IOVSupport", value)
}

// GetIOVSupport gets the value of IOVSupport for the instance
func (instance *Msvm_ExternalEthernetPortCapabilities) GetPropertyIOVSupport() (value bool, err error) {
	retValue, err := instance.GetProperty("IOVSupport")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOVSupportReasons sets the value of IOVSupportReasons for the instance
func (instance *Msvm_ExternalEthernetPortCapabilities) SetPropertyIOVSupportReasons(value []string) (err error) {
	return instance.SetProperty("IOVSupportReasons", value)
}

// GetIOVSupportReasons gets the value of IOVSupportReasons for the instance
func (instance *Msvm_ExternalEthernetPortCapabilities) GetPropertyIOVSupportReasons() (value []string, err error) {
	retValue, err := instance.GetProperty("IOVSupportReasons")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_ExternalEthernetPortCapabilities) GetRelatedExternalEthernetPort() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ExternalEthernetPort")
}
