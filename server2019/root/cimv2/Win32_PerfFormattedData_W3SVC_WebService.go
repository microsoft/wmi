// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Win32_PerfFormattedData_W3SVC_WebService struct
type Win32_PerfFormattedData_W3SVC_WebService struct { 
	*Win32_PerfFormattedData

	// 
	AnonymousUsersPersec uint32

	// 
	BytesReceivedPersec uint64

	// 
	BytesSentPersec uint64

	// 
	BytesTotalPersec uint64

	// 
	CGIRequestsPersec uint32

	// 
	ConnectionAttemptsPersec uint32

	// 
	CopyRequestsPersec uint32

	// 
	CurrentAnonymousUsers uint32

	// 
	CurrentBlockedAsyncIORequests uint32

	// 
	Currentblockedbandwidthbytes uint32

	// 
	CurrentCALcountforauthenticatedusers uint32

	// 
	CurrentCALcountforSSLconnections uint32

	// 
	CurrentCGIRequests uint32

	// 
	CurrentConnections uint32

	// 
	CurrentISAPIExtensionRequests uint32

	// 
	CurrentNonAnonymousUsers uint32

	// 
	DeleteRequestsPersec uint32

	// 
	FilesPersec uint32

	// 
	FilesReceivedPersec uint32

	// 
	FilesSentPersec uint32

	// 
	GetRequestsPersec uint32

	// 
	HeadRequestsPersec uint32

	// 
	ISAPIExtensionRequestsPersec uint32

	// 
	LockedErrorsPersec uint32

	// 
	LockRequestsPersec uint32

	// 
	LogonAttemptsPersec uint32

	// 
	MaximumAnonymousUsers uint32

	// 
	MaximumCALcountforauthenticatedusers uint32

	// 
	MaximumCALcountforSSLconnections uint32

	// 
	MaximumCGIRequests uint32

	// 
	MaximumConnections uint32

	// 
	MaximumISAPIExtensionRequests uint32

	// 
	MaximumNonAnonymousUsers uint32

	// 
	MeasuredAsyncIOBandwidthUsage uint32

	// 
	MkcolRequestsPersec uint32

	// 
	MoveRequestsPersec uint32

	// 
	NonAnonymousUsersPersec uint32

	// 
	NotFoundErrorsPersec uint32

	// 
	OptionsRequestsPersec uint32

	// 
	OtherRequestMethodsPersec uint32

	// 
	PostRequestsPersec uint32

	// 
	PropfindRequestsPersec uint32

	// 
	ProppatchRequestsPersec uint32

	// 
	PutRequestsPersec uint32

	// 
	SearchRequestsPersec uint32

	// 
	ServiceUptime uint32

	// 
	TotalAllowedAsyncIORequests uint32

	// 
	TotalAnonymousUsers uint32

	// 
	TotalBlockedAsyncIORequests uint32

	// 
	Totalblockedbandwidthbytes uint32

	// 
	TotalBytesReceived uint64

	// 
	TotalBytesSent uint64

	// 
	TotalBytesTransferred uint64

	// 
	TotalCGIRequests uint32

	// 
	TotalConnectionAttemptsallinstances uint32

	// 
	TotalCopyRequests uint32

	// 
	TotalcountoffailedCALrequestsforauthenticatedusers uint32

	// 
	TotalcountoffailedCALrequestsforSSLconnections uint32

	// 
	TotalDeleteRequests uint32

	// 
	TotalFilesReceived uint32

	// 
	TotalFilesSent uint32

	// 
	TotalFilesTransferred uint32

	// 
	TotalGetRequests uint32

	// 
	TotalHeadRequests uint32

	// 
	TotalISAPIExtensionRequests uint32

	// 
	TotalLockedErrors uint32

	// 
	TotalLockRequests uint32

	// 
	TotalLogonAttempts uint32

	// 
	TotalMethodRequests uint32

	// 
	TotalMethodRequestsPersec uint32

	// 
	TotalMkcolRequests uint32

	// 
	TotalMoveRequests uint32

	// 
	TotalNonAnonymousUsers uint32

	// 
	TotalNotFoundErrors uint32

	// 
	TotalOptionsRequests uint32

	// 
	TotalOtherRequestMethods uint32

	// 
	TotalPostRequests uint32

	// 
	TotalPropfindRequests uint32

	// 
	TotalProppatchRequests uint32

	// 
	TotalPutRequests uint32

	// 
	TotalRejectedAsyncIORequests uint32

	// 
	TotalSearchRequests uint32

	// 
	TotalTraceRequests uint32

	// 
	TotalUnlockRequests uint32

	// 
	TraceRequestsPersec uint32

	// 
	UnlockRequestsPersec uint32
}

	func NewWin32_PerfFormattedData_W3SVC_WebServiceEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_W3SVC_WebService, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_W3SVC_WebService {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_W3SVC_WebServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_W3SVC_WebService, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_W3SVC_WebService {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetAnonymousUsersPersec sets the value of AnonymousUsersPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyAnonymousUsersPersec(value uint32) (err error) { 
    return instance.SetProperty("AnonymousUsersPersec", (value))
}

// GetAnonymousUsersPersec gets the value of AnonymousUsersPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyAnonymousUsersPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AnonymousUsersPersec")
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

// SetBytesReceivedPersec sets the value of BytesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyBytesReceivedPersec(value uint64) (err error) { 
    return instance.SetProperty("BytesReceivedPersec", (value))
}

// GetBytesReceivedPersec gets the value of BytesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyBytesReceivedPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BytesReceivedPersec")
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

// SetBytesSentPersec sets the value of BytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyBytesSentPersec(value uint64) (err error) { 
    return instance.SetProperty("BytesSentPersec", (value))
}

