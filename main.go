package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/nicored/avatar"
)

func main() {
	name := []string{"John", "Smith"}
	_ = name
	size := 200

	for _, n := range name {
		newAvatar, err := avatar.NewAvatarFromInitials([]byte(n), &avatar.InitialsOptions{
			FontPath:  "test_data/Arial.ttf",        // Required
			Size:      size,                         // default 300
			NInitials: 2,                            // default 1 - If 0, the whole text will be printed
			TextColor: color.White,                  // Default White
			BgColor:   color.RGBA{255, 0, 255, 255}, // Default color.RGBA{215, 0, 255, 255} (purple)
		})

		if err != nil {
			panic(err)
		}

		square, _ := newAvatar.Square()
		squareFile, _ := os.Create(fmt.Sprintf("output/square_%s_initials.png", n))
		defer squareFile.Close()
		squareFile.Write(square)

		round, _ := newAvatar.Circle()
		roundFile, _ := os.Create(fmt.Sprintf("output/round_%s_initials.png", n))
		defer roundFile.Close()
		roundFile.Write(round)
	}
}
