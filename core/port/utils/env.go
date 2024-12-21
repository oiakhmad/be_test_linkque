package utils

type Env interface {
	SanitizeEnv(envName string) (string, error)
	GetEnvString(envName string) (string, error)
	GetEnvBool(envName string) (bool, error)
	GetEnvInt(envName string) (int, error)
	GetEnvFloat32(envName string) (float32, error)
	GetEnvFloat64(envName string) (float64, error)
}
