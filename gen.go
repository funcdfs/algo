package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"sort"
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
	TestCase struct {
		Input  string `json:"input"`
		Output string `json:"output"`
	}
	// InputOutput represents input/output specification for competitive programming problems
	InputOutput struct {
		Type     string   `json:"type"`     // Data type (e.g. stdin, stdout)
		Variable string   `json:"variable"` // Variable name used in the problem
		Pattern  string   `json:"pattern"`  // Input/output pattern specification
		Lines    []string `json:"lines"`    // Raw line data
	}
	// Problem represents problem metadata received from Competitive Companion browser extension
	Problem struct {
		Name        string      `json:"name"`        // Problem title
		Group       string      `json:"group"`       // Contest or group name
		URL         string      `json:"url"`         // Problem URL
		MemoryLimit int         `json:"memoryLimit"` // Memory limit in MB
		TimeLimit   int         `json:"timeLimit"`   // Time limit in milliseconds
		Tests       []TestCase  `json:"tests"`       // Sample test cases
		TestType    string      `json:"testType"`    // Type of test (single, multiple)
		Input       InputOutput `json:"input"`       // Input format specification
		Output      InputOutput `json:"output"`      // Output format specification
	}
	// GenerateTestTemplateData contains data needed for test file generation
	GenerateTestTemplateData struct {
		Package     string
		Tests       []TestCase
		Problem     Problem
		CurrentTime string
		URL         string
	}

	// Config stores program configuration and runtime settings
	Config struct {
		TemplateType string  // Type of template to use
		Port         int     // Port number for the server
		Problem      Problem // Problem metadata
		ShouldListen bool    // Whether to start the listener server
	}

	// commandInfo stores command information with aliases
	commandInfo struct {
		main    string   // Main command name
		aliases []string // Alternative names/aliases
		desc    string   // Description
	}
)

const (
	// Directory constants
	voidDir       = "void"        // Directory for void/empty problems
	leetcodeDir   = "leetcode"    // Directory for LeetCode problems
	codeforcesDir = "cf"          // Directory for Codeforces problems
	atcoderDir    = "atcoder"     // Directory for AtCoder problems
	acwingDir     = "acwing"      // Directory for Acwing problems
	luoguDir      = "luogu"       // Directory for Luogu problems
	othersDir     = "others"      // Directory for problems from other platforms
	testDir       = "test_folder" // Directory for test files

	// Template type constants
	templateTypeVoid       = "void"       // Basic template without IO handling
	templateTypeSimple     = "simple"     // Simple input/output template
	templateTypeMultiTest  = "multitest"  // Template for multiple test cases
	templateTypeSingleTest = "singletest" // Template for single test case (default)

	// ANSI color codes for terminal output
	colorReset   = "\033[0m"
	colorRed     = "\033[38;2;161;43;43m"
	colorGreen   = "\033[38;2;41;148;75m"
	colorYellow  = "\033[38;2;163;166;124m"
	colorBlue    = "\033[38;2;74;121;217m"
	colorMagenta = "\033[38;2;152;124;166m"
	colorCyan    = "\033[38;2;101;194;167m"
	colorGray    = "\033[38;2;111;111;111m"

	// Custom hex color examples (requires terminal with true color support):
	// For hex color #00FF00 (bright green): "\033[38;2;0;255;0m"
	// For hex color #FF5733 (coral): "\033[38;2;255;87;51m"
	// Format: \033[38;2;R;G;Bm where R, G, B are decimal RGB values (0-255)

	// Network settings
	port = 10043 // Default listening port for Competitive Companion
)

var (
	// ErrInvalidTemplate Error definitions
	ErrInvalidTemplate = errors.New("invalid template type")
	ErrCreateDirectory = errors.New("failed to create directory")
	ErrGenerateFile    = errors.New("failed to generate file")
	ErrInvalidArgument = errors.New("invalid command line argument")

	// Debug mode flag
	debugMode = false
)

// enableDebugMode turns on debugging features
func enableDebugMode() {
	debugMode = true
	fmt.Printf("\n%s[DEBUG MODE ENABLED]%s\n", colorCyan, colorReset)
	fmt.Printf("%s[*]%s Listening for problem data on port %s%d%s...\n",
		colorGreen, colorReset, colorMagenta, colorReset)
	fmt.Printf("%s[*]%s JSON and struct data will be printed to console\n\n",
		colorGreen, colorReset)
}

