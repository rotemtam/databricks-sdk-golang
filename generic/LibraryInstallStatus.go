package generic

type LibraryInstallStatus string

const (
	PENDING              = "PENDING"
	RESOLVING            = "RESOLVING"
	INSTALLING           = "INSTALLING"
	INSTALLED            = "INSTALLED"
	FAILED               = "FAILED"
	UNINSTALL_ON_RESTART = "UNINSTALL_ON_RESTART"
)
