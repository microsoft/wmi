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

// MSiSCSI_InitiatorLoginStatistics struct
type MSiSCSI_InitiatorLoginStatistics struct { 
	*Win32_PerfRawData

	// 
	Active bool

	// 
	InstanceName string

	// 
	LoginAcceptRsps uint32

	// 
	LoginAuthenticateFails uint32

	// 
	LoginAuthFailRsps uint32

	// 
	LoginFailures uint32

	// 
	LoginNegotiateFails uint32

	// 
	LoginOtherFailRsps uint32

	// 
	LoginRedirectRsps uint32

	// 
	LogoutNormals uint32

	// 
	LogoutOtherCodes uint32

	// 
	UniqueAdapterId uint64
}

	func NewMSiSCSI_InitiatorLoginStatisticsEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_InitiatorLoginStatistics, err error) {tmp, err := NewWin32_PerfRawDataEx1(instance)
		
	if err != nil { return }
	newInstance = &MSiSCSI_InitiatorLoginStatistics {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

	func NewMSiSCSI_InitiatorLoginStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSiSCSI_InitiatorLoginStatistics, err error) {tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSiSCSI_InitiatorLoginStatistics {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) GetPropertyActive()(value bool, err error) { 
    retValue, err := instance.GetProperty("Active")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) GetPropertyInstanceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceName")
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

// SetLoginAcceptRsps sets the value of LoginAcceptRsps for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) SetPropertyLoginAcceptRsps(value uint32) (err error) { 
    return instance.SetProperty("LoginAcceptRsps", (value))
}

// GetLoginAcceptRsps gets the value of LoginAcceptRsps for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) GetPropertyLoginAcceptRsps()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LoginAcceptRsps")
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

// SetLoginAuthenticateFails sets the value of LoginAuthenticateFails for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) SetPropertyLoginAuthenticateFails(value uint32) (err error) { 
    return instance.SetProperty("LoginAuthenticateFails", (value))
}

// GetLoginAuthenticateFails gets the value of LoginAuthenticateFails for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) GetPropertyLoginAuthenticateFails()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LoginAuthenticateFails")
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

// SetLoginAuthFailRsps sets the value of LoginAuthFailRsps for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) SetPropertyLoginAuthFailRsps(value uint32) (err error) { 
    return instance.SetProperty("LoginAuthFailRsps", (value))
}

// GetLoginAuthFailRsps gets the value of LoginAuthFailRsps for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) GetPropertyLoginAuthFailRsps()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LoginAuthFailRsps")
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

// SetLoginFailures sets the value of LoginFailures for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) SetPropertyLoginFailures(value uint32) (err error) { 
    return instance.SetProperty("LoginFailures", (value))
}

// GetLoginFailures gets the value of LoginFailures for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) GetPropertyLoginFailures()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LoginFailures")
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

// SetLoginNegotiateFails sets the value of LoginNegotiateFails for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) SetPropertyLoginNegotiateFails(value uint32) (err error) { 
    return instance.SetProperty("LoginNegotiateFails", (value))
}

// GetLoginNegotiateFails gets the value of LoginNegotiateFails for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) GetPropertyLoginNegotiateFails()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LoginNegotiateFails")
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

// SetLoginOtherFailRsps sets the value of LoginOtherFailRsps for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) SetPropertyLoginOtherFailRsps(value uint32) (err error) { 
    return instance.SetProperty("LoginOtherFailRsps", (value))
}

// GetLoginOtherFailRsps gets the value of LoginOtherFailRsps for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) GetPropertyLoginOtherFailRsps()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LoginOtherFailRsps")
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

// SetLoginRedirectRsps sets the value of LoginRedirectRsps for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) SetPropertyLoginRedirectRsps(value uint32) (err error) { 
    return instance.SetProperty("LoginRedirectRsps", (value))
}

// GetLoginRedirectRsps gets the value of LoginRedirectRsps for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) GetPropertyLoginRedirectRsps()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LoginRedirectRsps")
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

// SetLogoutNormals sets the value of LogoutNormals for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) SetPropertyLogoutNormals(value uint32) (err error) { 
    return instance.SetProperty("LogoutNormals", (value))
}

// GetLogoutNormals gets the value of LogoutNormals for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) GetPropertyLogoutNormals()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LogoutNormals")
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

// SetLogoutOtherCodes sets the value of LogoutOtherCodes for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) SetPropertyLogoutOtherCodes(value uint32) (err error) { 
    return instance.SetProperty("LogoutOtherCodes", (value))
}

// GetLogoutOtherCodes gets the value of LogoutOtherCodes for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) GetPropertyLogoutOtherCodes()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LogoutOtherCodes")
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

// SetUniqueAdapterId sets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) SetPropertyUniqueAdapterId(value uint64) (err error) { 
    return instance.SetProperty("UniqueAdapterId", (value))
}

// GetUniqueAdapterId gets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_InitiatorLoginStatistics) GetPropertyUniqueAdapterId()(value uint64, err error) { 
    retValue, err := instance.GetProperty("UniqueAdapterId")
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