// debugPrintProblem prints problem data in debug mode
func debugPrintProblem(problem Problem) {
	if !debugMode {
		return
	}

	// Print problem details in a formatted way
	fmt.Printf("\n%s========== RECEIVED PROBLEM DATA ==========%s\n", colorYellow, colorReset)

	// Convert to JSON for display
	jsonData, err := json.MarshalIndent(problem, "", "  ")
	if err != nil {
		fmt.Printf("%s[ERROR]%s Failed to marshal JSON: %v\n", colorRed, colorReset, err)
	} else {
		fmt.Printf("%s[JSON]%s\n%s\n", colorMagenta, colorReset, string(jsonData))
	}

	// Print processed path
	platform, contestID, problemID := extractProblemInfo(problem.URL)
	dirPath := calculateDirPath(problem.URL, problem.Name)

	fmt.Printf("\n%s[EXTRACTED INFO]%s\n", colorMagenta, colorReset)
	fmt.Printf("Platform: %s%s%s\n", colorCyan, platform, colorReset)
	fmt.Printf("Contest ID: %s%s%s\n", colorYellow, contestID, colorReset)
	fmt.Printf("Problem ID: %s%s%s\n", colorGreen, problemID, colorReset)
	fmt.Printf("Directory Path: %s%s%s\n", colorBlue, dirPath, colorReset)

	// Determine more info based on the platform
	switch platform {
	case "codeforces":
		if contestID != "" {
			fmt.Printf("Contest URL: %shttps://codeforces.com/contest/%s%s\n", colorBlue, contestID, colorReset)
		}
		if problemID != "" {
			fmt.Printf("Problem URL: %shttps://codeforces.com/contest/%s/problem/%s%s\n",
				colorBlue, contestID, problemID, colorReset)
		}
	case "atcoder":
		if contestID != "" {
			fmt.Printf("Contest URL: %shttps://atcoder.jp/contests/%s%s\n", colorBlue, contestID, colorReset)
		}
		if problemID != "" {
			fmt.Printf("Problem URL: %shttps://atcoder.jp/contests/%s/tasks/%s%s\n",
				colorBlue, contestID, problemID, colorReset)
		}
	case "acwing":
		if problemID != "" {
			fmt.Printf("Problem URL: %shttps://www.acwing.com/problem/content/%s/%s\n",
				colorBlue, problemID, colorReset)
		}
	case "luogu":
		if problemID != "" {
			fmt.Printf("Problem URL: %shttps://www.luogu.com.cn/problem/%s%s\n",
				colorBlue, problemID, colorReset)
		}
	}

	// Display problem metadata
	fmt.Printf("\n%s[PROBLEM METADATA]%s\n", colorMagenta, colorReset)
	fmt.Printf("Name: %s%s%s\n", colorCyan, problem.Name, colorReset)
	fmt.Printf("Group: %s%s%s\n", colorYellow, problem.Group, colorReset)
	fmt.Printf("Time Limit: %s%d ms%s\n", colorGreen, problem.TimeLimit, colorReset)
	fmt.Printf("Memory Limit: %s%d MB%s\n", colorGreen, problem.MemoryLimit, colorReset)
	fmt.Printf("Test Cases: %s%d%s\n", colorMagenta, len(problem.Tests), colorReset)

	fmt.Printf("\n%s==========================================%s\n\n", colorYellow, colorReset)
}

func main() {
	config := NewConfig()

	// Parse command line arguments
	if err := parseArgs(config); err != nil {
		log.Fatal(err)
	}

	// Either start the listener server or generate files directly
	if config.ShouldListen {
		if err := runServer(config); err != nil {
			log.Fatal(err)
		}
	} else {
		// Generate files using the specified template type
		if err := generateProblemFile(config.TemplateType, config.Problem); err != nil {
			log.Fatal(err)
		}
	}
	return
}

// NewConfig creates a new configuration with default settings
func NewConfig() *Config {
	return &Config{
		TemplateType: templateTypeSingleTest,
		Port:         port,
		ShouldListen: true,
	}
}

