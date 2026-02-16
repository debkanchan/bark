package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/debkanchan/bark/internal/scanner"
)

func TestIntegrationDirty(t *testing.T) {
	// Test using the testdata/dirty directory (files with BARK comments)
	dirtyPath := filepath.Join(".", "testdata", "dirty")

	// Check if testdata/dirty exists
	if _, err := os.Stat(dirtyPath); os.IsNotExist(err) {
		t.Skip("testdata/dirty directory not found")
	}

	s := scanner.NewScanner()
	result := s.Scan(dirtyPath)

	findings := result.GetFindings()
	errors := result.GetErrors()

	// We should have multiple BARK comments across different test files
	if len(findings) == 0 {
		t.Error("Expected to find BARK comments in testdata/dirty")
	}

	t.Logf("Found %d BARK comments in testdata/dirty", len(findings))

	// Log findings for debugging
	for _, finding := range findings {
		t.Logf("  %s", finding.String())
	}

	// Log any errors
	if len(errors) > 0 {
		t.Logf("Encountered %d errors:", len(errors))
		for _, err := range errors {
			t.Logf("  %v", err)
		}
	}

	// Verify that findings have proper structure
	for _, finding := range findings {
		if finding.FilePath == "" {
			t.Error("Finding has empty FilePath")
		}
		if finding.Line == 0 {
			t.Error("Finding has line number 0")
		}
		if finding.Comment == "" {
			t.Error("Finding has empty Comment")
		}
	}
}

func TestIntegrationClean(t *testing.T) {
	// Test using the testdata/clean directory (files without BARK comments)
	cleanPath := filepath.Join(".", "testdata", "clean")

	// Check if testdata/clean exists
	if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
		t.Skip("testdata/clean directory not found")
	}

	s := scanner.NewScanner()
	result := s.Scan(cleanPath)

	findings := result.GetFindings()
	errors := result.GetErrors()

	// Should have NO BARK comments in clean directory
	if len(findings) > 0 {
		t.Errorf("Expected NO BARK comments in testdata/clean, but found %d", len(findings))
		for _, finding := range findings {
			t.Logf("  Unexpected finding: %s", finding.String())
		}
	}

	// Should have no errors
	if len(errors) > 0 {
		t.Errorf("Encountered %d unexpected errors:", len(errors))
		for _, err := range errors {
			t.Logf("  %v", err)
		}
	}

	t.Logf("✅ Clean directory scan passed - no BARK comments found")
}

func TestIntegrationFull(t *testing.T) {
	// Test scanning the entire testdata directory
	testdataPath := filepath.Join(".", "testdata")

	// Check if testdata exists
	if _, err := os.Stat(testdataPath); os.IsNotExist(err) {
		t.Skip("testdata directory not found")
	}

	s := scanner.NewScanner()
	result := s.Scan(testdataPath)

	findings := result.GetFindings()

	// Should find BARK comments (from dirty directory)
	if len(findings) == 0 {
		t.Error("Expected to find BARK comments in testdata (includes dirty/)")
	}

	t.Logf("Found %d total BARK comments in testdata (dirty + clean)", len(findings))
}

func TestIntegrationMultiplePaths(t *testing.T) {
	dirtyPath := filepath.Join(".", "testdata", "dirty")
	cleanPath := filepath.Join(".", "testdata", "clean")

	// Check if both directories exist
	if _, err := os.Stat(dirtyPath); os.IsNotExist(err) {
		t.Skip("testdata/dirty directory not found")
	}
	if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
		t.Skip("testdata/clean directory not found")
	}

	s := scanner.NewScanner()

	// ScanPaths with both dirty and clean should produce the same findings as
	// scanning the parent testdata directory, since clean has no BARK comments.
	resultMulti := s.ScanPaths([]string{dirtyPath, cleanPath})
	resultSingle := s.Scan(dirtyPath)

	multiFindings := resultMulti.GetFindings()
	singleFindings := resultSingle.GetFindings()

	if len(multiFindings) == 0 {
		t.Error("Expected to find BARK comments when scanning dirty + clean paths")
	}

	if len(multiFindings) != len(singleFindings) {
		t.Errorf(
			"ScanPaths([dirty, clean]) should find same count as Scan(dirty): multi=%d, single=%d",
			len(multiFindings),
			len(singleFindings),
		)
	}

	t.Logf("Found %d BARK comments via multi-path scan", len(multiFindings))
}
