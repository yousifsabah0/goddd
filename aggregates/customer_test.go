package aggregates_test

import (
	"testing"

	"github.com/yousifsabah0/goddd/aggregates"
)

type testCase struct {
	test        string
	name        string
	age         int
	expectedErr error
}

func TestNewCustomer(t *testing.T) {
	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			age:         19,
			expectedErr: aggregates.ErrInvalidName,
		},
		{
			test:        "Empty Age validation",
			name:        "",
			age:         0,
			expectedErr: aggregates.ErrInvalidAge,
		},
		{
			test:        "Valid Name",
			name:        "Percy Bolmer",
			age:         18,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregates.NewCustomer(tc.name, tc.age)
			if err != tc.expectedErr {
				t.Errorf("expected %v; got %v", tc.expectedErr, err)
			}
		})
	}

}
