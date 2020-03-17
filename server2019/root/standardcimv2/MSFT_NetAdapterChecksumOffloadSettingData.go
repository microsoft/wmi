// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetAdapterChecksumOffloadSettingData struct
type MSFT_NetAdapterChecksumOffloadSettingData struct {
	MSFT_NetAdapterSettingData

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

// SetChecksumOffloadHardwareCapabilities sets the value of ChecksumOffloadHardwareCapabilities for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) SetPropertyChecksumOffloadHardwareCapabilities(value MSFT_NetAdapterChecksumOffloadCapabilities) (err error) {
	return instance.SetProperty("ChecksumOffloadHardwareCapabilities", value)
}

// GetChecksumOffloadHardwareCapabilities gets the value of ChecksumOffloadHardwareCapabilities for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) GetPropertyChecksumOffloadHardwareCapabilities() (value MSFT_NetAdapterChecksumOffloadCapabilities, err error) {
	retValue, err := instance.GetProperty("ChecksumOffloadHardwareCapabilities")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetAdapterChecksumOffloadCapabilities)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIpIPv4Enabled sets the value of IpIPv4Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) SetPropertyIpIPv4Enabled(value uint32) (err error) {
	return instance.SetProperty("IpIPv4Enabled", value)
}

// GetIpIPv4Enabled gets the value of IpIPv4Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) GetPropertyIpIPv4Enabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpIPv4Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTcpIPv4Enabled sets the value of TcpIPv4Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) SetPropertyTcpIPv4Enabled(value uint32) (err error) {
	return instance.SetProperty("TcpIPv4Enabled", value)
}

// GetTcpIPv4Enabled gets the value of TcpIPv4Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) GetPropertyTcpIPv4Enabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpIPv4Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTcpIPv6Enabled sets the value of TcpIPv6Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) SetPropertyTcpIPv6Enabled(value uint32) (err error) {
	return instance.SetProperty("TcpIPv6Enabled", value)
}

// GetTcpIPv6Enabled gets the value of TcpIPv6Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) GetPropertyTcpIPv6Enabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpIPv6Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUdpIPv4Enabled sets the value of UdpIPv4Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) SetPropertyUdpIPv4Enabled(value uint32) (err error) {
	return instance.SetProperty("UdpIPv4Enabled", value)
}

// GetUdpIPv4Enabled gets the value of UdpIPv4Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) GetPropertyUdpIPv4Enabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("UdpIPv4Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUdpIPv6Enabled sets the value of UdpIPv6Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) SetPropertyUdpIPv6Enabled(value uint32) (err error) {
	return instance.SetProperty("UdpIPv6Enabled", value)
}

// GetUdpIPv6Enabled gets the value of UdpIPv6Enabled for the instance
func (instance *MSFT_NetAdapterChecksumOffloadSettingData) GetPropertyUdpIPv6Enabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("UdpIPv6Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
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
