# go-labs
A collection of hands-on labs to help learning the fundamentals of Golang Language. Each lab has:

## Index

[Lab-001 - Hello World](lab-001)

The goal of this lab is to launch a single Hello World to start the language.

[Lab-002 - A Single EC2 Instance in a Private Subnet + Bastion Host](lab-002)

The goal of this lab is to demonstrate how to access an EC2 instance launched in a private subnet using a bastion host.

[Lab-003 - A Single EC2 Instance in a Private Subnet + Bastion Host + NAT Gateway](lab-003)

This lab illustrates how an EC2 instance running in a *private* subnet can be configured to access the internet with the help of a service called [Nat gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html).

[Lab-004 - A Single EC2 Instance in a Private Subnet + Bastion Host + Nat Instance](lab-004)

This lab is similar to [lab-003](lab-003) but instead of using a *Nat gateway* you are asked to use a [Nat instance](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_NAT_Instance.html).

[Lab-005 - An HTTP Load Balancer](lab-005)

The goal of this lab is to illustrate how to use an [Application Load Balancer](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html) to distribute requests over two web servers running on distinct Availability Zones.

[Lab-006 - A Simple EC2 Auto Scaling](lab-006)

This goal of this lab is to illustrate how to setup a simple EC2 auto scaling service. You will create an auto scaling group with initially only one EC2 instance running. A rule will be setup to automatically launch a second EC2 instance when the CPU usage of the first instance reaches 75%.

[Lab-007 - A Simple EC2 Auto Scaling + Load Balancing](lab-007)

This lab is a continuation of [lab-006](lab-006) with a Load Balancer added.

[Lab-008 - Secure EC2 Auto Scaling + Load Balancing](lab-008)

This is lab is similar to [lab-007](lab-007) with a caveat: the auto scaling group is configured to create the EC2 instances in the private subnet. Also, two security groups control traffic in and out of the application load balancer and the EC2 instances (in the private subnets). To enable the EC2 instances to respond to requests coming from the internet, NAT gateways were configured on the public subnets of the VPC.

[Lab-009 - A Standalone RDS Instance Running MySQL](lab-009)

This lab illustrates how to launch a standalone RDS instance running a MySQL database. An EC2 instance will also be launched to allow access to the database using phpMyAdmin.

[Lab-010 - An RDS Instance with a Standby](lab-010)

This lab is similar to [lab-009](lab-009) with the addition of a standby RDS instance in another AZ to simulate a failover scenario.

[Lab-011 - An RDS Instance with a Read Replica](lab-011)

This lab is similar to [lab-009](lab-009) with the addition of a *read replica* RDS instance, a read-only copy of the database.

[Lab-012 - A Simple Dynamodb Service](lab-012)
This lab illustrates how to create and remotely access a dynamodb table.  Dynamodb is a NoSQL managed database service.   

[Lab-013 - An S3 Bucket Accessed via AWS CLI](lab-013)

The goal of this lab is to illustrate how to create an S3 bucket and access it from a client computer via AWS CLI (command-line interface).

[Lab-014 - An S3 Bucket Accessed via an S3 Gateway Endpoint](lab-014)

The goal of this lab is to illustrate how to access an S3 bucket from an EC2 instance in a private subnet using an S3 gateway endpoint.

[Lab-015 - An S3 Bucket Accessed via AWS CLI + Resource-based Policy](lab-015)

This lab is similar to [lab-013](../lab-013). However, we will be creating a resource-based policy to only allow access to the bucket if requests are sent from a specific IP address.  

[Lab-016 - Moving an EBS Volume Between EC2 Instances](lab-016)

The goal of this lab is to illustrate how to move an EBS volume from one EC2 instance to another.

[Lab-017 - Snapshot an EBS Volume](lab-017)

This lab's goal is similar to [lab-016](lab-016) but because we want to move a volume from one AZ to another we will have to first create a snapshot of the volume and then create a new volume from the snapshot.
