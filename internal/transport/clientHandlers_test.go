package transport

import (
	"CarFix/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateReq(t *testing.T) {
	h := &ClientHandler{}
	in := &models.Client{
		Name:        "Petr",
		LastName:    "Petrov",
		PhoneNumber: "+7 952 745 81 94",
	}
	err := h.validateReq(in)

	require.NoError(t, err)
}

func TestValidateReqErr(t *testing.T) {
	cases := []struct {
		name   string
		in     *models.Client
		expErr error
	}{
		{
			name:   "empty_name",
			in:     &models.Client{LastName: "Petrov"},
			expErr: ErrEmptyName,
		},
		{
			name:   "empty_last_name",
			in:     &models.Client{Name: "Ivan"},
			expErr: ErrEmptyLastName,
		},
		{
			name:   "bad_name",
			in:     &models.Client{LastName: "Petrov", Name: "Ivan1"},
			expErr: ErrInvalidName,
		},
		{
			name:   "bad_last_name",
			in:     &models.Client{Name: "Ivan", LastName: "!Petrov"},
			expErr: ErrInvalidLastName,
		},
		// TODO: phone number testing
	}

	h := ClientHandler{}
	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := h.validateReq(tCase.in)
			require.Error(t, err)
			require.EqualError(t, tCase.expErr, err.Error())
		})
	}
}
