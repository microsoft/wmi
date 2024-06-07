// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_SmbBandwidthLimit struct
type MSFT_SmbBandwidthLimit struct { 
	*cim.WmiInstance

	// 
	BytesPerSecond uint64

	// 
	Category SmbBandwidthLimit_Category
}

	func NewMSFT_SmbBandwidthLimitEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbBandwidthLimit, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_SmbBandwidthLimit {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_SmbBandwidthLimitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_SmbBandwidthLimit, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_SmbBandwidthLimit {
	WmiInstance: tmp,
	}
	return
	}
	

// SetBytesPerSecond sets the value of BytesPerSecond for the instance
func (instance *MSFT_SmbBandwidthLimit) SetPropertyBytesPerSecond(value uint64) (err error) { 
    return instance.SetProperty("BytesPerSecond", (value))
}

// GetBytesPerSecond gets the value of BytesPerSecond for the instance
func (instance *MSFT_SmbBandwidthLimit) GetPropertyBytesPerSecond()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BytesPerSecond")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetCategory sets the value of Category for the instance
func (instance *MSFT_SmbBandwidthLimit) SetPropertyCategory(value SmbBandwidthLimit_Category) (err error) { 
    return instance.SetProperty("Category", (value))
}

// GetCategory gets the value of Category for the instance
func (instance *MSFT_SmbBandwidthLimit) GetPropertyCategory()(value SmbBandwidthLimit_Category, err error) { 
    retValue, err := instance.GetProperty("Category")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = SmbBandwidthLimit_Category(valuetmp)

    return
}

// 

// <param name="BytesPerSecond" type="uint64 "></param>
// <param name="Category" type="uint32 "></param>

// <param name="Output" type="MSFT_SmbBandwidthLimit []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbBandwidthLimit) Set( /* IN */ Category uint32,
 /* IN */ BytesPerSecond uint64,
 /* OUT */ Output []MSFT_SmbBandwidthLimit) (result uint32, err error) {retVal, err := instance.InvokeMethod("Set" , Category, BytesPerSecond)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


