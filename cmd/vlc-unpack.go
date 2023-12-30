package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/bazdalaz/archiver/lib/vlc"
	"github.com/spf13/cobra"
)

// TODO: take extension from file
const unpackedExtension = ".txt"

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable-length encoding",
	Run:   unpack,
}

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		handelError(ErrEmptyPath)
	}

	filePath := args[0] //BUG: check if file exists

	r, err := os.Open(filePath)
	if err != nil {
		handelError(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handelError(err)
	}

	packed := vlc.Decode(string(data))

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handelError(err)
	}
}

// TODO: refactor this function
func unpackedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + unpackedExtension
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}
