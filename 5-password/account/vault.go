package account

import (
	"demo/password/encrypter"
	"demo/password/file"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db  file.JsonDb
	enc encrypter.Encrypter
}

func GetVault(db *file.JsonDb, enc *encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read()

	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  *db,
			enc: *enc,
		}
	}

	existingVaultJson := enc.Decrypt(file)

	var vault Vault
	err = json.Unmarshal(existingVaultJson, &vault)

	fmt.Printf("Найдено %d аккаунтов", len(vault.Accounts))

	if err != nil {
		fmt.Print("Не удалось разобрать data.json")

		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  *db,
			enc: *enc,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    *db,
		enc:   *enc,
	}
}

func (vault *VaultWithDb) AddAccount(acc *Account) {
	vault.Accounts = append(vault.Accounts, *acc)

	err := vault.writeVault()

	if err != nil {
		fmt.Print("Не удалось преобразовать")
	}
}

func (vault *VaultWithDb) writeVault() error {
	vault.UpdatedAt = time.Now()

	vaultJson, err := json.Marshal(vault.Vault)
	encData := vault.enc.Encrypt(vaultJson)

	if err != nil {
		return errors.New("VAULT_WRITE_ERROR")
	}

	vault.db.Write(encData)

	return nil
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account

	for _, acc := range vault.Accounts {
		if checker(acc, str) {
			accounts = append(accounts, acc)
		}
	}

	return accounts
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) {
	for index, acc := range vault.Accounts {
		if acc.Url == url {
			vault.Accounts = append(vault.Accounts[:index], vault.Accounts[index+1:]...)
			err := vault.writeVault()

			if err != nil {
				fmt.Print("Не удалось записать")
				return
			}

			fmt.Printf("Аккаунт %v успешно удален", url)
			return
		}
	}

	fmt.Printf("Аккаунт %v не найден", url)
}
