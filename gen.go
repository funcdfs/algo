package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	texttemplate "text/template"
	"time"
	"unicode"

	"github.com/funcdfs/algo/tool/template"
	"github.com/funcdfs/algo/tool/testutil"
)

type (
	// TestCase represents a test case
	TestCase struct {
		Input  string `json:"input"`  // Input data
		Output string `json:"output"` // Expected output
	}

	// InputOutput represents input/output specification
	InputOutput struct {
		Type     string   `json:"type"`     // Data type
		Variable string   `json:"variable"` // Variable name
		Pattern  string   `json:"pattern"`  // Pattern
		Lines    []string `json:"lines"`    // Line data
	}

	// Problem represents problem information received from Competitive Companion
	Problem struct {
		Name        string      `json:"name"`        // Problem name
		Group       string      `json:"group"`       // Contest/group name
		URL         string      `json:"url"`         // Problem URL
		MemoryLimit int         `json:"memoryLimit"` // Memory limit (MB)
		TimeLimit   int         `json:"timeLimit"`   // Time limit (ms)
		Tests       []TestCase  `json:"tests"`       // Sample test cases
		TestType    string      `json:"testType"`    // Test type
		Input       InputOutput `json:"input"`       // Input specification
		Output      InputOutput `json:"output"`      // Output specification
	}

	// GenerateTestTemplateData template data structure
	GenerateTestTemplateData struct {
		Package     string
		Tests       []TestCase
		Problem     Problem
		CurrentTime string
		URL         string
	}

	// Config program configuration
	Config struct {
		TemplateType string
		Port         int
		Problem      Problem
		ShouldListen bool
	}
)

const (
	// Directory related
	voidDir       = "void"        // Void problem directory
	leetcodeDir   = "leetcode"    // LeetCode directory
	codeforcesDir = "cf"          // Codeforces directory
	atcoderDir    = "atcoder"     // AtCoder directory
	othersDir     = "others"      // Other problems directory
	testDir       = "test_folder" // Test files directory

	// Template types
	templateTypeVoid       = "void"       // Void template
	templateTypeSimple     = "simple"     // Simple IO template
	templateTypeMultiTest  = "multitest"  // Multiple test cases template
	templateTypeSingleTest = "singletest" // Single test case template (default)

	// Display related
	colorReset   = "\033[0m"
	colorRed     = "\033[31m" // Error messages
	colorGreen   = "\033[32m" // Success messages
	colorYellow  = "\033[33m" // Warnings and borders
	colorBlue    = "\033[34m" // Links and progress indicators
	colorMagenta = "\033[35m" // Numbers and problem IDs
	colorCyan    = "\033[36m" // Titles and paths
	colorGray    = "\033[90m" // Labels

	// Box characters
	boxHorizontal = "─"

	// Network configuration
	port = 10043 // Listening port

	// Field labels
	fieldTitleLabel   = "Problem:"
	fieldContestLabel = "Contest:"
	fieldURLLabel     = "URL:"
	fieldTimeLabel    = "Time:"
	fieldMemoryLabel  = "Memory:"
	fieldTestsLabel   = "Tests:"
	fieldPathLabel    = "Path:"
)

var (
	ErrInvalidTemplate  = errors.New("invalid template type")
	ErrCreateDirectory  = errors.New("failed to create directory")
	ErrGenerateFile    = errors.New("failed to generate file")
	ErrInvalidArgument = errors.New("invalid command line argument")
)

func main() {
	config := NewConfig()

	// Parse command line arguments
	if err := parseArgs(config); err != nil {
		log.Fatal(err)
	}

	// Handle listening mode
	if config.ShouldListen {
		if err := runServer(config); err != nil {
			log.Fatal(err)
		}
		return
	}

	// Non-listening mode: generate file directly
	if err := generateProblemFile(config.TemplateType, config.Problem); err != nil {
		log.Fatal(err)
	}
}

