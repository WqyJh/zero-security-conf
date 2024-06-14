package confz_test

import (
	"os"
	"testing"

	"github.com/WqyJh/confcrypt"
	confz "github.com/WqyJh/zero-security-conf"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/hash"
)

func TestSecurityLoad(t *testing.T) {
	key := "testkey"
	type testConfig struct {
		confz.SecurityConf

		User   string `json:"user"`
		Pass   string `json:"pass"`
		Secret string `json:"secret"`
	}
	expected := testConfig{
		SecurityConf: confz.SecurityConf{
			Enable: true,
			Env:    "CONFIG_KEY",
		},
		User:   "testuser",
		Pass:   "testpass",
		Secret: "testsecret",
	}
	encryptedPass, err := confcrypt.EncryptString(expected.Pass, key)
	assert.Nil(t, err)
	encryptedSecret, err := confcrypt.EncryptString(expected.Secret, key)
	assert.Nil(t, err)
	text := `{
		"user": "testuser",
		"pass": "` + encryptedPass + `",
		"secret": "` + encryptedSecret + `"
}`
	tmpfile, err := createTempFile(".json", text)
	assert.Nil(t, err)
	defer os.Remove(tmpfile)

	os.Setenv("CONFIG_KEY", key)
	var config testConfig
	err = confz.SecurityLoad(tmpfile, &config)
	assert.Nil(t, err)
	assert.NotEqual(t, encryptedPass, config.Pass)
	assert.NotEqual(t, encryptedPass, config.Secret)
	assert.Equal(t, expected.Pass, config.Pass)
	assert.Equal(t, expected.Secret, config.Secret)
}

func TestSecurityLoadRecursive(t *testing.T) {
	key := "testkey"
	type NestedConfig struct {
		Security confz.SecurityConf
	}
	type testConfig struct {
		NestedConfig

		User   string `json:"user"`
		Pass   string `json:"pass"`
		Secret string `json:"secret"`
	}
	expected := testConfig{
		NestedConfig: NestedConfig{
			Security: confz.SecurityConf{
				Enable: true,
				Env:    "CONFIG_KEY",
			},
		},
		User:   "testuser",
		Pass:   "testpass",
		Secret: "testsecret",
	}
	encryptedPass, err := confcrypt.EncryptString(expected.Pass, key)
	assert.Nil(t, err)
	encryptedSecret, err := confcrypt.EncryptString(expected.Secret, key)
	assert.Nil(t, err)
	text := `{
		"user": "testuser",
		"pass": "` + encryptedPass + `",
		"secret": "` + encryptedSecret + `"
}`
	tmpfile, err := createTempFile(".json", text)
	assert.Nil(t, err)
	defer os.Remove(tmpfile)

	os.Setenv("CONFIG_KEY", key)
	var config testConfig
	err = confz.SecurityLoad(tmpfile, &config)
	assert.Nil(t, err)
	assert.NotEqual(t, encryptedPass, config.Pass)
	assert.NotEqual(t, encryptedPass, config.Secret)
	assert.Equal(t, expected.Pass, config.Pass)
	assert.Equal(t, expected.Secret, config.Secret)
}

func createTempFile(ext, text string) (string, error) {
	tmpFile, err := os.CreateTemp(os.TempDir(), hash.Md5Hex([]byte(text))+"*"+ext)
	if err != nil {
		return "", err
	}

	if err := os.WriteFile(tmpFile.Name(), []byte(text), os.ModeTemporary); err != nil {
		return "", err
	}

	filename := tmpFile.Name()
	if err = tmpFile.Close(); err != nil {
		return "", err
	}

	return filename, nil
}
