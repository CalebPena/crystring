package cmd

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crystring",
	Short: "Generate a password",
	Long: `
Generate a 16 digit password and copies it to the clipboard.

Usage:
crystring <lenght=16>
`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		length := 16

		if len(args) == 1 {
			argLen, err := strconv.Atoi(args[0])

			if err != nil {
				return err
			}

			length = argLen
		}

		password, err := genPass(length)

		if err != nil {
			return err
		}

		fmt.Println("copied to clipboard")
		return clipboard.WriteAll(password)
	},
}

const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

func genPass(length int) (string, error) {
	password := ""

	charsLen := big.NewInt((int64(len(chars))))

	for i := 0; i < length; i++ {
		randomNum, err := rand.Int(rand.Reader, charsLen)

		if err != nil {
			return "", errors.New("failed to generate random number")
		}

		randomInt := int(randomNum.Int64())

		password += string(chars[randomInt])
	}

	return password, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
