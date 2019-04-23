package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // connect ไปที่ protocol rabbitMQ
	if err != nil {
		log.Fatalf("%s : %s", "Failed to connect to RabbitMQ", err)
	}
	defer connection.Close()

	channel, err := connection.Channel() // access channel เข้าไปใน connection
	if err != nil {
		log.Fatalf("%s : %s", "Failed to connect to RabbitMQ", err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare( // คือการสร้าง Queue ใหม่
		"hello", // name
		false,   // durable * เป็น queue ถาวรไหม (survive a broker)
		false,   // delete when unused * ไม่มีคนดึงข้อมูลใน queue จะลบทิ้งเลยหรือไม่
		false,   // exclusive * ถ้าคนที่ส่ง message หยุดแล้วจะลบ queue เลยหรือไม่
		false,   // no-wait * ไม่สนว่า queue มีอยู่หรือไม่ แต่จะใช้งานทันทีเลย
		nil,     // arguments * optional จะใช้โดย plugins ที่ใส่เพิ่มเข้าไป เช่น queue length limit (ความจุของคิว), message TTL (เวลาของคิว)
	)
	if err != nil {
		log.Fatalf("%s : %s", "Failed to declare a queue", err)
	}

	messags, err := channel.Consume( // worker ที่จะหยิบ queue ไปทำ
		queue.Name, // queue
		"",         // consumer * ชื่อของ worker/consumer ต่อไปที่ queue
		true,       // auto-ack * เป็นตัวที่บอกว่า consumer นี้ทำ queue นี้ไปแล้วและให้เอา queue ที่ทำแล้วออกไป(จะสำเร็จหรือไม่สำเร็จไม่รู้) ถ้าอยากรู้ให้ใช้ false
		false,      // exclusive
		false,      // no-local * ไม่ support ใน rabbitMQ
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatalf("%s : %s", "Failed to register a consumer", err)
	}

	forever := make(chan bool)

	go func() {
		for message := range messags {
			log.Printf("Received a message: %s", message.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
