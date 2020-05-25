package main

// Sections reperesents resume sections
type Sections []struct {
	Name   string `yaml:"name"`
	ID     string `yaml:"id"`
	Layout string `yaml:"layout"`
	Data   interface{}
}

// Blocks reperesents block layout
type Blocks []struct {
	Org      string   `yaml:"org"`
	Role     string   `yaml:"role"`
	Location string   `yaml:"location,omitempty"`
	When     string   `yaml:"when"`
	Bullets  []string `yaml:"bullets"`
}

// List reperesents list layout
type List []struct {
	Title string   `yaml:"title"`
	Item  []string `yaml:"item"`
}
