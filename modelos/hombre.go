package modelos

type Hombre struct {
	Edad       int
	Altura     float32
	Peso       float32
	Respirando bool
	Pensando   bool
	Comiendo   bool
	vivo       bool
}

func (this *Hombre) Respirar()      { this.Respirando = true }
func (this *Hombre) Comer()         { this.Comiendo = true }
func (this *Hombre) Pensar()        { this.Pensando = true }
func (this *Hombre) Sexo() string   { return "Hombre" }
func (this *Hombre) EstaVivo() bool { return this.vivo }