// GetBytesSentPersec gets the value of BytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyBytesSentPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BytesSentPersec")
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

// SetBytesTotalPersec sets the value of BytesTotalPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyBytesTotalPersec(value uint64) (err error) { 
    return instance.SetProperty("BytesTotalPersec", (value))
}

// GetBytesTotalPersec gets the value of BytesTotalPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyBytesTotalPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BytesTotalPersec")
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

// SetCGIRequestsPersec sets the value of CGIRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyCGIRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("CGIRequestsPersec", (value))
}

// GetCGIRequestsPersec gets the value of CGIRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyCGIRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CGIRequestsPersec")
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

// SetConnectionAttemptsPersec sets the value of ConnectionAttemptsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyConnectionAttemptsPersec(value uint32) (err error) { 
    return instance.SetProperty("ConnectionAttemptsPersec", (value))
}

// GetConnectionAttemptsPersec gets the value of ConnectionAttemptsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyConnectionAttemptsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ConnectionAttemptsPersec")
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

// SetCopyRequestsPersec sets the value of CopyRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyCopyRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("CopyRequestsPersec", (value))
}

// GetCopyRequestsPersec gets the value of CopyRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyCopyRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CopyRequestsPersec")
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

// SetCurrentAnonymousUsers sets the value of CurrentAnonymousUsers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyCurrentAnonymousUsers(value uint32) (err error) { 
    return instance.SetProperty("CurrentAnonymousUsers", (value))
}

// GetCurrentAnonymousUsers gets the value of CurrentAnonymousUsers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyCurrentAnonymousUsers()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentAnonymousUsers")
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

// SetCurrentBlockedAsyncIORequests sets the value of CurrentBlockedAsyncIORequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyCurrentBlockedAsyncIORequests(value uint32) (err error) { 
    return instance.SetProperty("CurrentBlockedAsyncIORequests", (value))
}

// GetCurrentBlockedAsyncIORequests gets the value of CurrentBlockedAsyncIORequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyCurrentBlockedAsyncIORequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentBlockedAsyncIORequests")
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

// SetCurrentblockedbandwidthbytes sets the value of Currentblockedbandwidthbytes for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyCurrentblockedbandwidthbytes(value uint32) (err error) { 
    return instance.SetProperty("Currentblockedbandwidthbytes", (value))
}

