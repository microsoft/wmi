// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapterChecksumOffloadCapabilities struct
type MSFT_NetAdapterChecksumOffloadCapabilities struct {
	*cim.WmiInstance

	//
	IPv4ReceiveEncapsulation MSFT_NetAdapterChecksumOffloadEncapsulationTypes

	//
	IPv4ReceiveIpChecksumSupported bool

	//
	IPv4ReceiveIpOptionsSupported bool

	//
	IPv4ReceiveTcpChecksumSupported bool

	//
	IPv4ReceiveTcpOptionsSupported bool

	//
	IPv4ReceiveUdpChecksumSupported bool

	//
	IPv4TransmitEncapsulation MSFT_NetAdapterChecksumOffloadEncapsulationTypes

	//
	IPv4TransmitIpChecksumSupported bool

	//
	IPv4TransmitIpOptionsSupported bool

	//
	IPv4TransmitTcpChecksumSupported bool

	//
	IPv4TransmitTcpOptionsSupported bool

	//
	IPv4TransmitUdpChecksumSupported bool

	//
	IPv6ReceiveEncapsulation MSFT_NetAdapterChecksumOffloadEncapsulationTypes

	//
	IPv6ReceiveIpExtensionHeadersSupported bool

	//
	IPv6ReceiveTcpChecksumSupported bool

	//
	IPv6ReceiveTcpOptionsSupported bool

	//
	IPv6ReceiveUdpChecksumSupported bool

	//
	IPv6TransmitEncapsulation MSFT_NetAdapterChecksumOffloadEncapsulationTypes

	//
	IPv6TransmitIpExtensionHeadersSupported bool

	//
	IPv6TransmitTcpChecksumSupported bool

	//
	IPv6TransmitTcpOptionsSupported bool

	//
	IPv6TransmitUdpChecksumSupported bool
}

func NewMSFT_NetAdapterChecksumOffloadCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterChecksumOffloadCapabilities, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterChecksumOffloadCapabilities{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapterChecksumOffloadCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterChecksumOffloadCapabilities, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterChecksumOffloadCapabilities{
		WmiInstance: tmp,
	}
	return
}

// SetIPv4ReceiveEncapsulation sets the value of IPv4ReceiveEncapsulation for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv4ReceiveEncapsulation(value MSFT_NetAdapterChecksumOffloadEncapsulationTypes) (err error) {
	return instance.SetProperty("IPv4ReceiveEncapsulation", (value))
}

// GetIPv4ReceiveEncapsulation gets the value of IPv4ReceiveEncapsulation for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv4ReceiveEncapsulation() (value MSFT_NetAdapterChecksumOffloadEncapsulationTypes, err error) {
	retValue, err := instance.GetProperty("IPv4ReceiveEncapsulation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterChecksumOffloadEncapsulationTypes)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterChecksumOffloadEncapsulationTypes is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterChecksumOffloadEncapsulationTypes(valuetmp)

	return
}

// SetIPv4ReceiveIpChecksumSupported sets the value of IPv4ReceiveIpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv4ReceiveIpChecksumSupported(value bool) (err error) {
	return instance.SetProperty("IPv4ReceiveIpChecksumSupported", (value))
}

// GetIPv4ReceiveIpChecksumSupported gets the value of IPv4ReceiveIpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv4ReceiveIpChecksumSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4ReceiveIpChecksumSupported")
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

// SetIPv4ReceiveIpOptionsSupported sets the value of IPv4ReceiveIpOptionsSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv4ReceiveIpOptionsSupported(value bool) (err error) {
	return instance.SetProperty("IPv4ReceiveIpOptionsSupported", (value))
}

// GetIPv4ReceiveIpOptionsSupported gets the value of IPv4ReceiveIpOptionsSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv4ReceiveIpOptionsSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4ReceiveIpOptionsSupported")
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

// SetIPv4ReceiveTcpChecksumSupported sets the value of IPv4ReceiveTcpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv4ReceiveTcpChecksumSupported(value bool) (err error) {
	return instance.SetProperty("IPv4ReceiveTcpChecksumSupported", (value))
}

// GetIPv4ReceiveTcpChecksumSupported gets the value of IPv4ReceiveTcpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv4ReceiveTcpChecksumSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4ReceiveTcpChecksumSupported")
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

// SetIPv4ReceiveTcpOptionsSupported sets the value of IPv4ReceiveTcpOptionsSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv4ReceiveTcpOptionsSupported(value bool) (err error) {
	return instance.SetProperty("IPv4ReceiveTcpOptionsSupported", (value))
}

