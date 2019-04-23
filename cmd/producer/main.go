package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // connect ไปที่ rabbitMQ
	if err != nil {
		log.Fatalf("%s : %s", "Failed to connect to RabbitMQ", err)
	}
	defer connection.Close()

	for {

	}

	channel, err := connection.Channel() // access channel เข้าไปใน connection
	if err != nil {
		log.Fatalf("%s : %s", "Failed to connect to RabbitMQ", err)
	}
	defer channel.Close()

	_, err = channel.QueueDeclare( // คือการสร้าง Queue ใหม่
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
}
