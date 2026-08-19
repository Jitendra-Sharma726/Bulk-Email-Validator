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
  result   := ValidationResult{
    Email:    email,
    SyntaxValid:  false,
    MXValid:      false,
    Deliverable:  false,
  }

//1. Check Syntax
if !checkSyntax(email) {
  return result
}
result.SyntaxValid = true

//2. check Domain MX Records
if !checkDomain(email) {
  return result
}
result.MXValid = true

//3. if both pass, it's valid and deliverable
result.Deliverable = true
return result
}



//checkSyntax uses Regex to ensure standard email formatting 
func checkSyntax(email string) bool {
  pattern := `[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
  regex := regexp.MustCompile(pattern)
  return regex.MatchString(email)
}

//checkDomain extracts the domain and queries the network for Mail Exchange (MX) records.
func checkDomain(email string) bool {
parts := strings.Split(email, "@")
if len(parts) !=2 {
  return false
}

domain := parts[1]
  
//Query live DNS records
mxRecords, err := net.LookupMX(domain)
  
if err!= nil {
  return false
}

if len(mxRecords) == 0 {
  return false
}
  
return true
}

// --- 5.CONCURRENT PIPELINE ---

//processEmailCOncurrent uses Goroutines to Validate emails in parallel.
func processEmailsConcurrent(emails []string) []ValidationResult {
  Var results []ValidationResult

  var wg sync.WaitGroup
  var mu sync.Mutex

  for _, email := range emails {
    wg.Add(1)

    //Launch a Goroutine
    go func() {
      defer wg.Done()

      //Perform the network Validation using the loop variable directly
      result := validateEmail(email)

      //Lock the Mutex to safely write to the shared silce
      mu.Lock()
      results = append(results, result)
      mu.Unlock()
    }()
  }

  //Wait for all Goroutines to finish before continuing
  wg.Wait()

  return results
}


// --- 6. EXPORT ---
// exportTOCSV writes the Validation results into a spreadsheet.

func exportToCSV(results []ValidationResult, filename string) error {
  file, err := os.Create(filename)
  if err != nil {
        return err
  }
  defer file.Close()

  writer := csv.NewWriter(file)
  defer writer.Flush()

  //Write the header row
  headers := []string{"Email", "SyntaxValid", "MXValid", "Deliverable"}
  if err := writer.Write(headers); err != nil {
    return err
  }

  //Write Data rows
  for _, res := range results {
    row := []string{
         res.Email,
         fmt.Sprint(res.SyntaxValid),
         fmt.Sprint(res.MXValid),
         fmt.Sprint(res.Deliverable),
    }

    err := writer.Write(row);

    if err != nil {
      return err 
    }
  }

  return nil 

}

    










    















      






























  



  

























  
  













  

  
  





































    
























  









              
