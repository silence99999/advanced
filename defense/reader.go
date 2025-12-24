package defense

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Println(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
