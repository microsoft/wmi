// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// HNet_BackupIpConfiguration struct
type HNet_BackupIpConfiguration struct { 
	*cim.WmiInstance

	// 
	Connection HNet_Connection

	// 
	DefaultGateway string

	// 
	EnableDHCP uint32

	// 
	IPAddress string

	// 
	SubnetMask string
}

	func NewHNet_BackupIpConfigurationEx1(instance *cim.WmiInstance) (newInstance *HNet_BackupIpConfiguration, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &HNet_BackupIpConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewHNet_BackupIpConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HNet_BackupIpConfiguration, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HNet_BackupIpConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

// SetConnection sets the value of Connection for the instance
func (instance *HNet_BackupIpConfiguration) SetPropertyConnection(value HNet_Connection) (err error) { 
    return instance.SetProperty("Connection", (value))
}

// GetConnection gets the value of Connection for the instance
func (instance *HNet_BackupIpConfiguration) GetPropertyConnection()(value HNet_Connection, err error) { 
    retValue, err := instance.GetProperty("Connection")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(HNet_Connection); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " HNet_Connection is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = HNet_Connection(valuetmp)

    return
}

// SetDefaultGateway sets the value of DefaultGateway for the instance
func (instance *HNet_BackupIpConfiguration) SetPropertyDefaultGateway(value string) (err error) { 
    return instance.SetProperty("DefaultGateway", (value))
}

// GetDefaultGateway gets the value of DefaultGateway for the instance
func (instance *HNet_BackupIpConfiguration) GetPropertyDefaultGateway()(value string, err error) { 
    retValue, err := instance.GetProperty("DefaultGateway")
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

// SetEnableDHCP sets the value of EnableDHCP for the instance
func (instance *HNet_BackupIpConfiguration) SetPropertyEnableDHCP(value uint32) (err error) { 
    return instance.SetProperty("EnableDHCP", (value))
}

// GetEnableDHCP gets the value of EnableDHCP for the instance
func (instance *HNet_BackupIpConfiguration) GetPropertyEnableDHCP()(value uint32, err error) { 
    retValue, err := instance.GetProperty("EnableDHCP")
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

// SetIPAddress sets the value of IPAddress for the instance
func (instance *HNet_BackupIpConfiguration) SetPropertyIPAddress(value string) (err error) { 
    return instance.SetProperty("IPAddress", (value))
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *HNet_BackupIpConfiguration) GetPropertyIPAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("IPAddress")
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

// SetSubnetMask sets the value of SubnetMask for the instance
func (instance *HNet_BackupIpConfiguration) SetPropertySubnetMask(value string) (err error) { 
    return instance.SetProperty("SubnetMask", (value))
}

// GetSubnetMask gets the value of SubnetMask for the instance
func (instance *HNet_BackupIpConfiguration) GetPropertySubnetMask()(value string, err error) { 
    retValue, err := instance.GetProperty("SubnetMask")
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

