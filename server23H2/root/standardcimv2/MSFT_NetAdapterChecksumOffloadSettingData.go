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

// MSFT_NetAdapterChecksumOffloadSettingData struct
type MSFT_NetAdapterChecksumOffloadSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	ChecksumOffloadHardwareCapabilities MSFT_NetAdapterChecksumOffloadCapabilities

	//
	IpIPv4Enabled uint32

	//
	TcpIPv4Enabled uint32

	//
	TcpIPv6Enabled uint32

	//
	UdpIPv4Enabled uint32

	//
	UdpIPv6Enabled uint32
}

func NewMSFT_NetAdapterChecksumOffloadSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterChecksumOffloadSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterChecksumOffloadSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterChecksumOffloadSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterChecksumOffloadSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterChecksumOffloadSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetChecksumOffloadHardwareCapabilities sets the value of ChecksumOffloadHardwareCapabilities for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) SetPropertyChecksumOffloadHardwareCapabilities(value MSFT_NetAdapterChecksumOffloadCapabilities) (err error) {
	return instance.SetProperty("ChecksumOffloadHardwareCapabilities", (value))
}

// GetChecksumOffloadHardwareCapabilities gets the value of ChecksumOffloadHardwareCapabilities for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) GetPropertyChecksumOffloadHardwareCapabilities() (value MSFT_NetAdapterChecksumOffloadCapabilities, err error) {
	retValue, err := instance.GetProperty("ChecksumOffloadHardwareCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterChecksumOffloadCapabilities)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterChecksumOffloadCapabilities is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterChecksumOffloadCapabilities(valuetmp)

	return
}

// SetIpIPv4Enabled sets the value of IpIPv4Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) SetPropertyIpIPv4Enabled(value uint32) (err error) {
	return instance.SetProperty("IpIPv4Enabled", (value))
}

// GetIpIPv4Enabled gets the value of IpIPv4Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) GetPropertyIpIPv4Enabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpIPv4Enabled")
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

// SetTcpIPv4Enabled sets the value of TcpIPv4Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) SetPropertyTcpIPv4Enabled(value uint32) (err error) {
	return instance.SetProperty("TcpIPv4Enabled", (value))
}

// GetTcpIPv4Enabled gets the value of TcpIPv4Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) GetPropertyTcpIPv4Enabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpIPv4Enabled")
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

// SetTcpIPv6Enabled sets the value of TcpIPv6Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) SetPropertyTcpIPv6Enabled(value uint32) (err error) {
	return instance.SetProperty("TcpIPv6Enabled", (value))
}

// GetTcpIPv6Enabled gets the value of TcpIPv6Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) GetPropertyTcpIPv6Enabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpIPv6Enabled")
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

// SetUdpIPv4Enabled sets the value of UdpIPv4Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) SetPropertyUdpIPv4Enabled(value uint32) (err error) {
	return instance.SetProperty("UdpIPv4Enabled", (value))
}

// GetUdpIPv4Enabled gets the value of UdpIPv4Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) GetPropertyUdpIPv4Enabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("UdpIPv4Enabled")
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

// SetUdpIPv6Enabled sets the value of UdpIPv6Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) SetPropertyUdpIPv6Enabled(value uint32) (err error) {
	return instance.SetProperty("UdpIPv6Enabled", (value))
}

// GetUdpIPv6Enabled gets the value of UdpIPv6Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) GetPropertyUdpIPv6Enabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("UdpIPv6Enabled")
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

//

// <param name="IpIPv4" type="bool "></param>
// <param name="RxTxControl" type="uint32 "></param>
// <param name="TcpIPv4" type="bool "></param>
// <param name="TcpIPv6" type="bool "></param>
// <param name="UdpIPv4" type="bool "></param>
// <param name="UdpIPv6" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterChecksumOffloadSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) Enable( /* IN */ IpIPv4 bool,
	/* IN */ TcpIPv4 bool,
	/* IN */ TcpIPv6 bool,
	/* IN */ UdpIPv4 bool,
	/* IN */ UdpIPv6 bool,
	/* IN */ RxTxControl uint32,
	/* OUT */ cmdletOutput MSFT_NetAdapterChecksumOffloadSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable", IpIPv4, TcpIPv4, TcpIPv6, UdpIPv4, UdpIPv6, RxTxControl)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IpIPv4" type="bool "></param>
// <param name="RxTxControl" type="uint32 "></param>
// <param name="TcpIPv4" type="bool "></param>
// <param name="TcpIPv6" type="bool "></param>
// <param name="UdpIPv4" type="bool "></param>
// <param name="UdpIPv6" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterChecksumOffloadSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) Disable( /* IN */ IpIPv4 bool,
	/* IN */ TcpIPv4 bool,
	/* IN */ TcpIPv6 bool,
	/* IN */ UdpIPv4 bool,
	/* IN */ UdpIPv6 bool,
	/* IN */ RxTxControl uint32,
	/* OUT */ cmdletOutput MSFT_NetAdapterChecksumOffloadSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable", IpIPv4, TcpIPv4, TcpIPv6, UdpIPv4, UdpIPv6, RxTxControl)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
