// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.Gateway
//////////////////////////////////////////////
package gateway
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// PS_GatewayRoutingDomain struct
type PS_GatewayRoutingDomain struct { 
	*cim.WmiInstance

	// 
	InterfaceName string

	// 
	RoutingDomain string

	// 
	RoutingDomainID string

	// 
	RoutingDomainStatus string

	// 
	VpnS2SStatus uint32

	// 
	VpnStatus uint32
}

	func NewPS_GatewayRoutingDomainEx1(instance *cim.WmiInstance) (newInstance *PS_GatewayRoutingDomain, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_GatewayRoutingDomain {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_GatewayRoutingDomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_GatewayRoutingDomain, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_GatewayRoutingDomain {
	WmiInstance: tmp,
	}
	return
	}
	

// SetInterfaceName sets the value of InterfaceName for the instance
func (instance *PS_GatewayRoutingDomain) SetPropertyInterfaceName(value string) (err error) { 
    return instance.SetProperty("InterfaceName", (value))
}

// GetInterfaceName gets the value of InterfaceName for the instance
func (instance *PS_GatewayRoutingDomain) GetPropertyInterfaceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InterfaceName")
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

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *PS_GatewayRoutingDomain) SetPropertyRoutingDomain(value string) (err error) { 
    return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *PS_GatewayRoutingDomain) GetPropertyRoutingDomain()(value string, err error) { 
    retValue, err := instance.GetProperty("RoutingDomain")
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

// SetRoutingDomainID sets the value of RoutingDomainID for the instance
func (instance *PS_GatewayRoutingDomain) SetPropertyRoutingDomainID(value string) (err error) { 
    return instance.SetProperty("RoutingDomainID", (value))
}

// GetRoutingDomainID gets the value of RoutingDomainID for the instance
func (instance *PS_GatewayRoutingDomain) GetPropertyRoutingDomainID()(value string, err error) { 
    retValue, err := instance.GetProperty("RoutingDomainID")
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

// SetRoutingDomainStatus sets the value of RoutingDomainStatus for the instance
func (instance *PS_GatewayRoutingDomain) SetPropertyRoutingDomainStatus(value string) (err error) { 
    return instance.SetProperty("RoutingDomainStatus", (value))
}

// GetRoutingDomainStatus gets the value of RoutingDomainStatus for the instance
func (instance *PS_GatewayRoutingDomain) GetPropertyRoutingDomainStatus()(value string, err error) { 
    retValue, err := instance.GetProperty("RoutingDomainStatus")
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

// SetVpnS2SStatus sets the value of VpnS2SStatus for the instance
func (instance *PS_GatewayRoutingDomain) SetPropertyVpnS2SStatus(value uint32) (err error) { 
    return instance.SetProperty("VpnS2SStatus", (value))
}

// GetVpnS2SStatus gets the value of VpnS2SStatus for the instance
func (instance *PS_GatewayRoutingDomain) GetPropertyVpnS2SStatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("VpnS2SStatus")
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

// SetVpnStatus sets the value of VpnStatus for the instance
func (instance *PS_GatewayRoutingDomain) SetPropertyVpnStatus(value uint32) (err error) { 
    return instance.SetProperty("VpnStatus", (value))
}

// GetVpnStatus gets the value of VpnStatus for the instance
func (instance *PS_GatewayRoutingDomain) GetPropertyVpnStatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("VpnStatus")
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

// 

// <param name="Name" type="string "></param>

// <param name="cmdletOutput" type="GatewayRoutingDomainConfiguration []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_GatewayRoutingDomain) Get( /* IN */ Name string,
 /* OUT */ cmdletOutput []GatewayRoutingDomainConfiguration) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" , Name)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="InterfaceName" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Type" type="uint32 "></param>

// <param name="cmdletOutput" type="GatewayRoutingDomainConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_GatewayRoutingDomain) Enable( /* IN */ Name string,
 /* IN */ PassThru bool,
 /* IN */ Type uint32,
 /* IN */ InterfaceName string,
 /* OUT */ cmdletOutput GatewayRoutingDomainConfiguration) (result uint32, err error) {retVal, err := instance.InvokeMethod("Enable" , Name, PassThru, Type, InterfaceName)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Force" type="bool "></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Type" type="uint32 "></param>

// <param name="cmdletOutput" type="GatewayRoutingDomainConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_GatewayRoutingDomain) Disable( /* IN */ Name string,
 /* IN */ Force bool,
 /* IN */ PassThru bool,
 /* IN */ Type uint32,
 /* OUT */ cmdletOutput GatewayRoutingDomainConfiguration) (result uint32, err error) {retVal, err := instance.InvokeMethod("Disable" , Name, Force, PassThru, Type)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


