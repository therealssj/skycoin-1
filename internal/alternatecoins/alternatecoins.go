package alternatecoins

import "net/http"

// TODO(therealssj): define inputs
type AltCoin interface {
    CreateWallet() http.HandlerFunc
    RecoverWallet() http.HandlerFunc
    CreateTransaction() http.HandlerFunc
    SendTransaction() http.HandlerFunc
    SendRawTransaction() http.HandlerFunc
    SignTransaction() http.HandlerFunc
    GetTransaction() http.HandlerFunc
    GetBalance() http.HandlerFunc
    GetFeeEstimate() http.HandlerFunc
    NewAddresses() http.HandlerFunc
}
