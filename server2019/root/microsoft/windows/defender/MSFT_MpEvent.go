// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.Defender
//////////////////////////////////////////////
package defender
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_MpEvent struct
type MSFT_MpEvent struct { 
	*cim.WmiInstance

	// 
	AdditionalData int64

	// 
	CategoryDiscriminant uint32

	// 
	ComputerNotificationsValue uint32

	// 
	NotificationTime string

	// 
	ScanNotificationsValue uint32

	// 
	SignatureNotificationsValue uint32

	// 
	ThreatNotificationsValue uint32
}

	func NewMSFT_MpEventEx1(instance *cim.WmiInstance) (newInstance *MSFT_MpEvent, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_MpEvent {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_MpEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_MpEvent, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_MpEvent {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAdditionalData sets the value of AdditionalData for the instance
func (instance *MSFT_MpEvent) SetPropertyAdditionalData(value int64) (err error) { 
    return instance.SetProperty("AdditionalData", (value))
}

// GetAdditionalData gets the value of AdditionalData for the instance
func (instance *MSFT_MpEvent) GetPropertyAdditionalData()(value int64, err error) { 
    retValue, err := instance.GetProperty("AdditionalData")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int64(valuetmp)

    return
}

// SetCategoryDiscriminant sets the value of CategoryDiscriminant for the instance
func (instance *MSFT_MpEvent) SetPropertyCategoryDiscriminant(value uint32) (err error) { 
    return instance.SetProperty("CategoryDiscriminant", (value))
}

// GetCategoryDiscriminant gets the value of CategoryDiscriminant for the instance
func (instance *MSFT_MpEvent) GetPropertyCategoryDiscriminant()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CategoryDiscriminant")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetComputerNotificationsValue sets the value of ComputerNotificationsValue for the instance
func (instance *MSFT_MpEvent) SetPropertyComputerNotificationsValue(value uint32) (err error) { 
    return instance.SetProperty("ComputerNotificationsValue", (value))
}

// GetComputerNotificationsValue gets the value of ComputerNotificationsValue for the instance
func (instance *MSFT_MpEvent) GetPropertyComputerNotificationsValue()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ComputerNotificationsValue")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetNotificationTime sets the value of NotificationTime for the instance
func (instance *MSFT_MpEvent) SetPropertyNotificationTime(value string) (err error) { 
    return instance.SetProperty("NotificationTime", (value))
}

// GetNotificationTime gets the value of NotificationTime for the instance
func (instance *MSFT_MpEvent) GetPropertyNotificationTime()(value string, err error) { 
    retValue, err := instance.GetProperty("NotificationTime")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetScanNotificationsValue sets the value of ScanNotificationsValue for the instance
func (instance *MSFT_MpEvent) SetPropertyScanNotificationsValue(value uint32) (err error) { 
    return instance.SetProperty("ScanNotificationsValue", (value))
}

// GetScanNotificationsValue gets the value of ScanNotificationsValue for the instance
func (instance *MSFT_MpEvent) GetPropertyScanNotificationsValue()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ScanNotificationsValue")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetSignatureNotificationsValue sets the value of SignatureNotificationsValue for the instance
func (instance *MSFT_MpEvent) SetPropertySignatureNotificationsValue(value uint32) (err error) { 
    return instance.SetProperty("SignatureNotificationsValue", (value))
}

// GetSignatureNotificationsValue gets the value of SignatureNotificationsValue for the instance
func (instance *MSFT_MpEvent) GetPropertySignatureNotificationsValue()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SignatureNotificationsValue")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetThreatNotificationsValue sets the value of ThreatNotificationsValue for the instance
func (instance *MSFT_MpEvent) SetPropertyThreatNotificationsValue(value uint32) (err error) { 
    return instance.SetProperty("ThreatNotificationsValue", (value))
}

// GetThreatNotificationsValue gets the value of ThreatNotificationsValue for the instance
func (instance *MSFT_MpEvent) GetPropertyThreatNotificationsValue()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ThreatNotificationsValue")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

