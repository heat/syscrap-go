package client

import (
    "errors"
    "fmt"
    "github.com/go-resty/resty/v2"
    "github.com/heat/latinify"
    "github.com/heat/syscrapgo/cotacaosrv"
    "github.com/heat/syscrapgo/eventosrv"
    "log"
    "regexp"
    "strconv"
    "strings"
    "time"
)

type betsApiClient struct {
    URL string
    Key string
    logger *log.Logger
}

var (
    allowedMarket = map[string]bool{
        "asian_handicap": true,
        "goal_line": true,
        "alternative_asian_handicap": true,
        "alternative_goal_line": true,
        //"1st_half_asian_handicap": true,
        "1st_half_goal_line": true,
        //"alternative_1st_half_asian_handicap": true,
        "alternative_1st_half_goal_line": true,
        //"corners": true,
        //"total_corners": true,
        //"alternative_corners": true,
        //"corners_2_way": true,
        //"first_half_corners": true,
        //"asian_total_corners": true,
        //"1st_half_asian_corners": true,
        "goals_over_under": true,
        "alternative_total_goals": true,
        "result_total_goals": true,
        "total_goals_both_teams_to_score": true,
        "exact_total_goals_(bands)": true,
        "number_of_goals_in_match": true,
        "both_teams_to_score": true,
        "teams_to_score": true,
        "both_teams_to_score_in_1st_half": true,
        "both_teams_to_score_in_2nd_half": true,
        "both_teams_to_score_1st_half_2nd_half": true,
        "first_half_goals": true,
        "exact_1st_half_goals": true,
        //"total_goal_minutes": true,
        //"first_team_to_score": true,
        //"early_goal": true,
        //"late_goal": true,
        //"time_of_first_goal_brackets": true,
        "2nd_half_goals": true,
        "exact_2nd_half_goals": true,
        "half_with_most_goals": true,
        "home_team_highest_scoring_half": true,
        "away_team_highest_scoring_half": true,
        "clean_sheet": true,
        "team_total_goals": true,
        "home_team_exact_goals": true,
        "away_team_exact_goals": true,
        //"time_of_1st_team_goal": true,
        "goals_odd_even": true,
        "home_team_odd_even_goals": true,
        "away_team_odd_even_goals": true,
        "1st_half_goals_odd_even": true,
        //"last_team_to_score": true,
        //"first_10_minutes_(00:00_09:59)": true,
        "half_time_result": true,
        "half_time_double_chance": true,
        "half_time_result_both_teams_to_score": true,
        "half_time_result_total_goals": true,
        "half_time_correct_score": true,
        //"1st_half_handicap": true,
        //"alternative_1st_half_handicap_result": true,
        "first_half_corners": true,
        "to_score_in_half": true,
        "2nd_half_result": true,
        "2nd_half_goals_odd_even": true,
        "full_time_result": true,
        "double_chance": true,
        "correct_score": true,
        "half_time_full_time": true,
        //"goalscorers": true,
        //"multi_scorers": true,
        //"draw_no_bet": true,
        "result_both_teams_to_score": true,
        "handicap_result": true,
        "alternative_handicap_result": true,
        "winning_margin": true,
        //"goalscorers": true,
        //"multi_scorers": true,
        //"team_goalscorer": true,
        "main": true,
        "specials": true,
        //"to_score_in_half": true,
        //"own_goal": true,
    }
)
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

    // talvez e so talvez eu tenha o o_away o_home
    if s.OAway.Name != "" {
        s.Away = s.OAway
    }
    if s.OHome.Name != "" {
        s.Home = s.OHome
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
        //"LNG_ID":   lngID,
    })

    resp, err := client.R().
        SetResult(&betsApiResult{}).
        Get(resultsEndpoint)
    if err != nil {
        return nil, err
    }
    if !resp.IsSuccess() {
        return nil, errors.New(string(resp.Body()))
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

func betApiSpConvert(section string, markets *betsApiMarketOdd, logger *log.Logger) *cotacaosrv.Odds {

    c := cotacaosrv.Odds{}
    // adiciona as main odds
    for market, odds := range *markets {
        if _, ok := allowedMarket[market]; !ok {
            logger.Printf("market=%s not allowed!\n", market)
            continue
        }
        for _, odd := range odds {

            taxa, _ := strconv.ParseFloat(odd.Odds, 64)

            parts := []string {
                strings.TrimSpace(strings.ToLower(odd.Name)),
                strings.TrimSpace(strings.ToLower(odd.Header)),
                strings.TrimSpace(strings.ToLower(odd.Goals)),
                strings.TrimSpace(strings.ToLower(odd.Handicap)),
                strings.TrimSpace(strings.ToLower(odd.Opp)),
            }

            mulv := strings.TrimSpace(strings.ToLower(market))

            for _, part := range(parts) {
                if part == "" {
                    continue
                }
                mulv = mulv + "__" + part
            }

            mulv, _ = slugify(mulv)

            modd := cotacaosrv.Odd{
                Codigo: "",
                Evento: "",
                Fonte:  mulv,
                ID:     "",
                Linha:  0,
                Odd:    "",
                Ref:    cotacaosrv.OddRef{},
                Taxa:   taxa,
                Ts:     time.Now().Unix(),
            }
            c = append(c, modd)
        }
    }

    return &c;
}

func slugify(source string) (string, error) {
    latin, err := latinify.String(source)
    if err != nil {
        return "", err
    }

    spacers := regexp.MustCompile(`\s+`)
    notAWord := regexp.MustCompile(`[^a-zA-Z0-9\-_\.\+]+`)
    transformers := apply(
        strings.ToLower,
        strings.TrimSpace,
        func(src string) string {
            return spacers.ReplaceAllString(src, "-")
        },
        func(src string) string {
            return notAWord.ReplaceAllString(src, "")
        },
    )

    res := transformers(latin)
    return res, nil
}

type stringTransform func(string) string

func apply(funcs ...stringTransform) func(string) string {

    return func(src string) (res string) {
        res = src
        for _, f := range funcs {
            res = f(res)
        }
        return
    }
}

func (api *betsApiClient) oddConvert(s string, betsApiOdd *betsApiPrematchOdd) (*cotacaosrv.Odds, error) {

    //oddsr := []cotacaosrv.Odd{}

    c := cotacaosrv.Odds{}
    var t *cotacaosrv.Odds

    // adiciona as main odds
    t = betApiSpConvert("main", &betsApiOdd.Main.SP, api.logger)
    c = append(c, *t...)

    // adicion goals odds
    t = betApiSpConvert("goals", &betsApiOdd.Goals.SP, api.logger)
    c = append(c, *t...)

    // adicion half odds
    t = betApiSpConvert("half", &betsApiOdd.Half.SP, api.logger)
    c = append(c, *t...)

    // adicion specials odds
    t = betApiSpConvert("specials", &betsApiOdd.Specials.SP, api.logger)
    c = append(c, *t...)

    // adicion asian odds
    t = betApiSpConvert("asian", &betsApiOdd.AsianLines.SP, api.logger)
    c = append(c, *t...)

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
    OHome           betsApiFixture     `json:"o_home,omitempty"`
    Away            betsApiFixture     `json:"away"`
    OAway           betsApiFixture     `json:"o_away,omitempty"`
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
