package main
import (
    "github.com/alouca/gosnmp"
    "log"
    "os"
    "strings"
    "time"
    "strconv"
	"github.com/influxdata/influxdb/client/v2"


)

func main() {
	argsWithoutProg := os.Args[1:]
	oids := os.Args[3:]
	addr := strings.Split(argsWithoutProg[0], ":")
	s, err := gosnmp.NewGoSNMP(addr[0], addr[1], gosnmp.Version2c, 5)
	if err != nil {
		log.Fatal(err)
	}
	dur,err := strconv.Atoi(argsWithoutProg[1])
	if err != nil{
		log.Println("Give me an valid time interval to probe")
		os.Exit(1)
	}

	for {

		for _, y := range oids {

			resp, err := s.Get(y)
			if err == nil {
				for _, v := range resp.Variables {
					switch v.Type {
					case gosnmp.Integer:

					default:


						i:= float64(v.Value.(uint64))

						err = Influx_Write(addr[0],v.Name,i)
						if err != nil{
							log.Println("Problem: %s",err)
							os.Exit(1)
						}
						log.Printf("Response from %s: %s : %f \n",addr[0], v.Name, i)

					}
				}
			}


		}
		time.Sleep(time.Duration(dur) * time.Second)
	}
}

func Influx_Write(tablename string,oid string,value float64) error  {
    // Make client
    c, err := client.NewHTTPClient(client.HTTPConfig{
        Addr: "http://localhost:8086",
    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }

    // Create a new point batch
    bp, err := client.NewBatchPoints(client.BatchPointsConfig{
        Database:  "anm",
        Precision: "s",
    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }

    // Create a point and add to batch
    fields := map[string]interface{}{
        "value": value,
    }
	tags := map[string]string{"oid": oid}
    pt, err := client.NewPoint(tablename,tags, fields, time.Now())

    if err != nil {
        log.Fatalln("Error: ", err)
    }

    bp.AddPoint(pt)

    // Write the batch
    c.Write(bp)
	return nil
}


