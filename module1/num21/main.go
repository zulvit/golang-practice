package main

import "fmt"

// Реализую адаптер на примере музыкального проигрывателя
func main() {
	// Использование адаптера для проигрывания музыки через VinylPlayer
	player := VinylAdapter{Player: &VinylPlayer{}}
	result := player.PlayMusic("Stairway to Heaven")
	fmt.Println(result)
}

// MusicPlayer - целевой интерфейс для проигрывания музыки
type MusicPlayer interface {
	PlayMusic(song string) string
}

// VinylPlayer - сторонний класс для проигрывания виниловых пластинок
type VinylPlayer struct{}

func (v *VinylPlayer) PlayVinyl(record string) string {
	return fmt.Sprintf("Виниловая пластинка: %s", record)
}

// VinylAdapter - адаптер для интеграции VinylPlayer с MusicPlayer
type VinylAdapter struct {
	Player *VinylPlayer
}

func (va *VinylAdapter) PlayMusic(song string) string {
	return va.Player.PlayVinyl(song)
}
