package circuit_breaker

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsErrNotAllowed(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "not",
			args: args{err: errors.New("some err")},
			want: false,
		},
		{
			name: "just",
			args: args{err: errNotAllowed},
			want: true,
		},
		{
			name: "wrapped",
			args: args{fmt.Errorf("abc: %w", errNotAllowed)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsErrNotAllowed(tt.args.err); got != tt.want {
				t.Errorf("IsErrNotAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrNotAllowed(t *testing.T) {
	got := ErrNotAllowed()
	assert.ErrorIs(t, got, errNotAllowed)
}
