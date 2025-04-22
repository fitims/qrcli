package cmd

import (
	"fmt"
	"github.com/caiguanhao/readqr"
	"github.com/spf13/cobra"
	"os"
)

var qrDecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode text from a QR code PNG image.",
	Long:  `Decode text from an existing QR code PNG image.`,
	Run: func(cmd *cobra.Command, args []string) {
		qrFile, err := cmd.Flags().GetString("input")
		if err != nil {
			fmt.Printf("The flag has invalid contents: %s", err.Error())
			os.Exit(1)
			return
		}

		if len(qrFile) == 0 {
			fmt.Println("Flag input is mandatory!")
			cmd.Usage()
			os.Exit(1)
			return
		}

		f, err := os.Open(qrFile)
		if err != nil {
			fmt.Printf("Cannot read QR code file: %s", err.Error())
			os.Exit(1)
			return
		}
		defer f.Close()

		qr, err := readqr.Decode(f)
		if err != nil {
			fmt.Printf("Cannot decode QR code: %s", err.Error())
			os.Exit(1)
			return
		}

		fmt.Printf("%s\n", qr)
	},
}

func init() {
	rootCmd.AddCommand(qrDecodeCmd)
	qrDecodeCmd.Flags().StringP("input", "i", "", "The path to the PNG image containing the QR code")
}
