package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mahiro72/go-mysql-ca/domain/entity"
)

func TestHandler_GetAllTask(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
		wantBody string
	}{
		{
			name:     "正常に動作した場合",
			wantCode: http.StatusOK,
			wantBody: `{"data":[{"id":1,"name":"hoge","status":"todo"},{"id":1,"name":"hoge","status":"todo"},{"id":1,"name":"hoge","status":"todo"}]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.Default()

			taskHandler := NewTaskHandler(&TaskUsecaseMock{})
			router.GET("/task/all", taskHandler.GetAllTask)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/task/all", nil)
			router.ServeHTTP(w, req)

			if tt.wantCode != w.Code {
				t.Errorf("TestHandler_GetAllTask code Error : want %d but got %d", tt.wantCode, w.Code)
			}
			if tt.wantBody != w.Body.String() {
				t.Errorf("TestHandler_GetAllTask body Error : want %s but got %s", tt.wantBody, w.Body.String())
			}
		})
	}
}

func TestHandler_CreateTask(t *testing.T) {
	tests := []struct {
		name     string
		reqBody  string
		wantCode int
		wantBody string
	}{
		{
			name:     "正常に動作した場合",
			reqBody:  `{"name":"hoge"}`,
			wantCode: http.StatusOK,
			wantBody: `{"data":{"id":0,"name":"hoge","status":""}}`,
		},
		{
			name:     "request bodyが空だった場合、400エラーになる",
			reqBody:  ``,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"TaskHandler.CreateTask ShouldBindJSON Error : EOF"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.Default()

			taskHandler := NewTaskHandler(&TaskUsecaseMock{})
			router.POST("/task/create", taskHandler.CreateTask)

			body := bytes.NewBufferString(tt.reqBody)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/task/create", body)
			router.ServeHTTP(w, req)

			if tt.wantCode != w.Code {
				t.Errorf("TestHandler_CreateTask code Error : want %d but got %d", tt.wantCode, w.Code)
			}
			if tt.wantBody != w.Body.String() {
				t.Errorf("TestHandler_CreateTask body Error : want %s but got %s", tt.wantBody, w.Body.String())
			}
		})
	}
}

func TestHandler_ChangeTaskStatus(t *testing.T) {
	tests := []struct {
		name     string
		reqBody  string
		wantCode int
		wantBody string
	}{
		{
			name:     "正常に動作した場合",
			reqBody:  `{"id":1,"status":"done"}`,
			wantCode: http.StatusOK,
			wantBody: `{"data":{"id":0,"name":"","status":"done"}}`,
		},
		{
			name:     "request bodyが空だった場合、400エラーになる",
			reqBody:  ``,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"TaskHandler.ChangeTaskStatus ShouldBindJSON Error : EOF"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.Default()

			taskHandler := NewTaskHandler(&TaskUsecaseMock{})
			router.PUT("/task/status", taskHandler.ChangeTaskStatus)

			body := bytes.NewBufferString(tt.reqBody)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPut, "/task/status", body)
			router.ServeHTTP(w, req)

			if tt.wantCode != w.Code {
				t.Errorf("TestHandler_CreateTask code Error : want %d but got %d", tt.wantCode, w.Code)
			}
			if tt.wantBody != w.Body.String() {
				t.Errorf("TestHandler_CreateTask body Error : want %s but got %s", tt.wantBody, w.Body.String())
			}
		})
	}
}

// ここでテストデータの用意をします
var (
	testGetAllTask []*entity.Task = []*entity.Task{
		{Id: 1, Name: "hoge", Status: "todo"},
		{Id: 1, Name: "hoge", Status: "todo"},
		{Id: 1, Name: "hoge", Status: "todo"},
	}
	testCreateTask *entity.Task = &entity.Task{
		Id:     0,
		Name:   "hoge",
		Status: "",
	}
	testChangeTaskStatus *entity.Task = &entity.Task{
		Id:     0,
		Name:   "",
		Status: "done",
	}
)

type TaskUsecaseMock struct{}

func (h *TaskUsecaseMock) GetAllTask() (entity.Tasks, error) {
	return testGetAllTask, nil
}

func (h *TaskUsecaseMock) CreateTask(task *entity.Task) (*entity.Task, error) {
	return testCreateTask, nil
}

func (h *TaskUsecaseMock) ChangeTaskStatus(task *entity.Task) (*entity.Task, error) {
	return testChangeTaskStatus, nil
}
