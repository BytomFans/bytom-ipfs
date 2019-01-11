package bytom

//钱包打开的情况下
type Xpub struct {
	Xpub string `json:"xpub"`
}
type CreateKey struct {
	Alias    string `json:"alias"`
	Password string `json:"password"`
}

type DeleteKey struct {
	Xpub     Xpub.Xpub ` "xpub"`
	Password string    `json:"password"`
}

//重置密钥密码
type ResetKeyPassword struct {
	Xpub        Xpub.Xpub `json:"xpub"`
	OldPassword string    `json:"old_password"`
	NewPassword string    `json:"new_password"`
}

//创建账户
type CreateAccount struct {
	RootXpubs string `json:"root_xpubs"`
	Quorum    int64  `json:"quorum"`
	Alias     string `json:"alias"`
}

type UpdateAccountAlias struct {
	AccountId    string `json:"account_id"`
	AccountAlias string `json:"account_alias"`
	NewAlias     string `json:"new_alias"`
}

//删除现有帐户，请确保相关地址中没有余额
type DeleteAccount struct {
	AccountInfo string `json:"account_id"`
}

type CreateAccountReceiver struct {
	AccountId    string `json:"account_id"`
	AccountAlias string `json:"account_alias"`
}

//验证地址是否有效，并判断地址是否为自己
type ValidateAddress struct {
	Address string `json:"address"`
}

//创建资产定义，准备发行资产
type CreateAsset struct {
	Alias     string `json:"alias"`
	RootXpubs []Xpub `json:"root_xpubs"`
	Quorum    string `json:"quorum"`
}

type GetAsset struct {
	Id string `json:"id"`
}

type UpdateAssetAlias struct {
	Id    string `json:"id"`
	Alias string `json:"alias"`
}

type SignMessage struct {
	Address string `json:"address"`
	Message string `json:"message"`
}

type GetTransaction struct {
	TxId string `json:"tx_id"`
}

type BuildTransaction struct {
}

type SignTransaction struct {
}

//钱包没打开的情况下

type SubmitTransaction struct {
}

type EstimateTransactionGas struct {
}

type CheckAccessToken struct {
}

type GetBlockCount struct {
}

type GetBlockHash struct {
}

type GetBlock struct {
}

type GetBlockHeader struct {
}

type GetDifficulty struct {
}

type GetHashRate struct {
}

type NetInfo struct {
}

type IsMining struct {
}

type SetMining struct {
}
type GasRate struct {
}
type VerifyMessage struct {
}
type DecodeProgram struct {
}
type GetWork struct {
}
type SubmitWork struct {
}
