package main

func main() {
	config := loadData()
	genResume(config)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