func parseArgs(config *Config) error {
	if len(os.Args) <= 1 {
		config.TemplateType = templateTypeVoid // Use void template by default
		return nil
	}

	arg := os.Args[1]
	switch {
	case arg == "-h" || arg == "--help":
		printUsage()
		os.Exit(0)
	case strings.HasPrefix(arg, "https://leetcode.cn/contest/"):
		config.TemplateType = templateTypeVoid
		config.Problem = genLeetCode(arg)
		config.ShouldListen = false
	case arg == "test":
		config.TemplateType = templateTypeVoid
		config.Problem = newVoidProblem()
		config.ShouldListen = false
	case arg == "atcoder" || arg == "atc" || arg == "st":
		config.TemplateType = templateTypeSingleTest // Use singletest template
		config.ShouldListen = true
	case arg == "cf" || arg == "codeforces" || arg == "mt" || arg == "solve":
		config.TemplateType = templateTypeMultiTest
		config.ShouldListen = true
	case arg == "simple" || arg == "old":
		config.TemplateType = templateTypeSimple
		config.ShouldListen = true
	case arg == "void":
		config.TemplateType = templateTypeVoid
		config.ShouldListen = true
	default:
		return fmt.Errorf("%w: %s", ErrInvalidArgument, arg)
	}
	return nil
}

func runServer(config *Config) error {
	fmt.Printf("\n%s[*]%s Starting server with %s template...\n", colorCyan, colorReset, config.TemplateType)
	fmt.Printf("%s[+]%s Listening on port %s%d%s...\n\n", colorGreen, colorReset, colorMagenta, config.Port, colorReset)

	// Create channels
	done := make(chan os.Signal, 1)
	problemChan := make(chan Problem)

	// Set up signal handling
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	// Start server
	server := newProblemServer(problemChan)
	go func() {
		addr := fmt.Sprintf(":%d", config.Port)
		fmt.Printf("Listening on port %d...\n", config.Port)
		if err := http.ListenAndServe(addr, server); err != nil {
			log.Printf("Server error: %v", err)
			done <- syscall.SIGTERM
		}
	}()

	// Handle problems
	go handleProblems(config.TemplateType, problemChan)

	// Wait for exit signal
	<-done
	fmt.Println("\nShutting down server...")
	return nil
}

func handleProblems(templateType string, problemChan <-chan Problem) {
	for problem := range problemChan {
		if err := generateProblemFile(templateType, problem); err != nil {
			log.Printf("Failed to generate problem files: %v", err)
			continue
		}
		fmt.Println("\nWaiting for next problem...")
	}
}

func generateProblemFile(templateType string, problem Problem) error {
	// Special handling for LeetCode contests
	if strings.Contains(problem.URL, "leetcode.cn/contest/") {
		baseDir := filepath.Join("leetcode", problem.Group)
		problems := []string{"a", "b", "c", "d"}
		
		fmt.Printf("\n%s[*]%s Initializing LeetCode contest...\n", colorCyan, colorReset)
		fmt.Printf("%s[+]%s Target: %s%s%s\n", colorGreen, colorReset, colorYellow, baseDir, colorReset)
		
		for _, p := range problems {
			dirPath := filepath.Join(baseDir, p)
			if err := os.MkdirAll(dirPath, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %v", dirPath, err)
			}
			
			fmt.Printf("%s[>]%s Problem %s%s%s ", colorBlue, colorReset, colorMagenta, p, colorReset)
			
			if err := generateProblemInDir(templateType, dirPath, problem, false); err != nil {
				fmt.Printf("%s[✗]%s\n", colorRed, colorReset)
				return fmt.Errorf("failed to generate files in %s: %v", dirPath, err)
			}
			
			fmt.Printf("%s[✓]%s\n", colorGreen, colorReset)
		}
		
		fmt.Printf("%s[*]%s Generation completed%s\n\n", colorCyan, colorGreen, colorReset)
		return nil
	}

	// Other types of problems
	dirPath := calculateDirPath(problem.URL, problem.Name)
	fmt.Printf("\n%s[*]%s Generating problem files...\n", colorCyan, colorReset)
	fmt.Printf("%s[+]%s Target: %s%s%s\n", colorGreen, colorReset, colorYellow, dirPath, colorReset)
	
	if err := generateProblemInDir(templateType, dirPath, problem, true); err != nil {
		fmt.Printf("%s[✗]%s Failed to generate files\n", colorRed, colorReset)
		return fmt.Errorf("failed to generate problem files: %v", err)
	}
	
	fmt.Printf("%s[✓]%s Files generated successfully\n", colorGreen, colorReset)
	return nil
}