// GetCurrentblockedbandwidthbytes gets the value of Currentblockedbandwidthbytes for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyCurrentblockedbandwidthbytes()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Currentblockedbandwidthbytes")
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

// SetCurrentCALcountforauthenticatedusers sets the value of CurrentCALcountforauthenticatedusers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyCurrentCALcountforauthenticatedusers(value uint32) (err error) { 
    return instance.SetProperty("CurrentCALcountforauthenticatedusers", (value))
}

// GetCurrentCALcountforauthenticatedusers gets the value of CurrentCALcountforauthenticatedusers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyCurrentCALcountforauthenticatedusers()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentCALcountforauthenticatedusers")
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

// SetCurrentCALcountforSSLconnections sets the value of CurrentCALcountforSSLconnections for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyCurrentCALcountforSSLconnections(value uint32) (err error) { 
    return instance.SetProperty("CurrentCALcountforSSLconnections", (value))
}

// GetCurrentCALcountforSSLconnections gets the value of CurrentCALcountforSSLconnections for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyCurrentCALcountforSSLconnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentCALcountforSSLconnections")
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

// SetCurrentCGIRequests sets the value of CurrentCGIRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyCurrentCGIRequests(value uint32) (err error) { 
    return instance.SetProperty("CurrentCGIRequests", (value))
}

// GetCurrentCGIRequests gets the value of CurrentCGIRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyCurrentCGIRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentCGIRequests")
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

// SetCurrentConnections sets the value of CurrentConnections for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyCurrentConnections(value uint32) (err error) { 
    return instance.SetProperty("CurrentConnections", (value))
}

// GetCurrentConnections gets the value of CurrentConnections for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyCurrentConnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentConnections")
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

// SetCurrentISAPIExtensionRequests sets the value of CurrentISAPIExtensionRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyCurrentISAPIExtensionRequests(value uint32) (err error) { 
    return instance.SetProperty("CurrentISAPIExtensionRequests", (value))
}

// GetCurrentISAPIExtensionRequests gets the value of CurrentISAPIExtensionRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyCurrentISAPIExtensionRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentISAPIExtensionRequests")
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

// SetCurrentNonAnonymousUsers sets the value of CurrentNonAnonymousUsers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyCurrentNonAnonymousUsers(value uint32) (err error) { 
    return instance.SetProperty("CurrentNonAnonymousUsers", (value))
}

// GetCurrentNonAnonymousUsers gets the value of CurrentNonAnonymousUsers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyCurrentNonAnonymousUsers()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentNonAnonymousUsers")
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

// SetDeleteRequestsPersec sets the value of DeleteRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyDeleteRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("DeleteRequestsPersec", (value))
}

// GetDeleteRequestsPersec gets the value of DeleteRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyDeleteRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DeleteRequestsPersec")
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

// SetFilesPersec sets the value of FilesPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyFilesPersec(value uint32) (err error) { 
    return instance.SetProperty("FilesPersec", (value))
}

// GetFilesPersec gets the value of FilesPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyFilesPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FilesPersec")
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

// SetFilesReceivedPersec sets the value of FilesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyFilesReceivedPersec(value uint32) (err error) { 
    return instance.SetProperty("FilesReceivedPersec", (value))
}

// GetFilesReceivedPersec gets the value of FilesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyFilesReceivedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FilesReceivedPersec")
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

// SetFilesSentPersec sets the value of FilesSentPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyFilesSentPersec(value uint32) (err error) { 
    return instance.SetProperty("FilesSentPersec", (value))
}

// GetFilesSentPersec gets the value of FilesSentPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyFilesSentPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FilesSentPersec")
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

// SetGetRequestsPersec sets the value of GetRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyGetRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("GetRequestsPersec", (value))
}

// GetGetRequestsPersec gets the value of GetRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyGetRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("GetRequestsPersec")
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

// SetHeadRequestsPersec sets the value of HeadRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyHeadRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("HeadRequestsPersec", (value))
}

