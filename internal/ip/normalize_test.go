package ip

import "testing"

func TestNormalize(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid unnecessary normalize",
			args: args{
				host: "127.0.0.1",
			},
			want:    "127.0.0.1",
			wantErr: false,
		},
		{
			name: "valid necessary normalize",
			args: args{
				host: "127.0.0.1:3000",
			},
			want:    "127.0.0.1",
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				host: ":::::::",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Normalize(tt.args.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
