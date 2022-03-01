package db

import (
	"context"
	"testing"

	"github.com/Tanej98/minibank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		OwnerName: util.RandomName(),
		Balance:   util.RandomMoney(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.OwnerName, arg.OwnerName)
	require.Equal(t, account.Balance, arg.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
