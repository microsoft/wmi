// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSCluster_StorageNode struct
type MSCluster_StorageNode struct { 
	*cim.WmiInstance

	// 
	ConnectionString string

	// 
	Description string

	// 
	Id string

	// 
	Location string

	// 
	ManufacturerId string

	// 
	Name string

	// 
	ProductId string

	// 
	SerialNumber string

	// 
	State uint32

	// 
	StorageNodeHealth uint32

	// 
	StorageNodeOperationalStatus uint32
}

	func NewMSCluster_StorageNodeEx1(instance *cim.WmiInstance) (newInstance *MSCluster_StorageNode, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSCluster_StorageNode {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSCluster_StorageNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSCluster_StorageNode, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSCluster_StorageNode {
	WmiInstance: tmp,
	}
	return
	}
	

// SetConnectionString sets the value of ConnectionString for the instance
func (instance *MSCluster_StorageNode) SetPropertyConnectionString(value string) (err error) { 
    return instance.SetProperty("ConnectionString", (value))
}

// GetConnectionString gets the value of ConnectionString for the instance
func (instance *MSCluster_StorageNode) GetPropertyConnectionString()(value string, err error) { 
    retValue, err := instance.GetProperty("ConnectionString")
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

// SetDescription sets the value of Description for the instance
func (instance *MSCluster_StorageNode) SetPropertyDescription(value string) (err error) { 
    return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSCluster_StorageNode) GetPropertyDescription()(value string, err error) { 
    retValue, err := instance.GetProperty("Description")
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

// SetId sets the value of Id for the instance
func (instance *MSCluster_StorageNode) SetPropertyId(value string) (err error) { 
    return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSCluster_StorageNode) GetPropertyId()(value string, err error) { 
    retValue, err := instance.GetProperty("Id")
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

// SetLocation sets the value of Location for the instance
func (instance *MSCluster_StorageNode) SetPropertyLocation(value string) (err error) { 
    return instance.SetProperty("Location", (value))
}

// GetLocation gets the value of Location for the instance
func (instance *MSCluster_StorageNode) GetPropertyLocation()(value string, err error) { 
    retValue, err := instance.GetProperty("Location")
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

// SetManufacturerId sets the value of ManufacturerId for the instance
func (instance *MSCluster_StorageNode) SetPropertyManufacturerId(value string) (err error) { 
    return instance.SetProperty("ManufacturerId", (value))
}

// GetManufacturerId gets the value of ManufacturerId for the instance
func (instance *MSCluster_StorageNode) GetPropertyManufacturerId()(value string, err error) { 
    retValue, err := instance.GetProperty("ManufacturerId")
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

// SetName sets the value of Name for the instance
func (instance *MSCluster_StorageNode) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSCluster_StorageNode) GetPropertyName()(value string, err error) { 
    retValue, err := instance.GetProperty("Name")
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

// SetProductId sets the value of ProductId for the instance
func (instance *MSCluster_StorageNode) SetPropertyProductId(value string) (err error) { 
    return instance.SetProperty("ProductId", (value))
}

// GetProductId gets the value of ProductId for the instance
func (instance *MSCluster_StorageNode) GetPropertyProductId()(value string, err error) { 
    retValue, err := instance.GetProperty("ProductId")
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

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *MSCluster_StorageNode) SetPropertySerialNumber(value string) (err error) { 
    return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *MSCluster_StorageNode) GetPropertySerialNumber()(value string, err error) { 
    retValue, err := instance.GetProperty("SerialNumber")
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

// SetState sets the value of State for the instance
func (instance *MSCluster_StorageNode) SetPropertyState(value uint32) (err error) { 
    return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *MSCluster_StorageNode) GetPropertyState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("State")
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

// SetStorageNodeHealth sets the value of StorageNodeHealth for the instance
func (instance *MSCluster_StorageNode) SetPropertyStorageNodeHealth(value uint32) (err error) { 
    return instance.SetProperty("StorageNodeHealth", (value))
}

// GetStorageNodeHealth gets the value of StorageNodeHealth for the instance
func (instance *MSCluster_StorageNode) GetPropertyStorageNodeHealth()(value uint32, err error) { 
    retValue, err := instance.GetProperty("StorageNodeHealth")
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

// SetStorageNodeOperationalStatus sets the value of StorageNodeOperationalStatus for the instance
func (instance *MSCluster_StorageNode) SetPropertyStorageNodeOperationalStatus(value uint32) (err error) { 
    return instance.SetProperty("StorageNodeOperationalStatus", (value))
}

// GetStorageNodeOperationalStatus gets the value of StorageNodeOperationalStatus for the instance
func (instance *MSCluster_StorageNode) GetPropertyStorageNodeOperationalStatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("StorageNodeOperationalStatus")
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

// <param name="Description" type="string "></param>
// <param name="Flags" type="uint32 "></param>
// <param name="Location" type="string "></param>
// <param name="Name" type="string "></param>

// <param name="AddedStorageNode" type="MSCluster_StorageNode "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_StorageNode) AddStorageNode( /* IN */ Name string,
 /* IN */ Description string,
 /* IN */ Location string,
 /* IN */ Flags uint32,
 /* OUT */ AddedStorageNode MSCluster_StorageNode) (result uint32, err error) {retVal, err := instance.InvokeMethod("AddStorageNode" , Name, Description, Location, Flags)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Description" type="string "></param>
// <param name="Flags" type="uint32 "></param>
// <param name="Location" type="string "></param>
// <param name="NewName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_StorageNode) SetStorageNode( /* IN */ NewName string,
 /* IN */ Description string,
 /* IN */ Location string,
 /* IN */ Flags uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("SetStorageNode" , NewName, Description, Location, Flags);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="Flags" type="uint32 "></param>

// <param name="Parent" type="MSCluster_FaultDomain "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_StorageNode) GetParent( /* OUT */ Parent MSCluster_FaultDomain,
 /* OPTIONAL IN */ Flags uint32) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetParent" , Flags)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_StorageNode) RemoveStorageNode( /* IN */ Flags uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("RemoveStorageNode" , Flags);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


