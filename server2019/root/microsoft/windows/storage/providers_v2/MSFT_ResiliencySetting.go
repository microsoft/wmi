// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// MSFT_ResiliencySetting struct
type MSFT_ResiliencySetting struct {
	MSFT_StorageObject

	// A system set description of the capabilities of the resiliency setting, including (but not limited to) when a setting should be used, its strengths and drawbacks, performance information, and any other information that the vendor feels is helpful to the user.
	Description string

	// This field describes the desired number of bytes that will form a strip in common striping-based resiliency settings. The strip is defined as the size of the portion of a stripe that lies on one physical disk. Thus, Interleave * NumberOfColumns will yield the total size of one stripe.
	InterleaveDefault uint64

	// This field describes the maximum number of bytes that can form a strip in common striping-based resiliency settings. The strip is defined as the size of the portion of a stripe that lies on one physical disk.
	InterleaveMax uint64

	// This field describes the minimum number of bytes that can form a strip in common striping-based resiliency settings. The strip is defined as the size of the portion of a stripe that lies on one physical disk.
	InterleaveMin uint64

	// A system set, user-friendly, display-oriented string which describes the resiliency setting.
	Name string

	// This field is a user-settable preference for the number of underlying physical disks across which data should be striped.
	NumberOfColumnsDefault uint16

	// This field describes the maximum number of underlying physical disks across which data can be striped in the common striping-based resiliency settings.
	NumberOfColumnsMax uint16

	// This field describes the minimum number of underlying physical disks across which data can be striped in the common striping-based resiliency settings.
	NumberOfColumnsMin uint16

	// This field is a user-settable preference for the number of complete data copies to maintain. Its value must be within the range defined by NumberofDataCopiesMin and NumberOfDataCopiesMax (inclusive). For new concrete pools, the default should be inherited from the corresponding primordial pool's capability. In the case of the primordial pool, the initial value for this field is left to the Storage Management Provider software.
	NumberOfDataCopiesDefault uint16

	// This field reports the maximum number of complete copies of data that can be maintained by the storage pool.
	NumberOfDataCopiesMax uint16

	// This field reports the minimum number of complete copies of data that will be maintained by the storage pool.
	NumberOfDataCopiesMin uint16

	//
	NumberOfGroupsDefault uint16

	//
	NumberOfGroupsMax uint16

	//
	NumberOfGroupsMin uint16

	// This field specifies whether a parity-based resiliency setting is using a rotated or non-rotated parity layout. If the resiliency setting is not parity based, this field must be set to NULL
	ParityLayout ResiliencySetting_ParityLayout

	// This field is a user-settable preference for how many physical disk failures a virtual disk should be able to withstand before data loss occurs.
	PhysicalDiskRedundancyDefault uint16

	// This field reports the maximum number of tolerable physical disk failures that could occur before data loss would occur.
	PhysicalDiskRedundancyMax uint16

	// This field reports the minimum number of tolerable physical disk failures that can occur before data loss would occur.
	PhysicalDiskRedundancyMin uint16

	//
	RequestNoSinglePointOfFailure bool
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterleaveDefault sets the value of InterleaveDefault for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyInterleaveDefault(value uint64) (err error) {
	return instance.SetProperty("InterleaveDefault", value)
}

// GetInterleaveDefault gets the value of InterleaveDefault for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyInterleaveDefault() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterleaveDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterleaveMax sets the value of InterleaveMax for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyInterleaveMax(value uint64) (err error) {
	return instance.SetProperty("InterleaveMax", value)
}

// GetInterleaveMax gets the value of InterleaveMax for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyInterleaveMax() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterleaveMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterleaveMin sets the value of InterleaveMin for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyInterleaveMin(value uint64) (err error) {
	return instance.SetProperty("InterleaveMin", value)
}

// GetInterleaveMin gets the value of InterleaveMin for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyInterleaveMin() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterleaveMin")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfColumnsDefault sets the value of NumberOfColumnsDefault for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyNumberOfColumnsDefault(value uint16) (err error) {
	return instance.SetProperty("NumberOfColumnsDefault", value)
}

