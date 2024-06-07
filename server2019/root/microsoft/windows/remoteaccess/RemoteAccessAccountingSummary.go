// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RemoteAccessAccountingSummary struct
type RemoteAccessAccountingSummary struct { 
	*RemoteAccessConnectionSummary

	// 
	AverageSessionsPerDay uint32

	// 
	MaxConcurrentSessions uint32

	// 
	TotalDASessions uint32

	// 
	TotalSessions uint32

	// 
	TotalUniqueDAClients uint32

	// 
	TotalVpnSessions uint32
}

	func NewRemoteAccessAccountingSummaryEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessAccountingSummary, err error) {tmp, err := NewRemoteAccessConnectionSummaryEx1(instance)
		
	if err != nil { return }
	newInstance = &RemoteAccessAccountingSummary {
	RemoteAccessConnectionSummary: tmp,
	}
	return
	}
	

	func NewRemoteAccessAccountingSummaryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RemoteAccessAccountingSummary, err error) {tmp, err := NewRemoteAccessConnectionSummaryEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RemoteAccessAccountingSummary {
	RemoteAccessConnectionSummary: tmp,
	}
	return
	}
	

// SetAverageSessionsPerDay sets the value of AverageSessionsPerDay for the instance
func (instance *RemoteAccessAccountingSummary) SetPropertyAverageSessionsPerDay(value uint32) (err error) { 
    return instance.SetProperty("AverageSessionsPerDay", (value))
}

// GetAverageSessionsPerDay gets the value of AverageSessionsPerDay for the instance
func (instance *RemoteAccessAccountingSummary) GetPropertyAverageSessionsPerDay()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AverageSessionsPerDay")
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

// SetMaxConcurrentSessions sets the value of MaxConcurrentSessions for the instance
func (instance *RemoteAccessAccountingSummary) SetPropertyMaxConcurrentSessions(value uint32) (err error) { 
    return instance.SetProperty("MaxConcurrentSessions", (value))
}

// GetMaxConcurrentSessions gets the value of MaxConcurrentSessions for the instance
func (instance *RemoteAccessAccountingSummary) GetPropertyMaxConcurrentSessions()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxConcurrentSessions")
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

// SetTotalDASessions sets the value of TotalDASessions for the instance
func (instance *RemoteAccessAccountingSummary) SetPropertyTotalDASessions(value uint32) (err error) { 
    return instance.SetProperty("TotalDASessions", (value))
}

// GetTotalDASessions gets the value of TotalDASessions for the instance
func (instance *RemoteAccessAccountingSummary) GetPropertyTotalDASessions()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalDASessions")
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

// SetTotalSessions sets the value of TotalSessions for the instance
func (instance *RemoteAccessAccountingSummary) SetPropertyTotalSessions(value uint32) (err error) { 
    return instance.SetProperty("TotalSessions", (value))
}

// GetTotalSessions gets the value of TotalSessions for the instance
func (instance *RemoteAccessAccountingSummary) GetPropertyTotalSessions()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalSessions")
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

// SetTotalUniqueDAClients sets the value of TotalUniqueDAClients for the instance
func (instance *RemoteAccessAccountingSummary) SetPropertyTotalUniqueDAClients(value uint32) (err error) { 
    return instance.SetProperty("TotalUniqueDAClients", (value))
}

// GetTotalUniqueDAClients gets the value of TotalUniqueDAClients for the instance
func (instance *RemoteAccessAccountingSummary) GetPropertyTotalUniqueDAClients()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalUniqueDAClients")
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

// SetTotalVpnSessions sets the value of TotalVpnSessions for the instance
func (instance *RemoteAccessAccountingSummary) SetPropertyTotalVpnSessions(value uint32) (err error) { 
    return instance.SetProperty("TotalVpnSessions", (value))
}

// GetTotalVpnSessions gets the value of TotalVpnSessions for the instance
func (instance *RemoteAccessAccountingSummary) GetPropertyTotalVpnSessions()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalVpnSessions")
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

