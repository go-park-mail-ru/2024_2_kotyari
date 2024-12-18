package address

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/address/mocks"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const testContextRequestIDKey = "request-id"

var (
	testContextRequestIDValue = uuid.New()
)

func TestAddressDelivery_GetAddress(t *testing.T) {
	type want struct {
		statusCode int
		response   string
	}

	tests := []struct {
		name      string
		setupFunc func(ctrl *gomock.Controller) *AddressDelivery
		request   *http.Request
		want      want
	}{
		{
			name: "Успешное выполнение",
			setupFunc: func(ctrl *gomock.Controller) *AddressDelivery {
				addressManager := mocks.NewMockaddressManager(ctrl)

				addressManager.EXPECT().
					GetAddressByProfileID(gomock.Any(), uint32(1)).
					Return(model.Address{Text: "Main St"}, nil)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))
				errResolver := errs.NewErrorStore()
				return &AddressDelivery{
					addressManager: addressManager,
					log:            logger,
					errResolver:    errResolver,
				}
			},
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodGet, "/address", nil)
				ctx := utils.SetContextSessionUserID(req.Context(), 1)
				return req.WithContext(ctx)
			}(),
			want: want{
				statusCode: http.StatusOK,
				response:   `{"body":{"address":"Main St"}, "status":200}`,
			},
		},
		{
			name: "Пользователь не авторизован",
			setupFunc: func(ctrl *gomock.Controller) *AddressDelivery {
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))
				mockAddressManager := mocks.NewMockaddressManager(ctrl)

				//mockAddressManager.
				//	EXPECT().
				//	GetAddressByProfileID(gomock.Any(), gomock.Any()).
				//	Times(1).
				//	Return(model.Address{}, errs.UserNotAuthorized)

				return &AddressDelivery{
					log:            logger,
					errResolver:    errs.NewErrorStore(),
					addressManager: mockAddressManager,
				}
			},
			request: httptest.NewRequest(http.MethodGet, "/address", nil),
			want: want{
				statusCode: http.StatusUnauthorized,
				response:   `{"body":{"error_message":"Пользователь не авторизован"}, "status":401}`,
			},
		},

		{
			name: "Ошибка при получении адреса",
			setupFunc: func(ctrl *gomock.Controller) *AddressDelivery {
				addressManager := mocks.NewMockaddressManager(ctrl)

				addressManager.EXPECT().
					GetAddressByProfileID(gomock.Any(), uint32(1)).
					Return(model.Address{}, errors.New("internal error"))

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))
				return &AddressDelivery{
					addressManager: addressManager,
					log:            logger,
					errResolver:    errs.NewErrorStore(),
				}
			},
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodGet, "/address", nil)
				ctx := utils.SetContextSessionUserID(req.Context(), 1)
				return req.WithContext(ctx)
			}(),
			want: want{
				statusCode: http.StatusInternalServerError,
				response:   `{"body":{"error_message":"внутренняя ошибка сервера"}, "status":500}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			addressDelivery := tt.setupFunc(ctrl)
			recorder := httptest.NewRecorder()

			ctx := context.WithValue(tt.request.Context(), testContextRequestIDKey, testContextRequestIDValue)

			reqWithCtx := tt.request.WithContext(ctx)

			addressDelivery.GetAddress(recorder, reqWithCtx)

			assert.Equal(t, tt.want.statusCode, recorder.Code)
			assert.JSONEq(t, tt.want.response, recorder.Body.String())
		})
	}
}

func TestAddressDelivery_UpdateAddressData(t *testing.T) {
	type want struct {
		statusCode int
		response   string
	}

	tests := []struct {
		name      string
		setupFunc func(ctrl *gomock.Controller) *AddressDelivery
		request   *http.Request
		want      want
	}{
		{
			name: "Успешное выполнение",
			setupFunc: func(ctrl *gomock.Controller) *AddressDelivery {
				addressManager := mocks.NewMockaddressManager(ctrl)
				addressManager.EXPECT().
					UpdateUsersAddress(gomock.Any(), uint32(1), gomock.Any()).
					Return(nil)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))
				errResolver := errs.NewErrorStore()
				return &AddressDelivery{
					addressManager: addressManager,
					log:            logger,
					errResolver:    errResolver,
				}
			},
			request: func() *http.Request {
				body := `{"address":"New StreetNew City"}`
				req := httptest.NewRequest(http.MethodPost, "/address/update", io.NopCloser(strings.NewReader(body)))
				ctx := utils.SetContextSessionUserID(req.Context(), 1)
				return req.WithContext(ctx)
			}(),
			want: want{
				statusCode: http.StatusOK,
				response:   `{"body":{"address":"New StreetNew City"}, "status":200}`,
			},
		},
		{
			name: "Пользователь не авторизован",
			setupFunc: func(ctrl *gomock.Controller) *AddressDelivery {
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))
				return &AddressDelivery{
					log:            logger,
					errResolver:    errs.NewErrorStore(),
					addressManager: mocks.NewMockaddressManager(ctrl),
				}
			},
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/address/update", nil)
				return req // Не нужно добавлять сессию, чтобы эмулировать неавторизованного пользователя
			}(),
			want: want{
				statusCode: http.StatusUnauthorized,
				response:   `{"body":{"error_message":"Пользователь не авторизован"}, "status":401}`,
			},
		},

		{
			name: "Ошибка десериализации JSON",
			setupFunc: func(ctrl *gomock.Controller) *AddressDelivery {
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))
				errResolver := errs.NewErrorStore()
				return &AddressDelivery{
					log:            logger,
					errResolver:    errResolver,
					addressManager: mocks.NewMockaddressManager(ctrl),
				}
			},
			request: func() *http.Request {
				body := `{"invalid_json":}`
				req := httptest.NewRequest(http.MethodPost, "/address/update", io.NopCloser(strings.NewReader(body)))
				ctx := utils.SetContextSessionUserID(req.Context(), 1)
				return req.WithContext(ctx)
			}(),
			want: want{
				statusCode: http.StatusBadRequest,
				response:   `{"body":{"error_message":"Неверный формат JSON"}, "status":400}`,
			},
		},
		{
			name: "Ошибка на уровне менеджера",
			setupFunc: func(ctrl *gomock.Controller) *AddressDelivery {
				addressManager := mocks.NewMockaddressManager(ctrl)
				addressManager.EXPECT().
					UpdateUsersAddress(gomock.Any(), uint32(1), gomock.Any()).
					Return(errors.New("internal error"))

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))
				errResolver := errs.NewErrorStore()
				return &AddressDelivery{
					addressManager: addressManager,
					log:            logger,
					errResolver:    errResolver,
				}
			},
			request: func() *http.Request {
				body := `{"street":"New Street", "city":"New City"}`
				req := httptest.NewRequest(http.MethodPost, "/address/update", io.NopCloser(strings.NewReader(body)))
				ctx := utils.SetContextSessionUserID(req.Context(), 1)
				return req.WithContext(ctx)
			}(),
			want: want{
				statusCode: http.StatusInternalServerError,
				response:   `{"body":{"error_message":"Внутренняя ошибка сервера"}, "status":500}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			addressDelivery := tt.setupFunc(ctrl)
			recorder := httptest.NewRecorder()

			ctx := context.WithValue(tt.request.Context(), testContextRequestIDKey, testContextRequestIDValue)

			reqWithCtx := tt.request.WithContext(ctx)

			addressDelivery.UpdateAddressData(recorder, reqWithCtx)

			assert.Equal(t, tt.want.statusCode, recorder.Code)
			assert.JSONEq(t, tt.want.response, recorder.Body.String())
		})
	}
}
