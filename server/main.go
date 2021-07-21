package main

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	cache "github.com/helmutkemper/cache.thread.safe/memory"
	"log"
	"os"
	"sync"
)

type AuthGoFakeIt struct {
	Username string
	Password string
}

func (e *AuthGoFakeIt) Populate() {
	e.Username = gofakeit.Username()
	e.Password = gofakeit.Password(true, true, true, true, true, 20)
}

type CompanyGoFakeIt struct {
	BS            string
	BuzzWord      string
	Company       string
	CompanySuffix string
	Job           *gofakeit.JobInfo
	JobDescriptor string
	JobLevel      string
	JobTitle      string
}

func (e *CompanyGoFakeIt) Populate() {
	e.BS = gofakeit.BS()
	e.BuzzWord = gofakeit.BuzzWord()
	e.Company = gofakeit.Company()
	e.CompanySuffix = gofakeit.CompanySuffix()
	e.Job = gofakeit.Job()
	e.JobDescriptor = gofakeit.JobDescriptor()
	e.JobLevel = gofakeit.JobLevel()
	e.JobTitle = gofakeit.JobTitle()
}

type PaymentGoFakeIt struct {
	CreditCard        *gofakeit.CreditCardInfo
	CreditCardCvv     string
	CreditCardExp     string
	CreditCardNumber  string
	CreditCardType    string
	Currency          *gofakeit.CurrencyInfo
	CurrencyLong      string
	CurrencyShort     string
	AchRouting        string
	AchAccount        string
	BitcoinAddress    string
	BitcoinPrivateKey string
}

func (e *PaymentGoFakeIt) Populate() {
	e.CreditCard = gofakeit.CreditCard()
	e.CreditCardCvv = gofakeit.CreditCardCvv()
	e.CreditCardExp = gofakeit.CreditCardExp()
	e.CreditCardNumber = gofakeit.CreditCardNumber(nil)
	e.CreditCardType = gofakeit.CreditCardType()
	e.Currency = gofakeit.Currency()
	e.CurrencyLong = gofakeit.CurrencyLong()
	e.CurrencyShort = gofakeit.CurrencyShort()
	e.AchRouting = gofakeit.AchRouting()
	e.AchAccount = gofakeit.AchAccount()
	e.BitcoinAddress = gofakeit.BitcoinAddress()
	e.BitcoinPrivateKey = gofakeit.BitcoinPrivateKey()
}

type AddressGoFakeIt struct {
	Address      *gofakeit.AddressInfo
	City         string
	Country      string
	CountryAbr   string
	State        string
	StateAbr     string
	Street       string
	StreetName   string
	StreetNumber string
	StreetPrefix string
	StreetSuffix string
	Zip          string
	Latitude     float64
	Longitude    float64
}

func (e *AddressGoFakeIt) Populate() {
	e.Address = gofakeit.Address()
	e.City = gofakeit.City()
	e.Country = gofakeit.Country()
	e.CountryAbr = gofakeit.CountryAbr()
	e.State = gofakeit.State()
	e.StateAbr = gofakeit.StateAbr()
	e.Street = gofakeit.Street()
	e.StreetName = gofakeit.StreetName()
	e.StreetNumber = gofakeit.StreetNumber()
	e.StreetPrefix = gofakeit.StreetPrefix()
	e.StreetSuffix = gofakeit.StreetSuffix()
	e.Zip = gofakeit.Zip()
	e.Latitude = gofakeit.Latitude()
	e.Longitude = gofakeit.Longitude()
}

type PersonGoFakeIt struct {
	Id             string
	Person         *gofakeit.PersonInfo
	Name           string
	NamePrefix     string
	NameSuffix     string
	FirstName      string
	LastName       string
	Gender         string
	SSN            string
	Contact        *gofakeit.ContactInfo
	Email          string
	Phone          string
	PhoneFormatted string

	Auth    AuthGoFakeIt
	Company CompanyGoFakeIt
	Payment PaymentGoFakeIt
	Address AddressGoFakeIt

	Bytes int
}

func (e *PersonGoFakeIt) Populate() (err error) {
	e.Id = gofakeit.UUID()
	e.Person = gofakeit.Person()
	e.Name = gofakeit.Name()
	e.NamePrefix = gofakeit.NamePrefix()
	e.NameSuffix = gofakeit.NameSuffix()
	e.FirstName = gofakeit.FirstName()
	e.LastName = gofakeit.LastName()
	e.Gender = gofakeit.Gender()
	e.SSN = gofakeit.SSN()
	e.Contact = gofakeit.Contact()
	e.Email = gofakeit.Email()
	e.Phone = gofakeit.Phone()
	e.PhoneFormatted = gofakeit.PhoneFormatted()

	e.Auth.Populate()
	e.Company.Populate()
	e.Payment.Populate()
	e.Address.Populate()

	_, err = json.Marshal(e)
	return
}

func (e *PersonGoFakeIt) MarshalJSON() (data []byte, err error) {
	type personGoFakeIt struct {
		Id             string
		Person         *gofakeit.PersonInfo
		Name           string
		NamePrefix     string
		NameSuffix     string
		FirstName      string
		LastName       string
		Gender         string
		SSN            string
		Contact        *gofakeit.ContactInfo
		Email          string
		Phone          string
		PhoneFormatted string

		Auth    AuthGoFakeIt
		Company CompanyGoFakeIt
		Payment PaymentGoFakeIt
		Address AddressGoFakeIt

		Bytes int
	}

	var toJson = personGoFakeIt(*e)
	data, err = json.Marshal(&toJson)
	if err != nil {
		log.Printf("1.PersonGoFakeIt().json.Marshal(&toJson).error: %v", err)
		return
	}

	toJson.Bytes = len(data)
	e.Bytes = toJson.Bytes

	data, err = json.Marshal(&toJson)
	if err != nil {
		log.Printf("2.PersonGoFakeIt().json.Marshal(&toJson).error: %v", err)
		return
	}

	return
}

func main() {
	var err error
	var wg sync.WaitGroup

	cacheServer := &cache.Memory{}

	gofakeit.Seed(0)

	log.Printf("populate start")
	var user = PersonGoFakeIt{}
	for i := 0; i != 1; i += 800 * 1000 {
		err = user.Populate()
		if err != nil {
			log.Printf("user.Populate().error: %v", err)
			return
		}
	}
	log.Printf("populate end")

	err = cacheServer.InitServert()
	if err != nil {
		log.Printf("error: %v", err.Error())
	}

	log.Printf("starting server main(): %v", os.Getenv("DEBUG_NAME"))

	wg.Add(1)
	wg.Wait()
}
