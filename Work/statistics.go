package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func main() {

http.HandleFunc("/", homePage) //Receive 2 args: path and link to fucntion
	if err := http.ListenAndServe(":9001", nil); err != nil {
	log.Fatal("failed to start server", err)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // Должна вызываться перед записью в ответ
	fmt.Fprint(writer, pageTop, form)
		if err != nil {
			fmt.Fprintf(writer, anError, err)
		} else {
			if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
				fmt.Fprint(writer, formatStats(stats))
			} else if message != "" {
		fmt.Fprintf(writer, anError, message)
			}		
		}
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
		if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
			text := strings.Replace(slice[0], ",", " ", -1)
			for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
			return numbers, "’" + field + "’ is invalid", false
			} else {
				numbers = append(numbers, x)
				}		
			}
		}
	if len(numbers) == 0 {
	return numbers, "", false // при первом отображении данные отсутствуют
	}
	return numbers, "", true
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
	<tr><th colspan="2">Results</th></tr>
	<tr><td>Numbers</td><td>%v</td></tr>
	<tr><td>Count</td><td>%d</td></tr>
	<tr><td>Mean</td><td>%f</td></tr>
	<tr><td>Median</td><td>%f</td></tr>
	</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
	}

const form = `<form action="/" method="POST">
	<label for="numbers">Numbers (comma or space-separated):</label><br />
	<input type="text" name="numbers" size="30"><br />
	<input type="submit" value="Calculate">
	</form>`