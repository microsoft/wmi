// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_DNSClient struct
type MSFT_DNSClient struct { 
	*CIM_DNSProtocolEndpoint

	// 657
	ConnectionSpecificSuffix string

	// 658
	ConnectionSpecificSuffixSearchList []string

	// 656
	InterfaceAlias string

	// 655
	InterfaceIndex uint32

	// 659
	RegisterThisConnectionsAddress bool

	// 660
	UseSuffixWhenRegistering bool
}

	func NewMSFT_DNSClientEx1(instance *cim.WmiInstance) (newInstance *MSFT_DNSClient, err error) {tmp, err := NewCIM_DNSProtocolEndpointEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_DNSClient {
	CIM_DNSProtocolEndpoint: tmp,
	}
	return
	}
	

	func NewMSFT_DNSClientEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_DNSClient, err error) {tmp, err := NewCIM_DNSProtocolEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_DNSClient {
	CIM_DNSProtocolEndpoint: tmp,
	}
	return
	}
	

// SetConnectionSpecificSuffix sets the value of ConnectionSpecificSuffix for the instance
func (instance *MSFT_DNSClient) SetPropertyConnectionSpecificSuffix(value string) (err error) { 
    return instance.SetProperty("ConnectionSpecificSuffix", (value))
}

// GetConnectionSpecificSuffix gets the value of ConnectionSpecificSuffix for the instance
func (instance *MSFT_DNSClient) GetPropertyConnectionSpecificSuffix()(value string, err error) { 
    retValue, err := instance.GetProperty("ConnectionSpecificSuffix")
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

// SetConnectionSpecificSuffixSearchList sets the value of ConnectionSpecificSuffixSearchList for the instance
func (instance *MSFT_DNSClient) SetPropertyConnectionSpecificSuffixSearchList(value []string) (err error) { 
    return instance.SetProperty("ConnectionSpecificSuffixSearchList", (value))
}

// GetConnectionSpecificSuffixSearchList gets the value of ConnectionSpecificSuffixSearchList for the instance
func (instance *MSFT_DNSClient) GetPropertyConnectionSpecificSuffixSearchList()(value []string, err error) { 
    retValue, err := instance.GetProperty("ConnectionSpecificSuffixSearchList")
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

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_DNSClient) SetPropertyInterfaceAlias(value string) (err error) { 
    return instance.SetProperty("InterfaceAlias", (value))
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_DNSClient) GetPropertyInterfaceAlias()(value string, err error) { 
    retValue, err := instance.GetProperty("InterfaceAlias")
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

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_DNSClient) SetPropertyInterfaceIndex(value uint32) (err error) { 
    return instance.SetProperty("InterfaceIndex", (value))
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_DNSClient) GetPropertyInterfaceIndex()(value uint32, err error) { 
    retValue, err := instance.GetProperty("InterfaceIndex")
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

// SetRegisterThisConnectionsAddress sets the value of RegisterThisConnectionsAddress for the instance
func (instance *MSFT_DNSClient) SetPropertyRegisterThisConnectionsAddress(value bool) (err error) { 
    return instance.SetProperty("RegisterThisConnectionsAddress", (value))
}

// GetRegisterThisConnectionsAddress gets the value of RegisterThisConnectionsAddress for the instance
func (instance *MSFT_DNSClient) GetPropertyRegisterThisConnectionsAddress()(value bool, err error) { 
    retValue, err := instance.GetProperty("RegisterThisConnectionsAddress")
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

// SetUseSuffixWhenRegistering sets the value of UseSuffixWhenRegistering for the instance
func (instance *MSFT_DNSClient) SetPropertyUseSuffixWhenRegistering(value bool) (err error) { 
    return instance.SetProperty("UseSuffixWhenRegistering", (value))
}

// GetUseSuffixWhenRegistering gets the value of UseSuffixWhenRegistering for the instance
func (instance *MSFT_DNSClient) GetPropertyUseSuffixWhenRegistering()(value bool, err error) { 
    retValue, err := instance.GetProperty("UseSuffixWhenRegistering")
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

// 661

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DNSClient) Register() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Register" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