// GetHeadRequestsPersec gets the value of HeadRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyHeadRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HeadRequestsPersec")
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

// SetISAPIExtensionRequestsPersec sets the value of ISAPIExtensionRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyISAPIExtensionRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("ISAPIExtensionRequestsPersec", (value))
}

// GetISAPIExtensionRequestsPersec gets the value of ISAPIExtensionRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyISAPIExtensionRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ISAPIExtensionRequestsPersec")
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

// SetLockedErrorsPersec sets the value of LockedErrorsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyLockedErrorsPersec(value uint32) (err error) { 
    return instance.SetProperty("LockedErrorsPersec", (value))
}

// GetLockedErrorsPersec gets the value of LockedErrorsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyLockedErrorsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LockedErrorsPersec")
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

// SetLockRequestsPersec sets the value of LockRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyLockRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("LockRequestsPersec", (value))
}

// GetLockRequestsPersec gets the value of LockRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyLockRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LockRequestsPersec")
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

// SetLogonAttemptsPersec sets the value of LogonAttemptsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyLogonAttemptsPersec(value uint32) (err error) { 
    return instance.SetProperty("LogonAttemptsPersec", (value))
}

// GetLogonAttemptsPersec gets the value of LogonAttemptsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyLogonAttemptsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LogonAttemptsPersec")
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

// SetMaximumAnonymousUsers sets the value of MaximumAnonymousUsers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyMaximumAnonymousUsers(value uint32) (err error) { 
    return instance.SetProperty("MaximumAnonymousUsers", (value))
}

// GetMaximumAnonymousUsers gets the value of MaximumAnonymousUsers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyMaximumAnonymousUsers()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaximumAnonymousUsers")
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

// SetMaximumCALcountforauthenticatedusers sets the value of MaximumCALcountforauthenticatedusers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyMaximumCALcountforauthenticatedusers(value uint32) (err error) { 
    return instance.SetProperty("MaximumCALcountforauthenticatedusers", (value))
}

// GetMaximumCALcountforauthenticatedusers gets the value of MaximumCALcountforauthenticatedusers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyMaximumCALcountforauthenticatedusers()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaximumCALcountforauthenticatedusers")
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

// SetMaximumCALcountforSSLconnections sets the value of MaximumCALcountforSSLconnections for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyMaximumCALcountforSSLconnections(value uint32) (err error) { 
    return instance.SetProperty("MaximumCALcountforSSLconnections", (value))
}

// GetMaximumCALcountforSSLconnections gets the value of MaximumCALcountforSSLconnections for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyMaximumCALcountforSSLconnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaximumCALcountforSSLconnections")
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

// SetMaximumCGIRequests sets the value of MaximumCGIRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyMaximumCGIRequests(value uint32) (err error) { 
    return instance.SetProperty("MaximumCGIRequests", (value))
}

// GetMaximumCGIRequests gets the value of MaximumCGIRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyMaximumCGIRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaximumCGIRequests")
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

// SetMaximumConnections sets the value of MaximumConnections for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyMaximumConnections(value uint32) (err error) { 
    return instance.SetProperty("MaximumConnections", (value))
}

// GetMaximumConnections gets the value of MaximumConnections for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyMaximumConnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaximumConnections")
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

// SetMaximumISAPIExtensionRequests sets the value of MaximumISAPIExtensionRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyMaximumISAPIExtensionRequests(value uint32) (err error) { 
    return instance.SetProperty("MaximumISAPIExtensionRequests", (value))
}

// GetMaximumISAPIExtensionRequests gets the value of MaximumISAPIExtensionRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyMaximumISAPIExtensionRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaximumISAPIExtensionRequests")
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

// SetMaximumNonAnonymousUsers sets the value of MaximumNonAnonymousUsers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyMaximumNonAnonymousUsers(value uint32) (err error) { 
    return instance.SetProperty("MaximumNonAnonymousUsers", (value))
}

// GetMaximumNonAnonymousUsers gets the value of MaximumNonAnonymousUsers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyMaximumNonAnonymousUsers()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaximumNonAnonymousUsers")
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

