package form

type form1 struct{
	Nome string
	Lotacao string
	Sexo []string
	Matricula string
	DataNascimento string
	Idade int
	TipoDeficiencia []string
	Etnia []string
	Diagnostico Diagnostico
}

type Diagnostico struct{
	CID string
	Desc string
	LocalAval string
	Informante []string
	Historico string
}

func Form1()form1{
	form1:=form1{}
	form1.Sexo=[]string{"F","M"}
	form1.Etnia=[]string{"branca","preta","amarela","parda","indígena"}
	form1.TipoDeficiencia=[]string{"física","visual","auditiva","multipla","intelectual ou cognitiva"}
	diag:=Diagnostico{}
	diag.Informante=[]string{"a própria pessoa","pessoa de convívio próprio","ambos","outros"}
	form1.Diagnostico=diag
	return form1
}