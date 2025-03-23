package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessor_Process(t *testing.T) {
	type testCase struct {
		name             string
		key              string
		value            interface{}
		initMockStorage  func(*MockStorage)
		initMockNotifier func(*MockNotifier)
		expectedResult   interface{}
		expectedError    error
	}

	cases := []testCase{
		{
			name:             "empty key - validation error",
			key:              "",
			value:            "value_1",
			initMockStorage:  nil,
			initMockNotifier: nil,
			expectedError:    ErrInvalidInput,
		},
		{
			name:  "existing data - only storage.Get used",
			key:   "key_2",
			value: "value_2",
			initMockStorage: func(m *MockStorage) {
				m.EXPECT().Get("key_2").Return("value_2", nil)
			},
			initMockNotifier: nil,
			expectedResult:   "value_2",
		},
		{
			name:  "new data with nil value - storage.Get and storage.Set used",
			key:   "key_3",
			value: nil,
			initMockStorage: func(m *MockStorage) {
				m.EXPECT().Get("key_3").Return(nil, ErrDataNotFound)
				m.EXPECT().Set("key_3", nil)
			},
			initMockNotifier: nil,
			expectedResult:   nil,
		},
		{
			name:  "new data with value - all dependencies used",
			key:   "key_4",
			value: "value",
			initMockStorage: func(m *MockStorage) {
				m.EXPECT().Get("key_4").Return(nil, ErrDataNotFound)
				m.EXPECT().Set("key_4", "value")
			},
			initMockNotifier: func(m *MockNotifier) {
				m.EXPECT().Send("key_4")
			},
			expectedResult: "value",
		},
		{
			name: "storage error - only storage.Get used",
			key:  "key_5",
			initMockStorage: func(m *MockStorage) {
				m.EXPECT().Get("key_5").Return(nil, errors.New("some error"))
			},
			initMockNotifier: nil,
			expectedError:    errors.New("some error"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Инициализация моков
			mockStorage := NewMockStorage(ctrl)
			mockNotifier := NewMockNotifier(ctrl)

			// Настройка моков
			if tc.initMockStorage != nil {
				tc.initMockStorage(mockStorage)
			}

			if tc.initMockNotifier != nil {
				tc.initMockNotifier(mockNotifier)
			}

			// Создание процессора с моками
			processor := NewProcessor(mockStorage, mockNotifier)

			// Вызов тестируемого метода
			result, err := processor.Process(tc.key, tc.value)

			// Проверка результатов
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}

			// Проверка вызовов зависимостей
			//mockStorage.AssertExpectations(t)
			//mockNotifier.AssertExpectations(t)
		})
	}
}
