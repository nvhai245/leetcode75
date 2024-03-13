package numberofrecentcalls

import "testing"

func TestRecentCounter_Ping(t *testing.T) {
	type fields struct {
		reqs []int
	}
	type args struct {
		t int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"1", fields{reqs: []int{}}, args{1}, 1},
		{"2", fields{reqs: []int{1}}, args{100}, 2},
		{"3", fields{reqs: []int{1, 100, 3001, 3002}}, args{3003}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RecentCounter{
				reqs: tt.fields.reqs,
			}
			if got := r.Ping(tt.args.t); got != tt.want {
				t.Errorf("RecentCounter.Ping() = %v, want %v", got, tt.want)
			}
		})
	}
}
