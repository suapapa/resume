package main

// Resume contains basic config informations
type Resume struct {
	Name        string `yaml:"name"`
	Title       string `yaml:"title"`
	Title2      string `yaml:"title2"`
	Description string `yaml:"description"`
	URL         string `yaml:"url"`
	Baseurl     string `yaml:"baseurl"`
	Personal    string `yaml:"personal"`
	Email       string `yaml:"email"`
	Twitter     string `yaml:"twitter"`
	Copyright   string `yaml:"copyright"`
	Navigation  string `yaml:"navigation"`
	Analytics   string `yaml:"analytics"`
	// Assets      struct {
	// 	Dirname       string   `yaml:"dirname"`
	// 	Baseurl       string   `yaml:"baseurl"`
	// 	CSSCompressor string   `yaml:"css_compressor"`
	// 	JsCompressor  string   `yaml:"js_compressor"`
	// 	Cachebust     string   `yaml:"cachebust"`
	// 	Cache         bool     `yaml:"cache"`
	// 	Gzip          string   `yaml:"gzip"`
	// 	Sources       []string `yaml:"sources"`
	// } `yaml:"assets"`
	// Exclude []string `yaml:"exclude"`

	SectionsName []string
	Sections     Sections
}

// Sections reperesents resume sections
type Sections []struct {
	ID     string      `yaml:"id"`
	Name   string      `yaml:"name"`
	Layout string      `yaml:"layout"`
	Data   interface{} // this should be List or Blocks
}

// List reperesents list layout
type List []struct {
	Title string   `yaml:"title"`
	Item  []string `yaml:"item"`
}

// Blocks reperesents block layout
type Blocks []struct {
	Org      string   `yaml:"org"`
	Role     string   `yaml:"role"`
	Location string   `yaml:"location,omitempty"`
	When     string   `yaml:"when"`
	Bullets  []string `yaml:"bullets"`
}