func generateProblemInDir(templateType string, dirPath string, problem Problem, showDetails bool) error {
	// Create directory
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("%w: %v", ErrCreateDirectory, err)
	}

	// Prepare template data
	templateData := GenerateTestTemplateData{
		Package:     filepath.Base(dirPath),
		Tests:       problem.Tests,
		Problem:     problem,
		CurrentTime: time.Now().Format("2006-01-02 15:04:05"),
		URL:         problem.URL,
	}

	// Generate main.go
	mainPath := filepath.Join(dirPath, "main.go")
	var mainTemplate string
	switch templateType {
	case templateTypeVoid:
		mainTemplate = template.VoidTemplate
	case templateTypeMultiTest:
		mainTemplate = template.MultiTestTemplate
	case templateTypeSimple:
		mainTemplate = template.SimpleTemplate
	case templateTypeSingleTest:
		mainTemplate = template.SingleTestTemplate
	default:
		return fmt.Errorf("%w: %s", ErrInvalidTemplate, templateType)
	}

	if err := generateFile(mainPath, mainTemplate, templateData); err != nil {
		return fmt.Errorf("%w: main.go: %v", ErrGenerateFile, err)
	}

	// Generate main_test.go
	testPath := filepath.Join(dirPath, "main_test.go")
	if err := generateFile(testPath, testutil.TestTemplate, templateData); err != nil {
		return fmt.Errorf("failed to generate main_test.go: %v", err)
	}

	// Only print detailed log when needed
	if showDetails {
		fmt.Print(formatLog(problem, dirPath))
	}

	return nil
}

func generateFile(filePath, templateContent string, data interface{}) error {
	// Create file (overwrite if exists)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Printf("Warning: failed to close file: %v", err)
		}
	}(file)

	// Create template with custom functions
	tmpl := texttemplate.New("").Funcs(texttemplate.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	})

	// Parse template
	tmpl, err = tmpl.Parse(templateContent)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	// Execute template
	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	return nil
}

func calculateDirPath(url string, problemName string) string {
	// Clean problem name for default path
	cleanName := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_' {
			return unicode.ToLower(r)
		}
		return '_'
	}, problemName)

	// Extract platform and problem info from URL
	platform, contestID, problemID := extractProblemInfo(url)

	var path string
	switch platform {
	case "codeforces":
		// Extract contest number
		re := regexp.MustCompile(`\d+`)
		contestNum := re.FindString(contestID)
		if contestNum == "" {
			path = filepath.Join(codeforcesDir, "unknown", cleanName)
			break
		}

		// Convert to number to determine range directory
		num := atoi(contestNum)
		rangeStart := (num / 200) * 200
		rangeEnd := rangeStart + 200
		rangeDir := fmt.Sprintf("%d_%d", rangeStart, rangeEnd)

		path = filepath.Join(codeforcesDir, rangeDir, contestNum, fmt.Sprintf("%s_%s", problemID, cleanName))

	case "atcoder":
		if contestID == "" {
			path = filepath.Join(atcoderDir, "unknown", cleanName)
			break
		}
		path = filepath.Join(atcoderDir, strings.ToLower(contestID), fmt.Sprintf("%s_%s", problemID, cleanName))

	case "leetcode":
		if contestID == "" {
			path = filepath.Join(leetcodeDir, "unknown", cleanName)
			break
		}
		path = filepath.Join(leetcodeDir, contestID, fmt.Sprintf("%s_%s", problemID, cleanName))

	default:
		if url == "" {
			timestamp := time.Now().Format("20060102_150405")
			if problemName == "Void Problem" {
				// For test files, use special directory
				path = filepath.Join(testDir, timestamp)
			} else {
				path = filepath.Join(voidDir, fmt.Sprintf("%s_%s", timestamp, cleanName))
			}
		} else {
			path = filepath.Join(othersDir, cleanName)
		}
	}

	// Normalize the last path component
	return normalizeLastPathComponent(path)
}