// findSimilarCommands finds commands that are similar to the given input
func findSimilarCommands(input string) []string {
	// Define commands with their aliases and descriptions
	commands := []commandInfo{
		{main: "void", aliases: nil, desc: "Basic template without IO handling"},
		{main: "simple", aliases: []string{"pure", "old"}, desc: "Simple template for interview"},
		{main: "codeforces", aliases: []string{"cf", "c"}, desc: "Multitest template"},
		{main: "atcoder", aliases: []string{"atc", "a"}, desc: "Single test template"},
		{main: "test", aliases: []string{"t"}, desc: "Generate test file"},
	}

	type scoredCommand struct {
		name  string
		score float64
	}

	var candidates []scoredCommand
	input = strings.ToLower(strings.TrimSpace(input))

	// If input is empty, return all main commands
	if input == "" {
		var result []string
		for _, cmd := range commands {
			result = append(result, cmd.main)
		}
		return result
	}

	// Calculate scores for each command and its aliases
	for _, cmd := range commands {
		// Check main command
		mainScore := calculateSimilarity(input, cmd.main)
		if mainScore > 0 {
			candidates = append(candidates, scoredCommand{cmd.main, mainScore})
		}

		// Check aliases
		for _, alias := range cmd.aliases {
			aliasScore := calculateSimilarity(input, alias)
			if aliasScore > 0 {
				candidates = append(candidates, scoredCommand{alias, aliasScore})
			}
		}
	}

	// Sort candidates by score in descending order
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].score > candidates[j].score
	})

	// Get unique commands with highest scores (up to 3)
	seen := make(map[string]bool)
	var result []string
	for _, c := range candidates {
		if !seen[c.name] && len(result) < 3 {
			seen[c.name] = true
			result = append(result, c.name)
		}
	}

	return result
}

// calculateSimilarity returns a similarity score between input and command
func calculateSimilarity(input, command string) float64 {
	// Exact match gets highest score
	if input == command {
		return 1.0
	}

	// Prefix match gets high score
	if strings.HasPrefix(command, input) {
		return 0.9
	}

	// Calculate Levenshtein distance for strings of similar length
	if math.Abs(float64(len(input)-len(command))) <= 3 {
		distance := levenshteinDistance(input, command)
		maxLen := math.Max(float64(len(input)), float64(len(command)))
		similarity := 1.0 - float64(distance)/maxLen
		if similarity > 0.5 { // Only consider if similarity is high enough
			return 0.7 * similarity
		}
	}

	// Check for common characters (weighted by position)
	commonScore := 0.0
	inputRunes := []rune(input)
	cmdRunes := []rune(command)

	for i, ch := range inputRunes {
		for j, cmdCh := range cmdRunes {
			if ch == cmdCh {
				// Characters matching in the same position get higher score
				positionWeight := 1.0 - math.Abs(float64(i-j))/float64(len(command))
				commonScore += positionWeight / float64(len(input))
			}
		}
	}

	if commonScore > 0.3 { // Only consider if score is high enough
		return 0.5 * commonScore
	}

	return 0
}

// levenshteinDistance calculates the minimum number of single-character edits
func levenshteinDistance(s1, s2 string) int {
	if len(s1) == 0 {
		return len(s2)
	}
	if len(s2) == 0 {
		return len(s1)
	}

	// Create matrix
	matrix := make([][]int, len(s1)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(s2)+1)
	}

	// Initialize first row and column
	for i := 0; i <= len(s1); i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len(s2); j++ {
		matrix[0][j] = j
	}

	// Fill in the rest of the matrix
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = min(
					matrix[i-1][j]+1,   // deletion
					matrix[i][j-1]+1,   // insertion
					matrix[i-1][j-1]+1, // substitution
				)
			}
		}
	}

	return matrix[len(s1)][len(s2)]
}

