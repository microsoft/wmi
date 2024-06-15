// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_TcpOffloadParameters struct
type MSNdis_TcpOffloadParameters struct {
	*MSNdis

	//
	Flags uint32

	//
	Header MSNdis_ObjectHeader

	//
	IPsec uint8

	//
	IPv4Checksum uint8

	//
	LsoV1 uint8

	//
	LsoV2IPv4 uint8

	//
	LsoV2IPv6 uint8

	//
	TcpConnectionIPv4 uint8

	//
	TcpConnectionIPv6 uint8

	//
	TCPIPv4Checksum uint8

	//
	TCPIPv6Checksum uint8

	//
	UDPIPv4Checksum uint8

	//
	UDPIPv6Checksum uint8
}

func NewMSNdis_TcpOffloadParametersEx1(instance *cim.WmiInstance) (newInstance *MSNdis_TcpOffloadParameters, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_TcpOffloadParameters{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_TcpOffloadParametersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_TcpOffloadParameters, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_TcpOffloadParameters{
		MSNdis: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSNdis_TcpOffloadParameters) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_TcpOffloadParameters) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_TcpOffloadParameters) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_TcpOffloadParameters) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}

// SetIPsec sets the value of IPsec for the instance
func (instance *MSNdis_TcpOffloadParameters) SetPropertyIPsec(value uint8) (err error) {
	return instance.SetProperty("IPsec", (value))
}

// GetIPsec gets the value of IPsec for the instance
func (instance *MSNdis_TcpOffloadParameters) GetPropertyIPsec() (value uint8, err error) {
	retValue, err := instance.GetProperty("IPsec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetIPv4Checksum sets the value of IPv4Checksum for the instance
func (instance *MSNdis_TcpOffloadParameters) SetPropertyIPv4Checksum(value uint8) (err error) {
	return instance.SetProperty("IPv4Checksum", (value))
}

// GetIPv4Checksum gets the value of IPv4Checksum for the instance
func (instance *MSNdis_TcpOffloadParameters) GetPropertyIPv4Checksum() (value uint8, err error) {
	retValue, err := instance.GetProperty("IPv4Checksum")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetLsoV1 sets the value of LsoV1 for the instance
func (instance *MSNdis_TcpOffloadParameters) SetPropertyLsoV1(value uint8) (err error) {
	return instance.SetProperty("LsoV1", (value))
}

// GetLsoV1 gets the value of LsoV1 for the instance
func (instance *MSNdis_TcpOffloadParameters) GetPropertyLsoV1() (value uint8, err error) {
	retValue, err := instance.GetProperty("LsoV1")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetLsoV2IPv4 sets the value of LsoV2IPv4 for the instance
func (instance *MSNdis_TcpOffloadParameters) SetPropertyLsoV2IPv4(value uint8) (err error) {
	return instance.SetProperty("LsoV2IPv4", (value))
}

// GetLsoV2IPv4 gets the value of LsoV2IPv4 for the instance
func (instance *MSNdis_TcpOffloadParameters) GetPropertyLsoV2IPv4() (value uint8, err error) {
	retValue, err := instance.GetProperty("LsoV2IPv4")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetLsoV2IPv6 sets the value of LsoV2IPv6 for the instance
func (instance *MSNdis_TcpOffloadParameters) SetPropertyLsoV2IPv6(value uint8) (err error) {
	return instance.SetProperty("LsoV2IPv6", (value))
}

// GetLsoV2IPv6 gets the value of LsoV2IPv6 for the instance
func (instance *MSNdis_TcpOffloadParameters) GetPropertyLsoV2IPv6() (value uint8, err error) {
	retValue, err := instance.GetProperty("LsoV2IPv6")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetTcpConnectionIPv4 sets the value of TcpConnectionIPv4 for the instance
func (instance *MSNdis_TcpOffloadParameters) SetPropertyTcpConnectionIPv4(value uint8) (err error) {
	return instance.SetProperty("TcpConnectionIPv4", (value))
}

// GetTcpConnectionIPv4 gets the value of TcpConnectionIPv4 for the instance
func (instance *MSNdis_TcpOffloadParameters) GetPropertyTcpConnectionIPv4() (value uint8, err error) {
	retValue, err := instance.GetProperty("TcpConnectionIPv4")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetTcpConnectionIPv6 sets the value of TcpConnectionIPv6 for the instance
func (instance *MSNdis_TcpOffloadParameters) SetPropertyTcpConnectionIPv6(value uint8) (err error) {
	return instance.SetProperty("TcpConnectionIPv6", (value))
}

// GetTcpConnectionIPv6 gets the value of TcpConnectionIPv6 for the instance
func (instance *MSNdis_TcpOffloadParameters) GetPropertyTcpConnectionIPv6() (value uint8, err error) {
	retValue, err := instance.GetProperty("TcpConnectionIPv6")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetTCPIPv4Checksum sets the value of TCPIPv4Checksum for the instance
func (instance *MSNdis_TcpOffloadParameters) SetPropertyTCPIPv4Checksum(value uint8) (err error) {
	return instance.SetProperty("TCPIPv4Checksum", (value))
}

// GetTCPIPv4Checksum gets the value of TCPIPv4Checksum for the instance
func (instance *MSNdis_TcpOffloadParameters) GetPropertyTCPIPv4Checksum() (value uint8, err error) {
	retValue, err := instance.GetProperty("TCPIPv4Checksum")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetTCPIPv6Checksum sets the value of TCPIPv6Checksum for the instance
func (instance *MSNdis_TcpOffloadParameters) SetPropertyTCPIPv6Checksum(value uint8) (err error) {
	return instance.SetProperty("TCPIPv6Checksum", (value))
}

// GetTCPIPv6Checksum gets the value of TCPIPv6Checksum for the instance
func (instance *MSNdis_TcpOffloadParameters) GetPropertyTCPIPv6Checksum() (value uint8, err error) {
	retValue, err := instance.GetProperty("TCPIPv6Checksum")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetUDPIPv4Checksum sets the value of UDPIPv4Checksum for the instance
func (instance *MSNdis_TcpOffloadParameters) SetPropertyUDPIPv4Checksum(value uint8) (err error) {
	return instance.SetProperty("UDPIPv4Checksum", (value))
}

// GetUDPIPv4Checksum gets the value of UDPIPv4Checksum for the instance
func (instance *MSNdis_TcpOffloadParameters) GetPropertyUDPIPv4Checksum() (value uint8, err error) {
	retValue, err := instance.GetProperty("UDPIPv4Checksum")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetUDPIPv6Checksum sets the value of UDPIPv6Checksum for the instance
func (instance *MSNdis_TcpOffloadParameters) SetPropertyUDPIPv6Checksum(value uint8) (err error) {
	return instance.SetProperty("UDPIPv6Checksum", (value))
}

// GetUDPIPv6Checksum gets the value of UDPIPv6Checksum for the instance
func (instance *MSNdis_TcpOffloadParameters) GetPropertyUDPIPv6Checksum() (value uint8, err error) {
	retValue, err := instance.GetProperty("UDPIPv6Checksum")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
