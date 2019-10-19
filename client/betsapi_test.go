package client

import (
    "github.com/heat/syscrapgo/eventosrv"
    "testing"
)

func Test_betsApiClientGet(t *testing.T) {

    type fields struct {
        URL string
        Key string
    }
    type args struct {
        source string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    *eventosrv.Evento
        wantErr bool
    }{
        {
            name: "Teste client request",
            fields: fields{
              URL: "https://api.betsapi.com/v1",
              Key: "27014-b8AivmL87jt2Mk",
            },
            args: args{
                source: "83218706",
            },
            want: &eventosrv.Evento{
                ID:   "bbshaaip83218706",
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            api := &betsApiClient{
                URL: tt.fields.URL,
                Key: tt.fields.Key,
            }
            got, err := api.Get(tt.args.source)
            if (err != nil) != tt.wantErr {
                t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got.ID != tt.want.ID {
                t.Errorf("Get() got = %v, want %v", got.ID, tt.want.ID)
            }
        })
    }
}
