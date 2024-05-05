package utils

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func SendSMS(to string, message string) error {
	link := "https://api.africastalking.com/version1/messaging"
	method := "POST"

	data := url.Values{}
	data.Set("username", "enyumba")
	data.Set("to", to)
	data.Set("message", message)
	body := strings.NewReader(data.Encode())
	client := &http.Client{}

	request, error := http.NewRequest(method, link, body)
	if error != nil {
		return error
	}
	live := "637e51e6a99b335eafa7fd6ad08287a410a95460aeeafd9a1eab27c735a14975"
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("apiKey", live)

	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	return nil
}

func GenerateRandomNumber() int64 {
	// Generate 6 bytes of random data
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	// Truncate the random data to 6 digits
	num := int64(b[0])*100000 + int64(b[1])*10000 + int64(b[2])*1000 + int64(b[3])*100 + int64(b[4])*10 + int64(b[5])
	num = num % 900000 // Ensure num is between 0 and 899999
	num += 100000      // Add 100000 to get a number between 100000 and 999999

	return num
}

func GenerateAccountNumber() int64 {
	// Generate 7 bytes of random data
	b := make([]byte, 7)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	// Truncate the random data to 12 digits
	num := int64(b[0])*100000000000 + int64(b[1])*10000000000 + int64(b[2])*1000000000 + int64(b[3])*100000000 + int64(b[4])*10000000 + int64(b[5])*1000000 + int64(b[6])*100000
	num = num % 1000000000000

	return num
}
