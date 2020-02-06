package form

type Dominio struct{
	Desc string
	Dominios []dominio
}

type dominio struct{
	Dominio string
	Desc string
}

func Form3()[]Dominio{
	sensorial:=Dominio{
		"Domínio Sensorial",
		[]dominio{
			dominio{
				"Observar",
				"Enxergar, reconhecer, e interpretar o que enxerga.",
			},
			dominio{
				"Ouvir",
				"Perceber, discriminar, reconhecer e interpretar sons.",
			},
		},
	}
	
	comunicacao:=Dominio{
		"Domínio de comunicação",
		[]dominio{
			dominio{
				"Comunicar-se ",
				"Compreender mensagens.",
			},
			dominio{
				"Conversar ",
				"Iniciar, manter e finalizar uma troca de pensamentos e idéias, realizadas através da linguagem oral e ou de sinais",
			},
			dominio{
				"Conversar ",
				"Iniciar, manter e finalizar uma troca de pensamentos e idéias, realizada através da linguagem oral e ou de sinais",
			},
			dominio{
				"Discutir ",
				"Iniciar, manter e finalizar uma troca de pensamentos e idéias, realizada através da linguagem oral e ou de sinais",
			},
			dominio{
				"Utilização de dispositivos de comunicação à distância ",
				"A utilização de dispositivos de comunicação à distância",
			},
		},
	}
	
	mobilidade:=Dominio{
		"Domínio Mobilidade",
		[]dominio{
			dominio{
				"Mudar e manter a posição do corpo",
				"Sem detalhes",
			},
			dominio{
				"Alcançar, transportar e mover objetos",
				"Transportar e mover objeto de posição e alcançar acima da cabeça, à frente, ao lado e abaixo",
			},
			dominio{
				"Movimentos finos da mão",
				"Manusear objetos, manipulá-los e soltá-los utilizando as mãos, dedos e polega",
			},
			dominio{
				"Deslocar se dentro de casa",
				"Andar ou deslocar-se dentro da própria casa, em um ambiente, incluindo áreas anexas quando aplicável",
			},
			dominio{
				"Deslocar-se dentro de edifícios que não a própria casa",
				"Andar ou deslocar-se dentro de edifícios que não a própria residência, incluindo áreas anexas quando aplicável",
			},
			dominio{
				"Deslocar-se fora de sua casa e de outros edifícios",
				"Andar ou deslocar-se perto ou longe da própria casa e de outros edifícios, sem utilização de transporte público ou privado",
			},
			dominio{
				"Utilizar transporte coletivo",
				"Utilizar transporte coletivo para se deslocar, como passageiro, por meio terrestre, aqüaviário ou aéreo",
			},
			dominio{
				"Utilizar transporte individual como passageiro",
				"Utilizar transporte para se deslocar, como passageiro, por meio terrestre, aqüaviário ou aéreo",
			},
		},
	}
	
	autonomia:=Dominio{
		"Domínio Cuidados Pessoais",
		[]dominio{
			dominio{
				"Lavar-se",
				"Tomar banho e limpar o corpo de forma completa incluindo lavar e secar todas as partes do corpo da forma habitual",
			},
			dominio{
				"Cuidar de partes do corpo",
				"Realizar cuidados pessoais",
			},
			dominio{
				"Regulação da micção",
				"Reconhecer a necessidade de urinar, escolher a forma e o local apropriado, manipular a roupa, urinar e secar-se",
			},
			dominio{
				"Regulação da defecação",
				"Reconhecer a necessidade de evacuar, escolher a forma e o local apropriado, manipular a roupa, evacuar e limpar-se",
			},
			dominio{
				"Vestir-se",
				"Vestir e retirar peças habituais, incluíndo calçados",
			},
			dominio{
				"Comer",
				"Levar à boca e comer alimento preparado e servido de forma de forma habitual",
			},
			dominio{
				"Beber",
				"Levar à boca e beber bebida preparada e servida de forma de forma habitual e com deglutição segura",
			},
			dominio{
				"Capacidade de identificar agravos à saúde",
				"Identificar sinais e sintomas que possam potencialmente comprometer a saúde e a integridade física, reconhecer abusos e violência",
			},
		},
	}
	
	quotidiano:=Dominio{
		"Domínio Vida Doméstica",
		[]dominio{
			dominio{
				"Preparar refeições tipo lanches",
				"Preparar e escolher alimentos simples para lanches",
			},
			dominio{
				"Cozinhar",
				"Planejar, organizar e executar o preparo de refeições que exijam varios ingredientes ou utensílios, ou uma sequência de ações",
			},
			dominio{
				"Realizar tarefas doméstica",
				"",
			},
			dominio{
				"Manutenção e uso apropriado de objetos pessoais e utensílios da cas",
				"Utilizar, guardar e conservar objetos pessoais, cuidar da casa e dos utensílios domésticos",
			},
			dominio{
				"Cuidar dos Outro",
				"",
			},
		},
	}
	
	educacao:=Dominio{
		"Domínio Educação",
		[]dominio{
			dominio{
				"Educação",
				"Adquirir habilidades e conhecimentos educacionais, dentro e fora da escola, considerando acessibilidade e disponibilidade de recursos educacionais adequados",
			},
			dominio{
				"Qualificação Profissional",
				"Adquirir habilidades e conhecimentos específicos para atividade profissional, considerando acessibilidade e disponibilidade de recursos educacionais adequados",
			},
			dominio{
				"Trabalho Remunerado",
				"Exerce trabalho remunerado, considerando localizar, escolher, permanecer, progredir e sair de forma adequada",
			},
			dominio{
				"Fazer compras e contratar serviços",
				"Ober, em troca de dinheiros, bens e serviços necessários para a vida diária",
			},
			dominio{
				"Administração de recursos Econômicos Pessoais",
				"",
			},
		},
	}
	
	social:=Dominio{
		"Domínio Socialização e Vida Comunitária",
		[]dominio{
			dominio{
				"Regular o comportamento nas interações",
				"Controlar o próprio comportamento, emoções e impulsos, agressão verbal e física nas interações, de maneira contextual e socialmente apropriada",
			},
			dominio{
				"Interagir de acordo com as regras sociais",
				"Agir independentemente nas interações sociais e adaptar-se às convenções sociais que regem o papel, posição ou status social das pessoas nas interações com os outros",
			},
			dominio{
				"Relacionamentos com estranhos",
				"Estabelecer contatos e ligações temporários com estranhos para fins específicos como, quando aplicável",
			},
			dominio{
				"Relacionamentos familiares e com pessoas familiares",
				"Criar e manter relações de parentesco com membros do núcleo familiar ou pessoas próximas. Participar da rotina familiar",
			},
			dominio{
				"Relacionamentos Íntimos",
				"Verificar detalhe desse item",
			},
			dominio{
				"Socialização",
				"Participar de eventos sociais",
			},
			dominio{
				"Fazer as próprias escolhas",
				"Exercer a capacidade de tomar decisões sobre a sua própria vida",
			},
			dominio{
				"Vida política e cidadania",
				"Exercer a cidadania",
			},
		},
	}

	return []Dominio{
		sensorial,
		comunicacao,
		mobilidade,
		autonomia,
		quotidiano,
		educacao,
		social,
	}
}