// min returns the minimum of three integers
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func parseArgs(config *Config) error {
	if len(os.Args) <= 1 {
		config.TemplateType = templateTypeVoid // Default to void template
		return nil
	}

	arg := os.Args[1]

	switch {
	case arg == "-h" || arg == "--help": // Display help information
		printUsage()
		os.Exit(0)
	case arg == "debug": // Enable debug mode
		config.TemplateType = templateTypeSingleTest
		config.ShouldListen = true
		// Set global debug flag
		enableDebugMode()
		return nil
	case strings.HasPrefix(arg, "https://leetcode.cn/contest/"): // Handle LeetCode contest URL
		config.TemplateType = templateTypeVoid
		config.Problem = newLeetCodeProblem(arg)
		config.ShouldListen = false
	case arg == "test" || arg == "t": // Generate test file
		config.TemplateType = templateTypeVoid
		config.Problem = newVoidProblem()
		config.ShouldListen = false
	case arg == "atcoder" || arg == "atc" || arg == "a": // AtCoder problems use single test template
		config.TemplateType = templateTypeSingleTest
		config.ShouldListen = true
	case arg == "codeforces" || arg == "cf" || arg == "c": // Codeforces problems use multiple test template
		config.TemplateType = templateTypeMultiTest
		config.ShouldListen = true
	case arg == "pure" || arg == "old" || arg == "simple": // Simple template for interview or legacy platforms
		config.TemplateType = templateTypeSimple
		config.ShouldListen = true
	case arg == "void": // Basic template without IO handling
		config.TemplateType = templateTypeVoid
		config.ShouldListen = true
	default:
		// Find similar commands
		similar := findSimilarCommands(arg)
		if len(similar) > 0 {
			suggestions := strings.Join(similar, ", ")
			return fmt.Errorf("%w: %s\n%sDid you mean:%s %s%s%s?",
				ErrInvalidArgument,
				arg,
				colorYellow,
				colorReset,
				colorCyan,
				suggestions,
				colorReset,
			)
		}
		return fmt.Errorf("%w: %s\nUse -h or --help to see available commands", ErrInvalidArgument, arg)
	}
	return nil
}

// printUsage displays help information about available commands
func printUsage() {
	fmt.Printf(`Usage:
    go run gen.go [command]

Commands:
    void                  Use void template (default)
    simple, pure, old     Use simple template
    codeforces, cf, c     Use multitest template
    atcoder, atc, a       Use singletest template 
    test, t               Generate test file
    debug                 Start in debug mode (prints JSON data)
    <leetcode_url>        Generate LeetCode contest files
    -h, --help            Show this help message

NOTE: All commands except test and URL-based commands will start a listener server
`)
}

