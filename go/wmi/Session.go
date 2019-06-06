package wmi

// Session
type Session interface {
	Connect(serverName, wmiNamespace, userName, password, domain string) (bool, error)
	ConnectEx(serverName, wmiNamespace string, credentials Credentials) (bool, error)
	Dispose()
}
