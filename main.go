package main

import (
  "bufio"
  "encoding/csv"
  "fmt"
  "net"
  "os"
  "regexp"
  "strings"
  "sync"
)

// --- 1.DATA MODELS ---

type ValidationResult struct {
      Email            string
      SyntaxValid      bool
      MXValid          bool
      Deliverable      bool
  }

// --- 2.ENTRY POINT ---

func main() {
  fmt.Println("=== Bulk Email Validator ===")

  //1. Read input file 
  emails := loadEmailsFromFile("emails.txt")
  if len(emails) == 0 {
    fmt.Println("No emails found to process. Exiting.")
    os.Exit(0)
  }

  fmt.Printf("Found %d emails to process. \n\n", len(emails))
  fmt.Println("Starting concurrent validation...")

  //2. Process concurrently and time the Execution
  results := processEmailsConcurrent(emails)

  fmt.Printf("\nValidation complete. Processed %d emails\n\n", len(results))

  //3.Export results
  fmt.Println("Exporting report to results.csv...")
  err := exportTOCSV(results, "results.csv")
  if err != nil {
    fmt.Printf("Error saving report: %v\n", err)
    os.Exit(1)
  }

  fmt.Println("Success. Report saved.")
}



// ---3. SETUP AND FILE I/O ---
//loadEmailsFromFile reads a text file and returns a slice of strings
func loadEmailsFromFile(filename string) []string {
  file, err := os.Open(filename)
  if err != nil {
    fmt.Printf("Error opening file: %v\n", err)
    os.Exit(1)
  }
  defer file.Close()

  var emails []string
  scanner :=bufio.NewScanner(file)

  for scanner.Scan() {
    line := strings.TrimSpace(scanner.Text())
    if line != "" {
      emails = append(emails, line)
    }
}
return emails
}



// --- 4.VALIDATION LOGIC ---

//ValidateEmail orchestrates the syntax and domain checks.
func validateEmail(email string) ValidationResult {
  





































    
























  









              
