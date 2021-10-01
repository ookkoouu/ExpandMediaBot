package handler

import "testing"

func TestGetIdFromUrl(t *testing.T){
	type args struct{
		url string
	}
	tests := []struct{
		name string
		args args
		want int64
	}{
		{
			name: "normal",
			args: args{url: "https://twitter.com/FloodSocial/status/861627479294746624"},
			want: 861627479294746624,
		},
		{
			name: "query",
			args: args{url: "https://twitter.com/i/statuses/861627479294746624?s=20"},
			want: 861627479294746624,
		},
		{
			name: "short",
			args: args{url: "http://twitter.com/status/861627479294746624?s=afw"},
			want: 861627479294746624,
		},
	}

	for _, tt :=range tests{
		t.Run(tt.name, func(t *testing.T) {
			if got := getIdFromUrl(tt.args.url); got != tt.want{
				t.Errorf("getIdFromUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}