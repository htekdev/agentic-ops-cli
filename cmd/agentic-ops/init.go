package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize agentic-ops for a repository",
	Long: `Creates the necessary directory structure and hook configuration
for agentic-ops in the current repository.

This command creates:
- .github/agent-workflows/ directory for your workflow files
- .copilot/hooks.json to integrate with Copilot CLI hooks

After running init, you can create workflows using 'agentic-ops create'
or by manually creating YAML files in .github/agent-workflows/`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, _ := cmd.Flags().GetString("dir")
		force, _ := cmd.Flags().GetBool("force")

		if dir == "" {
			var err error
			dir, err = os.Getwd()
			if err != nil {
				return err
			}
		}

		return runInit(dir, force)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("dir", "d", "", "Directory to initialize (default: current directory)")
	initCmd.Flags().BoolP("force", "f", false, "Overwrite existing configuration")
}

func runInit(dir string, force bool) error {
	fmt.Printf("Initializing agentic-ops in %s\n", dir)

	// Create .github/agent-workflows directory
	workflowsDir := filepath.Join(dir, ".github", "agent-workflows")
	if err := os.MkdirAll(workflowsDir, 0755); err != nil {
		return fmt.Errorf("failed to create workflows directory: %w", err)
	}
	fmt.Printf("✓ Created %s\n", workflowsDir)

	// Create .copilot directory
	copilotDir := filepath.Join(dir, ".copilot")
	if err := os.MkdirAll(copilotDir, 0755); err != nil {
		return fmt.Errorf("failed to create .copilot directory: %w", err)
	}

	// Create hooks.json
	hooksFile := filepath.Join(copilotDir, "hooks.json")
	if _, err := os.Stat(hooksFile); err == nil && !force {
		fmt.Printf("⚠ %s already exists (use --force to overwrite)\n", hooksFile)
	} else {
		hooksContent := generateHooksJSON()
		if err := os.WriteFile(hooksFile, []byte(hooksContent), 0644); err != nil {
			return fmt.Errorf("failed to create hooks.json: %w", err)
		}
		fmt.Printf("✓ Created %s\n", hooksFile)
	}

	// Create .gitignore in agent-workflows if it doesn't exist
	gitignorePath := filepath.Join(workflowsDir, ".gitignore")
	if _, err := os.Stat(gitignorePath); os.IsNotExist(err) {
		gitignoreContent := "# Temporary files\n*.tmp\n*.log\n"
		if err := os.WriteFile(gitignorePath, []byte(gitignoreContent), 0644); err != nil {
			// Not critical, just warn
			fmt.Printf("⚠ Could not create .gitignore: %v\n", err)
		}
	}

	fmt.Println("\n✓ agentic-ops initialized successfully!")
	fmt.Println("\nNext steps:")
	fmt.Println("  1. Create a workflow: agentic-ops create \"block edits to .env files\"")
	fmt.Println("  2. Or manually create a workflow in .github/agent-workflows/")
	fmt.Println("  3. Commit the .copilot/hooks.json to enable for your team")

	return nil
}

// generateHooksJSON creates the hooks.json content that integrates with agentic-ops CLI
func generateHooksJSON() string {
	return `{
  "version": 1,
  "hooks": {
    "preToolUse": [
      {
        "type": "command",
        "bash": "agentic-ops run --raw --dir \"$PWD\"",
        "powershell": "agentic-ops run --raw --dir (Get-Location)",
        "timeoutSec": 60
      }
    ]
  }
}
`
}
