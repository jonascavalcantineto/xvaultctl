package help

func Help() string {
	headerMsg := "xvaultctl controls the vaults on AWS Cloud and xvaultctl\n" +
		"\n" +
		"Find more information at:  \n" +
		"\n" +
		"Basic Commands:\n" +
		"\tlogin	Login on providers to work with vaults\n" +
		"\tcreate	Create vault \n" +
		"\tupdate	Update vault \n" +
		"\tgrant	Grant vault \n"

	return headerMsg
}

func Login() string {
	listMsg := "Login Command\n" +
		"\n" +
		"Examples:\n" +
		"\t# Login on vault provider using nominal user\n" +
		"\xvaultctl login -u <myuser> provider xvaultctl environment qa\n" +
		"\n" +
		"\t# Login on vault provider using nominal user and password\n" +
		"\xvaultctl login -u <myuser> -p '<mypass>' provider xvaultctl environment <dev | qa | prod>\n" +
		"\n" +
		"\t# Login on vault provider using service user\n" +
		"\xvaultctl login -U <service user> vaultid <vault name id> provider xvaultctl environment qa\n" +
		"\n" +
		"\t# Login on vault provider using service user and passoword\n" +
		"\xvaultctl login -U <service user> -p '<service password>' vaultid <vault name id> provider xvaultctl environment qa\n" +
		"\n" +
		"xvaultctl Options:\n" +
		"Using Nominal User:\n" +
		"\t-u,	user		User (aws | xvaultctl)\n" +
		"\t-p,	password	User password \n" +
		"\t-P,	provider	Vault provider (aws | xvaultctl)\n" +
		"\t-e,	enviroment	Environment where will get the vault (qa | prod)\n" +
		"\n" +
		"Using Service User:\n" +
		"\t-U,	serviceUser Service User \n" +
		"\t-p,	password	User password \n" +
		"\t	vaultid		Vault Name Id \n" +
		"\t-P,	provider	Vault provider (aws | xvaultctl)\n" +
		"\t-e,	enviroment	Environment where will get the vault (qa | prod)\n" +
		"\n" +
		"AWS Options:\n" +
		"\tWill be implemented"

	return listMsg
}

func Update() string {
	listMsg := "Update Command\n" +
		"\n" +
		"Examples:\n" +
		"\t# Update vaults\n" +
		"\xvaultctl update provider xvaultctl environment qa\n" +
		"\n" +
		"Options Required:\n" +
		"\tprovider     Vault provider (aws | xvaultctl)\n" +
		"\tenviroment   Environment where will get the vault (qa | prod)\n"

	return listMsg
}

func Get() string {
	getMsg := "Get Command\n" +
		"\n" +
		"Examples:\n" +
		"\t# Get vault with key id\n" +
		"\xvaultctl get provider xvaultctl environment qa\n" +
		"\n" +
		"Options Required:\n" +
		"\tprovider     Vault provider (aws | xvaultctl)\n" +
		"\tenviroment   Environment where will get the vault (qa | prod)\n"

	return getMsg
}

func Create() string {
	createMsg := "Create Command\n" +
		"\n" +
		"Examples:\n" +
		"\t# Create vault\n" +
		"\xvaultctl create name myvaultname provider xvaultctl environment qa\n" +
		"\n" +
		"Options Required:\n" +
		"\tname         Vault name\n" +
		"\tprovider     Vault provider (aws | xvaultctl)\n" +
		"\tenviroment   Environment where will get the vault (qa | prod)\n"

	return createMsg
}

func Grant() string {
	deleteMsg := "Grant Command\n" +
		"\n" +
		"Examples:\n" +
		"\n" +
		"\t# Grant vault\n" +
		"\xvaultctl grant provider xvaultctl environment qa vault 2123123\n" +
		"\n" +
		"Options Required:\n" +
		"\tprovider     Vault provider (aws | xvaultctl)\n" +
		"\tenviroment   Environment where will get the vault (qa | prod)\n" +
		"\tid       	Vault id\n"

	return deleteMsg
}
