package form

import(
	"fmt"
)

var count int

// Funcao defines the basic structure coletions of functions
type Funcao struct{
	ID int
	Tipo string
	SubFuncao []subfuncao
}

type subfuncao struct{
	Tipo string
	Detalhe string
}

// Form2 defines and populates the collections of Functions
func Form2()[]Funcao{
	mentais:=Funcao{
		autoID(),
		"Funções Mentais",
		[]subfuncao{
			subfuncao{
				"Funções Mentais Globais",
				"Consciência, orientação, (tempo, lugar, pessoa), intelectuais (inclui desenvolvimento cognitivo e intelectual), psicossociais (inclui altimo), temperamento e personalidade, energia e impulsos, sono.",
			},
			subfuncao{
				"Funções Mentais Específicas",
				"Atenção, memória, psicomotoras, emocionais, percepção, pensamento, funções executivas, linguagem, cálculo, sequenciamento de movimentos complexos (inclui apraxia), experiência pessoal e do tempo.",
			},
		},
	}

	sensoriais:=Funcao{
		autoID(),
		"Funções Sensoriais e Dor",
		[]subfuncao{
			subfuncao{
				"Visão e Funções Relacionadas",
				"Acuidade Visual, Campos Visual, Funções dos músculos internos e externos do olho, da pálpebra, glândulas lacrimais.",
			},
			subfuncao{
				"Funções Auditivas",
				"Detecção, descriminação, localização, do som e da fala.",
			},
			subfuncao{
				"Funções Vestibulares",
				"Relacionadas à posição, equilíbrio e movimento.",
			},
			subfuncao{
				"Dor",
				"Sensação desagradável que indica lesão potencial ou real em alguma parte do corpo. Generalizada ou localizada.",
			},
			subfuncao{
				"Funções Sensoriais adicionais",
				"Gustativa, olfativa, proprioceptiva, tátil, à dor, temperatura.",
			},
		},
	}

	vozfala:=Funcao{
		autoID(),
		"Funções da Voz e da Fala",
		[]subfuncao{
			subfuncao{
				"Voz, articulação, fluência, ritmo da fala",
				"Sem mais detalhes.",
			},
		},
	}

	sangues:=Funcao{
		autoID(),
		"Funções dos Sistemas Cardiovasculares, Hematológico, Imunológico e Respiratório",
		[]subfuncao{
			subfuncao{
				"Funções do Sistema Cardiovascular",
				"Função do coração, vasos sanguíneos, pressão arterial.",
			},
			subfuncao{
				"Funções do Sistema Hematológico",
				"Produção de sangue, transporte de oxigênio e metabólicos e de coagulação.",
			},
			subfuncao{
				"Funções do Sistema Imunológico",
				"Resposta imunológica, reações de hipersensibilidade, funções do sistema linfático.",
			},
			subfuncao{
				"Funções do Sistema Respiratório",
				"Respiratórias, dos músculos respiratórios, da tolerância aos exercícios.",
			},
		},
	}

	digestivo:=Funcao{
		autoID(),
		"Funções dos Sistemas Digestivo, Metabólico e Endócrino",
		[]subfuncao{
			subfuncao{
				"Funções dos Sistemas Digestivos",
				"Ingestão, deglutição, digestivas, assimilação, defecação, manutenção de peso.",
			},
			subfuncao{
				"Funções do Metabolismo e Sistema Endócrino",
				"Funções Metabólicas gerais, equilíbrio hídrico, mineral e eletrolítico, termorreguladoras, das glândulas endócrinas.",
			},
		},
	}

	idosos:=Funcao{
		autoID(),
		"Funções Genitourinárias e Reprodutivas",
		[]subfuncao{
			subfuncao{
				"Funções Urinárias",
				"Funções de filtragem, coleta e excreção de urina.",
			},	
			subfuncao{
				"Funções Genitais e Reprodutivas",
				"Funções mentais e físicas/motoras relacionadas ao ato sexual, da menstruação e da procriação.",
			},
		},
	}
		
	musculo:=Funcao{
		autoID(),
		"Funções Neuromusculoesqueléticas e relacionadas a Movimento",
		[]subfuncao{
			subfuncao{
				"Funções das Articulações e dos Ossos",
				"Mobilidade, estabilidade das articulações e ossos.",
				},
				subfuncao{
					"Funções Musculares",
					"Força, tônus e resistência muscular.",
				},
				subfuncao{
					"Funções dos Movimentos",
					"Reflexo motor, movimentos involuntários, controle dos movimentos voluntários, padrão da marcha, sensações relacionadas aos músculos e funções do movimento.",
				},
		},
	}
	
	pele:=Funcao{
		autoID(),
		"Funções de Pele e Estruturas Relacionadas",
		[]subfuncao{
			subfuncao{
				"Funções dos Movimentos",
				"Protetora, reparadora, sensação relacionada à pele, pelos e unhas.",
			},
		},
	}
	output:= []Funcao{
		mentais,
		sensoriais,
		vozfala,
		sangues,
		digestivo,
		idosos,
		musculo,
		pele,		
	}
	return output
}

func autoID()int{
	count++
	return count
}

