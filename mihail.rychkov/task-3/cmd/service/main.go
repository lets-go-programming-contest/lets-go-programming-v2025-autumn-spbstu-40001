package main;

import "fmt";
import "os";

import "github.com/Rychmick/task-3/internal/config";
import "github.com/Rychmick/task-3/internal/currency";

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

	currencyList, err := currency.ParseXml(settings.InputFilePath);
	if (err != nil) {
		fmt.Println(err);
		return;
	}
	err = currency.Prepare(&currencyList);
	if (err != nil) {
		fmt.Println(err);
		return;
	}

	fmt.Println(currencyList);
}