// newLeetCodeProblem creates a Problem instance for LeetCode contests
func newLeetCodeProblem(url string) Problem {
	parts := strings.Split(url, "/")
	var contestID string
	var contestType string
	for i, part := range parts {
		if part == "contest" && i+1 < len(parts) {
			contestID = parts[i+1]
			if strings.Contains(url, "biweekly-contest") {
				contestType = "biweek"
			} else {
				contestType = "week"
			}
		}
	}
	re := regexp.MustCompile(`\d+`)
	contestNum := re.FindString(contestID)
	contestDir := fmt.Sprintf("%s%s", contestType, contestNum)

	return Problem{
		Name:        "LeetCode",
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

// newVoidProblem creates a Problem instance for void/empty problems
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

func runServer(config *Config) error {
	fmt.Printf("\n%s[*]%s Starting with %s[%s\033[1m%s%s\033[22m%s]%s template...\n",
		colorCyan, colorReset, colorGreen, colorGreen, config.TemplateType, colorReset, colorGreen, colorReset)
	fmt.Printf("%s[+]%s Listening on port %s\033[1m%d%s\033[22m...\n\n",
		colorGreen, colorReset, colorMagenta, config.Port, colorReset)

	// Create buffered channel for OS signals
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	// Create unbuffered channel for problem data
	problemChan := make(chan Problem)
	server := newProblemServer(problemChan)

	// Start HTTP server in a goroutine
	go func() {
		addr := fmt.Sprintf(":%d", config.Port)
		fmt.Printf("Listening on port %d...\n", config.Port)
		if err := http.ListenAndServe(addr, server); err != nil {
			log.Printf("Server error: %v", err)
			done <- syscall.SIGTERM
		}
	}()

	// Process incoming problems in a separate goroutine
	go handleProblems(config.TemplateType, problemChan)

	// Wait for shutdown signal
	<-done
	fmt.Println("\nShutting down server...")
	return nil
}

// newProblemServer creates an HTTP handler for receiving problem data
func newProblemServer(problemChan chan<- Problem) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		// Debug mode: Read and log the raw JSON
		if debugMode {
			// Read the body first for logging
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusBadRequest)
				return
			}

			// Log the raw JSON
			fmt.Printf("\n%s[RAW JSON RECEIVED]%s\n%s\n\n",
				colorCyan, colorReset, string(body))

			// Create a new reader from the body for json.Decoder
			r.Body = io.NopCloser(bytes.NewBuffer(body))
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

// handleProblems processes incoming problems from the channel
func handleProblems(templateType string, problemChan <-chan Problem) {
	for problem := range problemChan {
		// Print debug information if debug mode is enabled
		debugPrintProblem(problem)

		if err := generateProblemFile(templateType, problem); err != nil {
			log.Printf("Failed to generate problem files: %v", err)
			continue
		}
		fmt.Printf("\n%s[=]%s %sWaiting for next problem...%s\n", colorBlue, colorReset, colorGray, colorReset)
	}
}

// generateProblemFile creates problem files based on template type and problem data
func generateProblemFile(templateType string, problem Problem) error {
	// Handle LeetCode contest URLs specially
	if strings.Contains(problem.URL, "leetcode.cn/contest/") {
		baseDir := filepath.Join(leetcodeDir, problem.Group)
		problems := []string{"a", "b", "c", "d"}

		fmt.Printf("\n%s[*]%s Target: %s\033[1m%s%s\033[22m\n", colorCyan, colorReset, colorGreen, baseDir, colorReset)

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

		fmt.Printf("%s[*]%s %sGeneration completed%s\n\n", colorCyan, colorReset, colorGreen, colorReset)
		return nil
	} else {
		// Handle other problem types
		dirPath := calculateDirPath(problem.URL, problem.Name)
		fmt.Printf("\n%s[*]%s Target: %s\033[1m%s%s\033[22m\n", colorCyan, colorReset, colorGreen, dirPath, colorReset)
		if err := generateProblemInDir(templateType, dirPath, problem, true); err != nil {
			fmt.Printf("%s[✗]%s Failed to generate files\n", colorRed, colorReset)
			return fmt.Errorf("failed to generate problem files: %v", err)
		}
		fmt.Printf("%s[✓]%s %sFiles generated successfully%s\n", colorGreen, colorReset, colorGreen, colorReset)
		return nil
	}
}

// generateProblemInDir creates problem files in the specified directory
func generateProblemInDir(templateType string, dirPath string, problem Problem, showDetails bool) error {
	// Create directory if it doesn't exist
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

	// Generate main.go with appropriate template
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

	// Generate test file
	testPath := filepath.Join(dirPath, "main_test.go")
	if err := generateFile(testPath, testutil.TestTemplate, templateData); err != nil {
		return fmt.Errorf("failed to generate main_test.go: %v", err)
	}

	// Print detailed log if requested
	if showDetails {
		fmt.Print(formatLog(problem, dirPath))
	}

	return nil
}

// generateFile creates a file from a template with the given data
func generateFile(filePath, templateContent string, data interface{}) error {
	// Create or overwrite the file
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
	}) // for testutil package

	// Parse and execute template
	tmpl, err = tmpl.Parse(templateContent)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	return nil
}

