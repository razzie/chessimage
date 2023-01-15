package main

import (
	"flag"
	"image/png"
	"log"
	"os"

	"github.com/razzie/chessimage"
)

const demoFEN = "rnbqkbnr/ppppp2p/5p2/6pQ/3PP3/8/PPP2PPP/RNB1KBNR b KQkq - 1 3"

func main() {
	var fen, output string
	flag.StringVar(&fen, "fen", demoFEN, "Forsyth-Edwards Notation")
	flag.StringVar(&output, "o", "board.png", "Output path for the board PNG")
	flag.Parse()

	r, err := chessimage.NewRendererFromFEN(fen)
	if err != nil {
		log.Fatal(err)
	}

	if fen == demoFEN {
		r.SetLastMove(chessimage.LastMove{
			From: chessimage.D1,
			To:   chessimage.H5,
		})
		r.SetCheckTile(chessimage.E8)
	}

	board, err := r.Render(chessimage.Options{})
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := png.Encode(file, board); err != nil {
		log.Fatal(err)
	}
}