// SetMeasuredAsyncIOBandwidthUsage sets the value of MeasuredAsyncIOBandwidthUsage for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyMeasuredAsyncIOBandwidthUsage(value uint32) (err error) { 
    return instance.SetProperty("MeasuredAsyncIOBandwidthUsage", (value))
}

// GetMeasuredAsyncIOBandwidthUsage gets the value of MeasuredAsyncIOBandwidthUsage for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyMeasuredAsyncIOBandwidthUsage()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MeasuredAsyncIOBandwidthUsage")
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

// SetMkcolRequestsPersec sets the value of MkcolRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyMkcolRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("MkcolRequestsPersec", (value))
}

// GetMkcolRequestsPersec gets the value of MkcolRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyMkcolRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MkcolRequestsPersec")
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

// SetMoveRequestsPersec sets the value of MoveRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyMoveRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("MoveRequestsPersec", (value))
}

// GetMoveRequestsPersec gets the value of MoveRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyMoveRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MoveRequestsPersec")
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

// SetNonAnonymousUsersPersec sets the value of NonAnonymousUsersPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyNonAnonymousUsersPersec(value uint32) (err error) { 
    return instance.SetProperty("NonAnonymousUsersPersec", (value))
}

// GetNonAnonymousUsersPersec gets the value of NonAnonymousUsersPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyNonAnonymousUsersPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NonAnonymousUsersPersec")
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

// SetNotFoundErrorsPersec sets the value of NotFoundErrorsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyNotFoundErrorsPersec(value uint32) (err error) { 
    return instance.SetProperty("NotFoundErrorsPersec", (value))
}

// GetNotFoundErrorsPersec gets the value of NotFoundErrorsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyNotFoundErrorsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NotFoundErrorsPersec")
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

// SetOptionsRequestsPersec sets the value of OptionsRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyOptionsRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("OptionsRequestsPersec", (value))
}

// GetOptionsRequestsPersec gets the value of OptionsRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyOptionsRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OptionsRequestsPersec")
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

// SetOtherRequestMethodsPersec sets the value of OtherRequestMethodsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyOtherRequestMethodsPersec(value uint32) (err error) { 
    return instance.SetProperty("OtherRequestMethodsPersec", (value))
}

// GetOtherRequestMethodsPersec gets the value of OtherRequestMethodsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyOtherRequestMethodsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OtherRequestMethodsPersec")
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

// SetPostRequestsPersec sets the value of PostRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyPostRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("PostRequestsPersec", (value))
}

// GetPostRequestsPersec gets the value of PostRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyPostRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PostRequestsPersec")
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

// SetPropfindRequestsPersec sets the value of PropfindRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyPropfindRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("PropfindRequestsPersec", (value))
}

// GetPropfindRequestsPersec gets the value of PropfindRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyPropfindRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PropfindRequestsPersec")
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

// SetProppatchRequestsPersec sets the value of ProppatchRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyProppatchRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("ProppatchRequestsPersec", (value))
}

// GetProppatchRequestsPersec gets the value of ProppatchRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyProppatchRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ProppatchRequestsPersec")
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

// SetPutRequestsPersec sets the value of PutRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyPutRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("PutRequestsPersec", (value))
}

// GetPutRequestsPersec gets the value of PutRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyPutRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PutRequestsPersec")
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

// SetSearchRequestsPersec sets the value of SearchRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertySearchRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("SearchRequestsPersec", (value))
}

// GetSearchRequestsPersec gets the value of SearchRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertySearchRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SearchRequestsPersec")
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

// SetServiceUptime sets the value of ServiceUptime for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyServiceUptime(value uint32) (err error) { 
    return instance.SetProperty("ServiceUptime", (value))
}

// GetServiceUptime gets the value of ServiceUptime for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyServiceUptime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ServiceUptime")
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

// SetTotalAllowedAsyncIORequests sets the value of TotalAllowedAsyncIORequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalAllowedAsyncIORequests(value uint32) (err error) { 
    return instance.SetProperty("TotalAllowedAsyncIORequests", (value))
}

