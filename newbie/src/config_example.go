package main

import (
	"log"
	"github.com/spf13/viper"
	"reflect"
	"fmt"
)


func main() {
	//var isWriter bool
	//flag.BoolVar(&isWriter,"isWriteable",false,"")
	//writeable:=flag.Bool("writable",false,"")
	////for _,v:= range flag.Args(){
	////	//log.Println(v)
	////}
	//log.Println(&writeable)
	//flag.Usage()
	//
	//config:=viper.New()
	//config.AddConfigPath()
	viper.SetDefault("CenterBank","cb")
	viper.SetDefault("CenterBankId","cb_master")
	viper.SetConfigType("yaml")
	viper.SetConfigName("core")
	viper.AddConfigPath("/Users/yangguo/work/blockchain/hyperledger/fabric/sampleconfig/")
	if err:=viper.ReadInConfig();err!=nil{
		log.Fatalf("Load Config failed;%v",err)
	}
	for key,value:=range viper.AllSettings() {
		log.Printf("%s= %v,typeof[%v]",key,value,reflect.TypeOf(value))
		switch value.(type) {
		case map[interface{}]interface{}:
			m:=value.(map[interface{}]interface{})
			for k,v:=range m{
				fmt.Printf("parentKey=%s,%s=%v,typeOf[%v]\n",key,k,v,reflect.TypeOf(v))
			}

		}

	}
	//config:=viper.New()
	//config.AutomaticEnv()
	//config.
	//for key,value:=range config.AllSettings(){
	//	log.Printf("%s= %v,typeof[%v]",key,value,reflect.TypeOf(value))
	//}




}
