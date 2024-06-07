// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc struct
type Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc struct { 
	*Win32_PerfFormattedData

	// 
	BytesreceivedbyactiveDirectAccessclients uint64

	// 
	BytesreceivedbydisconnectedDirectAccessclients uint64

	// 
	BytestransmittedbyactiveDirectAccessclients uint64

	// 
	BytestransmittedbydisconnectedDirectAccessclients uint64

	// 
	CumulativenumberofDirectAccessconnectionssinceservicestart uint32

	// 
	Maximumnumberofremoteclientsconnectedtoserversinceservicestart uint32

	// 
	TotalnumberofactiveDirectAccessconnections uint32
}

	func NewWin32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvcEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvcEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetBytesreceivedbyactiveDirectAccessclients sets the value of BytesreceivedbyactiveDirectAccessclients for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) SetPropertyBytesreceivedbyactiveDirectAccessclients(value uint64) (err error) { 
    return instance.SetProperty("BytesreceivedbyactiveDirectAccessclients", (value))
}

// GetBytesreceivedbyactiveDirectAccessclients gets the value of BytesreceivedbyactiveDirectAccessclients for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) GetPropertyBytesreceivedbyactiveDirectAccessclients()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BytesreceivedbyactiveDirectAccessclients")
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

// SetBytesreceivedbydisconnectedDirectAccessclients sets the value of BytesreceivedbydisconnectedDirectAccessclients for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) SetPropertyBytesreceivedbydisconnectedDirectAccessclients(value uint64) (err error) { 
    return instance.SetProperty("BytesreceivedbydisconnectedDirectAccessclients", (value))
}

// GetBytesreceivedbydisconnectedDirectAccessclients gets the value of BytesreceivedbydisconnectedDirectAccessclients for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) GetPropertyBytesreceivedbydisconnectedDirectAccessclients()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BytesreceivedbydisconnectedDirectAccessclients")
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

// SetBytestransmittedbyactiveDirectAccessclients sets the value of BytestransmittedbyactiveDirectAccessclients for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) SetPropertyBytestransmittedbyactiveDirectAccessclients(value uint64) (err error) { 
    return instance.SetProperty("BytestransmittedbyactiveDirectAccessclients", (value))
}

// GetBytestransmittedbyactiveDirectAccessclients gets the value of BytestransmittedbyactiveDirectAccessclients for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) GetPropertyBytestransmittedbyactiveDirectAccessclients()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BytestransmittedbyactiveDirectAccessclients")
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

// SetBytestransmittedbydisconnectedDirectAccessclients sets the value of BytestransmittedbydisconnectedDirectAccessclients for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) SetPropertyBytestransmittedbydisconnectedDirectAccessclients(value uint64) (err error) { 
    return instance.SetProperty("BytestransmittedbydisconnectedDirectAccessclients", (value))
}

// GetBytestransmittedbydisconnectedDirectAccessclients gets the value of BytestransmittedbydisconnectedDirectAccessclients for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) GetPropertyBytestransmittedbydisconnectedDirectAccessclients()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BytestransmittedbydisconnectedDirectAccessclients")
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

// SetCumulativenumberofDirectAccessconnectionssinceservicestart sets the value of CumulativenumberofDirectAccessconnectionssinceservicestart for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) SetPropertyCumulativenumberofDirectAccessconnectionssinceservicestart(value uint32) (err error) { 
    return instance.SetProperty("CumulativenumberofDirectAccessconnectionssinceservicestart", (value))
}

// GetCumulativenumberofDirectAccessconnectionssinceservicestart gets the value of CumulativenumberofDirectAccessconnectionssinceservicestart for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) GetPropertyCumulativenumberofDirectAccessconnectionssinceservicestart()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CumulativenumberofDirectAccessconnectionssinceservicestart")
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

// SetMaximumnumberofremoteclientsconnectedtoserversinceservicestart sets the value of Maximumnumberofremoteclientsconnectedtoserversinceservicestart for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) SetPropertyMaximumnumberofremoteclientsconnectedtoserversinceservicestart(value uint32) (err error) { 
    return instance.SetProperty("Maximumnumberofremoteclientsconnectedtoserversinceservicestart", (value))
}

// GetMaximumnumberofremoteclientsconnectedtoserversinceservicestart gets the value of Maximumnumberofremoteclientsconnectedtoserversinceservicestart for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) GetPropertyMaximumnumberofremoteclientsconnectedtoserversinceservicestart()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Maximumnumberofremoteclientsconnectedtoserversinceservicestart")
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

// SetTotalnumberofactiveDirectAccessconnections sets the value of TotalnumberofactiveDirectAccessconnections for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) SetPropertyTotalnumberofactiveDirectAccessconnections(value uint32) (err error) { 
    return instance.SetProperty("TotalnumberofactiveDirectAccessconnections", (value))
}

// GetTotalnumberofactiveDirectAccessconnections gets the value of TotalnumberofactiveDirectAccessconnections for the instance
func (instance *Win32_PerfFormattedData_RamgmtSvcCounterProvider_RaMgmtSvc) GetPropertyTotalnumberofactiveDirectAccessconnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalnumberofactiveDirectAccessconnections")
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

