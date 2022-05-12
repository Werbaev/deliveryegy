package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/werbaev/deliveryegy/api/models"
	"github.com/werbaev/deliveryegy/config"
	"github.com/werbaev/deliveryegy/pkg/logger"
	"github.com/werbaev/deliveryegy/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handlerV1 struct {
	log     logger.Logger
	cfg     *config.Config
	storage storage.StorageI
}

type HandlerV1Options struct {
	Log     logger.Logger
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		log:     options.Log,
		cfg:     options.Cfg,
		storage: options.Storage,
	}
}

const (
	//ErrorCodeInvalidURL ...
	ErrorCodeInvalidURL = "INVALID_URL"
	//ErrorCodeInvalidJSON ...
	ErrorCodeInvalidJSON = "INVALID_JSON"
	//ErrorCodeInternal ...
	ErrorCodeInternal = "INTERNAL_SERVER_ERROR"
	//ErrorCodeUnauthorized ...
	ErrorCodeUnauthorized = "UNAUTHORIZED"
	//ErrorCodeAlreadyExists ...
	ErrorCodeAlreadyExists = "ALREADY_EXISTS"
	//ErrorCodeNotFound ...
	ErrorCodeNotFound = "NOT_FOUND"
	//ErrorCodeInvalidCode ...
	ErrorCodeInvalidCode = "INVALID_CODE"
	//ErrorBadRequest ...
	ErrorBadRequest = "BAD_REQUEST"
	//ErrorCodeForbidden ...
	ErrorCodeForbidden = "FORBIDDEN"
	//ErrorCodeNotApproved ...
	ErrorCodeNotApproved       = "NOT_APPROVED"
	ErrorCodePasswordsNotEqual = "PASSWORDS_NOT_EQUAL"
	ErrServiceUnavailable      = "SERVICE_UNAVAILABLE"
	ErrInvalidArgument         = "INVALID_ARGUMENT"
)

var (
	signingKey = []byte("FfLbN7pIEYe8@!EqrttOLiwa(H8)voxe")
)

func ParseQueryParam(c *gin.Context, key string, defaultValue string) (int, error) {
	valueStr := c.DefaultQuery(key, defaultValue)

	value, err := strconv.Atoi(valueStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return 0, err
	}

	return value, nil
}

//ParsePageQueryParam ...
func ParsePageQueryParam(c *gin.Context) (uint64, error) {
	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 10)
	if err != nil {
		return 0, err
	}
	if page < 0 {
		return 0, errors.New("page must be an positive integer")
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

//ParseLimitQueryParam ...
func ParseLimitQueryParam(c *gin.Context) (uint64, error) {
	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 10)
	if err != nil {
		return 0, err
	}
	if limit < 0 {
		return 0, errors.New("page_size must be an positive integer")
	}
	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}

func (h *handlerV1) handleError(c *gin.Context, err error, message string) bool {
	st, ok := status.FromError(err)

	switch st.Code() {
	case codes.Internal:
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternal,
			Message: st.Message(),
		})
		h.log.Error(message+", internal server error", logger.Error(err))

		return true

	case codes.Unavailable:
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrServiceUnavailable,
			Message: st.Message(),
		})
		h.log.Error(message+", service unavailable", logger.Error(err))

		return true

	case codes.AlreadyExists:
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeAlreadyExists,
			Message: st.Message(),
		})
		h.log.Error(message+", already exists", logger.Error(err))

		return true

	case codes.InvalidArgument:
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    ErrInvalidArgument,
			Message: st.Message(),
		})
		h.log.Error(message+", invalid field", logger.Error(err))

		return true

	case codes.NotFound:
		c.JSON(http.StatusNotFound, models.ResponseError{
			Code:    ErrorCodeNotFound,
			Message: st.Message(),
		})
		h.log.Error(message+", not found", logger.Error(err))

		return true

	default:
		if err != nil || !ok {
			c.JSON(http.StatusInternalServerError, models.ResponseError{
				Code:    ErrorCodeInternal,
				Message: st.Message(),
			})
			h.log.Error(message+", unknown error", logger.Error(err))
			return true
		}

	}

	return false
}

func handleInternalWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternal,
			Message: "Internal Server Error",
		})
		l.Error(message, logger.Error(err))
		return true
	}

	return false
}

func handleBadRequestErrWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    ErrorCodeInvalidJSON,
			Message: "Invalid Json",
		})
		l.Error(message, logger.Error(err))
		return true
	}
	return false
}

func (h *handlerV1) ResponseProtoJson(c *gin.Context, p proto.Message) {
	var (
		jsonbMarshal jsonpb.Marshaler
	)

	jsonbMarshal.OrigName = true
	jsonbMarshal.EmitDefaults = true

	js, err := jsonbMarshal.MarshalToString(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    models.ErrorCodeInvalidJSON,
			Message: err.Error(),
		})

		h.log.Error("marshal proto", logger.Error(err))

		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
