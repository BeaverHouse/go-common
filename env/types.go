package env

// EnvType is the enum type for environment types
type EnvType string

const (
	LocalEnv EnvType = "local"
	DevEnv   EnvType = "development"
	ProdEnv  EnvType = "production"
)
