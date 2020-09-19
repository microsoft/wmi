// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_TcpIpPrinterPort struct
type MSFT_TcpIpPrinterPort struct {
	*MSFT_PrinterPort

	//
	LprByteCounting bool

	//
	LprQueueName string

	//
	PortNumber uint32

	//
	PrinterHostAddress string

	//
	PrinterHostIP string

	//
	Protocol uint32

	//
	SNMPCommunity string

	//
	SNMPEnabled bool

	//
	SNMPIndex uint32
}

func NewMSFT_TcpIpPrinterPortEx1(instance *cim.WmiInstance) (newInstance *MSFT_TcpIpPrinterPort, err error) {
	tmp, err := NewMSFT_PrinterPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TcpIpPrinterPort{
		MSFT_PrinterPort: tmp,
	}
	return
}

func NewMSFT_TcpIpPrinterPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TcpIpPrinterPort, err error) {
	tmp, err := NewMSFT_PrinterPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TcpIpPrinterPort{
		MSFT_PrinterPort: tmp,
	}
	return
}

// SetLprByteCounting sets the value of LprByteCounting for the instance
func (instance *MSFT_TcpIpPrinterPort) SetPropertyLprByteCounting(value bool) (err error) {
	return instance.SetProperty("LprByteCounting", (value))
}

// GetLprByteCounting gets the value of LprByteCounting for the instance
func (instance *MSFT_TcpIpPrinterPort) GetPropertyLprByteCounting() (value bool, err error) {
	retValue, err := instance.GetProperty("LprByteCounting")
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

// SetLprQueueName sets the value of LprQueueName for the instance
func (instance *MSFT_TcpIpPrinterPort) SetPropertyLprQueueName(value string) (err error) {
	return instance.SetProperty("LprQueueName", (value))
}

// GetLprQueueName gets the value of LprQueueName for the instance
func (instance *MSFT_TcpIpPrinterPort) GetPropertyLprQueueName() (value string, err error) {
	retValue, err := instance.GetProperty("LprQueueName")
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

// SetPortNumber sets the value of PortNumber for the instance
func (instance *MSFT_TcpIpPrinterPort) SetPropertyPortNumber(value uint32) (err error) {
	return instance.SetProperty("PortNumber", (value))
}

// GetPortNumber gets the value of PortNumber for the instance
func (instance *MSFT_TcpIpPrinterPort) GetPropertyPortNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortNumber")
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

// SetPrinterHostAddress sets the value of PrinterHostAddress for the instance
func (instance *MSFT_TcpIpPrinterPort) SetPropertyPrinterHostAddress(value string) (err error) {
	return instance.SetProperty("PrinterHostAddress", (value))
}

// GetPrinterHostAddress gets the value of PrinterHostAddress for the instance
func (instance *MSFT_TcpIpPrinterPort) GetPropertyPrinterHostAddress() (value string, err error) {
	retValue, err := instance.GetProperty("PrinterHostAddress")
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

// SetPrinterHostIP sets the value of PrinterHostIP for the instance
func (instance *MSFT_TcpIpPrinterPort) SetPropertyPrinterHostIP(value string) (err error) {
	return instance.SetProperty("PrinterHostIP", (value))
}

// GetPrinterHostIP gets the value of PrinterHostIP for the instance
func (instance *MSFT_TcpIpPrinterPort) GetPropertyPrinterHostIP() (value string, err error) {
	retValue, err := instance.GetProperty("PrinterHostIP")
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

// SetProtocol sets the value of Protocol for the instance
func (instance *MSFT_TcpIpPrinterPort) SetPropertyProtocol(value uint32) (err error) {
	return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *MSFT_TcpIpPrinterPort) GetPropertyProtocol() (value uint32, err error) {
	retValue, err := instance.GetProperty("Protocol")
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

// SetSNMPCommunity sets the value of SNMPCommunity for the instance
func (instance *MSFT_TcpIpPrinterPort) SetPropertySNMPCommunity(value string) (err error) {
	return instance.SetProperty("SNMPCommunity", (value))
}

// GetSNMPCommunity gets the value of SNMPCommunity for the instance
func (instance *MSFT_TcpIpPrinterPort) GetPropertySNMPCommunity() (value string, err error) {
	retValue, err := instance.GetProperty("SNMPCommunity")
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

// SetSNMPEnabled sets the value of SNMPEnabled for the instance
func (instance *MSFT_TcpIpPrinterPort) SetPropertySNMPEnabled(value bool) (err error) {
	return instance.SetProperty("SNMPEnabled", (value))
}

// GetSNMPEnabled gets the value of SNMPEnabled for the instance
func (instance *MSFT_TcpIpPrinterPort) GetPropertySNMPEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("SNMPEnabled")
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

// SetSNMPIndex sets the value of SNMPIndex for the instance
func (instance *MSFT_TcpIpPrinterPort) SetPropertySNMPIndex(value uint32) (err error) {
	return instance.SetProperty("SNMPIndex", (value))
}

// GetSNMPIndex gets the value of SNMPIndex for the instance
func (instance *MSFT_TcpIpPrinterPort) GetPropertySNMPIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("SNMPIndex")
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
