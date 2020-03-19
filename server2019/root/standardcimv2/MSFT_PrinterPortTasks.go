// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_PrinterPortTasks struct
type MSFT_PrinterPortTasks struct {
	*cim.WmiInstance
}

func NewMSFT_PrinterPortTasksEx1(instance *cim.WmiInstance) (newInstance *MSFT_PrinterPortTasks, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_PrinterPortTasks{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_PrinterPortTasksEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_PrinterPortTasks, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_PrinterPortTasks{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterPortTasks) AddByLocalPort( /* IN */ ComputerName string,
	/* IN */ Name string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddByLocalPort", ComputerName, Name)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="HostName" type="string "></param>
// <param name="PrinterName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterPortTasks) AddByLprPort( /* IN */ ComputerName string,
	/* IN */ HostName string,
	/* IN */ PrinterName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddByLprPort", ComputerName, HostName, PrinterName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="PortNumber" type="uint32 "></param>
// <param name="PrinterHostAddress" type="string "></param>
// <param name="SNMP" type="uint32 "></param>
// <param name="SNMPCommunity" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterPortTasks) AddByTcpPort( /* IN */ ComputerName string,
	/* IN */ Name string,
	/* IN */ PortNumber uint32,
	/* IN */ PrinterHostAddress string,
	/* IN */ SNMP uint32,
	/* IN */ SNMPCommunity string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddByTcpPort", ComputerName, Name, PortNumber, PrinterHostAddress, SNMP, SNMPCommunity)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="LprByteCounting" type="bool "></param>
// <param name="LprHostAddress" type="string "></param>
// <param name="LprQueueName" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="PortNumber" type="uint32 "></param>
// <param name="SNMP" type="uint32 "></param>
// <param name="SNMPCommunity" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterPortTasks) AddByTcpPortLprMode( /* IN */ ComputerName string,
	/* IN */ LprByteCounting bool,
	/* IN */ LprHostAddress string,
	/* IN */ LprQueueName string,
	/* IN */ Name string,
	/* IN */ PortNumber uint32,
	/* IN */ SNMP uint32,
	/* IN */ SNMPCommunity string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddByTcpPortLprMode", ComputerName, LprByteCounting, LprHostAddress, LprQueueName, Name, PortNumber, SNMP, SNMPCommunity)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
