package game

import "space_shooter/assets"

const (
	screenWidth   = 800
	screenHeight  = 600
	titleFontSize = fontSize * 1.5
	fontSize      = 24
	smallFontSize = fontSize / 2
)

var (
	ScoreFontFace    = assets.GetFontFace(24)
	HealthFontFace   = assets.GetFontFace(24)
	MenuItemFontFace = assets.GetFontFace(32)
)

// Vector representa o objeto com os atributos das posições X e Y de cada entidade dentro do jogo
type Vector struct {
	X float64
	Y float64
}

// Rect representa a hitbox de cada entidade dentro do jogo
type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// NewRect cria uma nova instância de um regângulo
func NewRect(x, y, width, height float64) Rect {
	return Rect{
		x,
		y,
		width,
		height,
	}
}

// Intersects verifica se os retangulos de duas entidades se encostaram
func (rect Rect) Intersects(other Rect) bool {
	return rect.X <= other.MaxX() &&
		other.X <= rect.MaxX() &&
		rect.Y <= other.MaxY() &&
		other.Y <= rect.MaxY()
}

// MaxX retorna a posição final do eixo X do retangulo
func (rect Rect) MaxX() float64 {
	return rect.X + rect.Width
}

// MaxY retorna a posição final do eixo Y do retangulo
func (rect Rect) MaxY() float64 {
	return rect.Y + rect.Height
}
