// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// FtpFileHandling struct
type FtpFileHandling struct { 
	*EmbeddedObject

	// 
	AllowReadUploadsInProgress bool

	// 
	AllowReplaceOnRename bool

	// 
	KeepPartialUploads bool
}

	func NewFtpFileHandlingEx1(instance *cim.WmiInstance) (newInstance *FtpFileHandling, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &FtpFileHandling {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewFtpFileHandlingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FtpFileHandling, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FtpFileHandling {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetAllowReadUploadsInProgress sets the value of AllowReadUploadsInProgress for the instance
func (instance *FtpFileHandling) SetPropertyAllowReadUploadsInProgress(value bool) (err error) { 
    return instance.SetProperty("AllowReadUploadsInProgress", (value))
}

// GetAllowReadUploadsInProgress gets the value of AllowReadUploadsInProgress for the instance
func (instance *FtpFileHandling) GetPropertyAllowReadUploadsInProgress()(value bool, err error) { 
    retValue, err := instance.GetProperty("AllowReadUploadsInProgress")
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

// SetAllowReplaceOnRename sets the value of AllowReplaceOnRename for the instance
func (instance *FtpFileHandling) SetPropertyAllowReplaceOnRename(value bool) (err error) { 
    return instance.SetProperty("AllowReplaceOnRename", (value))
}

// GetAllowReplaceOnRename gets the value of AllowReplaceOnRename for the instance
func (instance *FtpFileHandling) GetPropertyAllowReplaceOnRename()(value bool, err error) { 
    retValue, err := instance.GetProperty("AllowReplaceOnRename")
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

// SetKeepPartialUploads sets the value of KeepPartialUploads for the instance
func (instance *FtpFileHandling) SetPropertyKeepPartialUploads(value bool) (err error) { 
    return instance.SetProperty("KeepPartialUploads", (value))
}

// GetKeepPartialUploads gets the value of KeepPartialUploads for the instance
func (instance *FtpFileHandling) GetPropertyKeepPartialUploads()(value bool, err error) { 
    retValue, err := instance.GetProperty("KeepPartialUploads")
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

