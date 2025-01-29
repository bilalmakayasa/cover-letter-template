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
	ID      int
	Value   string
	Message string
}

// Define option lists
var industryOptions = []Option{
	{ID: 1, Value: "e-commerce", Message: "Given my background in backend optimization, system integrations, and cost-efficient engineering solutions, I am excited about the opportunity to contribute to %s. I am particularly interested in how your company leverages technology to enhance eCommerce experiences, optimize logistics, and drive customer engagement. \n\nI would love to explore how my skills in API development, system architecture, and engineering optimization can help streamline and enhance %s’s eCommerce platform."},

	{ID: 2, Value: "saas", Message: "I am excited about the opportunity to contribute to %s, a company that is pushing innovation in the SaaS industry. With 6 years of experience in backend development, system optimization, and third-party integrations, I have a strong foundation in building scalable and high-performance software solutions. \n\nAt my previous roles, I worked on optimizing CI/CD pipelines, reducing deployment times from 40+ seconds to 18 seconds, and integrating critical services like Sentry, Algolia, and Respond.io to enhance performance and observability. Additionally, I have experience handling cross-team collaborations, stepping in to assist teams when engineering resources are stretched. \n\nI am particularly interested in how %s is solving problems in the SaaS space, and I would love to explore how my expertise in system architecture, API integrations, and performance optimizations can contribute to your company’s growth and success."},

	{ID: 3, Value: "Fintech", Message: "I am particularly interested in how %s is transforming the financial sector through technology. With my experience in backend optimization, secure API development, and performance tuning, I am confident I can contribute to %s by building scalable and reliable financial solutions."},

	{ID: 4, Value: "Logistics", Message: "I am excited about the opportunity to contribute to %s by leveraging my experience in API integrations, real-time tracking, and backend optimization. With my experience optimizing Google Maps API costs and building scalable platforms, I believe I can help enhance %s's supply chain efficiency and logistics automation."},

	{ID: 5, Value: "Property Technology", Message: "I see %s as a company that is redefining how people interact with real estate through technology. My background in system architecture, API integrations, and cost-efficient engineering would allow me to contribute to building a seamless and high-performing %s's property platform."},

	{ID: 6, Value: "Medtech", Message: "With the rise of digital transformation in healthcare, %s is at the forefront of building innovative solutions. My experience in building secure, scalable, and optimized backend systems can help enhance %s’s reliability and performance."},
}

var header = "Hi, my name is Bilal. I am a software engineer with 6 years of experience, specializing in backend development. I have worked extensively with Go and JavaScript, and on the frontend, I have experience with Vue and React.js.\n\nPreviously, I worked at a software house in Bandung, where I handled multiple projects for various clients. Later, I joined a Jakarta-based startup called 'WeHelpYou,' a company focused on orchestrating digital platforms to help users easily find service providers. During my time there, I was promoted to Senior Software Engineer after optimizing Google Maps' operational expenses, reducing costs by 30 percent through efficient caching mechanisms. I also managed two junior developers.\n\nMy most recent role was as a Senior Software Engineer at TreeDots Enterprise, where I played a key role in improving engineering efficiency. I optimized the CI/CD workflow pipeline, reducing build times from 40+ seconds to just 18 seconds. Additionally, I integrated multiple third-party services such as Sentry, Algolia, and Respond.io to enhance system performance and observability.\n\n"

var footer = "I appreciate your time and consideration. I look forward to your response. If you are interested for further discussion please contact me at bilal.makayasa@gmail.com"

func main() {
	reader := bufio.NewReader(os.Stdin)

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
	generateMessage(selectedIndustryOption, companyName)

	// Display the final message
	fmt.Println("\nMessage generated successfully.")
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

// WriteToFile writes text content to a file
func writeToFile(filename, content string) error {
	// Create or overwrite the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

// generateMessage creates a personalized message based on user selection
func generateMessage(industry *Option, companyName string) string {
	// Format the message with the company name
	var text = fmt.Sprintf("Dear %s\n\n", companyName+" HR Team,")
	var body = fmt.Sprintf(industry.Message, companyName, companyName)

	// Combined message
	completeMessage := text + header + body + "\n\n" + footer

	// Write the message to a file
	err := writeToFile("output.txt", completeMessage)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Message successfully written to output.txt")
	}

	return completeMessage
}
