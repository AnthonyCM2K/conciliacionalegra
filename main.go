package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// // Función para deserializar JSON a un slice de FacturasCMS
// func UnmarshalFacturasCMS1(data []byte) ([]FacturasCMS, error) {
// 	var r []FacturasCMS
// 	err := json.Unmarshal(data, &r)
// 	return r, err
// }

// Función para deserializar JSON a un slice de FacturasCMS
func UnmarshalFacturasCMS(data []byte) ([]FacturasCMS, error) {
	var response Response
	err := json.Unmarshal(data, &response)
	return response.Data, err
}

// Método para serializar un slice de FacturasCMS a JSON
func (r FacturasCMS) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FacturasCMS struct {
	ID                int64             `json:"id"`
	InvoiceDate       time.Time         `json:"invoice_date"`
	CashboxID         int64             `json:"cashbox_id"`
	AlegraID          string            `json:"alegra_id"`
	AlegraTransaction AlegraTransaction `json:"alegra_transaction"`
	Country           string            `json:"country"`
	Account           string            `json:"account"`
	ApprovedBy        int64             `json:"approved_by"`
	Approver          string            `json:"approver"`
	State             string            `json:"state"`
	Support           string            `json:"support"`
	PayerEmail        string            `json:"payer_email"`
	Total             int64             `json:"total"`
	PaymentSupportID  int64             `json:"payment_support_id"`
	CouponGiftStatus  string            `json:"coupon_gift_status"`
	IsGift            bool              `json:"is_gift"`
	Items             []Item            `json:"items"`
}

type AlegraTransaction struct {
	Status          string          `json:"status"`
	Currency        Currency        `json:"currency"`
	Anotation       string          `json:"anotation"`
	Categories      []Category      `json:"categories"`
	AlegraData      AlegraData      `json:"alegra_data"`
	InvoiceRelation InvoiceRelation `json:"invoice_relation"`
}

type AlegraData struct {
	ID            int64  `json:"id"`
	BankAccount   string `json:"bank_account"`
	PaymentMethod string `json:"payment_method"`
}

type Category struct {
	ID           int64   `json:"id"`
	Tax          Tax     `json:"tax"`
	Price        float64 `json:"price"`
	Quantity     int64   `json:"quantity"`
	Observations string  `json:"observations"`
}

type Tax struct {
	ID int64 `json:"id"`
}

type Currency struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Description  string    `json:"description"`
	ExchangeRate float64   `json:"exchange_rate"`
}

type InvoiceRelation struct {
	Currency      Currency      `json:"currency"`
	InvoiceItems  []InvoiceItem `json:"invoice_items"`
	InvoiceHeader InvoiceHeader `json:"invoice_header"`
}

type InvoiceHeader struct {
	ID                  int64     `json:"id"`
	State               string    `json:"state"`
	Detail              Detail    `json:"detail"`
	Support             string    `json:"support"`
	UserID              int64     `json:"user_id"`
	GroupID             int64     `json:"group_id"`
	OrderID             int64     `json:"order_id"`
	Sequence            int64     `json:"sequence"`
	AlegraID            string    `json:"alegra_id"`
	CashboxID           int64     `json:"cashbox_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	ApprovedBy          int64     `json:"approved_by"`
	InvoiceDate         time.Time `json:"invoice_date"`
	Observations        string    `json:"observations"`
	ProratedDays        int64     `json:"prorated_days"`
	PaymentMethod       string    `json:"payment_method"`
	SubscriptionOrderID int64     `json:"subscription_order_id"`
}

type Detail struct {
	Totals                      Totals                     `json:"totals"`
	CreatedAtUser               time.Time                  `json:"created_at_user"`
	FirstSaleDate               time.Time                  `json:"first_sale_date"`
	GroupCreatedAt              time.Time                  `json:"group_created_at"`
	IsFirstPurchase             bool                       `json:"is_first_purchase"`
	DetailsBeforeFirstPurchase  DetailsBeforeFirstPurchase `json:"details_before_first_purchase"`
	DaysElapsedPreviousPurchase int64                      `json:"days_elapsed_previous_purchase"`
}

type DetailsBeforeFirstPurchase struct {
	DaysElapsed   int64 `json:"days_elapsed"`
	SubjectsViews int64 `json:"subjects_views"`
}

type Totals struct {
	Money   float64 `json:"money"`
	Tickets int64   `json:"tickets"`
}

type InvoiceItem struct {
	ID                  int64     `json:"id"`
	Quantity            int64     `json:"quantity"`
	CourseID            int64     `json:"course_id"`
	BasePrice           int64     `json:"base_price"`
	CreatedAt           time.Time `json:"created_at"`
	InvoiceID           int64     `json:"invoice_id"`
	UnitPrice           int64     `json:"unit_price"`
	UpdatedAt           time.Time `json:"updated_at"`
	CurrencyID          int64     `json:"currency_id"`
	ExchangeRate        int64     `json:"exchange_rate"`
	DiscountPrice       int64     `json:"discount_price"`
	OriginalPrice       int64     `json:"original_price"`
	IsSubscription      bool      `json:"is_subscription"`
	IsAttendingCourse   bool      `json:"is_attending_course"`
	ProfessorPercentage float64   `json:"professor_percentage"`
}

type Item struct {
	IsSubscription      bool    `json:"is_subscription"`
	CourseID            int64   `json:"course_id"`
	CourseName          string  `json:"course_name"`
	CurrencyID          int64   `json:"currency_id"`
	Price               int64   `json:"price"`
	ProfessorPercentage float64 `json:"professor_percentage"`
	CoursePicture       string  `json:"course_picture"`
	ExchangeRate        int64   `json:"exchange_rate"`
	Months              int64   `json:"months"`
	BeginsAt            string  `json:"begins_at"`
	EndsAt              string  `json:"ends_at"`
	IsRenewal           bool    `json:"is_renewal"`
}

// Estructura auxiliar que contiene solo el campo `data`
type Response struct {
	Data []FacturasCMS `json:"data"`
}

func main() {
	// Abre el archivo JSON
	file, err := os.Open("response.json")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Lee el contenido del archivo
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	// Imprimir el contenido del archivo para depuración
	//fmt.Println("Contenido del archivo JSON:", string(data))

	//=================================================================
	// Deserializa el contenido del archivo en un slice de FacturasCMS
	facturas, err := UnmarshalFacturasCMS(data)
	if err != nil {
		fmt.Println("Error al deserializar el JSON:", err)
		return
	}
	if len(facturas) > 0 {
		fmt.Println("Número de factura: ", facturas[0].ID)
		fmt.Println("Cliente: ", facturas[0].Items)
		fmt.Println("Total: ", facturas[0].AlegraTransaction.Currency.Name)
		fmt.Println("Fecha de facturación: ", facturas[0].InvoiceDate)
	}

	// Muestra los datos deserializados
	//fmt.Printf("Datos deserializados: %+v\n", facturas)

	// Si quieres volver a serializar y guardar el archivo
	serializedData, err := json.Marshal(facturas)
	if err != nil {
		fmt.Println("Error al serializar los datos:", err)
		return
	}

	err = ioutil.WriteFile("output.json", serializedData, 0777) //0644
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
		return
	}

	//fmt.Println("Datos serializados guardados en output.json")
}
