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

// MSCluster_GroupSet struct
type MSCluster_GroupSet struct { 
	*cim.WmiInstance

	// 
	ClusterNodeObjectReturnedFrom string

	// 
	FaultDomains uint32

	// 
	GroupNames []string

	// 
	IsAvailabilitySet bool

	// 
	IsGlobal bool

	// 
	Name string

	// 
	NodeDomainInfo []string

	// 
	ProviderNames []string

	// 
	ReserveSpareNode bool

	// 
	StartupCount uint32

	// 
	StartupDelay uint32

	// 
	StartupDelayTrigger uint32

	// 
	StatusInformation uint64

	// 
	UpdateDomains uint32
}

	func NewMSCluster_GroupSetEx1(instance *cim.WmiInstance) (newInstance *MSCluster_GroupSet, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSCluster_GroupSet {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSCluster_GroupSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSCluster_GroupSet, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSCluster_GroupSet {
	WmiInstance: tmp,
	}
	return
	}
	

// SetClusterNodeObjectReturnedFrom sets the value of ClusterNodeObjectReturnedFrom for the instance
func (instance *MSCluster_GroupSet) SetPropertyClusterNodeObjectReturnedFrom(value string) (err error) { 
    return instance.SetProperty("ClusterNodeObjectReturnedFrom", (value))
}

// GetClusterNodeObjectReturnedFrom gets the value of ClusterNodeObjectReturnedFrom for the instance
func (instance *MSCluster_GroupSet) GetPropertyClusterNodeObjectReturnedFrom()(value string, err error) { 
    retValue, err := instance.GetProperty("ClusterNodeObjectReturnedFrom")
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

// SetFaultDomains sets the value of FaultDomains for the instance
func (instance *MSCluster_GroupSet) SetPropertyFaultDomains(value uint32) (err error) { 
    return instance.SetProperty("FaultDomains", (value))
}

// GetFaultDomains gets the value of FaultDomains for the instance
func (instance *MSCluster_GroupSet) GetPropertyFaultDomains()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FaultDomains")
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

// SetGroupNames sets the value of GroupNames for the instance
func (instance *MSCluster_GroupSet) SetPropertyGroupNames(value []string) (err error) { 
    return instance.SetProperty("GroupNames", (value))
}

// GetGroupNames gets the value of GroupNames for the instance
func (instance *MSCluster_GroupSet) GetPropertyGroupNames()(value []string, err error) { 
    retValue, err := instance.GetProperty("GroupNames")
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

// SetIsAvailabilitySet sets the value of IsAvailabilitySet for the instance
func (instance *MSCluster_GroupSet) SetPropertyIsAvailabilitySet(value bool) (err error) { 
    return instance.SetProperty("IsAvailabilitySet", (value))
}

// GetIsAvailabilitySet gets the value of IsAvailabilitySet for the instance
func (instance *MSCluster_GroupSet) GetPropertyIsAvailabilitySet()(value bool, err error) { 
    retValue, err := instance.GetProperty("IsAvailabilitySet")
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

// SetIsGlobal sets the value of IsGlobal for the instance
func (instance *MSCluster_GroupSet) SetPropertyIsGlobal(value bool) (err error) { 
    return instance.SetProperty("IsGlobal", (value))
}

// GetIsGlobal gets the value of IsGlobal for the instance
func (instance *MSCluster_GroupSet) GetPropertyIsGlobal()(value bool, err error) { 
    retValue, err := instance.GetProperty("IsGlobal")
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

// SetName sets the value of Name for the instance
func (instance *MSCluster_GroupSet) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSCluster_GroupSet) GetPropertyName()(value string, err error) { 
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

// SetNodeDomainInfo sets the value of NodeDomainInfo for the instance
func (instance *MSCluster_GroupSet) SetPropertyNodeDomainInfo(value []string) (err error) { 
    return instance.SetProperty("NodeDomainInfo", (value))
}

// GetNodeDomainInfo gets the value of NodeDomainInfo for the instance
func (instance *MSCluster_GroupSet) GetPropertyNodeDomainInfo()(value []string, err error) { 
    retValue, err := instance.GetProperty("NodeDomainInfo")
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

// SetProviderNames sets the value of ProviderNames for the instance
func (instance *MSCluster_GroupSet) SetPropertyProviderNames(value []string) (err error) { 
    return instance.SetProperty("ProviderNames", (value))
}

// GetProviderNames gets the value of ProviderNames for the instance
func (instance *MSCluster_GroupSet) GetPropertyProviderNames()(value []string, err error) { 
    retValue, err := instance.GetProperty("ProviderNames")
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

// SetReserveSpareNode sets the value of ReserveSpareNode for the instance
func (instance *MSCluster_GroupSet) SetPropertyReserveSpareNode(value bool) (err error) { 
    return instance.SetProperty("ReserveSpareNode", (value))
}

// GetReserveSpareNode gets the value of ReserveSpareNode for the instance
func (instance *MSCluster_GroupSet) GetPropertyReserveSpareNode()(value bool, err error) { 
    retValue, err := instance.GetProperty("ReserveSpareNode")
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

// SetStartupCount sets the value of StartupCount for the instance
func (instance *MSCluster_GroupSet) SetPropertyStartupCount(value uint32) (err error) { 
    return instance.SetProperty("StartupCount", (value))
}

// GetStartupCount gets the value of StartupCount for the instance
func (instance *MSCluster_GroupSet) GetPropertyStartupCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("StartupCount")
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

// SetStartupDelay sets the value of StartupDelay for the instance
func (instance *MSCluster_GroupSet) SetPropertyStartupDelay(value uint32) (err error) { 
    return instance.SetProperty("StartupDelay", (value))
}

// GetStartupDelay gets the value of StartupDelay for the instance
func (instance *MSCluster_GroupSet) GetPropertyStartupDelay()(value uint32, err error) { 
    retValue, err := instance.GetProperty("StartupDelay")
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

// SetStartupDelayTrigger sets the value of StartupDelayTrigger for the instance
func (instance *MSCluster_GroupSet) SetPropertyStartupDelayTrigger(value uint32) (err error) { 
    return instance.SetProperty("StartupDelayTrigger", (value))
}

// GetStartupDelayTrigger gets the value of StartupDelayTrigger for the instance
func (instance *MSCluster_GroupSet) GetPropertyStartupDelayTrigger()(value uint32, err error) { 
    retValue, err := instance.GetProperty("StartupDelayTrigger")
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

// SetStatusInformation sets the value of StatusInformation for the instance
func (instance *MSCluster_GroupSet) SetPropertyStatusInformation(value uint64) (err error) { 
    return instance.SetProperty("StatusInformation", (value))
}

// GetStatusInformation gets the value of StatusInformation for the instance
func (instance *MSCluster_GroupSet) GetPropertyStatusInformation()(value uint64, err error) { 
    retValue, err := instance.GetProperty("StatusInformation")
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

// SetUpdateDomains sets the value of UpdateDomains for the instance
func (instance *MSCluster_GroupSet) SetPropertyUpdateDomains(value uint32) (err error) { 
    return instance.SetProperty("UpdateDomains", (value))
}

// GetUpdateDomains gets the value of UpdateDomains for the instance
func (instance *MSCluster_GroupSet) GetPropertyUpdateDomains()(value uint32, err error) { 
    retValue, err := instance.GetProperty("UpdateDomains")
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

// <param name="Group" type="string []"></param>
// <param name="Name" type="string "></param>
// <param name="provider" type="string []"></param>

// <param name="CreatedSet" type="MSCluster_GroupSet "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_GroupSet) CreateSet( /* IN */ Name string,
 /* IN */ Group []string,
 /* IN */ provider []string,
 /* OUT */ CreatedSet MSCluster_GroupSet) (result uint32, err error) {retVal, err := instance.InvokeMethod("CreateSet" , Name, Group, provider)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="FaultDomains" type="uint32 "></param>
// <param name="Group" type="string []"></param>
// <param name="Name" type="string "></param>
// <param name="ReserveSpareNode" type="bool "></param>
// <param name="UpdateDomains" type="uint32 "></param>

// <param name="CreatedSet" type="MSCluster_GroupSet "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_GroupSet) CreateAvailabilitySet( /* IN */ Name string,
 /* IN */ Group []string,
 /* IN */ UpdateDomains uint32,
 /* IN */ FaultDomains uint32,
 /* IN */ ReserveSpareNode bool,
 /* OUT */ CreatedSet MSCluster_GroupSet) (result uint32, err error) {retVal, err := instance.InvokeMethod("CreateAvailabilitySet" , Name, Group, UpdateDomains, FaultDomains, ReserveSpareNode)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ContainedGroup" type="string "></param>
// <param name="DependentGroup" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="provider" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Sets" type="MSCluster_GroupSet []"></param>
func (instance *MSCluster_GroupSet) GetSetFrom( /* IN */ ContainedGroup string,
 /* IN */ Name string,
 /* IN */ provider string,
 /* IN */ DependentGroup string,
 /* OUT */ Sets []MSCluster_GroupSet) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetSetFrom" , ContainedGroup, Name, provider, DependentGroup)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="provider" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_GroupSet) AddSetProvider( /* IN */ provider string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("AddSetProvider" , provider);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="provider" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_GroupSet) RemoveSetProvider( /* IN */ provider string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("RemoveSetProvider" , provider);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="Group" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_GroupSet) AddGroupToSet( /* IN */ Group string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("AddGroupToSet" , Group);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="FaultDomain" type="uint32 "></param>
// <param name="Group" type="string "></param>
// <param name="Reserved" type="uint64 "></param>
// <param name="UpdateDomain" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_GroupSet) AddGroupToSetEx( /* IN */ Group string,
 /* IN */ FaultDomain uint32,
 /* IN */ UpdateDomain uint32,
 /* IN */ Reserved uint64) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("AddGroupToSetEx" , Group, FaultDomain, UpdateDomain, Reserved);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="Group" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_GroupSet) RemoveGroupFromSet( /* IN */ Group string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("RemoveGroupFromSet" , Group);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_GroupSet) Remove() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Remove" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="IsGlobal" type="bool "></param>
// <param name="StartupCount" type="uint32 "></param>
// <param name="StartupDelay" type="uint32 "></param>
// <param name="StartupDelayTrigger" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_GroupSet) SetSet( /* IN */ StartupDelayTrigger uint32,
 /* IN */ StartupCount uint32,
 /* IN */ IsGlobal bool,
 /* IN */ StartupDelay uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("SetSet" , StartupDelayTrigger, StartupCount, IsGlobal, StartupDelay);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="DependentGroup" type="string "></param>
// <param name="ProviderGroup" type="string "></param>

// <param name="Groups" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_GroupSet) GetGroups( /* IN */ DependentGroup string,
 /* IN */ ProviderGroup string,
 /* OUT */ Groups []string) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetGroups" , DependentGroup, ProviderGroup)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Group" type="string "></param>
// <param name="provider" type="string "></param>
// <param name="ProviderGroup" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_GroupSet) RemoveProviderForGroup( /* IN */ Group string,
 /* IN */ ProviderGroup string,
 /* IN */ provider string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("RemoveProviderForGroup" , Group, ProviderGroup, provider);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="Group" type="string "></param>
// <param name="provider" type="string "></param>
// <param name="ProviderGroup" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_GroupSet) AddProviderForGroup( /* IN */ Group string,
 /* IN */ ProviderGroup string,
 /* IN */ provider string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("AddProviderForGroup" , Group, ProviderGroup, provider);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


