package cli

import (
    "encoding/json"
    "fmt"
    apiclient "github.com/heat/syscrapgo/client"
    "github.com/heat/syscrapgo/cotacaosrv"
    "github.com/heat/syscrapgo/eventosrv"
    "github.com/spf13/cobra"
    jww "github.com/spf13/jwalterweatherman"
    "github.com/spf13/viper"
    "log"
)

var (
    api = "betsapi"
    apikey = ""
    apis = map[string]string{
        "betsapi": "https://api.betsapi.com/v1",
    }
)

const (

)
var analyzeCmd = &cobra.Command{
    Use:   "analyze [source id]",
    Short: "Analisa as informacoes do jogo da api",
    Long:  ``,
    Args: cobra.ExactArgs(1),
    Run:   func(cmd *cobra.Command, args []string) {

        jww.SetLogThreshold(jww.LevelDebug)
        jww.SetStdoutThreshold(jww.LevelDebug)

        logger := jww.LOG

        kk := viper.GetString("apikey")

        api, err := apiclient.MakeApiClient(api, logger, &apiclient.ApiClientOptions{
            URL: apis[api],
            Key: kk,
        })

        if err != nil {
            log.Panic(err)
        }
        source := args[0]

        evento, err := api.Get(source)
        if err != nil {
            jww.CRITICAL.Fatal(err.Error())
        }

        odds, err := api.GetOddsRaw(source)
        if err != nil {
            jww.CRITICAL.Fatal(err.Error())
        }

        oddsFound := make(cotacaosrv.Odds, 0)
        oddsNotFound := make(cotacaosrv.Odds, 0)
        // passada contabilizacao
        for _, odd := range(*odds) {
            ref, found := cotacaosrv.GetOddRef(odd.Fonte)
            if found {
                jww.DEBUG.Printf("odd=%s ref=%s linha=%.2f\n", odd.Fonte, ref.Codigo, ref.V)
                odd.Ref = *ref
                odd.Codigo = fmt.Sprintf("odd_%s", ref.Codigo)
                odd.Odd = ref.Codigo
                odd.Linha = ref.V
                oddsFound = append(oddsFound, odd)
            } else {
                jww.DEBUG.Printf("odd=%s not found!", odd.Fonte)
                oddsNotFound = append(oddsNotFound, odd)
            }
        }

        jsons, err := json.Marshal(analyzeResult{evento, &oddsFound, &oddsNotFound})
        if err != nil {
            jww.CRITICAL.Fatal(err.Error())
        }
        jww.FEEDBACK.Println(string(jsons))
    },
}

type analyzeResult struct {
    Evento *eventosrv.Evento `json:"evento"`
    Odds *cotacaosrv.Odds `json:"odds"`
    OddsNotFound *cotacaosrv.Odds `json:"not_found"`
}
func init() {


    analyzeCmd.PersistentFlags().StringVar(&api, "api",  "betsapi", "define uma api para fazer analyze [ 'betsapi' ]")
    analyzeCmd.PersistentFlags().StringVar(&apikey, "apikey",  "", "define a chave para chamda da api")

    viper.BindPFlag("apikey", analyzeCmd.PersistentFlags().Lookup("apikey"))
    viper.AutomaticEnv()
}