func normalizeLastPathComponent(path string) string {
	dir, lastPart := filepath.Split(path)
	parts := strings.Split(lastPart, "_")
	if len(parts) < 2 {
		return path // If no underscore, keep original
	}

	// Process first part as problem ID
	first := strings.ToLower(parts[0])
	var problemID string

	// Determine problem ID
	if num, err := strconv.Atoi(first); err == nil {
		problemID = strconv.Itoa(num)
	} else if len(first) >= 1 && unicode.IsLetter([]rune(first)[0]) {
		problemID = string(unicode.ToLower([]rune(first)[0]))
	} else {
		problemID = "a"
	}

	// Check if second part matches problem ID
	if strings.ToLower(parts[1]) == problemID {
		// If matches, remove second part
		parts = append([]string{parts[0]}, parts[2:]...)
	}

	// Normalize name parts
	name := strings.Join(parts[1:], "_")
	name = strings.Map(func(r rune) rune {
		switch {
		case unicode.IsLetter(r):
			return unicode.ToLower(r)
		case unicode.IsNumber(r):
			return r
		default:
			return '_'
		}
	}, name)

	// Merge duplicate underscores and trim
	name = regexp.MustCompile(`_+`).ReplaceAllString(name, "_")
	name = strings.Trim(name, "_")

	// If name is empty, use "problem"
	if name == "" {
		name = "problem"
	}

	// Combine new path
	return filepath.Join(dir, fmt.Sprintf("%s_%s", problemID, name))
}

func newProblemServer(problemChan chan<- Problem) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var problem Problem
		if err := json.NewDecoder(r.Body).Decode(&problem); err != nil {
			http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		problemChan <- problem
		w.WriteHeader(http.StatusOK)
	})
	return mux
}

func newVoidProblem() Problem {
	return Problem{
		Name:        "Void Problem",
		Group:       "Practice",
		URL:         "",
		MemoryLimit: 256,
		TimeLimit:   1000,
		Tests: []TestCase{{
			Input:  "// TODO: Add test case input",
			Output: "// TODO: Add expected output",
		}},
		TestType: "single",
		Input:    InputOutput{Type: "stdin"},
		Output:   InputOutput{Type: "stdout"},
	}
}

