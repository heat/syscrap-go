package cotacaosrv

type OddRef struct {
    Codigo  string `json:"codigo"`
    Linha   bool   `json:"linha"`
    Posicao int    `json:"posicao"`
    Tipo    string `json:"tipo"`
}

type Odd struct {
    Codigo string  `json:"codigo"`
    Evento string  `json:"evento"`
    Fonte  string  `json:"fonte"`
    ID     string  `json:"id"`
    Linha  float64 `json:"linha"`
    Odd    string  `json:"odd"`
    Ref    OddRef  `json:"ref,omitempty"`
    Taxa   float64 `json:"taxa"`
    Mulv   string  `json:"mulv"`
    Ts     int64   `json:"ts"`
}

type Odds []Odd

type OddResult struct {
    Evento string `json:"evento"`
    Taxas  Odds   `json:"taxas"`
    Total  int64  `json:"total"`
    Ts     int64  `json:"ts"`
}
