package ink

import "fmt"
import "github.com/ttacon/chalk"

func Success(s string) {
	fmt.Println(chalk.Green.Color(s))
}

func Warning(s string) {
	fmt.Println(chalk.Yellow.Color(s))
}

func Error(s string) {
	fmt.Println(chalk.Red.Color(s))
}
