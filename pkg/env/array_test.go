package env

import (
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestGetArrayString(t *testing.T) {
	type args struct {
		key string
		sep string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				key: "HOSTS",
				sep: ",",
			},
			want:    []string{"tcp://admin:123", "tcp://localhost"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if !tt.wantErr {
			assert.Nil(t, os.Setenv(tt.args.key, strings.Join(tt.want, tt.args.sep)))
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := GetArrayString(tt.args.key, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArrayString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArrayString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