// GetTotalAllowedAsyncIORequests gets the value of TotalAllowedAsyncIORequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalAllowedAsyncIORequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalAllowedAsyncIORequests")
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

// SetTotalAnonymousUsers sets the value of TotalAnonymousUsers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalAnonymousUsers(value uint32) (err error) { 
    return instance.SetProperty("TotalAnonymousUsers", (value))
}

// GetTotalAnonymousUsers gets the value of TotalAnonymousUsers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalAnonymousUsers()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalAnonymousUsers")
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

// SetTotalBlockedAsyncIORequests sets the value of TotalBlockedAsyncIORequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalBlockedAsyncIORequests(value uint32) (err error) { 
    return instance.SetProperty("TotalBlockedAsyncIORequests", (value))
}

// GetTotalBlockedAsyncIORequests gets the value of TotalBlockedAsyncIORequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalBlockedAsyncIORequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalBlockedAsyncIORequests")
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

// SetTotalblockedbandwidthbytes sets the value of Totalblockedbandwidthbytes for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalblockedbandwidthbytes(value uint32) (err error) { 
    return instance.SetProperty("Totalblockedbandwidthbytes", (value))
}

// GetTotalblockedbandwidthbytes gets the value of Totalblockedbandwidthbytes for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalblockedbandwidthbytes()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Totalblockedbandwidthbytes")
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

// SetTotalBytesReceived sets the value of TotalBytesReceived for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalBytesReceived(value uint64) (err error) { 
    return instance.SetProperty("TotalBytesReceived", (value))
}

// GetTotalBytesReceived gets the value of TotalBytesReceived for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalBytesReceived()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TotalBytesReceived")
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

// SetTotalBytesSent sets the value of TotalBytesSent for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalBytesSent(value uint64) (err error) { 
    return instance.SetProperty("TotalBytesSent", (value))
}

// GetTotalBytesSent gets the value of TotalBytesSent for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalBytesSent()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TotalBytesSent")
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

// SetTotalBytesTransferred sets the value of TotalBytesTransferred for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalBytesTransferred(value uint64) (err error) { 
    return instance.SetProperty("TotalBytesTransferred", (value))
}

// GetTotalBytesTransferred gets the value of TotalBytesTransferred for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalBytesTransferred()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TotalBytesTransferred")
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

// SetTotalCGIRequests sets the value of TotalCGIRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalCGIRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalCGIRequests", (value))
}

// GetTotalCGIRequests gets the value of TotalCGIRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalCGIRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalCGIRequests")
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

// SetTotalConnectionAttemptsallinstances sets the value of TotalConnectionAttemptsallinstances for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalConnectionAttemptsallinstances(value uint32) (err error) { 
    return instance.SetProperty("TotalConnectionAttemptsallinstances", (value))
}

// GetTotalConnectionAttemptsallinstances gets the value of TotalConnectionAttemptsallinstances for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalConnectionAttemptsallinstances()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalConnectionAttemptsallinstances")
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

// SetTotalCopyRequests sets the value of TotalCopyRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalCopyRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalCopyRequests", (value))
}

// GetTotalCopyRequests gets the value of TotalCopyRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalCopyRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalCopyRequests")
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

// SetTotalcountoffailedCALrequestsforauthenticatedusers sets the value of TotalcountoffailedCALrequestsforauthenticatedusers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalcountoffailedCALrequestsforauthenticatedusers(value uint32) (err error) { 
    return instance.SetProperty("TotalcountoffailedCALrequestsforauthenticatedusers", (value))
}

// GetTotalcountoffailedCALrequestsforauthenticatedusers gets the value of TotalcountoffailedCALrequestsforauthenticatedusers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalcountoffailedCALrequestsforauthenticatedusers()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalcountoffailedCALrequestsforauthenticatedusers")
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

