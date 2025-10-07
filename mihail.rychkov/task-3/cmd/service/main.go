package main;

import "fmt";
import "os";

import "github.com/Rychmick/task-3/internal/config";

func main() {
	args := os.Args;

	if ((len(args) != 3) || (args[1] != "--config")) {
		fmt.Println("missing --config <file> in cmd args");
		return;
	}

	settings, err := config.Parse(args[2]);
	if (err != nil) {
		fmt.Println(err);
		return;
	}

	fmt.Println(settings.InputFilePath, settings.OutputFilePath);
}
