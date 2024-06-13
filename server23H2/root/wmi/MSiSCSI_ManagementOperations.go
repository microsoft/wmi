// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSI_ManagementOperations struct
type MSiSCSI_ManagementOperations struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewMSiSCSI_ManagementOperationsEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_ManagementOperations, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_ManagementOperations{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSI_ManagementOperationsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_ManagementOperations, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_ManagementOperations{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_ManagementOperations) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_ManagementOperations) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_ManagementOperations) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_ManagementOperations) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// Perform an ICMP ping

// <param name="Address" type="ISCSI_IP_Address ">IP address to ping</param>
// <param name="RequestCount" type="uint32 ">Number of requests to send</param>
// <param name="RequestSize" type="uint32 ">Number of bytes in each request</param>
// <param name="Timeout" type="uint32 ">Number of ms to wait for response</param>

// <param name="ResponsesReceived" type="uint32 ">Number of responses received</param>
// <param name="Status" type="ManagementOperations_Status ">Status code resulting from operation</param>
func (instance *MSiSCSI_ManagementOperations) PingIPAddress( /* IN */ RequestCount uint32,
	/* IN */ RequestSize uint32,
	/* IN */ Timeout uint32,
	/* IN */ Address ISCSI_IP_Address,
	/* OUT */ Status ManagementOperations_Status,
	/* OUT */ ResponsesReceived uint32) (err error) {
	_, err = instance.InvokeMethod("PingIPAddress", RequestCount, RequestSize, Timeout, Address)
	if err != nil {
		return
	}
	return

}
