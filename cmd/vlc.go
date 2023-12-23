package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/bazdalaz/archiver/lib/vlc"
)

const packedExtension = ".vlc"
var ErrEmptyPath = errors.New("no file specified")

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "using variable-length encoding",
	Run:   pack,
}

func pack(_ *cobra.Command, args []string) {
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

	packed := vlc.Encode(string(data))


	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handelError(err)
	} 
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + packedExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
