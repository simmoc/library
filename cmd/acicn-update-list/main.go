package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/acicn/library"
)

func main() {
	var err error
	defer func(err *error) {
		if *err != nil {
			log.Println("exited with error:", (*err).Error())
			os.Exit(1)
		}
	}(&err)

	var names []string

	for _, task := range library.Builds {
		for _, name := range task.Names {
			names = append(names, name)
		}
	}
	for _, task := range library.Mirrors {
		var subTasks []library.MirrorSubTask
		if subTasks, err = task.SubTasks(context.Background()); err != nil {
			return
		}
		for _, subTask := range subTasks {
			names = append(names, subTask.To)
		}
	}

	for _, name := range names {
		log.Println("Canonical Name:", name)
	}
	sort.Strings(names)
	if err = ioutil.WriteFile("IMAGES.txt", []byte(strings.Join(names, "\n")), 0644); err != nil {
		return
	}
}
