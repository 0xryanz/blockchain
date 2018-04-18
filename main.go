package main

import (
  "log"
)

func logf(f string, v ...interface{}) {
  log.Printf(f, v...)
}

func main() {

  bc := NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