// SetTotalcountoffailedCALrequestsforSSLconnections sets the value of TotalcountoffailedCALrequestsforSSLconnections for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalcountoffailedCALrequestsforSSLconnections(value uint32) (err error) { 
    return instance.SetProperty("TotalcountoffailedCALrequestsforSSLconnections", (value))
}

// GetTotalcountoffailedCALrequestsforSSLconnections gets the value of TotalcountoffailedCALrequestsforSSLconnections for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalcountoffailedCALrequestsforSSLconnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalcountoffailedCALrequestsforSSLconnections")
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

// SetTotalDeleteRequests sets the value of TotalDeleteRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalDeleteRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalDeleteRequests", (value))
}

// GetTotalDeleteRequests gets the value of TotalDeleteRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalDeleteRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalDeleteRequests")
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

// SetTotalFilesReceived sets the value of TotalFilesReceived for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalFilesReceived(value uint32) (err error) { 
    return instance.SetProperty("TotalFilesReceived", (value))
}

// GetTotalFilesReceived gets the value of TotalFilesReceived for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalFilesReceived()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalFilesReceived")
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

// SetTotalFilesSent sets the value of TotalFilesSent for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalFilesSent(value uint32) (err error) { 
    return instance.SetProperty("TotalFilesSent", (value))
}

// GetTotalFilesSent gets the value of TotalFilesSent for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalFilesSent()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalFilesSent")
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

// SetTotalFilesTransferred sets the value of TotalFilesTransferred for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalFilesTransferred(value uint32) (err error) { 
    return instance.SetProperty("TotalFilesTransferred", (value))
}

// GetTotalFilesTransferred gets the value of TotalFilesTransferred for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalFilesTransferred()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalFilesTransferred")
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

// SetTotalGetRequests sets the value of TotalGetRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalGetRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalGetRequests", (value))
}

// GetTotalGetRequests gets the value of TotalGetRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalGetRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalGetRequests")
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

// SetTotalHeadRequests sets the value of TotalHeadRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalHeadRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalHeadRequests", (value))
}

// GetTotalHeadRequests gets the value of TotalHeadRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalHeadRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalHeadRequests")
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

// SetTotalISAPIExtensionRequests sets the value of TotalISAPIExtensionRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalISAPIExtensionRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalISAPIExtensionRequests", (value))
}

// GetTotalISAPIExtensionRequests gets the value of TotalISAPIExtensionRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalISAPIExtensionRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalISAPIExtensionRequests")
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

// SetTotalLockedErrors sets the value of TotalLockedErrors for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalLockedErrors(value uint32) (err error) { 
    return instance.SetProperty("TotalLockedErrors", (value))
}

// GetTotalLockedErrors gets the value of TotalLockedErrors for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalLockedErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalLockedErrors")
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

// SetTotalLockRequests sets the value of TotalLockRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalLockRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalLockRequests", (value))
}

// GetTotalLockRequests gets the value of TotalLockRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalLockRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalLockRequests")
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

// SetTotalLogonAttempts sets the value of TotalLogonAttempts for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalLogonAttempts(value uint32) (err error) { 
    return instance.SetProperty("TotalLogonAttempts", (value))
}

// GetTotalLogonAttempts gets the value of TotalLogonAttempts for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalLogonAttempts()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalLogonAttempts")
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

// SetTotalMethodRequests sets the value of TotalMethodRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalMethodRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalMethodRequests", (value))
}

// GetTotalMethodRequests gets the value of TotalMethodRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalMethodRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalMethodRequests")
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

// SetTotalMethodRequestsPersec sets the value of TotalMethodRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalMethodRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("TotalMethodRequestsPersec", (value))
}

// GetTotalMethodRequestsPersec gets the value of TotalMethodRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalMethodRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalMethodRequestsPersec")
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

// SetTotalMkcolRequests sets the value of TotalMkcolRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalMkcolRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalMkcolRequests", (value))
}

// GetTotalMkcolRequests gets the value of TotalMkcolRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalMkcolRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalMkcolRequests")
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

// SetTotalMoveRequests sets the value of TotalMoveRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalMoveRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalMoveRequests", (value))
}

