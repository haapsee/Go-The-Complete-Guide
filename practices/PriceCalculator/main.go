package main

import (
  "fmt"

  "example.com/price-calculator/prices"
  "example.com/price-calculator/filemanager"
//  "example.com/price-calculator/cmdmanager"
)

func main() {
  taxRates := []float64{0,0.07,0.1,0.15}
  doneChans := make([]chan bool, len(taxRates))
  errorChans := make([]chan error, len(taxRates))

  for index, taxRate := range taxRates {
    doneChans[index] = make(chan bool)
    errorChans[index] = make(chan error)

//    cmdm := cmdmanager.New()
//    priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)

    fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
    priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)


    go priceJob.Process(doneChans[index], errorChans[index])

  }


  for index, _ := range doneChans {
    select {
      case err  := <-errorChans[index]:
        if err != nil {
          fmt.Println(err)
        }
      case <-doneChans[index]:
        fmt.Println("Done!")
    }
  }
}
