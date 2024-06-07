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

// DACorpPrefixConfiguration struct
type DACorpPrefixConfiguration struct { 
	*cim.WmiInstance

	// 
	basePrefix string

	// 
	corpIPv6Prefix []string

	// 
	IPv6State uint32
}

	func NewDACorpPrefixConfigurationEx1(instance *cim.WmiInstance) (newInstance *DACorpPrefixConfiguration, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &DACorpPrefixConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewDACorpPrefixConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DACorpPrefixConfiguration, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DACorpPrefixConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

// SetbasePrefix sets the value of basePrefix for the instance
func (instance *DACorpPrefixConfiguration) SetPropertybasePrefix(value string) (err error) { 
    return instance.SetProperty("basePrefix", (value))
}

// GetbasePrefix gets the value of basePrefix for the instance
func (instance *DACorpPrefixConfiguration) GetPropertybasePrefix()(value string, err error) { 
    retValue, err := instance.GetProperty("basePrefix")
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

// SetcorpIPv6Prefix sets the value of corpIPv6Prefix for the instance
func (instance *DACorpPrefixConfiguration) SetPropertycorpIPv6Prefix(value []string) (err error) { 
    return instance.SetProperty("corpIPv6Prefix", (value))
}

// GetcorpIPv6Prefix gets the value of corpIPv6Prefix for the instance
func (instance *DACorpPrefixConfiguration) GetPropertycorpIPv6Prefix()(value []string, err error) { 
    retValue, err := instance.GetProperty("corpIPv6Prefix")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetIPv6State sets the value of IPv6State for the instance
func (instance *DACorpPrefixConfiguration) SetPropertyIPv6State(value uint32) (err error) { 
    return instance.SetProperty("IPv6State", (value))
}

// GetIPv6State gets the value of IPv6State for the instance
func (instance *DACorpPrefixConfiguration) GetPropertyIPv6State()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IPv6State")
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

