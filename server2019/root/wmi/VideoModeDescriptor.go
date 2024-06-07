// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// VideoModeDescriptor struct
type VideoModeDescriptor struct { 
	*cim.WmiInstance

	// 
	CompositePolarityType uint8

	// 
	HorizontalActivePixels uint16

	// 
	HorizontalBlankingPixels uint16

	// 
	HorizontalBorder uint16

	// 
	HorizontalImageSize uint16

	// 
	HorizontalPolarityType uint8

	// 
	HorizontalRefreshRateDenominator uint32

	// 
	HorizontalRefreshRateNumerator uint32

	// 
	HorizontalSyncOffset uint16

	// 
	HorizontalSyncPulseWidth uint16

	// 
	IsInterlaced bool

	// 
	IsSerrationRequired uint8

	// 
	IsSyncOnRGB uint8

	// 
	Origin uint8

	// 
	PixelClockRate uint32

	// 
	StereoModeType uint8

	// 
	SyncSignalType uint8

	// 
	TimingType uint8

	// 
	VerticalActivePixels uint16

	// 
	VerticalBlankingPixels uint16

	// 
	VerticalBorder uint16

	// 
	VerticalImageSize uint16

	// 
	VerticalPolarityType uint8

	// 
	VerticalRefreshRateDenominator uint32

	// 
	VerticalRefreshRateNumerator uint32

	// 
	VerticalSyncOffset uint16

	// 
	VerticalSyncPulseWidth uint16

	// 
	VideoStandardType uint8
}

	func NewVideoModeDescriptorEx1(instance *cim.WmiInstance) (newInstance *VideoModeDescriptor, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &VideoModeDescriptor {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewVideoModeDescriptorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *VideoModeDescriptor, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &VideoModeDescriptor {
	WmiInstance: tmp,
	}
	return
	}
	

// SetCompositePolarityType sets the value of CompositePolarityType for the instance
func (instance *VideoModeDescriptor) SetPropertyCompositePolarityType(value uint8) (err error) { 
    return instance.SetProperty("CompositePolarityType", (value))
}

// GetCompositePolarityType gets the value of CompositePolarityType for the instance
func (instance *VideoModeDescriptor) GetPropertyCompositePolarityType()(value uint8, err error) { 
    retValue, err := instance.GetProperty("CompositePolarityType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetHorizontalActivePixels sets the value of HorizontalActivePixels for the instance
func (instance *VideoModeDescriptor) SetPropertyHorizontalActivePixels(value uint16) (err error) { 
    return instance.SetProperty("HorizontalActivePixels", (value))
}

// GetHorizontalActivePixels gets the value of HorizontalActivePixels for the instance
func (instance *VideoModeDescriptor) GetPropertyHorizontalActivePixels()(value uint16, err error) { 
    retValue, err := instance.GetProperty("HorizontalActivePixels")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetHorizontalBlankingPixels sets the value of HorizontalBlankingPixels for the instance
func (instance *VideoModeDescriptor) SetPropertyHorizontalBlankingPixels(value uint16) (err error) { 
    return instance.SetProperty("HorizontalBlankingPixels", (value))
}

// GetHorizontalBlankingPixels gets the value of HorizontalBlankingPixels for the instance
func (instance *VideoModeDescriptor) GetPropertyHorizontalBlankingPixels()(value uint16, err error) { 
    retValue, err := instance.GetProperty("HorizontalBlankingPixels")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetHorizontalBorder sets the value of HorizontalBorder for the instance
func (instance *VideoModeDescriptor) SetPropertyHorizontalBorder(value uint16) (err error) { 
    return instance.SetProperty("HorizontalBorder", (value))
}

// GetHorizontalBorder gets the value of HorizontalBorder for the instance
func (instance *VideoModeDescriptor) GetPropertyHorizontalBorder()(value uint16, err error) { 
    retValue, err := instance.GetProperty("HorizontalBorder")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetHorizontalImageSize sets the value of HorizontalImageSize for the instance
func (instance *VideoModeDescriptor) SetPropertyHorizontalImageSize(value uint16) (err error) { 
    return instance.SetProperty("HorizontalImageSize", (value))
}

// GetHorizontalImageSize gets the value of HorizontalImageSize for the instance
func (instance *VideoModeDescriptor) GetPropertyHorizontalImageSize()(value uint16, err error) { 
    retValue, err := instance.GetProperty("HorizontalImageSize")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetHorizontalPolarityType sets the value of HorizontalPolarityType for the instance
func (instance *VideoModeDescriptor) SetPropertyHorizontalPolarityType(value uint8) (err error) { 
    return instance.SetProperty("HorizontalPolarityType", (value))
}

// GetHorizontalPolarityType gets the value of HorizontalPolarityType for the instance
func (instance *VideoModeDescriptor) GetPropertyHorizontalPolarityType()(value uint8, err error) { 
    retValue, err := instance.GetProperty("HorizontalPolarityType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetHorizontalRefreshRateDenominator sets the value of HorizontalRefreshRateDenominator for the instance
func (instance *VideoModeDescriptor) SetPropertyHorizontalRefreshRateDenominator(value uint32) (err error) { 
    return instance.SetProperty("HorizontalRefreshRateDenominator", (value))
}

// GetHorizontalRefreshRateDenominator gets the value of HorizontalRefreshRateDenominator for the instance
func (instance *VideoModeDescriptor) GetPropertyHorizontalRefreshRateDenominator()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HorizontalRefreshRateDenominator")
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

// SetHorizontalRefreshRateNumerator sets the value of HorizontalRefreshRateNumerator for the instance
func (instance *VideoModeDescriptor) SetPropertyHorizontalRefreshRateNumerator(value uint32) (err error) { 
    return instance.SetProperty("HorizontalRefreshRateNumerator", (value))
}

// GetHorizontalRefreshRateNumerator gets the value of HorizontalRefreshRateNumerator for the instance
func (instance *VideoModeDescriptor) GetPropertyHorizontalRefreshRateNumerator()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HorizontalRefreshRateNumerator")
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

// SetHorizontalSyncOffset sets the value of HorizontalSyncOffset for the instance
func (instance *VideoModeDescriptor) SetPropertyHorizontalSyncOffset(value uint16) (err error) { 
    return instance.SetProperty("HorizontalSyncOffset", (value))
}

// GetHorizontalSyncOffset gets the value of HorizontalSyncOffset for the instance
func (instance *VideoModeDescriptor) GetPropertyHorizontalSyncOffset()(value uint16, err error) { 
    retValue, err := instance.GetProperty("HorizontalSyncOffset")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetHorizontalSyncPulseWidth sets the value of HorizontalSyncPulseWidth for the instance
func (instance *VideoModeDescriptor) SetPropertyHorizontalSyncPulseWidth(value uint16) (err error) { 
    return instance.SetProperty("HorizontalSyncPulseWidth", (value))
}

// GetHorizontalSyncPulseWidth gets the value of HorizontalSyncPulseWidth for the instance
func (instance *VideoModeDescriptor) GetPropertyHorizontalSyncPulseWidth()(value uint16, err error) { 
    retValue, err := instance.GetProperty("HorizontalSyncPulseWidth")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetIsInterlaced sets the value of IsInterlaced for the instance
func (instance *VideoModeDescriptor) SetPropertyIsInterlaced(value bool) (err error) { 
    return instance.SetProperty("IsInterlaced", (value))
}

// GetIsInterlaced gets the value of IsInterlaced for the instance
func (instance *VideoModeDescriptor) GetPropertyIsInterlaced()(value bool, err error) { 
    retValue, err := instance.GetProperty("IsInterlaced")
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

// SetIsSerrationRequired sets the value of IsSerrationRequired for the instance
func (instance *VideoModeDescriptor) SetPropertyIsSerrationRequired(value uint8) (err error) { 
    return instance.SetProperty("IsSerrationRequired", (value))
}

// GetIsSerrationRequired gets the value of IsSerrationRequired for the instance
func (instance *VideoModeDescriptor) GetPropertyIsSerrationRequired()(value uint8, err error) { 
    retValue, err := instance.GetProperty("IsSerrationRequired")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetIsSyncOnRGB sets the value of IsSyncOnRGB for the instance
func (instance *VideoModeDescriptor) SetPropertyIsSyncOnRGB(value uint8) (err error) { 
    return instance.SetProperty("IsSyncOnRGB", (value))
}

// GetIsSyncOnRGB gets the value of IsSyncOnRGB for the instance
func (instance *VideoModeDescriptor) GetPropertyIsSyncOnRGB()(value uint8, err error) { 
    retValue, err := instance.GetProperty("IsSyncOnRGB")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetOrigin sets the value of Origin for the instance
func (instance *VideoModeDescriptor) SetPropertyOrigin(value uint8) (err error) { 
    return instance.SetProperty("Origin", (value))
}

// GetOrigin gets the value of Origin for the instance
func (instance *VideoModeDescriptor) GetPropertyOrigin()(value uint8, err error) { 
    retValue, err := instance.GetProperty("Origin")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetPixelClockRate sets the value of PixelClockRate for the instance
func (instance *VideoModeDescriptor) SetPropertyPixelClockRate(value uint32) (err error) { 
    return instance.SetProperty("PixelClockRate", (value))
}

// GetPixelClockRate gets the value of PixelClockRate for the instance
func (instance *VideoModeDescriptor) GetPropertyPixelClockRate()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PixelClockRate")
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

// SetStereoModeType sets the value of StereoModeType for the instance
func (instance *VideoModeDescriptor) SetPropertyStereoModeType(value uint8) (err error) { 
    return instance.SetProperty("StereoModeType", (value))
}

// GetStereoModeType gets the value of StereoModeType for the instance
func (instance *VideoModeDescriptor) GetPropertyStereoModeType()(value uint8, err error) { 
    retValue, err := instance.GetProperty("StereoModeType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetSyncSignalType sets the value of SyncSignalType for the instance
func (instance *VideoModeDescriptor) SetPropertySyncSignalType(value uint8) (err error) { 
    return instance.SetProperty("SyncSignalType", (value))
}

// GetSyncSignalType gets the value of SyncSignalType for the instance
func (instance *VideoModeDescriptor) GetPropertySyncSignalType()(value uint8, err error) { 
    retValue, err := instance.GetProperty("SyncSignalType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetTimingType sets the value of TimingType for the instance
func (instance *VideoModeDescriptor) SetPropertyTimingType(value uint8) (err error) { 
    return instance.SetProperty("TimingType", (value))
}

// GetTimingType gets the value of TimingType for the instance
func (instance *VideoModeDescriptor) GetPropertyTimingType()(value uint8, err error) { 
    retValue, err := instance.GetProperty("TimingType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetVerticalActivePixels sets the value of VerticalActivePixels for the instance
func (instance *VideoModeDescriptor) SetPropertyVerticalActivePixels(value uint16) (err error) { 
    return instance.SetProperty("VerticalActivePixels", (value))
}

// GetVerticalActivePixels gets the value of VerticalActivePixels for the instance
func (instance *VideoModeDescriptor) GetPropertyVerticalActivePixels()(value uint16, err error) { 
    retValue, err := instance.GetProperty("VerticalActivePixels")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetVerticalBlankingPixels sets the value of VerticalBlankingPixels for the instance
func (instance *VideoModeDescriptor) SetPropertyVerticalBlankingPixels(value uint16) (err error) { 
    return instance.SetProperty("VerticalBlankingPixels", (value))
}

// GetVerticalBlankingPixels gets the value of VerticalBlankingPixels for the instance
func (instance *VideoModeDescriptor) GetPropertyVerticalBlankingPixels()(value uint16, err error) { 
    retValue, err := instance.GetProperty("VerticalBlankingPixels")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetVerticalBorder sets the value of VerticalBorder for the instance
func (instance *VideoModeDescriptor) SetPropertyVerticalBorder(value uint16) (err error) { 
    return instance.SetProperty("VerticalBorder", (value))
}

// GetVerticalBorder gets the value of VerticalBorder for the instance
func (instance *VideoModeDescriptor) GetPropertyVerticalBorder()(value uint16, err error) { 
    retValue, err := instance.GetProperty("VerticalBorder")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetVerticalImageSize sets the value of VerticalImageSize for the instance
func (instance *VideoModeDescriptor) SetPropertyVerticalImageSize(value uint16) (err error) { 
    return instance.SetProperty("VerticalImageSize", (value))
}

// GetVerticalImageSize gets the value of VerticalImageSize for the instance
func (instance *VideoModeDescriptor) GetPropertyVerticalImageSize()(value uint16, err error) { 
    retValue, err := instance.GetProperty("VerticalImageSize")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetVerticalPolarityType sets the value of VerticalPolarityType for the instance
func (instance *VideoModeDescriptor) SetPropertyVerticalPolarityType(value uint8) (err error) { 
    return instance.SetProperty("VerticalPolarityType", (value))
}

// GetVerticalPolarityType gets the value of VerticalPolarityType for the instance
func (instance *VideoModeDescriptor) GetPropertyVerticalPolarityType()(value uint8, err error) { 
    retValue, err := instance.GetProperty("VerticalPolarityType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetVerticalRefreshRateDenominator sets the value of VerticalRefreshRateDenominator for the instance
func (instance *VideoModeDescriptor) SetPropertyVerticalRefreshRateDenominator(value uint32) (err error) { 
    return instance.SetProperty("VerticalRefreshRateDenominator", (value))
}

// GetVerticalRefreshRateDenominator gets the value of VerticalRefreshRateDenominator for the instance
func (instance *VideoModeDescriptor) GetPropertyVerticalRefreshRateDenominator()(value uint32, err error) { 
    retValue, err := instance.GetProperty("VerticalRefreshRateDenominator")
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

// SetVerticalRefreshRateNumerator sets the value of VerticalRefreshRateNumerator for the instance
func (instance *VideoModeDescriptor) SetPropertyVerticalRefreshRateNumerator(value uint32) (err error) { 
    return instance.SetProperty("VerticalRefreshRateNumerator", (value))
}

// GetVerticalRefreshRateNumerator gets the value of VerticalRefreshRateNumerator for the instance
func (instance *VideoModeDescriptor) GetPropertyVerticalRefreshRateNumerator()(value uint32, err error) { 
    retValue, err := instance.GetProperty("VerticalRefreshRateNumerator")
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

// SetVerticalSyncOffset sets the value of VerticalSyncOffset for the instance
func (instance *VideoModeDescriptor) SetPropertyVerticalSyncOffset(value uint16) (err error) { 
    return instance.SetProperty("VerticalSyncOffset", (value))
}

// GetVerticalSyncOffset gets the value of VerticalSyncOffset for the instance
func (instance *VideoModeDescriptor) GetPropertyVerticalSyncOffset()(value uint16, err error) { 
    retValue, err := instance.GetProperty("VerticalSyncOffset")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetVerticalSyncPulseWidth sets the value of VerticalSyncPulseWidth for the instance
func (instance *VideoModeDescriptor) SetPropertyVerticalSyncPulseWidth(value uint16) (err error) { 
    return instance.SetProperty("VerticalSyncPulseWidth", (value))
}

// GetVerticalSyncPulseWidth gets the value of VerticalSyncPulseWidth for the instance
func (instance *VideoModeDescriptor) GetPropertyVerticalSyncPulseWidth()(value uint16, err error) { 
    retValue, err := instance.GetProperty("VerticalSyncPulseWidth")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetVideoStandardType sets the value of VideoStandardType for the instance
func (instance *VideoModeDescriptor) SetPropertyVideoStandardType(value uint8) (err error) { 
    return instance.SetProperty("VideoStandardType", (value))
}

// GetVideoStandardType gets the value of VideoStandardType for the instance
func (instance *VideoModeDescriptor) GetPropertyVideoStandardType()(value uint8, err error) { 
    retValue, err := instance.GetProperty("VideoStandardType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

