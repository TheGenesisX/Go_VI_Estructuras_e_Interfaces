package main

import (
	"bufio"
	"fmt"
	"os"

	"./multimedia"
)

func main() {
	var opc, frames, duracion int64
	var formato string
	var webContent multimedia.ContenidoWeb
	scanner := bufio.NewScanner(os.Stdin) //Para el Titulo.

	for {
		fmt.Println("---------- Menu Multimedia ----------")

		fmt.Println("1) Capturar Imagen.")
		fmt.Println("2) Capturar Audio.")
		fmt.Println("3) Capturar Video.")
		fmt.Println("4) Mostrar arreglo de Multimedia.")
		fmt.Println("0) Salir.")

		fmt.Print("\nElija una opcion: ")
		fmt.Scan(&opc)

		if opc == 0 {
			break
		}

		switch opc {
		case 1:
			var canales []string
			var canal string
			img := new(multimedia.Imagen)

			fmt.Print("\nTitulo de la imagen: ")
			scanner.Scan()
			img.Titulo = scanner.Text()

			fmt.Print("Formato de la imagen: ")
			fmt.Scan(&formato)
			img.Formato = formato

			fmt.Println("Canales de la imagen: (para salir, ingresar z)")

			for {
				fmt.Scan(&canal)
				if canal == "z" {
					break
				} else {
					canales = append(canales, canal)
				}
			}
			img.Canales = canales

			webContent.Multimedios = append(webContent.Multimedios, img)
		case 2:
			aud := new(multimedia.Audio)

			fmt.Print("\nTitulo del audio: ")
			scanner.Scan()
			aud.Titulo = scanner.Text()

			fmt.Print("Formato del audio: ")
			fmt.Scan(&formato)
			aud.Formato = formato

			fmt.Print("Duracion del audio (segundos): ")
			fmt.Scan(&duracion)
			aud.Duracion = duracion
			webContent.Multimedios = append(webContent.Multimedios, aud)
		case 3:
			vid := new(multimedia.Video)

			fmt.Print("\nTitulo del video: ")
			scanner.Scan()
			vid.Titulo = scanner.Text()

			fmt.Print("Formato del video: ")
			fmt.Scan(&formato)
			vid.Formato = formato

			fmt.Print("Frames del video: ")
			fmt.Scan(&frames)
			vid.Frames = frames

			webContent.Multimedios = append(webContent.Multimedios, vid)
		case 4:
			fmt.Println("\n----------------------")
			webContent.Mostrar()
			fmt.Println("\n----------------------")
		default:
			fmt.Println("Opcion No Valida.")
		}
	}
}
