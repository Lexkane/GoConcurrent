package main

import "fmt"

func main(){

po:=new(PurchaseOrder)
po.Value=43.34
ch:=make(chan *PurchaseOrder)
go SavePurchase(po,ch)
newPo:=<-ch
fmt.Printf("Purchase Order is %v",newPo)
}

type PurchaseOrder struct{
	Number int
	Value float64
}

func SavePurchase(po* PurchaseOrder, callback chan *PurchaseOrder){
	po.Number=1234
	callback<-po

}