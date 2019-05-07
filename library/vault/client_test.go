package vault

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNewVault(t *testing.T) {
  c, e := NewVault()
  assert.NotNil(t, c)
  assert.Nil(t, e)
  assert.NotEmpty(t, c.Token())
}

func TestReadWrite(t *testing.T) {
  c, _ := NewVault()
  secretData := map[string]interface{}{
    "foo": "bar",
    "age": "-1",
  }

  _, err := c.Logical().Write("kv1/my-secret", secretData)
  assert.Nil(t, err)

  s, err := c.Logical().Read("kv1/my-secret")
  assert.Nil(t, err)
  if assert.NotNil(t, s) {
    assert.Equal(t, "bar", s.Data["foo"], )
  }
}
