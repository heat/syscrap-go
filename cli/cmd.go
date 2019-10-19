package cli

import (
    "github.com/spf13/cobra"
    "log"
)
var rootCommand = &cobra.Command{
    Use: "syscrap",
    Short: "Comanda a leitura de odds da api",
    Long: ` Utilize com moderacao
    `,

    Run: func(cmd *cobra.Command, args[]string) {

    },
}

func init() {
    cobra.OnInitialize();

    rootCommand.AddCommand(analyzeCmd);

}

// Execute main command syscrap
func Execute() {

    if err := rootCommand.Execute(); err != nil {
        log.Fatal(err);
    }
}
