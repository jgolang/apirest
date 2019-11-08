package apirest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/jgolang/log"
)

func TestRequestBasic_UnmarshalBody(t *testing.T) {
	type fields struct {
		JSONStruct interface{}
		SessionID  string
		UserID     string
		TraceID    string
		Tools      ToolBasic
	}
	type args struct {
		r *http.Request
		v interface{}
	}

	type Request struct {
		Name string `json:"nombre"`
	}

	request := Request{
		Name: "Josue Giron",
	}

	rawRequest, _ := json.Marshal(request)
	httpRequest, err := http.NewRequest("POST", "/test", bytes.NewBuffer(rawRequest))
	if err != nil {
		log.Error(err)
		return
	}

	test := Request{}
	var test2 string

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Response
	}{
		{
			name: "Test1",
			fields: fields{
				JSONStruct: test,
			},
			args: args{
				r: httpRequest,
				v: &test,
			},
			want: nil,
		},
		{
			name: "Error",
			fields: fields{
				JSONStruct: test,
			},
			args: args{
				r: httpRequest,
				v: &test2,
			},
			want: Error{Title: "Estructura JSON invalida", Message: "No se ha leído la estructura..."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := RequestBasic{
				JSONStruct: tt.fields.JSONStruct,
				SessionID:  tt.fields.SessionID,
				UserID:     tt.fields.UserID,
				TraceID:    tt.fields.TraceID,
				Tools:      tt.fields.Tools,
			}
			if got := request.UnmarshalBody(tt.args.r, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestBasic.UnmarshalBody() = %v, want %v", got, tt.want)

				t.Log(got)
				res := got.setResponse()
				log.Info(res)
			}
		})
	}
}

func TestRequestBasic_GetSessionInfo(t *testing.T) {
	type fields struct {
		JSONStruct interface{}
		SessionID  string
		UserID     string
		TraceID    string
		Tools      ToolBasic
	}
	type args struct {
		r *http.Request
	}

	type Req struct {
		Name string `json:"nombre"`
	}

	request := Req{
		Name: "Test",
	}

	rawRequest, _ := json.Marshal(request)
	httpRequest, err := http.NewRequest("POST", "/test", bytes.NewBuffer(rawRequest))
	if err != nil {
		log.Error(err)
		return
	}

	httpRequest.Header.Add("UserID", "45")
	httpRequest.Header.Add("SessionID", "AJJLADSJLQI3JL3JJDFSLJK2KJ3K")
	httpRequest.Header.Add("TraceID", "evento.test")

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Request
		want1  Response
	}{
		{
			name:   "Primer error",
			fields: fields{},
			args: args{
				r: httpRequest,
			},
			want: RequestBasic{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := RequestBasic{
				JSONStruct: tt.fields.JSONStruct,
				SessionID:  tt.fields.SessionID,
				UserID:     tt.fields.UserID,
				TraceID:    tt.fields.TraceID,
				Tools:      tt.fields.Tools,
			}
			got, got1 := request.GetSessionInfo(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestBasic.GetSessionInfo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RequestBasic.GetSessionInfo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
