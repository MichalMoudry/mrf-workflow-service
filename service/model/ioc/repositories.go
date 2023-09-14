package ioc

// An interface for a repository connected to Application structure.
type IApplicationRepository interface {
	AddApplication() error
}
