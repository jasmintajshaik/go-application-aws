# Go Web Application - Deploy and Host on AWS

### Create an EC2 instance on AWS
* Go to AWS Console
* EC2 dashboard
* Launch Instance (with Ubuntu image)
### Log in to the EC2 instance
### Install Go
* sudo apt update
* Head to the official Go documentation here https://go.dev/doc/install and follow the steps to install Go.
### Clone this Repo
* git clone "https://github.com/jasmintajshaik/go-application-aws"
### Start the Web application
* git run main.go
#### Note: 
By default your application will not be accessible to the external world. Make sure you add an inbound traffic rule on your Ec2 instance. This allows the people to access the application on the port you have specified.
* Go to your EC2 instance
* Click on the instance
* Scroll below and select Security and then Security Groups
* Select Edit inbound rules and then add an inbound rule as below.
<img width="745" alt="image" src="https://github.com/jasmintajshaik/go-application-aws/assets/47131213/a7bc71c8-1dc4-4e77-9c9a-99592cb1e127">

#### Access your application
* http://ec2-instance-piblic-ip:8080/
