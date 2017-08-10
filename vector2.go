package main

type Vector2 struct {
	x float64
	y float64
}

func (this *Vector2) x() {
	return this.x
}

func (this *Vector2) y() {
	return this.y
}

func (this *Vector2) set(x, y float64) {
	this.x = x
	this.y = y
}

func (this *Vector2) add(xy Vector2) {
	this.x += xy.x
	this.y += xy.y
}

func (this *Vector2) multiply(xy Vector2) {
	this.x *= xy.x
	this.y *= xy.y
}

func (this *Vector2) subtract(xy Vector2) {
	this.x -= xy.x
	this.y -= xy.y
}

func (this *Vector2) divide(xy Vector2) {
	this.x /= xy.x
	this.y /= xy.y
}

func (this *Vector2) reset() {
	this.x = 0
	this.y = 0
}
