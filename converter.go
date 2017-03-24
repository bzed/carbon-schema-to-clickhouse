package main

import (
	"flag"
	"fmt"
	"github.com/lomik/go-carbon/persister"
	"os"
)

func printSchemaXML(schema persister.Schema, rollupFunction string, default_schema bool) {

	pattern_text := "pattern"
	if default_schema {
		pattern_text = "default"
	}
	fmt.Printf("\t<!-- %s -->\n", schema.Name)
	fmt.Printf("\t<%s>\n", pattern_text)
	if !default_schema {
		fmt.Printf("\t\t<regexp>%s</regexp>\n", schema.Pattern)
	}
	fmt.Printf("\t\t<function>%s</function>\n", rollupFunction)
	lastAge := 0
	for j := 0; j < len(schema.Retentions); j++ {
		fmt.Printf("\t\t<%s>\n", "retention")
		retention := schema.Retentions[j]
		fmt.Printf("\t\t\t<age>%d</age>\n", lastAge)
		lastAge = retention.NumberOfPoints()
		fmt.Printf("\t\t\t<precision>%d</precision>\n", retention.SecondsPerPoint())
		fmt.Printf("\t\t</%s>\n", "retention")
	}
	fmt.Printf("\t</%s>\n", pattern_text)
}

func main() {
	schemaFile := flag.String(
		"schemafile",
		"/etc/carbon/storage-schemas.conf",
		"storage schema file to convert")
	rollupFunction := flag.String(
		"rollupfunction",
		"any",
		"clickhouse rollup function")
	flag.Parse()
	schemas, err := persister.ReadWhisperSchemas(*schemaFile)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Println("<graphite_rollup>")
	default_schema := -1

	for i := 0; i < len(schemas); i++ {
		schema := schemas[i]
		if schema.Name == "default" {
			default_schema = i
			continue
		}
		printSchemaXML(schema, *rollupFunction, false)
	}
	if default_schema >= 0 {
		schema := schemas[default_schema]
		printSchemaXML(schema, *rollupFunction, true)
	}
	fmt.Println("</graphite_rollup>")
}
