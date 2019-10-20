package client

import (
    "errors"
    "fmt"
    "github.com/heat/syscrapgo/cotacaosrv"
    "github.com/heat/syscrapgo/eventosrv"
    "log"
)

type ApiClient interface {
    Get(source string) (*eventosrv.Evento, error)
    Result(source string) (*eventosrv.ScoreResult, error)
    // GetOddsRaw consulta as odd de forma crua para usa possiveo transformacaos o -> T -> o'
    GetOddsRaw(source string) (*cotacaosrv.Odds, error )
}

type ApiClientOptions struct {
    URL string
    Key string
}

func MakeApiClient(kind string, logger *log.Logger, conf *ApiClientOptions) (ApiClient, error) {
    switch kind {
    case "betsapi":
        return &betsApiClient{
            URL: conf.URL,
            Key: conf.Key,
            logger: logger,
        }, nil
    }

    return nil, errors.New(fmt.Sprintf("kind=%s desconhecido", kind))
}
