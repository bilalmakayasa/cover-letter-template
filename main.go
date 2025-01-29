package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Option struct
type Option struct {
	ID    int
	Value string
}

// Define option lists
var industryOptions = []Option{
	{ID: 1, Value: "e-commerce"},
	{ID: 2, Value: "saas"},
}

var nameOptions = []Option{
	{ID: 1, Value: "Bilal"},
	{ID: 2, Value: "Wiya"},
}

var header = []Option{
	{ID: 1, Value: "Hi, my name is Bilal. I am a software engineer with 6 years of experience. I work mostly on backend systems. Worked with Go, and Javascript. Previosuly worked at software house company in Bandung, handle a lot of projects from various clients. Then I was working on a startup company in Jakarta called 'WeHelpYou'. This company focus on orchestrating digital platform so user can easilty find a service provide, I was promoted to senior in this company after I saved the operation expense for google maps, I was cutting 40% by implementing user location using cache, managing 2 junior developers. My last role was senior engineer in Treedots enterprise, I optimize the CI/CD workflow pipeline from 40+ s to 18s."},
	{ID: 2, Value: "Hi, my name is Nuwiya. I am a product designer."},
}

var body = []Option{
	{ID: 1, Value: "I am reaching out to you because I am interested in your company."},
	{ID: 2, Value: "I am reaching out to you because I am interested in your company 2."},
}

var footer = "I hope to hear from you soon."

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Select Name
	selectedNameOption := getUserSelection(reader, "Please select your name:", nameOptions)
	if selectedNameOption == nil {
		fmt.Println("Invalid selection. Exiting...")
		return
	}

	// Select Industry
	selectedIndustryOption := getUserSelection(reader, "Please select your industry:", industryOptions)
	if selectedIndustryOption == nil {
		fmt.Println("Invalid selection. Exiting...")
		return
	}

	// Get company name
	fmt.Print("\nPlease input the company name: ")
	companyName, _ := reader.ReadString('\n')
	companyName = strings.TrimSpace(companyName)

	// Generate message
	message := generateMessage(selectedNameOption.ID, selectedIndustryOption.ID, companyName)

	// Display the final message
	fmt.Println("\n----------------------Generated Message--------------------------\n")
	fmt.Println(message)
}

// getUserSelection displays options and retrieves user selection
func getUserSelection(reader *bufio.Reader, prompt string, options []Option) *Option {
	fmt.Println("\n" + prompt)
	for _, option := range options {
		fmt.Printf("[%d] %s\n", option.ID, option.Value)
	}

	// Read user input
	fmt.Print("Enter the option number: ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)

	// Convert input to integer
	choice, err := strconv.Atoi(choiceStr)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return nil
	}

	// Find the matching option
	for _, option := range options {
		if option.ID == choice {
			return &option
		}
	}

	fmt.Println("Invalid choice. Please select a valid option.")
	return nil
}

// generateMessage creates a personalized message based on user selection
func generateMessage(nameID, industryID int, companyName string) string {
	var text = fmt.Sprintf("Dear %s\n\n", companyName+" HR Team,")

	// Get header based on name selection
	for _, option := range header {
		if option.ID == nameID {
			text += option.Value + "\n\n"
			break
		}
	}

	// Get body based on industry selection
	for _, option := range body {
		if option.ID == industryID {
			text += option.Value + "\n\n"
			break
		}
	}

	text += footer
	return text
}
