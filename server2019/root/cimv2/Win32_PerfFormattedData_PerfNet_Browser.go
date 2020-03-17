// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_PerfNet_Browser struct
type Win32_PerfFormattedData_PerfNet_Browser struct {
	Win32_PerfFormattedData

	//
	AnnouncementsDomainPersec uint64

	//
	AnnouncementsServerPersec uint64

	//
	AnnouncementsTotalPersec uint64

	//
	DuplicateMasterAnnouncements uint32

	//
	ElectionPacketsPersec uint32

	//
	EnumerationsDomainPersec uint32

	//
	EnumerationsOtherPersec uint32

	//
	EnumerationsServerPersec uint32

	//
	EnumerationsTotalPersec uint32

	//
	IllegalDatagramsPersec uint64

	//
	MailslotAllocationsFailed uint32

	//
	MailslotOpensFailedPersec uint32

	//
	MailslotReceivesFailed uint32

	//
	MailslotWritesFailed uint32

	//
	MailslotWritesPersec uint32

	//
	MissedMailslotDatagrams uint32

	//
	MissedServerAnnouncements uint32

	//
	MissedServerListRequests uint32

	//
	ServerAnnounceAllocationsFailedPersec uint32

	//
	ServerListRequestsPersec uint32
}

// SetAnnouncementsDomainPersec sets the value of AnnouncementsDomainPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyAnnouncementsDomainPersec(value uint64) (err error) {
	return instance.SetProperty("AnnouncementsDomainPersec", value)
}

// GetAnnouncementsDomainPersec gets the value of AnnouncementsDomainPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyAnnouncementsDomainPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AnnouncementsDomainPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAnnouncementsServerPersec sets the value of AnnouncementsServerPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyAnnouncementsServerPersec(value uint64) (err error) {
	return instance.SetProperty("AnnouncementsServerPersec", value)
}

// GetAnnouncementsServerPersec gets the value of AnnouncementsServerPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyAnnouncementsServerPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AnnouncementsServerPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAnnouncementsTotalPersec sets the value of AnnouncementsTotalPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyAnnouncementsTotalPersec(value uint64) (err error) {
	return instance.SetProperty("AnnouncementsTotalPersec", value)
}

// GetAnnouncementsTotalPersec gets the value of AnnouncementsTotalPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyAnnouncementsTotalPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AnnouncementsTotalPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDuplicateMasterAnnouncements sets the value of DuplicateMasterAnnouncements for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyDuplicateMasterAnnouncements(value uint32) (err error) {
	return instance.SetProperty("DuplicateMasterAnnouncements", value)
}

// GetDuplicateMasterAnnouncements gets the value of DuplicateMasterAnnouncements for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyDuplicateMasterAnnouncements() (value uint32, err error) {
	retValue, err := instance.GetProperty("DuplicateMasterAnnouncements")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetElectionPacketsPersec sets the value of ElectionPacketsPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyElectionPacketsPersec(value uint32) (err error) {
	return instance.SetProperty("ElectionPacketsPersec", value)
}

// GetElectionPacketsPersec gets the value of ElectionPacketsPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyElectionPacketsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ElectionPacketsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnumerationsDomainPersec sets the value of EnumerationsDomainPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyEnumerationsDomainPersec(value uint32) (err error) {
	return instance.SetProperty("EnumerationsDomainPersec", value)
}

// GetEnumerationsDomainPersec gets the value of EnumerationsDomainPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyEnumerationsDomainPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnumerationsDomainPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnumerationsOtherPersec sets the value of EnumerationsOtherPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyEnumerationsOtherPersec(value uint32) (err error) {
	return instance.SetProperty("EnumerationsOtherPersec", value)
}

// GetEnumerationsOtherPersec gets the value of EnumerationsOtherPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyEnumerationsOtherPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnumerationsOtherPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnumerationsServerPersec sets the value of EnumerationsServerPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyEnumerationsServerPersec(value uint32) (err error) {
	return instance.SetProperty("EnumerationsServerPersec", value)
}

// GetEnumerationsServerPersec gets the value of EnumerationsServerPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyEnumerationsServerPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnumerationsServerPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnumerationsTotalPersec sets the value of EnumerationsTotalPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyEnumerationsTotalPersec(value uint32) (err error) {
	return instance.SetProperty("EnumerationsTotalPersec", value)
}

