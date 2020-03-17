package alternatecoins

import (
    "fmt"
    "github.com/SkycoinProject/skycoin/src/api"
    "github.com/SkycoinProject/skycoin/src/util/logging"
    "net/http"
)

type Ticker string

var (
    logger = logging.MustGetLogger("altcoins")
)

// AltManager is a manager for altcoins
type AltManager struct {
    Alts map[Ticker]AltCoin
}

// NewAltManager constructs new manager according to the config
func NewAltManager(alts map[Ticker]AltCoin) (*AltManager, error) {
    logger.Info("Creating new KVStorage manager")

    m := &AltManager{
        Alts: alts,
    }

    return m, nil
}

func (am *AltManager) SetupAltcoinRoutes(prefix string, webHandler func(endpoint string, handler http.Handler, methodAPISets map[string][]string)) {
    // TODO(therealssj): add all routes
    for ticker, alt := range am.Alts {
        webHandler(fmt.Sprintf("%s/%s/getbalance", prefix, ticker), alt.GetBalance(), map[string][]string{
            http.MethodGet: []string{api.EndpointsRead},
        })
    }
}
