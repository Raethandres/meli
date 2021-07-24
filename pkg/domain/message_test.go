package domain

import (
    "testing"
    "github.com/matryer/is"
)

func TestGetMessage(t *testing.T) {

    t.Run("When get encoded message with symmetrical messages", func(t *testing.T) {
        is := is.New(t)
        kenobi := []string{"este", "", "", "mensaje", ""}
        skywalker := []string{"", "es", "", "", "secreto"}
        sato := []string{"este", "", "un", "", ""}

        decodedMessage := GetMessage(kenobi, skywalker, sato)
        
        is.Equal("este es un mensaje secreto", decodedMessage)
    })

    t.Run("When get encoded message with asymmetrical messages", func(t *testing.T) {
        is := is.New(t)
        kenobi := []string{"", "este", "es", "un", "mensaje"}
        skywalker := []string{"este", "", "un", "mensaje"}
        sato := []string{"", "", "es", "", "mensaje"}

        decodedMessage := GetMessage(kenobi, skywalker, sato)
        
        is.Equal("este es un mensaje", decodedMessage)
    })
}
