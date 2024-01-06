package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/bazdalaz/archiver/lib/compression"
	"github.com/bazdalaz/archiver/lib/compression/vlc"
	"github.com/bazdalaz/archiver/lib/compression/vlc/table/shannon_fano"
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}


const packedExtension = ".vlc"

var ErrEmptyPath = errors.New("no file specified")

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder

	if len(args) == 0 {
		handleErr(ErrEmptyPath)
	}

	method  := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New(shannon_fano.Generator{})

	default:
		cmd.PrintErr("Unknown method")
	}

	filePath := args[0] //BUG: check if file exists

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := encoder.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handleErr(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + packedExtension
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("method", "m", "", "compression method: vlc")
	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
