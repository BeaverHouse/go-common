package env

// EnvType is the enum type for environment types
type EnvType string

const (
	// LocalEnv is the local development environment.
	LocalEnv EnvType = "local"
	// DevEnv is the remote development/staging environment.
	DevEnv EnvType = "development"
	// ProdEnv is the production environment.
	ProdEnv EnvType = "production"
)
