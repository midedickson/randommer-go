package randommer

const (
	// cards
	CARD_TYPES   = "Card/Types"
	CARD         = "Card"
	CARD_BY_TYPE = "Card?type=%s"
	// name
	NAMES = "Name?nameType=%s&quantity=%d"
	// finance
	CRYPTOADRESS_TYPES     = "Finance/CryptoAddress/Types"
	CRYPTO_ADDRESS_BY_TYPE = "Finance/CryptoAddress?cryptoType=%s"
	AVAILABLE_COUNTRIES    = "Finance/Countries"
	BANK_BY_COUNTRIES      = "Finance/Iban/%s"
	VAT_VALIDATOR          = "Finance/Vat/Validator?country=%s&vat=%s"

	// misc
	CULTURES       = "Misc/Cultures"
	RANDOM_ADDRESS = "Misc/Random-Address?number=%d&culture=%s"
)