// formatLog creates a formatted log message for problem details
func formatLog(problem Problem, dirPath string) string {
	var b strings.Builder
	maxWidth := 60 // Increased width for more details
	padding := 2

	// Create box drawing characters
	const (
		topLeft     = "╔"
		topRight    = "╗"
		bottomLeft  = "╚"
		bottomRight = "╝"
		horizontal  = "═"
		vertical    = "║"
		teeRight    = "╠"
		teeLeft     = "╣"
	)

	// ANSI escape codes for text formatting
	const (
		boldOn  = "\033[1m"
		boldOff = "\033[22m"
	)

	// Helper function to count display width of a string (Chinese characters count as 2)
	runeDisplayWidth := func(r rune) int {
		if unicode.Is(unicode.Han, r) || unicode.Is(unicode.Hiragana, r) ||
			unicode.Is(unicode.Katakana, r) || unicode.Is(unicode.Hangul, r) {
			return 2
		}
		return 1
	}

	// Calculate display width of string considering Chinese characters
	stringDisplayWidth := func(s string) int {
		width := 0
		for _, r := range s {
			width += runeDisplayWidth(r)
		}
		return width
	}

	// Truncate long URL with ellipsis
	truncateURL := func(url string, maxLen int) string {
		if stringDisplayWidth(url) <= maxLen {
			return url
		}

		// Find the last occurrence of "problem" or similar keywords
		keywords := []string{"/problem/", "/contest/", "/tasks/"}
		lastKeywordPos := -1

		for _, keyword := range keywords {
			if pos := strings.LastIndex(url, keyword); pos > lastKeywordPos {
				lastKeywordPos = pos
			}
		}

		if lastKeywordPos == -1 {
			// If no keyword found, use simple truncation
			var truncated string
			runWidth := 0
			for _, r := range url {
				w := runeDisplayWidth(r)
				if runWidth+w > maxLen-3 {
					break
				}
				truncated += string(r)
				runWidth += w
			}
			return truncated + "..."
		}

		// Keep the domain and the last important part
		domainEnd := strings.Index(url, "//")
		if domainEnd == -1 {
			domainEnd = 0
		} else {
			domainEnd += 2
		}

		nextSlash := strings.Index(url[domainEnd:], "/")
		if nextSlash == -1 {
			nextSlash = len(url[domainEnd:])
		}
		domain := url[:domainEnd+nextSlash]

		importantPart := url[lastKeywordPos:]

		// Calculate available space
		availSpace := maxLen - stringDisplayWidth(domain) - 3 // 3 for "..."
		var truncatedImportant string
		runWidth := 0
		for _, r := range importantPart {
			w := runeDisplayWidth(r)
			if runWidth+w > availSpace {
				break
			}
			truncatedImportant += string(r)
			runWidth += w
		}

		return domain + "..." + truncatedImportant
	}

	// Create horizontal separator line
	createLine := func(special bool) {
		b.WriteString(strings.Repeat(" ", padding))
		b.WriteString(colorYellow)
		if special {
			b.WriteString(topLeft + strings.Repeat(horizontal, maxWidth-padding*2-2) + topRight)
		} else {
			b.WriteString(teeRight + strings.Repeat(horizontal, maxWidth-padding*2-2) + teeLeft)
		}
		b.WriteString(colorReset)
		b.WriteString("\n")
	}

	// Write a field with label and value
	writeField := func(label, value, valueColor string, bold bool) {
		b.WriteString(strings.Repeat(" ", padding))
		b.WriteString(colorYellow)
		b.WriteString(vertical)
		b.WriteString(colorReset)
		b.WriteString(" ")

		// Make label brighter and bolder
		b.WriteString("\033[1m")
		b.WriteString(colorCyan)
		b.WriteString(label)
		b.WriteString("\033[22m")
		b.WriteString(colorReset)

		// Calculate maximum value length based on available space
		maxValueLen := maxWidth - stringDisplayWidth(label) - padding*2 - 4
		displayValue := value
		if label == "URL" {
			displayValue = truncateURL(value, maxValueLen)
		} else {
			// Truncate if needed based on display width
			var truncated string
			runWidth := 0
			for _, r := range displayValue {
				w := runeDisplayWidth(r)
				if runWidth+w > maxValueLen {
					break
				}
				truncated += string(r)
				runWidth += w
			}
			displayValue = truncated
		}

		// Ensure spacing is never negative
		displayWidth := stringDisplayWidth(displayValue)
		spacing := maxWidth - stringDisplayWidth(label) - displayWidth - padding*2 - 4
		if spacing < 0 {
			spacing = 0
		}
		b.WriteString(strings.Repeat(" ", spacing))

		b.WriteString(valueColor)
		if bold {
			b.WriteString(boldOn)
		}
		b.WriteString(displayValue)
		if bold {
			b.WriteString(boldOff)
		}
		b.WriteString(colorReset)

		b.WriteString(" ")
		b.WriteString(colorYellow)
		b.WriteString(vertical)
		b.WriteString(colorReset)
		b.WriteString("\n")
	}

	// Format timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Build the log message - removed header section with PROBLEM DETAILS
	createLine(true)
	writeField("URL", problem.URL, colorBlue, false)
	writeField("NAME", problem.Name, colorGreen, false)
	writeField("PATH", dirPath, colorCyan, true) // PATH is bold
	writeField("TIME", timestamp, colorGray, false)
	writeField("MEM", fmt.Sprintf("%d MB", problem.MemoryLimit), colorMagenta, false)
	writeField("TIME", fmt.Sprintf("%d ms", problem.TimeLimit), colorMagenta, false)

	// Close the box
	b.WriteString(strings.Repeat(" ", padding))
	b.WriteString(colorYellow)
	b.WriteString(bottomLeft + strings.Repeat(horizontal, maxWidth-padding*2-2) + bottomRight)
	b.WriteString(colorReset)
	b.WriteString("\n")

	return b.String()
}

