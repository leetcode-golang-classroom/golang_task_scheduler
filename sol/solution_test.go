package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	for idx := 0; idx < b.N; idx++ {
		leastInterval(tasks, n)
	}
}
func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tasks = [\"A\",\"A\",\"A\",\"B\",\"B\",\"B\"], n = 2",
			args: args{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 2},
			want: 8,
		},
		{
			name: "tasks = [\"A\",\"A\",\"A\",\"B\",\"B\",\"B\"], n = 0",
			args: args{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 0},
			want: 6,
		},
		{
			name: "tasks = [\"A\",\"A\",\"A\",\"A\",\"A\",\"A\",\"B\",\"C\",\"D\",\"E\",\"F\",\"G\"], n = 2",
			args: args{tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, n: 2},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
