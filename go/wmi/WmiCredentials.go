package wmi

// WmiCredentials
type WmiCredentials struct {
	UserName string;
	Password string;
	Domain string;
}

// GetSecureString
func (cred WmiCredentials) GetSecureString() (string, error) {
	panic ("not implemented");
}
