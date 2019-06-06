package wmi;

type WmiSession interface {
	Connect(serverName, wmiNamespace, userName, password, domain string) (bool, error);
	Connect(serverName, wmiNamespace string, credentials WmiCredentials) (bool, error);
	Dispose();
}
