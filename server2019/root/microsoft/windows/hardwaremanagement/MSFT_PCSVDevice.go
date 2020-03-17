// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.HardwareManagement
//////////////////////////////////////////////
package hardwaremanagement

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_PCSVDevice struct
type MSFT_PCSVDevice struct {
	CIM_PhysicalComputerSystemView

	//
	IPv4Address string

	//
	IPv4AddressOrigin uint16

	//
	IPv4DefaultGateway string

	//
	IPv4SubnetMask string

	//
	LogFreeSpace uint16

	//
	MacAddress string

	//
	SMBIOSGuid string

	//
	TargetAddress string
}

// SetIPv4Address sets the value of IPv4Address for the instance
func (instance *MSFT_PCSVDevice) SetPropertyIPv4Address(value string) (err error) {
	return instance.SetProperty("IPv4Address", value)
}

// GetIPv4Address gets the value of IPv4Address for the instance
func (instance *MSFT_PCSVDevice) GetPropertyIPv4Address() (value string, err error) {
	retValue, err := instance.GetProperty("IPv4Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4AddressOrigin sets the value of IPv4AddressOrigin for the instance
func (instance *MSFT_PCSVDevice) SetPropertyIPv4AddressOrigin(value uint16) (err error) {
	return instance.SetProperty("IPv4AddressOrigin", value)
}

// GetIPv4AddressOrigin gets the value of IPv4AddressOrigin for the instance
func (instance *MSFT_PCSVDevice) GetPropertyIPv4AddressOrigin() (value uint16, err error) {
	retValue, err := instance.GetProperty("IPv4AddressOrigin")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4DefaultGateway sets the value of IPv4DefaultGateway for the instance
func (instance *MSFT_PCSVDevice) SetPropertyIPv4DefaultGateway(value string) (err error) {
	return instance.SetProperty("IPv4DefaultGateway", value)
}

// GetIPv4DefaultGateway gets the value of IPv4DefaultGateway for the instance
func (instance *MSFT_PCSVDevice) GetPropertyIPv4DefaultGateway() (value string, err error) {
	retValue, err := instance.GetProperty("IPv4DefaultGateway")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4SubnetMask sets the value of IPv4SubnetMask for the instance
func (instance *MSFT_PCSVDevice) SetPropertyIPv4SubnetMask(value string) (err error) {
	return instance.SetProperty("IPv4SubnetMask", value)
}

// GetIPv4SubnetMask gets the value of IPv4SubnetMask for the instance
func (instance *MSFT_PCSVDevice) GetPropertyIPv4SubnetMask() (value string, err error) {
	retValue, err := instance.GetProperty("IPv4SubnetMask")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogFreeSpace sets the value of LogFreeSpace for the instance
func (instance *MSFT_PCSVDevice) SetPropertyLogFreeSpace(value uint16) (err error) {
	return instance.SetProperty("LogFreeSpace", value)
}

// GetLogFreeSpace gets the value of LogFreeSpace for the instance
func (instance *MSFT_PCSVDevice) GetPropertyLogFreeSpace() (value uint16, err error) {
	retValue, err := instance.GetProperty("LogFreeSpace")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMacAddress sets the value of MacAddress for the instance
func (instance *MSFT_PCSVDevice) SetPropertyMacAddress(value string) (err error) {
	return instance.SetProperty("MacAddress", value)
}

// GetMacAddress gets the value of MacAddress for the instance
func (instance *MSFT_PCSVDevice) GetPropertyMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MacAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSMBIOSGuid sets the value of SMBIOSGuid for the instance
func (instance *MSFT_PCSVDevice) SetPropertySMBIOSGuid(value string) (err error) {
	return instance.SetProperty("SMBIOSGuid", value)
}

// GetSMBIOSGuid gets the value of SMBIOSGuid for the instance
func (instance *MSFT_PCSVDevice) GetPropertySMBIOSGuid() (value string, err error) {
	retValue, err := instance.GetProperty("SMBIOSGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetAddress sets the value of TargetAddress for the instance
func (instance *MSFT_PCSVDevice) SetPropertyTargetAddress(value string) (err error) {
	return instance.SetProperty("TargetAddress", value)
}

// GetTargetAddress gets the value of TargetAddress for the instance
func (instance *MSFT_PCSVDevice) GetPropertyTargetAddress() (value string, err error) {
	retValue, err := instance.GetProperty("TargetAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="OneTimeBootSource" type="string "></param>
// <param name="PersistentBootSource" type="string []"></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PCSVDevice) ChangeBootConfiguration( /* IN */ OneTimeBootSource string,
	/* IN */ PersistentBootSource []string,
	/* IN/OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ChangeBootConfiguration", Action, PercentComplete, Timeout, OneTimeBootSource, PersistentBootSource)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IPv4Address" type="string "></param>
// <param name="IPv4AddressOrigin" type="uint16 "></param>
// <param name="IPv4DefaultGateway" type="string "></param>
// <param name="IPv4SubnetMask" type="string "></param>
// <param name="Job" type="CIM_ConcreteJob "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PCSVDevice) ChangeNetworkConfiguration( /* IN */ IPv4AddressOrigin uint16,
	/* IN */ IPv4Address string,
	/* IN */ IPv4SubnetMask string,
	/* IN */ IPv4DefaultGateway string,
	/* IN/OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ChangeNetworkConfiguration", Action, PercentComplete, Timeout, IPv4AddressOrigin, IPv4Address, IPv4SubnetMask, IPv4DefaultGateway)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CurrentCredential" type="string "></param>
// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="NewPassword" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PCSVDevice) ChangeUserPassword( /* IN */ CurrentCredential string,
	/* IN */ NewPassword string,
	/* IN/OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ChangeUserPassword", Action, PercentComplete, Timeout, CurrentCredential, NewPassword)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Job" type="CIM_ConcreteJob "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="LogRecords" type="MSFT_PCSVLogRecord []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PCSVDevice) ReadLog( /* OUT */ LogRecords []MSFT_PCSVLogRecord,
	/* IN/OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ReadLog", Action, PercentComplete, Timeout)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