// GetEnumerationsTotalPersec gets the value of EnumerationsTotalPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyEnumerationsTotalPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnumerationsTotalPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIllegalDatagramsPersec sets the value of IllegalDatagramsPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyIllegalDatagramsPersec(value uint64) (err error) {
	return instance.SetProperty("IllegalDatagramsPersec", value)
}

// GetIllegalDatagramsPersec gets the value of IllegalDatagramsPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyIllegalDatagramsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IllegalDatagramsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMailslotAllocationsFailed sets the value of MailslotAllocationsFailed for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyMailslotAllocationsFailed(value uint32) (err error) {
	return instance.SetProperty("MailslotAllocationsFailed", value)
}

// GetMailslotAllocationsFailed gets the value of MailslotAllocationsFailed for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyMailslotAllocationsFailed() (value uint32, err error) {
	retValue, err := instance.GetProperty("MailslotAllocationsFailed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMailslotOpensFailedPersec sets the value of MailslotOpensFailedPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyMailslotOpensFailedPersec(value uint32) (err error) {
	return instance.SetProperty("MailslotOpensFailedPersec", value)
}

// GetMailslotOpensFailedPersec gets the value of MailslotOpensFailedPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyMailslotOpensFailedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("MailslotOpensFailedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMailslotReceivesFailed sets the value of MailslotReceivesFailed for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyMailslotReceivesFailed(value uint32) (err error) {
	return instance.SetProperty("MailslotReceivesFailed", value)
}

// GetMailslotReceivesFailed gets the value of MailslotReceivesFailed for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyMailslotReceivesFailed() (value uint32, err error) {
	retValue, err := instance.GetProperty("MailslotReceivesFailed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMailslotWritesFailed sets the value of MailslotWritesFailed for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyMailslotWritesFailed(value uint32) (err error) {
	return instance.SetProperty("MailslotWritesFailed", value)
}

// GetMailslotWritesFailed gets the value of MailslotWritesFailed for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyMailslotWritesFailed() (value uint32, err error) {
	retValue, err := instance.GetProperty("MailslotWritesFailed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMailslotWritesPersec sets the value of MailslotWritesPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyMailslotWritesPersec(value uint32) (err error) {
	return instance.SetProperty("MailslotWritesPersec", value)
}

// GetMailslotWritesPersec gets the value of MailslotWritesPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyMailslotWritesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("MailslotWritesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMissedMailslotDatagrams sets the value of MissedMailslotDatagrams for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyMissedMailslotDatagrams(value uint32) (err error) {
	return instance.SetProperty("MissedMailslotDatagrams", value)
}

// GetMissedMailslotDatagrams gets the value of MissedMailslotDatagrams for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyMissedMailslotDatagrams() (value uint32, err error) {
	retValue, err := instance.GetProperty("MissedMailslotDatagrams")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMissedServerAnnouncements sets the value of MissedServerAnnouncements for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyMissedServerAnnouncements(value uint32) (err error) {
	return instance.SetProperty("MissedServerAnnouncements", value)
}

// GetMissedServerAnnouncements gets the value of MissedServerAnnouncements for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyMissedServerAnnouncements() (value uint32, err error) {
	retValue, err := instance.GetProperty("MissedServerAnnouncements")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMissedServerListRequests sets the value of MissedServerListRequests for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyMissedServerListRequests(value uint32) (err error) {
	return instance.SetProperty("MissedServerListRequests", value)
}

// GetMissedServerListRequests gets the value of MissedServerListRequests for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyMissedServerListRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("MissedServerListRequests")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerAnnounceAllocationsFailedPersec sets the value of ServerAnnounceAllocationsFailedPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyServerAnnounceAllocationsFailedPersec(value uint32) (err error) {
	return instance.SetProperty("ServerAnnounceAllocationsFailedPersec", value)
}

// GetServerAnnounceAllocationsFailedPersec gets the value of ServerAnnounceAllocationsFailedPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyServerAnnounceAllocationsFailedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ServerAnnounceAllocationsFailedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerListRequestsPersec sets the value of ServerListRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) SetPropertyServerListRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("ServerListRequestsPersec", value)
}

// GetServerListRequestsPersec gets the value of ServerListRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_PerfNet_Browser) GetPropertyServerListRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ServerListRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
