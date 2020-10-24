package multimedia

import "fmt"

// Imagen ...
type Imagen struct {
	Titulo  string
	Formato string
	Canales []string
}

// Mostrar ...
func (img *Imagen) Mostrar() {
	fmt.Println("Titulo:", img.Titulo)
	fmt.Println("Formato:", img.Formato)
	fmt.Print("Canales: ")
	for _, v := range img.Canales {
		fmt.Print(v, " - ")
	}
	fmt.Print("\n\n")
}

// Audio ...
type Audio struct {
	Titulo   string
	Formato  string
	Duracion int64 //Segundos.
}

// Mostrar ...
func (aud *Audio) Mostrar() {
	fmt.Println("Titulo:", aud.Titulo)
	fmt.Println("Formato:", aud.Formato)
	fmt.Print("Duracion: ", aud.Duracion, "\n\n")
}

// Video ...
type Video struct {
	Titulo  string
	Formato string
	Frames  int64
}

// Mostrar ...
func (vid *Video) Mostrar() {
	fmt.Println("Titulo:", vid.Titulo)
	fmt.Println("Formato:", vid.Formato)
	fmt.Print("Frames: ", vid.Frames, "\n\n")
}

// Multimedia ...
type Multimedia interface {
	Mostrar()
}

// ContenidoWeb ...
type ContenidoWeb struct {
	Multimedios []Multimedia
}

// Mostrar ...
func (contWeb *ContenidoWeb) Mostrar() {
	for _, multimedio := range contWeb.Multimedios {
		multimedio.Mostrar()
	}
}
