package main

type Vector2 struct {
	x float64
	y float64
}

func (this *Vector2) set(x, y float64) {
	this.x = x
	this.y = y
}

func (this *Vector2) set_x(x float64) {
	this.x = x
}

func (this *Vector2) set_y(y float64) {
	this.y = y
}

func (this *Vector2) add(xy Vector2) {
	this.x += xy.x
	this.y += xy.y
}

func (this *Vector2) add_x(x float64) {
	this.x += x
}

func (this *Vector2) add_y(y float64) {
	this.y += y
}

func (this *Vector2) subtract(xy Vector2) {
	this.x -= xy.x
	this.y -= xy.y
}

func (this *Vector2) subtract_x(x float64) {
	this.x -= x
}

func (this *Vector2) subtract_y(y float64) {
	this.y -= y
}

func (this *Vector2) multiply(scalar float64) {
	this.x *= scalar
	this.y *= scalar
}

func (this *Vector2) divide(scalar float64) {
	this.x /= scalar
	this.y /= scalar
}

func (this *Vector2) reset() {
	this.x = 0
	this.y = 0
}
