package main

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/go-bip39"
	"github.com/spf13/cobra"
)

func keysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "keys",
		Aliases: []string{"k"},
		Short:   "Manage key held by the oracle process",
	}
	cmd.AddCommand(keysAddCmd())
	cmd.AddCommand(keysListCmd())
	return cmd
}

func keysAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add [name]",
		Aliases: []string{"a"},
		Short:   "Add a new key to the keychain",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: Allow mnemonic import
			seed, err := bip39.NewEntropy(256)
			if err != nil {
				return err
			}

			mnemonic, err := bip39.NewMnemonic(seed)
			if err != nil {
				return err
			}

			info, err := keybase.NewAccount(
				args[0], mnemonic, "", hd.CreateHDPath(494, 0, 0).String(), hd.Secp256k1,
			)
			if err != nil {
				return err
			}

			logger.Info("📝 Mnemonic: %s", mnemonic)
			logger.Info("📮 Address: %s", info.GetAddress().String())
			return nil
		},
	}
	return cmd
}

func keysListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "List all the keys in the keychain",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			keys, err := keybase.List()
			if err != nil {
				return err
			}

			for _, key := range keys {
				logger.Info("👨‍⚖️ %s => %s", key.GetName(), key.GetAddress().String())
			}
			return nil
		},
	}
	return cmd
}
