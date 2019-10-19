package client

import (
    "errors"
    "fmt"
    "github.com/go-resty/resty/v2"
    "github.com/heat/syscrapgo/cotacaosrv"
    "github.com/heat/syscrapgo/eventosrv"
    "strconv"
    "strings"
    "time"
    "github.com/heat/latinify"
)

type betsApiClient struct {
    URL string
    Key string
}

const (
    resultsEndpoint     = "/bet365/result"
    lngID               = "22"
    prematchOddEndpoint = "/bet365/prematch"
)

func (api *betsApiClient) convert(id string, s *betsApiGame) (*eventosrv.Evento, error) {

    eventTs, err := strconv.Atoi(s.Time)
    if err != nil {
        return nil, err
    }

    ideven := fmt.Sprint("bbshaaip", id)

    r := &eventosrv.Evento{
        CampID:     fmt.Sprint("bbshaaip-", s.League.ID),
        Campeonato: s.League.Name,
        Casa:       s.Home.Name,
        Esporte:    s.SportID,
        EventoTs:   int64(eventTs),
        Fora:       s.Away.Name,
        ID:         ideven,
        Inicio:     time.Unix(int64(eventTs), 0),
        Sha:        ideven,
        SourceID:   id,
        Ciclo:      0,
    }
    return r, nil
}

func (api *betsApiClient) Get(source string) (*eventosrv.Evento, error) {
    client := resty.New()

    client.SetHostURL(api.URL)

    client.SetQueryParams(map[string]string{
        "token":    api.Key,
        "event_id": source,
        "LNG_ID":   lngID,
    })

    resp, err := client.R().
        SetResult(&betsApiResult{}).
        Get(resultsEndpoint)
    if err != nil {
        return nil, err
    }

    result := resp.Result().(*betsApiResult)

    if result.Success != 1 {
        return nil, errors.New(result.Error)
    }
    // get first result
    if len(result.Results) == 0 {
        return nil, errors.New("evento nao econtrado")
    }
    ev := result.Results[0]
    if ev.Success != nil {
        return nil, errors.New("evento nao econtrado")
    }

    return api.convert(source, &ev)
}

func (api *betsApiClient) Result(source string) (*eventosrv.ScoreResult, error) {
    return nil, errors.New("nao implementado")
}

func (api *betsApiClient) GetOddsRaw(source string) (*cotacaosrv.Odds, error) {
    client := resty.New()

    client.SetHostURL(api.URL)

    client.SetQueryParams(map[string]string{
        "token":  api.Key,
        "FI":     source,
        "LNG_ID": lngID,
    })

    resp, err := client.R().
        SetResult(&betsApiPrematchOddResult{}).
        Get(prematchOddEndpoint)
    if err != nil {
        return nil, err
    }

    result := resp.Result().(*betsApiPrematchOddResult)

    if result.Success != 1 {
        return nil, errors.New(result.Error)
    }
    // get first result
    if len(result.Results) == 0 {
        return nil, errors.New("evento nao econtrado")
    }

    odds := result.Results[0]

    return api.oddConvert(source, &odds)
}

func (api *betsApiClient) oddConvert(s string, betApiOdd *betsApiPrematchOdd) (*cotacaosrv.Odds, error) {

    //oddsr := []cotacaosrv.Odd{}

    c := cotacaosrv.Odds{}
    // adiciona as main odds
    for market, odds := range betApiOdd.Main.SP {
        for _, odd := range odds {

            taxa, _ := strconv.ParseFloat(odd.Odds, 64)
            mulv := fmt.Sprintf("%s__%s__%s__%s__%s",
                strings.TrimSpace(strings.ToLower(market)),
                strings.TrimSpace(strings.ToLower(odd.Name)),
                strings.TrimSpace(strings.ToLower(odd.Header)),
                strings.TrimSpace(strings.ToLower(odd.Goals)),
                strings.TrimSpace(strings.ToLower(odd.Opp)), )
            mulv, _ = latinify.Slugify(mulv)
            modd := cotacaosrv.Odd{
                Codigo: "",
                Evento: "",
                Fonte:  "",
                ID:     "",
                Linha:  0,
                Odd:    "",
                Ref:    cotacaosrv.OddRef{},
                Taxa:   taxa,
                Mulv:   mulv,
                Ts:     time.Now().Unix(),
            }
            c = append(c, modd)
        }
    }
    return &c, nil
}

type betsApiFixture struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Cc      string `json:"cc"`
    ImageID string `json:"image_id"`
}

type betsApiScores struct {
    HtScore betsApiTimeScore `json:"1"`
    FtScore betsApiTimeScore `json:"2"`
}

type betsApiTimeScore struct {
    Home string `json:"home"`
    Away string `json:"away"`
}

type betsApiGameEvent struct {
    ID   string `json:"id"`
    Text string `json:"text"`
}

type betsApiGame struct {
    ID              string             `json:"id"`
    SportID         string             `json:"sport_id"`
    Time            string             `json:"time"`
    TimeStatus      string             `json:"time_status"`
    League          betsApiFixture     `json:"league"`
    Home            betsApiFixture     `json:"home"`
    Away            betsApiFixture     `json:"away"`
    Ss              string             `json:"ss"`
    Scores          betsApiScores      `json:"scores"`
    Events          []betsApiGameEvent `json:"events"`
    Success         interface{}        `json:"success"`
    HasLineup       int                `json:"has_lineup"`
    InplayCreatedAt string             `json:"inplay_created_at"`
    InplayUpdatedAt string             `json:"inplay_updated_at"`
    ConfirmedAt     string             `json:"confirmed_at"`
}

type betsApiResult struct {
    Success int           `json:"success"`
    Error   string        `json:"error,omitempty"`
    Results []betsApiGame `json:"results"`
}

type betsApiPrematchOdd struct {
    FI         string      `json:"FI"`
    EventID    string      `json:"event_id"`
    AsianLines betsApiSectionOdd `json:"asian_lines"`
    // DISABLED Corners    betsApiSectionOdd `json:"corners"`
    Goals      betsApiSectionOdd `json:"goals"`
    Half       betsApiSectionOdd `json:"half"`
    Main       betsApiSectionOdd `json:"main"`
    // DISABLED Player     interface{} `json:"player"`
    Schedule   betsApiSectionOdd `json:"schedule"`
    Specials   betsApiSectionOdd `json:"specials"`
}

type betsApiSectionOdd struct {
    UpdatedAt string `json:"updated_at"`
    SP betsApiMarketOdd `json:"sp"`
}
type betsApiGameOdd struct {
    Goals    string `json:"goals,omitempty"`
    Odds     string `json:"odds"`
    Header   string `json:"header"`
    Name     string `json:"name"`
    Handicap string `json:"handicap"`
    Opp      string `json:"opp"`
}

type betsApiGameOdds []betsApiGameOdd

type betsApiMarketOdd map[string]betsApiGameOdds

type betsApiPrematchOddResult struct {
    Success int                  `json:"success"`
    Error   string               `json:"error,omitempty"`
    Results []betsApiPrematchOdd `json:"results"`
}
