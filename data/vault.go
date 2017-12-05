package data

import (
	//"fmt"
	vault "github.com/hashicorp/vault/api"
)

type Vault struct {
	Client *vault.Logical
}

func NewVault() (*Vault, error) {
	/*
		TODO:
		- create vault config from params
		- return vault
	*/

	var vault *Vault 
	return vault, nil
}

func (v *Vault) Create(path string, data map[string]interface{}) error {
	_, err := v.Client.Write(path, data)
	if err != nil {
		return err
	}
	return nil
}

func (v *Vault) Get(path string) (map[string]interface{}, error) {
	var data map[string]interface{}
	s, err := v.Client.Read(path)
	if err != nil {
		return data, err
	}
	return s.Data, nil
}

func (v *Vault) List(path string) (map[string]interface{}, error) {
	var data map[string]interface{}
	l, err := v.Client.List(path)
	if err != nil {
		return data, err
	}
	return l.Data, nil
} 

func (v *Vault) Update(path string, data map[string]interface{}) error {
	_, err := v.Client.Write(path, data)
	if err != nil {
		return err
	}
	return nil
}

func (v *Vault) Delete(path string) error {
	_, err := v.Client.Delete(path)
	if err != nil {
		return err
	}
	return nil
}