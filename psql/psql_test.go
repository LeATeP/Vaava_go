package psql

import (
	// "errors"
	"testing"
)

func TestDbConnection(t *testing.T) {
	tests := []struct {
		description string
		wantErr error
	}{
		{
		description: "Establishing Connection",
		wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_, err := Psql_connect()
			if err != nil {
				t.Errorf("failed Psql_connect, got error: %v", err.Error())
			}
		})
	}
}

func TestHello(t *testing.T) {
}