func formatLog(problem Problem, dirPath string) string {
	var b strings.Builder
	maxWidth := 68  // Adjust width for tighter fit
	padding := 2
	labelWidth := 10

	// Create horizontal line
	createLine := func(special bool) {
		b.WriteString(strings.Repeat(" ", padding))
		b.WriteString(colorYellow)
		if special {
			b.WriteString(strings.Repeat("═", maxWidth-padding))
		} else {
			b.WriteString(strings.Repeat("─", maxWidth-padding))
		}
		b.WriteString(colorReset)
		b.WriteString("\n")
	}

	writeField := func(label, value, valueColor string) {
		b.WriteString(strings.Repeat(" ", padding))
		b.WriteString(colorCyan)
		b.WriteString("[")
		b.WriteString(colorGray)
		paddedLabel := label + strings.Repeat(" ", labelWidth-len(label))
		b.WriteString(paddedLabel)
		b.WriteString(colorCyan)
		b.WriteString("]")
		b.WriteString(colorReset)
		b.WriteString(" ")
		b.WriteString(valueColor)
		b.WriteString(value)
		b.WriteString(colorReset)
		b.WriteString("\n")
	}

	// Top border and title
	createLine(false)
	b.WriteString(strings.Repeat(" ", padding))
	titleText := ">> Problem Details <<"
	leftPadding := (maxWidth - len(titleText)) / 2
	b.WriteString(strings.Repeat(" ", leftPadding-padding))
	b.WriteString(colorCyan)
	b.WriteString(titleText)
	b.WriteString(colorReset)
	b.WriteString("\n")
	createLine(true)

	// Field information
	writeField("NAME", problem.Name, colorReset)
	writeField("CONTEST", problem.Group, colorReset)
	writeField("URL", problem.URL, colorBlue)
	writeField("TIME", fmt.Sprintf("%d ms", problem.TimeLimit), colorGreen)
	writeField("MEMORY", fmt.Sprintf("%d MB", problem.MemoryLimit), colorGreen)
	writeField("SAMPLES", fmt.Sprintf("%d", len(problem.Tests)), colorMagenta)
	writeField("PATH", dirPath, colorCyan)

	// Bottom border
	createLine(false)
	return b.String()
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func extractProblemInfo(url string) (platform, contestID, problemID string) {
	parts := strings.Split(url, "/")

	switch {
	case strings.Contains(url, "codeforces.com"):
		platform = "codeforces"
		for i, part := range parts {
			if part == "contest" && i+1 < len(parts) {
				contestID = parts[i+1]
			}
			if part == "problem" && i+1 < len(parts) {
				problemID = parts[i+1]
			}
		}

	case strings.Contains(url, "atcoder.jp"):
		platform = "atcoder"
		for i, part := range parts {
			if part == "contests" && i+1 < len(parts) {
				contestID = parts[i+1]
			}
			if part == "tasks" && i+1 < len(parts) {
				problemID = parts[i+1]
			}
		}

	case strings.Contains(url, "leetcode"):
		platform = "leetcode"
		for i, part := range parts {
			if part == "contest" && i+1 < len(parts) {
				contestID = parts[i+1]
			}
			if part == "problems" && i+1 < len(parts) {
				problemID = parts[i+1]
			}
		}
	}

	return
}

func genLeetCode(url string) Problem {
	parts := strings.Split(url, "/")
	var contestID string
	var contestType string
	for i, part := range parts {
		if part == "contest" && i+1 < len(parts) {
			contestID = parts[i+1]
			if strings.Contains(url, "biweekly-contest") {
				contestType = "b"
			} else {
				contestType = "w"
			}
		}
	}

	// Extract pure numeric part
	re := regexp.MustCompile(`\d+`)
	contestNum := re.FindString(contestID)

	// Build directory name: w421 or b141 format
	contestDir := fmt.Sprintf("%s%s", contestType, contestNum)

	return Problem{
		Name:        "LeetCode Contest",
		Group:       contestDir,
		URL:         url,
		MemoryLimit: 256,
		TimeLimit:   1000,
		Tests: []TestCase{{
			Input:  "// TODO: Add test case input",
			Output: "// TODO: Add expected output",
		}},
		TestType: "single",
		Input:    InputOutput{Type: "stdin"},
		Output:   InputOutput{Type: "stdout"},
	}
}

func printUsage() {
	fmt.Printf(`Usage:
    go run gen.go [command]

Commands:
    void               Use void template (default)
    simple, old        Use simple template
    cf, mt, solve      Use multitest template
    atcoder, atc, st   Use singletest template 
    test               Generate test file
    <leetcode_url>     Generate LeetCode contest files
    -h, --help         Show this help message
                       NOTE: 
All commands except test and <leetcode_url> will start a listener server
`)
}

// NewConfig creates default configuration
func NewConfig() *Config {
	return &Config{
		TemplateType:  templateTypeSingleTest,
		Port:         port,
		ShouldListen: true,
	}
}