// GetTotalMoveRequests gets the value of TotalMoveRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalMoveRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalMoveRequests")
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

// SetTotalNonAnonymousUsers sets the value of TotalNonAnonymousUsers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalNonAnonymousUsers(value uint32) (err error) { 
    return instance.SetProperty("TotalNonAnonymousUsers", (value))
}

// GetTotalNonAnonymousUsers gets the value of TotalNonAnonymousUsers for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalNonAnonymousUsers()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalNonAnonymousUsers")
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

// SetTotalNotFoundErrors sets the value of TotalNotFoundErrors for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalNotFoundErrors(value uint32) (err error) { 
    return instance.SetProperty("TotalNotFoundErrors", (value))
}

// GetTotalNotFoundErrors gets the value of TotalNotFoundErrors for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalNotFoundErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalNotFoundErrors")
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

// SetTotalOptionsRequests sets the value of TotalOptionsRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalOptionsRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalOptionsRequests", (value))
}

// GetTotalOptionsRequests gets the value of TotalOptionsRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalOptionsRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalOptionsRequests")
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

// SetTotalOtherRequestMethods sets the value of TotalOtherRequestMethods for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalOtherRequestMethods(value uint32) (err error) { 
    return instance.SetProperty("TotalOtherRequestMethods", (value))
}

// GetTotalOtherRequestMethods gets the value of TotalOtherRequestMethods for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalOtherRequestMethods()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalOtherRequestMethods")
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

// SetTotalPostRequests sets the value of TotalPostRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalPostRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalPostRequests", (value))
}

// GetTotalPostRequests gets the value of TotalPostRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalPostRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalPostRequests")
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

// SetTotalPropfindRequests sets the value of TotalPropfindRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalPropfindRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalPropfindRequests", (value))
}

// GetTotalPropfindRequests gets the value of TotalPropfindRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalPropfindRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalPropfindRequests")
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

// SetTotalProppatchRequests sets the value of TotalProppatchRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalProppatchRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalProppatchRequests", (value))
}

// GetTotalProppatchRequests gets the value of TotalProppatchRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalProppatchRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalProppatchRequests")
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

// SetTotalPutRequests sets the value of TotalPutRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalPutRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalPutRequests", (value))
}

// GetTotalPutRequests gets the value of TotalPutRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalPutRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalPutRequests")
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

// SetTotalRejectedAsyncIORequests sets the value of TotalRejectedAsyncIORequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalRejectedAsyncIORequests(value uint32) (err error) { 
    return instance.SetProperty("TotalRejectedAsyncIORequests", (value))
}

// GetTotalRejectedAsyncIORequests gets the value of TotalRejectedAsyncIORequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalRejectedAsyncIORequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalRejectedAsyncIORequests")
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

// SetTotalSearchRequests sets the value of TotalSearchRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalSearchRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalSearchRequests", (value))
}

// GetTotalSearchRequests gets the value of TotalSearchRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalSearchRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalSearchRequests")
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

// SetTotalTraceRequests sets the value of TotalTraceRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalTraceRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalTraceRequests", (value))
}

// GetTotalTraceRequests gets the value of TotalTraceRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalTraceRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalTraceRequests")
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

// SetTotalUnlockRequests sets the value of TotalUnlockRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTotalUnlockRequests(value uint32) (err error) { 
    return instance.SetProperty("TotalUnlockRequests", (value))
}

// GetTotalUnlockRequests gets the value of TotalUnlockRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTotalUnlockRequests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalUnlockRequests")
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

// SetTraceRequestsPersec sets the value of TraceRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyTraceRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("TraceRequestsPersec", (value))
}

// GetTraceRequestsPersec gets the value of TraceRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyTraceRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TraceRequestsPersec")
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

// SetUnlockRequestsPersec sets the value of UnlockRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) SetPropertyUnlockRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("UnlockRequestsPersec", (value))
}

// GetUnlockRequestsPersec gets the value of UnlockRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVC_WebService) GetPropertyUnlockRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("UnlockRequestsPersec")
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