// GetNumberOfColumnsDefault gets the value of NumberOfColumnsDefault for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyNumberOfColumnsDefault() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfColumnsDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfColumnsMax sets the value of NumberOfColumnsMax for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyNumberOfColumnsMax(value uint16) (err error) {
	return instance.SetProperty("NumberOfColumnsMax", value)
}

// GetNumberOfColumnsMax gets the value of NumberOfColumnsMax for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyNumberOfColumnsMax() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfColumnsMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfColumnsMin sets the value of NumberOfColumnsMin for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyNumberOfColumnsMin(value uint16) (err error) {
	return instance.SetProperty("NumberOfColumnsMin", value)
}

// GetNumberOfColumnsMin gets the value of NumberOfColumnsMin for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyNumberOfColumnsMin() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfColumnsMin")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfDataCopiesDefault sets the value of NumberOfDataCopiesDefault for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyNumberOfDataCopiesDefault(value uint16) (err error) {
	return instance.SetProperty("NumberOfDataCopiesDefault", value)
}

// GetNumberOfDataCopiesDefault gets the value of NumberOfDataCopiesDefault for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyNumberOfDataCopiesDefault() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfDataCopiesDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfDataCopiesMax sets the value of NumberOfDataCopiesMax for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyNumberOfDataCopiesMax(value uint16) (err error) {
	return instance.SetProperty("NumberOfDataCopiesMax", value)
}

// GetNumberOfDataCopiesMax gets the value of NumberOfDataCopiesMax for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyNumberOfDataCopiesMax() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfDataCopiesMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfDataCopiesMin sets the value of NumberOfDataCopiesMin for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyNumberOfDataCopiesMin(value uint16) (err error) {
	return instance.SetProperty("NumberOfDataCopiesMin", value)
}

// GetNumberOfDataCopiesMin gets the value of NumberOfDataCopiesMin for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyNumberOfDataCopiesMin() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfDataCopiesMin")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfGroupsDefault sets the value of NumberOfGroupsDefault for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyNumberOfGroupsDefault(value uint16) (err error) {
	return instance.SetProperty("NumberOfGroupsDefault", value)
}

// GetNumberOfGroupsDefault gets the value of NumberOfGroupsDefault for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyNumberOfGroupsDefault() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfGroupsDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfGroupsMax sets the value of NumberOfGroupsMax for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyNumberOfGroupsMax(value uint16) (err error) {
	return instance.SetProperty("NumberOfGroupsMax", value)
}

// GetNumberOfGroupsMax gets the value of NumberOfGroupsMax for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyNumberOfGroupsMax() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfGroupsMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfGroupsMin sets the value of NumberOfGroupsMin for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyNumberOfGroupsMin(value uint16) (err error) {
	return instance.SetProperty("NumberOfGroupsMin", value)
}

// GetNumberOfGroupsMin gets the value of NumberOfGroupsMin for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyNumberOfGroupsMin() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfGroupsMin")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParityLayout sets the value of ParityLayout for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyParityLayout(value ResiliencySetting_ParityLayout) (err error) {
	return instance.SetProperty("ParityLayout", value)
}

// GetParityLayout gets the value of ParityLayout for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyParityLayout() (value ResiliencySetting_ParityLayout, err error) {
	retValue, err := instance.GetProperty("ParityLayout")
	if err != nil {
		return
	}
	value, ok := retValue.(ResiliencySetting_ParityLayout)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhysicalDiskRedundancyDefault sets the value of PhysicalDiskRedundancyDefault for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyPhysicalDiskRedundancyDefault(value uint16) (err error) {
	return instance.SetProperty("PhysicalDiskRedundancyDefault", value)
}

// GetPhysicalDiskRedundancyDefault gets the value of PhysicalDiskRedundancyDefault for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyPhysicalDiskRedundancyDefault() (value uint16, err error) {
	retValue, err := instance.GetProperty("PhysicalDiskRedundancyDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhysicalDiskRedundancyMax sets the value of PhysicalDiskRedundancyMax for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyPhysicalDiskRedundancyMax(value uint16) (err error) {
	return instance.SetProperty("PhysicalDiskRedundancyMax", value)
}

