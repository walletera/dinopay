package tests

import (
    "context"
    "github.com/google/uuid"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
    "github.com/walletera/dinopay/api"
    "net/http/httptest"
    "testing"
)

func TestClient(t *testing.T) {
    paymentID, err := uuid.NewUUID()
    require.NoError(t, err)

    payment := &api.Payment{
        ID: api.OptUUID{
            Value: paymentID,
            Set:   true,
        },
        Amount:                0,
        Currency:              "",
        SourceAccount:         api.Account{},
        DestinationAccount:    api.Account{},
        Status:                api.OptPaymentStatus{},
        CustomerTransactionId: api.OptString{},
        CreatedAt:             api.OptDate{},
        UpdatedAt:             api.OptDate{},
    }

    handlerMock := NewMockHandler(t)
    handlerMock.EXPECT().CreatePayment(mock.Anything, mock.Anything).Return(
        payment, nil)

    dinopayServer, err := api.NewServer(handlerMock)
    require.NoError(t, err)

    ts := httptest.NewServer(dinopayServer)
    defer ts.Close()

    dinopayClient, err := api.NewClient(ts.URL)
    require.NoError(t, err)

    res, err := dinopayClient.CreatePayment(context.Background(), payment)
    require.NoError(t, err)

    resPayment, ok := res.(*api.Payment)
    require.True(t, ok)

    require.Equal(t, payment.ID, resPayment.ID)
}
