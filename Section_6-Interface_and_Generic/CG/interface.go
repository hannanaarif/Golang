package main

import "fmt"

type paymenter interface{
	pay(amount float64)
}

type payment struct{

}

func (p payment) makePayment(amount float64){
    // razorpayGw:=razorpay{}
	// razorpayGw.pay(amount)
	stripeGw:=stripe{}
	stripeGw.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float64){

	fmt.Println("making the payment in process",amount)
}

type stripe struct{}

func (s stripe) pay(amount float64){
	fmt.Println("making payment using stripe",amount)
}
func main(){
	newPayment:=payment{}
 }