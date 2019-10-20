package cotacaosrv


func GetOddRef(codigo string)(*OddRef, bool) {
    ref, ok := references[codigo]
    if !ok {
        return nil, false
    }

    return &ref, true
}

var references = map[string]OddRef{

    "1st_half_goals_odd_even__even": {Codigo: "igualdade-gols-1t_par", Tipo: "IH", Posicao: 1, Linha: false},       
    "1st_half_goals_odd_even__odd":  {Codigo: "igualdade-gols-1t_impar", Tipo: "IH", Posicao: 2, Linha: false},     
    "2nd_half_goals__over__0.5":     {Codigo: "total-gols-2t_acima", Tipo: "BT", Posicao: 1, Linha: true, V: 0.5},  
    "2nd_half_goals__over__1.5":     {Codigo: "total-gols-2t_acima", Tipo: "BT", Posicao: 1, Linha: true, V: 1.5},  
    "2nd_half_goals__over__2.5":     {Codigo: "total-gols-2t_acima", Tipo: "BT", Posicao: 1, Linha: true, V: 2.5},  
    "2nd_half_goals__over__3.5":     {Codigo: "total-gols-2t_acima", Tipo: "BT", Posicao: 1, Linha: true, V: 3.5},  
    "2nd_half_goals__over__4.5":     {Codigo: "total-gols-2t_acima", Tipo: "BT", Posicao: 1, Linha: true, V: 4.5},  
    "2nd_half_goals__over__5.5":     {Codigo: "total-gols-2t_acima", Tipo: "BT", Posicao: 1, Linha: true, V: 5.5},  
    "2nd_half_goals__over__6.5":     {Codigo: "total-gols-2t_acima", Tipo: "BT", Posicao: 1, Linha: true, V: 6.5},  
    "2nd_half_goals__under__0.5":    {Codigo: "total-gols-2t_abaixo", Tipo: "BT", Posicao: 2, Linha: true, V: 0.5}, 
    "2nd_half_goals__under__1.5":    {Codigo: "total-gols-2t_abaixo", Tipo: "BT", Posicao: 2, Linha: true, V: 1.5}, 
    "2nd_half_goals__under__2.5":    {Codigo: "total-gols-2t_abaixo", Tipo: "BT", Posicao: 2, Linha: true, V: 2.5}, 
    "2nd_half_goals__under__3.5":    {Codigo: "total-gols-2t_abaixo", Tipo: "BT", Posicao: 2, Linha: true, V: 3.5}, 
    "2nd_half_goals__under__4.5":    {Codigo: "total-gols-2t_abaixo", Tipo: "BT", Posicao: 2, Linha: true, V: 4.5}, 
    "2nd_half_goals__under__5.5":    {Codigo: "total-gols-2t_abaixo", Tipo: "BT", Posicao: 2, Linha: true, V: 5.5}, 
    "2nd_half_goals__under__6.5":    {Codigo: "total-gols-2t_abaixo", Tipo: "BT", Posicao: 2, Linha: true, V: 6.5}, 

    "2nd_half_goals_odd_even__even": {Codigo: "igualdade-gols-2t_impar", Tipo: "IS", Posicao: 2, Linha: false}, 
    "2nd_half_goals_odd_even__odd":  {Codigo: "igualdade-gols-2t_par", Tipo: "IS", Posicao: 1, Linha: false},   

    "2nd_half_result__1": {Codigo: "resultado-final-2t_casa", Tipo: "RT", Posicao: 1, Linha: false},   
    "2nd_half_result__2": {Codigo: "resultado-final-2t_fora", Tipo: "RT", Posicao: 2, Linha: false},   
    "2nd_half_result__x": {Codigo: "resultado-final-2t_empate", Tipo: "RT", Posicao: 3, Linha: false}, 

    
    "alternative_asian_handicap__away__+1.5": {Codigo: "resultado-final-handicap_fora", Tipo: "HR", Posicao: 2, Linha: true, V: 1.5},
    "alternative_asian_handicap__home__-1.5": {Codigo: "resultado-final-handicap_casa", Tipo: "HR", Posicao: 1, Linha: true, V: -1.5},

    "goals_over_under__over__2.5":  {Codigo: "total-gols_acima", Tipo: "GT", Posicao: 1, Linha: true, V: 2.5},  
    "goals_over_under__under__2.5": {Codigo: "total-gols_abaixo", Tipo: "GT", Posicao: 2, Linha: true, V: 2.5}, 
    "goals_over_under__over__1.5":  {Codigo: "total-gols_acima", Tipo: "GT", Posicao: 1, Linha: true, V: 1.5},  
    "goals_over_under__under__1.5": {Codigo: "total-gols_abaixo", Tipo: "GT", Posicao: 2, Linha: true, V: 1.5}, 

    "alternative_total_goals__over__0.5":  {Codigo: "total-gols_acima", Tipo: "GT", Posicao: 1, Linha: true, V: 0.5},  
    "alternative_total_goals__over__1.5":  {Codigo: "total-gols_acima", Tipo: "GT", Posicao: 1, Linha: true, V: 1.5},  
    "alternative_total_goals__over__2.5":  {Codigo: "total-gols_acima", Tipo: "GT", Posicao: 1, Linha: true, V: 2.5},  
    "alternative_total_goals__over__3.5":  {Codigo: "total-gols_acima", Tipo: "GT", Posicao: 1, Linha: true, V: 3.5},  
    "alternative_total_goals__over__4.5":  {Codigo: "total-gols_acima", Tipo: "GT", Posicao: 1, Linha: true, V: 4.5},  
    "alternative_total_goals__over__5.5":  {Codigo: "total-gols_acima", Tipo: "GT", Posicao: 1, Linha: true, V: 5.5},  
    "alternative_total_goals__over__6.5":  {Codigo: "total-gols_acima", Tipo: "GT", Posicao: 1, Linha: true, V: 6.5},  
    "alternative_total_goals__under__0.5": {Codigo: "total-gols_abaixo", Tipo: "GT", Posicao: 2, Linha: true, V: 0.5}, 
    "alternative_total_goals__under__1.5": {Codigo: "total-gols_abaixo", Tipo: "GT", Posicao: 2, Linha: true, V: 1.5}, 
    "alternative_total_goals__under__2.5": {Codigo: "total-gols_abaixo", Tipo: "GT", Posicao: 2, Linha: true, V: 2.5}, 
    "alternative_total_goals__under__3.5": {Codigo: "total-gols_abaixo", Tipo: "GT", Posicao: 2, Linha: true, V: 3.5}, 
    "alternative_total_goals__under__4.5": {Codigo: "total-gols_abaixo", Tipo: "GT", Posicao: 2, Linha: true, V: 4.5}, 
    "alternative_total_goals__under__5.5": {Codigo: "total-gols_abaixo", Tipo: "GT", Posicao: 2, Linha: true, V: 5.5}, 
    "alternative_total_goals__under__6.5": {Codigo: "total-gols_abaixo", Tipo: "GT", Posicao: 2, Linha: true, V: 6.5}, 

    "away_team_exact_goals__away---0-goals":  {Codigo: "total-gol-exato-fora_0", Tipo: "TF", Posicao: 1, Linha: false},  
    "away_team_exact_goals__away---1-goal":   {Codigo: "total-gol-exato-fora_1", Tipo: "TF", Posicao: 2, Linha: false},  
    "away_team_exact_goals__away---2-goals":  {Codigo: "total-gol-exato-fora_2", Tipo: "TF", Posicao: 3, Linha: false},  
    "away_team_exact_goals__away---3+-goals": {Codigo: "total-gol-exato-fora_3+", Tipo: "TF", Posicao: 4, Linha: false}, 

    
    "away_team_highest_scoring_half__away---1st-half": {Codigo: "total-gols_acima", Tipo: "XX", Posicao: 1, Linha: false},
    "away_team_highest_scoring_half__away---2nd-half": {Codigo: "total-gols_acima", Tipo: "XX", Posicao: 1, Linha: false},
    "away_team_highest_scoring_half__away---tie":      {Codigo: "total-gols_acima", Tipo: "XX", Posicao: 1, Linha: false},

    "away_team_odd_even_goals__away---even": {Codigo: "igualdade-gol-time_fora-impar", Tipo: "IT", Posicao: 4, Linha: false}, 
    "away_team_odd_even_goals__away---odd":  {Codigo: "igualdade-gol-time_fora-par", Tipo: "IT", Posicao: 3, Linha: false},   

    "both_teams_to_score__yes":                        {Codigo: "marcacao-gols_ambos", Tipo: "MG", Posicao: 1, Linha: false},             
    "both_teams_to_score__no":                         {Codigo: "marcacao-gols_unico", Tipo: "MG", Posicao: 2, Linha: false},             
    "both_teams_to_score_1st_half_2nd_half__no--no":   {Codigo: "marcacao-gols-1t-2t_unico_unico", Tipo: "MT", Posicao: 1, Linha: false}, 
    "both_teams_to_score_1st_half_2nd_half__no--yes":  {Codigo: "marcacao-gols-1t-2t_unico_ambos", Tipo: "MT", Posicao: 2, Linha: false}, 
    "both_teams_to_score_1st_half_2nd_half__yes--no":  {Codigo: "marcacao-gols-1t-2t_ambos-unico", Tipo: "MT", Posicao: 3, Linha: false}, 
    "both_teams_to_score_1st_half_2nd_half__yes--yes": {Codigo: "marcacao-gols-1t-2t_ambos-ambos", Tipo: "MT", Posicao: 4, Linha: false}, 

    "both_teams_to_score_in_1st_half__no":  {Codigo: "marcacao-gols-1t_unico", Tipo: "MH", Posicao: 2, Linha: false}, 
    "both_teams_to_score_in_1st_half__yes": {Codigo: "marcacao-gols-1t_ambos", Tipo: "MH", Posicao: 1, Linha: false}, 

    "both_teams_to_score_in_2nd_half__no":  {Codigo: "marcacao-gols-2t_unico", Tipo: "MS", Posicao: 2, Linha: false}, 
    "both_teams_to_score_in_2nd_half__yes": {Codigo: "marcacao-gols-2t_ambos", Tipo: "MS", Posicao: 1, Linha: false}, 

    "clean_sheet__home__yes": {Codigo: "marcacao-gols-time_casa-toma-nao", Tipo: "LC", Posicao: 1, Linha: false}, 
    "clean_sheet__home__no":  {Codigo: "marcacao-gols-time_casa-toma-sim", Tipo: "LC", Posicao: 2, Linha: false}, 
    "clean_sheet__away__yes": {Codigo: "marcacao-gols-time_fora-toma-nao", Tipo: "LC", Posicao: 3, Linha: false}, 
    "clean_sheet__away__no":  {Codigo: "marcacao-gols-time_fora-toma-sim", Tipo: "LC", Posicao: 4, Linha: false}, 

    "correct_score__away__1-0": {Codigo: "resultado-exato_0x1", Tipo: "XF", Posicao: 1, Linha: false},  
    "correct_score__away__2-0": {Codigo: "resultado-exato_0x2", Tipo: "XF", Posicao: 2, Linha: false},  
    "correct_score__away__2-1": {Codigo: "resultado-exato_1x2", Tipo: "XF", Posicao: 12, Linha: false}, 
    "correct_score__away__3-0": {Codigo: "resultado-exato_0x3", Tipo: "XF", Posicao: 3, Linha: false},  
    "correct_score__away__3-1": {Codigo: "resultado-exato_1x3", Tipo: "XF", Posicao: 13, Linha: false}, 
    "correct_score__away__3-2": {Codigo: "resultado-exato_2x3", Tipo: "XF", Posicao: 23, Linha: false}, 
    "correct_score__away__4-0": {Codigo: "resultado-exato_0x4", Tipo: "XF", Posicao: 4, Linha: false},  
    "correct_score__away__4-1": {Codigo: "resultado-exato_1x4", Tipo: "XF", Posicao: 14, Linha: false}, 
    "correct_score__away__4-2": {Codigo: "resultado-exato_2x4", Tipo: "XF", Posicao: 24, Linha: false}, 
    "correct_score__away__4-3": {Codigo: "resultado-exato_3x4", Tipo: "XF", Posicao: 34, Linha: false}, 
    "correct_score__away__5-0": {Codigo: "resultado-exato_0x5", Tipo: "XF", Posicao: 5, Linha: false},  
    "correct_score__away__5-1": {Codigo: "resultado-exato_1x5", Tipo: "XF", Posicao: 15, Linha: false}, 
    "correct_score__away__5-2": {Codigo: "resultado-exato_2x5", Tipo: "XF", Posicao: 25, Linha: false}, 
    "correct_score__away__5-3": {Codigo: "resultado-exato_3x5", Tipo: "XF", Posicao: 35, Linha: false}, 
    "correct_score__away__5-4": {Codigo: "resultado-exato_4x5", Tipo: "XF", Posicao: 45, Linha: false}, 
    "correct_score__home__1-0": {Codigo: "resultado-exato_1x0", Tipo: "XF", Posicao: 10, Linha: false}, 
    "correct_score__home__2-0": {Codigo: "resultado-exato_2x0", Tipo: "XF", Posicao: 20, Linha: false}, 
    "correct_score__home__2-1": {Codigo: "resultado-exato_2x1", Tipo: "XF", Posicao: 21, Linha: false}, 
    "correct_score__home__3-0": {Codigo: "resultado-exato_3x0", Tipo: "XF", Posicao: 30, Linha: false}, 
    "correct_score__home__3-1": {Codigo: "resultado-exato_3x1", Tipo: "XF", Posicao: 31, Linha: false}, 
    "correct_score__home__3-2": {Codigo: "resultado-exato_3x2", Tipo: "XF", Posicao: 32, Linha: false}, 
    "correct_score__home__4-0": {Codigo: "resultado-exato_4x0", Tipo: "XF", Posicao: 40, Linha: false}, 
    "correct_score__home__4-1": {Codigo: "resultado-exato_4x1", Tipo: "XF", Posicao: 41, Linha: false}, 
    "correct_score__home__4-2": {Codigo: "resultado-exato_4x2", Tipo: "XF", Posicao: 42, Linha: false}, 
    "correct_score__home__4-3": {Codigo: "resultado-exato_4x3", Tipo: "XF", Posicao: 43, Linha: false}, 
    "correct_score__home__5-0": {Codigo: "resultado-exato_5x0", Tipo: "XF", Posicao: 50, Linha: false}, 
    "correct_score__home__5-1": {Codigo: "resultado-exato_5x1", Tipo: "XF", Posicao: 51, Linha: false}, 
    "correct_score__home__5-2": {Codigo: "resultado-exato_5x2", Tipo: "XF", Posicao: 52, Linha: false}, 
    "correct_score__home__5-3": {Codigo: "resultado-exato_5x3", Tipo: "XF", Posicao: 53, Linha: false}, 
    "correct_score__home__5-4": {Codigo: "resultado-exato_5x4", Tipo: "XF", Posicao: 54, Linha: false}, 
    "correct_score__x__0-0":    {Codigo: "resultado-exato_0x0", Tipo: "XF", Posicao: 0, Linha: false},  
    "correct_score__x__1-1":    {Codigo: "resultado-exato_1x1", Tipo: "XF", Posicao: 11, Linha: false}, 
    "correct_score__x__2-2":    {Codigo: "resultado-exato_2x2", Tipo: "XF", Posicao: 22, Linha: false}, 
    "correct_score__x__3-3":    {Codigo: "resultado-exato_3x3", Tipo: "XF", Posicao: 33, Linha: false}, 
    "correct_score__x__4-4":    {Codigo: "resultado-exato_4x4", Tipo: "XF", Posicao: 44, Linha: false}, 
    "correct_score__x__5-5":    {Codigo: "resultado-exato_5x5", Tipo: "XF", Posicao: 55, Linha: false}, 

    "double_chance__12": {Codigo: "resultado-duplo_casa-fora", Tipo: "RD", Posicao: 2, Linha: false},   
    "double_chance__1x": {Codigo: "resultado-duplo_casa-empate", Tipo: "RD", Posicao: 1, Linha: false}, 
    "double_chance__x2": {Codigo: "resultado-duplo_fora-empate", Tipo: "RD", Posicao: 3, Linha: false}, 

    "exact_1st_half_goals__0-goals": {Codigo: "total-gol-exato-1t_0", Tipo: "TA", Posicao: 1, Linha: false}, 
    "exact_1st_half_goals__1-goal":  {Codigo: "total-gol-exato-1t_1", Tipo: "TA", Posicao: 2, Linha: false}, 
    "exact_1st_half_goals__2-goals": {Codigo: "total-gol-exato-1t_2", Tipo: "TA", Posicao: 3, Linha: false}, 

    "exact_2nd_half_goals__0-goals": {Codigo: "total-gol-exato-2t_0", Tipo: "TB", Posicao: 1, Linha: false}, 
    "exact_2nd_half_goals__1-goal":  {Codigo: "total-gol-exato-2t_1", Tipo: "TB", Posicao: 2, Linha: false}, 
    "exact_2nd_half_goals__2-goals": {Codigo: "total-gol-exato-2t_2", Tipo: "TB", Posicao: 3, Linha: false}, 

    "exact_total_goals_bands__0-goals": {Codigo: "total-gol-exato-resultado_0", Tipo: "TR", Posicao: 0, Linha: false}, 
    "exact_total_goals_bands__1-goal":  {Codigo: "total-gol-exato-resultado_1", Tipo: "TR", Posicao: 1, Linha: false}, 
    "exact_total_goals_bands__2-goals": {Codigo: "total-gol-exato-resultado_2", Tipo: "TR", Posicao: 2, Linha: false}, 
    "exact_total_goals_bands__3-goals": {Codigo: "total-gol-exato-resultado_3", Tipo: "TR", Posicao: 3, Linha: false}, 
    "exact_total_goals_bands__4-goals": {Codigo: "total-gol-exato-resultado_4", Tipo: "TR", Posicao: 4, Linha: false}, 
    "exact_total_goals_bands__5-goals": {Codigo: "total-gol-exato-resultado_5", Tipo: "TR", Posicao: 5, Linha: false}, 

    "first_half_goals__over__0.5":  {Codigo: "total-gols-1t_acima", Tipo: "AH", Posicao: 1, Linha: true, V: 0.5},  
    "first_half_goals__over__1.5":  {Codigo: "total-gols-1t_acima", Tipo: "AH", Posicao: 1, Linha: true, V: 1.5},  
    "first_half_goals__over__2.5":  {Codigo: "total-gols-1t_acima", Tipo: "AH", Posicao: 1, Linha: true, V: 2.5},  
    "first_half_goals__over__3.5":  {Codigo: "total-gols-1t_acima", Tipo: "AH", Posicao: 1, Linha: true, V: 3.5},  
    "first_half_goals__over__4.5":  {Codigo: "total-gols-1t_acima", Tipo: "AH", Posicao: 1, Linha: true, V: 4.5},  
    "first_half_goals__over__5.5":  {Codigo: "total-gols-1t_acima", Tipo: "AH", Posicao: 1, Linha: true, V: 5.5},  
    "first_half_goals__under__0.5": {Codigo: "total-gols-1t_abaixo", Tipo: "AH", Posicao: 2, Linha: true, V: 0.5}, 
    "first_half_goals__under__1.5": {Codigo: "total-gols-1t_abaixo", Tipo: "AH", Posicao: 2, Linha: true, V: 1.5}, 
    "first_half_goals__under__2.5": {Codigo: "total-gols-1t_abaixo", Tipo: "AH", Posicao: 2, Linha: true, V: 2.5}, 
    "first_half_goals__under__3.5": {Codigo: "total-gols-1t_abaixo", Tipo: "AH", Posicao: 2, Linha: true, V: 3.5}, 
    "first_half_goals__under__4.5": {Codigo: "total-gols-1t_abaixo", Tipo: "AH", Posicao: 2, Linha: true, V: 4.5}, 
    "first_half_goals__under__5.5": {Codigo: "total-gols-1t_abaixo", Tipo: "AH", Posicao: 2, Linha: true, V: 5.5}, 

    "full_time_result__1": {Codigo: "resultado-final_casa", Tipo: "RF", Posicao: 1, Linha: false},   
    "full_time_result__2": {Codigo: "resultado-final_fora", Tipo: "RF", Posicao: 2, Linha: false},   
    "full_time_result__x": {Codigo: "resultado-final_empate", Tipo: "RF", Posicao: 3, Linha: false}, 

    "goals_odd_even__even": {Codigo: "igualdade-gols_impar", Tipo: "IG", Posicao: 2, Linha: false}, 
    "goals_odd_even__odd":  {Codigo: "igualdade-gols_par", Tipo: "IG", Posicao: 1, Linha: false},   

    "half_time_correct_score__away__1-0": {Codigo: "resultado-exato-1t_0x1", Tipo: "WH", Posicao: 1, Linha: false},  
    "half_time_correct_score__away__2-0": {Codigo: "resultado-exato-1t_0x2", Tipo: "WH", Posicao: 2, Linha: false},  
    "half_time_correct_score__away__2-1": {Codigo: "resultado-exato-1t_1x2", Tipo: "WH", Posicao: 12, Linha: false}, 
    "half_time_correct_score__away__3-0": {Codigo: "resultado-exato-1t_0x3", Tipo: "WH", Posicao: 3, Linha: false},  
    "half_time_correct_score__away__3-1": {Codigo: "resultado-exato-1t_1x3", Tipo: "WH", Posicao: 13, Linha: false}, 
    "half_time_correct_score__away__3-2": {Codigo: "resultado-exato-1t_2x3", Tipo: "WH", Posicao: 23, Linha: false}, 
    "half_time_correct_score__away__4-0": {Codigo: "resultado-exato-1t_0x4", Tipo: "WH", Posicao: 4, Linha: false},  
    "half_time_correct_score__away__4-1": {Codigo: "resultado-exato-1t_1x4", Tipo: "WH", Posicao: 14, Linha: false}, 
    "half_time_correct_score__away__4-2": {Codigo: "resultado-exato-1t_2x4", Tipo: "WH", Posicao: 24, Linha: false}, 
    "half_time_correct_score__away__4-3": {Codigo: "resultado-exato-1t_3x4", Tipo: "WH", Posicao: 34, Linha: false}, 
    "half_time_correct_score__away__5-0": {Codigo: "resultado-exato-1t_0x5", Tipo: "WH", Posicao: 5, Linha: false},  
    "half_time_correct_score__away__5-1": {Codigo: "resultado-exato-1t_1x5", Tipo: "WH", Posicao: 15, Linha: false}, 
    "half_time_correct_score__away__5-2": {Codigo: "resultado-exato-1t_2x5", Tipo: "WH", Posicao: 25, Linha: false}, 
    "half_time_correct_score__away__5-3": {Codigo: "resultado-exato-1t_3x5", Tipo: "WH", Posicao: 35, Linha: false}, 
    "half_time_correct_score__away__5-4": {Codigo: "resultado-exato-1t_4x5", Tipo: "WH", Posicao: 45, Linha: false}, 

    "half_time_correct_score__home__1-0": {Codigo: "resultado-exato-1t_1x0", Tipo: "WH", Posicao: 10, Linha: false}, 
    "half_time_correct_score__home__2-0": {Codigo: "resultado-exato-1t_2x0", Tipo: "WH", Posicao: 20, Linha: false}, 
    "half_time_correct_score__home__2-1": {Codigo: "resultado-exato-1t_2x1", Tipo: "WH", Posicao: 21, Linha: false}, 
    "half_time_correct_score__home__3-0": {Codigo: "resultado-exato-1t_3x0", Tipo: "WH", Posicao: 30, Linha: false}, 
    "half_time_correct_score__home__3-1": {Codigo: "resultado-exato-1t_3x1", Tipo: "WH", Posicao: 21, Linha: false}, 
    "half_time_correct_score__home__3-2": {Codigo: "resultado-exato-1t_3x2", Tipo: "WH", Posicao: 32, Linha: false}, 
    "half_time_correct_score__home__4-0": {Codigo: "resultado-exato-1t_4x0", Tipo: "WH", Posicao: 40, Linha: false}, 
    "half_time_correct_score__home__4-1": {Codigo: "resultado-exato-1t_4x1", Tipo: "WH", Posicao: 41, Linha: false}, 
    "half_time_correct_score__home__4-2": {Codigo: "resultado-exato-1t_4x2", Tipo: "WH", Posicao: 42, Linha: false}, 
    "half_time_correct_score__home__4-3": {Codigo: "resultado-exato-1t_4x3", Tipo: "WH", Posicao: 43, Linha: false}, 
    "half_time_correct_score__home__5-0": {Codigo: "resultado-exato-1t_5x0", Tipo: "WH", Posicao: 50, Linha: false}, 
    "half_time_correct_score__home__5-1": {Codigo: "resultado-exato-1t_5x1", Tipo: "WH", Posicao: 51, Linha: false}, 
    "half_time_correct_score__home__5-2": {Codigo: "resultado-exato-1t_5x2", Tipo: "WH", Posicao: 52, Linha: false}, 
    "half_time_correct_score__home__5-3": {Codigo: "resultado-exato-1t_5x3", Tipo: "WH", Posicao: 53, Linha: false}, 
    "half_time_correct_score__home__5-4": {Codigo: "resultado-exato-1t_5x4", Tipo: "WH", Posicao: 54, Linha: false}, 

    "half_time_correct_score__x__0-0": {Codigo: "resultado-exato-1t_0x0", Tipo: "WH", Posicao: 0, Linha: false},  
    "half_time_correct_score__x__1-1": {Codigo: "resultado-exato-1t_1x1", Tipo: "WH", Posicao: 11, Linha: false}, 
    "half_time_correct_score__x__2-2": {Codigo: "resultado-exato-1t_2x2", Tipo: "WH", Posicao: 22, Linha: false}, 
    "half_time_correct_score__x__3-3": {Codigo: "resultado-exato-1t_3x3", Tipo: "WH", Posicao: 33, Linha: false}, 
    "half_time_correct_score__x__4-4": {Codigo: "resultado-exato-1t_4x4", Tipo: "WH", Posicao: 44, Linha: false}, 
    "half_time_correct_score__x__5-5": {Codigo: "resultado-exato-1t_5x5", Tipo: "WH", Posicao: 55, Linha: false}, 

    "half_time_double_chance__12": {Codigo: "resultado-duplo-1t_casa-fora", Tipo: "RH", Posicao: 2, Linha: false},   
    "half_time_double_chance__1x": {Codigo: "resultado-duplo-1t_casa-empate", Tipo: "RH", Posicao: 1, Linha: false}, 
    "half_time_double_chance__x2": {Codigo: "resultado-duplo-1t_fora-empate", Tipo: "RH", Posicao: 3, Linha: false}, 

    "half_time_result__1": {Codigo: "resultado-intervalo_casa", Tipo: "TH", Posicao: 1, Linha: false},   
    "half_time_result__2": {Codigo: "resultado-intervalo_fora", Tipo: "TH", Posicao: 2, Linha: false},   
    "half_time_result__x": {Codigo: "resultado-intervalo_empate", Tipo: "TH", Posicao: 3, Linha: false}, 

    "home_team_exact_goals__home---0-goals":  {Codigo: "total-gol-exato-casa_0", Tipo: "TC", Posicao: 1, Linha: false},  
    "home_team_exact_goals__home---1-goal":   {Codigo: "total-gol-exato-casa_1", Tipo: "TC", Posicao: 2, Linha: false},  
    "home_team_exact_goals__home---2-goals":  {Codigo: "total-gol-exato-casa_2", Tipo: "TC", Posicao: 3, Linha: false},  
    "home_team_exact_goals__home---3+-goals": {Codigo: "total-gol-exato-casa_3+", Tipo: "TC", Posicao: 4, Linha: false}, 

    "home_team_odd_even_goals__home---even": {Codigo: "igualdade-gol-time_casa-impar", Tipo: "IT", Posicao: 2, Linha: false}, 
    "home_team_odd_even_goals__home---odd":  {Codigo: "igualdade-gol-time_casa-par", Tipo: "IT", Posicao: 1, Linha: false},   

    "result_both_teams_to_score__away__no":  {Codigo: "resultado-marcacao_fora-unico-marca", Tipo: "RM", Posicao: 4, Linha: false},     
    "result_both_teams_to_score__away__yes": {Codigo: "resultado-marcacao_fora-ambos-marca", Tipo: "RM", Posicao: 3, Linha: false},     
    "result_both_teams_to_score__home__no":  {Codigo: "resultado-marcacao_casa-unico-marca", Tipo: "RM", Posicao: 2, Linha: false},     
    "result_both_teams_to_score__home__yes": {Codigo: "resultado-marcacao_casa-ambos-marca", Tipo: "RM", Posicao: 1, Linha: false},     
    "result_both_teams_to_score__draw__yes": {Codigo: "resultado-marcacao_empate-ambos-marca", Tipo: "RM", Posicao: 5, Linha: false},   
    "result_both_teams_to_score__draw__no":  {Codigo: "resultado-marcacao_empate-ninguem-marca", Tipo: "RM", Posicao: 6, Linha: false}, 

    "result_total_goals__away--over-2.5":  {Codigo: "resultado-final-total-gols_fora-acima", Tipo: "KG", Posicao: 6, Linha: true, V: 2.5},    
    "result_total_goals__away--under-2.5": {Codigo: "resultado-final-total-gols_fora-abaixo", Tipo: "KG", Posicao: 5, Linha: true, V: 2.5},   
    "result_total_goals__draw--over-2.5":  {Codigo: "resultado-final-total-gols_empate-acima", Tipo: "KG", Posicao: 4, Linha: true, V: 2.5},  
    "result_total_goals__draw--under-2.5": {Codigo: "resultado-final-total-gols_empate-abaixo", Tipo: "KG", Posicao: 3, Linha: true, V: 2.5}, 
    "result_total_goals__home--over-2.5":  {Codigo: "resultado-final-total-gols_casa-acima", Tipo: "KG", Posicao: 2, Linha: true, V: 2.5},    
    "result_total_goals__home--under-2.5": {Codigo: "resultado-final-total-gols_casa-abaixo", Tipo: "KG", Posicao: 1, Linha: true, V: 2.5},   

    "team_total_goals__away__o-1.5": {Codigo: "total-gols-time_fora-acima", Tipo: "JF", Posicao: 1, Linha: true, V: 1.5},  
    "team_total_goals__away__o-2.5": {Codigo: "total-gols-time_fora-acima", Tipo: "JF", Posicao: 1, Linha: true, V: 2.5},  
    "team_total_goals__away__o-3.5": {Codigo: "total-gols-time_fora-acima", Tipo: "JF", Posicao: 1, Linha: true, V: 3.5},  
    "team_total_goals__away__o-4.5": {Codigo: "total-gols-time_fora-acima", Tipo: "JF", Posicao: 1, Linha: true, V: 4.5},  
    "team_total_goals__away__u-1.5": {Codigo: "total-gols-time_fora-abaixo", Tipo: "JF", Posicao: 2, Linha: true, V: 1.5}, 
    "team_total_goals__away__u-2.5": {Codigo: "total-gols-time_fora-abaixo", Tipo: "JF", Posicao: 2, Linha: true, V: 2.5}, 
    "team_total_goals__away__u-3.5": {Codigo: "total-gols-time_fora-abaixo", Tipo: "JF", Posicao: 2, Linha: true, V: 3.5}, 
    "team_total_goals__away__u-4.5": {Codigo: "total-gols-time_fora-abaixo", Tipo: "JF", Posicao: 2, Linha: true, V: 4.5}, 

    "team_total_goals__home__o-1.5": {Codigo: "total-gols-time_casa-acima", Tipo: "KC", Posicao: 1, Linha: true, V: 1.5},  
    "team_total_goals__home__o-2.5": {Codigo: "total-gols-time_casa-acima", Tipo: "KC", Posicao: 1, Linha: true, V: 2.5},  
    "team_total_goals__home__o-3.5": {Codigo: "total-gols-time_casa-acima", Tipo: "KC", Posicao: 1, Linha: true, V: 3.5},  
    "team_total_goals__home__o-4.5": {Codigo: "total-gols-time_casa-acima", Tipo: "KC", Posicao: 1, Linha: true, V: 4.5},  
    "team_total_goals__home__u-1.5": {Codigo: "total-gols-time_casa-abaixo", Tipo: "KC", Posicao: 2, Linha: true, V: 1.5}, 
    "team_total_goals__home__u-2.5": {Codigo: "total-gols-time_casa-abaixo", Tipo: "KC", Posicao: 2, Linha: true, V: 2.5}, 
    "team_total_goals__home__u-3.5": {Codigo: "total-gols-time_casa-abaixo", Tipo: "KC", Posicao: 2, Linha: true, V: 3.5}, 
    "team_total_goals__home__u-4.5": {Codigo: "total-gols-time_casa-abaixo", Tipo: "KC", Posicao: 2, Linha: true, V: 4.5}, 

    "teams_to_score__away-only": {Codigo: "resultado-time-marca_fora-casa-nao", Tipo: "NT", Posicao: 2, Linha: false}, 
    "teams_to_score__home-only": {Codigo: "resultado-time-marca_casa-fora-nao", Tipo: "NT", Posicao: 1, Linha: false}, 
    
    "teams_to_score__home-teams_to_score__no-goal": {Codigo: "resultado-time-marca_casa-nao-fora-nao", Tipo: "NT", Posicao: 3, Linha: false}, 

    "specials__to-win-both-halves__home": {Codigo: "resultado-tempo_ambos-casa", Tipo: "RE", Posicao: 1, Linha: false}, 
    "specials__to-win-both-halves__away": {Codigo: "resultado-tempo_ambos-fora", Tipo: "RE", Posicao: 2, Linha: false}, 

    "total_goals_both_teams_to_score__over-2.5--yes":  {Codigo: "total-gols-marcacao_acima-sim-25", Tipo: "TM", Posicao: 2, Linha: true, V: 2.5},  
    "total_goals_both_teams_to_score__over-2.5--no":   {Codigo: "total-gols-marcacao_acima-nao-25", Tipo: "TM", Posicao: 1, Linha: true, V: 2.5},  
    "total_goals_both_teams_to_score__under-2.5--yes": {Codigo: "total-gols-marcacao_abaixo-sim-25", Tipo: "TM", Posicao: 4, Linha: true, V: 2.5}, 
    "total_goals_both_teams_to_score__under-2.5--no":  {Codigo: "total-gols-marcacao_abaixo-nao-25", Tipo: "TM", Posicao: 3, Linha: true, V: 2.5}, 
}
