package errortrace

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestTrace(t *testing.T) {
	err := Errorf("%d", 42)

	if err == nil {
		t.Fatal("Failed to build error object.")
	}
	if tErr, ok := err.(tracedError); ok {
		if tErr.line != 10 || filepath.Base(tErr.file) != "trace_test.go" {
			t.Fatalf("Wrong wrapper built. Got %#v", tErr)
		}

		if tErr.err.Error() != "42" {
			t.Fatalf("Inner error corrupted. Got %v", tErr.err)
		}
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		expectedErrStr string
	}{
		{
			name:    "Nil returns nil",
			args:    args{nil},
			wantErr: false,
		},
		{
			name:           "Real err gets wrapped.",
			args:           args{errors.New("foobar")},
			wantErr:        true,
			expectedErrStr: "[trace_test.go:50] foobar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Wrap(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("Wrap() error = %v, wantErr %v", err, tt.wantErr)
			} else if tt.wantErr && err.Error() != tt.expectedErrStr {
				t.Errorf("Expected Error Mismatch. error = %v, expect = %v", err, tt.expectedErrStr)
			}
		})
	}
}
