// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Server
//////////////////////////////////////////////
package server
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RemoteAccessConfigurationVersion struct
type RemoteAccessConfigurationVersion struct { 
	*cim.WmiInstance

	// 
	AppliedConfigGlobalVersion string

	// 
	AppliedConfigSiteVersion string

	// 
	AppliedConfigTimestamp string

	// 
	ReceivedConfigGlobalVersion string

	// 
	ReceivedConfigSiteVersion string

	// 
	ReceivedConfigTimestamp string
}

	func NewRemoteAccessConfigurationVersionEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessConfigurationVersion, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RemoteAccessConfigurationVersion {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRemoteAccessConfigurationVersionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RemoteAccessConfigurationVersion, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RemoteAccessConfigurationVersion {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAppliedConfigGlobalVersion sets the value of AppliedConfigGlobalVersion for the instance
func (instance *RemoteAccessConfigurationVersion) SetPropertyAppliedConfigGlobalVersion(value string) (err error) { 
    return instance.SetProperty("AppliedConfigGlobalVersion", (value))
}

// GetAppliedConfigGlobalVersion gets the value of AppliedConfigGlobalVersion for the instance
func (instance *RemoteAccessConfigurationVersion) GetPropertyAppliedConfigGlobalVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("AppliedConfigGlobalVersion")
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

// SetAppliedConfigSiteVersion sets the value of AppliedConfigSiteVersion for the instance
func (instance *RemoteAccessConfigurationVersion) SetPropertyAppliedConfigSiteVersion(value string) (err error) { 
    return instance.SetProperty("AppliedConfigSiteVersion", (value))
}

// GetAppliedConfigSiteVersion gets the value of AppliedConfigSiteVersion for the instance
func (instance *RemoteAccessConfigurationVersion) GetPropertyAppliedConfigSiteVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("AppliedConfigSiteVersion")
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

// SetAppliedConfigTimestamp sets the value of AppliedConfigTimestamp for the instance
func (instance *RemoteAccessConfigurationVersion) SetPropertyAppliedConfigTimestamp(value string) (err error) { 
    return instance.SetProperty("AppliedConfigTimestamp", (value))
}

// GetAppliedConfigTimestamp gets the value of AppliedConfigTimestamp for the instance
func (instance *RemoteAccessConfigurationVersion) GetPropertyAppliedConfigTimestamp()(value string, err error) { 
    retValue, err := instance.GetProperty("AppliedConfigTimestamp")
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

// SetReceivedConfigGlobalVersion sets the value of ReceivedConfigGlobalVersion for the instance
func (instance *RemoteAccessConfigurationVersion) SetPropertyReceivedConfigGlobalVersion(value string) (err error) { 
    return instance.SetProperty("ReceivedConfigGlobalVersion", (value))
}

// GetReceivedConfigGlobalVersion gets the value of ReceivedConfigGlobalVersion for the instance
func (instance *RemoteAccessConfigurationVersion) GetPropertyReceivedConfigGlobalVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("ReceivedConfigGlobalVersion")
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

// SetReceivedConfigSiteVersion sets the value of ReceivedConfigSiteVersion for the instance
func (instance *RemoteAccessConfigurationVersion) SetPropertyReceivedConfigSiteVersion(value string) (err error) { 
    return instance.SetProperty("ReceivedConfigSiteVersion", (value))
}

// GetReceivedConfigSiteVersion gets the value of ReceivedConfigSiteVersion for the instance
func (instance *RemoteAccessConfigurationVersion) GetPropertyReceivedConfigSiteVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("ReceivedConfigSiteVersion")
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

// SetReceivedConfigTimestamp sets the value of ReceivedConfigTimestamp for the instance
func (instance *RemoteAccessConfigurationVersion) SetPropertyReceivedConfigTimestamp(value string) (err error) { 
    return instance.SetProperty("ReceivedConfigTimestamp", (value))
}

// GetReceivedConfigTimestamp gets the value of ReceivedConfigTimestamp for the instance
func (instance *RemoteAccessConfigurationVersion) GetPropertyReceivedConfigTimestamp()(value string, err error) { 
    retValue, err := instance.GetProperty("ReceivedConfigTimestamp")
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

