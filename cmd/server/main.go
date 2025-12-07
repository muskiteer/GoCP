// ...existing code...
package main

import (
    "fmt"
    "log"
    "reflect"

    "github.com/muskiteer/GoCP/tools"
)

func callFunc(fn interface{}, args ...interface{}) ([]interface{}, error) {
    v := reflect.ValueOf(fn)
    if v.Kind() != reflect.Func {
        return nil, fmt.Errorf("not a function")
    }
    if len(args) != v.Type().NumIn() {
        return nil, fmt.Errorf("argument count mismatch: want %d got %d", v.Type().NumIn(), len(args))
    }

    in := make([]reflect.Value, len(args))
    for i, a := range args {
        if a == nil {
            in[i] = reflect.Zero(v.Type().In(i))
            continue
        }
        av := reflect.ValueOf(a)
        target := v.Type().In(i)
        if !av.Type().AssignableTo(target) {
            if av.Type().ConvertibleTo(target) {
                av = av.Convert(target)
            } else {
                return nil, fmt.Errorf("arg %d: cannot convert %s to %s", i, av.Type(), target)
            }
        }
        in[i] = av
    }

    out := v.Call(in)
    res := make([]interface{}, len(out))
    for i, rv := range out {
        if rv.IsValid() && rv.CanInterface() {
            res[i] = rv.Interface()
        } else {
            res[i] = nil
        }
    }
    return res, nil
}

func main() {
    var toolMap = make(map[string]interface{})
    toolMap["FetchCryptoData"] = tools.FetchCryptoData

    // dynamic call
    res, err := callFunc(toolMap["FetchCryptoData"], "bitcoin", "usd")
    if err != nil {
        log.Fatalf("call error: %v", err)
    }

    // expected signature: (float64, error)
    price, _ := res[0].(float64)
    var callErr error
    if len(res) > 1 && res[1] != nil {
        callErr = res[1].(error)
    }
    if callErr != nil {
        log.Fatalf("Error fetching crypto data: %v", callErr)
    }
    log.Printf("The current price of Bitcoin in USD is: %f", price)
}
