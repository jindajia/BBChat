package handlers

import (
	// 	"bytes"
	// 	"encoding/json"
	// 	"github.com/gorilla/mux"
	// 	"github.com/stretchr/testify/assert"
	// 	"log"
	"fmt"
	// 	"net/http"
	// 	"net/http/httptest"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
}

func TestStoreNewChatMessages(t *testing.T) {

	messagePayload_True := MessagePayloadStruct{
		FromUserID: "62476975f2b1f88baeffb354",
		Message:    "1 April Test for TestStoreNewChatMessages func",
		ToUserID:   "62476a7af2b1f88baeffb355",
	}
	if ans := StoreNewChatMessages(messagePayload_True); ans != true {
		fmt.Println(" expected be true, but got false")
	}

	messagePayload_False := MessagePayloadStruct{
		FromUserID: "62476975f2b1f88baeffb354",
		Message:    "1 April Test for TestStoreNewChatMessages func",
		ToUserID:   "65", // nil
	}
	if ans := StoreNewChatMessages(messagePayload_False); ans != true {
		fmt.Println(" expected be false since the to user ID is invalid value, and got false")
	}
}

func TestStoreNewDriftBottles(t *testing.T) {

	driftBottlePacket_True := DriftBottlePayloadStruct{
		FromUserID: "62476975f2b1f88baeffb354",
		Message:    "1 April Test for test func",
		ToUserID:   "62476a7af2b1f88baeffb355",
	}
	if ans := StoreNewDriftBottles(driftBottlePacket_True); ans != true {
		t.Errorf(" expected be true, but got false")
	}

	driftBottlePacket_False := DriftBottlePayloadStruct{
		FromUserID: "62476975f2b1f88baeffb354",
		Message:    "1 April Test for test func",
		ToUserID:   "656598565",
	}
	if StoreNewDriftBottles(driftBottlePacket_False) != true {
		fmt.Println(" expected be false since the to user ID is invalid value, and got false")
	}
}

func TestStoreNewChatImages(t *testing.T) {

	imagePayload_True := ImagePayloadStruct{
		FromUserID: "62476975f2b1f88baeffb354",
		Image:      "01ef011790175",
		ToUserID:   "62476a7af2b1f88baeffb355",
	}
	if StoreNewChatImages(imagePayload_True) != true {
		t.Errorf(" expected be true, but got false")
	}
	fmt.Println("imagePayload_True passed")

	imagePayload_False := ImagePayloadStruct{
		FromUserID: "62476975f2b1f88baeffb354",
		Image:      "1 April Test for test func",
		ToUserID:   "65", // nil
	}
	if StoreNewChatImages(imagePayload_False) != true {
		fmt.Println(" expected be false since the to user ID is invalid value, and got false")
	}
}

func TestStoreNewBroadcastMessages(t *testing.T) {

	messagePayload_True1 := MessagePayloadStruct{
		FromUserID: "62476975f2b1f88baeffb354",
		Image:      "01ef011790175",
	}
	if StoreNewBroadcastMessages(messagePayload_True1) != true {
		t.Errorf(" expected be true, but got false")
	}
	fmt.Println("imagePayload_True1 passed")

	messagePayload_True2 := MessagePayloadStruct{
		FromUserID: "62476975f2b1f88baeffb354",
		Message:    "This is a test message for unit test of StoreNewBroadcastMessages",
	}
	if StoreNewBroadcastMessages(messagePayload_True2) != true {
		t.Errorf(" expected be true, but got false")
	}
	fmt.Println("imagePayload_True2 passed")

	messagePayload_False := MessagePayloadStruct{
		FromUserID: "wrfe", // nil
		Image:      "1 April Test for test func",
	}
	if StoreNewBroadcastMessages(messagePayload_False) != true {
		fmt.Println(" expected be false since the from user ID is invalid value, and got false")
	}
}
