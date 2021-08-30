package logx

import "go.uber.org/zap"

var sugar *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	if sugar != nil {
		return sugar
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	sugar = logger.Sugar()

	return sugar
}
