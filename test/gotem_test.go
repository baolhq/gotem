package main

import (
	"bytes"
	"testing"

	"baolhq/gotem/cmd"
)

func TestSingleAddCmd(t *testing.T) {
	cmd := cmd.AddCmd()

	// Use a buffer to capture the output of the command
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	// Test with single file argument
	cmd.SetArgs([]string{"~/.bashrc"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the output is correct
	expected := "Adding file ~/.bashrc...\n"
	if buf.String() != expected {
		t.Errorf("Expected %v, got %v", expected, buf.String())
	}
}

func TestMultipleAddCmd(t *testing.T) {
	cmd := cmd.AddCmd()

	// Use a buffer to capture the output of the command
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	// Test with multiple file arguments
	cmd.SetArgs([]string{"~/.bashrc", "~/.vimrc"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the output is correct for multiple files
	expected := "Adding file ~/.bashrc...\nAdding file ~/.vimrc...\n"
	if buf.String() != expected {
		t.Errorf("Expected %v, got %v", expected, buf.String())
	}
}

func TestSingleRemoveCmd(t *testing.T) {
	cmd := cmd.RemoveCmd()

	// Use a buffer to capture the output of the command
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	// Test with no files, should print a default message
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test with one file argument
	buf.Reset()
	cmd.SetArgs([]string{"~/.bashrc"})
	err = cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check for the correct removal output
  expected := "Removing files ~/.bashrc...\n"
	if buf.String() != expected {
		t.Errorf("Expected %v, got %v", expected, buf.String())
	}
}

func TestMultipleRemoveCmd(t *testing.T) {
	cmd := cmd.RemoveCmd()

	// Use a buffer to capture the output of the command
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	// Test with no files, should print a default message
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check for the "Removing all files?" message
	expected := "Removing all files?\n"
	if buf.String() != expected {
		t.Errorf("Expected %v, got %v", expected, buf.String())
	}
}

// TestLinkCmd tests the "link" command.
func TestLinkCmd(t *testing.T) {
	cmd := cmd.LinkCmd()

	// Use a buffer to capture the output of the command
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	// Test with no arguments
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the output for linking all files
	expected := "Linking all files...\n"
	if buf.String() != expected {
		t.Errorf("Expected %v, got %v", expected, buf.String())
	}

	// Test with a file argument
	buf.Reset()
	cmd.SetArgs([]string{"~/.bashrc"})
	err = cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check for correct linking output
	expected = "Linking ~/.bashrc...\n"
	if buf.String() != expected {
		t.Errorf("Expected %v, got %v", expected, buf.String())
	}
}

// TestStatusCmd tests the "status" command.
func TestStatusCmd(t *testing.T) {
	cmd := cmd.StatusCmd()

	// Use a buffer to capture the output of the command
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	// Test with no files
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the "Checking status..." output
	expected := "Checking status...\n"
	if buf.String() != expected {
		t.Errorf("Expected %v, got %v", expected, buf.String())
	}

	// Test with a file argument
	buf.Reset()
	cmd.SetArgs([]string{"~/.bashrc"})
	err = cmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check for correct status output
	expected = "Checking status of ~/.bashrc...\n"
	if buf.String() != expected {
		t.Errorf("Expected %v, got %v", expected, buf.String())
	}
}
