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

// ads_msdfs_namespacev2 struct
type ads_msdfs_namespacev2 struct { 
	*ds_top

	// 
	DS_msDFS_Commentv2 string

	// 
	DS_msDFS_GenerationGUIDv2 Uint8Array

	// 
	DS_msDFS_LastModifiedv2 string

	// 
	DS_msDFS_NamespaceIdentityGUIDv2 Uint8Array

	// 
	DS_msDFS_Propertiesv2 []string

	// 
	DS_msDFS_SchemaMajorVersion int32

	// 
	DS_msDFS_SchemaMinorVersion int32

	// 
	DS_msDFS_TargetListv2 Uint8Array

	// 
	DS_msDFS_Ttlv2 int32
}

	func Newads_msdfs_namespacev2Ex1(instance *cim.WmiInstance) (newInstance *ads_msdfs_namespacev2, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_msdfs_namespacev2 {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_msdfs_namespacev2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_msdfs_namespacev2, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_msdfs_namespacev2 {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_msDFS_Commentv2 sets the value of DS_msDFS_Commentv2 for the instance
func (instance *ads_msdfs_namespacev2) SetPropertyDS_msDFS_Commentv2(value string) (err error) { 
    return instance.SetProperty("DS_msDFS_Commentv2", (value))
}

// GetDS_msDFS_Commentv2 gets the value of DS_msDFS_Commentv2 for the instance
func (instance *ads_msdfs_namespacev2) GetPropertyDS_msDFS_Commentv2()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msDFS_Commentv2")
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

// SetDS_msDFS_GenerationGUIDv2 sets the value of DS_msDFS_GenerationGUIDv2 for the instance
func (instance *ads_msdfs_namespacev2) SetPropertyDS_msDFS_GenerationGUIDv2(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_msDFS_GenerationGUIDv2", (value))
}

// GetDS_msDFS_GenerationGUIDv2 gets the value of DS_msDFS_GenerationGUIDv2 for the instance
func (instance *ads_msdfs_namespacev2) GetPropertyDS_msDFS_GenerationGUIDv2()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_msDFS_GenerationGUIDv2")
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

// SetDS_msDFS_LastModifiedv2 sets the value of DS_msDFS_LastModifiedv2 for the instance
func (instance *ads_msdfs_namespacev2) SetPropertyDS_msDFS_LastModifiedv2(value string) (err error) { 
    return instance.SetProperty("DS_msDFS_LastModifiedv2", (value))
}

// GetDS_msDFS_LastModifiedv2 gets the value of DS_msDFS_LastModifiedv2 for the instance
func (instance *ads_msdfs_namespacev2) GetPropertyDS_msDFS_LastModifiedv2()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msDFS_LastModifiedv2")
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

// SetDS_msDFS_NamespaceIdentityGUIDv2 sets the value of DS_msDFS_NamespaceIdentityGUIDv2 for the instance
func (instance *ads_msdfs_namespacev2) SetPropertyDS_msDFS_NamespaceIdentityGUIDv2(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_msDFS_NamespaceIdentityGUIDv2", (value))
}

// GetDS_msDFS_NamespaceIdentityGUIDv2 gets the value of DS_msDFS_NamespaceIdentityGUIDv2 for the instance
func (instance *ads_msdfs_namespacev2) GetPropertyDS_msDFS_NamespaceIdentityGUIDv2()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_msDFS_NamespaceIdentityGUIDv2")
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

// SetDS_msDFS_Propertiesv2 sets the value of DS_msDFS_Propertiesv2 for the instance
func (instance *ads_msdfs_namespacev2) SetPropertyDS_msDFS_Propertiesv2(value []string) (err error) { 
    return instance.SetProperty("DS_msDFS_Propertiesv2", (value))
}

// GetDS_msDFS_Propertiesv2 gets the value of DS_msDFS_Propertiesv2 for the instance
func (instance *ads_msdfs_namespacev2) GetPropertyDS_msDFS_Propertiesv2()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msDFS_Propertiesv2")
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

// SetDS_msDFS_SchemaMajorVersion sets the value of DS_msDFS_SchemaMajorVersion for the instance
func (instance *ads_msdfs_namespacev2) SetPropertyDS_msDFS_SchemaMajorVersion(value int32) (err error) { 
    return instance.SetProperty("DS_msDFS_SchemaMajorVersion", (value))
}

// GetDS_msDFS_SchemaMajorVersion gets the value of DS_msDFS_SchemaMajorVersion for the instance
func (instance *ads_msdfs_namespacev2) GetPropertyDS_msDFS_SchemaMajorVersion()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msDFS_SchemaMajorVersion")
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

// SetDS_msDFS_SchemaMinorVersion sets the value of DS_msDFS_SchemaMinorVersion for the instance
func (instance *ads_msdfs_namespacev2) SetPropertyDS_msDFS_SchemaMinorVersion(value int32) (err error) { 
    return instance.SetProperty("DS_msDFS_SchemaMinorVersion", (value))
}

// GetDS_msDFS_SchemaMinorVersion gets the value of DS_msDFS_SchemaMinorVersion for the instance
func (instance *ads_msdfs_namespacev2) GetPropertyDS_msDFS_SchemaMinorVersion()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msDFS_SchemaMinorVersion")
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

// SetDS_msDFS_TargetListv2 sets the value of DS_msDFS_TargetListv2 for the instance
func (instance *ads_msdfs_namespacev2) SetPropertyDS_msDFS_TargetListv2(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_msDFS_TargetListv2", (value))
}

// GetDS_msDFS_TargetListv2 gets the value of DS_msDFS_TargetListv2 for the instance
func (instance *ads_msdfs_namespacev2) GetPropertyDS_msDFS_TargetListv2()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_msDFS_TargetListv2")
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

// SetDS_msDFS_Ttlv2 sets the value of DS_msDFS_Ttlv2 for the instance
func (instance *ads_msdfs_namespacev2) SetPropertyDS_msDFS_Ttlv2(value int32) (err error) { 
    return instance.SetProperty("DS_msDFS_Ttlv2", (value))
}

// GetDS_msDFS_Ttlv2 gets the value of DS_msDFS_Ttlv2 for the instance
func (instance *ads_msdfs_namespacev2) GetPropertyDS_msDFS_Ttlv2()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msDFS_Ttlv2")
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

