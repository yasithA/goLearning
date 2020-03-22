package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	var correct int = 0
	var incorrect int = 0
	var answer int
	dat, err := ioutil.ReadFile("questions.csv")
	if err != nil {
		log.Fatal(err)
	}
	content := csv.NewReader(strings.NewReader(string(dat)))
	for {
		record, err := content.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record[0] + ", " + record[1])

		//text, _ := reader.ReadString('\n')
		//text = strings.Replace(text, "\n", "", -1)
		fmt.Scanf("%d\n", &answer)
		ans, _ := strconv.Atoi(record[1])
		if ans == answer {
			correct++
		} else {
			incorrect++
		}
	}
	fmt.Printf("Total Number of questions: %v \n", (correct + incorrect))
	fmt.Printf("Correctly Answered: %v \n", correct)
	fmt.Printf("Incorrectly Answered: %v \n", incorrect)
}
