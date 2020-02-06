package form

type Deficiencia struct{
	Desc string
	Agravante bool
	Obs []string
}

func Form4() []Deficiencia {
	auditiva:=Deficiencia{
		"Auditiva",
		false,
		[]string{
			"A surdez ocorreu antes dos 6 anos",
			"Não dispõe de auxílio de terceiros sempre que necessário",
		},
	}
	intelectual:=Deficiencia{
		"Intelectual (Cognitiva ou Mental)",
		false,
		[]string{
			"Não pode ficar sozinho em segurança",
			"Não dispõe de auxílio de terceiros sempre que necessário",
		},
	} 
	motora:=Deficiencia{
		"Motora",
		false,
		[]string{
			"Desloca-se exclusivamente em cadeira de rodas",
			"Não dispõe de auxílio de terceiros sempre que necessário",
		},
	}
	visual:=Deficiencia{
		"Visual",
		false,
		[]string{
			"A pessoa já não enxergava ao nascer",
			"Não dispõe de auxílio de terceiros sempre que necessário",
		},
	}
	return []Deficiencia{
		auditiva,
		intelectual,
		motora,
		visual,
	}
}
