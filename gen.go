package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"
)

//go:generate go run -tags=generate gen.go

func main() {
	fmt.Println("Generate code..")
	t := os.Getenv("GENERATOR_TYPE")
	routeName := os.Getenv("GENERATOR_ROUTE")
	handlerName := os.Getenv("GENERATOR_HANDLER")
	structName := os.Getenv("GENERATOR_STRUCT")
	jsonFile := os.Getenv("GENERATOR_JSON")

	// Check if the required parameters are provided
	if t == "" || handlerName == "" {
		fmt.Println("Missing required parameters.")
		return
	}

	// Use the parameters to generate code
	if t == "jwt" {
		// Generate JWT related code
		fmt.Printf("Generating JWT code with file name '%s'\n", handlerName)
		processTemplate("template/jwt.go", handlerName)
	} else if t == "handler" {
		fmt.Printf("Generating handler code with function name '%s'\n", handlerName)

		t, err := template.New("handler.tpl").ParseFiles("template/handler.tpl")
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}

		// Define the output file path
		outputFile := handlerName + ".go"

		// Create or open the output file
		f, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer f.Close()

		// Execute the template and write the generated code to the output file
		err = t.Execute(f, struct {
			Handler string
		}{
			Handler: handlerName, // Use the first parameter as the handler name
		})
		if err != nil {
			fmt.Println("Error executing template:", err)
			return
		}

		fmt.Println("Handler code generated successfully.")
	} else if t == "route" {
		fmt.Printf("Generating routing...")
		t, err := template.New("route.tpl").ParseFiles("template/route.tpl")
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}

		// Define the output file path
		outputFile := "routes.go"

		// Create or open the output file
		f, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer f.Close()

		// Execute the template and write the generated code to the output file
		err = t.Execute(f, struct {
			Handler string
		}{
			Handler: handlerName, // Use the first parameter as the handler name
		})
		if err != nil {
			fmt.Println("Error executing template:", err)
			return
		}

		fmt.Println("Routing code generated successfully.")
	} else if t == "new-route" {
		fmt.Printf("Generating route...")

		// Define the output file path
		outputFile := "routes.go"

		// Create or open the output file
		f, err := os.OpenFile(outputFile, os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer f.Close()

		// Read the contents of the existing file
		existingCode, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println("Error reading existing file:", err)
			return
		}

		// Find the placeholder comment in the existing file
		placeholder := "// Placeholder for generated code. Do not remove or modify this comment."
		index := strings.Index(string(existingCode), placeholder)
		if index == -1 {
			fmt.Println("Placeholder not found in the existing file")
			return
		}

		// Create a new buffer to store the modified code
		var generatedCode bytes.Buffer

		// Write the code before the placeholder
		generatedCode.Write(existingCode[:index])

		t, err := template.ParseFiles("template/create-route.tpl")
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}

		// Execute the template and write the generated code to the output file
		err = t.Execute(&generatedCode, struct {
			RouteName   string
			HandlerName string
		}{
			RouteName:   routeName,
			HandlerName: handlerName, // Use the first parameter as the handler name
		})
		if err != nil {
			fmt.Println("Error executing template:", err)
			return
		}

		// Write the code after the placeholder
		generatedCode.Write(existingCode[index:])

		// Truncate the existing file
		err = f.Truncate(0)
		if err != nil {
			fmt.Println("Error truncating file:", err)
			return
		}

		// Rewind to the beginning of the file
		_, err = f.Seek(0, 0)
		if err != nil {
			fmt.Println("Error seeking file:", err)
			return
		}

		// Write the modified code to the existing file
		_, err = generatedCode.WriteTo(f)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("Route code generated successfully.")
	} else if t == "entity" {
		fmt.Printf("Generating struct...")

		if _, err := os.Stat("entity"); os.IsNotExist(err) {
			log.Print("not exist folder")
			err = os.Mkdir("entity", 0755)
			if err != nil {
				log.Fatal(err)
			}
		}

		var resp map[string]any
		in, err := os.Open("coda/" + jsonFile)
		if err != nil {
			log.Fatal(err)
		}

		b, _ := ioutil.ReadAll(in)
		err = json.Unmarshal(b, &resp)
		if err != nil {
			log.Fatal(err)
		}

		log.Print(resp)

		data := struct {
			Name   string
			Fields map[string]any
		}{
			structName,
			resp,
		}

		t, err := template.New("struct.tpl").Funcs(template.FuncMap{
			"Title": strings.Title,
			"TypeOf": func(v any) string {
				if v == nil {
					return "string"
				}
				return strings.ToLower(reflect.TypeOf(v).String())
			},
		}).ParseFiles("template/struct.tpl")
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}

		// Define the output file path
		outputFile := "entity/" + structName + ".go"

		// Create or open the output file
		f, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer f.Close()

		// Execute the template and write the generated code to the output file
		err = t.Execute(f, data)
		if err != nil {
			fmt.Println("Error executing template:", err)
			return
		}

		fmt.Println("Route code generated successfully.")
	} else {
		fmt.Println("Invalid 'type' parameter.")
		return
	}

}

func processTemplate(fileName string, outputFile string) {
	if _, err := os.Stat("jwt"); os.IsNotExist(err) {
		log.Print("not exist folder")
		err = os.Mkdir("jwt", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	outputPath := "jwt/" + outputFile
	fmt.Println("Writing file: ", outputPath)
	fout, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)
	if err != nil {
		log.Fatal(err)
	}
}
