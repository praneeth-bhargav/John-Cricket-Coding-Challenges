package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"io"
)

func main(){
	bytes_flag := flag.Bool("b",false,"used to count bytes");
	char_flag := flag.Bool("c", false, "used to count charecters");
	word_flag := flag.Bool("w", false, "used to count workds");
	line_flag := flag.Bool("l", false, "used to count no of lines");

	flag.Parse();

	if len(flag.Args()) < 1{
		// fmt.Printf("Usage: ccwc [-c] [-t] [-w] <file_name>")
		// return;
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		PrintResult(data, bytes_flag, char_flag, word_flag, line_flag);
	}else{
		file_name := flag.Arg(0);

		_, err := os.Stat(file_name);

		if os.IsNotExist(err) {
			log.Fatal("unable to read file",err);
		}

		content, err := os.ReadFile(file_name);

		if err != nil{
			log.Fatal("error ",err);
		}

		PrintResult(content,bytes_flag,char_flag,word_flag,line_flag);
		fmt.Printf("%s",strings.Split(file_name, "/")[len(strings.Split(file_name, "/"))-1]);
	}


	

}

func PrintResult(content []byte, bytes_flag *bool, char_flag *bool, word_flag *bool, line_flag *bool){
	all_no_condition := !(*bytes_flag) && !(*char_flag) && !(*word_flag) && !(*line_flag)

	if *bytes_flag || all_no_condition{
		fmt.Printf("%d ",len(content));
	}

	if *char_flag || all_no_condition{
		fmt.Printf("%d ", len(string(content)))
	}

	if *word_flag || all_no_condition{
		fmt.Printf("%d ", len(strings.Fields(string(content))))
	}

	if *line_flag || all_no_condition{
		fmt.Printf("%d ", len(bytes.Split(content,[]byte("\n"))))
	}
}