package boostrap

import "go.uber.org/zap"

func newLogger() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	return logger.Sugar(), nil
}
