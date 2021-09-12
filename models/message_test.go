package models_test

import (
	"encoding/json"
	"testing"

	"github.com/cjkao/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
)

func Test_Marshell(t *testing.T) {
	m := models.Message{UnRead: true}
	str, _ := json.MarshalIndent(m, "\t", "\t")
	println(string(str))
	assert.Contains(t, string(str), "true")

	m = models.Message{UnRead: false}
	str, _ = json.MarshalIndent(m, "\t", "\t")
	println(string(str))
	assert.Contains(t, string(str), "false")
	// assert.Nil(t, str, "Couldn't create realtime client")
	// assert.Nil(t, str, "Couldn't create realtime client")

}
