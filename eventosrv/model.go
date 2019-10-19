package eventosrv

import "time"

type Evento struct {
    CampID     string    `json:"camp_id"`
    Campeonato string    `json:"campeonato"`
    Casa       string    `json:"casa"`
    Esporte    string    `json:"esporte"`
    EventoTs   int64     `json:"evento_ts"`
    Fora       string    `json:"fora"`
    ID         string    `json:"id"`
    Inicio     time.Time `json:"inicio"`
    Sha        string    `json:"sha"`
    SourceID   string    `json:"source_id"`
    Ciclo      int64     `json:"ciclo"`
}

type Score struct {
    HtScore    ScoreResult `json:"ht_score"`
    FtScore    ScoreResult `json:"ft_score"`
    Calculated bool        `json:"valid"`
}

type ScoreResult struct {
    LocalteamScore   int
    VisitorteamScore int
}

