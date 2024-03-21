package helper

import (
	"log"
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type MdtClient struct {
	c coreapi.Client
}

func NewMdtDriver() MdtClient {
	return MdtClient{c: coreapi.Client{}}
}

func (c *MdtClient) CreateTransaction(orderID string, guideID string, ticketPrice int, ticketIDs []string, guidePrice int, taxGuide float64, taxTicket float64, taxAmount float64, guideName string, placeName string, custEmail string,
	custName string, custHP string, bank string) (*coreapi.ChargeResponse, error) {
	c.c.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	var items []midtrans.ItemDetails
	var totalAmount int64

	// Create items and calculate total amount for tickets
	for _, ticketID := range ticketIDs {
		items = append(items, midtrans.ItemDetails{
			ID:    ticketID,
			Name:  "Ticket" + " " + placeName,
			Price: int64(ticketPrice), // Assuming each ticket has the same price
			Qty:   1,
		})
		totalAmount += int64(ticketPrice)
	}

	if guideID != "" {
		items = append(items, midtrans.ItemDetails{
			ID:    guideID,
			Name:  "Tour Guide " + placeName + " : " + guideName,
			Price: int64(guidePrice),
			Qty:   1,
		})
		totalAmount += int64(guidePrice)

		items = append(items, midtrans.ItemDetails{
			ID:    "TAX",
			Name:  "Guide Tax for : " + placeName,
			Price: int64(taxGuide),
			Qty:   1,
		})
	}

	items = append(items, midtrans.ItemDetails{
		ID:    "TAX",
		Name:  "Ticket Tax " + placeName,
		Price: int64(taxTicket),
		Qty:   int32(len(ticketIDs)),
	})
	totalTax := int64(taxTicket * float64(len(ticketIDs)) + taxGuide)
    totalAmount += totalTax

	log.Println(taxTicket)
	// Create charge request
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: totalAmount,
		},
		Items: &items,
		CustomerDetails: &midtrans.CustomerDetails{
			FName: custName,
			Phone: custHP,
			Email: custEmail,
		},
	}

	// Set bank transfer details based on the selected bank
	switch bank {
	case "bca":
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{Bank: midtrans.BankBca}
	case "bri":
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{Bank: midtrans.BankBri}
	case "bni":
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{Bank: midtrans.BankBni}
	}

	// Charge the transaction
	coreApiRes, err := c.c.ChargeTransaction(chargeReq)
	if err != nil {
		return nil, err
	}

	return coreApiRes, nil
}

func (c *MdtClient) NotifHandler(id string) (*coreapi.TransactionStatusResponse, error) {
	c.c.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	response, err := c.c.CheckTransaction(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}
