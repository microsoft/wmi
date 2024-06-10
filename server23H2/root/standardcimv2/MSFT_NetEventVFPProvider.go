// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetEventVFPProvider struct
type MSFT_NetEventVFPProvider struct {
	*MSFT_NetEventProviderBase

	//
	DestinationIPAddresses []string

	//
	DestinationMACAddresses []string

	//
	GREKeys []uint32

	//
	IPProtocols []uint8

	//
	PortIds []uint32

	//
	SourceIPAddresses []string

	//
	SourceMACAddresses []string

	//
	SwitchName string

	//
	TCPPorts []uint16

	//
	TenantIds []uint32

	//
	UDPPorts []uint16

	//
	VFPFlowDirection uint32

	//
	VLANIds []uint16
}

func NewMSFT_NetEventVFPProviderEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetEventVFPProvider, err error) {
	tmp, err := NewMSFT_NetEventProviderBaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventVFPProvider{
		MSFT_NetEventProviderBase: tmp,
	}
	return
}

func NewMSFT_NetEventVFPProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetEventVFPProvider, err error) {
	tmp, err := NewMSFT_NetEventProviderBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventVFPProvider{
		MSFT_NetEventProviderBase: tmp,
	}
	return
}

// SetDestinationIPAddresses sets the value of DestinationIPAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyDestinationIPAddresses(value []string) (err error) {
	return instance.SetProperty("DestinationIPAddresses", (value))
}

// GetDestinationIPAddresses gets the value of DestinationIPAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyDestinationIPAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("DestinationIPAddresses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDestinationMACAddresses sets the value of DestinationMACAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyDestinationMACAddresses(value []string) (err error) {
	return instance.SetProperty("DestinationMACAddresses", (value))
}

// GetDestinationMACAddresses gets the value of DestinationMACAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyDestinationMACAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("DestinationMACAddresses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetGREKeys sets the value of GREKeys for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyGREKeys(value []uint32) (err error) {
	return instance.SetProperty("GREKeys", (value))
}

// GetGREKeys gets the value of GREKeys for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyGREKeys() (value []uint32, err error) {
	retValue, err := instance.GetProperty("GREKeys")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetIPProtocols sets the value of IPProtocols for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyIPProtocols(value []uint8) (err error) {
	return instance.SetProperty("IPProtocols", (value))
}

// GetIPProtocols gets the value of IPProtocols for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyIPProtocols() (value []uint8, err error) {
	retValue, err := instance.GetProperty("IPProtocols")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetPortIds sets the value of PortIds for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyPortIds(value []uint32) (err error) {
	return instance.SetProperty("PortIds", (value))
}

// GetPortIds gets the value of PortIds for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyPortIds() (value []uint32, err error) {
	retValue, err := instance.GetProperty("PortIds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetSourceIPAddresses sets the value of SourceIPAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertySourceIPAddresses(value []string) (err error) {
	return instance.SetProperty("SourceIPAddresses", (value))
}

// GetSourceIPAddresses gets the value of SourceIPAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertySourceIPAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("SourceIPAddresses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetSourceMACAddresses sets the value of SourceMACAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertySourceMACAddresses(value []string) (err error) {
	return instance.SetProperty("SourceMACAddresses", (value))
}

// GetSourceMACAddresses gets the value of SourceMACAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertySourceMACAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("SourceMACAddresses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetSwitchName sets the value of SwitchName for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertySwitchName(value string) (err error) {
	return instance.SetProperty("SwitchName", (value))
}

// GetSwitchName gets the value of SwitchName for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertySwitchName() (value string, err error) {
	retValue, err := instance.GetProperty("SwitchName")
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

// SetTCPPorts sets the value of TCPPorts for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyTCPPorts(value []uint16) (err error) {
	return instance.SetProperty("TCPPorts", (value))
}

// GetTCPPorts gets the value of TCPPorts for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyTCPPorts() (value []uint16, err error) {
	retValue, err := instance.GetProperty("TCPPorts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetTenantIds sets the value of TenantIds for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyTenantIds(value []uint32) (err error) {
	return instance.SetProperty("TenantIds", (value))
}

// GetTenantIds gets the value of TenantIds for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyTenantIds() (value []uint32, err error) {
	retValue, err := instance.GetProperty("TenantIds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetUDPPorts sets the value of UDPPorts for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyUDPPorts(value []uint16) (err error) {
	return instance.SetProperty("UDPPorts", (value))
}

// GetUDPPorts gets the value of UDPPorts for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyUDPPorts() (value []uint16, err error) {
	retValue, err := instance.GetProperty("UDPPorts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetVFPFlowDirection sets the value of VFPFlowDirection for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyVFPFlowDirection(value uint32) (err error) {
	return instance.SetProperty("VFPFlowDirection", (value))
}

// GetVFPFlowDirection gets the value of VFPFlowDirection for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyVFPFlowDirection() (value uint32, err error) {
	retValue, err := instance.GetProperty("VFPFlowDirection")
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

// SetVLANIds sets the value of VLANIds for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyVLANIds(value []uint16) (err error) {
	return instance.SetProperty("VLANIds", (value))
}

// GetVLANIds gets the value of VLANIds for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyVLANIds() (value []uint16, err error) {
	retValue, err := instance.GetProperty("VLANIds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}