// cleanInputName sanitizes input strings for use in file/directory names
// Preserves Chinese characters while cleaning up other characters
func cleanInputName(str string) string {
	// Replace spaces and special characters with underscores, but preserve Chinese characters
	re := regexp.MustCompile(`[^\p{Han}\p{Latin}0-9]+`)
	str = re.ReplaceAllString(str, "_")

	// Replace three or more consecutive underscores with two
	reUnderscores := regexp.MustCompile(`_{3,}`)
	str = reUnderscores.ReplaceAllString(str, "__")

	// Trim leading and trailing underscores
	str = strings.Trim(str, "_")

	return str
}

// extractProblemInfo extracts platform-specific information from problem URLs
func extractProblemInfo(url string) (platform, contestID, problemID string) {
	parts := strings.Split(url, "/")

	switch {
	case strings.Contains(url, "codeforces.com"):
		platform = "codeforces"
		for i, part := range parts {
			if part == "contest" && i+1 < len(parts) {
				contestID = parts[i+1]
			} else if part == "problemset" && i+2 < len(parts) {
				contestID = parts[i+2]
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

	case strings.Contains(url, "vjudge.net"):
		platform = "vjudge"

	case strings.Contains(url, "acwing.com"):
		platform = "acwing"
		// Extract problem ID from URL like https://www.acwing.com/problem/content/1/
		for i, part := range parts {
			if part == "problem" && i+2 < len(parts) && parts[i+1] == "content" {
				problemID = strings.TrimSuffix(parts[i+2], "/")
			}
		}

	case strings.Contains(url, "luogu.com.cn"):
		platform = "luogu"
		// Extract problem ID from URL like https://www.luogu.com.cn/problem/B2002
		for i, part := range parts {
			if part == "problem" && i+1 < len(parts) {
				problemID = parts[i+1]
			}
		}

	case strings.Contains(url, "nowcoder.com"):
		platform = "nowcoder"
	}

	return
}

// calculateDirPath determines the appropriate directory path for a problem
func calculateDirPath(url string, problemName string) string {
	cleanName := cleanInputName(problemName)
	platform, contestID, problemID := extractProblemInfo(url)
	var problemPathStr string

	switch platform {
	case "codeforces":
		re := regexp.MustCompile(`\d+`)
		contestNum := re.FindString(contestID)
		if contestNum == "" {
			problemPathStr = filepath.Join(codeforcesDir, "unknown", cleanName)
			break
		}
		num, _ := strconv.Atoi(contestNum)
		rangeStart := (num / 200) * 200
		rangeEnd := rangeStart + 200
		rangeDir := fmt.Sprintf("%d_%d", rangeStart, rangeEnd)
		problemPathStr = filepath.Join(codeforcesDir, rangeDir, contestNum, fmt.Sprintf("%s_%s", problemID, cleanName))

	case "atcoder":
		if contestID == "" {
			problemPathStr = filepath.Join(atcoderDir, "unknown", cleanName)
			break
		}
		problemPathStr = filepath.Join(atcoderDir, strings.ToLower(contestID), fmt.Sprintf("%s", cleanName))

	case "acwing":
		// Create a dedicated acwing directory with problem ID as prefix
		if problemID == "" {
			problemPathStr = filepath.Join(acwingDir, cleanName)
		} else {
			problemPathStr = filepath.Join(acwingDir, fmt.Sprintf("%s_%s", problemID, cleanName))
		}

	case "luogu":
		// Create a dedicated luogu directory with problem ID as prefix
		// luogu's name already contains the problemID infomation so, the process is same
		if problemID == "" {
			problemPathStr = filepath.Join(luoguDir, cleanName)
		} else {
			problemPathStr = filepath.Join(luoguDir, fmt.Sprintf("%s", cleanName))
		}

	case "vjudge":
		problemPathStr = filepath.Join(othersDir, "vjudge", cleanName)

	case "nowcoder":
		problemPathStr = filepath.Join(othersDir, "nowcoder", cleanName)

	default:
		if url == "" {
			timestamp := time.Now().Format("20060102_150405")
			if problemName == "Void Problem" {
				problemPathStr = filepath.Join(testDir, timestamp)
			} else {
				problemPathStr = filepath.Join(voidDir, fmt.Sprintf("%s_%s", timestamp, cleanName))
			}
		} else {
			problemPathStr = filepath.Join(othersDir, cleanName)
		}
	}

	return problemPathStr
}
