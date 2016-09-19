package pingpong

import (
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestPingPong_para_o_numero_3_o_retorno_deve_ser_ping(t *testing.T){
    pingPingo := ping(3)
    assert.Equal(t, pingPingo, "Ping")
}

func TestPingPong_para_o_numero_6_o_retorno_deve_ser_ping(t *testing.T){
    pingPingo := ping(6)
    assert.Equal(t, pingPingo, "Ping")
}

func TestPingPong_para_o_numero_5_o_retorno_deve_ser_pong(t *testing.T){
    pingPingo := ping(5)
    assert.Equal(t, pingPingo, "Pong")
}

func TestPingPong_para_o_numero_25_o_retorno_deve_ser_pong(t *testing.T){
    pingPingo := ping(25)
    assert.Equal(t, pingPingo, "Pong")
}

func TestPingPong_para_o_numero_15_o_retorno_deve_ser_pingpong(t *testing.T){
    pingPingo := ping(15)
    assert.Equal(t, pingPingo, "PingPong")
}

func TestPingPong_para_o_numero_45_o_retorno_deve_ser_pingpong(t *testing.T){
    pingPingo := ping(45)
    assert.Equal(t, pingPingo, "PingPong")
}

func TestPingPong_para_o_numero_7_o_retorno_deve_ser_7(t *testing.T){
    pingPingo := ping(7)
    assert.Equal(t, pingPingo, "7")
}
