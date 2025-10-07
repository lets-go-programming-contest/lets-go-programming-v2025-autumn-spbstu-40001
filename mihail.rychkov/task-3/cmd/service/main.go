package main;

import "fmt";
import "os";

func main() {
	args := os.Args;

	if ((len(args) != 3) || (args[1] != "--config")) {
		fmt.Println("missing --config <file> in cmd args");
		return;
	}

	yamlFileData, err := os.ReadFile(args[2]);
	if (err != nil) {
		fmt.Println("failed to read config file:", err);
		return;
	}

	fmt.Println(string(yamlFileData));
}
