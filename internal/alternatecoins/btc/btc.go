package btc

import (
    "github.com/btcsuite/btcd/rpcclient"
    "log"
    "net/http"
)

// TODO(therealssj): finish rpc setup
// TODO(therealssj): finish function bodies

type BTC struct {
    rpc *rpcclient.Client
}

func New() *BTC {
    connCfg := &rpcclient.ConnConfig{
        Host:         "localhost:8332",
        User:         "yourrpcuser",
        Pass:         "yourrpcpass",
        HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
        DisableTLS:   true, // Bitcoin core does not provide TLS by default
    }

    // Notice the notification parameter is nil since notifications are
    // not supported in HTTP POST mode.
    client, err := rpcclient.New(connCfg, nil)
    if err != nil {
        log.Fatal(err)
    }

    return &BTC{
        rpc: client,
    }
}

func (b *BTC) CreateWallet() http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {

    }
}

func (b *BTC) RecoverWallet() http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {

    }
}

func (b *BTC) CreateTransaction() http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {

    }
}

func (b *BTC) SendTransaction() http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {

    }
}

func (b *BTC) SendRawTransaction() http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {

    }
}

func (b *BTC) SignTransaction() http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {

    }
}

func (b *BTC) GetTransaction() http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {

    }
}

func (b *BTC) GetBalance() http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {

    }
}

func (b *BTC) GetFeeEstimate() http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {

    }
}

func (b *BTC) NewAddresses() http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {

    }
}
