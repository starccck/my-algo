package sort

import "testing"

func TestIQS_Sort(t *testing.T) {
	type args struct {
		data []MyIntType
	}
	tests := []struct {
		name string
		q    IQS
		args args
	}{
		// TODO: Add test cases.
		{
			name: "1",
			q:    IQS{},
			args: args{
				data: []MyIntType{1, 34, 2, 45, 23, 134, 65, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := IQS{}
			q.Sort(tt.args.data)
			t.Logf("data: %v\n", tt.args.data)
		})
	}
}
