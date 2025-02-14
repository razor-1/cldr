package cldr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOperandsString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *PluralOperands
		wantErr bool
	}{
		{
			name:  "integer 1",
			input: "1",
			want: &PluralOperands{
				N: 1,
				I: 1,
				V: 0,
				W: 0,
				F: 0,
				T: 0,
				C: 0,
			},
		},
		{
			name:  "decimal 1.0",
			input: "1.0",
			want: &PluralOperands{
				N: 1,
				I: 1,
				V: 1,
				W: 0,
				F: 0,
				T: 0,
				C: 0,
			},
		},
		{
			name:  "decimal 1.00",
			input: "1.00",
			want: &PluralOperands{
				N: 1,
				I: 1,
				V: 2,
				W: 0,
				F: 0,
				T: 0,
				C: 0,
			},
		},
		{
			name:  "decimal 1.3",
			input: "1.3",
			want: &PluralOperands{
				N: 1.3,
				I: 1,
				V: 1,
				W: 1,
				F: 3,
				T: 3,
				C: 0,
			},
		},
		{
			name:  "decimal 1.30",
			input: "1.30",
			want: &PluralOperands{
				N: 1.3,
				I: 1,
				V: 2,
				W: 1,
				F: 30,
				T: 3,
				C: 0,
			},
		},
		{
			name:  "decimal 1.03",
			input: "1.03",
			want: &PluralOperands{
				N: 1.03,
				I: 1,
				V: 2,
				W: 2,
				F: 3,
				T: 3,
				C: 0,
			},
		},
		{
			name:  "decimal 1.230",
			input: "1.230",
			want: &PluralOperands{
				N: 1.23,
				I: 1,
				V: 3,
				W: 2,
				F: 230,
				T: 23,
				C: 0,
			},
		},
		{
			name:  "large integer",
			input: "1200000",
			want: &PluralOperands{
				N: 1200000,
				I: 1200000,
				V: 0,
				W: 0,
				F: 0,
				T: 0,
				C: 0,
			},
		},
		{
			name:  "compact decimal 1.2c6",
			input: "1.2c6",
			want: &PluralOperands{
				N: 1200000,
				I: 1200000,
				V: 0,
				W: 0,
				F: 0,
				T: 0,
				C: 6,
			},
		},
		{
			name:  "compact decimal 123c6",
			input: "123c6",
			want: &PluralOperands{
				N: 123000000,
				I: 123000000,
				V: 0,
				W: 0,
				F: 0,
				T: 0,
				C: 6,
			},
		},
		{
			name:  "compact decimal 123c5",
			input: "123c5",
			want: &PluralOperands{
				N: 12300000,
				I: 12300000,
				V: 0,
				W: 0,
				F: 0,
				T: 0,
				C: 5,
			},
		},
		{
			name:  "decimal with fraction 1200.50",
			input: "1200.50",
			want: &PluralOperands{
				N: 1200.5,
				I: 1200,
				V: 2,
				W: 1,
				F: 50,
				T: 5,
				C: 0,
			},
		},
		{
			name:  "compact decimal with fraction 1.20050c3",
			input: "1.20050c3",
			want: &PluralOperands{
				N: 1200.5,
				I: 1200,
				V: 2,
				W: 1,
				F: 50,
				T: 5,
				C: 3,
			},
		},
		{
			name:    "invalid input",
			input:   "invalid",
			wantErr: true,
		},
		{
			name:    "invalid compact decimal",
			input:   "1.2cx",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newOperandsString(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
