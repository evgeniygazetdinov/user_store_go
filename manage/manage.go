package manage

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func CreateFile() {
	if os.Args[1] == "startapp" && len(os.Args) > 2 && os.Args[2] != "" {

		packageName := os.Args[2]
		os.MkdirAll(packageName, os.ModePerm)
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
		myPackage := []byte("package " + packageName)
		x := []string{"handlers.go", "models.go", "views.go"}
		for i := 0; i < len(x); i++ {
			fmt.Printf("%x ", x[i])
			f, err := os.Create(x[i])

			if err != nil {
				log.Fatal(err)
			}

			err = os.WriteFile(x[i], myPackage, 0644)
			if err != nil {
				log.Fatal(err)
				defer f.Close()
			}
		}

	}
}

func RunServer() {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)
	app := "docker-compose"

	arg0 := "-f"
	arg1 := "docker-compose.yml"
	arg2 := "up -d"
	arg3 := "db"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))

}
