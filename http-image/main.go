package main

import (
    "bytes"
    "net/http"
    "image"
    "log"
    "image/color"
	"image/draw"
	"image/jpeg"
    "strconv"
)


func imageHandler(request http.ResponseWriter, response *http.Request){
    rgb := image.NewRGBA(image.Rect(0, 0, 1, 1))
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(rgb, rgb.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	var img image.Image = rgb

    buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	request.Header().Set("Content-Type", "image/jpeg")
	request.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := request.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}

    log.Println(response.Header)
}


func main(){
    http.HandleFunc("/image.jpg", imageHandler)

    log.Println("Listening :5000")
    err := http.ListenAndServe(":5000", nil)

    if err != nil {
        log.Fatal("Erro server", err)
    }
}
