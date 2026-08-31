package modelos

type Mujer struct {
	Hombre
}

func (this *Mujer) Sexo() string { return "Mujer" }
