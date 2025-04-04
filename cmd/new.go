package cmd

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var embeddedTemplates embed.FS

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Gin project",
	Long: `Create a new Gin project with a standard structure.
Example:
  gincraft new myapp
  gincraft new myapp --module github.com/username/myapp`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		if err := validateProjectName(projectName); err != nil {
			fmt.Println("❌ Error:", err)
			return
		}
		if err := scaffoldProject(projectName); err != nil {
			fmt.Println("❌ Error:", err)
		} else {
			fmt.Printf("✅ Project created successfully: %s\n", projectName)
			fmt.Println("\n📁 Project structure:")
			fmt.Println("  ├── controllers/")
			fmt.Println("  ├── routes/")
			fmt.Println("  ├── services/")
			fmt.Println("  ├── models/")
			fmt.Println("  ├── main.go")
			fmt.Println("  ├── go.mod")
			fmt.Println("  ├── .gitignore")
			fmt.Println("  └── README.md")
			fmt.Println("\n🚀 Next steps:")
			fmt.Println("  1. cd", projectName)
			fmt.Println("  2. go mod tidy")
			fmt.Println("  3. go run main.go")
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func validateProjectName(name string) error {
	if len(name) < 3 {
		return fmt.Errorf("project name must be at least 3 characters long")
	}
	if strings.ContainsAny(name, "\\/:*?\"<>|") {
		return fmt.Errorf("project name contains invalid characters")
	}
	return nil
}

func scaffoldProject(name string) error {
	projectDir := filepath.Join(".", name)
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Initialize the module first
	initCmd := exec.Command("go", "mod", "init", name)
	initCmd.Dir = projectDir
	output, err := initCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to initialize module: %s\n%s", err, string(output))
	}

	// Create README.md from template
	tmpl, err := template.ParseFS(embeddedTemplates, "templates/README.md.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse README template: %w", err)
	}

	readmeFile, err := os.Create(filepath.Join(projectDir, "README.md"))
	if err != nil {
		return fmt.Errorf("failed to create README.md: %w", err)
	}
	defer readmeFile.Close()

	if err := tmpl.Execute(readmeFile, map[string]string{"ProjectName": name}); err != nil {
		return fmt.Errorf("failed to execute README template: %w", err)
	}

	// Create .gitignore from template
	tmpl, err = template.ParseFS(embeddedTemplates, "templates/.gitignore.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse .gitignore template: %w", err)
	}

	gitignoreFile, err := os.Create(filepath.Join(projectDir, ".gitignore"))
	if err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}
	defer gitignoreFile.Close()

	if err := tmpl.Execute(gitignoreFile, map[string]string{"ProjectName": name}); err != nil {
		return fmt.Errorf("failed to execute .gitignore template: %w", err)
	}

	// Create main.go from template
	tmpl, err = template.ParseFS(embeddedTemplates, "templates/main.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse main template: %w", err)
	}

	mainFile, err := os.Create(filepath.Join(projectDir, "main.go"))
	if err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}
	defer mainFile.Close()

	if err := tmpl.Execute(mainFile, map[string]string{"ProjectName": name}); err != nil {
		return fmt.Errorf("failed to execute main template: %w", err)
	}

	// Create basic folders and copy template files
	templateFiles := map[string]string{
		"controllers": "example_controller.go",
		"routes":     "routes.go",
		"services":   "example_service.go",
		"models":     "example.go",
	}

	for dir, file := range templateFiles {
		dirPath := filepath.Join(projectDir, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}

		// Read template file from embedded filesystem
		content, err := embeddedTemplates.ReadFile(filepath.Join("templates", dir, file))
		if err != nil {
			return fmt.Errorf("failed to read template file %s: %w", file, err)
		}

		// Replace template variables
		contentStr := string(content)
		contentStr = strings.ReplaceAll(contentStr, "PROJECT_NAME", name)
		contentStr = strings.ReplaceAll(contentStr, "{{.ProjectName}}", name)

		if err := os.WriteFile(filepath.Join(dirPath, file), []byte(contentStr), 0644); err != nil {
			return fmt.Errorf("failed to write template file %s: %w", file, err)
		}
	}

	// Add gin dependency
	addCmd := exec.Command("go", "get", "github.com/gin-gonic/gin@latest")
	addCmd.Dir = projectDir
	output, err = addCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to add gin dependency: %s\n%s", err, string(output))
	}

	// Run go mod tidy
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = projectDir
	output, err = tidyCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run go mod tidy: %s\n%s", err, string(output))
	}

	return nil
}
