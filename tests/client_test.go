package tests

import (
    "context"
    "github.com/google/uuid"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
    "github.com/walletera/dinopay/api"
    "net/http/httptest"
    "testing"
    "time"
)

func TestClient_CreatePayment_Succeed(t *testing.T) {
    customerTransactionId, err := uuid.NewUUID()
    require.NoError(t, err)

    payment := &api.Payment{
        Amount:   100.5,
        Currency: "USD",
        SourceAccount: api.Account{
            AccountHolder: "John Doe",
            AccountNumber: "9A8SD7F98AS7D9F87",
        },
        DestinationAccount: api.Account{
            AccountHolder: "Jane Doe",
            AccountNumber: "A0S9D8F0A98SDF0A9",
        },
        Status: api.OptPaymentStatus{},
        CustomerTransactionId: api.OptString{
            Value: customerTransactionId.String(),
            Set:   true,
        },
        CreatedAt: api.OptDate{},
        UpdatedAt: api.OptDate{},
    }

    responsePaymentID, err := uuid.NewUUID()
    require.NoError(t, err)

    handlerMock := NewMockHandler(t)
    handlerMock.EXPECT().CreatePayment(mock.Anything, mock.Anything).Return(
        &api.Payment{
            ID: api.OptUUID{
                Value: responsePaymentID,
                Set:   true,
            },
            Amount:   payment.Amount,
            Currency: payment.Currency,
            SourceAccount: api.Account{
                AccountHolder: payment.SourceAccount.AccountHolder,
                AccountNumber: payment.SourceAccount.AccountNumber,
            },
            DestinationAccount: api.Account{
                AccountHolder: payment.DestinationAccount.AccountHolder,
                AccountNumber: payment.DestinationAccount.AccountNumber,
            },
            Status: api.OptPaymentStatus{
                Value: api.PaymentStatusConfirmed,
                Set:   true,
            },
            CustomerTransactionId: api.OptString{
                Value: customerTransactionId.String(),
                Set:   true,
            },
            CreatedAt: api.OptDate{
                Value: time.Now(),
                Set:   true,
            },
            UpdatedAt: api.OptDate{
                Value: time.Now(),
                Set:   true,
            },
        }, nil)

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

    require.Equal(t, responsePaymentID, resPayment.ID.Value)
    require.Equal(t, payment.CustomerTransactionId.Value, resPayment.CustomerTransactionId.Value)
    require.Equal(t, api.PaymentStatusConfirmed, resPayment.Status.Value)
}
