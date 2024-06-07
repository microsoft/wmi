// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ads_ntfrsreplicaset struct
type ads_ntfrsreplicaset struct { 
	*ds_top

	// 
	DS_fRSDirectoryFilter string

	// 
	DS_fRSDSPoll int32

	// 
	DS_fRSExtensions Uint8Array

	// 
	DS_fRSFileFilter string

	// 
	DS_fRSFlags int32

	// 
	DS_fRSLevelLimit int32

	// 
	DS_fRSPartnerAuthLevel int32

	// 
	DS_fRSPrimaryMember string

	// 
	DS_fRSReplicaSetGUID Uint8Array

	// 
	DS_fRSReplicaSetType int32

	// 
	DS_fRSRootSecurity Uint8Array

	// 
	DS_fRSServiceCommand string

	// 
	DS_fRSVersionGUID Uint8Array

	// 
	DS_managedBy string

	// 
	DS_msFRS_Hub_Member string

	// 
	DS_msFRS_Topology_Pref string

	// 
	DS_schedule Uint8Array
}

	func Newads_ntfrsreplicasetEx1(instance *cim.WmiInstance) (newInstance *ads_ntfrsreplicaset, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_ntfrsreplicaset {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_ntfrsreplicasetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_ntfrsreplicaset, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_ntfrsreplicaset {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_fRSDirectoryFilter sets the value of DS_fRSDirectoryFilter for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_fRSDirectoryFilter(value string) (err error) { 
    return instance.SetProperty("DS_fRSDirectoryFilter", (value))
}

// GetDS_fRSDirectoryFilter gets the value of DS_fRSDirectoryFilter for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_fRSDirectoryFilter()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_fRSDirectoryFilter")
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

// SetDS_fRSDSPoll sets the value of DS_fRSDSPoll for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_fRSDSPoll(value int32) (err error) { 
    return instance.SetProperty("DS_fRSDSPoll", (value))
}

// GetDS_fRSDSPoll gets the value of DS_fRSDSPoll for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_fRSDSPoll()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_fRSDSPoll")
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

// SetDS_fRSExtensions sets the value of DS_fRSExtensions for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_fRSExtensions(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_fRSExtensions", (value))
}

// GetDS_fRSExtensions gets the value of DS_fRSExtensions for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_fRSExtensions()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_fRSExtensions")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

// SetDS_fRSFileFilter sets the value of DS_fRSFileFilter for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_fRSFileFilter(value string) (err error) { 
    return instance.SetProperty("DS_fRSFileFilter", (value))
}

// GetDS_fRSFileFilter gets the value of DS_fRSFileFilter for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_fRSFileFilter()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_fRSFileFilter")
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

// SetDS_fRSFlags sets the value of DS_fRSFlags for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_fRSFlags(value int32) (err error) { 
    return instance.SetProperty("DS_fRSFlags", (value))
}

// GetDS_fRSFlags gets the value of DS_fRSFlags for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_fRSFlags()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_fRSFlags")
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

// SetDS_fRSLevelLimit sets the value of DS_fRSLevelLimit for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_fRSLevelLimit(value int32) (err error) { 
    return instance.SetProperty("DS_fRSLevelLimit", (value))
}

// GetDS_fRSLevelLimit gets the value of DS_fRSLevelLimit for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_fRSLevelLimit()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_fRSLevelLimit")
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

// SetDS_fRSPartnerAuthLevel sets the value of DS_fRSPartnerAuthLevel for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_fRSPartnerAuthLevel(value int32) (err error) { 
    return instance.SetProperty("DS_fRSPartnerAuthLevel", (value))
}

// GetDS_fRSPartnerAuthLevel gets the value of DS_fRSPartnerAuthLevel for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_fRSPartnerAuthLevel()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_fRSPartnerAuthLevel")
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

// SetDS_fRSPrimaryMember sets the value of DS_fRSPrimaryMember for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_fRSPrimaryMember(value string) (err error) { 
    return instance.SetProperty("DS_fRSPrimaryMember", (value))
}

// GetDS_fRSPrimaryMember gets the value of DS_fRSPrimaryMember for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_fRSPrimaryMember()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_fRSPrimaryMember")
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

// SetDS_fRSReplicaSetGUID sets the value of DS_fRSReplicaSetGUID for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_fRSReplicaSetGUID(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_fRSReplicaSetGUID", (value))
}

// GetDS_fRSReplicaSetGUID gets the value of DS_fRSReplicaSetGUID for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_fRSReplicaSetGUID()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_fRSReplicaSetGUID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

// SetDS_fRSReplicaSetType sets the value of DS_fRSReplicaSetType for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_fRSReplicaSetType(value int32) (err error) { 
    return instance.SetProperty("DS_fRSReplicaSetType", (value))
}

// GetDS_fRSReplicaSetType gets the value of DS_fRSReplicaSetType for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_fRSReplicaSetType()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_fRSReplicaSetType")
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

// SetDS_fRSRootSecurity sets the value of DS_fRSRootSecurity for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_fRSRootSecurity(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_fRSRootSecurity", (value))
}

// GetDS_fRSRootSecurity gets the value of DS_fRSRootSecurity for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_fRSRootSecurity()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_fRSRootSecurity")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

// SetDS_fRSServiceCommand sets the value of DS_fRSServiceCommand for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_fRSServiceCommand(value string) (err error) { 
    return instance.SetProperty("DS_fRSServiceCommand", (value))
}

// GetDS_fRSServiceCommand gets the value of DS_fRSServiceCommand for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_fRSServiceCommand()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_fRSServiceCommand")
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

// SetDS_fRSVersionGUID sets the value of DS_fRSVersionGUID for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_fRSVersionGUID(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_fRSVersionGUID", (value))
}

// GetDS_fRSVersionGUID gets the value of DS_fRSVersionGUID for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_fRSVersionGUID()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_fRSVersionGUID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_managedBy(value string) (err error) { 
    return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_managedBy()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_managedBy")
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

// SetDS_msFRS_Hub_Member sets the value of DS_msFRS_Hub_Member for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_msFRS_Hub_Member(value string) (err error) { 
    return instance.SetProperty("DS_msFRS_Hub_Member", (value))
}

// GetDS_msFRS_Hub_Member gets the value of DS_msFRS_Hub_Member for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_msFRS_Hub_Member()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msFRS_Hub_Member")
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

// SetDS_msFRS_Topology_Pref sets the value of DS_msFRS_Topology_Pref for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_msFRS_Topology_Pref(value string) (err error) { 
    return instance.SetProperty("DS_msFRS_Topology_Pref", (value))
}

// GetDS_msFRS_Topology_Pref gets the value of DS_msFRS_Topology_Pref for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_msFRS_Topology_Pref()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msFRS_Topology_Pref")
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

// SetDS_schedule sets the value of DS_schedule for the instance
func (instance *ads_ntfrsreplicaset) SetPropertyDS_schedule(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_schedule", (value))
}

// GetDS_schedule gets the value of DS_schedule for the instance
func (instance *ads_ntfrsreplicaset) GetPropertyDS_schedule()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_schedule")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

