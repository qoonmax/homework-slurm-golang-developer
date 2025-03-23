package main

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestSendRequest(t *testing.T) {
	testCases := map[string]struct {
		url             string
		initHttpClient  func(m *MockHttpClient)
		expectedSuccess bool
		expectedErr     error
	}{
		"200": {
			url: "https://google.com",
			initHttpClient: func(m *MockHttpClient) {
				m.EXPECT().Get("https://google.com").Return(&http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(`{"message": "ok"}`)), // тело ответа
				}, nil)
			},
			expectedSuccess: true,
			expectedErr:     nil,
		},
		"404": {
			url: "https://google.com/404",
			initHttpClient: func(m *MockHttpClient) {
				m.EXPECT().Get("https://google.com/404").Return(&http.Response{
					StatusCode: 404,
					Body:       io.NopCloser(strings.NewReader(`{"message": "page not found"}`)), // тело ответа
				}, nil)
			},
			expectedSuccess: false,
			expectedErr:     ErrUnexpectedStatusCode,
		},
		"err_request_failed": {
			url: "https://domainthatdoesnotexist.io",
			initHttpClient: func(m *MockHttpClient) {
				m.EXPECT().Get("https://domainthatdoesnotexist.io").Return(nil, ErrRequestFailed)
			},
			expectedSuccess: false,
			expectedErr:     ErrRequestFailed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			client := NewMockHttpClient(ctrl)
			if tc.initHttpClient != nil {
				tc.initHttpClient(client)
			}

			success, err := sendRequest(client, tc.url)

			require.Equal(t, tc.expectedSuccess, success, "they should be equal")
			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
