# super_suds
Willy and Jimmy's super suds site


Bavid Taylor fictional brother of the P&G CEO wants us to create a throwback website. A tribute to the Ivory Brand. Willy and Jimmy's Super Suds. (This is funny because the founders are William Proctor and James Gamble. Get it?)

We're going to build a single "micro-service" that will list out Willy and Jimmy's products. They are Soap Bar, Suds Bucket, and warsh cloth. (we can expand this with other relevant information). This service will be deployed to AWS Lambda with the golang runtime. The infrastructure will be managed with terraform.

The front end needs to be built in vue.js and consume this service.

## Steps to run
go run main.go