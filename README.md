# zero-security-conf [![GoDoc][doc-img]][doc]

A simple method for config encrypt/decrypt by AES-256 for go-zero.

## Usage

Add `confz.SecurityConf` to your `Config`.
```go
import (
	"github.com/zeromicro/go-zero/rest"
	confz "github.com/WqyJh/zero-security-conf"
)
type Config struct {
	rest.RestConf
 	Security confz.SecurityConf
 	SensitiveKey string
 	SensitiveValue string
}
```

Use the following code to encrypt your sensitive data, and replace the plain string in your config with the encrypted string.
```go
import (
	"github.com/WqyJh/confcrypt"
)

var (
	key = "12345678"
)

func TestEncrypt(t *testing.T) {
	plain := "sensitive_key"
	encrypted, err := confcrypt.EncryptString(plain, key)
	assert.NoError(t, err)
	t.Logf("encrypted: '%s'", encrypted) // encrypted: ENC~i1eiPez4IICS/iA+zIEyDk3UHQz9enP+kHG3X/LCJixtgEw4i3o=
}
```

This is the config file.
```yaml
Security:
  Enable: true
  Env: MY_SECRET_KEY

SensitiveKey: ENC~i1eiPez4IICS/iA+zIEyDk3UHQz9enP+kHG3X/LCJixtgEw4i3o=
SensitiveValue: ENC~KWLH5csxSeG3zgPMFYmgIslTrPaWUfZsLaAkJ9z9zwf6LXyHh5ddYeO5sCRH8xeLOXGWUaA=
```

Use `confz.SecurityLoad` or `confz.SecurityMustLoad` instead of `conf.Load` or `conf.MustLoad`.

Start the service with environment variable of you secret key.

```bash
export MY_SECRET_KEY=mysecretkey
```

All of the string config starts with `ENC~` would be decrypted.


## License

Released under the [MIT License](LICENSE).

[doc-img]: https://godoc.org/github.com/WqyJh/zero-security-conf?status.svg
[doc]: https://godoc.org/github.com/WqyJh/zero-security-conf
