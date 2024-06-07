// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MDM_Policy_User_Result01_EnterpriseCloudPrint02 struct
type MDM_Policy_User_Result01_EnterpriseCloudPrint02 struct { 
	*cim.WmiInstance

	// 
	CloudPrinterDiscoveryEndPoint string

	// 
	CloudPrintOAuthAuthority string

	// 
	CloudPrintOAuthClientId string

	// 
	CloudPrintResourceId string

	// 
	DiscoveryMaxPrinterLimit int32

	// 
	InstanceID string

	// 
	MopriaDiscoveryResourceId string

	// 
	ParentID string
}

	func NewMDM_Policy_User_Result01_EnterpriseCloudPrint02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_User_Result01_EnterpriseCloudPrint02, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_Policy_User_Result01_EnterpriseCloudPrint02 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_Policy_User_Result01_EnterpriseCloudPrint02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_Policy_User_Result01_EnterpriseCloudPrint02, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_Policy_User_Result01_EnterpriseCloudPrint02 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetCloudPrinterDiscoveryEndPoint sets the value of CloudPrinterDiscoveryEndPoint for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) SetPropertyCloudPrinterDiscoveryEndPoint(value string) (err error) { 
    return instance.SetProperty("CloudPrinterDiscoveryEndPoint", (value))
}

// GetCloudPrinterDiscoveryEndPoint gets the value of CloudPrinterDiscoveryEndPoint for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) GetPropertyCloudPrinterDiscoveryEndPoint()(value string, err error) { 
    retValue, err := instance.GetProperty("CloudPrinterDiscoveryEndPoint")
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

// SetCloudPrintOAuthAuthority sets the value of CloudPrintOAuthAuthority for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) SetPropertyCloudPrintOAuthAuthority(value string) (err error) { 
    return instance.SetProperty("CloudPrintOAuthAuthority", (value))
}

// GetCloudPrintOAuthAuthority gets the value of CloudPrintOAuthAuthority for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) GetPropertyCloudPrintOAuthAuthority()(value string, err error) { 
    retValue, err := instance.GetProperty("CloudPrintOAuthAuthority")
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

// SetCloudPrintOAuthClientId sets the value of CloudPrintOAuthClientId for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) SetPropertyCloudPrintOAuthClientId(value string) (err error) { 
    return instance.SetProperty("CloudPrintOAuthClientId", (value))
}

// GetCloudPrintOAuthClientId gets the value of CloudPrintOAuthClientId for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) GetPropertyCloudPrintOAuthClientId()(value string, err error) { 
    retValue, err := instance.GetProperty("CloudPrintOAuthClientId")
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

// SetCloudPrintResourceId sets the value of CloudPrintResourceId for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) SetPropertyCloudPrintResourceId(value string) (err error) { 
    return instance.SetProperty("CloudPrintResourceId", (value))
}

// GetCloudPrintResourceId gets the value of CloudPrintResourceId for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) GetPropertyCloudPrintResourceId()(value string, err error) { 
    retValue, err := instance.GetProperty("CloudPrintResourceId")
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

// SetDiscoveryMaxPrinterLimit sets the value of DiscoveryMaxPrinterLimit for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) SetPropertyDiscoveryMaxPrinterLimit(value int32) (err error) { 
    return instance.SetProperty("DiscoveryMaxPrinterLimit", (value))
}

// GetDiscoveryMaxPrinterLimit gets the value of DiscoveryMaxPrinterLimit for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) GetPropertyDiscoveryMaxPrinterLimit()(value int32, err error) { 
    retValue, err := instance.GetProperty("DiscoveryMaxPrinterLimit")
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

    value = int32(valuetmp)

    return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) GetPropertyInstanceID()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceID")
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

// SetMopriaDiscoveryResourceId sets the value of MopriaDiscoveryResourceId for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) SetPropertyMopriaDiscoveryResourceId(value string) (err error) { 
    return instance.SetProperty("MopriaDiscoveryResourceId", (value))
}

// GetMopriaDiscoveryResourceId gets the value of MopriaDiscoveryResourceId for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) GetPropertyMopriaDiscoveryResourceId()(value string, err error) { 
    retValue, err := instance.GetProperty("MopriaDiscoveryResourceId")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Result01_EnterpriseCloudPrint02) GetPropertyParentID()(value string, err error) { 
    retValue, err := instance.GetProperty("ParentID")
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