// GetIPv4ReceiveTcpOptionsSupported gets the value of IPv4ReceiveTcpOptionsSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv4ReceiveTcpOptionsSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4ReceiveTcpOptionsSupported")
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

// SetIPv4ReceiveUdpChecksumSupported sets the value of IPv4ReceiveUdpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv4ReceiveUdpChecksumSupported(value bool) (err error) {
	return instance.SetProperty("IPv4ReceiveUdpChecksumSupported", (value))
}

// GetIPv4ReceiveUdpChecksumSupported gets the value of IPv4ReceiveUdpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv4ReceiveUdpChecksumSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4ReceiveUdpChecksumSupported")
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

// SetIPv4TransmitEncapsulation sets the value of IPv4TransmitEncapsulation for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv4TransmitEncapsulation(value MSFT_NetAdapterChecksumOffloadEncapsulationTypes) (err error) {
	return instance.SetProperty("IPv4TransmitEncapsulation", (value))
}

// GetIPv4TransmitEncapsulation gets the value of IPv4TransmitEncapsulation for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv4TransmitEncapsulation() (value MSFT_NetAdapterChecksumOffloadEncapsulationTypes, err error) {
	retValue, err := instance.GetProperty("IPv4TransmitEncapsulation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterChecksumOffloadEncapsulationTypes)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterChecksumOffloadEncapsulationTypes is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterChecksumOffloadEncapsulationTypes(valuetmp)

	return
}

// SetIPv4TransmitIpChecksumSupported sets the value of IPv4TransmitIpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv4TransmitIpChecksumSupported(value bool) (err error) {
	return instance.SetProperty("IPv4TransmitIpChecksumSupported", (value))
}

// GetIPv4TransmitIpChecksumSupported gets the value of IPv4TransmitIpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv4TransmitIpChecksumSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4TransmitIpChecksumSupported")
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

// SetIPv4TransmitIpOptionsSupported sets the value of IPv4TransmitIpOptionsSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv4TransmitIpOptionsSupported(value bool) (err error) {
	return instance.SetProperty("IPv4TransmitIpOptionsSupported", (value))
}

// GetIPv4TransmitIpOptionsSupported gets the value of IPv4TransmitIpOptionsSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv4TransmitIpOptionsSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4TransmitIpOptionsSupported")
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

// SetIPv4TransmitTcpChecksumSupported sets the value of IPv4TransmitTcpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv4TransmitTcpChecksumSupported(value bool) (err error) {
	return instance.SetProperty("IPv4TransmitTcpChecksumSupported", (value))
}

// GetIPv4TransmitTcpChecksumSupported gets the value of IPv4TransmitTcpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv4TransmitTcpChecksumSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4TransmitTcpChecksumSupported")
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

// SetIPv4TransmitTcpOptionsSupported sets the value of IPv4TransmitTcpOptionsSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv4TransmitTcpOptionsSupported(value bool) (err error) {
	return instance.SetProperty("IPv4TransmitTcpOptionsSupported", (value))
}

// GetIPv4TransmitTcpOptionsSupported gets the value of IPv4TransmitTcpOptionsSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv4TransmitTcpOptionsSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4TransmitTcpOptionsSupported")
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

// SetIPv4TransmitUdpChecksumSupported sets the value of IPv4TransmitUdpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv4TransmitUdpChecksumSupported(value bool) (err error) {
	return instance.SetProperty("IPv4TransmitUdpChecksumSupported", (value))
}

// GetIPv4TransmitUdpChecksumSupported gets the value of IPv4TransmitUdpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv4TransmitUdpChecksumSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4TransmitUdpChecksumSupported")
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

// SetIPv6ReceiveEncapsulation sets the value of IPv6ReceiveEncapsulation for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv6ReceiveEncapsulation(value MSFT_NetAdapterChecksumOffloadEncapsulationTypes) (err error) {
	return instance.SetProperty("IPv6ReceiveEncapsulation", (value))
}

// GetIPv6ReceiveEncapsulation gets the value of IPv6ReceiveEncapsulation for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv6ReceiveEncapsulation() (value MSFT_NetAdapterChecksumOffloadEncapsulationTypes, err error) {
	retValue, err := instance.GetProperty("IPv6ReceiveEncapsulation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterChecksumOffloadEncapsulationTypes)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterChecksumOffloadEncapsulationTypes is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterChecksumOffloadEncapsulationTypes(valuetmp)

	return
}