// GetPhysicalDiskRedundancyMax gets the value of PhysicalDiskRedundancyMax for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyPhysicalDiskRedundancyMax() (value uint16, err error) {
	retValue, err := instance.GetProperty("PhysicalDiskRedundancyMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhysicalDiskRedundancyMin sets the value of PhysicalDiskRedundancyMin for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyPhysicalDiskRedundancyMin(value uint16) (err error) {
	return instance.SetProperty("PhysicalDiskRedundancyMin", value)
}

// GetPhysicalDiskRedundancyMin gets the value of PhysicalDiskRedundancyMin for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyPhysicalDiskRedundancyMin() (value uint16, err error) {
	retValue, err := instance.GetProperty("PhysicalDiskRedundancyMin")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestNoSinglePointOfFailure sets the value of RequestNoSinglePointOfFailure for the instance
func (instance *MSFT_ResiliencySetting) SetPropertyRequestNoSinglePointOfFailure(value bool) (err error) {
	return instance.SetProperty("RequestNoSinglePointOfFailure", value)
}

// GetRequestNoSinglePointOfFailure gets the value of RequestNoSinglePointOfFailure for the instance
func (instance *MSFT_ResiliencySetting) GetPropertyRequestNoSinglePointOfFailure() (value bool, err error) {
	retValue, err := instance.GetProperty("RequestNoSinglePointOfFailure")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// This method allows a user to modify the default values for this resiliency setting.The updated values will take effect only for subsequent virtual disk creations and are not retroactively applied.

// <param name="AutoNumberOfColumns" type="bool ">If TRUE, this field instructs the storage provider (or subsystem) to automatically pick what it determines to be the best number of columns for this resiliency setting. If this field is TRUE, then the NumberOfColumnsDefault parameter must be NULL.</param>
// <param name="InterleaveDefault" type="uint64 ">Specifies the desired size of a data strip on a single physical disk in a striping based resiliency setting. This value must be between InterleaveMin and InterleaveMax. </param>
// <param name="NumberOfColumnsDefault" type="uint16 ">Specifies the desired number of physical disks to stripe data across. This value must be between NumberOfColumnsMin and NumberofColumnsMax.</param>
// <param name="NumberOfDataCopiesDefault" type="uint16 ">The desired number of full data copies to maintain. This value must be between NumberofDataCopiesMin and NumberofDataCopiesMax.</param>
// <param name="PhysicalDiskRedundancyDefault" type="uint16 ">The desired level of physical disk failure tolerance. This value must be between PhyscialDiskRedundancyMin and PhysicalDiskRedundancyMax.</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ResiliencySetting) SetDefaults( /* IN */ NumberOfDataCopiesDefault uint16,
	/* IN */ PhysicalDiskRedundancyDefault uint16,
	/* IN */ NumberOfColumnsDefault uint16,
	/* IN */ AutoNumberOfColumns bool,
	/* IN */ InterleaveDefault uint64,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetDefaults", NumberOfDataCopiesDefault, PhysicalDiskRedundancyDefault, NumberOfColumnsDefault, AutoNumberOfColumns, InterleaveDefault)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AutoNumberOfColumns" type="bool "></param>
// <param name="InterleaveDefault" type="uint64 "></param>
// <param name="NumberOfColumnsDefault" type="uint16 "></param>
// <param name="NumberOfDataCopiesDefault" type="uint16 "></param>
// <param name="NumberOfGroupsDefault" type="uint16 "></param>
// <param name="PhysicalDiskRedundancyDefault" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ResiliencySetting) SetDefaults2( /* IN */ NumberOfDataCopiesDefault uint16,
	/* IN */ PhysicalDiskRedundancyDefault uint16,
	/* IN */ NumberOfColumnsDefault uint16,
	/* IN */ AutoNumberOfColumns bool,
	/* IN */ InterleaveDefault uint64,
	/* IN */ NumberOfGroupsDefault uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetDefaults2", NumberOfDataCopiesDefault, PhysicalDiskRedundancyDefault, NumberOfColumnsDefault, AutoNumberOfColumns, InterleaveDefault, NumberOfGroupsDefault)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
