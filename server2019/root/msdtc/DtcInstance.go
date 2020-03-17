// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// DtcInstance struct
type DtcInstance struct {
	cim.WmiInstance

	//
	DtcName string

	//
	KtmRmEndpointCid string

	//
	OleTxEndpointCid string

	//
	Status string

	//
	UisEndpointCid string

	//
	VirtualServerName string

	//
	XAEndpointCid string
}

// SetDtcName sets the value of DtcName for the instance
func (instance *DtcInstance) SetPropertyDtcName(value string) (err error) {
	return instance.SetProperty("DtcName", value)
}

// GetDtcName gets the value of DtcName for the instance
func (instance *DtcInstance) GetPropertyDtcName() (value string, err error) {
	retValue, err := instance.GetProperty("DtcName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKtmRmEndpointCid sets the value of KtmRmEndpointCid for the instance
func (instance *DtcInstance) SetPropertyKtmRmEndpointCid(value string) (err error) {
	return instance.SetProperty("KtmRmEndpointCid", value)
}

// GetKtmRmEndpointCid gets the value of KtmRmEndpointCid for the instance
func (instance *DtcInstance) GetPropertyKtmRmEndpointCid() (value string, err error) {
	retValue, err := instance.GetProperty("KtmRmEndpointCid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOleTxEndpointCid sets the value of OleTxEndpointCid for the instance
func (instance *DtcInstance) SetPropertyOleTxEndpointCid(value string) (err error) {
	return instance.SetProperty("OleTxEndpointCid", value)
}

// GetOleTxEndpointCid gets the value of OleTxEndpointCid for the instance
func (instance *DtcInstance) GetPropertyOleTxEndpointCid() (value string, err error) {
	retValue, err := instance.GetProperty("OleTxEndpointCid")
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
func (instance *DtcInstance) SetPropertyStatus(value string) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *DtcInstance) GetPropertyStatus() (value string, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUisEndpointCid sets the value of UisEndpointCid for the instance
func (instance *DtcInstance) SetPropertyUisEndpointCid(value string) (err error) {
	return instance.SetProperty("UisEndpointCid", value)
}

// GetUisEndpointCid gets the value of UisEndpointCid for the instance
func (instance *DtcInstance) GetPropertyUisEndpointCid() (value string, err error) {
	retValue, err := instance.GetProperty("UisEndpointCid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualServerName sets the value of VirtualServerName for the instance
func (instance *DtcInstance) SetPropertyVirtualServerName(value string) (err error) {
	return instance.SetProperty("VirtualServerName", value)
}

// GetVirtualServerName gets the value of VirtualServerName for the instance
func (instance *DtcInstance) GetPropertyVirtualServerName() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualServerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetXAEndpointCid sets the value of XAEndpointCid for the instance
func (instance *DtcInstance) SetPropertyXAEndpointCid(value string) (err error) {
	return instance.SetProperty("XAEndpointCid", value)
}

// GetXAEndpointCid gets the value of XAEndpointCid for the instance
func (instance *DtcInstance) GetPropertyXAEndpointCid() (value string, err error) {
	retValue, err := instance.GetProperty("XAEndpointCid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