// SetIPv6ReceiveIpExtensionHeadersSupported sets the value of IPv6ReceiveIpExtensionHeadersSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv6ReceiveIpExtensionHeadersSupported(value bool) (err error) {
	return instance.SetProperty("IPv6ReceiveIpExtensionHeadersSupported", (value))
}

// GetIPv6ReceiveIpExtensionHeadersSupported gets the value of IPv6ReceiveIpExtensionHeadersSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv6ReceiveIpExtensionHeadersSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6ReceiveIpExtensionHeadersSupported")
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

// SetIPv6ReceiveTcpChecksumSupported sets the value of IPv6ReceiveTcpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv6ReceiveTcpChecksumSupported(value bool) (err error) {
	return instance.SetProperty("IPv6ReceiveTcpChecksumSupported", (value))
}

// GetIPv6ReceiveTcpChecksumSupported gets the value of IPv6ReceiveTcpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv6ReceiveTcpChecksumSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6ReceiveTcpChecksumSupported")
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

// SetIPv6ReceiveTcpOptionsSupported sets the value of IPv6ReceiveTcpOptionsSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv6ReceiveTcpOptionsSupported(value bool) (err error) {
	return instance.SetProperty("IPv6ReceiveTcpOptionsSupported", (value))
}

// GetIPv6ReceiveTcpOptionsSupported gets the value of IPv6ReceiveTcpOptionsSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv6ReceiveTcpOptionsSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6ReceiveTcpOptionsSupported")
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

// SetIPv6ReceiveUdpChecksumSupported sets the value of IPv6ReceiveUdpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv6ReceiveUdpChecksumSupported(value bool) (err error) {
	return instance.SetProperty("IPv6ReceiveUdpChecksumSupported", (value))
}

// GetIPv6ReceiveUdpChecksumSupported gets the value of IPv6ReceiveUdpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv6ReceiveUdpChecksumSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6ReceiveUdpChecksumSupported")
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

// SetIPv6TransmitEncapsulation sets the value of IPv6TransmitEncapsulation for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv6TransmitEncapsulation(value MSFT_NetAdapterChecksumOffloadEncapsulationTypes) (err error) {
	return instance.SetProperty("IPv6TransmitEncapsulation", (value))
}

// GetIPv6TransmitEncapsulation gets the value of IPv6TransmitEncapsulation for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv6TransmitEncapsulation() (value MSFT_NetAdapterChecksumOffloadEncapsulationTypes, err error) {
	retValue, err := instance.GetProperty("IPv6TransmitEncapsulation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterChecksumOffloadEncapsulationTypes)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterChecksumOffloadEncapsulationTypes is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterChecksumOffloadEncapsulationTypes(valuetmp)

	return
}

// SetIPv6TransmitIpExtensionHeadersSupported sets the value of IPv6TransmitIpExtensionHeadersSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv6TransmitIpExtensionHeadersSupported(value bool) (err error) {
	return instance.SetProperty("IPv6TransmitIpExtensionHeadersSupported", (value))
}

// GetIPv6TransmitIpExtensionHeadersSupported gets the value of IPv6TransmitIpExtensionHeadersSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv6TransmitIpExtensionHeadersSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6TransmitIpExtensionHeadersSupported")
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

// SetIPv6TransmitTcpChecksumSupported sets the value of IPv6TransmitTcpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv6TransmitTcpChecksumSupported(value bool) (err error) {
	return instance.SetProperty("IPv6TransmitTcpChecksumSupported", (value))
}

// GetIPv6TransmitTcpChecksumSupported gets the value of IPv6TransmitTcpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv6TransmitTcpChecksumSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6TransmitTcpChecksumSupported")
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

// SetIPv6TransmitTcpOptionsSupported sets the value of IPv6TransmitTcpOptionsSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv6TransmitTcpOptionsSupported(value bool) (err error) {
	return instance.SetProperty("IPv6TransmitTcpOptionsSupported", (value))
}

// GetIPv6TransmitTcpOptionsSupported gets the value of IPv6TransmitTcpOptionsSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv6TransmitTcpOptionsSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6TransmitTcpOptionsSupported")
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

// SetIPv6TransmitUdpChecksumSupported sets the value of IPv6TransmitUdpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) SetPropertyIPv6TransmitUdpChecksumSupported(value bool) (err error) {
	return instance.SetProperty("IPv6TransmitUdpChecksumSupported", (value))
}

// GetIPv6TransmitUdpChecksumSupported gets the value of IPv6TransmitUdpChecksumSupported for the instance
func (instance *MSFT_NetAdapterChecksumOffloadCapabilities) GetPropertyIPv6TransmitUdpChecksumSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6TransmitUdpChecksumSupported")
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
