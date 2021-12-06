package main

import (
	"fmt"
	"os"
	"xvaultctl/utils"
)

type Options struct {
	Provider    string `example:"Provider (aws | gcp | hashcorp)"`
	Environment string `example:"Environment"`
	User        string `example:"User"`
	ServiceUser string `example:"ServiceUser"`
	Password    string `example:"Password"`
	ClientToken string `example:"ClientToken"`
	Tenant      string `example:"Tenant"`
	Credentials bool   `example:"false|true"`
	Vault       struct {
		Name  string `example:"Name"`
		Id    string `example:"Id"`
		Token string `example:"Id"`
	} `example:"Vault"`
}

func main() {

	fmt.Println("xVaultCTL Commands (CLI)")

	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "help" {
		fmt.Println(utils.Help())
	}

	options := new(Options)

	for index, opt := range os.Args {

		switch opt {

		case "provider", "-P":
			if len(os.Args[index:]) <= 1 {
				fmt.Println("You need to set Provider option")
				os.Exit(1)
			} else {
				options.Provider = os.Args[index+1]
			}
		case "environment", "-e":
			if len(os.Args[index:]) <= 1 {
				fmt.Println("You need to set Environment option")
				os.Exit(1)
			} else {
				options.Environment = os.Args[index+1]
			}
		case "user", "-u":
			if len(os.Args[index:]) <= 1 {
				fmt.Println("You need to set User option")
				os.Exit(1)
			} else {
				options.User = os.Args[index+1]
			}
		case "serviceuser", "-U":
			if len(os.Args[index:]) <= 1 {
				fmt.Println("You need to set Service User option")
				os.Exit(1)
			} else {
				options.ServiceUser = os.Args[index+1]
			}
		case "password", "-p":
			if len(os.Args[index:]) <= 1 {
				fmt.Println("You need to set User option")
				os.Exit(1)
			} else {
				options.Password = os.Args[index+1]
			}
		case "cliente_token":
			if len(os.Args[index:]) <= 1 {
				fmt.Println("You need to set the ClientToken")
				os.Exit(1)
			} else {
				options.ClientToken = os.Args[index+1]
			}
		case "vaultid":
			if len(os.Args[index:]) <= 1 {
				fmt.Println("You need to set the VaultID")
				os.Exit(1)
			} else {
				options.Vault.Id = os.Args[index+1]
			}
		case "credentials":
			if len(os.Args[index:]) <= 1 {
				fmt.Println("You need to set the Vault Name")
				os.Exit(1)
			} else {
				options.Credentials = true
				options.Vault.Name = os.Args[index+1]
			}
		}

	}

	for index, opt := range os.Args {

		switch opt {
		case "login":
			if len(os.Args[index:]) <= 1 || os.Args[index+1] == "--help" || os.Args[index+1] == "help" {
				fmt.Println(utils.Login())
			} else {
				options.ObjBuildMsg("Login on vault Provider: \n")
				options.Login()
			}
		case "update":
			if len(os.Args[index:]) <= 1 || os.Args[index+1] == "--help" || os.Args[index+1] == "help" {
				fmt.Println(utils.Update())
			} else {
				options.ObjBuildMsg("Update vaults with configurations below: \n")
				options.Update()
			}
		case "get":
			if len(os.Args[index:]) <= 1 || os.Args[index+1] == "--help" || os.Args[index+1] == "help" {
				fmt.Println(utils.Get())
			} else {
				options.ObjBuildMsg("Getting vault with configurations below: \n")
				options.Get()
			}
		case "create":
			if len(os.Args[index:]) <= 1 || os.Args[index+1] == "--help" || os.Args[index+1] == "help" {
				fmt.Println(utils.Create())
			} else {
				options.ObjBuildMsg("Creating vaults with configurations below: \n")
				options.Create()
			}
		case "grant":
			if len(os.Args[index:]) <= 1 || os.Args[index+1] == "--help" || os.Args[index+1] == "help" {
				fmt.Println(utils.Grant())
			} else {
				options.ObjBuildMsg("Deleting vault with configurations below: \n")
				options.Grant()
			}
		}
	}
}

func (options Options) ObjBuildMsg(msg string) {

	if options.Provider != "" {
		msg += "Provider: " + options.Provider + "\n"
	}
	if options.Environment != "" {
		msg += "Enviroment: " + options.Environment + "\n"
	}

	fmt.Println(msg)
}

func (options Options) Login() {
	fmt.Println("Login..")

	if options.Provider == "aws" {
		awsvault.Get("8212e123")
	} else if options.Provider == "gcp" {

		if options.Password == "" {
			options.Password = security.GenerateCredentialsPassword()
		}

		if options.ServiceUser != "" && options.Vault.Id != "" {
			options.Password = gcp.LoginServiceUser(options.ServiceUser, options.Password, options.Vault.Id)
			options.User = options.ServiceUser
		}

		options.Vault.Token = gcp.LoginVault(options.User, options.Password, options.Environment)
		fmt.Println("Vault Token:", options.Vault.Token)

	} else {
		fmt.Println(utils.Login())
	}
}

func (options Options) Update() {
	fmt.Println("Updating Vault")

}

func (options Options) Get() {
	fmt.Println("Getting Vault")
	if options.Credentials {

		fmt.Println("User:", os.Getenv("User"))
		gcp.GetRoleSecretID(options.User)
	}

}

func (options Options) Create() {
	fmt.Println("Creating")

}

func (options Options) Grant() {
	fmt.Println("Grant Vault")
}
