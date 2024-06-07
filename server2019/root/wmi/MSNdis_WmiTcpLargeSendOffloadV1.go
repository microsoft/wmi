// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSNdis_WmiTcpLargeSendOffloadV1 struct
type MSNdis_WmiTcpLargeSendOffloadV1 struct { 
	*MSNdis

	// 
	WmiIPv4 MSNdis_WmiTcpLargeSendOffloadV1_IPv4
}

	func NewMSNdis_WmiTcpLargeSendOffloadV1Ex1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiTcpLargeSendOffloadV1, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_WmiTcpLargeSendOffloadV1 {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_WmiTcpLargeSendOffloadV1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_WmiTcpLargeSendOffloadV1, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_WmiTcpLargeSendOffloadV1 {
	MSNdis: tmp,
	}
	return
	}
	

// SetWmiIPv4 sets the value of WmiIPv4 for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV1) SetPropertyWmiIPv4(value MSNdis_WmiTcpLargeSendOffloadV1_IPv4) (err error) { 
    return instance.SetProperty("WmiIPv4", (value))
}

// GetWmiIPv4 gets the value of WmiIPv4 for the instance
func (instance *MSNdis_WmiTcpLargeSendOffloadV1) GetPropertyWmiIPv4()(value MSNdis_WmiTcpLargeSendOffloadV1_IPv4, err error) { 
    retValue, err := instance.GetProperty("WmiIPv4")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_WmiTcpLargeSendOffloadV1_IPv4); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiTcpLargeSendOffloadV1_IPv4 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_WmiTcpLargeSendOffloadV1_IPv4(valuetmp)

    return
}

