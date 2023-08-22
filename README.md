# tbm-module
 Mock module for boardcasting and monitoring transaction


 ## Get Started

 Import tbmmodule to your package
 
```golang
 import (
	"github.com/patcharanant/tbm-module"
)
```

### Initiate()

Creates a new tbm instance with your http provider.

```golang

// change to your provider url
url := "https://demo.example.com"
tbm, err := tbmmodule.Initiate(url)
	if err != nil {
		panic(err)
	}

```

### TransactionPayload

create TrasactionPayload of the transaction you want to broadcast

```golang

payload := tbmmodule.TransactionPayload{
		Symbol:    "ETH",
		Price:     4500,
		Timestamp: uint64(time.Now().Unix()),
	}

```


### Broadcast()

Use TransactionPayload to Brodcast Trasaction
function will return TxHash struct

```golang

txHash, err := tbm.Broadcast(payload)
	if err != nil {
		panic(err)
	}
    fmt.Println("Trasaction hash : ", tx_hash.Hash)
```

### Monitor()

Monitor transaction status. provide txHash from Broadcast function or create with tbmmodule.TxHash
provide callback function when status change (ex. from pending to confirm will trigger a callback function)


```golang

tbm.Monitor(txHash, LogStatus) //txHash from Broadcast function

//OR

manualTxHash = tbmmodule.TxHash{Hash:"your hash here"}
tbm.Monitor(manualTxHash, LogStatus)

```
