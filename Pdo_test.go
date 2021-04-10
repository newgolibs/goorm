package goorm

import (
	"database/sql"
	"testing"
)

func TestPdo_Commit(t *testing.T) {
	p := NewPdoconfigFromBytes([]byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`))
	pdoconfigMiddleware := p.NewPdoconfigMiddleware()
	pdoconfigMiddleware.MakeDbPool()
	type fields struct {
		TX        *sql.Tx
		Pdoconfig *PdoconfigMiddleware
	}
	type args struct {
		recover interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "异常，测试错误提示",
			fields: fields{
				TX:        pdoconfigMiddleware.MakeTX(),
				Pdoconfig: pdoconfigMiddleware,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Pdo{
				TX:        tt.fields.TX,
				Pdoconfig: tt.fields.Pdoconfig,
			}
			this.Commit("errdata")
		})
	}
}
