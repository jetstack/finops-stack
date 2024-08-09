package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/lint/rules"
	"helm.sh/helm/v3/pkg/lint/support"
	"helm.sh/helm/v3/pkg/repo"
)

var (
	logFormat string
	logLevel  string
	logger    = logrus.New()
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "chart-dep-updater [Chart.yaml]",
		Short: "A tool to check for updates in Helm chart dependencies",
		Args:  cobra.ExactArgs(1),
		Run:   run,
	}

	rootCmd.Flags().StringVarP(&logFormat, "log-format", "f", "text", "Log format: text or json")
	rootCmd.Flags().StringVarP(&logLevel, "log-level", "l", "info", "Log level: debug, info, warn, error")

	if err := rootCmd.Execute(); err != nil {
		logger.Fatalf("Error executing command: %v", err)
	}
}

func run(cmd *cobra.Command, args []string) {
	configureLogger()

	chartPath := args[0]

	// Load the Chart.yaml using yaml.Node to preserve structure
	content, err := ioutil.ReadFile(chartPath)
	if err != nil {
		logger.Fatalf("Error reading Chart.yaml: %v", err)
	}

	var rootNode yaml.Node
	if err := yaml.Unmarshal(content, &rootNode); err != nil {
		logger.Fatalf("Error parsing Chart.yaml: %v", err)
	}

	updated := false

	// Find dependencies node and update versions
	for i, node := range rootNode.Content[0].Content {
		if node.Value == "dependencies" {
			dependencies := rootNode.Content[0].Content[i+1]
			for j := 0; j < len(dependencies.Content); j++ {
				dep := dependencies.Content[j]
				updateNode := false
				var name, version, repo string
				for k := 0; k < len(dep.Content); k += 2 {
					key := dep.Content[k].Value
					value := dep.Content[k+1].Value
					if key == "name" {
						name = value
					}
					if key == "version" {
						version = value
					}
					if key == "repository" {
						repo = value
					}
				}

				latestVersion, err := getLatestVersion(repo, name)
				if err != nil {
					logger.Errorf("Error fetching latest version for %s: %v", name, err)
					continue
				}

				constraint, err := semver.NewConstraint(version)
				if err != nil {
					logger.Errorf("Invalid version constraint for %s: %v", name, err)
					continue
				}

				latestVer, err := semver.NewVersion(latestVersion)
				if err != nil {
					logger.Errorf("Invalid latest version for %s: %v", name, err)
					continue
				}

				if constraint.Check(latestVer) {
					logger.Infof("Dependency %s is up-to-date: %s", name, latestVersion)
				} else {
					logger.Warnf("Dependency %s has a newer version available: %s", name, latestVersion)
					// Update the version in the YAML node
					for k := 0; k < len(dep.Content); k += 2 {
						if dep.Content[k].Value == "version" {
							dep.Content[k+1].Value = latestVersion
							updateNode = true
							break
						}
					}
					updated = updated || updateNode
				}
			}
			break
		}
	}

	if updated {
		// Write the updated content back to Chart.yaml
		out, err := yaml.Marshal(&rootNode)
		if err != nil {
			logger.Fatalf("Error marshaling updated Chart.yaml: %v", err)
		}

		if err := os.WriteFile(chartPath, out, 0644); err != nil {
			logger.Fatalf("Error writing updated Chart.yaml: %v", err)
		}
		logger.Infof("Chart.yaml updated successfully.")

		// Run Helm dependency update
		if err := runHelmDepUpdate(chartPath); err != nil {
			logger.Fatalf("Error running helm dependency update: %v", err)
		}

		// Run Helm lint
		if err := runHelmLint(chartPath); err != nil {
			logger.Fatalf("Helm linting failed: %v", err)
		}
	} else {
		logger.Infof("No updates were necessary.")
	}
}

func configureLogger() {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logger.Fatalf("Invalid log level: %v", err)
	}
	logger.SetLevel(level)

	switch logFormat {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{})
	default:
		logger.Fatalf("Invalid log format: %s", logFormat)
	}
}

func getLatestVersion(repoURL, chartName string) (string, error) {
	// Check if the repoURL uses the file protocol
	if strings.HasPrefix(repoURL, "file://") {
		localPath := strings.TrimPrefix(repoURL, "file://")
		return getLatestLocalVersion(localPath)
	}

	settings := cli.New()

	client := getter.All(settings)
	repoIndex, err := repo.NewChartRepository(&repo.Entry{
		URL: repoURL,
	}, client)
	if err != nil {
		return "", err
	}

	indexFile, err := repoIndex.DownloadIndexFile()
	if err != nil {
		return "", err
	}

	index, err := repo.LoadIndexFile(indexFile)
	if err != nil {
		return "", err
	}

	chartVersions, ok := index.Entries[chartName]
	if !ok || len(chartVersions) == 0 {
		return "", fmt.Errorf("no versions found for chart %s", chartName)
	}

	return chartVersions[0].Version, nil
}

func getLatestLocalVersion(localPath string) (string, error) {
	chart, err := loader.LoadDir(localPath)
	if err != nil {
		return "", fmt.Errorf("failed to load chart from %s: %v", localPath, err)
	}
	return chart.Metadata.Version, nil
}

// LoggerWriter is a custom writer that writes to the logger
type LoggerWriter struct {
	logger *logrus.Logger
}

func (lw *LoggerWriter) Write(p []byte) (n int, err error) {
	lw.logger.Info(string(p))
	return len(p), nil
}

func runHelmDepUpdate(chartPath string) error {
	chartDir := filepath.Dir(chartPath)
	// Create a custom writer for the logger
	loggerWriter := &LoggerWriter{logger: logger}

	man := &downloader.Manager{
		Out:             loggerWriter,
		Verify:          downloader.VerifyIfPossible,
		SkipUpdate:      false,
		ChartPath:       chartDir,
		Getters:         getter.All(cli.New()),
		RepositoryCache: os.TempDir(),
	}

	return man.Update()
}

func runHelmLint(chartPath string) error {
	chartDir := filepath.Dir(chartPath)

	linter := support.Linter{ChartDir: chartDir}
	rules.Chartfile(&linter)
	rules.Dependencies(&linter)

	if len(linter.Messages) > 0 {
		for _, msg := range linter.Messages {
			if msg.Severity == support.ErrorSev {
				return fmt.Errorf("linting error: %s", msg)
			}
			logger.Warnf("lint warning: %s", msg.Err)
		}
	}

	return nil
}
