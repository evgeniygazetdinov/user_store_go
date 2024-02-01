package manage

import (
	"fmt"
	"log"
	"os"
)

func CreateFile() {
	if os.Args[1] == "startapp" && len(os.Args) > 2 && os.Args[2] != "" {
		packageName := os.Args[2]
		os.MkdirAll(packageName, os.ModeDir)
		currDir, err := os.Getwd()
		fmt.Println(currDir)

		if err != nil {

			fmt.Println(err)

			return

		}
		fmt.Println(currDir)
		err = os.Chdir(currDir + "/" + packageName)
		if err != nil {

			fmt.Println(err)

			return

		}

		currDir, err = os.Getwd()

		fmt.Println(currDir)

		x := []string{"handlers.go", "models.go", "views.go"}
		for i := 0; i < len(x); i++ {
			fmt.Printf("%x ", x[i])
			f, err := os.Create(x[i])

			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
		}
	}

}
