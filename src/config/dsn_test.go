package config

import (
	"testing"
	"fmt"
)

func TestDsn_Dsn(t *testing.T){
	tests := []struct{
		name string
		wantDsn string
		wantErr error
	}{
		{
			name: "必要な環境変数がすべて設定されている場合、DSNを返す",
			wantDsn: "hoge:hoge@tcp(hoge:hoge)/hoge?parseTime=true&collation=utf8mb4_bin",
			wantErr: nil,
		},
	}

	for _,tt := range tests {
		t.Run(tt.name,func(t *testing.T) {
			t.Setenv("DB_USER", "hoge")
			t.Setenv("DB_PASSWORD", "hoge")
			t.Setenv("DB_HOST", "hoge")
			t.Setenv("DB_PORT", "hoge")
			t.Setenv("DB_DATABASE", "hoge")

			gotDSN,gotErr := DSN()
			if tt.wantDsn != gotDSN {
				t.Errorf("TestDsn_Dsn Error : wantDsb %s, but gotDsb %s", tt.wantDsn, gotDSN)
			}

			if tt.wantErr != gotErr{
				t.Errorf("TestDsn_Dsn Error : wantErr %s, but gotErr %s", tt.wantErr, gotErr)
			}
		})
	}
}

func TestDsn_DsnNotFound(t *testing.T){
	tests := []struct{
		name string
		wantDsn string
		wantErr error
	}{
		{
			name: "必要な環境変数が1つでも設定されていない場合、errorを返す",
			wantDsn: "",
			wantErr: fmt.Errorf("ERROR : required environment variable not found"),
		},
	}

	for _,tt := range tests {
		t.Run(tt.name,func(t *testing.T) {
			gotDSN,gotErr := DSN()
			if tt.wantDsn != gotDSN {
				t.Errorf("TestDsn_DsnNotFound Error : wantDsb %s, but gotDsb %s", tt.wantDsn, gotDSN)
			}

			if tt.wantErr.Error() != gotErr.Error() {
				t.Errorf("TestDsn_DsnNotFound Error : wantErr %s, but gotErr %s", tt.wantErr, gotErr)
			}
		})
	}
}