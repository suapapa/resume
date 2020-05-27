package main

func main() {
	genResume(loadData())
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
