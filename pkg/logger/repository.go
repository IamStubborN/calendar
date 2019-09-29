package logger

type Repository interface {
	Info(...interface{})
	Warn(...interface{})
	Fatal(...interface{})
	WithFields(level string, data map[string]interface{}, msg ...interface{})
}
