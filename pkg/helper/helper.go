package helper

import (
	"database/sql"
	"math"

	"github.com/anmimos/delivery/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleError(log logger.Logger, err error, message string, req interface{}) error {
	if err == sql.ErrNoRows {
		log.Error(message+", Not Found", logger.Error(err), logger.Any("req", req))
		return status.Error(codes.NotFound, "Not Found")
	} else if err != nil {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return status.Error(codes.Internal, message)
	}
	return nil
}

func RoundFloat2DecimalPrecison(f float64) float64 {
	return math.Floor(f*100) / 100
}

func RoundFloat4DecimalPrecison(f float64) float64 {
	return math.Floor(f*10000) / 10000
}
